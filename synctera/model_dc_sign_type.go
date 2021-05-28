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

// DcSignType Indication of debit or credit
type DcSignType string

// List of dc_sign_type
const (
	DCSIGNTYPE_DEBIT DcSignType = "DEBIT"
	DCSIGNTYPE_CREDIT DcSignType = "CREDIT"
)

var allowedDcSignTypeEnumValues = []DcSignType{
	"DEBIT",
	"CREDIT",
}

func (v *DcSignType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := DcSignType(value)
	for _, existing := range allowedDcSignTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid DcSignType", value)
}

// NewDcSignTypeFromValue returns a pointer to a valid DcSignType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewDcSignTypeFromValue(v string) (*DcSignType, error) {
	ev := DcSignType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for DcSignType: valid values are %v", v, allowedDcSignTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v DcSignType) IsValid() bool {
	for _, existing := range allowedDcSignTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to dc_sign_type value
func (v DcSignType) Ptr() *DcSignType {
	return &v
}

type NullableDcSignType struct {
	value *DcSignType
	isSet bool
}

func (v NullableDcSignType) Get() *DcSignType {
	return v.value
}

func (v *NullableDcSignType) Set(val *DcSignType) {
	v.value = val
	v.isSet = true
}

func (v NullableDcSignType) IsSet() bool {
	return v.isSet
}

func (v *NullableDcSignType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDcSignType(val *DcSignType) *NullableDcSignType {
	return &NullableDcSignType{value: val, isSet: true}
}

func (v NullableDcSignType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDcSignType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
