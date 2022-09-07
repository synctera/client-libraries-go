# \ApplicationsApi

All URIs are relative to *https://api.synctera.com/v0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateApplication**](ApplicationsApi.md#CreateApplication) | **Post** /applications | Create a credit application.
[**GetApplication**](ApplicationsApi.md#GetApplication) | **Get** /applications/{application_id} | Get Application
[**ListApplications**](ApplicationsApi.md#ListApplications) | **Get** /applications | List credit applications.
[**PatchApplication**](ApplicationsApi.md#PatchApplication) | **Patch** /applications/{application_id} | Patch Application



## CreateApplication

> ApplicationResponse CreateApplication(ctx).Application(application).IdempotencyKey(idempotencyKey).Execute()

Create a credit application.



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
    application := *openapiclient.NewApplication(map[string]interface{}({"annual_salary":50000,"fico_score":700}), "CustomerId_example", openapiclient.application_type("LINE_OF_CREDIT")) // Application | Submit application details for a credit account
    idempotencyKey := "df122e6f-2ba8-48a5-9508-4350bba5f27e" // string | An idempotency key is an arbitrary unique value generated by client to detect subsequent retries of the same request. It is recommended that a UUID or a similar random identifier be used as an idempotency key (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplicationsApi.CreateApplication(context.Background()).Application(application).IdempotencyKey(idempotencyKey).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplicationsApi.CreateApplication``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateApplication`: ApplicationResponse
    fmt.Fprintf(os.Stdout, "Response from `ApplicationsApi.CreateApplication`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateApplicationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **application** | [**Application**](Application.md) | Submit application details for a credit account | 
 **idempotencyKey** | **string** | An idempotency key is an arbitrary unique value generated by client to detect subsequent retries of the same request. It is recommended that a UUID or a similar random identifier be used as an idempotency key | 

### Return type

[**ApplicationResponse**](ApplicationResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetApplication

> ApplicationResponse GetApplication(ctx, applicationId).Execute()

Get Application



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
    applicationId := "2e3304dc-a1c2-427e-ac5a-a2586a95ce1f" // string | Unique identifier for the application.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplicationsApi.GetApplication(context.Background(), applicationId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplicationsApi.GetApplication``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetApplication`: ApplicationResponse
    fmt.Fprintf(os.Stdout, "Response from `ApplicationsApi.GetApplication`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **string** | Unique identifier for the application. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetApplicationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ApplicationResponse**](ApplicationResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListApplications

> ApplicationList ListApplications(ctx).CustomerId(customerId).Status(status).Execute()

List credit applications.



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
    customerId := []string{"4605deb4-3f8d-4566-9bf1-0dc558b63258"} // []string | A list of customer unique identifiers, with a comma separating any values. (optional)
    status := openapiclient.application_status("APPLICATION_SUBMITTED") // ApplicationStatus | Status of the credit application  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplicationsApi.ListApplications(context.Background()).CustomerId(customerId).Status(status).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplicationsApi.ListApplications``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListApplications`: ApplicationList
    fmt.Fprintf(os.Stdout, "Response from `ApplicationsApi.ListApplications`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListApplicationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **customerId** | **[]string** | A list of customer unique identifiers, with a comma separating any values. | 
 **status** | [**ApplicationStatus**](ApplicationStatus.md) | Status of the credit application  | 

### Return type

[**ApplicationList**](ApplicationList.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchApplication

> ApplicationResponse PatchApplication(ctx, applicationId).ApplicationPatch(applicationPatch).IdempotencyKey(idempotencyKey).Execute()

Patch Application



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
    applicationId := "2e3304dc-a1c2-427e-ac5a-a2586a95ce1f" // string | Unique identifier for the application.
    applicationPatch := *openapiclient.NewApplicationPatch() // ApplicationPatch | Application fields to be patched
    idempotencyKey := "df122e6f-2ba8-48a5-9508-4350bba5f27e" // string | An idempotency key is an arbitrary unique value generated by client to detect subsequent retries of the same request. It is recommended that a UUID or a similar random identifier be used as an idempotency key (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplicationsApi.PatchApplication(context.Background(), applicationId).ApplicationPatch(applicationPatch).IdempotencyKey(idempotencyKey).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplicationsApi.PatchApplication``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PatchApplication`: ApplicationResponse
    fmt.Fprintf(os.Stdout, "Response from `ApplicationsApi.PatchApplication`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **string** | Unique identifier for the application. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchApplicationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **applicationPatch** | [**ApplicationPatch**](ApplicationPatch.md) | Application fields to be patched | 
 **idempotencyKey** | **string** | An idempotency key is an arbitrary unique value generated by client to detect subsequent retries of the same request. It is recommended that a UUID or a similar random identifier be used as an idempotency key | 

### Return type

[**ApplicationResponse**](ApplicationResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)
