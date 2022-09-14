# AccountLineOfCredit

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessStatus** | Pointer to [**AccountAccessStatus**](AccountAccessStatus.md) |  | [optional] 
**AccountNumber** | Pointer to **string** | Account number | [optional] [readonly] 
**AccountNumberMasked** | Pointer to **string** | The response will contain the bank fintech ID (3 or 6 digits) plus the last 4 digits, with the digits in between replaced with * characters. Shadow mode account numbers will not be masked. | [optional] [readonly] 
**AccountPurpose** | Pointer to **string** | Purpose of the account | [optional] 
**AccountType** | Pointer to [**AccountType**](AccountType.md) |  | [optional] 
**ApplicationId** | Pointer to **string** | The application ID for this account.  | [optional] 
**Balances** | Pointer to [**[]Balance**](Balance.md) | A list of balances for account based on different type | [optional] [readonly] 
**BankRouting** | Pointer to **string** | Bank routing number | [optional] [readonly] 
**CreationTime** | Pointer to **time.Time** | Account creation timestamp in RFC3339 format | [optional] [readonly] 
**Currency** | Pointer to **string** | Account currency or account settlement currency. ISO 4217 alphabetic currency code. Default USD | [optional] 
**CustomerIds** | Pointer to **[]string** | A list of the customer IDs of the account holders. | [optional] [readonly] 
**CustomerType** | Pointer to [**CustomerType**](CustomerType.md) |  | [optional] 
**ExchangeRateType** | Pointer to **string** | Exchange rate type | [optional] 
**Iban** | Pointer to **string** | International bank account number | [optional] 
**Id** | Pointer to **string** | Account ID | [optional] [readonly] 
**IsAccountPool** | Pointer to **bool** | Account is investment (variable balance) account or a multi-balance account pool. Default false | [optional] 
**IsAchEnabled** | Pointer to **bool** | A flag to indicate whether ACH transactions are enabled. | [optional] [readonly] 
**IsCardEnabled** | Pointer to **bool** | A flag to indicate whether card transactions are enabled. | [optional] [readonly] 
**IsP2pEnabled** | Pointer to **bool** | A flag to indicate whether P2P transactions are enabled. | [optional] [readonly] 
**IsWireEnabled** | Pointer to **bool** | A flag to indicate whether wire transactions are enabled. | [optional] [readonly] 
**LastUpdatedTime** | Pointer to **time.Time** | Timestamp of the last account modification in RFC3339 format | [optional] [readonly] 
**Metadata** | Pointer to **map[string]interface{}** | User provided account metadata | [optional] 
**Nickname** | Pointer to **string** | User provided account nickname | [optional] 
**Status** | Pointer to [**Status**](Status.md) |  | [optional] 
**SwiftCode** | Pointer to **string** | SWIFT code | [optional] 
**ChargeoffPeriod** | Pointer to **int32** | The number of days an account can stay delinquent before marking an account as charged-off.  | [optional] [default to 90]
**CreditLimit** | Pointer to **int64** | The credit limit for this line of credit account in cents. Minimum is 0.  | [optional] 
**DelinquencyPeriod** | Pointer to **int32** | The number of days past the due date to wait for a minimum payment before marking an account as delinquent.  | [optional] [default to 30]
**GracePeriod** | Pointer to **int32** | The number of days past the billing period to allow for payment before it is considered due. This directly infers the due date for a payment.  | [optional] 
**InterestProductId** | Pointer to **string** | An interest account product that the current account associates with. The account product must have its calculation_method set to COMPOUNDED_DAILY.  | [optional] 
**MinimumPayment** | Pointer to [**MinimumPayment**](MinimumPayment.md) |  | [optional] 

## Methods

### NewAccountLineOfCredit

`func NewAccountLineOfCredit() *AccountLineOfCredit`

NewAccountLineOfCredit instantiates a new AccountLineOfCredit object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountLineOfCreditWithDefaults

