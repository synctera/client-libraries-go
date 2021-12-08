# BinAndDebitNetworkList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BinAndDebitNetworks** | [**[]BinAndDebitNetwork**](BinAndDebitNetwork.md) | Array of BINs and debit networks | 
**NextPageToken** | Pointer to **string** | If returned, use the next_page_token to query for the next page of results. Not returned if there are no more rows. | [optional] 

## Methods

### NewBinAndDebitNetworkList

`func NewBinAndDebitNetworkList(binAndDebitNetworks []BinAndDebitNetwork, ) *BinAndDebitNetworkList`

NewBinAndDebitNetworkList instantiates a new BinAndDebitNetworkList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBinAndDebitNetworkListWithDefaults

`func NewBinAndDebitNetworkListWithDefaults() *BinAndDebitNetworkList`

NewBinAndDebitNetworkListWithDefaults instantiates a new BinAndDebitNetworkList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBinAndDebitNetworks

`func (o *BinAndDebitNetworkList) GetBinAndDebitNetworks() []BinAndDebitNetwork`

GetBinAndDebitNetworks returns the BinAndDebitNetworks field if non-nil, zero value otherwise.

### GetBinAndDebitNetworksOk

`func (o *BinAndDebitNetworkList) GetBinAndDebitNetworksOk() (*[]BinAndDebitNetwork, bool)`

GetBinAndDebitNetworksOk returns a tuple with the BinAndDebitNetworks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBinAndDebitNetworks

`func (o *BinAndDebitNetworkList) SetBinAndDebitNetworks(v []BinAndDebitNetwork)`

SetBinAndDebitNetworks sets BinAndDebitNetworks field to given value.


### GetNextPageToken

`func (o *BinAndDebitNetworkList) GetNextPageToken() string`

GetNextPageToken returns the NextPageToken field if non-nil, zero value otherwise.

### GetNextPageTokenOk

`func (o *BinAndDebitNetworkList) GetNextPageTokenOk() (*string, bool)`

GetNextPageTokenOk returns a tuple with the NextPageToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextPageToken

`func (o *BinAndDebitNetworkList) SetNextPageToken(v string)`

SetNextPageToken sets NextPageToken field to given value.

### HasNextPageToken

`func (o *BinAndDebitNetworkList) HasNextPageToken() bool`

HasNextPageToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


