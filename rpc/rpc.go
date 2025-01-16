package rpc

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/hwsimmons17/solana-web3.go"
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
	var res string
	if err := r.send("getHealth", nil, &res); err != nil {
		return err
	}

	if res != "ok" {
		return fmt.Errorf("unexpected health response: %v", res)
	}

	return nil
}

func (r *RpcClient) GetIdentity() (solana.Pubkey, error) {
	var res struct {
		Identity solana.PubkeyStr `json:"identity"`
	}
	if err := r.send("getIdentity", nil, &res); err != nil {
		return nil, err
	}

	return &res.Identity, nil
}

func (r *RpcClient) GetVersion() (solana.Version, error) {
	var res solana.Version
	if err := r.send("getVersion", nil, &res); err != nil {
		return solana.Version{}, err
	}

	return res, nil
}

func (r *RpcClient) RequestAirdrop(destinationAddress solana.Pubkey, lamports uint, config ...solana.StandardCommitmentConfig) (string, error) {
	params := []interface{}{destinationAddress.String(), lamports}
	if len(config) > 0 {
		params = append(params, config[0])
	}
	var res string
	if err := r.send("requestAirdrop", params, &res); err != nil {
		return "", err
	}

	return res, nil
}

func (r *RpcClient) incrementID() {
	r.ID++
}

type rpcReq struct {
	ID      int    `json:"id"`
	Jsonrpc string `json:"jsonrpc"`
	Method  string `json:"method"`
	Params  any    `json:"params,omitempty"`
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

func (r *RpcClient) send(method string, params any, res interface{}) error {
	body := rpcReq{
		ID:      r.ID,
		Jsonrpc: "2.0",
		Method:  method,
		Params:  params,
	}
	r.incrementID()

	data, err := json.Marshal(body)
	if err != nil {
		return err
	}
	log.Println(string(data))

	client := http.Client{}
	req, err := http.NewRequest(http.MethodPost, string(r.Endpoint), bytes.NewBuffer([]byte(data)))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	var result rpcResp
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return err
	}

	if result.Error != nil {
		if result.Error.Code == 429 {
			retryAfter, err := time.ParseDuration(resp.Header.Get("Retry-After") + "s")
			if err == nil {
				time.Sleep(retryAfter)
				return r.send(method, params, res)
			}
		}
		return fmt.Errorf("rpc request failed. Code: %d, Message: %s, Data: %v", result.Error.Code, result.Error.Message, result.Error.Data)
	}

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("rpc request failed. Status code: %d", resp.StatusCode)
	}

	resJson, err := json.Marshal(result.Result)
	if err != nil {
		return err
	}
	if err := json.Unmarshal(resJson, res); err != nil {
		return err
	}

	return nil
}
