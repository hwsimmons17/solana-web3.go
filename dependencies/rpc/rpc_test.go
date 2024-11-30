package rpc

import (
	"solana"
	"testing"
)

func TestGetHealth(t *testing.T) {
	t.Skip("Skipping test that requires network access")
	client := NewRpcClient(solana.RpcEndpointDevnet)
	err := client.GetHealth()
	if err != nil {
		t.Fatal(err)
	}
}

func TestGetIdentity(t *testing.T) {
	t.Skip("Skipping test that requires network access")
	client := NewRpcClient(solana.RpcEndpointDevnet)
	identity, err := client.GetIdentity()
	if err != nil {
		t.Fatal(err)
	}
	if identity.String() != "39cvwUEpgka9bU7Sn4my82VViMDWaCxi4YoPevfZxLf3" {
		t.Fatal("Unexpected identity")
	}
}

func TestGetVersion(t *testing.T) {
	t.Skip("Skipping test that requires network access")
	client := NewRpcClient(solana.RpcEndpointDevnet)
	version, err := client.GetVersion()
	if err != nil {
		t.Fatal(err)
	}
	if version.SolanaCore != "2.0.15" {
		t.Fatal("Unexpected version", version)
	}
}

func TestRequestAirdrop(t *testing.T) {
	t.Skip("Skipping test that requires network access")
	client := NewRpcClient(solana.RpcEndpointDevnet)
	comit := solana.CommitmentConfirmed
	_, err := client.RequestAirdrop("5oNDL3swdJJF1g9DzJiZ4ynHXgszjAEpUkxVYejchzrY", 1000000000, solana.StandardCommitmentConfig{
		Commitment: &comit,
	})
	if err != nil {
		t.Fatal(err)
	}
}
