package solana

type Block struct {
	BlockHeight       *int                `json:"blockHeight"`       //The number of blocks beneath this block
	BlockTime         *int                `json:"blockTime"`         //Estimated production time, as Unix timestamp (seconds since the Unix epoch). null if not available
	Blockhash         string              `json:"blockhash"`         //The blockhash of this block, as base-58 encoded string
	ParentSlot        uint                `json:"parentSlot"`        //The slot index of this block's parent
	PreviousBlockhash string              `json:"previousBlockhash"` //The blockhash of this block's parent, as base-58 encoded string; if the parent block is not available due to ledger cleanup, this field will return "11111111111111111111111111111111"
	Transactions      []Transaction       `json:"transactions"`      //The list of transactions included in this block
	Signatures        []string            `json:"signatures"`        //Present if "signatures" are requested for transaction details; an array of signatures strings, corresponding to the transaction order in the block
	Rewards           []TransactionReward `json:"rewards"`           //Block-level rewards, present if rewards are requested; an array of JSON objects containing:
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
