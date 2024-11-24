package model

type UserTransaction struct {
	Data      map[string]interface{} `json:"Data" binding:"required"`
	Metadata  map[string]interface{} `json:"Metadata,omitempty"`
	Signature string                 `json:"Signature" binding:"required"`
}

type User struct {
	ID       string                 `json:"id"`
	AssetID  string                 `json:"asset_id"`
	Data     map[string]interface{} `json:"data"`
	Metadata map[string]interface{} `json:"metadata"`
	InputID  string                 `json:"input_id"`
}
