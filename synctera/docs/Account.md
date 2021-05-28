# Account

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Account ID | [optional] [readonly] 
**AccountNumber** | Pointer to **string** | Account number. Required for shadow mode. | [optional] [readonly] 
**Currency** | Pointer to **string** | Account currency or account settlement currency. ISO 4217 alphabetic currency code. Default USD | [optional] 
**Status** | Pointer to [**Status**](Status.md) |  | [optional] 
**ExchangeRateType** | Pointer to **string** | Exchange rate type | [optional] 
**Iban** | Pointer to **string** | International bank account number | [optional] 
**SwiftCode** | Pointer to **string** | SWIFT code | [optional] 
**IsAccountPool** | Pointer to **bool** | Account is investment (variable balance) account or a multi-balance account pool. Default false | [optional] 
**AccountTemplateCode** | Pointer to **string** | Account template code. Required for lead mode. | [optional] 
**AccountTemplateVersion** | Pointer to **float32** | Account template version. Required for lead mode. | [optional] 
**Relationships** | Pointer to [**[]SchemasRelationship**](SchemasRelationship.md) | List of the relationship for this account to the parties. Required for lead mode. | [optional] 
**Aliases** | Pointer to [**[]Alias**](Alias.md) | A list of the aliases for account. Account alias is the account number of different balance types to link to the same account ID | [optional] 
**Balances** | Pointer to [**[]Balance**](Balance.md) | A list of balances for account based on different type | [optional] [readonly] 
**RecentTransactions** | Pointer to [**[]Transaction**](Transaction.md) | The most recent 10 transactions of the account | [optional] [readonly] 

## Methods

### NewAccount

`func NewAccount() *Account`

NewAccount instantiates a new Account object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountWithDefaults

`func NewAccountWithDefaults() *Account`

NewAccountWithDefaults instantiates a new Account object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Account) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Account) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Account) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Account) HasId() bool`

HasId returns a boolean if a field has been set.

### GetAccountNumber

`func (o *Account) GetAccountNumber() string`

GetAccountNumber returns the AccountNumber field if non-nil, zero value otherwise.

### GetAccountNumberOk

`func (o *Account) GetAccountNumberOk() (*string, bool)`

GetAccountNumberOk returns a tuple with the AccountNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountNumber

`func (o *Account) SetAccountNumber(v string)`

SetAccountNumber sets AccountNumber field to given value.

### HasAccountNumber

`func (o *Account) HasAccountNumber() bool`

HasAccountNumber returns a boolean if a field has been set.

### GetCurrency

`func (o *Account) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *Account) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *Account) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *Account) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetStatus

`func (o *Account) GetStatus() Status`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Account) GetStatusOk() (*Status, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Account) SetStatus(v Status)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *Account) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetExchangeRateType

`func (o *Account) GetExchangeRateType() string`

GetExchangeRateType returns the ExchangeRateType field if non-nil, zero value otherwise.

### GetExchangeRateTypeOk

`func (o *Account) GetExchangeRateTypeOk() (*string, bool)`

GetExchangeRateTypeOk returns a tuple with the ExchangeRateType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExchangeRateType

`func (o *Account) SetExchangeRateType(v string)`

SetExchangeRateType sets ExchangeRateType field to given value.

### HasExchangeRateType

`func (o *Account) HasExchangeRateType() bool`

HasExchangeRateType returns a boolean if a field has been set.

### GetIban

`func (o *Account) GetIban() string`

GetIban returns the Iban field if non-nil, zero value otherwise.

### GetIbanOk

`func (o *Account) GetIbanOk() (*string, bool)`

GetIbanOk returns a tuple with the Iban field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIban

`func (o *Account) SetIban(v string)`

SetIban sets Iban field to given value.

### HasIban

`func (o *Account) HasIban() bool`

HasIban returns a boolean if a field has been set.

### GetSwiftCode

`func (o *Account) GetSwiftCode() string`

GetSwiftCode returns the SwiftCode field if non-nil, zero value otherwise.

### GetSwiftCodeOk

`func (o *Account) GetSwiftCodeOk() (*string, bool)`

GetSwiftCodeOk returns a tuple with the SwiftCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSwiftCode

