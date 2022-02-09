# \CardWebhooksApi

All URIs are relative to *https://api.synctera.com/v0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateWebhook**](CardWebhooksApi.md#CreateWebhook) | **Post** /cards/webhooks | Create a Webhook
[**EditWebhook**](CardWebhooksApi.md#EditWebhook) | **Put** /cards/webhooks/{webhook_id} | Webhook updates
[**EditWebhookCustomHeaders**](CardWebhooksApi.md#EditWebhookCustomHeaders) | **Put** /cards/webhooks/{webhook_id}/custom_headers | Webhook custom headers updates
[**GetWebhook**](CardWebhooksApi.md#GetWebhook) | **Get** /cards/webhooks/{webhook_id} | Get A Webhook
[**ListWebhooks**](CardWebhooksApi.md#ListWebhooks) | **Get** /cards/webhooks | List Webhooks
[**PingWebhook**](CardWebhooksApi.md#PingWebhook) | **Post** /cards/webhooks/{webhook_id} | Ping Webhook
[**ResendEventNotification**](CardWebhooksApi.md#ResendEventNotification) | **Post** /cards/webhooks/{webhook_id}/resend_event | Resend a Webhook Event Notification



## CreateWebhook

> WebhookResponse CreateWebhook(ctx).CreateWebhookRequest(createWebhookRequest).Execute()

Create a Webhook



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
    createWebhookRequest := *openapiclient.NewCreateWebhookRequest("Name_example", []string{"Events_example"}, *openapiclient.NewWebhookConfig()) // CreateWebhookRequest | Webhook to create

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CardWebhooksApi.CreateWebhook(context.Background()).CreateWebhookRequest(createWebhookRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CardWebhooksApi.CreateWebhook``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateWebhook`: WebhookResponse
    fmt.Fprintf(os.Stdout, "Response from `CardWebhooksApi.CreateWebhook`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateWebhookRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createWebhookRequest** | [**CreateWebhookRequest**](CreateWebhookRequest.md) | Webhook to create | 

### Return type

[**WebhookResponse**](WebhookResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EditWebhook

> WebhookResponse EditWebhook(ctx, webhookId).WebhookRequest(webhookRequest).Execute()

Webhook updates



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
    webhookId := TODO // string | The unique identifier of a webhook
    webhookRequest := *openapiclient.NewWebhookRequest() // WebhookRequest | Webhook edits

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CardWebhooksApi.EditWebhook(context.Background(), webhookId).WebhookRequest(webhookRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CardWebhooksApi.EditWebhook``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EditWebhook`: WebhookResponse
    fmt.Fprintf(os.Stdout, "Response from `CardWebhooksApi.EditWebhook`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**webhookId** | [**string**](.md) | The unique identifier of a webhook | 

### Other Parameters

Other parameters are passed through a pointer to a apiEditWebhookRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **webhookRequest** | [**WebhookRequest**](WebhookRequest.md) | Webhook edits | 

### Return type

[**WebhookResponse**](WebhookResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EditWebhookCustomHeaders

> WebhookResponse EditWebhookCustomHeaders(ctx, webhookId).CustomHeaders(customHeaders).Execute()

Webhook custom headers updates



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
    webhookId := TODO // string | The unique identifier of a webhook
    customHeaders := *openapiclient.NewCustomHeaders() // CustomHeaders | Webhook Custom Header Edits

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CardWebhooksApi.EditWebhookCustomHeaders(context.Background(), webhookId).CustomHeaders(customHeaders).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CardWebhooksApi.EditWebhookCustomHeaders``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EditWebhookCustomHeaders`: WebhookResponse
    fmt.Fprintf(os.Stdout, "Response from `CardWebhooksApi.EditWebhookCustomHeaders`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**webhookId** | [**string**](.md) | The unique identifier of a webhook | 

### Other Parameters

Other parameters are passed through a pointer to a apiEditWebhookCustomHeadersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **customHeaders** | [**CustomHeaders**](CustomHeaders.md) | Webhook Custom Header Edits | 

### Return type

[**WebhookResponse**](WebhookResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetWebhook

> WebhookResponse GetWebhook(ctx, webhookId).Execute()

Get A Webhook



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
    webhookId := TODO // string | The unique identifier of a webhook

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CardWebhooksApi.GetWebhook(context.Background(), webhookId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CardWebhooksApi.GetWebhook``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetWebhook`: WebhookResponse
    fmt.Fprintf(os.Stdout, "Response from `CardWebhooksApi.GetWebhook`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**webhookId** | [**string**](.md) | The unique identifier of a webhook | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetWebhookRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**WebhookResponse**](WebhookResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListWebhooks

> WebhookResponse ListWebhooks(ctx).Execute()

List Webhooks



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

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CardWebhooksApi.ListWebhooks(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CardWebhooksApi.ListWebhooks``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListWebhooks`: WebhookResponse
    fmt.Fprintf(os.Stdout, "Response from `CardWebhooksApi.ListWebhooks`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListWebhooksRequest struct via the builder pattern


### Return type

[**WebhookResponse**](WebhookResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PingWebhook

> PingResponse PingWebhook(ctx, webhookId).Execute()

Ping Webhook



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
    webhookId := TODO // string | The unique identifier of a webhook

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CardWebhooksApi.PingWebhook(context.Background(), webhookId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CardWebhooksApi.PingWebhook``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PingWebhook`: PingResponse
    fmt.Fprintf(os.Stdout, "Response from `CardWebhooksApi.PingWebhook`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**webhookId** | [**string**](.md) | The unique identifier of a webhook | 

### Other Parameters

Other parameters are passed through a pointer to a apiPingWebhookRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**PingResponse**](PingResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ResendEventNotification

> ResendResponse ResendEventNotification(ctx, webhookId).EventResend(eventResend).Execute()

Resend a Webhook Event Notification



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
    webhookId := TODO // string | The unique identifier of a webhook
    eventResend := *openapiclient.NewEventResend(openapiclient.event_type("usertransition"), "EventId_example") // EventResend | Resend event notification for webhook

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CardWebhooksApi.ResendEventNotification(context.Background(), webhookId).EventResend(eventResend).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CardWebhooksApi.ResendEventNotification``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ResendEventNotification`: ResendResponse
    fmt.Fprintf(os.Stdout, "Response from `CardWebhooksApi.ResendEventNotification`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**webhookId** | [**string**](.md) | The unique identifier of a webhook | 

### Other Parameters

Other parameters are passed through a pointer to a apiResendEventNotificationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **eventResend** | [**EventResend**](EventResend.md) | Resend event notification for webhook | 

### Return type

[**ResendResponse**](ResendResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

