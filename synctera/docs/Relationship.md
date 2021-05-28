# Relationship

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConnectId** | Pointer to **string** | Connection ID of the account | [optional] 
**RelationshipType** | **string** | Relationship type | 
**CustomerId** | **string** | Customer that the current account is associated with | 

## Methods

### NewRelationship

`func NewRelationship(relationshipType string, customerId string, ) *Relationship`

NewRelationship instantiates a new Relationship object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRelationshipWithDefaults

`func NewRelationshipWithDefaults() *Relationship`

NewRelationshipWithDefaults instantiates a new Relationship object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConnectId

`func (o *Relationship) GetConnectId() string`

GetConnectId returns the ConnectId field if non-nil, zero value otherwise.

### GetConnectIdOk

`func (o *Relationship) GetConnectIdOk() (*string, bool)`

GetConnectIdOk returns a tuple with the ConnectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectId

`func (o *Relationship) SetConnectId(v string)`

SetConnectId sets ConnectId field to given value.

### HasConnectId

`func (o *Relationship) HasConnectId() bool`

HasConnectId returns a boolean if a field has been set.

### GetRelationshipType

`func (o *Relationship) GetRelationshipType() string`

GetRelationshipType returns the RelationshipType field if non-nil, zero value otherwise.

### GetRelationshipTypeOk

`func (o *Relationship) GetRelationshipTypeOk() (*string, bool)`

GetRelationshipTypeOk returns a tuple with the RelationshipType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationshipType

`func (o *Relationship) SetRelationshipType(v string)`

SetRelationshipType sets RelationshipType field to given value.


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



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

