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

// InlineObject struct for InlineObject
type InlineObject struct {
	// Set true to let the current secret expire in the next 24 hours. Set false to let the current secret expire immediately.
	IsRollingSecret *bool `json:"is_rolling_secret,omitempty"`
}

// NewInlineObject instantiates a new InlineObject object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject() *InlineObject {
	this := InlineObject{}
	return &this
}

// NewInlineObjectWithDefaults instantiates a new InlineObject object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObjectWithDefaults() *InlineObject {
	this := InlineObject{}
	return &this
}

// GetIsRollingSecret returns the IsRollingSecret field value if set, zero value otherwise.
func (o *InlineObject) GetIsRollingSecret() bool {
	if o == nil || o.IsRollingSecret == nil {
		var ret bool
		return ret
	}
	return *o.IsRollingSecret
}

// GetIsRollingSecretOk returns a tuple with the IsRollingSecret field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject) GetIsRollingSecretOk() (*bool, bool) {
	if o == nil || o.IsRollingSecret == nil {
		return nil, false
	}
	return o.IsRollingSecret, true
}

// HasIsRollingSecret returns a boolean if a field has been set.
func (o *InlineObject) HasIsRollingSecret() bool {
	if o != nil && o.IsRollingSecret != nil {
		return true
	}

	return false
}

// SetIsRollingSecret gets a reference to the given bool and assigns it to the IsRollingSecret field.
func (o *InlineObject) SetIsRollingSecret(v bool) {
	o.IsRollingSecret = &v
}

func (o InlineObject) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.IsRollingSecret != nil {
		toSerialize["is_rolling_secret"] = o.IsRollingSecret
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject struct {
	value *InlineObject
	isSet bool
}

func (v NullableInlineObject) Get() *InlineObject {
	return v.value
}

func (v *NullableInlineObject) Set(val *InlineObject) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject(val *InlineObject) *NullableInlineObject {
	return &NullableInlineObject{value: val, isSet: true}
}

func (v NullableInlineObject) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
