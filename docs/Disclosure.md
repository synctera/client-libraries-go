# Disclosure

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The unique identifier for this resource. | [optional] [readonly] 
**PersonId** | Pointer to **string** | Unique ID for the person. Exactly one of &#x60;person_id&#x60; or &#x60;business_id&#x60; must be set.  | [optional] 
**BusinessId** | Pointer to **string** | Unique ID for the business. Exactly one of &#x60;business_id&#x60; or &#x60;person_id&#x60; must be set.  | [optional] 
**Type** | [**DisclosureType**](DisclosureType.md) |  | 
**Version** | **string** | Version of the disclosure document. | 
**EventType** | **string** |  | 
**DisclosureDate** | **time.Time** | Date and time the disclosure was made. | 
**CreationTime** | Pointer to **time.Time** | The date and time the resource was created. | [optional] [readonly] 
**LastUpdatedTime** | Pointer to **time.Time** | The date and time the resource was last updated. | [optional] [readonly] 

## Methods

### NewDisclosure

`func NewDisclosure(type_ DisclosureType, version string, eventType string, disclosureDate time.Time, ) *Disclosure`

NewDisclosure instantiates a new Disclosure object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDisclosureWithDefaults

`func NewDisclosureWithDefaults() *Disclosure`

NewDisclosureWithDefaults instantiates a new Disclosure object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Disclosure) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Disclosure) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Disclosure) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Disclosure) HasId() bool`

HasId returns a boolean if a field has been set.

### GetPersonId

`func (o *Disclosure) GetPersonId() string`

GetPersonId returns the PersonId field if non-nil, zero value otherwise.

### GetPersonIdOk

`func (o *Disclosure) GetPersonIdOk() (*string, bool)`

GetPersonIdOk returns a tuple with the PersonId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPersonId

`func (o *Disclosure) SetPersonId(v string)`

SetPersonId sets PersonId field to given value.

### HasPersonId

`func (o *Disclosure) HasPersonId() bool`

HasPersonId returns a boolean if a field has been set.

### GetBusinessId

`func (o *Disclosure) GetBusinessId() string`

GetBusinessId returns the BusinessId field if non-nil, zero value otherwise.

### GetBusinessIdOk

`func (o *Disclosure) GetBusinessIdOk() (*string, bool)`

GetBusinessIdOk returns a tuple with the BusinessId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessId

`func (o *Disclosure) SetBusinessId(v string)`

SetBusinessId sets BusinessId field to given value.

### HasBusinessId

`func (o *Disclosure) HasBusinessId() bool`

HasBusinessId returns a boolean if a field has been set.

### GetType

`func (o *Disclosure) GetType() DisclosureType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Disclosure) GetTypeOk() (*DisclosureType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Disclosure) SetType(v DisclosureType)`

SetType sets Type field to given value.


### GetVersion

`func (o *Disclosure) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *Disclosure) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *Disclosure) SetVersion(v string)`

SetVersion sets Version field to given value.


### GetEventType

`func (o *Disclosure) GetEventType() string`

GetEventType returns the EventType field if non-nil, zero value otherwise.

### GetEventTypeOk

`func (o *Disclosure) GetEventTypeOk() (*string, bool)`

GetEventTypeOk returns a tuple with the EventType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventType

`func (o *Disclosure) SetEventType(v string)`

SetEventType sets EventType field to given value.


### GetDisclosureDate

`func (o *Disclosure) GetDisclosureDate() time.Time`

GetDisclosureDate returns the DisclosureDate field if non-nil, zero value otherwise.

### GetDisclosureDateOk

`func (o *Disclosure) GetDisclosureDateOk() (*time.Time, bool)`

GetDisclosureDateOk returns a tuple with the DisclosureDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisclosureDate

`func (o *Disclosure) SetDisclosureDate(v time.Time)`

SetDisclosureDate sets DisclosureDate field to given value.


### GetCreationTime

`func (o *Disclosure) GetCreationTime() time.Time`

GetCreationTime returns the CreationTime field if non-nil, zero value otherwise.

### GetCreationTimeOk

`func (o *Disclosure) GetCreationTimeOk() (*time.Time, bool)`

GetCreationTimeOk returns a tuple with the CreationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationTime

`func (o *Disclosure) SetCreationTime(v time.Time)`

SetCreationTime sets CreationTime field to given value.

### HasCreationTime

`func (o *Disclosure) HasCreationTime() bool`

HasCreationTime returns a boolean if a field has been set.

### GetLastUpdatedTime

`func (o *Disclosure) GetLastUpdatedTime() time.Time`

GetLastUpdatedTime returns the LastUpdatedTime field if non-nil, zero value otherwise.

### GetLastUpdatedTimeOk

`func (o *Disclosure) GetLastUpdatedTimeOk() (*time.Time, bool)`

GetLastUpdatedTimeOk returns a tuple with the LastUpdatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdatedTime

`func (o *Disclosure) SetLastUpdatedTime(v time.Time)`

SetLastUpdatedTime sets LastUpdatedTime field to given value.

### HasLastUpdatedTime

`func (o *Disclosure) HasLastUpdatedTime() bool`

HasLastUpdatedTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


