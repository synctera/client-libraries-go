/*
Synctera API

<h2>Let's build something great.</h2><p>Welcome to the official reference documentation for Synctera APIs. Our APIs are the best way to automate your company's banking needs and are designed to be easy to understand and implement.</p><p>We're continuously growing this library and what you see here is just the start, but if you need something specific or have a question, <a class='text-blue-600' href='https://synctera.com/contact' target='_blank' rel='noreferrer'>contact us</a>.</p> 

API version: 0.5.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"fmt"
)

// AccountRelationshipType Relationship type
type AccountRelationshipType string

// List of account_relationship_type
const (
	ACCOUNTRELATIONSHIPTYPE_ACCOUNT_HOLDER AccountRelationshipType = "ACCOUNT_HOLDER"
	ACCOUNTRELATIONSHIPTYPE_EMPLOYEE AccountRelationshipType = "EMPLOYEE"
	ACCOUNTRELATIONSHIPTYPE_BRANCH AccountRelationshipType = "BRANCH"
	ACCOUNTRELATIONSHIPTYPE_ATM AccountRelationshipType = "ATM"
	ACCOUNTRELATIONSHIPTYPE_OFFICER AccountRelationshipType = "OFFICER"
	ACCOUNTRELATIONSHIPTYPE_PRIMARY_ACCOUNT_HOLDER AccountRelationshipType = "PRIMARY_ACCOUNT_HOLDER"
	ACCOUNTRELATIONSHIPTYPE_JOINT_ACCOUNT_HOLDER AccountRelationshipType = "JOINT_ACCOUNT_HOLDER"
	ACCOUNTRELATIONSHIPTYPE_AUTHORIZED_SIGNER AccountRelationshipType = "AUTHORIZED_SIGNER"
)

var allowedAccountRelationshipTypeEnumValues = []AccountRelationshipType{
	"ACCOUNT_HOLDER",
	"EMPLOYEE",
	"BRANCH",
	"ATM",
	"OFFICER",
	"PRIMARY_ACCOUNT_HOLDER",
	"JOINT_ACCOUNT_HOLDER",
	"AUTHORIZED_SIGNER",
}

func (v *AccountRelationshipType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := AccountRelationshipType(value)
	for _, existing := range allowedAccountRelationshipTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid AccountRelationshipType", value)
}

// NewAccountRelationshipTypeFromValue returns a pointer to a valid AccountRelationshipType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAccountRelationshipTypeFromValue(v string) (*AccountRelationshipType, error) {
	ev := AccountRelationshipType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AccountRelationshipType: valid values are %v", v, allowedAccountRelationshipTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AccountRelationshipType) IsValid() bool {
	for _, existing := range allowedAccountRelationshipTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to account_relationship_type value
func (v AccountRelationshipType) Ptr() *AccountRelationshipType {
	return &v
}

type NullableAccountRelationshipType struct {
	value *AccountRelationshipType
	isSet bool
}

func (v NullableAccountRelationshipType) Get() *AccountRelationshipType {
	return v.value
}

func (v *NullableAccountRelationshipType) Set(val *AccountRelationshipType) {
	v.value = val
	v.isSet = true
}

func (v NullableAccountRelationshipType) IsSet() bool {
	return v.isSet
}

func (v *NullableAccountRelationshipType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccountRelationshipType(val *AccountRelationshipType) *NullableAccountRelationshipType {
	return &NullableAccountRelationshipType{value: val, isSet: true}
}

func (v NullableAccountRelationshipType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccountRelationshipType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
