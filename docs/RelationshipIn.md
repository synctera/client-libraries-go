# RelationshipIn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdditionalData** | [**AdditionalOwnerData**](AdditionalOwnerData.md) |  | 
**CreationTime** | Pointer to **time.Time** | The date and time the resource was created. | [optional] [readonly] 
**FromPersonId** | **string** | Unique ID for the subject person.  | 
**Id** | Pointer to **string** | Relationship unique identifier. | [optional] [readonly] 
**LastUpdatedTime** | Pointer to **time.Time** | The date and time the resource was last updated. | [optional] [readonly] 
**Metadata** | Pointer to **map[string]interface{}** | Optional field to store additional information about the resource. Intended to be used by the integrator to store non-sensitive data.  | [optional] 
**RelationshipType** | **string** | The relationship type. One of the following: * &#x60;BENEFICIAL_OWNER_OF&#x60; – a person who directly or indirectly owns a portion of the business. * &#x60;MANAGING_PERSON_OF&#x60; – a person who is an officer, director, or other notable person of an organization. * &#x60;OWNER_OF&#x60; – a business with ownership of another business.  | 
**ToBusinessId** | **string** | Unique ID for the related business.  | 
**FromBusinessId** | **string** | Unique ID for the subject business.  | 

## Methods

### NewRelationshipIn

`func NewRelationshipIn(additionalData AdditionalOwnerData, fromPersonId string, relationshipType string, toBusinessId string, fromBusinessId string, ) *RelationshipIn`

NewRelationshipIn instantiates a new RelationshipIn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRelationshipInWithDefaults

`func NewRelationshipInWithDefaults() *RelationshipIn`

NewRelationshipInWithDefaults instantiates a new RelationshipIn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdditionalData

`func (o *RelationshipIn) GetAdditionalData() AdditionalOwnerData`

GetAdditionalData returns the AdditionalData field if non-nil, zero value otherwise.

### GetAdditionalDataOk

`func (o *RelationshipIn) GetAdditionalDataOk() (*AdditionalOwnerData, bool)`

GetAdditionalDataOk returns a tuple with the AdditionalData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalData

`func (o *RelationshipIn) SetAdditionalData(v AdditionalOwnerData)`

SetAdditionalData sets AdditionalData field to given value.


### GetCreationTime

`func (o *RelationshipIn) GetCreationTime() time.Time`

GetCreationTime returns the CreationTime field if non-nil, zero value otherwise.

### GetCreationTimeOk

`func (o *RelationshipIn) GetCreationTimeOk() (*time.Time, bool)`

GetCreationTimeOk returns a tuple with the CreationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationTime

`func (o *RelationshipIn) SetCreationTime(v time.Time)`

SetCreationTime sets CreationTime field to given value.

### HasCreationTime

`func (o *RelationshipIn) HasCreationTime() bool`

HasCreationTime returns a boolean if a field has been set.

### GetFromPersonId

`func (o *RelationshipIn) GetFromPersonId() string`

GetFromPersonId returns the FromPersonId field if non-nil, zero value otherwise.

### GetFromPersonIdOk

`func (o *RelationshipIn) GetFromPersonIdOk() (*string, bool)`

GetFromPersonIdOk returns a tuple with the FromPersonId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromPersonId

`func (o *RelationshipIn) SetFromPersonId(v string)`

SetFromPersonId sets FromPersonId field to given value.


### GetId

`func (o *RelationshipIn) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RelationshipIn) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RelationshipIn) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *RelationshipIn) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLastUpdatedTime

`func (o *RelationshipIn) GetLastUpdatedTime() time.Time`

GetLastUpdatedTime returns the LastUpdatedTime field if non-nil, zero value otherwise.

### GetLastUpdatedTimeOk

`func (o *RelationshipIn) GetLastUpdatedTimeOk() (*time.Time, bool)`

GetLastUpdatedTimeOk returns a tuple with the LastUpdatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdatedTime

`func (o *RelationshipIn) SetLastUpdatedTime(v time.Time)`

SetLastUpdatedTime sets LastUpdatedTime field to given value.

### HasLastUpdatedTime

`func (o *RelationshipIn) HasLastUpdatedTime() bool`

HasLastUpdatedTime returns a boolean if a field has been set.

### GetMetadata

`func (o *RelationshipIn) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *RelationshipIn) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *RelationshipIn) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *RelationshipIn) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetRelationshipType

`func (o *RelationshipIn) GetRelationshipType() string`

GetRelationshipType returns the RelationshipType field if non-nil, zero value otherwise.

### GetRelationshipTypeOk

`func (o *RelationshipIn) GetRelationshipTypeOk() (*string, bool)`

GetRelationshipTypeOk returns a tuple with the RelationshipType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationshipType

`func (o *RelationshipIn) SetRelationshipType(v string)`

SetRelationshipType sets RelationshipType field to given value.


### GetToBusinessId

`func (o *RelationshipIn) GetToBusinessId() string`

GetToBusinessId returns the ToBusinessId field if non-nil, zero value otherwise.

### GetToBusinessIdOk

`func (o *RelationshipIn) GetToBusinessIdOk() (*string, bool)`

GetToBusinessIdOk returns a tuple with the ToBusinessId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToBusinessId

`func (o *RelationshipIn) SetToBusinessId(v string)`

SetToBusinessId sets ToBusinessId field to given value.


### GetFromBusinessId

`func (o *RelationshipIn) GetFromBusinessId() string`

GetFromBusinessId returns the FromBusinessId field if non-nil, zero value otherwise.

### GetFromBusinessIdOk

`func (o *RelationshipIn) GetFromBusinessIdOk() (*string, bool)`

GetFromBusinessIdOk returns a tuple with the FromBusinessId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromBusinessId

`func (o *RelationshipIn) SetFromBusinessId(v string)`

SetFromBusinessId sets FromBusinessId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