`func NewAccountLineOfCreditWithDefaults() *AccountLineOfCredit`

NewAccountLineOfCreditWithDefaults instantiates a new AccountLineOfCredit object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessStatus

`func (o *AccountLineOfCredit) GetAccessStatus() AccountAccessStatus`

GetAccessStatus returns the AccessStatus field if non-nil, zero value otherwise.

### GetAccessStatusOk

`func (o *AccountLineOfCredit) GetAccessStatusOk() (*AccountAccessStatus, bool)`

GetAccessStatusOk returns a tuple with the AccessStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessStatus

`func (o *AccountLineOfCredit) SetAccessStatus(v AccountAccessStatus)`

SetAccessStatus sets AccessStatus field to given value.

### HasAccessStatus

`func (o *AccountLineOfCredit) HasAccessStatus() bool`

HasAccessStatus returns a boolean if a field has been set.

### GetAccountNumber

`func (o *AccountLineOfCredit) GetAccountNumber() string`

GetAccountNumber returns the AccountNumber field if non-nil, zero value otherwise.

### GetAccountNumberOk

`func (o *AccountLineOfCredit) GetAccountNumberOk() (*string, bool)`

GetAccountNumberOk returns a tuple with the AccountNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountNumber

`func (o *AccountLineOfCredit) SetAccountNumber(v string)`

SetAccountNumber sets AccountNumber field to given value.

### HasAccountNumber

`func (o *AccountLineOfCredit) HasAccountNumber() bool`

HasAccountNumber returns a boolean if a field has been set.

### GetAccountNumberMasked

`func (o *AccountLineOfCredit) GetAccountNumberMasked() string`

GetAccountNumberMasked returns the AccountNumberMasked field if non-nil, zero value otherwise.

### GetAccountNumberMaskedOk

`func (o *AccountLineOfCredit) GetAccountNumberMaskedOk() (*string, bool)`

GetAccountNumberMaskedOk returns a tuple with the AccountNumberMasked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountNumberMasked

`func (o *AccountLineOfCredit) SetAccountNumberMasked(v string)`

SetAccountNumberMasked sets AccountNumberMasked field to given value.

### HasAccountNumberMasked

`func (o *AccountLineOfCredit) HasAccountNumberMasked() bool`

HasAccountNumberMasked returns a boolean if a field has been set.

### GetAccountPurpose

`func (o *AccountLineOfCredit) GetAccountPurpose() string`

GetAccountPurpose returns the AccountPurpose field if non-nil, zero value otherwise.

### GetAccountPurposeOk

`func (o *AccountLineOfCredit) GetAccountPurposeOk() (*string, bool)`

GetAccountPurposeOk returns a tuple with the AccountPurpose field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountPurpose

`func (o *AccountLineOfCredit) SetAccountPurpose(v string)`

SetAccountPurpose sets AccountPurpose field to given value.

### HasAccountPurpose

`func (o *AccountLineOfCredit) HasAccountPurpose() bool`

HasAccountPurpose returns a boolean if a field has been set.

### GetAccountType

`func (o *AccountLineOfCredit) GetAccountType() AccountType`

GetAccountType returns the AccountType field if non-nil, zero value otherwise.

### GetAccountTypeOk

`func (o *AccountLineOfCredit) GetAccountTypeOk() (*AccountType, bool)`

GetAccountTypeOk returns a tuple with the AccountType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountType

`func (o *AccountLineOfCredit) SetAccountType(v AccountType)`

SetAccountType sets AccountType field to given value.

### HasAccountType

`func (o *AccountLineOfCredit) HasAccountType() bool`

HasAccountType returns a boolean if a field has been set.

### GetApplicationId

`func (o *AccountLineOfCredit) GetApplicationId() string`

GetApplicationId returns the ApplicationId field if non-nil, zero value otherwise.

### GetApplicationIdOk

