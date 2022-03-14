/*
Synctera API

<h2>Let's build something great.</h2><p>Welcome to the official reference documentation for Synctera APIs. Our APIs are the best way to automate your company's banking needs and are designed to be easy to understand and implement.</p><p>We're continuously growing this library and what you see here is just the start, but if you need something specific or have a question, <a class='text-blue-600' href='https://synctera.com/contact' target='_blank' rel='noreferrer'>contact us</a>.</p>

API version: 0.5.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package synctera

import (
	"encoding/json"
)

// CardPin struct for CardPin
type CardPin struct {
	// The new PIN
	Pin *string `json:"pin,omitempty"`
}

// NewCardPin instantiates a new CardPin object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCardPin() *CardPin {
	this := CardPin{}
	return &this
}

// NewCardPinWithDefaults instantiates a new CardPin object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCardPinWithDefaults() *CardPin {
	this := CardPin{}
	return &this
}

// GetPin returns the Pin field value if set, zero value otherwise.
func (o *CardPin) GetPin() string {
	if o == nil || o.Pin == nil {
		var ret string
		return ret
	}
	return *o.Pin
}

// GetPinOk returns a tuple with the Pin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CardPin) GetPinOk() (*string, bool) {
	if o == nil || o.Pin == nil {
		return nil, false
	}
	return o.Pin, true
}

// HasPin returns a boolean if a field has been set.
func (o *CardPin) HasPin() bool {
	if o != nil && o.Pin != nil {
		return true
	}

	return false
}

// SetPin gets a reference to the given string and assigns it to the Pin field.
func (o *CardPin) SetPin(v string) {
	o.Pin = &v
}

func (o CardPin) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Pin != nil {
		toSerialize["pin"] = o.Pin
	}
	return json.Marshal(toSerialize)
}

type NullableCardPin struct {
	value *CardPin
	isSet bool
}

func (v NullableCardPin) Get() *CardPin {
	return v.value
}

func (v *NullableCardPin) Set(val *CardPin) {
	v.value = val
	v.isSet = true
}

func (v NullableCardPin) IsSet() bool {
	return v.isSet
}

func (v *NullableCardPin) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCardPin(val *CardPin) *NullableCardPin {
	return &NullableCardPin{value: val, isSet: true}
}

func (v NullableCardPin) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCardPin) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
