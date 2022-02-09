# SpendingLimitWithTime

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Transactions** | Pointer to **int64** | Maximum number of transactions allowed within the time range | [optional] 
**Amount** | Pointer to **int64** | Maximum amount allowed within the time range | [optional] 

## Methods

### NewSpendingLimitWithTime

`func NewSpendingLimitWithTime() *SpendingLimitWithTime`

NewSpendingLimitWithTime instantiates a new SpendingLimitWithTime object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSpendingLimitWithTimeWithDefaults

`func NewSpendingLimitWithTimeWithDefaults() *SpendingLimitWithTime`

NewSpendingLimitWithTimeWithDefaults instantiates a new SpendingLimitWithTime object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTransactions

`func (o *SpendingLimitWithTime) GetTransactions() int64`

GetTransactions returns the Transactions field if non-nil, zero value otherwise.

### GetTransactionsOk

`func (o *SpendingLimitWithTime) GetTransactionsOk() (*int64, bool)`

GetTransactionsOk returns a tuple with the Transactions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactions

`func (o *SpendingLimitWithTime) SetTransactions(v int64)`

SetTransactions sets Transactions field to given value.

### HasTransactions

`func (o *SpendingLimitWithTime) HasTransactions() bool`

HasTransactions returns a boolean if a field has been set.

### GetAmount

`func (o *SpendingLimitWithTime) GetAmount() int64`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *SpendingLimitWithTime) GetAmountOk() (*int64, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *SpendingLimitWithTime) SetAmount(v int64)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *SpendingLimitWithTime) HasAmount() bool`

HasAmount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