`func (o *AccountLineOfCredit) GetApplicationIdOk() (*string, bool)`

GetApplicationIdOk returns a tuple with the ApplicationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationId

`func (o *AccountLineOfCredit) SetApplicationId(v string)`

SetApplicationId sets ApplicationId field to given value.

### HasApplicationId

`func (o *AccountLineOfCredit) HasApplicationId() bool`

HasApplicationId returns a boolean if a field has been set.

### GetBalances

`func (o *AccountLineOfCredit) GetBalances() []Balance`

GetBalances returns the Balances field if non-nil, zero value otherwise.

### GetBalancesOk

`func (o *AccountLineOfCredit) GetBalancesOk() (*[]Balance, bool)`

GetBalancesOk returns a tuple with the Balances field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalances

`func (o *AccountLineOfCredit) SetBalances(v []Balance)`

SetBalances sets Balances field to given value.

### HasBalances

`func (o *AccountLineOfCredit) HasBalances() bool`

HasBalances returns a boolean if a field has been set.

### GetBankRouting

`func (o *AccountLineOfCredit) GetBankRouting() string`

GetBankRouting returns the BankRouting field if non-nil, zero value otherwise.

### GetBankRoutingOk

`func (o *AccountLineOfCredit) GetBankRoutingOk() (*string, bool)`

GetBankRoutingOk returns a tuple with the BankRouting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBankRouting

`func (o *AccountLineOfCredit) SetBankRouting(v string)`

SetBankRouting sets BankRouting field to given value.

### HasBankRouting

`func (o *AccountLineOfCredit) HasBankRouting() bool`

HasBankRouting returns a boolean if a field has been set.

### GetCreationTime

`func (o *AccountLineOfCredit) GetCreationTime() time.Time`

GetCreationTime returns the CreationTime field if non-nil, zero value otherwise.

### GetCreationTimeOk

`func (o *AccountLineOfCredit) GetCreationTimeOk() (*time.Time, bool)`

GetCreationTimeOk returns a tuple with the CreationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationTime

`func (o *AccountLineOfCredit) SetCreationTime(v time.Time)`

SetCreationTime sets CreationTime field to given value.

### HasCreationTime

`func (o *AccountLineOfCredit) HasCreationTime() bool`

HasCreationTime returns a boolean if a field has been set.

### GetCurrency

`func (o *AccountLineOfCredit) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *AccountLineOfCredit) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *AccountLineOfCredit) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *AccountLineOfCredit) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetCustomerIds

`func (o *AccountLineOfCredit) GetCustomerIds() []string`

GetCustomerIds returns the CustomerIds field if non-nil, zero value otherwise.

### GetCustomerIdsOk

`func (o *AccountLineOfCredit) GetCustomerIdsOk() (*[]string, bool)`

GetCustomerIdsOk returns a tuple with the CustomerIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerIds

`func (o *AccountLineOfCredit) SetCustomerIds(v []string)`

SetCustomerIds sets CustomerIds field to given value.

### HasCustomerIds

`func (o *AccountLineOfCredit) HasCustomerIds() bool`

HasCustomerIds returns a boolean if a field has been set.

### GetCustomerType

`func (o *AccountLineOfCredit) GetCustomerType() CustomerType`

GetCustomerType returns the CustomerType field if non-nil, zero value otherwise.

### GetCustomerTypeOk

`func (o *AccountLineOfCredit) GetCustomerTypeOk() (*CustomerType, bool)`

GetCustomerTypeOk returns a tuple with the CustomerType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerType

`func (o *AccountLineOfCredit) SetCustomerType(v CustomerType)`

SetCustomerType sets CustomerType field to given value.

### HasCustomerType

`func (o *AccountLineOfCredit) HasCustomerType() bool`

HasCustomerType returns a boolean if a field has been set.

### GetExchangeRateType

`func (o *AccountLineOfCredit) GetExchangeRateType() string`

GetExchangeRateType returns the ExchangeRateType field if non-nil, zero value otherwise.

### GetExchangeRateTypeOk

`func (o *AccountLineOfCredit) GetExchangeRateTypeOk() (*string, bool)`

GetExchangeRateTypeOk returns a tuple with the ExchangeRateType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExchangeRateType

`func (o *AccountLineOfCredit) SetExchangeRateType(v string)`

SetExchangeRateType sets ExchangeRateType field to given value.

### HasExchangeRateType

`func (o *AccountLineOfCredit) HasExchangeRateType() bool`

HasExchangeRateType returns a boolean if a field has been set.

### GetIban

`func (o *AccountLineOfCredit) GetIban() string`

GetIban returns the Iban field if non-nil, zero value otherwise.

### GetIbanOk

`func (o *AccountLineOfCredit) GetIbanOk() (*string, bool)`

GetIbanOk returns a tuple with the Iban field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIban

`func (o *AccountLineOfCredit) SetIban(v string)`

SetIban sets Iban field to given value.

### HasIban

`func (o *AccountLineOfCredit) HasIban() bool`

HasIban returns a boolean if a field has been set.

### GetId

`func (o *AccountLineOfCredit) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AccountLineOfCredit) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AccountLineOfCredit) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *AccountLineOfCredit) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIsAccountPool

