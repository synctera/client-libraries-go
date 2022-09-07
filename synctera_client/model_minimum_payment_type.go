/*
Synctera API

<h2>Let's build something great.</h2><p>Welcome to the official reference documentation for Synctera APIs. Our APIs are the best way to automate your company's banking needs and are designed to be easy to understand and implement.</p><p>We're continuously growing this library and what you see here is just the start, but if you need something specific or have a question, <a class='text-blue-600' href='https://synctera.com/contact' target='_blank' rel='noreferrer'>contact us</a>.</p> 

API version: 0.19.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package synctera_client

import (
	"encoding/json"
	"fmt"
)

// MinimumPaymentType the model 'MinimumPaymentType'
type MinimumPaymentType string

// List of minimum_payment_type
const (
	MINIMUMPAYMENTTYPE_RATE_OR_AMOUNT MinimumPaymentType = "RATE_OR_AMOUNT"
)

// All allowed values of MinimumPaymentType enum
var AllowedMinimumPaymentTypeEnumValues = []MinimumPaymentType{
	"RATE_OR_AMOUNT",
}

func (v *MinimumPaymentType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := MinimumPaymentType(value)
	for _, existing := range AllowedMinimumPaymentTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid MinimumPaymentType", value)
}

// NewMinimumPaymentTypeFromValue returns a pointer to a valid MinimumPaymentType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewMinimumPaymentTypeFromValue(v string) (*MinimumPaymentType, error) {
	ev := MinimumPaymentType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for MinimumPaymentType: valid values are %v", v, AllowedMinimumPaymentTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v MinimumPaymentType) IsValid() bool {
	for _, existing := range AllowedMinimumPaymentTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to minimum_payment_type value
func (v MinimumPaymentType) Ptr() *MinimumPaymentType {
	return &v
}

type NullableMinimumPaymentType struct {
	value *MinimumPaymentType
	isSet bool
}

func (v NullableMinimumPaymentType) Get() *MinimumPaymentType {
	return v.value
}

func (v *NullableMinimumPaymentType) Set(val *MinimumPaymentType) {
	v.value = val
	v.isSet = true
}

func (v NullableMinimumPaymentType) IsSet() bool {
	return v.isSet
}

func (v *NullableMinimumPaymentType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMinimumPaymentType(val *MinimumPaymentType) *NullableMinimumPaymentType {
	return &NullableMinimumPaymentType{value: val, isSet: true}
}

func (v NullableMinimumPaymentType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMinimumPaymentType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

