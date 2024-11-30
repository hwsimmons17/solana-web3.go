package rpc

import "solana"

func (r *RpcClient) GetTokenAccountBalance(address solana.Pubkey, config ...solana.StandardCommitmentConfig) (solana.UiTokenAmount, error) {
	return solana.UiTokenAmount{}, nil
}

func (r *RpcClient) GetTokenAccountsByDelegate(delegateAddress solana.Pubkey, opts *solana.GetTokenAccountsByDelegateConfig, config ...solana.GetAccountInfoConfig) ([]solana.EncodedAccount, error) {
	return nil, nil
}

func (r *RpcClient) GetTokenAccountsByOwner(ownerAddress solana.Pubkey, opts *solana.GetTokenAccountsByDelegateConfig, config ...solana.GetAccountInfoConfig) ([]solana.EncodedAccount, error) {
	return nil, nil
}

func (r *RpcClient) GetTokenLargestAccounts(mintAddress solana.Pubkey, config ...solana.StandardCommitmentConfig) ([]solana.UiTokenAmount, error) {
	return nil, nil
}

func (r *RpcClient) GetTokenSupply(mintAddress solana.Pubkey, config ...solana.StandardCommitmentConfig) (solana.UiTokenAmount, error) {
	return solana.UiTokenAmount{}, nil
}
