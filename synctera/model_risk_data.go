/*
Synctera API

<h2>Let's build something great.</h2><p>Welcome to the official reference documentation for Synctera APIs. Our APIs are the best way to automate your company's banking needs and are designed to be easy to understand and implement.</p><p>We're continuously growing this library and what you see here is just the start, but if you need something specific or have a question, <a class='text-blue-600' href='https://synctera.com/contact' target='_blank' rel='noreferrer'>contact us</a>.</p> 

API version: 0.5.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// RiskData struct for RiskData
type RiskData struct {
	// Client IP
	ClientIp string `json:"client_ip"`
}

// NewRiskData instantiates a new RiskData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRiskData(clientIp string) *RiskData {
	this := RiskData{}
	this.ClientIp = clientIp
	return &this
}

// NewRiskDataWithDefaults instantiates a new RiskData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRiskDataWithDefaults() *RiskData {
	this := RiskData{}
	return &this
}

// GetClientIp returns the ClientIp field value
func (o *RiskData) GetClientIp() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClientIp
}

// GetClientIpOk returns a tuple with the ClientIp field value
// and a boolean to check if the value has been set.
func (o *RiskData) GetClientIpOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ClientIp, true
}

// SetClientIp sets field value
func (o *RiskData) SetClientIp(v string) {
	o.ClientIp = v
}

func (o RiskData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["client_ip"] = o.ClientIp
	}
	return json.Marshal(toSerialize)
}

type NullableRiskData struct {
	value *RiskData
	isSet bool
}

func (v NullableRiskData) Get() *RiskData {
	return v.value
}

func (v *NullableRiskData) Set(val *RiskData) {
	v.value = val
	v.isSet = true
}

func (v NullableRiskData) IsSet() bool {
	return v.isSet
}

func (v *NullableRiskData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRiskData(val *RiskData) *NullableRiskData {
	return &NullableRiskData{value: val, isSet: true}
}

func (v NullableRiskData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRiskData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

