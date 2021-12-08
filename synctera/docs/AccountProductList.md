# AccountProductList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountProducts** | [**[]AccountProduct**](AccountProduct.md) | Array of account products | 
**NextPageToken** | Pointer to **string** | If returned, use the next_page_token to query for the next page of results. Not returned if there are no more rows. | [optional] 

## Methods

### NewAccountProductList

`func NewAccountProductList(accountProducts []AccountProduct, ) *AccountProductList`

NewAccountProductList instantiates a new AccountProductList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountProductListWithDefaults

`func NewAccountProductListWithDefaults() *AccountProductList`

NewAccountProductListWithDefaults instantiates a new AccountProductList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountProducts

`func (o *AccountProductList) GetAccountProducts() []AccountProduct`

GetAccountProducts returns the AccountProducts field if non-nil, zero value otherwise.

### GetAccountProductsOk

`func (o *AccountProductList) GetAccountProductsOk() (*[]AccountProduct, bool)`

GetAccountProductsOk returns a tuple with the AccountProducts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountProducts

`func (o *AccountProductList) SetAccountProducts(v []AccountProduct)`

SetAccountProducts sets AccountProducts field to given value.


### GetNextPageToken

`func (o *AccountProductList) GetNextPageToken() string`

GetNextPageToken returns the NextPageToken field if non-nil, zero value otherwise.

### GetNextPageTokenOk

`func (o *AccountProductList) GetNextPageTokenOk() (*string, bool)`

GetNextPageTokenOk returns a tuple with the NextPageToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextPageToken

`func (o *AccountProductList) SetNextPageToken(v string)`

SetNextPageToken sets NextPageToken field to given value.

### HasNextPageToken

`func (o *AccountProductList) HasNextPageToken() bool`

HasNextPageToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


