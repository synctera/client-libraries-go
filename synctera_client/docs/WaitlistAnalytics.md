# WaitlistAnalytics

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Admitted** | Pointer to **int32** | Total number of prospects that were admitted on the given day | [optional] [readonly] 
**Created** | Pointer to **int32** | Total number of prospects that signed up on the given day | [optional] [readonly] 
**Date** | Pointer to **string** | Date | [optional] [readonly] 
**Deleted** | Pointer to **int32** | Total number of prospects that were deleted on the given day | [optional] [readonly] 
**Verified** | Pointer to **int32** | Total number of prospects that were verified on the given day | [optional] [readonly] 
**Withdrawn** | Pointer to **int32** | Total number of prospects that were withdrawn on the given day | [optional] [readonly] 

## Methods

### NewWaitlistAnalytics

`func NewWaitlistAnalytics() *WaitlistAnalytics`

NewWaitlistAnalytics instantiates a new WaitlistAnalytics object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWaitlistAnalyticsWithDefaults

`func NewWaitlistAnalyticsWithDefaults() *WaitlistAnalytics`

NewWaitlistAnalyticsWithDefaults instantiates a new WaitlistAnalytics object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdmitted

`func (o *WaitlistAnalytics) GetAdmitted() int32`

GetAdmitted returns the Admitted field if non-nil, zero value otherwise.

### GetAdmittedOk

`func (o *WaitlistAnalytics) GetAdmittedOk() (*int32, bool)`

GetAdmittedOk returns a tuple with the Admitted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdmitted

`func (o *WaitlistAnalytics) SetAdmitted(v int32)`

SetAdmitted sets Admitted field to given value.

### HasAdmitted

`func (o *WaitlistAnalytics) HasAdmitted() bool`

HasAdmitted returns a boolean if a field has been set.

### GetCreated

`func (o *WaitlistAnalytics) GetCreated() int32`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *WaitlistAnalytics) GetCreatedOk() (*int32, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *WaitlistAnalytics) SetCreated(v int32)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *WaitlistAnalytics) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetDate

`func (o *WaitlistAnalytics) GetDate() string`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *WaitlistAnalytics) GetDateOk() (*string, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *WaitlistAnalytics) SetDate(v string)`

SetDate sets Date field to given value.

### HasDate

`func (o *WaitlistAnalytics) HasDate() bool`

HasDate returns a boolean if a field has been set.

### GetDeleted

`func (o *WaitlistAnalytics) GetDeleted() int32`

GetDeleted returns the Deleted field if non-nil, zero value otherwise.

### GetDeletedOk

`func (o *WaitlistAnalytics) GetDeletedOk() (*int32, bool)`

GetDeletedOk returns a tuple with the Deleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleted

`func (o *WaitlistAnalytics) SetDeleted(v int32)`

SetDeleted sets Deleted field to given value.

### HasDeleted

`func (o *WaitlistAnalytics) HasDeleted() bool`

HasDeleted returns a boolean if a field has been set.

### GetVerified

`func (o *WaitlistAnalytics) GetVerified() int32`

GetVerified returns the Verified field if non-nil, zero value otherwise.

### GetVerifiedOk

`func (o *WaitlistAnalytics) GetVerifiedOk() (*int32, bool)`

GetVerifiedOk returns a tuple with the Verified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerified

`func (o *WaitlistAnalytics) SetVerified(v int32)`

SetVerified sets Verified field to given value.

### HasVerified

`func (o *WaitlistAnalytics) HasVerified() bool`

HasVerified returns a boolean if a field has been set.

### GetWithdrawn

`func (o *WaitlistAnalytics) GetWithdrawn() int32`

GetWithdrawn returns the Withdrawn field if non-nil, zero value otherwise.

### GetWithdrawnOk

`func (o *WaitlistAnalytics) GetWithdrawnOk() (*int32, bool)`

GetWithdrawnOk returns a tuple with the Withdrawn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWithdrawn

`func (o *WaitlistAnalytics) SetWithdrawn(v int32)`

SetWithdrawn sets Withdrawn field to given value.

### HasWithdrawn

`func (o *WaitlistAnalytics) HasWithdrawn() bool`

HasWithdrawn returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


