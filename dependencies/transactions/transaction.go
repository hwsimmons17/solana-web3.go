package transactions

import (
	"errors"
	"slices"
	"solana"
)

func RawTransaction(tx solana.Transaction) solana.RawTransaction {
	rawTransaction := solana.RawTransaction{
		Signatures: tx.Signatures,
	}

	accountKeys, header := populateAccountKeys(tx.Message)

	rawTransaction.Message = solana.RawTransactionMessage{
		AccountKeys:     accountKeys,
		Header:          header,
		Instructions:    getInstructions(tx.Message.Instructions, accountKeys),
		RecentBlockhash: tx.Message.RecentBlockhash,
	}

	return rawTransaction
}

func populateAccountKeys(msg solana.TransactionMessage) ([]solana.Pubkey, solana.TransactionHeader) {
	var accountKeys []solana.Pubkey
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
		if !slices.ContainsFunc(accountKeys, func(pubkey solana.Pubkey) bool {
			return pubkey.String() == instruction.ProgramID.String()
		}) {
			accountKeys = append(accountKeys, instruction.ProgramID)
		}
	}
	return accountKeys, solana.TransactionHeader{
		NumReadonlySignedAccounts:   readOnlySigned,
		NumReadonlyUnsignedAccounts: readOnlyUnsigned,
		NumRequiredSignatures:       signers,
	}
}

func getInstructions(instructions []solana.Instruction, accountKeys []solana.Pubkey) []solana.RawInstruction {
	rawInstructions := make([]solana.RawInstruction, len(instructions))
	for i, instruction := range instructions {
		rawInstructions[i] = solana.RawInstruction{
			Accounts:       getAccountIndices(instruction.Accounts, accountKeys),
			Data:           instruction.Data,
			ProgramIDIndex: getProgramIndex(instruction.ProgramID, accountKeys),
		}
	}
	return rawInstructions
}

func getAccountIndices(accounts []solana.InstructionAccount, accountKeys []solana.Pubkey) []int {
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

func getProgramIndex(programID solana.Pubkey, accountKeys []solana.Pubkey) int {
	accountKeysStr := make([]string, len(accountKeys))
	for i, key := range accountKeys {
		accountKeysStr[i] = key.String()
	}
	return slices.Index(accountKeysStr, programID.String())
}

func Transaction(rawTx solana.RawTransaction) (solana.Transaction, error) {
	var instructions []solana.Instruction

	instructionAccounts := getInstructionAccounts(rawTx.Message.AccountKeys, rawTx.Message.Header)

	for _, rawInstruction := range rawTx.Message.Instructions {
		if len(rawTx.Message.AccountKeys) <= rawInstruction.ProgramIDIndex {
			return solana.Transaction{}, errors.New("invalid program ID index, not enough account keys")
		}
		instruction := solana.Instruction{Data: rawInstruction.Data,
			ProgramID: rawTx.Message.AccountKeys[rawInstruction.ProgramIDIndex],
			Accounts:  make([]solana.InstructionAccount, len(rawInstruction.Accounts)),
		}
		for i, accountIndex := range rawInstruction.Accounts {
			if len(rawTx.Message.AccountKeys) <= accountIndex {
				return solana.Transaction{}, errors.New("invalid account index, not enough account keys")
			}
			instruction.Accounts[i] = instructionAccounts[accountIndex]
		}

		instructions = append(instructions, instruction)
	}

	return solana.Transaction{
		Signatures: rawTx.Signatures,
		Message: solana.TransactionMessage{
			Instructions:    instructions,
			RecentBlockhash: rawTx.Message.RecentBlockhash,
		},
	}, nil
}

func getInstructionAccounts(accountKeys []solana.Pubkey, header solana.TransactionHeader) []solana.InstructionAccount {
	accounts := make([]solana.InstructionAccount, len(accountKeys))
	for i, key := range accountKeys {
		accounts[i] = solana.InstructionAccount{
			Pubkey:   key,
			Signer:   i < header.NumRequiredSignatures,
			Writable: i < header.NumRequiredSignatures-header.NumReadonlySignedAccounts || (i >= header.NumRequiredSignatures && i < len(accountKeys)-header.NumReadonlyUnsignedAccounts),
		}
	}
	return accounts
}
