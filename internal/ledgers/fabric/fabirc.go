package fabric

import (
	"github.com/hyperledger/fabric-sdk-go/pkg/client/channel"
)

type FabricLedger struct {
	client *channel.Client
}

// NewFabricLedger creates a new FabricLedger
func NewFabricLedger(client *channel.Client) *FabricLedger {
	return &FabricLedger{
		client: client,
	}
}

// AddTransaction adds a transaction to the Fabric ledger
func (f *FabricLedger) AddTransaction(transactionData []byte) (string, error) {
	response, err := f.client.Execute(channel.Request{
		ChaincodeID: "yourChaincodeID",
		Fcn:         "addTransaction",
		Args:        [][]byte{transactionData},
	})
	if err != nil {
		return "", err
	}
	return string(response.TransactionID), nil
}

// GetTransaction retrieves a transaction from the Fabric ledger
func (f *FabricLedger) GetTransaction(transactionID string) ([]byte, error) {
	response, err := f.client.Query(channel.Request{
		ChaincodeID: "yourChaincodeID",
		Fcn:         "getTransaction",
		Args:        [][]byte{[]byte(transactionID)},
	})
	if err != nil {
		return nil, err
	}
	return response.Payload, nil
}
