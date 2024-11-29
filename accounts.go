package solana

import "errors"

type AccountWithBalance struct {
	Address  string `json:"address"` //Base-58 encoded address of the account
	Lamports uint   `json:"balance"` //Number of lamports in the account, as a u64
}

// The amount of bytes required to store the base account information without its data.
const BASE_ACCOUNT_SIZE = 128

// Describe the generic account details applicable to every account.
type BaseAccount struct {
	Executable bool    `json:"executable"`
	Lamports   uint    `json:"lamports"`
	Owner      Address `json:"owner"`
	RentEpoch  uint64  `json:"rentEpoch"`
	Space      int     `json:"space"`
}

// Defines a Solana account with its generic details or encoded data.
type Account[T any] struct {
	Address    Address `json:"address"`
	Data       []byte  `json:"data"`
	ParsedData T       `json:"parsedData"`
	BaseAccount
}

// Defines a Solana account with its generic details and encoded data.
type EncodedAccount struct {
	Address Address `json:"address"`
	Data    []byte  `json:"data"`
	BaseAccount
}

func DecodeAccount[T any](
	encodedAccount EncodedAccount,
	decoder Decoder,
) (Account[T], error) {
	parsedData, err := decoder.Decode(encodedAccount.Data)
	if err != nil {
		//TODO: Improve error messages
		return Account[T]{}, errors.New("SOLANA_ERROR__ACCOUNTS__FAILED_TO_DECODE_ACCOUNT")
	}

	data, ok := parsedData.(T)
	if !ok {
		return Account[T]{}, errors.New("SOLANA_ERROR__ACCOUNTS__FAILED_TO_DECODE_ACCOUNT")
	}

	return Account[T]{
		Address:     encodedAccount.Address,
		Data:        encodedAccount.Data,
		ParsedData:  data,
		BaseAccount: encodedAccount.BaseAccount,
	}, nil
}
