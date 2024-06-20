# AllowancesResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FromAllowances** | Pointer to [**[]Allowance**](Allowance.md) |  | [optional] 
**ToAllowances** | Pointer to [**[]Allowance**](Allowance.md) |  | [optional] 

## Methods

### NewAllowancesResponse

`func NewAllowancesResponse() *AllowancesResponse`

NewAllowancesResponse instantiates a new AllowancesResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAllowancesResponseWithDefaults

`func NewAllowancesResponseWithDefaults() *AllowancesResponse`

NewAllowancesResponseWithDefaults instantiates a new AllowancesResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFromAllowances

`func (o *AllowancesResponse) GetFromAllowances() []Allowance`

GetFromAllowances returns the FromAllowances field if non-nil, zero value otherwise.

### GetFromAllowancesOk

`func (o *AllowancesResponse) GetFromAllowancesOk() (*[]Allowance, bool)`

GetFromAllowancesOk returns a tuple with the FromAllowances field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromAllowances

`func (o *AllowancesResponse) SetFromAllowances(v []Allowance)`

SetFromAllowances sets FromAllowances field to given value.

### HasFromAllowances

`func (o *AllowancesResponse) HasFromAllowances() bool`

HasFromAllowances returns a boolean if a field has been set.

### GetToAllowances

`func (o *AllowancesResponse) GetToAllowances() []Allowance`

GetToAllowances returns the ToAllowances field if non-nil, zero value otherwise.

### GetToAllowancesOk

`func (o *AllowancesResponse) GetToAllowancesOk() (*[]Allowance, bool)`

GetToAllowancesOk returns a tuple with the ToAllowances field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToAllowances

`func (o *AllowancesResponse) SetToAllowances(v []Allowance)`

SetToAllowances sets ToAllowances field to given value.

### HasToAllowances

`func (o *AllowancesResponse) HasToAllowances() bool`

HasToAllowances returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


