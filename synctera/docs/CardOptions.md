# CardOptions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BillingAddress** | Pointer to [**BillingAddress**](BillingAddress.md) |  | [optional] 
**CardPresent** | Pointer to **bool** |  | [optional] [default to false]
**Cvv** | Pointer to **string** |  | [optional] 
**Expiration** | Pointer to **string** |  | [optional] 

## Methods

### NewCardOptions

`func NewCardOptions() *CardOptions`

NewCardOptions instantiates a new CardOptions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCardOptionsWithDefaults

`func NewCardOptionsWithDefaults() *CardOptions`

NewCardOptionsWithDefaults instantiates a new CardOptions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBillingAddress

`func (o *CardOptions) GetBillingAddress() BillingAddress`

GetBillingAddress returns the BillingAddress field if non-nil, zero value otherwise.

### GetBillingAddressOk

`func (o *CardOptions) GetBillingAddressOk() (*BillingAddress, bool)`

GetBillingAddressOk returns a tuple with the BillingAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingAddress

`func (o *CardOptions) SetBillingAddress(v BillingAddress)`

SetBillingAddress sets BillingAddress field to given value.

### HasBillingAddress

`func (o *CardOptions) HasBillingAddress() bool`

HasBillingAddress returns a boolean if a field has been set.

### GetCardPresent

`func (o *CardOptions) GetCardPresent() bool`

GetCardPresent returns the CardPresent field if non-nil, zero value otherwise.

### GetCardPresentOk

`func (o *CardOptions) GetCardPresentOk() (*bool, bool)`

GetCardPresentOk returns a tuple with the CardPresent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardPresent

`func (o *CardOptions) SetCardPresent(v bool)`

SetCardPresent sets CardPresent field to given value.

### HasCardPresent

`func (o *CardOptions) HasCardPresent() bool`

HasCardPresent returns a boolean if a field has been set.

### GetCvv

`func (o *CardOptions) GetCvv() string`

GetCvv returns the Cvv field if non-nil, zero value otherwise.

### GetCvvOk

`func (o *CardOptions) GetCvvOk() (*string, bool)`

GetCvvOk returns a tuple with the Cvv field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCvv

`func (o *CardOptions) SetCvv(v string)`

SetCvv sets Cvv field to given value.

### HasCvv

`func (o *CardOptions) HasCvv() bool`

HasCvv returns a boolean if a field has been set.

### GetExpiration

`func (o *CardOptions) GetExpiration() string`

GetExpiration returns the Expiration field if non-nil, zero value otherwise.

### GetExpirationOk

`func (o *CardOptions) GetExpirationOk() (*string, bool)`

GetExpirationOk returns a tuple with the Expiration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiration

`func (o *CardOptions) SetExpiration(v string)`

SetExpiration sets Expiration field to given value.

### HasExpiration

`func (o *CardOptions) HasExpiration() bool`

HasExpiration returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


