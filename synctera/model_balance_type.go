/*
Synctera API

<h2>Let's build something great.</h2><p>Welcome to the official reference documentation for Synctera APIs. Our APIs are the best way to automate your company's banking needs and are designed to be easy to understand and implement.</p><p>We're continuously growing this library and what you see here is just the start, but if you need something specific or have a question, <a class='text-blue-600' href='https://synctera.com/contact' target='_blank' rel='noreferrer'>contact us</a>.</p> 

API version: 0.5.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"fmt"
)

// BalanceType The type of the balance. GET operation default is for all types
type BalanceType string

// List of balance_type
const (
	BALANCETYPE_ACCOUNT_BALANCE BalanceType = "ACCOUNT_BALANCE"
	BALANCETYPE_AVAILABLE_BALANCE BalanceType = "AVAILABLE_BALANCE"
)

var allowedBalanceTypeEnumValues = []BalanceType{
	"ACCOUNT_BALANCE",
	"AVAILABLE_BALANCE",
}

func (v *BalanceType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := BalanceType(value)
	for _, existing := range allowedBalanceTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid BalanceType", value)
}

// NewBalanceTypeFromValue returns a pointer to a valid BalanceType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewBalanceTypeFromValue(v string) (*BalanceType, error) {
	ev := BalanceType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for BalanceType: valid values are %v", v, allowedBalanceTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v BalanceType) IsValid() bool {
	for _, existing := range allowedBalanceTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to balance_type value
func (v BalanceType) Ptr() *BalanceType {
	return &v
}

type NullableBalanceType struct {
	value *BalanceType
	isSet bool
}

func (v NullableBalanceType) Get() *BalanceType {
	return v.value
}

func (v *NullableBalanceType) Set(val *BalanceType) {
	v.value = val
	v.isSet = true
}

func (v NullableBalanceType) IsSet() bool {
	return v.isSet
}

func (v *NullableBalanceType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBalanceType(val *BalanceType) *NullableBalanceType {
	return &NullableBalanceType{value: val, isSet: true}
}

func (v NullableBalanceType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBalanceType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

