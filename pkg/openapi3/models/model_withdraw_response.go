package models

type WithdrawResponse struct {

	Sig string `json:"sig,omitempty"`

	Nonce int32 `json:"nonce,omitempty"`

	Amount string `json:"amount,omitempty"`
}
