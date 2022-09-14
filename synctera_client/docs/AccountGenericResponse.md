# AccountGenericResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessStatus** | Pointer to [**AccountAccessStatus**](AccountAccessStatus.md) |  | [optional] 
**AccessStatusLastUpdatedTime** | Pointer to **time.Time** | Timestamp of the last modification of the access_status. RFC3339 format. | [optional] [readonly] 
**AccountNumber** | Pointer to **string** | Account number | [optional] [readonly] 
**AccountNumberMasked** | Pointer to **string** | The response will contain the bank fintech ID (3 or 6 digits) plus the last 4 digits, with the digits in between replaced with * characters. Shadow mode account numbers will not be masked. | [optional] [readonly] 
**AccountPurpose** | Pointer to **string** | Purpose of the account | [optional] 
**AccountType** | Pointer to [**AccountType**](AccountType.md) |  | [optional] 
**ApplicationId** | Pointer to **string** | The application ID for this account.  | [optional] 
**BalanceCeiling** | Pointer to [**BalanceCeiling**](BalanceCeiling.md) |  | [optional] 
**BalanceFloor** | Pointer to [**BalanceFloor**](BalanceFloor.md) |  | [optional] 
**Balances** | Pointer to [**[]Balance**](Balance.md) | A list of balances for account based on different type | [optional] [readonly] 
**BankRouting** | Pointer to **string** | Bank routing number | [optional] [readonly] 
**BillingPeriod** | Pointer to [**BillingPeriod**](BillingPeriod.md) |  | [optional] 
**BusinessIds** | Pointer to **[]string** | A list of the business IDs of the account holders. | [optional] [readonly] 
**ChargeoffPeriod** | Pointer to **int32** | The number of days an account can stay delinquent before marking an account as charged-off.  | [optional] [default to 90]
**CreationTime** | Pointer to **time.Time** | Account creation timestamp in RFC3339 format | [optional] [readonly] 
**CreditLimit** | Pointer to **int64** | The credit limit for this line of credit account in cents. Minimum is 0.  | [optional] 
**Currency** | Pointer to **string** | Account currency or account settlement currency. ISO 4217 alphabetic currency code. Default USD | [optional] 
**CustomerIds** | Pointer to **[]string** | A list of the customer IDs of the account holders. | [optional] [readonly] 
**CustomerType** | Pointer to [**CustomerType**](CustomerType.md) |  | [optional] 
**DelinquencyPeriod** | Pointer to **int32** | The number of days past the due date to wait for a minimum payment before marking an account as delinquent.  | [optional] [default to 30]
**ExchangeRateType** | Pointer to **string** | Exchange rate type | [optional] 
**FeeProductIds** | Pointer to **[]string** | A list of fee account products that the current account associates with. | [optional] 
**GracePeriod** | Pointer to **int32** | The number of days past the billing period to allow for payment before it is considered due. This directly infers the due date for a payment.  | [optional] 
**Iban** | Pointer to **string** | International bank account number | [optional] 
**Id** | Pointer to **string** | Account ID | [optional] [readonly] 
**InterestProductId** | Pointer to **string** | An interest account product that the current account associates with. | [optional] 
**IsAccountPool** | Pointer to **bool** | Account is investment (variable balance) account or a multi-balance account pool. Default false | [optional] 
**IsAchEnabled** | Pointer to **bool** | A flag to indicate whether ACH transactions are enabled. | [optional] [readonly] 
**IsCardEnabled** | Pointer to **bool** | A flag to indicate whether card transactions are enabled. | [optional] [readonly] 
**IsP2pEnabled** | Pointer to **bool** | A flag to indicate whether P2P transactions are enabled. | [optional] [readonly] 
**IsWireEnabled** | Pointer to **bool** | A flag to indicate whether wire transactions are enabled. | [optional] [readonly] 
**LastUpdatedTime** | Pointer to **time.Time** | Timestamp of the last account modification in RFC3339 format | [optional] [readonly] 
**Metadata** | Pointer to **map[string]interface{}** | User provided account metadata | [optional] 
**MinimumPayment** | Pointer to [**MinimumPayment**](MinimumPayment.md) |  | [optional] 
**Nickname** | Pointer to **string** | User provided account nickname | [optional] 
**OverdraftLimit** | Pointer to **int64** | Account&#39;s overdraft limit | [optional] 
**SpendingLimits** | Pointer to [**SpendingLimits**](SpendingLimits.md) |  | [optional] 
**Status** | Pointer to [**Status**](Status.md) |  | [optional] 
**SwiftCode** | Pointer to **string** | SWIFT code | [optional] 

