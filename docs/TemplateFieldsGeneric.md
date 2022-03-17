# TemplateFieldsGeneric

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountType** | [**AccountType**](AccountType.md) |  | 
**BankCountry** | **string** | Bank country of the account | 
**Currency** | **string** | Account currency. ISO 4217 alphabetic currency code | 
**BalanceCeiling** | Pointer to [**BalanceCeiling**](BalanceCeiling.md) |  | [optional] 
**BalanceFloor** | Pointer to [**BalanceFloor**](BalanceFloor.md) |  | [optional] 
**FeeProductIds** | Pointer to **[]string** | A list of fee resources from account product that new accounts will associate with | [optional] 
**InterestProductId** | Pointer to **string** | Interest from account product that new accounts will associate with | [optional] 
**IsAchEnabled** | Pointer to **bool** | Enable ACH transaction on ledger. Default is false | [optional] 
**IsCardEnabled** | Pointer to **bool** | Enable card transaction on ledger. Default is false | [optional] 
**IsP2pEnabled** | Pointer to **bool** | Enable P2P transaction on ledger. Default is false | [optional] 
**OverdraftLimit** | Pointer to **int64** | Account&#39;s overdraft limit. Default is 0 | [optional] 
**SpendingLimits** | Pointer to [**SpendingLimits**](SpendingLimits.md) |  | [optional] 

## Methods

### NewTemplateFieldsGeneric

`func NewTemplateFieldsGeneric(accountType AccountType, bankCountry string, currency string, ) *TemplateFieldsGeneric`

NewTemplateFieldsGeneric instantiates a new TemplateFieldsGeneric object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTemplateFieldsGenericWithDefaults

`func NewTemplateFieldsGenericWithDefaults() *TemplateFieldsGeneric`

NewTemplateFieldsGenericWithDefaults instantiates a new TemplateFieldsGeneric object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountType

`func (o *TemplateFieldsGeneric) GetAccountType() AccountType`

GetAccountType returns the AccountType field if non-nil, zero value otherwise.

### GetAccountTypeOk

`func (o *TemplateFieldsGeneric) GetAccountTypeOk() (*AccountType, bool)`

GetAccountTypeOk returns a tuple with the AccountType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountType

`func (o *TemplateFieldsGeneric) SetAccountType(v AccountType)`

SetAccountType sets AccountType field to given value.


### GetBankCountry

`func (o *TemplateFieldsGeneric) GetBankCountry() string`

GetBankCountry returns the BankCountry field if non-nil, zero value otherwise.

### GetBankCountryOk

`func (o *TemplateFieldsGeneric) GetBankCountryOk() (*string, bool)`

GetBankCountryOk returns a tuple with the BankCountry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBankCountry

`func (o *TemplateFieldsGeneric) SetBankCountry(v string)`

SetBankCountry sets BankCountry field to given value.


### GetCurrency

`func (o *TemplateFieldsGeneric) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *TemplateFieldsGeneric) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *TemplateFieldsGeneric) SetCurrency(v string)`

SetCurrency sets Currency field to given value.


### GetBalanceCeiling

`func (o *TemplateFieldsGeneric) GetBalanceCeiling() BalanceCeiling`

GetBalanceCeiling returns the BalanceCeiling field if non-nil, zero value otherwise.

### GetBalanceCeilingOk

`func (o *TemplateFieldsGeneric) GetBalanceCeilingOk() (*BalanceCeiling, bool)`

GetBalanceCeilingOk returns a tuple with the BalanceCeiling field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalanceCeiling

`func (o *TemplateFieldsGeneric) SetBalanceCeiling(v BalanceCeiling)`

SetBalanceCeiling sets BalanceCeiling field to given value.

### HasBalanceCeiling

`func (o *TemplateFieldsGeneric) HasBalanceCeiling() bool`

HasBalanceCeiling returns a boolean if a field has been set.

### GetBalanceFloor

`func (o *TemplateFieldsGeneric) GetBalanceFloor() BalanceFloor`

GetBalanceFloor returns the BalanceFloor field if non-nil, zero value otherwise.

### GetBalanceFloorOk

`func (o *TemplateFieldsGeneric) GetBalanceFloorOk() (*BalanceFloor, bool)`

GetBalanceFloorOk returns a tuple with the BalanceFloor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalanceFloor

`func (o *TemplateFieldsGeneric) SetBalanceFloor(v BalanceFloor)`

SetBalanceFloor sets BalanceFloor field to given value.

### HasBalanceFloor

`func (o *TemplateFieldsGeneric) HasBalanceFloor() bool`

HasBalanceFloor returns a boolean if a field has been set.

### GetFeeProductIds

`func (o *TemplateFieldsGeneric) GetFeeProductIds() []string`

GetFeeProductIds returns the FeeProductIds field if non-nil, zero value otherwise.

### GetFeeProductIdsOk

`func (o *TemplateFieldsGeneric) GetFeeProductIdsOk() (*[]string, bool)`

