# \DisclosuresApi

All URIs are relative to *https://api.synctera.com/v0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateDisclosure**](DisclosuresApi.md#CreateDisclosure) | **Post** /disclosures | Create disclosure
[**GetDisclosure**](DisclosuresApi.md#GetDisclosure) | **Get** /disclosures/{disclosure_id} | Get disclosure
[**ListDisclosures**](DisclosuresApi.md#ListDisclosures) | **Get** /disclosures | List disclosures



## CreateDisclosure

> Disclosure CreateDisclosure(ctx).Disclosure(disclosure).Execute()

Create disclosure



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
    disclosure := *openapiclient.NewDisclosure(time.Now(), "ACKNOWLEDGED", openapiclient.disclosure_type("REG_DD"), "1.0") // Disclosure | Disclosure to create.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DisclosuresApi.CreateDisclosure(context.Background()).Disclosure(disclosure).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DisclosuresApi.CreateDisclosure``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateDisclosure`: Disclosure
    fmt.Fprintf(os.Stdout, "Response from `DisclosuresApi.CreateDisclosure`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateDisclosureRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **disclosure** | [**Disclosure**](Disclosure.md) | Disclosure to create. | 

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


## GetDisclosure

> Disclosure GetDisclosure(ctx, disclosureId).Execute()

Get disclosure



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
    disclosureId := TODO // string | The unique identifier for the disclosure.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DisclosuresApi.GetDisclosure(context.Background(), disclosureId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DisclosuresApi.GetDisclosure``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDisclosure`: Disclosure
    fmt.Fprintf(os.Stdout, "Response from `DisclosuresApi.GetDisclosure`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**disclosureId** | [**string**](.md) | The unique identifier for the disclosure. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDisclosureRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Disclosure**](Disclosure.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListDisclosures

> DisclosureList ListDisclosures(ctx).Id(id).PersonId(personId).BusinessId(businessId).Limit(limit).PageToken(pageToken).Execute()

List disclosures



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
    id := []string{"Inner_example"} // []string | Unique resource identifier (optional)
    personId := []string{"Inner_example"} // []string | Unique identifier for the person. Multiple IDs can be provided as a comma-separated list.  (optional)
    businessId := []string{"Inner_example"} // []string | Unique identifier for the business. Multiple IDs can be provided as a comma-separated list.  (optional)
    limit := int32(100) // int32 |  (optional) (default to 100)
    pageToken := "faker.random.alphaNumeric(10)" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DisclosuresApi.ListDisclosures(context.Background()).Id(id).PersonId(personId).BusinessId(businessId).Limit(limit).PageToken(pageToken).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DisclosuresApi.ListDisclosures``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListDisclosures`: DisclosureList
    fmt.Fprintf(os.Stdout, "Response from `DisclosuresApi.ListDisclosures`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListDisclosuresRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **[]string** | Unique resource identifier | 
 **personId** | **[]string** | Unique identifier for the person. Multiple IDs can be provided as a comma-separated list.  | 
 **businessId** | **[]string** | Unique identifier for the business. Multiple IDs can be provided as a comma-separated list.  | 
 **limit** | **int32** |  | [default to 100]
 **pageToken** | **string** |  | 

### Return type

[**DisclosureList**](DisclosureList.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

