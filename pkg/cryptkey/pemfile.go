package cryptkey

import (
	"encoding/pem"
	"github.com/pkg/errors"
	"os"
	"path/filepath"
)

type PemFile struct {
	filePath    string
	fileContent []byte
	decoded     *pem.Block
}

func (pf *PemFile) New(filePath string) error {
	pf.filePath = filepath.Clean(filePath)
	e := pf.readFromFile()
	if e != nil {
		return e
	}
	e = pf.decodePem()
	if e != nil {
		return e
	}
	return nil
}

func (pf *PemFile) readFromFile() error {
	buf, err := os.ReadFile(pf.filePath)
	if err != nil {
		e := errors.Wrap(err, "Could not read file")
		return e
	}
	pf.fileContent = buf
	return nil
}

func (pf *PemFile) decodePem() error {
	buf, _ := pem.Decode(pf.fileContent)
	if buf == nil {
		e := errors.New("PEM Decoding failed. File not in PEM format?")
		return e
	}
	pf.decoded = buf
	return nil
}
