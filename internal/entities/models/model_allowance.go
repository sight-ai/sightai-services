package models

type Allowance struct {

	Id int64 `json:"id,omitempty"`

	FromAccount int64 `json:"from_account,omitempty"`

	ToAccount int64 `json:"to_account,omitempty"`

	Allowance string `json:"allowance,omitempty"`

	Version int64 `json:"version,omitempty"`
}
