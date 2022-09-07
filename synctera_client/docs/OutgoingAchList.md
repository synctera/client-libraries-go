# OutgoingAchList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Transactions** | [**[]OutgoingAch**](OutgoingAch.md) | Array of sent ACH transactions. | 
**NextPageToken** | Pointer to **string** | If returned, use the next_page_token to query for the next page of results. Not returned if there are no more rows. | [optional] 

## Methods

### NewOutgoingAchList

`func NewOutgoingAchList(transactions []OutgoingAch, ) *OutgoingAchList`

NewOutgoingAchList instantiates a new OutgoingAchList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOutgoingAchListWithDefaults

`func NewOutgoingAchListWithDefaults() *OutgoingAchList`

NewOutgoingAchListWithDefaults instantiates a new OutgoingAchList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTransactions

`func (o *OutgoingAchList) GetTransactions() []OutgoingAch`

GetTransactions returns the Transactions field if non-nil, zero value otherwise.

### GetTransactionsOk

`func (o *OutgoingAchList) GetTransactionsOk() (*[]OutgoingAch, bool)`

GetTransactionsOk returns a tuple with the Transactions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactions

`func (o *OutgoingAchList) SetTransactions(v []OutgoingAch)`

SetTransactions sets Transactions field to given value.


### GetNextPageToken

`func (o *OutgoingAchList) GetNextPageToken() string`

GetNextPageToken returns the NextPageToken field if non-nil, zero value otherwise.

### GetNextPageTokenOk

`func (o *OutgoingAchList) GetNextPageTokenOk() (*string, bool)`

GetNextPageTokenOk returns a tuple with the NextPageToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextPageToken

`func (o *OutgoingAchList) SetNextPageToken(v string)`

SetNextPageToken sets NextPageToken field to given value.

### HasNextPageToken

`func (o *OutgoingAchList) HasNextPageToken() bool`

HasNextPageToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


