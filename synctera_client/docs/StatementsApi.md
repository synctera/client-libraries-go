# \StatementsApi

All URIs are relative to *https://api.synctera.com/v0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetStatement**](StatementsApi.md#GetStatement) | **Get** /statements/{statement_id} | Get a statement
[**ListStatements**](StatementsApi.md#ListStatements) | **Get** /statements | List statements



## GetStatement

> Statement GetStatement(ctx, statementId).Execute()

Get a statement



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
    statementId := "1dac8580-6815-4c9f-be84-ac5518a46832" // string | The unique identifier of a statement

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.StatementsApi.GetStatement(context.Background(), statementId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StatementsApi.GetStatement``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetStatement`: Statement
    fmt.Fprintf(os.Stdout, "Response from `StatementsApi.GetStatement`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**statementId** | **string** | The unique identifier of a statement | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetStatementRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Statement**](Statement.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListStatements

> StatementList ListStatements(ctx).AccountId(accountId).Execute()

List statements



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
    accountId := "e4b92656-ef1e-41f5-a85f-af09346611c0" // string | The account's unique identifier provided by Synctera

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.StatementsApi.ListStatements(context.Background()).AccountId(accountId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StatementsApi.ListStatements``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListStatements`: StatementList
    fmt.Fprintf(os.Stdout, "Response from `StatementsApi.ListStatements`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListStatementsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountId** | **string** | The account&#39;s unique identifier provided by Synctera | 

### Return type

[**StatementList**](StatementList.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