`func (o *AccountLineOfCredit) GetIsAccountPool() bool`

GetIsAccountPool returns the IsAccountPool field if non-nil, zero value otherwise.

### GetIsAccountPoolOk

`func (o *AccountLineOfCredit) GetIsAccountPoolOk() (*bool, bool)`

GetIsAccountPoolOk returns a tuple with the IsAccountPool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAccountPool

`func (o *AccountLineOfCredit) SetIsAccountPool(v bool)`

SetIsAccountPool sets IsAccountPool field to given value.

### HasIsAccountPool

`func (o *AccountLineOfCredit) HasIsAccountPool() bool`

HasIsAccountPool returns a boolean if a field has been set.

### GetIsAchEnabled

`func (o *AccountLineOfCredit) GetIsAchEnabled() bool`

GetIsAchEnabled returns the IsAchEnabled field if non-nil, zero value otherwise.

### GetIsAchEnabledOk

`func (o *AccountLineOfCredit) GetIsAchEnabledOk() (*bool, bool)`

GetIsAchEnabledOk returns a tuple with the IsAchEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAchEnabled

`func (o *AccountLineOfCredit) SetIsAchEnabled(v bool)`

SetIsAchEnabled sets IsAchEnabled field to given value.

### HasIsAchEnabled

`func (o *AccountLineOfCredit) HasIsAchEnabled() bool`

HasIsAchEnabled returns a boolean if a field has been set.

### GetIsCardEnabled

`func (o *AccountLineOfCredit) GetIsCardEnabled() bool`

GetIsCardEnabled returns the IsCardEnabled field if non-nil, zero value otherwise.

### GetIsCardEnabledOk

`func (o *AccountLineOfCredit) GetIsCardEnabledOk() (*bool, bool)`

GetIsCardEnabledOk returns a tuple with the IsCardEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsCardEnabled

`func (o *AccountLineOfCredit) SetIsCardEnabled(v bool)`

SetIsCardEnabled sets IsCardEnabled field to given value.

### HasIsCardEnabled

`func (o *AccountLineOfCredit) HasIsCardEnabled() bool`

HasIsCardEnabled returns a boolean if a field has been set.

### GetIsP2pEnabled

`func (o *AccountLineOfCredit) GetIsP2pEnabled() bool`

GetIsP2pEnabled returns the IsP2pEnabled field if non-nil, zero value otherwise.

### GetIsP2pEnabledOk

`func (o *AccountLineOfCredit) GetIsP2pEnabledOk() (*bool, bool)`

GetIsP2pEnabledOk returns a tuple with the IsP2pEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsP2pEnabled

