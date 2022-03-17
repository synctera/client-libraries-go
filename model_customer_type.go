/*
Synctera API

<h2>Let's build something great.</h2><p>Welcome to the official reference documentation for Synctera APIs. Our APIs are the best way to automate your company's banking needs and are designed to be easy to understand and implement.</p><p>We're continuously growing this library and what you see here is just the start, but if you need something specific or have a question, <a class='text-blue-600' href='https://synctera.com/contact' target='_blank' rel='noreferrer'>contact us</a>.</p>

API version: 0.5.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package synctera

import (
	"encoding/json"
	"fmt"
)

// CustomerType Customer type
type CustomerType string

// List of customer_type
const (
	CUSTOMERTYPE_BUSINESS CustomerType = "BUSINESS"
	CUSTOMERTYPE_PERSONAL CustomerType = "PERSONAL"
)

var allowedCustomerTypeEnumValues = []CustomerType{
	"BUSINESS",
	"PERSONAL",
}

func (v *CustomerType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := CustomerType(value)
	for _, existing := range allowedCustomerTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid CustomerType", value)
}

// NewCustomerTypeFromValue returns a pointer to a valid CustomerType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewCustomerTypeFromValue(v string) (*CustomerType, error) {
	ev := CustomerType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for CustomerType: valid values are %v", v, allowedCustomerTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v CustomerType) IsValid() bool {
	for _, existing := range allowedCustomerTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to customer_type value
func (v CustomerType) Ptr() *CustomerType {
	return &v
}

type NullableCustomerType struct {
	value *CustomerType
	isSet bool
}

func (v NullableCustomerType) Get() *CustomerType {
	return v.value
}

func (v *NullableCustomerType) Set(val *CustomerType) {
	v.value = val
	v.isSet = true
}

func (v NullableCustomerType) IsSet() bool {
	return v.isSet
}

func (v *NullableCustomerType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCustomerType(val *CustomerType) *NullableCustomerType {
	return &NullableCustomerType{value: val, isSet: true}
}

func (v NullableCustomerType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCustomerType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
