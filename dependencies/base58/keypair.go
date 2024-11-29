package base58

import (
	"crypto/ed25519"
	"errors"
	"solana"

	"github.com/mr-tron/base58"
)

func NewKeypair(privateKey []byte) (solana.Keypair, error) {
	signer, err := NewSigner(privateKey)
	if err != nil {
		return solana.Keypair{}, err
	}

	pub := ed25519.PrivateKey(privateKey).Public().(ed25519.PublicKey)
	pubkey, err := ParsePubkeyBytes(pub)
	if err != nil {
		return solana.Keypair{}, err
	}

	return solana.Keypair{
		PublicKey: pubkey,
		Signer:    signer,
	}, nil
}

func NewKeypairFromBase58(str string) (solana.Keypair, error) {
	signer, err := NewSignerFromBase58(str)
	if err != nil {
		return solana.Keypair{}, err
	}

	privateKey, err := base58.Decode(str)
	if err != nil {
		return solana.Keypair{}, errors.New("invalid base58 string")
	}
	pub := ed25519.PrivateKey(privateKey).Public().(ed25519.PublicKey)
	pubkey, err := ParsePubkeyBytes(pub)
	if err != nil {
		return solana.Keypair{}, err
	}

	return solana.Keypair{
		PublicKey: pubkey,
		Signer:    signer,
	}, nil
}
