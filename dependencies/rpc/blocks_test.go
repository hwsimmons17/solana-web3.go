package rpc

import (
	"solana"
	"testing"
)

func TestGetBlock(t *testing.T) {
	t.Skip("Skipping test that requires network access")
	client := NewRpcClient(solana.RpcEndpointDevnet)
	block, err := client.GetBlock(343820691)
	if err != nil {
		t.Fatal(err)
	}
	t.Fatal(block)
}
