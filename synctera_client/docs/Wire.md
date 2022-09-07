# Wire

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | **int32** | Transfer amount in cents ($100 would be 10000) | 
**BankMessage** | Pointer to **string** | Instructions intended for the financial institutions that are processing the wire. | [optional] 
**CreationTime** | **time.Time** |  | [readonly] 
**Currency** | **string** | 3-character currency code | 
**CustomerId** | **string** | The customer UUID representing the person initiating the Wire transfer | 
**Id** | **string** | wire ID | [readonly] 
**LastUpdatedTime** | **time.Time** |  | [readonly] 
**OriginatingAccountId** | **string** | Sender account ID | 
**ReceivingAccountId** | **string** | The external account uuid representing the recipient of the wire. | 
**RecipientMessage** | Pointer to **string** | Information from the originator to the beneficiary (recipient). | [optional] 
**SenderReferenceId** | **string** | Sender&#39;s id associated with fedwire transfer | [readonly] 
**Status** | **string** | The current status of the transfer | [readonly] 
**TransactionId** | **string** | ID of the resulting transaction resource | [readonly] 

## Methods

### NewWire

`func NewWire(amount int32, creationTime time.Time, currency string, customerId string, id string, lastUpdatedTime time.Time, originatingAccountId string, receivingAccountId string, senderReferenceId string, status string, transactionId string, ) *Wire`

NewWire instantiates a new Wire object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWireWithDefaults

`func NewWireWithDefaults() *Wire`

NewWireWithDefaults instantiates a new Wire object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmount

`func (o *Wire) GetAmount() int32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *Wire) GetAmountOk() (*int32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *Wire) SetAmount(v int32)`

SetAmount sets Amount field to given value.


### GetBankMessage

`func (o *Wire) GetBankMessage() string`

GetBankMessage returns the BankMessage field if non-nil, zero value otherwise.

### GetBankMessageOk

`func (o *Wire) GetBankMessageOk() (*string, bool)`

GetBankMessageOk returns a tuple with the BankMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBankMessage

`func (o *Wire) SetBankMessage(v string)`

SetBankMessage sets BankMessage field to given value.

### HasBankMessage

`func (o *Wire) HasBankMessage() bool`

HasBankMessage returns a boolean if a field has been set.

### GetCreationTime

`func (o *Wire) GetCreationTime() time.Time`

GetCreationTime returns the CreationTime field if non-nil, zero value otherwise.

### GetCreationTimeOk

`func (o *Wire) GetCreationTimeOk() (*time.Time, bool)`

GetCreationTimeOk returns a tuple with the CreationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationTime

`func (o *Wire) SetCreationTime(v time.Time)`

SetCreationTime sets CreationTime field to given value.


### GetCurrency

`func (o *Wire) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *Wire) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *Wire) SetCurrency(v string)`

SetCurrency sets Currency field to given value.


### GetCustomerId

`func (o *Wire) GetCustomerId() string`

GetCustomerId returns the CustomerId field if non-nil, zero value otherwise.

### GetCustomerIdOk

`func (o *Wire) GetCustomerIdOk() (*string, bool)`

GetCustomerIdOk returns a tuple with the CustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerId

`func (o *Wire) SetCustomerId(v string)`

SetCustomerId sets CustomerId field to given value.


### GetId

`func (o *Wire) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Wire) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Wire) SetId(v string)`

SetId sets Id field to given value.


### GetLastUpdatedTime

`func (o *Wire) GetLastUpdatedTime() time.Time`

GetLastUpdatedTime returns the LastUpdatedTime field if non-nil, zero value otherwise.

### GetLastUpdatedTimeOk

`func (o *Wire) GetLastUpdatedTimeOk() (*time.Time, bool)`

GetLastUpdatedTimeOk returns a tuple with the LastUpdatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdatedTime

`func (o *Wire) SetLastUpdatedTime(v time.Time)`

SetLastUpdatedTime sets LastUpdatedTime field to given value.


### GetOriginatingAccountId

`func (o *Wire) GetOriginatingAccountId() string`

GetOriginatingAccountId returns the OriginatingAccountId field if non-nil, zero value otherwise.

### GetOriginatingAccountIdOk

`func (o *Wire) GetOriginatingAccountIdOk() (*string, bool)`

GetOriginatingAccountIdOk returns a tuple with the OriginatingAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginatingAccountId

`func (o *Wire) SetOriginatingAccountId(v string)`

SetOriginatingAccountId sets OriginatingAccountId field to given value.


### GetReceivingAccountId

`func (o *Wire) GetReceivingAccountId() string`

GetReceivingAccountId returns the ReceivingAccountId field if non-nil, zero value otherwise.

### GetReceivingAccountIdOk

`func (o *Wire) GetReceivingAccountIdOk() (*string, bool)`

GetReceivingAccountIdOk returns a tuple with the ReceivingAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceivingAccountId

`func (o *Wire) SetReceivingAccountId(v string)`

SetReceivingAccountId sets ReceivingAccountId field to given value.


### GetRecipientMessage

`func (o *Wire) GetRecipientMessage() string`

GetRecipientMessage returns the RecipientMessage field if non-nil, zero value otherwise.

### GetRecipientMessageOk

`func (o *Wire) GetRecipientMessageOk() (*string, bool)`

GetRecipientMessageOk returns a tuple with the RecipientMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipientMessage

`func (o *Wire) SetRecipientMessage(v string)`

SetRecipientMessage sets RecipientMessage field to given value.

### HasRecipientMessage

`func (o *Wire) HasRecipientMessage() bool`

HasRecipientMessage returns a boolean if a field has been set.

### GetSenderReferenceId

`func (o *Wire) GetSenderReferenceId() string`

GetSenderReferenceId returns the SenderReferenceId field if non-nil, zero value otherwise.

### GetSenderReferenceIdOk

`func (o *Wire) GetSenderReferenceIdOk() (*string, bool)`

GetSenderReferenceIdOk returns a tuple with the SenderReferenceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSenderReferenceId

`func (o *Wire) SetSenderReferenceId(v string)`

SetSenderReferenceId sets SenderReferenceId field to given value.


### GetStatus

`func (o *Wire) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Wire) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Wire) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetTransactionId

`func (o *Wire) GetTransactionId() string`

GetTransactionId returns the TransactionId field if non-nil, zero value otherwise.

### GetTransactionIdOk

`func (o *Wire) GetTransactionIdOk() (*string, bool)`

GetTransactionIdOk returns a tuple with the TransactionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionId

`func (o *Wire) SetTransactionId(v string)`

SetTransactionId sets TransactionId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


