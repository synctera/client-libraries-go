# HoldCreateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountNo** | **string** | The account number associated with the hold | 
**AllowPartial** | **bool** |  | 
**Amount** | **int64** | The amount of the hold. | 
**AutoPostAt** | **time.Time** | The time the transaction will be automatically posted. | 
**Currency** | **string** | ISO 4217 alphabetic currency code of the transfer amount | 
**DcSign** | [**DcSign**](DcSign.md) |  | 
**DeclineReason** | **string** | The reason for the decline, if any | 
**Description** | Pointer to **string** | The description of the transaction | [optional] 
**EffectiveDate** | **time.Time** | The effective date of the transaction once it gets posted | 
**EnhancedTransaction** | Pointer to **map[string]interface{}** | An unstructured JSON blob representing additional transaction information specific to each payment rail. | [optional] 
**ExpiresAt** | **time.Time** | The date that at which this hold is no longer valid. | 
**ExternalData** | **map[string]interface{}** | an unstructured json blob representing additional transaction information supplied by the integrator. | 
**ForcePost** | **bool** | Whether or not the hold was forced (spending controls ignored) | 
**Memo** | **string** | A short note to the recipient | 
**Network** | **string** | The network this transaction is associated with | 
**OffsetDescription** | Pointer to **string** | The description of the offset transaction | [optional] 
**ReferenceId** | **NullableString** | An external ID provided by the payment network to represent this transaction. This will always be null for internal transfers. | 
**Status** | **string** | The status of the hold. | 
**Subtype** | **string** | The specific transaction type. For example, for &#x60;ach&#x60;, this may be \&quot;outgoing_debit\&quot;. | 
**TransactionTime** | **time.Time** | The time the transaction occurred. | 
**Type** | **string** | The general type of transaction. For example, \&quot;card\&quot; or \&quot;ach\&quot;. | 
**UserData** | **map[string]interface{}** | An unstructured JSON blob representing additional transaction information specific to each payment rail. | 

## Methods

### NewHoldCreateRequest

`func NewHoldCreateRequest(accountNo string, allowPartial bool, amount int64, autoPostAt time.Time, currency string, dcSign DcSign, declineReason string, effectiveDate time.Time, expiresAt time.Time, externalData map[string]interface{}, forcePost bool, memo string, network string, referenceId NullableString, status string, subtype string, transactionTime time.Time, type_ string, userData map[string]interface{}, ) *HoldCreateRequest`

NewHoldCreateRequest instantiates a new HoldCreateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHoldCreateRequestWithDefaults

`func NewHoldCreateRequestWithDefaults() *HoldCreateRequest`

NewHoldCreateRequestWithDefaults instantiates a new HoldCreateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountNo

`func (o *HoldCreateRequest) GetAccountNo() string`

GetAccountNo returns the AccountNo field if non-nil, zero value otherwise.

### GetAccountNoOk

`func (o *HoldCreateRequest) GetAccountNoOk() (*string, bool)`

GetAccountNoOk returns a tuple with the AccountNo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountNo

`func (o *HoldCreateRequest) SetAccountNo(v string)`

SetAccountNo sets AccountNo field to given value.


### GetAllowPartial

`func (o *HoldCreateRequest) GetAllowPartial() bool`

GetAllowPartial returns the AllowPartial field if non-nil, zero value otherwise.

### GetAllowPartialOk

`func (o *HoldCreateRequest) GetAllowPartialOk() (*bool, bool)`

GetAllowPartialOk returns a tuple with the AllowPartial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowPartial

`func (o *HoldCreateRequest) SetAllowPartial(v bool)`

SetAllowPartial sets AllowPartial field to given value.


### GetAmount

