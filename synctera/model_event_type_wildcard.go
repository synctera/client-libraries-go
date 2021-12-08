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

// EventTypeWildcard the model 'EventTypeWildcard'
type EventTypeWildcard string

// List of event_type_wildcard
const (
	EVENTTYPEWILDCARD_ACCOUNT EventTypeWildcard = "ACCOUNT.*"
	EVENTTYPEWILDCARD_CARD EventTypeWildcard = "CARD.*"
	EVENTTYPEWILDCARD_CUSTOMER EventTypeWildcard = "CUSTOMER.*"
	EVENTTYPEWILDCARD_INTEREST EventTypeWildcard = "INTEREST.*"
)

var allowedEventTypeWildcardEnumValues = []EventTypeWildcard{
	"ACCOUNT.*",
	"CARD.*",
	"CUSTOMER.*",
	"INTEREST.*",
}

func (v *EventTypeWildcard) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EventTypeWildcard(value)
	for _, existing := range allowedEventTypeWildcardEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid EventTypeWildcard", value)
}

// NewEventTypeWildcardFromValue returns a pointer to a valid EventTypeWildcard
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEventTypeWildcardFromValue(v string) (*EventTypeWildcard, error) {
	ev := EventTypeWildcard(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EventTypeWildcard: valid values are %v", v, allowedEventTypeWildcardEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EventTypeWildcard) IsValid() bool {
	for _, existing := range allowedEventTypeWildcardEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to event_type_wildcard value
func (v EventTypeWildcard) Ptr() *EventTypeWildcard {
	return &v
}

type NullableEventTypeWildcard struct {
	value *EventTypeWildcard
	isSet bool
}

func (v NullableEventTypeWildcard) Get() *EventTypeWildcard {
	return v.value
}

func (v *NullableEventTypeWildcard) Set(val *EventTypeWildcard) {
	v.value = val
	v.isSet = true
}

func (v NullableEventTypeWildcard) IsSet() bool {
	return v.isSet
}

func (v *NullableEventTypeWildcard) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEventTypeWildcard(val *EventTypeWildcard) *NullableEventTypeWildcard {
	return &NullableEventTypeWildcard{value: val, isSet: true}
}

func (v NullableEventTypeWildcard) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEventTypeWildcard) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

