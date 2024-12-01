package solana

import (
	"slices"
	"testing"
)

func TestParseTransactionData(t *testing.T) {
	data := []byte{1, 225, 123, 186, 177, 227, 37, 173, 118, 175, 177, 33, 183, 32, 42, 221, 1, 221, 157, 125, 95, 166, 203, 158, 122, 199, 20, 33, 150, 13, 137, 230, 73, 214, 249, 133, 58, 117, 149, 7, 172, 31, 82, 244, 204, 194, 174, 239, 106, 22, 110, 149, 21, 69, 251, 78, 198, 149, 210, 164, 166, 175, 5, 89, 14, 1, 0, 1, 3, 161, 68, 106, 177, 149, 255, 18, 122, 114, 25, 238, 115, 76, 66, 62, 224, 224, 53, 252, 245, 200, 239, 51, 242, 63, 210, 180, 118, 128, 206, 13, 23, 28, 53, 216, 230, 107, 42, 37, 88, 63, 140, 189, 215, 44, 74, 139, 105, 1, 234, 186, 80, 209, 29, 30, 151, 84, 217, 151, 212, 187, 30, 134, 92, 7, 97, 72, 29, 53, 116, 116, 187, 124, 77, 118, 36, 235, 211, 189, 179, 216, 53, 94, 115, 209, 16, 67, 252, 13, 163, 83, 128, 0, 0, 0, 0, 67, 38, 70, 223, 206, 169, 192, 25, 253, 85, 235, 130, 87, 163, 25, 137, 217, 2, 167, 14, 240, 241, 33, 120, 65, 63, 87, 91, 170, 240, 200, 39, 1, 2, 2, 1, 0, 116, 12, 0, 0, 0, 115, 73, 126, 20, 0, 0, 0, 0, 31, 1, 31, 1, 30, 1, 29, 1, 28, 1, 27, 1, 26, 1, 25, 1, 24, 1, 23, 1, 22, 1, 21, 1, 20, 1, 19, 1, 18, 1, 17, 1, 16, 1, 15, 1, 14, 1, 13, 1, 12, 1, 11, 1, 10, 1, 9, 1, 8, 1, 7, 1, 6, 1, 5, 1, 4, 1, 3, 1, 2, 1, 1, 23, 67, 38, 224, 121, 226, 210, 146, 191, 248, 41, 122, 192, 109, 199, 28, 87, 225, 33, 77, 211, 76, 139, 225, 161, 103, 27, 73, 101, 59, 71, 35, 1, 197, 178, 75, 103, 0, 0, 0, 0}

	transaction, err := ParseTransactionData(data)
	if err != nil {
		t.Fatal(err)
	}
	if transaction.Message.RecentBlockhash != "5X8Ak8LYQTdoXbDaEYUdBC5dZophA7fNbiSEgMFYd1Qa" {
		t.Fatal("Unexpected recent blockhash", transaction.Message.RecentBlockhash)
	}
}

func TestToData(t *testing.T) {
	data := []byte{1, 225, 123, 186, 177, 227, 37, 173, 118, 175, 177, 33, 183, 32, 42, 221, 1, 221, 157, 125, 95, 166, 203, 158, 122, 199, 20, 33, 150, 13, 137, 230, 73, 214, 249, 133, 58, 117, 149, 7, 172, 31, 82, 244, 204, 194, 174, 239, 106, 22, 110, 149, 21, 69, 251, 78, 198, 149, 210, 164, 166, 175, 5, 89, 14, 1, 0, 1, 3, 161, 68, 106, 177, 149, 255, 18, 122, 114, 25, 238, 115, 76, 66, 62, 224, 224, 53, 252, 245, 200, 239, 51, 242, 63, 210, 180, 118, 128, 206, 13, 23, 28, 53, 216, 230, 107, 42, 37, 88, 63, 140, 189, 215, 44, 74, 139, 105, 1, 234, 186, 80, 209, 29, 30, 151, 84, 217, 151, 212, 187, 30, 134, 92, 7, 97, 72, 29, 53, 116, 116, 187, 124, 77, 118, 36, 235, 211, 189, 179, 216, 53, 94, 115, 209, 16, 67, 252, 13, 163, 83, 128, 0, 0, 0, 0, 67, 38, 70, 223, 206, 169, 192, 25, 253, 85, 235, 130, 87, 163, 25, 137, 217, 2, 167, 14, 240, 241, 33, 120, 65, 63, 87, 91, 170, 240, 200, 39, 1, 2, 2, 1, 0, 116, 12, 0, 0, 0, 115, 73, 126, 20, 0, 0, 0, 0, 31, 1, 31, 1, 30, 1, 29, 1, 28, 1, 27, 1, 26, 1, 25, 1, 24, 1, 23, 1, 22, 1, 21, 1, 20, 1, 19, 1, 18, 1, 17, 1, 16, 1, 15, 1, 14, 1, 13, 1, 12, 1, 11, 1, 10, 1, 9, 1, 8, 1, 7, 1, 6, 1, 5, 1, 4, 1, 3, 1, 2, 1, 1, 23, 67, 38, 224, 121, 226, 210, 146, 191, 248, 41, 122, 192, 109, 199, 28, 87, 225, 33, 77, 211, 76, 139, 225, 161, 103, 27, 73, 101, 59, 71, 35, 1, 197, 178, 75, 103, 0, 0, 0, 0}

	transaction, err := ParseTransactionData(data)
	if err != nil {
		t.Fatal(err)
	}

	newData, err := transaction.Bytes()
	if err != nil {
		t.Fatal(err)
	}
	if slices.Compare(data, newData) != 0 {
		t.Fatal("Data does not match")
	}
}

