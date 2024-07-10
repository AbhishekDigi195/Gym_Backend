package store

import (
	"fmt"
	"gym/pkg/chain"
	"log"
	"os"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

var Session *StorageSession

func Init() {
	address := os.Getenv("TOKEN_ADDRESS")
	if address == "" {
		log.Fatal("Environment variable YOUR_ENV_VARIABLE_NAME not set")
	}

	// Convert the address to common.Address
	commonAddress := common.HexToAddress(address)
	vest, err := NewStorage(commonAddress, chain.Client)
	if err != nil {
		log.Fatal(err)
	}
	privateKeyHex := os.Getenv("PRIVATE_KEY")
	if privateKeyHex == "" {
		log.Fatal("Environment variable YOUR_PRIVATE_KEY_ENV_VARIABLE not set")
	}

	// Convert the private key from hex string to *ecdsa.PrivateKey
	privateKey, err := crypto.HexToECDSA(privateKeyHex)
	if err != nil {
		log.Fatal(err)
	}
	txtor, err := bind.NewKeyedTransactorWithChainID(privateKey, chain.ID)
	if err != nil {
		log.Fatal(err)
	}
	Session = &StorageSession{
		Contract:     vest,
		TransactOpts: *txtor,
		CallOpts: bind.CallOpts{
			Pending: true,
		},
	}
	fmt.Println("sess", Session)
}
