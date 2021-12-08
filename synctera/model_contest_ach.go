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

// ContestAch Contested return
type ContestAch struct {
	Type string `json:"type"`
}

// NewContestAch instantiates a new ContestAch object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewContestAch(type_ string) *ContestAch {
	this := ContestAch{}
	this.Type = type_
	return &this
}

// NewContestAchWithDefaults instantiates a new ContestAch object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewContestAchWithDefaults() *ContestAch {
	this := ContestAch{}
	return &this
}

// GetType returns the Type field value
func (o *ContestAch) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ContestAch) GetTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *ContestAch) SetType(v string) {
	o.Type = v
}

func (o ContestAch) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["type"] = o.Type
	}
	return json.Marshal(toSerialize)
}

type NullableContestAch struct {
	value *ContestAch
	isSet bool
}

func (v NullableContestAch) Get() *ContestAch {
	return v.value
}

func (v *NullableContestAch) Set(val *ContestAch) {
	v.value = val
	v.isSet = true
}

func (v NullableContestAch) IsSet() bool {
	return v.isSet
}

func (v *NullableContestAch) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableContestAch(val *ContestAch) *NullableContestAch {
	return &NullableContestAch{value: val, isSet: true}
}

func (v NullableContestAch) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableContestAch) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


