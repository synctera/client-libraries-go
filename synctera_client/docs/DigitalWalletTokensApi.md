# \DigitalWalletTokensApi

All URIs are relative to *https://api.synctera.com/v0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateDigitalWalletApple**](DigitalWalletTokensApi.md#CreateDigitalWalletApple) | **Post** /cards/{card_id}/digital_wallet_tokens/applepay | Create digital wallet token provision request for Apple Pay
[**CreateDigitalWalletGoogle**](DigitalWalletTokensApi.md#CreateDigitalWalletGoogle) | **Post** /cards/{card_id}/digital_wallet_tokens/googlepay | Create digital wallet token provision request for Google Pay
[**GetDigitalWalletToken**](DigitalWalletTokensApi.md#GetDigitalWalletToken) | **Get** /cards/digital_wallet_tokens/{digital_wallet_token_id} | Get Digital Wallet Token
[**ListDigitalWalletTokens**](DigitalWalletTokensApi.md#ListDigitalWalletTokens) | **Get** /cards/digital_wallet_tokens | List Digital Wallet Tokens
[**UpdateDigitalWalletTokenStatus**](DigitalWalletTokensApi.md#UpdateDigitalWalletTokenStatus) | **Patch** /cards/digital_wallet_tokens/{digital_wallet_token_id} | Update Digital Wallet Token&#39;s life cycle status



## CreateDigitalWalletApple

> AppleDigitalWalletProvisionResponse CreateDigitalWalletApple(ctx, cardId).AppleDigitalWalletProvisionRequest(appleDigitalWalletProvisionRequest).Execute()

Create digital wallet token provision request for Apple Pay

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
    cardId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
    appleDigitalWalletProvisionRequest := *openapiclient.NewAppleDigitalWalletProvisionRequest([]string{"Certificates_example"}, openapiclient.device_type("MOBILE_PHONE"), "Nonce_example", "NonceSignature_example", "ProvisioningAppVersion_example") // AppleDigitalWalletProvisionRequest | Request to provision digital wallet card data to pass to Apple Pay digital wallet

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DigitalWalletTokensApi.CreateDigitalWalletApple(context.Background(), cardId).AppleDigitalWalletProvisionRequest(appleDigitalWalletProvisionRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DigitalWalletTokensApi.CreateDigitalWalletApple``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateDigitalWalletApple`: AppleDigitalWalletProvisionResponse
    fmt.Fprintf(os.Stdout, "Response from `DigitalWalletTokensApi.CreateDigitalWalletApple`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cardId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateDigitalWalletAppleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **appleDigitalWalletProvisionRequest** | [**AppleDigitalWalletProvisionRequest**](AppleDigitalWalletProvisionRequest.md) | Request to provision digital wallet card data to pass to Apple Pay digital wallet | 

### Return type

[**AppleDigitalWalletProvisionResponse**](AppleDigitalWalletProvisionResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateDigitalWalletGoogle

> GoogleDigitalWalletProvisionResponse CreateDigitalWalletGoogle(ctx, cardId).GoogleDigitalWalletProvisionRequest(googleDigitalWalletProvisionRequest).Execute()

Create digital wallet token provision request for Google Pay

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
    cardId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
    googleDigitalWalletProvisionRequest := *openapiclient.NewGoogleDigitalWalletProvisionRequest("DeviceId_example", openapiclient.device_type("MOBILE_PHONE"), "ProvisioningAppVersion_example", "WalletAccountId_example") // GoogleDigitalWalletProvisionRequest | Request to provision digital wallet card data to pass to Google Pay digital wallet

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DigitalWalletTokensApi.CreateDigitalWalletGoogle(context.Background(), cardId).GoogleDigitalWalletProvisionRequest(googleDigitalWalletProvisionRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DigitalWalletTokensApi.CreateDigitalWalletGoogle``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateDigitalWalletGoogle`: GoogleDigitalWalletProvisionResponse
    fmt.Fprintf(os.Stdout, "Response from `DigitalWalletTokensApi.CreateDigitalWalletGoogle`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cardId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateDigitalWalletGoogleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **googleDigitalWalletProvisionRequest** | [**GoogleDigitalWalletProvisionRequest**](GoogleDigitalWalletProvisionRequest.md) | Request to provision digital wallet card data to pass to Google Pay digital wallet | 

### Return type

[**GoogleDigitalWalletProvisionResponse**](GoogleDigitalWalletProvisionResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


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
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DigitalWalletTokensApi.GetDigitalWalletToken(context.Background(), digitalWalletTokenId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DigitalWalletTokensApi.GetDigitalWalletToken``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDigitalWalletToken`: DigitalWalletTokenResponse
    fmt.Fprintf(os.Stdout, "Response from `DigitalWalletTokensApi.GetDigitalWalletToken`: %v\n", resp)
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

> TokenListResponse ListDigitalWalletTokens(ctx).CardId(cardId).TokenState(tokenState).Limit(limit).PageToken(pageToken).Execute()

List Digital Wallet Tokens



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
    cardId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string |  (optional)
    tokenState := openapiclient.digital_wallet_token_state("REQUESTED") // DigitalWalletTokenState | The status of the Digital Wallet Token (optional)
    limit := int32(100) // int32 |  (optional) (default to 100)
    pageToken := "h50ffqz9q5" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DigitalWalletTokensApi.ListDigitalWalletTokens(context.Background()).CardId(cardId).TokenState(tokenState).Limit(limit).PageToken(pageToken).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DigitalWalletTokensApi.ListDigitalWalletTokens``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListDigitalWalletTokens`: TokenListResponse
    fmt.Fprintf(os.Stdout, "Response from `DigitalWalletTokensApi.ListDigitalWalletTokens`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListDigitalWalletTokensRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cardId** | **string** |  | 
 **tokenState** | [**DigitalWalletTokenState**](DigitalWalletTokenState.md) | The status of the Digital Wallet Token | 
 **limit** | **int32** |  | [default to 100]
 **pageToken** | **string** |  | 

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

Update Digital Wallet Token's life cycle status



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
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DigitalWalletTokensApi.UpdateDigitalWalletTokenStatus(context.Background(), digitalWalletTokenId).DigitalWalletTokenEditRequest(digitalWalletTokenEditRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DigitalWalletTokensApi.UpdateDigitalWalletTokenStatus``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateDigitalWalletTokenStatus`: DigitalWalletTokenResponse
    fmt.Fprintf(os.Stdout, "Response from `DigitalWalletTokensApi.UpdateDigitalWalletTokenStatus`: %v\n", resp)
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

