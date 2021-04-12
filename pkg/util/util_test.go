package util

import "testing"

func Test_downloadFile(t *testing.T) {
	err := DownloadFile("https://c-ssl.duitang.com/uploads/item/201207/24/20120724171731_Cecsx.png", "/tmp/cat.png")
	if err != nil {
		t.Error(err)
	}
}
