package models

import (
	"time"
)

type Gateway struct {

	Id int64 `json:"id,omitempty"`

	AccountId int64 `json:"account_id,omitempty"`

	Address string `json:"address,omitempty"`

	Endpoint string `json:"endpoint,omitempty"`

	Name string `json:"name,omitempty"`

	DeletedAt time.Time `json:"deleted_at,omitempty"`
}
