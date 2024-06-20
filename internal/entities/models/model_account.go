package models

import (
	"time"
)

type Account struct {

	Id int64 `json:"id"`

	CreatedAt time.Time `json:"created_at,omitempty"`

	Address string `json:"address"`

	Hold string `json:"hold,omitempty"`

	Available string `json:"available,omitempty"`

	Nonce int64 `json:"nonce,omitempty"`

	Role string `json:"role"`
}
