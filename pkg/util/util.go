package util

import (
	"github.com/levigross/grequests"
	"github.com/pkg/errors"
	"io/ioutil"
)

func DownloadFile(uri, filename string) error {
	resp, err := grequests.Get(uri, nil)
	if err != nil {
		return errors.Wrapf(err, "unable to download uri %s\n", uri)
	}

	return resp.DownloadToFile(filename)
}

func WriteFile(data []byte, filename string) error {
	return ioutil.WriteFile(filename, data, 0644)
}
