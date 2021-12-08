# CardProductListResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CardProducts** | [**[]CardProductResponse**](CardProductResponse.md) | Array of Card Products | 
**NextPageToken** | Pointer to **string** | If returned, use the next_page_token to query for the next page of results. Not returned if there are no more rows. | [optional] 

## Methods

### NewCardProductListResponse

`func NewCardProductListResponse(cardProducts []CardProductResponse, ) *CardProductListResponse`

NewCardProductListResponse instantiates a new CardProductListResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCardProductListResponseWithDefaults

`func NewCardProductListResponseWithDefaults() *CardProductListResponse`

NewCardProductListResponseWithDefaults instantiates a new CardProductListResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCardProducts

`func (o *CardProductListResponse) GetCardProducts() []CardProductResponse`

GetCardProducts returns the CardProducts field if non-nil, zero value otherwise.

### GetCardProductsOk

`func (o *CardProductListResponse) GetCardProductsOk() (*[]CardProductResponse, bool)`

GetCardProductsOk returns a tuple with the CardProducts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardProducts

`func (o *CardProductListResponse) SetCardProducts(v []CardProductResponse)`

SetCardProducts sets CardProducts field to given value.


### GetNextPageToken

`func (o *CardProductListResponse) GetNextPageToken() string`

GetNextPageToken returns the NextPageToken field if non-nil, zero value otherwise.

### GetNextPageTokenOk

`func (o *CardProductListResponse) GetNextPageTokenOk() (*string, bool)`

GetNextPageTokenOk returns a tuple with the NextPageToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextPageToken

`func (o *CardProductListResponse) SetNextPageToken(v string)`

SetNextPageToken sets NextPageToken field to given value.

### HasNextPageToken

`func (o *CardProductListResponse) HasNextPageToken() bool`

HasNextPageToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


