# Disclosure

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreationTime** | Pointer to **time.Time** |  | [optional] [readonly] 
**EventType** | **string** |  | 
**Id** | Pointer to **string** | Disclosure ID | [optional] 
**LastUpdatedTime** | Pointer to **time.Time** |  | [optional] [readonly] 
**Timestamp** | **time.Time** | Date of disclosure | 
**Type** | **string** | Disclosure Type | 
**Version** | **string** | Disclosure Version | 

## Methods

### NewDisclosure

`func NewDisclosure(eventType string, timestamp time.Time, type_ string, version string, ) *Disclosure`

NewDisclosure instantiates a new Disclosure object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDisclosureWithDefaults

`func NewDisclosureWithDefaults() *Disclosure`

NewDisclosureWithDefaults instantiates a new Disclosure object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

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

### GetTimestamp

`func (o *Disclosure) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *Disclosure) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *Disclosure) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.


### GetType

`func (o *Disclosure) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Disclosure) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Disclosure) SetType(v string)`

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



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


