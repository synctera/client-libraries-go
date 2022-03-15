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

// ProviderType Name of the organization that provided this data.
type ProviderType string

// List of provider_type
const (
	PROVIDERTYPE_IDOLOGY ProviderType = "IDOLOGY"
	PROVIDERTYPE_SOCURE  ProviderType = "SOCURE"
)

// All allowed values of ProviderType enum
var AllowedProviderTypeEnumValues = []ProviderType{
	"IDOLOGY",
	"SOCURE",
}

func (v *ProviderType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ProviderType(value)
	for _, existing := range AllowedProviderTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ProviderType", value)
}

// NewProviderTypeFromValue returns a pointer to a valid ProviderType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewProviderTypeFromValue(v string) (*ProviderType, error) {
	ev := ProviderType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ProviderType: valid values are %v", v, AllowedProviderTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ProviderType) IsValid() bool {
	for _, existing := range AllowedProviderTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to provider_type value
func (v ProviderType) Ptr() *ProviderType {
	return &v
}

type NullableProviderType struct {
	value *ProviderType
	isSet bool
}

func (v NullableProviderType) Get() *ProviderType {
	return v.value
}

func (v *NullableProviderType) Set(val *ProviderType) {
	v.value = val
	v.isSet = true
}

func (v NullableProviderType) IsSet() bool {
	return v.isSet
}

func (v *NullableProviderType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProviderType(val *ProviderType) *NullableProviderType {
	return &NullableProviderType{value: val, isSet: true}
}

func (v NullableProviderType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProviderType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
