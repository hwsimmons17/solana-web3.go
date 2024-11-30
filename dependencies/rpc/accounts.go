package rpc

import (
	"solana"
)

func (r *RpcClient) GetAccountInfo(address solana.Pubkey, config ...solana.GetAccountInfoConfig) (*solana.EncodedAccount, error) { //Returns all information associated with the account of provided Pubkey
	//Set the encoding to base64 no matter what
	encoding := solana.EncodingBase64
	params := []interface{}{address.String()}
	if len(config) > 0 {
		config[0].Encoding = &encoding
		params = append(params, config[0])
	} else {
		params = append(params, solana.GetAccountInfoConfig{Encoding: &encoding})
	}
	res, err := r.send("getAccountInfo", params)
	if err != nil {
		return nil, err
	}
	//If the account does not exist, return nil
	if res == nil {
		return nil, nil
	}
	valueMap, err := getValueMap(res)
	if err != nil {
		return nil, err
	}

	data, err := getBytes(valueMap, "data")
	if err != nil {
		return nil, err
	}
	executable, err := getBool(valueMap, "executable")
	if err != nil {
		return nil, err
	}
	lamports, err := getUint(valueMap, "lamports")
	if err != nil {
		return nil, err
	}
	owner, err := getPubkey(valueMap, "owner")
	if err != nil {
		return nil, err
	}
	rentEpoch, err := getUint(valueMap, "rentEpoch")
	if err != nil {
		return nil, err
	}
	space, err := getInt(valueMap, "space")
	if err != nil {
		return nil, err
	}

	return &solana.EncodedAccount{
		Address:    address,
		Data:       data,
		Executable: executable,
		Lamports:   lamports,
		Owner:      owner,
		RentEpoch:  rentEpoch,
		Space:      space,
	}, nil
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
