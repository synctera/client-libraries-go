# PendingTransactionHistoryData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | **int64** | The amount of the hold. | 
**AutoPostAt** | **time.Time** | The time the transaction will be automatically posted. | 
**AvailBalance** | **int64** | The account \&quot;available balance\&quot; at the time this hold was created (to be deprecated) | 
**AvailableBalance** | **int64** | The account \&quot;available balance\&quot; at the time this hold was created | 
**Balance** | **int64** | The account balance at the time this hold was created | 
**Currency** | **string** | ISO 4217 alphabetic currency code of the transfer amount | 
**DcSign** | [**DcSign**](DcSign.md) |  | 
**EffectiveDate** | **time.Time** | The effective date of the transaction once it gets posted | 
**ExpiresAt** | **time.Time** | The date that at which this hold is no longer valid. | 
**ExternalData** | **map[string]interface{}** | an unstructured json blob representing additional transaction information supplied by the integrator. | 
**ForcePost** | **bool** | Whether or not the hold was forced (spending controls ignored) | 
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
**TransactionTime** | **time.Time** | The time the transaction occurred. | 
**Type** | **string** | The general type of transaction. For example, \&quot;card\&quot; or \&quot;ach\&quot;. | 
**UserData** | **map[string]interface{}** | An unstructured JSON blob representing additional transaction information specific to each payment rail. | 
**WasPartial** | **bool** | Does this hold represent a partial debit (or credit)? | 

## Methods

### NewPendingTransactionHistoryData

`func NewPendingTransactionHistoryData(amount int64, autoPostAt time.Time, availBalance int64, availableBalance int64, balance int64, currency string, dcSign DcSign, effectiveDate time.Time, expiresAt time.Time, externalData map[string]interface{}, forcePost bool, idemkey string, memo string, network string, operation string, reason string, reqAmount int64, riskInfo map[string]interface{}, status string, subtype string, totalAmount int64, transactionTime time.Time, type_ string, userData map[string]interface{}, wasPartial bool, ) *PendingTransactionHistoryData`

NewPendingTransactionHistoryData instantiates a new PendingTransactionHistoryData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPendingTransactionHistoryDataWithDefaults

`func NewPendingTransactionHistoryDataWithDefaults() *PendingTransactionHistoryData`

NewPendingTransactionHistoryDataWithDefaults instantiates a new PendingTransactionHistoryData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmount

`func (o *PendingTransactionHistoryData) GetAmount() int64`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *PendingTransactionHistoryData) GetAmountOk() (*int64, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *PendingTransactionHistoryData) SetAmount(v int64)`

SetAmount sets Amount field to given value.


### GetAutoPostAt

`func (o *PendingTransactionHistoryData) GetAutoPostAt() time.Time`

GetAutoPostAt returns the AutoPostAt field if non-nil, zero value otherwise.

### GetAutoPostAtOk

`func (o *PendingTransactionHistoryData) GetAutoPostAtOk() (*time.Time, bool)`

GetAutoPostAtOk returns a tuple with the AutoPostAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoPostAt

`func (o *PendingTransactionHistoryData) SetAutoPostAt(v time.Time)`

SetAutoPostAt sets AutoPostAt field to given value.


### GetAvailBalance

`func (o *PendingTransactionHistoryData) GetAvailBalance() int64`

GetAvailBalance returns the AvailBalance field if non-nil, zero value otherwise.

### GetAvailBalanceOk

`func (o *PendingTransactionHistoryData) GetAvailBalanceOk() (*int64, bool)`

GetAvailBalanceOk returns a tuple with the AvailBalance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailBalance

`func (o *PendingTransactionHistoryData) SetAvailBalance(v int64)`

SetAvailBalance sets AvailBalance field to given value.


### GetAvailableBalance

`func (o *PendingTransactionHistoryData) GetAvailableBalance() int64`

GetAvailableBalance returns the AvailableBalance field if non-nil, zero value otherwise.

### GetAvailableBalanceOk

`func (o *PendingTransactionHistoryData) GetAvailableBalanceOk() (*int64, bool)`

GetAvailableBalanceOk returns a tuple with the AvailableBalance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailableBalance

`func (o *PendingTransactionHistoryData) SetAvailableBalance(v int64)`

SetAvailableBalance sets AvailableBalance field to given value.


### GetBalance

`func (o *PendingTransactionHistoryData) GetBalance() int64`

GetBalance returns the Balance field if non-nil, zero value otherwise.

### GetBalanceOk

`func (o *PendingTransactionHistoryData) GetBalanceOk() (*int64, bool)`

GetBalanceOk returns a tuple with the Balance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalance

`func (o *PendingTransactionHistoryData) SetBalance(v int64)`

SetBalance sets Balance field to given value.


### GetCurrency

`func (o *PendingTransactionHistoryData) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *PendingTransactionHistoryData) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *PendingTransactionHistoryData) SetCurrency(v string)`

SetCurrency sets Currency field to given value.


### GetDcSign

`func (o *PendingTransactionHistoryData) GetDcSign() DcSign`

GetDcSign returns the DcSign field if non-nil, zero value otherwise.

### GetDcSignOk

`func (o *PendingTransactionHistoryData) GetDcSignOk() (*DcSign, bool)`

GetDcSignOk returns a tuple with the DcSign field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDcSign

`func (o *PendingTransactionHistoryData) SetDcSign(v DcSign)`

SetDcSign sets DcSign field to given value.


### GetEffectiveDate

`func (o *PendingTransactionHistoryData) GetEffectiveDate() time.Time`

GetEffectiveDate returns the EffectiveDate field if non-nil, zero value otherwise.

### GetEffectiveDateOk

`func (o *PendingTransactionHistoryData) GetEffectiveDateOk() (*time.Time, bool)`

GetEffectiveDateOk returns a tuple with the EffectiveDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEffectiveDate

`func (o *PendingTransactionHistoryData) SetEffectiveDate(v time.Time)`

SetEffectiveDate sets EffectiveDate field to given value.


### GetExpiresAt

`func (o *PendingTransactionHistoryData) GetExpiresAt() time.Time`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *PendingTransactionHistoryData) GetExpiresAtOk() (*time.Time, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *PendingTransactionHistoryData) SetExpiresAt(v time.Time)`

