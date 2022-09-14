# AccountToAccountTransferRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | **int64** | The amount (in cents) to transfer from originating account to receiving account. | 
**Currency** | **string** | ISO 4217 alphabetic currency code of the transfer amount | 
**Description** | Pointer to **string** | A description of the transaction | [optional] 
**DestAccountAlias** | **string** | An alias representing a GL account to credit. This is an alternative to specifying by account id | 
**DestAccountNo** | **string** | The account number of the account being credited | 
**EffectiveDate** | **time.Time** | The date the transaction should be effective | 
**ExternalData** | **map[string]interface{}** | an unstructured json blob representing additional transaction information supplied by the integrator. | 
**Memo** | **string** | A short note to the recipient | 
**OffsetDescription** | Pointer to **string** | A description of the offsetting transaction | [optional] 
**PostedDate** | Pointer to **time.Time** | The date the transaction was posted | [optional] 
**SourceAccountAlias** | **string** | An alias representing a GL account to debit. This is alternative to specifying by account id | 
**SourceAccountNo** | **string** | The account number of the account being debited | 
**Subtype** | **string** | The desired transaction subtype to use for this transfer | 
**Type** | **string** | The desired transaction type to use for this transfer | 
**UserData** | **map[string]interface{}** | An unstructured JSON blob representing additional transaction information specific to each payment rail. | 
**Uuid** | **string** | The UUID of the transaction | 

## Methods

### NewAccountToAccountTransferRequest

`func NewAccountToAccountTransferRequest(amount int64, currency string, destAccountAlias string, destAccountNo string, effectiveDate time.Time, externalData map[string]interface{}, memo string, sourceAccountAlias string, sourceAccountNo string, subtype string, type_ string, userData map[string]interface{}, uuid string, ) *AccountToAccountTransferRequest`

NewAccountToAccountTransferRequest instantiates a new AccountToAccountTransferRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountToAccountTransferRequestWithDefaults

`func NewAccountToAccountTransferRequestWithDefaults() *AccountToAccountTransferRequest`

NewAccountToAccountTransferRequestWithDefaults instantiates a new AccountToAccountTransferRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmount

