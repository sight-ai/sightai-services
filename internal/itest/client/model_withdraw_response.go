/*
SIGHTAI-SERVICES

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// WithdrawResponse struct for WithdrawResponse
type WithdrawResponse struct {
	Sig *string `json:"sig,omitempty"`
	Nonce *int32 `json:"nonce,omitempty"`
	Amount *string `json:"amount,omitempty"`
}

// NewWithdrawResponse instantiates a new WithdrawResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWithdrawResponse() *WithdrawResponse {
	this := WithdrawResponse{}
	return &this
}

// NewWithdrawResponseWithDefaults instantiates a new WithdrawResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWithdrawResponseWithDefaults() *WithdrawResponse {
	this := WithdrawResponse{}
	return &this
}

// GetSig returns the Sig field value if set, zero value otherwise.
func (o *WithdrawResponse) GetSig() string {
	if o == nil || o.Sig == nil {
		var ret string
		return ret
	}
	return *o.Sig
}

// GetSigOk returns a tuple with the Sig field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WithdrawResponse) GetSigOk() (*string, bool) {
	if o == nil || o.Sig == nil {
		return nil, false
	}
	return o.Sig, true
}

// HasSig returns a boolean if a field has been set.
func (o *WithdrawResponse) HasSig() bool {
	if o != nil && o.Sig != nil {
		return true
	}

	return false
}

// SetSig gets a reference to the given string and assigns it to the Sig field.
func (o *WithdrawResponse) SetSig(v string) {
	o.Sig = &v
}

// GetNonce returns the Nonce field value if set, zero value otherwise.
func (o *WithdrawResponse) GetNonce() int32 {
	if o == nil || o.Nonce == nil {
		var ret int32
		return ret
	}
	return *o.Nonce
}

// GetNonceOk returns a tuple with the Nonce field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WithdrawResponse) GetNonceOk() (*int32, bool) {
	if o == nil || o.Nonce == nil {
		return nil, false
	}
	return o.Nonce, true
}

// HasNonce returns a boolean if a field has been set.
func (o *WithdrawResponse) HasNonce() bool {
	if o != nil && o.Nonce != nil {
		return true
	}

	return false
}

// SetNonce gets a reference to the given int32 and assigns it to the Nonce field.
func (o *WithdrawResponse) SetNonce(v int32) {
	o.Nonce = &v
}

// GetAmount returns the Amount field value if set, zero value otherwise.
func (o *WithdrawResponse) GetAmount() string {
	if o == nil || o.Amount == nil {
		var ret string
		return ret
	}
	return *o.Amount
}

// GetAmountOk returns a tuple with the Amount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WithdrawResponse) GetAmountOk() (*string, bool) {
	if o == nil || o.Amount == nil {
		return nil, false
	}
	return o.Amount, true
}

// HasAmount returns a boolean if a field has been set.
func (o *WithdrawResponse) HasAmount() bool {
	if o != nil && o.Amount != nil {
		return true
	}

	return false
}

// SetAmount gets a reference to the given string and assigns it to the Amount field.
func (o *WithdrawResponse) SetAmount(v string) {
	o.Amount = &v
}

func (o WithdrawResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Sig != nil {
		toSerialize["sig"] = o.Sig
	}
	if o.Nonce != nil {
		toSerialize["nonce"] = o.Nonce
	}
	if o.Amount != nil {
		toSerialize["amount"] = o.Amount
	}
	return json.Marshal(toSerialize)
}

type NullableWithdrawResponse struct {
	value *WithdrawResponse
	isSet bool
}

func (v NullableWithdrawResponse) Get() *WithdrawResponse {
	return v.value
}

func (v *NullableWithdrawResponse) Set(val *WithdrawResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableWithdrawResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableWithdrawResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWithdrawResponse(val *WithdrawResponse) *NullableWithdrawResponse {
	return &NullableWithdrawResponse{value: val, isSet: true}
}

func (v NullableWithdrawResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWithdrawResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

