# HoldData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | **int32** |  | 
**AvailabilityTime** | **time.Time** |  | 

## Methods

### NewHoldData

`func NewHoldData(amount int32, availabilityTime time.Time, ) *HoldData`

NewHoldData instantiates a new HoldData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHoldDataWithDefaults

`func NewHoldDataWithDefaults() *HoldData`

NewHoldDataWithDefaults instantiates a new HoldData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmount

`func (o *HoldData) GetAmount() int32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *HoldData) GetAmountOk() (*int32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *HoldData) SetAmount(v int32)`

SetAmount sets Amount field to given value.


### GetAvailabilityTime

`func (o *HoldData) GetAvailabilityTime() time.Time`

GetAvailabilityTime returns the AvailabilityTime field if non-nil, zero value otherwise.

### GetAvailabilityTimeOk

`func (o *HoldData) GetAvailabilityTimeOk() (*time.Time, bool)`

GetAvailabilityTimeOk returns a tuple with the AvailabilityTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailabilityTime

`func (o *HoldData) SetAvailabilityTime(v time.Time)`

SetAvailabilityTime sets AvailabilityTime field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


