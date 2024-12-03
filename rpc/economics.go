package rpc

import "github.com/hwsimmons17/solana-web3.go"

func (r *RpcClient) GetInflationGovernor(config ...solana.StandardCommitmentConfig) (solana.InflationGovernor, error) {
	var res solana.InflationGovernor
	params := []interface{}{}
	if len(config) > 0 {
		params = append(params, config[0])
	}
	if err := r.send("getInflationGovernor", params, &res); err != nil {
		return solana.InflationGovernor{}, err
	}
	return res, nil
}

func (r *RpcClient) GetInflationRate(config ...solana.StandardCommitmentConfig) (solana.InflationRate, error) {
	var res solana.InflationRate
	params := []interface{}{}
	if len(config) > 0 {
		params = append(params, config[0])
	}
	if err := r.send("getInflationRate", params, &res); err != nil {
		return solana.InflationRate{}, err
	}
	return res, nil
}

func (r *RpcClient) GetInflationReward(addresses []solana.Pubkey, config ...solana.GetInflationRewardConfig) ([]*solana.InflationReward, error) {
	var res []*solana.InflationReward
	params := []interface{}{}
	var strs []string
	for _, address := range addresses {
		strs = append(strs, address.String())
	}
	params = append(params, strs)
	if len(config) > 0 {
		params = append(params, config[0])
	}
	if err := r.send("getInflationReward", params, &res); err != nil {
		return nil, err
	}
	return res, nil
}

func (r *RpcClient) GetStakeMinimumDelegation(config ...solana.StandardCommitmentConfig) (uint, error) {
	var res struct {
		Value uint `json:"value"`
	}
	params := []interface{}{}
	if len(config) > 0 {
		params = append(params, config[0])
	}
	if err := r.send("getStakeMinimumDelegation", params, &res); err != nil {
		return 0, err
	}
	return res.Value, nil
}

func (r *RpcClient) GetSupply(config ...solana.GetSupplyConfig) (solana.Supply, error) {
	var res struct {
		Value solana.Supply `json:"value"`
	}
	params := []interface{}{}
	if len(config) > 0 {
		params = append(params, config[0])
	}
	if err := r.send("getSupply", params, &res); err != nil {
		return solana.Supply{}, err
	}
	return res.Value, nil
}
