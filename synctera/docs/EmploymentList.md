# EmploymentList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Employment** | [**[]Employment**](Employment.md) | Array of customer employment records. | 
**NextPageToken** | Pointer to **string** | If returned, use the next_page_token to query for the next page of results. Not returned if there are no more rows. | [optional] 

## Methods

### NewEmploymentList

`func NewEmploymentList(employment []Employment, ) *EmploymentList`

NewEmploymentList instantiates a new EmploymentList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEmploymentListWithDefaults

`func NewEmploymentListWithDefaults() *EmploymentList`

NewEmploymentListWithDefaults instantiates a new EmploymentList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmployment

`func (o *EmploymentList) GetEmployment() []Employment`

GetEmployment returns the Employment field if non-nil, zero value otherwise.

### GetEmploymentOk

`func (o *EmploymentList) GetEmploymentOk() (*[]Employment, bool)`

GetEmploymentOk returns a tuple with the Employment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmployment

`func (o *EmploymentList) SetEmployment(v []Employment)`

SetEmployment sets Employment field to given value.


### GetNextPageToken

`func (o *EmploymentList) GetNextPageToken() string`

GetNextPageToken returns the NextPageToken field if non-nil, zero value otherwise.

### GetNextPageTokenOk

`func (o *EmploymentList) GetNextPageTokenOk() (*string, bool)`

GetNextPageTokenOk returns a tuple with the NextPageToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextPageToken

`func (o *EmploymentList) SetNextPageToken(v string)`

SetNextPageToken sets NextPageToken field to given value.

### HasNextPageToken

`func (o *EmploymentList) HasNextPageToken() bool`

HasNextPageToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


