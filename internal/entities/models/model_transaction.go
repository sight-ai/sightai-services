package models

import (
	"time"
)

type Transaction struct {

	Id int64 `json:"id"`

	CreatedAt time.Time `json:"created_at,omitempty"`

	AccountId int64 `json:"account_id"`

	AvailableDelta string `json:"available_delta,omitempty"`

	HoldDelta string `json:"hold_delta,omitempty"`

	Type string `json:"type,omitempty"`

	Notes string `json:"notes,omitempty"`
}
