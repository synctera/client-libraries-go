/*
 * Synctera API
 *
 * <h2>Let's build something great.</h2><p>Welcome to the official reference documentation for Synctera APIs. Our APIs are the best way to automate your company's banking needs and are designed to be easy to understand and implement.</p><p>We're continuously growing this library and what you see here is just the start, but if you need something specific or have a question, <a class='text-blue-600' href='https://synctera.com/contact' target='_blank' rel='noreferrer'>contact us</a>.</p> 
 *
 * API version: 0.5.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package synctera

import (
	"encoding/json"
)

// FinancialRequestModel struct for FinancialRequestModel
type FinancialRequestModel struct {
	Amount float32 `json:"amount"`
	CardAcceptor CardAcceptorModel `json:"card_acceptor"`
	CardToken string `json:"card_token"`
	CashBackAmount *float32 `json:"cash_back_amount,omitempty"`
	IsPreAuth *bool `json:"is_pre_auth,omitempty"`
	Mid string `json:"mid"`
	Pin *string `json:"pin,omitempty"`
	TransactionOptions *TransactionOptions `json:"transaction_options,omitempty"`
}

// NewFinancialRequestModel instantiates a new FinancialRequestModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFinancialRequestModel(amount float32, cardAcceptor CardAcceptorModel, cardToken string, mid string) *FinancialRequestModel {
	this := FinancialRequestModel{}
	this.Amount = amount
	this.CardAcceptor = cardAcceptor
	this.CardToken = cardToken
	var isPreAuth bool = false
	this.IsPreAuth = &isPreAuth
	this.Mid = mid
	return &this
}

// NewFinancialRequestModelWithDefaults instantiates a new FinancialRequestModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFinancialRequestModelWithDefaults() *FinancialRequestModel {
	this := FinancialRequestModel{}
	var isPreAuth bool = false
	this.IsPreAuth = &isPreAuth
	return &this
}

// GetAmount returns the Amount field value
func (o *FinancialRequestModel) GetAmount() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Amount
}

// GetAmountOk returns a tuple with the Amount field value
// and a boolean to check if the value has been set.
func (o *FinancialRequestModel) GetAmountOk() (*float32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Amount, true
}

// SetAmount sets field value
func (o *FinancialRequestModel) SetAmount(v float32) {
	o.Amount = v
}

// GetCardAcceptor returns the CardAcceptor field value
func (o *FinancialRequestModel) GetCardAcceptor() CardAcceptorModel {
	if o == nil {
		var ret CardAcceptorModel
		return ret
	}

	return o.CardAcceptor
}

// GetCardAcceptorOk returns a tuple with the CardAcceptor field value
// and a boolean to check if the value has been set.
func (o *FinancialRequestModel) GetCardAcceptorOk() (*CardAcceptorModel, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.CardAcceptor, true
}

// SetCardAcceptor sets field value
func (o *FinancialRequestModel) SetCardAcceptor(v CardAcceptorModel) {
	o.CardAcceptor = v
}

// GetCardToken returns the CardToken field value
func (o *FinancialRequestModel) GetCardToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CardToken
}

// GetCardTokenOk returns a tuple with the CardToken field value
// and a boolean to check if the value has been set.
func (o *FinancialRequestModel) GetCardTokenOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.CardToken, true
}

// SetCardToken sets field value
func (o *FinancialRequestModel) SetCardToken(v string) {
	o.CardToken = v
}

// GetCashBackAmount returns the CashBackAmount field value if set, zero value otherwise.
func (o *FinancialRequestModel) GetCashBackAmount() float32 {
	if o == nil || o.CashBackAmount == nil {
		var ret float32
		return ret
	}
	return *o.CashBackAmount
}

// GetCashBackAmountOk returns a tuple with the CashBackAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FinancialRequestModel) GetCashBackAmountOk() (*float32, bool) {
	if o == nil || o.CashBackAmount == nil {
		return nil, false
	}
	return o.CashBackAmount, true
}

// HasCashBackAmount returns a boolean if a field has been set.
func (o *FinancialRequestModel) HasCashBackAmount() bool {
	if o != nil && o.CashBackAmount != nil {
		return true
	}

	return false
}

// SetCashBackAmount gets a reference to the given float32 and assigns it to the CashBackAmount field.
func (o *FinancialRequestModel) SetCashBackAmount(v float32) {
	o.CashBackAmount = &v
}

// GetIsPreAuth returns the IsPreAuth field value if set, zero value otherwise.
func (o *FinancialRequestModel) GetIsPreAuth() bool {
	if o == nil || o.IsPreAuth == nil {
		var ret bool
		return ret
	}
	return *o.IsPreAuth
}

// GetIsPreAuthOk returns a tuple with the IsPreAuth field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FinancialRequestModel) GetIsPreAuthOk() (*bool, bool) {
	if o == nil || o.IsPreAuth == nil {
		return nil, false
	}
	return o.IsPreAuth, true
}

// HasIsPreAuth returns a boolean if a field has been set.
func (o *FinancialRequestModel) HasIsPreAuth() bool {
	if o != nil && o.IsPreAuth != nil {
		return true
	}

	return false
}

// SetIsPreAuth gets a reference to the given bool and assigns it to the IsPreAuth field.
func (o *FinancialRequestModel) SetIsPreAuth(v bool) {
	o.IsPreAuth = &v
}

// GetMid returns the Mid field value
func (o *FinancialRequestModel) GetMid() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Mid
}

// GetMidOk returns a tuple with the Mid field value
// and a boolean to check if the value has been set.
func (o *FinancialRequestModel) GetMidOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Mid, true
}

// SetMid sets field value
func (o *FinancialRequestModel) SetMid(v string) {
	o.Mid = v
}

// GetPin returns the Pin field value if set, zero value otherwise.
func (o *FinancialRequestModel) GetPin() string {
	if o == nil || o.Pin == nil {
		var ret string
		return ret
	}
	return *o.Pin
}

// GetPinOk returns a tuple with the Pin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FinancialRequestModel) GetPinOk() (*string, bool) {
	if o == nil || o.Pin == nil {
		return nil, false
	}
	return o.Pin, true
}

// HasPin returns a boolean if a field has been set.
func (o *FinancialRequestModel) HasPin() bool {
	if o != nil && o.Pin != nil {
		return true
	}

	return false
}

// SetPin gets a reference to the given string and assigns it to the Pin field.
func (o *FinancialRequestModel) SetPin(v string) {
	o.Pin = &v
}

// GetTransactionOptions returns the TransactionOptions field value if set, zero value otherwise.
func (o *FinancialRequestModel) GetTransactionOptions() TransactionOptions {
	if o == nil || o.TransactionOptions == nil {
		var ret TransactionOptions
		return ret
	}
	return *o.TransactionOptions
}

// GetTransactionOptionsOk returns a tuple with the TransactionOptions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FinancialRequestModel) GetTransactionOptionsOk() (*TransactionOptions, bool) {
	if o == nil || o.TransactionOptions == nil {
		return nil, false
	}
	return o.TransactionOptions, true
}

// HasTransactionOptions returns a boolean if a field has been set.
func (o *FinancialRequestModel) HasTransactionOptions() bool {
	if o != nil && o.TransactionOptions != nil {
		return true
	}

	return false
}

// SetTransactionOptions gets a reference to the given TransactionOptions and assigns it to the TransactionOptions field.
func (o *FinancialRequestModel) SetTransactionOptions(v TransactionOptions) {
	o.TransactionOptions = &v
}

func (o FinancialRequestModel) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["amount"] = o.Amount
	}
	if true {
		toSerialize["card_acceptor"] = o.CardAcceptor
	}
	if true {
		toSerialize["card_token"] = o.CardToken
	}
	if o.CashBackAmount != nil {
		toSerialize["cash_back_amount"] = o.CashBackAmount
	}
	if o.IsPreAuth != nil {
		toSerialize["is_pre_auth"] = o.IsPreAuth
	}
	if true {
		toSerialize["mid"] = o.Mid
	}
	if o.Pin != nil {
		toSerialize["pin"] = o.Pin
	}
	if o.TransactionOptions != nil {
		toSerialize["transaction_options"] = o.TransactionOptions
	}
	return json.Marshal(toSerialize)
}

type NullableFinancialRequestModel struct {
	value *FinancialRequestModel
	isSet bool
}

func (v NullableFinancialRequestModel) Get() *FinancialRequestModel {
	return v.value
}

func (v *NullableFinancialRequestModel) Set(val *FinancialRequestModel) {
	v.value = val
	v.isSet = true
}

func (v NullableFinancialRequestModel) IsSet() bool {
	return v.isSet
}

func (v *NullableFinancialRequestModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFinancialRequestModel(val *FinancialRequestModel) *NullableFinancialRequestModel {
	return &NullableFinancialRequestModel{value: val, isSet: true}
}

func (v NullableFinancialRequestModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFinancialRequestModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


