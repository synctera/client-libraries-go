# PatchAccountProduct

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccrualPayoutSchedule** | Pointer to [**AccrualPayoutSchedule**](AccrualPayoutSchedule.md) |  | [optional] 
**CalculationMethod** | Pointer to [**CalculationMethod**](CalculationMethod.md) |  | [optional] 
**Description** | Pointer to **string** | User provided description for the current interest. | [optional] 
**Id** | Pointer to **string** | Fee ID | [optional] [readonly] 
**ProductType** | **string** |  | 
**Rates** | Pointer to [**[]RateDetails**](RateDetails.md) | A list of interest rate. Date intervals between valid_from and valid_to expect to have no overlap.  | [optional] 
**Amount** | **int64** | Fee amount | 
**Currency** | **string** | Fee currency code in ISO 4217 | 
**FeeType** | **string** | Fee type | 

## Methods

### NewPatchAccountProduct

`func NewPatchAccountProduct(productType string, amount int64, currency string, feeType string, ) *PatchAccountProduct`

NewPatchAccountProduct instantiates a new PatchAccountProduct object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchAccountProductWithDefaults

`func NewPatchAccountProductWithDefaults() *PatchAccountProduct`

NewPatchAccountProductWithDefaults instantiates a new PatchAccountProduct object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccrualPayoutSchedule

`func (o *PatchAccountProduct) GetAccrualPayoutSchedule() AccrualPayoutSchedule`

GetAccrualPayoutSchedule returns the AccrualPayoutSchedule field if non-nil, zero value otherwise.

### GetAccrualPayoutScheduleOk

`func (o *PatchAccountProduct) GetAccrualPayoutScheduleOk() (*AccrualPayoutSchedule, bool)`

GetAccrualPayoutScheduleOk returns a tuple with the AccrualPayoutSchedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccrualPayoutSchedule

`func (o *PatchAccountProduct) SetAccrualPayoutSchedule(v AccrualPayoutSchedule)`

SetAccrualPayoutSchedule sets AccrualPayoutSchedule field to given value.

### HasAccrualPayoutSchedule

`func (o *PatchAccountProduct) HasAccrualPayoutSchedule() bool`

HasAccrualPayoutSchedule returns a boolean if a field has been set.

### GetCalculationMethod

`func (o *PatchAccountProduct) GetCalculationMethod() CalculationMethod`

GetCalculationMethod returns the CalculationMethod field if non-nil, zero value otherwise.

### GetCalculationMethodOk

`func (o *PatchAccountProduct) GetCalculationMethodOk() (*CalculationMethod, bool)`

GetCalculationMethodOk returns a tuple with the CalculationMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCalculationMethod

`func (o *PatchAccountProduct) SetCalculationMethod(v CalculationMethod)`

SetCalculationMethod sets CalculationMethod field to given value.

### HasCalculationMethod

`func (o *PatchAccountProduct) HasCalculationMethod() bool`

HasCalculationMethod returns a boolean if a field has been set.

### GetDescription

`func (o *PatchAccountProduct) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PatchAccountProduct) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PatchAccountProduct) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PatchAccountProduct) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetId

`func (o *PatchAccountProduct) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PatchAccountProduct) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PatchAccountProduct) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *PatchAccountProduct) HasId() bool`

HasId returns a boolean if a field has been set.

### GetProductType

`func (o *PatchAccountProduct) GetProductType() string`

GetProductType returns the ProductType field if non-nil, zero value otherwise.

### GetProductTypeOk

`func (o *PatchAccountProduct) GetProductTypeOk() (*string, bool)`

GetProductTypeOk returns a tuple with the ProductType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductType

`func (o *PatchAccountProduct) SetProductType(v string)`

SetProductType sets ProductType field to given value.


### GetRates

`func (o *PatchAccountProduct) GetRates() []RateDetails`

GetRates returns the Rates field if non-nil, zero value otherwise.

### GetRatesOk

`func (o *PatchAccountProduct) GetRatesOk() (*[]RateDetails, bool)`

GetRatesOk returns a tuple with the Rates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRates

`func (o *PatchAccountProduct) SetRates(v []RateDetails)`

SetRates sets Rates field to given value.

### HasRates

`func (o *PatchAccountProduct) HasRates() bool`

HasRates returns a boolean if a field has been set.

### GetAmount

`func (o *PatchAccountProduct) GetAmount() int64`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *PatchAccountProduct) GetAmountOk() (*int64, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *PatchAccountProduct) SetAmount(v int64)`

SetAmount sets Amount field to given value.


### GetCurrency

`func (o *PatchAccountProduct) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *PatchAccountProduct) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *PatchAccountProduct) SetCurrency(v string)`

SetCurrency sets Currency field to given value.


### GetFeeType

`func (o *PatchAccountProduct) GetFeeType() string`

GetFeeType returns the FeeType field if non-nil, zero value otherwise.

### GetFeeTypeOk

`func (o *PatchAccountProduct) GetFeeTypeOk() (*string, bool)`

GetFeeTypeOk returns a tuple with the FeeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeeType

`func (o *PatchAccountProduct) SetFeeType(v string)`

SetFeeType sets FeeType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


