package rpc

import (
	"errors"
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
	params := []interface{}{address.String()}
	if len(config) > 0 {
		params = append(params, config[0])
	}
	res, err := r.send("getBalance", params)
	if err != nil {
		return 0, err
	}

	return getUint(res.(map[string]interface{}), "value")
}

func (r *RpcClient) GetLargestAccounts(config ...solana.GetLargestAccountsConfig) ([]solana.AccountWithBalance, error) {
	params := []interface{}{}
	if len(config) > 0 {
		params = append(params, config[0])
	}
	res, err := r.send("getLargestAccounts", params)
	if err != nil {
		return nil, err
	}
	value, ok := res.(map[string]interface{})["value"].([]interface{})
	if !ok {
		return nil, errors.New("expected value to be []interface{}")
	}
	accounts := []solana.AccountWithBalance{}
	for _, account := range value {
		accountMap, ok := account.(map[string]interface{})
		if !ok {
			return nil, errors.New("expected account to be map[string]interface{}")
		}
		address, err := getPubkey(accountMap, "address")
		if err != nil {
			return nil, err
		}
		lamports, err := getUint(accountMap, "lamports")
		if err != nil {
			return nil, err
		}
		accounts = append(accounts, solana.AccountWithBalance{
			Address:  address,
			Lamports: lamports,
		})
	}

	return accounts, nil
}

func (r *RpcClient) GetMinimumBalanceForRentExemption(accountDataLength uint, config ...solana.StandardCommitmentConfig) (uint, error) {
	params := []interface{}{
		accountDataLength,
	}
	if len(config) > 0 {
		params = append(params, config[0])
	}
	res, err := r.send("getMinimumBalanceForRentExemption", params)
	if err != nil {
		return 0, err
	}
	result, ok := res.(float64)
	if !ok {
		return 0, errors.New("expected float64")
	}

	return uint(result), nil
}

func (r *RpcClient) GetMultipleAccounts(pubkeys []solana.Pubkey, config ...solana.GetAccountInfoConfig) ([]*solana.EncodedAccount, error) {
	//Set the encoding to base64 no matter what
	encoding := solana.EncodingBase64
	params := []interface{}{pubkeys}
	if len(config) > 0 {
		config[0].Encoding = &encoding
		params = append(params, config[0])
	} else {
		params = append(params, solana.GetAccountInfoConfig{Encoding: &encoding})
	}
	res, err := r.send("getMultipleAccounts", params)
	if err != nil {
		return nil, err
	}
	values, ok := res.(map[string]interface{})["value"].([]interface{})
	if !ok {
		return nil, errors.New("expected value to be []interface{}")
	}

	var accounts []*solana.EncodedAccount
	for i, value := range values {
		account, ok := value.(map[string]interface{})
		if !ok {
			return nil, errors.New("expected account to be map[string]interface{}")
		}

		if account == nil {
			accounts = append(accounts, nil)
			continue
		}

		data, err := getBytes(account, "data")
		if err != nil {
			return nil, err
		}
		executable, err := getBool(account, "executable")
		if err != nil {
			return nil, err
		}
		lamports, err := getUint(account, "lamports")
		if err != nil {
			return nil, err
		}
		owner, err := getPubkey(account, "owner")
		if err != nil {
			return nil, err
		}
		rentEpoch, err := getUint(account, "rentEpoch")
		if err != nil {
			return nil, err
		}
		space, err := getInt(account, "space")
		if err != nil {
			return nil, err
		}
		accounts = append(accounts, &solana.EncodedAccount{
			Address:    pubkeys[i],
			Data:       data,
			Executable: executable,
			Lamports:   lamports,
			Owner:      owner,
			RentEpoch:  rentEpoch,
			Space:      space,
		})
	}

	return accounts, nil
}

func (r *RpcClient) GetProgramAccounts(programPubkey solana.Pubkey, config ...solana.GetAccountInfoConfig) ([]solana.EncodedAccount, error) {
	//Set the encoding to base64 no matter what
	encoding := solana.EncodingBase64
	params := []interface{}{programPubkey}
	if len(config) > 0 {
		config[0].Encoding = &encoding
		params = append(params, config[0])
	} else {
		params = append(params, solana.GetAccountInfoConfig{Encoding: &encoding})
	}
	res, err := r.send("getProgramAccounts", params)
	if err != nil {
		return nil, err
	}

	values, ok := res.([]interface{})
	if !ok {
		return nil, errors.New("expected value to be []interface{}")
	}

	var accounts []solana.EncodedAccount
	for _, value := range values {
		val, ok := value.(map[string]interface{})
		if !ok {
			return nil, errors.New("expected account to be map[string]interface{}")
		}
		account, ok := val["account"].(map[string]interface{})
		if !ok {
			return nil, errors.New("expected account to be map[string]interface{}")
		}
		pubkey, err := getPubkey(val, "pubkey")
		if err != nil {
			return nil, err
		}

		data, err := getBytes(account, "data")
		if err != nil {
			return nil, err
		}
		executable, err := getBool(account, "executable")
		if err != nil {
			return nil, err
		}
		lamports, err := getUint(account, "lamports")
		if err != nil {
			return nil, err
		}
		owner, err := getPubkey(account, "owner")
		if err != nil {
			return nil, err
		}
		rentEpoch, err := getUint(account, "rentEpoch")
		if err != nil {
			return nil, err
		}
		space, err := getInt(account, "space")
		if err != nil {
			return nil, err
		}
		accounts = append(accounts, solana.EncodedAccount{
			Address:    pubkey,
			Data:       data,
			Executable: executable,
			Lamports:   lamports,
			Owner:      owner,
			RentEpoch:  rentEpoch,
			Space:      space,
		})
	}

	return accounts, nil
}
