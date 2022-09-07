# Fee

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | **int64** | Fee amount | 
**Currency** | **string** | Fee currency code in ISO 4217 | 
**FeeType** | **string** | Fee type | 
**Id** | Pointer to **string** | Fee ID | [optional] [readonly] 
**ProductType** | **string** |  | 

## Methods

### NewFee

`func NewFee(amount int64, currency string, feeType string, productType string, ) *Fee`

NewFee instantiates a new Fee object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFeeWithDefaults

`func NewFeeWithDefaults() *Fee`

NewFeeWithDefaults instantiates a new Fee object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmount

`func (o *Fee) GetAmount() int64`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *Fee) GetAmountOk() (*int64, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *Fee) SetAmount(v int64)`

SetAmount sets Amount field to given value.


### GetCurrency

`func (o *Fee) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *Fee) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *Fee) SetCurrency(v string)`

SetCurrency sets Currency field to given value.


### GetFeeType

`func (o *Fee) GetFeeType() string`

GetFeeType returns the FeeType field if non-nil, zero value otherwise.

### GetFeeTypeOk

`func (o *Fee) GetFeeTypeOk() (*string, bool)`

GetFeeTypeOk returns a tuple with the FeeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeeType

`func (o *Fee) SetFeeType(v string)`

SetFeeType sets FeeType field to given value.


### GetId

`func (o *Fee) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Fee) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Fee) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Fee) HasId() bool`

HasId returns a boolean if a field has been set.

### GetProductType

`func (o *Fee) GetProductType() string`

GetProductType returns the ProductType field if non-nil, zero value otherwise.

### GetProductTypeOk

`func (o *Fee) GetProductTypeOk() (*string, bool)`

GetProductTypeOk returns a tuple with the ProductType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductType

`func (o *Fee) SetProductType(v string)`

SetProductType sets ProductType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


