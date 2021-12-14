# \RDCDepositsApi

All URIs are relative to *https://api.synctera.com/v0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateRdcDeposit**](RDCDepositsApi.md#CreateRdcDeposit) | **Post** /accounts/{account_id}/rdc/deposit | Create an RDC Deposit
[**CreateRdcImage**](RDCDepositsApi.md#CreateRdcImage) | **Post** /accounts/{account_id}/rdc/images | Create an RDC Image
[**CreateRdcScan**](RDCDepositsApi.md#CreateRdcScan) | **Post** /accounts/{account_id}/rdc/scans | Create an RDC Scan
[**GetRdcDeposit**](RDCDepositsApi.md#GetRdcDeposit) | **Get** /accounts/{account_id}/rdc/deposit/{deposit_id} | Get RDC Deposit
[**GetRdcImage**](RDCDepositsApi.md#GetRdcImage) | **Get** /accounts/{account_id}/rdc/images/{image_id} | Get RDC Image
[**GetRdcScan**](RDCDepositsApi.md#GetRdcScan) | **Get** /accounts/{account_id}/rdc/scans/{scan_id} | Get RDC Scan
[**ListRdcDeposits**](RDCDepositsApi.md#ListRdcDeposits) | **Get** /accounts/{account_id}/rdc/deposit | List RDC Deposits
[**ListRdcImageIds**](RDCDepositsApi.md#ListRdcImageIds) | **Get** /accounts/{account_id}/rdc/images | List RDC Image Ids
[**ListRdcScans**](RDCDepositsApi.md#ListRdcScans) | **Get** /accounts/{account_id}/rdc/scans | List RDC Scans



## CreateRdcDeposit

> Deposit CreateRdcDeposit(ctx, accountId).Deposit(deposit).Execute()

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
    accountId := TODO // string | Unique identifier for the account.
    deposit := *openapiclient.NewDeposit("zl03l21ecqlmuk", "a9cf61a6-459c-4775-a1fe-24478c491d7e", int32(5176), "USD", "2019-01-02T03:04:05.678Z", "2019-01-02T03:04:05.678Z", int32(3561), "USD", "1a433f80-ff63-4153-8d7d-7b67cdc5ceb2", "c62bc781-8978-49d3-9109-27aadc2d71e4", "123456789", "c6860cd0-3da4-4556-86b4-5be044f56d90") // Deposit | Attributes of the RDC deposit to create (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.RDCDepositsApi.CreateRdcDeposit(context.Background(), accountId).Deposit(deposit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RDCDepositsApi.CreateRdcDeposit``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateRdcDeposit`: Deposit
    fmt.Fprintf(os.Stdout, "Response from `RDCDepositsApi.CreateRdcDeposit`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | [**string**](.md) | Unique identifier for the account. | 

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


## CreateRdcImage

> Image CreateRdcImage(ctx, accountId).Image(image).Execute()

Create an RDC Image



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
    accountId := TODO // string | Unique identifier for the account.
    image := *openapiclient.NewImage(string(123), "2019-01-02T03:04:05.678Z", "f18babb4-1dc2-498a-a05d-992083a99cd1", openapiclient.rdc_media_type("PDF")) // Image | RDC image to create (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.RDCDepositsApi.CreateRdcImage(context.Background(), accountId).Image(image).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RDCDepositsApi.CreateRdcImage``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateRdcImage`: Image
    fmt.Fprintf(os.Stdout, "Response from `RDCDepositsApi.CreateRdcImage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | [**string**](.md) | Unique identifier for the account. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateRdcImageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **image** | [**Image**](Image.md) | RDC image to create | 

### Return type

[**Image**](Image.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateRdcScan

> Scan CreateRdcScan(ctx, accountId).Scan(scan).Execute()

Create an RDC Scan



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
    accountId := TODO // string | Unique identifier for the account.
    scan := *openapiclient.NewScan("ezyyq4k2sf3wdn", "8463a333-b216-4121-818a-8788978355ac", int32(9002), "USD", "2019-01-02T03:04:05.678Z", "5615ff98-f948-45a9-83c5-ae1f4a57cc19", "6e2c215f-3011-4cdf-892e-8cfe2ee1df74", "123456789") // Scan | RDC scan to create (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.RDCDepositsApi.CreateRdcScan(context.Background(), accountId).Scan(scan).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RDCDepositsApi.CreateRdcScan``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateRdcScan`: Scan
    fmt.Fprintf(os.Stdout, "Response from `RDCDepositsApi.CreateRdcScan`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | [**string**](.md) | Unique identifier for the account. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateRdcScanRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **scan** | [**Scan**](Scan.md) | RDC scan to create | 

### Return type

[**Scan**](Scan.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRdcDeposit

> Deposit GetRdcDeposit(ctx, accountId, depositId).Execute()

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
    accountId := TODO // string | Unique identifier for the account.
    depositId := TODO // string | ID of a deposit for a remote deposit capture

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.RDCDepositsApi.GetRdcDeposit(context.Background(), accountId, depositId).Execute()
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
**accountId** | [**string**](.md) | Unique identifier for the account. | 
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


## GetRdcImage

> Image GetRdcImage(ctx, accountId, imageId).Execute()

Get RDC Image



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
    accountId := TODO // string | Unique identifier for the account.
    imageId := TODO // string | ID of an image uploaded for a remote deposit capture

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.RDCDepositsApi.GetRdcImage(context.Background(), accountId, imageId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RDCDepositsApi.GetRdcImage``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetRdcImage`: Image
    fmt.Fprintf(os.Stdout, "Response from `RDCDepositsApi.GetRdcImage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | [**string**](.md) | Unique identifier for the account. | 
**imageId** | [**string**](.md) | ID of an image uploaded for a remote deposit capture | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRdcImageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**Image**](Image.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRdcScan

> Scan GetRdcScan(ctx, accountId, scanId).Execute()

Get RDC Scan



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
    accountId := TODO // string | Unique identifier for the account.
    scanId := TODO // string | ID of an OCR scan of a check image uploaded for a remote deposit capture

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.RDCDepositsApi.GetRdcScan(context.Background(), accountId, scanId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RDCDepositsApi.GetRdcScan``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetRdcScan`: Scan
    fmt.Fprintf(os.Stdout, "Response from `RDCDepositsApi.GetRdcScan`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | [**string**](.md) | Unique identifier for the account. | 
**scanId** | [**string**](.md) | ID of an OCR scan of a check image uploaded for a remote deposit capture | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRdcScanRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**Scan**](Scan.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListRdcDeposits

> DepositList ListRdcDeposits(ctx, accountId).Limit(limit).PageToken(pageToken).Execute()

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
    accountId := TODO // string | Unique identifier for the account.
    limit := int32(100) // int32 |  (optional) (default to 100)
    pageToken := "0w75x33ztx" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.RDCDepositsApi.ListRdcDeposits(context.Background(), accountId).Limit(limit).PageToken(pageToken).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RDCDepositsApi.ListRdcDeposits``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListRdcDeposits`: DepositList
    fmt.Fprintf(os.Stdout, "Response from `RDCDepositsApi.ListRdcDeposits`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | [**string**](.md) | Unique identifier for the account. | 

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


## ListRdcImageIds

> ImageList ListRdcImageIds(ctx, accountId).Limit(limit).PageToken(pageToken).Execute()

List RDC Image Ids



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
    accountId := TODO // string | Unique identifier for the account.
    limit := int32(100) // int32 |  (optional) (default to 100)
    pageToken := "0w75x33ztx" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.RDCDepositsApi.ListRdcImageIds(context.Background(), accountId).Limit(limit).PageToken(pageToken).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RDCDepositsApi.ListRdcImageIds``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListRdcImageIds`: ImageList
    fmt.Fprintf(os.Stdout, "Response from `RDCDepositsApi.ListRdcImageIds`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | [**string**](.md) | Unique identifier for the account. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListRdcImageIdsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** |  | [default to 100]
 **pageToken** | **string** |  | 

### Return type

[**ImageList**](ImageList.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListRdcScans

> ScanList ListRdcScans(ctx, accountId).Limit(limit).PageToken(pageToken).Execute()

List RDC Scans



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
    accountId := TODO // string | Unique identifier for the account.
    limit := int32(100) // int32 |  (optional) (default to 100)
    pageToken := "0w75x33ztx" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.RDCDepositsApi.ListRdcScans(context.Background(), accountId).Limit(limit).PageToken(pageToken).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RDCDepositsApi.ListRdcScans``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListRdcScans`: ScanList
    fmt.Fprintf(os.Stdout, "Response from `RDCDepositsApi.ListRdcScans`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | [**string**](.md) | Unique identifier for the account. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListRdcScansRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** |  | [default to 100]
 **pageToken** | **string** |  | 

### Return type

[**ScanList**](ScanList.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

