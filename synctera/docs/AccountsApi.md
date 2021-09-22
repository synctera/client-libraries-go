# \AccountsApi

All URIs are relative to *https://api.synctera.com/v0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateAccount**](AccountsApi.md#CreateAccount) | **Post** /accounts | Create an account
[**CreateAccountAddress**](AccountsApi.md#CreateAccountAddress) | **Post** /accounts/{account_id}/addresses | Create address for an account
[**CreateAccountRelationship**](AccountsApi.md#CreateAccountRelationship) | **Post** /accounts/{account_id}/relationships | Create account relationship
[**DeleteAccount**](AccountsApi.md#DeleteAccount) | **Delete** /accounts/{account_id} | Delete account
[**DeleteAccountAddress**](AccountsApi.md#DeleteAccountAddress) | **Delete** /accounts/{account_id}/addresses/{connect_id} | Delete address for an account
[**DeleteAccountRelationship**](AccountsApi.md#DeleteAccountRelationship) | **Delete** /accounts/{account_id}/relationships/{relationship_id} | Delete account relationship
[**GetAccount**](AccountsApi.md#GetAccount) | **Get** /accounts/{account_id} | Get account
[**GetAccountAddress**](AccountsApi.md#GetAccountAddress) | **Get** /accounts/{account_id}/addresses/{connect_id} | Get address for an account
[**GetAccountBalance**](AccountsApi.md#GetAccountBalance) | **Get** /accounts/{account_id}/balance | Get account balance
[**GetAccountRelationship**](AccountsApi.md#GetAccountRelationship) | **Get** /accounts/{account_id}/relationships/{relationship_id} | Get account relationship
[**GetAccountTransactions**](AccountsApi.md#GetAccountTransactions) | **Get** /accounts/{account_id}/transactions | Get account transactions
[**ListAccountRelationship**](AccountsApi.md#ListAccountRelationship) | **Get** /accounts/{account_id}/relationships | List account relationships
[**ListAccounts**](AccountsApi.md#ListAccounts) | **Get** /accounts | List accounts
[**PatchAccount**](AccountsApi.md#PatchAccount) | **Patch** /accounts/{account_id} | Patch account
[**UpdateAccount**](AccountsApi.md#UpdateAccount) | **Put** /accounts/{account_id} | Update account
[**UpdateAccountAddress**](AccountsApi.md#UpdateAccountAddress) | **Put** /accounts/{account_id}/addresses/{connect_id} | Update address for an account
[**UpdateAccountRelationship**](AccountsApi.md#UpdateAccountRelationship) | **Put** /accounts/{account_id}/relationships/{relationship_id} | Update account relationship



## CreateAccount

> Account CreateAccount(ctx).AccountCreation(accountCreation).Execute()

Create an account



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
    accountCreation := *openapiclient.NewAccountCreation() // AccountCreation | Account to create

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AccountsApi.CreateAccount(context.Background()).AccountCreation(accountCreation).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountsApi.CreateAccount``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateAccount`: Account
    fmt.Fprintf(os.Stdout, "Response from `AccountsApi.CreateAccount`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateAccountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountCreation** | [**AccountCreation**](AccountCreation.md) | Account to create | 

### Return type

[**Account**](Account.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateAccountAddress

> AccountAddress CreateAccountAddress(ctx, accountId).AccountAddress(accountAddress).Execute()

Create address for an account



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
    accountId := TODO // string | Account ID
    accountAddress := *openapiclient.NewAccountAddress() // AccountAddress | Account address to create

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AccountsApi.CreateAccountAddress(context.Background(), accountId).AccountAddress(accountAddress).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountsApi.CreateAccountAddress``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateAccountAddress`: AccountAddress
    fmt.Fprintf(os.Stdout, "Response from `AccountsApi.CreateAccountAddress`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | [**string**](.md) | Account ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateAccountAddressRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **accountAddress** | [**AccountAddress**](AccountAddress.md) | Account address to create | 

### Return type

[**AccountAddress**](AccountAddress.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateAccountRelationship

> Relationship CreateAccountRelationship(ctx, accountId).Relationship(relationship).Execute()

Create account relationship



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
    accountId := TODO // string | Account ID
    relationship := *openapiclient.NewRelationship("CustomerId_example", "RelationshipType_example") // Relationship | Account relationship object

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AccountsApi.CreateAccountRelationship(context.Background(), accountId).Relationship(relationship).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountsApi.CreateAccountRelationship``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateAccountRelationship`: Relationship
    fmt.Fprintf(os.Stdout, "Response from `AccountsApi.CreateAccountRelationship`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | [**string**](.md) | Account ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateAccountRelationshipRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **relationship** | [**Relationship**](Relationship.md) | Account relationship object | 

### Return type

[**Relationship**](Relationship.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteAccount

> DeleteResponse DeleteAccount(ctx, accountId).Execute()

Delete account



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
    accountId := TODO // string | Account ID

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AccountsApi.DeleteAccount(context.Background(), accountId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountsApi.DeleteAccount``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteAccount`: DeleteResponse
    fmt.Fprintf(os.Stdout, "Response from `AccountsApi.DeleteAccount`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | [**string**](.md) | Account ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAccountRequest struct via the builder pattern


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


## DeleteAccountAddress

> DeleteResponse DeleteAccountAddress(ctx, accountId, connectId).Execute()

Delete address for an account



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
    accountId := TODO // string | Account ID
    connectId := TODO // string | Connect ID of the account associate with the account entity

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AccountsApi.DeleteAccountAddress(context.Background(), accountId, connectId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountsApi.DeleteAccountAddress``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteAccountAddress`: DeleteResponse
    fmt.Fprintf(os.Stdout, "Response from `AccountsApi.DeleteAccountAddress`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | [**string**](.md) | Account ID | 
**connectId** | [**string**](.md) | Connect ID of the account associate with the account entity | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAccountAddressRequest struct via the builder pattern


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


## DeleteAccountRelationship

> DeleteResponse DeleteAccountRelationship(ctx, accountId, relationshipId).Execute()

Delete account relationship



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
    accountId := TODO // string | Account ID
    relationshipId := TODO // string | Relationship ID of the account associate with the account entity

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AccountsApi.DeleteAccountRelationship(context.Background(), accountId, relationshipId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountsApi.DeleteAccountRelationship``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteAccountRelationship`: DeleteResponse
    fmt.Fprintf(os.Stdout, "Response from `AccountsApi.DeleteAccountRelationship`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | [**string**](.md) | Account ID | 
**relationshipId** | [**string**](.md) | Relationship ID of the account associate with the account entity | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAccountRelationshipRequest struct via the builder pattern


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


## GetAccount

> Account GetAccount(ctx, accountId).Execute()

Get account



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
    accountId := TODO // string | Account ID

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AccountsApi.GetAccount(context.Background(), accountId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountsApi.GetAccount``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAccount`: Account
    fmt.Fprintf(os.Stdout, "Response from `AccountsApi.GetAccount`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | [**string**](.md) | Account ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAccountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Account**](Account.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAccountAddress

> AccountAddress GetAccountAddress(ctx, accountId, connectId).InlineObject(inlineObject).Execute()

Get address for an account



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
    accountId := TODO // string | Account ID
    connectId := TODO // string | Connect ID of the account associate with the account entity
    inlineObject := *openapiclient.NewInlineObject() // InlineObject |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AccountsApi.GetAccountAddress(context.Background(), accountId, connectId).InlineObject(inlineObject).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountsApi.GetAccountAddress``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAccountAddress`: AccountAddress
    fmt.Fprintf(os.Stdout, "Response from `AccountsApi.GetAccountAddress`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | [**string**](.md) | Account ID | 
**connectId** | [**string**](.md) | Connect ID of the account associate with the account entity | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAccountAddressRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **inlineObject** | [**InlineObject**](InlineObject.md) |  | 

### Return type

[**AccountAddress**](AccountAddress.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAccountBalance

> Balance GetAccountBalance(ctx, accountId).InlineObject1(inlineObject1).Execute()

Get account balance



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
    accountId := TODO // string | Account ID
    inlineObject1 := *openapiclient.NewInlineObject1() // InlineObject1 | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AccountsApi.GetAccountBalance(context.Background(), accountId).InlineObject1(inlineObject1).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountsApi.GetAccountBalance``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAccountBalance`: Balance
    fmt.Fprintf(os.Stdout, "Response from `AccountsApi.GetAccountBalance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | [**string**](.md) | Account ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAccountBalanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **inlineObject1** | [**InlineObject1**](InlineObject1.md) |  | 

### Return type

[**Balance**](Balance.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAccountRelationship

> Relationship GetAccountRelationship(ctx, accountId, relationshipId).Execute()

Get account relationship



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
    accountId := TODO // string | Account ID
    relationshipId := TODO // string | Relationship ID of the account associate with the account entity

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AccountsApi.GetAccountRelationship(context.Background(), accountId, relationshipId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountsApi.GetAccountRelationship``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAccountRelationship`: Relationship
    fmt.Fprintf(os.Stdout, "Response from `AccountsApi.GetAccountRelationship`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | [**string**](.md) | Account ID | 
**relationshipId** | [**string**](.md) | Relationship ID of the account associate with the account entity | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAccountRelationshipRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**Relationship**](Relationship.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAccountTransactions

> TransactionList GetAccountTransactions(ctx, accountId).InlineObject2(inlineObject2).Limit(limit).PageToken(pageToken).Execute()

Get account transactions



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
    accountId := TODO // string | Account ID
    inlineObject2 := *openapiclient.NewInlineObject2() // InlineObject2 | 
    limit := int32(100) // int32 |  (optional) (default to 100)
    pageToken := "bnw3qvoyid" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AccountsApi.GetAccountTransactions(context.Background(), accountId).InlineObject2(inlineObject2).Limit(limit).PageToken(pageToken).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountsApi.GetAccountTransactions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAccountTransactions`: TransactionList
    fmt.Fprintf(os.Stdout, "Response from `AccountsApi.GetAccountTransactions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | [**string**](.md) | Account ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAccountTransactionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **inlineObject2** | [**InlineObject2**](InlineObject2.md) |  | 
 **limit** | **int32** |  | [default to 100]
 **pageToken** | **string** |  | 

### Return type

[**TransactionList**](TransactionList.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListAccountRelationship

> RelationshipList ListAccountRelationship(ctx, accountId).Limit(limit).PageToken(pageToken).Execute()

List account relationships



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
    accountId := TODO // string | Account ID
    limit := int32(100) // int32 |  (optional) (default to 100)
    pageToken := "bnw3qvoyid" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AccountsApi.ListAccountRelationship(context.Background(), accountId).Limit(limit).PageToken(pageToken).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountsApi.ListAccountRelationship``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListAccountRelationship`: RelationshipList
    fmt.Fprintf(os.Stdout, "Response from `AccountsApi.ListAccountRelationship`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | [**string**](.md) | Account ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiListAccountRelationshipRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** |  | [default to 100]
 **pageToken** | **string** |  | 

### Return type

[**RelationshipList**](RelationshipList.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListAccounts

> AccountList ListAccounts(ctx).Tenant(tenant).Limit(limit).PageToken(pageToken).CustomerId(customerId).HasDetails(hasDetails).Execute()

List accounts



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
    tenant := "2_3" // string |  (optional)
    limit := int32(100) // int32 |  (optional) (default to 100)
    pageToken := "bnw3qvoyid" // string |  (optional)
    customerId := TODO // string | Customer ID (optional)
    hasDetails := true // bool | If true, the returned accounts will include relationships, aliases, balances and recent transactions. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AccountsApi.ListAccounts(context.Background()).Tenant(tenant).Limit(limit).PageToken(pageToken).CustomerId(customerId).HasDetails(hasDetails).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountsApi.ListAccounts``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListAccounts`: AccountList
    fmt.Fprintf(os.Stdout, "Response from `AccountsApi.ListAccounts`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListAccountsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenant** | **string** |  | 
 **limit** | **int32** |  | [default to 100]
 **pageToken** | **string** |  | 
 **customerId** | [**string**](string.md) | Customer ID | 
 **hasDetails** | **bool** | If true, the returned accounts will include relationships, aliases, balances and recent transactions. | 

### Return type

[**AccountList**](AccountList.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchAccount

> Account PatchAccount(ctx, accountId).Account(account).Execute()

Patch account



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
    accountId := TODO // string | Account ID
    account := *openapiclient.NewAccount() // Account | Account fields to be patched

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AccountsApi.PatchAccount(context.Background(), accountId).Account(account).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountsApi.PatchAccount``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PatchAccount`: Account
    fmt.Fprintf(os.Stdout, "Response from `AccountsApi.PatchAccount`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | [**string**](.md) | Account ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchAccountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **account** | [**Account**](Account.md) | Account fields to be patched | 

### Return type

[**Account**](Account.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateAccount

> Account UpdateAccount(ctx, accountId).Account(account).Execute()

Update account



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
    accountId := TODO // string | Account ID
    account := *openapiclient.NewAccount() // Account | Account to update

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AccountsApi.UpdateAccount(context.Background(), accountId).Account(account).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountsApi.UpdateAccount``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateAccount`: Account
    fmt.Fprintf(os.Stdout, "Response from `AccountsApi.UpdateAccount`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | [**string**](.md) | Account ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateAccountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **account** | [**Account**](Account.md) | Account to update | 

### Return type

[**Account**](Account.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateAccountAddress

> AccountAddress UpdateAccountAddress(ctx, accountId, connectId).AccountAddress(accountAddress).Execute()

Update address for an account



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
    accountId := TODO // string | Account ID
    connectId := TODO // string | Connect ID of the account associate with the account entity
    accountAddress := *openapiclient.NewAccountAddress() // AccountAddress | Account address to be updated

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AccountsApi.UpdateAccountAddress(context.Background(), accountId, connectId).AccountAddress(accountAddress).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountsApi.UpdateAccountAddress``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateAccountAddress`: AccountAddress
    fmt.Fprintf(os.Stdout, "Response from `AccountsApi.UpdateAccountAddress`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | [**string**](.md) | Account ID | 
**connectId** | [**string**](.md) | Connect ID of the account associate with the account entity | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateAccountAddressRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **accountAddress** | [**AccountAddress**](AccountAddress.md) | Account address to be updated | 

### Return type

[**AccountAddress**](AccountAddress.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateAccountRelationship

> Relationship UpdateAccountRelationship(ctx, accountId, relationshipId).Relationship(relationship).Execute()

Update account relationship



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
    accountId := TODO // string | Account ID
    relationshipId := TODO // string | Relationship ID of the account associate with the account entity
    relationship := *openapiclient.NewRelationship("CustomerId_example", "RelationshipType_example") // Relationship | Account relationship to be updated

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AccountsApi.UpdateAccountRelationship(context.Background(), accountId, relationshipId).Relationship(relationship).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountsApi.UpdateAccountRelationship``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateAccountRelationship`: Relationship
    fmt.Fprintf(os.Stdout, "Response from `AccountsApi.UpdateAccountRelationship`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | [**string**](.md) | Account ID | 
**relationshipId** | [**string**](.md) | Relationship ID of the account associate with the account entity | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateAccountRelationshipRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **relationship** | [**Relationship**](Relationship.md) | Account relationship to be updated | 

### Return type

[**Relationship**](Relationship.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