## Methods

### NewAccountGenericResponse

`func NewAccountGenericResponse() *AccountGenericResponse`

NewAccountGenericResponse instantiates a new AccountGenericResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountGenericResponseWithDefaults

`func NewAccountGenericResponseWithDefaults() *AccountGenericResponse`

NewAccountGenericResponseWithDefaults instantiates a new AccountGenericResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessStatus

`func (o *AccountGenericResponse) GetAccessStatus() AccountAccessStatus`

GetAccessStatus returns the AccessStatus field if non-nil, zero value otherwise.

### GetAccessStatusOk

`func (o *AccountGenericResponse) GetAccessStatusOk() (*AccountAccessStatus, bool)`

GetAccessStatusOk returns a tuple with the AccessStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessStatus

`func (o *AccountGenericResponse) SetAccessStatus(v AccountAccessStatus)`

SetAccessStatus sets AccessStatus field to given value.

### HasAccessStatus

`func (o *AccountGenericResponse) HasAccessStatus() bool`

HasAccessStatus returns a boolean if a field has been set.

### GetAccessStatusLastUpdatedTime

`func (o *AccountGenericResponse) GetAccessStatusLastUpdatedTime() time.Time`

GetAccessStatusLastUpdatedTime returns the AccessStatusLastUpdatedTime field if non-nil, zero value otherwise.

### GetAccessStatusLastUpdatedTimeOk

`func (o *AccountGenericResponse) GetAccessStatusLastUpdatedTimeOk() (*time.Time, bool)`

GetAccessStatusLastUpdatedTimeOk returns a tuple with the AccessStatusLastUpdatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessStatusLastUpdatedTime

`func (o *AccountGenericResponse) SetAccessStatusLastUpdatedTime(v time.Time)`

SetAccessStatusLastUpdatedTime sets AccessStatusLastUpdatedTime field to given value.

### HasAccessStatusLastUpdatedTime

`func (o *AccountGenericResponse) HasAccessStatusLastUpdatedTime() bool`

HasAccessStatusLastUpdatedTime returns a boolean if a field has been set.

### GetAccountNumber

`func (o *AccountGenericResponse) GetAccountNumber() string`

GetAccountNumber returns the AccountNumber field if non-nil, zero value otherwise.

### GetAccountNumberOk

`func (o *AccountGenericResponse) GetAccountNumberOk() (*string, bool)`

GetAccountNumberOk returns a tuple with the AccountNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountNumber

`func (o *AccountGenericResponse) SetAccountNumber(v string)`

SetAccountNumber sets AccountNumber field to given value.

### HasAccountNumber

`func (o *AccountGenericResponse) HasAccountNumber() bool`

HasAccountNumber returns a boolean if a field has been set.

### GetAccountNumberMasked

`func (o *AccountGenericResponse) GetAccountNumberMasked() string`

GetAccountNumberMasked returns the AccountNumberMasked field if non-nil, zero value otherwise.

### GetAccountNumberMaskedOk

`func (o *AccountGenericResponse) GetAccountNumberMaskedOk() (*string, bool)`

GetAccountNumberMaskedOk returns a tuple with the AccountNumberMasked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountNumberMasked

`func (o *AccountGenericResponse) SetAccountNumberMasked(v string)`

SetAccountNumberMasked sets AccountNumberMasked field to given value.

### HasAccountNumberMasked

`func (o *AccountGenericResponse) HasAccountNumberMasked() bool`

HasAccountNumberMasked returns a boolean if a field has been set.

### GetAccountPurpose

`func (o *AccountGenericResponse) GetAccountPurpose() string`

GetAccountPurpose returns the AccountPurpose field if non-nil, zero value otherwise.

### GetAccountPurposeOk

`func (o *AccountGenericResponse) GetAccountPurposeOk() (*string, bool)`

GetAccountPurposeOk returns a tuple with the AccountPurpose field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountPurpose

`func (o *AccountGenericResponse) SetAccountPurpose(v string)`

