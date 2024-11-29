package solana

// Represents a public key in the Solana blockchain.
type Pubkey interface {
	String() string
	Bytes() []byte
	MarshalJSON() ([]byte, error)
	UnmarshalJSON(data []byte) error
}

// Represents a keypair in the Solana blockchain.
type Keypair struct {
	PublicKey Pubkey `json:"publicKey"`
	Signer
}

// Represents a signer in the Solana blockchain.
type Signer interface {
	Sign(message []byte) ([]byte, error)
}