GetFeeProductIdsOk returns a tuple with the FeeProductIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeeProductIds

`func (o *TemplateFieldsGeneric) SetFeeProductIds(v []string)`

SetFeeProductIds sets FeeProductIds field to given value.

### HasFeeProductIds

`func (o *TemplateFieldsGeneric) HasFeeProductIds() bool`

HasFeeProductIds returns a boolean if a field has been set.

### GetInterestProductId

`func (o *TemplateFieldsGeneric) GetInterestProductId() string`

GetInterestProductId returns the InterestProductId field if non-nil, zero value otherwise.

### GetInterestProductIdOk

`func (o *TemplateFieldsGeneric) GetInterestProductIdOk() (*string, bool)`

GetInterestProductIdOk returns a tuple with the InterestProductId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterestProductId

`func (o *TemplateFieldsGeneric) SetInterestProductId(v string)`

SetInterestProductId sets InterestProductId field to given value.

### HasInterestProductId

`func (o *TemplateFieldsGeneric) HasInterestProductId() bool`

HasInterestProductId returns a boolean if a field has been set.

### GetIsAchEnabled

`func (o *TemplateFieldsGeneric) GetIsAchEnabled() bool`

GetIsAchEnabled returns the IsAchEnabled field if non-nil, zero value otherwise.

### GetIsAchEnabledOk

`func (o *TemplateFieldsGeneric) GetIsAchEnabledOk() (*bool, bool)`

GetIsAchEnabledOk returns a tuple with the IsAchEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAchEnabled

`func (o *TemplateFieldsGeneric) SetIsAchEnabled(v bool)`

SetIsAchEnabled sets IsAchEnabled field to given value.

### HasIsAchEnabled

`func (o *TemplateFieldsGeneric) HasIsAchEnabled() bool`

HasIsAchEnabled returns a boolean if a field has been set.

### GetIsCardEnabled

`func (o *TemplateFieldsGeneric) GetIsCardEnabled() bool`

GetIsCardEnabled returns the IsCardEnabled field if non-nil, zero value otherwise.

### GetIsCardEnabledOk

`func (o *TemplateFieldsGeneric) GetIsCardEnabledOk() (*bool, bool)`

GetIsCardEnabledOk returns a tuple with the IsCardEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsCardEnabled

`func (o *TemplateFieldsGeneric) SetIsCardEnabled(v bool)`

SetIsCardEnabled sets IsCardEnabled field to given value.

### HasIsCardEnabled

`func (o *TemplateFieldsGeneric) HasIsCardEnabled() bool`

HasIsCardEnabled returns a boolean if a field has been set.

### GetIsP2pEnabled

`func (o *TemplateFieldsGeneric) GetIsP2pEnabled() bool`

GetIsP2pEnabled returns the IsP2pEnabled field if non-nil, zero value otherwise.

### GetIsP2pEnabledOk

`func (o *TemplateFieldsGeneric) GetIsP2pEnabledOk() (*bool, bool)`

GetIsP2pEnabledOk returns a tuple with the IsP2pEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsP2pEnabled

`func (o *TemplateFieldsGeneric) SetIsP2pEnabled(v bool)`

SetIsP2pEnabled sets IsP2pEnabled field to given value.

### HasIsP2pEnabled

`func (o *TemplateFieldsGeneric) HasIsP2pEnabled() bool`

HasIsP2pEnabled returns a boolean if a field has been set.

### GetOverdraftLimit

`func (o *TemplateFieldsGeneric) GetOverdraftLimit() int64`

GetOverdraftLimit returns the OverdraftLimit field if non-nil, zero value otherwise.

### GetOverdraftLimitOk

`func (o *TemplateFieldsGeneric) GetOverdraftLimitOk() (*int64, bool)`

GetOverdraftLimitOk returns a tuple with the OverdraftLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverdraftLimit

`func (o *TemplateFieldsGeneric) SetOverdraftLimit(v int64)`

SetOverdraftLimit sets OverdraftLimit field to given value.

### HasOverdraftLimit

`func (o *TemplateFieldsGeneric) HasOverdraftLimit() bool`

HasOverdraftLimit returns a boolean if a field has been set.

### GetSpendingLimits

`func (o *TemplateFieldsGeneric) GetSpendingLimits() SpendingLimits`

GetSpendingLimits returns the SpendingLimits field if non-nil, zero value otherwise.

### GetSpendingLimitsOk

`func (o *TemplateFieldsGeneric) GetSpendingLimitsOk() (*SpendingLimits, bool)`

GetSpendingLimitsOk returns a tuple with the SpendingLimits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpendingLimits

`func (o *TemplateFieldsGeneric) SetSpendingLimits(v SpendingLimits)`

SetSpendingLimits sets SpendingLimits field to given value.

### HasSpendingLimits

`func (o *TemplateFieldsGeneric) HasSpendingLimits() bool`

HasSpendingLimits returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


