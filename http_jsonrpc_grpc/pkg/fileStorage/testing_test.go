package fileStorage

import (
	"github.com/GlobchanskyDenis/benchmarks.git/pkg/dto"
	"10.10.11.220/ursgis/u_conf.git"
	"testing"
)

const (
	RED       = "\033[31m"
	RED_BG    = "\033[41;30m"
	GREEN     = "\033[32m"
	GREEN_BG  = "\033[42;30m"
	YELLOW    = "\033[33m"
	YELLOW_BG = "\033[43;30m"
	NO_COLOR  = "\033[m"
)

var gConf *dto.Config

func initConfig(t *testing.T) {
	if err := u_conf.SetConfigFile("../../test.json"); err != nil {
		t.Errorf(RED_BG + "Cannot initialize package: " + err.Error() + NO_COLOR)
		t.FailNow()
	}
	gConf = &dto.Config{}
	if err := u_conf.ParsePackageConfig(gConf, "Storage"); err != nil {
		t.Errorf(RED_BG + "Cannot initialize package: " + err.Error() + NO_COLOR)
		t.FailNow()
	}
	t.Log(GREEN + "package configured successfully" + NO_COLOR)
}
