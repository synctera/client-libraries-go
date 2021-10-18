# \ExternalAccountsApi

All URIs are relative to *https://api.synctera.com/v0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddVendorExternalAccounts**](ExternalAccountsApi.md#AddVendorExternalAccounts) | **Post** /external_accounts/add_vendor_accounts | Add external accounts through a vendor, such as Plaid.
[**CreateAccessToken**](ExternalAccountsApi.md#CreateAccessToken) | **Post** /external_accounts/access_tokens | Create a permanent access token for an external account
[**CreateVerificationLinkToken**](ExternalAccountsApi.md#CreateVerificationLinkToken) | **Post** /external_accounts/link_tokens | Create a link token to verify an external account
[**GetExternalAccount**](ExternalAccountsApi.md#GetExternalAccount) | **Get** /external_accounts/{external_account_id} | Get an external account
[**ListExternalAccounts**](ExternalAccountsApi.md#ListExternalAccounts) | **Get** /external_accounts | List external accounts



## AddVendorExternalAccounts

> AddVendorAccountsResponse AddVendorExternalAccounts(ctx).AddVendorAccountsRequest(addVendorAccountsRequest).Execute()

Add external accounts through a vendor, such as Plaid.



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
    addVendorAccountsRequest := *openapiclient.NewAddVendorAccountsRequest("0a50cd1d-afe2-4b85-abfe-d9cfbbad2b98", "CustomerType_example", openapiclient.external_account_vendor_values("PLAID"), "access-sandbox-de3ce8ef-33f8-452c-a685-8671031fc0f6", []string{"blgvvBlXw3cq5GMPwqB6s6q4dLKB9WcVqGDGo"}) // AddVendorAccountsRequest | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ExternalAccountsApi.AddVendorExternalAccounts(context.Background()).AddVendorAccountsRequest(addVendorAccountsRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ExternalAccountsApi.AddVendorExternalAccounts``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddVendorExternalAccounts`: AddVendorAccountsResponse
    fmt.Fprintf(os.Stdout, "Response from `ExternalAccountsApi.AddVendorExternalAccounts`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAddVendorExternalAccountsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **addVendorAccountsRequest** | [**AddVendorAccountsRequest**](AddVendorAccountsRequest.md) |  | 

### Return type

[**AddVendorAccountsResponse**](AddVendorAccountsResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateAccessToken

> ExternalAccountAccessToken CreateAccessToken(ctx).ExternalAccountAccessToken(externalAccountAccessToken).Execute()

Create a permanent access token for an external account

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
    externalAccountAccessToken := *openapiclient.NewExternalAccountAccessToken("VendorCustomerId_example", "VendorPublicToken_example") // ExternalAccountAccessToken | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ExternalAccountsApi.CreateAccessToken(context.Background()).ExternalAccountAccessToken(externalAccountAccessToken).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ExternalAccountsApi.CreateAccessToken``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateAccessToken`: ExternalAccountAccessToken
    fmt.Fprintf(os.Stdout, "Response from `ExternalAccountsApi.CreateAccessToken`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateAccessTokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **externalAccountAccessToken** | [**ExternalAccountAccessToken**](ExternalAccountAccessToken.md) |  | 

### Return type

[**ExternalAccountAccessToken**](ExternalAccountAccessToken.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateVerificationLinkToken

> ExternalAccountLinkToken CreateVerificationLinkToken(ctx).ExternalAccountLinkToken(externalAccountLinkToken).Execute()

Create a link token to verify an external account

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
    externalAccountLinkToken := *openapiclient.NewExternalAccountLinkToken("ClientName_example", []string{"CountryCodes_example"}, *openapiclient.NewExternalAccountUser("927d0e3e-3c91-4c82-9681-2808e503eede"), "Language_example") // ExternalAccountLinkToken | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ExternalAccountsApi.CreateVerificationLinkToken(context.Background()).ExternalAccountLinkToken(externalAccountLinkToken).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ExternalAccountsApi.CreateVerificationLinkToken``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateVerificationLinkToken`: ExternalAccountLinkToken
    fmt.Fprintf(os.Stdout, "Response from `ExternalAccountsApi.CreateVerificationLinkToken`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateVerificationLinkTokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **externalAccountLinkToken** | [**ExternalAccountLinkToken**](ExternalAccountLinkToken.md) |  | 

### Return type

[**ExternalAccountLinkToken**](ExternalAccountLinkToken.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetExternalAccount

> ExternalAccount GetExternalAccount(ctx, externalAccountId).Execute()

Get an external account



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
    externalAccountId := TODO // string | External Account ID

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ExternalAccountsApi.GetExternalAccount(context.Background(), externalAccountId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ExternalAccountsApi.GetExternalAccount``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetExternalAccount`: ExternalAccount
    fmt.Fprintf(os.Stdout, "Response from `ExternalAccountsApi.GetExternalAccount`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**externalAccountId** | [**string**](.md) | External Account ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetExternalAccountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ExternalAccount**](ExternalAccount.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListExternalAccounts

> ExternalAccountsList ListExternalAccounts(ctx).CustomerId(customerId).Limit(limit).PageToken(pageToken).Execute()

List external accounts



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
    customerId := []string{"26c66e9e-50c5-4841-868d-5efbdf95d574"} // []string | A list of customer unique identifiers (optional)
    limit := int32(100) // int32 |  (optional) (default to 100)
    pageToken := "0pqid5u7lx" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ExternalAccountsApi.ListExternalAccounts(context.Background()).CustomerId(customerId).Limit(limit).PageToken(pageToken).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ExternalAccountsApi.ListExternalAccounts``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListExternalAccounts`: ExternalAccountsList
    fmt.Fprintf(os.Stdout, "Response from `ExternalAccountsApi.ListExternalAccounts`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListExternalAccountsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **customerId** | **[]string** | A list of customer unique identifiers | 
 **limit** | **int32** |  | [default to 100]
 **pageToken** | **string** |  | 

### Return type

[**ExternalAccountsList**](ExternalAccountsList.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