SetAccountPurpose sets AccountPurpose field to given value.

### HasAccountPurpose

`func (o *AccountGenericResponse) HasAccountPurpose() bool`

HasAccountPurpose returns a boolean if a field has been set.

### GetAccountType

`func (o *AccountGenericResponse) GetAccountType() AccountType`

GetAccountType returns the AccountType field if non-nil, zero value otherwise.

### GetAccountTypeOk

`func (o *AccountGenericResponse) GetAccountTypeOk() (*AccountType, bool)`

GetAccountTypeOk returns a tuple with the AccountType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountType

`func (o *AccountGenericResponse) SetAccountType(v AccountType)`

SetAccountType sets AccountType field to given value.

### HasAccountType

`func (o *AccountGenericResponse) HasAccountType() bool`

HasAccountType returns a boolean if a field has been set.

### GetApplicationId

`func (o *AccountGenericResponse) GetApplicationId() string`

GetApplicationId returns the ApplicationId field if non-nil, zero value otherwise.

### GetApplicationIdOk

`func (o *AccountGenericResponse) GetApplicationIdOk() (*string, bool)`

GetApplicationIdOk returns a tuple with the ApplicationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationId

`func (o *AccountGenericResponse) SetApplicationId(v string)`

SetApplicationId sets ApplicationId field to given value.

### HasApplicationId

`func (o *AccountGenericResponse) HasApplicationId() bool`

HasApplicationId returns a boolean if a field has been set.

### GetBalanceCeiling

`func (o *AccountGenericResponse) GetBalanceCeiling() BalanceCeiling`

GetBalanceCeiling returns the BalanceCeiling field if non-nil, zero value otherwise.

### GetBalanceCeilingOk

`func (o *AccountGenericResponse) GetBalanceCeilingOk() (*BalanceCeiling, bool)`

GetBalanceCeilingOk returns a tuple with the BalanceCeiling field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalanceCeiling

`func (o *AccountGenericResponse) SetBalanceCeiling(v BalanceCeiling)`

SetBalanceCeiling sets BalanceCeiling field to given value.

### HasBalanceCeiling

`func (o *AccountGenericResponse) HasBalanceCeiling() bool`

HasBalanceCeiling returns a boolean if a field has been set.

### GetBalanceFloor

`func (o *AccountGenericResponse) GetBalanceFloor() BalanceFloor`

GetBalanceFloor returns the BalanceFloor field if non-nil, zero value otherwise.

### GetBalanceFloorOk

`func (o *AccountGenericResponse) GetBalanceFloorOk() (*BalanceFloor, bool)`

GetBalanceFloorOk returns a tuple with the BalanceFloor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalanceFloor

`func (o *AccountGenericResponse) SetBalanceFloor(v BalanceFloor)`

SetBalanceFloor sets BalanceFloor field to given value.

### HasBalanceFloor

`func (o *AccountGenericResponse) HasBalanceFloor() bool`

HasBalanceFloor returns a boolean if a field has been set.

### GetBalances

`func (o *AccountGenericResponse) GetBalances() []Balance`

GetBalances returns the Balances field if non-nil, zero value otherwise.

### GetBalancesOk

`func (o *AccountGenericResponse) GetBalancesOk() (*[]Balance, bool)`

GetBalancesOk returns a tuple with the Balances field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalances

`func (o *AccountGenericResponse) SetBalances(v []Balance)`

SetBalances sets Balances field to given value.

### HasBalances

`func (o *AccountGenericResponse) HasBalances() bool`

HasBalances returns a boolean if a field has been set.

### GetBankRouting

`func (o *AccountGenericResponse) GetBankRouting() string`

GetBankRouting returns the BankRouting field if non-nil, zero value otherwise.

### GetBankRoutingOk

`func (o *AccountGenericResponse) GetBankRoutingOk() (*string, bool)`

GetBankRoutingOk returns a tuple with the BankRouting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBankRouting

`func (o *AccountGenericResponse) SetBankRouting(v string)`

SetBankRouting sets BankRouting field to given value.

### HasBankRouting

`func (o *AccountGenericResponse) HasBankRouting() bool`

HasBankRouting returns a boolean if a field has been set.

### GetBillingPeriod

`func (o *AccountGenericResponse) GetBillingPeriod() BillingPeriod`

