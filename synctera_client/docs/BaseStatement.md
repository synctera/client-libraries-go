# BaseStatement

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

## Methods

### NewBaseStatement

`func NewBaseStatement() *BaseStatement`

NewBaseStatement instantiates a new BaseStatement object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBaseStatementWithDefaults

`func NewBaseStatementWithDefaults() *BaseStatement`

NewBaseStatementWithDefaults instantiates a new BaseStatement object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountId

`func (o *BaseStatement) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *BaseStatement) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *BaseStatement) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.

### HasAccountId

`func (o *BaseStatement) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.

### GetDueDate

`func (o *BaseStatement) GetDueDate() string`

GetDueDate returns the DueDate field if non-nil, zero value otherwise.

### GetDueDateOk

`func (o *BaseStatement) GetDueDateOk() (*string, bool)`

GetDueDateOk returns a tuple with the DueDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDueDate

`func (o *BaseStatement) SetDueDate(v string)`

SetDueDate sets DueDate field to given value.

### HasDueDate

`func (o *BaseStatement) HasDueDate() bool`

HasDueDate returns a boolean if a field has been set.

### GetEndDate

`func (o *BaseStatement) GetEndDate() string`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *BaseStatement) GetEndDateOk() (*string, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *BaseStatement) SetEndDate(v string)`

SetEndDate sets EndDate field to given value.

### HasEndDate

`func (o *BaseStatement) HasEndDate() bool`

HasEndDate returns a boolean if a field has been set.

### GetId

`func (o *BaseStatement) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BaseStatement) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BaseStatement) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *BaseStatement) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIssueDate

`func (o *BaseStatement) GetIssueDate() string`

GetIssueDate returns the IssueDate field if non-nil, zero value otherwise.

### GetIssueDateOk

`func (o *BaseStatement) GetIssueDateOk() (*string, bool)`

GetIssueDateOk returns a tuple with the IssueDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssueDate

`func (o *BaseStatement) SetIssueDate(v string)`

SetIssueDate sets IssueDate field to given value.

### HasIssueDate

`func (o *BaseStatement) HasIssueDate() bool`

HasIssueDate returns a boolean if a field has been set.

### GetStartDate

`func (o *BaseStatement) GetStartDate() string`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *BaseStatement) GetStartDateOk() (*string, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *BaseStatement) SetStartDate(v string)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *BaseStatement) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### GetAccountSummary

`func (o *BaseStatement) GetAccountSummary() AccountSummary`

GetAccountSummary returns the AccountSummary field if non-nil, zero value otherwise.

### GetAccountSummaryOk

`func (o *BaseStatement) GetAccountSummaryOk() (*AccountSummary, bool)`

GetAccountSummaryOk returns a tuple with the AccountSummary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountSummary

`func (o *BaseStatement) SetAccountSummary(v AccountSummary)`

SetAccountSummary sets AccountSummary field to given value.

### HasAccountSummary

`func (o *BaseStatement) HasAccountSummary() bool`

HasAccountSummary returns a boolean if a field has been set.

### GetAuthorizedSigner

`func (o *BaseStatement) GetAuthorizedSigner() []Person1`

GetAuthorizedSigner returns the AuthorizedSigner field if non-nil, zero value otherwise.

### GetAuthorizedSignerOk

`func (o *BaseStatement) GetAuthorizedSignerOk() (*[]Person1, bool)`

GetAuthorizedSignerOk returns a tuple with the AuthorizedSigner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorizedSigner

`func (o *BaseStatement) SetAuthorizedSigner(v []Person1)`

SetAuthorizedSigner sets AuthorizedSigner field to given value.

### HasAuthorizedSigner

`func (o *BaseStatement) HasAuthorizedSigner() bool`

HasAuthorizedSigner returns a boolean if a field has been set.

### GetDisclosure

`func (o *BaseStatement) GetDisclosure() string`

GetDisclosure returns the Disclosure field if non-nil, zero value otherwise.

### GetDisclosureOk

`func (o *BaseStatement) GetDisclosureOk() (*string, bool)`

GetDisclosureOk returns a tuple with the Disclosure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisclosure

`func (o *BaseStatement) SetDisclosure(v string)`

SetDisclosure sets Disclosure field to given value.

### HasDisclosure

`func (o *BaseStatement) HasDisclosure() bool`

HasDisclosure returns a boolean if a field has been set.

### GetJointAccountHolders

`func (o *BaseStatement) GetJointAccountHolders() []Person1`

GetJointAccountHolders returns the JointAccountHolders field if non-nil, zero value otherwise.

### GetJointAccountHoldersOk

`func (o *BaseStatement) GetJointAccountHoldersOk() (*[]Person1, bool)`

GetJointAccountHoldersOk returns a tuple with the JointAccountHolders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJointAccountHolders

`func (o *BaseStatement) SetJointAccountHolders(v []Person1)`

SetJointAccountHolders sets JointAccountHolders field to given value.

### HasJointAccountHolders

`func (o *BaseStatement) HasJointAccountHolders() bool`

HasJointAccountHolders returns a boolean if a field has been set.

### GetPrimaryAccountHolderBusiness

`func (o *BaseStatement) GetPrimaryAccountHolderBusiness() Business1`

GetPrimaryAccountHolderBusiness returns the PrimaryAccountHolderBusiness field if non-nil, zero value otherwise.

### GetPrimaryAccountHolderBusinessOk

`func (o *BaseStatement) GetPrimaryAccountHolderBusinessOk() (*Business1, bool)`

GetPrimaryAccountHolderBusinessOk returns a tuple with the PrimaryAccountHolderBusiness field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryAccountHolderBusiness

`func (o *BaseStatement) SetPrimaryAccountHolderBusiness(v Business1)`

SetPrimaryAccountHolderBusiness sets PrimaryAccountHolderBusiness field to given value.

### HasPrimaryAccountHolderBusiness

`func (o *BaseStatement) HasPrimaryAccountHolderBusiness() bool`

HasPrimaryAccountHolderBusiness returns a boolean if a field has been set.

### GetPrimaryAccountHolderPersonal

`func (o *BaseStatement) GetPrimaryAccountHolderPersonal() Person1`

GetPrimaryAccountHolderPersonal returns the PrimaryAccountHolderPersonal field if non-nil, zero value otherwise.

### GetPrimaryAccountHolderPersonalOk

`func (o *BaseStatement) GetPrimaryAccountHolderPersonalOk() (*Person1, bool)`

GetPrimaryAccountHolderPersonalOk returns a tuple with the PrimaryAccountHolderPersonal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryAccountHolderPersonal

`func (o *BaseStatement) SetPrimaryAccountHolderPersonal(v Person1)`

SetPrimaryAccountHolderPersonal sets PrimaryAccountHolderPersonal field to given value.

### HasPrimaryAccountHolderPersonal

`func (o *BaseStatement) HasPrimaryAccountHolderPersonal() bool`

HasPrimaryAccountHolderPersonal returns a boolean if a field has been set.

### GetTransactions

`func (o *BaseStatement) GetTransactions() []Transaction`

GetTransactions returns the Transactions field if non-nil, zero value otherwise.

### GetTransactionsOk

`func (o *BaseStatement) GetTransactionsOk() (*[]Transaction, bool)`

GetTransactionsOk returns a tuple with the Transactions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactions

`func (o *BaseStatement) SetTransactions(v []Transaction)`

SetTransactions sets Transactions field to given value.

### HasTransactions

`func (o *BaseStatement) HasTransactions() bool`

HasTransactions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


