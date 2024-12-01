package rpc

import (
	"encoding/base64"
	"errors"
	"fmt"
	"solana"
	"solana/dependencies/transactions"
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
