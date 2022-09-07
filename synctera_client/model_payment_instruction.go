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

// PaymentInstruction - struct for PaymentInstruction
type PaymentInstruction struct {
	AchInstruction *AchInstruction
	InternalTransferInstruction *InternalTransferInstruction
}

// AchInstructionAsPaymentInstruction is a convenience function that returns AchInstruction wrapped in PaymentInstruction
func AchInstructionAsPaymentInstruction(v *AchInstruction) PaymentInstruction {
	return PaymentInstruction{
		AchInstruction: v,
	}
}

// InternalTransferInstructionAsPaymentInstruction is a convenience function that returns InternalTransferInstruction wrapped in PaymentInstruction
func InternalTransferInstructionAsPaymentInstruction(v *InternalTransferInstruction) PaymentInstruction {
	return PaymentInstruction{
		InternalTransferInstruction: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *PaymentInstruction) UnmarshalJSON(data []byte) error {
	var err error
	// use discriminator value to speed up the lookup
	var jsonDict map[string]interface{}
	err = newStrictDecoder(data).Decode(&jsonDict)
	if err != nil {
		return fmt.Errorf("Failed to unmarshal JSON into map for the discriminator lookup.")
	}

	// check if the discriminator value is 'ACH'
	if jsonDict["type"] == "ACH" {
		// try to unmarshal JSON data into AchInstruction
		err = json.Unmarshal(data, &dst.AchInstruction)
		if err == nil {
			return nil // data stored in dst.AchInstruction, return on the first match
		} else {
			dst.AchInstruction = nil
			return fmt.Errorf("Failed to unmarshal PaymentInstruction as AchInstruction: %s", err.Error())
		}
	}

	// check if the discriminator value is 'INTERNAL_TRANSFER'
	if jsonDict["type"] == "INTERNAL_TRANSFER" {
		// try to unmarshal JSON data into InternalTransferInstruction
		err = json.Unmarshal(data, &dst.InternalTransferInstruction)
		if err == nil {
			return nil // data stored in dst.InternalTransferInstruction, return on the first match
		} else {
			dst.InternalTransferInstruction = nil
			return fmt.Errorf("Failed to unmarshal PaymentInstruction as InternalTransferInstruction: %s", err.Error())
		}
	}

	// check if the discriminator value is 'ach_instruction'
	if jsonDict["type"] == "ach_instruction" {
		// try to unmarshal JSON data into AchInstruction
		err = json.Unmarshal(data, &dst.AchInstruction)
		if err == nil {
			return nil // data stored in dst.AchInstruction, return on the first match
		} else {
			dst.AchInstruction = nil
			return fmt.Errorf("Failed to unmarshal PaymentInstruction as AchInstruction: %s", err.Error())
		}
	}

	// check if the discriminator value is 'internal_transfer_instruction'
	if jsonDict["type"] == "internal_transfer_instruction" {
		// try to unmarshal JSON data into InternalTransferInstruction
		err = json.Unmarshal(data, &dst.InternalTransferInstruction)
		if err == nil {
			return nil // data stored in dst.InternalTransferInstruction, return on the first match
		} else {
			dst.InternalTransferInstruction = nil
			return fmt.Errorf("Failed to unmarshal PaymentInstruction as InternalTransferInstruction: %s", err.Error())
		}
	}

	return nil
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src PaymentInstruction) MarshalJSON() ([]byte, error) {
	if src.AchInstruction != nil {
		return json.Marshal(&src.AchInstruction)
	}

	if src.InternalTransferInstruction != nil {
		return json.Marshal(&src.InternalTransferInstruction)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *PaymentInstruction) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.AchInstruction != nil {
		return obj.AchInstruction
	}

	if obj.InternalTransferInstruction != nil {
		return obj.InternalTransferInstruction
	}

	// all schemas are nil
	return nil
}

type NullablePaymentInstruction struct {
	value *PaymentInstruction
	isSet bool
}

func (v NullablePaymentInstruction) Get() *PaymentInstruction {
	return v.value
}

func (v *NullablePaymentInstruction) Set(val *PaymentInstruction) {
	v.value = val
	v.isSet = true
}

func (v NullablePaymentInstruction) IsSet() bool {
	return v.isSet
}

func (v *NullablePaymentInstruction) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaymentInstruction(val *PaymentInstruction) *NullablePaymentInstruction {
	return &NullablePaymentInstruction{value: val, isSet: true}
}

func (v NullablePaymentInstruction) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaymentInstruction) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


