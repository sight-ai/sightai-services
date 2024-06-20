# Receipt

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int64** |  | 
**UserAddress** | **string** |  | 
**GatewayAddress** | **string** |  | 
**FinishedAt** | Pointer to **time.Time** |  | [optional] 
**Cost** | Pointer to **string** |  | [optional] 
**Proof** | Pointer to **string** |  | [optional] 
**TxnId** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 

## Methods

### NewReceipt

`func NewReceipt(id int64, userAddress string, gatewayAddress string, ) *Receipt`

NewReceipt instantiates a new Receipt object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReceiptWithDefaults

`func NewReceiptWithDefaults() *Receipt`

NewReceiptWithDefaults instantiates a new Receipt object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Receipt) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Receipt) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Receipt) SetId(v int64)`

SetId sets Id field to given value.


### GetUserAddress

`func (o *Receipt) GetUserAddress() string`

GetUserAddress returns the UserAddress field if non-nil, zero value otherwise.

### GetUserAddressOk

`func (o *Receipt) GetUserAddressOk() (*string, bool)`

GetUserAddressOk returns a tuple with the UserAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserAddress

`func (o *Receipt) SetUserAddress(v string)`

SetUserAddress sets UserAddress field to given value.


### GetGatewayAddress

`func (o *Receipt) GetGatewayAddress() string`

GetGatewayAddress returns the GatewayAddress field if non-nil, zero value otherwise.

### GetGatewayAddressOk

`func (o *Receipt) GetGatewayAddressOk() (*string, bool)`

GetGatewayAddressOk returns a tuple with the GatewayAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewayAddress

`func (o *Receipt) SetGatewayAddress(v string)`

SetGatewayAddress sets GatewayAddress field to given value.


### GetFinishedAt

`func (o *Receipt) GetFinishedAt() time.Time`

GetFinishedAt returns the FinishedAt field if non-nil, zero value otherwise.

### GetFinishedAtOk

`func (o *Receipt) GetFinishedAtOk() (*time.Time, bool)`

GetFinishedAtOk returns a tuple with the FinishedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFinishedAt

`func (o *Receipt) SetFinishedAt(v time.Time)`

SetFinishedAt sets FinishedAt field to given value.

### HasFinishedAt

`func (o *Receipt) HasFinishedAt() bool`

HasFinishedAt returns a boolean if a field has been set.

### GetCost

`func (o *Receipt) GetCost() string`

GetCost returns the Cost field if non-nil, zero value otherwise.

### GetCostOk

`func (o *Receipt) GetCostOk() (*string, bool)`

GetCostOk returns a tuple with the Cost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCost

`func (o *Receipt) SetCost(v string)`

SetCost sets Cost field to given value.

### HasCost

`func (o *Receipt) HasCost() bool`

HasCost returns a boolean if a field has been set.

### GetProof

`func (o *Receipt) GetProof() string`

GetProof returns the Proof field if non-nil, zero value otherwise.

### GetProofOk

`func (o *Receipt) GetProofOk() (*string, bool)`

GetProofOk returns a tuple with the Proof field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProof

`func (o *Receipt) SetProof(v string)`

SetProof sets Proof field to given value.

### HasProof

`func (o *Receipt) HasProof() bool`

HasProof returns a boolean if a field has been set.

### GetTxnId

`func (o *Receipt) GetTxnId() string`

GetTxnId returns the TxnId field if non-nil, zero value otherwise.

### GetTxnIdOk

`func (o *Receipt) GetTxnIdOk() (*string, bool)`

GetTxnIdOk returns a tuple with the TxnId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxnId

`func (o *Receipt) SetTxnId(v string)`

SetTxnId sets TxnId field to given value.

### HasTxnId

`func (o *Receipt) HasTxnId() bool`

HasTxnId returns a boolean if a field has been set.

### GetStatus

`func (o *Receipt) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Receipt) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Receipt) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *Receipt) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


