# PendingTransactionHistoryData

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

### NewPendingTransactionHistoryData

`func NewPendingTransactionHistoryData(amount int32, availBalance int32, balance int32, currency string, dcSign DcSign, effectiveDate time.Time, expiresAt time.Time, forcePost bool, network string, operation string, reason string, reqAmount int32, riskInfo map[string]interface{}, status string, subtype string, totalAmount int32, type_ string, userData map[string]interface{}, wasPartial bool, ) *PendingTransactionHistoryData`

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

`func (o *PendingTransactionHistoryData) GetAmount() int32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *PendingTransactionHistoryData) GetAmountOk() (*int32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *PendingTransactionHistoryData) SetAmount(v int32)`

SetAmount sets Amount field to given value.


### GetAvailBalance

`func (o *PendingTransactionHistoryData) GetAvailBalance() int32`

GetAvailBalance returns the AvailBalance field if non-nil, zero value otherwise.

### GetAvailBalanceOk

`func (o *PendingTransactionHistoryData) GetAvailBalanceOk() (*int32, bool)`

GetAvailBalanceOk returns a tuple with the AvailBalance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailBalance

`func (o *PendingTransactionHistoryData) SetAvailBalance(v int32)`

SetAvailBalance sets AvailBalance field to given value.


### GetBalance

`func (o *PendingTransactionHistoryData) GetBalance() int32`

GetBalance returns the Balance field if non-nil, zero value otherwise.

### GetBalanceOk

`func (o *PendingTransactionHistoryData) GetBalanceOk() (*int32, bool)`

GetBalanceOk returns a tuple with the Balance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalance

`func (o *PendingTransactionHistoryData) SetBalance(v int32)`

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

`func (o *PendingTransactionHistoryData) GetReqAmount() int32`

GetReqAmount returns the ReqAmount field if non-nil, zero value otherwise.

### GetReqAmountOk

`func (o *PendingTransactionHistoryData) GetReqAmountOk() (*int32, bool)`

GetReqAmountOk returns a tuple with the ReqAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReqAmount

`func (o *PendingTransactionHistoryData) SetReqAmount(v int32)`

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

`func (o *PendingTransactionHistoryData) GetTotalAmount() int32`

GetTotalAmount returns the TotalAmount field if non-nil, zero value otherwise.

### GetTotalAmountOk

`func (o *PendingTransactionHistoryData) GetTotalAmountOk() (*int32, bool)`

GetTotalAmountOk returns a tuple with the TotalAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalAmount

`func (o *PendingTransactionHistoryData) SetTotalAmount(v int32)`

SetTotalAmount sets TotalAmount field to given value.


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


