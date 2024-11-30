package rpc

import "solana"

func (r *RpcClient) GetFeeForMessage(msg string, config ...solana.StandardRpcConfig) (*uint, error) {
	return nil, nil
}

func (r *RpcClient) GetLatestBlockhash(config ...solana.StandardRpcConfig) (solana.LatestBlockhash, error) {
	return solana.LatestBlockhash{}, nil
}

func (r *RpcClient) GetSignatureStatuses(signatures []string, config ...solana.GetSignatureStatusesConfig) ([]*solana.SignatureStatus, error) {
	return nil, nil
}

func (r *RpcClient) GetSignaturesForAddress(address solana.Pubkey, config ...solana.GetSignaturesForAddressConfig) ([]solana.TransactionSignature, error) {
	return nil, nil
}

func (r *RpcClient) GetTransaction(transactionSignature string, config ...solana.GetTransactionSignatureConfig) (*solana.Transaction, error) {
	return nil, nil
}

func (r *RpcClient) IsBlockhashValid(blockhash string, config ...solana.StandardRpcConfig) (bool, error) {
	return false, nil
}

func (r *RpcClient) SendTransaction(fullySignedTransaction string, config ...solana.SendTransactionConfig) (string, error) {
	return "", nil
}

func (r *RpcClient) SimulateTransaction(transaction string, config ...solana.SimulateTransactionConfig) (solana.SimulateTransactionResult, error) {
	return solana.SimulateTransactionResult{}, nil
}
