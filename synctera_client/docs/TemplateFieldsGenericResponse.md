# TemplateFieldsGenericResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountType** | [**AccountType**](AccountType.md) |  | 
**BalanceCeiling** | Pointer to [**BalanceCeiling**](BalanceCeiling.md) |  | [optional] 
**BalanceFloor** | Pointer to [**BalanceFloor**](BalanceFloor.md) |  | [optional] 
**BankCountry** | **string** | Bank country of the account | 
**BillingPeriod** | Pointer to [**BillingPeriod**](BillingPeriod.md) |  | [optional] 
**ChargeoffPeriod** | Pointer to **int32** | The number of days an account can stay delinquent before marking an account as charged-off.  | [optional] [default to 90]
**Currency** | **string** | Account currency. ISO 4217 alphabetic currency code | 
**DelinquencyPeriod** | Pointer to **int32** | The number of days past the due date to wait for a minimum payment before marking an account as delinquent.  | [optional] [default to 30]
**FeeProductIds** | Pointer to **[]string** | A list of fee account products that the current account associates with. | [optional] 
**GracePeriod** | Pointer to **int32** | The number of days past the billing period to allow for payment before it is considered due. This directly infers the due date for a payment.  | [optional] [default to 30]
**InterestProductId** | Pointer to **string** | An interest account product that the current account associates with. | [optional] 
**IsAchEnabled** | Pointer to **bool** | Enable ACH transaction on ledger. | [optional] [default to false]
**IsCardEnabled** | Pointer to **bool** | Enable card transaction on ledger. | [optional] [default to false]
**IsP2pEnabled** | Pointer to **bool** | Enable P2P transaction on ledger. | [optional] [default to false]
**IsWireEnabled** | Pointer to **bool** | Enable wire transaction on ledger. | [optional] [default to false]
**MinimumPayment** | Pointer to [**MinimumPayment**](MinimumPayment.md) |  | [optional] 
**OverdraftLimit** | Pointer to **int64** | Account&#39;s overdraft limit. Default is 0. Unit in cents. | [optional] 
**SpendingLimits** | Pointer to [**SpendingLimits**](SpendingLimits.md) |  | [optional] 

## Methods

### NewTemplateFieldsGenericResponse

`func NewTemplateFieldsGenericResponse(accountType AccountType, bankCountry string, currency string, ) *TemplateFieldsGenericResponse`

NewTemplateFieldsGenericResponse instantiates a new TemplateFieldsGenericResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTemplateFieldsGenericResponseWithDefaults

`func NewTemplateFieldsGenericResponseWithDefaults() *TemplateFieldsGenericResponse`

NewTemplateFieldsGenericResponseWithDefaults instantiates a new TemplateFieldsGenericResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountType

`func (o *TemplateFieldsGenericResponse) GetAccountType() AccountType`

GetAccountType returns the AccountType field if non-nil, zero value otherwise.

### GetAccountTypeOk

`func (o *TemplateFieldsGenericResponse) GetAccountTypeOk() (*AccountType, bool)`

GetAccountTypeOk returns a tuple with the AccountType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountType

`func (o *TemplateFieldsGenericResponse) SetAccountType(v AccountType)`

SetAccountType sets AccountType field to given value.


### GetBalanceCeiling

`func (o *TemplateFieldsGenericResponse) GetBalanceCeiling() BalanceCeiling`

GetBalanceCeiling returns the BalanceCeiling field if non-nil, zero value otherwise.

### GetBalanceCeilingOk

`func (o *TemplateFieldsGenericResponse) GetBalanceCeilingOk() (*BalanceCeiling, bool)`

GetBalanceCeilingOk returns a tuple with the BalanceCeiling field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalanceCeiling

`func (o *TemplateFieldsGenericResponse) SetBalanceCeiling(v BalanceCeiling)`

SetBalanceCeiling sets BalanceCeiling field to given value.

### HasBalanceCeiling

