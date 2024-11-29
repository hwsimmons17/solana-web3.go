/*
Keypair is a package for creating and managing Solana keypairs.

This package has a dependency on the `base58` package, and on ed25519 from the standard library.

Example:

	privateKey := os.GetEnv("PRIVATE_KEY")
	keypair, err := keypair.NewKeypairFromBase58(privateKey)
	if err != nil {
		panic(err)
	}
	signature, err := keypair.Signer.Sign([]byte("hello world"))
	if err != nil {
		panic(err)
	}
	fmt.Println(signature)
*/
package keypair

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
