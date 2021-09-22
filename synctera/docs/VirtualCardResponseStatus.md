# VirtualCardResponseStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to [**CardStatusObject**](CardStatusObject.md) |  | [optional] 

## Methods

### NewVirtualCardResponseStatus

`func NewVirtualCardResponseStatus() *VirtualCardResponseStatus`

NewVirtualCardResponseStatus instantiates a new VirtualCardResponseStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVirtualCardResponseStatusWithDefaults

`func NewVirtualCardResponseStatusWithDefaults() *VirtualCardResponseStatus`

NewVirtualCardResponseStatusWithDefaults instantiates a new VirtualCardResponseStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *VirtualCardResponseStatus) GetStatus() CardStatusObject`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *VirtualCardResponseStatus) GetStatusOk() (*CardStatusObject, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *VirtualCardResponseStatus) SetStatus(v CardStatusObject)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *VirtualCardResponseStatus) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