`func (o *TemplateFieldsGenericResponse) HasBalanceCeiling() bool`

HasBalanceCeiling returns a boolean if a field has been set.

### GetBalanceFloor

`func (o *TemplateFieldsGenericResponse) GetBalanceFloor() BalanceFloor`

GetBalanceFloor returns the BalanceFloor field if non-nil, zero value otherwise.

### GetBalanceFloorOk

`func (o *TemplateFieldsGenericResponse) GetBalanceFloorOk() (*BalanceFloor, bool)`

GetBalanceFloorOk returns a tuple with the BalanceFloor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalanceFloor

`func (o *TemplateFieldsGenericResponse) SetBalanceFloor(v BalanceFloor)`

SetBalanceFloor sets BalanceFloor field to given value.

### HasBalanceFloor

`func (o *TemplateFieldsGenericResponse) HasBalanceFloor() bool`

HasBalanceFloor returns a boolean if a field has been set.

### GetBankCountry

`func (o *TemplateFieldsGenericResponse) GetBankCountry() string`

GetBankCountry returns the BankCountry field if non-nil, zero value otherwise.

### GetBankCountryOk

`func (o *TemplateFieldsGenericResponse) GetBankCountryOk() (*string, bool)`

GetBankCountryOk returns a tuple with the BankCountry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBankCountry

`func (o *TemplateFieldsGenericResponse) SetBankCountry(v string)`

SetBankCountry sets BankCountry field to given value.


### GetBillingPeriod

`func (o *TemplateFieldsGenericResponse) GetBillingPeriod() BillingPeriod`

GetBillingPeriod returns the BillingPeriod field if non-nil, zero value otherwise.

### GetBillingPeriodOk

`func (o *TemplateFieldsGenericResponse) GetBillingPeriodOk() (*BillingPeriod, bool)`

GetBillingPeriodOk returns a tuple with the BillingPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingPeriod

`func (o *TemplateFieldsGenericResponse) SetBillingPeriod(v BillingPeriod)`

SetBillingPeriod sets BillingPeriod field to given value.

### HasBillingPeriod

`func (o *TemplateFieldsGenericResponse) HasBillingPeriod() bool`

HasBillingPeriod returns a boolean if a field has been set.

### GetChargeoffPeriod

`func (o *TemplateFieldsGenericResponse) GetChargeoffPeriod() int32`

GetChargeoffPeriod returns the ChargeoffPeriod field if non-nil, zero value otherwise.

### GetChargeoffPeriodOk

`func (o *TemplateFieldsGenericResponse) GetChargeoffPeriodOk() (*int32, bool)`

GetChargeoffPeriodOk returns a tuple with the ChargeoffPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChargeoffPeriod

`func (o *TemplateFieldsGenericResponse) SetChargeoffPeriod(v int32)`

SetChargeoffPeriod sets ChargeoffPeriod field to given value.

### HasChargeoffPeriod

`func (o *TemplateFieldsGenericResponse) HasChargeoffPeriod() bool`

HasChargeoffPeriod returns a boolean if a field has been set.

### GetCurrency

`func (o *TemplateFieldsGenericResponse) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *TemplateFieldsGenericResponse) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *TemplateFieldsGenericResponse) SetCurrency(v string)`

SetCurrency sets Currency field to given value.


### GetDelinquencyPeriod

`func (o *TemplateFieldsGenericResponse) GetDelinquencyPeriod() int32`

GetDelinquencyPeriod returns the DelinquencyPeriod field if non-nil, zero value otherwise.

### GetDelinquencyPeriodOk

`func (o *TemplateFieldsGenericResponse) GetDelinquencyPeriodOk() (*int32, bool)`

GetDelinquencyPeriodOk returns a tuple with the DelinquencyPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelinquencyPeriod

`func (o *TemplateFieldsGenericResponse) SetDelinquencyPeriod(v int32)`

SetDelinquencyPeriod sets DelinquencyPeriod field to given value.

