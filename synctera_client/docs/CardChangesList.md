# CardChangesList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Changes** | [**[]CardChange**](CardChange.md) | List of changes in descending chronological order | 
**NextPageToken** | Pointer to **string** | If returned, use the next_page_token to query for the next page of results. Not returned if there are no more rows. | [optional] 

## Methods

### NewCardChangesList

`func NewCardChangesList(changes []CardChange, ) *CardChangesList`

NewCardChangesList instantiates a new CardChangesList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCardChangesListWithDefaults

`func NewCardChangesListWithDefaults() *CardChangesList`

NewCardChangesListWithDefaults instantiates a new CardChangesList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChanges

`func (o *CardChangesList) GetChanges() []CardChange`

GetChanges returns the Changes field if non-nil, zero value otherwise.

### GetChangesOk

`func (o *CardChangesList) GetChangesOk() (*[]CardChange, bool)`

GetChangesOk returns a tuple with the Changes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChanges

`func (o *CardChangesList) SetChanges(v []CardChange)`

SetChanges sets Changes field to given value.


### GetNextPageToken

`func (o *CardChangesList) GetNextPageToken() string`

GetNextPageToken returns the NextPageToken field if non-nil, zero value otherwise.

### GetNextPageTokenOk

`func (o *CardChangesList) GetNextPageTokenOk() (*string, bool)`

GetNextPageTokenOk returns a tuple with the NextPageToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextPageToken

`func (o *CardChangesList) SetNextPageToken(v string)`

SetNextPageToken sets NextPageToken field to given value.

### HasNextPageToken

`func (o *CardChangesList) HasNextPageToken() bool`

HasNextPageToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


