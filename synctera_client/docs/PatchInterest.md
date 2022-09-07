# PatchInterest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccrualPayoutSchedule** | Pointer to [**AccrualPayoutSchedule**](AccrualPayoutSchedule.md) |  | [optional] 
**CalculationMethod** | Pointer to [**CalculationMethod**](CalculationMethod.md) |  | [optional] 
**Description** | Pointer to **string** | User provided description for the current interest. | [optional] 
**Id** | Pointer to **string** | Interest ID | [optional] [readonly] 
**ProductType** | **string** |  | 
**Rates** | Pointer to [**[]RateDetails**](RateDetails.md) | A list of interest rate. Date intervals between valid_from and valid_to expect to have no overlap.  | [optional] 

## Methods

### NewPatchInterest

`func NewPatchInterest(productType string, ) *PatchInterest`

NewPatchInterest instantiates a new PatchInterest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchInterestWithDefaults

`func NewPatchInterestWithDefaults() *PatchInterest`

NewPatchInterestWithDefaults instantiates a new PatchInterest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccrualPayoutSchedule

`func (o *PatchInterest) GetAccrualPayoutSchedule() AccrualPayoutSchedule`

GetAccrualPayoutSchedule returns the AccrualPayoutSchedule field if non-nil, zero value otherwise.

### GetAccrualPayoutScheduleOk

`func (o *PatchInterest) GetAccrualPayoutScheduleOk() (*AccrualPayoutSchedule, bool)`

GetAccrualPayoutScheduleOk returns a tuple with the AccrualPayoutSchedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccrualPayoutSchedule

`func (o *PatchInterest) SetAccrualPayoutSchedule(v AccrualPayoutSchedule)`

SetAccrualPayoutSchedule sets AccrualPayoutSchedule field to given value.

### HasAccrualPayoutSchedule

`func (o *PatchInterest) HasAccrualPayoutSchedule() bool`

HasAccrualPayoutSchedule returns a boolean if a field has been set.

### GetCalculationMethod

`func (o *PatchInterest) GetCalculationMethod() CalculationMethod`

GetCalculationMethod returns the CalculationMethod field if non-nil, zero value otherwise.

### GetCalculationMethodOk

`func (o *PatchInterest) GetCalculationMethodOk() (*CalculationMethod, bool)`

GetCalculationMethodOk returns a tuple with the CalculationMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCalculationMethod

`func (o *PatchInterest) SetCalculationMethod(v CalculationMethod)`

SetCalculationMethod sets CalculationMethod field to given value.

### HasCalculationMethod

`func (o *PatchInterest) HasCalculationMethod() bool`

HasCalculationMethod returns a boolean if a field has been set.

### GetDescription

`func (o *PatchInterest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PatchInterest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PatchInterest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PatchInterest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetId

`func (o *PatchInterest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PatchInterest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PatchInterest) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *PatchInterest) HasId() bool`

HasId returns a boolean if a field has been set.

### GetProductType

`func (o *PatchInterest) GetProductType() string`

GetProductType returns the ProductType field if non-nil, zero value otherwise.

### GetProductTypeOk

`func (o *PatchInterest) GetProductTypeOk() (*string, bool)`

GetProductTypeOk returns a tuple with the ProductType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductType

`func (o *PatchInterest) SetProductType(v string)`

SetProductType sets ProductType field to given value.


### GetRates

`func (o *PatchInterest) GetRates() []RateDetails`

GetRates returns the Rates field if non-nil, zero value otherwise.

### GetRatesOk

`func (o *PatchInterest) GetRatesOk() (*[]RateDetails, bool)`

GetRatesOk returns a tuple with the Rates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRates

`func (o *PatchInterest) SetRates(v []RateDetails)`

SetRates sets Rates field to given value.

### HasRates

`func (o *PatchInterest) HasRates() bool`

HasRates returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


