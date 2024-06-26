/*
SIGHTAI-SERVICES

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"time"
)

// CreateReceiptRequest struct for CreateReceiptRequest
type CreateReceiptRequest struct {
	UserAddress string `json:"user_address"`
	GatewayAddress string `json:"gateway_address"`
	FinishedAt time.Time `json:"finished_at"`
	Cost string `json:"cost"`
	Proof string `json:"proof"`
	TxnId string `json:"txn_id"`
	Status string `json:"status"`
}

// NewCreateReceiptRequest instantiates a new CreateReceiptRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateReceiptRequest(userAddress string, gatewayAddress string, finishedAt time.Time, cost string, proof string, txnId string, status string) *CreateReceiptRequest {
	this := CreateReceiptRequest{}
	this.UserAddress = userAddress
	this.GatewayAddress = gatewayAddress
	this.FinishedAt = finishedAt
	this.Cost = cost
	this.Proof = proof
	this.TxnId = txnId
	this.Status = status
	return &this
}

// NewCreateReceiptRequestWithDefaults instantiates a new CreateReceiptRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateReceiptRequestWithDefaults() *CreateReceiptRequest {
	this := CreateReceiptRequest{}
	return &this
}

// GetUserAddress returns the UserAddress field value
func (o *CreateReceiptRequest) GetUserAddress() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UserAddress
}

// GetUserAddressOk returns a tuple with the UserAddress field value
// and a boolean to check if the value has been set.
func (o *CreateReceiptRequest) GetUserAddressOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UserAddress, true
}

// SetUserAddress sets field value
func (o *CreateReceiptRequest) SetUserAddress(v string) {
	o.UserAddress = v
}

// GetGatewayAddress returns the GatewayAddress field value
func (o *CreateReceiptRequest) GetGatewayAddress() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.GatewayAddress
}

// GetGatewayAddressOk returns a tuple with the GatewayAddress field value
// and a boolean to check if the value has been set.
func (o *CreateReceiptRequest) GetGatewayAddressOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.GatewayAddress, true
}

// SetGatewayAddress sets field value
func (o *CreateReceiptRequest) SetGatewayAddress(v string) {
	o.GatewayAddress = v
}

// GetFinishedAt returns the FinishedAt field value
func (o *CreateReceiptRequest) GetFinishedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.FinishedAt
}

// GetFinishedAtOk returns a tuple with the FinishedAt field value
// and a boolean to check if the value has been set.
func (o *CreateReceiptRequest) GetFinishedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FinishedAt, true
}

// SetFinishedAt sets field value
func (o *CreateReceiptRequest) SetFinishedAt(v time.Time) {
	o.FinishedAt = v
}

// GetCost returns the Cost field value
func (o *CreateReceiptRequest) GetCost() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Cost
}

// GetCostOk returns a tuple with the Cost field value
// and a boolean to check if the value has been set.
func (o *CreateReceiptRequest) GetCostOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Cost, true
}

// SetCost sets field value
func (o *CreateReceiptRequest) SetCost(v string) {
	o.Cost = v
}

// GetProof returns the Proof field value
func (o *CreateReceiptRequest) GetProof() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Proof
}

// GetProofOk returns a tuple with the Proof field value
// and a boolean to check if the value has been set.
func (o *CreateReceiptRequest) GetProofOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Proof, true
}

// SetProof sets field value
func (o *CreateReceiptRequest) SetProof(v string) {
	o.Proof = v
}

// GetTxnId returns the TxnId field value
func (o *CreateReceiptRequest) GetTxnId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TxnId
}

// GetTxnIdOk returns a tuple with the TxnId field value
// and a boolean to check if the value has been set.
func (o *CreateReceiptRequest) GetTxnIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TxnId, true
}

// SetTxnId sets field value
func (o *CreateReceiptRequest) SetTxnId(v string) {
	o.TxnId = v
}

// GetStatus returns the Status field value
func (o *CreateReceiptRequest) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *CreateReceiptRequest) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *CreateReceiptRequest) SetStatus(v string) {
	o.Status = v
}

func (o CreateReceiptRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["user_address"] = o.UserAddress
	}
	if true {
		toSerialize["gateway_address"] = o.GatewayAddress
	}
	if true {
		toSerialize["finished_at"] = o.FinishedAt
	}
	if true {
		toSerialize["cost"] = o.Cost
	}
	if true {
		toSerialize["proof"] = o.Proof
	}
	if true {
		toSerialize["txn_id"] = o.TxnId
	}
	if true {
		toSerialize["status"] = o.Status
	}
	return json.Marshal(toSerialize)
}

type NullableCreateReceiptRequest struct {
	value *CreateReceiptRequest
	isSet bool
}

func (v NullableCreateReceiptRequest) Get() *CreateReceiptRequest {
	return v.value
}

func (v *NullableCreateReceiptRequest) Set(val *CreateReceiptRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateReceiptRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateReceiptRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateReceiptRequest(val *CreateReceiptRequest) *NullableCreateReceiptRequest {
	return &NullableCreateReceiptRequest{value: val, isSet: true}
}

func (v NullableCreateReceiptRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateReceiptRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


