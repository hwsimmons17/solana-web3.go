package solana

type TokenBalance struct {
	AccountIndex  int           `json:"accountIndex"` //Index of the account in which the token balance is provided for.
	Mint          string        `json:"mint"`         //Pubkey of the token's mint.
	Owner         *string       `json:"owner"`        //Pubkey of token balance's owner.
	ProgramID     *string       `json:"programId"`    //Pubkey of the Token program that owns the account.
	UiTokenAmount UiTokenAmount `json:"uiTokenAmount"`
}

type UiTokenAmount struct {
	Amount         string `json:"amount"`         //Raw amount of tokens as a string, ignoring decimals.
	Decimals       int    `json:"decimals"`       //Number of decimals configured for token's mint.
	UiAmountString string `json:"uiAmountString"` //Token amount as a string, accounting for decimals.
	UiAmount       string `json:"uiAmount"`       //Deprecated: Token amount as a float, accounting for decimals.
}
