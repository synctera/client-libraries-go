# PaymentErrorDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | Pointer to **string** |  | [optional] 
**Details** | Pointer to **string** |  | [optional] 

## Methods

### NewPaymentErrorDetails

`func NewPaymentErrorDetails() *PaymentErrorDetails`

NewPaymentErrorDetails instantiates a new PaymentErrorDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaymentErrorDetailsWithDefaults

`func NewPaymentErrorDetailsWithDefaults() *PaymentErrorDetails`

NewPaymentErrorDetailsWithDefaults instantiates a new PaymentErrorDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *PaymentErrorDetails) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *PaymentErrorDetails) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *PaymentErrorDetails) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *PaymentErrorDetails) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetDetails

`func (o *PaymentErrorDetails) GetDetails() string`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *PaymentErrorDetails) GetDetailsOk() (*string, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetails

`func (o *PaymentErrorDetails) SetDetails(v string)`

SetDetails sets Details field to given value.

### HasDetails

`func (o *PaymentErrorDetails) HasDetails() bool`

HasDetails returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