GetBillingPeriod returns the BillingPeriod field if non-nil, zero value otherwise.

### GetBillingPeriodOk

`func (o *AccountGenericResponse) GetBillingPeriodOk() (*BillingPeriod, bool)`

GetBillingPeriodOk returns a tuple with the BillingPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingPeriod

`func (o *AccountGenericResponse) SetBillingPeriod(v BillingPeriod)`

SetBillingPeriod sets BillingPeriod field to given value.

### HasBillingPeriod

`func (o *AccountGenericResponse) HasBillingPeriod() bool`

HasBillingPeriod returns a boolean if a field has been set.

### GetBusinessIds

`func (o *AccountGenericResponse) GetBusinessIds() []string`

GetBusinessIds returns the BusinessIds field if non-nil, zero value otherwise.

### GetBusinessIdsOk

`func (o *AccountGenericResponse) GetBusinessIdsOk() (*[]string, bool)`

GetBusinessIdsOk returns a tuple with the BusinessIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessIds

`func (o *AccountGenericResponse) SetBusinessIds(v []string)`

SetBusinessIds sets BusinessIds field to given value.

### HasBusinessIds

`func (o *AccountGenericResponse) HasBusinessIds() bool`

HasBusinessIds returns a boolean if a field has been set.

### GetChargeoffPeriod

`func (o *AccountGenericResponse) GetChargeoffPeriod() int32`

GetChargeoffPeriod returns the ChargeoffPeriod field if non-nil, zero value otherwise.

### GetChargeoffPeriodOk

`func (o *AccountGenericResponse) GetChargeoffPeriodOk() (*int32, bool)`

GetChargeoffPeriodOk returns a tuple with the ChargeoffPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChargeoffPeriod

`func (o *AccountGenericResponse) SetChargeoffPeriod(v int32)`

SetChargeoffPeriod sets ChargeoffPeriod field to given value.

### HasChargeoffPeriod

`func (o *AccountGenericResponse) HasChargeoffPeriod() bool`

HasChargeoffPeriod returns a boolean if a field has been set.

### GetCreationTime

`func (o *AccountGenericResponse) GetCreationTime() time.Time`

GetCreationTime returns the CreationTime field if non-nil, zero value otherwise.

### GetCreationTimeOk

`func (o *AccountGenericResponse) GetCreationTimeOk() (*time.Time, bool)`

GetCreationTimeOk returns a tuple with the CreationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationTime

`func (o *AccountGenericResponse) SetCreationTime(v time.Time)`

SetCreationTime sets CreationTime field to given value.

### HasCreationTime

`func (o *AccountGenericResponse) HasCreationTime() bool`

HasCreationTime returns a boolean if a field has been set.

### GetCreditLimit

`func (o *AccountGenericResponse) GetCreditLimit() int64`

GetCreditLimit returns the CreditLimit field if non-nil, zero value otherwise.

### GetCreditLimitOk

`func (o *AccountGenericResponse) GetCreditLimitOk() (*int64, bool)`

GetCreditLimitOk returns a tuple with the CreditLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreditLimit

`func (o *AccountGenericResponse) SetCreditLimit(v int64)`

SetCreditLimit sets CreditLimit field to given value.

### HasCreditLimit

`func (o *AccountGenericResponse) HasCreditLimit() bool`

HasCreditLimit returns a boolean if a field has been set.

### GetCurrency

