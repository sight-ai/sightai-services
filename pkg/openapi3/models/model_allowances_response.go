package models

type AllowancesResponse struct {

	FromAllowances []Allowance `json:"from_allowances,omitempty"`

	ToAllowances []Allowance `json:"to_allowances,omitempty"`
}
