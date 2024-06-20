package models

type SignAllowanceRequest struct {

	ToAccountId int64 `json:"to_account_id,omitempty"`

	Allowance string `json:"allowance"`

	Version int64 `json:"version,omitempty"`
}
