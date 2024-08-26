package model

type UserTransaction struct {
	Data      map[string]interface{} `json:"data" binding:"required"`
	Metadata  map[string]interface{} `json:"metadata,omitempty"`
	Signature string                 `json:"signature" binding:"required"`
}

type User struct {
	ID       string                 `json:"id"`
	AssetID  string                 `json:"asset_id"`
	Data     map[string]interface{} `json:"data"`
	Metadata map[string]interface{} `json:"metadata"`
	InputID  string                 `json:"input_id"`
}
