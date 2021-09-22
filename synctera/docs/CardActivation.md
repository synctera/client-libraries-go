# CardActivation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ActivationCode** | Pointer to **string** | An activation code provided with the card required to prove possession of the card | [optional] 

## Methods

### NewCardActivation

`func NewCardActivation() *CardActivation`

NewCardActivation instantiates a new CardActivation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCardActivationWithDefaults

`func NewCardActivationWithDefaults() *CardActivation`

NewCardActivationWithDefaults instantiates a new CardActivation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActivationCode

`func (o *CardActivation) GetActivationCode() string`

GetActivationCode returns the ActivationCode field if non-nil, zero value otherwise.

### GetActivationCodeOk

`func (o *CardActivation) GetActivationCodeOk() (*string, bool)`

GetActivationCodeOk returns a tuple with the ActivationCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivationCode

`func (o *CardActivation) SetActivationCode(v string)`

SetActivationCode sets ActivationCode field to given value.

### HasActivationCode

`func (o *CardActivation) HasActivationCode() bool`

HasActivationCode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


