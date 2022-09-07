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

// ApplicationType Type of Credit Application
type ApplicationType string

// List of application_type
const (
	APPLICATIONTYPE_LINE_OF_CREDIT ApplicationType = "LINE_OF_CREDIT"
)

// All allowed values of ApplicationType enum
var AllowedApplicationTypeEnumValues = []ApplicationType{
	"LINE_OF_CREDIT",
}

func (v *ApplicationType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ApplicationType(value)
	for _, existing := range AllowedApplicationTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ApplicationType", value)
}

// NewApplicationTypeFromValue returns a pointer to a valid ApplicationType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewApplicationTypeFromValue(v string) (*ApplicationType, error) {
	ev := ApplicationType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ApplicationType: valid values are %v", v, AllowedApplicationTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ApplicationType) IsValid() bool {
	for _, existing := range AllowedApplicationTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to application_type value
func (v ApplicationType) Ptr() *ApplicationType {
	return &v
}

type NullableApplicationType struct {
	value *ApplicationType
	isSet bool
}

func (v NullableApplicationType) Get() *ApplicationType {
	return v.value
}

func (v *NullableApplicationType) Set(val *ApplicationType) {
	v.value = val
	v.isSet = true
}

func (v NullableApplicationType) IsSet() bool {
	return v.isSet
}

func (v *NullableApplicationType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApplicationType(val *ApplicationType) *NullableApplicationType {
	return &NullableApplicationType{value: val, isSet: true}
}

func (v NullableApplicationType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApplicationType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