`func (o *AccountToAccountTransferRequest) GetAmount() int64`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *AccountToAccountTransferRequest) GetAmountOk() (*int64, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *AccountToAccountTransferRequest) SetAmount(v int64)`

SetAmount sets Amount field to given value.


### GetCurrency

`func (o *AccountToAccountTransferRequest) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *AccountToAccountTransferRequest) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *AccountToAccountTransferRequest) SetCurrency(v string)`

SetCurrency sets Currency field to given value.


### GetDescription

`func (o *AccountToAccountTransferRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AccountToAccountTransferRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AccountToAccountTransferRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AccountToAccountTransferRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDestAccountAlias

`func (o *AccountToAccountTransferRequest) GetDestAccountAlias() string`

GetDestAccountAlias returns the DestAccountAlias field if non-nil, zero value otherwise.

### GetDestAccountAliasOk

`func (o *AccountToAccountTransferRequest) GetDestAccountAliasOk() (*string, bool)`

GetDestAccountAliasOk returns a tuple with the DestAccountAlias field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestAccountAlias

`func (o *AccountToAccountTransferRequest) SetDestAccountAlias(v string)`

SetDestAccountAlias sets DestAccountAlias field to given value.


### GetDestAccountNo

`func (o *AccountToAccountTransferRequest) GetDestAccountNo() string`

GetDestAccountNo returns the DestAccountNo field if non-nil, zero value otherwise.

### GetDestAccountNoOk

`func (o *AccountToAccountTransferRequest) GetDestAccountNoOk() (*string, bool)`

GetDestAccountNoOk returns a tuple with the DestAccountNo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestAccountNo

`func (o *AccountToAccountTransferRequest) SetDestAccountNo(v string)`

SetDestAccountNo sets DestAccountNo field to given value.


### GetEffectiveDate

`func (o *AccountToAccountTransferRequest) GetEffectiveDate() time.Time`

GetEffectiveDate returns the EffectiveDate field if non-nil, zero value otherwise.

### GetEffectiveDateOk

`func (o *AccountToAccountTransferRequest) GetEffectiveDateOk() (*time.Time, bool)`

GetEffectiveDateOk returns a tuple with the EffectiveDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEffectiveDate

`func (o *AccountToAccountTransferRequest) SetEffectiveDate(v time.Time)`

SetEffectiveDate sets EffectiveDate field to given value.


### GetExternalData

`func (o *AccountToAccountTransferRequest) GetExternalData() map[string]interface{}`

GetExternalData returns the ExternalData field if non-nil, zero value otherwise.

### GetExternalDataOk

`func (o *AccountToAccountTransferRequest) GetExternalDataOk() (*map[string]interface{}, bool)`

GetExternalDataOk returns a tuple with the ExternalData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalData

`func (o *AccountToAccountTransferRequest) SetExternalData(v map[string]interface{})`

SetExternalData sets ExternalData field to given value.


### SetExternalDataNil

`func (o *AccountToAccountTransferRequest) SetExternalDataNil(b bool)`

 SetExternalDataNil sets the value for ExternalData to be an explicit nil

### UnsetExternalData
`func (o *AccountToAccountTransferRequest) UnsetExternalData()`

UnsetExternalData ensures that no value is present for ExternalData, not even an explicit nil
### GetMemo

`func (o *AccountToAccountTransferRequest) GetMemo() string`

GetMemo returns the Memo field if non-nil, zero value otherwise.

### GetMemoOk

`func (o *AccountToAccountTransferRequest) GetMemoOk() (*string, bool)`

GetMemoOk returns a tuple with the Memo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemo

`func (o *AccountToAccountTransferRequest) SetMemo(v string)`

SetMemo sets Memo field to given value.


### GetOffsetDescription

`func (o *AccountToAccountTransferRequest) GetOffsetDescription() string`

GetOffsetDescription returns the OffsetDescription field if non-nil, zero value otherwise.

### GetOffsetDescriptionOk

`func (o *AccountToAccountTransferRequest) GetOffsetDescriptionOk() (*string, bool)`

GetOffsetDescriptionOk returns a tuple with the OffsetDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffsetDescription

`func (o *AccountToAccountTransferRequest) SetOffsetDescription(v string)`

SetOffsetDescription sets OffsetDescription field to given value.

### HasOffsetDescription

`func (o *AccountToAccountTransferRequest) HasOffsetDescription() bool`

HasOffsetDescription returns a boolean if a field has been set.

### GetPostedDate

`func (o *AccountToAccountTransferRequest) GetPostedDate() time.Time`

GetPostedDate returns the PostedDate field if non-nil, zero value otherwise.

### GetPostedDateOk

`func (o *AccountToAccountTransferRequest) GetPostedDateOk() (*time.Time, bool)`

GetPostedDateOk returns a tuple with the PostedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostedDate

`func (o *AccountToAccountTransferRequest) SetPostedDate(v time.Time)`

SetPostedDate sets PostedDate field to given value.

### HasPostedDate

`func (o *AccountToAccountTransferRequest) HasPostedDate() bool`

HasPostedDate returns a boolean if a field has been set.

### GetSourceAccountAlias

`func (o *AccountToAccountTransferRequest) GetSourceAccountAlias() string`

GetSourceAccountAlias returns the SourceAccountAlias field if non-nil, zero value otherwise.

### GetSourceAccountAliasOk

`func (o *AccountToAccountTransferRequest) GetSourceAccountAliasOk() (*string, bool)`

GetSourceAccountAliasOk returns a tuple with the SourceAccountAlias field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceAccountAlias

`func (o *AccountToAccountTransferRequest) SetSourceAccountAlias(v string)`

SetSourceAccountAlias sets SourceAccountAlias field to given value.


### GetSourceAccountNo

`func (o *AccountToAccountTransferRequest) GetSourceAccountNo() string`

GetSourceAccountNo returns the SourceAccountNo field if non-nil, zero value otherwise.

### GetSourceAccountNoOk

`func (o *AccountToAccountTransferRequest) GetSourceAccountNoOk() (*string, bool)`

GetSourceAccountNoOk returns a tuple with the SourceAccountNo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceAccountNo

`func (o *AccountToAccountTransferRequest) SetSourceAccountNo(v string)`

SetSourceAccountNo sets SourceAccountNo field to given value.


### GetSubtype

`func (o *AccountToAccountTransferRequest) GetSubtype() string`

GetSubtype returns the Subtype field if non-nil, zero value otherwise.

### GetSubtypeOk

`func (o *AccountToAccountTransferRequest) GetSubtypeOk() (*string, bool)`

GetSubtypeOk returns a tuple with the Subtype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubtype

`func (o *AccountToAccountTransferRequest) SetSubtype(v string)`

SetSubtype sets Subtype field to given value.


### GetType

`func (o *AccountToAccountTransferRequest) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AccountToAccountTransferRequest) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AccountToAccountTransferRequest) SetType(v string)`

SetType sets Type field to given value.


### GetUserData

`func (o *AccountToAccountTransferRequest) GetUserData() map[string]interface{}`

GetUserData returns the UserData field if non-nil, zero value otherwise.

### GetUserDataOk

`func (o *AccountToAccountTransferRequest) GetUserDataOk() (*map[string]interface{}, bool)`

GetUserDataOk returns a tuple with the UserData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserData

`func (o *AccountToAccountTransferRequest) SetUserData(v map[string]interface{})`

SetUserData sets UserData field to given value.


### SetUserDataNil

`func (o *AccountToAccountTransferRequest) SetUserDataNil(b bool)`

 SetUserDataNil sets the value for UserData to be an explicit nil

### UnsetUserData
`func (o *AccountToAccountTransferRequest) UnsetUserData()`

UnsetUserData ensures that no value is present for UserData, not even an explicit nil
### GetUuid

`func (o *AccountToAccountTransferRequest) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *AccountToAccountTransferRequest) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *AccountToAccountTransferRequest) SetUuid(v string)`

SetUuid sets Uuid field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


