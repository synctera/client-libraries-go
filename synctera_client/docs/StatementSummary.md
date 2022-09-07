# StatementSummary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountId** | Pointer to **string** | The unique identifier of the account the statement belongs to | [optional] [readonly] 
**DueDate** | Pointer to **string** | The limit date when the due amount indicated on the statement should be paid | [optional] [readonly] 
**EndDate** | Pointer to **string** | The date indicating the ending of the time interval covered by the statement | [optional] [readonly] 
**Id** | Pointer to **string** | statement ID | [optional] [readonly] 
**IssueDate** | Pointer to **string** | The date when the statement has been issued | [optional] [readonly] 
**StartDate** | Pointer to **string** | The date indicating the beginning of the time interval covered by the statement | [optional] [readonly] 

## Methods

### NewStatementSummary

`func NewStatementSummary() *StatementSummary`

NewStatementSummary instantiates a new StatementSummary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStatementSummaryWithDefaults

`func NewStatementSummaryWithDefaults() *StatementSummary`

NewStatementSummaryWithDefaults instantiates a new StatementSummary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountId

`func (o *StatementSummary) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *StatementSummary) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *StatementSummary) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.

### HasAccountId

`func (o *StatementSummary) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.

### GetDueDate

`func (o *StatementSummary) GetDueDate() string`

GetDueDate returns the DueDate field if non-nil, zero value otherwise.

### GetDueDateOk

`func (o *StatementSummary) GetDueDateOk() (*string, bool)`

GetDueDateOk returns a tuple with the DueDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDueDate

`func (o *StatementSummary) SetDueDate(v string)`

SetDueDate sets DueDate field to given value.

### HasDueDate

`func (o *StatementSummary) HasDueDate() bool`

HasDueDate returns a boolean if a field has been set.

### GetEndDate

`func (o *StatementSummary) GetEndDate() string`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *StatementSummary) GetEndDateOk() (*string, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *StatementSummary) SetEndDate(v string)`

SetEndDate sets EndDate field to given value.

### HasEndDate

`func (o *StatementSummary) HasEndDate() bool`

HasEndDate returns a boolean if a field has been set.

### GetId

`func (o *StatementSummary) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *StatementSummary) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *StatementSummary) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *StatementSummary) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIssueDate

`func (o *StatementSummary) GetIssueDate() string`

GetIssueDate returns the IssueDate field if non-nil, zero value otherwise.

### GetIssueDateOk

`func (o *StatementSummary) GetIssueDateOk() (*string, bool)`

GetIssueDateOk returns a tuple with the IssueDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssueDate

`func (o *StatementSummary) SetIssueDate(v string)`

SetIssueDate sets IssueDate field to given value.

### HasIssueDate

`func (o *StatementSummary) HasIssueDate() bool`

HasIssueDate returns a boolean if a field has been set.

### GetStartDate

`func (o *StatementSummary) GetStartDate() string`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *StatementSummary) GetStartDateOk() (*string, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *StatementSummary) SetStartDate(v string)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *StatementSummary) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


