package rpc

import (
	"solana"
	"testing"
)

func TestGetTokenAccountBalance(t *testing.T) {
	t.Skip("Skipping test that requires network access")
	client := NewRpcClient(solana.RpcEndpointDevnet)
	balance, err := client.GetTokenAccountBalance(solana.MustParsePubkey("7Fg8XQBVY4z7gPzecGo7abbHZbHj3iFfGozXsz1VcvKk"))
	if err != nil {
		t.Fatal(err)
	}
	t.Fatal(balance)
}

func TestGetTokenAccountsByDelegate(t *testing.T) {
	t.Skip("Skipping test that requires network access")
	client := NewRpcClient(solana.RpcEndpointDevnet)
	programID := solana.MustParsePubkey("TokenkegQfeZyiNwAJbNbGKPFXCWuBvf9Ss623VQ5DA")
	accounts, err := client.GetTokenAccountsByDelegate(solana.MustParsePubkey("GR16g49y2fEjRQD612ryaXjNomRF2TCWoiMgspKtXqya"), solana.GetTokenAccountsByDelegateConfig{
		ProgramID: &programID,
	})
	if err != nil {
		t.Fatal(err)
	}
	if len(accounts) == 0 {
		t.Fatal("Expected accounts")
	}
}

func TestGetTokenAccountsByOwner(t *testing.T) {
	t.Skip("Skipping test that requires network access")
	client := NewRpcClient(solana.RpcEndpointDevnet)
	programID := solana.MustParsePubkey("TokenkegQfeZyiNwAJbNbGKPFXCWuBvf9Ss623VQ5DA")
	accounts, err := client.GetTokenAccountsByOwner(solana.MustParsePubkey("GR16g49y2fEjRQD612ryaXjNomRF2TCWoiMgspKtXqya"), solana.GetTokenAccountsByDelegateConfig{
		ProgramID: &programID,
	})
	if err != nil {
		t.Fatal(err)
	}
	t.Fatal(accounts)
}

func TestGetTokenLargestAccounts(t *testing.T) {
	t.Skip("Skipping test that requires network access")
	client := NewRpcClient(solana.RpcEndpointDevnet)
	accounts, err := client.GetTokenLargestAccounts(solana.MustParsePubkey("BLrD8HqBy4vKNvkb28Bijg4y6s8tE49jyVFbfZnmesjY"))
	if err != nil {
		t.Fatal(err)
	}
	t.Fatal(accounts)
}

func TestGetTokenSupply(t *testing.T) {
	t.Skip("Skipping test that requires network access")
	client := NewRpcClient(solana.RpcEndpointDevnet)
	supply, err := client.GetTokenSupply(solana.MustParsePubkey("BLrD8HqBy4vKNvkb28Bijg4y6s8tE49jyVFbfZnmesjY"))
	if err != nil {
		t.Fatal(err)
	}
	t.Fatal(supply)
}
