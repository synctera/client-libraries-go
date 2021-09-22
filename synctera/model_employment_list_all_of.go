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

// EmploymentListAllOf struct for EmploymentListAllOf
type EmploymentListAllOf struct {
	// Array of customer employment records
	Employment []Employment `json:"employment"`
}

// NewEmploymentListAllOf instantiates a new EmploymentListAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEmploymentListAllOf(employment []Employment) *EmploymentListAllOf {
	this := EmploymentListAllOf{}
	this.Employment = employment
	return &this
}

// NewEmploymentListAllOfWithDefaults instantiates a new EmploymentListAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEmploymentListAllOfWithDefaults() *EmploymentListAllOf {
	this := EmploymentListAllOf{}
	return &this
}

// GetEmployment returns the Employment field value
func (o *EmploymentListAllOf) GetEmployment() []Employment {
	if o == nil {
		var ret []Employment
		return ret
	}

	return o.Employment
}

// GetEmploymentOk returns a tuple with the Employment field value
// and a boolean to check if the value has been set.
func (o *EmploymentListAllOf) GetEmploymentOk() (*[]Employment, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Employment, true
}

// SetEmployment sets field value
func (o *EmploymentListAllOf) SetEmployment(v []Employment) {
	o.Employment = v
}

func (o EmploymentListAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["employment"] = o.Employment
	}
	return json.Marshal(toSerialize)
}

type NullableEmploymentListAllOf struct {
	value *EmploymentListAllOf
	isSet bool
}

func (v NullableEmploymentListAllOf) Get() *EmploymentListAllOf {
	return v.value
}

func (v *NullableEmploymentListAllOf) Set(val *EmploymentListAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableEmploymentListAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableEmploymentListAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEmploymentListAllOf(val *EmploymentListAllOf) *NullableEmploymentListAllOf {
	return &NullableEmploymentListAllOf{value: val, isSet: true}
}

func (v NullableEmploymentListAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEmploymentListAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


