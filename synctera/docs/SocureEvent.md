# SocureEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GlobalWatchlist** | Pointer to [**SocureGlobalWatchlist**](SocureGlobalWatchlist.md) |  | [optional] 
**ReferenceId** | **string** | A 36 character reference ID included with every ID+ response. | 

## Methods

### NewSocureEvent

`func NewSocureEvent(referenceId string, ) *SocureEvent`

NewSocureEvent instantiates a new SocureEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSocureEventWithDefaults

`func NewSocureEventWithDefaults() *SocureEvent`

NewSocureEventWithDefaults instantiates a new SocureEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGlobalWatchlist

`func (o *SocureEvent) GetGlobalWatchlist() SocureGlobalWatchlist`

GetGlobalWatchlist returns the GlobalWatchlist field if non-nil, zero value otherwise.

### GetGlobalWatchlistOk

`func (o *SocureEvent) GetGlobalWatchlistOk() (*SocureGlobalWatchlist, bool)`

GetGlobalWatchlistOk returns a tuple with the GlobalWatchlist field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlobalWatchlist

`func (o *SocureEvent) SetGlobalWatchlist(v SocureGlobalWatchlist)`

SetGlobalWatchlist sets GlobalWatchlist field to given value.

### HasGlobalWatchlist

`func (o *SocureEvent) HasGlobalWatchlist() bool`

HasGlobalWatchlist returns a boolean if a field has been set.

### GetReferenceId

`func (o *SocureEvent) GetReferenceId() string`

GetReferenceId returns the ReferenceId field if non-nil, zero value otherwise.

### GetReferenceIdOk

`func (o *SocureEvent) GetReferenceIdOk() (*string, bool)`

GetReferenceIdOk returns a tuple with the ReferenceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceId

`func (o *SocureEvent) SetReferenceId(v string)`

SetReferenceId sets ReferenceId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


