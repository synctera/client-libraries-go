# WaitlistAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreationTime** | Pointer to **time.Time** | Creation time | [optional] [readonly] 
**Id** | Pointer to **string** | Waitlist ID | [optional] [readonly] 
**LastUpdatedTime** | Pointer to **time.Time** | Most recent updated time | [optional] [readonly] 
**NumAdmitted** | Pointer to **int32** | Current number of prospects in this waitlist with a status of admitted | [optional] [readonly] 
**NumCreated** | Pointer to **int32** | Current number of prospects in this waitlist with a status of created | [optional] [readonly] 
**NumProspects** | Pointer to **int32** | Current number of prospects in this waitlist, in any state | [optional] [readonly] 
**NumVerified** | Pointer to **int32** | Current number of prospects in this waitlist with a status of verified | [optional] [readonly] 
**NumWithdrawn** | Pointer to **int32** | Current number of prospects in this waitlist with a status of withdrawn | [optional] [readonly] 

## Methods

### NewWaitlistAllOf

`func NewWaitlistAllOf() *WaitlistAllOf`

NewWaitlistAllOf instantiates a new WaitlistAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWaitlistAllOfWithDefaults

`func NewWaitlistAllOfWithDefaults() *WaitlistAllOf`

NewWaitlistAllOfWithDefaults instantiates a new WaitlistAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreationTime

`func (o *WaitlistAllOf) GetCreationTime() time.Time`

GetCreationTime returns the CreationTime field if non-nil, zero value otherwise.

### GetCreationTimeOk

`func (o *WaitlistAllOf) GetCreationTimeOk() (*time.Time, bool)`

GetCreationTimeOk returns a tuple with the CreationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationTime

`func (o *WaitlistAllOf) SetCreationTime(v time.Time)`

SetCreationTime sets CreationTime field to given value.

### HasCreationTime

`func (o *WaitlistAllOf) HasCreationTime() bool`

HasCreationTime returns a boolean if a field has been set.

### GetId

`func (o *WaitlistAllOf) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *WaitlistAllOf) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *WaitlistAllOf) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *WaitlistAllOf) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLastUpdatedTime

`func (o *WaitlistAllOf) GetLastUpdatedTime() time.Time`

GetLastUpdatedTime returns the LastUpdatedTime field if non-nil, zero value otherwise.

### GetLastUpdatedTimeOk

`func (o *WaitlistAllOf) GetLastUpdatedTimeOk() (*time.Time, bool)`

GetLastUpdatedTimeOk returns a tuple with the LastUpdatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdatedTime

`func (o *WaitlistAllOf) SetLastUpdatedTime(v time.Time)`

SetLastUpdatedTime sets LastUpdatedTime field to given value.

### HasLastUpdatedTime

`func (o *WaitlistAllOf) HasLastUpdatedTime() bool`

HasLastUpdatedTime returns a boolean if a field has been set.

### GetNumAdmitted

`func (o *WaitlistAllOf) GetNumAdmitted() int32`

GetNumAdmitted returns the NumAdmitted field if non-nil, zero value otherwise.

### GetNumAdmittedOk

`func (o *WaitlistAllOf) GetNumAdmittedOk() (*int32, bool)`

GetNumAdmittedOk returns a tuple with the NumAdmitted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumAdmitted

`func (o *WaitlistAllOf) SetNumAdmitted(v int32)`

SetNumAdmitted sets NumAdmitted field to given value.

### HasNumAdmitted

`func (o *WaitlistAllOf) HasNumAdmitted() bool`

HasNumAdmitted returns a boolean if a field has been set.

### GetNumCreated

`func (o *WaitlistAllOf) GetNumCreated() int32`

GetNumCreated returns the NumCreated field if non-nil, zero value otherwise.

### GetNumCreatedOk

`func (o *WaitlistAllOf) GetNumCreatedOk() (*int32, bool)`

GetNumCreatedOk returns a tuple with the NumCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumCreated

`func (o *WaitlistAllOf) SetNumCreated(v int32)`

SetNumCreated sets NumCreated field to given value.

### HasNumCreated

`func (o *WaitlistAllOf) HasNumCreated() bool`

HasNumCreated returns a boolean if a field has been set.

### GetNumProspects

`func (o *WaitlistAllOf) GetNumProspects() int32`

GetNumProspects returns the NumProspects field if non-nil, zero value otherwise.

### GetNumProspectsOk

`func (o *WaitlistAllOf) GetNumProspectsOk() (*int32, bool)`

GetNumProspectsOk returns a tuple with the NumProspects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumProspects

`func (o *WaitlistAllOf) SetNumProspects(v int32)`

SetNumProspects sets NumProspects field to given value.

### HasNumProspects

`func (o *WaitlistAllOf) HasNumProspects() bool`

HasNumProspects returns a boolean if a field has been set.

### GetNumVerified

`func (o *WaitlistAllOf) GetNumVerified() int32`

GetNumVerified returns the NumVerified field if non-nil, zero value otherwise.

### GetNumVerifiedOk

`func (o *WaitlistAllOf) GetNumVerifiedOk() (*int32, bool)`

GetNumVerifiedOk returns a tuple with the NumVerified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumVerified

`func (o *WaitlistAllOf) SetNumVerified(v int32)`

SetNumVerified sets NumVerified field to given value.

### HasNumVerified

`func (o *WaitlistAllOf) HasNumVerified() bool`

HasNumVerified returns a boolean if a field has been set.

### GetNumWithdrawn

`func (o *WaitlistAllOf) GetNumWithdrawn() int32`

GetNumWithdrawn returns the NumWithdrawn field if non-nil, zero value otherwise.

### GetNumWithdrawnOk

`func (o *WaitlistAllOf) GetNumWithdrawnOk() (*int32, bool)`

GetNumWithdrawnOk returns a tuple with the NumWithdrawn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumWithdrawn

`func (o *WaitlistAllOf) SetNumWithdrawn(v int32)`

SetNumWithdrawn sets NumWithdrawn field to given value.

### HasNumWithdrawn

`func (o *WaitlistAllOf) HasNumWithdrawn() bool`

HasNumWithdrawn returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


