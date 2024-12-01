package solana

import (
	"errors"
	"slices"

	"github.com/mr-tron/base58"
)

type Transaction struct {
	Signatures []string `json:"signatures"`
	Message    Message  `json:"message"`
}

type Message struct {
	Instructions    []Instruction `json:"instructions"`
	RecentBlockhash string        `json:"recentBlockhash"`
}

type Instruction struct {
	Accounts  []InstructionAccount `json:"accounts"`
	Data      []byte               `json:"data"`
	ProgramID Pubkey               `json:"programId"`
}

type InstructionAccount struct {
	Pubkey   Pubkey `json:"pubkey"`   //Public key of the account
	Signer   bool   `json:"signer"`   //Boolean indicating if the account is a signer
	Writable bool   `json:"writable"` //Boolean indicating if the account is writable
}

type RawTransaction struct {
	Message    RawMessage `json:"message"`
	Signatures []string   `json:"signatures"`
}

type RawMessage struct {
	AccountKeys     []Pubkey         `json:"accountKeys"`
	Header          MessageHeader    `json:"header"`
	Instructions    []RawInstruction `json:"instructions"`
	RecentBlockhash string           `json:"recentBlockhash"`
}

type MessageHeader struct {
	NumReadonlySignedAccounts   int `json:"numReadonlySignedAccounts"`
	NumReadonlyUnsignedAccounts int `json:"numReadonlyUnsignedAccounts"`
	NumRequiredSignatures       int `json:"numRequiredSignatures"`
}

type RawInstruction struct {
	Accounts       []int  `json:"accounts"`
	Data           []byte `json:"data"`
	ProgramIDIndex int    `json:"programIdIndex"`
}

type TransactionWithMeta struct {
	Meta        *TransactionMeta `json:"meta"`
	Version     *int             `json:"version"` //Transaction version. Undefined if maxSupportedTransactionVersion is not set in request params. --note can also be "legacy"
	Transaction Transaction      `json:"transaction"`
}

