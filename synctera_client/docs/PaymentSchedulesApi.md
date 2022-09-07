# \PaymentSchedulesApi

All URIs are relative to *https://api.synctera.com/v0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreatePaymentSchedule**](PaymentSchedulesApi.md#CreatePaymentSchedule) | **Post** /payment_schedules | Create a payment schedule
[**ListPaymentSchedules**](PaymentSchedulesApi.md#ListPaymentSchedules) | **Get** /payment_schedules | List payment schedules
[**ListPayments**](PaymentSchedulesApi.md#ListPayments) | **Get** /payment_schedules/payments | List payments
[**PatchPaymentSchedule**](PaymentSchedulesApi.md#PatchPaymentSchedule) | **Patch** /payment_schedules/{payment_schedule_id} | Update a payment schedule



## CreatePaymentSchedule

> PaymentSchedule CreatePaymentSchedule(ctx).PaymentSchedule(paymentSchedule).Execute()

Create a payment schedule



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
    paymentSchedule := *openapiclient.NewPaymentSchedule("Description_example", openapiclient.payment_instruction{AchInstruction: openapiclient.NewAchInstruction(*openapiclient.NewOutgoingAchRequest(int32(607), "USD", "2071f55a-0aeb-4f62-85a9-68f72856d463", "debit", "4394f57f-3396-4661-bd03-27684791611f", "18b1f30b-227f-4720-9956-4c6805e5cdfa"), "Type_example")}, *openapiclient.NewScheduleConfig("Frequency_example", int32(123), time.Now())) // PaymentSchedule | payment schedule to create

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PaymentSchedulesApi.CreatePaymentSchedule(context.Background()).PaymentSchedule(paymentSchedule).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PaymentSchedulesApi.CreatePaymentSchedule``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreatePaymentSchedule`: PaymentSchedule
    fmt.Fprintf(os.Stdout, "Response from `PaymentSchedulesApi.CreatePaymentSchedule`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreatePaymentScheduleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **paymentSchedule** | [**PaymentSchedule**](PaymentSchedule.md) | payment schedule to create | 

### Return type

[**PaymentSchedule**](PaymentSchedule.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListPaymentSchedules

> PaymentScheduleList ListPaymentSchedules(ctx).Limit(limit).PageToken(pageToken).Id(id).AccountId(accountId).CustomerId(customerId).Execute()

List payment schedules



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
    id := []string{"Inner_example"} // []string | IDs. Multiple IDs can be provided as a comma-separated list. (optional)
    accountId := []string{"Inner_example"} // []string | Originating account IDs. Multiple IDs can be provided as a comma-separated list. (optional)
    customerId := []string{"Inner_example"} // []string | The IDs of customers who created the payment schedules. Multiple IDs can be provided as a comma-separated list. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PaymentSchedulesApi.ListPaymentSchedules(context.Background()).Limit(limit).PageToken(pageToken).Id(id).AccountId(accountId).CustomerId(customerId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PaymentSchedulesApi.ListPaymentSchedules``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListPaymentSchedules`: PaymentScheduleList
    fmt.Fprintf(os.Stdout, "Response from `PaymentSchedulesApi.ListPaymentSchedules`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListPaymentSchedulesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** |  | [default to 100]
 **pageToken** | **string** |  | 
 **id** | **[]string** | IDs. Multiple IDs can be provided as a comma-separated list. | 
 **accountId** | **[]string** | Originating account IDs. Multiple IDs can be provided as a comma-separated list. | 
 **customerId** | **[]string** | The IDs of customers who created the payment schedules. Multiple IDs can be provided as a comma-separated list. | 

### Return type

[**PaymentScheduleList**](PaymentScheduleList.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListPayments

> PaymentList ListPayments(ctx).Limit(limit).PageToken(pageToken).Id(id).ScheduleId(scheduleId).AccountId(accountId).CustomerId(customerId).Execute()

List payments



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
    id := []string{"Inner_example"} // []string | IDs. Multiple IDs can be provided as a comma-separated list. (optional)
    scheduleId := []string{"Inner_example"} // []string | Payment schedule IDs. Multiple IDs can be provided as a comma-separated list. (optional)
    accountId := []string{"Inner_example"} // []string | Originating account IDs. Multiple IDs can be provided as a comma-separated list. (optional)
    customerId := []string{"Inner_example"} // []string | The IDs of customers who created the payment schedules. Multiple IDs can be provided as a comma-separated list. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PaymentSchedulesApi.ListPayments(context.Background()).Limit(limit).PageToken(pageToken).Id(id).ScheduleId(scheduleId).AccountId(accountId).CustomerId(customerId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PaymentSchedulesApi.ListPayments``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListPayments`: PaymentList
    fmt.Fprintf(os.Stdout, "Response from `PaymentSchedulesApi.ListPayments`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListPaymentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** |  | [default to 100]
 **pageToken** | **string** |  | 
 **id** | **[]string** | IDs. Multiple IDs can be provided as a comma-separated list. | 
 **scheduleId** | **[]string** | Payment schedule IDs. Multiple IDs can be provided as a comma-separated list. | 
 **accountId** | **[]string** | Originating account IDs. Multiple IDs can be provided as a comma-separated list. | 
 **customerId** | **[]string** | The IDs of customers who created the payment schedules. Multiple IDs can be provided as a comma-separated list. | 

### Return type

[**PaymentList**](PaymentList.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchPaymentSchedule

> PaymentSchedule PatchPaymentSchedule(ctx, paymentScheduleId).PatchPaymentSchedule(patchPaymentSchedule).Execute()

Update a payment schedule



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
    paymentScheduleId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Payment schedule ID
    patchPaymentSchedule := *openapiclient.NewPatchPaymentSchedule() // PatchPaymentSchedule | payment schedule to update

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PaymentSchedulesApi.PatchPaymentSchedule(context.Background(), paymentScheduleId).PatchPaymentSchedule(patchPaymentSchedule).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PaymentSchedulesApi.PatchPaymentSchedule``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PatchPaymentSchedule`: PaymentSchedule
    fmt.Fprintf(os.Stdout, "Response from `PaymentSchedulesApi.PatchPaymentSchedule`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**paymentScheduleId** | **string** | Payment schedule ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchPaymentScheduleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **patchPaymentSchedule** | [**PatchPaymentSchedule**](PatchPaymentSchedule.md) | payment schedule to update | 

### Return type

[**PaymentSchedule**](PaymentSchedule.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

