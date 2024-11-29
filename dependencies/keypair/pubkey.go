package keypair

import (
	"errors"
	"solana"

	"github.com/mr-tron/base58"
)

type Pubkey string

func ParsePubkey(str string) (solana.Pubkey, error) {
	bytes, err := base58.Decode(str)
	if err != nil {
		return nil, errors.New("invalid base58 string")
	}
	if len(bytes) != 32 {
		return nil, errors.New("invalid pubkey length, expected 32 bytes")
	}

	key := Pubkey(str)
	return &key, nil
}

func ParsePubkeyBytes(bytes []byte) (solana.Pubkey, error) {
	str := base58.Encode(bytes)
	return ParsePubkey(str)
}

func (p *Pubkey) String() string {
	return string(*p)
}

func (p *Pubkey) Bytes() []byte {
	bytes, _ := base58.Decode(p.String())
	return bytes
}

func (p *Pubkey) IsOnCurve() bool {
	return IsOnCurve(p.Bytes())
}

func (p *Pubkey) MarshalJSON() ([]byte, error) {
	return []byte(`"` + p.String() + `"`), nil
}

func (p *Pubkey) UnmarshalJSON(data []byte) error {
	decoded, err := ParsePubkey(string(data[1 : len(data)-1]))
	if err != nil {
		return err
	}
	*p = Pubkey(decoded.String())
	return nil
}
