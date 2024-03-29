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

// CardFormat struct for CardFormat
type CardFormat struct {
	// PHYSICAL or VIRTUAL.
	Form string `json:"form"`
}

// NewCardFormat instantiates a new CardFormat object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCardFormat(form string) *CardFormat {
	this := CardFormat{}
	this.Form = form
	return &this
}

// NewCardFormatWithDefaults instantiates a new CardFormat object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCardFormatWithDefaults() *CardFormat {
	this := CardFormat{}
	return &this
}

// GetForm returns the Form field value
func (o *CardFormat) GetForm() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Form
}

// GetFormOk returns a tuple with the Form field value
// and a boolean to check if the value has been set.
func (o *CardFormat) GetFormOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Form, true
}

// SetForm sets field value
func (o *CardFormat) SetForm(v string) {
	o.Form = v
}

func (o CardFormat) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["form"] = o.Form
	}
	return json.Marshal(toSerialize)
}

type NullableCardFormat struct {
	value *CardFormat
	isSet bool
}

func (v NullableCardFormat) Get() *CardFormat {
	return v.value
}

func (v *NullableCardFormat) Set(val *CardFormat) {
	v.value = val
	v.isSet = true
}

func (v NullableCardFormat) IsSet() bool {
	return v.isSet
}

func (v *NullableCardFormat) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCardFormat(val *CardFormat) *NullableCardFormat {
	return &NullableCardFormat{value: val, isSet: true}
}

func (v NullableCardFormat) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCardFormat) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
