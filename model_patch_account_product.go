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

// PatchAccountProduct - struct for PatchAccountProduct
type PatchAccountProduct struct {
	Fee           *Fee
	PatchInterest *PatchInterest
}

// FeeAsPatchAccountProduct is a convenience function that returns Fee wrapped in PatchAccountProduct
func FeeAsPatchAccountProduct(v *Fee) PatchAccountProduct {
	return PatchAccountProduct{Fee: v}
}

// PatchInterestAsPatchAccountProduct is a convenience function that returns PatchInterest wrapped in PatchAccountProduct
func PatchInterestAsPatchAccountProduct(v *PatchInterest) PatchAccountProduct {
	return PatchAccountProduct{PatchInterest: v}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *PatchAccountProduct) UnmarshalJSON(data []byte) error {
	var err error
	// use discriminator value to speed up the lookup
	var jsonDict map[string]interface{}
	err = json.Unmarshal(data, &jsonDict)
	if err != nil {
		return fmt.Errorf("Failed to unmarshal JSON into map for the discriminator lookup.")
	}

	// check if the discriminator value is 'FEE'
	if jsonDict["product_type"] == "FEE" {
		// try to unmarshal JSON data into Fee
		err = json.Unmarshal(data, &dst.Fee)
		if err == nil {
			return nil // data stored in dst.Fee, return on the first match
		} else {
			dst.Fee = nil
			return fmt.Errorf("Failed to unmarshal PatchAccountProduct as Fee: %s", err.Error())
		}
	}

	// check if the discriminator value is 'INTEREST'
	if jsonDict["product_type"] == "INTEREST" {
		// try to unmarshal JSON data into PatchInterest
		err = json.Unmarshal(data, &dst.PatchInterest)
		if err == nil {
			return nil // data stored in dst.PatchInterest, return on the first match
		} else {
			dst.PatchInterest = nil
			return fmt.Errorf("Failed to unmarshal PatchAccountProduct as PatchInterest: %s", err.Error())
		}
	}

	// check if the discriminator value is 'fee'
	if jsonDict["product_type"] == "fee" {
		// try to unmarshal JSON data into Fee
		err = json.Unmarshal(data, &dst.Fee)
		if err == nil {
			return nil // data stored in dst.Fee, return on the first match
		} else {
			dst.Fee = nil
			return fmt.Errorf("Failed to unmarshal PatchAccountProduct as Fee: %s", err.Error())
		}
	}

	// check if the discriminator value is 'patch_interest'
	if jsonDict["product_type"] == "patch_interest" {
		// try to unmarshal JSON data into PatchInterest
		err = json.Unmarshal(data, &dst.PatchInterest)
		if err == nil {
			return nil // data stored in dst.PatchInterest, return on the first match
		} else {
			dst.PatchInterest = nil
			return fmt.Errorf("Failed to unmarshal PatchAccountProduct as PatchInterest: %s", err.Error())
		}
	}

	return nil
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src PatchAccountProduct) MarshalJSON() ([]byte, error) {
	if src.Fee != nil {
		return json.Marshal(&src.Fee)
	}

	if src.PatchInterest != nil {
		return json.Marshal(&src.PatchInterest)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *PatchAccountProduct) GetActualInstance() interface{} {
	if obj.Fee != nil {
		return obj.Fee
	}

	if obj.PatchInterest != nil {
		return obj.PatchInterest
	}

	// all schemas are nil
	return nil
}

type NullablePatchAccountProduct struct {
	value *PatchAccountProduct
	isSet bool
}

func (v NullablePatchAccountProduct) Get() *PatchAccountProduct {
	return v.value
}

func (v *NullablePatchAccountProduct) Set(val *PatchAccountProduct) {
	v.value = val
	v.isSet = true
}

func (v NullablePatchAccountProduct) IsSet() bool {
	return v.isSet
}

func (v *NullablePatchAccountProduct) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePatchAccountProduct(val *PatchAccountProduct) *NullablePatchAccountProduct {
	return &NullablePatchAccountProduct{value: val, isSet: true}
}

func (v NullablePatchAccountProduct) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePatchAccountProduct) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