func TestRawTransaction(t *testing.T) {
	tx := Transaction{
		Signatures: []string{"5WUMzKkDaSuLUj3RpbufHi2PLRPgkjkHhsJpLE7Q3a3fMNDkV579zDmWLfMTnw4my5cbHKicRhYDTQsoAidv8nYD"},
		Message: Message{
			Instructions: []Instruction{
				{
					Accounts: []AccountMeta{
						{
							Pubkey:   MustParsePubkey("2u83Dx5qPV4QnujjJQv8v2SoqG1ixuAxPK5Jwhtkovd1"),
							Signer:   false,
							Writable: false,
						},
						{
							Pubkey:   MustParsePubkey("BrX9Z85BbmXYMjvvuAWU8imwsAqutVQiDg9uNfTGkzrJ"),
							Signer:   true,
							Writable: true,
						}},
					Data:      []byte{12, 0, 0, 0, 115},
					ProgramID: MustParsePubkey("Vote111111111111111111111111111111111111111"),
				},
			},
			RecentBlockhash: "5X8Ak8LYQTdoXbDaEYUdBC5dZophA7fNbiSEgMFYd1Qa",
		},
	}

	rawTx := tx.Serialize()
	if rawTx.Message.AccountKeys[0].String() != "BrX9Z85BbmXYMjvvuAWU8imwsAqutVQiDg9uNfTGkzrJ" {
		t.Fatal("Unexpected account key")
	}
}

func TestTransaction(t *testing.T) {
	rawTx := RawTransaction{
		Message: RawMessage{
			AccountKeys: []Pubkey{
				MustParsePubkey("BrX9Z85BbmXYMjvvuAWU8imwsAqutVQiDg9uNfTGkzrJ"),
				MustParsePubkey("2u83Dx5qPV4QnujjJQv8v2SoqG1ixuAxPK5Jwhtkovd1"),
				MustParsePubkey("Vote111111111111111111111111111111111111111"),
			},
			Header: MessageHeader{
				NumReadonlySignedAccounts:   0,
				NumReadonlyUnsignedAccounts: 1,
				NumRequiredSignatures:       1,
			},
			Instructions: []RawInstruction{
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
	tx, err := rawTx.Transaction()
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
}

func TestRawTxToTransaction(t *testing.T) {
	rawTx := RawTransaction{
		Message: RawMessage{
			AccountKeys: []Pubkey{
				MustParsePubkey("GR16g49y2fEjRQD612ryaXjNomRF2TCWoiMgspKtXqya"),
				MustParsePubkey("uaRGm8MX21msZNqwW3AQWYBXheLPcheK6CZmdx788uN"),
				MustParsePubkey("7Fg8XQBVY4z7gPzecGo7abbHZbHj3iFfGozXsz1VcvKk"),
				MustParsePubkey("TokenkegQfeZyiNwAJbNbGKPFXCWuBvf9Ss623VQ5DA"),
				MustParsePubkey("BLrD8HqBy4vKNvkb28Bijg4y6s8tE49jyVFbfZnmesjY")},
			Header: MessageHeader{NumReadonlySignedAccounts: 0, NumReadonlyUnsignedAccounts: 2, NumRequiredSignatures: 1},
			Instructions: []RawInstruction{{
				Accounts:       []int{2, 4, 1, 0},
				Data:           []byte{12, 0, 16, 141, 190, 28, 0, 0, 0, 9},
				ProgramIDIndex: 3,
			}},
			RecentBlockhash: "FLwNEQozzqBBaeHf17JA83PUi3GPEJZkneDm548wMUSj"},
		Signatures: []string{"2WPaFF2Jy51ioWknrfAaraueRLwbJfo5ciEgK9t2ccNvR6LQiT7UarkQTH4EKpXw9TRVce9dMrDiuANYMQzE6sTg"}}

	tx, err := rawTx.Transaction()
	if err != nil {
		t.Fatal(err)
	}

	t.Fatalf("%+v", tx)
}
