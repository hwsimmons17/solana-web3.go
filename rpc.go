package solana

import "fmt"

type Rpc interface {
	GetAccountInfo(address Address, config ...GetAccountInfoConfig) (EncodedAccount, error) //Returns all information associated with the account of provided Pubkey
	GetBalance(address Address, config ...GetBalanceConfig) (uint, error)                   //Returns the lamport balance of the account of provided Pubkey
	GetBlock(slotNumber uint, config ...GetBlockConfig) (*Block, error)                     //Returns identity and transaction information about a confirmed block in the ledger

}

type GetAccountInfoConfig struct {
	Commitment     *Commitment `json:"commitment"`     //For preflight checks and transaction processing, Solana nodes choose which bank state to query based on a commitment requirement set by the client. The commitment describes how finalized a block is at that point in time. When querying the ledger state, it's recommended to use lower levels of commitment to report progress and higher levels to ensure the state will not be rolled back.
	Encoding       *Encoding   `json:"encoding"`       //Encoding format for Account data
	DataSlice      *DataSlice  `json:"dataSlice"`      //Request a slice of the account's data.
	MinContextSlot *int        `json:"minContextSlot"` //The minimum slot that the request can be evaluated at
}

type GetBalanceConfig struct {
	Commitment     *Commitment `json:"commitment"`     //For preflight checks and transaction processing, Solana nodes choose which bank state to query based on a commitment requirement set by the client. The commitment describes how finalized a block is at that point in time. When querying the ledger state, it's recommended to use lower levels of commitment to report progress and higher levels to ensure the state will not be rolled back.
	MinContextSlot *int        `json:"minContextSlot"` //The minimum slot that the request can be evaluated at
}

type GetBlockConfig struct {
	Commitment                     *Commitment         `json:"commitment"`                     //For preflight checks and transaction processing, Solana nodes choose which bank state to query based on a commitment requirement set by the client. The commitment describes how finalized a block is at that point in time. When querying the ledger state, it's recommended to use lower levels of commitment to report progress and higher levels to ensure the state will not be rolled back.
	Encoding                       *Encoding           `json:"encoding"`                       //Encoding format for each returned Transaction
	TransactionDetails             *TransactionDetails `json:"transactionDetails"`             //Level of transaction detail to return -- default is "full"
	MaxSupportedTransactionVersion int                 `json:"maxSupportedTransactionVersion"` //If this parameter is omitted, only legacy transactions will be returned, and a block containing any versioned transaction will prompt the error.
	Rewards                        bool                `json:"rewards"`                        //Whether to populate the rewards array. If parameter not provided, the default includes rewards.
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

const LocalRpcUrl = "http://localhost:8899"
const DevnetRpcUrl = "https://api.devnet.solana.com"
const TestnetRpcUrl = "https://api.testnet.solana.com"
const MainnetRpcUrl = "https://api.mainnet-beta.solana.com"

type RpcRequest[T any] struct {
	MethodName string `json:"methodName"`
	Params     T      `json:"params"`
}

type RpcMessage[T any] struct {
	ID      string `json:"id"`
	Jsonrpc string `json:"jsonrpc"`
	Method  string `json:"method"`
	Params  T      `json:"params"`
}

// Required for the JSON-RPC request.
var nextMessageId = 0

func getNextMessageId() string {
	id := nextMessageId
	nextMessageId++
	return fmt.Sprintf("%d", id)
}

func CreateRpcMessage(request RpcRequest[any]) RpcMessage[any] {
	return RpcMessage[any]{
		ID:      getNextMessageId(),
		Jsonrpc: "2.0",
		Method:  request.MethodName,
		Params:  request.Params,
	}
}

type RpcErrorResponsePayload struct {
	Code    int         `json:"code"`
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
}
