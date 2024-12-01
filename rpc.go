package solana

import (
	"math/big"
	"time"
)

type Rpc interface {
	GetAccountInfo(address Pubkey, config ...GetAccountInfoConfig) (*EncodedAccount, error)                                                             //Returns all information associated with the account of provided Pubkey
	GetBalance(address Pubkey, config ...StandardRpcConfig) (uint, error)                                                                               //Returns the lamport balance of the account of provided Pubkey
	GetBlock(slotNumber uint, config ...GetBlockConfig) (*Block, error)                                                                                 //Returns identity and transaction information about a confirmed block in the ledger
	GetBlockCommitment(slotNumber uint) (BlockCommitment, error)                                                                                        //Returns the current block height and the estimated production time of a block
	GetBlockHeight(config ...StandardRpcConfig) (uint, error)                                                                                           //Returns the current block height of the node
	GetBlockProduction(config ...GetBlockProductionConfig) (BlockProduction, error)                                                                     //Returns recent block production information from the current or previous epoch.
	GetBlockTime(slotNumber uint) (time.Time, error)                                                                                                    //Returns the estimated production time of a block as a unix timestamp.
	GetBlocks(startSlot uint, endSlot *uint, config ...GetBlockConfig) ([]uint, error)                                                                  //Returns a list of confirmed blocks between two slots
	GetBlocksWithLimit(startSlot uint, limit uint, config ...GetBlockConfig) ([]uint, error)                                                            //Returns a list of confirmed blocks starting at the given slot
	GetClusterNodes() ([]ClusterNode, error)                                                                                                            //Returns information about all the nodes participating in the cluster
	GetEpochInfo(...StandardRpcConfig) (EpochInfo, error)                                                                                               //Returns information about the current epoch
	GetEpochSchedule() (EpochSchedule, error)                                                                                                           //Returns the epoch schedule information from this cluster's genesis config
	GetFeeForMessage(msg string, config ...StandardRpcConfig) (*uint, error)                                                                            //Get the fee the network will charge for a particular Message. NOTE: This method is only available in solana-core v1.9 or newer. Please use getFees for solana-core v1.8 and below.
	GetFirstAvailableBlock() (uint, error)                                                                                                              //Returns the slot of the lowest confirmed block that has not been purged from the ledger
	GetGenesisHash() (string, error)                                                                                                                    //Returns the genesis hash
	GetHealth() error                                                                                                                                   //Returns the current health of the node. A healthy node is one that is within HEALTH_CHECK_SLOT_DISTANCE slots of the latest cluster confirmed slot.
	GetHighestSnapshotSlot() (HighestSnapshotSlot, error)                                                                                               //Returns the highest slot information that the node has snapshots for. This will find the highest full snapshot slot, and the highest incremental snapshot slot based on the full snapshot slot, if there is one.
	GetIdentity() (Pubkey, error)                                                                                                                       //Returns the identity pubkey for the current node
	GetInflationGovernor(config ...StandardCommitmentConfig) (InflationGovernor, error)                                                                 //Returns the current inflation governor
	GetInflationRate(config ...StandardCommitmentConfig) (InflationRate, error)                                                                         //Returns the specific inflation values for the current epoch
	GetInflationReward(addresses []Pubkey, config ...GetInflationRewardConfig) ([]*InflationReward, error)                                              //Returns the inflation / staking reward for a list of addresses for an epoch
	GetLargestAccounts(config ...GetLargestAccountsConfig) ([]AccountWithBalance, error)                                                                //Returns the 20 largest accounts, by lamport balance (results may be cached up to two hours)
	GetLatestBlockhash(config ...StandardRpcConfig) (LatestBlockhash, error)                                                                            //Returns the latest blockhash. NOTE: This method is only available in solana-core v1.9 or newer. Please use getRecentBlockhash for solana-core v1.8 and below.
	GetLeaderSchedule(slot *uint, config ...GetLeaderScheduleConfig) (*LeaderSchedule, error)                                                           //Returns the leader schedule for an epoch
	GetMaxRetransmitSlot() (uint, error)                                                                                                                //Get the max slot seen from retransmit stage.
	GetMaxShredInsertSlot() (uint, error)                                                                                                               //Get the max slot seen from after shred insert.
	GetMinimumBalanceForRentExemption(accountDataLength uint, config ...StandardCommitmentConfig) (uint, error)                                         //Returns minimum balance required to make account rent exempt.
	GetMultipleAccounts(pubkeys []Pubkey, config ...GetAccountInfoConfig) ([]*EncodedAccount, error)                                                    //Returns the account information for a list of Pubkeys.
	GetProgramAccounts(programPubkey Pubkey, config ...GetAccountInfoConfig) ([]EncodedAccount, error)                                                  //Returns all accounts owned by the provided program Pubkey
	GetRecentPerformanceSamples(limit uint) ([]PerformanceSample, error)                                                                                //Returns a list of recent performance samples, in reverse slot order. Performance samples are taken every 60 seconds and include the number of transactions and slots that occur in a given time window. -- NOTE max limit is 720
	GetRecentPrioritizationFees(addresses []Pubkey) ([]PrioritizationFee, error)                                                                        //Returns a list of prioritization fees from recent blocks.
	GetSignatureStatuses(signatures []string, config ...GetSignatureStatusesConfig) ([]*SignatureStatus, error)                                         //Returns the statuses of a list of signatures. Each signature must be a txid, the first signature of a transaction. Unless the searchTransactionHistory configuration parameter is included, this method only searches the recent status cache of signatures, which retains statuses for all active slots plus MAX_RECENT_BLOCKHASHES rooted slots.
	GetSignaturesForAddress(address Pubkey, config ...GetSignaturesForAddressConfig) ([]TransactionSignature, error)                                    //Returns signatures for confirmed transactions that include the given address in their accountKeys list. Returns signatures backwards in time from the provided signature or most recent confirmed block
	GetSlot(config ...StandardRpcConfig) (uint, error)                                                                                                  //Returns the slot that has reached the given or default commitment level
	GetSlotLeader(config ...StandardRpcConfig) (Pubkey, error)                                                                                          //Returns the current slot leader
	GetSlotLeaders(start, limit *uint) ([]Pubkey, error)                                                                                                //Returns the slot leaders for a given slot range
	GetStakeMinimumDelegation(config ...StandardCommitmentConfig) (uint, error)                                                                         //Returns the stake minimum delegation, in lamports.
	GetSupply(config ...GetSupplyConfig) (Supply, error)                                                                                                //Returns information about the current supply.
	GetTokenAccountBalance(address Pubkey, config ...StandardCommitmentConfig) (UiTokenAmount, error)                                                   //Returns the token balance of an SPL Token account.
	GetTokenAccountsByDelegate(delegateAddress Pubkey, opts GetTokenAccountsByDelegateConfig, config ...GetAccountInfoConfig) ([]EncodedAccount, error) //Returns all SPL Token accounts by approved Delegate.
	GetTokenAccountsByOwner(ownerAddress Pubkey, opts GetTokenAccountsByDelegateConfig, config ...GetAccountInfoConfig) ([]EncodedAccount, error)       //Returns all SPL Token accounts by token owner.
	GetTokenLargestAccounts(mintAddress Pubkey, config ...StandardCommitmentConfig) ([]UiTokenAmount, error)                                            //Returns the 20 largest accounts of a particular SPL Token type.
	GetTokenSupply(mintAddress Pubkey, config ...StandardCommitmentConfig) (UiTokenAmount, error)                                                       //Returns the total supply of an SPL Token type.
	GetTransaction(transactionSignature string, config ...GetTransactionSignatureConfig) (*TransactionWithMeta, error)                                  //Returns transaction details for a confirmed transaction
	GetTransactionCount(config ...StandardRpcConfig) (uint, error)                                                                                      //Returns the current transaction count from the ledger
	GetVersion() (Version, error)                                                                                                                       //Returns the current Solana version running on the node
	GetVoteAccounts(config ...GetVoteAccountsConfig) (VoteAccounts, error)                                                                              //Returns the account info and associated stake for all the voting accounts in the current bank.
	IsBlockhashValid(blockhash string, config ...StandardRpcConfig) (bool, error)                                                                       //Returns whether a blockhash is valid
	MinimumLedgerSlot() (uint, error)                                                                                                                   //Returns the lowest slot that the node has information about in its ledger.
	RequestAirdrop(destinationAddress Pubkey, lamports uint, config ...StandardCommitmentConfig) (string, error)                                        //Requests an airdrop of lamports to a Solana account
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

type RpcEndpoint string

const (
	RpcEndpointMainnetBeta RpcEndpoint = "https://api.mainnet-beta.solana.com"
	RpcEndpointDevnet      RpcEndpoint = "https://api.devnet.solana.com"
	RpcEndpointTestnet     RpcEndpoint = "https://api.testnet.solana.com"
	RpcEndpointLocalnet    RpcEndpoint = "http://locahost:8899"
)

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

type Filter string

const (
	FilterCirculating    Filter = "circulating"
	FilterNonCirculating Filter = "nonCirculating"
)

type StandardRpcConfig struct {
	Commitment     *Commitment `json:"commitment,omitempty"`     //For preflight checks and transaction processing, Solana nodes choose which bank state to query based on a commitment requirement set by the client. The commitment describes how finalized a block is at that point in time. When querying the ledger state, it's recommended to use lower levels of commitment to report progress and higher levels to ensure the state will not be rolled back.
	MinContextSlot *int        `json:"minContextSlot,omitempty"` //The minimum slot that the request can be evaluated at
}

type StandardCommitmentConfig struct {
	Commitment *Commitment `json:"commitment,omitempty"` //For preflight checks and transaction processing, Solana nodes choose which bank state to query based on a commitment requirement set by the client. The commitment describes how finalized a block is at that point in time. When querying the ledger state, it's recommended to use lower levels of commitment to report progress and higher levels to ensure the state will not be rolled back.
}

type GetAccountInfoConfig struct {
	Commitment     *Commitment `json:"commitment,omitempty"`     //For preflight checks and transaction processing, Solana nodes choose which bank state to query based on a commitment requirement set by the client. The commitment describes how finalized a block is at that point in time. When querying the ledger state, it's recommended to use lower levels of commitment to report progress and higher levels to ensure the state will not be rolled back.
	Encoding       *Encoding   `json:"encoding,omitempty"`       //Encoding format for Account data
	DataSlice      *DataSlice  `json:"dataSlice,omitempty"`      //Request a slice of the account's data.
	MinContextSlot *int        `json:"minContextSlot,omitempty"` //The minimum slot that the request can be evaluated at
}

type GetBlockConfig struct {
	Commitment                     *Commitment         `json:"commitment,omitempty"`                     //For preflight checks and transaction processing, Solana nodes choose which bank state to query based on a commitment requirement set by the client. The commitment describes how finalized a block is at that point in time. When querying the ledger state, it's recommended to use lower levels of commitment to report progress and higher levels to ensure the state will not be rolled back.
	Encoding                       *Encoding           `json:"encoding,omitempty"`                       //Encoding format for each returned Transaction
	TransactionDetails             *TransactionDetails `json:"transactionDetails,omitempty"`             //Level of transaction detail to return -- default is "full"
	MaxSupportedTransactionVersion int                 `json:"maxSupportedTransactionVersion,omitempty"` //If this parameter is omitted, only legacy transactions will be returned, and a block containing any versioned transaction will prompt the error.
	Rewards                        bool                `json:"rewards,omitempty"`                        //Whether to populate the rewards array. If parameter not provided, the default includes rewards.
}

type GetBlockProductionConfig struct {
	Commitment *Commitment `json:"commitment,omitempty"` //For preflight checks and transaction processing, Solana nodes choose which bank state to query based on a commitment requirement set by the client. The commitment describes how finalized a block is at that point in time. When querying the ledger state, it's recommended to use lower levels of commitment to report progress and higher levels to ensure the state will not be rolled back.
	Identity   *string     `json:"identity,omitempty"`   //Only return results for this validator identity (base-58 encoded)
	Range      *BlockRange `json:"range,omitempty"`      //Slot range to return block production for. If parameter not provided, defaults to current epoch.
}

type GetBlocksConfig struct {
	Commitment *Commitment `json:"commitment,omitempty"` //"processed" is not supported
}

type GetInflationRewardConfig struct {
	MinContextSlot *int  `json:"minContextSlot,omitempty"` //The minimum slot that the request can be evaluated at
	Epoch          *uint `json:"epoch,omitempty"`          //An epoch for which the reward occurs. If omitted, the previous epoch will be used
}

type GetLargestAccountsConfig struct {
	Commitment *Commitment `json:"commitment,omitempty"` //For preflight checks and transaction processing, Solana nodes choose which bank state to query based on a commitment requirement set by the client. The commitment describes how finalized a block is at that point in time. When querying the ledger state, it's recommended to use lower levels of commitment to report progress and higher levels to ensure the state will not be rolled back.
	Filter     *string     `json:"filter,omitempty"`     //Filter results by account type
}

type GetLeaderScheduleConfig struct {
	Commitment *Commitment `json:"commitment,omitempty"` //For preflight checks and transaction processing, Solana nodes choose which bank state to query based on a commitment requirement set by the client. The commitment describes how finalized a block is at that point in time. When querying the ledger state, it's recommended to use lower levels of commitment to report progress and higher levels to ensure the state will not be rolled back.
	Identity   *string     `json:"identity,omitempty"`   //Only return results for this validator identity (base-58 encoded)
}

type GetSignatureStatusesConfig struct {
	SearchTransactionHistory bool `json:"searchTransactionHistory"` //If true - a Solana node will search its ledger cache for any signatures not found in the recent status cache
}

type GetSignaturesForAddressConfig struct {
	Commitment     *Commitment `json:"commitment,omitempty"`     //For preflight checks and transaction processing, Solana nodes choose which bank state to query based on a commitment requirement set by the client. The commitment describes how finalized a block is at that point in time. When querying the ledger state, it's recommended to use lower levels of commitment to report progress and higher levels to ensure the state will not be rolled back.
	MinContextSlot *int        `json:"minContextSlot,omitempty"` //The minimum slot that the request can be evaluated at
	Limit          *uint       `json:"limit,omitempty"`          //Maximum transaction signatures to return (between 1 and 1,000). Default is 1000.
	Before         *string     `json:"before,omitempty"`         //Start searching backwards from this transaction signature. If not provided the search starts from the top of the highest max confirmed block.
	After          *string     `json:"after,omitempty"`          //Search until this transaction signature, if found before limit reached
}

type GetSupplyConfig struct {
	Commitment                        *Commitment `json:"commitment,omitempty"`                        //For preflight checks and transaction processing, Solana nodes choose which bank state to query based on a commitment requirement set by the client. The commitment describes how finalized a block is at that point in time. When querying the ledger state, it's recommended to use lower levels of commitment to report progress and higher levels to ensure the state will not be rolled back.
	ExcludeNonCirculatingAccountsList *bool       `json:"excludeNonCirculatingAccountsList,omitempty"` //Exclude non circulating accounts list from response
}

// Supply one of the following fields
type GetTokenAccountsByDelegateConfig struct {
	Mint      *Pubkey `json:"mint,omitempty"`      //Pubkey of the specific token Mint to limit accounts to, as base-58 encoded string; or
	ProgramID *Pubkey `json:"programId,omitempty"` //Pubkey of the Token program that owns the accounts, as base-58 encoded string
}

type GetTransactionSignatureConfig struct {
	Commitment                     *Commitment `json:"commitment,omitempty"`                     //For preflight checks and transaction processing, Solana nodes choose which bank state to query based on a commitment requirement set by the client. The commitment describes how finalized a block is at that point in time. When querying the ledger state, it's recommended to use lower levels of commitment to report progress and higher levels to ensure the state will not be rolled back.
	Encoding                       *Encoding   `json:"encoding,omitempty"`                       //Encoding format for each returned Transaction
	MaxSupportedTransactionVersion int         `json:"maxSupportedTransactionVersion,omitempty"` //Set the max transaction version to return in responses. If the requested transaction is a higher version, an error will be returned. If this parameter is omitted, only legacy transactions will be returned, and any versioned transaction will prompt the error.
}

type GetVoteAccountsConfig struct {
	Commitment *Commitment `json:"commitment,omitempty"` //For preflight checks and transaction processing, Solana nodes choose which bank state to query based on a commitment requirement set by the client. The commitment describes how finalized a block is at that point in time. When querying the ledger state, it's recommended to use lower levels of commitment to report progress and higher levels to ensure the state will not be rolled back.
	VotePubkey *string     `json:"votePubkey,omitempty"` //Only return results for this validator vote address (base-58 encoded)

}

type SendTransactionConfig struct {
	Encoding            *Encoding   `json:"encoding,omitempty"`            //Encoding used for the transaction data. Values: base58 (slow, DEPRECATED), or base64.
	SkipPreflight       *bool       `json:"skipPreflight,omitempty"`       //When true, skip the preflight transaction checks
	PreflightCommitment *Commitment `json:"preflightCommitment,omitempty"` //Default: finalized. Commitment level to use for preflight.
	MaxRetries          *uint       `json:"maxRetries,omitempty"`          //Maximum number of times for the RPC node to retry sending the transaction to the leader. If this parameter not provided, the RPC node will retry the transaction until it is finalized or until the blockhash expires.
	MinContextSlot      *uint       `json:"minContextSlot,omitempty"`      //The minimum slot that the request can be evaluated at
}

type SimulateTransactionConfig struct {
	Commitment             *Commitment `json:"commitment,omitempty"`             //Commitment level to simulate the transaction at
	SigVerify              *bool       `json:"sigVerify,omitempty"`              //If true the transaction signatures will be verified (conflicts with replaceRecentBlockhash)
	ReplaceRecentBlockhash *bool       `json:"replaceRecentBlockhash,omitempty"` //If true the transaction recent blockhash will be replaced with the most recent blockhash. (conflicts with sigVerify)
	MinContextSlot         *uint       `json:"minContextSlot,omitempty"`         //The minimum slot that the request can be evaluated at
	Encoding               *Encoding   `json:"encoding,omitempty"`               //Encoding used for the transaction data. Values: base58 (slow, DEPRECATED), or base64.
	InnerInstructions      *bool       `json:"innerInstructions,omitempty"`      //If true the response will include inner instructions. These inner instructions will be jsonParsed where possible, otherwise json.
	Accounts               *struct {
		Addresses []string `json:"addresses"` //An array of accounts to return, as base-58 encoded strings
		Encoding  Encoding `json:"encoding"`  //Encoding for returned Account data
	} `json:"accounts,omitempty"` //Accounts configuration object
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

// Request a slice of the account's data.
type DataSlice struct {
	Offset int `json:"offset"`
	Length int `json:"length"`
}

type Block struct {
	BlockHeight       *int                  `json:"blockHeight"`       //The number of blocks beneath this block
	BlockTime         *int                  `json:"blockTime"`         //Estimated production time, as Unix timestamp (seconds since the Unix epoch). null if not available
	Blockhash         string                `json:"blockhash"`         //The blockhash of this block, as base-58 encoded string
	ParentSlot        uint                  `json:"parentSlot"`        //The slot index of this block's parent
	PreviousBlockhash string                `json:"previousBlockhash"` //The blockhash of this block's parent, as base-58 encoded string; if the parent block is not available due to ledger cleanup, this field will return "11111111111111111111111111111111"
	Transactions      []TransactionWithMeta `json:"transactions"`      //The list of transactions included in this block
	Signatures        []string              `json:"signatures"`        //Present if "signatures" are requested for transaction details; an array of signatures strings, corresponding to the transaction order in the block
	Rewards           []TransactionReward   `json:"rewards"`           //Block-level rewards, present if rewards are requested; an array of JSON objects containing:
}

type BlockCommitment struct {
	Commitment []uint //Commitment, array of u64 integers logging the amount of cluster stake in lamports that has voted on the block at each depth from 0 to MAX_LOCKOUT_HISTORY + 1
	TotalStake uint   //Total active stake, in lamports, of the current epoch
}

type BlockRange struct {
	FirstSlot uint  `json:"firstSlot"` //first slot to return block production information for (inclusive)
	LastSlot  *uint `json:"lastSlot"`  //Last slot to return block production information for (inclusive). If parameter not provided, defaults to the highest slot
}

type BlockProduction struct {
	ByIdentity map[string][]uint `json:"byIdentity"` //A dictionary of validator identities, as base-58 encoded strings. Value is a two element array containing the number of leader slots and the number of blocks produced.
	Range      BlockRange        `json:"range"`      //Block production slot range
}

type ClusterNode struct {
	Pubkey       string  `json:"pubkey"`       //Node public key, as base-58 encoded string
	Gossip       *string `json:"gossip"`       //Gossip network address for the node
	Tpu          *string `json:"tpu"`          //TPU network address for the node
	Rpc          *string `json:"rpc"`          //JSON RPC network address for the node, or null if the JSON RPC service is not enabled
	Version      *string `json:"version"`      //The software version of the node, or null if the version information is not available
	FeatureSet   *uint   `json:"featureSet"`   //The unique identifier of the node's feature set
	ShredVersion *uint   `json:"shredVersion"` //The shred version the node has been configured to use
}

type EpochInfo struct {
	AbsoluteSlot     uint  `json:"absoluteSlot"`     //The current slot
	BlockHeight      uint  `json:"blockHeight"`      //The current block height
	Epoch            uint  `json:"epoch"`            //The current epoch
	SlotIndex        uint  `json:"slotIndex"`        //The current slot relative to the start of the current epoch
	SlotsInEpoch     uint  `json:"slotsInEpoch"`     //The number of slots in this epoch
	TransactionCount *uint `json:"transactionCount"` //Total number of transactions processed without error since genesis
}

type EpochSchedule struct {
	SlotsPerEpoch            uint `json:"slotsPerEpoch"`            //The maximum number of slots in each epoch
	LeaderScheduleSlotOffset uint `json:"leaderScheduleSlotOffset"` //The number of slots before beginning of an epoch to calculate a leader schedule for that epoch
	Warmup                   bool `json:"warmup"`                   //Whether epochs start short and grow
	FirstNormalEpoch         uint `json:"firstNormalEpoch"`         //First normal-length epoch, log2(slotsPerEpoch) - log2(MINIMUM_SLOTS_PER_EPOCH)
	FirstNormalSlot          uint `json:"firstNormalSlot"`          //MINIMUM_SLOTS_PER_EPOCH * (2.pow(firstNormalEpoch) - 1)
}

type HighestSnapshotSlot struct {
	Full        uint  `json:"full"`        //Highest full snapshot slot
	Incremental *uint `json:"incremental"` //Highest incremental snapshot slot based on full
}

type InflationGovernor struct {
	Initial        float64 `json:"initial"`        //The initial inflation percentage from time 0
	Terminal       float64 `json:"terminal"`       //The terminal inflation percentage
	Taper          float64 `json:"taper"`          //Rate per year at which inflation is lowered. (Rate reduction is derived using the target slot time in genesis config)
	Foundation     float64 `json:"foundation"`     //Percentage of total inflation allocated to the foundation
	FoundationTerm float64 `json:"foundationTerm"` //Duration of foundation pool inflation in years
}

type InflationRate struct {
	Total      float64 `json:"total"`      //Total inflation
	Validator  float64 `json:"validator"`  //Inflation awarded to validators
	Foundation float64 `json:"foundation"` //Inflation awarded to the foundation
	Epoch      uint    `json:"epoch"`      //Epoch for which these values are valid
}

type InflationReward struct {
	Epoch         uint `json:"epoch"`         //The epoch for which rewards are calculated
	EffectiveSlot uint `json:"effectiveSlot"` //The slot in which the rewards are effective
	Amount        uint `json:"amount"`        //Reward amount in lamports
	PostBalance   uint `json:"postBalance"`   //Post balance of the account in lamports
	Comission     uint `json:"comission"`     //Vote account commission when the reward was credited
}

type LeaderSchedule map[string][]uint //A dictionary of validator identities, as base-58 encoded strings, and their corresponding leader slot indices as values (indices are relative to the first slot in the requested epoch)

type PerformanceSample struct {
	Slot                   uint `json:"slot"`                   //Slot in which sample was taken at
	NumTransactions        uint `json:"numTransactions"`        //Number of transactions processed during the sample period
	NumSlots               uint `json:"numSlots"`               //Number of slots completed during the sample period
	SamplePeriodSecs       uint `json:"samplePeriodSecs"`       //Number of seconds in a sample window
	NumNonVoteTransactions uint `json:"numNonVoteTransactions"` //Number of non-vote transactions processed during the sample period
}

type PrioritizationFee struct {
	Slot              uint `json:"slot"`              //The slot for which the fee applies
	PrioritizationFee uint `json:"prioritizationFee"` //The per-compute-unit fee paid by at least one successfully landed transaction, specified in increments of micro-lamports (0.000001 lamports)
}

type SignatureStatus struct {
	Slot               uint              `json:"slot"`               //The slot in which the transaction was processed
	Confirmations      *uint             `json:"confirmations"`      //Number of blocks since signature confirmation, null if rooted, as well as finalized by a supermajority of the cluster
	Err                any               `json:"err"`                //Error if transaction failed
	ConfirmationStatus *Commitment       `json:"confirmationStatus"` //The transaction's cluster confirmation status
	Status             TransactionStatus `json:"status"`             //Deprecated: Transaction status
}

type TransactionSignature struct {
	Signature          string      `json:"signature"`          //The transaction signature, as base-58 encoded string
	Slot               uint        `json:"slot"`               //The slot that contains the block with the transaction
	Err                any         `json:"err"`                //Error if transaction failed
	Memo               *string     `json:"memo"`               //Memo associated with the transaction, null if no memo is present
	BlockTime          *int        `json:"blockTime"`          //Estimated production time, as Unix timestamp (seconds since the Unix epoch) of when transaction was processed. null if not available.
	ConfirmationStatus *Commitment `json:"confirmationStatus"` //The transaction's cluster confirmation status
}

type Version struct {
	SolanaCore string `json:"solana-core"` //software version of solana-core as a string
	FeatureSet uint   `json:"feature-set"` //unique identifier of the feature set
}

type VoteAccounts struct {
	Current    []VoteAccount `json:"current"`    //The current vote accounts
	Delinquent []VoteAccount `json:"delinquent"` //The delinquent vote accounts
}

type VoteAccount struct {
	VotePubkey       string   `json:"votePubkey"`       //Vote account public key, as base-58 encoded string
	NodePubkey       string   `json:"nodePubkey"`       //Validator identity, as base-58 encoded string
	ActivatedStake   uint     `json:"activatedStake"`   //The stake, in lamports, delegated to this vote account and activated
	EpochVoteAccount bool     `json:"epochVoteAccount"` //Whether the vote account is staked for this epoch
	Comission        uint     `json:"comission"`        //Percentage (0-100) of rewards payout owed to the vote account
	LastVote         uint     `json:"lastVote"`         //Most recent slot voted on by this vote account
	EpochCredits     [][]uint `json:"epochCredits"`     //Latest history of earned credits for up to five epochs, as an array of arrays containing: [epoch, credits, previousCredits].
	RootSlot         uint     `json:"rootSlot"`         //Current root slot for this vote account
}

type AccountWithBalance struct {
	Address  Pubkey `json:"address"` //Base-58 encoded address of the account
	Lamports uint   `json:"balance"` //Number of lamports in the account, as a u64
}

type Supply struct {
	Total                  uint     `json:"total"`                  //Total supply in lamports
	Circulating            uint     `json:"circulating"`            //Circulating supply in lamports
	NonCirculating         uint     `json:"nonCirculating"`         //Non-circulating supply in lamports
	NonCirculatingAccounts []string `json:"nonCirculatingAccounts"` //An array of account addresses of non-circulating accounts, as strings. If excludeNonCirculatingAccountsList is enabled, the returned array will be empty.
}

// Defines a Solana account with its generic details and encoded data.
type EncodedAccount struct {
	Address    Pubkey  `json:"address"`
	Data       []byte  `json:"data"`
	Executable bool    `json:"executable"`
	Lamports   uint    `json:"lamports"`
	Owner      Pubkey  `json:"owner"`
	RentEpoch  big.Int `json:"rentEpoch"`
	Space      int     `json:"space"`
}
