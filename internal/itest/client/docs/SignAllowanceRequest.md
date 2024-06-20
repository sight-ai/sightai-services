# SignAllowanceRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ToAccountId** | Pointer to **int64** |  | [optional] 
**Allowance** | **string** |  | 
**Version** | Pointer to **int64** |  | [optional] 

## Methods

### NewSignAllowanceRequest

`func NewSignAllowanceRequest(allowance string, ) *SignAllowanceRequest`

NewSignAllowanceRequest instantiates a new SignAllowanceRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSignAllowanceRequestWithDefaults

`func NewSignAllowanceRequestWithDefaults() *SignAllowanceRequest`

NewSignAllowanceRequestWithDefaults instantiates a new SignAllowanceRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetToAccountId

`func (o *SignAllowanceRequest) GetToAccountId() int64`

GetToAccountId returns the ToAccountId field if non-nil, zero value otherwise.

### GetToAccountIdOk

`func (o *SignAllowanceRequest) GetToAccountIdOk() (*int64, bool)`

GetToAccountIdOk returns a tuple with the ToAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToAccountId

`func (o *SignAllowanceRequest) SetToAccountId(v int64)`

SetToAccountId sets ToAccountId field to given value.

### HasToAccountId

`func (o *SignAllowanceRequest) HasToAccountId() bool`

HasToAccountId returns a boolean if a field has been set.

### GetAllowance

`func (o *SignAllowanceRequest) GetAllowance() string`

GetAllowance returns the Allowance field if non-nil, zero value otherwise.

### GetAllowanceOk

`func (o *SignAllowanceRequest) GetAllowanceOk() (*string, bool)`

GetAllowanceOk returns a tuple with the Allowance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowance

`func (o *SignAllowanceRequest) SetAllowance(v string)`

SetAllowance sets Allowance field to given value.


### GetVersion

`func (o *SignAllowanceRequest) GetVersion() int64`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *SignAllowanceRequest) GetVersionOk() (*int64, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *SignAllowanceRequest) SetVersion(v int64)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *SignAllowanceRequest) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


