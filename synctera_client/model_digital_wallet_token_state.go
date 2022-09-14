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

// DigitalWalletTokenState Current status of the Digital Wallet Token
type DigitalWalletTokenState string

// List of digital_wallet_token_state
const (
	DIGITALWALLETTOKENSTATE_REQUESTED        DigitalWalletTokenState = "REQUESTED"
	DIGITALWALLETTOKENSTATE_REQUEST_DECLINED DigitalWalletTokenState = "REQUEST_DECLINED"
	DIGITALWALLETTOKENSTATE_ACTIVE           DigitalWalletTokenState = "ACTIVE"
	DIGITALWALLETTOKENSTATE_SUSPENDED        DigitalWalletTokenState = "SUSPENDED"
	DIGITALWALLETTOKENSTATE_TERMINATED       DigitalWalletTokenState = "TERMINATED"
)

// All allowed values of DigitalWalletTokenState enum
var AllowedDigitalWalletTokenStateEnumValues = []DigitalWalletTokenState{
	"REQUESTED",
	"REQUEST_DECLINED",
	"ACTIVE",
	"SUSPENDED",
	"TERMINATED",
}

func (v *DigitalWalletTokenState) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := DigitalWalletTokenState(value)
	for _, existing := range AllowedDigitalWalletTokenStateEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid DigitalWalletTokenState", value)
}

// NewDigitalWalletTokenStateFromValue returns a pointer to a valid DigitalWalletTokenState
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewDigitalWalletTokenStateFromValue(v string) (*DigitalWalletTokenState, error) {
	ev := DigitalWalletTokenState(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for DigitalWalletTokenState: valid values are %v", v, AllowedDigitalWalletTokenStateEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v DigitalWalletTokenState) IsValid() bool {
	for _, existing := range AllowedDigitalWalletTokenStateEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to digital_wallet_token_state value
func (v DigitalWalletTokenState) Ptr() *DigitalWalletTokenState {
	return &v
}

type NullableDigitalWalletTokenState struct {
	value *DigitalWalletTokenState
	isSet bool
}

func (v NullableDigitalWalletTokenState) Get() *DigitalWalletTokenState {
	return v.value
}

func (v *NullableDigitalWalletTokenState) Set(val *DigitalWalletTokenState) {
	v.value = val
	v.isSet = true
}

func (v NullableDigitalWalletTokenState) IsSet() bool {
	return v.isSet
}

func (v *NullableDigitalWalletTokenState) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDigitalWalletTokenState(val *DigitalWalletTokenState) *NullableDigitalWalletTokenState {
	return &NullableDigitalWalletTokenState{value: val, isSet: true}
}

func (v NullableDigitalWalletTokenState) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDigitalWalletTokenState) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
