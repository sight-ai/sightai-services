# CreateReceiptRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UserAddress** | **string** |  | 
**GatewayAddress** | **string** |  | 
**FinishedAt** | **time.Time** |  | 
**Cost** | **string** |  | 
**Proof** | **string** |  | 
**TxnId** | **string** |  | 
**Status** | **string** |  | 

## Methods

### NewCreateReceiptRequest

`func NewCreateReceiptRequest(userAddress string, gatewayAddress string, finishedAt time.Time, cost string, proof string, txnId string, status string, ) *CreateReceiptRequest`

NewCreateReceiptRequest instantiates a new CreateReceiptRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateReceiptRequestWithDefaults

`func NewCreateReceiptRequestWithDefaults() *CreateReceiptRequest`

NewCreateReceiptRequestWithDefaults instantiates a new CreateReceiptRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUserAddress

`func (o *CreateReceiptRequest) GetUserAddress() string`

GetUserAddress returns the UserAddress field if non-nil, zero value otherwise.

### GetUserAddressOk

`func (o *CreateReceiptRequest) GetUserAddressOk() (*string, bool)`

GetUserAddressOk returns a tuple with the UserAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserAddress

`func (o *CreateReceiptRequest) SetUserAddress(v string)`

SetUserAddress sets UserAddress field to given value.


### GetGatewayAddress

`func (o *CreateReceiptRequest) GetGatewayAddress() string`

GetGatewayAddress returns the GatewayAddress field if non-nil, zero value otherwise.

### GetGatewayAddressOk

`func (o *CreateReceiptRequest) GetGatewayAddressOk() (*string, bool)`

GetGatewayAddressOk returns a tuple with the GatewayAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewayAddress

`func (o *CreateReceiptRequest) SetGatewayAddress(v string)`

SetGatewayAddress sets GatewayAddress field to given value.


### GetFinishedAt

`func (o *CreateReceiptRequest) GetFinishedAt() time.Time`

GetFinishedAt returns the FinishedAt field if non-nil, zero value otherwise.

### GetFinishedAtOk

`func (o *CreateReceiptRequest) GetFinishedAtOk() (*time.Time, bool)`

GetFinishedAtOk returns a tuple with the FinishedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFinishedAt

`func (o *CreateReceiptRequest) SetFinishedAt(v time.Time)`

SetFinishedAt sets FinishedAt field to given value.


### GetCost

`func (o *CreateReceiptRequest) GetCost() string`

GetCost returns the Cost field if non-nil, zero value otherwise.

### GetCostOk

`func (o *CreateReceiptRequest) GetCostOk() (*string, bool)`

GetCostOk returns a tuple with the Cost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCost

`func (o *CreateReceiptRequest) SetCost(v string)`

SetCost sets Cost field to given value.


### GetProof

`func (o *CreateReceiptRequest) GetProof() string`

GetProof returns the Proof field if non-nil, zero value otherwise.

### GetProofOk

`func (o *CreateReceiptRequest) GetProofOk() (*string, bool)`

GetProofOk returns a tuple with the Proof field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProof

`func (o *CreateReceiptRequest) SetProof(v string)`

SetProof sets Proof field to given value.


### GetTxnId

`func (o *CreateReceiptRequest) GetTxnId() string`

GetTxnId returns the TxnId field if non-nil, zero value otherwise.

### GetTxnIdOk

`func (o *CreateReceiptRequest) GetTxnIdOk() (*string, bool)`

GetTxnIdOk returns a tuple with the TxnId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxnId

`func (o *CreateReceiptRequest) SetTxnId(v string)`

SetTxnId sets TxnId field to given value.


### GetStatus

`func (o *CreateReceiptRequest) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CreateReceiptRequest) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CreateReceiptRequest) SetStatus(v string)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


