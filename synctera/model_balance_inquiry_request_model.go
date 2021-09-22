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

// BalanceInquiryRequestModel struct for BalanceInquiryRequestModel
type BalanceInquiryRequestModel struct {
	AccountType string `json:"account_type"`
	CardAcceptor CardAcceptorModel `json:"card_acceptor"`
	CardToken string `json:"card_token"`
	Mid string `json:"mid"`
	NetworkFees *[]NetworkFeeModel `json:"network_fees,omitempty"`
	Pin *string `json:"pin,omitempty"`
}

// NewBalanceInquiryRequestModel instantiates a new BalanceInquiryRequestModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBalanceInquiryRequestModel(accountType string, cardAcceptor CardAcceptorModel, cardToken string, mid string) *BalanceInquiryRequestModel {
	this := BalanceInquiryRequestModel{}
	this.AccountType = accountType
	this.CardAcceptor = cardAcceptor
	this.CardToken = cardToken
	this.Mid = mid
	return &this
}

// NewBalanceInquiryRequestModelWithDefaults instantiates a new BalanceInquiryRequestModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBalanceInquiryRequestModelWithDefaults() *BalanceInquiryRequestModel {
	this := BalanceInquiryRequestModel{}
	return &this
}

// GetAccountType returns the AccountType field value
func (o *BalanceInquiryRequestModel) GetAccountType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AccountType
}

// GetAccountTypeOk returns a tuple with the AccountType field value
// and a boolean to check if the value has been set.
func (o *BalanceInquiryRequestModel) GetAccountTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.AccountType, true
}

// SetAccountType sets field value
func (o *BalanceInquiryRequestModel) SetAccountType(v string) {
	o.AccountType = v
}

// GetCardAcceptor returns the CardAcceptor field value
func (o *BalanceInquiryRequestModel) GetCardAcceptor() CardAcceptorModel {
	if o == nil {
		var ret CardAcceptorModel
		return ret
	}

	return o.CardAcceptor
}

// GetCardAcceptorOk returns a tuple with the CardAcceptor field value
// and a boolean to check if the value has been set.
func (o *BalanceInquiryRequestModel) GetCardAcceptorOk() (*CardAcceptorModel, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.CardAcceptor, true
}

// SetCardAcceptor sets field value
func (o *BalanceInquiryRequestModel) SetCardAcceptor(v CardAcceptorModel) {
	o.CardAcceptor = v
}

// GetCardToken returns the CardToken field value
func (o *BalanceInquiryRequestModel) GetCardToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CardToken
}

// GetCardTokenOk returns a tuple with the CardToken field value
// and a boolean to check if the value has been set.
func (o *BalanceInquiryRequestModel) GetCardTokenOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.CardToken, true
}

// SetCardToken sets field value
func (o *BalanceInquiryRequestModel) SetCardToken(v string) {
	o.CardToken = v
}

// GetMid returns the Mid field value
func (o *BalanceInquiryRequestModel) GetMid() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Mid
}

// GetMidOk returns a tuple with the Mid field value
// and a boolean to check if the value has been set.
func (o *BalanceInquiryRequestModel) GetMidOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Mid, true
}

// SetMid sets field value
func (o *BalanceInquiryRequestModel) SetMid(v string) {
	o.Mid = v
}

// GetNetworkFees returns the NetworkFees field value if set, zero value otherwise.
func (o *BalanceInquiryRequestModel) GetNetworkFees() []NetworkFeeModel {
	if o == nil || o.NetworkFees == nil {
		var ret []NetworkFeeModel
		return ret
	}
	return *o.NetworkFees
}

// GetNetworkFeesOk returns a tuple with the NetworkFees field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BalanceInquiryRequestModel) GetNetworkFeesOk() (*[]NetworkFeeModel, bool) {
	if o == nil || o.NetworkFees == nil {
		return nil, false
	}
	return o.NetworkFees, true
}

// HasNetworkFees returns a boolean if a field has been set.
func (o *BalanceInquiryRequestModel) HasNetworkFees() bool {
	if o != nil && o.NetworkFees != nil {
		return true
	}

	return false
}

// SetNetworkFees gets a reference to the given []NetworkFeeModel and assigns it to the NetworkFees field.
func (o *BalanceInquiryRequestModel) SetNetworkFees(v []NetworkFeeModel) {
	o.NetworkFees = &v
}

// GetPin returns the Pin field value if set, zero value otherwise.
func (o *BalanceInquiryRequestModel) GetPin() string {
	if o == nil || o.Pin == nil {
		var ret string
		return ret
	}
	return *o.Pin
}

// GetPinOk returns a tuple with the Pin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BalanceInquiryRequestModel) GetPinOk() (*string, bool) {
	if o == nil || o.Pin == nil {
		return nil, false
	}
	return o.Pin, true
}

// HasPin returns a boolean if a field has been set.
func (o *BalanceInquiryRequestModel) HasPin() bool {
	if o != nil && o.Pin != nil {
		return true
	}

	return false
}

// SetPin gets a reference to the given string and assigns it to the Pin field.
func (o *BalanceInquiryRequestModel) SetPin(v string) {
	o.Pin = &v
}

func (o BalanceInquiryRequestModel) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["account_type"] = o.AccountType
	}
	if true {
		toSerialize["card_acceptor"] = o.CardAcceptor
	}
	if true {
		toSerialize["card_token"] = o.CardToken
	}
	if true {
		toSerialize["mid"] = o.Mid
	}
	if o.NetworkFees != nil {
		toSerialize["network_fees"] = o.NetworkFees
	}
	if o.Pin != nil {
		toSerialize["pin"] = o.Pin
	}
	return json.Marshal(toSerialize)
}

type NullableBalanceInquiryRequestModel struct {
	value *BalanceInquiryRequestModel
	isSet bool
}

func (v NullableBalanceInquiryRequestModel) Get() *BalanceInquiryRequestModel {
	return v.value
}

func (v *NullableBalanceInquiryRequestModel) Set(val *BalanceInquiryRequestModel) {
	v.value = val
	v.isSet = true
}

func (v NullableBalanceInquiryRequestModel) IsSet() bool {
	return v.isSet
}

func (v *NullableBalanceInquiryRequestModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBalanceInquiryRequestModel(val *BalanceInquiryRequestModel) *NullableBalanceInquiryRequestModel {
	return &NullableBalanceInquiryRequestModel{value: val, isSet: true}
}

func (v NullableBalanceInquiryRequestModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBalanceInquiryRequestModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

