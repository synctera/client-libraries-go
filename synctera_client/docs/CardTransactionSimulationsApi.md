# \CardTransactionSimulationsApi

All URIs are relative to *https://api.synctera.com/v0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SimulateAuthorization**](CardTransactionSimulationsApi.md#SimulateAuthorization) | **Post** /cards/transaction_simulations/authorization | Simulate authorization
[**SimulateAuthorizationAdvice**](CardTransactionSimulationsApi.md#SimulateAuthorizationAdvice) | **Post** /cards/transaction_simulations/authorization/advice | Simulate authorization advice
[**SimulateBalanceInquiry**](CardTransactionSimulationsApi.md#SimulateBalanceInquiry) | **Post** /cards/transaction_simulations/financial/balance_inquiry | Simulate balance inquiry
[**SimulateClearing**](CardTransactionSimulationsApi.md#SimulateClearing) | **Post** /cards/transaction_simulations/clearing | Simulate clearing or refund
[**SimulateFinancial**](CardTransactionSimulationsApi.md#SimulateFinancial) | **Post** /cards/transaction_simulations/financial | Simulate financial
[**SimulateFinancialAdvice**](CardTransactionSimulationsApi.md#SimulateFinancialAdvice) | **Post** /cards/transaction_simulations/financial/advice | Simulate financial advice
[**SimulateOriginalCredit**](CardTransactionSimulationsApi.md#SimulateOriginalCredit) | **Post** /cards/transaction_simulations/financial/original_credit | Simulate OCT
[**SimulateReversal**](CardTransactionSimulationsApi.md#SimulateReversal) | **Post** /cards/transaction_simulations/reversal | Simulate reversal
[**SimulateWithdrawal**](CardTransactionSimulationsApi.md#SimulateWithdrawal) | **Post** /cards/transaction_simulations/financial/withdrawal | Simulate ATM withdrawal



## SimulateAuthorization

> map[string]interface{} SimulateAuthorization(ctx).AuthRequestModel(authRequestModel).Execute()

Simulate authorization



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
    authRequestModel := *openapiclient.NewAuthRequestModel(int32(123), "CardId_example", "Mid_example") // AuthRequestModel | Authorization details (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CardTransactionSimulationsApi.SimulateAuthorization(context.Background()).AuthRequestModel(authRequestModel).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CardTransactionSimulationsApi.SimulateAuthorization``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SimulateAuthorization`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `CardTransactionSimulationsApi.SimulateAuthorization`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSimulateAuthorizationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authRequestModel** | [**AuthRequestModel**](AuthRequestModel.md) | Authorization details | 

### Return type

**map[string]interface{}**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SimulateAuthorizationAdvice

> map[string]interface{} SimulateAuthorizationAdvice(ctx).AuthorizationAdviceModel(authorizationAdviceModel).Execute()

Simulate authorization advice



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
    authorizationAdviceModel := *openapiclient.NewAuthorizationAdviceModel(int32(123), "OriginalTransactionId_example") // AuthorizationAdviceModel | Authorization advice details (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CardTransactionSimulationsApi.SimulateAuthorizationAdvice(context.Background()).AuthorizationAdviceModel(authorizationAdviceModel).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CardTransactionSimulationsApi.SimulateAuthorizationAdvice``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SimulateAuthorizationAdvice`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `CardTransactionSimulationsApi.SimulateAuthorizationAdvice`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSimulateAuthorizationAdviceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorizationAdviceModel** | [**AuthorizationAdviceModel**](AuthorizationAdviceModel.md) | Authorization advice details | 

### Return type

**map[string]interface{}**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SimulateBalanceInquiry

> map[string]interface{} SimulateBalanceInquiry(ctx).BalanceInquiryRequestModel(balanceInquiryRequestModel).Execute()

Simulate balance inquiry



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
    balanceInquiryRequestModel := *openapiclient.NewBalanceInquiryRequestModel("AccountType_example", *openapiclient.NewCardAcceptorModel(), "CardId_example", "Mid_example") // BalanceInquiryRequestModel | Balance inquiry details

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CardTransactionSimulationsApi.SimulateBalanceInquiry(context.Background()).BalanceInquiryRequestModel(balanceInquiryRequestModel).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CardTransactionSimulationsApi.SimulateBalanceInquiry``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SimulateBalanceInquiry`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `CardTransactionSimulationsApi.SimulateBalanceInquiry`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSimulateBalanceInquiryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **balanceInquiryRequestModel** | [**BalanceInquiryRequestModel**](BalanceInquiryRequestModel.md) | Balance inquiry details | 

### Return type

**map[string]interface{}**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SimulateClearing

> map[string]interface{} SimulateClearing(ctx).ClearingModel(clearingModel).Execute()

Simulate clearing or refund



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
    clearingModel := *openapiclient.NewClearingModel(int32(123), "OriginalTransactionId_example") // ClearingModel | Transaction clearing or refund details (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CardTransactionSimulationsApi.SimulateClearing(context.Background()).ClearingModel(clearingModel).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CardTransactionSimulationsApi.SimulateClearing``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SimulateClearing`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `CardTransactionSimulationsApi.SimulateClearing`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSimulateClearingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **clearingModel** | [**ClearingModel**](ClearingModel.md) | Transaction clearing or refund details | 

### Return type

**map[string]interface{}**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SimulateFinancial

> map[string]interface{} SimulateFinancial(ctx).FinancialRequestModel(financialRequestModel).Execute()

Simulate financial



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
    financialRequestModel := *openapiclient.NewFinancialRequestModel(int32(123), *openapiclient.NewCardAcceptorModel(), "CardId_example", "Mid_example") // FinancialRequestModel | Financial details

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CardTransactionSimulationsApi.SimulateFinancial(context.Background()).FinancialRequestModel(financialRequestModel).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CardTransactionSimulationsApi.SimulateFinancial``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SimulateFinancial`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `CardTransactionSimulationsApi.SimulateFinancial`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSimulateFinancialRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **financialRequestModel** | [**FinancialRequestModel**](FinancialRequestModel.md) | Financial details | 

### Return type

**map[string]interface{}**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SimulateFinancialAdvice

> map[string]interface{} SimulateFinancialAdvice(ctx).AuthorizationAdviceModel(authorizationAdviceModel).Execute()

Simulate financial advice



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
    authorizationAdviceModel := *openapiclient.NewAuthorizationAdviceModel(int32(123), "OriginalTransactionId_example") // AuthorizationAdviceModel | Financial advice details

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CardTransactionSimulationsApi.SimulateFinancialAdvice(context.Background()).AuthorizationAdviceModel(authorizationAdviceModel).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CardTransactionSimulationsApi.SimulateFinancialAdvice``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SimulateFinancialAdvice`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `CardTransactionSimulationsApi.SimulateFinancialAdvice`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSimulateFinancialAdviceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorizationAdviceModel** | [**AuthorizationAdviceModel**](AuthorizationAdviceModel.md) | Financial advice details | 

### Return type

**map[string]interface{}**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SimulateOriginalCredit

> map[string]interface{} SimulateOriginalCredit(ctx).OriginalCreditRequestModel(originalCreditRequestModel).Execute()

Simulate OCT



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
    originalCreditRequestModel := *openapiclient.NewOriginalCreditRequestModel(int32(123), "CardId_example", "Mid_example", "Type_example") // OriginalCreditRequestModel | OCT details

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CardTransactionSimulationsApi.SimulateOriginalCredit(context.Background()).OriginalCreditRequestModel(originalCreditRequestModel).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CardTransactionSimulationsApi.SimulateOriginalCredit``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SimulateOriginalCredit`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `CardTransactionSimulationsApi.SimulateOriginalCredit`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSimulateOriginalCreditRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **originalCreditRequestModel** | [**OriginalCreditRequestModel**](OriginalCreditRequestModel.md) | OCT details | 

### Return type

**map[string]interface{}**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SimulateReversal

> map[string]interface{} SimulateReversal(ctx).ReversalModel(reversalModel).Execute()

Simulate reversal



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
    reversalModel := *openapiclient.NewReversalModel(int32(123), "OriginalTransactionId_example") // ReversalModel | Reversal details

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CardTransactionSimulationsApi.SimulateReversal(context.Background()).ReversalModel(reversalModel).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CardTransactionSimulationsApi.SimulateReversal``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SimulateReversal`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `CardTransactionSimulationsApi.SimulateReversal`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSimulateReversalRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **reversalModel** | [**ReversalModel**](ReversalModel.md) | Reversal details | 

### Return type

**map[string]interface{}**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SimulateWithdrawal

> map[string]interface{} SimulateWithdrawal(ctx).WithdrawalRequestModel(withdrawalRequestModel).Execute()

Simulate ATM withdrawal



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
    withdrawalRequestModel := *openapiclient.NewWithdrawalRequestModel(int32(123), "CardId_example", "Mid_example") // WithdrawalRequestModel | ATM withdrawal details

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CardTransactionSimulationsApi.SimulateWithdrawal(context.Background()).WithdrawalRequestModel(withdrawalRequestModel).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CardTransactionSimulationsApi.SimulateWithdrawal``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SimulateWithdrawal`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `CardTransactionSimulationsApi.SimulateWithdrawal`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSimulateWithdrawalRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **withdrawalRequestModel** | [**WithdrawalRequestModel**](WithdrawalRequestModel.md) | ATM withdrawal details | 

### Return type

**map[string]interface{}**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

