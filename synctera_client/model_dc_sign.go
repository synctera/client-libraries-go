/*
Synctera API

<h2>Let's build something great.</h2><p>Welcome to the official reference documentation for Synctera APIs. Our APIs are the best way to automate your company's banking needs and are designed to be easy to understand and implement.</p><p>We're continuously growing this library and what you see here is just the start, but if you need something specific or have a question, <a class='text-blue-600' href='https://synctera.com/contact' target='_blank' rel='noreferrer'>contact us</a>.</p>

API version: 0.20.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package synctera_client

import (
	"encoding/json"
	"fmt"
)

// DcSign The `dc_sign` represents the direction money was moved.
type DcSign string

// List of dc_sign
const (
	DCSIGN_DEBIT  DcSign = "debit"
	DCSIGN_CREDIT DcSign = "credit"
)

// All allowed values of DcSign enum
var AllowedDcSignEnumValues = []DcSign{
	"debit",
	"credit",
}

func (v *DcSign) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := DcSign(value)
	for _, existing := range AllowedDcSignEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid DcSign", value)
}

// NewDcSignFromValue returns a pointer to a valid DcSign
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewDcSignFromValue(v string) (*DcSign, error) {
	ev := DcSign(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for DcSign: valid values are %v", v, AllowedDcSignEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v DcSign) IsValid() bool {
	for _, existing := range AllowedDcSignEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to dc_sign value
func (v DcSign) Ptr() *DcSign {
	return &v
}

type NullableDcSign struct {
	value *DcSign
	isSet bool
}

func (v NullableDcSign) Get() *DcSign {
	return v.value
}

func (v *NullableDcSign) Set(val *DcSign) {
	v.value = val
	v.isSet = true
}

func (v NullableDcSign) IsSet() bool {
	return v.isSet
}

func (v *NullableDcSign) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDcSign(val *DcSign) *NullableDcSign {
	return &NullableDcSign{value: val, isSet: true}
}

func (v NullableDcSign) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDcSign) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
