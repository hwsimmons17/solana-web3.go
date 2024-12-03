package examples

import (
	"os"
	"solana"
	"solana/rpc"

	"github.com/joho/godotenv"
)

func SendAirdrop() {
	godotenv.Load("../.env")
	rpc := rpc.NewRpcClient(solana.RpcEndpointDevnet)
	keypair, err := solana.NewKeypairFromBase58(os.Getenv("KEYPAIR"))
	if err != nil {
		panic(err)
	}
	client := solana.NewClient(rpc, keypair)

	if _, err := client.RequestAirdrop(keypair.Pubkey, solana.SolInLamports(1)); err != nil {
		panic(err)
	}
}
