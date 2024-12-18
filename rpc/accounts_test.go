package rpc

import (
	"testing"

	"github.com/hwsimmons17/solana-web3.go"
)

func TestGetAccountInfo(t *testing.T) {
	t.Skip("Skipping test that requires network access")
	client := NewRpcClient(solana.RpcEndpointDevnet)
	pubkey := solana.MustParsePubkey("BLrD8HqBy4vKNvkb28Bijg4y6s8tE49jyVFbfZnmesjY")
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
	t.Fatal(account)
}

func TestGetBalance(t *testing.T) {
	t.Skip("Skipping test that requires network access")
	client := NewRpcClient(solana.RpcEndpointDevnet)
	pubkey := solana.MustParsePubkey("BLrD8HqBy4vKNvkb28Bijg4y6s8tE49jyVFbfZnmesjY")
	comit := solana.CommitmentConfirmed
	balance, err := client.GetBalance(pubkey, solana.StandardRpcConfig{
		Commitment: &comit,
	})
	if err != nil {
		t.Fatal(err)
	}
	t.Fatal(balance)
}

func TestGetLargestAccounts(t *testing.T) {
	t.Skip("Skipping test that requires network access")
	client := NewRpcClient(solana.RpcEndpointDevnet)
	accounts, err := client.GetLargestAccounts()
	if err != nil {
		t.Fatal(err)
	}
	if len(accounts) == 0 {
		t.Fatal("Expected accounts")
	}
}

func TestGetMiminumBalanceForRentExemption(t *testing.T) {
	t.Skip("Skipping test that requires network access")
	client := NewRpcClient(solana.RpcEndpointDevnet)
	minimum, err := client.GetMinimumBalanceForRentExemption(8200000)
	if err != nil {
		t.Fatal(err)
	}
	t.Fatal(minimum)
}

func TestGetMultipleAccounts(t *testing.T) {
	t.Skip("Skipping test that requires network access")
	client := NewRpcClient(solana.RpcEndpointDevnet)
	pubkey1 := solana.MustParsePubkey("BLrD8HqBy4vKNvkb28Bijg4y6s8tE49jyVFbfZnmesjY")
	pubkey2 := solana.MustParsePubkey("5oNDL3swdJJF1g9DzJiZ4ynHXgszjAEpUkxVYejchzrY")
	pubkeys := []solana.Pubkey{pubkey1, pubkey2}
	comit := solana.CommitmentConfirmed
	accounts, err := client.GetMultipleAccounts(pubkeys, solana.GetAccountInfoConfig{
		Commitment: &comit,
	})
	if err != nil {
		t.Fatal(err)
	}
	t.Fatal(*accounts[0])
}

func TestGetProgramAccounts(t *testing.T) {
	t.Skip("Skipping test that requires network access")
	client := NewRpcClient(solana.RpcEndpointDevnet)
	programId := solana.MustParsePubkey("NativeLoader1111111111111111111111111111111")
	accounts, err := client.GetProgramAccounts(programId)
	if err != nil {
		t.Fatal(err)
	}
	if len(accounts) == 0 {
		t.Fatal("Expected accounts")
	}
	t.Fatal(accounts)
}
