# ExternalAccountBalance

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Available** | Pointer to **NullableInt64** | The available balance of the account | [optional] 
**Currency** | **string** | ISO 4217 alphabetic currency code | 
**Current** | Pointer to **NullableInt64** | The current balance of the account | [optional] 
**Limit** | Pointer to **NullableInt64** | The limit on the balance amount | [optional] 

## Methods

### NewExternalAccountBalance

`func NewExternalAccountBalance(currency string, ) *ExternalAccountBalance`

NewExternalAccountBalance instantiates a new ExternalAccountBalance object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExternalAccountBalanceWithDefaults

`func NewExternalAccountBalanceWithDefaults() *ExternalAccountBalance`

NewExternalAccountBalanceWithDefaults instantiates a new ExternalAccountBalance object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAvailable

`func (o *ExternalAccountBalance) GetAvailable() int64`

GetAvailable returns the Available field if non-nil, zero value otherwise.

### GetAvailableOk

`func (o *ExternalAccountBalance) GetAvailableOk() (*int64, bool)`

GetAvailableOk returns a tuple with the Available field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailable

`func (o *ExternalAccountBalance) SetAvailable(v int64)`

SetAvailable sets Available field to given value.

### HasAvailable

`func (o *ExternalAccountBalance) HasAvailable() bool`

HasAvailable returns a boolean if a field has been set.

### SetAvailableNil

`func (o *ExternalAccountBalance) SetAvailableNil(b bool)`

 SetAvailableNil sets the value for Available to be an explicit nil

### UnsetAvailable
`func (o *ExternalAccountBalance) UnsetAvailable()`

UnsetAvailable ensures that no value is present for Available, not even an explicit nil
### GetCurrency

`func (o *ExternalAccountBalance) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *ExternalAccountBalance) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *ExternalAccountBalance) SetCurrency(v string)`

SetCurrency sets Currency field to given value.


### GetCurrent

`func (o *ExternalAccountBalance) GetCurrent() int64`

GetCurrent returns the Current field if non-nil, zero value otherwise.

### GetCurrentOk

`func (o *ExternalAccountBalance) GetCurrentOk() (*int64, bool)`

GetCurrentOk returns a tuple with the Current field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrent

`func (o *ExternalAccountBalance) SetCurrent(v int64)`

SetCurrent sets Current field to given value.

### HasCurrent

`func (o *ExternalAccountBalance) HasCurrent() bool`

HasCurrent returns a boolean if a field has been set.

### SetCurrentNil

`func (o *ExternalAccountBalance) SetCurrentNil(b bool)`

 SetCurrentNil sets the value for Current to be an explicit nil

### UnsetCurrent
`func (o *ExternalAccountBalance) UnsetCurrent()`

UnsetCurrent ensures that no value is present for Current, not even an explicit nil
### GetLimit

`func (o *ExternalAccountBalance) GetLimit() int64`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *ExternalAccountBalance) GetLimitOk() (*int64, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *ExternalAccountBalance) SetLimit(v int64)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *ExternalAccountBalance) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### SetLimitNil

`func (o *ExternalAccountBalance) SetLimitNil(b bool)`

 SetLimitNil sets the value for Limit to be an explicit nil

### UnsetLimit
`func (o *ExternalAccountBalance) UnsetLimit()`

UnsetLimit ensures that no value is present for Limit, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


