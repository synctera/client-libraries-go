# SpendingLimitsTransaction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | Pointer to **int64** | Maximum amount allowed | [optional] 

## Methods

### NewSpendingLimitsTransaction

`func NewSpendingLimitsTransaction() *SpendingLimitsTransaction`

NewSpendingLimitsTransaction instantiates a new SpendingLimitsTransaction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSpendingLimitsTransactionWithDefaults

`func NewSpendingLimitsTransactionWithDefaults() *SpendingLimitsTransaction`

NewSpendingLimitsTransactionWithDefaults instantiates a new SpendingLimitsTransaction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmount

`func (o *SpendingLimitsTransaction) GetAmount() int64`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *SpendingLimitsTransaction) GetAmountOk() (*int64, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *SpendingLimitsTransaction) SetAmount(v int64)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *SpendingLimitsTransaction) HasAmount() bool`

HasAmount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


