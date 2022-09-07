# \ReconciliationsApi

All URIs are relative to *https://api.synctera.com/v0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateReconciliation**](ReconciliationsApi.md#CreateReconciliation) | **Post** /reconciliations | Create a reconciliation
[**GetReconciliation**](ReconciliationsApi.md#GetReconciliation) | **Get** /reconciliations/{reconciliation_id} | Get reconciliation
[**ListReconciliations**](ReconciliationsApi.md#ListReconciliations) | **Get** /reconciliations | List reconciliations



## CreateReconciliation

> Reconciliation CreateReconciliation(ctx).ReconciliationInput(reconciliationInput).Execute()

Create a reconciliation



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
    reconciliationInput := *openapiclient.NewReconciliationInput(string(123), "sg67wa98toiudl") // ReconciliationInput | Reconciliation to perform (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReconciliationsApi.CreateReconciliation(context.Background()).ReconciliationInput(reconciliationInput).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReconciliationsApi.CreateReconciliation``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateReconciliation`: Reconciliation
    fmt.Fprintf(os.Stdout, "Response from `ReconciliationsApi.CreateReconciliation`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateReconciliationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **reconciliationInput** | [**ReconciliationInput**](ReconciliationInput.md) | Reconciliation to perform | 

### Return type

[**Reconciliation**](Reconciliation.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetReconciliation

> Reconciliation GetReconciliation(ctx, reconciliationId).Execute()

Get reconciliation



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
    reconciliationId := "6917767f-50b9-475e-91c0-ed7d1efc8e13" // string | Reconciliation id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReconciliationsApi.GetReconciliation(context.Background(), reconciliationId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReconciliationsApi.GetReconciliation``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetReconciliation`: Reconciliation
    fmt.Fprintf(os.Stdout, "Response from `ReconciliationsApi.GetReconciliation`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reconciliationId** | **string** | Reconciliation id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetReconciliationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Reconciliation**](Reconciliation.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListReconciliations

> ReconciliationList ListReconciliations(ctx).Limit(limit).PageToken(pageToken).Execute()

List reconciliations



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
    resp, r, err := apiClient.ReconciliationsApi.ListReconciliations(context.Background()).Limit(limit).PageToken(pageToken).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReconciliationsApi.ListReconciliations``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListReconciliations`: ReconciliationList
    fmt.Fprintf(os.Stdout, "Response from `ReconciliationsApi.ListReconciliations`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListReconciliationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** |  | [default to 100]
 **pageToken** | **string** |  | 

### Return type

[**ReconciliationList**](ReconciliationList.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

