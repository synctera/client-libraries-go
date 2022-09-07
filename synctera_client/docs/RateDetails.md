# RateDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccrualPeriod** | **string** |  | 
**Rate** | **int32** | Rate in basis points. E.g. 5 represents 0.05% | 
**ValidFrom** | **string** | Rate effective start date. Inclusive. | 
**ValidTo** | Pointer to **string** | Rate effective end date. Exclusive. | [optional] 

## Methods

### NewRateDetails

`func NewRateDetails(accrualPeriod string, rate int32, validFrom string, ) *RateDetails`

NewRateDetails instantiates a new RateDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRateDetailsWithDefaults

`func NewRateDetailsWithDefaults() *RateDetails`

NewRateDetailsWithDefaults instantiates a new RateDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccrualPeriod

`func (o *RateDetails) GetAccrualPeriod() string`

GetAccrualPeriod returns the AccrualPeriod field if non-nil, zero value otherwise.

### GetAccrualPeriodOk

`func (o *RateDetails) GetAccrualPeriodOk() (*string, bool)`

GetAccrualPeriodOk returns a tuple with the AccrualPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccrualPeriod

`func (o *RateDetails) SetAccrualPeriod(v string)`

SetAccrualPeriod sets AccrualPeriod field to given value.


### GetRate

`func (o *RateDetails) GetRate() int32`

GetRate returns the Rate field if non-nil, zero value otherwise.

### GetRateOk

`func (o *RateDetails) GetRateOk() (*int32, bool)`

GetRateOk returns a tuple with the Rate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRate

`func (o *RateDetails) SetRate(v int32)`

SetRate sets Rate field to given value.


### GetValidFrom

`func (o *RateDetails) GetValidFrom() string`

GetValidFrom returns the ValidFrom field if non-nil, zero value otherwise.

### GetValidFromOk

`func (o *RateDetails) GetValidFromOk() (*string, bool)`

GetValidFromOk returns a tuple with the ValidFrom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidFrom

`func (o *RateDetails) SetValidFrom(v string)`

SetValidFrom sets ValidFrom field to given value.


### GetValidTo

`func (o *RateDetails) GetValidTo() string`

GetValidTo returns the ValidTo field if non-nil, zero value otherwise.

### GetValidToOk

`func (o *RateDetails) GetValidToOk() (*string, bool)`

GetValidToOk returns a tuple with the ValidTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidTo

`func (o *RateDetails) SetValidTo(v string)`

SetValidTo sets ValidTo field to given value.

### HasValidTo

`func (o *RateDetails) HasValidTo() bool`

HasValidTo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


