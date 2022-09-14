# AccountDepositoryAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BalanceCeiling** | Pointer to [**BalanceCeiling**](BalanceCeiling.md) |  | [optional] 
**BalanceFloor** | Pointer to [**BalanceFloor**](BalanceFloor.md) |  | [optional] 
**FeeProductIds** | Pointer to **[]string** | A list of fee account products that the current account associates with. | [optional] 
**InterestProductId** | Pointer to **string** | An interest account product that the current account associates with. The account product must have its calculation_method set to COMPOUNDED_MONTHLY.  | [optional] 
**OverdraftLimit** | Pointer to **int64** | Account&#39;s overdraft limit | [optional] 
**SpendingLimits** | Pointer to [**SpendingLimits**](SpendingLimits.md) |  | [optional] 

## Methods

### NewAccountDepositoryAllOf

`func NewAccountDepositoryAllOf() *AccountDepositoryAllOf`

NewAccountDepositoryAllOf instantiates a new AccountDepositoryAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountDepositoryAllOfWithDefaults

`func NewAccountDepositoryAllOfWithDefaults() *AccountDepositoryAllOf`

NewAccountDepositoryAllOfWithDefaults instantiates a new AccountDepositoryAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBalanceCeiling

`func (o *AccountDepositoryAllOf) GetBalanceCeiling() BalanceCeiling`

GetBalanceCeiling returns the BalanceCeiling field if non-nil, zero value otherwise.

### GetBalanceCeilingOk

`func (o *AccountDepositoryAllOf) GetBalanceCeilingOk() (*BalanceCeiling, bool)`

GetBalanceCeilingOk returns a tuple with the BalanceCeiling field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalanceCeiling

`func (o *AccountDepositoryAllOf) SetBalanceCeiling(v BalanceCeiling)`

SetBalanceCeiling sets BalanceCeiling field to given value.

### HasBalanceCeiling

`func (o *AccountDepositoryAllOf) HasBalanceCeiling() bool`

HasBalanceCeiling returns a boolean if a field has been set.

### GetBalanceFloor

`func (o *AccountDepositoryAllOf) GetBalanceFloor() BalanceFloor`

GetBalanceFloor returns the BalanceFloor field if non-nil, zero value otherwise.

### GetBalanceFloorOk

`func (o *AccountDepositoryAllOf) GetBalanceFloorOk() (*BalanceFloor, bool)`

GetBalanceFloorOk returns a tuple with the BalanceFloor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalanceFloor

`func (o *AccountDepositoryAllOf) SetBalanceFloor(v BalanceFloor)`

SetBalanceFloor sets BalanceFloor field to given value.

### HasBalanceFloor

`func (o *AccountDepositoryAllOf) HasBalanceFloor() bool`

HasBalanceFloor returns a boolean if a field has been set.

### GetFeeProductIds

`func (o *AccountDepositoryAllOf) GetFeeProductIds() []string`

GetFeeProductIds returns the FeeProductIds field if non-nil, zero value otherwise.

### GetFeeProductIdsOk

`func (o *AccountDepositoryAllOf) GetFeeProductIdsOk() (*[]string, bool)`

GetFeeProductIdsOk returns a tuple with the FeeProductIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeeProductIds

`func (o *AccountDepositoryAllOf) SetFeeProductIds(v []string)`

SetFeeProductIds sets FeeProductIds field to given value.

### HasFeeProductIds

`func (o *AccountDepositoryAllOf) HasFeeProductIds() bool`

HasFeeProductIds returns a boolean if a field has been set.

### GetInterestProductId

`func (o *AccountDepositoryAllOf) GetInterestProductId() string`

GetInterestProductId returns the InterestProductId field if non-nil, zero value otherwise.

### GetInterestProductIdOk

`func (o *AccountDepositoryAllOf) GetInterestProductIdOk() (*string, bool)`

GetInterestProductIdOk returns a tuple with the InterestProductId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterestProductId

`func (o *AccountDepositoryAllOf) SetInterestProductId(v string)`

SetInterestProductId sets InterestProductId field to given value.

### HasInterestProductId

`func (o *AccountDepositoryAllOf) HasInterestProductId() bool`

HasInterestProductId returns a boolean if a field has been set.

### GetOverdraftLimit

`func (o *AccountDepositoryAllOf) GetOverdraftLimit() int64`

GetOverdraftLimit returns the OverdraftLimit field if non-nil, zero value otherwise.

### GetOverdraftLimitOk

`func (o *AccountDepositoryAllOf) GetOverdraftLimitOk() (*int64, bool)`

GetOverdraftLimitOk returns a tuple with the OverdraftLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverdraftLimit

`func (o *AccountDepositoryAllOf) SetOverdraftLimit(v int64)`

SetOverdraftLimit sets OverdraftLimit field to given value.

### HasOverdraftLimit

`func (o *AccountDepositoryAllOf) HasOverdraftLimit() bool`

HasOverdraftLimit returns a boolean if a field has been set.

### GetSpendingLimits

`func (o *AccountDepositoryAllOf) GetSpendingLimits() SpendingLimits`

GetSpendingLimits returns the SpendingLimits field if non-nil, zero value otherwise.

### GetSpendingLimitsOk

`func (o *AccountDepositoryAllOf) GetSpendingLimitsOk() (*SpendingLimits, bool)`

GetSpendingLimitsOk returns a tuple with the SpendingLimits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpendingLimits

`func (o *AccountDepositoryAllOf) SetSpendingLimits(v SpendingLimits)`

SetSpendingLimits sets SpendingLimits field to given value.

### HasSpendingLimits

`func (o *AccountDepositoryAllOf) HasSpendingLimits() bool`

HasSpendingLimits returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


