package notion

import (
	"bytes"
	"fmt"
	"github.com/kjk/notionapi"
	"github.com/kjk/notionapi/tomarkdown"
	"github.com/sirupsen/logrus"
	"github.com/xujiahua/notion-md/pkg/util"
	"path/filepath"
	"strings"
	"time"
)

type Manager struct {
	*notionapi.Client
	outputDir  string
	rootPageID string
}

func New(token, rootPageID, outputDir string) *Manager {
	logrus.Debugf("token: %s\n", token)
	logrus.Debugf("rootPageID: %s\n", rootPageID)
	logrus.Debugf("outputDir: %s\n", outputDir)

	client := &notionapi.Client{}
	client.AuthToken = token

	return &Manager{
		Client:     client,
		outputDir:  outputDir,
		rootPageID: rootPageID,
	}
}

func (m Manager) Do() {
	rootPage, err := m.DownloadPage(m.rootPageID)
	if err != nil {
		logrus.Errorf("unable to download rootPage %s, %v", m.rootPageID, err)
		return
	}

	for _, pageID := range rootPage.GetSubPages() {
		logrus.Infof("downloading page %s", pageID)
		page, err := m.DownloadPage(pageID)
		if err != nil {
			logrus.Errorf("unable to download page %s, %v", pageID, err)
			continue
		}
		err = m.toMarkdown(page)
		if err != nil {
			logrus.Errorf("unable to save markdown %s, %v", pageID, err)
			continue
		}
		err = m.downloadImages(page)
		if err != nil {
			logrus.Errorf("unable to download images %s, %v", pageID, err)
			continue
		}
	}
}

func (m Manager) toMarkdown(page *notionapi.Page) error {
	filename := tomarkdown.MarkdownFileNameForPage(page)
	filename = filepath.Join(m.outputDir, filename)
	metadata := extractMetadata(page)
	data := tomarkdown.ToMarkdown(page)
	// NOTE: trim title line at the beginning
	for i, c := range data {
		// meet first line
		if c == '\n' {
			data = data[i+1:]
			break
		}
	}
	data = append(metadata, data...)
	return util.WriteFile(data, filename)
}

func extractMetadata(page *notionapi.Page) []byte {
	var buf bytes.Buffer
	buf.WriteString("---\n")
	title, date := extractTitleAndDate(page)
	buf.WriteString(fmt.Sprintf("title: \"%s\"\n", title))
	buf.WriteString(fmt.Sprintf("date: \"%s\"\n", date))
	buf.WriteString("draft: true\n")
	buf.WriteString("toc: true\n")
	buf.WriteString("autoCollapseToc: false\n")
	buf.WriteString("comment: true\n")
	buf.WriteString("---\n")
	return buf.Bytes()
}

func extractTitleAndDate(page *notionapi.Page) (string, string) {
	var title string
	var date string
	page.ForEachBlock(func(block *notionapi.Block) {
		if block.Type == notionapi.BlockPage {
			title = block.Title
			date = block.CreatedOn().Format(time.RFC3339)
		}
	})
	return title, date
}

func (m Manager) downloadImages(page *notionapi.Page) error {
	var errstrings []string

	page.ForEachBlock(func(block *notionapi.Block) {
		if block.IsImage() {
			filename := getImageFilename(block)
			filename = filepath.Join(m.outputDir, filename)
			downloadFileResponse, err := m.DownloadFile(block.Source, block.ID)
			if err != nil {
				errstrings = append(errstrings, err.Error())
				return
			}

			err = util.WriteFile(downloadFileResponse.Data, filename)
			if err != nil {
				errstrings = append(errstrings, err.Error())
				return
			}
		}
	})

	if len(errstrings) == 0 {
		return nil
	}

	return fmt.Errorf(strings.Join(errstrings, "\n"))
}

// reference:
// func (c *Converter) RenderImage(block *notionapi.Block)
func getImageFilename(block *notionapi.Block) string {
	if len(block.FileIDs) == 0 {
		logrus.Warnf("RenderImage when len(FileIDs) == 0 NYI\n")
		return ""
	}
	source := block.Source // also present in block.Format.DisplaySource
	// source looks like: "https://s3-us-west-2.amazonaws.com/secure.notion-static.com/e5470cfd-08f0-4fb8-8ec2-452ca1a3f05e/Schermafbeelding2018-06-19om09.52.45.png"
	fileID := block.FileIDs[0]
	parts := strings.Split(source, "/")
	fileName := parts[len(parts)-1]
	parts = strings.SplitN(fileName, ".", 2)
	ext := ""
	if len(parts) == 2 {
		fileName = parts[0]
		ext = "." + parts[1]
	}
	return fmt.Sprintf("%s-%s%s", fileName, fileID, ext)
}
