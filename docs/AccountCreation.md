# AccountCreation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessStatus** | Pointer to [**AccountAccessStatus**](AccountAccessStatus.md) |  | [optional] 
**AccountNumber** | Pointer to **string** | Account number | [optional] [readonly] 
**AccountPurpose** | Pointer to **string** | Purpose of the account | [optional] 
**AccountType** | Pointer to [**AccountType**](AccountType.md) |  | [optional] 
**BalanceCeiling** | Pointer to [**BalanceCeiling**](BalanceCeiling.md) |  | [optional] 
**BalanceFloor** | Pointer to [**BalanceFloor**](BalanceFloor.md) |  | [optional] 
**Balances** | Pointer to [**[]Balance**](Balance.md) | A list of balances for account based on different type | [optional] [readonly] 
**BankRouting** | Pointer to **string** | Bank routing number | [optional] [readonly] 
**CreationTime** | Pointer to **time.Time** | Account creation timestamp in RFC3337 format | [optional] [readonly] 
**Currency** | Pointer to **string** | Account currency or account settlement currency. ISO 4217 alphabetic currency code. Default USD | [optional] 
**CustomerIds** | Pointer to **[]string** | A list of the customer IDs of the account holders. | [optional] [readonly] 
**CustomerType** | Pointer to [**CustomerType**](CustomerType.md) |  | [optional] 
**ExchangeRateType** | Pointer to **string** | Exchange rate type | [optional] 
**FeeProductIds** | Pointer to **[]string** | A list of fee resources from account product that the current account associate with | [optional] 
**Iban** | Pointer to **string** | International bank account number | [optional] 
**Id** | Pointer to **string** | Account ID | [optional] [readonly] 
**InterestProductId** | Pointer to **string** | An interest from account product that the current account associate with | [optional] 
**IsAccountPool** | Pointer to **bool** | Account is investment (variable balance) account or a multi-balance account pool. Default false | [optional] 
**LastUpdatedTime** | Pointer to **time.Time** | Timestamp of the last account modification in RFC3337 format | [optional] [readonly] 
**Metadata** | Pointer to **map[string]interface{}** | User provided account metadata | [optional] 
**Nickname** | Pointer to **string** | User provided account nickname | [optional] 
**OverdraftLimit** | Pointer to **int64** | Account&#39;s overdraft limit | [optional] 
**SpendingLimits** | Pointer to [**SpendingLimits**](SpendingLimits.md) |  | [optional] 
**Status** | Pointer to [**Status**](Status.md) |  | [optional] 
**SwiftCode** | Pointer to **string** | SWIFT code | [optional] 
**AccountTemplateId** | Pointer to **string** | Account template ID | [optional] 
**Relationships** | Pointer to [**[]Relationship**](Relationship.md) | List of the relationship for this account to the parties | [optional] 

## Methods

### NewAccountCreation

`func NewAccountCreation() *AccountCreation`

NewAccountCreation instantiates a new AccountCreation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountCreationWithDefaults

`func NewAccountCreationWithDefaults() *AccountCreation`

NewAccountCreationWithDefaults instantiates a new AccountCreation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessStatus

`func (o *AccountCreation) GetAccessStatus() AccountAccessStatus`

GetAccessStatus returns the AccessStatus field if non-nil, zero value otherwise.

### GetAccessStatusOk

`func (o *AccountCreation) GetAccessStatusOk() (*AccountAccessStatus, bool)`

GetAccessStatusOk returns a tuple with the AccessStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessStatus

`func (o *AccountCreation) SetAccessStatus(v AccountAccessStatus)`

SetAccessStatus sets AccessStatus field to given value.

### HasAccessStatus

`func (o *AccountCreation) HasAccessStatus() bool`

HasAccessStatus returns a boolean if a field has been set.

### GetAccountNumber

`func (o *AccountCreation) GetAccountNumber() string`

GetAccountNumber returns the AccountNumber field if non-nil, zero value otherwise.

### GetAccountNumberOk

`func (o *AccountCreation) GetAccountNumberOk() (*string, bool)`

GetAccountNumberOk returns a tuple with the AccountNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountNumber

`func (o *AccountCreation) SetAccountNumber(v string)`

SetAccountNumber sets AccountNumber field to given value.

### HasAccountNumber

`func (o *AccountCreation) HasAccountNumber() bool`

HasAccountNumber returns a boolean if a field has been set.

### GetAccountPurpose

`func (o *AccountCreation) GetAccountPurpose() string`

GetAccountPurpose returns the AccountPurpose field if non-nil, zero value otherwise.

