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
