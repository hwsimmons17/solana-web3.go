package transactions

import (
	"errors"
	"solana"
	"solana/dependencies/keypair"

	"github.com/mr-tron/base58"
)

func ParseTransactionData(data []byte) (solana.Transaction, error) {
	signatures, messageData, err := getSignatures(data)
	if err != nil {
		return solana.Transaction{}, err
	}
	if len(messageData) < 3 {
		return solana.Transaction{}, errors.New("not enough data to read message header")
	}
	numRequiredSignatures, numReadonlySignedAccounts, numReadonlyUnsignedAccounts, messageData, err := readMessageHeader(messageData)
	if err != nil {
		return solana.Transaction{}, err
	}

	accounts, messageData, err := getAccounts(messageData)
	if err != nil {
		return solana.Transaction{}, err
	}

	recentBlockhash, instructionsData, err := getRecentBlockhash(messageData)
	if err != nil {
		return solana.Transaction{}, err
	}

	instructions, err := parseInstructions(instructionsData)
	if err != nil {
		return solana.Transaction{}, err
	}

	return solana.Transaction{
		Signatures: signatures,
		Message: solana.TransactionMessage{
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
		pubkey, err := keypair.ParsePubkeyBytes(accountData)
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

func parseInstructions(data []byte) ([]solana.TransactionInstruction, error) {
	if len(data) < 1 {
		return nil, errors.New("not enough data to read number of instructions")
	}

	numInstructions := int(data[0])
	data = data[1:]
	instructions := make([]solana.TransactionInstruction, numInstructions)
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

func parseInstruction(data []byte) (solana.TransactionInstruction, []byte, error) {
	if len(data) < 1 {
		return solana.TransactionInstruction{}, nil, errors.New("not enough data to read program id index")
	}
	programIDIndex := int(data[0])
	data = data[1:]

	if len(data) < 1 {
		return solana.TransactionInstruction{}, nil, errors.New("not enough data to read accounts")
	}
	numAccounts := int(data[0])
	data = data[1:]

	if len(data) < numAccounts {
		return solana.TransactionInstruction{}, nil, errors.New("not enough data to read accounts")
	}
	accounts := make([]int, numAccounts)
	for i := 0; i < numAccounts; i++ {
		accounts[i] = int(data[i])
	}
	data = data[numAccounts:]

	if len(data) < 1 {
		return solana.TransactionInstruction{}, nil, errors.New("not enough data to read data length")
	}
	dataLength := int(data[0])
	data = data[1:]

	if len(data) < dataLength {
		return solana.TransactionInstruction{}, nil, errors.New("not enough data to read data")
	}
	instructionData := data[:dataLength]
	data = data[dataLength:]

	return solana.TransactionInstruction{
		ProgramIDIndex: programIDIndex,
		Accounts:       accounts,
		Data:           instructionData,
	}, data, nil
}
