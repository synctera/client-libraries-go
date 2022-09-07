# AccountRangeResponseList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountRanges** | [**[]AccountRangeResponse**](AccountRangeResponse.md) | Array of Account Ranges | 
**NextPageToken** | Pointer to **string** | If returned, use the next_page_token to query for the next page of results. Not returned if there are no more rows. | [optional] 

## Methods

### NewAccountRangeResponseList

`func NewAccountRangeResponseList(accountRanges []AccountRangeResponse, ) *AccountRangeResponseList`

NewAccountRangeResponseList instantiates a new AccountRangeResponseList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountRangeResponseListWithDefaults

`func NewAccountRangeResponseListWithDefaults() *AccountRangeResponseList`

NewAccountRangeResponseListWithDefaults instantiates a new AccountRangeResponseList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountRanges

`func (o *AccountRangeResponseList) GetAccountRanges() []AccountRangeResponse`

GetAccountRanges returns the AccountRanges field if non-nil, zero value otherwise.

### GetAccountRangesOk

`func (o *AccountRangeResponseList) GetAccountRangesOk() (*[]AccountRangeResponse, bool)`

GetAccountRangesOk returns a tuple with the AccountRanges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountRanges

`func (o *AccountRangeResponseList) SetAccountRanges(v []AccountRangeResponse)`

SetAccountRanges sets AccountRanges field to given value.


### GetNextPageToken

`func (o *AccountRangeResponseList) GetNextPageToken() string`

GetNextPageToken returns the NextPageToken field if non-nil, zero value otherwise.

### GetNextPageTokenOk

`func (o *AccountRangeResponseList) GetNextPageTokenOk() (*string, bool)`

GetNextPageTokenOk returns a tuple with the NextPageToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextPageToken

`func (o *AccountRangeResponseList) SetNextPageToken(v string)`

SetNextPageToken sets NextPageToken field to given value.

### HasNextPageToken

`func (o *AccountRangeResponseList) HasNextPageToken() bool`

HasNextPageToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


