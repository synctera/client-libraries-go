# AccountSummary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountNumber** | Pointer to **string** | Account number | [optional] 
**AccountStatus** | Pointer to **string** | Account Status | [optional] 
**AccountType** | Pointer to **string** | The type of the account. In lead mode, this always takes the value of the template. If not specified in shadow mode, CHECKING will be assumed.  | [optional] 
**BalanceCeiling** | Pointer to [**AccountSummaryBalanceCeiling**](AccountSummaryBalanceCeiling.md) |  | [optional] 
**BalanceFloor** | Pointer to [**AccountSummaryBalanceFloor**](AccountSummaryBalanceFloor.md) |  | [optional] 
**CreationTime** | Pointer to **time.Time** | Account creation time | [optional] 
**Currency** | Pointer to **string** | Account currency or account settlement currency. ISO 4217 alphabetic currency code. Default USD | [optional] 
**CustomerType** | Pointer to **string** | Customer type | [optional] 
**FinancialInstitution** | Pointer to [**FinancialInstitution**](FinancialInstitution.md) |  | [optional] 
**Id** | Pointer to **string** | The unique identifier of the account the statement belongs to | [optional] 
**LastUpdatedTime** | Pointer to **time.Time** | Account last modification time | [optional] 
**Nickname** | Pointer to **string** | User provided account nickname | [optional] 

## Methods

### NewAccountSummary

`func NewAccountSummary() *AccountSummary`

NewAccountSummary instantiates a new AccountSummary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountSummaryWithDefaults

`func NewAccountSummaryWithDefaults() *AccountSummary`

NewAccountSummaryWithDefaults instantiates a new AccountSummary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountNumber

`func (o *AccountSummary) GetAccountNumber() string`

GetAccountNumber returns the AccountNumber field if non-nil, zero value otherwise.

### GetAccountNumberOk

`func (o *AccountSummary) GetAccountNumberOk() (*string, bool)`

GetAccountNumberOk returns a tuple with the AccountNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountNumber

`func (o *AccountSummary) SetAccountNumber(v string)`

SetAccountNumber sets AccountNumber field to given value.

### HasAccountNumber

`func (o *AccountSummary) HasAccountNumber() bool`

HasAccountNumber returns a boolean if a field has been set.

### GetAccountStatus

`func (o *AccountSummary) GetAccountStatus() string`

GetAccountStatus returns the AccountStatus field if non-nil, zero value otherwise.

### GetAccountStatusOk

`func (o *AccountSummary) GetAccountStatusOk() (*string, bool)`

GetAccountStatusOk returns a tuple with the AccountStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountStatus

`func (o *AccountSummary) SetAccountStatus(v string)`

SetAccountStatus sets AccountStatus field to given value.

### HasAccountStatus

`func (o *AccountSummary) HasAccountStatus() bool`

HasAccountStatus returns a boolean if a field has been set.

### GetAccountType

`func (o *AccountSummary) GetAccountType() string`

GetAccountType returns the AccountType field if non-nil, zero value otherwise.

### GetAccountTypeOk

`func (o *AccountSummary) GetAccountTypeOk() (*string, bool)`

GetAccountTypeOk returns a tuple with the AccountType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountType

`func (o *AccountSummary) SetAccountType(v string)`

SetAccountType sets AccountType field to given value.

### HasAccountType

`func (o *AccountSummary) HasAccountType() bool`

HasAccountType returns a boolean if a field has been set.

### GetBalanceCeiling

`func (o *AccountSummary) GetBalanceCeiling() AccountSummaryBalanceCeiling`

GetBalanceCeiling returns the BalanceCeiling field if non-nil, zero value otherwise.

### GetBalanceCeilingOk

`func (o *AccountSummary) GetBalanceCeilingOk() (*AccountSummaryBalanceCeiling, bool)`

GetBalanceCeilingOk returns a tuple with the BalanceCeiling field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalanceCeiling

`func (o *AccountSummary) SetBalanceCeiling(v AccountSummaryBalanceCeiling)`

SetBalanceCeiling sets BalanceCeiling field to given value.

### HasBalanceCeiling

`func (o *AccountSummary) HasBalanceCeiling() bool`

HasBalanceCeiling returns a boolean if a field has been set.

