package rpc

import (
	"solana"
	"solana/dependencies/keypair"
	"testing"
)

func TestGetInflationGovernor(t *testing.T) {
	t.Skip("Skipping test that requires network access")
	client := NewRpcClient(solana.RpcEndpointMainnetBeta)
	governor, err := client.GetInflationGovernor()
	if err != nil {
		t.Fatal(err)
	}
	t.Fatal(governor)
}

func TestGetInflationRate(t *testing.T) {
	t.Skip("Skipping test that requires network access")
	client := NewRpcClient(solana.RpcEndpointMainnetBeta)
	rate, err := client.GetInflationRate()
	if err != nil {
		t.Fatal(err)
	}
	t.Fatal(rate)
}

func TestGetInflationReward(t *testing.T) {
	t.Skip("Skipping test that requires network access")
	client := NewRpcClient(solana.RpcEndpointMainnetBeta)
	reward, err := client.GetInflationReward([]solana.Pubkey{keypair.MustParsePubkey("9jxgosAfHgHzwnxsHw4RAZYaLVokMbnYtmiZBreynGFP")})
	if err != nil {
		t.Fatal(err)
	}
	t.Fatal(reward)
}

func TestGetStakeMinimumDelegation(t *testing.T) {
	t.Skip("Skipping test that requires network access")
	client := NewRpcClient(solana.RpcEndpointMainnetBeta)
	minimum, err := client.GetStakeMinimumDelegation()
	if err != nil {
		t.Fatal(err)
	}
	t.Fatal(minimum)
}

func TestGetSupply(t *testing.T) {
	t.Skip("Skipping test that requires network access")
	client := NewRpcClient(solana.RpcEndpointMainnetBeta)
	supply, err := client.GetSupply()
	if err != nil {
		t.Fatal(err)
	}
	t.Fatal(supply)
}
