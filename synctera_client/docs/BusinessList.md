# BusinessList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Businesses** | [**[]Business**](Business.md) | Array of businesses. | 
**NextPageToken** | Pointer to **string** | If returned, use the next_page_token to query for the next page of results. Not returned if there are no more rows. | [optional] 

## Methods

### NewBusinessList

`func NewBusinessList(businesses []Business, ) *BusinessList`

NewBusinessList instantiates a new BusinessList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBusinessListWithDefaults

`func NewBusinessListWithDefaults() *BusinessList`

NewBusinessListWithDefaults instantiates a new BusinessList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBusinesses

`func (o *BusinessList) GetBusinesses() []Business`

GetBusinesses returns the Businesses field if non-nil, zero value otherwise.

### GetBusinessesOk

`func (o *BusinessList) GetBusinessesOk() (*[]Business, bool)`

GetBusinessesOk returns a tuple with the Businesses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinesses

`func (o *BusinessList) SetBusinesses(v []Business)`

SetBusinesses sets Businesses field to given value.


### GetNextPageToken

`func (o *BusinessList) GetNextPageToken() string`

GetNextPageToken returns the NextPageToken field if non-nil, zero value otherwise.

### GetNextPageTokenOk

`func (o *BusinessList) GetNextPageTokenOk() (*string, bool)`

GetNextPageTokenOk returns a tuple with the NextPageToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextPageToken

`func (o *BusinessList) SetNextPageToken(v string)`

SetNextPageToken sets NextPageToken field to given value.

### HasNextPageToken

`func (o *BusinessList) HasNextPageToken() bool`

HasNextPageToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


