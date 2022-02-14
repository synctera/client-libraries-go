# Relationship

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CustomerId** | **string** | Customer that the current account is associated with | 
**Id** | Pointer to **string** | ID of account relationship | [optional] [readonly] 
**RelationshipType** | [**AccountRelationshipType**](AccountRelationshipType.md) |  | 

## Methods

### NewRelationship

`func NewRelationship(customerId string, relationshipType AccountRelationshipType, ) *Relationship`

NewRelationship instantiates a new Relationship object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRelationshipWithDefaults

`func NewRelationshipWithDefaults() *Relationship`

NewRelationshipWithDefaults instantiates a new Relationship object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCustomerId

`func (o *Relationship) GetCustomerId() string`

GetCustomerId returns the CustomerId field if non-nil, zero value otherwise.

### GetCustomerIdOk

`func (o *Relationship) GetCustomerIdOk() (*string, bool)`

GetCustomerIdOk returns a tuple with the CustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerId

`func (o *Relationship) SetCustomerId(v string)`

SetCustomerId sets CustomerId field to given value.


### GetId

`func (o *Relationship) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Relationship) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Relationship) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Relationship) HasId() bool`

HasId returns a boolean if a field has been set.

### GetRelationshipType

`func (o *Relationship) GetRelationshipType() AccountRelationshipType`

GetRelationshipType returns the RelationshipType field if non-nil, zero value otherwise.

### GetRelationshipTypeOk

`func (o *Relationship) GetRelationshipTypeOk() (*AccountRelationshipType, bool)`

GetRelationshipTypeOk returns a tuple with the RelationshipType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationshipType

`func (o *Relationship) SetRelationshipType(v AccountRelationshipType)`

SetRelationshipType sets RelationshipType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

