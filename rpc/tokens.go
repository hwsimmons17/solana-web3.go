package rpc

import (
	"log"

	"github.com/hwsimmons17/solana-web3.go"
)

func (r *RpcClient) GetTokenAccountBalance(address solana.Pubkey, config ...solana.StandardCommitmentConfig) (solana.UiTokenAmount, error) {
	var res struct {
		Value solana.UiTokenAmount `json:"value"`
	}
	params := []interface{}{address.String()}
	if len(config) > 0 {
		params = append(params, config[0])
	}
	if err := r.send("getTokenAccountBalance", params, &res); err != nil {
		return solana.UiTokenAmount{}, err
	}
	return res.Value, nil
}

func (r *RpcClient) GetTokenAccountsByDelegate(delegateAddress solana.Pubkey, opts solana.GetTokenAccountsByDelegateConfig, config ...solana.GetAccountInfoConfig) ([]solana.Account, error) {
	var res struct {
		Value []struct {
			Account encodedAccount `json:"account"`
			Pubkey  string         `json:"pubkey"`
		} `json:"value"`
	}
	params := []interface{}{delegateAddress.String(), opts}
	encoding := solana.EncodingBase64
	if len(config) > 0 {
		config[0].Encoding = &encoding
		params = append(params, config[0])
	} else {
		params = append(params, solana.GetAccountInfoConfig{Encoding: &encoding})
	}
	if err := r.send("getTokenAccountsByDelegate", params, &res); err != nil {
		return nil, err
	}

	var accounts []solana.Account
	for _, account := range res.Value {
		pubkey, err := solana.ParsePubkey(account.Pubkey)
		if err != nil {
			return nil, err
		}

		accounts = append(accounts, solana.Account{
			Address:    pubkey,
			Data:       account.Account.Data,
			Executable: account.Account.Executable,
			Lamports:   account.Account.Lamports,
			Owner:      &account.Account.Owner,
			RentEpoch:  account.Account.RentEpoch,
			Space:      account.Account.Space,
		})
	}
	return accounts, nil
}

func (r *RpcClient) GetTokenAccountsByOwner(ownerAddress solana.Pubkey, opts solana.GetTokenAccountsByDelegateConfig, config ...solana.GetAccountInfoConfig) ([]solana.Account, error) {
	var res struct {
		Value []struct {
			Account encodedAccount `json:"account"`
			Pubkey  string         `json:"pubkey"`
		} `json:"value"`
	}
	params := []interface{}{ownerAddress.String(), opts}
	encoding := solana.EncodingBase64
	if len(config) > 0 {
		config[0].Encoding = &encoding
		params = append(params, config[0])
	} else {
		params = append(params, solana.GetAccountInfoConfig{Encoding: &encoding})
	}
	if err := r.send("getTokenAccountsByOwner", params, &res); err != nil {
		return nil, err
	}

	var accounts []solana.Account
	for _, account := range res.Value {
		pubkey, err := solana.ParsePubkey(account.Pubkey)
		if err != nil {
			return nil, err
		}

		accounts = append(accounts, solana.Account{
			Address:    pubkey,
			Data:       account.Account.Data,
			Executable: account.Account.Executable,
			Lamports:   account.Account.Lamports,
			Owner:      &account.Account.Owner,
			RentEpoch:  account.Account.RentEpoch,
			Space:      account.Account.Space,
		})
	}
	return accounts, nil
}

func (r *RpcClient) GetTokenLargestAccounts(mintAddress solana.Pubkey, config ...solana.StandardCommitmentConfig) ([]solana.UiTokenAmount, error) {
	var res struct {
		Value []solana.UiTokenAmount `json:"value"`
	}
	params := []interface{}{mintAddress.String()}
	if len(config) > 0 {
		params = append(params, config[0])
	}
	if err := r.send("getTokenLargestAccounts", params, &res); err != nil {
		return nil, err
	}

	return res.Value, nil
}

func (r *RpcClient) GetTokenSupply(mintAddress solana.Pubkey, config ...solana.StandardCommitmentConfig) (solana.UiTokenAmount, error) {
	var res struct {
		Value solana.UiTokenAmount `json:"value"`
	}
	params := []interface{}{mintAddress.String()}
	if len(config) > 0 {
		params = append(params, config[0])
	}
	if err := r.send("getTokenSupply", params, &res); err != nil {
		return solana.UiTokenAmount{}, err
	}
	return res.Value, nil
}

func (r *RpcClient) GetAsset(pubkey solana.Pubkey, config ...solana.GetAssetConfig) (solana.GetAssetResult, error) {
	var res solana.GetAssetResult

	type getAssetParams struct {
		ID             string                 `json:"id"`
		DisplayOptions *solana.GetAssetConfig `json:"displayOptions,omitempty"`
	}
	params := getAssetParams{
		ID: pubkey.String(),
	}
	if len(config) > 0 {
		params.DisplayOptions = &config[0]
	}
	log.Println(params)
	if err := r.send("getAsset", params, &res); err != nil {
		return res, err
	}
	return res, nil
}
