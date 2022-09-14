# TransactionDirectPostRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountAlias** | **string** |  | 
**AccountNo** | **string** | The account number associated with the transaction | 
**Amount** | **int64** | The amount of the hold. | 
**Currency** | **string** | ISO 4217 alphabetic currency code of the transfer amount | 
**DcSign** | [**DcSign**](DcSign.md) |  | 
**Description** | Pointer to **string** | The description of the transaction | [optional] 
**EffectiveDate** | **time.Time** | The effective date of the transaction once it gets posted | 
**EnhancedTransaction** | Pointer to **map[string]interface{}** | An unstructured JSON blob representing additional transaction information specific to each payment rail. | [optional] 
**ExternalData** | **map[string]interface{}** | an unstructured json blob representing additional transaction information supplied by the integrator. | 
**ForcePost** | **bool** | Whether or not the hold was forced (spending controls ignored) | 
**InfoOnly** | **bool** | Whether or not this transaction represents a purely informational operation or an actual money movement | 
**Memo** | **string** | A short note to the recipient | 
**Network** | **string** | The network this transaction is associated with | 
**OffsetDescription** | Pointer to **string** | The description of the offset transaction | [optional] 
**ReferenceId** | **NullableString** | An external ID provided by the payment network to represent this transaction. This will always be null for internal transfers. | 
**RiskInfo** | **map[string]interface{}** | Information received by the transaction risk/fraud service related to this transaction | 
**Subtype** | **string** | The specific transaction type. For example, for &#x60;ach&#x60;, this may be \&quot;outgoing_debit\&quot;. | 
**TransactionTime** | **time.Time** | The time the transaction occurred. | 
**Type** | **string** | The general type of transaction. For example, \&quot;card\&quot; or \&quot;ach\&quot;. | 
**UserData** | **map[string]interface{}** | An unstructured JSON blob representing additional transaction information specific to each payment rail. | 
**Uuid** | **string** | The unique identifier of the transaction. | 

## Methods

### NewTransactionDirectPostRequest

`func NewTransactionDirectPostRequest(accountAlias string, accountNo string, amount int64, currency string, dcSign DcSign, effectiveDate time.Time, externalData map[string]interface{}, forcePost bool, infoOnly bool, memo string, network string, referenceId NullableString, riskInfo map[string]interface{}, subtype string, transactionTime time.Time, type_ string, userData map[string]interface{}, uuid string, ) *TransactionDirectPostRequest`

NewTransactionDirectPostRequest instantiates a new TransactionDirectPostRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransactionDirectPostRequestWithDefaults

`func NewTransactionDirectPostRequestWithDefaults() *TransactionDirectPostRequest`

NewTransactionDirectPostRequestWithDefaults instantiates a new TransactionDirectPostRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountAlias

`func (o *TransactionDirectPostRequest) GetAccountAlias() string`

GetAccountAlias returns the AccountAlias field if non-nil, zero value otherwise.

### GetAccountAliasOk

`func (o *TransactionDirectPostRequest) GetAccountAliasOk() (*string, bool)`

GetAccountAliasOk returns a tuple with the AccountAlias field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountAlias

`func (o *TransactionDirectPostRequest) SetAccountAlias(v string)`

SetAccountAlias sets AccountAlias field to given value.


### GetAccountNo

`func (o *TransactionDirectPostRequest) GetAccountNo() string`

GetAccountNo returns the AccountNo field if non-nil, zero value otherwise.

### GetAccountNoOk

`func (o *TransactionDirectPostRequest) GetAccountNoOk() (*string, bool)`

GetAccountNoOk returns a tuple with the AccountNo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountNo

`func (o *TransactionDirectPostRequest) SetAccountNo(v string)`

SetAccountNo sets AccountNo field to given value.


### GetAmount