SetExpiresAt sets ExpiresAt field to given value.


### GetExternalData

`func (o *PendingTransactionHistoryData) GetExternalData() map[string]interface{}`

GetExternalData returns the ExternalData field if non-nil, zero value otherwise.

### GetExternalDataOk

`func (o *PendingTransactionHistoryData) GetExternalDataOk() (*map[string]interface{}, bool)`

GetExternalDataOk returns a tuple with the ExternalData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalData

`func (o *PendingTransactionHistoryData) SetExternalData(v map[string]interface{})`

SetExternalData sets ExternalData field to given value.


### SetExternalDataNil

`func (o *PendingTransactionHistoryData) SetExternalDataNil(b bool)`

 SetExternalDataNil sets the value for ExternalData to be an explicit nil

### UnsetExternalData
`func (o *PendingTransactionHistoryData) UnsetExternalData()`

UnsetExternalData ensures that no value is present for ExternalData, not even an explicit nil
### GetForcePost

`func (o *PendingTransactionHistoryData) GetForcePost() bool`

GetForcePost returns the ForcePost field if non-nil, zero value otherwise.

### GetForcePostOk

`func (o *PendingTransactionHistoryData) GetForcePostOk() (*bool, bool)`

GetForcePostOk returns a tuple with the ForcePost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForcePost

`func (o *PendingTransactionHistoryData) SetForcePost(v bool)`

SetForcePost sets ForcePost field to given value.


### GetIdemkey

`func (o *PendingTransactionHistoryData) GetIdemkey() string`

GetIdemkey returns the Idemkey field if non-nil, zero value otherwise.

### GetIdemkeyOk

`func (o *PendingTransactionHistoryData) GetIdemkeyOk() (*string, bool)`

GetIdemkeyOk returns a tuple with the Idemkey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdemkey

`func (o *PendingTransactionHistoryData) SetIdemkey(v string)`

SetIdemkey sets Idemkey field to given value.


### GetMemo

`func (o *PendingTransactionHistoryData) GetMemo() string`

GetMemo returns the Memo field if non-nil, zero value otherwise.

### GetMemoOk

`func (o *PendingTransactionHistoryData) GetMemoOk() (*string, bool)`

GetMemoOk returns a tuple with the Memo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemo

`func (o *PendingTransactionHistoryData) SetMemo(v string)`

SetMemo sets Memo field to given value.


### GetNetwork