### GetBalanceFloor

`func (o *AccountSummary) GetBalanceFloor() AccountSummaryBalanceFloor`

GetBalanceFloor returns the BalanceFloor field if non-nil, zero value otherwise.

### GetBalanceFloorOk

`func (o *AccountSummary) GetBalanceFloorOk() (*AccountSummaryBalanceFloor, bool)`

GetBalanceFloorOk returns a tuple with the BalanceFloor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalanceFloor

`func (o *AccountSummary) SetBalanceFloor(v AccountSummaryBalanceFloor)`

SetBalanceFloor sets BalanceFloor field to given value.

### HasBalanceFloor

`func (o *AccountSummary) HasBalanceFloor() bool`

HasBalanceFloor returns a boolean if a field has been set.

### GetCreationTime

`func (o *AccountSummary) GetCreationTime() time.Time`

GetCreationTime returns the CreationTime field if non-nil, zero value otherwise.

### GetCreationTimeOk

`func (o *AccountSummary) GetCreationTimeOk() (*time.Time, bool)`

GetCreationTimeOk returns a tuple with the CreationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationTime

`func (o *AccountSummary) SetCreationTime(v time.Time)`

SetCreationTime sets CreationTime field to given value.

### HasCreationTime

`func (o *AccountSummary) HasCreationTime() bool`

HasCreationTime returns a boolean if a field has been set.

### GetCurrency

`func (o *AccountSummary) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *AccountSummary) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *AccountSummary) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *AccountSummary) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetCustomerType

`func (o *AccountSummary) GetCustomerType() string`

GetCustomerType returns the CustomerType field if non-nil, zero value otherwise.

### GetCustomerTypeOk

`func (o *AccountSummary) GetCustomerTypeOk() (*string, bool)`

GetCustomerTypeOk returns a tuple with the CustomerType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerType

`func (o *AccountSummary) SetCustomerType(v string)`

SetCustomerType sets CustomerType field to given value.

### HasCustomerType

`func (o *AccountSummary) HasCustomerType() bool`

HasCustomerType returns a boolean if a field has been set.

### GetFinancialInstitution

`func (o *AccountSummary) GetFinancialInstitution() FinancialInstitution`

GetFinancialInstitution returns the FinancialInstitution field if non-nil, zero value otherwise.

### GetFinancialInstitutionOk

`func (o *AccountSummary) GetFinancialInstitutionOk() (*FinancialInstitution, bool)`

GetFinancialInstitutionOk returns a tuple with the FinancialInstitution field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFinancialInstitution

`func (o *AccountSummary) SetFinancialInstitution(v FinancialInstitution)`

SetFinancialInstitution sets FinancialInstitution field to given value.

### HasFinancialInstitution

`func (o *AccountSummary) HasFinancialInstitution() bool`

HasFinancialInstitution returns a boolean if a field has been set.

### GetId

`func (o *AccountSummary) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AccountSummary) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AccountSummary) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *AccountSummary) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLastUpdatedTime

`func (o *AccountSummary) GetLastUpdatedTime() time.Time`

GetLastUpdatedTime returns the LastUpdatedTime field if non-nil, zero value otherwise.

### GetLastUpdatedTimeOk

`func (o *AccountSummary) GetLastUpdatedTimeOk() (*time.Time, bool)`

GetLastUpdatedTimeOk returns a tuple with the LastUpdatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdatedTime

`func (o *AccountSummary) SetLastUpdatedTime(v time.Time)`

SetLastUpdatedTime sets LastUpdatedTime field to given value.

### HasLastUpdatedTime

`func (o *AccountSummary) HasLastUpdatedTime() bool`

HasLastUpdatedTime returns a boolean if a field has been set.

### GetNickname

`func (o *AccountSummary) GetNickname() string`

GetNickname returns the Nickname field if non-nil, zero value otherwise.

### GetNicknameOk

`func (o *AccountSummary) GetNicknameOk() (*string, bool)`

GetNicknameOk returns a tuple with the Nickname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNickname

`func (o *AccountSummary) SetNickname(v string)`

SetNickname sets Nickname field to given value.

### HasNickname

`func (o *AccountSummary) HasNickname() bool`

HasNickname returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


