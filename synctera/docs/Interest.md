# Interest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccrualPayoutSchedule** | **string** |  | 
**CalculationMethod** | **string** |  | 
**Description** | Pointer to **string** | User provided description for the current interest. | [optional] 
**Id** | Pointer to **string** | Interest ID | [optional] [readonly] 
**ProductType** | **string** |  | 
**Rates** | [**[]RateDetails**](RateDetails.md) | A list of interest rate. Date intervals between valid_from and valid_to expect to have no overlap.  | 

## Methods

### NewInterest

`func NewInterest(accrualPayoutSchedule string, calculationMethod string, productType string, rates []RateDetails, ) *Interest`

NewInterest instantiates a new Interest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInterestWithDefaults

`func NewInterestWithDefaults() *Interest`

NewInterestWithDefaults instantiates a new Interest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccrualPayoutSchedule

`func (o *Interest) GetAccrualPayoutSchedule() string`

GetAccrualPayoutSchedule returns the AccrualPayoutSchedule field if non-nil, zero value otherwise.

### GetAccrualPayoutScheduleOk

`func (o *Interest) GetAccrualPayoutScheduleOk() (*string, bool)`

GetAccrualPayoutScheduleOk returns a tuple with the AccrualPayoutSchedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccrualPayoutSchedule

`func (o *Interest) SetAccrualPayoutSchedule(v string)`

SetAccrualPayoutSchedule sets AccrualPayoutSchedule field to given value.


### GetCalculationMethod

`func (o *Interest) GetCalculationMethod() string`

GetCalculationMethod returns the CalculationMethod field if non-nil, zero value otherwise.

### GetCalculationMethodOk

`func (o *Interest) GetCalculationMethodOk() (*string, bool)`

GetCalculationMethodOk returns a tuple with the CalculationMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCalculationMethod

`func (o *Interest) SetCalculationMethod(v string)`

SetCalculationMethod sets CalculationMethod field to given value.


### GetDescription

`func (o *Interest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Interest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Interest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Interest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetId

`func (o *Interest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Interest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Interest) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Interest) HasId() bool`

HasId returns a boolean if a field has been set.

### GetProductType

`func (o *Interest) GetProductType() string`

GetProductType returns the ProductType field if non-nil, zero value otherwise.

### GetProductTypeOk

`func (o *Interest) GetProductTypeOk() (*string, bool)`

GetProductTypeOk returns a tuple with the ProductType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductType

`func (o *Interest) SetProductType(v string)`

SetProductType sets ProductType field to given value.


### GetRates

`func (o *Interest) GetRates() []RateDetails`

GetRates returns the Rates field if non-nil, zero value otherwise.

### GetRatesOk

`func (o *Interest) GetRatesOk() (*[]RateDetails, bool)`

GetRatesOk returns a tuple with the Rates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRates

`func (o *Interest) SetRates(v []RateDetails)`

SetRates sets Rates field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


