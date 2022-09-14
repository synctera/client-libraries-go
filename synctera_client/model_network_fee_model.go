/*
Synctera API

<h2>Let's build something great.</h2><p>Welcome to the official reference documentation for Synctera APIs. Our APIs are the best way to automate your company's banking needs and are designed to be easy to understand and implement.</p><p>We're continuously growing this library and what you see here is just the start, but if you need something specific or have a question, <a class='text-blue-600' href='https://synctera.com/contact' target='_blank' rel='noreferrer'>contact us</a>.</p>

API version: 0.20.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package synctera_client

import (
	"encoding/json"
)

// NetworkFeeModel struct for NetworkFeeModel
type NetworkFeeModel struct {
	// The amount of the fee in the smallest whole denomination of the applicable currency (eg. For USD use cents)
	Amount int32 `json:"amount"`
	// C = credit; D = debit
	CreditDebit *string `json:"credit_debit,omitempty"`
	Type        string  `json:"type"`
}

// NewNetworkFeeModel instantiates a new NetworkFeeModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworkFeeModel(amount int32, type_ string) *NetworkFeeModel {
	this := NetworkFeeModel{}
	this.Amount = amount
	this.Type = type_
	return &this
}

// NewNetworkFeeModelWithDefaults instantiates a new NetworkFeeModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworkFeeModelWithDefaults() *NetworkFeeModel {
	this := NetworkFeeModel{}
	return &this
}

// GetAmount returns the Amount field value
func (o *NetworkFeeModel) GetAmount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Amount
}

// GetAmountOk returns a tuple with the Amount field value
// and a boolean to check if the value has been set.
func (o *NetworkFeeModel) GetAmountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Amount, true
}

// SetAmount sets field value
func (o *NetworkFeeModel) SetAmount(v int32) {
	o.Amount = v
}

// GetCreditDebit returns the CreditDebit field value if set, zero value otherwise.
func (o *NetworkFeeModel) GetCreditDebit() string {
	if o == nil || o.CreditDebit == nil {
		var ret string
		return ret
	}
	return *o.CreditDebit
}

// GetCreditDebitOk returns a tuple with the CreditDebit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkFeeModel) GetCreditDebitOk() (*string, bool) {
	if o == nil || o.CreditDebit == nil {
		return nil, false
	}
	return o.CreditDebit, true
}

// HasCreditDebit returns a boolean if a field has been set.
func (o *NetworkFeeModel) HasCreditDebit() bool {
	if o != nil && o.CreditDebit != nil {
		return true
	}

	return false
}

// SetCreditDebit gets a reference to the given string and assigns it to the CreditDebit field.
func (o *NetworkFeeModel) SetCreditDebit(v string) {
	o.CreditDebit = &v
}

// GetType returns the Type field value
func (o *NetworkFeeModel) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *NetworkFeeModel) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *NetworkFeeModel) SetType(v string) {
	o.Type = v
}

func (o NetworkFeeModel) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["amount"] = o.Amount
	}
	if o.CreditDebit != nil {
		toSerialize["credit_debit"] = o.CreditDebit
	}
	if true {
		toSerialize["type"] = o.Type
	}
	return json.Marshal(toSerialize)
}

type NullableNetworkFeeModel struct {
	value *NetworkFeeModel
	isSet bool
}

func (v NullableNetworkFeeModel) Get() *NetworkFeeModel {
	return v.value
}

func (v *NullableNetworkFeeModel) Set(val *NetworkFeeModel) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworkFeeModel) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworkFeeModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworkFeeModel(val *NetworkFeeModel) *NullableNetworkFeeModel {
	return &NullableNetworkFeeModel{value: val, isSet: true}
}

func (v NullableNetworkFeeModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworkFeeModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
