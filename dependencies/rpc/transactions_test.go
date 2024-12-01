package rpc

import (
	"solana"
	"testing"
)

func TestGetFeeForMessage(t *testing.T) {
	t.Skip("Skipping test that requires network access")
	client := NewRpcClient(solana.RpcEndpointMainnetBeta)
	msg := []byte{1, 0, 1, 3, 161, 68, 106, 177, 149, 255, 18, 122, 114, 25, 238, 115, 76, 66, 62, 224, 224, 53, 252, 245, 200, 239, 51, 242, 63, 210, 180, 118, 128, 206, 13, 23, 28, 53, 216, 230, 107, 42, 37, 88, 63, 140, 189, 215, 44, 74, 139, 105, 1, 234, 186, 80, 209, 29, 30, 151, 84, 217, 151, 212, 187, 30, 134, 92, 7, 97, 72, 29, 53, 116, 116, 187, 124, 77, 118, 36, 235, 211, 189, 179, 216, 53, 94, 115, 209, 16, 67, 252, 13, 163, 83, 128, 0, 0, 0, 0, 67, 38, 70, 223, 206, 169, 192, 25, 253, 85, 235, 130, 87, 163, 25, 137, 217, 2, 167, 14, 240, 241, 33, 120, 65, 63, 87, 91, 170, 240, 200, 39, 1, 2, 2, 1, 0, 116, 12, 0, 0, 0, 115, 73, 126, 20, 0, 0, 0, 0, 31, 1, 31, 1, 30, 1, 29, 1, 28, 1, 27, 1, 26, 1, 25, 1, 24, 1, 23, 1, 22, 1, 21, 1, 20, 1, 19, 1, 18, 1, 17, 1, 16, 1, 15, 1, 14, 1, 13, 1, 12, 1, 11, 1, 10, 1, 9, 1, 8, 1, 7, 1, 6, 1, 5, 1, 4, 1, 3, 1, 2, 1, 1, 23, 67, 38, 224, 121, 226, 210, 146, 191, 248, 41, 122, 192, 109, 199, 28, 87, 225, 33, 77, 211, 76, 139, 225, 161, 103, 27, 73, 101, 59, 71, 35, 1, 197, 178, 75, 103, 0, 0, 0, 0}
	fee, err := client.GetFeeForMessage(msg)
	if err != nil {
		t.Fatal(err)
	}
	t.Fatal(fee)
}

func TestGetLatestBlockhash(t *testing.T) {
	t.Skip("Skipping test that requires network access")
	client := NewRpcClient(solana.RpcEndpointMainnetBeta)
	blockhash, err := client.GetLatestBlockhash()
	if err != nil {
		t.Fatal(err)
	}
	t.Fatal(blockhash)
}

func TestGetSignatureStatuses(t *testing.T) {
	t.Skip("Skipping test that requires network access")
	client := NewRpcClient(solana.RpcEndpointDevnet)
	signatures, err := client.GetSignatureStatuses([]string{"2WPaFF2Jy51ioWknrfAaraueRLwbJfo5ciEgK9t2ccNvR6LQiT7UarkQTH4EKpXw9TRVce9dMrDiuANYMQzE6sTg"}, solana.GetSignatureStatusesConfig{
		SearchTransactionHistory: true,
	})
	if err != nil {
		t.Fatal(err)
	}
	t.Fatal(*signatures[0])
}

func TestGetSignaturesForAddress(t *testing.T) {
	t.Skip("Skipping test that requires network access")
	client := NewRpcClient(solana.RpcEndpointDevnet)
	signatures, err := client.GetSignaturesForAddress(solana.MustParsePubkey("7Fg8XQBVY4z7gPzecGo7abbHZbHj3iFfGozXsz1VcvKk"))
	if err != nil {
		t.Fatal(err)
	}
	t.Fatal(signatures)
}

func TestGetTransaction(t *testing.T) {
	t.Skip("Skipping test that requires network access")
	client := NewRpcClient(solana.RpcEndpointDevnet)
	tx, err := client.GetTransaction("2WPaFF2Jy51ioWknrfAaraueRLwbJfo5ciEgK9t2ccNvR6LQiT7UarkQTH4EKpXw9TRVce9dMrDiuANYMQzE6sTg")
	if err != nil {
		t.Fatal(err)
	}
	t.Fatalf("%+v", tx)
}

func TestIsBlockhashValid(t *testing.T) {
	t.Skip("Skipping test that requires network access")
	client := NewRpcClient(solana.RpcEndpointMainnetBeta)
	blockhash, err := client.GetLatestBlockhash()
	if err != nil {
		t.Fatal(err)
	}
	valid, err := client.IsBlockhashValid(blockhash.Blockhash)
	if err != nil {
		t.Fatal(err)
	}
	t.Fatal(valid)
}
