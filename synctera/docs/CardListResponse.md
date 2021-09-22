# CardListResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cards** | [**[]CardResponse**](CardResponse.md) | Array of Cards | 
**NextPageToken** | Pointer to **string** | If returned, use the next_page_token to query for the next page of results. Not returned if there are no more rows. | [optional] 

## Methods

### NewCardListResponse

`func NewCardListResponse(cards []CardResponse, ) *CardListResponse`

NewCardListResponse instantiates a new CardListResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCardListResponseWithDefaults

`func NewCardListResponseWithDefaults() *CardListResponse`

NewCardListResponseWithDefaults instantiates a new CardListResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCards

`func (o *CardListResponse) GetCards() []CardResponse`

GetCards returns the Cards field if non-nil, zero value otherwise.

### GetCardsOk

`func (o *CardListResponse) GetCardsOk() (*[]CardResponse, bool)`

GetCardsOk returns a tuple with the Cards field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCards

`func (o *CardListResponse) SetCards(v []CardResponse)`

SetCards sets Cards field to given value.


### GetNextPageToken

`func (o *CardListResponse) GetNextPageToken() string`

GetNextPageToken returns the NextPageToken field if non-nil, zero value otherwise.

### GetNextPageTokenOk

`func (o *CardListResponse) GetNextPageTokenOk() (*string, bool)`

GetNextPageTokenOk returns a tuple with the NextPageToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextPageToken

`func (o *CardListResponse) SetNextPageToken(v string)`

SetNextPageToken sets NextPageToken field to given value.

### HasNextPageToken

`func (o *CardListResponse) HasNextPageToken() bool`

HasNextPageToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


