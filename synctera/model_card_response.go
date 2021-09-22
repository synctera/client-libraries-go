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
	"time"
	"fmt"
)

// CardResponse - struct for CardResponse
type CardResponse struct {
	PhysicalCardResponse *PhysicalCardResponse
	VirtualCardResponse *VirtualCardResponse
}

// PhysicalCardResponseAsCardResponse is a convenience function that returns PhysicalCardResponse wrapped in CardResponse
func PhysicalCardResponseAsCardResponse(v *PhysicalCardResponse) CardResponse {
	return CardResponse{ PhysicalCardResponse: v}
}

// VirtualCardResponseAsCardResponse is a convenience function that returns VirtualCardResponse wrapped in CardResponse
func VirtualCardResponseAsCardResponse(v *VirtualCardResponse) CardResponse {
	return CardResponse{ VirtualCardResponse: v}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *CardResponse) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into PhysicalCardResponse
	err = json.Unmarshal(data, &dst.PhysicalCardResponse)
	if err == nil {
		jsonPhysicalCardResponse, _ := json.Marshal(dst.PhysicalCardResponse)
		if string(jsonPhysicalCardResponse) == "{}" { // empty struct
			dst.PhysicalCardResponse = nil
		} else {
			match++
		}
	} else {
		dst.PhysicalCardResponse = nil
	}

	// try to unmarshal data into VirtualCardResponse
	err = json.Unmarshal(data, &dst.VirtualCardResponse)
	if err == nil {
		jsonVirtualCardResponse, _ := json.Marshal(dst.VirtualCardResponse)
		if string(jsonVirtualCardResponse) == "{}" { // empty struct
			dst.VirtualCardResponse = nil
		} else {
			match++
		}
	} else {
		dst.VirtualCardResponse = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.PhysicalCardResponse = nil
		dst.VirtualCardResponse = nil

		return fmt.Errorf("Data matches more than one schema in oneOf(CardResponse)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("Data failed to match schemas in oneOf(CardResponse)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src CardResponse) MarshalJSON() ([]byte, error) {
	if src.PhysicalCardResponse != nil {
		return json.Marshal(&src.PhysicalCardResponse)
	}

	if src.VirtualCardResponse != nil {
		return json.Marshal(&src.VirtualCardResponse)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *CardResponse) GetActualInstance() (interface{}) {
	if obj.PhysicalCardResponse != nil {
		return obj.PhysicalCardResponse
	}

	if obj.VirtualCardResponse != nil {
		return obj.VirtualCardResponse
	}

	// all schemas are nil
	return nil
}

type NullableCardResponse struct {
	value *CardResponse
	isSet bool
}

func (v NullableCardResponse) Get() *CardResponse {
	return v.value
}

func (v *NullableCardResponse) Set(val *CardResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCardResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCardResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCardResponse(val *CardResponse) *NullableCardResponse {
	return &NullableCardResponse{value: val, isSet: true}
}

func (v NullableCardResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCardResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

