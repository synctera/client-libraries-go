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
)

// ReconciliationInput Input data for a reconciliation
type ReconciliationInput struct {
	// Base64url encoded image
	ByteData string `json:"byte_data"`
	// Filename of the data to be reconciled
	FileName string `json:"file_name"`
}

// NewReconciliationInput instantiates a new ReconciliationInput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReconciliationInput(byteData string, fileName string) *ReconciliationInput {
	this := ReconciliationInput{}
	this.ByteData = byteData
	this.FileName = fileName
	return &this
}

// NewReconciliationInputWithDefaults instantiates a new ReconciliationInput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReconciliationInputWithDefaults() *ReconciliationInput {
	this := ReconciliationInput{}
	return &this
}

// GetByteData returns the ByteData field value
func (o *ReconciliationInput) GetByteData() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ByteData
}

// GetByteDataOk returns a tuple with the ByteData field value
// and a boolean to check if the value has been set.
func (o *ReconciliationInput) GetByteDataOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ByteData, true
}

// SetByteData sets field value
func (o *ReconciliationInput) SetByteData(v string) {
	o.ByteData = v
}

// GetFileName returns the FileName field value
func (o *ReconciliationInput) GetFileName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FileName
}

// GetFileNameOk returns a tuple with the FileName field value
// and a boolean to check if the value has been set.
func (o *ReconciliationInput) GetFileNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FileName, true
}

// SetFileName sets field value
func (o *ReconciliationInput) SetFileName(v string) {
	o.FileName = v
}

func (o ReconciliationInput) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["byte_data"] = o.ByteData
	}
	if true {
		toSerialize["file_name"] = o.FileName
	}
	return json.Marshal(toSerialize)
}

type NullableReconciliationInput struct {
	value *ReconciliationInput
	isSet bool
}

func (v NullableReconciliationInput) Get() *ReconciliationInput {
	return v.value
}

func (v *NullableReconciliationInput) Set(val *ReconciliationInput) {
	v.value = val
	v.isSet = true
}

func (v NullableReconciliationInput) IsSet() bool {
	return v.isSet
}

func (v *NullableReconciliationInput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReconciliationInput(val *ReconciliationInput) *NullableReconciliationInput {
	return &NullableReconciliationInput{value: val, isSet: true}
}

func (v NullableReconciliationInput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReconciliationInput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
