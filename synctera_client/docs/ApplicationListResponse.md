# ApplicationListResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExternalApplications** | [**[]ApplicationResponse1**](ApplicationResponse1.md) | Array of External Applications | 
**NextPageToken** | Pointer to **string** | If returned, use the next_page_token to query for the next page of results. Not returned if there are no more rows. | [optional] 

## Methods

### NewApplicationListResponse

`func NewApplicationListResponse(externalApplications []ApplicationResponse1, ) *ApplicationListResponse`

NewApplicationListResponse instantiates a new ApplicationListResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplicationListResponseWithDefaults

`func NewApplicationListResponseWithDefaults() *ApplicationListResponse`

NewApplicationListResponseWithDefaults instantiates a new ApplicationListResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExternalApplications

`func (o *ApplicationListResponse) GetExternalApplications() []ApplicationResponse1`

GetExternalApplications returns the ExternalApplications field if non-nil, zero value otherwise.

### GetExternalApplicationsOk

`func (o *ApplicationListResponse) GetExternalApplicationsOk() (*[]ApplicationResponse1, bool)`

GetExternalApplicationsOk returns a tuple with the ExternalApplications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalApplications

`func (o *ApplicationListResponse) SetExternalApplications(v []ApplicationResponse1)`

SetExternalApplications sets ExternalApplications field to given value.


### GetNextPageToken

`func (o *ApplicationListResponse) GetNextPageToken() string`

GetNextPageToken returns the NextPageToken field if non-nil, zero value otherwise.

### GetNextPageTokenOk

`func (o *ApplicationListResponse) GetNextPageTokenOk() (*string, bool)`

GetNextPageTokenOk returns a tuple with the NextPageToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextPageToken

`func (o *ApplicationListResponse) SetNextPageToken(v string)`

SetNextPageToken sets NextPageToken field to given value.

### HasNextPageToken

`func (o *ApplicationListResponse) HasNextPageToken() bool`

HasNextPageToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


