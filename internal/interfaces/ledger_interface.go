package interfaces

// Ledger defines a generic interface for interacting with a ledger
type Ledger interface {
	AddTransaction(transactionData []byte) (string, error) // Adds a transaction to the ledger
	GetTransaction(transactionID string) ([]byte, error)   // Retrieves a transaction from the ledger by ID
}

// type Ledger interface {
// 	CreateUser(data map[string]interface{}, metadata map[string]interface{}, signature string) (string, error)
// 	GetUser(userID string) (map[string]interface{}, error)
// 	CreateContext(data map[string]interface{}, metadata map[string]interface{}, signature string) (string, error)
// 	GetContext(contextID string) (map[string]interface{}, error)
// 	GetAllContexts() ([]map[string]interface{}, error)
// 	GetContextState(contextID string) (map[string]interface{}, error)
// 	GetAllContextStates() ([]map[string]interface{}, error)
// 	CreateTransaction(data map[string]interface{}, metadata map[string]interface{}, signature string) (string, error)
// 	GetTransaction(transactionID string) (map[string]interface{}, error)
// 	GetTransactionsByAssetID(assetID string) ([]map[string]interface{}, error)
// 	GetTransactionsByContextID(contextID string) ([]map[string]interface{}, error)
// 	GetTransactionsByParentID(parentID string) ([]map[string]interface{}, error)
// 	GetTransactionStateByAssetID(assetID string) ([]map[string]interface{}, error)
// 	GetTransactionStateByContextID(contextID string) ([]map[string]interface{}, error)
// 	GetTransactionStateByParentID(parentID string) ([]map[string]interface{}, error)
// 	GetTransactionStateByAllRelatedTo(allRelatedTo string) ([]map[string]interface{}, error)
// }
