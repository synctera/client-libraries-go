/*
Synctera API

<h2>Let's build something great.</h2><p>Welcome to the official reference documentation for Synctera APIs. Our APIs are the best way to automate your company's banking needs and are designed to be easy to understand and implement.</p><p>We're continuously growing this library and what you see here is just the start, but if you need something specific or have a question, <a class='text-blue-600' href='https://synctera.com/contact' target='_blank' rel='noreferrer'>contact us</a>.</p> 

API version: 0.19.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package synctera_client

import (
	"encoding/json"
)

// PaymentListAllOf struct for PaymentListAllOf
type PaymentListAllOf struct {
	// Array of payments
	Payments []Payment `json:"payments"`
}

// NewPaymentListAllOf instantiates a new PaymentListAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaymentListAllOf(payments []Payment) *PaymentListAllOf {
	this := PaymentListAllOf{}
	this.Payments = payments
	return &this
}

// NewPaymentListAllOfWithDefaults instantiates a new PaymentListAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaymentListAllOfWithDefaults() *PaymentListAllOf {
	this := PaymentListAllOf{}
	return &this
}

// GetPayments returns the Payments field value
func (o *PaymentListAllOf) GetPayments() []Payment {
	if o == nil {
		var ret []Payment
		return ret
	}

	return o.Payments
}

// GetPaymentsOk returns a tuple with the Payments field value
// and a boolean to check if the value has been set.
func (o *PaymentListAllOf) GetPaymentsOk() ([]Payment, bool) {
	if o == nil {
		return nil, false
	}
	return o.Payments, true
}

// SetPayments sets field value
func (o *PaymentListAllOf) SetPayments(v []Payment) {
	o.Payments = v
}

func (o PaymentListAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["payments"] = o.Payments
	}
	return json.Marshal(toSerialize)
}

type NullablePaymentListAllOf struct {
	value *PaymentListAllOf
	isSet bool
}

func (v NullablePaymentListAllOf) Get() *PaymentListAllOf {
	return v.value
}

func (v *NullablePaymentListAllOf) Set(val *PaymentListAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullablePaymentListAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullablePaymentListAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaymentListAllOf(val *PaymentListAllOf) *NullablePaymentListAllOf {
	return &NullablePaymentListAllOf{value: val, isSet: true}
}

func (v NullablePaymentListAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaymentListAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


