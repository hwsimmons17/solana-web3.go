package solana

type Rpc interface {
	GetAccountInfo(address Address, config ...GetAccountInfoConfig) (EncodedAccount, error)                                                              //Returns all information associated with the account of provided Pubkey
	GetBalance(address Address, config ...StandardRpcConfig) (uint, error)                                                                               //Returns the lamport balance of the account of provided Pubkey
	GetBlock(slotNumber uint, config ...GetBlockConfig) (*Block, error)                                                                                  //Returns identity and transaction information about a confirmed block in the ledger
	GetBlockCommitment(slotNumber uint) (BlockCommitment, error)                                                                                         //Returns the current block height and the estimated production time of a block
	GetBlockHeight(config ...StandardRpcConfig) (uint, error)                                                                                            //Returns the current block height of the node
	GetBlockProduction(config ...GetBlockProductionConfig) (BlockProduction, error)                                                                      //Returns recent block production information from the current or previous epoch.
	GetBlockTime(slotNumber uint) (int, error)                                                                                                           //Returns the estimated production time of a block as a unix timestamp.
	GetBlocks(startSlot uint, endSlot *uint, config ...GetBlockConfig) ([]uint, error)                                                                   //Returns a list of confirmed blocks between two slots
	GetBlocksWithLimit(startSlot uint, limit uint, config ...GetBlockConfig) ([]uint, error)                                                             //Returns a list of confirmed blocks starting at the given slot
	GetClusterNodes() ([]ClusterNode, error)                                                                                                             //Returns information about all the nodes participating in the cluster
	GetEpochInfo(...StandardRpcConfig) (EpochInfo, error)                                                                                                //Returns information about the current epoch
	GetEpochSchedule() (EpochSchedule, error)                                                                                                            //Returns the epoch schedule information from this cluster's genesis config
	GetFeeForMessage(msg string, config ...StandardRpcConfig) (*uint, error)                                                                             //Get the fee the network will charge for a particular Message. NOTE: This method is only available in solana-core v1.9 or newer. Please use getFees for solana-core v1.8 and below.
	GetFirstAvailableBlock() (uint, error)                                                                                                               //Returns the slot of the lowest confirmed block that has not been purged from the ledger
	GetGenesisHash() (string, error)                                                                                                                     //Returns the genesis hash
	GetHealth() error                                                                                                                                    //Returns the current health of the node. A healthy node is one that is within HEALTH_CHECK_SLOT_DISTANCE slots of the latest cluster confirmed slot.
	GetHighestSnapshotSlot() (HighestSnapshotSlot, error)                                                                                                //Returns the highest slot information that the node has snapshots for. This will find the highest full snapshot slot, and the highest incremental snapshot slot based on the full snapshot slot, if there is one.
	GetIdentity() (string, error)                                                                                                                        //Returns the identity pubkey for the current node
	GetInflationGovernor(config ...StandardCommitmentConfig) (InflationGovernor, error)                                                                  //Returns the current inflation governor
	GetInflationRate(config ...StandardCommitmentConfig) (InflationRate, error)                                                                          //Returns the specific inflation values for the current epoch
	GetInflationReward(addresses []string, config ...GetInflationRewardConfig) (InflationReward, error)                                                  //Returns the inflation / staking reward for a list of addresses for an epoch
	GetLargestAccounts(config ...GetLargestAccountsConfig) ([]AccountWithBalance, error)                                                                 //Returns the 20 largest accounts, by lamport balance (results may be cached up to two hours)
	GetLatestBlockhash(config ...StandardRpcConfig) (LatestBlockhash, error)                                                                             //Returns the latest blockhash. NOTE: This method is only available in solana-core v1.9 or newer. Please use getRecentBlockhash for solana-core v1.8 and below.
	GetLeaderSchedule(slot *uint, config ...GetLeaderScheduleConfig) (*LeaderSchedule, error)                                                            //Returns the leader schedule for an epoch
	GetMaxRetransmitSlots() (uint, error)                                                                                                                //Get the max slot seen from retransmit stage.
	GetMaxShredInsertSlot() (uint, error)                                                                                                                //Get the max slot seen from after shred insert.
	GetMinimumBalanceForRentExemption(accountDataLength uint, config ...StandardCommitmentConfig) (uint, error)                                          //Returns minimum balance required to make account rent exempt.
	GetMultipleAccounts(pubkeys []string, config ...GetAccountInfoConfig) ([]EncodedAccount, error)                                                      //Returns the account information for a list of Pubkeys.
	GetProgramAccounts(programPubkey string, config ...GetAccountInfoConfig) ([]EncodedAccount, error)                                                   //Returns all accounts owned by the provided program Pubkey
	GetRecentPerformanceSamples(limit uint) ([]PerformanceSample, error)                                                                                 //Returns a list of recent performance samples, in reverse slot order. Performance samples are taken every 60 seconds and include the number of transactions and slots that occur in a given time window. -- NOTE max limit is 720
	GetRecentPrioritizationFees(addresses []string) (PrioritizationFee, error)                                                                           //Returns a list of prioritization fees from recent blocks.
	GetSignatureStatuses(signatures []string, config ...GetSignatureStatusesConfig) ([]*SignatureStatus, error)                                          //Returns the statuses of a list of signatures. Each signature must be a txid, the first signature of a transaction. Unless the searchTransactionHistory configuration parameter is included, this method only searches the recent status cache of signatures, which retains statuses for all active slots plus MAX_RECENT_BLOCKHASHES rooted slots.
	GetSignaturesForAddress(account string, config ...GetSignaturesForAddressConfig) ([]TransactionSignature, error)                                     //Returns signatures for confirmed transactions that include the given address in their accountKeys list. Returns signatures backwards in time from the provided signature or most recent confirmed block
	GetSlot(config ...StandardRpcConfig) (uint, error)                                                                                                   //Returns the slot that has reached the given or default commitment level
	GetSlotLeader(config ...StandardRpcConfig) (string, error)                                                                                           //Returns the current slot leader
	GetSlotLeaders(start, end uint) ([]string, error)                                                                                                    //Returns the slot leaders for a given slot range
	GetStakeMinimumDelegation(config ...StandardCommitmentConfig) (uint, error)                                                                          //Returns the stake minimum delegation, in lamports.
	GetSupply(config ...GetSupplyConfig) (Supply, error)                                                                                                 //Returns information about the current supply.
	GetTokenAccountBalance(account string, config ...StandardCommitmentConfig) (UiTokenAmount, error)                                                    //Returns the token balance of an SPL Token account.
	GetTokenAccountsByDelegate(delegateAccount string, opts *GetTokenAccountsByDelegateConfig, config ...GetAccountInfoConfig) ([]EncodedAccount, error) //Returns all SPL Token accounts by approved Delegate.
	GetTokenAccountsByOwner(ownerAccount string, opts *GetTokenAccountsByDelegateConfig, config ...GetAccountInfoConfig) ([]EncodedAccount, error)       //Returns all SPL Token accounts by token owner.
	GetTokenLargestAccounts(mintAccount string, config ...StandardCommitmentConfig) ([]UiTokenAmount, error)                                             //Returns the 20 largest accounts of a particular SPL Token type.
	GetTokenSupply(mintAccount string, config ...StandardCommitmentConfig) (UiTokenAmount, error)                                                        //Returns the total supply of an SPL Token type.
	GetTransaction(transactionSignature string, config ...GetTransactionSignatureConfig) (*Transaction, error)                                           //Returns transaction details for a confirmed transaction
	GetTransactionCount(config ...StandardRpcConfig) (uint, error)                                                                                       //Returns the current transaction count from the ledger
	GetVersion() (Version, error)                                                                                                                        //Returns the current Solana version running on the node
	GetVoteAccounts(config ...GetVoteAccountsConfig) (VoteAccounts, error)                                                                               //Returns the account info and associated stake for all the voting accounts in the current bank.
	IsBlockhashValid(blockhash string, config ...StandardRpcConfig) (bool, error)                                                                        //Returns whether a blockhash is valid
	MinimumLedgerSlot() (uint, error)                                                                                                                    //Returns the lowest slot that the node has information about in its ledger.
	RequestAirdrop(destinationAddress string, lamports uint, config ...StandardCommitmentConfig) (string, error)                                         //Requests an airdrop of lamports to a Solana account
	/*
		Submits a signed transaction to the cluster for processing.

		This method does not alter the transaction in any way; it relays the transaction created by clients to the node as-is.

		If the node's rpc service receives the transaction, this method immediately succeeds, without waiting for any confirmations. A successful response from this method does not guarantee the transaction is processed or confirmed by the cluster.

		While the rpc service will reasonably retry to submit it, the transaction could be rejected if transaction's recent_blockhash expires before it lands.

		Use getSignatureStatuses to ensure a transaction is processed and confirmed.

		Before submitting, the following preflight checks are performed:

		The transaction signatures are verified
		The transaction is simulated against the bank slot specified by the preflight commitment. On failure an error will be returned. Preflight checks may be disabled if desired. It is recommended to specify the same commitment and preflight commitment to avoid confusing behavior.
		The returned signature is the first signature in the transaction, which is used to identify the transaction (transaction id). This identifier can be easily extracted from the transaction data before submission.
	*/
	SendTransaction(fullySignedTransaction string, config ...SendTransactionConfig) (string, error)
	SimulateTransaction(transaction string, config ...SimulateTransactionConfig) (SimulateTransactionResult, error) //Simulate sending a transaction. NOTE: Transaction needs a valid recent blockhash, but does not need to be signed
}

