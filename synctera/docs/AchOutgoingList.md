# AchOutgoingList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AchOutgoings** | [**[]AchOutgoing**](AchOutgoing.md) | Array of ACH | 
**NextPageToken** | Pointer to **string** | If returned, use the next_page_token to query for the next page of results. | [optional] 

## Methods

### NewAchOutgoingList

`func NewAchOutgoingList(achOutgoings []AchOutgoing, ) *AchOutgoingList`

NewAchOutgoingList instantiates a new AchOutgoingList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAchOutgoingListWithDefaults

`func NewAchOutgoingListWithDefaults() *AchOutgoingList`

NewAchOutgoingListWithDefaults instantiates a new AchOutgoingList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAchOutgoings

`func (o *AchOutgoingList) GetAchOutgoings() []AchOutgoing`

GetAchOutgoings returns the AchOutgoings field if non-nil, zero value otherwise.

### GetAchOutgoingsOk

`func (o *AchOutgoingList) GetAchOutgoingsOk() (*[]AchOutgoing, bool)`

GetAchOutgoingsOk returns a tuple with the AchOutgoings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAchOutgoings

`func (o *AchOutgoingList) SetAchOutgoings(v []AchOutgoing)`

SetAchOutgoings sets AchOutgoings field to given value.


### GetNextPageToken

`func (o *AchOutgoingList) GetNextPageToken() string`

GetNextPageToken returns the NextPageToken field if non-nil, zero value otherwise.

### GetNextPageTokenOk

`func (o *AchOutgoingList) GetNextPageTokenOk() (*string, bool)`

GetNextPageTokenOk returns a tuple with the NextPageToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextPageToken

`func (o *AchOutgoingList) SetNextPageToken(v string)`

SetNextPageToken sets NextPageToken field to given value.

### HasNextPageToken

`func (o *AchOutgoingList) HasNextPageToken() bool`

HasNextPageToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


