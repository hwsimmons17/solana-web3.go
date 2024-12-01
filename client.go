package solana

type Client interface {
	GetAccountInfo(pubkey Pubkey) (Account, error)
	GetBalance(pubkey Pubkey) (uint, error)
	GetProgramAccounts(programId Pubkey) ([]Account, error)
	RequestAirdrop(pubkey Pubkey, lamports uint) (string, error)
	GetTokenAccountBalance(pubkey Pubkey) (UiTokenAmount, error)
	SendTransaction(transaction Transaction) (string, error)
	SendAndSignTransaction(transaction Transaction, signers ...Signer) (string, error)

	Rpc() Rpc
}

// type client struct {
// 	rpc Rpc
// }
