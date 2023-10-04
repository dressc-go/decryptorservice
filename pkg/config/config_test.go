package config

import (
	"github.com/dressc-go/zlogger"
	"os"
	"path"
	"runtime"
	"strings"
	"testing"
)

var testdataPath string

func init() {
	_, filename, _, _ := runtime.Caller(0)
	testdataPath = path.Join(path.Dir(filename), "testdata")
	zlogger.SetGlobalLevel(zlogger.DebugLevel)
}

func TestConfig_New_Fail(t *testing.T) {
	expectErrPrefix := "Config file does not exist"
	gotError := ""
	cnf := new(Config)
	err := cnf.New()
	if err != nil {
		gotError = err.Error()
		if strings.HasPrefix(gotError, expectErrPrefix) {
			return
		}
	}
	t.Errorf("Failed. Expected Error: " + expectErrPrefix + " got: " + gotError)
}

func TestConfig_New_Explicit(t *testing.T) {
	teardown := func() {
		_ = os.Setenv("ConfigFile", "")
	}
	defer teardown()
	_ = os.Setenv("ConfigFile", path.Join(testdataPath, "config.yml"))
	cnf := new(Config)
	err := cnf.New()
	if err != nil {
		t.Errorf("Failed: " + err.Error())
	}
	if cnf.IpAddress != "127.0.1.15" {
		t.Errorf("Failed: got wrong data")
	}
	if cnf.DecPrivateKeyFile != "privkey.encrypted.pem" {
		t.Errorf("Failed: got wrong data")
	}
}

func TestConfig_New_Explicit_Fail(t *testing.T) {
	teardown := func() {
		_ = os.Setenv("ConfigFile", "")
	}
	defer teardown()
	expectErrPrefix := "Config file does not exist"
	gotError := ""
	_ = os.Setenv("ConfigFile", path.Join(testdataPath, "does_not_exist.yml"))
	cnf := new(Config)
	err := cnf.New()
	if err != nil {
		gotError = err.Error()
		if strings.HasPrefix(gotError, expectErrPrefix) {
			return
		}
	}
	t.Errorf("Failed. Expected Error: " + expectErrPrefix + " got: " + gotError)
}

func TestConfig_New_Automatic(t *testing.T) {
	cwd, _ := os.Getwd()
	teardown := func() {
		_ = os.Chdir(cwd)
	}
	defer teardown()
	_ = os.Chdir(testdataPath)
	cnf := new(Config)
	err := cnf.New()
	if err != nil {
		t.Errorf("Failed: " + err.Error())
	}
}
