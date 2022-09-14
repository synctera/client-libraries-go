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

// AdditionalData Contains additional information about the relationship.
type AdditionalData struct {
	// The professional role or position the person holds at the related organization.
	Title string `json:"title"`
}

// NewAdditionalData instantiates a new AdditionalData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAdditionalData(title string) *AdditionalData {
	this := AdditionalData{}
	this.Title = title
	return &this
}

// NewAdditionalDataWithDefaults instantiates a new AdditionalData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAdditionalDataWithDefaults() *AdditionalData {
	this := AdditionalData{}
	return &this
}

// GetTitle returns the Title field value
func (o *AdditionalData) GetTitle() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Title
}

// GetTitleOk returns a tuple with the Title field value
// and a boolean to check if the value has been set.
func (o *AdditionalData) GetTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Title, true
}

// SetTitle sets field value
func (o *AdditionalData) SetTitle(v string) {
	o.Title = v
}

func (o AdditionalData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["title"] = o.Title
	}
	return json.Marshal(toSerialize)
}

type NullableAdditionalData struct {
	value *AdditionalData
	isSet bool
}

func (v NullableAdditionalData) Get() *AdditionalData {
	return v.value
}

func (v *NullableAdditionalData) Set(val *AdditionalData) {
	v.value = val
	v.isSet = true
}

func (v NullableAdditionalData) IsSet() bool {
	return v.isSet
}

func (v *NullableAdditionalData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAdditionalData(val *AdditionalData) *NullableAdditionalData {
	return &NullableAdditionalData{value: val, isSet: true}
}

func (v NullableAdditionalData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAdditionalData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
