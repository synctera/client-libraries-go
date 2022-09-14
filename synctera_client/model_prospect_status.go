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

// ProspectStatus Status of the prospect
type ProspectStatus string

// List of prospect_status
const (
	PROSPECTSTATUS_CREATED   ProspectStatus = "CREATED"
	PROSPECTSTATUS_VERIFIED  ProspectStatus = "VERIFIED"
	PROSPECTSTATUS_ADMITTED  ProspectStatus = "ADMITTED"
	PROSPECTSTATUS_WITHDRAWN ProspectStatus = "WITHDRAWN"
)

// All allowed values of ProspectStatus enum
var AllowedProspectStatusEnumValues = []ProspectStatus{
	"CREATED",
	"VERIFIED",
	"ADMITTED",
	"WITHDRAWN",
}

func (v *ProspectStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ProspectStatus(value)
	for _, existing := range AllowedProspectStatusEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ProspectStatus", value)
}

// NewProspectStatusFromValue returns a pointer to a valid ProspectStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewProspectStatusFromValue(v string) (*ProspectStatus, error) {
	ev := ProspectStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ProspectStatus: valid values are %v", v, AllowedProspectStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ProspectStatus) IsValid() bool {
	for _, existing := range AllowedProspectStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to prospect_status value
func (v ProspectStatus) Ptr() *ProspectStatus {
	return &v
}

type NullableProspectStatus struct {
	value *ProspectStatus
	isSet bool
}

func (v NullableProspectStatus) Get() *ProspectStatus {
	return v.value
}

func (v *NullableProspectStatus) Set(val *ProspectStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableProspectStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableProspectStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProspectStatus(val *ProspectStatus) *NullableProspectStatus {
	return &NullableProspectStatus{value: val, isSet: true}
}

func (v NullableProspectStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProspectStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
