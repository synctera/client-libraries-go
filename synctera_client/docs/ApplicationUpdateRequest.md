# ApplicationUpdateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | Pointer to **bool** | To enable or disable aft/oct feature | [optional] 
**Processor** | Pointer to [**Processor**](Processor.md) |  | [optional] 

## Methods

### NewApplicationUpdateRequest

`func NewApplicationUpdateRequest() *ApplicationUpdateRequest`

NewApplicationUpdateRequest instantiates a new ApplicationUpdateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplicationUpdateRequestWithDefaults

`func NewApplicationUpdateRequestWithDefaults() *ApplicationUpdateRequest`

NewApplicationUpdateRequestWithDefaults instantiates a new ApplicationUpdateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *ApplicationUpdateRequest) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *ApplicationUpdateRequest) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *ApplicationUpdateRequest) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *ApplicationUpdateRequest) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetProcessor

`func (o *ApplicationUpdateRequest) GetProcessor() Processor`

GetProcessor returns the Processor field if non-nil, zero value otherwise.

### GetProcessorOk

`func (o *ApplicationUpdateRequest) GetProcessorOk() (*Processor, bool)`

GetProcessorOk returns a tuple with the Processor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessor

`func (o *ApplicationUpdateRequest) SetProcessor(v Processor)`

SetProcessor sets Processor field to given value.

### HasProcessor

`func (o *ApplicationUpdateRequest) HasProcessor() bool`

HasProcessor returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


