package rpc

import (
	"encoding/base64"
	"errors"
	"fmt"
	"solana"
)

type encodedTransaction struct {
	Meta        *solana.TransactionMeta `json:"meta"`
	Version     *int                    `json:"version"` //Transaction version. Undefined if maxSupportedTransactionVersion is not set in request params. --note can also be "legacy"
	Transaction []string                `json:"transaction"`
}

func (r *RpcClient) GetFeeForMessage(msg []byte, config ...solana.StandardRpcConfig) (*uint, error) {
	var res struct {
		Value *uint `json:"value"`
	}
	params := []interface{}{base64.StdEncoding.EncodeToString([]byte(msg))}
	if len(config) > 0 {
		params = append(params, config[0])
	}
	if err := r.send("getFeeForMessage", params, &res); err != nil {
		return nil, err
	}
	return res.Value, nil
}

func (r *RpcClient) GetLatestBlockhash(config ...solana.StandardRpcConfig) (solana.LatestBlockhash, error) {
	var res struct {
		Value solana.LatestBlockhash `json:"value"`
	}
	params := []interface{}{}
	if len(config) > 0 {
		params = append(params, config[0])
	}
	if err := r.send("getLatestBlockhash", params, &res); err != nil {
		return solana.LatestBlockhash{}, err
	}
	return res.Value, nil
}

func (r *RpcClient) GetSignatureStatuses(signatures []string, config ...solana.GetSignatureStatusesConfig) ([]*solana.SignatureStatus, error) {
	var res struct {
		Value []*solana.SignatureStatus `json:"value"`
	}
	params := []interface{}{signatures}
	if len(config) > 0 {
		params = append(params, config[0])
	}
	if err := r.send("getSignatureStatuses", params, &res); err != nil {
		return nil, err
	}
	return res.Value, nil
}

func (r *RpcClient) GetSignaturesForAddress(address solana.Pubkey, config ...solana.GetSignaturesForAddressConfig) ([]solana.TransactionSignature, error) {
	var res []solana.TransactionSignature
	params := []interface{}{address.String()}
	if len(config) > 0 {
		params = append(params, config[0])
	}
	if err := r.send("getSignaturesForAddress", params, &res); err != nil {
		return nil, err
	}
	return res, nil
}

func (r *RpcClient) GetTransaction(transactionSignature string, config ...solana.GetTransactionSignatureConfig) (*solana.TransactionWithMeta, error) {
	var res *encodedTransaction
	params := []interface{}{transactionSignature}
	encoding := solana.EncodingBase64
	if len(config) > 0 {
		config[0].Encoding = &encoding
		params = append(params, config[0])
	} else {
		params = append(params, solana.GetTransactionSignatureConfig{Encoding: &encoding})
	}
	if err := r.send("getTransaction", params, &res); err != nil {
		return nil, err
	}
	if res == nil {
		return nil, nil
	}
	if len(res.Transaction) == 0 {
		return nil, errors.New("transaction not found")
	}
	transactionData, err := base64.StdEncoding.DecodeString(res.Transaction[0])
	if err != nil {
		return nil, fmt.Errorf("failed to decode transaction data: %v", err)
	}
	rawTx, err := solana.ParseTransactionData(transactionData)
	if err != nil {
		return nil, fmt.Errorf("failed to parse transaction data: %v", err)
	}
	transaction, err := rawTx.Transaction()
	if err != nil {
		return nil, fmt.Errorf("failed to parse transaction: %v", err)
	}

	return &solana.TransactionWithMeta{
		Meta:        res.Meta,
		Version:     res.Version,
		Transaction: transaction,
	}, nil
}

func (r *RpcClient) IsBlockhashValid(blockhash string, config ...solana.StandardRpcConfig) (bool, error) {
	var res struct {
		Value bool `json:"value"`
	}
	params := []interface{}{blockhash}
	if len(config) > 0 {
		params = append(params, config[0])
	}
	if err := r.send("isBlockhashValid", params, &res); err != nil {
		return false, err
	}
	return res.Value, nil
}

// NOTE: This function is not tested yet
func (r *RpcClient) SendTransaction(fullySignedTransaction string, config ...solana.SendTransactionConfig) (string, error) {
	var res string
	params := []interface{}{fullySignedTransaction}
	if len(config) > 0 {
		params = append(params, config[0])
	}
	if err := r.send("sendTransaction", params, &res); err != nil {
		return "", err
	}
	return res, nil
}

// NOTE: This function is not tested yet
func (r *RpcClient) SimulateTransaction(transaction string, config ...solana.SimulateTransactionConfig) (solana.SimulateTransactionResult, error) {
	var res solana.SimulateTransactionResult
	params := []interface{}{transaction}
	if len(config) > 0 {
		params = append(params, config[0])
	}
	if err := r.send("simulateTransaction", params, &res); err != nil {
		return solana.SimulateTransactionResult{}, err
	}
	return res, nil
}
