# PendingTransactionData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | **int32** | The amount of the hold. | 
**AvailBalance** | **int32** | The account \&quot;available balance\&quot; at the time this hold was created | 
**Balance** | **int32** | The account balance at the time this hold was created | 
**Currency** | **string** | ISO 4217 alphabetic currency code of the transfer amount | 
**DcSign** | [**DcSign**](DcSign.md) |  | 
**EffectiveDate** | **time.Time** | The effective date of the transaction once it gets posted | 
**ExpiresAt** | **time.Time** | The date that at which this hold is no longer valid. | 
**ForcePost** | **bool** | Whether or not the hold was forced (spending controls ignored) | 
**History** | [**[]PendingTransactionHistory**](PendingTransactionHistory.md) | An array representing any previous states of the hold, if it has been modified (For example, increasing or decreasing the hold amount). | 
**Network** | **string** | The network this transaction is associated with | 
**Operation** | **string** |  | 
**Reason** | **string** | If a hold has been declined or modified, this will include the reason. | 
**ReqAmount** | **int32** | The requested amount, in the case of hold modifications. | 
**RiskInfo** | **map[string]interface{}** | Information received by the transaction risk/fraud service related to this transaction | 
**Status** | **string** | The status of the hold. | 
**Subtype** | **string** | The specific transaction type. For example, for &#x60;ach&#x60;, this may be \&quot;outgoing_debit\&quot;. | 
**TotalAmount** | **int32** | The total amount of the hold. This may be different than &#x60;amount&#x60; in the case where a hold increase or decrease was requested. | 
**Type** | **string** | The general type of transaction. For example, \&quot;card\&quot; or \&quot;ach\&quot;. | 
**UserData** | **map[string]interface{}** | An unstructured JSON blob representing additional transaction information specific to each payment rail. | 
**WasPartial** | **bool** | Does this hold represent a partial debit (or credit)? | 

## Methods

### NewPendingTransactionData

`func NewPendingTransactionData(amount int32, availBalance int32, balance int32, currency string, dcSign DcSign, effectiveDate time.Time, expiresAt time.Time, forcePost bool, history []PendingTransactionHistory, network string, operation string, reason string, reqAmount int32, riskInfo map[string]interface{}, status string, subtype string, totalAmount int32, type_ string, userData map[string]interface{}, wasPartial bool, ) *PendingTransactionData`

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

`func (o *PendingTransactionData) GetAmount() int32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *PendingTransactionData) GetAmountOk() (*int32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *PendingTransactionData) SetAmount(v int32)`

SetAmount sets Amount field to given value.


### GetAvailBalance

`func (o *PendingTransactionData) GetAvailBalance() int32`

GetAvailBalance returns the AvailBalance field if non-nil, zero value otherwise.

### GetAvailBalanceOk

`func (o *PendingTransactionData) GetAvailBalanceOk() (*int32, bool)`

GetAvailBalanceOk returns a tuple with the AvailBalance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailBalance

`func (o *PendingTransactionData) SetAvailBalance(v int32)`

SetAvailBalance sets AvailBalance field to given value.


### GetBalance

`func (o *PendingTransactionData) GetBalance() int32`

GetBalance returns the Balance field if non-nil, zero value otherwise.

### GetBalanceOk

`func (o *PendingTransactionData) GetBalanceOk() (*int32, bool)`

GetBalanceOk returns a tuple with the Balance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalance

`func (o *PendingTransactionData) SetBalance(v int32)`

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

`func (o *PendingTransactionData) GetReqAmount() int32`

GetReqAmount returns the ReqAmount field if non-nil, zero value otherwise.

### GetReqAmountOk

`func (o *PendingTransactionData) GetReqAmountOk() (*int32, bool)`

GetReqAmountOk returns a tuple with the ReqAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReqAmount

`func (o *PendingTransactionData) SetReqAmount(v int32)`

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

`func (o *PendingTransactionData) GetTotalAmount() int32`

GetTotalAmount returns the TotalAmount field if non-nil, zero value otherwise.

### GetTotalAmountOk

`func (o *PendingTransactionData) GetTotalAmountOk() (*int32, bool)`

GetTotalAmountOk returns a tuple with the TotalAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalAmount

`func (o *PendingTransactionData) SetTotalAmount(v int32)`

SetTotalAmount sets TotalAmount field to given value.


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


