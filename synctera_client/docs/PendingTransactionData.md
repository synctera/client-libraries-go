# PendingTransactionData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | **int64** | The amount of the hold. | 
**AutoPostAt** | **time.Time** | The account \&quot;available balance\&quot; at the time this hold was created | 
**AvailBalance** | **int64** | The account \&quot;available balance\&quot; at the time this hold was created (to be deprecated) | 
**AvailableBalance** | **int64** | The account \&quot;available balance\&quot; at the time this hold was created | 
**Balance** | **int64** | The account balance at the time this hold was created | 
**Currency** | **string** | ISO 4217 alphabetic currency code of the transfer amount | 
**DcSign** | [**DcSign**](DcSign.md) |  | 
**EffectiveDate** | **time.Time** | The effective date of the transaction once it gets posted | 
**ExpiresAt** | **time.Time** | The date that at which this hold is no longer valid. | 
**ExternalData** | **map[string]interface{}** | an unstructured json blob representing additional transaction information supplied by the integrator. | 
**ForcePost** | **bool** | Whether or not the hold was forced (spending controls ignored) | 
**History** | [**[]PendingTransactionHistory**](PendingTransactionHistory.md) | An array representing any previous states of the hold, if it has been modified (For example, increasing or decreasing the hold amount). | 
**Idemkey** | **string** | The idempotency key used when initially creating this hold. | 
**Memo** | **string** | A short note to the recipient | 
**Network** | **string** | The network this transaction is associated with | 
**Operation** | **string** |  | 
**Reason** | **string** | If a hold has been declined or modified, this will include the reason. | 
**ReqAmount** | **int64** | The requested amount, in the case of hold modifications. | 
**RiskInfo** | **map[string]interface{}** | Information received by the transaction risk/fraud service related to this transaction | 
**Status** | **string** | The status of the hold. | 
**Subtype** | **string** | The specific transaction type. For example, for &#x60;ach&#x60;, this may be \&quot;outgoing_debit\&quot;. | 
**TotalAmount** | **int64** | The total amount of the hold. This may be different than &#x60;amount&#x60; in the case where a hold increase or decrease was requested. | 
**TransactionId** | Pointer to **string** | The uuid of the transaction that this pending transaction originated from, if any. This is primary used when a transaction \&quot;posts\&quot;, but a subset of the amount reserved until a future settlement date. | [optional] 
**TransactionTime** | **time.Time** | The time that the transaction was created | 
**Type** | **string** | The general type of transaction. For example, \&quot;card\&quot; or \&quot;ach\&quot;. | 
**UserData** | **map[string]interface{}** | An unstructured JSON blob representing additional transaction information specific to each payment rail. | 
**WasPartial** | **bool** | Does this hold represent a partial debit (or credit)? | 

## Methods

### NewPendingTransactionData

`func NewPendingTransactionData(amount int64, autoPostAt time.Time, availBalance int64, availableBalance int64, balance int64, currency string, dcSign DcSign, effectiveDate time.Time, expiresAt time.Time, externalData map[string]interface{}, forcePost bool, history []PendingTransactionHistory, idemkey string, memo string, network string, operation string, reason string, reqAmount int64, riskInfo map[string]interface{}, status string, subtype string, totalAmount int64, transactionTime time.Time, type_ string, userData map[string]interface{}, wasPartial bool, ) *PendingTransactionData`

NewPendingTransactionData instantiates a new PendingTransactionData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPendingTransactionDataWithDefaults

`func NewPendingTransactionDataWithDefaults() *PendingTransactionData`

NewPendingTransactionDataWithDefaults instantiates a new PendingTransactionData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmount

