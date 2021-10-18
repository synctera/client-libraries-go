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

// DocumentType Type of document: `LICENSE_FRONT` - Government issued license front, `LICENSE_BACK` - Government issued license back, `ID_CARD_FRONT` - Government issued ID front, `ID_CARD_BACK` - Government issued ID back, `PASSPORT` - Passport, `SELF_PORTRAIT` - Selfie, `PROOF_OF_ADDRESS` - Document to prove address such as a utility bill
type DocumentType string

// List of document_type
const (
	DOCUMENTTYPE_LICENSE_FRONT DocumentType = "LICENSE_FRONT"
	DOCUMENTTYPE_LICENSE_BACK DocumentType = "LICENSE_BACK"
	DOCUMENTTYPE_ID_CARD_FRONT DocumentType = "ID_CARD_FRONT"
	DOCUMENTTYPE_ID_CARD_BACK DocumentType = "ID_CARD_BACK"
	DOCUMENTTYPE_PASSPORT DocumentType = "PASSPORT"
	DOCUMENTTYPE_SELF_PORTRAIT DocumentType = "SELF_PORTRAIT"
	DOCUMENTTYPE_PROOF_OF_ADDRESS DocumentType = "PROOF_OF_ADDRESS"
)

var allowedDocumentTypeEnumValues = []DocumentType{
	"LICENSE_FRONT",
	"LICENSE_BACK",
	"ID_CARD_FRONT",
	"ID_CARD_BACK",
	"PASSPORT",
	"SELF_PORTRAIT",
	"PROOF_OF_ADDRESS",
}

func (v *DocumentType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := DocumentType(value)
	for _, existing := range allowedDocumentTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid DocumentType", value)
}

// NewDocumentTypeFromValue returns a pointer to a valid DocumentType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewDocumentTypeFromValue(v string) (*DocumentType, error) {
	ev := DocumentType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for DocumentType: valid values are %v", v, allowedDocumentTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v DocumentType) IsValid() bool {
	for _, existing := range allowedDocumentTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to document_type value
func (v DocumentType) Ptr() *DocumentType {
	return &v
}

type NullableDocumentType struct {
	value *DocumentType
	isSet bool
}

func (v NullableDocumentType) Get() *DocumentType {
	return v.value
}

func (v *NullableDocumentType) Set(val *DocumentType) {
	v.value = val
	v.isSet = true
}

func (v NullableDocumentType) IsSet() bool {
	return v.isSet
}

func (v *NullableDocumentType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDocumentType(val *DocumentType) *NullableDocumentType {
	return &NullableDocumentType{value: val, isSet: true}
}

func (v NullableDocumentType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDocumentType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

