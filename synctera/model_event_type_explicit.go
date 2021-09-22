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

// EventTypeExplicit All the webhook event types
type EventTypeExplicit string

// List of event_type_explicit
const (
	EVENTTYPEEXPLICIT_ACCOUNT_CREATED EventTypeExplicit = "ACCOUNT.CREATED"
	EVENTTYPEEXPLICIT_CARD_STATUS_CHANGE EventTypeExplicit = "CARD.STATUS_CHANGE"
)

var allowedEventTypeExplicitEnumValues = []EventTypeExplicit{
	"ACCOUNT.CREATED",
	"CARD.STATUS_CHANGE",
}

func (v *EventTypeExplicit) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EventTypeExplicit(value)
	for _, existing := range allowedEventTypeExplicitEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid EventTypeExplicit", value)
}

// NewEventTypeExplicitFromValue returns a pointer to a valid EventTypeExplicit
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEventTypeExplicitFromValue(v string) (*EventTypeExplicit, error) {
	ev := EventTypeExplicit(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EventTypeExplicit: valid values are %v", v, allowedEventTypeExplicitEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EventTypeExplicit) IsValid() bool {
	for _, existing := range allowedEventTypeExplicitEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to event_type_explicit value
func (v EventTypeExplicit) Ptr() *EventTypeExplicit {
	return &v
}

type NullableEventTypeExplicit struct {
	value *EventTypeExplicit
	isSet bool
}

func (v NullableEventTypeExplicit) Get() *EventTypeExplicit {
	return v.value
}

func (v *NullableEventTypeExplicit) Set(val *EventTypeExplicit) {
	v.value = val
	v.isSet = true
}

func (v NullableEventTypeExplicit) IsSet() bool {
	return v.isSet
}

func (v *NullableEventTypeExplicit) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEventTypeExplicit(val *EventTypeExplicit) *NullableEventTypeExplicit {
	return &NullableEventTypeExplicit{value: val, isSet: true}
}

func (v NullableEventTypeExplicit) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEventTypeExplicit) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
