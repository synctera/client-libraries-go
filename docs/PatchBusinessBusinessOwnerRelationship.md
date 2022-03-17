# PatchBusinessBusinessOwnerRelationship

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdditionalData** | Pointer to [**AdditionalOwnerData**](AdditionalOwnerData.md) |  | [optional] 
**CreationTime** | Pointer to **time.Time** | The date and time the resource was created. | [optional] [readonly] 
**FromBusinessId** | Pointer to **string** | Unique ID for the subject business.  | [optional] 
**Id** | Pointer to **string** | Relationship unique identifier. | [optional] [readonly] 
**LastUpdatedTime** | Pointer to **time.Time** | The date and time the resource was last updated. | [optional] [readonly] 
**Metadata** | Pointer to **map[string]interface{}** | Optional field to store additional information about the resource. Intended to be used by the integrator to store non-sensitive data.  | [optional] 
**RelationshipType** | **string** | The relationship type. One of the following: * &#x60;BENEFICIAL_OWNER_OF&#x60; – a person who directly or indirectly owns a portion of the business. * &#x60;MANAGING_PERSON_OF&#x60; – a person who is an officer, director, or other notable person of an organization. * &#x60;OWNER_OF&#x60; – a business with ownership of another business.  | 
**ToBusinessId** | Pointer to **string** | Unique ID for the related business.  | [optional] 

## Methods

### NewPatchBusinessBusinessOwnerRelationship

`func NewPatchBusinessBusinessOwnerRelationship(relationshipType string, ) *PatchBusinessBusinessOwnerRelationship`

NewPatchBusinessBusinessOwnerRelationship instantiates a new PatchBusinessBusinessOwnerRelationship object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchBusinessBusinessOwnerRelationshipWithDefaults

`func NewPatchBusinessBusinessOwnerRelationshipWithDefaults() *PatchBusinessBusinessOwnerRelationship`

NewPatchBusinessBusinessOwnerRelationshipWithDefaults instantiates a new PatchBusinessBusinessOwnerRelationship object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdditionalData

`func (o *PatchBusinessBusinessOwnerRelationship) GetAdditionalData() AdditionalOwnerData`

GetAdditionalData returns the AdditionalData field if non-nil, zero value otherwise.

### GetAdditionalDataOk

`func (o *PatchBusinessBusinessOwnerRelationship) GetAdditionalDataOk() (*AdditionalOwnerData, bool)`

GetAdditionalDataOk returns a tuple with the AdditionalData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalData

`func (o *PatchBusinessBusinessOwnerRelationship) SetAdditionalData(v AdditionalOwnerData)`

SetAdditionalData sets AdditionalData field to given value.

### HasAdditionalData

`func (o *PatchBusinessBusinessOwnerRelationship) HasAdditionalData() bool`

HasAdditionalData returns a boolean if a field has been set.

### GetCreationTime

`func (o *PatchBusinessBusinessOwnerRelationship) GetCreationTime() time.Time`

GetCreationTime returns the CreationTime field if non-nil, zero value otherwise.

### GetCreationTimeOk

`func (o *PatchBusinessBusinessOwnerRelationship) GetCreationTimeOk() (*time.Time, bool)`

GetCreationTimeOk returns a tuple with the CreationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationTime

`func (o *PatchBusinessBusinessOwnerRelationship) SetCreationTime(v time.Time)`

SetCreationTime sets CreationTime field to given value.

### HasCreationTime

`func (o *PatchBusinessBusinessOwnerRelationship) HasCreationTime() bool`

HasCreationTime returns a boolean if a field has been set.

### GetFromBusinessId

`func (o *PatchBusinessBusinessOwnerRelationship) GetFromBusinessId() string`

GetFromBusinessId returns the FromBusinessId field if non-nil, zero value otherwise.

### GetFromBusinessIdOk

`func (o *PatchBusinessBusinessOwnerRelationship) GetFromBusinessIdOk() (*string, bool)`

GetFromBusinessIdOk returns a tuple with the FromBusinessId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromBusinessId

`func (o *PatchBusinessBusinessOwnerRelationship) SetFromBusinessId(v string)`

SetFromBusinessId sets FromBusinessId field to given value.

### HasFromBusinessId

`func (o *PatchBusinessBusinessOwnerRelationship) HasFromBusinessId() bool`

HasFromBusinessId returns a boolean if a field has been set.

### GetId

`func (o *PatchBusinessBusinessOwnerRelationship) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PatchBusinessBusinessOwnerRelationship) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PatchBusinessBusinessOwnerRelationship) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *PatchBusinessBusinessOwnerRelationship) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLastUpdatedTime

`func (o *PatchBusinessBusinessOwnerRelationship) GetLastUpdatedTime() time.Time`

GetLastUpdatedTime returns the LastUpdatedTime field if non-nil, zero value otherwise.

### GetLastUpdatedTimeOk

`func (o *PatchBusinessBusinessOwnerRelationship) GetLastUpdatedTimeOk() (*time.Time, bool)`

GetLastUpdatedTimeOk returns a tuple with the LastUpdatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdatedTime

`func (o *PatchBusinessBusinessOwnerRelationship) SetLastUpdatedTime(v time.Time)`

SetLastUpdatedTime sets LastUpdatedTime field to given value.

### HasLastUpdatedTime

`func (o *PatchBusinessBusinessOwnerRelationship) HasLastUpdatedTime() bool`

HasLastUpdatedTime returns a boolean if a field has been set.

### GetMetadata

`func (o *PatchBusinessBusinessOwnerRelationship) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *PatchBusinessBusinessOwnerRelationship) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *PatchBusinessBusinessOwnerRelationship) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *PatchBusinessBusinessOwnerRelationship) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetRelationshipType

`func (o *PatchBusinessBusinessOwnerRelationship) GetRelationshipType() string`

GetRelationshipType returns the RelationshipType field if non-nil, zero value otherwise.

### GetRelationshipTypeOk

`func (o *PatchBusinessBusinessOwnerRelationship) GetRelationshipTypeOk() (*string, bool)`

GetRelationshipTypeOk returns a tuple with the RelationshipType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationshipType

`func (o *PatchBusinessBusinessOwnerRelationship) SetRelationshipType(v string)`

SetRelationshipType sets RelationshipType field to given value.


### GetToBusinessId

`func (o *PatchBusinessBusinessOwnerRelationship) GetToBusinessId() string`

GetToBusinessId returns the ToBusinessId field if non-nil, zero value otherwise.

### GetToBusinessIdOk

`func (o *PatchBusinessBusinessOwnerRelationship) GetToBusinessIdOk() (*string, bool)`

GetToBusinessIdOk returns a tuple with the ToBusinessId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToBusinessId

`func (o *PatchBusinessBusinessOwnerRelationship) SetToBusinessId(v string)`

SetToBusinessId sets ToBusinessId field to given value.

### HasToBusinessId

`func (o *PatchBusinessBusinessOwnerRelationship) HasToBusinessId() bool`

HasToBusinessId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


