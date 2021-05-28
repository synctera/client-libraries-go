# Relationship1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | ID of related entity | 
**RelationshipRole** | [**RelationshipRole**](RelationshipRole.md) |  | 

## Methods

### NewRelationship1

`func NewRelationship1(id string, relationshipRole RelationshipRole, ) *Relationship1`

NewRelationship1 instantiates a new Relationship1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRelationship1WithDefaults

`func NewRelationship1WithDefaults() *Relationship1`

NewRelationship1WithDefaults instantiates a new Relationship1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Relationship1) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Relationship1) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Relationship1) SetId(v string)`

SetId sets Id field to given value.


### GetRelationshipRole

`func (o *Relationship1) GetRelationshipRole() RelationshipRole`

GetRelationshipRole returns the RelationshipRole field if non-nil, zero value otherwise.

### GetRelationshipRoleOk

`func (o *Relationship1) GetRelationshipRoleOk() (*RelationshipRole, bool)`

GetRelationshipRoleOk returns a tuple with the RelationshipRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationshipRole

`func (o *Relationship1) SetRelationshipRole(v RelationshipRole)`

SetRelationshipRole sets RelationshipRole field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


