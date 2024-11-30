package rpc

import "solana"

func (r *RpcClient) GetInflationGovernor(config ...solana.StandardCommitmentConfig) (solana.InflationGovernor, error) {
	return solana.InflationGovernor{}, nil
}

func (r *RpcClient) GetInflationRate(config ...solana.StandardCommitmentConfig) (solana.InflationRate, error) {
	return solana.InflationRate{}, nil
}

func (r *RpcClient) GetInflationReward(addresses []string, config ...solana.GetInflationRewardConfig) (solana.InflationReward, error) {
	return solana.InflationReward{}, nil
}

func (r *RpcClient) GetStakeMinimumDelegation(config ...solana.StandardCommitmentConfig) (uint, error) {
	return 0, nil
}

func (r *RpcClient) GetSupply(config ...solana.GetSupplyConfig) (solana.Supply, error) {
	return solana.Supply{}, nil
}
