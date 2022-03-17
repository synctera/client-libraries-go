/*
Synctera API

<h2>Let's build something great.</h2><p>Welcome to the official reference documentation for Synctera APIs. Our APIs are the best way to automate your company's banking needs and are designed to be easy to understand and implement.</p><p>We're continuously growing this library and what you see here is just the start, but if you need something specific or have a question, <a class='text-blue-600' href='https://synctera.com/contact' target='_blank' rel='noreferrer'>contact us</a>.</p>

API version: 0.5.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package synctera

import (
	"encoding/json"
)

// DisclosureResponse Represents all customer disclosures
type DisclosureResponse struct {
	// List of the customer's disclosures
	Disclosures *[]Disclosure1 `json:"disclosures,omitempty"`
}

// NewDisclosureResponse instantiates a new DisclosureResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDisclosureResponse() *DisclosureResponse {
	this := DisclosureResponse{}
	return &this
}

// NewDisclosureResponseWithDefaults instantiates a new DisclosureResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDisclosureResponseWithDefaults() *DisclosureResponse {
	this := DisclosureResponse{}
	return &this
}

// GetDisclosures returns the Disclosures field value if set, zero value otherwise.
func (o *DisclosureResponse) GetDisclosures() []Disclosure1 {
	if o == nil || o.Disclosures == nil {
		var ret []Disclosure1
		return ret
	}
	return *o.Disclosures
}

// GetDisclosuresOk returns a tuple with the Disclosures field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DisclosureResponse) GetDisclosuresOk() (*[]Disclosure1, bool) {
	if o == nil || o.Disclosures == nil {
		return nil, false
	}
	return o.Disclosures, true
}

// HasDisclosures returns a boolean if a field has been set.
func (o *DisclosureResponse) HasDisclosures() bool {
	if o != nil && o.Disclosures != nil {
		return true
	}

	return false
}

// SetDisclosures gets a reference to the given []Disclosure1 and assigns it to the Disclosures field.
func (o *DisclosureResponse) SetDisclosures(v []Disclosure1) {
	o.Disclosures = &v
}

func (o DisclosureResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Disclosures != nil {
		toSerialize["disclosures"] = o.Disclosures
	}
	return json.Marshal(toSerialize)
}

type NullableDisclosureResponse struct {
	value *DisclosureResponse
	isSet bool
}

func (v NullableDisclosureResponse) Get() *DisclosureResponse {
	return v.value
}

func (v *NullableDisclosureResponse) Set(val *DisclosureResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableDisclosureResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableDisclosureResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDisclosureResponse(val *DisclosureResponse) *NullableDisclosureResponse {
	return &NullableDisclosureResponse{value: val, isSet: true}
}

func (v NullableDisclosureResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDisclosureResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
