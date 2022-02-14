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
	"fmt"
)

// AccountAccessStatus Access status for account. Default ACTIVE
type AccountAccessStatus string

// List of account_access_status
const (
	ACCOUNTACCESSSTATUS_ACTIVE AccountAccessStatus = "ACTIVE"
	ACCOUNTACCESSSTATUS_FROZEN AccountAccessStatus = "FROZEN"
)

var allowedAccountAccessStatusEnumValues = []AccountAccessStatus{
	"ACTIVE",
	"FROZEN",
}

func (v *AccountAccessStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := AccountAccessStatus(value)
	for _, existing := range allowedAccountAccessStatusEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid AccountAccessStatus", value)
}

// NewAccountAccessStatusFromValue returns a pointer to a valid AccountAccessStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAccountAccessStatusFromValue(v string) (*AccountAccessStatus, error) {
	ev := AccountAccessStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AccountAccessStatus: valid values are %v", v, allowedAccountAccessStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AccountAccessStatus) IsValid() bool {
	for _, existing := range allowedAccountAccessStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to account_access_status value
func (v AccountAccessStatus) Ptr() *AccountAccessStatus {
	return &v
}

type NullableAccountAccessStatus struct {
	value *AccountAccessStatus
	isSet bool
}

func (v NullableAccountAccessStatus) Get() *AccountAccessStatus {
	return v.value
}

func (v *NullableAccountAccessStatus) Set(val *AccountAccessStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableAccountAccessStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableAccountAccessStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccountAccessStatus(val *AccountAccessStatus) *NullableAccountAccessStatus {
	return &NullableAccountAccessStatus{value: val, isSet: true}
}

func (v NullableAccountAccessStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccountAccessStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}