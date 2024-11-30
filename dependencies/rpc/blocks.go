package rpc

import "solana"

func (r *RpcClient) GetBlock(slotNumber uint, config ...solana.GetBlockConfig) (*solana.Block, error) {
	//Set the encoding to base64 no matter what
	// encoding := solana.EncodingBase64
	// params := []interface{}{slotNumber}
	// if len(config) > 0 {
	// 	config[0].Encoding = &encoding
	// 	params = append(params, config[0])
	// } else {
	// 	params = append(params, solana.GetAccountInfoConfig{Encoding: &encoding})
	// }
	// res, err := r.send("getBlock", params)
	// if err != nil {
	// 	return nil, err
	// }

	return nil, nil
}

func (r *RpcClient) GetBlockCommitment(slotNumber uint) (solana.BlockCommitment, error) {
	return solana.BlockCommitment{}, nil
}

func (r *RpcClient) GetBlockHeight(config ...solana.StandardRpcConfig) (uint, error) {
	return 0, nil
}

func (r *RpcClient) GetBlockProduction(config ...solana.GetBlockProductionConfig) (solana.BlockProduction, error) {
	return solana.BlockProduction{}, nil
}

func (r *RpcClient) GetBlockTime(slotNumber uint) (int, error) {
	return 0, nil
}

func (r *RpcClient) GetBlocks(startSlot uint, endSlot *uint, config ...solana.GetBlockConfig) ([]uint, error) {
	return nil, nil
}

func (r *RpcClient) GetBlocksWithLimit(startSlot uint, limit uint, config ...solana.GetBlockConfig) ([]uint, error) {
	return nil, nil
}

func (r *RpcClient) GetClusterNodes() ([]solana.ClusterNode, error) {
	return nil, nil
}

func (r *RpcClient) GetEpochInfo(...solana.StandardRpcConfig) (solana.EpochInfo, error) {
	return solana.EpochInfo{}, nil
}

func (r *RpcClient) GetEpochSchedule() (solana.EpochSchedule, error) {
	return solana.EpochSchedule{}, nil
}

func (r *RpcClient) GetFirstAvailableBlock() (uint, error) {
	return 0, nil
}

func (r *RpcClient) GetGenesisHash() (string, error) {
	return "", nil
}

func (r *RpcClient) GetHighestSnapshotSlot() (solana.HighestSnapshotSlot, error) {
	return solana.HighestSnapshotSlot{}, nil
}

func (r *RpcClient) GetLeaderSchedule(slot *uint, config ...solana.GetLeaderScheduleConfig) (*solana.LeaderSchedule, error) {
	return nil, nil
}

func (r *RpcClient) GetMaxRetransmitSlots() (uint, error) {
	return 0, nil
}

func (r *RpcClient) GetMaxShredInsertSlot() (uint, error) {
	return 0, nil
}

func (r *RpcClient) GetRecentPerformanceSamples(limit uint) ([]solana.PerformanceSample, error) {
	return nil, nil
}

func (r *RpcClient) GetRecentPrioritizationFees(addresses []string) (solana.PrioritizationFee, error) {
	return solana.PrioritizationFee{}, nil
}

func (r *RpcClient) GetSlot(config ...solana.StandardRpcConfig) (uint, error) {
	return 0, nil
}

func (r *RpcClient) GetSlotLeader(config ...solana.StandardRpcConfig) (string, error) {
	return "", nil
}

func (r *RpcClient) GetSlotLeaders(start, end uint) ([]string, error) {
	return nil, nil
}

func (r *RpcClient) GetTransactionCount(config ...solana.StandardRpcConfig) (uint, error) {
	return 0, nil
}

func (r *RpcClient) GetVoteAccounts(config ...solana.GetVoteAccountsConfig) (solana.VoteAccounts, error) {
	return solana.VoteAccounts{}, nil
}

func (r *RpcClient) MinimumLedgerSlot() (uint, error) {
	return 0, nil
}
