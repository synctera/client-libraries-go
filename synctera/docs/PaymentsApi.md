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
    outgoingAch := *openapiclient.NewOutgoingAch(int32(607), "USD", "2710879c-96fa-4663-96bc-0b4f3fb1eafb", "debit", "3497087c-e3cc-4d4d-9edf-7061653e9d08", "d06494f7-0450-4b1d-b46d-ad72926e474d", *openapiclient.NewRiskData("ClientIp_example")) // OutgoingAch | Outgoing ACH

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