type StandardRpcConfig struct {
	Commitment     *Commitment `json:"commitment"`     //For preflight checks and transaction processing, Solana nodes choose which bank state to query based on a commitment requirement set by the client. The commitment describes how finalized a block is at that point in time. When querying the ledger state, it's recommended to use lower levels of commitment to report progress and higher levels to ensure the state will not be rolled back.
	MinContextSlot *int        `json:"minContextSlot"` //The minimum slot that the request can be evaluated at
}

type StandardCommitmentConfig struct {
	Commitment *Commitment `json:"commitment"` //For preflight checks and transaction processing, Solana nodes choose which bank state to query based on a commitment requirement set by the client. The commitment describes how finalized a block is at that point in time. When querying the ledger state, it's recommended to use lower levels of commitment to report progress and higher levels to ensure the state will not be rolled back.
}

type GetAccountInfoConfig struct {
	Commitment     *Commitment `json:"commitment"`     //For preflight checks and transaction processing, Solana nodes choose which bank state to query based on a commitment requirement set by the client. The commitment describes how finalized a block is at that point in time. When querying the ledger state, it's recommended to use lower levels of commitment to report progress and higher levels to ensure the state will not be rolled back.
	Encoding       *Encoding   `json:"encoding"`       //Encoding format for Account data
	DataSlice      *DataSlice  `json:"dataSlice"`      //Request a slice of the account's data.
	MinContextSlot *int        `json:"minContextSlot"` //The minimum slot that the request can be evaluated at
}

