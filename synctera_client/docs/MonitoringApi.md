# \MonitoringApi

All URIs are relative to *https://api.synctera.com/v0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateSubscription**](MonitoringApi.md#CreateSubscription) | **Post** /monitoring/subscriptions | Subscribe a customer or business to monitoring
[**DeleteSubscription**](MonitoringApi.md#DeleteSubscription) | **Delete** /monitoring/subscriptions/{subscription_id} | Delete monitoring subscription
[**GetAlert**](MonitoringApi.md#GetAlert) | **Get** /monitoring/alerts/{alert_id} | Retrieve a monitoring alert
[**GetSubscription**](MonitoringApi.md#GetSubscription) | **Get** /monitoring/subscriptions/{subscription_id} | Retrieve monitoring subscription
[**ListAlerts**](MonitoringApi.md#ListAlerts) | **Get** /monitoring/alerts | List monitoring alerts
[**ListSubscriptions**](MonitoringApi.md#ListSubscriptions) | **Get** /monitoring/subscriptions | List monitoring subscriptions
[**UpdateAlert**](MonitoringApi.md#UpdateAlert) | **Patch** /monitoring/alerts/{alert_id} | Update a monitoring alert



## CreateSubscription

> MonitoringSubscription CreateSubscription(ctx).MonitoringSubscription(monitoringSubscription).Execute()

Subscribe a customer or business to monitoring

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
    monitoringSubscription := *openapiclient.NewMonitoringSubscription() // MonitoringSubscription | The monitoring subscription.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MonitoringApi.CreateSubscription(context.Background()).MonitoringSubscription(monitoringSubscription).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MonitoringApi.CreateSubscription``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateSubscription`: MonitoringSubscription
    fmt.Fprintf(os.Stdout, "Response from `MonitoringApi.CreateSubscription`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateSubscriptionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **monitoringSubscription** | [**MonitoringSubscription**](MonitoringSubscription.md) | The monitoring subscription. | 

### Return type

[**MonitoringSubscription**](MonitoringSubscription.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteSubscription

> DeleteResponse DeleteSubscription(ctx, subscriptionId).Execute()

Delete monitoring subscription

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
    subscriptionId := "d3682b94-3a25-41ee-8358-2308cb4a71b1" // string | Unique identifier for monitoring subscription.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MonitoringApi.DeleteSubscription(context.Background(), subscriptionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MonitoringApi.DeleteSubscription``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteSubscription`: DeleteResponse
    fmt.Fprintf(os.Stdout, "Response from `MonitoringApi.DeleteSubscription`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subscriptionId** | **string** | Unique identifier for monitoring subscription. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSubscriptionRequest struct via the builder pattern


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


## GetAlert

> MonitoringAlert GetAlert(ctx, alertId).Execute()

Retrieve a monitoring alert

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
    alertId := "4a09a067-2f01-475d-bc62-50f245536017" // string | Unique identifier for this monitoring alert.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MonitoringApi.GetAlert(context.Background(), alertId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MonitoringApi.GetAlert``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAlert`: MonitoringAlert
    fmt.Fprintf(os.Stdout, "Response from `MonitoringApi.GetAlert`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**alertId** | **string** | Unique identifier for this monitoring alert. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAlertRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**MonitoringAlert**](MonitoringAlert.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSubscription

> MonitoringSubscription GetSubscription(ctx, subscriptionId).Execute()

Retrieve monitoring subscription

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
    subscriptionId := "d3682b94-3a25-41ee-8358-2308cb4a71b1" // string | Unique identifier for monitoring subscription.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MonitoringApi.GetSubscription(context.Background(), subscriptionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MonitoringApi.GetSubscription``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSubscription`: MonitoringSubscription
    fmt.Fprintf(os.Stdout, "Response from `MonitoringApi.GetSubscription`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subscriptionId** | **string** | Unique identifier for monitoring subscription. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSubscriptionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**MonitoringSubscription**](MonitoringSubscription.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListAlerts

> MonitoringAlertList ListAlerts(ctx).Id(id).PersonId(personId).BusinessId(businessId).Limit(limit).PageToken(pageToken).Execute()

List monitoring alerts

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
    id := []string{"Inner_example"} // []string | Unique identifier for the resource. Multiple IDs can be provided as a comma-separated list.  (optional)
    personId := []string{"Inner_example"} // []string | Unique identifier for the person. Multiple IDs can be provided as a comma-separated list.  (optional)
    businessId := []string{"Inner_example"} // []string | Unique identifier for the business. Multiple IDs can be provided as a comma-separated list.  (optional)
    limit := int32(100) // int32 |  (optional) (default to 100)
    pageToken := "h50ffqz9q5" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MonitoringApi.ListAlerts(context.Background()).Id(id).PersonId(personId).BusinessId(businessId).Limit(limit).PageToken(pageToken).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MonitoringApi.ListAlerts``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListAlerts`: MonitoringAlertList
    fmt.Fprintf(os.Stdout, "Response from `MonitoringApi.ListAlerts`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListAlertsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **[]string** | Unique identifier for the resource. Multiple IDs can be provided as a comma-separated list.  | 
 **personId** | **[]string** | Unique identifier for the person. Multiple IDs can be provided as a comma-separated list.  | 
 **businessId** | **[]string** | Unique identifier for the business. Multiple IDs can be provided as a comma-separated list.  | 
 **limit** | **int32** |  | [default to 100]
 **pageToken** | **string** |  | 

### Return type

[**MonitoringAlertList**](MonitoringAlertList.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListSubscriptions

> MonitoringSubscriptionList ListSubscriptions(ctx).Id(id).PersonId(personId).BusinessId(businessId).Limit(limit).PageToken(pageToken).Execute()

List monitoring subscriptions

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
    id := []string{"Inner_example"} // []string | Unique identifier for the resource. Multiple IDs can be provided as a comma-separated list.  (optional)
    personId := []string{"Inner_example"} // []string | Unique identifier for the person. Multiple IDs can be provided as a comma-separated list.  (optional)
    businessId := []string{"Inner_example"} // []string | Unique identifier for the business. Multiple IDs can be provided as a comma-separated list.  (optional)
    limit := int32(100) // int32 |  (optional) (default to 100)
    pageToken := "h50ffqz9q5" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MonitoringApi.ListSubscriptions(context.Background()).Id(id).PersonId(personId).BusinessId(businessId).Limit(limit).PageToken(pageToken).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MonitoringApi.ListSubscriptions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListSubscriptions`: MonitoringSubscriptionList
    fmt.Fprintf(os.Stdout, "Response from `MonitoringApi.ListSubscriptions`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListSubscriptionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **[]string** | Unique identifier for the resource. Multiple IDs can be provided as a comma-separated list.  | 
 **personId** | **[]string** | Unique identifier for the person. Multiple IDs can be provided as a comma-separated list.  | 
 **businessId** | **[]string** | Unique identifier for the business. Multiple IDs can be provided as a comma-separated list.  | 
 **limit** | **int32** |  | [default to 100]
 **pageToken** | **string** |  | 

### Return type

[**MonitoringSubscriptionList**](MonitoringSubscriptionList.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateAlert

> UpdateAlert(ctx, alertId).MonitoringAlert(monitoringAlert).Execute()

Update a monitoring alert

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
    alertId := "4a09a067-2f01-475d-bc62-50f245536017" // string | Unique identifier for this monitoring alert.
    monitoringAlert := *openapiclient.NewMonitoringAlert() // MonitoringAlert | Monitoring alert to update.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MonitoringApi.UpdateAlert(context.Background(), alertId).MonitoringAlert(monitoringAlert).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MonitoringApi.UpdateAlert``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**alertId** | **string** | Unique identifier for this monitoring alert. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateAlertRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **monitoringAlert** | [**MonitoringAlert**](MonitoringAlert.md) | Monitoring alert to update. | 

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

