# \AccountTemplatesApi

All URIs are relative to *https://api.synctera.com/v0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteAccountTemplate**](AccountTemplatesApi.md#DeleteAccountTemplate) | **Delete** /accounts/templates/{template_id} | Delete account template



## DeleteAccountTemplate

> DeleteResponse DeleteAccountTemplate(ctx, templateId).Execute()

Delete account template



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
    templateId := TODO // string | Account Template ID

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AccountTemplatesApi.DeleteAccountTemplate(context.Background(), templateId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountTemplatesApi.DeleteAccountTemplate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteAccountTemplate`: DeleteResponse
    fmt.Fprintf(os.Stdout, "Response from `AccountTemplatesApi.DeleteAccountTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**templateId** | [**string**](.md) | Account Template ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAccountTemplateRequest struct via the builder pattern


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

