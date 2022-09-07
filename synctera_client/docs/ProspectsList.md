# ProspectsList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Prospects** | [**[]Prospect1**](Prospect1.md) | Array of Prospects | 
**NextPageToken** | Pointer to **string** | If returned, use the next_page_token to query for the next page of results. Not returned if there are no more rows. | [optional] 

## Methods

### NewProspectsList

`func NewProspectsList(prospects []Prospect1, ) *ProspectsList`

NewProspectsList instantiates a new ProspectsList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProspectsListWithDefaults

`func NewProspectsListWithDefaults() *ProspectsList`

NewProspectsListWithDefaults instantiates a new ProspectsList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProspects

`func (o *ProspectsList) GetProspects() []Prospect1`

GetProspects returns the Prospects field if non-nil, zero value otherwise.

### GetProspectsOk

`func (o *ProspectsList) GetProspectsOk() (*[]Prospect1, bool)`

GetProspectsOk returns a tuple with the Prospects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProspects

`func (o *ProspectsList) SetProspects(v []Prospect1)`

SetProspects sets Prospects field to given value.


### GetNextPageToken

`func (o *ProspectsList) GetNextPageToken() string`

GetNextPageToken returns the NextPageToken field if non-nil, zero value otherwise.

### GetNextPageTokenOk

`func (o *ProspectsList) GetNextPageTokenOk() (*string, bool)`

GetNextPageTokenOk returns a tuple with the NextPageToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextPageToken

`func (o *ProspectsList) SetNextPageToken(v string)`

SetNextPageToken sets NextPageToken field to given value.

### HasNextPageToken

`func (o *ProspectsList) HasNextPageToken() bool`

HasNextPageToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


