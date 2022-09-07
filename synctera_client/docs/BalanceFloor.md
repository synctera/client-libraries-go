# BalanceFloor

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Balance** | **int64** | Minimum balance in the account&#39;s currency. Unit in cents. | 
**LinkedAccountId** | Pointer to **string** | ID of linked backing account for just-in-time (JIT) funding of transactions to maintain the balance floor  | [optional] 
**OverdraftAccountId** | Pointer to **string** | ID of linked backing account for just-in-time (JIT) funding of transactions to maintain the balance floor This attribute is a deprecated alias for linked_account_id.  | [optional] 

## Methods

### NewBalanceFloor

`func NewBalanceFloor(balance int64, ) *BalanceFloor`

NewBalanceFloor instantiates a new BalanceFloor object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBalanceFloorWithDefaults

`func NewBalanceFloorWithDefaults() *BalanceFloor`

NewBalanceFloorWithDefaults instantiates a new BalanceFloor object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBalance

`func (o *BalanceFloor) GetBalance() int64`

GetBalance returns the Balance field if non-nil, zero value otherwise.

### GetBalanceOk

`func (o *BalanceFloor) GetBalanceOk() (*int64, bool)`

GetBalanceOk returns a tuple with the Balance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalance

`func (o *BalanceFloor) SetBalance(v int64)`

SetBalance sets Balance field to given value.


### GetLinkedAccountId

`func (o *BalanceFloor) GetLinkedAccountId() string`

GetLinkedAccountId returns the LinkedAccountId field if non-nil, zero value otherwise.

### GetLinkedAccountIdOk

`func (o *BalanceFloor) GetLinkedAccountIdOk() (*string, bool)`

GetLinkedAccountIdOk returns a tuple with the LinkedAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkedAccountId

`func (o *BalanceFloor) SetLinkedAccountId(v string)`

SetLinkedAccountId sets LinkedAccountId field to given value.

### HasLinkedAccountId

`func (o *BalanceFloor) HasLinkedAccountId() bool`

HasLinkedAccountId returns a boolean if a field has been set.

### GetOverdraftAccountId

`func (o *BalanceFloor) GetOverdraftAccountId() string`

GetOverdraftAccountId returns the OverdraftAccountId field if non-nil, zero value otherwise.

### GetOverdraftAccountIdOk

`func (o *BalanceFloor) GetOverdraftAccountIdOk() (*string, bool)`

GetOverdraftAccountIdOk returns a tuple with the OverdraftAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverdraftAccountId

`func (o *BalanceFloor) SetOverdraftAccountId(v string)`

SetOverdraftAccountId sets OverdraftAccountId field to given value.

### HasOverdraftAccountId

`func (o *BalanceFloor) HasOverdraftAccountId() bool`

HasOverdraftAccountId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


