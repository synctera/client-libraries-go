# \RDCDepositsApi

All URIs are relative to *https://api.synctera.com/v0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateRdcDeposit**](RDCDepositsApi.md#CreateRdcDeposit) | **Post** /accounts/rdc/deposits | Create an RDC Deposit
[**GetRdcDeposit**](RDCDepositsApi.md#GetRdcDeposit) | **Get** /accounts/rdc/deposits/{deposit_id} | Get RDC Deposit
[**ListRdcDeposits**](RDCDepositsApi.md#ListRdcDeposits) | **Get** /accounts/rdc/deposits | List RDC Deposits



## CreateRdcDeposit

> Deposit CreateRdcDeposit(ctx).Deposit(deposit).Execute()

Create an RDC Deposit



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
    deposit := *openapiclient.NewDeposit("83cf7723-58a3-4d6c-bb2a-9f24775a6ba0", "1b2bafc8-db39-4cfa-a064-5a30c36edb96", int32(1153), "2022-01-02T03:04:05.678Z", "2022-01-02T03:04:05.678Z", int32(4931), "USD", "a6b7b53d-e4ec-40e7-a3a2-22a1f1cb8894", "b0b9528e-05d1-4f93-97e3-5e0f3c42b0c3", openapiclient.vendor_info{VendorJson: openapiclient.NewVendorJson("ContentType_example", map[string]interface{}(123), "SOCURE")}) // Deposit | Attributes of the RDC deposit to create (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.RDCDepositsApi.CreateRdcDeposit(context.Background()).Deposit(deposit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RDCDepositsApi.CreateRdcDeposit``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateRdcDeposit`: Deposit
    fmt.Fprintf(os.Stdout, "Response from `RDCDepositsApi.CreateRdcDeposit`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateRdcDepositRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **deposit** | [**Deposit**](Deposit.md) | Attributes of the RDC deposit to create | 

### Return type

[**Deposit**](Deposit.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRdcDeposit

> Deposit GetRdcDeposit(ctx, depositId).Execute()

Get RDC Deposit



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
    depositId := TODO // string | ID of a deposit for a remote deposit capture

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.RDCDepositsApi.GetRdcDeposit(context.Background(), depositId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RDCDepositsApi.GetRdcDeposit``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetRdcDeposit`: Deposit
    fmt.Fprintf(os.Stdout, "Response from `RDCDepositsApi.GetRdcDeposit`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**depositId** | [**string**](.md) | ID of a deposit for a remote deposit capture | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRdcDepositRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Deposit**](Deposit.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListRdcDeposits

> DepositList ListRdcDeposits(ctx).Limit(limit).PageToken(pageToken).Execute()

List RDC Deposits



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
    limit := int32(100) // int32 |  (optional) (default to 100)
    pageToken := "2ss64pmfs8" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.RDCDepositsApi.ListRdcDeposits(context.Background()).Limit(limit).PageToken(pageToken).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RDCDepositsApi.ListRdcDeposits``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListRdcDeposits`: DepositList
    fmt.Fprintf(os.Stdout, "Response from `RDCDepositsApi.ListRdcDeposits`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListRdcDepositsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** |  | [default to 100]
 **pageToken** | **string** |  | 

### Return type

[**DepositList**](DepositList.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

