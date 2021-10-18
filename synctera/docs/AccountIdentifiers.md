# AccountIdentifiers

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Iban** | **string** | The IBAN of the account | 
**Number** | **string** | The account number | 

## Methods

### NewAccountIdentifiers

`func NewAccountIdentifiers(iban string, number string, ) *AccountIdentifiers`

NewAccountIdentifiers instantiates a new AccountIdentifiers object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountIdentifiersWithDefaults

`func NewAccountIdentifiersWithDefaults() *AccountIdentifiers`

NewAccountIdentifiersWithDefaults instantiates a new AccountIdentifiers object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIban

`func (o *AccountIdentifiers) GetIban() string`

GetIban returns the Iban field if non-nil, zero value otherwise.

### GetIbanOk

`func (o *AccountIdentifiers) GetIbanOk() (*string, bool)`

GetIbanOk returns a tuple with the Iban field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIban

`func (o *AccountIdentifiers) SetIban(v string)`

SetIban sets Iban field to given value.


### GetNumber

`func (o *AccountIdentifiers) GetNumber() string`

GetNumber returns the Number field if non-nil, zero value otherwise.

### GetNumberOk

`func (o *AccountIdentifiers) GetNumberOk() (*string, bool)`

GetNumberOk returns a tuple with the Number field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumber

`func (o *AccountIdentifiers) SetNumber(v string)`

SetNumber sets Number field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