### GetAccountPurposeOk

`func (o *AccountCreation) GetAccountPurposeOk() (*string, bool)`

GetAccountPurposeOk returns a tuple with the AccountPurpose field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountPurpose

`func (o *AccountCreation) SetAccountPurpose(v string)`

SetAccountPurpose sets AccountPurpose field to given value.

### HasAccountPurpose

`func (o *AccountCreation) HasAccountPurpose() bool`

HasAccountPurpose returns a boolean if a field has been set.

### GetAccountType

`func (o *AccountCreation) GetAccountType() AccountType`

GetAccountType returns the AccountType field if non-nil, zero value otherwise.

### GetAccountTypeOk

`func (o *AccountCreation) GetAccountTypeOk() (*AccountType, bool)`

GetAccountTypeOk returns a tuple with the AccountType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountType

`func (o *AccountCreation) SetAccountType(v AccountType)`

SetAccountType sets AccountType field to given value.

### HasAccountType

`func (o *AccountCreation) HasAccountType() bool`

HasAccountType returns a boolean if a field has been set.

### GetBalanceCeiling

`func (o *AccountCreation) GetBalanceCeiling() BalanceCeiling`

GetBalanceCeiling returns the BalanceCeiling field if non-nil, zero value otherwise.

### GetBalanceCeilingOk

`func (o *AccountCreation) GetBalanceCeilingOk() (*BalanceCeiling, bool)`

GetBalanceCeilingOk returns a tuple with the BalanceCeiling field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalanceCeiling

`func (o *AccountCreation) SetBalanceCeiling(v BalanceCeiling)`

SetBalanceCeiling sets BalanceCeiling field to given value.

### HasBalanceCeiling

`func (o *AccountCreation) HasBalanceCeiling() bool`

HasBalanceCeiling returns a boolean if a field has been set.

### GetBalanceFloor

`func (o *AccountCreation) GetBalanceFloor() BalanceFloor`

GetBalanceFloor returns the BalanceFloor field if non-nil, zero value otherwise.

### GetBalanceFloorOk

`func (o *AccountCreation) GetBalanceFloorOk() (*BalanceFloor, bool)`

GetBalanceFloorOk returns a tuple with the BalanceFloor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalanceFloor

`func (o *AccountCreation) SetBalanceFloor(v BalanceFloor)`

SetBalanceFloor sets BalanceFloor field to given value.

### HasBalanceFloor

`func (o *AccountCreation) HasBalanceFloor() bool`

HasBalanceFloor returns a boolean if a field has been set.

### GetBalances

`func (o *AccountCreation) GetBalances() []Balance`

GetBalances returns the Balances field if non-nil, zero value otherwise.

### GetBalancesOk

`func (o *AccountCreation) GetBalancesOk() (*[]Balance, bool)`

GetBalancesOk returns a tuple with the Balances field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalances

`func (o *AccountCreation) SetBalances(v []Balance)`

SetBalances sets Balances field to given value.

### HasBalances

`func (o *AccountCreation) HasBalances() bool`

HasBalances returns a boolean if a field has been set.

### GetBankRouting

`func (o *AccountCreation) GetBankRouting() string`

GetBankRouting returns the BankRouting field if non-nil, zero value otherwise.

### GetBankRoutingOk

`func (o *AccountCreation) GetBankRoutingOk() (*string, bool)`

GetBankRoutingOk returns a tuple with the BankRouting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBankRouting

`func (o *AccountCreation) SetBankRouting(v string)`

SetBankRouting sets BankRouting field to given value.

### HasBankRouting

`func (o *AccountCreation) HasBankRouting() bool`

HasBankRouting returns a boolean if a field has been set.

### GetCreationTime

`func (o *AccountCreation) GetCreationTime() time.Time`

GetCreationTime returns the CreationTime field if non-nil, zero value otherwise.

### GetCreationTimeOk

`func (o *AccountCreation) GetCreationTimeOk() (*time.Time, bool)`

GetCreationTimeOk returns a tuple with the CreationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationTime

`func (o *AccountCreation) SetCreationTime(v time.Time)`

SetCreationTime sets CreationTime field to given value.

### HasCreationTime

`func (o *AccountCreation) HasCreationTime() bool`

HasCreationTime returns a boolean if a field has been set.

### GetCurrency