type GetBlockConfig struct {
	Commitment                     *Commitment         `json:"commitment"`                     //For preflight checks and transaction processing, Solana nodes choose which bank state to query based on a commitment requirement set by the client. The commitment describes how finalized a block is at that point in time. When querying the ledger state, it's recommended to use lower levels of commitment to report progress and higher levels to ensure the state will not be rolled back.
	Encoding                       *Encoding           `json:"encoding"`                       //Encoding format for each returned Transaction
	TransactionDetails             *TransactionDetails `json:"transactionDetails"`             //Level of transaction detail to return -- default is "full"
	MaxSupportedTransactionVersion int                 `json:"maxSupportedTransactionVersion"` //If this parameter is omitted, only legacy transactions will be returned, and a block containing any versioned transaction will prompt the error.
	Rewards                        bool                `json:"rewards"`                        //Whether to populate the rewards array. If parameter not provided, the default includes rewards.
}

type GetBlockProductionConfig struct {
	Commitment *Commitment `json:"commitment"` //For preflight checks and transaction processing, Solana nodes choose which bank state to query based on a commitment requirement set by the client. The commitment describes how finalized a block is at that point in time. When querying the ledger state, it's recommended to use lower levels of commitment to report progress and higher levels to ensure the state will not be rolled back.
	Identity   *string     `json:"identity"`   //Only return results for this validator identity (base-58 encoded)
	Range      *BlockRange `json:"range"`      //Slot range to return block production for. If parameter not provided, defaults to current epoch.
}

