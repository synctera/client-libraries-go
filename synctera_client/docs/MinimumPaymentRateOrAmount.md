# MinimumPaymentRateOrAmount

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MinAmount** | **int64** | The maximum amount to charge as a minimum payment, in cents. For example, to set the maximum to $30, set this value to 3000.  | 
**Rate** | **int32** | The percentage of the balance to use, in basis points. For example, to set 12.5% of the balance, set this value to 1250.  | 
**Type** | [**MinimumPaymentType**](MinimumPaymentType.md) |  | 

## Methods

### NewMinimumPaymentRateOrAmount

`func NewMinimumPaymentRateOrAmount(minAmount int64, rate int32, type_ MinimumPaymentType, ) *MinimumPaymentRateOrAmount`

NewMinimumPaymentRateOrAmount instantiates a new MinimumPaymentRateOrAmount object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMinimumPaymentRateOrAmountWithDefaults

`func NewMinimumPaymentRateOrAmountWithDefaults() *MinimumPaymentRateOrAmount`

NewMinimumPaymentRateOrAmountWithDefaults instantiates a new MinimumPaymentRateOrAmount object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMinAmount

`func (o *MinimumPaymentRateOrAmount) GetMinAmount() int64`

GetMinAmount returns the MinAmount field if non-nil, zero value otherwise.

### GetMinAmountOk

`func (o *MinimumPaymentRateOrAmount) GetMinAmountOk() (*int64, bool)`

GetMinAmountOk returns a tuple with the MinAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinAmount

`func (o *MinimumPaymentRateOrAmount) SetMinAmount(v int64)`

SetMinAmount sets MinAmount field to given value.


### GetRate

`func (o *MinimumPaymentRateOrAmount) GetRate() int32`

GetRate returns the Rate field if non-nil, zero value otherwise.

### GetRateOk

`func (o *MinimumPaymentRateOrAmount) GetRateOk() (*int32, bool)`

GetRateOk returns a tuple with the Rate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRate

`func (o *MinimumPaymentRateOrAmount) SetRate(v int32)`

SetRate sets Rate field to given value.


### GetType

`func (o *MinimumPaymentRateOrAmount) GetType() MinimumPaymentType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *MinimumPaymentRateOrAmount) GetTypeOk() (*MinimumPaymentType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *MinimumPaymentRateOrAmount) SetType(v MinimumPaymentType)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


