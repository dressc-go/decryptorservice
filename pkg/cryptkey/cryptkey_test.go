package cryptkey

import (
	"path"
	"runtime"
	"strings"
	"testing"

	"github.com/dressc-go/zlogger"
)

var testdataPathCryptkey string

func init() {
	_, filename, _, _ := runtime.Caller(0)
	testdataPathCryptkey = path.Join(path.Dir(filename), "testdata")
	zlogger.SetGlobalLevel(zlogger.DebugLevel)
}

func TestCryptKey_New_DoesNotExist(t *testing.T) {
	pf := new(CryptKey)
	err := pf.New("does_not_exist", []byte{0})
	if err == nil {
		t.Errorf("Failed: FileNotFound expected")
	}
}

func TestCryptFile_New_Privkey(t *testing.T) {
	testFilePath := path.Join(testdataPathCryptkey, "privkey.pem")
	pf := new(CryptKey)
	no_passw := []byte{0}
	password := []byte("secret")
	err := pf.NewPrivate(testFilePath, no_passw)
	if err != nil {
		t.Errorf("Failed: with" + err.Error())
	}
	testFilePath = path.Join(testdataPathCryptkey, "privkey.encrypted.pem")
	if err := pf.NewPrivate(testFilePath, password); err != nil {
		t.Errorf("Failed: with" + err.Error())
	}
	testFilePath = path.Join(testdataPathCryptkey, "privkey.pkcs1.pem")
	if err := pf.NewPrivate(testFilePath, no_passw); err != nil {
		t.Errorf("Failed: with" + err.Error())
	}
	testFilePath = path.Join(testdataPathCryptkey, "privkey.encrypted.pkcs1.pem")
	if err := pf.NewPrivate(testFilePath, password); err != nil {
		t.Errorf("Failed: with" + err.Error())
	}
	if _, err := pf.GetPrivateKey(); err != nil {
		t.Errorf("Failed: Key returns no private key")
	}
}

func TestCryptFile_New_Empty(t *testing.T) {
	testFilePath := path.Join(testdataPathCryptkey, "empty.pem")
	expectErrPrefix := "reading PEM failed"
	gotError := ""
	pf := new(CryptKey)
	err := pf.New(testFilePath, []byte{0})
	if err != nil {
		gotError = err.Error()
		if strings.HasPrefix(gotError, expectErrPrefix) {
			return
		}
	}
	t.Errorf("Failed. Expected Error: " + expectErrPrefix + " got: " + gotError)
}

func TestCryptFile_New_Damaged(t *testing.T) {
	testFilePath := path.Join(testdataPathCryptkey, "pubkey.damaged.pem")
	expectErrContains := "Can't parse PKIX from PEM"
	gotError := ""
	pf := new(CryptKey)
	err := pf.NewPublic(testFilePath)
	if err != nil {
		gotError = err.Error()
		if strings.Contains(gotError, expectErrContains) {
			return
		}
	}
	t.Errorf("Failed. Expected Error: " + expectErrContains + " got: " + gotError)
}

func TestCryptFile_New_Invalid(t *testing.T) {
	testFilePath := path.Join(testdataPathCryptkey, "privkey.invalid.pem")
	expectErrContains := "Can't parse PKCS8 from PEM"
	gotError := ""
	pf := new(CryptKey)
	err := pf.NewPrivate(testFilePath, []byte{0})
	if err != nil {
		gotError = err.Error()
		if strings.Contains(gotError, expectErrContains) {
			return
		}
	}
	t.Errorf("Failed. Expected Error: " + expectErrContains + " got: " + gotError)
}

func TestCryptFile_New_Invalid_PKCS1(t *testing.T) {
	testFilePath := path.Join(testdataPathCryptkey, "privkey.invalid.pkcs1.pem")
	expectErrContains := "Can't parse PKCS1 from PEM"
	gotError := ""
	pf := new(CryptKey)
	err := pf.NewPrivate(testFilePath, []byte{0})
	if err != nil {
		gotError = err.Error()
		if strings.Contains(gotError, expectErrContains) {
			return
		}
	}
	t.Errorf("Failed. Expected Error: " + expectErrContains + " got: " + gotError)
}

func TestCryptFile_New_Pubkey(t *testing.T) {
	testFilePath := path.Join(testdataPathCryptkey, "pubkey.pem")
	pf := new(CryptKey)
	err := pf.NewPublic(testFilePath)
	if err != nil {
		t.Errorf("Failed: with" + err.Error())
	}
	if pf.pubkey == nil {
		t.Errorf("Got no Pubkey")
	}
	testFilePath = path.Join(testdataPathCryptkey, "pubkey.pkcs1.pem")
	err = pf.NewPublic(testFilePath)
	if err != nil {
		t.Errorf("Failed: with" + err.Error())
	}
	if pf.pubkey == nil {
		t.Errorf("Got no Pubkey")
	}
	if _, err := pf.GetPrivateKey(); err == nil {
		t.Errorf("Failed: Public key returns private key")
	}
}

func TestCryptFile_New_Invalid_Pubkey(t *testing.T) {
	testFilePath := path.Join(testdataPathCryptkey, "pubkey.invalid.pem")
	expectErrContains := "Can't parse PKIX from PEM"
	gotError := ""
	pf := new(CryptKey)
	err := pf.NewPublic(testFilePath)
	if err != nil {
		gotError = err.Error()
		if strings.Contains(gotError, expectErrContains) {
			return
		}
	}
	t.Errorf("Failed. Expected Error: " + expectErrContains + " got: " + gotError)
}

func TestCryptFile_New_Invalid_Pubkey_PKCS1(t *testing.T) {
	testFilePath := path.Join(testdataPathCryptkey, "pubkey.invalid.pkcs1.pem")
	expectErrContains := "Can't parse PKCS1 from PEM"
	gotError := ""
	pf := new(CryptKey)
	err := pf.NewPublic(testFilePath)
	if err != nil {
		gotError = err.Error()
		if strings.Contains(gotError, expectErrContains) {
			return
		}
	}
	t.Errorf("Failed. Expected Error: " + expectErrContains + " got: " + gotError)
}
