package ethereum

import (
    "context"
    "crypto/ecdsa"
    "log"
    "math/big"

    "github.com/ethereum/go-ethereum/accounts/abi/bind"
    "github.com/ethereum/go-ethereum/common"
    "github.com/ethereum/go-ethereum/crypto"
    "github.com/ethereum/go-ethereum/ethclient"
    "github.com/your_project/ledger"
)

type EthereumLedger struct {
    client        *ethclient.Client
    contractAddr  common.Address
    privateKey    *ecdsa.PrivateKey
    auth          *bind.TransactOpts
    contract      *bind.BoundContract
}

// NewEthereumLedger creates a new EthereumLedger
func NewEthereumLedger(client *ethclient.Client, contractAddr string, privateKeyHex string, abiJSON string) (*EthereumLedger, error) {
    privateKey, err := crypto.HexToECDSA(privateKeyHex)
    if err != nil {
        return nil, err
    }

    publicKey := privateKey.Public()
    publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
    if !ok {
        return nil, err
    }

    fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
    chainID, err := client.NetworkID(context.Background())
    if err != nil {
        return nil, err
    }

    auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
    if err != nil {
        return nil, err
    }

    contract := bind.NewBoundContract(common.HexToAddress(contractAddr), abiJSON, client, client, client)

    return &EthereumLedger{
        client:       client,
        contractAddr: common.HexToAddress(contractAddr),
        privateKey:   privateKey,
        auth:         auth,
        contract:     contract,
    }, nil
}

// AddTransaction adds a transaction to the Ethereum ledger
func (e *EthereumLedger) AddTransaction(transactionData []byte) (string, error) {
    tx, err := e.contract.Transact(e.auth, "addTransaction", transactionData)
    if err != nil {
        return "", err
    }
    return tx.Hash().Hex(), nil
}

// GetTransaction retrieves a transaction from the Ethereum ledger
func (e *EthereumLedger) GetTransaction(transactionID string) ([]byte, error) {
    var result []byte
    err := e.contract.Call(&bind.CallOpts{}, &result, "getTransaction", common.HexToHash(transactionID))
    if err != nil {
        return nil, err
    }
    return result, nil
}
