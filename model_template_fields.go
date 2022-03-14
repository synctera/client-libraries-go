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

// TemplateFields - struct for TemplateFields
type TemplateFields struct {
	TemplateFieldsGeneric   *TemplateFieldsGeneric
	TemplateFieldsOverdraft *TemplateFieldsOverdraft
}

// TemplateFieldsGenericAsTemplateFields is a convenience function that returns TemplateFieldsGeneric wrapped in TemplateFields
func TemplateFieldsGenericAsTemplateFields(v *TemplateFieldsGeneric) TemplateFields {
	return TemplateFields{TemplateFieldsGeneric: v}
}

// TemplateFieldsOverdraftAsTemplateFields is a convenience function that returns TemplateFieldsOverdraft wrapped in TemplateFields
func TemplateFieldsOverdraftAsTemplateFields(v *TemplateFieldsOverdraft) TemplateFields {
	return TemplateFields{TemplateFieldsOverdraft: v}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *TemplateFields) UnmarshalJSON(data []byte) error {
	var err error
	// use discriminator value to speed up the lookup
	var jsonDict map[string]interface{}
	err = json.Unmarshal(data, &jsonDict)
	if err != nil {
		return fmt.Errorf("Failed to unmarshal JSON into map for the discriminator lookup.")
	}

	// check if the discriminator value is 'CHECKING'
	if jsonDict["account_type"] == "CHECKING" {
		// try to unmarshal JSON data into TemplateFieldsGeneric
		err = json.Unmarshal(data, &dst.TemplateFieldsGeneric)
		if err == nil {
			return nil // data stored in dst.TemplateFieldsGeneric, return on the first match
		} else {
			dst.TemplateFieldsGeneric = nil
			return fmt.Errorf("Failed to unmarshal TemplateFields as TemplateFieldsGeneric: %s", err.Error())
		}
	}

	// check if the discriminator value is 'LINE_OF_CREDIT'
	if jsonDict["account_type"] == "LINE_OF_CREDIT" {
		// try to unmarshal JSON data into TemplateFieldsGeneric
		err = json.Unmarshal(data, &dst.TemplateFieldsGeneric)
		if err == nil {
			return nil // data stored in dst.TemplateFieldsGeneric, return on the first match
		} else {
			dst.TemplateFieldsGeneric = nil
			return fmt.Errorf("Failed to unmarshal TemplateFields as TemplateFieldsGeneric: %s", err.Error())
		}
	}

	// check if the discriminator value is 'ODP'
	if jsonDict["account_type"] == "ODP" {
		// try to unmarshal JSON data into TemplateFieldsOverdraft
		err = json.Unmarshal(data, &dst.TemplateFieldsOverdraft)
		if err == nil {
			return nil // data stored in dst.TemplateFieldsOverdraft, return on the first match
		} else {
			dst.TemplateFieldsOverdraft = nil
			return fmt.Errorf("Failed to unmarshal TemplateFields as TemplateFieldsOverdraft: %s", err.Error())
		}
	}

	// check if the discriminator value is 'SAVING'
	if jsonDict["account_type"] == "SAVING" {
		// try to unmarshal JSON data into TemplateFieldsGeneric
		err = json.Unmarshal(data, &dst.TemplateFieldsGeneric)
		if err == nil {
			return nil // data stored in dst.TemplateFieldsGeneric, return on the first match
		} else {
			dst.TemplateFieldsGeneric = nil
			return fmt.Errorf("Failed to unmarshal TemplateFields as TemplateFieldsGeneric: %s", err.Error())
		}
	}

	// check if the discriminator value is 'template_fields_generic'
	if jsonDict["account_type"] == "template_fields_generic" {
		// try to unmarshal JSON data into TemplateFieldsGeneric
		err = json.Unmarshal(data, &dst.TemplateFieldsGeneric)
		if err == nil {
			return nil // data stored in dst.TemplateFieldsGeneric, return on the first match
		} else {
			dst.TemplateFieldsGeneric = nil
			return fmt.Errorf("Failed to unmarshal TemplateFields as TemplateFieldsGeneric: %s", err.Error())
		}
	}

	// check if the discriminator value is 'template_fields_overdraft'
	if jsonDict["account_type"] == "template_fields_overdraft" {
		// try to unmarshal JSON data into TemplateFieldsOverdraft
		err = json.Unmarshal(data, &dst.TemplateFieldsOverdraft)
		if err == nil {
			return nil // data stored in dst.TemplateFieldsOverdraft, return on the first match
		} else {
			dst.TemplateFieldsOverdraft = nil
			return fmt.Errorf("Failed to unmarshal TemplateFields as TemplateFieldsOverdraft: %s", err.Error())
		}
	}

	return nil
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src TemplateFields) MarshalJSON() ([]byte, error) {
	if src.TemplateFieldsGeneric != nil {
		return json.Marshal(&src.TemplateFieldsGeneric)
	}

	if src.TemplateFieldsOverdraft != nil {
		return json.Marshal(&src.TemplateFieldsOverdraft)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *TemplateFields) GetActualInstance() interface{} {
	if obj.TemplateFieldsGeneric != nil {
		return obj.TemplateFieldsGeneric
	}

	if obj.TemplateFieldsOverdraft != nil {
		return obj.TemplateFieldsOverdraft
	}

	// all schemas are nil
	return nil
}

type NullableTemplateFields struct {
	value *TemplateFields
	isSet bool
}

func (v NullableTemplateFields) Get() *TemplateFields {
	return v.value
}

func (v *NullableTemplateFields) Set(val *TemplateFields) {
	v.value = val
	v.isSet = true
}

func (v NullableTemplateFields) IsSet() bool {
	return v.isSet
}

func (v *NullableTemplateFields) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTemplateFields(val *TemplateFields) *NullableTemplateFields {
	return &NullableTemplateFields{value: val, isSet: true}
}

func (v NullableTemplateFields) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTemplateFields) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