`func (o *AccountGenericResponse) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *AccountGenericResponse) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *AccountGenericResponse) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *AccountGenericResponse) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetCustomerIds

`func (o *AccountGenericResponse) GetCustomerIds() []string`

GetCustomerIds returns the CustomerIds field if non-nil, zero value otherwise.

### GetCustomerIdsOk

`func (o *AccountGenericResponse) GetCustomerIdsOk() (*[]string, bool)`

GetCustomerIdsOk returns a tuple with the CustomerIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerIds

`func (o *AccountGenericResponse) SetCustomerIds(v []string)`

SetCustomerIds sets CustomerIds field to given value.

### HasCustomerIds

`func (o *AccountGenericResponse) HasCustomerIds() bool`

HasCustomerIds returns a boolean if a field has been set.

### GetCustomerType

`func (o *AccountGenericResponse) GetCustomerType() CustomerType`

GetCustomerType returns the CustomerType field if non-nil, zero value otherwise.

### GetCustomerTypeOk

`func (o *AccountGenericResponse) GetCustomerTypeOk() (*CustomerType, bool)`

GetCustomerTypeOk returns a tuple with the CustomerType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerType

`func (o *AccountGenericResponse) SetCustomerType(v CustomerType)`

SetCustomerType sets CustomerType field to given value.

### HasCustomerType

`func (o *AccountGenericResponse) HasCustomerType() bool`

HasCustomerType returns a boolean if a field has been set.

### GetDelinquencyPeriod

`func (o *AccountGenericResponse) GetDelinquencyPeriod() int32`

GetDelinquencyPeriod returns the DelinquencyPeriod field if non-nil, zero value otherwise.

### GetDelinquencyPeriodOk

`func (o *AccountGenericResponse) GetDelinquencyPeriodOk() (*int32, bool)`

GetDelinquencyPeriodOk returns a tuple with the DelinquencyPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelinquencyPeriod

`func (o *AccountGenericResponse) SetDelinquencyPeriod(v int32)`

SetDelinquencyPeriod sets DelinquencyPeriod field to given value.

### HasDelinquencyPeriod

`func (o *AccountGenericResponse) HasDelinquencyPeriod() bool`

HasDelinquencyPeriod returns a boolean if a field has been set.

### GetExchangeRateType

`func (o *AccountGenericResponse) GetExchangeRateType() string`

GetExchangeRateType returns the ExchangeRateType field if non-nil, zero value otherwise.

### GetExchangeRateTypeOk

`func (o *AccountGenericResponse) GetExchangeRateTypeOk() (*string, bool)`

GetExchangeRateTypeOk returns a tuple with the ExchangeRateType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExchangeRateType

`func (o *AccountGenericResponse) SetExchangeRateType(v string)`

SetExchangeRateType sets ExchangeRateType field to given value.

### HasExchangeRateType

`func (o *AccountGenericResponse) HasExchangeRateType() bool`

HasExchangeRateType returns a boolean if a field has been set.

### GetFeeProductIds

`func (o *AccountGenericResponse) GetFeeProductIds() []string`

GetFeeProductIds returns the FeeProductIds field if non-nil, zero value otherwise.

### GetFeeProductIdsOk

`func (o *AccountGenericResponse) GetFeeProductIdsOk() (*[]string, bool)`

GetFeeProductIdsOk returns a tuple with the FeeProductIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeeProductIds

`func (o *AccountGenericResponse) SetFeeProductIds(v []string)`

SetFeeProductIds sets FeeProductIds field to given value.

### HasFeeProductIds

`func (o *AccountGenericResponse) HasFeeProductIds() bool`

HasFeeProductIds returns a boolean if a field has been set.

### GetGracePeriod

`func (o *AccountGenericResponse) GetGracePeriod() int32`

GetGracePeriod returns the GracePeriod field if non-nil, zero value otherwise.

### GetGracePeriodOk

`func (o *AccountGenericResponse) GetGracePeriodOk() (*int32, bool)`

GetGracePeriodOk returns a tuple with the GracePeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGracePeriod

`func (o *AccountGenericResponse) SetGracePeriod(v int32)`

SetGracePeriod sets GracePeriod field to given value.

### HasGracePeriod

`func (o *AccountGenericResponse) HasGracePeriod() bool`

HasGracePeriod returns a boolean if a field has been set.

### GetIban

`func (o *AccountGenericResponse) GetIban() string`

GetIban returns the Iban field if non-nil, zero value otherwise.

### GetIbanOk

`func (o *AccountGenericResponse) GetIbanOk() (*string, bool)`

GetIbanOk returns a tuple with the Iban field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIban

`func (o *AccountGenericResponse) SetIban(v string)`

SetIban sets Iban field to given value.

### HasIban

`func (o *AccountGenericResponse) HasIban() bool`

HasIban returns a boolean if a field has been set.

### GetId

`func (o *AccountGenericResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AccountGenericResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AccountGenericResponse) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *AccountGenericResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### GetInterestProductId

`func (o *AccountGenericResponse) GetInterestProductId() string`

GetInterestProductId returns the InterestProductId field if non-nil, zero value otherwise.

### GetInterestProductIdOk

`func (o *AccountGenericResponse) GetInterestProductIdOk() (*string, bool)`

