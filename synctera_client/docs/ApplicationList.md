# ApplicationList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Applications** | [**[]ApplicationResponse**](ApplicationResponse.md) | Array of credit applications. | 
**NextPageToken** | Pointer to **string** | If returned, use the next_page_token to query for the next page of results. Not returned if there are no more rows. | [optional] 

## Methods

### NewApplicationList

`func NewApplicationList(applications []ApplicationResponse, ) *ApplicationList`

NewApplicationList instantiates a new ApplicationList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplicationListWithDefaults

`func NewApplicationListWithDefaults() *ApplicationList`

NewApplicationListWithDefaults instantiates a new ApplicationList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApplications

`func (o *ApplicationList) GetApplications() []ApplicationResponse`

GetApplications returns the Applications field if non-nil, zero value otherwise.

### GetApplicationsOk

`func (o *ApplicationList) GetApplicationsOk() (*[]ApplicationResponse, bool)`

GetApplicationsOk returns a tuple with the Applications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplications

`func (o *ApplicationList) SetApplications(v []ApplicationResponse)`

SetApplications sets Applications field to given value.


### GetNextPageToken

`func (o *ApplicationList) GetNextPageToken() string`

GetNextPageToken returns the NextPageToken field if non-nil, zero value otherwise.

### GetNextPageTokenOk

`func (o *ApplicationList) GetNextPageTokenOk() (*string, bool)`

GetNextPageTokenOk returns a tuple with the NextPageToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextPageToken

`func (o *ApplicationList) SetNextPageToken(v string)`

SetNextPageToken sets NextPageToken field to given value.

### HasNextPageToken

`func (o *ApplicationList) HasNextPageToken() bool`

HasNextPageToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


