package rpc

import (
	"encoding/base64"
	"errors"
	"fmt"
	"solana"
	"solana/dependencies/keypair"
	"solana/dependencies/transactions"
	"time"
)

type block struct {
	BlockHeight       *int                       `json:"blockHeight"`       //The number of blocks beneath this block
	BlockTime         *int                       `json:"blockTime"`         //Estimated production time, as Unix timestamp (seconds since the Unix epoch). null if not available
	Blockhash         string                     `json:"blockhash"`         //The blockhash of this block, as base-58 encoded string
	ParentSlot        uint                       `json:"parentSlot"`        //The slot index of this block's parent
	PreviousBlockhash string                     `json:"previousBlockhash"` //The blockhash of this block's parent, as base-58 encoded string; if the parent block is not available due to ledger cleanup, this field will return "11111111111111111111111111111111"
	Transactions      []encodedTransaction       `json:"transactions"`      //The list of transactions included in this block
	Signatures        []string                   `json:"signatures"`        //Present if "signatures" are requested for transaction details; an array of signatures strings, corresponding to the transaction order in the block
	Rewards           []solana.TransactionReward `json:"rewards"`           //Block-level rewards, present if rewards are requested; an array of JSON objects containing:
}

func (r *RpcClient) GetBlock(slotNumber uint, config ...solana.GetBlockConfig) (*solana.Block, error) {
	var res *block

	// Set the encoding to base64 no matter what
	encoding := solana.EncodingBase64
	params := []interface{}{slotNumber}
	if len(config) > 0 {
		config[0].Encoding = &encoding
		params = append(params, config[0])
	} else {
		params = append(params, solana.GetAccountInfoConfig{Encoding: &encoding})
	}
	if err := r.send("getBlock", params, &res); err != nil {
		return nil, err
	}
	if res == nil {
		return nil, nil
	}

	var txs []solana.TransactionWithMeta
	for _, encodedTransaction := range res.Transactions {
		if len(encodedTransaction.Transaction) == 0 {
			return nil, errors.New("transaction not found")
		}
		transactionData, err := base64.StdEncoding.DecodeString(encodedTransaction.Transaction[0])
		if err != nil {
			return nil, fmt.Errorf("failed to decode transaction data: %v", err)
		}
		transaction, err := transactions.ParseTransactionData(transactionData)
		if err != nil {
			return nil, fmt.Errorf("failed to parse transaction data: %v", err)
		}

		txs = append(txs, solana.TransactionWithMeta{
			Meta:        encodedTransaction.Meta,
			Version:     encodedTransaction.Version,
			Transaction: transaction,
		})
	}

	return &solana.Block{
		BlockHeight:       res.BlockHeight,
		BlockTime:         res.BlockTime,
		Blockhash:         res.Blockhash,
		ParentSlot:        res.ParentSlot,
		PreviousBlockhash: res.PreviousBlockhash,
		Transactions:      txs,
		Signatures:        res.Signatures,
		Rewards:           res.Rewards,
	}, nil
}

func (r *RpcClient) GetBlockCommitment(slotNumber uint) (solana.BlockCommitment, error) {
	var res solana.BlockCommitment
	params := []interface{}{slotNumber}
	if err := r.send("getBlockCommitment", params, &res); err != nil {
		return solana.BlockCommitment{}, err
	}

	return res, nil
}

func (r *RpcClient) GetBlockHeight(config ...solana.StandardRpcConfig) (uint, error) {
	var res uint
	if err := r.send("getBlockHeight", nil, &res); err != nil {
		return 0, err
	}

	return res, nil
}

func (r *RpcClient) GetBlockProduction(config ...solana.GetBlockProductionConfig) (solana.BlockProduction, error) {
	var res struct {
		Value solana.BlockProduction `json:"value"`
	}
	params := []interface{}{}
	if len(config) > 0 {
		params = append(params, config[0])
	}
	if err := r.send("getBlockProduction", params, &res); err != nil {
		return solana.BlockProduction{}, err
	}
	return res.Value, nil
}

func (r *RpcClient) GetBlockTime(slotNumber uint) (time.Time, error) {
	var res int
	params := []interface{}{slotNumber}
	if err := r.send("getBlockTime", params, &res); err != nil {
		return time.Time{}, err
	}

	return time.Unix(int64(res), 0), nil
}

func (r *RpcClient) GetBlocks(startSlot uint, endSlot *uint, config ...solana.GetBlockConfig) ([]uint, error) {
	var res []uint
	params := []interface{}{startSlot, endSlot}
	if len(config) > 0 {
		params = append(params, config[0])
	}
	if err := r.send("getBlocks", params, &res); err != nil {
		return nil, err
	}
	return res, nil
}

func (r *RpcClient) GetBlocksWithLimit(startSlot uint, limit uint, config ...solana.GetBlockConfig) ([]uint, error) {
	var res []uint
	params := []interface{}{startSlot, limit}
	if len(config) > 0 {
		params = append(params, config[0])
	}
	if err := r.send("getBlocksWithLimit", params, &res); err != nil {
		return nil, err
	}

	return res, nil
}

func (r *RpcClient) GetClusterNodes() ([]solana.ClusterNode, error) {
	var res []solana.ClusterNode
	if err := r.send("getClusterNodes", nil, &res); err != nil {
		return nil, err
	}
	return res, nil
}