type RawTransactionWithMeta struct {
	Meta        *TransactionMeta `json:"meta"`
	Version     *int             `json:"version"` //Transaction version. Undefined if maxSupportedTransactionVersion is not set in request params. --note can also be "legacy"
	Transaction RawTransaction   `json:"transaction"`
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

// The Solana runtime records the cross-program instructions that are invoked during transaction processing and makes these available for greater transparency of what was executed on-chain per transaction instruction. Invoked instructions are grouped by the originating transaction instruction and are listed in order of processing.
type InnerInstructions struct {
	Index        int              `json:"index"`        //Index of the transaction instruction from which the inner instruction(s) originated
	Instructions []RawInstruction `json:"instructions"` //Ordered list of inner program instructions that were invoked during a single transaction instruction.
}

type LoadedAddresses struct {
	Writable []string `json:"writable"` //Writable account addresses
	Readonly []string `json:"readonly"` //Readonly account addresses
}

type TransactionReturnData struct {
	ProgramID string `json:"programId"` //The program that generated the return data, as base-58 encoded Pubkey
	Data      string `json:"data"`      //the return data itself, as base-64 encoded binary data
}

func ParseTransactionData(data []byte) (RawTransaction, error) {
	signatures, messageData, err := getSignatures(data)
	if err != nil {
		return RawTransaction{}, err
	}
	if len(messageData) < 3 {
		return RawTransaction{}, errors.New("not enough data to read message header")
	}
	numRequiredSignatures, numReadonlySignedAccounts, numReadonlyUnsignedAccounts, messageData, err := readMessageHeader(messageData)
	if err != nil {
		return RawTransaction{}, err
	}

	accounts, messageData, err := getAccounts(messageData)
	if err != nil {
		return RawTransaction{}, err
	}

	recentBlockhash, instructionsData, err := getRecentBlockhash(messageData)
	if err != nil {
		return RawTransaction{}, err
	}

	instructions, err := parseInstructions(instructionsData)
	if err != nil {
		return RawTransaction{}, err
	}

	return RawTransaction{
		Signatures: signatures,
		Message: RawMessage{
			AccountKeys: accounts,
			Header: MessageHeader{
				NumRequiredSignatures:       numRequiredSignatures,
				NumReadonlySignedAccounts:   numReadonlySignedAccounts,
				NumReadonlyUnsignedAccounts: numReadonlyUnsignedAccounts,
			},
			Instructions:    instructions,
			RecentBlockhash: recentBlockhash,
		},
	}, nil
}

func getSignatures(data []byte) ([]string, []byte, error) {
	if len(data) < 1 {
		return nil, nil, errors.New("not enough data to read number of signatures")
	}

	numSignatures := int(data[0])
	signatures := make([]string, numSignatures)
	for i := 0; i < numSignatures; i++ {
		if len(data) < 1+(i+1)*64 {
			return nil, nil, errors.New("not enough data to read signature")
		}
		signatureData := data[1+i*64 : 1+(i+1)*64]
		signatures[i] = base58.Encode(signatureData)
	}
	remainingData := data[1+numSignatures*64:]
	return signatures, remainingData, nil
}

func readMessageHeader(data []byte) (int, int, int, []byte, error) {
	if len(data) < 3 {
		return 0, 0, 0, nil, errors.New("not enough data to read message header")
	}
	numRequiredSignatures := int(data[0])
	numReadonlySignedAccounts := int(data[1])
	numReadonlyUnsignedAccounts := int(data[2])
	messageData := data[3:]
	return numRequiredSignatures, numReadonlySignedAccounts, numReadonlyUnsignedAccounts, messageData, nil
}

func getAccounts(data []byte) ([]Pubkey, []byte, error) {
	if len(data) < 1 {
		return nil, nil, errors.New("not enough data to read number of accounts")
	}

	totalNumAccounts := data[0]
	data = data[1:]
	if len(data) < int(totalNumAccounts)*32 {
		return nil, nil, errors.New("not enough data to read accounts")
	}

	accounts := make([]Pubkey, totalNumAccounts)
	for i := 0; i < int(totalNumAccounts); i++ {
		accountData := data[i*32 : (i+1)*32]
		pubkey, err := ParsePubkeyBytes(accountData)
		if err != nil {
			return nil, nil, err
		}
		accounts[i] = pubkey
	}
	remainingData := data[int(totalNumAccounts)*32:]
	return accounts, remainingData, nil
}

func getRecentBlockhash(data []byte) (string, []byte, error) {
	if len(data) < 32 {
		return "", nil, errors.New("not enough data to read recent blockhash")
	}
	blockhashData := data[:32]
	blockhash := base58.Encode(blockhashData)
	remainingData := data[32:]
	return blockhash, remainingData, nil
}

func parseInstructions(data []byte) ([]RawInstruction, error) {
	if len(data) < 1 {
		return nil, errors.New("not enough data to read number of instructions")
	}

	numInstructions := int(data[0])
	data = data[1:]
	instructions := make([]RawInstruction, numInstructions)
	for i := 0; i < numInstructions; i++ {
		instruction, remainingData, err := parseInstruction(data)
		if err != nil {
			return nil, err
		}
		instructions[i] = instruction
		data = remainingData
	}
	return instructions, nil
}

func parseInstruction(data []byte) (RawInstruction, []byte, error) {
	if len(data) < 1 {
		return RawInstruction{}, nil, errors.New("not enough data to read program id index")
	}
	programIDIndex := int(data[0])
	data = data[1:]

	if len(data) < 1 {
		return RawInstruction{}, nil, errors.New("not enough data to read accounts")
	}
	numAccounts := int(data[0])
	data = data[1:]

	if len(data) < numAccounts {
		return RawInstruction{}, nil, errors.New("not enough data to read accounts")
	}
	accounts := make([]int, numAccounts)
	for i := 0; i < numAccounts; i++ {
		accounts[i] = int(data[i])
	}
	data = data[numAccounts:]

	if len(data) < 1 {
		return RawInstruction{}, nil, errors.New("not enough data to read data length")
	}
	dataLength := int(data[0])
	data = data[1:]

	if len(data) < dataLength {
		return RawInstruction{}, nil, errors.New("not enough data to read data")
	}
	instructionData := data[:dataLength]
	data = data[dataLength:]

	return RawInstruction{
		ProgramIDIndex: programIDIndex,
		Accounts:       accounts,
		Data:           instructionData,
	}, data, nil
}

func (transaction RawTransaction) Data() ([]byte, error) {
	signaturesData, err := getSignaturesData(transaction.Signatures)
	if err != nil {
		return nil, err
	}

	messageHeaderData := []byte{
		byte(transaction.Message.Header.NumRequiredSignatures),
		byte(transaction.Message.Header.NumReadonlySignedAccounts),
		byte(transaction.Message.Header.NumReadonlyUnsignedAccounts),
	}

	accountsData, err := getAccountsData(transaction.Message.AccountKeys)
	if err != nil {
		return nil, err
	}

	recentBlockhashData, err := getRecentBlockhashData(transaction.Message.RecentBlockhash)
	if err != nil {
		return nil, err
	}

	instructionsData, err := getInstructionsData(transaction.Message.Instructions)
	if err != nil {
		return nil, err
	}

	return append(append(append(append(signaturesData, messageHeaderData...), accountsData...), recentBlockhashData...), instructionsData...), nil
}

func getSignaturesData(signatures []string) ([]byte, error) {
	signaturesData := []byte{byte(len(signatures))}
	for _, signature := range signatures {
		signatureData, err := base58.Decode(signature)
		if err != nil {
			return nil, err
		}
		signaturesData = append(signaturesData, signatureData...)
	}
	return signaturesData, nil
}

func getAccountsData(accounts []Pubkey) ([]byte, error) {
	accountsData := []byte{byte(len(accounts))}
	for _, account := range accounts {
		accountsData = append(accountsData, account.Bytes()...)
	}
	return accountsData, nil
}

func getRecentBlockhashData(recentBlockhash string) ([]byte, error) {
	blockhashData, err := base58.Decode(recentBlockhash)
	if err != nil {
		return nil, err
	}
	return blockhashData, nil
}

func getInstructionsData(instructions []RawInstruction) ([]byte, error) {
	instructionsData := []byte{byte(len(instructions))}
	for _, instruction := range instructions {
		instructionData, err := getInstructionData(instruction)
		if err != nil {
			return nil, err
		}
		instructionsData = append(instructionsData, instructionData...)
	}
	return instructionsData, nil
}

func getInstructionData(instruction RawInstruction) ([]byte, error) {
	instructionData := []byte{
		byte(instruction.ProgramIDIndex),
		byte(len(instruction.Accounts)),
	}
	for _, account := range instruction.Accounts {
		instructionData = append(instructionData, byte(account))
	}
	instructionData = append(instructionData, byte(len(instruction.Data)))
	instructionData = append(instructionData, instruction.Data...)
	return instructionData, nil
}

func (tx Transaction) RawTransaction() RawTransaction {
	rawTransaction := RawTransaction{
		Signatures: tx.Signatures,
	}

	accountKeys, header := populateAccountKeys(tx.Message)

	rawTransaction.Message = RawMessage{
		AccountKeys:     accountKeys,
		Header:          header,
		Instructions:    getInstructions(tx.Message.Instructions, accountKeys),
		RecentBlockhash: tx.Message.RecentBlockhash,
	}

	return rawTransaction
}

func populateAccountKeys(msg Message) ([]Pubkey, MessageHeader) {
	var accountKeys []Pubkey
	var readOnlySigned int
	var readOnlyUnsigned int
	var signers int

	//First we add the signers + writable accounts
	for _, instruction := range msg.Instructions {
		for _, account := range instruction.Accounts {
			if account.Signer && account.Writable {
				accountKeys = append(accountKeys, account.Pubkey)
				signers++
			}
		}
	}
	//Then we add the signers + readonly accounts
	for _, instruction := range msg.Instructions {
		for _, account := range instruction.Accounts {
			if account.Signer && !account.Writable {
				accountKeys = append(accountKeys, account.Pubkey)
				signers++
				readOnlySigned++
			}
		}
	}
	//Then we add the writable accounts
	for _, instruction := range msg.Instructions {
		for _, account := range instruction.Accounts {
			if !account.Signer && account.Writable {
				accountKeys = append(accountKeys, account.Pubkey)
			}
		}
	}
	//Finally we add the readonly accounts
	for _, instruction := range msg.Instructions {
		for _, account := range instruction.Accounts {
			if !account.Signer && !account.Writable {
				accountKeys = append(accountKeys, account.Pubkey)
				readOnlyUnsigned++
			}
		}
	}
	for _, instruction := range msg.Instructions {
		if !slices.ContainsFunc(accountKeys, func(pubkey Pubkey) bool {
			return pubkey.String() == instruction.ProgramID.String()
		}) {
			accountKeys = append(accountKeys, instruction.ProgramID)
		}
	}
	return accountKeys, MessageHeader{
		NumReadonlySignedAccounts:   readOnlySigned,
		NumReadonlyUnsignedAccounts: readOnlyUnsigned,
		NumRequiredSignatures:       signers,
	}
}

func getInstructions(instructions []Instruction, accountKeys []Pubkey) []RawInstruction {
	rawInstructions := make([]RawInstruction, len(instructions))
	for i, instruction := range instructions {
		rawInstructions[i] = RawInstruction{
			Accounts:       getAccountIndices(instruction.Accounts, accountKeys),
			Data:           instruction.Data,
			ProgramIDIndex: getProgramIndex(instruction.ProgramID, accountKeys),
		}
	}
	return rawInstructions
}

func getAccountIndices(accounts []InstructionAccount, accountKeys []Pubkey) []int {
	indices := make([]int, len(accounts))
	accountKeysStr := make([]string, len(accountKeys))
	for i, key := range accountKeys {
		accountKeysStr[i] = key.String()
	}
	for i, account := range accounts {
		indices[i] = slices.Index(accountKeysStr, account.Pubkey.String())
	}
	return indices
}

func getProgramIndex(programID Pubkey, accountKeys []Pubkey) int {
	accountKeysStr := make([]string, len(accountKeys))
	for i, key := range accountKeys {
		accountKeysStr[i] = key.String()
	}
	return slices.Index(accountKeysStr, programID.String())
}

func (rawTx RawTransaction) Transaction() (Transaction, error) {
	var instructions []Instruction

	instructionAccounts := getInstructionAccounts(rawTx.Message.AccountKeys, rawTx.Message.Header)

	for _, rawInstruction := range rawTx.Message.Instructions {
		if len(rawTx.Message.AccountKeys) <= rawInstruction.ProgramIDIndex {
			return Transaction{}, errors.New("invalid program ID index, not enough account keys")
		}
		instruction := Instruction{Data: rawInstruction.Data,
			ProgramID: rawTx.Message.AccountKeys[rawInstruction.ProgramIDIndex],
			Accounts:  make([]InstructionAccount, len(rawInstruction.Accounts)),
		}
		for i, accountIndex := range rawInstruction.Accounts {
			if len(rawTx.Message.AccountKeys) <= accountIndex {
				return Transaction{}, errors.New("invalid account index, not enough account keys")
			}
			instruction.Accounts[i] = instructionAccounts[accountIndex]
		}

		instructions = append(instructions, instruction)
	}

	return Transaction{
		Signatures: rawTx.Signatures,
		Message: Message{
			Instructions:    instructions,
			RecentBlockhash: rawTx.Message.RecentBlockhash,
		},
	}, nil
}

func getInstructionAccounts(accountKeys []Pubkey, header MessageHeader) []InstructionAccount {
	accounts := make([]InstructionAccount, len(accountKeys))
	for i, key := range accountKeys {
		accounts[i] = InstructionAccount{
			Pubkey:   key,
			Signer:   i < header.NumRequiredSignatures,
			Writable: i < header.NumRequiredSignatures-header.NumReadonlySignedAccounts || (i >= header.NumRequiredSignatures && i < len(accountKeys)-header.NumReadonlyUnsignedAccounts),
		}
	}
	return accounts
}
