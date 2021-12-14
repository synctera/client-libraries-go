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

// BinStatus Bin status
type BinStatus string

// List of bin_status
const (
	BINSTATUS_ACTIVE BinStatus = "ACTIVE"
	BINSTATUS_INACTIVE BinStatus = "INACTIVE"
)

var allowedBinStatusEnumValues = []BinStatus{
	"ACTIVE",
	"INACTIVE",
}

func (v *BinStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := BinStatus(value)
	for _, existing := range allowedBinStatusEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid BinStatus", value)
}

// NewBinStatusFromValue returns a pointer to a valid BinStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewBinStatusFromValue(v string) (*BinStatus, error) {
	ev := BinStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for BinStatus: valid values are %v", v, allowedBinStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v BinStatus) IsValid() bool {
	for _, existing := range allowedBinStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to bin_status value
func (v BinStatus) Ptr() *BinStatus {
	return &v
}

type NullableBinStatus struct {
	value *BinStatus
	isSet bool
}

func (v NullableBinStatus) Get() *BinStatus {
	return v.value
}

func (v *NullableBinStatus) Set(val *BinStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableBinStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableBinStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBinStatus(val *BinStatus) *NullableBinStatus {
	return &NullableBinStatus{value: val, isSet: true}
}

func (v NullableBinStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBinStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