GetInterestProductIdOk returns a tuple with the InterestProductId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterestProductId

`func (o *AccountGenericResponse) SetInterestProductId(v string)`

SetInterestProductId sets InterestProductId field to given value.

### HasInterestProductId

`func (o *AccountGenericResponse) HasInterestProductId() bool`

HasInterestProductId returns a boolean if a field has been set.

### GetIsAccountPool

`func (o *AccountGenericResponse) GetIsAccountPool() bool`

GetIsAccountPool returns the IsAccountPool field if non-nil, zero value otherwise.

### GetIsAccountPoolOk

`func (o *AccountGenericResponse) GetIsAccountPoolOk() (*bool, bool)`

GetIsAccountPoolOk returns a tuple with the IsAccountPool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAccountPool

`func (o *AccountGenericResponse) SetIsAccountPool(v bool)`

SetIsAccountPool sets IsAccountPool field to given value.

### HasIsAccountPool

`func (o *AccountGenericResponse) HasIsAccountPool() bool`

HasIsAccountPool returns a boolean if a field has been set.

### GetIsAchEnabled

`func (o *AccountGenericResponse) GetIsAchEnabled() bool`

GetIsAchEnabled returns the IsAchEnabled field if non-nil, zero value otherwise.

### GetIsAchEnabledOk

`func (o *AccountGenericResponse) GetIsAchEnabledOk() (*bool, bool)`

GetIsAchEnabledOk returns a tuple with the IsAchEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAchEnabled

`func (o *AccountGenericResponse) SetIsAchEnabled(v bool)`

SetIsAchEnabled sets IsAchEnabled field to given value.

### HasIsAchEnabled

`func (o *AccountGenericResponse) HasIsAchEnabled() bool`

HasIsAchEnabled returns a boolean if a field has been set.

### GetIsCardEnabled

`func (o *AccountGenericResponse) GetIsCardEnabled() bool`

GetIsCardEnabled returns the IsCardEnabled field if non-nil, zero value otherwise.

### GetIsCardEnabledOk

`func (o *AccountGenericResponse) GetIsCardEnabledOk() (*bool, bool)`

GetIsCardEnabledOk returns a tuple with the IsCardEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsCardEnabled

`func (o *AccountGenericResponse) SetIsCardEnabled(v bool)`

SetIsCardEnabled sets IsCardEnabled field to given value.

### HasIsCardEnabled

`func (o *AccountGenericResponse) HasIsCardEnabled() bool`

HasIsCardEnabled returns a boolean if a field has been set.

### GetIsP2pEnabled

`func (o *AccountGenericResponse) GetIsP2pEnabled() bool`

GetIsP2pEnabled returns the IsP2pEnabled field if non-nil, zero value otherwise.

### GetIsP2pEnabledOk

`func (o *AccountGenericResponse) GetIsP2pEnabledOk() (*bool, bool)`

GetIsP2pEnabledOk returns a tuple with the IsP2pEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsP2pEnabled

`func (o *AccountGenericResponse) SetIsP2pEnabled(v bool)`

SetIsP2pEnabled sets IsP2pEnabled field to given value.

### HasIsP2pEnabled

`func (o *AccountGenericResponse) HasIsP2pEnabled() bool`

HasIsP2pEnabled returns a boolean if a field has been set.

### GetIsWireEnabled

`func (o *AccountGenericResponse) GetIsWireEnabled() bool`

GetIsWireEnabled returns the IsWireEnabled field if non-nil, zero value otherwise.

### GetIsWireEnabledOk

`func (o *AccountGenericResponse) GetIsWireEnabledOk() (*bool, bool)`

GetIsWireEnabledOk returns a tuple with the IsWireEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsWireEnabled

`func (o *AccountGenericResponse) SetIsWireEnabled(v bool)`

SetIsWireEnabled sets IsWireEnabled field to given value.

### HasIsWireEnabled

`func (o *AccountGenericResponse) HasIsWireEnabled() bool`

HasIsWireEnabled returns a boolean if a field has been set.

### GetLastUpdatedTime

`func (o *AccountGenericResponse) GetLastUpdatedTime() time.Time`

GetLastUpdatedTime returns the LastUpdatedTime field if non-nil, zero value otherwise.

### GetLastUpdatedTimeOk

