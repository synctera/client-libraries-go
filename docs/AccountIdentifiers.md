# AccountIdentifiers

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Number** | **string** | The account number. Value may be masked, in which case only the last four digits are returned.  | 
**Iban** | Pointer to **string** | The IBAN of the account. Value may be masked, in which case only the last four digits are returned.  | [optional] 

## Methods

### NewAccountIdentifiers

`func NewAccountIdentifiers(number string, ) *AccountIdentifiers`

NewAccountIdentifiers instantiates a new AccountIdentifiers object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountIdentifiersWithDefaults

`func NewAccountIdentifiersWithDefaults() *AccountIdentifiers`

NewAccountIdentifiersWithDefaults instantiates a new AccountIdentifiers object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

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

### HasIban

`func (o *AccountIdentifiers) HasIban() bool`

HasIban returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


