package transactions

import (
	"errors"
	"solana"

	"github.com/mr-tron/base58"
)

func ParseTransactionData(data []byte) (solana.RawTransaction, error) {
	signatures, messageData, err := getSignatures(data)
	if err != nil {
		return solana.RawTransaction{}, err
	}
	if len(messageData) < 3 {
		return solana.RawTransaction{}, errors.New("not enough data to read message header")
	}
	numRequiredSignatures, numReadonlySignedAccounts, numReadonlyUnsignedAccounts, messageData, err := readMessageHeader(messageData)
	if err != nil {
		return solana.RawTransaction{}, err
	}

	accounts, messageData, err := getAccounts(messageData)
	if err != nil {
		return solana.RawTransaction{}, err
	}

	recentBlockhash, instructionsData, err := getRecentBlockhash(messageData)
	if err != nil {
		return solana.RawTransaction{}, err
	}

	instructions, err := parseInstructions(instructionsData)
	if err != nil {
		return solana.RawTransaction{}, err
	}

	return solana.RawTransaction{
		Signatures: signatures,
		Message: solana.RawTransactionMessage{
			AccountKeys: accounts,
			Header: solana.TransactionHeader{
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

func getAccounts(data []byte) ([]solana.Pubkey, []byte, error) {
	if len(data) < 1 {
		return nil, nil, errors.New("not enough data to read number of accounts")
	}

	totalNumAccounts := data[0]
	data = data[1:]
	if len(data) < int(totalNumAccounts)*32 {
		return nil, nil, errors.New("not enough data to read accounts")
	}

	accounts := make([]solana.Pubkey, totalNumAccounts)
	for i := 0; i < int(totalNumAccounts); i++ {
		accountData := data[i*32 : (i+1)*32]
		pubkey, err := solana.ParsePubkeyBytes(accountData)
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

func parseInstructions(data []byte) ([]solana.RawInstruction, error) {
	if len(data) < 1 {
		return nil, errors.New("not enough data to read number of instructions")
	}

	numInstructions := int(data[0])
	data = data[1:]
	instructions := make([]solana.RawInstruction, numInstructions)
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

func parseInstruction(data []byte) (solana.RawInstruction, []byte, error) {
	if len(data) < 1 {
		return solana.RawInstruction{}, nil, errors.New("not enough data to read program id index")
	}
	programIDIndex := int(data[0])
	data = data[1:]

	if len(data) < 1 {
		return solana.RawInstruction{}, nil, errors.New("not enough data to read accounts")
	}
	numAccounts := int(data[0])
	data = data[1:]

	if len(data) < numAccounts {
		return solana.RawInstruction{}, nil, errors.New("not enough data to read accounts")
	}
	accounts := make([]int, numAccounts)
	for i := 0; i < numAccounts; i++ {
		accounts[i] = int(data[i])
	}
	data = data[numAccounts:]

	if len(data) < 1 {
		return solana.RawInstruction{}, nil, errors.New("not enough data to read data length")
	}
	dataLength := int(data[0])
	data = data[1:]

	if len(data) < dataLength {
		return solana.RawInstruction{}, nil, errors.New("not enough data to read data")
	}
	instructionData := data[:dataLength]
	data = data[dataLength:]

	return solana.RawInstruction{
		ProgramIDIndex: programIDIndex,
		Accounts:       accounts,
		Data:           instructionData,
	}, data, nil
}

func ToData(transaction solana.RawTransaction) ([]byte, error) {
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

func getAccountsData(accounts []solana.Pubkey) ([]byte, error) {
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

func getInstructionsData(instructions []solana.RawInstruction) ([]byte, error) {
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

func getInstructionData(instruction solana.RawInstruction) ([]byte, error) {
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
