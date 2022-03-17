# \CardsApi

All URIs are relative to *https://api.synctera.com/v0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ActivateCard**](CardsApi.md#ActivateCard) | **Post** /cards/activate | Activate a card
[**CreateCardImage**](CardsApi.md#CreateCardImage) | **Post** /cards/images | Create Card Image
[**GetCard**](CardsApi.md#GetCard) | **Get** /cards/{card_id} | Get Card
[**GetCardBarcode**](CardsApi.md#GetCardBarcode) | **Get** /cards/{card_id}/barcodes | Get Card Barcode
[**GetCardImageData**](CardsApi.md#GetCardImageData) | **Get** /cards/images/{card_image_id}/data | Get Card Image Data
[**GetCardImageDetails**](CardsApi.md#GetCardImageDetails) | **Get** /cards/images/{card_image_id} | Get Card Image Details
[**GetCardWidgetURL**](CardsApi.md#GetCardWidgetURL) | **Get** /cards/card_widget_url | Get card widget URL
[**GetClientAccessToken**](CardsApi.md#GetClientAccessToken) | **Post** /cards/{card_id}/client_token | Get a client token
[**GetClientSingleUseToken**](CardsApi.md#GetClientSingleUseToken) | **Post** /cards/single_use_token | Get single-use token
[**IssueCard**](CardsApi.md#IssueCard) | **Post** /cards | Issue a Card
[**ListCardImageDetails**](CardsApi.md#ListCardImageDetails) | **Get** /cards/images | List Card Image Details
[**ListCards**](CardsApi.md#ListCards) | **Get** /cards | List Cards
[**ListChanges**](CardsApi.md#ListChanges) | **Get** /cards/{card_id}/changes | List Card Changes
[**UpdateAccountRange**](CardsApi.md#UpdateAccountRange) | **Patch** /cards/account_ranges/{account_range_id} | Update Account Range
[**UpdateBin**](CardsApi.md#UpdateBin) | **Patch** /cards/bins/{bin_id} | Update BIN
[**UpdateCard**](CardsApi.md#UpdateCard) | **Patch** /cards/{card_id} | Update Card
[**UpdateCardImageDetails**](CardsApi.md#UpdateCardImageDetails) | **Patch** /cards/images/{card_image_id} | Update Card Image Details
[**UploadCardImageData**](CardsApi.md#UploadCardImageData) | **Post** /cards/images/{card_image_id}/data | Upload Card Image



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
    cardActivationRequest := *openapiclient.NewCardActivationRequest("ActivationCode_example", "b30d9a2a-d4b6-4007-b4af-62926957be91") // CardActivationRequest | Card activation code

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


## CreateCardImage

> CardImageDetails CreateCardImage(ctx).CreateCardImageRequest(createCardImageRequest).Execute()

Create Card Image



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
    createCardImageRequest := *openapiclient.NewCreateCardImageRequest("b6f6b35a-56f3-4960-b4cb-f66c565150ba") // CreateCardImageRequest | Details of the image to create

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CardsApi.CreateCardImage(context.Background()).CreateCardImageRequest(createCardImageRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CardsApi.CreateCardImage``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateCardImage`: CardImageDetails
    fmt.Fprintf(os.Stdout, "Response from `CardsApi.CreateCardImage`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateCardImageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createCardImageRequest** | [**CreateCardImageRequest**](CreateCardImageRequest.md) | Details of the image to create | 

### Return type

[**CardImageDetails**](CardImageDetails.md)

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


## GetCardBarcode

> InlineResponse200 GetCardBarcode(ctx, cardId).Execute()

Get Card Barcode



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
    resp, r, err := api_client.CardsApi.GetCardBarcode(context.Background(), cardId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CardsApi.GetCardBarcode``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCardBarcode`: InlineResponse200
    fmt.Fprintf(os.Stdout, "Response from `CardsApi.GetCardBarcode`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cardId** | [**string**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCardBarcodeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**InlineResponse200**](InlineResponse200.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCardImageData

> *os.File GetCardImageData(ctx, cardImageId).Execute()

Get Card Image Data



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
    cardImageId := TODO // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CardsApi.GetCardImageData(context.Background(), cardImageId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CardsApi.GetCardImageData``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCardImageData`: *os.File
    fmt.Fprintf(os.Stdout, "Response from `CardsApi.GetCardImageData`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cardImageId** | [**string**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCardImageDataRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[***os.File**](*os.File.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: image/jpeg, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCardImageDetails

> CardImageDetails GetCardImageDetails(ctx, cardImageId).Execute()

Get Card Image Details



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
    cardImageId := TODO // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CardsApi.GetCardImageDetails(context.Background(), cardImageId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CardsApi.GetCardImageDetails``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCardImageDetails`: CardImageDetails
    fmt.Fprintf(os.Stdout, "Response from `CardsApi.GetCardImageDetails`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cardImageId** | [**string**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCardImageDetailsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CardImageDetails**](CardImageDetails.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCardWidgetURL

> CardWidgetUrlResponse GetCardWidgetURL(ctx).WidgetType(widgetType).CustomerId(customerId).AccountId(accountId).CardId(cardId).Execute()

Get card widget URL



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
    widgetType := openapiclient.widget_type("set_pin") // WidgetType | The type of widget for which to construct the URL
    customerId := TODO // string | 
    accountId := TODO // string | 
    cardId := TODO // string | The ID of the card (required for set PIN widget) (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CardsApi.GetCardWidgetURL(context.Background()).WidgetType(widgetType).CustomerId(customerId).AccountId(accountId).CardId(cardId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CardsApi.GetCardWidgetURL``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCardWidgetURL`: CardWidgetUrlResponse
    fmt.Fprintf(os.Stdout, "Response from `CardsApi.GetCardWidgetURL`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetCardWidgetURLRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **widgetType** | [**WidgetType**](WidgetType.md) | The type of widget for which to construct the URL | 
 **customerId** | [**string**](string.md) |  | 
 **accountId** | [**string**](string.md) |  | 
 **cardId** | [**string**](string.md) | The ID of the card (required for set PIN widget) | 

### Return type

[**CardWidgetUrlResponse**](CardWidgetUrlResponse.md)

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
    singleUseTokenRequest := *openapiclient.NewSingleUseTokenRequest("7c1ffd48-b4d0-4a61-bbac-b0c240fc7d33", "50f496ee-8a19-4b96-9e8d-04dcc5af5e48") // SingleUseTokenRequest | User token details

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
    cardIssuanceRequest := openapiclient.card_issuance_request{PhysicalCardIssuanceRequest: openapiclient.NewPhysicalCardIssuanceRequest("Form_example", "b7912244-dedc-4150-86ee-e1b970c3f564", "cdd7cdc7-e9e3-457d-aa66-368dc28f9434", "29e0655f-b1fe-4211-ba2f-6a05acb16faa", "Type_example")} // CardIssuanceRequest | Card to issue

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


## ListCardImageDetails

> CardImageDetailsList ListCardImageDetails(ctx).CustomerId(customerId).Execute()

List Card Image Details



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
    customerId := TODO // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CardsApi.ListCardImageDetails(context.Background()).CustomerId(customerId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CardsApi.ListCardImageDetails``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListCardImageDetails`: CardImageDetailsList
    fmt.Fprintf(os.Stdout, "Response from `CardsApi.ListCardImageDetails`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListCardImageDetailsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **customerId** | [**string**](string.md) |  | 

### Return type

[**CardImageDetailsList**](CardImageDetailsList.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListCards

> CardListResponse ListCards(ctx).Tenant(tenant).CustomerId(customerId).AccountId(accountId).EmbossName(embossName).LastFour(lastFour).ExpirationDate(expirationDate).CardType(cardType).CardBrand(cardBrand).Form(form).CardProductId(cardProductId).CardStatus(cardStatus).PostalCode(postalCode).Limit(limit).PageToken(pageToken).SortBy(sortBy).Execute()

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
    expirationDate := time.Now() // oapi.Date | The date representing when the card would expire at (optional)
    cardType := "cardType_example" // string | Indicates the type of card (optional)
    cardBrand := openapiclient.card_brand("MASTERCARD") // CardBrand | The brand of a card product (optional)
    form := openapiclient.form("PHYSICAL") // Form | The format of the card (optional)
    cardProductId := TODO // string |  (optional)
    cardStatus := openapiclient.card_status("ACTIVE") // CardStatus | The status of a card (optional)
    postalCode := "49633" // string | The postal code of a card user (optional)
    limit := int32(100) // int32 |  (optional) (default to 100)
    pageToken := "2ss64pmfs8" // string |  (optional)
    sortBy := []string{"SortBy_example"} // []string | Specifies the sort order for the returned cards.  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CardsApi.ListCards(context.Background()).Tenant(tenant).CustomerId(customerId).AccountId(accountId).EmbossName(embossName).LastFour(lastFour).ExpirationDate(expirationDate).CardType(cardType).CardBrand(cardBrand).Form(form).CardProductId(cardProductId).CardStatus(cardStatus).PostalCode(postalCode).Limit(limit).PageToken(pageToken).SortBy(sortBy).Execute()
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
 **expirationDate** | **oapi.Date** | The date representing when the card would expire at | 
 **cardType** | **string** | Indicates the type of card | 
 **cardBrand** | [**CardBrand**](CardBrand.md) | The brand of a card product | 
 **form** | [**Form**](Form.md) | The format of the card | 
 **cardProductId** | [**string**](string.md) |  | 
 **cardStatus** | [**CardStatus**](CardStatus.md) | The status of a card | 
 **postalCode** | **string** | The postal code of a card user | 
 **limit** | **int32** |  | [default to 100]
 **pageToken** | **string** |  | 
 **sortBy** | **[]string** | Specifies the sort order for the returned cards.  | 

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


## UpdateAccountRange

> AccountRangeResponse UpdateAccountRange(ctx, accountRangeId).AccountRangeUpdateRequest(accountRangeUpdateRequest).Execute()

Update Account Range



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
    accountRangeId := TODO // string | 
    accountRangeUpdateRequest := *openapiclient.NewAccountRangeUpdateRequest() // AccountRangeUpdateRequest | Fields to update

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CardsApi.UpdateAccountRange(context.Background(), accountRangeId).AccountRangeUpdateRequest(accountRangeUpdateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CardsApi.UpdateAccountRange``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateAccountRange`: AccountRangeResponse
    fmt.Fprintf(os.Stdout, "Response from `CardsApi.UpdateAccountRange`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountRangeId** | [**string**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateAccountRangeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **accountRangeUpdateRequest** | [**AccountRangeUpdateRequest**](AccountRangeUpdateRequest.md) | Fields to update | 

### Return type

[**AccountRangeResponse**](AccountRangeResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateBin

> BinResponse UpdateBin(ctx, binId).BinUpdateRequest(binUpdateRequest).Execute()

Update BIN



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
    binId := TODO // string | 
    binUpdateRequest := *openapiclient.NewBinUpdateRequest() // BinUpdateRequest | Fields to update

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CardsApi.UpdateBin(context.Background(), binId).BinUpdateRequest(binUpdateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CardsApi.UpdateBin``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateBin`: BinResponse
    fmt.Fprintf(os.Stdout, "Response from `CardsApi.UpdateBin`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**binId** | [**string**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateBinRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **binUpdateRequest** | [**BinUpdateRequest**](BinUpdateRequest.md) | Fields to update | 

### Return type

[**BinResponse**](BinResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
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
    cardEditRequest := *openapiclient.NewCardEditRequest(openapiclient.card_status_request("ACTIVE"), openapiclient.card_status_reason_code("NEW")) // CardEditRequest | Card edits

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


## UpdateCardImageDetails

> CardImageDetails UpdateCardImageDetails(ctx, cardImageId).UpdateCardImageRequest(updateCardImageRequest).Execute()

Update Card Image Details



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
    cardImageId := TODO // string | 
    updateCardImageRequest := *openapiclient.NewUpdateCardImageRequest(openapiclient.card_image_status("NOT_UPLOADED")) // UpdateCardImageRequest | Details of the image to create

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CardsApi.UpdateCardImageDetails(context.Background(), cardImageId).UpdateCardImageRequest(updateCardImageRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CardsApi.UpdateCardImageDetails``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateCardImageDetails`: CardImageDetails
    fmt.Fprintf(os.Stdout, "Response from `CardsApi.UpdateCardImageDetails`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cardImageId** | [**string**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateCardImageDetailsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateCardImageRequest** | [**UpdateCardImageRequest**](UpdateCardImageRequest.md) | Details of the image to create | 

### Return type

[**CardImageDetails**](CardImageDetails.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UploadCardImageData

> CardImageDetails UploadCardImageData(ctx, cardImageId).Body(body).Execute()

Upload Card Image



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
    cardImageId := TODO // string | 
    body := os.NewFile(1234, "some_file") // *os.File | Binary image data

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CardsApi.UploadCardImageData(context.Background(), cardImageId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CardsApi.UploadCardImageData``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UploadCardImageData`: CardImageDetails
    fmt.Fprintf(os.Stdout, "Response from `CardsApi.UploadCardImageData`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cardImageId** | [**string**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUploadCardImageDataRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | ***os.File** | Binary image data | 

### Return type

[**CardImageDetails**](CardImageDetails.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: image/jpeg
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

