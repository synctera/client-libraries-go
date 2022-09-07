# DepositList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Deposits** | [**[]Deposit**](Deposit.md) | Array of  Remote Check Deposits | 
**NextPageToken** | Pointer to **string** | If returned, use the next_page_token to query for the next page of results. Not returned if there are no more rows. | [optional] 

## Methods

### NewDepositList

`func NewDepositList(deposits []Deposit, ) *DepositList`

NewDepositList instantiates a new DepositList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDepositListWithDefaults

`func NewDepositListWithDefaults() *DepositList`

NewDepositListWithDefaults instantiates a new DepositList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeposits

`func (o *DepositList) GetDeposits() []Deposit`

GetDeposits returns the Deposits field if non-nil, zero value otherwise.

### GetDepositsOk

`func (o *DepositList) GetDepositsOk() (*[]Deposit, bool)`

GetDepositsOk returns a tuple with the Deposits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeposits

`func (o *DepositList) SetDeposits(v []Deposit)`

SetDeposits sets Deposits field to given value.


### GetNextPageToken

`func (o *DepositList) GetNextPageToken() string`

GetNextPageToken returns the NextPageToken field if non-nil, zero value otherwise.

### GetNextPageTokenOk

`func (o *DepositList) GetNextPageTokenOk() (*string, bool)`

GetNextPageTokenOk returns a tuple with the NextPageToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextPageToken

`func (o *DepositList) SetNextPageToken(v string)`

SetNextPageToken sets NextPageToken field to given value.

### HasNextPageToken

`func (o *DepositList) HasNextPageToken() bool`

HasNextPageToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


