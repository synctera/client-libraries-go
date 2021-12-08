# \DisclosuresApi

All URIs are relative to *https://api.synctera.com/v0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateDisclosure**](DisclosuresApi.md#CreateDisclosure) | **Post** /customers/{customer_id}/disclosures | Create a Disclosure
[**ListDisclosures**](DisclosuresApi.md#ListDisclosures) | **Get** /customers/{customer_id}/disclosures | List Disclosures



## CreateDisclosure

> Disclosure CreateDisclosure(ctx, customerId).Disclosure(disclosure).Execute()

Create a Disclosure



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
    customerId := TODO // string | The customer's unique identifier
    disclosure := *openapiclient.NewDisclosure("VIEWED", time.Now(), "REG_DD", "v1.1") // Disclosure | Disclosure to create

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DisclosuresApi.CreateDisclosure(context.Background(), customerId).Disclosure(disclosure).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DisclosuresApi.CreateDisclosure``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateDisclosure`: Disclosure
    fmt.Fprintf(os.Stdout, "Response from `DisclosuresApi.CreateDisclosure`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customerId** | [**string**](.md) | The customer&#39;s unique identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateDisclosureRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **disclosure** | [**Disclosure**](Disclosure.md) | Disclosure to create | 

### Return type

[**Disclosure**](Disclosure.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListDisclosures

> DisclosureResponse ListDisclosures(ctx, customerId).Limit(limit).PageToken(pageToken).Execute()

List Disclosures



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
    customerId := TODO // string | The customer's unique identifier
    limit := int32(100) // int32 |  (optional) (default to 100)
    pageToken := "19waxl0g93" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DisclosuresApi.ListDisclosures(context.Background(), customerId).Limit(limit).PageToken(pageToken).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DisclosuresApi.ListDisclosures``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListDisclosures`: DisclosureResponse
    fmt.Fprintf(os.Stdout, "Response from `DisclosuresApi.ListDisclosures`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customerId** | [**string**](.md) | The customer&#39;s unique identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiListDisclosuresRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** |  | [default to 100]
 **pageToken** | **string** |  | 

### Return type

[**DisclosureResponse**](DisclosureResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

