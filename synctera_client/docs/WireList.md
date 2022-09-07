# WireList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Wires** | [**[]Wire**](Wire.md) | Array of wires | 
**NextPageToken** | Pointer to **string** | If returned, use the next_page_token to query for the next page of results. Not returned if there are no more rows. | [optional] 

## Methods

### NewWireList

`func NewWireList(wires []Wire, ) *WireList`

NewWireList instantiates a new WireList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWireListWithDefaults

`func NewWireListWithDefaults() *WireList`

NewWireListWithDefaults instantiates a new WireList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWires

`func (o *WireList) GetWires() []Wire`

GetWires returns the Wires field if non-nil, zero value otherwise.

### GetWiresOk

`func (o *WireList) GetWiresOk() (*[]Wire, bool)`

GetWiresOk returns a tuple with the Wires field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWires

`func (o *WireList) SetWires(v []Wire)`

SetWires sets Wires field to given value.


### GetNextPageToken

`func (o *WireList) GetNextPageToken() string`

GetNextPageToken returns the NextPageToken field if non-nil, zero value otherwise.

### GetNextPageTokenOk

`func (o *WireList) GetNextPageTokenOk() (*string, bool)`

GetNextPageTokenOk returns a tuple with the NextPageToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextPageToken

`func (o *WireList) SetNextPageToken(v string)`

SetNextPageToken sets NextPageToken field to given value.

### HasNextPageToken

`func (o *WireList) HasNextPageToken() bool`

HasNextPageToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


