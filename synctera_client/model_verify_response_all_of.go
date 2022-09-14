/*
Synctera API

<h2>Let's build something great.</h2><p>Welcome to the official reference documentation for Synctera APIs. Our APIs are the best way to automate your company's banking needs and are designed to be easy to understand and implement.</p><p>We're continuously growing this library and what you see here is just the start, but if you need something specific or have a question, <a class='text-blue-600' href='https://synctera.com/contact' target='_blank' rel='noreferrer'>contact us</a>.</p>

API version: 0.20.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package synctera_client

import (
	"encoding/json"
)

// VerifyResponseAllOf struct for VerifyResponseAllOf
type VerifyResponseAllOf struct {
	VerificationStatus VerificationStatus `json:"verification_status"`
	// Array of verification results.
	Verifications []Verification `json:"verifications"`
}

// NewVerifyResponseAllOf instantiates a new VerifyResponseAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVerifyResponseAllOf(verificationStatus VerificationStatus, verifications []Verification) *VerifyResponseAllOf {
	this := VerifyResponseAllOf{}
	this.VerificationStatus = verificationStatus
	this.Verifications = verifications
	return &this
}

// NewVerifyResponseAllOfWithDefaults instantiates a new VerifyResponseAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVerifyResponseAllOfWithDefaults() *VerifyResponseAllOf {
	this := VerifyResponseAllOf{}
	return &this
}

// GetVerificationStatus returns the VerificationStatus field value
func (o *VerifyResponseAllOf) GetVerificationStatus() VerificationStatus {
	if o == nil {
		var ret VerificationStatus
		return ret
	}

	return o.VerificationStatus
}

// GetVerificationStatusOk returns a tuple with the VerificationStatus field value
// and a boolean to check if the value has been set.
func (o *VerifyResponseAllOf) GetVerificationStatusOk() (*VerificationStatus, bool) {
	if o == nil {
		return nil, false
	}
	return &o.VerificationStatus, true
}

// SetVerificationStatus sets field value
func (o *VerifyResponseAllOf) SetVerificationStatus(v VerificationStatus) {
	o.VerificationStatus = v
}

// GetVerifications returns the Verifications field value
func (o *VerifyResponseAllOf) GetVerifications() []Verification {
	if o == nil {
		var ret []Verification
		return ret
	}

	return o.Verifications
}

// GetVerificationsOk returns a tuple with the Verifications field value
// and a boolean to check if the value has been set.
func (o *VerifyResponseAllOf) GetVerificationsOk() ([]Verification, bool) {
	if o == nil {
		return nil, false
	}
	return o.Verifications, true
}

// SetVerifications sets field value
func (o *VerifyResponseAllOf) SetVerifications(v []Verification) {
	o.Verifications = v
}

func (o VerifyResponseAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["verification_status"] = o.VerificationStatus
	}
	if true {
		toSerialize["verifications"] = o.Verifications
	}
	return json.Marshal(toSerialize)
}

type NullableVerifyResponseAllOf struct {
	value *VerifyResponseAllOf
	isSet bool
}

func (v NullableVerifyResponseAllOf) Get() *VerifyResponseAllOf {
	return v.value
}

func (v *NullableVerifyResponseAllOf) Set(val *VerifyResponseAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableVerifyResponseAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableVerifyResponseAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVerifyResponseAllOf(val *VerifyResponseAllOf) *NullableVerifyResponseAllOf {
	return &NullableVerifyResponseAllOf{value: val, isSet: true}
}

func (v NullableVerifyResponseAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVerifyResponseAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