### HasDelinquencyPeriod

`func (o *TemplateFieldsGenericResponse) HasDelinquencyPeriod() bool`

HasDelinquencyPeriod returns a boolean if a field has been set.

### GetFeeProductIds

`func (o *TemplateFieldsGenericResponse) GetFeeProductIds() []string`

GetFeeProductIds returns the FeeProductIds field if non-nil, zero value otherwise.

### GetFeeProductIdsOk

`func (o *TemplateFieldsGenericResponse) GetFeeProductIdsOk() (*[]string, bool)`

GetFeeProductIdsOk returns a tuple with the FeeProductIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeeProductIds

`func (o *TemplateFieldsGenericResponse) SetFeeProductIds(v []string)`

SetFeeProductIds sets FeeProductIds field to given value.

### HasFeeProductIds

`func (o *TemplateFieldsGenericResponse) HasFeeProductIds() bool`

HasFeeProductIds returns a boolean if a field has been set.

### GetGracePeriod

`func (o *TemplateFieldsGenericResponse) GetGracePeriod() int32`

GetGracePeriod returns the GracePeriod field if non-nil, zero value otherwise.

### GetGracePeriodOk

`func (o *TemplateFieldsGenericResponse) GetGracePeriodOk() (*int32, bool)`

GetGracePeriodOk returns a tuple with the GracePeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGracePeriod

`func (o *TemplateFieldsGenericResponse) SetGracePeriod(v int32)`

SetGracePeriod sets GracePeriod field to given value.

### HasGracePeriod

`func (o *TemplateFieldsGenericResponse) HasGracePeriod() bool`

HasGracePeriod returns a boolean if a field has been set.

### GetInterestProductId

`func (o *TemplateFieldsGenericResponse) GetInterestProductId() string`

GetInterestProductId returns the InterestProductId field if non-nil, zero value otherwise.

### GetInterestProductIdOk

`func (o *TemplateFieldsGenericResponse) GetInterestProductIdOk() (*string, bool)`

GetInterestProductIdOk returns a tuple with the InterestProductId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterestProductId

`func (o *TemplateFieldsGenericResponse) SetInterestProductId(v string)`

SetInterestProductId sets InterestProductId field to given value.

### HasInterestProductId

`func (o *TemplateFieldsGenericResponse) HasInterestProductId() bool`

HasInterestProductId returns a boolean if a field has been set.

### GetIsAchEnabled

`func (o *TemplateFieldsGenericResponse) GetIsAchEnabled() bool`

GetIsAchEnabled returns the IsAchEnabled field if non-nil, zero value otherwise.

### GetIsAchEnabledOk

`func (o *TemplateFieldsGenericResponse) GetIsAchEnabledOk() (*bool, bool)`

GetIsAchEnabledOk returns a tuple with the IsAchEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAchEnabled

`func (o *TemplateFieldsGenericResponse) SetIsAchEnabled(v bool)`

SetIsAchEnabled sets IsAchEnabled field to given value.

### HasIsAchEnabled

`func (o *TemplateFieldsGenericResponse) HasIsAchEnabled() bool`

HasIsAchEnabled returns a boolean if a field has been set.

### GetIsCardEnabled

`func (o *TemplateFieldsGenericResponse) GetIsCardEnabled() bool`

GetIsCardEnabled returns the IsCardEnabled field if non-nil, zero value otherwise.

### GetIsCardEnabledOk

`func (o *TemplateFieldsGenericResponse) GetIsCardEnabledOk() (*bool, bool)`

GetIsCardEnabledOk returns a tuple with the IsCardEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsCardEnabled

`func (o *TemplateFieldsGenericResponse) SetIsCardEnabled(v bool)`

SetIsCardEnabled sets IsCardEnabled field to given value.

### HasIsCardEnabled

`func (o *TemplateFieldsGenericResponse) HasIsCardEnabled() bool`

HasIsCardEnabled returns a boolean if a field has been set.

