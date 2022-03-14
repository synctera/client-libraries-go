# BalanceCeiling

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Balance** | **int64** | Maximum balance in the account&#39;s currency | 
**OverflowAccountId** | Pointer to **string** | ID of overflow account for JIT funding of transactions | [optional] 

## Methods

### NewBalanceCeiling

`func NewBalanceCeiling(balance int64, ) *BalanceCeiling`

NewBalanceCeiling instantiates a new BalanceCeiling object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBalanceCeilingWithDefaults

`func NewBalanceCeilingWithDefaults() *BalanceCeiling`

NewBalanceCeilingWithDefaults instantiates a new BalanceCeiling object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBalance

`func (o *BalanceCeiling) GetBalance() int64`

GetBalance returns the Balance field if non-nil, zero value otherwise.

### GetBalanceOk

`func (o *BalanceCeiling) GetBalanceOk() (*int64, bool)`

GetBalanceOk returns a tuple with the Balance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalance

`func (o *BalanceCeiling) SetBalance(v int64)`

SetBalance sets Balance field to given value.


### GetOverflowAccountId

`func (o *BalanceCeiling) GetOverflowAccountId() string`

GetOverflowAccountId returns the OverflowAccountId field if non-nil, zero value otherwise.

### GetOverflowAccountIdOk

`func (o *BalanceCeiling) GetOverflowAccountIdOk() (*string, bool)`

GetOverflowAccountIdOk returns a tuple with the OverflowAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverflowAccountId

`func (o *BalanceCeiling) SetOverflowAccountId(v string)`

SetOverflowAccountId sets OverflowAccountId field to given value.

### HasOverflowAccountId

`func (o *BalanceCeiling) HasOverflowAccountId() bool`

HasOverflowAccountId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