`func (o *PendingTransactionHistoryData) GetNetwork() string`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *PendingTransactionHistoryData) GetNetworkOk() (*string, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *PendingTransactionHistoryData) SetNetwork(v string)`

SetNetwork sets Network field to given value.


### GetOperation

`func (o *PendingTransactionHistoryData) GetOperation() string`

GetOperation returns the Operation field if non-nil, zero value otherwise.

### GetOperationOk

`func (o *PendingTransactionHistoryData) GetOperationOk() (*string, bool)`

GetOperationOk returns a tuple with the Operation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperation

`func (o *PendingTransactionHistoryData) SetOperation(v string)`

SetOperation sets Operation field to given value.


### GetReason

`func (o *PendingTransactionHistoryData) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *PendingTransactionHistoryData) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *PendingTransactionHistoryData) SetReason(v string)`

SetReason sets Reason field to given value.


### GetReqAmount

`func (o *PendingTransactionHistoryData) GetReqAmount() int64`

GetReqAmount returns the ReqAmount field if non-nil, zero value otherwise.

### GetReqAmountOk

`func (o *PendingTransactionHistoryData) GetReqAmountOk() (*int64, bool)`

GetReqAmountOk returns a tuple with the ReqAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReqAmount

`func (o *PendingTransactionHistoryData) SetReqAmount(v int64)`

SetReqAmount sets ReqAmount field to given value.


### GetRiskInfo

`func (o *PendingTransactionHistoryData) GetRiskInfo() map[string]interface{}`

GetRiskInfo returns the RiskInfo field if non-nil, zero value otherwise.

### GetRiskInfoOk

`func (o *PendingTransactionHistoryData) GetRiskInfoOk() (*map[string]interface{}, bool)`

GetRiskInfoOk returns a tuple with the RiskInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRiskInfo

`func (o *PendingTransactionHistoryData) SetRiskInfo(v map[string]interface{})`

SetRiskInfo sets RiskInfo field to given value.


### SetRiskInfoNil

`func (o *PendingTransactionHistoryData) SetRiskInfoNil(b bool)`

 SetRiskInfoNil sets the value for RiskInfo to be an explicit nil

### UnsetRiskInfo
`func (o *PendingTransactionHistoryData) UnsetRiskInfo()`

UnsetRiskInfo ensures that no value is present for RiskInfo, not even an explicit nil
### GetStatus

`func (o *PendingTransactionHistoryData) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *PendingTransactionHistoryData) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *PendingTransactionHistoryData) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetSubtype

`func (o *PendingTransactionHistoryData) GetSubtype() string`

GetSubtype returns the Subtype field if non-nil, zero value otherwise.

### GetSubtypeOk

`func (o *PendingTransactionHistoryData) GetSubtypeOk() (*string, bool)`

GetSubtypeOk returns a tuple with the Subtype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubtype

`func (o *PendingTransactionHistoryData) SetSubtype(v string)`

SetSubtype sets Subtype field to given value.


### GetTotalAmount

`func (o *PendingTransactionHistoryData) GetTotalAmount() int64`

GetTotalAmount returns the TotalAmount field if non-nil, zero value otherwise.

### GetTotalAmountOk

`func (o *PendingTransactionHistoryData) GetTotalAmountOk() (*int64, bool)`

GetTotalAmountOk returns a tuple with the TotalAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalAmount

`func (o *PendingTransactionHistoryData) SetTotalAmount(v int64)`

SetTotalAmount sets TotalAmount field to given value.


### GetTransactionId

`func (o *PendingTransactionHistoryData) GetTransactionId() string`

GetTransactionId returns the TransactionId field if non-nil, zero value otherwise.

### GetTransactionIdOk

`func (o *PendingTransactionHistoryData) GetTransactionIdOk() (*string, bool)`

GetTransactionIdOk returns a tuple with the TransactionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionId

`func (o *PendingTransactionHistoryData) SetTransactionId(v string)`

SetTransactionId sets TransactionId field to given value.

### HasTransactionId

`func (o *PendingTransactionHistoryData) HasTransactionId() bool`

HasTransactionId returns a boolean if a field has been set.

### GetTransactionTime

`func (o *PendingTransactionHistoryData) GetTransactionTime() time.Time`

GetTransactionTime returns the TransactionTime field if non-nil, zero value otherwise.

### GetTransactionTimeOk

`func (o *PendingTransactionHistoryData) GetTransactionTimeOk() (*time.Time, bool)`

GetTransactionTimeOk returns a tuple with the TransactionTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionTime

`func (o *PendingTransactionHistoryData) SetTransactionTime(v time.Time)`

SetTransactionTime sets TransactionTime field to given value.


### GetType

`func (o *PendingTransactionHistoryData) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PendingTransactionHistoryData) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PendingTransactionHistoryData) SetType(v string)`

SetType sets Type field to given value.


### GetUserData

`func (o *PendingTransactionHistoryData) GetUserData() map[string]interface{}`

GetUserData returns the UserData field if non-nil, zero value otherwise.

### GetUserDataOk

`func (o *PendingTransactionHistoryData) GetUserDataOk() (*map[string]interface{}, bool)`

GetUserDataOk returns a tuple with the UserData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserData

`func (o *PendingTransactionHistoryData) SetUserData(v map[string]interface{})`

SetUserData sets UserData field to given value.


### SetUserDataNil

`func (o *PendingTransactionHistoryData) SetUserDataNil(b bool)`

 SetUserDataNil sets the value for UserData to be an explicit nil

### UnsetUserData
`func (o *PendingTransactionHistoryData) UnsetUserData()`

UnsetUserData ensures that no value is present for UserData, not even an explicit nil
### GetWasPartial

`func (o *PendingTransactionHistoryData) GetWasPartial() bool`

GetWasPartial returns the WasPartial field if non-nil, zero value otherwise.

### GetWasPartialOk

`func (o *PendingTransactionHistoryData) GetWasPartialOk() (*bool, bool)`

GetWasPartialOk returns a tuple with the WasPartial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWasPartial

`func (o *PendingTransactionHistoryData) SetWasPartial(v bool)`

SetWasPartial sets WasPartial field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


