/*
Synctera API

<h2>Let's build something great.</h2><p>Welcome to the official reference documentation for Synctera APIs. Our APIs are the best way to automate your company's banking needs and are designed to be easy to understand and implement.</p><p>We're continuously growing this library and what you see here is just the start, but if you need something specific or have a question, <a class='text-blue-600' href='https://synctera.com/contact' target='_blank' rel='noreferrer'>contact us</a>.</p> 

API version: 0.5.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package synctera

import (
	"encoding/json"
	"time"
)

// PendingTransactionData struct for PendingTransactionData
type PendingTransactionData struct {
	// The amount of the hold.
	Amount int32 `json:"amount"`
	// The account \"available balance\" at the time this hold was created
	AvailBalance int32 `json:"avail_balance"`
	// The account balance at the time this hold was created
	Balance int32 `json:"balance"`
	// ISO 4217 alphabetic currency code of the transfer amount
	Currency string `json:"currency"`
	DcSign DcSign `json:"dc_sign"`
	// The effective date of the transaction once it gets posted
	EffectiveDate time.Time `json:"effective_date"`
	// The date that at which this hold is no longer valid.
	ExpiresAt time.Time `json:"expires_at"`
	// Whether or not the hold was forced (spending controls ignored)
	ForcePost bool `json:"force_post"`
	// An array representing any previous states of the hold, if it has been modified (For example, increasing or decreasing the hold amount).
	History []PendingTransactionHistory `json:"history"`
	// The network this transaction is associated with
	Network string `json:"network"`
	Operation string `json:"operation"`
	// If a hold has been declined or modified, this will include the reason.
	Reason string `json:"reason"`
	// The requested amount, in the case of hold modifications.
	ReqAmount int32 `json:"req_amount"`
	// Information received by the transaction risk/fraud service related to this transaction
	RiskInfo map[string]interface{} `json:"risk_info"`
	// The status of the hold.
	Status string `json:"status"`
	// The specific transaction type. For example, for `ach`, this may be \"outgoing_debit\".
	Subtype string `json:"subtype"`
	// The total amount of the hold. This may be different than `amount` in the case where a hold increase or decrease was requested.
	TotalAmount int32 `json:"total_amount"`
	// The general type of transaction. For example, \"card\" or \"ach\".
	Type string `json:"type"`
	// An unstructured JSON blob representing additional transaction information specific to each payment rail.
	UserData map[string]interface{} `json:"user_data"`
	// Does this hold represent a partial debit (or credit)?
	WasPartial bool `json:"was_partial"`
}

// NewPendingTransactionData instantiates a new PendingTransactionData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPendingTransactionData(amount int32, availBalance int32, balance int32, currency string, dcSign DcSign, effectiveDate time.Time, expiresAt time.Time, forcePost bool, history []PendingTransactionHistory, network string, operation string, reason string, reqAmount int32, riskInfo map[string]interface{}, status string, subtype string, totalAmount int32, type_ string, userData map[string]interface{}, wasPartial bool) *PendingTransactionData {
	this := PendingTransactionData{}
	this.Amount = amount
	this.AvailBalance = availBalance
	this.Balance = balance
	this.Currency = currency
	this.DcSign = dcSign
	this.EffectiveDate = effectiveDate
	this.ExpiresAt = expiresAt
	this.ForcePost = forcePost
	this.History = history
	this.Network = network
	this.Operation = operation
	this.Reason = reason
	this.ReqAmount = reqAmount
	this.RiskInfo = riskInfo
	this.Status = status
	this.Subtype = subtype
	this.TotalAmount = totalAmount
	this.Type = type_
	this.UserData = userData
	this.WasPartial = wasPartial
	return &this
}

// NewPendingTransactionDataWithDefaults instantiates a new PendingTransactionData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPendingTransactionDataWithDefaults() *PendingTransactionData {
	this := PendingTransactionData{}
	return &this
}

// GetAmount returns the Amount field value
func (o *PendingTransactionData) GetAmount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Amount
}

// GetAmountOk returns a tuple with the Amount field value
// and a boolean to check if the value has been set.
func (o *PendingTransactionData) GetAmountOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Amount, true
}

// SetAmount sets field value
func (o *PendingTransactionData) SetAmount(v int32) {
	o.Amount = v
}

// GetAvailBalance returns the AvailBalance field value
func (o *PendingTransactionData) GetAvailBalance() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.AvailBalance
}

// GetAvailBalanceOk returns a tuple with the AvailBalance field value
// and a boolean to check if the value has been set.
func (o *PendingTransactionData) GetAvailBalanceOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.AvailBalance, true
}

// SetAvailBalance sets field value
func (o *PendingTransactionData) SetAvailBalance(v int32) {
	o.AvailBalance = v
}

// GetBalance returns the Balance field value
func (o *PendingTransactionData) GetBalance() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Balance
}

// GetBalanceOk returns a tuple with the Balance field value
// and a boolean to check if the value has been set.
func (o *PendingTransactionData) GetBalanceOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Balance, true
}

// SetBalance sets field value
func (o *PendingTransactionData) SetBalance(v int32) {
	o.Balance = v
}

// GetCurrency returns the Currency field value
func (o *PendingTransactionData) GetCurrency() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Currency
}

// GetCurrencyOk returns a tuple with the Currency field value
// and a boolean to check if the value has been set.
func (o *PendingTransactionData) GetCurrencyOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Currency, true
}

// SetCurrency sets field value
func (o *PendingTransactionData) SetCurrency(v string) {
	o.Currency = v
}

// GetDcSign returns the DcSign field value
func (o *PendingTransactionData) GetDcSign() DcSign {
	if o == nil {
		var ret DcSign
		return ret
	}

	return o.DcSign
}

// GetDcSignOk returns a tuple with the DcSign field value
// and a boolean to check if the value has been set.
func (o *PendingTransactionData) GetDcSignOk() (*DcSign, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.DcSign, true
}

// SetDcSign sets field value
func (o *PendingTransactionData) SetDcSign(v DcSign) {
	o.DcSign = v
}

// GetEffectiveDate returns the EffectiveDate field value
func (o *PendingTransactionData) GetEffectiveDate() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.EffectiveDate
}

// GetEffectiveDateOk returns a tuple with the EffectiveDate field value
// and a boolean to check if the value has been set.
func (o *PendingTransactionData) GetEffectiveDateOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.EffectiveDate, true
}

// SetEffectiveDate sets field value
func (o *PendingTransactionData) SetEffectiveDate(v time.Time) {
	o.EffectiveDate = v
}

// GetExpiresAt returns the ExpiresAt field value
func (o *PendingTransactionData) GetExpiresAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.ExpiresAt
}

// GetExpiresAtOk returns a tuple with the ExpiresAt field value
// and a boolean to check if the value has been set.
func (o *PendingTransactionData) GetExpiresAtOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ExpiresAt, true
}

// SetExpiresAt sets field value
func (o *PendingTransactionData) SetExpiresAt(v time.Time) {
	o.ExpiresAt = v
}

// GetForcePost returns the ForcePost field value
func (o *PendingTransactionData) GetForcePost() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.ForcePost
}

// GetForcePostOk returns a tuple with the ForcePost field value
// and a boolean to check if the value has been set.
func (o *PendingTransactionData) GetForcePostOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ForcePost, true
}

// SetForcePost sets field value
func (o *PendingTransactionData) SetForcePost(v bool) {
	o.ForcePost = v
}

// GetHistory returns the History field value
func (o *PendingTransactionData) GetHistory() []PendingTransactionHistory {
	if o == nil {
		var ret []PendingTransactionHistory
		return ret
	}

	return o.History
}

// GetHistoryOk returns a tuple with the History field value
// and a boolean to check if the value has been set.
func (o *PendingTransactionData) GetHistoryOk() (*[]PendingTransactionHistory, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.History, true
}

// SetHistory sets field value
func (o *PendingTransactionData) SetHistory(v []PendingTransactionHistory) {
	o.History = v
}

// GetNetwork returns the Network field value
func (o *PendingTransactionData) GetNetwork() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Network
}

// GetNetworkOk returns a tuple with the Network field value
// and a boolean to check if the value has been set.
func (o *PendingTransactionData) GetNetworkOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Network, true
}

// SetNetwork sets field value
func (o *PendingTransactionData) SetNetwork(v string) {
	o.Network = v
}

// GetOperation returns the Operation field value
func (o *PendingTransactionData) GetOperation() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Operation
}

// GetOperationOk returns a tuple with the Operation field value
// and a boolean to check if the value has been set.
func (o *PendingTransactionData) GetOperationOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Operation, true
}

// SetOperation sets field value
func (o *PendingTransactionData) SetOperation(v string) {
	o.Operation = v
}

// GetReason returns the Reason field value
func (o *PendingTransactionData) GetReason() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Reason
}

// GetReasonOk returns a tuple with the Reason field value
// and a boolean to check if the value has been set.
func (o *PendingTransactionData) GetReasonOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Reason, true
}

// SetReason sets field value
func (o *PendingTransactionData) SetReason(v string) {
	o.Reason = v
}

// GetReqAmount returns the ReqAmount field value
func (o *PendingTransactionData) GetReqAmount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.ReqAmount
}

// GetReqAmountOk returns a tuple with the ReqAmount field value
// and a boolean to check if the value has been set.
func (o *PendingTransactionData) GetReqAmountOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ReqAmount, true
}

// SetReqAmount sets field value
func (o *PendingTransactionData) SetReqAmount(v int32) {
	o.ReqAmount = v
}

// GetRiskInfo returns the RiskInfo field value
func (o *PendingTransactionData) GetRiskInfo() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.RiskInfo
}

// GetRiskInfoOk returns a tuple with the RiskInfo field value
// and a boolean to check if the value has been set.
func (o *PendingTransactionData) GetRiskInfoOk() (*map[string]interface{}, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.RiskInfo, true
}

// SetRiskInfo sets field value
func (o *PendingTransactionData) SetRiskInfo(v map[string]interface{}) {
	o.RiskInfo = v
}

// GetStatus returns the Status field value
func (o *PendingTransactionData) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *PendingTransactionData) GetStatusOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *PendingTransactionData) SetStatus(v string) {
	o.Status = v
}

// GetSubtype returns the Subtype field value
func (o *PendingTransactionData) GetSubtype() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Subtype
}

// GetSubtypeOk returns a tuple with the Subtype field value
// and a boolean to check if the value has been set.
func (o *PendingTransactionData) GetSubtypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Subtype, true
}

// SetSubtype sets field value
func (o *PendingTransactionData) SetSubtype(v string) {
	o.Subtype = v
}

// GetTotalAmount returns the TotalAmount field value
func (o *PendingTransactionData) GetTotalAmount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.TotalAmount
}

// GetTotalAmountOk returns a tuple with the TotalAmount field value
// and a boolean to check if the value has been set.
func (o *PendingTransactionData) GetTotalAmountOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.TotalAmount, true
}

// SetTotalAmount sets field value
func (o *PendingTransactionData) SetTotalAmount(v int32) {
	o.TotalAmount = v
}

// GetType returns the Type field value
func (o *PendingTransactionData) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *PendingTransactionData) GetTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *PendingTransactionData) SetType(v string) {
	o.Type = v
}

// GetUserData returns the UserData field value
func (o *PendingTransactionData) GetUserData() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.UserData
}

// GetUserDataOk returns a tuple with the UserData field value
// and a boolean to check if the value has been set.
func (o *PendingTransactionData) GetUserDataOk() (*map[string]interface{}, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.UserData, true
}

// SetUserData sets field value
func (o *PendingTransactionData) SetUserData(v map[string]interface{}) {
	o.UserData = v
}

// GetWasPartial returns the WasPartial field value
func (o *PendingTransactionData) GetWasPartial() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.WasPartial
}

// GetWasPartialOk returns a tuple with the WasPartial field value
// and a boolean to check if the value has been set.
func (o *PendingTransactionData) GetWasPartialOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.WasPartial, true
}

// SetWasPartial sets field value
func (o *PendingTransactionData) SetWasPartial(v bool) {
	o.WasPartial = v
}

func (o PendingTransactionData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["amount"] = o.Amount
	}
	if true {
		toSerialize["avail_balance"] = o.AvailBalance
	}
	if true {
		toSerialize["balance"] = o.Balance
	}
	if true {
		toSerialize["currency"] = o.Currency
	}
	if true {
		toSerialize["dc_sign"] = o.DcSign
	}
	if true {
		toSerialize["effective_date"] = o.EffectiveDate
	}
	if true {
		toSerialize["expires_at"] = o.ExpiresAt
	}
	if true {
		toSerialize["force_post"] = o.ForcePost
	}
	if true {
		toSerialize["history"] = o.History
	}
	if true {
		toSerialize["network"] = o.Network
	}
	if true {
		toSerialize["operation"] = o.Operation
	}
	if true {
		toSerialize["reason"] = o.Reason
	}
	if true {
		toSerialize["req_amount"] = o.ReqAmount
	}
	if true {
		toSerialize["risk_info"] = o.RiskInfo
	}
	if true {
		toSerialize["status"] = o.Status
	}
	if true {
		toSerialize["subtype"] = o.Subtype
	}
	if true {
		toSerialize["total_amount"] = o.TotalAmount
	}
	if true {
		toSerialize["type"] = o.Type
	}
	if true {
		toSerialize["user_data"] = o.UserData
	}
	if true {
		toSerialize["was_partial"] = o.WasPartial
	}
	return json.Marshal(toSerialize)
}

type NullablePendingTransactionData struct {
	value *PendingTransactionData
	isSet bool
}

func (v NullablePendingTransactionData) Get() *PendingTransactionData {
	return v.value
}

func (v *NullablePendingTransactionData) Set(val *PendingTransactionData) {
	v.value = val
	v.isSet = true
}

func (v NullablePendingTransactionData) IsSet() bool {
	return v.isSet
}

func (v *NullablePendingTransactionData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePendingTransactionData(val *PendingTransactionData) *NullablePendingTransactionData {
	return &NullablePendingTransactionData{value: val, isSet: true}
}

func (v NullablePendingTransactionData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePendingTransactionData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