`func (o *Account) SetSwiftCode(v string)`

SetSwiftCode sets SwiftCode field to given value.

### HasSwiftCode

`func (o *Account) HasSwiftCode() bool`

HasSwiftCode returns a boolean if a field has been set.

### GetIsAccountPool

`func (o *Account) GetIsAccountPool() bool`

GetIsAccountPool returns the IsAccountPool field if non-nil, zero value otherwise.

### GetIsAccountPoolOk

`func (o *Account) GetIsAccountPoolOk() (*bool, bool)`

GetIsAccountPoolOk returns a tuple with the IsAccountPool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAccountPool

`func (o *Account) SetIsAccountPool(v bool)`

SetIsAccountPool sets IsAccountPool field to given value.

### HasIsAccountPool

`func (o *Account) HasIsAccountPool() bool`

HasIsAccountPool returns a boolean if a field has been set.

### GetAccountTemplateCode

`func (o *Account) GetAccountTemplateCode() string`

GetAccountTemplateCode returns the AccountTemplateCode field if non-nil, zero value otherwise.

### GetAccountTemplateCodeOk

`func (o *Account) GetAccountTemplateCodeOk() (*string, bool)`

GetAccountTemplateCodeOk returns a tuple with the AccountTemplateCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountTemplateCode

`func (o *Account) SetAccountTemplateCode(v string)`

SetAccountTemplateCode sets AccountTemplateCode field to given value.

### HasAccountTemplateCode

`func (o *Account) HasAccountTemplateCode() bool`

HasAccountTemplateCode returns a boolean if a field has been set.

### GetAccountTemplateVersion

`func (o *Account) GetAccountTemplateVersion() float32`

GetAccountTemplateVersion returns the AccountTemplateVersion field if non-nil, zero value otherwise.

### GetAccountTemplateVersionOk

`func (o *Account) GetAccountTemplateVersionOk() (*float32, bool)`

GetAccountTemplateVersionOk returns a tuple with the AccountTemplateVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountTemplateVersion

`func (o *Account) SetAccountTemplateVersion(v float32)`

SetAccountTemplateVersion sets AccountTemplateVersion field to given value.

### HasAccountTemplateVersion

`func (o *Account) HasAccountTemplateVersion() bool`

HasAccountTemplateVersion returns a boolean if a field has been set.

### GetRelationships

`func (o *Account) GetRelationships() []SchemasRelationship`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *Account) GetRelationshipsOk() (*[]SchemasRelationship, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *Account) SetRelationships(v []SchemasRelationship)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *Account) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.

### GetAliases

`func (o *Account) GetAliases() []Alias`

GetAliases returns the Aliases field if non-nil, zero value otherwise.

### GetAliasesOk

`func (o *Account) GetAliasesOk() (*[]Alias, bool)`

GetAliasesOk returns a tuple with the Aliases field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAliases

`func (o *Account) SetAliases(v []Alias)`

SetAliases sets Aliases field to given value.

### HasAliases

`func (o *Account) HasAliases() bool`

HasAliases returns a boolean if a field has been set.

### GetBalances

`func (o *Account) GetBalances() []Balance`

GetBalances returns the Balances field if non-nil, zero value otherwise.

### GetBalancesOk

`func (o *Account) GetBalancesOk() (*[]Balance, bool)`

GetBalancesOk returns a tuple with the Balances field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalances

`func (o *Account) SetBalances(v []Balance)`

SetBalances sets Balances field to given value.

### HasBalances

`func (o *Account) HasBalances() bool`

HasBalances returns a boolean if a field has been set.

### GetRecentTransactions

`func (o *Account) GetRecentTransactions() []Transaction`

GetRecentTransactions returns the RecentTransactions field if non-nil, zero value otherwise.

### GetRecentTransactionsOk

`func (o *Account) GetRecentTransactionsOk() (*[]Transaction, bool)`

GetRecentTransactionsOk returns a tuple with the RecentTransactions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecentTransactions

`func (o *Account) SetRecentTransactions(v []Transaction)`

SetRecentTransactions sets RecentTransactions field to given value.

### HasRecentTransactions

`func (o *Account) HasRecentTransactions() bool`

HasRecentTransactions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


