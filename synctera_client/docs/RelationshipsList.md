# RelationshipsList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Relationships** | [**[]RelationshipIn**](RelationshipIn.md) | Array of business/person relationships. | 
**NextPageToken** | Pointer to **string** | If returned, use the next_page_token to query for the next page of results. Not returned if there are no more rows. | [optional] 

## Methods

### NewRelationshipsList

`func NewRelationshipsList(relationships []RelationshipIn, ) *RelationshipsList`

NewRelationshipsList instantiates a new RelationshipsList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRelationshipsListWithDefaults

`func NewRelationshipsListWithDefaults() *RelationshipsList`

NewRelationshipsListWithDefaults instantiates a new RelationshipsList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRelationships

`func (o *RelationshipsList) GetRelationships() []RelationshipIn`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *RelationshipsList) GetRelationshipsOk() (*[]RelationshipIn, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *RelationshipsList) SetRelationships(v []RelationshipIn)`

SetRelationships sets Relationships field to given value.


### GetNextPageToken

`func (o *RelationshipsList) GetNextPageToken() string`

GetNextPageToken returns the NextPageToken field if non-nil, zero value otherwise.

### GetNextPageTokenOk

`func (o *RelationshipsList) GetNextPageTokenOk() (*string, bool)`

GetNextPageTokenOk returns a tuple with the NextPageToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextPageToken

`func (o *RelationshipsList) SetNextPageToken(v string)`

SetNextPageToken sets NextPageToken field to given value.

### HasNextPageToken

`func (o *RelationshipsList) HasNextPageToken() bool`

HasNextPageToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