`func (o *TransactionDirectPostRequest) GetAmount() int64`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *TransactionDirectPostRequest) GetAmountOk() (*int64, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *TransactionDirectPostRequest) SetAmount(v int64)`

SetAmount sets Amount field to given value.


### GetCurrency

`func (o *TransactionDirectPostRequest) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *TransactionDirectPostRequest) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *TransactionDirectPostRequest) SetCurrency(v string)`

SetCurrency sets Currency field to given value.


### GetDcSign

`func (o *TransactionDirectPostRequest) GetDcSign() DcSign`

GetDcSign returns the DcSign field if non-nil, zero value otherwise.

### GetDcSignOk

`func (o *TransactionDirectPostRequest) GetDcSignOk() (*DcSign, bool)`

GetDcSignOk returns a tuple with the DcSign field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDcSign

`func (o *TransactionDirectPostRequest) SetDcSign(v DcSign)`

SetDcSign sets DcSign field to given value.


### GetDescription

`func (o *TransactionDirectPostRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *TransactionDirectPostRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *TransactionDirectPostRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *TransactionDirectPostRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEffectiveDate

`func (o *TransactionDirectPostRequest) GetEffectiveDate() time.Time`

GetEffectiveDate returns the EffectiveDate field if non-nil, zero value otherwise.

### GetEffectiveDateOk

`func (o *TransactionDirectPostRequest) GetEffectiveDateOk() (*time.Time, bool)`

GetEffectiveDateOk returns a tuple with the EffectiveDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEffectiveDate

`func (o *TransactionDirectPostRequest) SetEffectiveDate(v time.Time)`

SetEffectiveDate sets EffectiveDate field to given value.


### GetEnhancedTransaction

`func (o *TransactionDirectPostRequest) GetEnhancedTransaction() map[string]interface{}`

GetEnhancedTransaction returns the EnhancedTransaction field if non-nil, zero value otherwise.

### GetEnhancedTransactionOk

`func (o *TransactionDirectPostRequest) GetEnhancedTransactionOk() (*map[string]interface{}, bool)`

GetEnhancedTransactionOk returns a tuple with the EnhancedTransaction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnhancedTransaction

`func (o *TransactionDirectPostRequest) SetEnhancedTransaction(v map[string]interface{})`

SetEnhancedTransaction sets EnhancedTransaction field to given value.

### HasEnhancedTransaction

`func (o *TransactionDirectPostRequest) HasEnhancedTransaction() bool`

HasEnhancedTransaction returns a boolean if a field has been set.

### SetEnhancedTransactionNil

`func (o *TransactionDirectPostRequest) SetEnhancedTransactionNil(b bool)`

 SetEnhancedTransactionNil sets the value for EnhancedTransaction to be an explicit nil

### UnsetEnhancedTransaction
`func (o *TransactionDirectPostRequest) UnsetEnhancedTransaction()`

UnsetEnhancedTransaction ensures that no value is present for EnhancedTransaction, not even an explicit nil
### GetExternalData

`func (o *TransactionDirectPostRequest) GetExternalData() map[string]interface{}`

GetExternalData returns the ExternalData field if non-nil, zero value otherwise.

### GetExternalDataOk

`func (o *TransactionDirectPostRequest) GetExternalDataOk() (*map[string]interface{}, bool)`

GetExternalDataOk returns a tuple with the ExternalData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalData

`func (o *TransactionDirectPostRequest) SetExternalData(v map[string]interface{})`

SetExternalData sets ExternalData field to given value.


### SetExternalDataNil

`func (o *TransactionDirectPostRequest) SetExternalDataNil(b bool)`

 SetExternalDataNil sets the value for ExternalData to be an explicit nil

### UnsetExternalData
`func (o *TransactionDirectPostRequest) UnsetExternalData()`

UnsetExternalData ensures that no value is present for ExternalData, not even an explicit nil
### GetForcePost

`func (o *TransactionDirectPostRequest) GetForcePost() bool`

GetForcePost returns the ForcePost field if non-nil, zero value otherwise.

### GetForcePostOk

`func (o *TransactionDirectPostRequest) GetForcePostOk() (*bool, bool)`

GetForcePostOk returns a tuple with the ForcePost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForcePost

`func (o *TransactionDirectPostRequest) SetForcePost(v bool)`

SetForcePost sets ForcePost field to given value.


### GetInfoOnly

`func (o *TransactionDirectPostRequest) GetInfoOnly() bool`

GetInfoOnly returns the InfoOnly field if non-nil, zero value otherwise.

### GetInfoOnlyOk

`func (o *TransactionDirectPostRequest) GetInfoOnlyOk() (*bool, bool)`

GetInfoOnlyOk returns a tuple with the InfoOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfoOnly

`func (o *TransactionDirectPostRequest) SetInfoOnly(v bool)`

SetInfoOnly sets InfoOnly field to given value.


### GetMemo

`func (o *TransactionDirectPostRequest) GetMemo() string`

GetMemo returns the Memo field if non-nil, zero value otherwise.

### GetMemoOk

`func (o *TransactionDirectPostRequest) GetMemoOk() (*string, bool)`

GetMemoOk returns a tuple with the Memo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemo

`func (o *TransactionDirectPostRequest) SetMemo(v string)`

SetMemo sets Memo field to given value.


### GetNetwork

`func (o *TransactionDirectPostRequest) GetNetwork() string`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *TransactionDirectPostRequest) GetNetworkOk() (*string, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *TransactionDirectPostRequest) SetNetwork(v string)`

SetNetwork sets Network field to given value.


### GetOffsetDescription

`func (o *TransactionDirectPostRequest) GetOffsetDescription() string`

GetOffsetDescription returns the OffsetDescription field if non-nil, zero value otherwise.

### GetOffsetDescriptionOk

`func (o *TransactionDirectPostRequest) GetOffsetDescriptionOk() (*string, bool)`

GetOffsetDescriptionOk returns a tuple with the OffsetDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffsetDescription

`func (o *TransactionDirectPostRequest) SetOffsetDescription(v string)`

SetOffsetDescription sets OffsetDescription field to given value.

### HasOffsetDescription

`func (o *TransactionDirectPostRequest) HasOffsetDescription() bool`

HasOffsetDescription returns a boolean if a field has been set.

### GetReferenceId

`func (o *TransactionDirectPostRequest) GetReferenceId() string`

GetReferenceId returns the ReferenceId field if non-nil, zero value otherwise.

### GetReferenceIdOk

`func (o *TransactionDirectPostRequest) GetReferenceIdOk() (*string, bool)`

GetReferenceIdOk returns a tuple with the ReferenceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceId

`func (o *TransactionDirectPostRequest) SetReferenceId(v string)`

SetReferenceId sets ReferenceId field to given value.


### SetReferenceIdNil

`func (o *TransactionDirectPostRequest) SetReferenceIdNil(b bool)`

 SetReferenceIdNil sets the value for ReferenceId to be an explicit nil

### UnsetReferenceId
`func (o *TransactionDirectPostRequest) UnsetReferenceId()`

UnsetReferenceId ensures that no value is present for ReferenceId, not even an explicit nil
### GetRiskInfo

`func (o *TransactionDirectPostRequest) GetRiskInfo() map[string]interface{}`

GetRiskInfo returns the RiskInfo field if non-nil, zero value otherwise.

### GetRiskInfoOk

`func (o *TransactionDirectPostRequest) GetRiskInfoOk() (*map[string]interface{}, bool)`

GetRiskInfoOk returns a tuple with the RiskInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRiskInfo

`func (o *TransactionDirectPostRequest) SetRiskInfo(v map[string]interface{})`

SetRiskInfo sets RiskInfo field to given value.


### SetRiskInfoNil

`func (o *TransactionDirectPostRequest) SetRiskInfoNil(b bool)`

 SetRiskInfoNil sets the value for RiskInfo to be an explicit nil

### UnsetRiskInfo
`func (o *TransactionDirectPostRequest) UnsetRiskInfo()`

UnsetRiskInfo ensures that no value is present for RiskInfo, not even an explicit nil
### GetSubtype

`func (o *TransactionDirectPostRequest) GetSubtype() string`

GetSubtype returns the Subtype field if non-nil, zero value otherwise.

### GetSubtypeOk

`func (o *TransactionDirectPostRequest) GetSubtypeOk() (*string, bool)`

GetSubtypeOk returns a tuple with the Subtype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubtype

`func (o *TransactionDirectPostRequest) SetSubtype(v string)`

SetSubtype sets Subtype field to given value.


### GetTransactionTime

`func (o *TransactionDirectPostRequest) GetTransactionTime() time.Time`

GetTransactionTime returns the TransactionTime field if non-nil, zero value otherwise.

### GetTransactionTimeOk

`func (o *TransactionDirectPostRequest) GetTransactionTimeOk() (*time.Time, bool)`

GetTransactionTimeOk returns a tuple with the TransactionTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionTime

`func (o *TransactionDirectPostRequest) SetTransactionTime(v time.Time)`

SetTransactionTime sets TransactionTime field to given value.


### GetType

`func (o *TransactionDirectPostRequest) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *TransactionDirectPostRequest) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *TransactionDirectPostRequest) SetType(v string)`

SetType sets Type field to given value.


### GetUserData

`func (o *TransactionDirectPostRequest) GetUserData() map[string]interface{}`

GetUserData returns the UserData field if non-nil, zero value otherwise.

### GetUserDataOk

`func (o *TransactionDirectPostRequest) GetUserDataOk() (*map[string]interface{}, bool)`

GetUserDataOk returns a tuple with the UserData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserData

`func (o *TransactionDirectPostRequest) SetUserData(v map[string]interface{})`

SetUserData sets UserData field to given value.


### SetUserDataNil

`func (o *TransactionDirectPostRequest) SetUserDataNil(b bool)`

 SetUserDataNil sets the value for UserData to be an explicit nil

### UnsetUserData
`func (o *TransactionDirectPostRequest) UnsetUserData()`

UnsetUserData ensures that no value is present for UserData, not even an explicit nil
### GetUuid

`func (o *TransactionDirectPostRequest) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *TransactionDirectPostRequest) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *TransactionDirectPostRequest) SetUuid(v string)`

SetUuid sets Uuid field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


