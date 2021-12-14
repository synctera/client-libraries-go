# InternalTransferResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | **int32** | The amount (in cents) to transfer from originating account to receiving account. | 
**Currency** | **string** | ISO 4217 alphabetic currency code of the transfer amount | 
**Memo** | Pointer to **string** | A short note to the recipient | [optional] 
**OriginatingAccountAlias** | Pointer to **string** | An alias representing a GL account to debit. This is alternative to specifying by account id | [optional] 
**OriginatingAccountId** | Pointer to **string** | The UUID of the account being debited | [optional] 
**ReceivingAccountAlias** | Pointer to **string** | An alias representing a GL account to credit. This is an alternative to specifying by account id | [optional] 
**ReceivingAccountCustomerId** | Pointer to **string** | The customer id of the owner of the receiving account. Only required when type is \&quot;outgoing_remittance\&quot; | [optional] 
**ReceivingAccountId** | Pointer to **string** | The UUID of the account being credited | [optional] 
**Type** | **string** | The desired transaction type to use for this transfer | 
**Id** | **string** | The transaction id associated with the transfer | 

## Methods

### NewInternalTransferResponse

`func NewInternalTransferResponse(amount int32, currency string, type_ string, id string, ) *InternalTransferResponse`

NewInternalTransferResponse instantiates a new InternalTransferResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInternalTransferResponseWithDefaults

`func NewInternalTransferResponseWithDefaults() *InternalTransferResponse`

NewInternalTransferResponseWithDefaults instantiates a new InternalTransferResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmount

`func (o *InternalTransferResponse) GetAmount() int32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *InternalTransferResponse) GetAmountOk() (*int32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *InternalTransferResponse) SetAmount(v int32)`

SetAmount sets Amount field to given value.


### GetCurrency

`func (o *InternalTransferResponse) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *InternalTransferResponse) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *InternalTransferResponse) SetCurrency(v string)`

SetCurrency sets Currency field to given value.


### GetMemo

`func (o *InternalTransferResponse) GetMemo() string`

GetMemo returns the Memo field if non-nil, zero value otherwise.

### GetMemoOk

`func (o *InternalTransferResponse) GetMemoOk() (*string, bool)`

GetMemoOk returns a tuple with the Memo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemo

`func (o *InternalTransferResponse) SetMemo(v string)`

SetMemo sets Memo field to given value.

### HasMemo

`func (o *InternalTransferResponse) HasMemo() bool`

HasMemo returns a boolean if a field has been set.

### GetOriginatingAccountAlias

`func (o *InternalTransferResponse) GetOriginatingAccountAlias() string`

GetOriginatingAccountAlias returns the OriginatingAccountAlias field if non-nil, zero value otherwise.

### GetOriginatingAccountAliasOk

`func (o *InternalTransferResponse) GetOriginatingAccountAliasOk() (*string, bool)`

GetOriginatingAccountAliasOk returns a tuple with the OriginatingAccountAlias field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginatingAccountAlias

`func (o *InternalTransferResponse) SetOriginatingAccountAlias(v string)`

SetOriginatingAccountAlias sets OriginatingAccountAlias field to given value.

### HasOriginatingAccountAlias

`func (o *InternalTransferResponse) HasOriginatingAccountAlias() bool`

HasOriginatingAccountAlias returns a boolean if a field has been set.

### GetOriginatingAccountId

`func (o *InternalTransferResponse) GetOriginatingAccountId() string`

GetOriginatingAccountId returns the OriginatingAccountId field if non-nil, zero value otherwise.

### GetOriginatingAccountIdOk

`func (o *InternalTransferResponse) GetOriginatingAccountIdOk() (*string, bool)`

GetOriginatingAccountIdOk returns a tuple with the OriginatingAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginatingAccountId

`func (o *InternalTransferResponse) SetOriginatingAccountId(v string)`

SetOriginatingAccountId sets OriginatingAccountId field to given value.

### HasOriginatingAccountId

`func (o *InternalTransferResponse) HasOriginatingAccountId() bool`

HasOriginatingAccountId returns a boolean if a field has been set.

### GetReceivingAccountAlias

`func (o *InternalTransferResponse) GetReceivingAccountAlias() string`

GetReceivingAccountAlias returns the ReceivingAccountAlias field if non-nil, zero value otherwise.

### GetReceivingAccountAliasOk

`func (o *InternalTransferResponse) GetReceivingAccountAliasOk() (*string, bool)`

GetReceivingAccountAliasOk returns a tuple with the ReceivingAccountAlias field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceivingAccountAlias

`func (o *InternalTransferResponse) SetReceivingAccountAlias(v string)`

SetReceivingAccountAlias sets ReceivingAccountAlias field to given value.

### HasReceivingAccountAlias

`func (o *InternalTransferResponse) HasReceivingAccountAlias() bool`

HasReceivingAccountAlias returns a boolean if a field has been set.

### GetReceivingAccountCustomerId

`func (o *InternalTransferResponse) GetReceivingAccountCustomerId() string`

GetReceivingAccountCustomerId returns the ReceivingAccountCustomerId field if non-nil, zero value otherwise.

### GetReceivingAccountCustomerIdOk

`func (o *InternalTransferResponse) GetReceivingAccountCustomerIdOk() (*string, bool)`

GetReceivingAccountCustomerIdOk returns a tuple with the ReceivingAccountCustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceivingAccountCustomerId

`func (o *InternalTransferResponse) SetReceivingAccountCustomerId(v string)`

SetReceivingAccountCustomerId sets ReceivingAccountCustomerId field to given value.

### HasReceivingAccountCustomerId

`func (o *InternalTransferResponse) HasReceivingAccountCustomerId() bool`

HasReceivingAccountCustomerId returns a boolean if a field has been set.

### GetReceivingAccountId

`func (o *InternalTransferResponse) GetReceivingAccountId() string`

GetReceivingAccountId returns the ReceivingAccountId field if non-nil, zero value otherwise.

### GetReceivingAccountIdOk

`func (o *InternalTransferResponse) GetReceivingAccountIdOk() (*string, bool)`

GetReceivingAccountIdOk returns a tuple with the ReceivingAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceivingAccountId

`func (o *InternalTransferResponse) SetReceivingAccountId(v string)`

SetReceivingAccountId sets ReceivingAccountId field to given value.

### HasReceivingAccountId

`func (o *InternalTransferResponse) HasReceivingAccountId() bool`

HasReceivingAccountId returns a boolean if a field has been set.

### GetType

`func (o *InternalTransferResponse) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *InternalTransferResponse) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *InternalTransferResponse) SetType(v string)`

SetType sets Type field to given value.


### GetId

`func (o *InternalTransferResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *InternalTransferResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *InternalTransferResponse) SetId(v string)`

SetId sets Id field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


