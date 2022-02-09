# BaseDisclosure

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The unique identifier for this resource. | [optional] [readonly] 
**PersonId** | Pointer to **string** | Unique ID for the person. Exactly one of &#x60;person_id&#x60; or &#x60;business_id&#x60; must be set.  | [optional] 
**BusinessId** | Pointer to **string** | Unique ID for the business. Exactly one of &#x60;business_id&#x60; or &#x60;person_id&#x60; must be set.  | [optional] 
**Type** | Pointer to [**DisclosureType**](DisclosureType.md) |  | [optional] 
**Version** | Pointer to **string** | Version of the disclosure document. | [optional] 
**EventType** | Pointer to **string** |  | [optional] 
**DisclosureDate** | Pointer to **time.Time** | Date and time the disclosure was made. | [optional] 
**CreationTime** | Pointer to **time.Time** | The date and time the resource was created. | [optional] [readonly] 
**LastUpdatedTime** | Pointer to **time.Time** | The date and time the resource was last updated. | [optional] [readonly] 

## Methods

### NewBaseDisclosure

`func NewBaseDisclosure() *BaseDisclosure`

NewBaseDisclosure instantiates a new BaseDisclosure object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBaseDisclosureWithDefaults

`func NewBaseDisclosureWithDefaults() *BaseDisclosure`

NewBaseDisclosureWithDefaults instantiates a new BaseDisclosure object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BaseDisclosure) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BaseDisclosure) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BaseDisclosure) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *BaseDisclosure) HasId() bool`

HasId returns a boolean if a field has been set.

### GetPersonId

`func (o *BaseDisclosure) GetPersonId() string`

GetPersonId returns the PersonId field if non-nil, zero value otherwise.

### GetPersonIdOk

`func (o *BaseDisclosure) GetPersonIdOk() (*string, bool)`

GetPersonIdOk returns a tuple with the PersonId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPersonId

`func (o *BaseDisclosure) SetPersonId(v string)`

SetPersonId sets PersonId field to given value.

### HasPersonId

`func (o *BaseDisclosure) HasPersonId() bool`

HasPersonId returns a boolean if a field has been set.

### GetBusinessId

`func (o *BaseDisclosure) GetBusinessId() string`

GetBusinessId returns the BusinessId field if non-nil, zero value otherwise.

### GetBusinessIdOk

`func (o *BaseDisclosure) GetBusinessIdOk() (*string, bool)`

GetBusinessIdOk returns a tuple with the BusinessId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessId

`func (o *BaseDisclosure) SetBusinessId(v string)`

SetBusinessId sets BusinessId field to given value.

### HasBusinessId

`func (o *BaseDisclosure) HasBusinessId() bool`

HasBusinessId returns a boolean if a field has been set.

### GetType

`func (o *BaseDisclosure) GetType() DisclosureType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *BaseDisclosure) GetTypeOk() (*DisclosureType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *BaseDisclosure) SetType(v DisclosureType)`

SetType sets Type field to given value.

### HasType

`func (o *BaseDisclosure) HasType() bool`

HasType returns a boolean if a field has been set.

### GetVersion

`func (o *BaseDisclosure) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *BaseDisclosure) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *BaseDisclosure) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *BaseDisclosure) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetEventType

`func (o *BaseDisclosure) GetEventType() string`

GetEventType returns the EventType field if non-nil, zero value otherwise.

### GetEventTypeOk

`func (o *BaseDisclosure) GetEventTypeOk() (*string, bool)`

GetEventTypeOk returns a tuple with the EventType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventType

`func (o *BaseDisclosure) SetEventType(v string)`

SetEventType sets EventType field to given value.

### HasEventType

`func (o *BaseDisclosure) HasEventType() bool`

HasEventType returns a boolean if a field has been set.

### GetDisclosureDate

`func (o *BaseDisclosure) GetDisclosureDate() time.Time`

GetDisclosureDate returns the DisclosureDate field if non-nil, zero value otherwise.

### GetDisclosureDateOk

`func (o *BaseDisclosure) GetDisclosureDateOk() (*time.Time, bool)`

GetDisclosureDateOk returns a tuple with the DisclosureDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisclosureDate

`func (o *BaseDisclosure) SetDisclosureDate(v time.Time)`

SetDisclosureDate sets DisclosureDate field to given value.

### HasDisclosureDate

`func (o *BaseDisclosure) HasDisclosureDate() bool`

HasDisclosureDate returns a boolean if a field has been set.

### GetCreationTime

`func (o *BaseDisclosure) GetCreationTime() time.Time`

GetCreationTime returns the CreationTime field if non-nil, zero value otherwise.

### GetCreationTimeOk

`func (o *BaseDisclosure) GetCreationTimeOk() (*time.Time, bool)`

GetCreationTimeOk returns a tuple with the CreationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationTime

`func (o *BaseDisclosure) SetCreationTime(v time.Time)`

SetCreationTime sets CreationTime field to given value.

### HasCreationTime

`func (o *BaseDisclosure) HasCreationTime() bool`

HasCreationTime returns a boolean if a field has been set.

### GetLastUpdatedTime

`func (o *BaseDisclosure) GetLastUpdatedTime() time.Time`

GetLastUpdatedTime returns the LastUpdatedTime field if non-nil, zero value otherwise.

### GetLastUpdatedTimeOk

`func (o *BaseDisclosure) GetLastUpdatedTimeOk() (*time.Time, bool)`

GetLastUpdatedTimeOk returns a tuple with the LastUpdatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdatedTime

`func (o *BaseDisclosure) SetLastUpdatedTime(v time.Time)`

SetLastUpdatedTime sets LastUpdatedTime field to given value.

### HasLastUpdatedTime

`func (o *BaseDisclosure) HasLastUpdatedTime() bool`

HasLastUpdatedTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


