# Statement

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountSummary** | Pointer to [**AccountSummary**](AccountSummary.md) |  | [optional] 
**AuthorizedSigners** | Pointer to [**[]Person1**](Person1.md) |  | [optional] [readonly] 
**Disclosure** | Pointer to **string** |  | [optional] 
**JointAccountHolders** | Pointer to [**[]Person1**](Person1.md) |  | [optional] [readonly] 
**PrimaryAccountHolderBusiness** | Pointer to [**Business1**](Business1.md) |  | [optional] 
**PrimaryAccountHolderPersonal** | Pointer to [**Person1**](Person1.md) |  | [optional] 
**StatementSummary** | Pointer to [**StatementSummary**](StatementSummary.md) |  | [optional] 
**Transactions** | Pointer to [**[]Transaction**](Transaction.md) |  | [optional] 
**SavingsSummary** | [**SavingsSummary**](SavingsSummary.md) |  | 

## Methods

### NewStatement

`func NewStatement(savingsSummary SavingsSummary, ) *Statement`

NewStatement instantiates a new Statement object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStatementWithDefaults

`func NewStatementWithDefaults() *Statement`

NewStatementWithDefaults instantiates a new Statement object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountSummary

`func (o *Statement) GetAccountSummary() AccountSummary`

GetAccountSummary returns the AccountSummary field if non-nil, zero value otherwise.

### GetAccountSummaryOk

`func (o *Statement) GetAccountSummaryOk() (*AccountSummary, bool)`

GetAccountSummaryOk returns a tuple with the AccountSummary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountSummary

`func (o *Statement) SetAccountSummary(v AccountSummary)`

SetAccountSummary sets AccountSummary field to given value.

### HasAccountSummary

`func (o *Statement) HasAccountSummary() bool`

HasAccountSummary returns a boolean if a field has been set.

### GetAuthorizedSigners

`func (o *Statement) GetAuthorizedSigners() []Person1`

GetAuthorizedSigners returns the AuthorizedSigners field if non-nil, zero value otherwise.

### GetAuthorizedSignersOk

`func (o *Statement) GetAuthorizedSignersOk() (*[]Person1, bool)`

GetAuthorizedSignersOk returns a tuple with the AuthorizedSigners field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorizedSigners

`func (o *Statement) SetAuthorizedSigners(v []Person1)`

SetAuthorizedSigners sets AuthorizedSigners field to given value.

### HasAuthorizedSigners

`func (o *Statement) HasAuthorizedSigners() bool`

HasAuthorizedSigners returns a boolean if a field has been set.

### GetDisclosure

`func (o *Statement) GetDisclosure() string`

GetDisclosure returns the Disclosure field if non-nil, zero value otherwise.

### GetDisclosureOk

`func (o *Statement) GetDisclosureOk() (*string, bool)`

GetDisclosureOk returns a tuple with the Disclosure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisclosure

`func (o *Statement) SetDisclosure(v string)`

SetDisclosure sets Disclosure field to given value.

### HasDisclosure

`func (o *Statement) HasDisclosure() bool`

HasDisclosure returns a boolean if a field has been set.

### GetJointAccountHolders

`func (o *Statement) GetJointAccountHolders() []Person1`

GetJointAccountHolders returns the JointAccountHolders field if non-nil, zero value otherwise.

### GetJointAccountHoldersOk

`func (o *Statement) GetJointAccountHoldersOk() (*[]Person1, bool)`

GetJointAccountHoldersOk returns a tuple with the JointAccountHolders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJointAccountHolders

`func (o *Statement) SetJointAccountHolders(v []Person1)`

SetJointAccountHolders sets JointAccountHolders field to given value.

### HasJointAccountHolders

`func (o *Statement) HasJointAccountHolders() bool`

HasJointAccountHolders returns a boolean if a field has been set.

### GetPrimaryAccountHolderBusiness

`func (o *Statement) GetPrimaryAccountHolderBusiness() Business1`

GetPrimaryAccountHolderBusiness returns the PrimaryAccountHolderBusiness field if non-nil, zero value otherwise.

### GetPrimaryAccountHolderBusinessOk

`func (o *Statement) GetPrimaryAccountHolderBusinessOk() (*Business1, bool)`

GetPrimaryAccountHolderBusinessOk returns a tuple with the PrimaryAccountHolderBusiness field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryAccountHolderBusiness

`func (o *Statement) SetPrimaryAccountHolderBusiness(v Business1)`

SetPrimaryAccountHolderBusiness sets PrimaryAccountHolderBusiness field to given value.

### HasPrimaryAccountHolderBusiness

`func (o *Statement) HasPrimaryAccountHolderBusiness() bool`

HasPrimaryAccountHolderBusiness returns a boolean if a field has been set.

### GetPrimaryAccountHolderPersonal

`func (o *Statement) GetPrimaryAccountHolderPersonal() Person1`

GetPrimaryAccountHolderPersonal returns the PrimaryAccountHolderPersonal field if non-nil, zero value otherwise.

### GetPrimaryAccountHolderPersonalOk

`func (o *Statement) GetPrimaryAccountHolderPersonalOk() (*Person1, bool)`

GetPrimaryAccountHolderPersonalOk returns a tuple with the PrimaryAccountHolderPersonal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryAccountHolderPersonal

`func (o *Statement) SetPrimaryAccountHolderPersonal(v Person1)`

SetPrimaryAccountHolderPersonal sets PrimaryAccountHolderPersonal field to given value.

### HasPrimaryAccountHolderPersonal

`func (o *Statement) HasPrimaryAccountHolderPersonal() bool`

HasPrimaryAccountHolderPersonal returns a boolean if a field has been set.

### GetStatementSummary

`func (o *Statement) GetStatementSummary() StatementSummary`

GetStatementSummary returns the StatementSummary field if non-nil, zero value otherwise.

### GetStatementSummaryOk

`func (o *Statement) GetStatementSummaryOk() (*StatementSummary, bool)`

GetStatementSummaryOk returns a tuple with the StatementSummary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatementSummary

`func (o *Statement) SetStatementSummary(v StatementSummary)`

SetStatementSummary sets StatementSummary field to given value.

### HasStatementSummary

`func (o *Statement) HasStatementSummary() bool`

HasStatementSummary returns a boolean if a field has been set.

### GetTransactions

`func (o *Statement) GetTransactions() []Transaction`

GetTransactions returns the Transactions field if non-nil, zero value otherwise.

### GetTransactionsOk

`func (o *Statement) GetTransactionsOk() (*[]Transaction, bool)`

GetTransactionsOk returns a tuple with the Transactions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactions

`func (o *Statement) SetTransactions(v []Transaction)`

SetTransactions sets Transactions field to given value.

### HasTransactions

`func (o *Statement) HasTransactions() bool`

HasTransactions returns a boolean if a field has been set.

### GetSavingsSummary

`func (o *Statement) GetSavingsSummary() SavingsSummary`

GetSavingsSummary returns the SavingsSummary field if non-nil, zero value otherwise.

### GetSavingsSummaryOk

`func (o *Statement) GetSavingsSummaryOk() (*SavingsSummary, bool)`

GetSavingsSummaryOk returns a tuple with the SavingsSummary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSavingsSummary

`func (o *Statement) SetSavingsSummary(v SavingsSummary)`

SetSavingsSummary sets SavingsSummary field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


