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

// AliasList struct for AliasList
type AliasList struct {
	// Array of account alias
	Aliases []Alias `json:"aliases"`
}

// NewAliasList instantiates a new AliasList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAliasList(aliases []Alias) *AliasList {
	this := AliasList{}
	this.Aliases = aliases
	return &this
}

// NewAliasListWithDefaults instantiates a new AliasList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAliasListWithDefaults() *AliasList {
	this := AliasList{}
	return &this
}

// GetAliases returns the Aliases field value
func (o *AliasList) GetAliases() []Alias {
	if o == nil {
		var ret []Alias
		return ret
	}

	return o.Aliases
}

// GetAliasesOk returns a tuple with the Aliases field value
// and a boolean to check if the value has been set.
func (o *AliasList) GetAliasesOk() ([]Alias, bool) {
	if o == nil {
		return nil, false
	}
	return o.Aliases, true
}

// SetAliases sets field value
func (o *AliasList) SetAliases(v []Alias) {
	o.Aliases = v
}

func (o AliasList) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["aliases"] = o.Aliases
	}
	return json.Marshal(toSerialize)
}

type NullableAliasList struct {
	value *AliasList
	isSet bool
}

func (v NullableAliasList) Get() *AliasList {
	return v.value
}

func (v *NullableAliasList) Set(val *AliasList) {
	v.value = val
	v.isSet = true
}

func (v NullableAliasList) IsSet() bool {
	return v.isSet
}

func (v *NullableAliasList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAliasList(val *AliasList) *NullableAliasList {
	return &NullableAliasList{value: val, isSet: true}
}

func (v NullableAliasList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAliasList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
