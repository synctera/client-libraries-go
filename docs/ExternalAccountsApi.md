# \ExternalAccountsApi

All URIs are relative to *https://api.synctera.com/v0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddExternalAccounts**](ExternalAccountsApi.md#AddExternalAccounts) | **Post** /external_accounts | Add external accounts
[**AddVendorExternalAccounts**](ExternalAccountsApi.md#AddVendorExternalAccounts) | **Post** /external_accounts/add_vendor_accounts | Add external accounts through a vendor, such as Plaid.
[**CreateAccessToken**](ExternalAccountsApi.md#CreateAccessToken) | **Post** /external_accounts/access_tokens | Create a permanent access token for an external account
[**CreateVerificationLinkToken**](ExternalAccountsApi.md#CreateVerificationLinkToken) | **Post** /external_accounts/link_tokens | Create a link token to verify an external account
[**GetExternalAccount**](ExternalAccountsApi.md#GetExternalAccount) | **Get** /external_accounts/{external_account_id} | Get an external account
[**GetExternalAccountBalance**](ExternalAccountsApi.md#GetExternalAccountBalance) | **Get** /external_accounts/{external_account_id}/balance | Get an external account balance
[**GetExternalAccountTransactions**](ExternalAccountsApi.md#GetExternalAccountTransactions) | **Get** /external_accounts/{external_account_id}/transactions | List transactions of a given external account
[**ListExternalAccounts**](ExternalAccountsApi.md#ListExternalAccounts) | **Get** /external_accounts | List external accounts
[**UpdateExternalAccount**](ExternalAccountsApi.md#UpdateExternalAccount) | **Patch** /external_accounts/{external_account_id} | Patch an external account



## AddExternalAccounts

> ExternalAccount AddExternalAccounts(ctx).AddAccountsRequest(addAccountsRequest).Execute()

Add external accounts



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
    addAccountsRequest := *openapiclient.NewAddAccountsRequest(*openapiclient.NewAddAccountsRequestAccountIdentifiers("78277121"), []string{"faker.name.findName"}, "CustomerId_example", "CustomerType_example", *openapiclient.NewAddAccountsRequestRoutingIdentifiers("756392185", []string{"BankCountries_example"}, "Chase"), "Type_example") // AddAccountsRequest | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ExternalAccountsApi.AddExternalAccounts(context.Background()).AddAccountsRequest(addAccountsRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ExternalAccountsApi.AddExternalAccounts``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddExternalAccounts`: ExternalAccount
    fmt.Fprintf(os.Stdout, "Response from `ExternalAccountsApi.AddExternalAccounts`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAddExternalAccountsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **addAccountsRequest** | [**AddAccountsRequest**](AddAccountsRequest.md) |  | 

### Return type

[**ExternalAccount**](ExternalAccount.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


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
    addVendorAccountsRequest := *openapiclient.NewAddVendorAccountsRequest("CustomerId_example", "CustomerType_example", openapiclient.external_account_vendor_values("PLAID"), "access-sandbox-de3ce8ef-33f8-452c-a685-8671031fc0f6", []string{"blgvvBlXw3cq5GMPwqB6s6q4dLKB9WcVqGDGo"}) // AddVendorAccountsRequest | 

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
    externalAccountAccessToken := *openapiclient.NewExternalAccountAccessToken("VendorInstitutionId_example", "VendorPublicToken_example") // ExternalAccountAccessToken | 

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
    externalAccountLinkToken := *openapiclient.NewExternalAccountLinkToken("ClientName_example", []string{"CountryCodes_example"}, "CustomerId_example", "Language_example", "Type_example") // ExternalAccountLinkToken | 

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


## GetExternalAccountBalance

> ExternalAccountBalance GetExternalAccountBalance(ctx, externalAccountId).Execute()

Get an external account balance



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
    resp, r, err := api_client.ExternalAccountsApi.GetExternalAccountBalance(context.Background(), externalAccountId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ExternalAccountsApi.GetExternalAccountBalance``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetExternalAccountBalance`: ExternalAccountBalance
    fmt.Fprintf(os.Stdout, "Response from `ExternalAccountsApi.GetExternalAccountBalance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**externalAccountId** | [**string**](.md) | External Account ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetExternalAccountBalanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ExternalAccountBalance**](ExternalAccountBalance.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetExternalAccountTransactions

> ExternalAccountsTransactionList GetExternalAccountTransactions(ctx, externalAccountId).StartDate(startDate).EndDate(endDate).Execute()

List transactions of a given external account



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
    externalAccountId := TODO // string | External Account ID
    startDate := time.Now() // oapi.Date | Date range filtering for transactions. Date is inclusive.
    endDate := time.Now() // oapi.Date | Date range filtering for transactions. Date is exclusive.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ExternalAccountsApi.GetExternalAccountTransactions(context.Background(), externalAccountId).StartDate(startDate).EndDate(endDate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ExternalAccountsApi.GetExternalAccountTransactions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetExternalAccountTransactions`: ExternalAccountsTransactionList
    fmt.Fprintf(os.Stdout, "Response from `ExternalAccountsApi.GetExternalAccountTransactions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**externalAccountId** | [**string**](.md) | External Account ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetExternalAccountTransactionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **startDate** | **oapi.Date** | Date range filtering for transactions. Date is inclusive. | 
 **endDate** | **oapi.Date** | Date range filtering for transactions. Date is exclusive. | 

### Return type

[**ExternalAccountsTransactionList**](ExternalAccountsTransactionList.md)

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
    customerId := []string{"Inner_example"} // []string | A list of customer unique identifiers, with a comma separating any values. (optional)
    limit := int32(100) // int32 |  (optional) (default to 100)
    pageToken := "faker.random.alphaNumeric(10)" // string |  (optional)

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
 **customerId** | **[]string** | A list of customer unique identifiers, with a comma separating any values. | 
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


## UpdateExternalAccount

> ExternalAccount UpdateExternalAccount(ctx, externalAccountId).PatchExternalAccount(patchExternalAccount).Execute()

Patch an external account



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
    patchExternalAccount := *openapiclient.NewPatchExternalAccount() // PatchExternalAccount | External account to be updated

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ExternalAccountsApi.UpdateExternalAccount(context.Background(), externalAccountId).PatchExternalAccount(patchExternalAccount).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ExternalAccountsApi.UpdateExternalAccount``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateExternalAccount`: ExternalAccount
    fmt.Fprintf(os.Stdout, "Response from `ExternalAccountsApi.UpdateExternalAccount`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**externalAccountId** | [**string**](.md) | External Account ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateExternalAccountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **patchExternalAccount** | [**PatchExternalAccount**](PatchExternalAccount.md) | External account to be updated | 

### Return type

[**ExternalAccount**](ExternalAccount.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)
