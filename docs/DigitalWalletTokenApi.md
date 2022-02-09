# \DigitalWalletTokenApi

All URIs are relative to *https://api.synctera.com/v0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetDigitalWalletToken**](DigitalWalletTokenApi.md#GetDigitalWalletToken) | **Get** /cards/digitalwallets/{digital_wallet_token_id} | Get Digital Wallet Token
[**ListDigitalWalletTokens**](DigitalWalletTokenApi.md#ListDigitalWalletTokens) | **Get** /cards/{card_id}/digitalwallets | List Digital Wallet Tokens of a card
[**UpdateDigitalWalletTokenStatus**](DigitalWalletTokenApi.md#UpdateDigitalWalletTokenStatus) | **Patch** /cards/digitalwallets/{digital_wallet_token_id} | Update Digital wallet token id&#39;s life cycle status



## GetDigitalWalletToken

> DigitalWalletTokenResponse GetDigitalWalletToken(ctx, digitalWalletTokenId).Execute()

Get Digital Wallet Token



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
    digitalWalletTokenId := "digitalWalletTokenId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DigitalWalletTokenApi.GetDigitalWalletToken(context.Background(), digitalWalletTokenId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DigitalWalletTokenApi.GetDigitalWalletToken``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDigitalWalletToken`: DigitalWalletTokenResponse
    fmt.Fprintf(os.Stdout, "Response from `DigitalWalletTokenApi.GetDigitalWalletToken`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**digitalWalletTokenId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDigitalWalletTokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DigitalWalletTokenResponse**](DigitalWalletTokenResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListDigitalWalletTokens

> TokenListResponse ListDigitalWalletTokens(ctx, cardId).Execute()

List Digital Wallet Tokens of a card



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
    cardId := TODO // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DigitalWalletTokenApi.ListDigitalWalletTokens(context.Background(), cardId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DigitalWalletTokenApi.ListDigitalWalletTokens``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListDigitalWalletTokens`: TokenListResponse
    fmt.Fprintf(os.Stdout, "Response from `DigitalWalletTokenApi.ListDigitalWalletTokens`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cardId** | [**string**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListDigitalWalletTokensRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**TokenListResponse**](TokenListResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateDigitalWalletTokenStatus

> DigitalWalletTokenResponse UpdateDigitalWalletTokenStatus(ctx, digitalWalletTokenId).DigitalWalletTokenEditRequest(digitalWalletTokenEditRequest).Execute()

Update Digital wallet token id's life cycle status



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
    digitalWalletTokenId := "digitalWalletTokenId_example" // string | 
    digitalWalletTokenEditRequest := *openapiclient.NewDigitalWalletTokenEditRequest("TokenStatus_example") // DigitalWalletTokenEditRequest | Update Digital wallet token status

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DigitalWalletTokenApi.UpdateDigitalWalletTokenStatus(context.Background(), digitalWalletTokenId).DigitalWalletTokenEditRequest(digitalWalletTokenEditRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DigitalWalletTokenApi.UpdateDigitalWalletTokenStatus``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateDigitalWalletTokenStatus`: DigitalWalletTokenResponse
    fmt.Fprintf(os.Stdout, "Response from `DigitalWalletTokenApi.UpdateDigitalWalletTokenStatus`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**digitalWalletTokenId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateDigitalWalletTokenStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **digitalWalletTokenEditRequest** | [**DigitalWalletTokenEditRequest**](DigitalWalletTokenEditRequest.md) | Update Digital wallet token status | 

### Return type

[**DigitalWalletTokenResponse**](DigitalWalletTokenResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

