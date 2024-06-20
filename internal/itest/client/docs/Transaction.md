# Transaction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int64** |  | 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**AccountId** | **int64** |  | 
**AvailableDelta** | Pointer to **string** |  | [optional] 
**HoldDelta** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**Notes** | Pointer to **string** |  | [optional] 

## Methods

### NewTransaction

`func NewTransaction(id int64, accountId int64, ) *Transaction`

NewTransaction instantiates a new Transaction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransactionWithDefaults

`func NewTransactionWithDefaults() *Transaction`

NewTransactionWithDefaults instantiates a new Transaction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Transaction) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Transaction) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Transaction) SetId(v int64)`

SetId sets Id field to given value.


### GetCreatedAt

`func (o *Transaction) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Transaction) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Transaction) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *Transaction) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetAccountId

`func (o *Transaction) GetAccountId() int64`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *Transaction) GetAccountIdOk() (*int64, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *Transaction) SetAccountId(v int64)`

SetAccountId sets AccountId field to given value.


### GetAvailableDelta

`func (o *Transaction) GetAvailableDelta() string`

GetAvailableDelta returns the AvailableDelta field if non-nil, zero value otherwise.

### GetAvailableDeltaOk

`func (o *Transaction) GetAvailableDeltaOk() (*string, bool)`

GetAvailableDeltaOk returns a tuple with the AvailableDelta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailableDelta

`func (o *Transaction) SetAvailableDelta(v string)`

SetAvailableDelta sets AvailableDelta field to given value.

### HasAvailableDelta

`func (o *Transaction) HasAvailableDelta() bool`

HasAvailableDelta returns a boolean if a field has been set.

### GetHoldDelta

`func (o *Transaction) GetHoldDelta() string`

GetHoldDelta returns the HoldDelta field if non-nil, zero value otherwise.

### GetHoldDeltaOk

`func (o *Transaction) GetHoldDeltaOk() (*string, bool)`

GetHoldDeltaOk returns a tuple with the HoldDelta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHoldDelta

`func (o *Transaction) SetHoldDelta(v string)`

SetHoldDelta sets HoldDelta field to given value.

### HasHoldDelta

`func (o *Transaction) HasHoldDelta() bool`

HasHoldDelta returns a boolean if a field has been set.

### GetType

`func (o *Transaction) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Transaction) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Transaction) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *Transaction) HasType() bool`

HasType returns a boolean if a field has been set.

### GetNotes

`func (o *Transaction) GetNotes() string`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *Transaction) GetNotesOk() (*string, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *Transaction) SetNotes(v string)`

SetNotes sets Notes field to given value.

### HasNotes

`func (o *Transaction) HasNotes() bool`

HasNotes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


