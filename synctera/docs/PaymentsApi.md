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
    "time"
    openapiclient "./openapi"
)

func main() {
    outgoingAch := *openapiclient.NewOutgoingAch(int32(607), "USD", "b89508a8-d402-4244-a9c7-4b922c708f47", "debit", time.Now(), "34520cae-2e8a-438d-a233-2b18abb648ad", "653af5db-4c36-4b8e-87f8-9e976fbf5b64", *openapiclient.NewRiskData("ClientIp_example")) // OutgoingAch | Outgoing ACH

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PaymentsApi.AddOutgoingACHTransaction(context.Background()).OutgoingAch(outgoingAch).Execute()
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

