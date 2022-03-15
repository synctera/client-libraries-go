/*
Synctera API

<h2>Let's build something great.</h2><p>Welcome to the official reference documentation for Synctera APIs. Our APIs are the best way to automate your company's banking needs and are designed to be easy to understand and implement.</p><p>We're continuously growing this library and what you see here is just the start, but if you need something specific or have a question, <a class='text-blue-600' href='https://synctera.com/contact' target='_blank' rel='noreferrer'>contact us</a>.</p>

API version: 0.5.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package synctera

import (
	"encoding/json"

	oapi "github.com/deepmap/oapi-codegen/pkg/types"
)

// OutgoingAch Create an outgoing ACH
type OutgoingAch struct {
	// Amount to transfer in ISO 4217 minor currency units
	Amount int32 `json:"amount"`
	// ISO 4217 alphabetic currency code of the transfer amount
	Currency string `json:"currency"`
	// The customer's unique identifier
	CustomerId string `json:"customer_id"`
	// The type of transaction (debit or credit). A debit is a transfer in and a credit is a transfer out of the originating account
	DcSign string `json:"dc_sign"`
	// Effective date transaction proccesses (is_same_day needs to be false or not present at all)
	EffectiveDate *oapi.Date `json:"effective_date,omitempty"`
	// Additional transfer metadata structured as key-value pairs
	ExternalData map[string]interface{} `json:"external_data,omitempty"`
	// ID of the international customer that receives the final remittance transfer (required for OFAC enabled payments)
	FinalCustomerId *string `json:"final_customer_id,omitempty"`
	Id              *string `json:"id,omitempty"`
	// Send as same day ACH transaction (use only is_same_day without specific effective_date)
	IsSameDay *bool `json:"is_same_day,omitempty"`
	// Memo for the payment
	Memo *string `json:"memo,omitempty"`
	// The unique identifier for an originating account
	OriginatingAccountId string `json:"originating_account_id"`
	// The unique identifier for an receiving account
	ReceivingAccountId string `json:"receiving_account_id"`
	// Will be sent to the ACH network and maps to Addenda record 05 - the recipient bank will receive this info
	ReferenceInfo *string  `json:"reference_info,omitempty"`
	Risk          RiskData `json:"risk"`
}

// NewOutgoingAch instantiates a new OutgoingAch object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOutgoingAch(amount int32, currency string, customerId string, dcSign string, originatingAccountId string, receivingAccountId string, risk RiskData) *OutgoingAch {
	this := OutgoingAch{}
	this.Amount = amount
	this.Currency = currency
	this.CustomerId = customerId
	this.DcSign = dcSign
	this.OriginatingAccountId = originatingAccountId
	this.ReceivingAccountId = receivingAccountId
	this.Risk = risk
	return &this
}

// NewOutgoingAchWithDefaults instantiates a new OutgoingAch object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOutgoingAchWithDefaults() *OutgoingAch {
	this := OutgoingAch{}
	return &this
}

// GetAmount returns the Amount field value
func (o *OutgoingAch) GetAmount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Amount
}

// GetAmountOk returns a tuple with the Amount field value
// and a boolean to check if the value has been set.
func (o *OutgoingAch) GetAmountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Amount, true
}

// SetAmount sets field value
func (o *OutgoingAch) SetAmount(v int32) {
	o.Amount = v
}

// GetCurrency returns the Currency field value
func (o *OutgoingAch) GetCurrency() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Currency
}

// GetCurrencyOk returns a tuple with the Currency field value
// and a boolean to check if the value has been set.
func (o *OutgoingAch) GetCurrencyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Currency, true
}

// SetCurrency sets field value
func (o *OutgoingAch) SetCurrency(v string) {
	o.Currency = v
}

// GetCustomerId returns the CustomerId field value
func (o *OutgoingAch) GetCustomerId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CustomerId
}

// GetCustomerIdOk returns a tuple with the CustomerId field value
// and a boolean to check if the value has been set.
func (o *OutgoingAch) GetCustomerIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CustomerId, true
}

// SetCustomerId sets field value
func (o *OutgoingAch) SetCustomerId(v string) {
	o.CustomerId = v
}

// GetDcSign returns the DcSign field value
func (o *OutgoingAch) GetDcSign() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DcSign
}

// GetDcSignOk returns a tuple with the DcSign field value
// and a boolean to check if the value has been set.
func (o *OutgoingAch) GetDcSignOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DcSign, true
}

// SetDcSign sets field value
func (o *OutgoingAch) SetDcSign(v string) {
	o.DcSign = v
}

// GetEffectiveDate returns the EffectiveDate field value if set, zero value otherwise.
func (o *OutgoingAch) GetEffectiveDate() oapi.Date {
	if o == nil || o.EffectiveDate == nil {
		var ret oapi.Date
		return ret
	}
	return *o.EffectiveDate
}

// GetEffectiveDateOk returns a tuple with the EffectiveDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OutgoingAch) GetEffectiveDateOk() (*oapi.Date, bool) {
	if o == nil || o.EffectiveDate == nil {
		return nil, false
	}
	return o.EffectiveDate, true
}

// HasEffectiveDate returns a boolean if a field has been set.
func (o *OutgoingAch) HasEffectiveDate() bool {
	if o != nil && o.EffectiveDate != nil {
		return true
	}

	return false
}

// SetEffectiveDate gets a reference to the given oapi.Date and assigns it to the EffectiveDate field.
func (o *OutgoingAch) SetEffectiveDate(v oapi.Date) {
	o.EffectiveDate = &v
}

// GetExternalData returns the ExternalData field value if set, zero value otherwise.
func (o *OutgoingAch) GetExternalData() map[string]interface{} {
	if o == nil || o.ExternalData == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.ExternalData
}

// GetExternalDataOk returns a tuple with the ExternalData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OutgoingAch) GetExternalDataOk() (map[string]interface{}, bool) {
	if o == nil || o.ExternalData == nil {
		return nil, false
	}
	return o.ExternalData, true
}

// HasExternalData returns a boolean if a field has been set.
func (o *OutgoingAch) HasExternalData() bool {
	if o != nil && o.ExternalData != nil {
		return true
	}

	return false
}

// SetExternalData gets a reference to the given map[string]interface{} and assigns it to the ExternalData field.
func (o *OutgoingAch) SetExternalData(v map[string]interface{}) {
	o.ExternalData = v
}

// GetFinalCustomerId returns the FinalCustomerId field value if set, zero value otherwise.
func (o *OutgoingAch) GetFinalCustomerId() string {
	if o == nil || o.FinalCustomerId == nil {
		var ret string
		return ret
	}
	return *o.FinalCustomerId
}

// GetFinalCustomerIdOk returns a tuple with the FinalCustomerId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OutgoingAch) GetFinalCustomerIdOk() (*string, bool) {
	if o == nil || o.FinalCustomerId == nil {
		return nil, false
	}
	return o.FinalCustomerId, true
}

// HasFinalCustomerId returns a boolean if a field has been set.
func (o *OutgoingAch) HasFinalCustomerId() bool {
	if o != nil && o.FinalCustomerId != nil {
		return true
	}

	return false
}

// SetFinalCustomerId gets a reference to the given string and assigns it to the FinalCustomerId field.
func (o *OutgoingAch) SetFinalCustomerId(v string) {
	o.FinalCustomerId = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *OutgoingAch) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OutgoingAch) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *OutgoingAch) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *OutgoingAch) SetId(v string) {
	o.Id = &v
}

// GetIsSameDay returns the IsSameDay field value if set, zero value otherwise.
func (o *OutgoingAch) GetIsSameDay() bool {
	if o == nil || o.IsSameDay == nil {
		var ret bool
		return ret
	}
	return *o.IsSameDay
}

// GetIsSameDayOk returns a tuple with the IsSameDay field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OutgoingAch) GetIsSameDayOk() (*bool, bool) {
	if o == nil || o.IsSameDay == nil {
		return nil, false
	}
	return o.IsSameDay, true
}

// HasIsSameDay returns a boolean if a field has been set.
func (o *OutgoingAch) HasIsSameDay() bool {
	if o != nil && o.IsSameDay != nil {
		return true
	}

	return false
}

// SetIsSameDay gets a reference to the given bool and assigns it to the IsSameDay field.
func (o *OutgoingAch) SetIsSameDay(v bool) {
	o.IsSameDay = &v
}

// GetMemo returns the Memo field value if set, zero value otherwise.
func (o *OutgoingAch) GetMemo() string {
	if o == nil || o.Memo == nil {
		var ret string
		return ret
	}
	return *o.Memo
}

// GetMemoOk returns a tuple with the Memo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OutgoingAch) GetMemoOk() (*string, bool) {
	if o == nil || o.Memo == nil {
		return nil, false
	}
	return o.Memo, true
}

// HasMemo returns a boolean if a field has been set.
func (o *OutgoingAch) HasMemo() bool {
	if o != nil && o.Memo != nil {
		return true
	}

	return false
}

// SetMemo gets a reference to the given string and assigns it to the Memo field.
func (o *OutgoingAch) SetMemo(v string) {
	o.Memo = &v
}

// GetOriginatingAccountId returns the OriginatingAccountId field value
func (o *OutgoingAch) GetOriginatingAccountId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OriginatingAccountId
}

// GetOriginatingAccountIdOk returns a tuple with the OriginatingAccountId field value
// and a boolean to check if the value has been set.
func (o *OutgoingAch) GetOriginatingAccountIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OriginatingAccountId, true
}

// SetOriginatingAccountId sets field value
func (o *OutgoingAch) SetOriginatingAccountId(v string) {
	o.OriginatingAccountId = v
}

// GetReceivingAccountId returns the ReceivingAccountId field value
func (o *OutgoingAch) GetReceivingAccountId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ReceivingAccountId
}

// GetReceivingAccountIdOk returns a tuple with the ReceivingAccountId field value
// and a boolean to check if the value has been set.
func (o *OutgoingAch) GetReceivingAccountIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ReceivingAccountId, true
}

// SetReceivingAccountId sets field value
func (o *OutgoingAch) SetReceivingAccountId(v string) {
	o.ReceivingAccountId = v
}

// GetReferenceInfo returns the ReferenceInfo field value if set, zero value otherwise.
func (o *OutgoingAch) GetReferenceInfo() string {
	if o == nil || o.ReferenceInfo == nil {
		var ret string
		return ret
	}
	return *o.ReferenceInfo
}

// GetReferenceInfoOk returns a tuple with the ReferenceInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OutgoingAch) GetReferenceInfoOk() (*string, bool) {
	if o == nil || o.ReferenceInfo == nil {
		return nil, false
	}
	return o.ReferenceInfo, true
}

// HasReferenceInfo returns a boolean if a field has been set.
func (o *OutgoingAch) HasReferenceInfo() bool {
	if o != nil && o.ReferenceInfo != nil {
		return true
	}

	return false
}

// SetReferenceInfo gets a reference to the given string and assigns it to the ReferenceInfo field.
func (o *OutgoingAch) SetReferenceInfo(v string) {
	o.ReferenceInfo = &v
}

// GetRisk returns the Risk field value
func (o *OutgoingAch) GetRisk() RiskData {
	if o == nil {
		var ret RiskData
		return ret
	}

	return o.Risk
}

// GetRiskOk returns a tuple with the Risk field value
// and a boolean to check if the value has been set.
func (o *OutgoingAch) GetRiskOk() (*RiskData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Risk, true
}

// SetRisk sets field value
func (o *OutgoingAch) SetRisk(v RiskData) {
	o.Risk = v
}

func (o OutgoingAch) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["amount"] = o.Amount
	}
	if true {
		toSerialize["currency"] = o.Currency
	}
	if true {
		toSerialize["customer_id"] = o.CustomerId
	}
	if true {
		toSerialize["dc_sign"] = o.DcSign
	}
	if o.EffectiveDate != nil {
		toSerialize["effective_date"] = o.EffectiveDate
	}
	if o.ExternalData != nil {
		toSerialize["external_data"] = o.ExternalData
	}
	if o.FinalCustomerId != nil {
		toSerialize["final_customer_id"] = o.FinalCustomerId
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.IsSameDay != nil {
		toSerialize["is_same_day"] = o.IsSameDay
	}
	if o.Memo != nil {
		toSerialize["memo"] = o.Memo
	}
	if true {
		toSerialize["originating_account_id"] = o.OriginatingAccountId
	}
	if true {
		toSerialize["receiving_account_id"] = o.ReceivingAccountId
	}
	if o.ReferenceInfo != nil {
		toSerialize["reference_info"] = o.ReferenceInfo
	}
	if true {
		toSerialize["risk"] = o.Risk
	}
	return json.Marshal(toSerialize)
}

type NullableOutgoingAch struct {
	value *OutgoingAch
	isSet bool
}

func (v NullableOutgoingAch) Get() *OutgoingAch {
	return v.value
}

func (v *NullableOutgoingAch) Set(val *OutgoingAch) {
	v.value = val
	v.isSet = true
}

func (v NullableOutgoingAch) IsSet() bool {
	return v.isSet
}

func (v *NullableOutgoingAch) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOutgoingAch(val *OutgoingAch) *NullableOutgoingAch {
	return &NullableOutgoingAch{value: val, isSet: true}
}

func (v NullableOutgoingAch) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOutgoingAch) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
