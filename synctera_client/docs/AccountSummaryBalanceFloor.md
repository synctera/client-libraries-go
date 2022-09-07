# AccountSummaryBalanceFloor

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Balance** | Pointer to **int64** | Minimum balance in the account&#39;s currency. Unit in cents. | [optional] 

## Methods

### NewAccountSummaryBalanceFloor

`func NewAccountSummaryBalanceFloor() *AccountSummaryBalanceFloor`

NewAccountSummaryBalanceFloor instantiates a new AccountSummaryBalanceFloor object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountSummaryBalanceFloorWithDefaults

`func NewAccountSummaryBalanceFloorWithDefaults() *AccountSummaryBalanceFloor`

NewAccountSummaryBalanceFloorWithDefaults instantiates a new AccountSummaryBalanceFloor object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBalance

`func (o *AccountSummaryBalanceFloor) GetBalance() int64`

GetBalance returns the Balance field if non-nil, zero value otherwise.

### GetBalanceOk

`func (o *AccountSummaryBalanceFloor) GetBalanceOk() (*int64, bool)`

GetBalanceOk returns a tuple with the Balance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalance

`func (o *AccountSummaryBalanceFloor) SetBalance(v int64)`

SetBalance sets Balance field to given value.

### HasBalance

`func (o *AccountSummaryBalanceFloor) HasBalance() bool`

HasBalance returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


