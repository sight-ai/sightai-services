# \UserApi

All URIs are relative to *http://localhost:10101/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAccountAllowances**](UserApi.md#GetAccountAllowances) | **Get** /accounts/{account_id}/allowances | 
[**GetAccountInfo**](UserApi.md#GetAccountInfo) | **Get** /accounts/{account_id} | 
[**GetAccountTransactions**](UserApi.md#GetAccountTransactions) | **Get** /accounts/{account_id}/transactions | 
[**GetGateways**](UserApi.md#GetGateways) | **Get** /gateways | 
[**GetNextNonce**](UserApi.md#GetNextNonce) | **Get** /next_nonce | 
[**SignAllowance**](UserApi.md#SignAllowance) | **Post** /sign_allowance | 
[**SignIn**](UserApi.md#SignIn) | **Post** /sign_in | 
[**UserGetReceipts**](UserApi.md#UserGetReceipts) | **Get** /accounts/{account_id}/receipts | 
[**Withdraw**](UserApi.md#Withdraw) | **Post** /withdraw | 



## GetAccountAllowances

> AllowancesResponse GetAccountAllowances(ctx, accountId).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    accountId := int64(789) // int64 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UserApi.GetAccountAllowances(context.Background(), accountId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserApi.GetAccountAllowances``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAccountAllowances`: AllowancesResponse
    fmt.Fprintf(os.Stdout, "Response from `UserApi.GetAccountAllowances`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **int64** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAccountAllowancesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AllowancesResponse**](AllowancesResponse.md)

### Authorization

[userJwtToken](../README.md#userJwtToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAccountInfo

> Account GetAccountInfo(ctx, accountId).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    accountId := int64(789) // int64 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UserApi.GetAccountInfo(context.Background(), accountId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserApi.GetAccountInfo``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAccountInfo`: Account
    fmt.Fprintf(os.Stdout, "Response from `UserApi.GetAccountInfo`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **int64** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAccountInfoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Account**](Account.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAccountTransactions

> TransactionsResponse GetAccountTransactions(ctx, accountId).Page(page).PageSize(pageSize).Type_(type_).Before(before).After(after).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    openapiclient "./openapi"
)

func main() {
    accountId := int64(789) // int64 | 
    page := int64(789) // int64 | 
    pageSize := int64(789) // int64 | 
    type_ := "type__example" // string |  (optional)
    before := time.Now() // time.Time |  (optional)
    after := time.Now() // time.Time |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UserApi.GetAccountTransactions(context.Background(), accountId).Page(page).PageSize(pageSize).Type_(type_).Before(before).After(after).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserApi.GetAccountTransactions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAccountTransactions`: TransactionsResponse
    fmt.Fprintf(os.Stdout, "Response from `UserApi.GetAccountTransactions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **int64** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAccountTransactionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **int64** |  | 
 **pageSize** | **int64** |  | 
 **type_** | **string** |  | 
 **before** | **time.Time** |  | 
 **after** | **time.Time** |  | 

### Return type

[**TransactionsResponse**](TransactionsResponse.md)

### Authorization

[userJwtToken](../README.md#userJwtToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetGateways

> GatewaysResponse GetGateways(ctx).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UserApi.GetGateways(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserApi.GetGateways``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetGateways`: GatewaysResponse
    fmt.Fprintf(os.Stdout, "Response from `UserApi.GetGateways`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetGatewaysRequest struct via the builder pattern


### Return type

[**GatewaysResponse**](GatewaysResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNextNonce

> AccountNextNonceResponse GetNextNonce(ctx).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UserApi.GetNextNonce(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserApi.GetNextNonce``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNextNonce`: AccountNextNonceResponse
    fmt.Fprintf(os.Stdout, "Response from `UserApi.GetNextNonce`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetNextNonceRequest struct via the builder pattern


### Return type

[**AccountNextNonceResponse**](AccountNextNonceResponse.md)

### Authorization

[userJwtToken](../README.md#userJwtToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SignAllowance

> SimpleMessageResponse SignAllowance(ctx).SignAllowanceRequest(signAllowanceRequest).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    signAllowanceRequest := *openapiclient.NewSignAllowanceRequest("Allowance_example") // SignAllowanceRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UserApi.SignAllowance(context.Background()).SignAllowanceRequest(signAllowanceRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserApi.SignAllowance``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SignAllowance`: SimpleMessageResponse
    fmt.Fprintf(os.Stdout, "Response from `UserApi.SignAllowance`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSignAllowanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **signAllowanceRequest** | [**SignAllowanceRequest**](SignAllowanceRequest.md) |  | 

### Return type

[**SimpleMessageResponse**](SimpleMessageResponse.md)

### Authorization

[userJwtToken](../README.md#userJwtToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SignIn

> SignInResponse SignIn(ctx).SignInRequest(signInRequest).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    signInRequest := *openapiclient.NewSignInRequest("Address_example") // SignInRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UserApi.SignIn(context.Background()).SignInRequest(signInRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserApi.SignIn``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SignIn`: SignInResponse
    fmt.Fprintf(os.Stdout, "Response from `UserApi.SignIn`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSignInRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **signInRequest** | [**SignInRequest**](SignInRequest.md) |  | 

### Return type

[**SignInResponse**](SignInResponse.md)

### Authorization

[signature](../README.md#signature)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UserGetReceipts

> ReceiptsResponse UserGetReceipts(ctx, accountId).GatewayAddress(gatewayAddress).Page(page).PageSize(pageSize).Status(status).Before(before).After(after).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    openapiclient "./openapi"
)

func main() {
    accountId := int64(789) // int64 | 
    gatewayAddress := "gatewayAddress_example" // string | 
    page := int64(789) // int64 | 
    pageSize := int64(789) // int64 | 
    status := "status_example" // string |  (optional)
    before := time.Now() // time.Time |  (optional)
    after := time.Now() // time.Time |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UserApi.UserGetReceipts(context.Background(), accountId).GatewayAddress(gatewayAddress).Page(page).PageSize(pageSize).Status(status).Before(before).After(after).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserApi.UserGetReceipts``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UserGetReceipts`: ReceiptsResponse
    fmt.Fprintf(os.Stdout, "Response from `UserApi.UserGetReceipts`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **int64** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUserGetReceiptsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **gatewayAddress** | **string** |  | 
 **page** | **int64** |  | 
 **pageSize** | **int64** |  | 
 **status** | **string** |  | 
 **before** | **time.Time** |  | 
 **after** | **time.Time** |  | 

### Return type

[**ReceiptsResponse**](ReceiptsResponse.md)

### Authorization

[userJwtToken](../README.md#userJwtToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Withdraw

> WithdrawResponse Withdraw(ctx).WithdrawRequest(withdrawRequest).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    withdrawRequest := *openapiclient.NewWithdrawRequest("Amount_example") // WithdrawRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UserApi.Withdraw(context.Background()).WithdrawRequest(withdrawRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserApi.Withdraw``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Withdraw`: WithdrawResponse
    fmt.Fprintf(os.Stdout, "Response from `UserApi.Withdraw`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiWithdrawRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **withdrawRequest** | [**WithdrawRequest**](WithdrawRequest.md) |  | 

### Return type

[**WithdrawResponse**](WithdrawResponse.md)

### Authorization

[signature](../README.md#signature), [userJwtToken](../README.md#userJwtToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

