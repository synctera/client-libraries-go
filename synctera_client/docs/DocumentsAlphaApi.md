# \DocumentsAlphaApi

All URIs are relative to *https://api.synctera.com/v0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateDocument**](DocumentsAlphaApi.md#CreateDocument) | **Post** /documents | Create a document
[**GetDocument**](DocumentsAlphaApi.md#GetDocument) | **Get** /documents/{document_id} | Get a document
[**GetDocumentContents**](DocumentsAlphaApi.md#GetDocumentContents) | **Get** /documents/{document_id}/contents | Get a document
[**ListDocuments**](DocumentsAlphaApi.md#ListDocuments) | **Get** /documents | List documents
[**UpdateDocument**](DocumentsAlphaApi.md#UpdateDocument) | **Patch** /documents/{document_id} | Update a document



## CreateDocument

> Document CreateDocument(ctx).File(file).Description(description).Encryption(encryption).Metadata(metadata).Name(name).RelatedResourceId(relatedResourceId).RelatedResourceType(relatedResourceType).Execute()

Create a document



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
    file := os.NewFile(1234, "some_file") // *os.File | 
    description := "description_example" // string |  (optional)
    encryption := openapiclient.encryption("REQUIRED") // Encryption |  (optional) (default to "NOT_REQUIRED")
    metadata := map[string]interface{}{ ... } // map[string]interface{} | Optional field to store additional informaton about the resource.  Intended to be used by the integrator to store non-sensitive data.  (optional)
    name := "name_example" // string |  (optional)
    relatedResourceId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Related resource ID (optional)
    relatedResourceType := openapiclient.related_resource_type("CUSTOMER") // RelatedResourceType |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DocumentsAlphaApi.CreateDocument(context.Background()).File(file).Description(description).Encryption(encryption).Metadata(metadata).Name(name).RelatedResourceId(relatedResourceId).RelatedResourceType(relatedResourceType).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DocumentsAlphaApi.CreateDocument``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateDocument`: Document
    fmt.Fprintf(os.Stdout, "Response from `DocumentsAlphaApi.CreateDocument`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateDocumentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **file** | ***os.File** |  | 
 **description** | **string** |  | 
 **encryption** | [**Encryption**](Encryption.md) |  | [default to &quot;NOT_REQUIRED&quot;]
 **metadata** | [**map[string]interface{}**](map[string]interface{}.md) | Optional field to store additional informaton about the resource.  Intended to be used by the integrator to store non-sensitive data.  | 
 **name** | **string** |  | 
 **relatedResourceId** | **string** | Related resource ID | 
 **relatedResourceType** | [**RelatedResourceType**](RelatedResourceType.md) |  | 

### Return type

[**Document**](Document.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDocument

> *os.File GetDocument(ctx, documentId).Execute()

Get a document



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
    documentId := "dd8cd509-ce52-4990-8f84-316558e68e9a" // string | The unique identifier of a document.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DocumentsAlphaApi.GetDocument(context.Background(), documentId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DocumentsAlphaApi.GetDocument``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDocument`: *os.File
    fmt.Fprintf(os.Stdout, "Response from `DocumentsAlphaApi.GetDocument`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**documentId** | **string** | The unique identifier of a document. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDocumentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[***os.File**](*os.File.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDocumentContents

> *os.File GetDocumentContents(ctx, documentId).Execute()

Get a document



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
    documentId := "dd8cd509-ce52-4990-8f84-316558e68e9a" // string | The unique identifier of a document.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DocumentsAlphaApi.GetDocumentContents(context.Background(), documentId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DocumentsAlphaApi.GetDocumentContents``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDocumentContents`: *os.File
    fmt.Fprintf(os.Stdout, "Response from `DocumentsAlphaApi.GetDocumentContents`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**documentId** | **string** | The unique identifier of a document. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDocumentContentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[***os.File**](*os.File.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListDocuments

> DocumentList ListDocuments(ctx).Id(id).Limit(limit).PageToken(pageToken).Execute()

List documents



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
    id := []string{"81026fb3-d06c-4b37-80da-2b17b4749a3f"} // []string | Unique resource identifier (optional)
    limit := int32(100) // int32 |  (optional) (default to 100)
    pageToken := "h50ffqz9q5" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DocumentsAlphaApi.ListDocuments(context.Background()).Id(id).Limit(limit).PageToken(pageToken).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DocumentsAlphaApi.ListDocuments``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListDocuments`: DocumentList
    fmt.Fprintf(os.Stdout, "Response from `DocumentsAlphaApi.ListDocuments`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListDocumentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **[]string** | Unique resource identifier | 
 **limit** | **int32** |  | [default to 100]
 **pageToken** | **string** |  | 

### Return type

[**DocumentList**](DocumentList.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateDocument

> Document UpdateDocument(ctx, documentId).PatchDocument(patchDocument).Execute()

Update a document



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
    documentId := "dd8cd509-ce52-4990-8f84-316558e68e9a" // string | The unique identifier of a document.
    patchDocument := *openapiclient.NewPatchDocument() // PatchDocument | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DocumentsAlphaApi.UpdateDocument(context.Background(), documentId).PatchDocument(patchDocument).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DocumentsAlphaApi.UpdateDocument``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateDocument`: Document
    fmt.Fprintf(os.Stdout, "Response from `DocumentsAlphaApi.UpdateDocument`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**documentId** | **string** | The unique identifier of a document. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateDocumentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **patchDocument** | [**PatchDocument**](PatchDocument.md) |  | 

### Return type

[**Document**](Document.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

