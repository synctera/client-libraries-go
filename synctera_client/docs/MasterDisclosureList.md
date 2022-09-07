# MasterDisclosureList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MasterDisclosures** | [**[]MasterDisclosure**](MasterDisclosure.md) | Array of master disclosures. | 
**NextPageToken** | Pointer to **string** | If returned, use the next_page_token to query for the next page of results. Not returned if there are no more rows. | [optional] 

## Methods

### NewMasterDisclosureList

`func NewMasterDisclosureList(masterDisclosures []MasterDisclosure, ) *MasterDisclosureList`

NewMasterDisclosureList instantiates a new MasterDisclosureList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMasterDisclosureListWithDefaults

`func NewMasterDisclosureListWithDefaults() *MasterDisclosureList`

NewMasterDisclosureListWithDefaults instantiates a new MasterDisclosureList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMasterDisclosures

`func (o *MasterDisclosureList) GetMasterDisclosures() []MasterDisclosure`

GetMasterDisclosures returns the MasterDisclosures field if non-nil, zero value otherwise.

### GetMasterDisclosuresOk

`func (o *MasterDisclosureList) GetMasterDisclosuresOk() (*[]MasterDisclosure, bool)`

GetMasterDisclosuresOk returns a tuple with the MasterDisclosures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMasterDisclosures

`func (o *MasterDisclosureList) SetMasterDisclosures(v []MasterDisclosure)`

SetMasterDisclosures sets MasterDisclosures field to given value.


### GetNextPageToken

`func (o *MasterDisclosureList) GetNextPageToken() string`

GetNextPageToken returns the NextPageToken field if non-nil, zero value otherwise.

### GetNextPageTokenOk

`func (o *MasterDisclosureList) GetNextPageTokenOk() (*string, bool)`

GetNextPageTokenOk returns a tuple with the NextPageToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextPageToken

`func (o *MasterDisclosureList) SetNextPageToken(v string)`

SetNextPageToken sets NextPageToken field to given value.

### HasNextPageToken

`func (o *MasterDisclosureList) HasNextPageToken() bool`

HasNextPageToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


