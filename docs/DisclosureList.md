# DisclosureList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Disclosures** | [**[]Disclosure**](Disclosure.md) | Array of disclosures. | 
**NextPageToken** | Pointer to **string** | If returned, use the next_page_token to query for the next page of results. Not returned if there are no more rows. | [optional] 

## Methods

### NewDisclosureList

`func NewDisclosureList(disclosures []Disclosure, ) *DisclosureList`

NewDisclosureList instantiates a new DisclosureList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDisclosureListWithDefaults

`func NewDisclosureListWithDefaults() *DisclosureList`

NewDisclosureListWithDefaults instantiates a new DisclosureList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDisclosures

`func (o *DisclosureList) GetDisclosures() []Disclosure`

GetDisclosures returns the Disclosures field if non-nil, zero value otherwise.

### GetDisclosuresOk

`func (o *DisclosureList) GetDisclosuresOk() (*[]Disclosure, bool)`

GetDisclosuresOk returns a tuple with the Disclosures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisclosures

`func (o *DisclosureList) SetDisclosures(v []Disclosure)`

SetDisclosures sets Disclosures field to given value.


### GetNextPageToken

`func (o *DisclosureList) GetNextPageToken() string`

GetNextPageToken returns the NextPageToken field if non-nil, zero value otherwise.

### GetNextPageTokenOk

`func (o *DisclosureList) GetNextPageTokenOk() (*string, bool)`

GetNextPageTokenOk returns a tuple with the NextPageToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextPageToken

`func (o *DisclosureList) SetNextPageToken(v string)`

SetNextPageToken sets NextPageToken field to given value.

### HasNextPageToken

`func (o *DisclosureList) HasNextPageToken() bool`

HasNextPageToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


