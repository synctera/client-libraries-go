# AccountCreationAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Relationships** | Pointer to [**[]Relationship**](Relationship.md) | List of the relationship for this account to the parties | [optional] 
**AccountTemplateId** | Pointer to **string** | Account template ID | [optional] 

## Methods

### NewAccountCreationAllOf

`func NewAccountCreationAllOf() *AccountCreationAllOf`

NewAccountCreationAllOf instantiates a new AccountCreationAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountCreationAllOfWithDefaults

`func NewAccountCreationAllOfWithDefaults() *AccountCreationAllOf`

NewAccountCreationAllOfWithDefaults instantiates a new AccountCreationAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRelationships

`func (o *AccountCreationAllOf) GetRelationships() []Relationship`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *AccountCreationAllOf) GetRelationshipsOk() (*[]Relationship, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *AccountCreationAllOf) SetRelationships(v []Relationship)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *AccountCreationAllOf) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.

### GetAccountTemplateId

`func (o *AccountCreationAllOf) GetAccountTemplateId() string`

GetAccountTemplateId returns the AccountTemplateId field if non-nil, zero value otherwise.

### GetAccountTemplateIdOk

`func (o *AccountCreationAllOf) GetAccountTemplateIdOk() (*string, bool)`

GetAccountTemplateIdOk returns a tuple with the AccountTemplateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountTemplateId

`func (o *AccountCreationAllOf) SetAccountTemplateId(v string)`

SetAccountTemplateId sets AccountTemplateId field to given value.

### HasAccountTemplateId

`func (o *AccountCreationAllOf) HasAccountTemplateId() bool`

HasAccountTemplateId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


