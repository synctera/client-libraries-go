# \PaymentsApi

All URIs are relative to *https://api.synctera.com/v0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddOutgoingACHTransaction**](PaymentsApi.md#AddOutgoingACHTransaction) | **Post** /ach | Create an outgoing ACH



## AddOutgoingACHTransaction

> OutgoingAch AddOutgoingACHTransaction(ctx).OutgoingAch(outgoingAch).Execute()

Create an outgoing ACH



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
    outgoingAch := *openapiclient.NewOutgoingAch(int32(607), "USD", "7e64ebf9-31ab-4cbc-9562-b6e4c2cb988c", "debit", "852f90e6-3327-4b20-a0cd-0bf2a469515a", "d29e4b65-8999-4a51-8fec-1b64c93d1328", *openapiclient.NewRiskData()) // OutgoingAch | Outgoing ACH

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PaymentsApi.AddOutgoingACHTransaction(context.Background()).OutgoingAch(outgoingAch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PaymentsApi.AddOutgoingACHTransaction``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddOutgoingACHTransaction`: OutgoingAch
    fmt.Fprintf(os.Stdout, "Response from `PaymentsApi.AddOutgoingACHTransaction`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAddOutgoingACHTransactionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **outgoingAch** | [**OutgoingAch**](OutgoingAch.md) | Outgoing ACH | 

### Return type

[**OutgoingAch**](OutgoingAch.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

