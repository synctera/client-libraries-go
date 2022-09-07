# ScheduleConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Count** | Pointer to **int32** | Number of times to recur. Exactly one of end_date or count must be provided | [optional] 
**EndDate** | Pointer to **string** | End date of the schedule (exclusive). Exactly one of end_date or count must be provided | [optional] 
**Frequency** | **string** |  | 
**Interval** | **int32** | Interval between recurrences, e.g. interval &#x3D; 2 with frequency &#x3D; WEEKLY means every other week. | 
**StartDate** | **string** | Start date of the schedule (inclusive) | 

## Methods

### NewScheduleConfig

`func NewScheduleConfig(frequency string, interval int32, startDate string, ) *ScheduleConfig`

NewScheduleConfig instantiates a new ScheduleConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewScheduleConfigWithDefaults

`func NewScheduleConfigWithDefaults() *ScheduleConfig`

NewScheduleConfigWithDefaults instantiates a new ScheduleConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCount

`func (o *ScheduleConfig) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *ScheduleConfig) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *ScheduleConfig) SetCount(v int32)`

SetCount sets Count field to given value.

### HasCount

`func (o *ScheduleConfig) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetEndDate

`func (o *ScheduleConfig) GetEndDate() string`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *ScheduleConfig) GetEndDateOk() (*string, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *ScheduleConfig) SetEndDate(v string)`

SetEndDate sets EndDate field to given value.

### HasEndDate

`func (o *ScheduleConfig) HasEndDate() bool`

HasEndDate returns a boolean if a field has been set.

### GetFrequency

`func (o *ScheduleConfig) GetFrequency() string`

GetFrequency returns the Frequency field if non-nil, zero value otherwise.

### GetFrequencyOk

`func (o *ScheduleConfig) GetFrequencyOk() (*string, bool)`

GetFrequencyOk returns a tuple with the Frequency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrequency

`func (o *ScheduleConfig) SetFrequency(v string)`

SetFrequency sets Frequency field to given value.


### GetInterval

`func (o *ScheduleConfig) GetInterval() int32`

GetInterval returns the Interval field if non-nil, zero value otherwise.

### GetIntervalOk

`func (o *ScheduleConfig) GetIntervalOk() (*int32, bool)`

GetIntervalOk returns a tuple with the Interval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterval

`func (o *ScheduleConfig) SetInterval(v int32)`

SetInterval sets Interval field to given value.


### GetStartDate

`func (o *ScheduleConfig) GetStartDate() string`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *ScheduleConfig) GetStartDateOk() (*string, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *ScheduleConfig) SetStartDate(v string)`

SetStartDate sets StartDate field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


