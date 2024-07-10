package chain

import (
	"context"
	"log"
	"math/big"
	"os"

	"github.com/ethereum/go-ethereum/ethclient"
)

var Client *ethclient.Client
var ID *big.Int

func Init() {
	var err error
	Client, err = ethclient.Dial(os.Getenv(("RPC")))
	if err != nil {
		log.Fatal(err)
	}
	ID, err = Client.ChainID(context.Background())
	if err != nil {
		log.Fatal(err)
	}
}