`func (o *AccountCreation) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *AccountCreation) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *AccountCreation) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *AccountCreation) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetCustomerIds

`func (o *AccountCreation) GetCustomerIds() []string`

GetCustomerIds returns the CustomerIds field if non-nil, zero value otherwise.

### GetCustomerIdsOk

`func (o *AccountCreation) GetCustomerIdsOk() (*[]string, bool)`

GetCustomerIdsOk returns a tuple with the CustomerIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerIds

`func (o *AccountCreation) SetCustomerIds(v []string)`

SetCustomerIds sets CustomerIds field to given value.

### HasCustomerIds

`func (o *AccountCreation) HasCustomerIds() bool`

HasCustomerIds returns a boolean if a field has been set.

### GetCustomerType

`func (o *AccountCreation) GetCustomerType() CustomerType`

GetCustomerType returns the CustomerType field if non-nil, zero value otherwise.

### GetCustomerTypeOk

`func (o *AccountCreation) GetCustomerTypeOk() (*CustomerType, bool)`

GetCustomerTypeOk returns a tuple with the CustomerType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerType

`func (o *AccountCreation) SetCustomerType(v CustomerType)`

SetCustomerType sets CustomerType field to given value.

### HasCustomerType

`func (o *AccountCreation) HasCustomerType() bool`

HasCustomerType returns a boolean if a field has been set.

### GetExchangeRateType

`func (o *AccountCreation) GetExchangeRateType() string`

GetExchangeRateType returns the ExchangeRateType field if non-nil, zero value otherwise.

### GetExchangeRateTypeOk

`func (o *AccountCreation) GetExchangeRateTypeOk() (*string, bool)`

GetExchangeRateTypeOk returns a tuple with the ExchangeRateType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExchangeRateType

`func (o *AccountCreation) SetExchangeRateType(v string)`

SetExchangeRateType sets ExchangeRateType field to given value.

### HasExchangeRateType

`func (o *AccountCreation) HasExchangeRateType() bool`

HasExchangeRateType returns a boolean if a field has been set.

### GetFeeProductIds

`func (o *AccountCreation) GetFeeProductIds() []string`

GetFeeProductIds returns the FeeProductIds field if non-nil, zero value otherwise.

### GetFeeProductIdsOk

`func (o *AccountCreation) GetFeeProductIdsOk() (*[]string, bool)`

GetFeeProductIdsOk returns a tuple with the FeeProductIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeeProductIds

`func (o *AccountCreation) SetFeeProductIds(v []string)`

SetFeeProductIds sets FeeProductIds field to given value.

### HasFeeProductIds

`func (o *AccountCreation) HasFeeProductIds() bool`

HasFeeProductIds returns a boolean if a field has been set.

### GetIban

`func (o *AccountCreation) GetIban() string`

GetIban returns the Iban field if non-nil, zero value otherwise.

### GetIbanOk

`func (o *AccountCreation) GetIbanOk() (*string, bool)`

GetIbanOk returns a tuple with the Iban field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIban

`func (o *AccountCreation) SetIban(v string)`

SetIban sets Iban field to given value.

### HasIban

`func (o *AccountCreation) HasIban() bool`

HasIban returns a boolean if a field has been set.

### GetId

`func (o *AccountCreation) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AccountCreation) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AccountCreation) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *AccountCreation) HasId() bool`

HasId returns a boolean if a field has been set.

### GetInterestProductId

`func (o *AccountCreation) GetInterestProductId() string`

GetInterestProductId returns the InterestProductId field if non-nil, zero value otherwise.

### GetInterestProductIdOk

`func (o *AccountCreation) GetInterestProductIdOk() (*string, bool)`

GetInterestProductIdOk returns a tuple with the InterestProductId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterestProductId

`func (o *AccountCreation) SetInterestProductId(v string)`

SetInterestProductId sets InterestProductId field to given value.

### HasInterestProductId

`func (o *AccountCreation) HasInterestProductId() bool`

HasInterestProductId returns a boolean if a field has been set.

### GetIsAccountPool

`func (o *AccountCreation) GetIsAccountPool() bool`

GetIsAccountPool returns the IsAccountPool field if non-nil, zero value otherwise.

### GetIsAccountPoolOk

`func (o *AccountCreation) GetIsAccountPoolOk() (*bool, bool)`

GetIsAccountPoolOk returns a tuple with the IsAccountPool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAccountPool

`func (o *AccountCreation) SetIsAccountPool(v bool)`

SetIsAccountPool sets IsAccountPool field to given value.

### HasIsAccountPool

`func (o *AccountCreation) HasIsAccountPool() bool`

HasIsAccountPool returns a boolean if a field has been set.

### GetLastUpdatedTime

`func (o *AccountCreation) GetLastUpdatedTime() time.Time`

GetLastUpdatedTime returns the LastUpdatedTime field if non-nil, zero value otherwise.

### GetLastUpdatedTimeOk

`func (o *AccountCreation) GetLastUpdatedTimeOk() (*time.Time, bool)`

GetLastUpdatedTimeOk returns a tuple with the LastUpdatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdatedTime

`func (o *AccountCreation) SetLastUpdatedTime(v time.Time)`

SetLastUpdatedTime sets LastUpdatedTime field to given value.

### HasLastUpdatedTime

`func (o *AccountCreation) HasLastUpdatedTime() bool`

HasLastUpdatedTime returns a boolean if a field has been set.

### GetMetadata

`func (o *AccountCreation) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *AccountCreation) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *AccountCreation) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *AccountCreation) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetNickname

