package keypair

import (
	"log"
	"testing"
)

func TestNewSigner(t *testing.T) {
	invalidLength := []byte{0, 1, 2, 3, 4, 5}
	if _, err := NewSigner(invalidLength); err == nil {
		log.Fatal("Expected error for invalid pubkey length")
	}

	notOnCurve := []byte{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36, 37, 38, 39, 40, 41, 42, 43, 44, 45, 46, 47, 48, 49, 50, 51, 52, 53, 54, 55, 56, 57, 58, 59, 60, 61, 62, 63}
	if _, err := NewSigner(notOnCurve); err == nil {
		log.Fatal("Expected no error for valid pubkey", err)
	}

	//NOTE: Not to ever be used anywhere
	valid := []byte{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36, 37, 38, 39, 40, 41, 42, 43, 44, 45, 46, 47, 48, 49, 50, 51, 52, 53, 54, 55, 56, 57, 58, 59, 60, 61, 62, 61}
	if _, err := NewSigner(valid); err != nil {
		log.Fatal("Expected no error for valid pubkey", err)
	}
}

func TestNewSignerFromBase58(t *testing.T) {
	invalidBase58 := "_($#!@#$)"
	if _, err := NewSignerFromBase58(invalidBase58); err == nil {
		log.Fatal("Expected error for invalid base58 string")
	}

	//NOTE: This is the private key from above. Not to ever be used anywhere
	valid := "1GMkH3brNXiNNs1tiFZHu4yZSRrzJwxi5wB9bHFtMinfCXNnR1adh8Vo8NTheK4evneedH4qmvjeqcBBNAefgQ"
	if _, err := NewSignerFromBase58(valid); err != nil {
		log.Fatal("Expected no error for valid base58", err)
	}
}

func TestSign(t *testing.T) {
	//NOTE: This is the private key from above. Not to ever be used anywhere
	signer, _ := NewSignerFromBase58("1GMkH3brNXiNNs1tiFZHu4yZSRrzJwxi5wB9bHFtMinfCXNnR1adh8Vo8NTheK4evneedH4qmvjeqcBBNAefgQ")
	message := []byte("hello world")
	signature, err := signer.Sign(message)
	if err != nil {
		log.Fatal("Expected no error signing message", err)
	}
	if len(signature) != 64 {
		log.Fatal("Expected signature to be 64 bytes")
	}
}
