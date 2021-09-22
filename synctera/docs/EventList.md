# EventList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EventList** | [**[]Event**](Event.md) | Array of events | 
**NextPageToken** | Pointer to **string** | If returned, use the next_page_token to query for the next page of results. Not returned if there are no more rows. | [optional] 

## Methods

### NewEventList

`func NewEventList(eventList []Event, ) *EventList`

NewEventList instantiates a new EventList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEventListWithDefaults

`func NewEventListWithDefaults() *EventList`

NewEventListWithDefaults instantiates a new EventList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEventList

`func (o *EventList) GetEventList() []Event`

GetEventList returns the EventList field if non-nil, zero value otherwise.

### GetEventListOk

`func (o *EventList) GetEventListOk() (*[]Event, bool)`

GetEventListOk returns a tuple with the EventList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventList

`func (o *EventList) SetEventList(v []Event)`

SetEventList sets EventList field to given value.


### GetNextPageToken

`func (o *EventList) GetNextPageToken() string`

GetNextPageToken returns the NextPageToken field if non-nil, zero value otherwise.

### GetNextPageTokenOk

`func (o *EventList) GetNextPageTokenOk() (*string, bool)`

GetNextPageTokenOk returns a tuple with the NextPageToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextPageToken

`func (o *EventList) SetNextPageToken(v string)`

SetNextPageToken sets NextPageToken field to given value.

### HasNextPageToken

`func (o *EventList) HasNextPageToken() bool`

HasNextPageToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