`func (o *AccountGenericResponse) GetLastUpdatedTimeOk() (*time.Time, bool)`

GetLastUpdatedTimeOk returns a tuple with the LastUpdatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdatedTime

`func (o *AccountGenericResponse) SetLastUpdatedTime(v time.Time)`

SetLastUpdatedTime sets LastUpdatedTime field to given value.

### HasLastUpdatedTime

`func (o *AccountGenericResponse) HasLastUpdatedTime() bool`

HasLastUpdatedTime returns a boolean if a field has been set.

### GetMetadata

`func (o *AccountGenericResponse) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *AccountGenericResponse) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *AccountGenericResponse) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *AccountGenericResponse) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetMinimumPayment

`func (o *AccountGenericResponse) GetMinimumPayment() MinimumPayment`

GetMinimumPayment returns the MinimumPayment field if non-nil, zero value otherwise.

### GetMinimumPaymentOk

`func (o *AccountGenericResponse) GetMinimumPaymentOk() (*MinimumPayment, bool)`

GetMinimumPaymentOk returns a tuple with the MinimumPayment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinimumPayment

`func (o *AccountGenericResponse) SetMinimumPayment(v MinimumPayment)`

SetMinimumPayment sets MinimumPayment field to given value.

### HasMinimumPayment

`func (o *AccountGenericResponse) HasMinimumPayment() bool`

HasMinimumPayment returns a boolean if a field has been set.

### GetNickname

`func (o *AccountGenericResponse) GetNickname() string`

GetNickname returns the Nickname field if non-nil, zero value otherwise.

### GetNicknameOk

`func (o *AccountGenericResponse) GetNicknameOk() (*string, bool)`

GetNicknameOk returns a tuple with the Nickname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNickname

`func (o *AccountGenericResponse) SetNickname(v string)`

SetNickname sets Nickname field to given value.

### HasNickname

`func (o *AccountGenericResponse) HasNickname() bool`

HasNickname returns a boolean if a field has been set.

### GetOverdraftLimit

`func (o *AccountGenericResponse) GetOverdraftLimit() int64`

GetOverdraftLimit returns the OverdraftLimit field if non-nil, zero value otherwise.

### GetOverdraftLimitOk

`func (o *AccountGenericResponse) GetOverdraftLimitOk() (*int64, bool)`

GetOverdraftLimitOk returns a tuple with the OverdraftLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverdraftLimit

`func (o *AccountGenericResponse) SetOverdraftLimit(v int64)`

SetOverdraftLimit sets OverdraftLimit field to given value.

### HasOverdraftLimit

`func (o *AccountGenericResponse) HasOverdraftLimit() bool`

HasOverdraftLimit returns a boolean if a field has been set.

### GetSpendingLimits

`func (o *AccountGenericResponse) GetSpendingLimits() SpendingLimits`

GetSpendingLimits returns the SpendingLimits field if non-nil, zero value otherwise.

### GetSpendingLimitsOk

`func (o *AccountGenericResponse) GetSpendingLimitsOk() (*SpendingLimits, bool)`

GetSpendingLimitsOk returns a tuple with the SpendingLimits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpendingLimits

`func (o *AccountGenericResponse) SetSpendingLimits(v SpendingLimits)`

SetSpendingLimits sets SpendingLimits field to given value.

### HasSpendingLimits

`func (o *AccountGenericResponse) HasSpendingLimits() bool`

HasSpendingLimits returns a boolean if a field has been set.

### GetStatus

`func (o *AccountGenericResponse) GetStatus() Status`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *AccountGenericResponse) GetStatusOk() (*Status, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *AccountGenericResponse) SetStatus(v Status)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *AccountGenericResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSwiftCode

`func (o *AccountGenericResponse) GetSwiftCode() string`

GetSwiftCode returns the SwiftCode field if non-nil, zero value otherwise.

### GetSwiftCodeOk

`func (o *AccountGenericResponse) GetSwiftCodeOk() (*string, bool)`

GetSwiftCodeOk returns a tuple with the SwiftCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSwiftCode

`func (o *AccountGenericResponse) SetSwiftCode(v string)`

SetSwiftCode sets SwiftCode field to given value.

### HasSwiftCode

`func (o *AccountGenericResponse) HasSwiftCode() bool`

HasSwiftCode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


