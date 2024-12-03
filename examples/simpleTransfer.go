package examples

import (
	"log"
	"os"

	"github.com/hwsimmons17/solana-web3.go"
	"github.com/hwsimmons17/solana-web3.go/rpc"
	"github.com/joho/godotenv"
)

func SimpleTransfer() {
	godotenv.Load("../.env")
	rpc := rpc.NewRpcClient(solana.RpcEndpointDevnet)
	keypair, err := solana.NewKeypairFromBase58(os.Getenv("KEYPAIR"))
	if err != nil {
		panic(err)
	}
	client := solana.NewClient(rpc, keypair)

	transferIx := solana.SystemProgramInstructions().Transfer(keypair.Pubkey, solana.MustParsePubkey("E4GJZbM77LwkUhCzh2jbdmBWSRktsQuz1SRYRujZTAmu"), 500_000_000)
	tx := solana.Transaction{
		Message: solana.Message{
			Instructions: []solana.Instruction{transferIx},
		},
	}

	//SendAndSignTransaction signs the transaction with the Client's default signer and handles getting the recent blockhash
	if txStr, err := client.SendAndSignTransaction(tx); err != nil {
		log.Println(txStr)
		panic(err)
	}
}