type GetBlocksConfig struct {
	Commitment *Commitment `json:"commitment"` //"processed" is not supported
}

type GetInflationRewardConfig struct {
	MinContextSlot *int  `json:"minContextSlot"` //The minimum slot that the request can be evaluated at
	Epoch          *uint `json:"epoch"`          //An epoch for which the reward occurs. If omitted, the previous epoch will be used
}

type GetLargestAccountsConfig struct {
	Commitment *Commitment `json:"commitment"` //For preflight checks and transaction processing, Solana nodes choose which bank state to query based on a commitment requirement set by the client. The commitment describes how finalized a block is at that point in time. When querying the ledger state, it's recommended to use lower levels of commitment to report progress and higher levels to ensure the state will not be rolled back.
	Filter     *string     `json:"filter"`     //Filter results by account type
}

type GetLeaderScheduleConfig struct {
	Commitment *Commitment `json:"commitment"` //For preflight checks and transaction processing, Solana nodes choose which bank state to query based on a commitment requirement set by the client. The commitment describes how finalized a block is at that point in time. When querying the ledger state, it's recommended to use lower levels of commitment to report progress and higher levels to ensure the state will not be rolled back.
	Identity   *string     `json:"identity"`   //Only return results for this validator identity (base-58 encoded)
}

type GetSignatureStatusesConfig struct {
	SearchTransactionHistory bool `json:"searchTransactionHistory"` //If true - a Solana node will search its ledger cache for any signatures not found in the recent status cache
}

type GetSignaturesForAddressConfig struct {
	Commitment     *Commitment `json:"commitment"`     //For preflight checks and transaction processing, Solana nodes choose which bank state to query based on a commitment requirement set by the client. The commitment describes how finalized a block is at that point in time. When querying the ledger state, it's recommended to use lower levels of commitment to report progress and higher levels to ensure the state will not be rolled back.
	MinContextSlot *int        `json:"minContextSlot"` //The minimum slot that the request can be evaluated at
	Limit          *uint       `json:"limit"`          //Maximum transaction signatures to return (between 1 and 1,000). Default is 1000.
	Before         *string     `json:"before"`         //Start searching backwards from this transaction signature. If not provided the search starts from the top of the highest max confirmed block.
	After          *string     `json:"after"`          //Search until this transaction signature, if found before limit reached
}

type GetSupplyConfig struct {
	Commitment                        *Commitment `json:"commitment"`                        //For preflight checks and transaction processing, Solana nodes choose which bank state to query based on a commitment requirement set by the client. The commitment describes how finalized a block is at that point in time. When querying the ledger state, it's recommended to use lower levels of commitment to report progress and higher levels to ensure the state will not be rolled back.
	ExcludeNonCirculatingAccountsList *bool       `json:"excludeNonCirculatingAccountsList"` //Exclude non circulating accounts list from response
}

