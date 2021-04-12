package core

import (
	"crypto"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"io"
)

type RSASigner struct {
	PublicKey *rsa.PublicKey
}

func (s *RSASigner) Public() crypto.PublicKey {
	return s.PublicKey
}

func (s *RSASigner) Sign(rand io.Reader, digest []byte, opts crypto.SignerOpts) (signature []byte, err error) {
	panic("implement me")
}

func ParseRSAPubKeyFromPEMStr(pemStr string) (*rsa.PublicKey, error) {
	block, _ := pem.Decode([]byte(pemStr))
	if block == nil {
		return nil, errors.New("failed to parse PEM block containing the key")
	}

	return x509.ParsePKCS1PublicKey(block.Bytes)
}
