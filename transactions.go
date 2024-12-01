package solana

type TransactionWithMeta struct {
	Meta        *TransactionMeta `json:"meta"`
	Version     *int             `json:"version"` //Transaction version. Undefined if maxSupportedTransactionVersion is not set in request params. --note can also be "legacy"
	Transaction Transaction      `json:"transaction"`
}

type Transaction struct {
	Message    TransactionMessage `json:"message"`
	Signatures []string           `json:"signatures"`
}

type TransactionMeta struct {
	Err                  any                    `json:"err"`                  //Error if transaction failed, null if transaction succeeded. TransactionError definitions
	Fee                  uint                   `json:"fee"`                  //Ree this transaction was charged, as u64 integer
	InnerInstructions    []InnerInstructions    `json:"innerInstructions"`    //List of inner instructions or null if inner instruction recording was not enabled during this transaction
	LogMessages          []string               `json:"logMessages"`          //Array of string log messages or null if log message recording was not enabled during this transaction
	PostBalances         []uint                 `json:"postBalances"`         //Array of u64 account balances after the transaction was processed
	PostTokenBalances    []TokenBalance         `json:"postTokenBalances"`    //List of token balances from after the transaction was processed or omitted if token balance recording was not yet enabled during this transaction
	PreBalances          []uint                 `json:"preBalances"`          //Array of u64 account balances from before the transaction was processed
	PreTokenBalances     []TokenBalance         `json:"preTokenBalances"`     //List of token balances from before the transaction was processed or omitted if token balance recording was not yet enabled during this transaction
	Rewards              []TransactionReward    `json:"rewards"`              //Transaction-level rewards, populated if rewards are requested; an array of JSON objects containing:
	Status               TransactionStatus      `json:"status"`               //Deprecated: Transaction status
	LoadedAddresses      LoadedAddresses        `json:"loadedAddresses"`      //Transaction addresses loaded from address lookup tables. Undefined if maxSupportedTransactionVersion is not set in request params, or if jsonParsed encoding is set in request params.
	ReturnData           *TransactionReturnData `json:"returnData"`           //The most-recent return data generated by an instruction in the transaction.
	ComputeUnitsConsumed *uint                  `json:"computeUnitsConsumed"` //The number of compute units consumed during the execution of the transaction
}

type TransactionMessage struct {
	AccountKeys     []Pubkey                 `json:"accountKeys"`
	Header          TransactionHeader        `json:"header"`
	Instructions    []TransactionInstruction `json:"instructions"`
	RecentBlockhash string                   `json:"recentBlockhash"`
}

type TransactionInstruction struct {
	Accounts       []int  `json:"accounts"`
	Data           []byte `json:"data"`
	ProgramIDIndex int    `json:"programIdIndex"`
}

type TransactionHeader struct {
	NumReadonlySignedAccounts   int `json:"numReadonlySignedAccounts"`
	NumReadonlyUnsignedAccounts int `json:"numReadonlyUnsignedAccounts"`
	NumRequiredSignatures       int `json:"numRequiredSignatures"`
}

// The Solana runtime records the cross-program instructions that are invoked during transaction processing and makes these available for greater transparency of what was executed on-chain per transaction instruction. Invoked instructions are grouped by the originating transaction instruction and are listed in order of processing.
type InnerInstructions struct {
	Index        int                      `json:"index"`        //Index of the transaction instruction from which the inner instruction(s) originated
	Instructions []TransactionInstruction `json:"instructions"` //Ordered list of inner program instructions that were invoked during a single transaction instruction.
}

type TransactionReward struct {
	Pubkey      string      `json:"pubkey"`      //The public key, as base-58 encoded string, of the account that received the reward
	Lamports    int         `json:"lamports"`    //number of reward lamports credited or debited by the account, as a i64
	PostBalance uint        `json:"postBalance"` //Account balance in lamports after the reward was applied
	RewardType  *RewardType `json:"rewardType"`  //Type of reward: "fee", "rent", "voting", "staking"
	Commission  *uint       `json:"commission"`  //Vote account commission when the reward was credited, only present for voting and staking rewards
}

type RewardType string

const (
	RewardTypeFee     RewardType = "fee"
	RewardTypeRent    RewardType = "rent"
	RewardTypeVoting  RewardType = "voting"
	RewardTypeStaking RewardType = "staking"
)

// Deprecated: Transaction status
type TransactionStatus struct {
	Ok  any `json:"Ok"`
	Err any `json:"Err"`
}

type LoadedAddresses struct {
	Writable []string `json:"writable"` //Writable account addresses
	Readonly []string `json:"readonly"` //Readonly account addresses
}

type TransactionReturnData struct {
	ProgramID string `json:"programId"` //The program that generated the return data, as base-58 encoded Pubkey
	Data      string `json:"data"`      //the return data itself, as base-64 encoded binary data
}

type LatestBlockhash struct {
	Blockhash            string `json:"blockhash"`            //A Hash as base-58 encoded string
	LastValidBlockHeight uint   `json:"lastValidBlockHeight"` //Last block height at which the blockhash will be valid
}
