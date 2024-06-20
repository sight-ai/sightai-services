# \GatewayApi

All URIs are relative to *http://localhost:10101/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateReceipt**](GatewayApi.md#CreateReceipt) | **Post** /receipt | 
[**GatewayGetReceipts**](GatewayApi.md#GatewayGetReceipts) | **Get** /gateways/{account_id}/receipts | 
[**GatewayWithdraw**](GatewayApi.md#GatewayWithdraw) | **Post** /gateways/{account_id}/withdraw | 



## CreateReceipt

> SimpleMessageResponse CreateReceipt(ctx).CreateReceiptRequest(createReceiptRequest).Execute()





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
    createReceiptRequest := *openapiclient.NewCreateReceiptRequest("UserAddress_example", "GatewayAddress_example", time.Now(), "Cost_example", "Proof_example", "TxnId_example", "Status_example") // CreateReceiptRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GatewayApi.CreateReceipt(context.Background()).CreateReceiptRequest(createReceiptRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GatewayApi.CreateReceipt``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateReceipt`: SimpleMessageResponse
    fmt.Fprintf(os.Stdout, "Response from `GatewayApi.CreateReceipt`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateReceiptRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createReceiptRequest** | [**CreateReceiptRequest**](CreateReceiptRequest.md) |  | 

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


## GatewayGetReceipts

> ReceiptsResponse GatewayGetReceipts(ctx, accountId).Page(page).PageSize(pageSize).UserAddress(userAddress).Status(status).Before(before).After(after).Execute()





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
    userAddress := "userAddress_example" // string |  (optional)
    status := "status_example" // string |  (optional)
    before := time.Now() // time.Time |  (optional)
    after := time.Now() // time.Time |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GatewayApi.GatewayGetReceipts(context.Background(), accountId).Page(page).PageSize(pageSize).UserAddress(userAddress).Status(status).Before(before).After(after).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GatewayApi.GatewayGetReceipts``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GatewayGetReceipts`: ReceiptsResponse
    fmt.Fprintf(os.Stdout, "Response from `GatewayApi.GatewayGetReceipts`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **int64** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGatewayGetReceiptsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **int64** |  | 
 **pageSize** | **int64** |  | 
 **userAddress** | **string** |  | 
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


## GatewayWithdraw

> WithdrawResponse GatewayWithdraw(ctx, accountId).GatewayWithdrawRequest(gatewayWithdrawRequest).Execute()





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
    gatewayWithdrawRequest := *openapiclient.NewGatewayWithdrawRequest([]int64{int64(123)}) // GatewayWithdrawRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GatewayApi.GatewayWithdraw(context.Background(), accountId).GatewayWithdrawRequest(gatewayWithdrawRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GatewayApi.GatewayWithdraw``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GatewayWithdraw`: WithdrawResponse
    fmt.Fprintf(os.Stdout, "Response from `GatewayApi.GatewayWithdraw`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **int64** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGatewayWithdrawRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **gatewayWithdrawRequest** | [**GatewayWithdrawRequest**](GatewayWithdrawRequest.md) |  | 

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

