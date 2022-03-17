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

// WidgetType The widget type
type WidgetType string

// List of widget_type
const (
	WIDGETTYPE_SET_PIN       WidgetType = "set_pin"
	WIDGETTYPE_ACTIVATE_CARD WidgetType = "activate_card"
)

var allowedWidgetTypeEnumValues = []WidgetType{
	"set_pin",
	"activate_card",
}

func (v *WidgetType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := WidgetType(value)
	for _, existing := range allowedWidgetTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid WidgetType", value)
}

// NewWidgetTypeFromValue returns a pointer to a valid WidgetType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewWidgetTypeFromValue(v string) (*WidgetType, error) {
	ev := WidgetType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for WidgetType: valid values are %v", v, allowedWidgetTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v WidgetType) IsValid() bool {
	for _, existing := range allowedWidgetTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to widget_type value
func (v WidgetType) Ptr() *WidgetType {
	return &v
}

type NullableWidgetType struct {
	value *WidgetType
	isSet bool
}

func (v NullableWidgetType) Get() *WidgetType {
	return v.value
}

func (v *NullableWidgetType) Set(val *WidgetType) {
	v.value = val
	v.isSet = true
}

func (v NullableWidgetType) IsSet() bool {
	return v.isSet
}

func (v *NullableWidgetType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWidgetType(val *WidgetType) *NullableWidgetType {
	return &NullableWidgetType{value: val, isSet: true}
}

func (v NullableWidgetType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWidgetType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
