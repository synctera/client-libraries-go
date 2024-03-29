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

// SpendingLimitsTransaction Individual transaction limit
type SpendingLimitsTransaction struct {
	// Maximum amount allowed. Unit in cents.
	Amount *int64 `json:"amount,omitempty"`
}

// NewSpendingLimitsTransaction instantiates a new SpendingLimitsTransaction object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSpendingLimitsTransaction() *SpendingLimitsTransaction {
	this := SpendingLimitsTransaction{}
	return &this
}

// NewSpendingLimitsTransactionWithDefaults instantiates a new SpendingLimitsTransaction object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSpendingLimitsTransactionWithDefaults() *SpendingLimitsTransaction {
	this := SpendingLimitsTransaction{}
	return &this
}

// GetAmount returns the Amount field value if set, zero value otherwise.
func (o *SpendingLimitsTransaction) GetAmount() int64 {
	if o == nil || o.Amount == nil {
		var ret int64
		return ret
	}
	return *o.Amount
}

// GetAmountOk returns a tuple with the Amount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpendingLimitsTransaction) GetAmountOk() (*int64, bool) {
	if o == nil || o.Amount == nil {
		return nil, false
	}
	return o.Amount, true
}

// HasAmount returns a boolean if a field has been set.
func (o *SpendingLimitsTransaction) HasAmount() bool {
	if o != nil && o.Amount != nil {
		return true
	}

	return false
}

// SetAmount gets a reference to the given int64 and assigns it to the Amount field.
func (o *SpendingLimitsTransaction) SetAmount(v int64) {
	o.Amount = &v
}

func (o SpendingLimitsTransaction) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Amount != nil {
		toSerialize["amount"] = o.Amount
	}
	return json.Marshal(toSerialize)
}

type NullableSpendingLimitsTransaction struct {
	value *SpendingLimitsTransaction
	isSet bool
}

func (v NullableSpendingLimitsTransaction) Get() *SpendingLimitsTransaction {
	return v.value
}

func (v *NullableSpendingLimitsTransaction) Set(val *SpendingLimitsTransaction) {
	v.value = val
	v.isSet = true
}

func (v NullableSpendingLimitsTransaction) IsSet() bool {
	return v.isSet
}

func (v *NullableSpendingLimitsTransaction) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSpendingLimitsTransaction(val *SpendingLimitsTransaction) *NullableSpendingLimitsTransaction {
	return &NullableSpendingLimitsTransaction{value: val, isSet: true}
}

func (v NullableSpendingLimitsTransaction) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSpendingLimitsTransaction) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
