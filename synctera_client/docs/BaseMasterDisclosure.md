# BaseMasterDisclosure

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreationTime** | Pointer to **time.Time** | The date and time the resource was created. | [optional] [readonly] 
**Id** | Pointer to **string** | The unique identifier for this resource. | [optional] [readonly] 
**LastUpdatedTime** | Pointer to **time.Time** | The date and time the resource was last updated. | [optional] [readonly] 
**Metadata** | Pointer to **map[string]interface{}** | Optional field to store additional information about the resource. Intended to be used by the integrator to store non-sensitive data.  | [optional] 
**Type** | Pointer to [**DisclosureType**](DisclosureType.md) |  | [optional] 
**Version** | Pointer to **string** | Version of the disclosure document. | [optional] 

## Methods

### NewBaseMasterDisclosure

`func NewBaseMasterDisclosure() *BaseMasterDisclosure`

NewBaseMasterDisclosure instantiates a new BaseMasterDisclosure object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBaseMasterDisclosureWithDefaults

`func NewBaseMasterDisclosureWithDefaults() *BaseMasterDisclosure`

NewBaseMasterDisclosureWithDefaults instantiates a new BaseMasterDisclosure object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreationTime

`func (o *BaseMasterDisclosure) GetCreationTime() time.Time`

GetCreationTime returns the CreationTime field if non-nil, zero value otherwise.

### GetCreationTimeOk

`func (o *BaseMasterDisclosure) GetCreationTimeOk() (*time.Time, bool)`

GetCreationTimeOk returns a tuple with the CreationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationTime

`func (o *BaseMasterDisclosure) SetCreationTime(v time.Time)`

SetCreationTime sets CreationTime field to given value.

### HasCreationTime

`func (o *BaseMasterDisclosure) HasCreationTime() bool`

HasCreationTime returns a boolean if a field has been set.

### GetId

`func (o *BaseMasterDisclosure) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BaseMasterDisclosure) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BaseMasterDisclosure) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *BaseMasterDisclosure) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLastUpdatedTime

`func (o *BaseMasterDisclosure) GetLastUpdatedTime() time.Time`

GetLastUpdatedTime returns the LastUpdatedTime field if non-nil, zero value otherwise.

### GetLastUpdatedTimeOk

`func (o *BaseMasterDisclosure) GetLastUpdatedTimeOk() (*time.Time, bool)`

GetLastUpdatedTimeOk returns a tuple with the LastUpdatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdatedTime

`func (o *BaseMasterDisclosure) SetLastUpdatedTime(v time.Time)`

SetLastUpdatedTime sets LastUpdatedTime field to given value.

### HasLastUpdatedTime

`func (o *BaseMasterDisclosure) HasLastUpdatedTime() bool`

HasLastUpdatedTime returns a boolean if a field has been set.

### GetMetadata

`func (o *BaseMasterDisclosure) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *BaseMasterDisclosure) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *BaseMasterDisclosure) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *BaseMasterDisclosure) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetType

`func (o *BaseMasterDisclosure) GetType() DisclosureType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *BaseMasterDisclosure) GetTypeOk() (*DisclosureType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *BaseMasterDisclosure) SetType(v DisclosureType)`

SetType sets Type field to given value.

### HasType

`func (o *BaseMasterDisclosure) HasType() bool`

HasType returns a boolean if a field has been set.

### GetVersion

`func (o *BaseMasterDisclosure) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *BaseMasterDisclosure) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *BaseMasterDisclosure) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *BaseMasterDisclosure) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