`func (o *PendingTransactionData) GetAmount() int64`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *PendingTransactionData) GetAmountOk() (*int64, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *PendingTransactionData) SetAmount(v int64)`

SetAmount sets Amount field to given value.


### GetAutoPostAt

`func (o *PendingTransactionData) GetAutoPostAt() time.Time`

GetAutoPostAt returns the AutoPostAt field if non-nil, zero value otherwise.

### GetAutoPostAtOk

`func (o *PendingTransactionData) GetAutoPostAtOk() (*time.Time, bool)`

GetAutoPostAtOk returns a tuple with the AutoPostAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoPostAt

`func (o *PendingTransactionData) SetAutoPostAt(v time.Time)`

SetAutoPostAt sets AutoPostAt field to given value.


### GetAvailBalance

`func (o *PendingTransactionData) GetAvailBalance() int64`

GetAvailBalance returns the AvailBalance field if non-nil, zero value otherwise.

### GetAvailBalanceOk

`func (o *PendingTransactionData) GetAvailBalanceOk() (*int64, bool)`

GetAvailBalanceOk returns a tuple with the AvailBalance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailBalance

`func (o *PendingTransactionData) SetAvailBalance(v int64)`

SetAvailBalance sets AvailBalance field to given value.


### GetAvailableBalance

`func (o *PendingTransactionData) GetAvailableBalance() int64`

GetAvailableBalance returns the AvailableBalance field if non-nil, zero value otherwise.

### GetAvailableBalanceOk

`func (o *PendingTransactionData) GetAvailableBalanceOk() (*int64, bool)`

GetAvailableBalanceOk returns a tuple with the AvailableBalance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailableBalance

`func (o *PendingTransactionData) SetAvailableBalance(v int64)`

SetAvailableBalance sets AvailableBalance field to given value.


### GetBalance

`func (o *PendingTransactionData) GetBalance() int64`

GetBalance returns the Balance field if non-nil, zero value otherwise.

### GetBalanceOk

`func (o *PendingTransactionData) GetBalanceOk() (*int64, bool)`

GetBalanceOk returns a tuple with the Balance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalance

`func (o *PendingTransactionData) SetBalance(v int64)`

SetBalance sets Balance field to given value.


### GetCurrency

`func (o *PendingTransactionData) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *PendingTransactionData) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *PendingTransactionData) SetCurrency(v string)`

SetCurrency sets Currency field to given value.


### GetDcSign

`func (o *PendingTransactionData) GetDcSign() DcSign`

GetDcSign returns the DcSign field if non-nil, zero value otherwise.

### GetDcSignOk

`func (o *PendingTransactionData) GetDcSignOk() (*DcSign, bool)`

GetDcSignOk returns a tuple with the DcSign field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDcSign

`func (o *PendingTransactionData) SetDcSign(v DcSign)`

SetDcSign sets DcSign field to given value.


### GetEffectiveDate

`func (o *PendingTransactionData) GetEffectiveDate() time.Time`

GetEffectiveDate returns the EffectiveDate field if non-nil, zero value otherwise.

### GetEffectiveDateOk

`func (o *PendingTransactionData) GetEffectiveDateOk() (*time.Time, bool)`

GetEffectiveDateOk returns a tuple with the EffectiveDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEffectiveDate

`func (o *PendingTransactionData) SetEffectiveDate(v time.Time)`

SetEffectiveDate sets EffectiveDate field to given value.


### GetExpiresAt

`func (o *PendingTransactionData) GetExpiresAt() time.Time`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *PendingTransactionData) GetExpiresAtOk() (*time.Time, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *PendingTransactionData) SetExpiresAt(v time.Time)`

SetExpiresAt sets ExpiresAt field to given value.


### GetExternalData

`func (o *PendingTransactionData) GetExternalData() map[string]interface{}`

GetExternalData returns the ExternalData field if non-nil, zero value otherwise.

### GetExternalDataOk

`func (o *PendingTransactionData) GetExternalDataOk() (*map[string]interface{}, bool)`

GetExternalDataOk returns a tuple with the ExternalData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalData

`func (o *PendingTransactionData) SetExternalData(v map[string]interface{})`

SetExternalData sets ExternalData field to given value.


### SetExternalDataNil

