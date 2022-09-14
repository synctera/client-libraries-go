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

// CardChangeState - struct for CardChangeState
type CardChangeState struct {
	CardFulfillmentStatus *CardFulfillmentStatus
	CardPinStatus         *CardPinStatus
	CardStatus            *CardStatus
}

// CardFulfillmentStatusAsCardChangeState is a convenience function that returns CardFulfillmentStatus wrapped in CardChangeState
func CardFulfillmentStatusAsCardChangeState(v *CardFulfillmentStatus) CardChangeState {
	return CardChangeState{
		CardFulfillmentStatus: v,
	}
}

// CardPinStatusAsCardChangeState is a convenience function that returns CardPinStatus wrapped in CardChangeState
func CardPinStatusAsCardChangeState(v *CardPinStatus) CardChangeState {
	return CardChangeState{
		CardPinStatus: v,
	}
}

// CardStatusAsCardChangeState is a convenience function that returns CardStatus wrapped in CardChangeState
func CardStatusAsCardChangeState(v *CardStatus) CardChangeState {
	return CardChangeState{
		CardStatus: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *CardChangeState) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into CardFulfillmentStatus
	err = json.Unmarshal(data, &dst.CardFulfillmentStatus)
	if err == nil {
		jsonCardFulfillmentStatus, _ := json.Marshal(dst.CardFulfillmentStatus)
		if string(jsonCardFulfillmentStatus) == "{}" { // empty struct
			dst.CardFulfillmentStatus = nil
		} else {
			match++
		}
	} else {
		dst.CardFulfillmentStatus = nil
	}

	// try to unmarshal data into CardPinStatus
	err = json.Unmarshal(data, &dst.CardPinStatus)
	if err == nil {
		jsonCardPinStatus, _ := json.Marshal(dst.CardPinStatus)
		if string(jsonCardPinStatus) == "{}" { // empty struct
			dst.CardPinStatus = nil
		} else {
			match++
		}
	} else {
		dst.CardPinStatus = nil
	}

	// try to unmarshal data into CardStatus
	err = json.Unmarshal(data, &dst.CardStatus)
	if err == nil {
		jsonCardStatus, _ := json.Marshal(dst.CardStatus)
		if string(jsonCardStatus) == "{}" { // empty struct
			dst.CardStatus = nil
		} else {
			match++
		}
	} else {
		dst.CardStatus = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.CardFulfillmentStatus = nil
		dst.CardPinStatus = nil
		dst.CardStatus = nil

		return fmt.Errorf("Data matches more than one schema in oneOf(CardChangeState)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("Data failed to match schemas in oneOf(CardChangeState)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src CardChangeState) MarshalJSON() ([]byte, error) {
	if src.CardFulfillmentStatus != nil {
		return json.Marshal(&src.CardFulfillmentStatus)
	}

	if src.CardPinStatus != nil {
		return json.Marshal(&src.CardPinStatus)
	}

	if src.CardStatus != nil {
		return json.Marshal(&src.CardStatus)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *CardChangeState) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.CardFulfillmentStatus != nil {
		return obj.CardFulfillmentStatus
	}

	if obj.CardPinStatus != nil {
		return obj.CardPinStatus
	}

	if obj.CardStatus != nil {
		return obj.CardStatus
	}

	// all schemas are nil
	return nil
}

type NullableCardChangeState struct {
	value *CardChangeState
	isSet bool
}

func (v NullableCardChangeState) Get() *CardChangeState {
	return v.value
}

func (v *NullableCardChangeState) Set(val *CardChangeState) {
	v.value = val
	v.isSet = true
}

func (v NullableCardChangeState) IsSet() bool {
	return v.isSet
}

func (v *NullableCardChangeState) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCardChangeState(val *CardChangeState) *NullableCardChangeState {
	return &NullableCardChangeState{value: val, isSet: true}
}

func (v NullableCardChangeState) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCardChangeState) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
