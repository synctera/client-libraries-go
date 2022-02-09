# SocureMatch

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EntityId** | **string** |  | 
**Status** | **string** |  | 
**MatchFields** | Pointer to **[]string** |  | [optional] 
**SourceUrls** | Pointer to **[]string** |  | [optional] 
**Comments** | Pointer to [**SocureMatchComments**](SocureMatchComments.md) |  | [optional] 

## Methods

### NewSocureMatch

`func NewSocureMatch(entityId string, status string, ) *SocureMatch`

NewSocureMatch instantiates a new SocureMatch object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSocureMatchWithDefaults

`func NewSocureMatchWithDefaults() *SocureMatch`

NewSocureMatchWithDefaults instantiates a new SocureMatch object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEntityId

`func (o *SocureMatch) GetEntityId() string`

GetEntityId returns the EntityId field if non-nil, zero value otherwise.

### GetEntityIdOk

`func (o *SocureMatch) GetEntityIdOk() (*string, bool)`

GetEntityIdOk returns a tuple with the EntityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityId

`func (o *SocureMatch) SetEntityId(v string)`

SetEntityId sets EntityId field to given value.


### GetStatus

`func (o *SocureMatch) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *SocureMatch) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *SocureMatch) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetMatchFields

`func (o *SocureMatch) GetMatchFields() []string`

GetMatchFields returns the MatchFields field if non-nil, zero value otherwise.

### GetMatchFieldsOk

`func (o *SocureMatch) GetMatchFieldsOk() (*[]string, bool)`

GetMatchFieldsOk returns a tuple with the MatchFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatchFields

`func (o *SocureMatch) SetMatchFields(v []string)`

SetMatchFields sets MatchFields field to given value.

### HasMatchFields

`func (o *SocureMatch) HasMatchFields() bool`

HasMatchFields returns a boolean if a field has been set.

### GetSourceUrls

`func (o *SocureMatch) GetSourceUrls() []string`

GetSourceUrls returns the SourceUrls field if non-nil, zero value otherwise.

### GetSourceUrlsOk

`func (o *SocureMatch) GetSourceUrlsOk() (*[]string, bool)`

GetSourceUrlsOk returns a tuple with the SourceUrls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceUrls

`func (o *SocureMatch) SetSourceUrls(v []string)`

SetSourceUrls sets SourceUrls field to given value.

### HasSourceUrls

`func (o *SocureMatch) HasSourceUrls() bool`

HasSourceUrls returns a boolean if a field has been set.

### GetComments

`func (o *SocureMatch) GetComments() SocureMatchComments`

GetComments returns the Comments field if non-nil, zero value otherwise.

### GetCommentsOk

`func (o *SocureMatch) GetCommentsOk() (*SocureMatchComments, bool)`

GetCommentsOk returns a tuple with the Comments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComments

`func (o *SocureMatch) SetComments(v SocureMatchComments)`

SetComments sets Comments field to given value.

### HasComments

`func (o *SocureMatch) HasComments() bool`

HasComments returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


