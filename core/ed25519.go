package core

import (
	"crypto"
	"crypto/ed25519"
	"encoding/base64"
	"io"
)

type ED25519Signer struct {
	PublicKey []byte
}

func (s *ED25519Signer) Public() crypto.PublicKey {
	return (ed25519.PublicKey)(s.PublicKey)
}

func (s *ED25519Signer) Sign(rand io.Reader, digest []byte, opts crypto.SignerOpts) (signature []byte, err error) {
	panic("implement me")
}

func ParseED25519PublicKeyFromStr(str string) (ed25519.PublicKey, error) {
	content, err := base64.StdEncoding.DecodeString(str)
	if err != nil {
		content, err := base64.URLEncoding.DecodeString(str)
		if err != nil {
			return nil, err
		}
		return content, nil
	}

	return content, nil
}
