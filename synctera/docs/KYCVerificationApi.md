# \KYCVerificationApi

All URIs are relative to *https://api.synctera.com/v0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateCustomerVerificationResult**](KYCVerificationApi.md#CreateCustomerVerificationResult) | **Post** /customers/{customer_id}/verifications | Create a Customer Verification Result
[**CreateDocument**](KYCVerificationApi.md#CreateDocument) | **Post** /customers/{customer_id}/documents | Create a Document
[**GetDocument**](KYCVerificationApi.md#GetDocument) | **Get** /customers/{customer_id}/documents/{document_id} | Get Document
[**GetVerification**](KYCVerificationApi.md#GetVerification) | **Get** /customers/{customer_id}/verifications/{verification_id} | Get Verification Result
[**ListDocuments**](KYCVerificationApi.md#ListDocuments) | **Get** /customers/{customer_id}/documents | List Documents
[**ListVerifications**](KYCVerificationApi.md#ListVerifications) | **Get** /customers/{customer_id}/verifications | List Verification Results
[**PatchDocument**](KYCVerificationApi.md#PatchDocument) | **Patch** /customers/{customer_id}/documents/{document_id} | Patch Document
[**UpdateDocument**](KYCVerificationApi.md#UpdateDocument) | **Put** /customers/{customer_id}/documents/{document_id} | Update Document
[**VerifyCustomer**](KYCVerificationApi.md#VerifyCustomer) | **Post** /customers/{customer_id}/verify | Verify a Customer



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
    customerVerificationResult := *openapiclient.NewCustomerVerificationResult([]openapiclient.VerificationType{openapiclient.verification_type("fraud")}, "PASS", time.Now()) // CustomerVerificationResult | Customer verification result to create

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
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateDocument

> Document CreateDocument(ctx, customerId).Document(document).Execute()

Create a Document



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
    document := *openapiclient.NewDocument(openapiclient.kyc_media_type("PDF"), openapiclient.document_type("LICENSE_FRONT")) // Document | Document to create

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.KYCVerificationApi.CreateDocument(context.Background(), customerId).Document(document).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KYCVerificationApi.CreateDocument``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateDocument`: Document
    fmt.Fprintf(os.Stdout, "Response from `KYCVerificationApi.CreateDocument`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customerId** | [**string**](.md) | Customer ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateDocumentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **document** | [**Document**](Document.md) | Document to create | 

### Return type

[**Document**](Document.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDocument

> Document GetDocument(ctx, customerId, documentId).Execute()

Get Document



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
    documentId := TODO // string | Document ID

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.KYCVerificationApi.GetDocument(context.Background(), customerId, documentId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KYCVerificationApi.GetDocument``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDocument`: Document
    fmt.Fprintf(os.Stdout, "Response from `KYCVerificationApi.GetDocument`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customerId** | [**string**](.md) | Customer ID | 
**documentId** | [**string**](.md) | Document ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDocumentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**Document**](Document.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

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
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListDocuments

> DocumentList ListDocuments(ctx, customerId).Limit(limit).PageToken(pageToken).Execute()

List Documents



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
    pageToken := "faker.random.alphaNumeric(10)" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.KYCVerificationApi.ListDocuments(context.Background(), customerId).Limit(limit).PageToken(pageToken).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KYCVerificationApi.ListDocuments``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListDocuments`: DocumentList
    fmt.Fprintf(os.Stdout, "Response from `KYCVerificationApi.ListDocuments`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customerId** | [**string**](.md) | Customer ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiListDocumentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** |  | [default to 100]
 **pageToken** | **string** |  | 

### Return type

[**DocumentList**](DocumentList.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListVerifications

> CustomerVerificationResultList ListVerifications(ctx, customerId).Limit(limit).PageToken(pageToken).Execute()

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
    limit := int32(100) // int32 |  (optional) (default to 100)
    pageToken := "faker.random.alphaNumeric(10)" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.KYCVerificationApi.ListVerifications(context.Background(), customerId).Limit(limit).PageToken(pageToken).Execute()
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

 **limit** | **int32** |  | [default to 100]
 **pageToken** | **string** |  | 

### Return type

[**CustomerVerificationResultList**](CustomerVerificationResultList.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchDocument

> Document PatchDocument(ctx, customerId, documentId).Document(document).Execute()

Patch Document

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
    documentId := TODO // string | Document ID
    document := *openapiclient.NewDocument(openapiclient.kyc_media_type("PDF"), openapiclient.document_type("LICENSE_FRONT")) // Document | Document to patch

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.KYCVerificationApi.PatchDocument(context.Background(), customerId, documentId).Document(document).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KYCVerificationApi.PatchDocument``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PatchDocument`: Document
    fmt.Fprintf(os.Stdout, "Response from `KYCVerificationApi.PatchDocument`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customerId** | [**string**](.md) | Customer ID | 
**documentId** | [**string**](.md) | Document ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchDocumentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **document** | [**Document**](Document.md) | Document to patch | 

### Return type

[**Document**](Document.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateDocument

> Document UpdateDocument(ctx, customerId, documentId).Document(document).Execute()

Update Document

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
    documentId := TODO // string | Document ID
    document := *openapiclient.NewDocument(openapiclient.kyc_media_type("PDF"), openapiclient.document_type("LICENSE_FRONT")) // Document | Document to update

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.KYCVerificationApi.UpdateDocument(context.Background(), customerId, documentId).Document(document).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KYCVerificationApi.UpdateDocument``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateDocument`: Document
    fmt.Fprintf(os.Stdout, "Response from `KYCVerificationApi.UpdateDocument`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customerId** | [**string**](.md) | Customer ID | 
**documentId** | [**string**](.md) | Document ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateDocumentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **document** | [**Document**](Document.md) | Document to update | 

### Return type

[**Document**](Document.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## VerifyCustomer

> CustomerVerification VerifyCustomer(ctx, customerId).CustomerVerification(customerVerification).Execute()

Verify a Customer



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
    customerVerification := *openapiclient.NewCustomerVerification([]openapiclient.VerificationType{openapiclient.verification_type("fraud")}) // CustomerVerification | Customer verification

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.KYCVerificationApi.VerifyCustomer(context.Background(), customerId).CustomerVerification(customerVerification).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KYCVerificationApi.VerifyCustomer``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `VerifyCustomer`: CustomerVerification
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

 **customerVerification** | [**CustomerVerification**](CustomerVerification.md) | Customer verification | 

### Return type

[**CustomerVerification**](CustomerVerification.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

