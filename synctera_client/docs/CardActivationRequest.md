# CardActivationRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ActivationCode** | **string** | An activation code provided with the card required to prove possession of the card | 
**CustomerId** | **string** | The ID of the customer for which card is being activated | 

## Methods

### NewCardActivationRequest

`func NewCardActivationRequest(activationCode string, customerId string, ) *CardActivationRequest`

NewCardActivationRequest instantiates a new CardActivationRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCardActivationRequestWithDefaults

`func NewCardActivationRequestWithDefaults() *CardActivationRequest`

NewCardActivationRequestWithDefaults instantiates a new CardActivationRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActivationCode

`func (o *CardActivationRequest) GetActivationCode() string`

GetActivationCode returns the ActivationCode field if non-nil, zero value otherwise.

### GetActivationCodeOk

`func (o *CardActivationRequest) GetActivationCodeOk() (*string, bool)`

GetActivationCodeOk returns a tuple with the ActivationCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivationCode

`func (o *CardActivationRequest) SetActivationCode(v string)`

SetActivationCode sets ActivationCode field to given value.


### GetCustomerId

`func (o *CardActivationRequest) GetCustomerId() string`

GetCustomerId returns the CustomerId field if non-nil, zero value otherwise.

### GetCustomerIdOk

`func (o *CardActivationRequest) GetCustomerIdOk() (*string, bool)`

GetCustomerIdOk returns a tuple with the CustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerId

`func (o *CardActivationRequest) SetCustomerId(v string)`

SetCustomerId sets CustomerId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