// Supply one of the following fields
type GetTokenAccountsByDelegateConfig struct {
	Mint      *string `json:"mint"`      //Pubkey of the specific token Mint to limit accounts to, as base-58 encoded string; or
	ProgramID *string `json:"programId"` //Pubkey of the Token program that owns the accounts, as base-58 encoded string
}

type GetTransactionSignatureConfig struct {
	Commitment                     *Commitment `json:"commitment"`                     //For preflight checks and transaction processing, Solana nodes choose which bank state to query based on a commitment requirement set by the client. The commitment describes how finalized a block is at that point in time. When querying the ledger state, it's recommended to use lower levels of commitment to report progress and higher levels to ensure the state will not be rolled back.
	Encoding                       *Encoding   `json:"encoding"`                       //Encoding format for each returned Transaction
	MaxSupportedTransactionVersion int         `json:"maxSupportedTransactionVersion"` //Set the max transaction version to return in responses. If the requested transaction is a higher version, an error will be returned. If this parameter is omitted, only legacy transactions will be returned, and any versioned transaction will prompt the error.
}

type GetVoteAccountsConfig struct {
	Commitment *Commitment `json:"commitment"` //For preflight checks and transaction processing, Solana nodes choose which bank state to query based on a commitment requirement set by the client. The commitment describes how finalized a block is at that point in time. When querying the ledger state, it's recommended to use lower levels of commitment to report progress and higher levels to ensure the state will not be rolled back.
	VotePubkey *string     `json:"votePubkey"` //Only return results for this validator vote address (base-58 encoded)

}

type SendTransactionConfig struct {
	Encoding            *Encoding   `json:"encoding"`            //Encoding used for the transaction data. Values: base58 (slow, DEPRECATED), or base64.
	SkipPreflight       *bool       `json:"skipPreflight"`       //When true, skip the preflight transaction checks
	PreflightCommitment *Commitment `json:"preflightCommitment"` //Default: finalized. Commitment level to use for preflight.
	MaxRetries          *uint       `json:"maxRetries"`          //Maximum number of times for the RPC node to retry sending the transaction to the leader. If this parameter not provided, the RPC node will retry the transaction until it is finalized or until the blockhash expires.
	MinContextSlot      *uint       `json:"minContextSlot"`      //The minimum slot that the request can be evaluated at
}

type SimulateTransactionConfig struct {
	Commitment             *Commitment `json:"commitment"`             //Commitment level to simulate the transaction at
	SigVerify              *bool       `json:"sigVerify"`              //If true the transaction signatures will be verified (conflicts with replaceRecentBlockhash)
	ReplaceRecentBlockhash *bool       `json:"replaceRecentBlockhash"` //If true the transaction recent blockhash will be replaced with the most recent blockhash. (conflicts with sigVerify)
	MinContextSlot         *uint       `json:"minContextSlot"`         //The minimum slot that the request can be evaluated at
	Encoding               *Encoding   `json:"encoding"`               //Encoding used for the transaction data. Values: base58 (slow, DEPRECATED), or base64.
	InnerInstructions      *bool       `json:"innerInstructions"`      //If true the response will include inner instructions. These inner instructions will be jsonParsed where possible, otherwise json.
	Accounts               *struct {
		Addresses []string `json:"addresses"` //An array of accounts to return, as base-58 encoded strings
		Encoding  Encoding `json:"encoding"`  //Encoding for returned Account data
	} `json:"accounts"` //Accounts configuration object
}

type SimulateTransactionResult struct {
	Err           any              `json:"err"`           //Error if transaction failed, null if transaction succeeded.
	Logs          []string         `json:"logs"`          //Array of string log messages or null if log message recording was not enabled during this transaction
	Accounts      []EncodedAccount `json:"accounts"`      //array of accounts with the same length as the accounts.addresses array in the request
	UnitsConsumed *uint            `json:"unitsConsumed"` //The number of compute budget units consumed during the processing of this transaction
	ReturnData    *struct {
		ProgramID string `json:"programId"` //The program that generated the return data, as base-58 encoded Pubkey
		Data      string `json:"data"`      //The return data, as base64 encoded string
	} //the most-recent return data generated by an instruction in the transaction
	InnerInstructions []InnerInstructions `json:"innerInstructions"` //Defined only if innerInstructions was set to true. The value is a list of inner instructions.
}

