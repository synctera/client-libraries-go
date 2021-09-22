# AccountCreation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessStatus** | Pointer to **string** | Access status for account. Default ACTIVE | [optional] 
**AccountNumber** | Pointer to **string** | Account number | [optional] 
**AccountTemplateId** | Pointer to **string** | Account template ID | [optional] 
**AccountTemplateVersion** | Pointer to **float32** | Account template version | [optional] 
**AccountType** | Pointer to [**AccountType**](AccountType.md) |  | [optional] 
**Balances** | Pointer to [**[]Balance**](Balance.md) | A list of balances for account based on different type | [optional] [readonly] 
**Currency** | Pointer to **string** | Account currency or account settlement currency. ISO 4217 alphabetic currency code. Default USD | [optional] 
**ExchangeRateType** | Pointer to **string** | Exchange rate type | [optional] 
**Iban** | Pointer to **string** | International bank account number | [optional] 
**Id** | Pointer to **string** | Account ID | [optional] [readonly] 
**IsAccountPool** | Pointer to **bool** | Account is investment (variable balance) account or a multi-balance account pool. Default false | [optional] 
**RecentTransactions** | Pointer to [**[]Transaction**](Transaction.md) | The most recent 10 transactions of the account | [optional] [readonly] 
**Status** | Pointer to [**Status**](Status.md) |  | [optional] 
**SwiftCode** | Pointer to **string** | SWIFT code | [optional] 
**Aliases** | Pointer to [**[]Alias**](Alias.md) | A list of the aliases for account. Account alias is the account number of different balance types to link to the same account ID | [optional] 
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

`func (o *AccountCreation) GetAccessStatus() string`

GetAccessStatus returns the AccessStatus field if non-nil, zero value otherwise.

### GetAccessStatusOk

`func (o *AccountCreation) GetAccessStatusOk() (*string, bool)`

GetAccessStatusOk returns a tuple with the AccessStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessStatus

`func (o *AccountCreation) SetAccessStatus(v string)`

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

### GetAccountTemplateVersion

`func (o *AccountCreation) GetAccountTemplateVersion() float32`

GetAccountTemplateVersion returns the AccountTemplateVersion field if non-nil, zero value otherwise.

### GetAccountTemplateVersionOk

`func (o *AccountCreation) GetAccountTemplateVersionOk() (*float32, bool)`

GetAccountTemplateVersionOk returns a tuple with the AccountTemplateVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountTemplateVersion

`func (o *AccountCreation) SetAccountTemplateVersion(v float32)`

SetAccountTemplateVersion sets AccountTemplateVersion field to given value.

### HasAccountTemplateVersion

`func (o *AccountCreation) HasAccountTemplateVersion() bool`

HasAccountTemplateVersion returns a boolean if a field has been set.

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

### GetRecentTransactions

`func (o *AccountCreation) GetRecentTransactions() []Transaction`

GetRecentTransactions returns the RecentTransactions field if non-nil, zero value otherwise.

### GetRecentTransactionsOk

`func (o *AccountCreation) GetRecentTransactionsOk() (*[]Transaction, bool)`

GetRecentTransactionsOk returns a tuple with the RecentTransactions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecentTransactions

`func (o *AccountCreation) SetRecentTransactions(v []Transaction)`

SetRecentTransactions sets RecentTransactions field to given value.

### HasRecentTransactions

`func (o *AccountCreation) HasRecentTransactions() bool`

HasRecentTransactions returns a boolean if a field has been set.

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

### GetAliases

`func (o *AccountCreation) GetAliases() []Alias`

GetAliases returns the Aliases field if non-nil, zero value otherwise.

### GetAliasesOk

`func (o *AccountCreation) GetAliasesOk() (*[]Alias, bool)`

GetAliasesOk returns a tuple with the Aliases field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAliases

`func (o *AccountCreation) SetAliases(v []Alias)`

SetAliases sets Aliases field to given value.

### HasAliases

`func (o *AccountCreation) HasAliases() bool`

HasAliases returns a boolean if a field has been set.

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


