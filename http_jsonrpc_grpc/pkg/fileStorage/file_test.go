package fileStorage

import (
	"testing"
)

func TestFile(t *testing.T) {
	initConfig(t)

	var (
		fileBody         = []byte("КАЖДЫЙ охотник желает знать")
		pathPrefix       = gConf.PathPrefix // Требуется initConfig(t)
		filePathWithName string
	)

	t.Run("FileSaver", func(t *testing.T) {
		fileResult, err := New(pathPrefix).SaveRequest(100500, fileBody)
		if err != nil {
			t.Errorf("%sError: %s%s", RED_BG, err, NO_COLOR)
			t.FailNow()
		}
		filePathWithName = fileResult.GetFilePath()

		t.Logf("filePathWithName %s", filePathWithName)
	})

	t.Run("FileGetter", func(t *testing.T) {
		fileResult, err := New(pathPrefix).LoadFromStorage(filePathWithName)
		if err != nil {
			t.Errorf("%sError: %s%s", RED_BG, err, NO_COLOR)
			t.FailNow()
		}

		body := fileResult.GetFileBody()

		if string(body) != string(fileBody) {
			t.Errorf("%sFail: body expected %s got %s%s", RED_BG, string(fileBody), string(body), NO_COLOR)
		}
	})

	if t.Failed() == false {
		t.Logf("%sSuccess%s", GREEN, NO_COLOR)
	}
}
