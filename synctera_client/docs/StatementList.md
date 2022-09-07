# StatementList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Statements** | [**[]StatementSummary**](StatementSummary.md) | Array of statements | 
**NextPageToken** | Pointer to **string** | If returned, use the next_page_token to query for the next page of results. Not returned if there are no more rows. | [optional] 

## Methods

### NewStatementList

`func NewStatementList(statements []StatementSummary, ) *StatementList`

NewStatementList instantiates a new StatementList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStatementListWithDefaults

`func NewStatementListWithDefaults() *StatementList`

NewStatementListWithDefaults instantiates a new StatementList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatements

`func (o *StatementList) GetStatements() []StatementSummary`

GetStatements returns the Statements field if non-nil, zero value otherwise.

### GetStatementsOk

`func (o *StatementList) GetStatementsOk() (*[]StatementSummary, bool)`

GetStatementsOk returns a tuple with the Statements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatements

`func (o *StatementList) SetStatements(v []StatementSummary)`

SetStatements sets Statements field to given value.


### GetNextPageToken

`func (o *StatementList) GetNextPageToken() string`

GetNextPageToken returns the NextPageToken field if non-nil, zero value otherwise.

### GetNextPageTokenOk

`func (o *StatementList) GetNextPageTokenOk() (*string, bool)`

GetNextPageTokenOk returns a tuple with the NextPageToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextPageToken

`func (o *StatementList) SetNextPageToken(v string)`

SetNextPageToken sets NextPageToken field to given value.

### HasNextPageToken

`func (o *StatementList) HasNextPageToken() bool`

HasNextPageToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


