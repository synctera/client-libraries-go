# WireRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | **int32** | Transfer amount in cents ($100 would be 10000) | 
**BankMessage** | Pointer to **string** | Instructions intended for the financial institutions that are processing the wire. | [optional] 
**Currency** | **string** | 3-character currency code | 
**CustomerId** | **string** | The customer UUID representing the person initiating the Wire transfer | 
**Metadata** | Pointer to **map[string]interface{}** | Additional transfer metadata structured as key-value pairs | [optional] 
**OriginatingAccountId** | **string** | Sender account ID | 
**ReceivingAccountId** | **string** | The external account uuid representing the recipient of the wire. | 
**RecipientMessage** | Pointer to **string** | Information from the originator to the beneficiary (recipient). | [optional] 

## Methods

### NewWireRequest

`func NewWireRequest(amount int32, currency string, customerId string, originatingAccountId string, receivingAccountId string, ) *WireRequest`

NewWireRequest instantiates a new WireRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWireRequestWithDefaults

`func NewWireRequestWithDefaults() *WireRequest`

NewWireRequestWithDefaults instantiates a new WireRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmount

`func (o *WireRequest) GetAmount() int32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *WireRequest) GetAmountOk() (*int32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *WireRequest) SetAmount(v int32)`

SetAmount sets Amount field to given value.


### GetBankMessage

`func (o *WireRequest) GetBankMessage() string`

GetBankMessage returns the BankMessage field if non-nil, zero value otherwise.

### GetBankMessageOk

`func (o *WireRequest) GetBankMessageOk() (*string, bool)`

GetBankMessageOk returns a tuple with the BankMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBankMessage

`func (o *WireRequest) SetBankMessage(v string)`

SetBankMessage sets BankMessage field to given value.

### HasBankMessage

`func (o *WireRequest) HasBankMessage() bool`

HasBankMessage returns a boolean if a field has been set.

### GetCurrency

`func (o *WireRequest) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *WireRequest) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *WireRequest) SetCurrency(v string)`

SetCurrency sets Currency field to given value.


### GetCustomerId

`func (o *WireRequest) GetCustomerId() string`

GetCustomerId returns the CustomerId field if non-nil, zero value otherwise.

### GetCustomerIdOk

`func (o *WireRequest) GetCustomerIdOk() (*string, bool)`

GetCustomerIdOk returns a tuple with the CustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerId

`func (o *WireRequest) SetCustomerId(v string)`

SetCustomerId sets CustomerId field to given value.


### GetMetadata

`func (o *WireRequest) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *WireRequest) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *WireRequest) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *WireRequest) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetOriginatingAccountId

`func (o *WireRequest) GetOriginatingAccountId() string`

GetOriginatingAccountId returns the OriginatingAccountId field if non-nil, zero value otherwise.

### GetOriginatingAccountIdOk

`func (o *WireRequest) GetOriginatingAccountIdOk() (*string, bool)`

GetOriginatingAccountIdOk returns a tuple with the OriginatingAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginatingAccountId

`func (o *WireRequest) SetOriginatingAccountId(v string)`

SetOriginatingAccountId sets OriginatingAccountId field to given value.


### GetReceivingAccountId

`func (o *WireRequest) GetReceivingAccountId() string`

GetReceivingAccountId returns the ReceivingAccountId field if non-nil, zero value otherwise.

### GetReceivingAccountIdOk

`func (o *WireRequest) GetReceivingAccountIdOk() (*string, bool)`

GetReceivingAccountIdOk returns a tuple with the ReceivingAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceivingAccountId

`func (o *WireRequest) SetReceivingAccountId(v string)`

SetReceivingAccountId sets ReceivingAccountId field to given value.


### GetRecipientMessage

`func (o *WireRequest) GetRecipientMessage() string`

GetRecipientMessage returns the RecipientMessage field if non-nil, zero value otherwise.

### GetRecipientMessageOk

`func (o *WireRequest) GetRecipientMessageOk() (*string, bool)`

GetRecipientMessageOk returns a tuple with the RecipientMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipientMessage

`func (o *WireRequest) SetRecipientMessage(v string)`

SetRecipientMessage sets RecipientMessage field to given value.

### HasRecipientMessage

`func (o *WireRequest) HasRecipientMessage() bool`

HasRecipientMessage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


