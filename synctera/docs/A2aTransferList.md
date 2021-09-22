# A2aTransferList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**A2aTransfers** | [**[]A2aTransfer**](A2aTransfer.md) | Array of account to account transfers | 
**NextPageToken** | Pointer to **string** | If returned, use the next_page_token to query for the next page of results. Not returned if there are no more rows. | [optional] 

## Methods

### NewA2aTransferList

`func NewA2aTransferList(a2aTransfers []A2aTransfer, ) *A2aTransferList`

NewA2aTransferList instantiates a new A2aTransferList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewA2aTransferListWithDefaults

`func NewA2aTransferListWithDefaults() *A2aTransferList`

NewA2aTransferListWithDefaults instantiates a new A2aTransferList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetA2aTransfers

`func (o *A2aTransferList) GetA2aTransfers() []A2aTransfer`

GetA2aTransfers returns the A2aTransfers field if non-nil, zero value otherwise.

### GetA2aTransfersOk

`func (o *A2aTransferList) GetA2aTransfersOk() (*[]A2aTransfer, bool)`

GetA2aTransfersOk returns a tuple with the A2aTransfers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetA2aTransfers

`func (o *A2aTransferList) SetA2aTransfers(v []A2aTransfer)`

SetA2aTransfers sets A2aTransfers field to given value.


### GetNextPageToken

`func (o *A2aTransferList) GetNextPageToken() string`

GetNextPageToken returns the NextPageToken field if non-nil, zero value otherwise.

### GetNextPageTokenOk

`func (o *A2aTransferList) GetNextPageTokenOk() (*string, bool)`

GetNextPageTokenOk returns a tuple with the NextPageToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextPageToken

`func (o *A2aTransferList) SetNextPageToken(v string)`

SetNextPageToken sets NextPageToken field to given value.

### HasNextPageToken

`func (o *A2aTransferList) HasNextPageToken() bool`

HasNextPageToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