`func (o *PendingTransactionData) SetExternalDataNil(b bool)`

 SetExternalDataNil sets the value for ExternalData to be an explicit nil

### UnsetExternalData
`func (o *PendingTransactionData) UnsetExternalData()`

UnsetExternalData ensures that no value is present for ExternalData, not even an explicit nil
### GetForcePost

`func (o *PendingTransactionData) GetForcePost() bool`

GetForcePost returns the ForcePost field if non-nil, zero value otherwise.

### GetForcePostOk

`func (o *PendingTransactionData) GetForcePostOk() (*bool, bool)`

GetForcePostOk returns a tuple with the ForcePost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForcePost

`func (o *PendingTransactionData) SetForcePost(v bool)`

SetForcePost sets ForcePost field to given value.


### GetHistory

`func (o *PendingTransactionData) GetHistory() []PendingTransactionHistory`

GetHistory returns the History field if non-nil, zero value otherwise.

### GetHistoryOk

`func (o *PendingTransactionData) GetHistoryOk() (*[]PendingTransactionHistory, bool)`

GetHistoryOk returns a tuple with the History field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHistory

`func (o *PendingTransactionData) SetHistory(v []PendingTransactionHistory)`

SetHistory sets History field to given value.


### GetIdemkey

`func (o *PendingTransactionData) GetIdemkey() string`

GetIdemkey returns the Idemkey field if non-nil, zero value otherwise.

### GetIdemkeyOk

`func (o *PendingTransactionData) GetIdemkeyOk() (*string, bool)`

GetIdemkeyOk returns a tuple with the Idemkey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdemkey

`func (o *PendingTransactionData) SetIdemkey(v string)`

SetIdemkey sets Idemkey field to given value.


### GetMemo

`func (o *PendingTransactionData) GetMemo() string`

GetMemo returns the Memo field if non-nil, zero value otherwise.

### GetMemoOk

`func (o *PendingTransactionData) GetMemoOk() (*string, bool)`

GetMemoOk returns a tuple with the Memo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemo

`func (o *PendingTransactionData) SetMemo(v string)`

SetMemo sets Memo field to given value.


### GetNetwork

`func (o *PendingTransactionData) GetNetwork() string`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *PendingTransactionData) GetNetworkOk() (*string, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *PendingTransactionData) SetNetwork(v string)`

SetNetwork sets Network field to given value.


### GetOperation

`func (o *PendingTransactionData) GetOperation() string`

GetOperation returns the Operation field if non-nil, zero value otherwise.

### GetOperationOk

`func (o *PendingTransactionData) GetOperationOk() (*string, bool)`

GetOperationOk returns a tuple with the Operation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperation

`func (o *PendingTransactionData) SetOperation(v string)`

SetOperation sets Operation field to given value.


### GetReason

`func (o *PendingTransactionData) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *PendingTransactionData) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *PendingTransactionData) SetReason(v string)`

SetReason sets Reason field to given value.


### GetReqAmount

`func (o *PendingTransactionData) GetReqAmount() int64`

GetReqAmount returns the ReqAmount field if non-nil, zero value otherwise.

### GetReqAmountOk

`func (o *PendingTransactionData) GetReqAmountOk() (*int64, bool)`

GetReqAmountOk returns a tuple with the ReqAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReqAmount

`func (o *PendingTransactionData) SetReqAmount(v int64)`

SetReqAmount sets ReqAmount field to given value.


### GetRiskInfo

`func (o *PendingTransactionData) GetRiskInfo() map[string]interface{}`

GetRiskInfo returns the RiskInfo field if non-nil, zero value otherwise.

### GetRiskInfoOk

`func (o *PendingTransactionData) GetRiskInfoOk() (*map[string]interface{}, bool)`

GetRiskInfoOk returns a tuple with the RiskInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRiskInfo

`func (o *PendingTransactionData) SetRiskInfo(v map[string]interface{})`

SetRiskInfo sets RiskInfo field to given value.


### SetRiskInfoNil

`func (o *PendingTransactionData) SetRiskInfoNil(b bool)`

 SetRiskInfoNil sets the value for RiskInfo to be an explicit nil

### UnsetRiskInfo
`func (o *PendingTransactionData) UnsetRiskInfo()`

UnsetRiskInfo ensures that no value is present for RiskInfo, not even an explicit nil
### GetStatus

`func (o *PendingTransactionData) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *PendingTransactionData) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *PendingTransactionData) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetSubtype

