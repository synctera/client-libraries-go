# SocureWatchlistResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ReferenceId** | **string** | A 36 character reference ID included with every ID+ response. | 
**GlobalWatchlist** | Pointer to [**SocureGlobalWatchlist**](SocureGlobalWatchlist.md) |  | [optional] 

## Methods

### NewSocureWatchlistResult

`func NewSocureWatchlistResult(referenceId string, ) *SocureWatchlistResult`

NewSocureWatchlistResult instantiates a new SocureWatchlistResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSocureWatchlistResultWithDefaults

`func NewSocureWatchlistResultWithDefaults() *SocureWatchlistResult`

NewSocureWatchlistResultWithDefaults instantiates a new SocureWatchlistResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReferenceId

`func (o *SocureWatchlistResult) GetReferenceId() string`

GetReferenceId returns the ReferenceId field if non-nil, zero value otherwise.

### GetReferenceIdOk

`func (o *SocureWatchlistResult) GetReferenceIdOk() (*string, bool)`

GetReferenceIdOk returns a tuple with the ReferenceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceId

`func (o *SocureWatchlistResult) SetReferenceId(v string)`

SetReferenceId sets ReferenceId field to given value.


### GetGlobalWatchlist

`func (o *SocureWatchlistResult) GetGlobalWatchlist() SocureGlobalWatchlist`

GetGlobalWatchlist returns the GlobalWatchlist field if non-nil, zero value otherwise.

### GetGlobalWatchlistOk

`func (o *SocureWatchlistResult) GetGlobalWatchlistOk() (*SocureGlobalWatchlist, bool)`

GetGlobalWatchlistOk returns a tuple with the GlobalWatchlist field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlobalWatchlist

`func (o *SocureWatchlistResult) SetGlobalWatchlist(v SocureGlobalWatchlist)`

SetGlobalWatchlist sets GlobalWatchlist field to given value.

### HasGlobalWatchlist

`func (o *SocureWatchlistResult) HasGlobalWatchlist() bool`

HasGlobalWatchlist returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


