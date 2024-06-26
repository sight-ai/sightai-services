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

// SignAllowanceRequest struct for SignAllowanceRequest
type SignAllowanceRequest struct {
	ToAccountId *int64 `json:"to_account_id,omitempty"`
	Allowance string `json:"allowance"`
	Version *int64 `json:"version,omitempty"`
}

// NewSignAllowanceRequest instantiates a new SignAllowanceRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSignAllowanceRequest(allowance string) *SignAllowanceRequest {
	this := SignAllowanceRequest{}
	this.Allowance = allowance
	return &this
}

// NewSignAllowanceRequestWithDefaults instantiates a new SignAllowanceRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSignAllowanceRequestWithDefaults() *SignAllowanceRequest {
	this := SignAllowanceRequest{}
	return &this
}

// GetToAccountId returns the ToAccountId field value if set, zero value otherwise.
func (o *SignAllowanceRequest) GetToAccountId() int64 {
	if o == nil || o.ToAccountId == nil {
		var ret int64
		return ret
	}
	return *o.ToAccountId
}

// GetToAccountIdOk returns a tuple with the ToAccountId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SignAllowanceRequest) GetToAccountIdOk() (*int64, bool) {
	if o == nil || o.ToAccountId == nil {
		return nil, false
	}
	return o.ToAccountId, true
}

// HasToAccountId returns a boolean if a field has been set.
func (o *SignAllowanceRequest) HasToAccountId() bool {
	if o != nil && o.ToAccountId != nil {
		return true
	}

	return false
}

// SetToAccountId gets a reference to the given int64 and assigns it to the ToAccountId field.
func (o *SignAllowanceRequest) SetToAccountId(v int64) {
	o.ToAccountId = &v
}

// GetAllowance returns the Allowance field value
func (o *SignAllowanceRequest) GetAllowance() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Allowance
}

// GetAllowanceOk returns a tuple with the Allowance field value
// and a boolean to check if the value has been set.
func (o *SignAllowanceRequest) GetAllowanceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Allowance, true
}

// SetAllowance sets field value
func (o *SignAllowanceRequest) SetAllowance(v string) {
	o.Allowance = v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *SignAllowanceRequest) GetVersion() int64 {
	if o == nil || o.Version == nil {
		var ret int64
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SignAllowanceRequest) GetVersionOk() (*int64, bool) {
	if o == nil || o.Version == nil {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *SignAllowanceRequest) HasVersion() bool {
	if o != nil && o.Version != nil {
		return true
	}

	return false
}

// SetVersion gets a reference to the given int64 and assigns it to the Version field.
func (o *SignAllowanceRequest) SetVersion(v int64) {
	o.Version = &v
}

func (o SignAllowanceRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ToAccountId != nil {
		toSerialize["to_account_id"] = o.ToAccountId
	}
	if true {
		toSerialize["allowance"] = o.Allowance
	}
	if o.Version != nil {
		toSerialize["version"] = o.Version
	}
	return json.Marshal(toSerialize)
}

type NullableSignAllowanceRequest struct {
	value *SignAllowanceRequest
	isSet bool
}

func (v NullableSignAllowanceRequest) Get() *SignAllowanceRequest {
	return v.value
}

func (v *NullableSignAllowanceRequest) Set(val *SignAllowanceRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableSignAllowanceRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableSignAllowanceRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSignAllowanceRequest(val *SignAllowanceRequest) *NullableSignAllowanceRequest {
	return &NullableSignAllowanceRequest{value: val, isSet: true}
}

func (v NullableSignAllowanceRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSignAllowanceRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


