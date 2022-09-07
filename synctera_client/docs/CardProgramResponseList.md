# CardProgramResponseList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Programs** | [**[]CardProgramResponse**](CardProgramResponse.md) | Array of Card Programs | 
**NextPageToken** | Pointer to **string** | If returned, use the next_page_token to query for the next page of results. Not returned if there are no more rows. | [optional] 

## Methods

### NewCardProgramResponseList

`func NewCardProgramResponseList(programs []CardProgramResponse, ) *CardProgramResponseList`

NewCardProgramResponseList instantiates a new CardProgramResponseList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCardProgramResponseListWithDefaults

`func NewCardProgramResponseListWithDefaults() *CardProgramResponseList`

NewCardProgramResponseListWithDefaults instantiates a new CardProgramResponseList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPrograms

`func (o *CardProgramResponseList) GetPrograms() []CardProgramResponse`

GetPrograms returns the Programs field if non-nil, zero value otherwise.

### GetProgramsOk

`func (o *CardProgramResponseList) GetProgramsOk() (*[]CardProgramResponse, bool)`

GetProgramsOk returns a tuple with the Programs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrograms

`func (o *CardProgramResponseList) SetPrograms(v []CardProgramResponse)`

SetPrograms sets Programs field to given value.


### GetNextPageToken

`func (o *CardProgramResponseList) GetNextPageToken() string`

GetNextPageToken returns the NextPageToken field if non-nil, zero value otherwise.

### GetNextPageTokenOk

`func (o *CardProgramResponseList) GetNextPageTokenOk() (*string, bool)`

GetNextPageTokenOk returns a tuple with the NextPageToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextPageToken

`func (o *CardProgramResponseList) SetNextPageToken(v string)`

SetNextPageToken sets NextPageToken field to given value.

### HasNextPageToken

`func (o *CardProgramResponseList) HasNextPageToken() bool`

HasNextPageToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


