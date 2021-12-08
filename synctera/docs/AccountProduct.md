# AccountProduct

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccrualPayoutSchedule** | **string** |  | 
**CalculationMethod** | **string** |  | 
**Description** | Pointer to **string** | User provided description for the current interest. | [optional] 
**Id** | Pointer to **string** | Fee ID | [optional] [readonly] 
**ProductType** | **string** |  | 
**Rates** | [**[]RateDetails**](RateDetails.md) | A list of interest rate. Date intervals between valid_from and valid_to expect to have no overlap.  | 
**Amount** | **int64** | Fee amount | 
**Currency** | **string** | Fee currency code in ISO 4217 | 
**FeeType** | **string** | Fee type | 

## Methods

### NewAccountProduct

`func NewAccountProduct(accrualPayoutSchedule string, calculationMethod string, productType string, rates []RateDetails, amount int64, currency string, feeType string, ) *AccountProduct`

NewAccountProduct instantiates a new AccountProduct object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountProductWithDefaults

`func NewAccountProductWithDefaults() *AccountProduct`

NewAccountProductWithDefaults instantiates a new AccountProduct object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccrualPayoutSchedule

`func (o *AccountProduct) GetAccrualPayoutSchedule() string`

GetAccrualPayoutSchedule returns the AccrualPayoutSchedule field if non-nil, zero value otherwise.

### GetAccrualPayoutScheduleOk

`func (o *AccountProduct) GetAccrualPayoutScheduleOk() (*string, bool)`

GetAccrualPayoutScheduleOk returns a tuple with the AccrualPayoutSchedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccrualPayoutSchedule

`func (o *AccountProduct) SetAccrualPayoutSchedule(v string)`

SetAccrualPayoutSchedule sets AccrualPayoutSchedule field to given value.


### GetCalculationMethod

`func (o *AccountProduct) GetCalculationMethod() string`

GetCalculationMethod returns the CalculationMethod field if non-nil, zero value otherwise.

### GetCalculationMethodOk

`func (o *AccountProduct) GetCalculationMethodOk() (*string, bool)`

GetCalculationMethodOk returns a tuple with the CalculationMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCalculationMethod

`func (o *AccountProduct) SetCalculationMethod(v string)`

SetCalculationMethod sets CalculationMethod field to given value.


### GetDescription

`func (o *AccountProduct) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AccountProduct) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AccountProduct) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AccountProduct) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetId

`func (o *AccountProduct) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AccountProduct) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AccountProduct) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *AccountProduct) HasId() bool`

HasId returns a boolean if a field has been set.

### GetProductType

`func (o *AccountProduct) GetProductType() string`

GetProductType returns the ProductType field if non-nil, zero value otherwise.

### GetProductTypeOk

`func (o *AccountProduct) GetProductTypeOk() (*string, bool)`

GetProductTypeOk returns a tuple with the ProductType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductType

`func (o *AccountProduct) SetProductType(v string)`

SetProductType sets ProductType field to given value.


### GetRates

`func (o *AccountProduct) GetRates() []RateDetails`

GetRates returns the Rates field if non-nil, zero value otherwise.

### GetRatesOk

`func (o *AccountProduct) GetRatesOk() (*[]RateDetails, bool)`

GetRatesOk returns a tuple with the Rates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRates

`func (o *AccountProduct) SetRates(v []RateDetails)`

SetRates sets Rates field to given value.


### GetAmount

`func (o *AccountProduct) GetAmount() int64`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *AccountProduct) GetAmountOk() (*int64, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *AccountProduct) SetAmount(v int64)`

SetAmount sets Amount field to given value.


### GetCurrency

`func (o *AccountProduct) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *AccountProduct) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *AccountProduct) SetCurrency(v string)`

SetCurrency sets Currency field to given value.


### GetFeeType

`func (o *AccountProduct) GetFeeType() string`

GetFeeType returns the FeeType field if non-nil, zero value otherwise.

### GetFeeTypeOk

`func (o *AccountProduct) GetFeeTypeOk() (*string, bool)`

GetFeeTypeOk returns a tuple with the FeeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeeType

`func (o *AccountProduct) SetFeeType(v string)`

SetFeeType sets FeeType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


