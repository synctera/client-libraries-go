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

// ProspectsListAllOf struct for ProspectsListAllOf
type ProspectsListAllOf struct {
	// Array of Prospects
	Prospects []Prospect1 `json:"prospects"`
}

// NewProspectsListAllOf instantiates a new ProspectsListAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProspectsListAllOf(prospects []Prospect1) *ProspectsListAllOf {
	this := ProspectsListAllOf{}
	this.Prospects = prospects
	return &this
}

// NewProspectsListAllOfWithDefaults instantiates a new ProspectsListAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProspectsListAllOfWithDefaults() *ProspectsListAllOf {
	this := ProspectsListAllOf{}
	return &this
}

// GetProspects returns the Prospects field value
func (o *ProspectsListAllOf) GetProspects() []Prospect1 {
	if o == nil {
		var ret []Prospect1
		return ret
	}

	return o.Prospects
}

// GetProspectsOk returns a tuple with the Prospects field value
// and a boolean to check if the value has been set.
func (o *ProspectsListAllOf) GetProspectsOk() ([]Prospect1, bool) {
	if o == nil {
		return nil, false
	}
	return o.Prospects, true
}

// SetProspects sets field value
func (o *ProspectsListAllOf) SetProspects(v []Prospect1) {
	o.Prospects = v
}

func (o ProspectsListAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["prospects"] = o.Prospects
	}
	return json.Marshal(toSerialize)
}

type NullableProspectsListAllOf struct {
	value *ProspectsListAllOf
	isSet bool
}

func (v NullableProspectsListAllOf) Get() *ProspectsListAllOf {
	return v.value
}

func (v *NullableProspectsListAllOf) Set(val *ProspectsListAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableProspectsListAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableProspectsListAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProspectsListAllOf(val *ProspectsListAllOf) *NullableProspectsListAllOf {
	return &NullableProspectsListAllOf{value: val, isSet: true}
}

func (v NullableProspectsListAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProspectsListAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


