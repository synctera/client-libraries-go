# InternalAccount

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountNumber** | Pointer to **string** | Generated internal account number | [optional] [readonly] 
**Balances** | Pointer to [**[]Balance**](Balance.md) | A list of balances for internal account based on different type | [optional] [readonly] 
**BankRouting** | Pointer to **string** | Bank routing number | [optional] [readonly] 
**CreationTime** | Pointer to **time.Time** |  | [optional] [readonly] 
**Currency** | **string** | Account currency or account settlement currency. ISO 4217 alphabetic currency code. | 
**Description** | Pointer to **string** | A user provided description for the current account | [optional] 
**Id** | Pointer to **string** | Generated ID for internal account | [optional] [readonly] 
**LastUpdatedTime** | Pointer to **time.Time** |  | [optional] [readonly] 
**Status** | **string** |  | 

## Methods

### NewInternalAccount

`func NewInternalAccount(currency string, status string, ) *InternalAccount`

NewInternalAccount instantiates a new InternalAccount object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInternalAccountWithDefaults

`func NewInternalAccountWithDefaults() *InternalAccount`

NewInternalAccountWithDefaults instantiates a new InternalAccount object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountNumber

`func (o *InternalAccount) GetAccountNumber() string`

GetAccountNumber returns the AccountNumber field if non-nil, zero value otherwise.

### GetAccountNumberOk

`func (o *InternalAccount) GetAccountNumberOk() (*string, bool)`

GetAccountNumberOk returns a tuple with the AccountNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountNumber

`func (o *InternalAccount) SetAccountNumber(v string)`

SetAccountNumber sets AccountNumber field to given value.

### HasAccountNumber

`func (o *InternalAccount) HasAccountNumber() bool`

HasAccountNumber returns a boolean if a field has been set.

### GetBalances

`func (o *InternalAccount) GetBalances() []Balance`

GetBalances returns the Balances field if non-nil, zero value otherwise.

### GetBalancesOk

`func (o *InternalAccount) GetBalancesOk() (*[]Balance, bool)`

GetBalancesOk returns a tuple with the Balances field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalances

`func (o *InternalAccount) SetBalances(v []Balance)`

SetBalances sets Balances field to given value.

### HasBalances

`func (o *InternalAccount) HasBalances() bool`

HasBalances returns a boolean if a field has been set.

### GetBankRouting

`func (o *InternalAccount) GetBankRouting() string`

GetBankRouting returns the BankRouting field if non-nil, zero value otherwise.

### GetBankRoutingOk

`func (o *InternalAccount) GetBankRoutingOk() (*string, bool)`

GetBankRoutingOk returns a tuple with the BankRouting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBankRouting

`func (o *InternalAccount) SetBankRouting(v string)`

SetBankRouting sets BankRouting field to given value.

### HasBankRouting

`func (o *InternalAccount) HasBankRouting() bool`

HasBankRouting returns a boolean if a field has been set.

### GetCreationTime

`func (o *InternalAccount) GetCreationTime() time.Time`

GetCreationTime returns the CreationTime field if non-nil, zero value otherwise.

### GetCreationTimeOk

`func (o *InternalAccount) GetCreationTimeOk() (*time.Time, bool)`

GetCreationTimeOk returns a tuple with the CreationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationTime

`func (o *InternalAccount) SetCreationTime(v time.Time)`

SetCreationTime sets CreationTime field to given value.

### HasCreationTime

`func (o *InternalAccount) HasCreationTime() bool`

HasCreationTime returns a boolean if a field has been set.

### GetCurrency

`func (o *InternalAccount) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *InternalAccount) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *InternalAccount) SetCurrency(v string)`

SetCurrency sets Currency field to given value.


### GetDescription

`func (o *InternalAccount) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *InternalAccount) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *InternalAccount) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *InternalAccount) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetId

`func (o *InternalAccount) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *InternalAccount) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *InternalAccount) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *InternalAccount) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLastUpdatedTime

`func (o *InternalAccount) GetLastUpdatedTime() time.Time`

GetLastUpdatedTime returns the LastUpdatedTime field if non-nil, zero value otherwise.

### GetLastUpdatedTimeOk

`func (o *InternalAccount) GetLastUpdatedTimeOk() (*time.Time, bool)`

GetLastUpdatedTimeOk returns a tuple with the LastUpdatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdatedTime

`func (o *InternalAccount) SetLastUpdatedTime(v time.Time)`

SetLastUpdatedTime sets LastUpdatedTime field to given value.

### HasLastUpdatedTime

`func (o *InternalAccount) HasLastUpdatedTime() bool`

HasLastUpdatedTime returns a boolean if a field has been set.

### GetStatus

`func (o *InternalAccount) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *InternalAccount) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *InternalAccount) SetStatus(v string)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


