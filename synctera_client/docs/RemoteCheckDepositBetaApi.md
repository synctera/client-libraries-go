# \RemoteCheckDepositBetaApi

All URIs are relative to *https://api.synctera.com/v0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateRdcDeposit**](RemoteCheckDepositBetaApi.md#CreateRdcDeposit) | **Post** /rdc/deposits | Create a Remote Check Deposit
[**GetRdcDeposit**](RemoteCheckDepositBetaApi.md#GetRdcDeposit) | **Get** /rdc/deposits/{deposit_id} | Get Remote Check Deposit
[**ListRdcDeposits**](RemoteCheckDepositBetaApi.md#ListRdcDeposits) | **Get** /rdc/deposits | List Remote Check Deposits



## CreateRdcDeposit

> Deposit CreateRdcDeposit(ctx).Deposit(deposit).Execute()

Create a Remote Check Deposit



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
    deposit := *openapiclient.NewDeposit("ab144f71-3708-4625-8df3-d9307804475e", "a745e775-9d35-46a5-812d-4a7663178a3b", int32(8032), time.Now(), time.Now(), int32(6207), "USD", "a6be36a2-1211-4c7f-b38e-14923e009262", "69170eba-daa5-4c06-9a6d-abc1d4733f67", "Status_example", "261343ed-2752-4e07-b14f-5cafcd38361c", openapiclient.vendor_info1{VendorJson: openapiclient.NewVendorJson("ContentType_example", map[string]interface{}(123), "SOCURE")}) // Deposit | Attributes of the Remote Check Deposit to create (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RemoteCheckDepositBetaApi.CreateRdcDeposit(context.Background()).Deposit(deposit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RemoteCheckDepositBetaApi.CreateRdcDeposit``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateRdcDeposit`: Deposit
    fmt.Fprintf(os.Stdout, "Response from `RemoteCheckDepositBetaApi.CreateRdcDeposit`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateRdcDepositRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **deposit** | [**Deposit**](Deposit.md) | Attributes of the Remote Check Deposit to create | 

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

Get Remote Check Deposit



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
    depositId := "1985f769-dd31-4128-95db-f765355e6631" // string | ID of a deposit for a remote deposit capture

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RemoteCheckDepositBetaApi.GetRdcDeposit(context.Background(), depositId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RemoteCheckDepositBetaApi.GetRdcDeposit``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetRdcDeposit`: Deposit
    fmt.Fprintf(os.Stdout, "Response from `RemoteCheckDepositBetaApi.GetRdcDeposit`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**depositId** | **string** | ID of a deposit for a remote deposit capture | 

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

List Remote Check Deposits



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
    pageToken := "h50ffqz9q5" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RemoteCheckDepositBetaApi.ListRdcDeposits(context.Background()).Limit(limit).PageToken(pageToken).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RemoteCheckDepositBetaApi.ListRdcDeposits``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListRdcDeposits`: DepositList
    fmt.Fprintf(os.Stdout, "Response from `RemoteCheckDepositBetaApi.ListRdcDeposits`: %v\n", resp)
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

