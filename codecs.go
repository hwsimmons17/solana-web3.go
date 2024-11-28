package solana

type Decoder interface {
	Decode([]byte) (interface{}, error)
}
