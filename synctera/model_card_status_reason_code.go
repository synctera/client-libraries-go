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

// CardStatusReasonCode The reason for the card status  Code | Description --- | --- NEW | Card activated REQ | Requested by you INA | Dormant UNK | Invalid shipping address NEG | Negative account balance REV | Account under review SUS | Suspicious activity OUT | Activity outside program parameters FRD | Confirmed fraud MAT | Matched with an OFAC list LOS | Card reported lost CLO | Card was cloned COM | Account or card was compromised TMP | Awaiting customer confirmation PRC | Initiated by Processor ISS | Initiated by Issuer EXP | Card expired KYC | Failed KYC INF | Information was validated ACT | Account activity was validated AUX | Initiated by a third party PIN | PIN try limit reached STO | Card reported stolen ADD | Address issue NAM | Name issue SSN | SSN issue DOB | DOB issue EML | Email issue PHO | Phone issue FUL | Account/fulfillment mismatch OTH | Other 
type CardStatusReasonCode string

// List of card_status_reason_code
const (
	CARDSTATUSREASONCODE_NEW CardStatusReasonCode = "NEW"
	CARDSTATUSREASONCODE_REQ CardStatusReasonCode = "REQ"
	CARDSTATUSREASONCODE_INA CardStatusReasonCode = "INA"
	CARDSTATUSREASONCODE_UNK CardStatusReasonCode = "UNK"
	CARDSTATUSREASONCODE_NEG CardStatusReasonCode = "NEG"
	CARDSTATUSREASONCODE_REV CardStatusReasonCode = "REV"
	CARDSTATUSREASONCODE_SUS CardStatusReasonCode = "SUS"
	CARDSTATUSREASONCODE_OUT CardStatusReasonCode = "OUT"
	CARDSTATUSREASONCODE_FRD CardStatusReasonCode = "FRD"
	CARDSTATUSREASONCODE_MAT CardStatusReasonCode = "MAT"
	CARDSTATUSREASONCODE_LOS CardStatusReasonCode = "LOS"
	CARDSTATUSREASONCODE_CLO CardStatusReasonCode = "CLO"
	CARDSTATUSREASONCODE_COM CardStatusReasonCode = "COM"
	CARDSTATUSREASONCODE_TMP CardStatusReasonCode = "TMP"
	CARDSTATUSREASONCODE_PRC CardStatusReasonCode = "PRC"
	CARDSTATUSREASONCODE_ISS CardStatusReasonCode = "ISS"
	CARDSTATUSREASONCODE_EXP CardStatusReasonCode = "EXP"
	CARDSTATUSREASONCODE_KYC CardStatusReasonCode = "KYC"
	CARDSTATUSREASONCODE_INF CardStatusReasonCode = "INF"
	CARDSTATUSREASONCODE_ACT CardStatusReasonCode = "ACT"
	CARDSTATUSREASONCODE_AUX CardStatusReasonCode = "AUX"
	CARDSTATUSREASONCODE_PIN CardStatusReasonCode = "PIN"
	CARDSTATUSREASONCODE_STO CardStatusReasonCode = "STO"
	CARDSTATUSREASONCODE_ADD CardStatusReasonCode = "ADD"
	CARDSTATUSREASONCODE_NAM CardStatusReasonCode = "NAM"
	CARDSTATUSREASONCODE_SSN CardStatusReasonCode = "SSN"
	CARDSTATUSREASONCODE_DOB CardStatusReasonCode = "DOB"
	CARDSTATUSREASONCODE_EML CardStatusReasonCode = "EML"
	CARDSTATUSREASONCODE_PHO CardStatusReasonCode = "PHO"
	CARDSTATUSREASONCODE_FUL CardStatusReasonCode = "FUL"
	CARDSTATUSREASONCODE_OTH CardStatusReasonCode = "OTH"
)

var allowedCardStatusReasonCodeEnumValues = []CardStatusReasonCode{
	"NEW",
	"REQ",
	"INA",
	"UNK",
	"NEG",
	"REV",
	"SUS",
	"OUT",
	"FRD",
	"MAT",
	"LOS",
	"CLO",
	"COM",
	"TMP",
	"PRC",
	"ISS",
	"EXP",
	"KYC",
	"INF",
	"ACT",
	"AUX",
	"PIN",
	"STO",
	"ADD",
	"NAM",
	"SSN",
	"DOB",
	"EML",
	"PHO",
	"FUL",
	"OTH",
}

func (v *CardStatusReasonCode) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := CardStatusReasonCode(value)
	for _, existing := range allowedCardStatusReasonCodeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid CardStatusReasonCode", value)
}

// NewCardStatusReasonCodeFromValue returns a pointer to a valid CardStatusReasonCode
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewCardStatusReasonCodeFromValue(v string) (*CardStatusReasonCode, error) {
	ev := CardStatusReasonCode(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for CardStatusReasonCode: valid values are %v", v, allowedCardStatusReasonCodeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v CardStatusReasonCode) IsValid() bool {
	for _, existing := range allowedCardStatusReasonCodeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to card_status_reason_code value
func (v CardStatusReasonCode) Ptr() *CardStatusReasonCode {
	return &v
}

type NullableCardStatusReasonCode struct {
	value *CardStatusReasonCode
	isSet bool
}

func (v NullableCardStatusReasonCode) Get() *CardStatusReasonCode {
	return v.value
}

func (v *NullableCardStatusReasonCode) Set(val *CardStatusReasonCode) {
	v.value = val
	v.isSet = true
}

func (v NullableCardStatusReasonCode) IsSet() bool {
	return v.isSet
}

func (v *NullableCardStatusReasonCode) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCardStatusReasonCode(val *CardStatusReasonCode) *NullableCardStatusReasonCode {
	return &NullableCardStatusReasonCode{value: val, isSet: true}
}

func (v NullableCardStatusReasonCode) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCardStatusReasonCode) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

