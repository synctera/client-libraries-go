# BinResponseList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Bins** | [**[]BinResponse**](BinResponse.md) | Array of Bins | 
**NextPageToken** | Pointer to **string** | If returned, use the next_page_token to query for the next page of results. Not returned if there are no more rows. | [optional] 

## Methods

### NewBinResponseList

`func NewBinResponseList(bins []BinResponse, ) *BinResponseList`

NewBinResponseList instantiates a new BinResponseList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBinResponseListWithDefaults

`func NewBinResponseListWithDefaults() *BinResponseList`

NewBinResponseListWithDefaults instantiates a new BinResponseList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBins

`func (o *BinResponseList) GetBins() []BinResponse`

GetBins returns the Bins field if non-nil, zero value otherwise.

### GetBinsOk

`func (o *BinResponseList) GetBinsOk() (*[]BinResponse, bool)`

GetBinsOk returns a tuple with the Bins field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBins

`func (o *BinResponseList) SetBins(v []BinResponse)`

SetBins sets Bins field to given value.


### GetNextPageToken

`func (o *BinResponseList) GetNextPageToken() string`

GetNextPageToken returns the NextPageToken field if non-nil, zero value otherwise.

### GetNextPageTokenOk

`func (o *BinResponseList) GetNextPageTokenOk() (*string, bool)`

GetNextPageTokenOk returns a tuple with the NextPageToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextPageToken

`func (o *BinResponseList) SetNextPageToken(v string)`

SetNextPageToken sets NextPageToken field to given value.

### HasNextPageToken

`func (o *BinResponseList) HasNextPageToken() bool`

HasNextPageToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


