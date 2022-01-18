# \InternalAccountsApi

All URIs are relative to *https://api.synctera.com/v0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddInternalAccounts**](InternalAccountsApi.md#AddInternalAccounts) | **Post** /internal_accounts | Add internal accounts
[**ListInternalAccounts**](InternalAccountsApi.md#ListInternalAccounts) | **Get** /internal_accounts | List internal accounts



## AddInternalAccounts

> InternalAccount AddInternalAccounts(ctx).InternalAccount(internalAccount).Execute()

Add internal accounts



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
    internalAccount := *openapiclient.NewInternalAccount("NOK", "Status_example") // InternalAccount | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.InternalAccountsApi.AddInternalAccounts(context.Background()).InternalAccount(internalAccount).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InternalAccountsApi.AddInternalAccounts``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddInternalAccounts`: InternalAccount
    fmt.Fprintf(os.Stdout, "Response from `InternalAccountsApi.AddInternalAccounts`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAddInternalAccountsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **internalAccount** | [**InternalAccount**](InternalAccount.md) |  | 

### Return type

[**InternalAccount**](InternalAccount.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListInternalAccounts

> InternalAccountsList ListInternalAccounts(ctx).Limit(limit).PageToken(pageToken).Execute()

List internal accounts



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
    pageToken := "akic8nczf2" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.InternalAccountsApi.ListInternalAccounts(context.Background()).Limit(limit).PageToken(pageToken).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InternalAccountsApi.ListInternalAccounts``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListInternalAccounts`: InternalAccountsList
    fmt.Fprintf(os.Stdout, "Response from `InternalAccountsApi.ListInternalAccounts`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListInternalAccountsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** |  | [default to 100]
 **pageToken** | **string** |  | 

### Return type

[**InternalAccountsList**](InternalAccountsList.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

