# \WatchlistApi

All URIs are relative to *https://api.synctera.com/v0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetWatchlistAlert**](WatchlistApi.md#GetWatchlistAlert) | **Get** /customers/{customer_id}/watchlists/alerts/{alert_id} | Retrieve watchlist monitoring alert
[**GetWatchlistSubscription**](WatchlistApi.md#GetWatchlistSubscription) | **Get** /customers/{customer_id}/watchlists/subscriptions/{subscription_id} | Retrieve watchlist monitoring subscription
[**ListWatchlistAlerts**](WatchlistApi.md#ListWatchlistAlerts) | **Get** /customers/{customer_id}/watchlists/alerts | List watchlist monitoring alerts for a customer
[**ListWatchlistSubscriptions**](WatchlistApi.md#ListWatchlistSubscriptions) | **Get** /customers/{customer_id}/watchlists/subscriptions | List watchlist monitoring subscriptions for a customer
[**SuppressWatchlistEntityAlert**](WatchlistApi.md#SuppressWatchlistEntityAlert) | **Post** /customers/{customer_id}/watchlists/suppressions | Suppress entity alert
[**UpdateWatchlistAlert**](WatchlistApi.md#UpdateWatchlistAlert) | **Put** /customers/{customer_id}/watchlists/alerts/{alert_id} | Update watchlist alert
[**UpdateWatchlistSubscription**](WatchlistApi.md#UpdateWatchlistSubscription) | **Put** /customers/{customer_id}/watchlists/subscriptions/{subscription_id} | Update watchlist monitoring subscription
[**WatchlistSubscribe**](WatchlistApi.md#WatchlistSubscribe) | **Post** /customers/{customer_id}/watchlists/subscriptions | Subscribe a customer to watchlist monitoring



## GetWatchlistAlert

> WatchlistAlert GetWatchlistAlert(ctx, customerId, alertId).Execute()

Retrieve watchlist monitoring alert

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
    customerId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The customer's unique identifier
    alertId := "ff23b9d0-4e64-4b98-9f4a-3591ed08121a" // string | Unique identifier for this watchlist alert.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WatchlistApi.GetWatchlistAlert(context.Background(), customerId, alertId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WatchlistApi.GetWatchlistAlert``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetWatchlistAlert`: WatchlistAlert
    fmt.Fprintf(os.Stdout, "Response from `WatchlistApi.GetWatchlistAlert`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customerId** | **string** | The customer&#39;s unique identifier | 
**alertId** | **string** | Unique identifier for this watchlist alert. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetWatchlistAlertRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**WatchlistAlert**](WatchlistAlert.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetWatchlistSubscription

> WatchlistSubscription GetWatchlistSubscription(ctx, customerId, subscriptionId).Execute()

Retrieve watchlist monitoring subscription

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
    customerId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The customer's unique identifier
    subscriptionId := "ddcacaa4-e0e4-4652-ae9f-5ef7f1b7d7e1" // string | Watchlist monitoring subscription ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WatchlistApi.GetWatchlistSubscription(context.Background(), customerId, subscriptionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WatchlistApi.GetWatchlistSubscription``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetWatchlistSubscription`: WatchlistSubscription
    fmt.Fprintf(os.Stdout, "Response from `WatchlistApi.GetWatchlistSubscription`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customerId** | **string** | The customer&#39;s unique identifier | 
**subscriptionId** | **string** | Watchlist monitoring subscription ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetWatchlistSubscriptionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**WatchlistSubscription**](WatchlistSubscription.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListWatchlistAlerts

> WatchlistAlertList ListWatchlistAlerts(ctx, customerId).Execute()

List watchlist monitoring alerts for a customer

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
    customerId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The customer's unique identifier

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WatchlistApi.ListWatchlistAlerts(context.Background(), customerId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WatchlistApi.ListWatchlistAlerts``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListWatchlistAlerts`: WatchlistAlertList
    fmt.Fprintf(os.Stdout, "Response from `WatchlistApi.ListWatchlistAlerts`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customerId** | **string** | The customer&#39;s unique identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiListWatchlistAlertsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**WatchlistAlertList**](WatchlistAlertList.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListWatchlistSubscriptions

> WatchlistSubscriptionList ListWatchlistSubscriptions(ctx, customerId).Execute()

List watchlist monitoring subscriptions for a customer

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
    customerId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The customer's unique identifier

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WatchlistApi.ListWatchlistSubscriptions(context.Background(), customerId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WatchlistApi.ListWatchlistSubscriptions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListWatchlistSubscriptions`: WatchlistSubscriptionList
    fmt.Fprintf(os.Stdout, "Response from `WatchlistApi.ListWatchlistSubscriptions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customerId** | **string** | The customer&#39;s unique identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiListWatchlistSubscriptionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**WatchlistSubscriptionList**](WatchlistSubscriptionList.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SuppressWatchlistEntityAlert

> SuppressWatchlistEntityAlert(ctx, customerId).WatchlistSuppress(watchlistSuppress).IdempotencyKey(idempotencyKey).Execute()

Suppress entity alert

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
    customerId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The customer's unique identifier
    watchlistSuppress := *openapiclient.NewWatchlistSuppress("ProviderSubjectId_example", "ProviderSubscriptionId_example", "Status_example") // WatchlistSuppress | A watchlist suppression object
    idempotencyKey := "df122e6f-2ba8-48a5-9508-4350bba5f27e" // string | An idempotency key is an arbitrary unique value generated by client to detect subsequent retries of the same request. It is recommended that a UUID or a similar random identifier be used as an idempotency key (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WatchlistApi.SuppressWatchlistEntityAlert(context.Background(), customerId).WatchlistSuppress(watchlistSuppress).IdempotencyKey(idempotencyKey).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WatchlistApi.SuppressWatchlistEntityAlert``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customerId** | **string** | The customer&#39;s unique identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiSuppressWatchlistEntityAlertRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **watchlistSuppress** | [**WatchlistSuppress**](WatchlistSuppress.md) | A watchlist suppression object | 
 **idempotencyKey** | **string** | An idempotency key is an arbitrary unique value generated by client to detect subsequent retries of the same request. It is recommended that a UUID or a similar random identifier be used as an idempotency key | 

### Return type

 (empty response body)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateWatchlistAlert

> UpdateWatchlistAlert(ctx, customerId, alertId).WatchlistAlert(watchlistAlert).Execute()

Update watchlist alert

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
    customerId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The customer's unique identifier
    alertId := "ff23b9d0-4e64-4b98-9f4a-3591ed08121a" // string | Unique identifier for this watchlist alert.
    watchlistAlert := *openapiclient.NewWatchlistAlert("Status_example") // WatchlistAlert | A watchlist body

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WatchlistApi.UpdateWatchlistAlert(context.Background(), customerId, alertId).WatchlistAlert(watchlistAlert).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WatchlistApi.UpdateWatchlistAlert``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customerId** | **string** | The customer&#39;s unique identifier | 
**alertId** | **string** | Unique identifier for this watchlist alert. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateWatchlistAlertRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **watchlistAlert** | [**WatchlistAlert**](WatchlistAlert.md) | A watchlist body | 

### Return type

 (empty response body)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateWatchlistSubscription

> WatchlistSubscription UpdateWatchlistSubscription(ctx, customerId, subscriptionId).WatchlistSubscription(watchlistSubscription).Execute()

Update watchlist monitoring subscription

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
    customerId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The customer's unique identifier
    subscriptionId := "ddcacaa4-e0e4-4652-ae9f-5ef7f1b7d7e1" // string | Watchlist monitoring subscription ID
    watchlistSubscription := *openapiclient.NewWatchlistSubscription(false) // WatchlistSubscription | Watchlist monitoring subscription to be updated. The only field that matters is `status`; all other fields are ignored. 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WatchlistApi.UpdateWatchlistSubscription(context.Background(), customerId, subscriptionId).WatchlistSubscription(watchlistSubscription).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WatchlistApi.UpdateWatchlistSubscription``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateWatchlistSubscription`: WatchlistSubscription
    fmt.Fprintf(os.Stdout, "Response from `WatchlistApi.UpdateWatchlistSubscription`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customerId** | **string** | The customer&#39;s unique identifier | 
**subscriptionId** | **string** | Watchlist monitoring subscription ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateWatchlistSubscriptionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **watchlistSubscription** | [**WatchlistSubscription**](WatchlistSubscription.md) | Watchlist monitoring subscription to be updated. The only field that matters is &#x60;status&#x60;; all other fields are ignored.  | 

### Return type

[**WatchlistSubscription**](WatchlistSubscription.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WatchlistSubscribe

> WatchlistSubscription WatchlistSubscribe(ctx, customerId).WatchlistSubscription(watchlistSubscription).IdempotencyKey(idempotencyKey).Execute()

Subscribe a customer to watchlist monitoring

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
    customerId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The customer's unique identifier
    watchlistSubscription := *openapiclient.NewWatchlistSubscription(false) // WatchlistSubscription | A watchlist subscription
    idempotencyKey := "df122e6f-2ba8-48a5-9508-4350bba5f27e" // string | An idempotency key is an arbitrary unique value generated by client to detect subsequent retries of the same request. It is recommended that a UUID or a similar random identifier be used as an idempotency key (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WatchlistApi.WatchlistSubscribe(context.Background(), customerId).WatchlistSubscription(watchlistSubscription).IdempotencyKey(idempotencyKey).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WatchlistApi.WatchlistSubscribe``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `WatchlistSubscribe`: WatchlistSubscription
    fmt.Fprintf(os.Stdout, "Response from `WatchlistApi.WatchlistSubscribe`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customerId** | **string** | The customer&#39;s unique identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiWatchlistSubscribeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **watchlistSubscription** | [**WatchlistSubscription**](WatchlistSubscription.md) | A watchlist subscription | 
 **idempotencyKey** | **string** | An idempotency key is an arbitrary unique value generated by client to detect subsequent retries of the same request. It is recommended that a UUID or a similar random identifier be used as an idempotency key | 

### Return type

[**WatchlistSubscription**](WatchlistSubscription.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

