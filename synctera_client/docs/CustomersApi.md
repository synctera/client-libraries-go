# \CustomersApi

All URIs are relative to *https://api.synctera.com/v0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateCustomer**](CustomersApi.md#CreateCustomer) | **Post** /customers | Create a Customer
[**CreateCustomerEmployment**](CustomersApi.md#CreateCustomerEmployment) | **Post** /customers/{customer_id}/employment | Create employment record
[**CreateCustomerRiskRating**](CustomersApi.md#CreateCustomerRiskRating) | **Post** /customers/{customer_id}/risk_ratings | Create customer risk rating
[**GetAllCustomerEmployment**](CustomersApi.md#GetAllCustomerEmployment) | **Get** /customers/{customer_id}/employment | List customer employment records
[**GetAllCustomerRiskRatings**](CustomersApi.md#GetAllCustomerRiskRatings) | **Get** /customers/{customer_id}/risk_ratings | List customer risk ratings
[**GetCustomer**](CustomersApi.md#GetCustomer) | **Get** /customers/{customer_id} | Get Customer
[**GetCustomerRiskRating**](CustomersApi.md#GetCustomerRiskRating) | **Get** /customers/{customer_id}/risk_ratings/{risk_rating_id} | Get customer risk rating
[**GetPartyEmployment**](CustomersApi.md#GetPartyEmployment) | **Get** /customers/{customer_id}/employment/{employment_id} | Get customer employment record
[**ListCustomers**](CustomersApi.md#ListCustomers) | **Get** /customers | List Customers
[**PatchCustomer**](CustomersApi.md#PatchCustomer) | **Patch** /customers/{customer_id} | Patch Customer
[**PrefillCustomer**](CustomersApi.md#PrefillCustomer) | **Post** /customers/{customer_id}/prefill | Prefill customer
[**UpdateCustomer**](CustomersApi.md#UpdateCustomer) | **Put** /customers/{customer_id} | Update Customer
[**UpdatePartyEmployment**](CustomersApi.md#UpdatePartyEmployment) | **Put** /customers/{customer_id}/employment/{employment_id} | Update customer employment record



## CreateCustomer

> CustomerInBody CreateCustomer(ctx).CustomerInBody(customerInBody).IdempotencyKey(idempotencyKey).Execute()

Create a Customer



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
    customerInBody := openapiclient.customer_in_body{Customer: openapiclient.NewCustomer("Status_example")} // CustomerInBody | Customer to create
    idempotencyKey := "df122e6f-2ba8-48a5-9508-4350bba5f27e" // string | An idempotency key is an arbitrary unique value generated by client to detect subsequent retries of the same request. It is recommended that a UUID or a similar random identifier be used as an idempotency key (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CustomersApi.CreateCustomer(context.Background()).CustomerInBody(customerInBody).IdempotencyKey(idempotencyKey).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomersApi.CreateCustomer``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateCustomer`: CustomerInBody
    fmt.Fprintf(os.Stdout, "Response from `CustomersApi.CreateCustomer`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateCustomerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **customerInBody** | [**CustomerInBody**](CustomerInBody.md) | Customer to create | 
 **idempotencyKey** | **string** | An idempotency key is an arbitrary unique value generated by client to detect subsequent retries of the same request. It is recommended that a UUID or a similar random identifier be used as an idempotency key | 

### Return type

[**CustomerInBody**](CustomerInBody.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateCustomerEmployment

> Employment CreateCustomerEmployment(ctx, customerId).Employment(employment).IdempotencyKey(idempotencyKey).Execute()

Create employment record



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
    employment := *openapiclient.NewEmployment("ABC, Inc.") // Employment | Customer employment record to create.
    idempotencyKey := "df122e6f-2ba8-48a5-9508-4350bba5f27e" // string | An idempotency key is an arbitrary unique value generated by client to detect subsequent retries of the same request. It is recommended that a UUID or a similar random identifier be used as an idempotency key (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CustomersApi.CreateCustomerEmployment(context.Background(), customerId).Employment(employment).IdempotencyKey(idempotencyKey).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomersApi.CreateCustomerEmployment``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateCustomerEmployment`: Employment
    fmt.Fprintf(os.Stdout, "Response from `CustomersApi.CreateCustomerEmployment`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customerId** | **string** | The customer&#39;s unique identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateCustomerEmploymentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **employment** | [**Employment**](Employment.md) | Customer employment record to create. | 
 **idempotencyKey** | **string** | An idempotency key is an arbitrary unique value generated by client to detect subsequent retries of the same request. It is recommended that a UUID or a similar random identifier be used as an idempotency key | 

### Return type

[**Employment**](Employment.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateCustomerRiskRating

> RiskRating CreateCustomerRiskRating(ctx, customerId).IdempotencyKey(idempotencyKey).RiskRating(riskRating).Execute()

Create customer risk rating



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
    idempotencyKey := "df122e6f-2ba8-48a5-9508-4350bba5f27e" // string | An idempotency key is an arbitrary unique value generated by client to detect subsequent retries of the same request. It is recommended that a UUID or a similar random identifier be used as an idempotency key (optional)
    riskRating := *openapiclient.NewRiskRating() // RiskRating | Create a customer risk rating. With a risk rating in request body, Synctera validates the data and saves it. With no request body, Synctera runs an automated risk analysis and saves the result.  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CustomersApi.CreateCustomerRiskRating(context.Background(), customerId).IdempotencyKey(idempotencyKey).RiskRating(riskRating).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomersApi.CreateCustomerRiskRating``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateCustomerRiskRating`: RiskRating
    fmt.Fprintf(os.Stdout, "Response from `CustomersApi.CreateCustomerRiskRating`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customerId** | **string** | The customer&#39;s unique identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateCustomerRiskRatingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **idempotencyKey** | **string** | An idempotency key is an arbitrary unique value generated by client to detect subsequent retries of the same request. It is recommended that a UUID or a similar random identifier be used as an idempotency key | 
 **riskRating** | [**RiskRating**](RiskRating.md) | Create a customer risk rating. With a risk rating in request body, Synctera validates the data and saves it. With no request body, Synctera runs an automated risk analysis and saves the result.  | 

### Return type

[**RiskRating**](RiskRating.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAllCustomerEmployment

> EmploymentList GetAllCustomerEmployment(ctx, customerId).IncludeHistory(includeHistory).Execute()

List customer employment records



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
    includeHistory := true // bool | If true, include old (inactive) records as well. (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CustomersApi.GetAllCustomerEmployment(context.Background(), customerId).IncludeHistory(includeHistory).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomersApi.GetAllCustomerEmployment``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAllCustomerEmployment`: EmploymentList
    fmt.Fprintf(os.Stdout, "Response from `CustomersApi.GetAllCustomerEmployment`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customerId** | **string** | The customer&#39;s unique identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAllCustomerEmploymentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **includeHistory** | **bool** | If true, include old (inactive) records as well. | [default to false]

### Return type

[**EmploymentList**](EmploymentList.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAllCustomerRiskRatings

> RiskRatingList GetAllCustomerRiskRatings(ctx, customerId).IncludeHistory(includeHistory).Execute()

List customer risk ratings



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
    includeHistory := true // bool | If true, include old (inactive) records as well. (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CustomersApi.GetAllCustomerRiskRatings(context.Background(), customerId).IncludeHistory(includeHistory).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomersApi.GetAllCustomerRiskRatings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAllCustomerRiskRatings`: RiskRatingList
    fmt.Fprintf(os.Stdout, "Response from `CustomersApi.GetAllCustomerRiskRatings`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customerId** | **string** | The customer&#39;s unique identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAllCustomerRiskRatingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **includeHistory** | **bool** | If true, include old (inactive) records as well. | [default to false]

### Return type

[**RiskRatingList**](RiskRatingList.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCustomer

> CustomerInBody GetCustomer(ctx, customerId).Execute()

Get Customer



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
    resp, r, err := apiClient.CustomersApi.GetCustomer(context.Background(), customerId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomersApi.GetCustomer``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCustomer`: CustomerInBody
    fmt.Fprintf(os.Stdout, "Response from `CustomersApi.GetCustomer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customerId** | **string** | The customer&#39;s unique identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCustomerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CustomerInBody**](CustomerInBody.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCustomerRiskRating

> RiskRating GetCustomerRiskRating(ctx, customerId, riskRatingId).Execute()

Get customer risk rating



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
    riskRatingId := "25ae35db-92b9-4b74-82d9-140a07eece71" // string | Risk Rating ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CustomersApi.GetCustomerRiskRating(context.Background(), customerId, riskRatingId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomersApi.GetCustomerRiskRating``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCustomerRiskRating`: RiskRating
    fmt.Fprintf(os.Stdout, "Response from `CustomersApi.GetCustomerRiskRating`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customerId** | **string** | The customer&#39;s unique identifier | 
**riskRatingId** | **string** | Risk Rating ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCustomerRiskRatingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**RiskRating**](RiskRating.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPartyEmployment

> Employment GetPartyEmployment(ctx, customerId, employmentId).Execute()

Get customer employment record



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
    employmentId := "4675ebf0-0691-4a2b-b1db-7ca6f4ff9ec5" // string | Unique ID for the employment record.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CustomersApi.GetPartyEmployment(context.Background(), customerId, employmentId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomersApi.GetPartyEmployment``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPartyEmployment`: Employment
    fmt.Fprintf(os.Stdout, "Response from `CustomersApi.GetPartyEmployment`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customerId** | **string** | The customer&#39;s unique identifier | 
**employmentId** | **string** | Unique ID for the employment record. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPartyEmploymentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**Employment**](Employment.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListCustomers

> CustomerList ListCustomers(ctx).Id(id).FirstName(firstName).LastName(lastName).PhoneNumber(phoneNumber).Email(email).SsnLast4(ssnLast4).Status(status).Limit(limit).PageToken(pageToken).SortBy(sortBy).Execute()

List Customers



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
    id := []string{"81026fb3-d06c-4b37-80da-2b17b4749a3f"} // []string | Unique resource identifier (optional)
    firstName := "Alice" // string |  (optional)
    lastName := "Smith" // string |  (optional)
    phoneNumber := "+12065550100" // string |  (optional)
    email := "john.doe@example.com" // string |  (optional)
    ssnLast4 := "6789" // string |  (optional)
    status := "ACTIVE" // string |  (optional)
    limit := int32(100) // int32 |  (optional) (default to 100)
    pageToken := "h50ffqz9q5" // string |  (optional)
    sortBy := []string{"SortBy_example"} // []string | Specifies the sort order for the returned customers.  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CustomersApi.ListCustomers(context.Background()).Id(id).FirstName(firstName).LastName(lastName).PhoneNumber(phoneNumber).Email(email).SsnLast4(ssnLast4).Status(status).Limit(limit).PageToken(pageToken).SortBy(sortBy).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomersApi.ListCustomers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListCustomers`: CustomerList
    fmt.Fprintf(os.Stdout, "Response from `CustomersApi.ListCustomers`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListCustomersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **[]string** | Unique resource identifier | 
 **firstName** | **string** |  | 
 **lastName** | **string** |  | 
 **phoneNumber** | **string** |  | 
 **email** | **string** |  | 
 **ssnLast4** | **string** |  | 
 **status** | **string** |  | 
 **limit** | **int32** |  | [default to 100]
 **pageToken** | **string** |  | 
 **sortBy** | **[]string** | Specifies the sort order for the returned customers.  | 

### Return type

[**CustomerList**](CustomerList.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchCustomer

> CustomerInBody PatchCustomer(ctx, customerId).PatchCustomer(patchCustomer).IdempotencyKey(idempotencyKey).Execute()

Patch Customer



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
    patchCustomer := *openapiclient.NewPatchCustomer() // PatchCustomer | Customer to be patched
    idempotencyKey := "df122e6f-2ba8-48a5-9508-4350bba5f27e" // string | An idempotency key is an arbitrary unique value generated by client to detect subsequent retries of the same request. It is recommended that a UUID or a similar random identifier be used as an idempotency key (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CustomersApi.PatchCustomer(context.Background(), customerId).PatchCustomer(patchCustomer).IdempotencyKey(idempotencyKey).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomersApi.PatchCustomer``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PatchCustomer`: CustomerInBody
    fmt.Fprintf(os.Stdout, "Response from `CustomersApi.PatchCustomer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customerId** | **string** | The customer&#39;s unique identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchCustomerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **patchCustomer** | [**PatchCustomer**](PatchCustomer.md) | Customer to be patched | 
 **idempotencyKey** | **string** | An idempotency key is an arbitrary unique value generated by client to detect subsequent retries of the same request. It is recommended that a UUID or a similar random identifier be used as an idempotency key | 

### Return type

[**CustomerInBody**](CustomerInBody.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PrefillCustomer

> PrefillRequest PrefillCustomer(ctx, customerId).PrefillRequest(prefillRequest).IdempotencyKey(idempotencyKey).Execute()

Prefill customer



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
    prefillRequest := *openapiclient.NewPrefillRequest() // PrefillRequest | 
    idempotencyKey := "df122e6f-2ba8-48a5-9508-4350bba5f27e" // string | An idempotency key is an arbitrary unique value generated by client to detect subsequent retries of the same request. It is recommended that a UUID or a similar random identifier be used as an idempotency key (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CustomersApi.PrefillCustomer(context.Background(), customerId).PrefillRequest(prefillRequest).IdempotencyKey(idempotencyKey).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomersApi.PrefillCustomer``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PrefillCustomer`: PrefillRequest
    fmt.Fprintf(os.Stdout, "Response from `CustomersApi.PrefillCustomer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customerId** | **string** | The customer&#39;s unique identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiPrefillCustomerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **prefillRequest** | [**PrefillRequest**](PrefillRequest.md) |  | 
 **idempotencyKey** | **string** | An idempotency key is an arbitrary unique value generated by client to detect subsequent retries of the same request. It is recommended that a UUID or a similar random identifier be used as an idempotency key | 

### Return type

[**PrefillRequest**](PrefillRequest.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateCustomer

> CustomerInBody UpdateCustomer(ctx, customerId).CustomerInBody(customerInBody).Execute()

Update Customer



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
    customerInBody := openapiclient.customer_in_body{Customer: openapiclient.NewCustomer("Status_example")} // CustomerInBody | Customer to be updated

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CustomersApi.UpdateCustomer(context.Background(), customerId).CustomerInBody(customerInBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomersApi.UpdateCustomer``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateCustomer`: CustomerInBody
    fmt.Fprintf(os.Stdout, "Response from `CustomersApi.UpdateCustomer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customerId** | **string** | The customer&#39;s unique identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateCustomerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **customerInBody** | [**CustomerInBody**](CustomerInBody.md) | Customer to be updated | 

### Return type

[**CustomerInBody**](CustomerInBody.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdatePartyEmployment

> Employment UpdatePartyEmployment(ctx, customerId, employmentId).Employment(employment).Execute()

Update customer employment record



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
    employmentId := "4675ebf0-0691-4a2b-b1db-7ca6f4ff9ec5" // string | Unique ID for the employment record.
    employment := *openapiclient.NewEmployment("ABC, Inc.") // Employment | Customer employment to update.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CustomersApi.UpdatePartyEmployment(context.Background(), customerId, employmentId).Employment(employment).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomersApi.UpdatePartyEmployment``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdatePartyEmployment`: Employment
    fmt.Fprintf(os.Stdout, "Response from `CustomersApi.UpdatePartyEmployment`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customerId** | **string** | The customer&#39;s unique identifier | 
**employmentId** | **string** | Unique ID for the employment record. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdatePartyEmploymentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **employment** | [**Employment**](Employment.md) | Customer employment to update. | 

### Return type

[**Employment**](Employment.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

