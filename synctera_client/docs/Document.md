# Document

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreationTime** | Pointer to **time.Time** | The date and time the resource was created. | [optional] [readonly] 
**Description** | Pointer to **string** | A description of the attached document. | [optional] 
**Encryption** | Pointer to [**Encryption**](Encryption.md) |  | [optional] [default to ENCRYPTION_NOT_REQUIRED]
**FileName** | Pointer to **string** | The name of the document. | [optional] [readonly] 
**Id** | Pointer to **string** | The unique identifier for this resource. | [optional] [readonly] 
**LastUpdatedTime** | Pointer to **time.Time** | The date and time the resource was last update. | [optional] [readonly] 
**Metadata** | Pointer to **map[string]interface{}** | Optional field to store additional informaton about the resource.  Intended to be used by the integrator to store non-sensitive data.  | [optional] 
**Name** | Pointer to **string** | A user-friendly name for the document. | [optional] 
**RelatedResourceId** | Pointer to **string** | Related resource ID | [optional] 
**RelatedResourceType** | Pointer to [**RelatedResourceType**](RelatedResourceType.md) |  | [optional] 

## Methods

### NewDocument

`func NewDocument() *Document`

NewDocument instantiates a new Document object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDocumentWithDefaults

`func NewDocumentWithDefaults() *Document`

NewDocumentWithDefaults instantiates a new Document object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreationTime

`func (o *Document) GetCreationTime() time.Time`

GetCreationTime returns the CreationTime field if non-nil, zero value otherwise.

### GetCreationTimeOk

`func (o *Document) GetCreationTimeOk() (*time.Time, bool)`

GetCreationTimeOk returns a tuple with the CreationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationTime

`func (o *Document) SetCreationTime(v time.Time)`

SetCreationTime sets CreationTime field to given value.

### HasCreationTime

`func (o *Document) HasCreationTime() bool`

HasCreationTime returns a boolean if a field has been set.

### GetDescription

`func (o *Document) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Document) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Document) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Document) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEncryption

`func (o *Document) GetEncryption() Encryption`

GetEncryption returns the Encryption field if non-nil, zero value otherwise.

### GetEncryptionOk

`func (o *Document) GetEncryptionOk() (*Encryption, bool)`

GetEncryptionOk returns a tuple with the Encryption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryption

`func (o *Document) SetEncryption(v Encryption)`

SetEncryption sets Encryption field to given value.

### HasEncryption

`func (o *Document) HasEncryption() bool`

HasEncryption returns a boolean if a field has been set.

### GetFileName

`func (o *Document) GetFileName() string`

GetFileName returns the FileName field if non-nil, zero value otherwise.

### GetFileNameOk

`func (o *Document) GetFileNameOk() (*string, bool)`

GetFileNameOk returns a tuple with the FileName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileName

`func (o *Document) SetFileName(v string)`

SetFileName sets FileName field to given value.

### HasFileName

`func (o *Document) HasFileName() bool`

HasFileName returns a boolean if a field has been set.

### GetId

`func (o *Document) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Document) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Document) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Document) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLastUpdatedTime

`func (o *Document) GetLastUpdatedTime() time.Time`

GetLastUpdatedTime returns the LastUpdatedTime field if non-nil, zero value otherwise.

### GetLastUpdatedTimeOk

`func (o *Document) GetLastUpdatedTimeOk() (*time.Time, bool)`

GetLastUpdatedTimeOk returns a tuple with the LastUpdatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdatedTime

`func (o *Document) SetLastUpdatedTime(v time.Time)`

SetLastUpdatedTime sets LastUpdatedTime field to given value.

### HasLastUpdatedTime

`func (o *Document) HasLastUpdatedTime() bool`

HasLastUpdatedTime returns a boolean if a field has been set.

### GetMetadata

`func (o *Document) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *Document) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *Document) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *Document) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetName

`func (o *Document) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Document) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Document) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Document) HasName() bool`

HasName returns a boolean if a field has been set.

### GetRelatedResourceId

`func (o *Document) GetRelatedResourceId() string`

GetRelatedResourceId returns the RelatedResourceId field if non-nil, zero value otherwise.

### GetRelatedResourceIdOk

`func (o *Document) GetRelatedResourceIdOk() (*string, bool)`

GetRelatedResourceIdOk returns a tuple with the RelatedResourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelatedResourceId

`func (o *Document) SetRelatedResourceId(v string)`

SetRelatedResourceId sets RelatedResourceId field to given value.

### HasRelatedResourceId

`func (o *Document) HasRelatedResourceId() bool`

HasRelatedResourceId returns a boolean if a field has been set.

### GetRelatedResourceType

`func (o *Document) GetRelatedResourceType() RelatedResourceType`

GetRelatedResourceType returns the RelatedResourceType field if non-nil, zero value otherwise.

### GetRelatedResourceTypeOk

`func (o *Document) GetRelatedResourceTypeOk() (*RelatedResourceType, bool)`

GetRelatedResourceTypeOk returns a tuple with the RelatedResourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelatedResourceType

`func (o *Document) SetRelatedResourceType(v RelatedResourceType)`

SetRelatedResourceType sets RelatedResourceType field to given value.

### HasRelatedResourceType

`func (o *Document) HasRelatedResourceType() bool`

HasRelatedResourceType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


