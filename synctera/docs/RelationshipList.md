# RelationshipList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Relationships** | [**[]Relationship**](Relationship.md) | Array of relationships | 
**NextPageToken** | Pointer to **string** | If returned, use the next_page_token to query for the next page of results. Not returned if there are no more rows. | [optional] 

## Methods

### NewRelationshipList

`func NewRelationshipList(relationships []Relationship, ) *RelationshipList`

NewRelationshipList instantiates a new RelationshipList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRelationshipListWithDefaults

`func NewRelationshipListWithDefaults() *RelationshipList`

NewRelationshipListWithDefaults instantiates a new RelationshipList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRelationships

`func (o *RelationshipList) GetRelationships() []Relationship`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *RelationshipList) GetRelationshipsOk() (*[]Relationship, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *RelationshipList) SetRelationships(v []Relationship)`

SetRelationships sets Relationships field to given value.


### GetNextPageToken

`func (o *RelationshipList) GetNextPageToken() string`

GetNextPageToken returns the NextPageToken field if non-nil, zero value otherwise.

### GetNextPageTokenOk

`func (o *RelationshipList) GetNextPageTokenOk() (*string, bool)`

GetNextPageTokenOk returns a tuple with the NextPageToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextPageToken

`func (o *RelationshipList) SetNextPageToken(v string)`

SetNextPageToken sets NextPageToken field to given value.

### HasNextPageToken

`func (o *RelationshipList) HasNextPageToken() bool`

HasNextPageToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


