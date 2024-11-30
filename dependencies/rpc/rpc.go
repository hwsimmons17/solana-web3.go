package rpc

import "solana"

type RpcClient struct {
	Endpoint solana.RpcEndpoint
}

func NewRpcClient(endpoint solana.RpcEndpoint) solana.Rpc {
	return &RpcClient{Endpoint: endpoint}
}

func NewRpcClientWithChecks(endpoint solana.RpcEndpoint) (solana.Rpc, error) {
	//TODO: Call GetHealth() and check if the RPC endpoint is healthy

	return &RpcClient{Endpoint: endpoint}, nil
}

func (r *RpcClient) GetHealth() error {
	return nil
}

func (r *RpcClient) GetIdentity() (string, error) {
	return "", nil
}

func (r *RpcClient) GetVersion() (solana.Version, error) {
	return solana.Version{}, nil
}

func (r *RpcClient) RequestAirdrop(destinationAddress string, lamports uint, config ...solana.StandardCommitmentConfig) (string, error) {
	return "", nil
}
