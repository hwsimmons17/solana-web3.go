package solana

import "errors"

type AccountWithBalance struct {
	Address  string `json:"address"` //Base-58 encoded address of the account
	Lamports uint   `json:"balance"` //Number of lamports in the account, as a u64
}

type Supply struct {
	Total                  uint     `json:"total"`                  //Total supply in lamports
	Circulating            uint     `json:"circulating"`            //Circulating supply in lamports
	NonCirculating         uint     `json:"nonCirculating"`         //Non-circulating supply in lamports
	NonCirculatingAccounts []string `json:"nonCirculatingAccounts"` //An array of account addresses of non-circulating accounts, as strings. If excludeNonCirculatingAccountsList is enabled, the returned array will be empty.
}

// The amount of bytes required to store the base account information without its data.
const BASE_ACCOUNT_SIZE = 128

// Describe the generic account details applicable to every account.
type BaseAccount struct {
	Executable bool   `json:"executable"`
	Lamports   uint   `json:"lamports"`
	Owner      Pubkey `json:"owner"`
	RentEpoch  uint64 `json:"rentEpoch"`
	Space      int    `json:"space"`
}

// Defines a Solana account with its generic details or encoded data.
type Account[T any] struct {
	Address    Pubkey `json:"address"`
	Data       []byte `json:"data"`
	ParsedData T      `json:"parsedData"`
	Executable bool   `json:"executable"`
	Lamports   uint   `json:"lamports"`
	Owner      Pubkey `json:"owner"`
	RentEpoch  uint64 `json:"rentEpoch"`
	Space      int    `json:"space"`
}

// Defines a Solana account with its generic details and encoded data.
type EncodedAccount struct {
	Address    Pubkey `json:"address"`
	Data       []byte `json:"data"`
	Executable bool   `json:"executable"`
	Lamports   uint   `json:"lamports"`
	Owner      Pubkey `json:"owner"`
	RentEpoch  uint64 `json:"rentEpoch"`
	Space      int    `json:"space"`
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
		Address:    encodedAccount.Address,
		Data:       encodedAccount.Data,
		ParsedData: data,
		Executable: encodedAccount.Executable,
		Lamports:   encodedAccount.Lamports,
		Owner:      encodedAccount.Owner,
		RentEpoch:  encodedAccount.RentEpoch,
		Space:      encodedAccount.Space,
	}, nil
}
