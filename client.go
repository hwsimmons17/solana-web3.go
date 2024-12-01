package solana

type SolanaClient interface {
	GetAccountInfo(pubkey Pubkey) (Account, error)
	GetBalance(pubkey Pubkey) (uint, error)
	GetProgramAccounts(programId Pubkey) ([]Account, error)
	RequestAirdrop(pubkey Pubkey, lamports uint) (string, error)
	GetTokenAccountBalance(pubkey Pubkey) (UiTokenAmount, error)
	SendTransaction(transaction RawTransaction) (string, error)
}