func (r *RpcClient) GetEpochInfo(config ...solana.StandardRpcConfig) (solana.EpochInfo, error) {
	var res solana.EpochInfo
	params := []interface{}{}
	if len(config) > 0 {
		params = append(params, config[0])
	}
	if err := r.send("getEpochInfo", params, &res); err != nil {
		return solana.EpochInfo{}, err
	}

	return res, nil
}

func (r *RpcClient) GetEpochSchedule() (solana.EpochSchedule, error) {
	var res solana.EpochSchedule
	if err := r.send("getEpochSchedule", nil, &res); err != nil {
		return solana.EpochSchedule{}, err
	}

	return res, nil
}

func (r *RpcClient) GetFirstAvailableBlock() (uint, error) {
	var res uint
	if err := r.send("getFirstAvailableBlock", nil, &res); err != nil {
		return 0, err
	}
	return res, nil
}

func (r *RpcClient) GetGenesisHash() (string, error) {
	var res string
	if err := r.send("getGenesisHash", nil, &res); err != nil {
		return "", err
	}
	return res, nil
}

func (r *RpcClient) GetHighestSnapshotSlot() (solana.HighestSnapshotSlot, error) {
	var res solana.HighestSnapshotSlot
	if err := r.send("getHighestSnapshotSlot", nil, &res); err != nil {
		return solana.HighestSnapshotSlot{}, err
	}

	return res, nil
}

func (r *RpcClient) GetLeaderSchedule(slot *uint, config ...solana.GetLeaderScheduleConfig) (*solana.LeaderSchedule, error) {
	var res *solana.LeaderSchedule
	params := []interface{}{slot}
	if len(config) > 0 {
		params = append(params, config[0])
	}
	if err := r.send("getLeaderSchedule", params, &res); err != nil {
		return nil, err
	}

	return res, nil
}

func (r *RpcClient) GetMaxRetransmitSlot() (uint, error) {
	var res uint
	if err := r.send("getMaxRetransmitSlot", nil, &res); err != nil {
		return 0, err
	}
	return res, nil
}

func (r *RpcClient) GetMaxShredInsertSlot() (uint, error) {
	var res uint
	if err := r.send("getMaxShredInsertSlot", nil, &res); err != nil {
		return 0, err
	}
	return res, nil
}

func (r *RpcClient) GetRecentPerformanceSamples(limit uint) ([]solana.PerformanceSample, error) {
	var res []solana.PerformanceSample
	params := []interface{}{limit}
	if err := r.send("getRecentPerformanceSamples", params, &res); err != nil {
		return nil, err
	}
	return res, nil
}

func (r *RpcClient) GetRecentPrioritizationFees(addresses []solana.Pubkey) ([]solana.PrioritizationFee, error) {
	var res []solana.PrioritizationFee
	params := []interface{}{}
	if len(addresses) > 0 {
		strs := make([]string, len(addresses))
		for i, address := range addresses {
			strs[i] = address.String()
		}
		params = append(params, strs)
	}
	if err := r.send("getRecentPrioritizationFees", params, &res); err != nil {
		return nil, err
	}
	return res, nil
}

func (r *RpcClient) GetSlot(config ...solana.StandardRpcConfig) (uint, error) {
	var res uint
	params := []interface{}{}
	if len(config) > 0 {
		params = append(params, config[0])
	}
	if err := r.send("getSlot", params, &res); err != nil {
		return 0, err
	}
	return res, nil
}

func (r *RpcClient) GetSlotLeader(config ...solana.StandardRpcConfig) (solana.Pubkey, error) {
	var res keypair.Pubkey
	params := []interface{}{}
	if len(config) > 0 {
		params = append(params, config[0])
	}
	if err := r.send("getSlotLeader", params, &res); err != nil {
		return nil, err
	}
	return &res, nil
}

func (r *RpcClient) GetSlotLeaders(start, limit *uint) ([]solana.Pubkey, error) {
	var res []keypair.Pubkey
	params := []interface{}{start, limit}
	if err := r.send("getSlotLeaders", params, &res); err != nil {
		return nil, err
	}
	var pubkeys []solana.Pubkey
	for _, pubkey := range res {
		pubkeys = append(pubkeys, &pubkey)
	}
	return pubkeys, nil
}

func (r *RpcClient) GetTransactionCount(config ...solana.StandardRpcConfig) (uint, error) {
	var res uint
	params := []interface{}{}
	if len(config) > 0 {
		params = append(params, config[0])
	}
	if err := r.send("getTransactionCount", params, &res); err != nil {
		return 0, err
	}
	return res, nil
}

func (r *RpcClient) GetVoteAccounts(config ...solana.GetVoteAccountsConfig) (solana.VoteAccounts, error) {
	var res solana.VoteAccounts
	params := []interface{}{}
	if len(config) > 0 {
		params = append(params, config[0])
	}
	if err := r.send("getVoteAccounts", params, &res); err != nil {
		return solana.VoteAccounts{}, err
	}
	return res, nil
}

func (r *RpcClient) MinimumLedgerSlot() (uint, error) {
	var res uint
	if err := r.send("minimumLedgerSlot", nil, &res); err != nil {
		return 0, err
	}
	return res, nil
}
