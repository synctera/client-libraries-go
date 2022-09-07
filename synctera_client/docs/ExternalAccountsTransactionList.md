# ExternalAccountsTransactionList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Transactions** | [**[]ExternalAccountTransaction**](ExternalAccountTransaction.md) | Array of transactions of a given external account | 

## Methods

### NewExternalAccountsTransactionList

`func NewExternalAccountsTransactionList(transactions []ExternalAccountTransaction, ) *ExternalAccountsTransactionList`

NewExternalAccountsTransactionList instantiates a new ExternalAccountsTransactionList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExternalAccountsTransactionListWithDefaults

`func NewExternalAccountsTransactionListWithDefaults() *ExternalAccountsTransactionList`

NewExternalAccountsTransactionListWithDefaults instantiates a new ExternalAccountsTransactionList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTransactions

`func (o *ExternalAccountsTransactionList) GetTransactions() []ExternalAccountTransaction`

GetTransactions returns the Transactions field if non-nil, zero value otherwise.

### GetTransactionsOk

`func (o *ExternalAccountsTransactionList) GetTransactionsOk() (*[]ExternalAccountTransaction, bool)`

GetTransactionsOk returns a tuple with the Transactions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactions

`func (o *ExternalAccountsTransactionList) SetTransactions(v []ExternalAccountTransaction)`

SetTransactions sets Transactions field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


