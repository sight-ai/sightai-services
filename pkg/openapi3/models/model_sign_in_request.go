package models

type SignInRequest struct {

	Address string `json:"address"`

	Domain string `json:"domain,omitempty"`
}
