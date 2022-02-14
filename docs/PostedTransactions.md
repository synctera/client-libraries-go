# PostedTransactions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NextPageToken** | **NullableString** | If returned, use the next_page_token to query for the next page of results. Not returned if there are no more rows. | 
**Result** | [**[]PostedTransaction**](PostedTransaction.md) | List of posted transactions | 

## Methods

### NewPostedTransactions

`func NewPostedTransactions(nextPageToken NullableString, result []PostedTransaction, ) *PostedTransactions`

NewPostedTransactions instantiates a new PostedTransactions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostedTransactionsWithDefaults

`func NewPostedTransactionsWithDefaults() *PostedTransactions`

NewPostedTransactionsWithDefaults instantiates a new PostedTransactions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNextPageToken

`func (o *PostedTransactions) GetNextPageToken() string`

GetNextPageToken returns the NextPageToken field if non-nil, zero value otherwise.

### GetNextPageTokenOk

`func (o *PostedTransactions) GetNextPageTokenOk() (*string, bool)`

GetNextPageTokenOk returns a tuple with the NextPageToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextPageToken

`func (o *PostedTransactions) SetNextPageToken(v string)`

SetNextPageToken sets NextPageToken field to given value.


### SetNextPageTokenNil

`func (o *PostedTransactions) SetNextPageTokenNil(b bool)`

 SetNextPageTokenNil sets the value for NextPageToken to be an explicit nil

### UnsetNextPageToken
`func (o *PostedTransactions) UnsetNextPageToken()`

UnsetNextPageToken ensures that no value is present for NextPageToken, not even an explicit nil
### GetResult

`func (o *PostedTransactions) GetResult() []PostedTransaction`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *PostedTransactions) GetResultOk() (*[]PostedTransaction, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *PostedTransactions) SetResult(v []PostedTransaction)`

SetResult sets Result field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