// For preflight checks and transaction processing, Solana nodes choose which bank state to query based on a commitment requirement set by the client. The commitment describes how finalized a block is at that point in time. When querying the ledger state, it's recommended to use lower levels of commitment to report progress and higher levels to ensure the state will not be rolled back.
type Commitment string

const (
	CommitmentFinalized Commitment = "finalized" // (Most finalized) The node will query the most recent block confirmed by supermajority of the cluster as having reached maximum lockout, meaning the cluster has recognized this block as finalized
	CommitmentConfirmed Commitment = "confirmed" // (Middle finalized) The node will query the most recent block that has been voted on by supermajority of the cluster.
	CommitmentProcessed Commitment = "processed" // (Least finalized) The node will query its most recent block. Note that the block may still be skipped by the cluster.
)

// Encoding format for Account data
type Encoding string

const (
	EncodingBase58     Encoding = "base58"      // Is slow and limited to less than 129 bytes of Account data.
	EncodingBase64     Encoding = "base64"      // Will return base64 encoded data for Account data of any size.
	EncodingBase64Zstd Encoding = "base64+zstd" // Compresses the Account data using Zstandard and base64-encodes the result.
	EncodingJson       Encoding = "json"        // Encoding attempts to use program-specific state parsers to return more human-readable and explicit account state data.
	EncodingJsonParsed Encoding = "jsonParsed"  // Encoding attempts to use program-specific state parsers to return more human-readable and explicit account state data.
)

// Level of transaction detail to return -- default is "full"
type TransactionDetails string

const (
	TransactionDetailsFull       TransactionDetails = "full"       // Return all transaction fields
	TransactionDetailsAccounts   TransactionDetails = "accounts"   // If accounts are requested, transaction details only include signatures and an annotated list of accounts in each transaction.
	TransactionDetailsSignatures TransactionDetails = "signatures" // If signatures are requested, transaction details only include signatures.
	TransactionDetailsNone       TransactionDetails = "none"
)

// Request a slice of the account's data.
type DataSlice struct {
	Offset int `json:"offset"`
	Length int `json:"length"`
}

type Filter string

const (
	FilterCirculating    Filter = "circulating"
	FilterNonCirculating Filter = "nonCirculating"
)

// const LocalRpcUrl = "http://localhost:8899"
// const DevnetRpcUrl = "https://api.devnet.solana.com"
// const TestnetRpcUrl = "https://api.testnet.solana.com"
// const MainnetRpcUrl = "https://api.mainnet-beta.solana.com"

// type RpcRequest[T any] struct {
// 	MethodName string `json:"methodName"`
// 	Params     T      `json:"params"`
// }

// type RpcMessage[T any] struct {
// 	ID      string `json:"id"`
// 	Jsonrpc string `json:"jsonrpc"`
// 	Method  string `json:"method"`
// 	Params  T      `json:"params"`
// }

// // Required for the JSON-RPC request.
// var nextMessageId = 0

// func getNextMessageId() string {
// 	id := nextMessageId
// 	nextMessageId++
// 	return fmt.Sprintf("%d", id)
// }

// func CreateRpcMessage(request RpcRequest[any]) RpcMessage[any] {
// 	return RpcMessage[any]{
// 		ID:      getNextMessageId(),
// 		Jsonrpc: "2.0",
// 		Method:  request.MethodName,
// 		Params:  request.Params,
// 	}
// }

// type RpcErrorResponsePayload struct {
// 	Code    int         `json:"code"`
// 	Data    interface{} `json:"data"`
// 	Message string      `json:"message"`
// }
