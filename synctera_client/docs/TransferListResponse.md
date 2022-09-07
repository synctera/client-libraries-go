# TransferListResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExternalTransfers** | [**[]TransferResponse**](TransferResponse.md) | Array of External transfer | 
**NextPageToken** | Pointer to **string** | If returned, use the next_page_token to query for the next page of results. Not returned if there are no more rows. | [optional] 

## Methods

### NewTransferListResponse

`func NewTransferListResponse(externalTransfers []TransferResponse, ) *TransferListResponse`

NewTransferListResponse instantiates a new TransferListResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransferListResponseWithDefaults

`func NewTransferListResponseWithDefaults() *TransferListResponse`

NewTransferListResponseWithDefaults instantiates a new TransferListResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExternalTransfers

`func (o *TransferListResponse) GetExternalTransfers() []TransferResponse`

GetExternalTransfers returns the ExternalTransfers field if non-nil, zero value otherwise.

### GetExternalTransfersOk

`func (o *TransferListResponse) GetExternalTransfersOk() (*[]TransferResponse, bool)`

GetExternalTransfersOk returns a tuple with the ExternalTransfers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalTransfers

`func (o *TransferListResponse) SetExternalTransfers(v []TransferResponse)`

SetExternalTransfers sets ExternalTransfers field to given value.


### GetNextPageToken

`func (o *TransferListResponse) GetNextPageToken() string`

GetNextPageToken returns the NextPageToken field if non-nil, zero value otherwise.

### GetNextPageTokenOk

`func (o *TransferListResponse) GetNextPageTokenOk() (*string, bool)`

GetNextPageTokenOk returns a tuple with the NextPageToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextPageToken

`func (o *TransferListResponse) SetNextPageToken(v string)`

SetNextPageToken sets NextPageToken field to given value.

### HasNextPageToken

`func (o *TransferListResponse) HasNextPageToken() bool`

HasNextPageToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


