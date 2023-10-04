package cryptkey

import (
	"crypto/rsa"
	"crypto/x509"
	"github.com/dressc-go/zlogger"
	"github.com/pkg/errors"
)

type CryptKey struct {
	pemfile *PemFile
	privkey *rsa.PrivateKey
	pubkey  *rsa.PublicKey
}

func (ck *CryptKey) New(filePath string) error {
	logger := zlogger.GetLogger("cryptkey.New")
	ck.pemfile = new(PemFile)
	e := ck.pemfile.New(filePath)
	if e != nil {
		err := errors.Wrap(e, "reading PEM failed")
		logger.Error().Err(err).Msg("")
		return err
	}
	if ck.pemfile.decoded.Type == "PRIVATE KEY" {
		e = ck.parsePKCS8PrivateKey()
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

func (ck *CryptKey) parsePKCS8PrivateKey() error {
	buf, e := x509.ParsePKCS8PrivateKey(ck.pemfile.decoded.Bytes)
	if e != nil {
		return errors.Wrap(e, "Can't parse PKCS8 from PEM")
	}
	ck.privkey = buf.(*rsa.PrivateKey)
	return nil
}

func (ck *CryptKey) parsePKCS1PrivateKey() error {
	buf, e := x509.ParsePKCS1PrivateKey(ck.pemfile.decoded.Bytes)
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

func (ck *CryptKey) GetPublicKey() (*rsa.PublicKey, error) {
	if ck.pubkey != nil {
		return ck.pubkey, nil
	}
	return nil, errors.New("Not a public key")
}
