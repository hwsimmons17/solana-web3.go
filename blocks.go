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
