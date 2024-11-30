package rpc

import (
	"solana"
	"solana/dependencies/keypair"
	"testing"
)

func TestGetAccountInfo(t *testing.T) {
	t.Skip("Skipping test that requires network access")
	client := NewRpcClient(solana.RpcEndpointDevnet)
	pubkey := keypair.MustParsePubkey("BLrD8HqBy4vKNvkb28Bijg4y6s8tE49jyVFbfZnmesjY")
	account, err := client.GetAccountInfo(pubkey)
	if err != nil {
		t.Fatal(err)
	}
	if account.Address.String() != "BLrD8HqBy4vKNvkb28Bijg4y6s8tE49jyVFbfZnmesjY" {
		t.Fatal("Unexpected account address")
	}
	if account.Space != 82 {
		t.Fatal("Unexpected account space")
	}
}
