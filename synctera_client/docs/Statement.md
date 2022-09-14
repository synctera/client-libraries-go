# Statement

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountId** | Pointer to **string** | The unique identifier of the account the statement belongs to | [optional] [readonly] 
**DueDate** | Pointer to **string** | The limit date when the due amount indicated on the statement should be paid | [optional] [readonly] 
**EndDate** | Pointer to **string** | The date indicating the ending of the time interval covered by the statement | [optional] [readonly] 
**Id** | Pointer to **string** | statement ID | [optional] [readonly] 
**IssueDate** | Pointer to **string** | The date when the statement has been issued | [optional] [readonly] 
**StartDate** | Pointer to **string** | The date indicating the beginning of the time interval covered by the statement | [optional] [readonly] 
**AccountSummary** | Pointer to [**AccountSummary**](AccountSummary.md) |  | [optional] 
**AuthorizedSigner** | Pointer to [**[]Person1**](Person1.md) |  | [optional] [readonly] 
**Disclosure** | Pointer to **string** |  | [optional] 
**JointAccountHolders** | Pointer to [**[]Person1**](Person1.md) |  | [optional] [readonly] 
**PrimaryAccountHolderBusiness** | Pointer to [**Business1**](Business1.md) |  | [optional] 
**PrimaryAccountHolderPersonal** | Pointer to [**Person1**](Person1.md) |  | [optional] 
**Transactions** | Pointer to [**[]Transaction**](Transaction.md) |  | [optional] 
**SavingsSummary** | Pointer to [**SavingsSummary**](SavingsSummary.md) |  | [optional] 

## Methods

### NewStatement

`func NewStatement() *Statement`

NewStatement instantiates a new Statement object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStatementWithDefaults

`func NewStatementWithDefaults() *Statement`

NewStatementWithDefaults instantiates a new Statement object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountId

`func (o *Statement) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *Statement) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *Statement) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.

### HasAccountId

`func (o *Statement) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.

### GetDueDate

`func (o *Statement) GetDueDate() string`

GetDueDate returns the DueDate field if non-nil, zero value otherwise.

### GetDueDateOk

`func (o *Statement) GetDueDateOk() (*string, bool)`

GetDueDateOk returns a tuple with the DueDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDueDate

`func (o *Statement) SetDueDate(v string)`

SetDueDate sets DueDate field to given value.

### HasDueDate

`func (o *Statement) HasDueDate() bool`

HasDueDate returns a boolean if a field has been set.

### GetEndDate

`func (o *Statement) GetEndDate() string`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *Statement) GetEndDateOk() (*string, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *Statement) SetEndDate(v string)`

SetEndDate sets EndDate field to given value.

### HasEndDate

`func (o *Statement) HasEndDate() bool`

HasEndDate returns a boolean if a field has been set.

### GetId

`func (o *Statement) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Statement) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Statement) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Statement) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIssueDate

`func (o *Statement) GetIssueDate() string`

GetIssueDate returns the IssueDate field if non-nil, zero value otherwise.

### GetIssueDateOk

`func (o *Statement) GetIssueDateOk() (*string, bool)`

GetIssueDateOk returns a tuple with the IssueDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssueDate

`func (o *Statement) SetIssueDate(v string)`

SetIssueDate sets IssueDate field to given value.

### HasIssueDate

`func (o *Statement) HasIssueDate() bool`

HasIssueDate returns a boolean if a field has been set.

### GetStartDate

`func (o *Statement) GetStartDate() string`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *Statement) GetStartDateOk() (*string, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *Statement) SetStartDate(v string)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *Statement) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

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

### GetAuthorizedSigner

`func (o *Statement) GetAuthorizedSigner() []Person1`

GetAuthorizedSigner returns the AuthorizedSigner field if non-nil, zero value otherwise.

### GetAuthorizedSignerOk

`func (o *Statement) GetAuthorizedSignerOk() (*[]Person1, bool)`

GetAuthorizedSignerOk returns a tuple with the AuthorizedSigner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorizedSigner

`func (o *Statement) SetAuthorizedSigner(v []Person1)`

SetAuthorizedSigner sets AuthorizedSigner field to given value.

### HasAuthorizedSigner

`func (o *Statement) HasAuthorizedSigner() bool`

HasAuthorizedSigner returns a boolean if a field has been set.

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

### HasSavingsSummary

`func (o *Statement) HasSavingsSummary() bool`

HasSavingsSummary returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