`func (o *HoldCreateRequest) GetAmount() int64`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *HoldCreateRequest) GetAmountOk() (*int64, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *HoldCreateRequest) SetAmount(v int64)`

SetAmount sets Amount field to given value.


### GetAutoPostAt

`func (o *HoldCreateRequest) GetAutoPostAt() time.Time`

GetAutoPostAt returns the AutoPostAt field if non-nil, zero value otherwise.

### GetAutoPostAtOk

`func (o *HoldCreateRequest) GetAutoPostAtOk() (*time.Time, bool)`

GetAutoPostAtOk returns a tuple with the AutoPostAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoPostAt

`func (o *HoldCreateRequest) SetAutoPostAt(v time.Time)`

SetAutoPostAt sets AutoPostAt field to given value.


### GetCurrency

`func (o *HoldCreateRequest) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *HoldCreateRequest) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *HoldCreateRequest) SetCurrency(v string)`

SetCurrency sets Currency field to given value.


### GetDcSign

`func (o *HoldCreateRequest) GetDcSign() DcSign`

GetDcSign returns the DcSign field if non-nil, zero value otherwise.

### GetDcSignOk

`func (o *HoldCreateRequest) GetDcSignOk() (*DcSign, bool)`

GetDcSignOk returns a tuple with the DcSign field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDcSign

`func (o *HoldCreateRequest) SetDcSign(v DcSign)`

SetDcSign sets DcSign field to given value.


### GetDeclineReason

`func (o *HoldCreateRequest) GetDeclineReason() string`

GetDeclineReason returns the DeclineReason field if non-nil, zero value otherwise.

### GetDeclineReasonOk

`func (o *HoldCreateRequest) GetDeclineReasonOk() (*string, bool)`

GetDeclineReasonOk returns a tuple with the DeclineReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeclineReason

`func (o *HoldCreateRequest) SetDeclineReason(v string)`

SetDeclineReason sets DeclineReason field to given value.


### GetDescription

`func (o *HoldCreateRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *HoldCreateRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *HoldCreateRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *HoldCreateRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEffectiveDate

`func (o *HoldCreateRequest) GetEffectiveDate() time.Time`

GetEffectiveDate returns the EffectiveDate field if non-nil, zero value otherwise.

### GetEffectiveDateOk

`func (o *HoldCreateRequest) GetEffectiveDateOk() (*time.Time, bool)`

GetEffectiveDateOk returns a tuple with the EffectiveDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEffectiveDate

`func (o *HoldCreateRequest) SetEffectiveDate(v time.Time)`

SetEffectiveDate sets EffectiveDate field to given value.


### GetEnhancedTransaction

`func (o *HoldCreateRequest) GetEnhancedTransaction() map[string]interface{}`

GetEnhancedTransaction returns the EnhancedTransaction field if non-nil, zero value otherwise.

### GetEnhancedTransactionOk

`func (o *HoldCreateRequest) GetEnhancedTransactionOk() (*map[string]interface{}, bool)`

GetEnhancedTransactionOk returns a tuple with the EnhancedTransaction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnhancedTransaction

`func (o *HoldCreateRequest) SetEnhancedTransaction(v map[string]interface{})`

SetEnhancedTransaction sets EnhancedTransaction field to given value.

### HasEnhancedTransaction

`func (o *HoldCreateRequest) HasEnhancedTransaction() bool`

HasEnhancedTransaction returns a boolean if a field has been set.

### SetEnhancedTransactionNil

`func (o *HoldCreateRequest) SetEnhancedTransactionNil(b bool)`

 SetEnhancedTransactionNil sets the value for EnhancedTransaction to be an explicit nil

### UnsetEnhancedTransaction
`func (o *HoldCreateRequest) UnsetEnhancedTransaction()`

UnsetEnhancedTransaction ensures that no value is present for EnhancedTransaction, not even an explicit nil
### GetExpiresAt

`func (o *HoldCreateRequest) GetExpiresAt() time.Time`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *HoldCreateRequest) GetExpiresAtOk() (*time.Time, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *HoldCreateRequest) SetExpiresAt(v time.Time)`

SetExpiresAt sets ExpiresAt field to given value.


### GetExternalData

`func (o *HoldCreateRequest) GetExternalData() map[string]interface{}`

GetExternalData returns the ExternalData field if non-nil, zero value otherwise.

### GetExternalDataOk

`func (o *HoldCreateRequest) GetExternalDataOk() (*map[string]interface{}, bool)`

GetExternalDataOk returns a tuple with the ExternalData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalData

`func (o *HoldCreateRequest) SetExternalData(v map[string]interface{})`

SetExternalData sets ExternalData field to given value.


### SetExternalDataNil

`func (o *HoldCreateRequest) SetExternalDataNil(b bool)`

 SetExternalDataNil sets the value for ExternalData to be an explicit nil

### UnsetExternalData
`func (o *HoldCreateRequest) UnsetExternalData()`

UnsetExternalData ensures that no value is present for ExternalData, not even an explicit nil
### GetForcePost

`func (o *HoldCreateRequest) GetForcePost() bool`

GetForcePost returns the ForcePost field if non-nil, zero value otherwise.

### GetForcePostOk

`func (o *HoldCreateRequest) GetForcePostOk() (*bool, bool)`

GetForcePostOk returns a tuple with the ForcePost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForcePost

`func (o *HoldCreateRequest) SetForcePost(v bool)`

SetForcePost sets ForcePost field to given value.


### GetMemo

`func (o *HoldCreateRequest) GetMemo() string`

GetMemo returns the Memo field if non-nil, zero value otherwise.

### GetMemoOk

`func (o *HoldCreateRequest) GetMemoOk() (*string, bool)`

GetMemoOk returns a tuple with the Memo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemo

`func (o *HoldCreateRequest) SetMemo(v string)`

SetMemo sets Memo field to given value.


### GetNetwork

`func (o *HoldCreateRequest) GetNetwork() string`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *HoldCreateRequest) GetNetworkOk() (*string, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *HoldCreateRequest) SetNetwork(v string)`

