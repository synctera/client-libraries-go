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
	"fmt"
)

// Form PHYSICAL or VIRTUAL.
type Form string

// List of form
const (
	FORM_PHYSICAL Form = "PHYSICAL"
	FORM_VIRTUAL Form = "VIRTUAL"
)

var allowedFormEnumValues = []Form{
	"PHYSICAL",
	"VIRTUAL",
}

func (v *Form) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := Form(value)
	for _, existing := range allowedFormEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid Form", value)
}

// NewFormFromValue returns a pointer to a valid Form
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewFormFromValue(v string) (*Form, error) {
	ev := Form(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for Form: valid values are %v", v, allowedFormEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v Form) IsValid() bool {
	for _, existing := range allowedFormEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to form value
func (v Form) Ptr() *Form {
	return &v
}

type NullableForm struct {
	value *Form
	isSet bool
}

func (v NullableForm) Get() *Form {
	return v.value
}

func (v *NullableForm) Set(val *Form) {
	v.value = val
	v.isSet = true
}

func (v NullableForm) IsSet() bool {
	return v.isSet
}

func (v *NullableForm) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableForm(val *Form) *NullableForm {
	return &NullableForm{value: val, isSet: true}
}

func (v NullableForm) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableForm) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