### GetIsP2pEnabled

`func (o *TemplateFieldsGenericResponse) GetIsP2pEnabled() bool`

GetIsP2pEnabled returns the IsP2pEnabled field if non-nil, zero value otherwise.

### GetIsP2pEnabledOk

`func (o *TemplateFieldsGenericResponse) GetIsP2pEnabledOk() (*bool, bool)`

GetIsP2pEnabledOk returns a tuple with the IsP2pEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsP2pEnabled

`func (o *TemplateFieldsGenericResponse) SetIsP2pEnabled(v bool)`

SetIsP2pEnabled sets IsP2pEnabled field to given value.

### HasIsP2pEnabled

`func (o *TemplateFieldsGenericResponse) HasIsP2pEnabled() bool`

HasIsP2pEnabled returns a boolean if a field has been set.

### GetIsWireEnabled

`func (o *TemplateFieldsGenericResponse) GetIsWireEnabled() bool`

GetIsWireEnabled returns the IsWireEnabled field if non-nil, zero value otherwise.

### GetIsWireEnabledOk

`func (o *TemplateFieldsGenericResponse) GetIsWireEnabledOk() (*bool, bool)`

GetIsWireEnabledOk returns a tuple with the IsWireEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsWireEnabled

`func (o *TemplateFieldsGenericResponse) SetIsWireEnabled(v bool)`

SetIsWireEnabled sets IsWireEnabled field to given value.

### HasIsWireEnabled

`func (o *TemplateFieldsGenericResponse) HasIsWireEnabled() bool`

HasIsWireEnabled returns a boolean if a field has been set.

### GetMinimumPayment

`func (o *TemplateFieldsGenericResponse) GetMinimumPayment() MinimumPayment`

GetMinimumPayment returns the MinimumPayment field if non-nil, zero value otherwise.

### GetMinimumPaymentOk

`func (o *TemplateFieldsGenericResponse) GetMinimumPaymentOk() (*MinimumPayment, bool)`

GetMinimumPaymentOk returns a tuple with the MinimumPayment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinimumPayment

`func (o *TemplateFieldsGenericResponse) SetMinimumPayment(v MinimumPayment)`

SetMinimumPayment sets MinimumPayment field to given value.

### HasMinimumPayment

`func (o *TemplateFieldsGenericResponse) HasMinimumPayment() bool`

HasMinimumPayment returns a boolean if a field has been set.

### GetOverdraftLimit

`func (o *TemplateFieldsGenericResponse) GetOverdraftLimit() int64`

GetOverdraftLimit returns the OverdraftLimit field if non-nil, zero value otherwise.

### GetOverdraftLimitOk

`func (o *TemplateFieldsGenericResponse) GetOverdraftLimitOk() (*int64, bool)`

GetOverdraftLimitOk returns a tuple with the OverdraftLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverdraftLimit

`func (o *TemplateFieldsGenericResponse) SetOverdraftLimit(v int64)`

SetOverdraftLimit sets OverdraftLimit field to given value.

### HasOverdraftLimit

`func (o *TemplateFieldsGenericResponse) HasOverdraftLimit() bool`

HasOverdraftLimit returns a boolean if a field has been set.

### GetSpendingLimits

`func (o *TemplateFieldsGenericResponse) GetSpendingLimits() SpendingLimits`

GetSpendingLimits returns the SpendingLimits field if non-nil, zero value otherwise.

### GetSpendingLimitsOk

`func (o *TemplateFieldsGenericResponse) GetSpendingLimitsOk() (*SpendingLimits, bool)`

GetSpendingLimitsOk returns a tuple with the SpendingLimits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpendingLimits

`func (o *TemplateFieldsGenericResponse) SetSpendingLimits(v SpendingLimits)`

SetSpendingLimits sets SpendingLimits field to given value.

### HasSpendingLimits

`func (o *TemplateFieldsGenericResponse) HasSpendingLimits() bool`

HasSpendingLimits returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


