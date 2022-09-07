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

// DeviceType The type of device into which the digital wallet token will be provisioned.
type DeviceType string

// List of device_type
const (
	DEVICETYPE_MOBILE_PHONE DeviceType = "MOBILE_PHONE"
	DEVICETYPE_WATCH DeviceType = "WATCH"
	DEVICETYPE_TABLET DeviceType = "TABLET"
)

// All allowed values of DeviceType enum
var AllowedDeviceTypeEnumValues = []DeviceType{
	"MOBILE_PHONE",
	"WATCH",
	"TABLET",
}

func (v *DeviceType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := DeviceType(value)
	for _, existing := range AllowedDeviceTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid DeviceType", value)
}

// NewDeviceTypeFromValue returns a pointer to a valid DeviceType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewDeviceTypeFromValue(v string) (*DeviceType, error) {
	ev := DeviceType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for DeviceType: valid values are %v", v, AllowedDeviceTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v DeviceType) IsValid() bool {
	for _, existing := range AllowedDeviceTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to device_type value
func (v DeviceType) Ptr() *DeviceType {
	return &v
}

type NullableDeviceType struct {
	value *DeviceType
	isSet bool
}

func (v NullableDeviceType) Get() *DeviceType {
	return v.value
}

func (v *NullableDeviceType) Set(val *DeviceType) {
	v.value = val
	v.isSet = true
}

func (v NullableDeviceType) IsSet() bool {
	return v.isSet
}

func (v *NullableDeviceType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeviceType(val *DeviceType) *NullableDeviceType {
	return &NullableDeviceType{value: val, isSet: true}
}

func (v NullableDeviceType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeviceType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
