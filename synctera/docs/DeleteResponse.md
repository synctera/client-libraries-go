# DeleteResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Object ID | [optional] 
**Resource** | Pointer to **string** | The resource name | [optional] 

## Methods

### NewDeleteResponse

`func NewDeleteResponse() *DeleteResponse`

NewDeleteResponse instantiates a new DeleteResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeleteResponseWithDefaults

`func NewDeleteResponseWithDefaults() *DeleteResponse`

NewDeleteResponseWithDefaults instantiates a new DeleteResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *DeleteResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DeleteResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DeleteResponse) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *DeleteResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### GetResource

`func (o *DeleteResponse) GetResource() string`

GetResource returns the Resource field if non-nil, zero value otherwise.

### GetResourceOk

`func (o *DeleteResponse) GetResourceOk() (*string, bool)`

GetResourceOk returns a tuple with the Resource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResource

`func (o *DeleteResponse) SetResource(v string)`

SetResource sets Resource field to given value.

### HasResource

`func (o *DeleteResponse) HasResource() bool`

HasResource returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


