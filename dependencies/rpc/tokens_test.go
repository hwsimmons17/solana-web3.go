package rpc

import (
	"solana"
	"solana/dependencies/keypair"
	"testing"
)

func TestGetTokenAccountBalance(t *testing.T) {
	t.Skip("Skipping test that requires network access")
	client := NewRpcClient(solana.RpcEndpointDevnet)
	balance, err := client.GetTokenAccountBalance(keypair.MustParsePubkey("7Fg8XQBVY4z7gPzecGo7abbHZbHj3iFfGozXsz1VcvKk"))
	if err != nil {
		t.Fatal(err)
	}
	t.Fatal(balance)
}

func TestGetTokenAccountsByDelegate(t *testing.T) {
	t.Skip("Skipping test that requires network access")
	client := NewRpcClient(solana.RpcEndpointDevnet)
	programID := keypair.MustParsePubkey("TokenkegQfeZyiNwAJbNbGKPFXCWuBvf9Ss623VQ5DA")
	accounts, err := client.GetTokenAccountsByDelegate(keypair.MustParsePubkey("GR16g49y2fEjRQD612ryaXjNomRF2TCWoiMgspKtXqya"), solana.GetTokenAccountsByDelegateConfig{
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
	programID := keypair.MustParsePubkey("TokenkegQfeZyiNwAJbNbGKPFXCWuBvf9Ss623VQ5DA")
	accounts, err := client.GetTokenAccountsByOwner(keypair.MustParsePubkey("GR16g49y2fEjRQD612ryaXjNomRF2TCWoiMgspKtXqya"), solana.GetTokenAccountsByDelegateConfig{
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
	accounts, err := client.GetTokenLargestAccounts(keypair.MustParsePubkey("BLrD8HqBy4vKNvkb28Bijg4y6s8tE49jyVFbfZnmesjY"))
	if err != nil {
		t.Fatal(err)
	}
	t.Fatal(accounts)
}

func TestGetTokenSupply(t *testing.T) {
	t.Skip("Skipping test that requires network access")
	client := NewRpcClient(solana.RpcEndpointDevnet)
	supply, err := client.GetTokenSupply(keypair.MustParsePubkey("BLrD8HqBy4vKNvkb28Bijg4y6s8tE49jyVFbfZnmesjY"))
	if err != nil {
		t.Fatal(err)
	}
	t.Fatal(supply)
}
