package main

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"log"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

// Replace these variables with your own contract details
const (
	infuraURL       = "https://mainnet.infura.io/v3/YOUR_INFURA_PROJECT_ID"
	contractAddress = "0xYourContractAddress"
	privateKey      = "0xYourPrivateKey" // Private key of the sender
	abiJSON         = `[{"constant":false,"inputs":[],"name":"yourFunction","outputs":[],"payable":false,"stateMutability":"nonpayable","type":"function"}]`
)

func main() {
	// Connect to the Ethereum network
	client, err := ethclient.Dial(infuraURL)
	if err != nil {
		log.Fatalf("Failed to connect to the Ethereum client: %v", err)
	}

	// Load the private key of the sender
	privateKey, err := crypto.HexToECDSA(privateKey[2:])
	if err != nil {
		log.Fatalf("Failed to load private key: %v", err)
	}

	// Get the public address of the sender
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatalf("Failed to get public key")
	}
	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)

	// Load the contract ABI
	contractABI, err := abi.JSON(strings.NewReader(abiJSON))
	if err != nil {
		log.Fatalf("Failed to load contract ABI: %v", err)
	}

	// Create an authorized transactor
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatalf("Failed to get nonce: %v", err)
	}

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatalf("Failed to get gas price: %v", err)
	}

	chainID, err := client.NetworkID(context.Background())
	if err != nil {
		log.Fatalf("Failed to get chain ID: %v", err)
	}

	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err != nil {
		log.Fatalf("Failed to create transactor: %v", err)
	}
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)      // Amount of ETH to send (0 in this case)
	auth.GasLimit = uint64(3000000) // Gas limit
	auth.GasPrice = gasPrice

	// Load the contract instance
	contractAddress := common.HexToAddress(contractAddress)
	parsed, err := abi.JSON(strings.NewReader(string(abiJSON)))
	if err != nil {
		log.Fatalf("Failed to parse ABI: %v", err)
	}
	contract := bind.NewBoundContract(contractAddress, parsed, client, client, client)

	// Invoke the contract function
	tx, err := contract.Transact(auth, "yourFunction")
	if err != nil {
		log.Fatalf("Failed to invoke contract function: %v", err)
	}

	fmt.Printf("Transaction submitted! Hash: %s\n", tx.Hash().Hex())
}
