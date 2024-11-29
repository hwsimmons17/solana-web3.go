package base58

import (
	"crypto"
	"crypto/ed25519"
	"crypto/rand"
	"errors"
	"solana"

	"filippo.io/edwards25519"
	"github.com/mr-tron/base58"
)

type Signer struct {
	privateKey []byte
}

func NewSigner(privateKey []byte) (solana.Signer, error) {
	if len(privateKey) != 64 {
		return nil, errors.New("invalid private key length, expected 64 bytes")
	}

	// check if the public key is on the ed25519 curve
	pub := ed25519.PrivateKey(privateKey).Public().(ed25519.PublicKey)
	if !IsOnCurve(pub) {
		return nil, errors.New("corresponding public key is NOT on the ed25519 curve")
	}

	return &Signer{privateKey}, nil
}

func IsOnCurve(b []byte) bool {
	if len(b) != ed25519.PublicKeySize {
		return false
	}
	_, err := new(edwards25519.Point).SetBytes(b)
	isOnCurve := err == nil
	return isOnCurve
}

func NewSignerFromBase58(str string) (solana.Signer, error) {
	bytes, err := base58.Decode(str)
	if err != nil {
		return nil, errors.New("invalid base58 string")
	}
	return NewSigner(bytes)
}

func (s *Signer) Sign(message []byte) ([]byte, error) {
	p := ed25519.PrivateKey(s.privateKey)
	signData, err := p.Sign(rand.Reader, message, crypto.Hash(0))
	if err != nil {
		return nil, err
	}
	return signData, nil
}
