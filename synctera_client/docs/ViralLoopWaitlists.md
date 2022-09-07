# ViralLoopWaitlists

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BankId** | Pointer to **int32** | Bank ID | [optional] [readonly] 
**CreationTime** | Pointer to **time.Time** | Creation time | [optional] [readonly] 
**Id** | **string** | Viral Loop Waitlist ID | [readonly] 
**LastUpdatedTime** | Pointer to **time.Time** | Most recent updated time | [optional] [readonly] 
**LeadCount** | Pointer to **int32** | Total number of leads/people on the waitlist | [optional] [readonly] 
**PartnerId** | Pointer to **int32** | Partner ID | [optional] [readonly] 
**ViralLoopApiToken** | **string** | viral loop api token | 

## Methods

### NewViralLoopWaitlists

`func NewViralLoopWaitlists(id string, viralLoopApiToken string, ) *ViralLoopWaitlists`

NewViralLoopWaitlists instantiates a new ViralLoopWaitlists object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewViralLoopWaitlistsWithDefaults

`func NewViralLoopWaitlistsWithDefaults() *ViralLoopWaitlists`

NewViralLoopWaitlistsWithDefaults instantiates a new ViralLoopWaitlists object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBankId

`func (o *ViralLoopWaitlists) GetBankId() int32`

GetBankId returns the BankId field if non-nil, zero value otherwise.

### GetBankIdOk

`func (o *ViralLoopWaitlists) GetBankIdOk() (*int32, bool)`

GetBankIdOk returns a tuple with the BankId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBankId

`func (o *ViralLoopWaitlists) SetBankId(v int32)`

SetBankId sets BankId field to given value.

### HasBankId

`func (o *ViralLoopWaitlists) HasBankId() bool`

HasBankId returns a boolean if a field has been set.

### GetCreationTime

`func (o *ViralLoopWaitlists) GetCreationTime() time.Time`

GetCreationTime returns the CreationTime field if non-nil, zero value otherwise.

### GetCreationTimeOk

`func (o *ViralLoopWaitlists) GetCreationTimeOk() (*time.Time, bool)`

GetCreationTimeOk returns a tuple with the CreationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationTime

`func (o *ViralLoopWaitlists) SetCreationTime(v time.Time)`

SetCreationTime sets CreationTime field to given value.

### HasCreationTime

`func (o *ViralLoopWaitlists) HasCreationTime() bool`

HasCreationTime returns a boolean if a field has been set.

### GetId

`func (o *ViralLoopWaitlists) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ViralLoopWaitlists) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ViralLoopWaitlists) SetId(v string)`

SetId sets Id field to given value.


### GetLastUpdatedTime

`func (o *ViralLoopWaitlists) GetLastUpdatedTime() time.Time`

GetLastUpdatedTime returns the LastUpdatedTime field if non-nil, zero value otherwise.

### GetLastUpdatedTimeOk

`func (o *ViralLoopWaitlists) GetLastUpdatedTimeOk() (*time.Time, bool)`

GetLastUpdatedTimeOk returns a tuple with the LastUpdatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdatedTime

`func (o *ViralLoopWaitlists) SetLastUpdatedTime(v time.Time)`

SetLastUpdatedTime sets LastUpdatedTime field to given value.

### HasLastUpdatedTime

`func (o *ViralLoopWaitlists) HasLastUpdatedTime() bool`

HasLastUpdatedTime returns a boolean if a field has been set.

### GetLeadCount

`func (o *ViralLoopWaitlists) GetLeadCount() int32`

GetLeadCount returns the LeadCount field if non-nil, zero value otherwise.

### GetLeadCountOk

`func (o *ViralLoopWaitlists) GetLeadCountOk() (*int32, bool)`

GetLeadCountOk returns a tuple with the LeadCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeadCount

`func (o *ViralLoopWaitlists) SetLeadCount(v int32)`

SetLeadCount sets LeadCount field to given value.

### HasLeadCount

`func (o *ViralLoopWaitlists) HasLeadCount() bool`

HasLeadCount returns a boolean if a field has been set.

### GetPartnerId

`func (o *ViralLoopWaitlists) GetPartnerId() int32`

GetPartnerId returns the PartnerId field if non-nil, zero value otherwise.

### GetPartnerIdOk

`func (o *ViralLoopWaitlists) GetPartnerIdOk() (*int32, bool)`

GetPartnerIdOk returns a tuple with the PartnerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartnerId

`func (o *ViralLoopWaitlists) SetPartnerId(v int32)`

SetPartnerId sets PartnerId field to given value.

### HasPartnerId

`func (o *ViralLoopWaitlists) HasPartnerId() bool`

HasPartnerId returns a boolean if a field has been set.

### GetViralLoopApiToken

`func (o *ViralLoopWaitlists) GetViralLoopApiToken() string`

GetViralLoopApiToken returns the ViralLoopApiToken field if non-nil, zero value otherwise.

### GetViralLoopApiTokenOk

`func (o *ViralLoopWaitlists) GetViralLoopApiTokenOk() (*string, bool)`

GetViralLoopApiTokenOk returns a tuple with the ViralLoopApiToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViralLoopApiToken

`func (o *ViralLoopWaitlists) SetViralLoopApiToken(v string)`

SetViralLoopApiToken sets ViralLoopApiToken field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


