/*
Synctera API

<h2>Let's build something great.</h2><p>Welcome to the official reference documentation for Synctera APIs. Our APIs are the best way to automate your company's banking needs and are designed to be easy to understand and implement.</p><p>We're continuously growing this library and what you see here is just the start, but if you need something specific or have a question, <a class='text-blue-600' href='https://synctera.com/contact' target='_blank' rel='noreferrer'>contact us</a>.</p> 

API version: 0.5.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package synctera

import (
	"encoding/json"
	"time"
	"fmt"
)

// CardIssuanceRequest - struct for CardIssuanceRequest
type CardIssuanceRequest struct {
	PhysicalCardIssuanceRequest *PhysicalCardIssuanceRequest
	VirtualCardIssuanceRequest *VirtualCardIssuanceRequest
}

// PhysicalCardIssuanceRequestAsCardIssuanceRequest is a convenience function that returns PhysicalCardIssuanceRequest wrapped in CardIssuanceRequest
func PhysicalCardIssuanceRequestAsCardIssuanceRequest(v *PhysicalCardIssuanceRequest) CardIssuanceRequest {
	return CardIssuanceRequest{ PhysicalCardIssuanceRequest: v}
}

// VirtualCardIssuanceRequestAsCardIssuanceRequest is a convenience function that returns VirtualCardIssuanceRequest wrapped in CardIssuanceRequest
func VirtualCardIssuanceRequestAsCardIssuanceRequest(v *VirtualCardIssuanceRequest) CardIssuanceRequest {
	return CardIssuanceRequest{ VirtualCardIssuanceRequest: v}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *CardIssuanceRequest) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into PhysicalCardIssuanceRequest
	err = json.Unmarshal(data, &dst.PhysicalCardIssuanceRequest)
	if err == nil {
		jsonPhysicalCardIssuanceRequest, _ := json.Marshal(dst.PhysicalCardIssuanceRequest)
		if string(jsonPhysicalCardIssuanceRequest) == "{}" { // empty struct
			dst.PhysicalCardIssuanceRequest = nil
		} else {
			match++
		}
	} else {
		dst.PhysicalCardIssuanceRequest = nil
	}

	// try to unmarshal data into VirtualCardIssuanceRequest
	err = json.Unmarshal(data, &dst.VirtualCardIssuanceRequest)
	if err == nil {
		jsonVirtualCardIssuanceRequest, _ := json.Marshal(dst.VirtualCardIssuanceRequest)
		if string(jsonVirtualCardIssuanceRequest) == "{}" { // empty struct
			dst.VirtualCardIssuanceRequest = nil
		} else {
			match++
		}
	} else {
		dst.VirtualCardIssuanceRequest = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.PhysicalCardIssuanceRequest = nil
		dst.VirtualCardIssuanceRequest = nil

		return fmt.Errorf("Data matches more than one schema in oneOf(CardIssuanceRequest)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("Data failed to match schemas in oneOf(CardIssuanceRequest)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src CardIssuanceRequest) MarshalJSON() ([]byte, error) {
	if src.PhysicalCardIssuanceRequest != nil {
		return json.Marshal(&src.PhysicalCardIssuanceRequest)
	}

	if src.VirtualCardIssuanceRequest != nil {
		return json.Marshal(&src.VirtualCardIssuanceRequest)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *CardIssuanceRequest) GetActualInstance() (interface{}) {
	if obj.PhysicalCardIssuanceRequest != nil {
		return obj.PhysicalCardIssuanceRequest
	}

	if obj.VirtualCardIssuanceRequest != nil {
		return obj.VirtualCardIssuanceRequest
	}

	// all schemas are nil
	return nil
}

type NullableCardIssuanceRequest struct {
	value *CardIssuanceRequest
	isSet bool
}

func (v NullableCardIssuanceRequest) Get() *CardIssuanceRequest {
	return v.value
}

func (v *NullableCardIssuanceRequest) Set(val *CardIssuanceRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCardIssuanceRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCardIssuanceRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCardIssuanceRequest(val *CardIssuanceRequest) *NullableCardIssuanceRequest {
	return &NullableCardIssuanceRequest{value: val, isSet: true}
}

func (v NullableCardIssuanceRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCardIssuanceRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


