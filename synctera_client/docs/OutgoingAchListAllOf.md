# OutgoingAchListAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Transactions** | [**[]OutgoingAch**](OutgoingAch.md) | Array of sent ACH transactions. | 

## Methods

### NewOutgoingAchListAllOf

`func NewOutgoingAchListAllOf(transactions []OutgoingAch, ) *OutgoingAchListAllOf`

NewOutgoingAchListAllOf instantiates a new OutgoingAchListAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOutgoingAchListAllOfWithDefaults

`func NewOutgoingAchListAllOfWithDefaults() *OutgoingAchListAllOf`

NewOutgoingAchListAllOfWithDefaults instantiates a new OutgoingAchListAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTransactions

`func (o *OutgoingAchListAllOf) GetTransactions() []OutgoingAch`

GetTransactions returns the Transactions field if non-nil, zero value otherwise.

### GetTransactionsOk

`func (o *OutgoingAchListAllOf) GetTransactionsOk() (*[]OutgoingAch, bool)`

GetTransactionsOk returns a tuple with the Transactions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactions

`func (o *OutgoingAchListAllOf) SetTransactions(v []OutgoingAch)`

SetTransactions sets Transactions field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


