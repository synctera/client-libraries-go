/*
Synctera API

<h2>Let's build something great.</h2><p>Welcome to the official reference documentation for Synctera APIs. Our APIs are the best way to automate your company's banking needs and are designed to be easy to understand and implement.</p><p>We're continuously growing this library and what you see here is just the start, but if you need something specific or have a question, <a class='text-blue-600' href='https://synctera.com/contact' target='_blank' rel='noreferrer'>contact us</a>.</p> 

API version: 0.19.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package synctera_client

import (
	"encoding/json"
)

// ApplicationListResponseAllOf struct for ApplicationListResponseAllOf
type ApplicationListResponseAllOf struct {
	// Array of External Applications
	ExternalApplications []ApplicationResponse1 `json:"external_applications"`
}

// NewApplicationListResponseAllOf instantiates a new ApplicationListResponseAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApplicationListResponseAllOf(externalApplications []ApplicationResponse1) *ApplicationListResponseAllOf {
	this := ApplicationListResponseAllOf{}
	this.ExternalApplications = externalApplications
	return &this
}

// NewApplicationListResponseAllOfWithDefaults instantiates a new ApplicationListResponseAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApplicationListResponseAllOfWithDefaults() *ApplicationListResponseAllOf {
	this := ApplicationListResponseAllOf{}
	return &this
}

// GetExternalApplications returns the ExternalApplications field value
func (o *ApplicationListResponseAllOf) GetExternalApplications() []ApplicationResponse1 {
	if o == nil {
		var ret []ApplicationResponse1
		return ret
	}

	return o.ExternalApplications
}

// GetExternalApplicationsOk returns a tuple with the ExternalApplications field value
// and a boolean to check if the value has been set.
func (o *ApplicationListResponseAllOf) GetExternalApplicationsOk() ([]ApplicationResponse1, bool) {
	if o == nil {
		return nil, false
	}
	return o.ExternalApplications, true
}

// SetExternalApplications sets field value
func (o *ApplicationListResponseAllOf) SetExternalApplications(v []ApplicationResponse1) {
	o.ExternalApplications = v
}

func (o ApplicationListResponseAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["external_applications"] = o.ExternalApplications
	}
	return json.Marshal(toSerialize)
}

type NullableApplicationListResponseAllOf struct {
	value *ApplicationListResponseAllOf
	isSet bool
}

func (v NullableApplicationListResponseAllOf) Get() *ApplicationListResponseAllOf {
	return v.value
}

func (v *NullableApplicationListResponseAllOf) Set(val *ApplicationListResponseAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableApplicationListResponseAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableApplicationListResponseAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApplicationListResponseAllOf(val *ApplicationListResponseAllOf) *NullableApplicationListResponseAllOf {
	return &NullableApplicationListResponseAllOf{value: val, isSet: true}
}

func (v NullableApplicationListResponseAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApplicationListResponseAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