`func (o *PendingTransactionData) GetSubtype() string`

GetSubtype returns the Subtype field if non-nil, zero value otherwise.

### GetSubtypeOk

`func (o *PendingTransactionData) GetSubtypeOk() (*string, bool)`

GetSubtypeOk returns a tuple with the Subtype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubtype

`func (o *PendingTransactionData) SetSubtype(v string)`

SetSubtype sets Subtype field to given value.


### GetTotalAmount

`func (o *PendingTransactionData) GetTotalAmount() int64`

GetTotalAmount returns the TotalAmount field if non-nil, zero value otherwise.

### GetTotalAmountOk

`func (o *PendingTransactionData) GetTotalAmountOk() (*int64, bool)`

GetTotalAmountOk returns a tuple with the TotalAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalAmount

`func (o *PendingTransactionData) SetTotalAmount(v int64)`

SetTotalAmount sets TotalAmount field to given value.


### GetTransactionId

`func (o *PendingTransactionData) GetTransactionId() string`

GetTransactionId returns the TransactionId field if non-nil, zero value otherwise.

### GetTransactionIdOk

`func (o *PendingTransactionData) GetTransactionIdOk() (*string, bool)`

GetTransactionIdOk returns a tuple with the TransactionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionId

`func (o *PendingTransactionData) SetTransactionId(v string)`

SetTransactionId sets TransactionId field to given value.

### HasTransactionId

`func (o *PendingTransactionData) HasTransactionId() bool`

HasTransactionId returns a boolean if a field has been set.

### GetTransactionTime

`func (o *PendingTransactionData) GetTransactionTime() time.Time`

GetTransactionTime returns the TransactionTime field if non-nil, zero value otherwise.

### GetTransactionTimeOk

`func (o *PendingTransactionData) GetTransactionTimeOk() (*time.Time, bool)`

GetTransactionTimeOk returns a tuple with the TransactionTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionTime

`func (o *PendingTransactionData) SetTransactionTime(v time.Time)`

SetTransactionTime sets TransactionTime field to given value.


### GetType

`func (o *PendingTransactionData) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PendingTransactionData) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PendingTransactionData) SetType(v string)`

SetType sets Type field to given value.


### GetUserData

`func (o *PendingTransactionData) GetUserData() map[string]interface{}`

GetUserData returns the UserData field if non-nil, zero value otherwise.

### GetUserDataOk

`func (o *PendingTransactionData) GetUserDataOk() (*map[string]interface{}, bool)`

GetUserDataOk returns a tuple with the UserData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserData

`func (o *PendingTransactionData) SetUserData(v map[string]interface{})`

SetUserData sets UserData field to given value.


### SetUserDataNil

`func (o *PendingTransactionData) SetUserDataNil(b bool)`

 SetUserDataNil sets the value for UserData to be an explicit nil

### UnsetUserData
`func (o *PendingTransactionData) UnsetUserData()`

UnsetUserData ensures that no value is present for UserData, not even an explicit nil
### GetWasPartial

`func (o *PendingTransactionData) GetWasPartial() bool`

GetWasPartial returns the WasPartial field if non-nil, zero value otherwise.

### GetWasPartialOk

`func (o *PendingTransactionData) GetWasPartialOk() (*bool, bool)`

GetWasPartialOk returns a tuple with the WasPartial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWasPartial

`func (o *PendingTransactionData) SetWasPartial(v bool)`

SetWasPartial sets WasPartial field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


