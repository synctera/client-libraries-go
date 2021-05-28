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

// InlineObject2 struct for InlineObject2
type InlineObject2 struct {
	// Include the address information (e.g. street number) if set to True. Address reference only if set to false. Default is false
	HasDetails *bool `json:"has_details,omitempty"`
}

// NewInlineObject2 instantiates a new InlineObject2 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject2() *InlineObject2 {
	this := InlineObject2{}
	return &this
}

// NewInlineObject2WithDefaults instantiates a new InlineObject2 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject2WithDefaults() *InlineObject2 {
	this := InlineObject2{}
	return &this
}

// GetHasDetails returns the HasDetails field value if set, zero value otherwise.
func (o *InlineObject2) GetHasDetails() bool {
	if o == nil || o.HasDetails == nil {
		var ret bool
		return ret
	}
	return *o.HasDetails
}

// GetHasDetailsOk returns a tuple with the HasDetails field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject2) GetHasDetailsOk() (*bool, bool) {
	if o == nil || o.HasDetails == nil {
		return nil, false
	}
	return o.HasDetails, true
}

// HasHasDetails returns a boolean if a field has been set.
func (o *InlineObject2) HasHasDetails() bool {
	if o != nil && o.HasDetails != nil {
		return true
	}

	return false
}

// SetHasDetails gets a reference to the given bool and assigns it to the HasDetails field.
func (o *InlineObject2) SetHasDetails(v bool) {
	o.HasDetails = &v
}

func (o InlineObject2) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.HasDetails != nil {
		toSerialize["has_details"] = o.HasDetails
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject2 struct {
	value *InlineObject2
	isSet bool
}

func (v NullableInlineObject2) Get() *InlineObject2 {
	return v.value
}

func (v *NullableInlineObject2) Set(val *InlineObject2) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject2) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject2) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject2(val *InlineObject2) *NullableInlineObject2 {
	return &NullableInlineObject2{value: val, isSet: true}
}

func (v NullableInlineObject2) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject2) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


