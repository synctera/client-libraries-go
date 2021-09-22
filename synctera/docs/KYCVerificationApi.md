# \KYCVerificationApi

All URIs are relative to *https://api.synctera.com/v0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateCustomerVerificationResult**](KYCVerificationApi.md#CreateCustomerVerificationResult) | **Post** /customers/{customer_id}/verifications | Create a Customer Verification Result
[**GetVerification**](KYCVerificationApi.md#GetVerification) | **Get** /customers/{customer_id}/verifications/{verification_id} | Get Verification Result
[**ListVerifications**](KYCVerificationApi.md#ListVerifications) | **Get** /customers/{customer_id}/verifications | List Verification Results
[**VerifyCustomer**](KYCVerificationApi.md#VerifyCustomer) | **Post** /customers/{customer_id}/verify | Verify a customer&#39;s identity



## CreateCustomerVerificationResult

> CustomerVerificationResult CreateCustomerVerificationResult(ctx, customerId).CustomerVerificationResult(customerVerificationResult).Execute()

Create a Customer Verification Result



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
    customerId := TODO // string | Customer ID
    customerVerificationResult := *openapiclient.NewCustomerVerificationResult("ACCEPTED", time.Now(), openapiclient.verification_type("fraud")) // CustomerVerificationResult | Customer verification result to create

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.KYCVerificationApi.CreateCustomerVerificationResult(context.Background(), customerId).CustomerVerificationResult(customerVerificationResult).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KYCVerificationApi.CreateCustomerVerificationResult``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateCustomerVerificationResult`: CustomerVerificationResult
    fmt.Fprintf(os.Stdout, "Response from `KYCVerificationApi.CreateCustomerVerificationResult`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customerId** | [**string**](.md) | Customer ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateCustomerVerificationResultRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **customerVerificationResult** | [**CustomerVerificationResult**](CustomerVerificationResult.md) | Customer verification result to create | 

### Return type

[**CustomerVerificationResult**](CustomerVerificationResult.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetVerification

> CustomerVerificationResult GetVerification(ctx, customerId, verificationId).Execute()

Get Verification Result



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
    verificationId := TODO // string | Verification ID

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.KYCVerificationApi.GetVerification(context.Background(), customerId, verificationId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KYCVerificationApi.GetVerification``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetVerification`: CustomerVerificationResult
    fmt.Fprintf(os.Stdout, "Response from `KYCVerificationApi.GetVerification`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customerId** | [**string**](.md) | Customer ID | 
**verificationId** | [**string**](.md) | Verification ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetVerificationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**CustomerVerificationResult**](CustomerVerificationResult.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListVerifications

> CustomerVerificationResultList ListVerifications(ctx, customerId).IncludeHistory(includeHistory).Limit(limit).PageToken(pageToken).Execute()

List Verification Results



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
    includeHistory := true // bool | Whether to include a history of records or not (optional)
    limit := int32(100) // int32 |  (optional) (default to 100)
    pageToken := "bnw3qvoyid" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.KYCVerificationApi.ListVerifications(context.Background(), customerId).IncludeHistory(includeHistory).Limit(limit).PageToken(pageToken).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KYCVerificationApi.ListVerifications``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListVerifications`: CustomerVerificationResultList
    fmt.Fprintf(os.Stdout, "Response from `KYCVerificationApi.ListVerifications`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customerId** | [**string**](.md) | Customer ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiListVerificationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **includeHistory** | **bool** | Whether to include a history of records or not | 
 **limit** | **int32** |  | [default to 100]
 **pageToken** | **string** |  | 

### Return type

[**CustomerVerificationResultList**](CustomerVerificationResultList.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## VerifyCustomer

> CustomerVerificationResult VerifyCustomer(ctx, customerId).CustomerVerification(customerVerification).Execute()

Verify a customer's identity



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
    customerVerification := *openapiclient.NewCustomerVerification([]openapiclient.VerificationType{openapiclient.verification_type("fraud")}) // CustomerVerification | Customer verification request

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.KYCVerificationApi.VerifyCustomer(context.Background(), customerId).CustomerVerification(customerVerification).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KYCVerificationApi.VerifyCustomer``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `VerifyCustomer`: CustomerVerificationResult
    fmt.Fprintf(os.Stdout, "Response from `KYCVerificationApi.VerifyCustomer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customerId** | [**string**](.md) | Customer ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiVerifyCustomerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **customerVerification** | [**CustomerVerification**](CustomerVerification.md) | Customer verification request | 

### Return type

[**CustomerVerificationResult**](CustomerVerificationResult.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

