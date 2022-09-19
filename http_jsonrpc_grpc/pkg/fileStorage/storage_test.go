package fileStorage

import (
	"path/filepath"
	"testing"
)

func TestPutGetFile(t *testing.T) {
	initConfig(t)

	var (
		fileBody          = []byte("КАЖДЫЙ охотник желает знать")
		fileName          = "9b7833f218712b471c7197e250c8eefb70eac244.txt"
		generatedFilePath = "api_test/123123213/12/2021/May/28"
		pathPrefix        = gConf.PathPrefix // Требуется initConfig(t)
	)

	t.Run("putFile", func(t *testing.T) {
		if err := putFile(pathPrefix, generatedFilePath, fileName, fileBody); err != nil {
			t.Errorf("%sError: %s%s", RED_BG, err, NO_COLOR)
		}
	})

	if t.Failed() == true {
		t.FailNow()
	}

	t.Run("getFile", func(t *testing.T) {
		body, err := getFile(filepath.Join(pathPrefix, generatedFilePath, fileName))
		if err != nil {
			t.Errorf("%sError: %s%s", RED_BG, err, NO_COLOR)
			t.FailNow()
		}
		if string(body) != string(fileBody) {
			t.Errorf("%sFail: expected %s got %s%s", RED_BG, string(fileBody), string(body), NO_COLOR)
		}
	})

	if t.Failed() == false {
		t.Logf("%sSuccess%s", GREEN, NO_COLOR)
	}
}
