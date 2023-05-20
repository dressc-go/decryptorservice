package cryptkey

import (
	"crypto/rsa"
	"crypto/x509"

	"github.com/dressc-go/zlogger"
	"github.com/pkg/errors"
	"go.step.sm/crypto/pemutil"
)

type CryptKey struct {
	pemfile *PemFile
	privkey *rsa.PrivateKey
	pubkey  *rsa.PublicKey
}

func (ck *CryptKey) NewPublic(filePath string) error {
	return ck.New(filePath, []byte{0})
}

func (ck *CryptKey) NewPrivate(filePath string, password []byte) error {
	return ck.New(filePath, password)
}

func (ck *CryptKey) New(filePath string, password []byte) error {
	logger := zlogger.GetLogger("cryptkey.New")
	ck.pemfile = new(PemFile)
	e := ck.pemfile.New(filePath)
	if e != nil {
		err := errors.Wrap(e, "reading PEM failed")
		logger.Error().Err(err).Msg("")
		return err
	}
	if ck.pemfile.decoded.Type == "PRIVATE KEY" {
		password := []byte{0}
		e = ck.parsePKCS8PrivateKey(password)
	} else if ck.pemfile.decoded.Type == "ENCRYPTED PRIVATE KEY" {
		e = ck.parsePKCS8PrivateKey(password)
	} else if ck.pemfile.decoded.Type == "RSA PRIVATE KEY" {
		e = ck.parsePKCS1PrivateKey()
	} else if ck.pemfile.decoded.Type == "PUBLIC KEY" {
		e = ck.parsePKIXPublicKey()
	} else if ck.pemfile.decoded.Type == "RSA PUBLIC KEY" {
		e = ck.parsePKCS1PublicKey()
	} else {
		e = errors.New("PEM Type not implemented")
	}
	if e != nil {
		err := errors.Wrap(e, "parsing PEM failed")
		logger.Error().Err(err).Msg("")
		return err
	}
	return nil
}

func (ck *CryptKey) parsePKIXPublicKey() error {
	buf, e := x509.ParsePKIXPublicKey(ck.pemfile.decoded.Bytes)
	if e != nil {
		return errors.Wrap(e, "Can't parse PKIX from PEM")
	}
	ck.pubkey = buf.(*rsa.PublicKey)
	return nil
}

func (ck *CryptKey) parsePKCS1PublicKey() error {
	buf, e := x509.ParsePKCS1PublicKey(ck.pemfile.decoded.Bytes)
	if e != nil {
		return errors.Wrap(e, "Can't parse PKCS1 from PEM")
	}
	ck.pubkey = buf
	return nil
}

func (ck *CryptKey) parsePKCS8PrivateKey(password []byte) error {
	var e error

	inbuf := ck.pemfile.decoded.Bytes
	if password[0] != 0 {
		inbuf, e = pemutil.DecryptPEMBlock(ck.pemfile.decoded, password)
	}
	buf, e := x509.ParsePKCS8PrivateKey(inbuf)
	if e != nil {
		return errors.Wrap(e, "Can't parse PKCS8 from PEM")
	}
	ck.privkey = buf.(*rsa.PrivateKey)
	return nil
}

func (ck *CryptKey) parsePKCS1PrivateKey() error {
	var e error
	var passwd []byte
	passwd = []byte("secret")
	inbuf := ck.pemfile.decoded.Bytes
	if x509.IsEncryptedPEMBlock(ck.pemfile.decoded) {
		inbuf, e = x509.DecryptPEMBlock(ck.pemfile.decoded, passwd)
	}
	buf, e := x509.ParsePKCS1PrivateKey(inbuf)
	if e != nil {
		return errors.Wrap(e, "Can't parse PKCS1 from PEM")
	}
	ck.privkey = buf
	return nil
}

func (ck *CryptKey) GetPrivateKey() (*rsa.PrivateKey, error) {
	if ck.privkey != nil {
		return ck.privkey, nil
	}
	return nil, errors.New("Not a private key")
}