`func (o *AccountLineOfCredit) SetIsP2pEnabled(v bool)`

SetIsP2pEnabled sets IsP2pEnabled field to given value.

### HasIsP2pEnabled

`func (o *AccountLineOfCredit) HasIsP2pEnabled() bool`

HasIsP2pEnabled returns a boolean if a field has been set.

### GetIsWireEnabled

`func (o *AccountLineOfCredit) GetIsWireEnabled() bool`

GetIsWireEnabled returns the IsWireEnabled field if non-nil, zero value otherwise.

### GetIsWireEnabledOk

`func (o *AccountLineOfCredit) GetIsWireEnabledOk() (*bool, bool)`

GetIsWireEnabledOk returns a tuple with the IsWireEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsWireEnabled

`func (o *AccountLineOfCredit) SetIsWireEnabled(v bool)`

SetIsWireEnabled sets IsWireEnabled field to given value.

### HasIsWireEnabled

`func (o *AccountLineOfCredit) HasIsWireEnabled() bool`

HasIsWireEnabled returns a boolean if a field has been set.

### GetLastUpdatedTime

`func (o *AccountLineOfCredit) GetLastUpdatedTime() time.Time`

GetLastUpdatedTime returns the LastUpdatedTime field if non-nil, zero value otherwise.

### GetLastUpdatedTimeOk

`func (o *AccountLineOfCredit) GetLastUpdatedTimeOk() (*time.Time, bool)`

GetLastUpdatedTimeOk returns a tuple with the LastUpdatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdatedTime

`func (o *AccountLineOfCredit) SetLastUpdatedTime(v time.Time)`

SetLastUpdatedTime sets LastUpdatedTime field to given value.

### HasLastUpdatedTime

`func (o *AccountLineOfCredit) HasLastUpdatedTime() bool`

HasLastUpdatedTime returns a boolean if a field has been set.

### GetMetadata

