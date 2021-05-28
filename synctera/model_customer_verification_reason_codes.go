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

// CustomerVerificationReasonCodes Specific verification issue.  Based on rules defined with bank this may
type CustomerVerificationReasonCodes string

// List of customer_verification_reason_codes
const (
	CUSTOMERVERIFICATIONREASONCODES_NAME CustomerVerificationReasonCodes = "NAME"
	CUSTOMERVERIFICATIONREASONCODES_ADDRESS CustomerVerificationReasonCodes = "ADDRESS"
	CUSTOMERVERIFICATIONREASONCODES_SSN_ADDRESS_MISMATCH CustomerVerificationReasonCodes = "SSN_ADDRESS_MISMATCH"
	CUSTOMERVERIFICATIONREASONCODES_NON_US_CITIZEN CustomerVerificationReasonCodes = "NON_US_CITIZEN"
	CUSTOMERVERIFICATIONREASONCODES_IDENTITY_TOO_RECENT CustomerVerificationReasonCodes = "IDENTITY_TOO_RECENT"
	CUSTOMERVERIFICATIONREASONCODES_GLOBAL_WATCHLIST CustomerVerificationReasonCodes = "GLOBAL_WATCHLIST"
	CUSTOMERVERIFICATIONREASONCODES_EMAIL_FRAUD_LIST CustomerVerificationReasonCodes = "EMAIL_FRAUD_LIST"
	CUSTOMERVERIFICATIONREASONCODES_SSN_FRAUD_LIST CustomerVerificationReasonCodes = "SSN_FRAUD_LIST"
	CUSTOMERVERIFICATIONREASONCODES_PHONE_FRAUD_LIST CustomerVerificationReasonCodes = "PHONE_FRAUD_LIST"
	CUSTOMERVERIFICATIONREASONCODES_DOB CustomerVerificationReasonCodes = "DOB"
	CUSTOMERVERIFICATIONREASONCODES_SSN_DECEASED CustomerVerificationReasonCodes = "SSN_DECEASED"
	CUSTOMERVERIFICATIONREASONCODES_SSN CustomerVerificationReasonCodes = "SSN"
	CUSTOMERVERIFICATIONREASONCODES_SSN_NAME_MISMATCH CustomerVerificationReasonCodes = "SSN_NAME_MISMATCH"
	CUSTOMERVERIFICATIONREASONCODES_ACTIVE_DUTY CustomerVerificationReasonCodes = "ACTIVE_DUTY"
	CUSTOMERVERIFICATIONREASONCODES_FRAUD_VICTIM CustomerVerificationReasonCodes = "FRAUD_VICTIM"
)

var allowedCustomerVerificationReasonCodesEnumValues = []CustomerVerificationReasonCodes{
	"NAME",
	"ADDRESS",
	"SSN_ADDRESS_MISMATCH",
	"NON_US_CITIZEN",
	"IDENTITY_TOO_RECENT",
	"GLOBAL_WATCHLIST",
	"EMAIL_FRAUD_LIST",
	"SSN_FRAUD_LIST",
	"PHONE_FRAUD_LIST",
	"DOB",
	"SSN_DECEASED",
	"SSN",
	"SSN_NAME_MISMATCH",
	"ACTIVE_DUTY",
	"FRAUD_VICTIM",
}

func (v *CustomerVerificationReasonCodes) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := CustomerVerificationReasonCodes(value)
	for _, existing := range allowedCustomerVerificationReasonCodesEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid CustomerVerificationReasonCodes", value)
}

// NewCustomerVerificationReasonCodesFromValue returns a pointer to a valid CustomerVerificationReasonCodes
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewCustomerVerificationReasonCodesFromValue(v string) (*CustomerVerificationReasonCodes, error) {
	ev := CustomerVerificationReasonCodes(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for CustomerVerificationReasonCodes: valid values are %v", v, allowedCustomerVerificationReasonCodesEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v CustomerVerificationReasonCodes) IsValid() bool {
	for _, existing := range allowedCustomerVerificationReasonCodesEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to customer_verification_reason_codes value
func (v CustomerVerificationReasonCodes) Ptr() *CustomerVerificationReasonCodes {
	return &v
}

type NullableCustomerVerificationReasonCodes struct {
	value *CustomerVerificationReasonCodes
	isSet bool
}

func (v NullableCustomerVerificationReasonCodes) Get() *CustomerVerificationReasonCodes {
	return v.value
}

func (v *NullableCustomerVerificationReasonCodes) Set(val *CustomerVerificationReasonCodes) {
	v.value = val
	v.isSet = true
}

func (v NullableCustomerVerificationReasonCodes) IsSet() bool {
	return v.isSet
}

func (v *NullableCustomerVerificationReasonCodes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCustomerVerificationReasonCodes(val *CustomerVerificationReasonCodes) *NullableCustomerVerificationReasonCodes {
	return &NullableCustomerVerificationReasonCodes{value: val, isSet: true}
}

func (v NullableCustomerVerificationReasonCodes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCustomerVerificationReasonCodes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

