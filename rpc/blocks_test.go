package rpc

import (
	"testing"

	"github.com/hwsimmons17/solana-web3.go"
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

func TestGetBlockCommitment(t *testing.T) {
	t.Skip("Skipping test that requires network access")
	client := NewRpcClient(solana.RpcEndpointMainnetBeta)
	commitment, err := client.GetBlockCommitment(283079039)
	if err != nil {
		t.Fatal(err)
	}
	t.Fatal(commitment)
}

func TestGetBlockHeight(t *testing.T) {
	t.Skip("Skipping test that requires network access")
	client := NewRpcClient(solana.RpcEndpointMainnetBeta)
	height, err := client.GetBlockHeight()
	if err != nil {
		t.Fatal(err)
	}
	t.Fatal(height)
}

func TestGetBlockProduction(t *testing.T) {
	t.Skip("Skipping test that requires network access")
	client := NewRpcClient(solana.RpcEndpointMainnetBeta)
	production, err := client.GetBlockProduction()
	if err != nil {
		t.Fatal(err)
	}
	t.Fatal(production)
}

func TestGetBlockTime(t *testing.T) {
	t.Skip("Skipping test that requires network access")
	client := NewRpcClient(solana.RpcEndpointMainnetBeta)
	time, err := client.GetBlockTime(283079039)
	if err != nil {
		t.Fatal(err)
	}
	t.Fatal(time)
}

func TestGetBlocks(t *testing.T) {
	t.Skip("Skipping test that requires network access")
	client := NewRpcClient(solana.RpcEndpointMainnetBeta)
	end := uint(283079050)
	blocks, err := client.GetBlocks(283079039, &end)
	if err != nil {
		t.Fatal(err)
	}
	t.Fatal(blocks)
}

func TestGetBlocksWithLimit(t *testing.T) {
	t.Skip("Skipping test that requires network access")
	client := NewRpcClient(solana.RpcEndpointMainnetBeta)
	blocks, err := client.GetBlocksWithLimit(283079039, 5)
	if err != nil {
		t.Fatal(err)
	}
	t.Fatal(blocks)
}

func TestGetClusterNodes(t *testing.T) {
	t.Skip("Skipping test that requires network access")
	client := NewRpcClient(solana.RpcEndpointDevnet)
	nodes, err := client.GetClusterNodes()
	if err != nil {
		t.Fatal(err)
	}
	t.Fatal(nodes)
}

func TestGetEpochInfo(t *testing.T) {
	t.Skip("Skipping test that requires network access")
	client := NewRpcClient(solana.RpcEndpointMainnetBeta)
	info, err := client.GetEpochInfo()
	if err != nil {
		t.Fatal(err)
	}
	t.Fatal(info)
}

func TestGetEpochSchedule(t *testing.T) {
	t.Skip("Skipping test that requires network access")
	client := NewRpcClient(solana.RpcEndpointMainnetBeta)
	schedule, err := client.GetEpochSchedule()
	if err != nil {
		t.Fatal(err)
	}
	t.Fatal(schedule)
}

func TestGetFirstAvailableBlock(t *testing.T) {
	t.Skip("Skipping test that requires network access")
	client := NewRpcClient(solana.RpcEndpointMainnetBeta)
	block, err := client.GetFirstAvailableBlock()
	if err != nil {
		t.Fatal(err)
	}
	t.Fatal(block)
}

func TestGetGenesisHash(t *testing.T) {
	t.Skip("Skipping test that requires network access")
	client := NewRpcClient(solana.RpcEndpointMainnetBeta)
	hash, err := client.GetGenesisHash()
	if err != nil {
		t.Fatal(err)
	}
	t.Fatal(hash)
}

func TestGetHighestSnapshotSlot(t *testing.T) {
	t.Skip("Skipping test that requires network access")
	client := NewRpcClient(solana.RpcEndpointMainnetBeta)
	slot, err := client.GetHighestSnapshotSlot()
	if err != nil {
		t.Fatal(err)
	}
	t.Fatal(slot)
}

func TestGetLeaderSchedule(t *testing.T) {
	t.Skip("Skipping test that requires network access")
	client := NewRpcClient(solana.RpcEndpointMainnetBeta)
	schedule, err := client.GetLeaderSchedule(nil)
	if err != nil {
		t.Fatal(err)
	}
	t.Fatal(schedule)
}

func TestGetMaxRetransmitSlots(t *testing.T) {
	t.Skip("Skipping test that requires network access")
	client := NewRpcClient(solana.RpcEndpointMainnetBeta)
	slots, err := client.GetMaxRetransmitSlot()
	if err != nil {
		t.Fatal(err)
	}
	t.Fatal(slots)
}

func TestGetMaxShredInsertSlot(t *testing.T) {
	t.Skip("Skipping test that requires network access")
	client := NewRpcClient(solana.RpcEndpointMainnetBeta)
	slot, err := client.GetMaxShredInsertSlot()
	if err != nil {
		t.Fatal(err)
	}
	t.Fatal(slot)
}

func TestGetRecentPerformanceSamples(t *testing.T) {
	t.Skip("Skipping test that requires network access")
	client := NewRpcClient(solana.RpcEndpointMainnetBeta)
	samples, err := client.GetRecentPerformanceSamples(3)
	if err != nil {
		t.Fatal(err)
	}
	t.Fatal(samples)
}

func TestGetRecentPrioritizationFees(t *testing.T) {
	t.Skip("Skipping test that requires network access")
	client := NewRpcClient(solana.RpcEndpointMainnetBeta)
	fees, err := client.GetRecentPrioritizationFees(nil)
	if err != nil {
		t.Fatal(err)
	}
	t.Fatal(fees)
}

func TestGetSlot(t *testing.T) {
	t.Skip("Skipping test that requires network access")
	client := NewRpcClient(solana.RpcEndpointMainnetBeta)
	slot, err := client.GetSlot()
	if err != nil {
		t.Fatal(err)
	}
	t.Fatal(slot)
}

func TestGetSlotLeader(t *testing.T) {
	t.Skip("Skipping test that requires network access")
	client := NewRpcClient(solana.RpcEndpointMainnetBeta)
	leader, err := client.GetSlotLeader()
	if err != nil {
		t.Fatal(err)
	}
	t.Fatal(leader)
}

func TestGetSlotLeaders(t *testing.T) {
	t.Skip("Skipping test that requires network access")
	client := NewRpcClient(solana.RpcEndpointMainnetBeta)
	start := uint(304661650)
	limit := uint(100)
	leaders, err := client.GetSlotLeaders(&start, &limit)
	if err != nil {
		t.Fatal(err)
	}
	t.Fatal(leaders)
}

func TestGetTransactionCount(t *testing.T) {
	t.Skip("Skipping test that requires network access")
	client := NewRpcClient(solana.RpcEndpointMainnetBeta)
	count, err := client.GetTransactionCount()
	if err != nil {
		t.Fatal(err)
	}
	t.Fatal(count)
}

func TestGetVoteAccounts(t *testing.T) {
	t.Skip("Skipping test that requires network access")
	client := NewRpcClient(solana.RpcEndpointMainnetBeta)
	accounts, err := client.GetVoteAccounts()
	if err != nil {
		t.Fatal(err)
	}
	t.Fatal(accounts)
}

func TestMinimumLedgerSlot(t *testing.T) {
	t.Skip("Skipping test that requires network access")
	client := NewRpcClient(solana.RpcEndpointMainnetBeta)
	slot, err := client.MinimumLedgerSlot()
	if err != nil {
		t.Fatal(err)
	}
	t.Fatal(slot)
}
