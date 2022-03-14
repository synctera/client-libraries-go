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

// CustomerKycStatus Customer's KYC status
type CustomerKycStatus string

// List of customer_kyc_status
const (
	CUSTOMERKYCSTATUS_UNVERIFIED       CustomerKycStatus = "UNVERIFIED"
	CUSTOMERKYCSTATUS_REVIEW           CustomerKycStatus = "REVIEW"
	CUSTOMERKYCSTATUS_PROVIDER_FAILURE CustomerKycStatus = "PROVIDER_FAILURE"
	CUSTOMERKYCSTATUS_ACCEPTED         CustomerKycStatus = "ACCEPTED"
	CUSTOMERKYCSTATUS_REJECTED         CustomerKycStatus = "REJECTED"
)

// All allowed values of CustomerKycStatus enum
var AllowedCustomerKycStatusEnumValues = []CustomerKycStatus{
	"UNVERIFIED",
	"REVIEW",
	"PROVIDER_FAILURE",
	"ACCEPTED",
	"REJECTED",
}

func (v *CustomerKycStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := CustomerKycStatus(value)
	for _, existing := range AllowedCustomerKycStatusEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid CustomerKycStatus", value)
}

// NewCustomerKycStatusFromValue returns a pointer to a valid CustomerKycStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewCustomerKycStatusFromValue(v string) (*CustomerKycStatus, error) {
	ev := CustomerKycStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for CustomerKycStatus: valid values are %v", v, AllowedCustomerKycStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v CustomerKycStatus) IsValid() bool {
	for _, existing := range AllowedCustomerKycStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to customer_kyc_status value
func (v CustomerKycStatus) Ptr() *CustomerKycStatus {
	return &v
}

type NullableCustomerKycStatus struct {
	value *CustomerKycStatus
	isSet bool
}

func (v NullableCustomerKycStatus) Get() *CustomerKycStatus {
	return v.value
}

func (v *NullableCustomerKycStatus) Set(val *CustomerKycStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableCustomerKycStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableCustomerKycStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCustomerKycStatus(val *CustomerKycStatus) *NullableCustomerKycStatus {
	return &NullableCustomerKycStatus{value: val, isSet: true}
}

func (v NullableCustomerKycStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCustomerKycStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
