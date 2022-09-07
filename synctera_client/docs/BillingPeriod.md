# BillingPeriod

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Frequency** | **string** | The frequency of billing cycles. Along with the start date, this will determine the start and end of each cycle.  | 
**StartDate** | **time.Time** | The first day of the first billing cycle for this account. For a monthly billing cycle, this would determine the day of the month each billing cycle will start on.  | 

## Methods

### NewBillingPeriod

`func NewBillingPeriod(frequency string, startDate time.Time, ) *BillingPeriod`

NewBillingPeriod instantiates a new BillingPeriod object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBillingPeriodWithDefaults

`func NewBillingPeriodWithDefaults() *BillingPeriod`

NewBillingPeriodWithDefaults instantiates a new BillingPeriod object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFrequency

`func (o *BillingPeriod) GetFrequency() string`

GetFrequency returns the Frequency field if non-nil, zero value otherwise.

### GetFrequencyOk

`func (o *BillingPeriod) GetFrequencyOk() (*string, bool)`

GetFrequencyOk returns a tuple with the Frequency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrequency

`func (o *BillingPeriod) SetFrequency(v string)`

SetFrequency sets Frequency field to given value.


### GetStartDate

`func (o *BillingPeriod) GetStartDate() time.Time`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *BillingPeriod) GetStartDateOk() (*time.Time, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *BillingPeriod) SetStartDate(v time.Time)`

SetStartDate sets StartDate field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


