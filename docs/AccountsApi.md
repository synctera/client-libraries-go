# \AccountsApi

All URIs are relative to *https://api.synctera.com/v0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateAccount**](AccountsApi.md#CreateAccount) | **Post** /accounts | Create an account
[**CreateAccountRelationship**](AccountsApi.md#CreateAccountRelationship) | **Post** /accounts/{account_id}/relationships | Create account relationship
[**CreateAccountResourceProduct**](AccountsApi.md#CreateAccountResourceProduct) | **Post** /accounts/products | Create an account product
[**CreateAccountTemplate**](AccountsApi.md#CreateAccountTemplate) | **Post** /accounts/templates | Create an account template
[**DeleteAccount**](AccountsApi.md#DeleteAccount) | **Delete** /accounts/{account_id} | Delete account
[**DeleteAccountRelationship**](AccountsApi.md#DeleteAccountRelationship) | **Delete** /accounts/{account_id}/relationships/{relationship_id} | Delete account relationship
[**DeleteAccountTemplate**](AccountsApi.md#DeleteAccountTemplate) | **Delete** /accounts/templates/{template_id} | Delete account template
[**GetAccount**](AccountsApi.md#GetAccount) | **Get** /accounts/{account_id} | Get account
[**GetAccountRelationship**](AccountsApi.md#GetAccountRelationship) | **Get** /accounts/{account_id}/relationships/{relationship_id} | Get account relationship
[**GetAccountTemplate**](AccountsApi.md#GetAccountTemplate) | **Get** /accounts/templates/{template_id} | Get account template
[**ListAccountRelationship**](AccountsApi.md#ListAccountRelationship) | **Get** /accounts/{account_id}/relationships | List account relationships
[**ListAccountResourceProducts**](AccountsApi.md#ListAccountResourceProducts) | **Get** /accounts/products | List account products
[**ListAccountTemplates**](AccountsApi.md#ListAccountTemplates) | **Get** /accounts/templates | List account templates
[**ListAccounts**](AccountsApi.md#ListAccounts) | **Get** /accounts | List accounts
[**PatchAccount**](AccountsApi.md#PatchAccount) | **Patch** /accounts/{account_id} | Patch account
[**PatchAccountProduct**](AccountsApi.md#PatchAccountProduct) | **Patch** /accounts/products/{product_id} | Patch account product
[**UpdateAccount**](AccountsApi.md#UpdateAccount) | **Put** /accounts/{account_id} | Update account
[**UpdateAccountRelationship**](AccountsApi.md#UpdateAccountRelationship) | **Put** /accounts/{account_id}/relationships/{relationship_id} | Update account relationship
[**UpdateAccountTemplate**](AccountsApi.md#UpdateAccountTemplate) | **Put** /accounts/templates/{template_id} | Update account template



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
    accountId := TODO // string | Unique identifier for the account.
    relationship := *openapiclient.NewRelationship("CustomerId_example", openapiclient.account_relationship_type("ACCOUNT_HOLDER")) // Relationship | Account relationship object

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
**accountId** | [**string**](.md) | Unique identifier for the account. | 

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


## CreateAccountResourceProduct

> AccountProduct CreateAccountResourceProduct(ctx).AccountProduct(accountProduct).Execute()

Create an account product



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
    accountProduct := openapiclient.account_product{Fee: openapiclient.NewFee(int64(123), "Currency_example", "FeeType_example", "ProductType_example")} // AccountProduct | Account product to create

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AccountsApi.CreateAccountResourceProduct(context.Background()).AccountProduct(accountProduct).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountsApi.CreateAccountResourceProduct``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateAccountResourceProduct`: AccountProduct
    fmt.Fprintf(os.Stdout, "Response from `AccountsApi.CreateAccountResourceProduct`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateAccountResourceProductRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountProduct** | [**AccountProduct**](AccountProduct.md) | Account product to create | 

### Return type

[**AccountProduct**](AccountProduct.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateAccountTemplate

> AccountTemplate CreateAccountTemplate(ctx).AccountTemplate(accountTemplate).Execute()

Create an account template



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
    accountTemplate := *openapiclient.NewAccountTemplate(false, "Name_example", *openapiclient.NewTemplateFields(openapiclient.account_type("SAVING"), "US", "USD")) // AccountTemplate | Account template to create

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AccountsApi.CreateAccountTemplate(context.Background()).AccountTemplate(accountTemplate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountsApi.CreateAccountTemplate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateAccountTemplate`: AccountTemplate
    fmt.Fprintf(os.Stdout, "Response from `AccountsApi.CreateAccountTemplate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateAccountTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountTemplate** | [**AccountTemplate**](AccountTemplate.md) | Account template to create | 

### Return type

[**AccountTemplate**](AccountTemplate.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

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
    accountId := TODO // string | Unique identifier for the account.

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
**accountId** | [**string**](.md) | Unique identifier for the account. | 

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
    accountId := TODO // string | Unique identifier for the account.
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
**accountId** | [**string**](.md) | Unique identifier for the account. | 
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
    resp, r, err := api_client.AccountsApi.DeleteAccountTemplate(context.Background(), templateId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountsApi.DeleteAccountTemplate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteAccountTemplate`: DeleteResponse
    fmt.Fprintf(os.Stdout, "Response from `AccountsApi.DeleteAccountTemplate`: %v\n", resp)
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
    accountId := TODO // string | Unique identifier for the account.

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
**accountId** | [**string**](.md) | Unique identifier for the account. | 

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
    accountId := TODO // string | Unique identifier for the account.
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
**accountId** | [**string**](.md) | Unique identifier for the account. | 
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


## GetAccountTemplate

> AccountTemplate GetAccountTemplate(ctx, templateId).Execute()

Get account template



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
    resp, r, err := api_client.AccountsApi.GetAccountTemplate(context.Background(), templateId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountsApi.GetAccountTemplate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAccountTemplate`: AccountTemplate
    fmt.Fprintf(os.Stdout, "Response from `AccountsApi.GetAccountTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**templateId** | [**string**](.md) | Account Template ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAccountTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AccountTemplate**](AccountTemplate.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

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
    accountId := TODO // string | Unique identifier for the account.
    limit := int32(100) // int32 |  (optional) (default to 100)
    pageToken := "faker.random.alphaNumeric(10)" // string |  (optional)

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
**accountId** | [**string**](.md) | Unique identifier for the account. | 

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


## ListAccountResourceProducts

> AccountProductList ListAccountResourceProducts(ctx).ProductType(productType).Limit(limit).PageToken(pageToken).StartDate(startDate).EndDate(endDate).Execute()

List account products



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
    productType := "productType_example" // string | Type of account product
    limit := int32(100) // int32 |  (optional) (default to 100)
    pageToken := "faker.random.alphaNumeric(10)" // string |  (optional)
    startDate := time.Now() // oapi.Date | Date range filtering for type INTEREST. All rates in interest resource have to have valid_from later or equal to start_date. (optional)
    endDate := time.Now() // oapi.Date | Date range filtering for type INTEREST. All rates in interest resource have to have valid_to earlier or equal to end_date. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AccountsApi.ListAccountResourceProducts(context.Background()).ProductType(productType).Limit(limit).PageToken(pageToken).StartDate(startDate).EndDate(endDate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountsApi.ListAccountResourceProducts``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListAccountResourceProducts`: AccountProductList
    fmt.Fprintf(os.Stdout, "Response from `AccountsApi.ListAccountResourceProducts`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListAccountResourceProductsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **productType** | **string** | Type of account product | 
 **limit** | **int32** |  | [default to 100]
 **pageToken** | **string** |  | 
 **startDate** | **oapi.Date** | Date range filtering for type INTEREST. All rates in interest resource have to have valid_from later or equal to start_date. | 
 **endDate** | **oapi.Date** | Date range filtering for type INTEREST. All rates in interest resource have to have valid_to earlier or equal to end_date. | 

### Return type

[**AccountProductList**](AccountProductList.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListAccountTemplates

> TemplateList ListAccountTemplates(ctx).Limit(limit).PageToken(pageToken).Execute()

List account templates



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
    pageToken := "faker.random.alphaNumeric(10)" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AccountsApi.ListAccountTemplates(context.Background()).Limit(limit).PageToken(pageToken).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountsApi.ListAccountTemplates``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListAccountTemplates`: TemplateList
    fmt.Fprintf(os.Stdout, "Response from `AccountsApi.ListAccountTemplates`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListAccountTemplatesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** |  | [default to 100]
 **pageToken** | **string** |  | 

### Return type

[**TemplateList**](TemplateList.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListAccounts

> AccountList ListAccounts(ctx).Id(id).OverdraftAccountId(overdraftAccountId).OverflowAccountId(overflowAccountId).AccountNumber(accountNumber).Status(status).InterestProductId(interestProductId).CustomerId(customerId).FirstName(firstName).LastName(lastName).Limit(limit).PageToken(pageToken).SortBy(sortBy).Execute()

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
    id := []string{"Inner_example"} // []string | Account ID(s). Multiple IDs can be provided as a comma-separated list.  (optional)
    overdraftAccountId := []string{"Inner_example"} // []string | Overdraft account ID(s). Multiple IDs can be provided as a comma-separated list.  (optional)
    overflowAccountId := []string{"Inner_example"} // []string | Overflow account ID(s). Multiple IDs can be provided as a comma-separated list.  (optional)
    accountNumber := []string{"Inner_example"} // []string | Account number(s). Multiple account numbers can be provided as a comma-separated list. When only a single account number is provided, any * characters in the string are wildcards, and match any characters.  (optional)
    status := openapiclient.status("APPLICATION_SUBMITTED") // Status |  (optional)
    interestProductId := "interestProductId_example" // string | Interest product ID that accounts associate with. Multiple IDs can be provided as a comma-separated list. (optional)
    customerId := TODO // string | The customer's unique identifier (optional)
    firstName := "Alice" // string |  (optional)
    lastName := "Smith" // string |  (optional)
    limit := int32(100) // int32 |  (optional) (default to 100)
    pageToken := "faker.random.alphaNumeric(10)" // string |  (optional)
    sortBy := []string{"SortBy_example"} // []string | Specifies the sort order for the returned accounts.  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AccountsApi.ListAccounts(context.Background()).Id(id).OverdraftAccountId(overdraftAccountId).OverflowAccountId(overflowAccountId).AccountNumber(accountNumber).Status(status).InterestProductId(interestProductId).CustomerId(customerId).FirstName(firstName).LastName(lastName).Limit(limit).PageToken(pageToken).SortBy(sortBy).Execute()
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
 **id** | **[]string** | Account ID(s). Multiple IDs can be provided as a comma-separated list.  | 
 **overdraftAccountId** | **[]string** | Overdraft account ID(s). Multiple IDs can be provided as a comma-separated list.  | 
 **overflowAccountId** | **[]string** | Overflow account ID(s). Multiple IDs can be provided as a comma-separated list.  | 
 **accountNumber** | **[]string** | Account number(s). Multiple account numbers can be provided as a comma-separated list. When only a single account number is provided, any * characters in the string are wildcards, and match any characters.  | 
 **status** | [**Status**](Status.md) |  | 
 **interestProductId** | **string** | Interest product ID that accounts associate with. Multiple IDs can be provided as a comma-separated list. | 
 **customerId** | [**string**](string.md) | The customer&#39;s unique identifier | 
 **firstName** | **string** |  | 
 **lastName** | **string** |  | 
 **limit** | **int32** |  | [default to 100]
 **pageToken** | **string** |  | 
 **sortBy** | **[]string** | Specifies the sort order for the returned accounts.  | 

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
    accountId := TODO // string | Unique identifier for the account.
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
**accountId** | [**string**](.md) | Unique identifier for the account. | 

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


## PatchAccountProduct

> AccountProduct PatchAccountProduct(ctx, productId).PatchAccountProduct(patchAccountProduct).Execute()

Patch account product



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
    productId := TODO // string | Account Product ID
    patchAccountProduct := openapiclient.patch_account_product{Fee: openapiclient.NewFee(int64(123), "Currency_example", "FeeType_example", "ProductType_example")} // PatchAccountProduct | Account product fields to be patched

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AccountsApi.PatchAccountProduct(context.Background(), productId).PatchAccountProduct(patchAccountProduct).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountsApi.PatchAccountProduct``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PatchAccountProduct`: AccountProduct
    fmt.Fprintf(os.Stdout, "Response from `AccountsApi.PatchAccountProduct`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**productId** | [**string**](.md) | Account Product ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchAccountProductRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **patchAccountProduct** | [**PatchAccountProduct**](PatchAccountProduct.md) | Account product fields to be patched | 

### Return type

[**AccountProduct**](AccountProduct.md)

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
    accountId := TODO // string | Unique identifier for the account.
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
**accountId** | [**string**](.md) | Unique identifier for the account. | 

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
    accountId := TODO // string | Unique identifier for the account.
    relationshipId := TODO // string | Relationship ID of the account associate with the account entity
    relationship := *openapiclient.NewRelationship("CustomerId_example", openapiclient.account_relationship_type("ACCOUNT_HOLDER")) // Relationship | Account relationship to be updated

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
**accountId** | [**string**](.md) | Unique identifier for the account. | 
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


## UpdateAccountTemplate

> AccountTemplate UpdateAccountTemplate(ctx, templateId).AccountTemplate(accountTemplate).Execute()

Update account template



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
    accountTemplate := *openapiclient.NewAccountTemplate(false, "Name_example", *openapiclient.NewTemplateFields(openapiclient.account_type("SAVING"), "US", "USD")) // AccountTemplate | Account template to update

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AccountsApi.UpdateAccountTemplate(context.Background(), templateId).AccountTemplate(accountTemplate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountsApi.UpdateAccountTemplate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateAccountTemplate`: AccountTemplate
    fmt.Fprintf(os.Stdout, "Response from `AccountsApi.UpdateAccountTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**templateId** | [**string**](.md) | Account Template ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateAccountTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **accountTemplate** | [**AccountTemplate**](AccountTemplate.md) | Account template to update | 

### Return type

[**AccountTemplate**](AccountTemplate.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

