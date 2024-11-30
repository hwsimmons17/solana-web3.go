package rpc

import "solana"

func (r *RpcClient) GetAccountInfo(address solana.Pubkey, config ...solana.GetAccountInfoConfig) (solana.EncodedAccount, error) { //Returns all information associated with the account of provided Pubkey

	return solana.EncodedAccount{}, nil
}

func (r *RpcClient) GetBalance(address solana.Pubkey, config ...solana.StandardRpcConfig) (uint, error) {
	return 0, nil
}

func (r *RpcClient) GetLargestAccounts(config ...solana.GetLargestAccountsConfig) ([]solana.AccountWithBalance, error) {
	return nil, nil
}

func (r *RpcClient) GetMinimumBalanceForRentExemption(accountDataLength uint, config ...solana.StandardCommitmentConfig) (uint, error) {
	return 0, nil
}

func (r *RpcClient) GetMultipleAccounts(pubkeys []string, config ...solana.GetAccountInfoConfig) ([]solana.EncodedAccount, error) {
	return nil, nil
}

func (r *RpcClient) GetProgramAccounts(programPubkey string, config ...solana.GetAccountInfoConfig) ([]solana.EncodedAccount, error) {
	return nil, nil
}
