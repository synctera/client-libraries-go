# WaitlistsList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Waitlists** | [**[]Waitlist**](Waitlist.md) | Array of Waitlists | 
**NextPageToken** | Pointer to **string** | If returned, use the next_page_token to query for the next page of results. Not returned if there are no more rows. | [optional] 

## Methods

### NewWaitlistsList

`func NewWaitlistsList(waitlists []Waitlist, ) *WaitlistsList`

NewWaitlistsList instantiates a new WaitlistsList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWaitlistsListWithDefaults

`func NewWaitlistsListWithDefaults() *WaitlistsList`

NewWaitlistsListWithDefaults instantiates a new WaitlistsList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWaitlists

`func (o *WaitlistsList) GetWaitlists() []Waitlist`

GetWaitlists returns the Waitlists field if non-nil, zero value otherwise.

### GetWaitlistsOk

`func (o *WaitlistsList) GetWaitlistsOk() (*[]Waitlist, bool)`

GetWaitlistsOk returns a tuple with the Waitlists field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWaitlists

`func (o *WaitlistsList) SetWaitlists(v []Waitlist)`

SetWaitlists sets Waitlists field to given value.


### GetNextPageToken

`func (o *WaitlistsList) GetNextPageToken() string`

GetNextPageToken returns the NextPageToken field if non-nil, zero value otherwise.

### GetNextPageTokenOk

`func (o *WaitlistsList) GetNextPageTokenOk() (*string, bool)`

GetNextPageTokenOk returns a tuple with the NextPageToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextPageToken

`func (o *WaitlistsList) SetNextPageToken(v string)`

SetNextPageToken sets NextPageToken field to given value.

### HasNextPageToken

`func (o *WaitlistsList) HasNextPageToken() bool`

HasNextPageToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


