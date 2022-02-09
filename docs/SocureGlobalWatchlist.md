# SocureGlobalWatchlist

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ReasonCodes** | Pointer to **[]string** | Array of reason codes | [optional] 
**Matches** | [**map[string][]SocureMatch**](array.md) | Contains key-value pair of the Source list name and an array of details about that match. | 

## Methods

### NewSocureGlobalWatchlist

`func NewSocureGlobalWatchlist(matches map[string][]SocureMatch, ) *SocureGlobalWatchlist`

NewSocureGlobalWatchlist instantiates a new SocureGlobalWatchlist object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSocureGlobalWatchlistWithDefaults

`func NewSocureGlobalWatchlistWithDefaults() *SocureGlobalWatchlist`

NewSocureGlobalWatchlistWithDefaults instantiates a new SocureGlobalWatchlist object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReasonCodes

`func (o *SocureGlobalWatchlist) GetReasonCodes() []string`

GetReasonCodes returns the ReasonCodes field if non-nil, zero value otherwise.

### GetReasonCodesOk

`func (o *SocureGlobalWatchlist) GetReasonCodesOk() (*[]string, bool)`

GetReasonCodesOk returns a tuple with the ReasonCodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReasonCodes

`func (o *SocureGlobalWatchlist) SetReasonCodes(v []string)`

SetReasonCodes sets ReasonCodes field to given value.

### HasReasonCodes

`func (o *SocureGlobalWatchlist) HasReasonCodes() bool`

HasReasonCodes returns a boolean if a field has been set.

### GetMatches

`func (o *SocureGlobalWatchlist) GetMatches() map[string][]SocureMatch`

GetMatches returns the Matches field if non-nil, zero value otherwise.

### GetMatchesOk

`func (o *SocureGlobalWatchlist) GetMatchesOk() (*map[string][]SocureMatch, bool)`

GetMatchesOk returns a tuple with the Matches field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatches

`func (o *SocureGlobalWatchlist) SetMatches(v map[string][]SocureMatch)`

SetMatches sets Matches field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


