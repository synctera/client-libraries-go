# \CardsApi

All URIs are relative to *https://api.synctera.com/v0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ActivateCard**](CardsApi.md#ActivateCard) | **Post** /cards/activate | Activate a card
[**GetCard**](CardsApi.md#GetCard) | **Get** /cards/{card_id} | Get Card
[**GetClientAccessToken**](CardsApi.md#GetClientAccessToken) | **Post** /cards/{card_id}/client_token | Get a client token
[**GetClientSingleUseToken**](CardsApi.md#GetClientSingleUseToken) | **Post** /cards/single_use_token | Get single-use token
[**IssueCard**](CardsApi.md#IssueCard) | **Post** /cards | Issue a Card
[**ListCards**](CardsApi.md#ListCards) | **Get** /cards | List Cards
[**ListChanges**](CardsApi.md#ListChanges) | **Get** /cards/{card_id}/changes | List Card Changes
[**UpdateCard**](CardsApi.md#UpdateCard) | **Patch** /cards/{card_id} | Update Card



## ActivateCard

> CardResponse ActivateCard(ctx).CardActivationRequest(cardActivationRequest).Execute()

Activate a card



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
    cardActivationRequest := *openapiclient.NewCardActivationRequest("ActivationCode_example", "fe1fa42f-10e3-48d2-ab7d-e3324950a42e") // CardActivationRequest | Card activation code

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CardsApi.ActivateCard(context.Background()).CardActivationRequest(cardActivationRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CardsApi.ActivateCard``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ActivateCard`: CardResponse
    fmt.Fprintf(os.Stdout, "Response from `CardsApi.ActivateCard`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiActivateCardRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cardActivationRequest** | [**CardActivationRequest**](CardActivationRequest.md) | Card activation code | 

### Return type

[**CardResponse**](CardResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCard

> CardResponse GetCard(ctx, cardId).Execute()

Get Card



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
    resp, r, err := api_client.CardsApi.GetCard(context.Background(), cardId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CardsApi.GetCard``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCard`: CardResponse
    fmt.Fprintf(os.Stdout, "Response from `CardsApi.GetCard`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cardId** | [**string**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCardRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CardResponse**](CardResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetClientAccessToken

> ClientToken GetClientAccessToken(ctx, cardId).Execute()

Get a client token



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
    resp, r, err := api_client.CardsApi.GetClientAccessToken(context.Background(), cardId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CardsApi.GetClientAccessToken``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetClientAccessToken`: ClientToken
    fmt.Fprintf(os.Stdout, "Response from `CardsApi.GetClientAccessToken`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cardId** | [**string**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetClientAccessTokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ClientToken**](ClientToken.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetClientSingleUseToken

> SingleUseTokenResponse GetClientSingleUseToken(ctx).SingleUseTokenRequest(singleUseTokenRequest).Execute()

Get single-use token



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
    singleUseTokenRequest := *openapiclient.NewSingleUseTokenRequest("409302b6-a448-4697-89bc-bacb678f103f", "538f0f18-df33-4903-8d5a-bc0f719d3ef1") // SingleUseTokenRequest | User token details

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CardsApi.GetClientSingleUseToken(context.Background()).SingleUseTokenRequest(singleUseTokenRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CardsApi.GetClientSingleUseToken``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetClientSingleUseToken`: SingleUseTokenResponse
    fmt.Fprintf(os.Stdout, "Response from `CardsApi.GetClientSingleUseToken`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetClientSingleUseTokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **singleUseTokenRequest** | [**SingleUseTokenRequest**](SingleUseTokenRequest.md) | User token details | 

### Return type

[**SingleUseTokenResponse**](SingleUseTokenResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IssueCard

> CardResponse IssueCard(ctx).CardIssuanceRequest(cardIssuanceRequest).Execute()

Issue a Card



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
    cardIssuanceRequest := openapiclient.card_issuance_request{PhysicalCardIssuanceRequest: openapiclient.NewPhysicalCardIssuanceRequest("Form_example", "4903b265-ca1f-404b-b2be-9d8a783bb580", "f88e0735-5190-42c8-9baf-5bcaffa8e486", "14ffe289-490f-40be-9881-60c84d6e36bc", "Type_example")} // CardIssuanceRequest | Card to issue

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CardsApi.IssueCard(context.Background()).CardIssuanceRequest(cardIssuanceRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CardsApi.IssueCard``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IssueCard`: CardResponse
    fmt.Fprintf(os.Stdout, "Response from `CardsApi.IssueCard`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiIssueCardRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cardIssuanceRequest** | [**CardIssuanceRequest**](CardIssuanceRequest.md) | Card to issue | 

### Return type

[**CardResponse**](CardResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListCards

> CardListResponse ListCards(ctx).Tenant(tenant).CustomerId(customerId).AccountId(accountId).EmbossName(embossName).LastFour(lastFour).ExpirationDate(expirationDate).CardType(cardType).CardBrand(cardBrand).Form(form).CardProductId(cardProductId).CardStatus(cardStatus).PostalCode(postalCode).Limit(limit).PageToken(pageToken).Execute()

List Cards



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
    tenant := "tenant_example" // string |  (optional)
    customerId := TODO // string |  (optional)
    accountId := TODO // string |  (optional)
    embossName := "embossName_example" // string | emboss name (optional)
    lastFour := "1234" // string | The last 4 digits of the card PAN (optional)
    expirationDate := time.Now() // string | The date representing when the card would expire at (optional)
    cardType := "cardType_example" // string | Indicates the type of card (optional)
    cardBrand := openapiclient.card_brand("MASTERCARD") // CardBrand | The brand of a card product (optional)
    form := openapiclient.form("PHYSICAL") // Form | The format of the card (optional)
    cardProductId := TODO // string |  (optional)
    cardStatus := openapiclient.card_status("ACTIVE") // CardStatus | The status of a card (optional)
    postalCode := "49633" // string | The postal code of a card user (optional)
    limit := int32(100) // int32 |  (optional) (default to 100)
    pageToken := "0w75x33ztx" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CardsApi.ListCards(context.Background()).Tenant(tenant).CustomerId(customerId).AccountId(accountId).EmbossName(embossName).LastFour(lastFour).ExpirationDate(expirationDate).CardType(cardType).CardBrand(cardBrand).Form(form).CardProductId(cardProductId).CardStatus(cardStatus).PostalCode(postalCode).Limit(limit).PageToken(pageToken).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CardsApi.ListCards``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListCards`: CardListResponse
    fmt.Fprintf(os.Stdout, "Response from `CardsApi.ListCards`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListCardsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenant** | **string** |  | 
 **customerId** | [**string**](string.md) |  | 
 **accountId** | [**string**](string.md) |  | 
 **embossName** | **string** | emboss name | 
 **lastFour** | **string** | The last 4 digits of the card PAN | 
 **expirationDate** | **string** | The date representing when the card would expire at | 
 **cardType** | **string** | Indicates the type of card | 
 **cardBrand** | [**CardBrand**](CardBrand.md) | The brand of a card product | 
 **form** | [**Form**](Form.md) | The format of the card | 
 **cardProductId** | [**string**](string.md) |  | 
 **cardStatus** | [**CardStatus**](CardStatus.md) | The status of a card | 
 **postalCode** | **string** | The postal code of a card user | 
 **limit** | **int32** |  | [default to 100]
 **pageToken** | **string** |  | 

### Return type

[**CardListResponse**](CardListResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListChanges

> CardChangesList ListChanges(ctx, cardId).Execute()

List Card Changes



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
    resp, r, err := api_client.CardsApi.ListChanges(context.Background(), cardId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CardsApi.ListChanges``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListChanges`: CardChangesList
    fmt.Fprintf(os.Stdout, "Response from `CardsApi.ListChanges`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cardId** | [**string**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListChangesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CardChangesList**](CardChangesList.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateCard

> CardResponse UpdateCard(ctx, cardId).CardEditRequest(cardEditRequest).Execute()

Update Card



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
    cardEditRequest := *openapiclient.NewCardEditRequest(openapiclient.card_status("ACTIVE"), openapiclient.card_status_reason_code("NEW")) // CardEditRequest | Card edits

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CardsApi.UpdateCard(context.Background(), cardId).CardEditRequest(cardEditRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CardsApi.UpdateCard``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateCard`: CardResponse
    fmt.Fprintf(os.Stdout, "Response from `CardsApi.UpdateCard`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cardId** | [**string**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateCardRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **cardEditRequest** | [**CardEditRequest**](CardEditRequest.md) | Card edits | 

### Return type

[**CardResponse**](CardResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

