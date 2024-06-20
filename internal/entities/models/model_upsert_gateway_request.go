package models

type UpsertGatewayRequest struct {

	Id int64 `json:"id,omitempty"`

	Address string `json:"address"`

	Endpoint string `json:"endpoint,omitempty"`

	Name string `json:"name,omitempty"`

	Deleted bool `json:"deleted,omitempty"`
}
