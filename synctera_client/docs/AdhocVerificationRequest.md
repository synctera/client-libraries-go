# AdhocVerificationRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PayeeId** | **string** | Synctera party (non-customer) who is receiving money from a customer (the payer) | 
**PayerId** | **string** | Synctera customer who is sending money to a non-customer (the payee) | 

## Methods

### NewAdhocVerificationRequest

`func NewAdhocVerificationRequest(payeeId string, payerId string, ) *AdhocVerificationRequest`

NewAdhocVerificationRequest instantiates a new AdhocVerificationRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdhocVerificationRequestWithDefaults

`func NewAdhocVerificationRequestWithDefaults() *AdhocVerificationRequest`

NewAdhocVerificationRequestWithDefaults instantiates a new AdhocVerificationRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPayeeId

`func (o *AdhocVerificationRequest) GetPayeeId() string`

GetPayeeId returns the PayeeId field if non-nil, zero value otherwise.

### GetPayeeIdOk

`func (o *AdhocVerificationRequest) GetPayeeIdOk() (*string, bool)`

GetPayeeIdOk returns a tuple with the PayeeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayeeId

`func (o *AdhocVerificationRequest) SetPayeeId(v string)`

SetPayeeId sets PayeeId field to given value.


### GetPayerId

`func (o *AdhocVerificationRequest) GetPayerId() string`

GetPayerId returns the PayerId field if non-nil, zero value otherwise.

### GetPayerIdOk

`func (o *AdhocVerificationRequest) GetPayerIdOk() (*string, bool)`

GetPayerIdOk returns a tuple with the PayerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayerId

`func (o *AdhocVerificationRequest) SetPayerId(v string)`

SetPayerId sets PayerId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


