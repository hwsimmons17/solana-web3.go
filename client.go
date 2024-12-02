package solana

import "encoding/base64"

type Client interface {
	GetAccountInfo(pubkey Pubkey) (*Account, error)
	GetBalance(pubkey Pubkey) (uint, error)
	GetProgramAccounts(programId Pubkey) ([]Account, error)
	RequestAirdrop(pubkey Pubkey, lamports uint) (string, error)
	GetTokenAccountBalance(pubkey Pubkey) (UiTokenAmount, error)
	RecentBlockhash() (string, error)
	SendTransaction(transaction Transaction) (string, error)
	SendAndSignTransaction(transaction Transaction) (string, error) //Signs the transaction with the Client's default signer and handles getting the recent blockhash

	Rpc() Rpc
	Signer
	Pubkey
}

type client struct {
	rpc Rpc
	Keypair

	DefaultCommitment Commitment
}

func NewClient(rpc Rpc, keypair Keypair) Client {
	return &client{rpc: rpc, Keypair: keypair, DefaultCommitment: CommitmentFinalized}
}

func (c *client) Rpc() Rpc {
	return c.rpc
}

func (c *client) GetAccountInfo(pubkey Pubkey) (*Account, error) {
	return c.rpc.GetAccountInfo(pubkey, GetAccountInfoConfig{Commitment: &c.DefaultCommitment})
}

func (c *client) GetBalance(pubkey Pubkey) (uint, error) {
	return c.rpc.GetBalance(pubkey, StandardRpcConfig{Commitment: &c.DefaultCommitment})
}

func (c *client) GetProgramAccounts(programId Pubkey) ([]Account, error) {
	return c.rpc.GetProgramAccounts(programId, GetAccountInfoConfig{Commitment: &c.DefaultCommitment})
}

func (c *client) RequestAirdrop(pubkey Pubkey, lamports uint) (string, error) {
	return c.rpc.RequestAirdrop(pubkey, lamports, StandardCommitmentConfig{Commitment: &c.DefaultCommitment})
}

func (c *client) GetTokenAccountBalance(pubkey Pubkey) (UiTokenAmount, error) {
	return c.rpc.GetTokenAccountBalance(pubkey, StandardCommitmentConfig{Commitment: &c.DefaultCommitment})
}

func (c *client) RecentBlockhash() (string, error) {
	blockhash, err := c.rpc.GetLatestBlockhash(StandardRpcConfig{Commitment: &c.DefaultCommitment})
	if err != nil {
		return "", err
	}
	return blockhash.Blockhash, nil
}

func (c *client) SendTransaction(transaction Transaction) (string, error) {
	rawTx := transaction.Serialize()
	txBytes, err := rawTx.Bytes()
	if err != nil {
		return "", err
	}
	encodedTx := base64.StdEncoding.EncodeToString(txBytes)

	encoding := EncodingBase64
	return c.rpc.SendTransaction(encodedTx, SendTransactionConfig{PreflightCommitment: &c.DefaultCommitment, Encoding: &encoding})
}

func (c *client) SendAndSignTransaction(transaction Transaction) (string, error) {
	blockhash, err := c.RecentBlockhash()
	if err != nil {
		return "", err
	}

	transaction.Message.RecentBlockhash = blockhash
	transaction.Sign(c)

	return c.SendTransaction(transaction)
}
