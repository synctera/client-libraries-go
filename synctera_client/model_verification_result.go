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

// VerificationResult The determination of this verification. One of the following: * `PENDING` – verification is in progress for this customer. * `PROVISIONAL` – partially verified or verified with restrictions. * `ACCEPTED` – the customer has been verified. * `REVIEW` – verification has run and issues have been identified and require review. * `VENDOR_ERROR` – verification did not successfully run due to an unexpected error or failure. * `REJECTED` – the customer was rejected and should not be allowed to take certain actions e.g., open an account.
type VerificationResult string

// List of verification_result
const (
	VERIFICATIONRESULT_PENDING      VerificationResult = "PENDING"
	VERIFICATIONRESULT_PROVISIONAL  VerificationResult = "PROVISIONAL"
	VERIFICATIONRESULT_ACCEPTED     VerificationResult = "ACCEPTED"
	VERIFICATIONRESULT_REVIEW       VerificationResult = "REVIEW"
	VERIFICATIONRESULT_VENDOR_ERROR VerificationResult = "VENDOR_ERROR"
	VERIFICATIONRESULT_REJECTED     VerificationResult = "REJECTED"
)

// All allowed values of VerificationResult enum
var AllowedVerificationResultEnumValues = []VerificationResult{
	"PENDING",
	"PROVISIONAL",
	"ACCEPTED",
	"REVIEW",
	"VENDOR_ERROR",
	"REJECTED",
}

func (v *VerificationResult) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := VerificationResult(value)
	for _, existing := range AllowedVerificationResultEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid VerificationResult", value)
}

// NewVerificationResultFromValue returns a pointer to a valid VerificationResult
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewVerificationResultFromValue(v string) (*VerificationResult, error) {
	ev := VerificationResult(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for VerificationResult: valid values are %v", v, AllowedVerificationResultEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v VerificationResult) IsValid() bool {
	for _, existing := range AllowedVerificationResultEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to verification_result value
func (v VerificationResult) Ptr() *VerificationResult {
	return &v
}

type NullableVerificationResult struct {
	value *VerificationResult
	isSet bool
}

func (v NullableVerificationResult) Get() *VerificationResult {
	return v.value
}

func (v *NullableVerificationResult) Set(val *VerificationResult) {
	v.value = val
	v.isSet = true
}

func (v NullableVerificationResult) IsSet() bool {
	return v.isSet
}

func (v *NullableVerificationResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVerificationResult(val *VerificationResult) *NullableVerificationResult {
	return &NullableVerificationResult{value: val, isSet: true}
}

func (v NullableVerificationResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVerificationResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
