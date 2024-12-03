package solana

import (
	"crypto"
	"crypto/ed25519"
	"crypto/rand"
	"crypto/sha256"
	"errors"
	"math"

	"filippo.io/edwards25519"
	"github.com/mr-tron/base58"
)

// Represents a public key in the Solana blockchain.
type Pubkey interface {
	String() string
	Bytes() []byte
	MarshalJSON() ([]byte, error)
	UnmarshalJSON(data []byte) error
	IsOnCurve() bool
}

// Represents a keypair in the Solana blockchain.
type Keypair struct {
	Pubkey
	Signer
}

// Represents a signer in the Solana blockchain.
type Signer interface {
	Sign(message []byte) ([]byte, error)
}

func NewKeypair(privateKey []byte) (Keypair, error) {
	signer, err := NewSigner(privateKey)
	if err != nil {
		return Keypair{}, err
	}

	pub := ed25519.PrivateKey(privateKey).Public().(ed25519.PublicKey)
	pubkey, err := ParsePubkeyBytes(pub)
	if err != nil {
		return Keypair{}, err
	}

	return Keypair{
		Pubkey: pubkey,
		Signer: signer,
	}, nil
}

func NewKeypairFromBase58(str string) (Keypair, error) {
	signer, err := NewSignerFromBase58(str)
	if err != nil {
		return Keypair{}, err
	}

	privateKey, err := base58.Decode(str)
	if err != nil {
		return Keypair{}, errors.New("invalid base58 string")
	}
	pub := ed25519.PrivateKey(privateKey).Public().(ed25519.PublicKey)
	pubkey, err := ParsePubkeyBytes(pub)
	if err != nil {
		return Keypair{}, err
	}

	return Keypair{
		Pubkey: pubkey,
		Signer: signer,
	}, nil
}

type PubkeyStr string

func ParsePubkey(str string) (Pubkey, error) {
	bytes, err := base58.Decode(str)
	if err != nil {
		return nil, errors.New("invalid base58 string")
	}
	if len(bytes) != 32 {
		return nil, errors.New("invalid pubkey length, expected 32 bytes")
	}

	key := PubkeyStr(str)
	return &key, nil
}

func MustParsePubkey(str string) Pubkey {
	key, err := ParsePubkey(str)
	if err != nil {
		panic(err)
	}
	return key
}

func ParsePubkeyBytes(bytes []byte) (Pubkey, error) {
	str := base58.Encode(bytes)
	return ParsePubkey(str)
}

func (p *PubkeyStr) String() string {
	return string(*p)
}

func (p *PubkeyStr) Bytes() []byte {
	bytes, _ := base58.Decode(p.String())
	return bytes
}

func (p *PubkeyStr) IsOnCurve() bool {
	return IsOnCurve(p.Bytes())
}

func (p *PubkeyStr) MarshalJSON() ([]byte, error) {
	return []byte(`"` + p.String() + `"`), nil
}

func (p *PubkeyStr) UnmarshalJSON(data []byte) error {
	decoded, err := ParsePubkey(string(data[1 : len(data)-1]))
	if err != nil {
		return err
	}
	*p = PubkeyStr(decoded.String())
	return nil
}

type signer struct {
	privateKey []byte
}

func NewSigner(privateKey []byte) (Signer, error) {
	if len(privateKey) != 64 {
		return nil, errors.New("invalid private key length, expected 64 bytes")
	}

	// check if the public key is on the ed25519 curve
	pub := ed25519.PrivateKey(privateKey).Public().(ed25519.PublicKey)
	if !IsOnCurve(pub) {
		return nil, errors.New("corresponding public key is NOT on the ed25519 curve")
	}

	return &signer{privateKey}, nil
}

func IsOnCurve(b []byte) bool {
	if len(b) != ed25519.PublicKeySize {
		return false
	}
	_, err := new(edwards25519.Point).SetBytes(b)
	isOnCurve := err == nil
	return isOnCurve
}

func NewSignerFromBase58(str string) (Signer, error) {
	bytes, err := base58.Decode(str)
	if err != nil {
		return nil, errors.New("invalid base58 string")
	}
	return NewSigner(bytes)
}

func (s *signer) Sign(message []byte) ([]byte, error) {
	p := ed25519.PrivateKey(s.privateKey)
	signData, err := p.Sign(rand.Reader, message, crypto.Hash(0))
	if err != nil {
		return nil, err
	}
	return signData, nil
}

const PDA_MARKER = "ProgramDerivedAddress"

// Find a valid program address and its corresponding bump seed.
func Pda(seed [][]byte, programID Pubkey) (Pubkey, uint8, error) {
	var address Pubkey
	var err error
	bumpSeed := uint8(math.MaxUint8)
	for bumpSeed != 0 {
		address, err = generatePdaCandidate(append(seed, []byte{byte(bumpSeed)}), programID)
		if err == nil {
			return address, bumpSeed, nil
		}
		if err.Error() != "PDA candidate is on the ed25519 curve" {
			return nil, bumpSeed, err
		}
		bumpSeed--
	}
	return nil, bumpSeed, errors.New("unable to find a valid program address")
}
func generatePdaCandidate(seeds [][]byte, programID Pubkey) (Pubkey, error) {
	if len(seeds) > 16 {
		return nil, errors.New("too many seeds, expected 16 or fewer")
	}
	for _, seed := range seeds {
		if len(seed) > 16 {
			return nil, errors.New("too many seeds, expected 16 or fewer")
		}
	}
	buf := []byte{}
	for _, seed := range seeds {
		buf = append(buf, seed...)
	}
	buf = append(buf, programID.Bytes()[:]...)
	buf = append(buf, []byte(PDA_MARKER)...)
	hash := sha256.Sum256(buf)
	pdaCandidate, err := ParsePubkeyBytes(hash[:])
	if err != nil {
		return nil, err
	}
	if pdaCandidate.IsOnCurve() {
		return nil, errors.New("PDA candidate is on the ed25519 curve")
	}
	return pdaCandidate, nil
}
