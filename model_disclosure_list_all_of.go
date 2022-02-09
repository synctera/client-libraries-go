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

// DisclosureListAllOf struct for DisclosureListAllOf
type DisclosureListAllOf struct {
	// Array of disclosures.
	Disclosures []Disclosure `json:"disclosures"`
}

// NewDisclosureListAllOf instantiates a new DisclosureListAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDisclosureListAllOf(disclosures []Disclosure) *DisclosureListAllOf {
	this := DisclosureListAllOf{}
	this.Disclosures = disclosures
	return &this
}

// NewDisclosureListAllOfWithDefaults instantiates a new DisclosureListAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDisclosureListAllOfWithDefaults() *DisclosureListAllOf {
	this := DisclosureListAllOf{}
	return &this
}

// GetDisclosures returns the Disclosures field value
func (o *DisclosureListAllOf) GetDisclosures() []Disclosure {
	if o == nil {
		var ret []Disclosure
		return ret
	}

	return o.Disclosures
}

// GetDisclosuresOk returns a tuple with the Disclosures field value
// and a boolean to check if the value has been set.
func (o *DisclosureListAllOf) GetDisclosuresOk() (*[]Disclosure, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Disclosures, true
}

// SetDisclosures sets field value
func (o *DisclosureListAllOf) SetDisclosures(v []Disclosure) {
	o.Disclosures = v
}

func (o DisclosureListAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["disclosures"] = o.Disclosures
	}
	return json.Marshal(toSerialize)
}

type NullableDisclosureListAllOf struct {
	value *DisclosureListAllOf
	isSet bool
}

func (v NullableDisclosureListAllOf) Get() *DisclosureListAllOf {
	return v.value
}

func (v *NullableDisclosureListAllOf) Set(val *DisclosureListAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableDisclosureListAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableDisclosureListAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDisclosureListAllOf(val *DisclosureListAllOf) *NullableDisclosureListAllOf {
	return &NullableDisclosureListAllOf{value: val, isSet: true}
}

func (v NullableDisclosureListAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDisclosureListAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
