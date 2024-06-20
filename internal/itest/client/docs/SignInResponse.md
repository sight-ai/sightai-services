# SignInResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UserToken** | **string** | user jwt token | 

## Methods

### NewSignInResponse

`func NewSignInResponse(userToken string, ) *SignInResponse`

NewSignInResponse instantiates a new SignInResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSignInResponseWithDefaults

`func NewSignInResponseWithDefaults() *SignInResponse`

NewSignInResponseWithDefaults instantiates a new SignInResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUserToken

`func (o *SignInResponse) GetUserToken() string`

GetUserToken returns the UserToken field if non-nil, zero value otherwise.

### GetUserTokenOk

`func (o *SignInResponse) GetUserTokenOk() (*string, bool)`

GetUserTokenOk returns a tuple with the UserToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserToken

`func (o *SignInResponse) SetUserToken(v string)`

SetUserToken sets UserToken field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


