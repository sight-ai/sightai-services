# SignInRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address** | **string** |  | 
**Domain** | Pointer to **string** |  | [optional] 

## Methods

### NewSignInRequest

`func NewSignInRequest(address string, ) *SignInRequest`

NewSignInRequest instantiates a new SignInRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSignInRequestWithDefaults

`func NewSignInRequestWithDefaults() *SignInRequest`

NewSignInRequestWithDefaults instantiates a new SignInRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddress

`func (o *SignInRequest) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *SignInRequest) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *SignInRequest) SetAddress(v string)`

SetAddress sets Address field to given value.


### GetDomain

`func (o *SignInRequest) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *SignInRequest) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *SignInRequest) SetDomain(v string)`

SetDomain sets Domain field to given value.

### HasDomain

`func (o *SignInRequest) HasDomain() bool`

HasDomain returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


