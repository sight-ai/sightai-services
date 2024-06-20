# \AdminApi

All URIs are relative to *http://localhost:10101/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AdminDeposit**](AdminApi.md#AdminDeposit) | **Post** /deposit | 
[**UpsertGateway**](AdminApi.md#UpsertGateway) | **Put** /gateway | 



## AdminDeposit

> SimpleMessageResponse AdminDeposit(ctx).DepositRequest(depositRequest).Execute()





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
    depositRequest := *openapiclient.NewDepositRequest("Address_example", "Amount_example") // DepositRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AdminApi.AdminDeposit(context.Background()).DepositRequest(depositRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AdminApi.AdminDeposit``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AdminDeposit`: SimpleMessageResponse
    fmt.Fprintf(os.Stdout, "Response from `AdminApi.AdminDeposit`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAdminDepositRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **depositRequest** | [**DepositRequest**](DepositRequest.md) |  | 

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


## UpsertGateway

> SimpleMessageResponse UpsertGateway(ctx).UpsertGatewayRequest(upsertGatewayRequest).Execute()





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
    upsertGatewayRequest := *openapiclient.NewUpsertGatewayRequest("Address_example") // UpsertGatewayRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AdminApi.UpsertGateway(context.Background()).UpsertGatewayRequest(upsertGatewayRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AdminApi.UpsertGateway``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpsertGateway`: SimpleMessageResponse
    fmt.Fprintf(os.Stdout, "Response from `AdminApi.UpsertGateway`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpsertGatewayRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **upsertGatewayRequest** | [**UpsertGatewayRequest**](UpsertGatewayRequest.md) |  | 

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

