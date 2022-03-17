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

// CardPinStatus The status of the card PIN
type CardPinStatus string

// List of card_pin_status
const (
	CARDPINSTATUS_SET     CardPinStatus = "SET"
	CARDPINSTATUS_CHANGED CardPinStatus = "CHANGED"
)

var allowedCardPinStatusEnumValues = []CardPinStatus{
	"SET",
	"CHANGED",
}

func (v *CardPinStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := CardPinStatus(value)
	for _, existing := range allowedCardPinStatusEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid CardPinStatus", value)
}

// NewCardPinStatusFromValue returns a pointer to a valid CardPinStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewCardPinStatusFromValue(v string) (*CardPinStatus, error) {
	ev := CardPinStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for CardPinStatus: valid values are %v", v, allowedCardPinStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v CardPinStatus) IsValid() bool {
	for _, existing := range allowedCardPinStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to card_pin_status value
func (v CardPinStatus) Ptr() *CardPinStatus {
	return &v
}

type NullableCardPinStatus struct {
	value *CardPinStatus
	isSet bool
}

func (v NullableCardPinStatus) Get() *CardPinStatus {
	return v.value
}

func (v *NullableCardPinStatus) Set(val *CardPinStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableCardPinStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableCardPinStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCardPinStatus(val *CardPinStatus) *NullableCardPinStatus {
	return &NullableCardPinStatus{value: val, isSet: true}
}

func (v NullableCardPinStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCardPinStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
