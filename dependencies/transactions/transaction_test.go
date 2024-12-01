package transactions

import (
	"solana"
	"solana/dependencies/keypair"
	"testing"
)

func TestRawTransaction(t *testing.T) {
	tx := solana.Transaction{
		Signatures: []string{"5WUMzKkDaSuLUj3RpbufHi2PLRPgkjkHhsJpLE7Q3a3fMNDkV579zDmWLfMTnw4my5cbHKicRhYDTQsoAidv8nYD"},
		Message: solana.TransactionMessage{
			Instructions: []solana.Instruction{
				{
					Accounts: []solana.InstructionAccount{
						{
							Pubkey:   keypair.MustParsePubkey("2u83Dx5qPV4QnujjJQv8v2SoqG1ixuAxPK5Jwhtkovd1"),
							Signer:   false,
							Writable: false,
						},
						{
							Pubkey:   keypair.MustParsePubkey("BrX9Z85BbmXYMjvvuAWU8imwsAqutVQiDg9uNfTGkzrJ"),
							Signer:   true,
							Writable: true,
						}},
					Data:      []byte{12, 0, 0, 0, 115},
					ProgramID: keypair.MustParsePubkey("Vote111111111111111111111111111111111111111"),
				},
			},
			RecentBlockhash: "5X8Ak8LYQTdoXbDaEYUdBC5dZophA7fNbiSEgMFYd1Qa",
		},
	}

	rawTx := RawTransaction(tx)
	if rawTx.Message.AccountKeys[0].String() != "BrX9Z85BbmXYMjvvuAWU8imwsAqutVQiDg9uNfTGkzrJ" {
		t.Fatal("Unexpected account key")
	}
}

func TestTransaction(t *testing.T) {
	rawTx := solana.RawTransaction{
		Message: solana.RawTransactionMessage{
			AccountKeys: []solana.Pubkey{
				keypair.MustParsePubkey("BrX9Z85BbmXYMjvvuAWU8imwsAqutVQiDg9uNfTGkzrJ"),
				keypair.MustParsePubkey("2u83Dx5qPV4QnujjJQv8v2SoqG1ixuAxPK5Jwhtkovd1"),
				keypair.MustParsePubkey("Vote111111111111111111111111111111111111111"),
			},
			Header: solana.TransactionHeader{
				NumReadonlySignedAccounts:   0,
				NumReadonlyUnsignedAccounts: 1,
				NumRequiredSignatures:       1,
			},
			Instructions: []solana.RawInstruction{
				{
					Accounts:       []int{1, 0},
					Data:           []byte{12, 0, 0, 0, 115},
					ProgramIDIndex: 2,
				},
			},
			RecentBlockhash: "5X8Ak8LYQTdoXbDaEYUdBC5dZophA7fNbiSEgMFYd1Qa",
		},
		Signatures: []string{"5WUMzKkDaSuLUj3RpbufHi2PLRPgkjkHhsJpLE7Q3a3fMNDkV579zDmWLfMTnw4my5cbHKicRhYDTQsoAidv8nYD"},
	}
	tx, err := Transaction(rawTx)
	if err != nil {
		t.Fatal(err)
	}

	if tx.Message.Instructions[0].Accounts[0].Pubkey.String() != "2u83Dx5qPV4QnujjJQv8v2SoqG1ixuAxPK5Jwhtkovd1" {
		t.Fatal("Unexpected account")
	}
	if tx.Message.Instructions[0].Accounts[0].Signer {
		t.Fatal("Unexpected account signer")
	}
	if !tx.Message.Instructions[0].Accounts[0].Writable {
		t.Fatal("Unexpected account writable")
	}
	if !tx.Message.Instructions[0].Accounts[1].Signer {
		t.Fatal("Unexpected account signer")
	}
	if !tx.Message.Instructions[0].Accounts[1].Writable {
		t.Fatal("Unexpected account writable")
	}
	t.Fatalf("%+v", tx)
}