`func (o *AccountLineOfCredit) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *AccountLineOfCredit) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *AccountLineOfCredit) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *AccountLineOfCredit) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetNickname

`func (o *AccountLineOfCredit) GetNickname() string`

GetNickname returns the Nickname field if non-nil, zero value otherwise.

### GetNicknameOk

`func (o *AccountLineOfCredit) GetNicknameOk() (*string, bool)`

GetNicknameOk returns a tuple with the Nickname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNickname

`func (o *AccountLineOfCredit) SetNickname(v string)`

SetNickname sets Nickname field to given value.

### HasNickname

`func (o *AccountLineOfCredit) HasNickname() bool`

HasNickname returns a boolean if a field has been set.

### GetStatus

`func (o *AccountLineOfCredit) GetStatus() Status`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *AccountLineOfCredit) GetStatusOk() (*Status, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *AccountLineOfCredit) SetStatus(v Status)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *AccountLineOfCredit) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSwiftCode

`func (o *AccountLineOfCredit) GetSwiftCode() string`

GetSwiftCode returns the SwiftCode field if non-nil, zero value otherwise.

### GetSwiftCodeOk

`func (o *AccountLineOfCredit) GetSwiftCodeOk() (*string, bool)`

GetSwiftCodeOk returns a tuple with the SwiftCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSwiftCode

`func (o *AccountLineOfCredit) SetSwiftCode(v string)`

SetSwiftCode sets SwiftCode field to given value.

### HasSwiftCode

`func (o *AccountLineOfCredit) HasSwiftCode() bool`

HasSwiftCode returns a boolean if a field has been set.

### GetChargeoffPeriod

`func (o *AccountLineOfCredit) GetChargeoffPeriod() int32`

GetChargeoffPeriod returns the ChargeoffPeriod field if non-nil, zero value otherwise.

### GetChargeoffPeriodOk

`func (o *AccountLineOfCredit) GetChargeoffPeriodOk() (*int32, bool)`

GetChargeoffPeriodOk returns a tuple with the ChargeoffPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChargeoffPeriod

`func (o *AccountLineOfCredit) SetChargeoffPeriod(v int32)`

SetChargeoffPeriod sets ChargeoffPeriod field to given value.

### HasChargeoffPeriod

`func (o *AccountLineOfCredit) HasChargeoffPeriod() bool`

HasChargeoffPeriod returns a boolean if a field has been set.

### GetCreditLimit

`func (o *AccountLineOfCredit) GetCreditLimit() int64`

GetCreditLimit returns the CreditLimit field if non-nil, zero value otherwise.

### GetCreditLimitOk

`func (o *AccountLineOfCredit) GetCreditLimitOk() (*int64, bool)`

GetCreditLimitOk returns a tuple with the CreditLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreditLimit

`func (o *AccountLineOfCredit) SetCreditLimit(v int64)`

SetCreditLimit sets CreditLimit field to given value.

### HasCreditLimit

`func (o *AccountLineOfCredit) HasCreditLimit() bool`

HasCreditLimit returns a boolean if a field has been set.

### GetDelinquencyPeriod

`func (o *AccountLineOfCredit) GetDelinquencyPeriod() int32`

GetDelinquencyPeriod returns the DelinquencyPeriod field if non-nil, zero value otherwise.

### GetDelinquencyPeriodOk

`func (o *AccountLineOfCredit) GetDelinquencyPeriodOk() (*int32, bool)`

GetDelinquencyPeriodOk returns a tuple with the DelinquencyPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelinquencyPeriod

`func (o *AccountLineOfCredit) SetDelinquencyPeriod(v int32)`

SetDelinquencyPeriod sets DelinquencyPeriod field to given value.

### HasDelinquencyPeriod

`func (o *AccountLineOfCredit) HasDelinquencyPeriod() bool`

HasDelinquencyPeriod returns a boolean if a field has been set.

### GetGracePeriod

`func (o *AccountLineOfCredit) GetGracePeriod() int32`

GetGracePeriod returns the GracePeriod field if non-nil, zero value otherwise.

### GetGracePeriodOk

`func (o *AccountLineOfCredit) GetGracePeriodOk() (*int32, bool)`

GetGracePeriodOk returns a tuple with the GracePeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGracePeriod

`func (o *AccountLineOfCredit) SetGracePeriod(v int32)`

SetGracePeriod sets GracePeriod field to given value.

### HasGracePeriod

`func (o *AccountLineOfCredit) HasGracePeriod() bool`

HasGracePeriod returns a boolean if a field has been set.

### GetInterestProductId

`func (o *AccountLineOfCredit) GetInterestProductId() string`

GetInterestProductId returns the InterestProductId field if non-nil, zero value otherwise.

### GetInterestProductIdOk

`func (o *AccountLineOfCredit) GetInterestProductIdOk() (*string, bool)`

GetInterestProductIdOk returns a tuple with the InterestProductId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterestProductId

`func (o *AccountLineOfCredit) SetInterestProductId(v string)`

SetInterestProductId sets InterestProductId field to given value.

### HasInterestProductId

`func (o *AccountLineOfCredit) HasInterestProductId() bool`

HasInterestProductId returns a boolean if a field has been set.

### GetMinimumPayment

`func (o *AccountLineOfCredit) GetMinimumPayment() MinimumPayment`

GetMinimumPayment returns the MinimumPayment field if non-nil, zero value otherwise.

### GetMinimumPaymentOk

`func (o *AccountLineOfCredit) GetMinimumPaymentOk() (*MinimumPayment, bool)`

GetMinimumPaymentOk returns a tuple with the MinimumPayment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinimumPayment

`func (o *AccountLineOfCredit) SetMinimumPayment(v MinimumPayment)`

SetMinimumPayment sets MinimumPayment field to given value.

### HasMinimumPayment

`func (o *AccountLineOfCredit) HasMinimumPayment() bool`

HasMinimumPayment returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


