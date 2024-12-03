package rpc

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"math/big"

	"github.com/hwsimmons17/solana-web3.go"
)

type encodedAccount struct {
	Address    solana.PubkeyStr `json:"address"`
	Data       []byte           `json:"data"`
	Executable bool             `json:"executable"`
	Lamports   uint             `json:"lamports"`
	Owner      solana.PubkeyStr `json:"owner"`
	RentEpoch  big.Int          `json:"rentEpoch"`
	Space      int              `json:"space"`
}

func (a *encodedAccount) UnmarshalJSON(data []byte) error {
	aux := &struct {
		Address    string   `json:"address"`
		Data       []string `json:"data"`
		Executable bool     `json:"executable"`
		Lamports   uint     `json:"lamports"`
		Owner      string   `json:"owner"`
		RentEpoch  big.Int  `json:"rentEpoch"`
		Space      int      `json:"space"`
	}{}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}

	a.Address = solana.PubkeyStr(aux.Address)
	a.Executable = aux.Executable
	a.Lamports = uint(aux.Lamports)
	a.Owner = solana.PubkeyStr(aux.Owner)
	a.RentEpoch = aux.RentEpoch
	a.Space = aux.Space
	if len(aux.Data) != 2 {
		return fmt.Errorf("invalid data length, expected 2, got %d", len(aux.Data))
	}
	data, err := base64.StdEncoding.DecodeString(aux.Data[0])
	if err != nil {
		return err
	}
	a.Data = data
	return nil
}

func (r *RpcClient) GetAccountInfo(address solana.Pubkey, config ...solana.GetAccountInfoConfig) (*solana.Account, error) { //Returns all information associated with the account of provided Pubkey
	var res struct {
		Value *encodedAccount `json:"value"`
	}

	//Set the encoding to base64 no matter what
	encoding := solana.EncodingBase64
	params := []interface{}{address.String()}
	if len(config) > 0 {
		config[0].Encoding = &encoding
		params = append(params, config[0])
	} else {
		params = append(params, solana.GetAccountInfoConfig{Encoding: &encoding})
	}
	if err := r.send("getAccountInfo", params, &res); err != nil {
		return nil, err
	}
	//If the account does not exist, return nil
	if res.Value == nil {
		return nil, nil
	}

	return &solana.Account{
		Address:    address,
		Data:       res.Value.Data,
		Executable: res.Value.Executable,
		Lamports:   res.Value.Lamports,
		Owner:      &res.Value.Owner,
		RentEpoch:  res.Value.RentEpoch,
		Space:      res.Value.Space,
	}, nil
}

func (r *RpcClient) GetBalance(address solana.Pubkey, config ...solana.StandardRpcConfig) (uint, error) {
	var res struct {
		Value uint `json:"value"`
	}

	params := []interface{}{address.String()}
	if len(config) > 0 {
		params = append(params, config[0])
	}
	if err := r.send("getBalance", params, &res); err != nil {
		return 0, err
	}

	return res.Value, nil
}

func (r *RpcClient) GetLargestAccounts(config ...solana.GetLargestAccountsConfig) ([]solana.AccountWithBalance, error) {
	var res struct {
		Value []solana.AccountWithBalance `json:"value"`
	}

	params := []interface{}{}
	if len(config) > 0 {
		params = append(params, config[0])
	}
	if err := r.send("getLargestAccounts", params, &res); err != nil {
		return nil, err
	}

	return res.Value, nil
}

func (r *RpcClient) GetMinimumBalanceForRentExemption(accountDataLength uint, config ...solana.StandardCommitmentConfig) (uint, error) {
	var res uint
	params := []interface{}{
		accountDataLength,
	}
	if len(config) > 0 {
		params = append(params, config[0])
	}
	if err := r.send("getMinimumBalanceForRentExemption", params, &res); err != nil {
		return 0, err
	}

	return res, nil
}

func (r *RpcClient) GetMultipleAccounts(pubkeys []solana.Pubkey, config ...solana.GetAccountInfoConfig) ([]*solana.Account, error) {
	var res struct {
		Value []*encodedAccount `json:"value"`
	}

	//Set the encoding to base64 no matter what
	encoding := solana.EncodingBase64
	params := []interface{}{pubkeys}
	if len(config) > 0 {
		config[0].Encoding = &encoding
		params = append(params, config[0])
	} else {
		params = append(params, solana.GetAccountInfoConfig{Encoding: &encoding})
	}
	if err := r.send("getMultipleAccounts", params, &res); err != nil {
		return nil, err
	}
	var accounts []*solana.Account
	for _, value := range res.Value {
		if value == nil {
			accounts = append(accounts, nil)
			continue
		}
		accounts = append(accounts, &solana.Account{
			Address:    &value.Address,
			Data:       value.Data,
			Executable: value.Executable,
			Lamports:   value.Lamports,
			Owner:      &value.Owner,
			RentEpoch:  value.RentEpoch,
			Space:      value.Space,
		})
	}

	return accounts, nil
}

func (r *RpcClient) GetProgramAccounts(programPubkey solana.Pubkey, config ...solana.GetAccountInfoConfig) ([]solana.Account, error) {
	var res []struct {
		Account encodedAccount   `json:"account"`
		Pubkey  solana.PubkeyStr `json:"pubkey"`
	}

	//Set the encoding to base64 no matter what
	encoding := solana.EncodingBase64
	params := []interface{}{programPubkey}
	if len(config) > 0 {
		config[0].Encoding = &encoding
		params = append(params, config[0])
	} else {
		params = append(params, solana.GetAccountInfoConfig{Encoding: &encoding})
	}
	if err := r.send("getProgramAccounts", params, &res); err != nil {
		return nil, err
	}

	var accounts []solana.Account
	for _, account := range res {
		accounts = append(accounts, solana.Account{
			Address:    &account.Pubkey,
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
