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

// Encryption the model 'Encryption'
type Encryption string

// List of encryption
const (
	ENCRYPTION_REQUIRED Encryption = "REQUIRED"
	ENCRYPTION_NOT_REQUIRED Encryption = "NOT_REQUIRED"
)

// All allowed values of Encryption enum
var AllowedEncryptionEnumValues = []Encryption{
	"REQUIRED",
	"NOT_REQUIRED",
}

func (v *Encryption) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := Encryption(value)
	for _, existing := range AllowedEncryptionEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid Encryption", value)
}

// NewEncryptionFromValue returns a pointer to a valid Encryption
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEncryptionFromValue(v string) (*Encryption, error) {
	ev := Encryption(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for Encryption: valid values are %v", v, AllowedEncryptionEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v Encryption) IsValid() bool {
	for _, existing := range AllowedEncryptionEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to encryption value
func (v Encryption) Ptr() *Encryption {
	return &v
}

type NullableEncryption struct {
	value *Encryption
	isSet bool
}

func (v NullableEncryption) Get() *Encryption {
	return v.value
}

func (v *NullableEncryption) Set(val *Encryption) {
	v.value = val
	v.isSet = true
}

func (v NullableEncryption) IsSet() bool {
	return v.isSet
}

func (v *NullableEncryption) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEncryption(val *Encryption) *NullableEncryption {
	return &NullableEncryption{value: val, isSet: true}
}

func (v NullableEncryption) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEncryption) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
