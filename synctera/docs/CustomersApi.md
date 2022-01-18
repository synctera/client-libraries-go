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
[**UpdateCustomer**](CustomersApi.md#UpdateCustomer) | **Put** /customers/{customer_id} | Update Customer
[**UpdatePartyEmployment**](CustomersApi.md#UpdatePartyEmployment) | **Put** /customers/{customer_id}/employment/{employment_id} | Update customer employment record



## CreateCustomer

> CustomerInPath CreateCustomer(ctx).CustomerInPath(customerInPath).Execute()

Create a Customer



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
    customerInPath := openapiclient.customer_in_path{Customer: openapiclient.NewCustomer(time.Now(), "Sallie", "Abernathy", "Status_example")} // CustomerInPath | Customer to create

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CustomersApi.CreateCustomer(context.Background()).CustomerInPath(customerInPath).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomersApi.CreateCustomer``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateCustomer`: CustomerInPath
    fmt.Fprintf(os.Stdout, "Response from `CustomersApi.CreateCustomer`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateCustomerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **customerInPath** | [**CustomerInPath**](CustomerInPath.md) | Customer to create | 

### Return type

[**CustomerInPath**](CustomerInPath.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateCustomerEmployment

> Employment CreateCustomerEmployment(ctx, customerId).Employment(employment).Execute()

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
    customerId := TODO // string | The customer's unique identifier
    employment := *openapiclient.NewEmployment("ABC, Inc.") // Employment | Customer employment record to create.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CustomersApi.CreateCustomerEmployment(context.Background(), customerId).Employment(employment).Execute()
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
**customerId** | [**string**](.md) | The customer&#39;s unique identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateCustomerEmploymentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **employment** | [**Employment**](Employment.md) | Customer employment record to create. | 

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

> RiskRating CreateCustomerRiskRating(ctx, customerId).RiskRating(riskRating).Execute()

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
    customerId := TODO // string | The customer's unique identifier
    riskRating := *openapiclient.NewRiskRating() // RiskRating | Create a customer risk rating. With a risk rating in request body, Synctera validates the data and saves it. With no request body, Synctera runs an automated risk analysis and saves the result.  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CustomersApi.CreateCustomerRiskRating(context.Background(), customerId).RiskRating(riskRating).Execute()
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
**customerId** | [**string**](.md) | The customer&#39;s unique identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateCustomerRiskRatingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

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
    customerId := TODO // string | The customer's unique identifier
    includeHistory := true // bool | If true, include old (inactive) records as well. (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CustomersApi.GetAllCustomerEmployment(context.Background(), customerId).IncludeHistory(includeHistory).Execute()
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
**customerId** | [**string**](.md) | The customer&#39;s unique identifier | 

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
    customerId := TODO // string | The customer's unique identifier
    includeHistory := true // bool | If true, include old (inactive) records as well. (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CustomersApi.GetAllCustomerRiskRatings(context.Background(), customerId).IncludeHistory(includeHistory).Execute()
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
**customerId** | [**string**](.md) | The customer&#39;s unique identifier | 

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

> CustomerInPath GetCustomer(ctx, customerId).Execute()

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
    customerId := TODO // string | The customer's unique identifier

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CustomersApi.GetCustomer(context.Background(), customerId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomersApi.GetCustomer``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCustomer`: CustomerInPath
    fmt.Fprintf(os.Stdout, "Response from `CustomersApi.GetCustomer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customerId** | [**string**](.md) | The customer&#39;s unique identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCustomerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CustomerInPath**](CustomerInPath.md)

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
    customerId := TODO // string | The customer's unique identifier
    riskRatingId := TODO // string | Risk Rating ID

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CustomersApi.GetCustomerRiskRating(context.Background(), customerId, riskRatingId).Execute()
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
**customerId** | [**string**](.md) | The customer&#39;s unique identifier | 
**riskRatingId** | [**string**](.md) | Risk Rating ID | 

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
    customerId := TODO // string | The customer's unique identifier
    employmentId := TODO // string | Unique ID for the employment record.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CustomersApi.GetPartyEmployment(context.Background(), customerId, employmentId).Execute()
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
**customerId** | [**string**](.md) | The customer&#39;s unique identifier | 
**employmentId** | [**string**](.md) | Unique ID for the employment record. | 

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
    id := []string{"Inner_example"} // []string | Only return resources with the specified customer ID(s). Multiple IDs can be provided as a comma-separated list.  (optional)
    firstName := "Alice" // string |  (optional)
    lastName := "Smith" // string |  (optional)
    phoneNumber := "+12065550100" // string |  (optional)
    email := "john.doe@example.com" // string |  (optional)
    ssnLast4 := "6789" // string |  (optional)
    status := "ACTIVE" // string |  (optional)
    limit := int32(100) // int32 |  (optional) (default to 100)
    pageToken := "akic8nczf2" // string |  (optional)
    sortBy := []string{"SortBy_example"} // []string | Specifies the sort order for the returned customers.  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CustomersApi.ListCustomers(context.Background()).Id(id).FirstName(firstName).LastName(lastName).PhoneNumber(phoneNumber).Email(email).SsnLast4(ssnLast4).Status(status).Limit(limit).PageToken(pageToken).SortBy(sortBy).Execute()
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
 **id** | **[]string** | Only return resources with the specified customer ID(s). Multiple IDs can be provided as a comma-separated list.  | 
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

> CustomerInPath PatchCustomer(ctx, customerId).PatchCustomer(patchCustomer).Execute()

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
    customerId := TODO // string | The customer's unique identifier
    patchCustomer := *openapiclient.NewPatchCustomer() // PatchCustomer | Customer to be patched

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CustomersApi.PatchCustomer(context.Background(), customerId).PatchCustomer(patchCustomer).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomersApi.PatchCustomer``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PatchCustomer`: CustomerInPath
    fmt.Fprintf(os.Stdout, "Response from `CustomersApi.PatchCustomer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customerId** | [**string**](.md) | The customer&#39;s unique identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchCustomerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **patchCustomer** | [**PatchCustomer**](PatchCustomer.md) | Customer to be patched | 

### Return type

[**CustomerInPath**](CustomerInPath.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateCustomer

> CustomerInPath UpdateCustomer(ctx, customerId).CustomerInPath(customerInPath).Execute()

Update Customer



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
    customerId := TODO // string | The customer's unique identifier
    customerInPath := openapiclient.customer_in_path{Customer: openapiclient.NewCustomer(time.Now(), "Sallie", "Abernathy", "Status_example")} // CustomerInPath | Customer to be updated

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CustomersApi.UpdateCustomer(context.Background(), customerId).CustomerInPath(customerInPath).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomersApi.UpdateCustomer``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateCustomer`: CustomerInPath
    fmt.Fprintf(os.Stdout, "Response from `CustomersApi.UpdateCustomer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customerId** | [**string**](.md) | The customer&#39;s unique identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateCustomerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **customerInPath** | [**CustomerInPath**](CustomerInPath.md) | Customer to be updated | 

### Return type

[**CustomerInPath**](CustomerInPath.md)

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
    customerId := TODO // string | The customer's unique identifier
    employmentId := TODO // string | Unique ID for the employment record.
    employment := *openapiclient.NewEmployment("ABC, Inc.") // Employment | Customer employment to update.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CustomersApi.UpdatePartyEmployment(context.Background(), customerId, employmentId).Employment(employment).Execute()
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
**customerId** | [**string**](.md) | The customer&#39;s unique identifier | 
**employmentId** | [**string**](.md) | Unique ID for the employment record. | 

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