SetNetwork sets Network field to given value.


### GetOffsetDescription

`func (o *HoldCreateRequest) GetOffsetDescription() string`

GetOffsetDescription returns the OffsetDescription field if non-nil, zero value otherwise.

### GetOffsetDescriptionOk

`func (o *HoldCreateRequest) GetOffsetDescriptionOk() (*string, bool)`

GetOffsetDescriptionOk returns a tuple with the OffsetDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffsetDescription

`func (o *HoldCreateRequest) SetOffsetDescription(v string)`

SetOffsetDescription sets OffsetDescription field to given value.

### HasOffsetDescription

`func (o *HoldCreateRequest) HasOffsetDescription() bool`

HasOffsetDescription returns a boolean if a field has been set.

### GetReferenceId

`func (o *HoldCreateRequest) GetReferenceId() string`

GetReferenceId returns the ReferenceId field if non-nil, zero value otherwise.

### GetReferenceIdOk

`func (o *HoldCreateRequest) GetReferenceIdOk() (*string, bool)`

GetReferenceIdOk returns a tuple with the ReferenceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceId

`func (o *HoldCreateRequest) SetReferenceId(v string)`

SetReferenceId sets ReferenceId field to given value.


### SetReferenceIdNil

`func (o *HoldCreateRequest) SetReferenceIdNil(b bool)`

 SetReferenceIdNil sets the value for ReferenceId to be an explicit nil

### UnsetReferenceId
`func (o *HoldCreateRequest) UnsetReferenceId()`

UnsetReferenceId ensures that no value is present for ReferenceId, not even an explicit nil
### GetStatus

`func (o *HoldCreateRequest) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *HoldCreateRequest) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *HoldCreateRequest) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetSubtype

`func (o *HoldCreateRequest) GetSubtype() string`

GetSubtype returns the Subtype field if non-nil, zero value otherwise.

### GetSubtypeOk

`func (o *HoldCreateRequest) GetSubtypeOk() (*string, bool)`

GetSubtypeOk returns a tuple with the Subtype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubtype

`func (o *HoldCreateRequest) SetSubtype(v string)`

SetSubtype sets Subtype field to given value.


### GetTransactionTime

`func (o *HoldCreateRequest) GetTransactionTime() time.Time`

GetTransactionTime returns the TransactionTime field if non-nil, zero value otherwise.

### GetTransactionTimeOk

`func (o *HoldCreateRequest) GetTransactionTimeOk() (*time.Time, bool)`

GetTransactionTimeOk returns a tuple with the TransactionTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionTime

`func (o *HoldCreateRequest) SetTransactionTime(v time.Time)`

SetTransactionTime sets TransactionTime field to given value.


### GetType

`func (o *HoldCreateRequest) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *HoldCreateRequest) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *HoldCreateRequest) SetType(v string)`

SetType sets Type field to given value.


### GetUserData

`func (o *HoldCreateRequest) GetUserData() map[string]interface{}`

GetUserData returns the UserData field if non-nil, zero value otherwise.

### GetUserDataOk

`func (o *HoldCreateRequest) GetUserDataOk() (*map[string]interface{}, bool)`

GetUserDataOk returns a tuple with the UserData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserData

`func (o *HoldCreateRequest) SetUserData(v map[string]interface{})`

SetUserData sets UserData field to given value.


### SetUserDataNil

`func (o *HoldCreateRequest) SetUserDataNil(b bool)`

 SetUserDataNil sets the value for UserData to be an explicit nil

### UnsetUserData
`func (o *HoldCreateRequest) UnsetUserData()`

UnsetUserData ensures that no value is present for UserData, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


