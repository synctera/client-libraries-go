# \CardsApi

All URIs are relative to *https://api.synctera.com/v0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ActivateCard**](CardsApi.md#ActivateCard) | **Post** /cards/activate | Activate a card
[**CreateCardImage**](CardsApi.md#CreateCardImage) | **Post** /cards/images | Create Card Image
[**CreateGateway**](CardsApi.md#CreateGateway) | **Post** /cards/gateways | Create Gateway
[**GetCard**](CardsApi.md#GetCard) | **Get** /cards/{card_id} | Get Card
[**GetCardBarcode**](CardsApi.md#GetCardBarcode) | **Get** /cards/{card_id}/barcodes | Get Card Barcode
[**GetCardImageData**](CardsApi.md#GetCardImageData) | **Get** /cards/images/{card_image_id}/data | Get Card Image Data
[**GetCardImageDetails**](CardsApi.md#GetCardImageDetails) | **Get** /cards/images/{card_image_id} | Get Card Image Details
[**GetCardWidgetURL**](CardsApi.md#GetCardWidgetURL) | **Get** /cards/card_widget_url | Get card widget URL
[**GetClientAccessToken**](CardsApi.md#GetClientAccessToken) | **Post** /cards/{card_id}/client_token | Get a client token
[**GetClientSingleUseToken**](CardsApi.md#GetClientSingleUseToken) | **Post** /cards/single_use_token | Get single-use token
[**GetGateway**](CardsApi.md#GetGateway) | **Get** /cards/gateways/{gateway_id} | Get Gateway
[**IssueCard**](CardsApi.md#IssueCard) | **Post** /cards | Issue a Card
[**ListCardImageDetails**](CardsApi.md#ListCardImageDetails) | **Get** /cards/images | List Card Image Details
[**ListCardProducts**](CardsApi.md#ListCardProducts) | **Get** /cards/products | List Card Products
[**ListCards**](CardsApi.md#ListCards) | **Get** /cards | List Cards
[**ListChanges**](CardsApi.md#ListChanges) | **Get** /cards/{card_id}/changes | List Card Changes
[**ListGateways**](CardsApi.md#ListGateways) | **Get** /cards/gateways | List Gateways
[**UpdateCard**](CardsApi.md#UpdateCard) | **Patch** /cards/{card_id} | Update Card
[**UpdateCardImageDetails**](CardsApi.md#UpdateCardImageDetails) | **Patch** /cards/images/{card_image_id} | Update Card Image Details
[**UpdateGateway**](CardsApi.md#UpdateGateway) | **Patch** /cards/gateways/{gateway_id} | Update Gateway
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
    cardActivationRequest := *openapiclient.NewCardActivationRequest("ActivationCode_example", "1acf447c-544b-41b4-94f9-bd96b03cc48e") // CardActivationRequest | Card activation code

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CardsApi.ActivateCard(context.Background()).CardActivationRequest(cardActivationRequest).Execute()
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
    createCardImageRequest := *openapiclient.NewCreateCardImageRequest("d6180286-ae04-4bdb-ac62-bce0b4b06f91", "8da2b02b-81f2-41f3-a32e-041eb8ccd825") // CreateCardImageRequest | Details of the image to create

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CardsApi.CreateCardImage(context.Background()).CreateCardImageRequest(createCardImageRequest).Execute()
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


## CreateGateway

> GatewayResponse CreateGateway(ctx).CreateGatewayRequest(createGatewayRequest).Execute()

Create Gateway



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
    createGatewayRequest := *openapiclient.NewCreateGatewayRequest([]string{"9faacdc6-bd75-4c79-b39f-898a95a21cac"}, "Url_example") // CreateGatewayRequest | Create a new Authorization Gateway Configuration

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CardsApi.CreateGateway(context.Background()).CreateGatewayRequest(createGatewayRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CardsApi.CreateGateway``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateGateway`: GatewayResponse
    fmt.Fprintf(os.Stdout, "Response from `CardsApi.CreateGateway`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateGatewayRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createGatewayRequest** | [**CreateGatewayRequest**](CreateGatewayRequest.md) | Create a new Authorization Gateway Configuration | 

### Return type

[**GatewayResponse**](GatewayResponse.md)

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
    cardId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CardsApi.GetCard(context.Background(), cardId).Execute()
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
**cardId** | **string** |  | 

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

> GetCardBarcode200Response GetCardBarcode(ctx, cardId).Execute()

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
    cardId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CardsApi.GetCardBarcode(context.Background(), cardId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CardsApi.GetCardBarcode``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCardBarcode`: GetCardBarcode200Response
    fmt.Fprintf(os.Stdout, "Response from `CardsApi.GetCardBarcode`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cardId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCardBarcodeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetCardBarcode200Response**](GetCardBarcode200Response.md)

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
    cardImageId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CardsApi.GetCardImageData(context.Background(), cardImageId).Execute()
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
**cardImageId** | **string** |  | 

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
    cardImageId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CardsApi.GetCardImageDetails(context.Background(), cardImageId).Execute()
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
**cardImageId** | **string** |  | 

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
    customerId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
    accountId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
    cardId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The ID of the card (required for set PIN widget) (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CardsApi.GetCardWidgetURL(context.Background()).WidgetType(widgetType).CustomerId(customerId).AccountId(accountId).CardId(cardId).Execute()
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
 **customerId** | **string** |  | 
 **accountId** | **string** |  | 
 **cardId** | **string** | The ID of the card (required for set PIN widget) | 

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
    cardId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CardsApi.GetClientAccessToken(context.Background(), cardId).Execute()
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
**cardId** | **string** |  | 

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
    singleUseTokenRequest := *openapiclient.NewSingleUseTokenRequest("df063c1c-cc31-4d60-8a54-ed7a6fdbc1be", "ffe4f5b1-c84d-4fbf-b595-821d3254fed8") // SingleUseTokenRequest | User token details

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CardsApi.GetClientSingleUseToken(context.Background()).SingleUseTokenRequest(singleUseTokenRequest).Execute()
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


## GetGateway

> GatewayResponse GetGateway(ctx, gatewayId).Execute()

Get Gateway



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
    gatewayId := "gatewayId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CardsApi.GetGateway(context.Background(), gatewayId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CardsApi.GetGateway``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetGateway`: GatewayResponse
    fmt.Fprintf(os.Stdout, "Response from `CardsApi.GetGateway`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**gatewayId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetGatewayRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GatewayResponse**](GatewayResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
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
    cardIssuanceRequest := openapiclient.card_issuance_request{PhysicalCardIssuanceRequest: openapiclient.NewPhysicalCardIssuanceRequest("Form_example", "98d3a7dd-8b7e-4b93-94b0-f5c9bcf88b39", "f453e97a-83f8-4810-bdbc-8df324fc4241", "a33e874c-4fa5-4c6f-850f-0778ffa097e9", "Type_example")} // CardIssuanceRequest | Card to issue

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CardsApi.IssueCard(context.Background()).CardIssuanceRequest(cardIssuanceRequest).Execute()
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
    customerId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CardsApi.ListCardImageDetails(context.Background()).CustomerId(customerId).Execute()
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
 **customerId** | **string** |  | 

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


## ListCardProducts

> CardProductListResponse ListCardProducts(ctx).Limit(limit).PageToken(pageToken).Execute()

List Card Products

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
    pageToken := "h50ffqz9q5" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CardsApi.ListCardProducts(context.Background()).Limit(limit).PageToken(pageToken).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CardsApi.ListCardProducts``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListCardProducts`: CardProductListResponse
    fmt.Fprintf(os.Stdout, "Response from `CardsApi.ListCardProducts`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListCardProductsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** |  | [default to 100]
 **pageToken** | **string** |  | 

### Return type

[**CardProductListResponse**](CardProductListResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListCards

> CardListResponse ListCards(ctx).CustomerId(customerId).AccountId(accountId).EmbossName(embossName).LastFour(lastFour).ExpirationDate(expirationDate).CardType(cardType).CardBrand(cardBrand).Form(form).CardProductId(cardProductId).CardStatus(cardStatus).PostalCode(postalCode).Limit(limit).PageToken(pageToken).SortBy(sortBy).Execute()

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
    customerId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string |  (optional)
    accountId := []string{"4403f645-c2e3-4f8d-9027-127dca09f353"} // []string | Account ID(s). Multiple IDs can be provided as a comma-separated list. (optional)
    embossName := "embossName_example" // string | emboss name (optional)
    lastFour := "1234" // string | The last 4 digits of the card PAN (optional)
    expirationDate := time.Now() // string | The date representing when the card would expire at (optional)
    cardType := "cardType_example" // string | Indicates the type of card (optional)
    cardBrand := openapiclient.card_brand("MASTERCARD") // CardBrand | The brand of a card product (optional)
    form := openapiclient.form("PHYSICAL") // Form | The format of the card (optional)
    cardProductId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string |  (optional)
    cardStatus := openapiclient.card_status("ACTIVE") // CardStatus | The status of a card (optional)
    postalCode := "49633" // string | The postal code of a card user (optional)
    limit := int32(100) // int32 |  (optional) (default to 100)
    pageToken := "h50ffqz9q5" // string |  (optional)
    sortBy := []string{"SortBy_example"} // []string | Specifies the sort order for the returned cards.  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CardsApi.ListCards(context.Background()).CustomerId(customerId).AccountId(accountId).EmbossName(embossName).LastFour(lastFour).ExpirationDate(expirationDate).CardType(cardType).CardBrand(cardBrand).Form(form).CardProductId(cardProductId).CardStatus(cardStatus).PostalCode(postalCode).Limit(limit).PageToken(pageToken).SortBy(sortBy).Execute()
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
 **customerId** | **string** |  | 
 **accountId** | **[]string** | Account ID(s). Multiple IDs can be provided as a comma-separated list. | 
 **embossName** | **string** | emboss name | 
 **lastFour** | **string** | The last 4 digits of the card PAN | 
 **expirationDate** | **string** | The date representing when the card would expire at | 
 **cardType** | **string** | Indicates the type of card | 
 **cardBrand** | [**CardBrand**](CardBrand.md) | The brand of a card product | 
 **form** | [**Form**](Form.md) | The format of the card | 
 **cardProductId** | **string** |  | 
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
    cardId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CardsApi.ListChanges(context.Background(), cardId).Execute()
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
**cardId** | **string** |  | 

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


## ListGateways

> GatewayListResponse ListGateways(ctx).Limit(limit).PageToken(pageToken).Execute()

List Gateways



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
    pageToken := "h50ffqz9q5" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CardsApi.ListGateways(context.Background()).Limit(limit).PageToken(pageToken).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CardsApi.ListGateways``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListGateways`: GatewayListResponse
    fmt.Fprintf(os.Stdout, "Response from `CardsApi.ListGateways`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListGatewaysRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** |  | [default to 100]
 **pageToken** | **string** |  | 

### Return type

[**GatewayListResponse**](GatewayListResponse.md)

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
    cardId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
    cardEditRequest := *openapiclient.NewCardEditRequest() // CardEditRequest | Card edits

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CardsApi.UpdateCard(context.Background(), cardId).CardEditRequest(cardEditRequest).Execute()
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
**cardId** | **string** |  | 

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
    cardImageId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
    updateCardImageRequest := *openapiclient.NewUpdateCardImageRequest(openapiclient.card_image_status("NOT_UPLOADED")) // UpdateCardImageRequest | Details of the image to create

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CardsApi.UpdateCardImageDetails(context.Background(), cardImageId).UpdateCardImageRequest(updateCardImageRequest).Execute()
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
**cardImageId** | **string** |  | 

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


## UpdateGateway

> GatewayResponse UpdateGateway(ctx, gatewayId).UpdateGatewayRequest(updateGatewayRequest).Execute()

Update Gateway



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
    gatewayId := "gatewayId_example" // string | 
    updateGatewayRequest := *openapiclient.NewUpdateGatewayRequest() // UpdateGatewayRequest | Gateway edits

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CardsApi.UpdateGateway(context.Background(), gatewayId).UpdateGatewayRequest(updateGatewayRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CardsApi.UpdateGateway``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateGateway`: GatewayResponse
    fmt.Fprintf(os.Stdout, "Response from `CardsApi.UpdateGateway`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**gatewayId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateGatewayRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateGatewayRequest** | [**UpdateGatewayRequest**](UpdateGatewayRequest.md) | Gateway edits | 

### Return type

[**GatewayResponse**](GatewayResponse.md)

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
    cardImageId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
    body := os.NewFile(1234, "some_file") // *os.File | Binary image data

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CardsApi.UploadCardImageData(context.Background(), cardImageId).Body(body).Execute()
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
**cardImageId** | **string** |  | 

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

