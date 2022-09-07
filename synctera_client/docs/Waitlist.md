# Waitlist

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
**Description** | Pointer to **string** | Waitlist Name | [optional] 
**Increment** | **int32** | Number of points earned per verified referral | 
**MaxProspects** | Pointer to **int32** | Max number of prospects allowed on this waitlist. Default is 10,000,000. | [optional] 
**WaitlistName** | **string** | Waitlist Name | 

## Methods

### NewWaitlist

`func NewWaitlist(increment int32, waitlistName string, ) *Waitlist`

NewWaitlist instantiates a new Waitlist object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWaitlistWithDefaults

`func NewWaitlistWithDefaults() *Waitlist`

NewWaitlistWithDefaults instantiates a new Waitlist object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreationTime

`func (o *Waitlist) GetCreationTime() time.Time`

GetCreationTime returns the CreationTime field if non-nil, zero value otherwise.

### GetCreationTimeOk

`func (o *Waitlist) GetCreationTimeOk() (*time.Time, bool)`

GetCreationTimeOk returns a tuple with the CreationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationTime

`func (o *Waitlist) SetCreationTime(v time.Time)`

SetCreationTime sets CreationTime field to given value.

### HasCreationTime

`func (o *Waitlist) HasCreationTime() bool`

HasCreationTime returns a boolean if a field has been set.

### GetId

`func (o *Waitlist) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Waitlist) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Waitlist) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Waitlist) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLastUpdatedTime

`func (o *Waitlist) GetLastUpdatedTime() time.Time`

GetLastUpdatedTime returns the LastUpdatedTime field if non-nil, zero value otherwise.

### GetLastUpdatedTimeOk

`func (o *Waitlist) GetLastUpdatedTimeOk() (*time.Time, bool)`

GetLastUpdatedTimeOk returns a tuple with the LastUpdatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdatedTime

`func (o *Waitlist) SetLastUpdatedTime(v time.Time)`

SetLastUpdatedTime sets LastUpdatedTime field to given value.

### HasLastUpdatedTime

`func (o *Waitlist) HasLastUpdatedTime() bool`

HasLastUpdatedTime returns a boolean if a field has been set.

### GetNumAdmitted

`func (o *Waitlist) GetNumAdmitted() int32`

GetNumAdmitted returns the NumAdmitted field if non-nil, zero value otherwise.

### GetNumAdmittedOk

`func (o *Waitlist) GetNumAdmittedOk() (*int32, bool)`

GetNumAdmittedOk returns a tuple with the NumAdmitted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumAdmitted

`func (o *Waitlist) SetNumAdmitted(v int32)`

SetNumAdmitted sets NumAdmitted field to given value.

### HasNumAdmitted

`func (o *Waitlist) HasNumAdmitted() bool`

HasNumAdmitted returns a boolean if a field has been set.

### GetNumCreated

`func (o *Waitlist) GetNumCreated() int32`

GetNumCreated returns the NumCreated field if non-nil, zero value otherwise.

### GetNumCreatedOk

`func (o *Waitlist) GetNumCreatedOk() (*int32, bool)`

GetNumCreatedOk returns a tuple with the NumCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumCreated

`func (o *Waitlist) SetNumCreated(v int32)`

SetNumCreated sets NumCreated field to given value.

### HasNumCreated

`func (o *Waitlist) HasNumCreated() bool`

HasNumCreated returns a boolean if a field has been set.

### GetNumProspects

`func (o *Waitlist) GetNumProspects() int32`

GetNumProspects returns the NumProspects field if non-nil, zero value otherwise.

### GetNumProspectsOk

`func (o *Waitlist) GetNumProspectsOk() (*int32, bool)`

GetNumProspectsOk returns a tuple with the NumProspects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumProspects

`func (o *Waitlist) SetNumProspects(v int32)`

SetNumProspects sets NumProspects field to given value.

### HasNumProspects

`func (o *Waitlist) HasNumProspects() bool`

HasNumProspects returns a boolean if a field has been set.

### GetNumVerified

`func (o *Waitlist) GetNumVerified() int32`

GetNumVerified returns the NumVerified field if non-nil, zero value otherwise.

### GetNumVerifiedOk

`func (o *Waitlist) GetNumVerifiedOk() (*int32, bool)`

GetNumVerifiedOk returns a tuple with the NumVerified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumVerified

`func (o *Waitlist) SetNumVerified(v int32)`

SetNumVerified sets NumVerified field to given value.

### HasNumVerified

`func (o *Waitlist) HasNumVerified() bool`

HasNumVerified returns a boolean if a field has been set.

### GetNumWithdrawn

`func (o *Waitlist) GetNumWithdrawn() int32`

GetNumWithdrawn returns the NumWithdrawn field if non-nil, zero value otherwise.

### GetNumWithdrawnOk

`func (o *Waitlist) GetNumWithdrawnOk() (*int32, bool)`

GetNumWithdrawnOk returns a tuple with the NumWithdrawn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumWithdrawn

`func (o *Waitlist) SetNumWithdrawn(v int32)`

SetNumWithdrawn sets NumWithdrawn field to given value.

### HasNumWithdrawn

`func (o *Waitlist) HasNumWithdrawn() bool`

HasNumWithdrawn returns a boolean if a field has been set.

### GetDescription

`func (o *Waitlist) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Waitlist) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Waitlist) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Waitlist) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetIncrement

`func (o *Waitlist) GetIncrement() int32`

GetIncrement returns the Increment field if non-nil, zero value otherwise.

### GetIncrementOk

`func (o *Waitlist) GetIncrementOk() (*int32, bool)`

GetIncrementOk returns a tuple with the Increment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncrement

`func (o *Waitlist) SetIncrement(v int32)`

SetIncrement sets Increment field to given value.


### GetMaxProspects

`func (o *Waitlist) GetMaxProspects() int32`

GetMaxProspects returns the MaxProspects field if non-nil, zero value otherwise.

### GetMaxProspectsOk

`func (o *Waitlist) GetMaxProspectsOk() (*int32, bool)`

GetMaxProspectsOk returns a tuple with the MaxProspects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxProspects

`func (o *Waitlist) SetMaxProspects(v int32)`

SetMaxProspects sets MaxProspects field to given value.

### HasMaxProspects

`func (o *Waitlist) HasMaxProspects() bool`

HasMaxProspects returns a boolean if a field has been set.

### GetWaitlistName

`func (o *Waitlist) GetWaitlistName() string`

GetWaitlistName returns the WaitlistName field if non-nil, zero value otherwise.

### GetWaitlistNameOk

`func (o *Waitlist) GetWaitlistNameOk() (*string, bool)`

GetWaitlistNameOk returns a tuple with the WaitlistName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWaitlistName

`func (o *Waitlist) SetWaitlistName(v string)`

SetWaitlistName sets WaitlistName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


