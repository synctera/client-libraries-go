# ExternalCardListResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExternalCards** | [**[]ExternalCardResponse**](ExternalCardResponse.md) | Array of External Cards | 
**NextPageToken** | Pointer to **string** | If returned, use the next_page_token to query for the next page of results. Not returned if there are no more rows. | [optional] 

## Methods

### NewExternalCardListResponse

`func NewExternalCardListResponse(externalCards []ExternalCardResponse, ) *ExternalCardListResponse`

NewExternalCardListResponse instantiates a new ExternalCardListResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExternalCardListResponseWithDefaults

`func NewExternalCardListResponseWithDefaults() *ExternalCardListResponse`

NewExternalCardListResponseWithDefaults instantiates a new ExternalCardListResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExternalCards

`func (o *ExternalCardListResponse) GetExternalCards() []ExternalCardResponse`

GetExternalCards returns the ExternalCards field if non-nil, zero value otherwise.

### GetExternalCardsOk

`func (o *ExternalCardListResponse) GetExternalCardsOk() (*[]ExternalCardResponse, bool)`

GetExternalCardsOk returns a tuple with the ExternalCards field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalCards

`func (o *ExternalCardListResponse) SetExternalCards(v []ExternalCardResponse)`

SetExternalCards sets ExternalCards field to given value.


### GetNextPageToken

`func (o *ExternalCardListResponse) GetNextPageToken() string`

GetNextPageToken returns the NextPageToken field if non-nil, zero value otherwise.

### GetNextPageTokenOk

`func (o *ExternalCardListResponse) GetNextPageTokenOk() (*string, bool)`

GetNextPageTokenOk returns a tuple with the NextPageToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextPageToken

`func (o *ExternalCardListResponse) SetNextPageToken(v string)`

SetNextPageToken sets NextPageToken field to given value.

### HasNextPageToken

`func (o *ExternalCardListResponse) HasNextPageToken() bool`

HasNextPageToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


