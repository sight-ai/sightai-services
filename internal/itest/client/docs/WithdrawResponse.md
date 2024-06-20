# WithdrawResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Sig** | Pointer to **string** |  | [optional] 
**Nonce** | Pointer to **int32** |  | [optional] 
**Amount** | Pointer to **string** |  | [optional] 

## Methods

### NewWithdrawResponse

`func NewWithdrawResponse() *WithdrawResponse`

NewWithdrawResponse instantiates a new WithdrawResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWithdrawResponseWithDefaults

`func NewWithdrawResponseWithDefaults() *WithdrawResponse`

NewWithdrawResponseWithDefaults instantiates a new WithdrawResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSig

`func (o *WithdrawResponse) GetSig() string`

GetSig returns the Sig field if non-nil, zero value otherwise.

### GetSigOk

`func (o *WithdrawResponse) GetSigOk() (*string, bool)`

GetSigOk returns a tuple with the Sig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSig

`func (o *WithdrawResponse) SetSig(v string)`

SetSig sets Sig field to given value.

### HasSig

`func (o *WithdrawResponse) HasSig() bool`

HasSig returns a boolean if a field has been set.

### GetNonce

`func (o *WithdrawResponse) GetNonce() int32`

GetNonce returns the Nonce field if non-nil, zero value otherwise.

### GetNonceOk

`func (o *WithdrawResponse) GetNonceOk() (*int32, bool)`

GetNonceOk returns a tuple with the Nonce field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNonce

`func (o *WithdrawResponse) SetNonce(v int32)`

SetNonce sets Nonce field to given value.

### HasNonce

`func (o *WithdrawResponse) HasNonce() bool`

HasNonce returns a boolean if a field has been set.

### GetAmount

`func (o *WithdrawResponse) GetAmount() string`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *WithdrawResponse) GetAmountOk() (*string, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *WithdrawResponse) SetAmount(v string)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *WithdrawResponse) HasAmount() bool`

HasAmount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


