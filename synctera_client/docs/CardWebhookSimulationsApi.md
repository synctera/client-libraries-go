# \CardWebhookSimulationsApi

All URIs are relative to *https://api.synctera.com/v0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SimulateCardFulfillmentEvent**](CardWebhookSimulationsApi.md#SimulateCardFulfillmentEvent) | **Post** /cards/{card_id}/webhook_simulations/fulfillment | Simulate Card Fulfillment Event



## SimulateCardFulfillmentEvent

> SimulateCardFulfillment SimulateCardFulfillmentEvent(ctx, cardId).SimulateCardFulfillment(simulateCardFulfillment).Execute()

Simulate Card Fulfillment Event



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
    simulateCardFulfillment := *openapiclient.NewSimulateCardFulfillment(openapiclient.card_fulfillment_status("ISSUED")) // SimulateCardFulfillment | Desired simulated fulfillment status change value

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CardWebhookSimulationsApi.SimulateCardFulfillmentEvent(context.Background(), cardId).SimulateCardFulfillment(simulateCardFulfillment).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CardWebhookSimulationsApi.SimulateCardFulfillmentEvent``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SimulateCardFulfillmentEvent`: SimulateCardFulfillment
    fmt.Fprintf(os.Stdout, "Response from `CardWebhookSimulationsApi.SimulateCardFulfillmentEvent`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cardId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSimulateCardFulfillmentEventRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **simulateCardFulfillment** | [**SimulateCardFulfillment**](SimulateCardFulfillment.md) | Desired simulated fulfillment status change value | 

### Return type

[**SimulateCardFulfillment**](SimulateCardFulfillment.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

