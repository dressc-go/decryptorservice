package cryptkey

import (
	"github.com/dressc-go/zlogger"
	"path"
	"runtime"
	"strings"
	"testing"
)

var testdataPath_pemfile string

func init() {
	_, filename, _, _ := runtime.Caller(0)
	testdataPath_pemfile = path.Join(path.Dir(filename), "testdata")
	zlogger.SetGlobalLevel(zlogger.WarnLevel)
}

func TestPemFile_New_DoesNotExist(t *testing.T) {
	pf := new(PemFile)
	err := pf.New("does_not_exist")
	if err == nil {
		t.Errorf("Failed: FileNotFound expected")
	}
}

func TestPemFile_New_Privkey(t *testing.T) {
	testFilePath := path.Join(testdataPath_pemfile, "privkey.pem")
	pf := new(PemFile)
	err := pf.New(testFilePath)
	if err != nil {
		t.Errorf("Failed: FileNotFound")
	}
	if pf.decoded.Type != "PRIVATE KEY" {
		t.Errorf("Failed: PEM Decoding wrong")
	}
}

func TestPemFile_New_NoPem(t *testing.T) {
	testFilePath := path.Join(testdataPath_pemfile, "nopem.pem")
	expectErrPrefix := "reading PEM failed: PEM Decoding failed."
	gotError := ""
	pf := new(CryptKey)
	err := pf.New(testFilePath)
	if err != nil {
		gotError = err.Error()
		if strings.HasPrefix(gotError, expectErrPrefix) {
			return
		}
	}
	t.Errorf("Failed. Expected Error: " + expectErrPrefix + " got: " + gotError)
}
