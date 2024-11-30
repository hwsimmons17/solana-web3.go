package rpc

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"solana"
	"solana/dependencies/keypair"
)

type RpcClient struct {
	Endpoint solana.RpcEndpoint
	ID       int
}

func NewRpcClient(endpoint solana.RpcEndpoint) solana.Rpc {
	return &RpcClient{Endpoint: endpoint, ID: 1}
}

func NewRpcClientWithHealthCheck(endpoint solana.RpcEndpoint) (solana.Rpc, error) {
	client := &RpcClient{Endpoint: endpoint}

	if err := client.GetHealth(); err != nil {
		return nil, err
	}

	return &RpcClient{Endpoint: endpoint}, nil
}

func (r *RpcClient) GetHealth() error {
	res, err := r.send("getHealth", nil)
	if err != nil {
		return err
	}

	if res != "ok" {
		return fmt.Errorf("unexpected health response: %v", res)
	}

	return nil
}

func (r *RpcClient) GetIdentity() (solana.Pubkey, error) {
	res, err := r.send("getIdentity", nil)
	if err != nil {
		return nil, err
	}

	resMap, ok := res.(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("expected map[string]interface{}, got %T", res)
	}
	identStr := resMap["identity"].(string)
	identity, err := keypair.ParsePubkey(identStr)
	if err != nil {
		return nil, err
	}

	return identity, nil
}

func (r *RpcClient) GetVersion() (solana.Version, error) {
	res, err := r.send("getVersion", nil)
	if err != nil {
		return solana.Version{}, err
	}

	resMap, ok := res.(map[string]interface{})
	if !ok {
		return solana.Version{}, fmt.Errorf("expected map[string]interface{}, got %T", res)
	}
	solanaCore, ok := resMap["solana-core"].(string)
	if !ok {
		return solana.Version{}, fmt.Errorf("expected solana-core to be string, got %T", resMap["solana-core"])
	}
	featureSetFloat, ok := resMap["feature-set"].(float64)
	if !ok {
		return solana.Version{}, fmt.Errorf("expected feature-set to be float64, got %T", resMap["feature-set"])
	}
	featureSet := uint(featureSetFloat)

	return solana.Version{
		SolanaCore: solanaCore,
		FeatureSet: featureSet,
	}, nil
}

func (r *RpcClient) RequestAirdrop(destinationAddress solana.Pubkey, lamports uint, config ...solana.StandardCommitmentConfig) (string, error) {
	params := []interface{}{destinationAddress.String(), lamports}
	if len(config) > 0 {
		params = append(params, config[0])
	}
	res, err := r.send("requestAirdrop", params)
	if err != nil {
		return "", err
	}

	result, ok := res.(string)
	if !ok {
		return "", fmt.Errorf("expected string, got %T", res)
	}

	return result, nil
}

func (r *RpcClient) incrementID() {
	r.ID++
}

type rpcReq struct {
	ID      int           `json:"id"`
	Jsonrpc string        `json:"jsonrpc"`
	Method  string        `json:"method"`
	Params  []interface{} `json:"params,omitempty"`
}

type rpcResp struct {
	Jsonrpc string      `json:"jsonrpc"`
	ID      int         `json:"id"`
	Result  interface{} `json:"result"`
	Error   *struct {
		Code    int         `json:"code"`
		Message string      `json:"message"`
		Data    interface{} `json:"data"`
	} `json:"error"`
}

func (r *RpcClient) send(method string, params []interface{}) (interface{}, error) {
	body := rpcReq{
		ID:      r.ID,
		Jsonrpc: "2.0",
		Method:  method,
		Params:  params,
	}
	r.incrementID()

	data, err := json.Marshal(body)
	if err != nil {
		return "", err
	}

	client := http.Client{}
	req, err := http.NewRequest(http.MethodPost, string(r.Endpoint), bytes.NewBuffer([]byte(data)))
	if err != nil {
		return "", err
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	var result rpcResp
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return "", err
	}

	if result.Error != nil {
		return "", fmt.Errorf("rpc request failed. Code: %d, Message: %s, Data: %v", result.Error.Code, result.Error.Message, result.Error.Data)
	}

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("rpc request failed. Status code: %d", resp.StatusCode)
	}

	return result.Result, nil
}

func getValueMap(res interface{}) (map[string]interface{}, error) {
	resMap, ok := res.(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("expected map[string]interface{}, got %T", res)
	}
	valueMap, ok := resMap["value"].(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("expected value to be map[string]interface{}, got %T", resMap["value"])
	}

	return valueMap, nil
}

func getBytes(resMap map[string]interface{}, key string) ([]byte, error) {
	dataMap, ok := resMap[key]
	if !ok {
		return nil, fmt.Errorf("expected data to be present in account info")
	}
	dataArr, ok := dataMap.([]interface{})
	if !ok {
		return nil, fmt.Errorf("expected data to be array, got %T", resMap[key])
	}
	if len(dataArr) != 2 {
		return nil, fmt.Errorf("expected data to be array of length 2, got %d", len(dataArr))
	}
	if dataArr[1] != "base64" {
		return nil, fmt.Errorf("expected data encoding to be base64, got %s", dataArr[1])
	}

	var bytes []byte
	if dataArr[0] != nil {
		dataStr, ok := dataArr[0].(string)
		if !ok {
			return nil, fmt.Errorf("expected data to be string, got %T", dataArr[0])
		}
		d, err := base64.StdEncoding.DecodeString(dataStr)
		if err != nil {
			return nil, fmt.Errorf("error decoding data: %v", err)
		}
		bytes = d
	}
	return bytes, nil
}

func getBool(resMap map[string]interface{}, key string) (bool, error) {
	val, ok := resMap[key]
	if !ok {
		return false, fmt.Errorf("expected %s to be present in account info, %v", key, resMap)
	}
	b, ok := val.(bool)
	if !ok {
		return false, fmt.Errorf("expected %s to be bool, got %T", key, val)
	}
	return b, nil
}

func getUint(resMap map[string]interface{}, key string) (uint, error) {
	val, ok := resMap[key]
	if !ok {
		return 0, fmt.Errorf("expected %s to be present in account info", key)
	}
	f, ok := val.(float64)
	if !ok {
		return 0, fmt.Errorf("expected %s to be float64, got %T", key, val)
	}
	return uint(f), nil
}

func getInt(resMap map[string]interface{}, key string) (int, error) {
	val, ok := resMap[key]
	if !ok {
		return 0, fmt.Errorf("expected %s to be present in account info", key)
	}
	f, ok := val.(float64)
	if !ok {
		return 0, fmt.Errorf("expected %s to be float64, got %T", key, val)
	}
	return int(f), nil
}

func getPubkey(resMap map[string]interface{}, key string) (solana.Pubkey, error) {
	val, ok := resMap[key]
	if !ok {
		return nil, fmt.Errorf("expected %s to be present in account info", key)
	}
	str, ok := val.(string)
	if !ok {
		return nil, fmt.Errorf("expected %s to be string, got %T", key, val)
	}
	pubkey, err := keypair.ParsePubkey(str)
	if err != nil {
		return nil, fmt.Errorf("error parsing %s pubkey: %v", key, err)
	}
	return pubkey, nil
}
