# \TransactionsApi

All URIs are relative to *https://api.synctera.com/v0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**A2aTransfer**](TransactionsApi.md#A2aTransfer) | **Post** /transactions/a2a_transfer | Account to account transfer
[**CreateOutgoingACH**](TransactionsApi.md#CreateOutgoingACH) | **Post** /transactions/ach | Create outgoing ACH
[**DeleteOutgoingACH**](TransactionsApi.md#DeleteOutgoingACH) | **Delete** /transactions/ach/{payment_id} | Delete pending outgoing ACH
[**GetA2ATransfer**](TransactionsApi.md#GetA2ATransfer) | **Get** /transactions/a2a_transfer/{payment_id} | Get account to account transfer
[**GetOutgoingACH**](TransactionsApi.md#GetOutgoingACH) | **Get** /transactions/ach | Get Pending ACH List
[**ListA2ATransfer**](TransactionsApi.md#ListA2ATransfer) | **Get** /transactions/a2a_transfer/list/{customer_id} | List account to account transfer
[**ReverseA2ATransfer**](TransactionsApi.md#ReverseA2ATransfer) | **Post** /transactions/a2a_transfer/{payment_id} | Reverse existing account to account transfer
[**UpdateOutgoingACH**](TransactionsApi.md#UpdateOutgoingACH) | **Put** /transactions/ach/{payment_id} | Update outgoing ACH



## A2aTransfer

> A2aTransfer A2aTransfer(ctx).A2aTransfer(a2aTransfer).Execute()

Account to account transfer



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
    a2aTransfer := *openapiclient.NewA2aTransfer(int32(441), "LBP", openapiclient.dc_sign_type("DEBIT"), "SourceAccount_example", "TargetAccount_example") // A2aTransfer | Account to account transfer to create

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TransactionsApi.A2aTransfer(context.Background()).A2aTransfer(a2aTransfer).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TransactionsApi.A2aTransfer``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `A2aTransfer`: A2aTransfer
    fmt.Fprintf(os.Stdout, "Response from `TransactionsApi.A2aTransfer`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiA2aTransferRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **a2aTransfer** | [**A2aTransfer**](A2aTransfer.md) | Account to account transfer to create | 

### Return type

[**A2aTransfer**](A2aTransfer.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateOutgoingACH

> AchOutgoing CreateOutgoingACH(ctx).AchOutgoing(achOutgoing).MfaToken(mfaToken).Execute()

Create outgoing ACH



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
    achOutgoing := *openapiclient.NewAchOutgoing("afc3532b-cc77-4bf2-b625-bb034263bdc7", int32(536), "KYD", "Terry Aufderhar", "4586598335580180", "472840906", "LY") // AchOutgoing | Outgoing ACH to create
    mfaToken := "409069" // string | Multi-Factor Authentication Token (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TransactionsApi.CreateOutgoingACH(context.Background()).AchOutgoing(achOutgoing).MfaToken(mfaToken).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TransactionsApi.CreateOutgoingACH``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateOutgoingACH`: AchOutgoing
    fmt.Fprintf(os.Stdout, "Response from `TransactionsApi.CreateOutgoingACH`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateOutgoingACHRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **achOutgoing** | [**AchOutgoing**](AchOutgoing.md) | Outgoing ACH to create | 
 **mfaToken** | **string** | Multi-Factor Authentication Token | 

### Return type

[**AchOutgoing**](AchOutgoing.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteOutgoingACH

> DeleteResponse DeleteOutgoingACH(ctx, paymentId).Execute()

Delete pending outgoing ACH



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
    paymentId := TODO // string | Payment ID

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TransactionsApi.DeleteOutgoingACH(context.Background(), paymentId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TransactionsApi.DeleteOutgoingACH``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteOutgoingACH`: DeleteResponse
    fmt.Fprintf(os.Stdout, "Response from `TransactionsApi.DeleteOutgoingACH`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**paymentId** | [**string**](.md) | Payment ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteOutgoingACHRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DeleteResponse**](DeleteResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetA2ATransfer

> A2aTransfer GetA2ATransfer(ctx, paymentId).Execute()

Get account to account transfer



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
    paymentId := TODO // string | Payment ID

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TransactionsApi.GetA2ATransfer(context.Background(), paymentId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TransactionsApi.GetA2ATransfer``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetA2ATransfer`: A2aTransfer
    fmt.Fprintf(os.Stdout, "Response from `TransactionsApi.GetA2ATransfer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**paymentId** | [**string**](.md) | Payment ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetA2ATransferRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**A2aTransfer**](A2aTransfer.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOutgoingACH

> AchOutgoingList GetOutgoingACH(ctx).ACHExecutionDate(aCHExecutionDate).Limit(limit).PageToken(pageToken).Execute()

Get Pending ACH List



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
    aCHExecutionDate := time.Now() // string | Execution Date
    limit := int32(100) // int32 |  (optional) (default to 100)
    pageToken := "bnw3qvoyid" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TransactionsApi.GetOutgoingACH(context.Background()).ACHExecutionDate(aCHExecutionDate).Limit(limit).PageToken(pageToken).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TransactionsApi.GetOutgoingACH``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOutgoingACH`: AchOutgoingList
    fmt.Fprintf(os.Stdout, "Response from `TransactionsApi.GetOutgoingACH`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetOutgoingACHRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **aCHExecutionDate** | **string** | Execution Date | 
 **limit** | **int32** |  | [default to 100]
 **pageToken** | **string** |  | 

### Return type

[**AchOutgoingList**](AchOutgoingList.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListA2ATransfer

> A2aTransferList ListA2ATransfer(ctx, customerId).Limit(limit).PageToken(pageToken).Execute()

List account to account transfer



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
    customerId := TODO // string | Customer ID
    limit := int32(100) // int32 |  (optional) (default to 100)
    pageToken := "bnw3qvoyid" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TransactionsApi.ListA2ATransfer(context.Background(), customerId).Limit(limit).PageToken(pageToken).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TransactionsApi.ListA2ATransfer``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListA2ATransfer`: A2aTransferList
    fmt.Fprintf(os.Stdout, "Response from `TransactionsApi.ListA2ATransfer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customerId** | [**string**](.md) | Customer ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiListA2ATransferRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** |  | [default to 100]
 **pageToken** | **string** |  | 

### Return type

[**A2aTransferList**](A2aTransferList.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReverseA2ATransfer

> A2aTransfer ReverseA2ATransfer(ctx, paymentId).Execute()

Reverse existing account to account transfer



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
    paymentId := TODO // string | Payment ID

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TransactionsApi.ReverseA2ATransfer(context.Background(), paymentId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TransactionsApi.ReverseA2ATransfer``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReverseA2ATransfer`: A2aTransfer
    fmt.Fprintf(os.Stdout, "Response from `TransactionsApi.ReverseA2ATransfer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**paymentId** | [**string**](.md) | Payment ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiReverseA2ATransferRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**A2aTransfer**](A2aTransfer.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateOutgoingACH

> AchOutgoing UpdateOutgoingACH(ctx, paymentId).AchOutgoing(achOutgoing).MfaToken(mfaToken).Execute()

Update outgoing ACH



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
    paymentId := TODO // string | Payment ID
    achOutgoing := *openapiclient.NewAchOutgoing("afc3532b-cc77-4bf2-b625-bb034263bdc7", int32(536), "KYD", "Terry Aufderhar", "4586598335580180", "472840906", "LY") // AchOutgoing | Outgoing ACH to update
    mfaToken := "409069" // string | Multi-Factor Authentication Token (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TransactionsApi.UpdateOutgoingACH(context.Background(), paymentId).AchOutgoing(achOutgoing).MfaToken(mfaToken).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TransactionsApi.UpdateOutgoingACH``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateOutgoingACH`: AchOutgoing
    fmt.Fprintf(os.Stdout, "Response from `TransactionsApi.UpdateOutgoingACH`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**paymentId** | [**string**](.md) | Payment ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateOutgoingACHRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **achOutgoing** | [**AchOutgoing**](AchOutgoing.md) | Outgoing ACH to update | 
 **mfaToken** | **string** | Multi-Factor Authentication Token | 

### Return type

[**AchOutgoing**](AchOutgoing.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

