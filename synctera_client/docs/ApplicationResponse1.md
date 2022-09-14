# ApplicationResponse1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedTime** | **time.Time** |  | 
**Enabled** | **bool** | To enable or disable aft/oct feature | 
**ExternalId** | Pointer to **string** | The id of the application from processor | [optional] 
**Id** | **string** | The id of the application | 
**LastModifiedTime** | **time.Time** |  | 
**Processor** | [**Processor**](Processor.md) |  | 

## Methods

### NewApplicationResponse1

`func NewApplicationResponse1(createdTime time.Time, enabled bool, id string, lastModifiedTime time.Time, processor Processor, ) *ApplicationResponse1`

NewApplicationResponse1 instantiates a new ApplicationResponse1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplicationResponse1WithDefaults

`func NewApplicationResponse1WithDefaults() *ApplicationResponse1`

NewApplicationResponse1WithDefaults instantiates a new ApplicationResponse1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedTime

`func (o *ApplicationResponse1) GetCreatedTime() time.Time`

GetCreatedTime returns the CreatedTime field if non-nil, zero value otherwise.

### GetCreatedTimeOk

`func (o *ApplicationResponse1) GetCreatedTimeOk() (*time.Time, bool)`

GetCreatedTimeOk returns a tuple with the CreatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTime

`func (o *ApplicationResponse1) SetCreatedTime(v time.Time)`

SetCreatedTime sets CreatedTime field to given value.


### GetEnabled

`func (o *ApplicationResponse1) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *ApplicationResponse1) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *ApplicationResponse1) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetExternalId

`func (o *ApplicationResponse1) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *ApplicationResponse1) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *ApplicationResponse1) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *ApplicationResponse1) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### GetId

`func (o *ApplicationResponse1) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ApplicationResponse1) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ApplicationResponse1) SetId(v string)`

SetId sets Id field to given value.


### GetLastModifiedTime

`func (o *ApplicationResponse1) GetLastModifiedTime() time.Time`

GetLastModifiedTime returns the LastModifiedTime field if non-nil, zero value otherwise.

### GetLastModifiedTimeOk

`func (o *ApplicationResponse1) GetLastModifiedTimeOk() (*time.Time, bool)`

GetLastModifiedTimeOk returns a tuple with the LastModifiedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModifiedTime

`func (o *ApplicationResponse1) SetLastModifiedTime(v time.Time)`

SetLastModifiedTime sets LastModifiedTime field to given value.


### GetProcessor

`func (o *ApplicationResponse1) GetProcessor() Processor`

GetProcessor returns the Processor field if non-nil, zero value otherwise.

### GetProcessorOk

`func (o *ApplicationResponse1) GetProcessorOk() (*Processor, bool)`

GetProcessorOk returns a tuple with the Processor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessor

`func (o *ApplicationResponse1) SetProcessor(v Processor)`

SetProcessor sets Processor field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