`func (o *AccountCreation) GetNickname() string`

GetNickname returns the Nickname field if non-nil, zero value otherwise.

### GetNicknameOk

`func (o *AccountCreation) GetNicknameOk() (*string, bool)`

GetNicknameOk returns a tuple with the Nickname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNickname

`func (o *AccountCreation) SetNickname(v string)`

SetNickname sets Nickname field to given value.

### HasNickname

`func (o *AccountCreation) HasNickname() bool`

HasNickname returns a boolean if a field has been set.

### GetOverdraftLimit

`func (o *AccountCreation) GetOverdraftLimit() int64`

GetOverdraftLimit returns the OverdraftLimit field if non-nil, zero value otherwise.

### GetOverdraftLimitOk

`func (o *AccountCreation) GetOverdraftLimitOk() (*int64, bool)`

GetOverdraftLimitOk returns a tuple with the OverdraftLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverdraftLimit

`func (o *AccountCreation) SetOverdraftLimit(v int64)`

SetOverdraftLimit sets OverdraftLimit field to given value.

### HasOverdraftLimit

`func (o *AccountCreation) HasOverdraftLimit() bool`

HasOverdraftLimit returns a boolean if a field has been set.

### GetSpendingLimits

`func (o *AccountCreation) GetSpendingLimits() SpendingLimits`

GetSpendingLimits returns the SpendingLimits field if non-nil, zero value otherwise.

### GetSpendingLimitsOk

`func (o *AccountCreation) GetSpendingLimitsOk() (*SpendingLimits, bool)`

GetSpendingLimitsOk returns a tuple with the SpendingLimits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpendingLimits

`func (o *AccountCreation) SetSpendingLimits(v SpendingLimits)`

SetSpendingLimits sets SpendingLimits field to given value.

### HasSpendingLimits

`func (o *AccountCreation) HasSpendingLimits() bool`

HasSpendingLimits returns a boolean if a field has been set.

### GetStatus

`func (o *AccountCreation) GetStatus() Status`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *AccountCreation) GetStatusOk() (*Status, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *AccountCreation) SetStatus(v Status)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *AccountCreation) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSwiftCode

`func (o *AccountCreation) GetSwiftCode() string`

GetSwiftCode returns the SwiftCode field if non-nil, zero value otherwise.

### GetSwiftCodeOk

`func (o *AccountCreation) GetSwiftCodeOk() (*string, bool)`

GetSwiftCodeOk returns a tuple with the SwiftCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSwiftCode

`func (o *AccountCreation) SetSwiftCode(v string)`

SetSwiftCode sets SwiftCode field to given value.

### HasSwiftCode

`func (o *AccountCreation) HasSwiftCode() bool`

HasSwiftCode returns a boolean if a field has been set.

### GetAccountTemplateId

`func (o *AccountCreation) GetAccountTemplateId() string`

GetAccountTemplateId returns the AccountTemplateId field if non-nil, zero value otherwise.

### GetAccountTemplateIdOk

`func (o *AccountCreation) GetAccountTemplateIdOk() (*string, bool)`

GetAccountTemplateIdOk returns a tuple with the AccountTemplateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountTemplateId

`func (o *AccountCreation) SetAccountTemplateId(v string)`

SetAccountTemplateId sets AccountTemplateId field to given value.

### HasAccountTemplateId

`func (o *AccountCreation) HasAccountTemplateId() bool`

HasAccountTemplateId returns a boolean if a field has been set.

### GetRelationships

`func (o *AccountCreation) GetRelationships() []Relationship`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *AccountCreation) GetRelationshipsOk() (*[]Relationship, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *AccountCreation) SetRelationships(v []Relationship)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *AccountCreation) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


