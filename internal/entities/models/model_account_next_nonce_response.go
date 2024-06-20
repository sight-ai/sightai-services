package models

type AccountNextNonceResponse struct {

	// the next available nonce
	NextNonce int32 `json:"next_nonce,omitempty"`
}
