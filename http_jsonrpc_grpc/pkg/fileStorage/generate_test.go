package fileStorage

import (
	"testing"
)

func TestGenerateHashFileNameThatNotExist(t *testing.T) {
	initConfig(t)

	t.Run("valid", func(t *testing.T) {
		var (
			fileNamePrefix    = "request"
			generatedFilePath = "api_test/123123213/12/2021/May/28"
			pathPrefix        = gConf.PathPrefix // Требуется initConfig(t)
		)

		if fileName, err := generateFileNameThatNotExist(pathPrefix, generatedFilePath, fileNamePrefix); err != nil {
			t.Errorf("%sError: %s%s", RED_BG, err, NO_COLOR)
		} else {
			t.Logf("New file name %s", fileName)
		}
	})

	if t.Failed() == false {
		t.Logf("%sSuccess%s", GREEN, NO_COLOR)
	}
}
