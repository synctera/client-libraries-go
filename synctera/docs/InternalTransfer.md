# InternalTransfer

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | **int32** | The amount (in cents) to transfer from originating account to receiving account. | 
**Currency** | **string** | ISO 4217 alphabetic currency code of the transfer amount | 
**Memo** | Pointer to **string** | A short note to the recipient | [optional] 
**OriginatingAccountAlias** | Pointer to **string** | An alias representing a GL account to debit | [optional] 
**OriginatingAccountId** | **string** | The UUID of the account being debited | 
**ReceivingAccountAlias** | Pointer to **string** | An alias representing a GL account to credit | [optional] 
**ReceivingAccountFullname** | Pointer to **string** | The full name of the owner of the receiving account. Only required when type is \&quot;outgoing_remittance\&quot; | [optional] 
**ReceivingAccountId** | **string** | The UUID of the account being credited | 
**Type** | **string** | The desired transaction type to use for this transfer | 

## Methods

### NewInternalTransfer

`func NewInternalTransfer(amount int32, currency string, originatingAccountId string, receivingAccountId string, type_ string, ) *InternalTransfer`

NewInternalTransfer instantiates a new InternalTransfer object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInternalTransferWithDefaults

`func NewInternalTransferWithDefaults() *InternalTransfer`

NewInternalTransferWithDefaults instantiates a new InternalTransfer object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmount

`func (o *InternalTransfer) GetAmount() int32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *InternalTransfer) GetAmountOk() (*int32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *InternalTransfer) SetAmount(v int32)`

SetAmount sets Amount field to given value.


### GetCurrency

`func (o *InternalTransfer) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *InternalTransfer) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *InternalTransfer) SetCurrency(v string)`

SetCurrency sets Currency field to given value.


### GetMemo

`func (o *InternalTransfer) GetMemo() string`

GetMemo returns the Memo field if non-nil, zero value otherwise.

### GetMemoOk

`func (o *InternalTransfer) GetMemoOk() (*string, bool)`

GetMemoOk returns a tuple with the Memo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemo

`func (o *InternalTransfer) SetMemo(v string)`

SetMemo sets Memo field to given value.

### HasMemo

`func (o *InternalTransfer) HasMemo() bool`

HasMemo returns a boolean if a field has been set.

### GetOriginatingAccountAlias

`func (o *InternalTransfer) GetOriginatingAccountAlias() string`

GetOriginatingAccountAlias returns the OriginatingAccountAlias field if non-nil, zero value otherwise.

### GetOriginatingAccountAliasOk

`func (o *InternalTransfer) GetOriginatingAccountAliasOk() (*string, bool)`

GetOriginatingAccountAliasOk returns a tuple with the OriginatingAccountAlias field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginatingAccountAlias

`func (o *InternalTransfer) SetOriginatingAccountAlias(v string)`

SetOriginatingAccountAlias sets OriginatingAccountAlias field to given value.

### HasOriginatingAccountAlias

`func (o *InternalTransfer) HasOriginatingAccountAlias() bool`

HasOriginatingAccountAlias returns a boolean if a field has been set.

### GetOriginatingAccountId

`func (o *InternalTransfer) GetOriginatingAccountId() string`

GetOriginatingAccountId returns the OriginatingAccountId field if non-nil, zero value otherwise.

### GetOriginatingAccountIdOk

`func (o *InternalTransfer) GetOriginatingAccountIdOk() (*string, bool)`

GetOriginatingAccountIdOk returns a tuple with the OriginatingAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginatingAccountId

`func (o *InternalTransfer) SetOriginatingAccountId(v string)`

SetOriginatingAccountId sets OriginatingAccountId field to given value.


### GetReceivingAccountAlias

`func (o *InternalTransfer) GetReceivingAccountAlias() string`

GetReceivingAccountAlias returns the ReceivingAccountAlias field if non-nil, zero value otherwise.

### GetReceivingAccountAliasOk

`func (o *InternalTransfer) GetReceivingAccountAliasOk() (*string, bool)`

GetReceivingAccountAliasOk returns a tuple with the ReceivingAccountAlias field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceivingAccountAlias

`func (o *InternalTransfer) SetReceivingAccountAlias(v string)`

SetReceivingAccountAlias sets ReceivingAccountAlias field to given value.

### HasReceivingAccountAlias

`func (o *InternalTransfer) HasReceivingAccountAlias() bool`

HasReceivingAccountAlias returns a boolean if a field has been set.

### GetReceivingAccountFullname

`func (o *InternalTransfer) GetReceivingAccountFullname() string`

GetReceivingAccountFullname returns the ReceivingAccountFullname field if non-nil, zero value otherwise.

### GetReceivingAccountFullnameOk

`func (o *InternalTransfer) GetReceivingAccountFullnameOk() (*string, bool)`

GetReceivingAccountFullnameOk returns a tuple with the ReceivingAccountFullname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceivingAccountFullname

`func (o *InternalTransfer) SetReceivingAccountFullname(v string)`

SetReceivingAccountFullname sets ReceivingAccountFullname field to given value.

### HasReceivingAccountFullname

`func (o *InternalTransfer) HasReceivingAccountFullname() bool`

HasReceivingAccountFullname returns a boolean if a field has been set.

### GetReceivingAccountId

`func (o *InternalTransfer) GetReceivingAccountId() string`

GetReceivingAccountId returns the ReceivingAccountId field if non-nil, zero value otherwise.

### GetReceivingAccountIdOk

`func (o *InternalTransfer) GetReceivingAccountIdOk() (*string, bool)`

GetReceivingAccountIdOk returns a tuple with the ReceivingAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceivingAccountId

`func (o *InternalTransfer) SetReceivingAccountId(v string)`

SetReceivingAccountId sets ReceivingAccountId field to given value.


### GetType

`func (o *InternalTransfer) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *InternalTransfer) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *InternalTransfer) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


