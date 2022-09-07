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

// GetViralLoopWaitlists Get Viral Loop Waitlists
type GetViralLoopWaitlists struct {
	// total number of leads/people on the waitlist
	LeadCount *int32 `json:"lead_count,omitempty"`
}

// NewGetViralLoopWaitlists instantiates a new GetViralLoopWaitlists object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetViralLoopWaitlists() *GetViralLoopWaitlists {
	this := GetViralLoopWaitlists{}
	return &this
}

// NewGetViralLoopWaitlistsWithDefaults instantiates a new GetViralLoopWaitlists object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetViralLoopWaitlistsWithDefaults() *GetViralLoopWaitlists {
	this := GetViralLoopWaitlists{}
	return &this
}

// GetLeadCount returns the LeadCount field value if set, zero value otherwise.
func (o *GetViralLoopWaitlists) GetLeadCount() int32 {
	if o == nil || o.LeadCount == nil {
		var ret int32
		return ret
	}
	return *o.LeadCount
}

// GetLeadCountOk returns a tuple with the LeadCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetViralLoopWaitlists) GetLeadCountOk() (*int32, bool) {
	if o == nil || o.LeadCount == nil {
		return nil, false
	}
	return o.LeadCount, true
}

// HasLeadCount returns a boolean if a field has been set.
func (o *GetViralLoopWaitlists) HasLeadCount() bool {
	if o != nil && o.LeadCount != nil {
		return true
	}

	return false
}

// SetLeadCount gets a reference to the given int32 and assigns it to the LeadCount field.
func (o *GetViralLoopWaitlists) SetLeadCount(v int32) {
	o.LeadCount = &v
}

func (o GetViralLoopWaitlists) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.LeadCount != nil {
		toSerialize["lead_count"] = o.LeadCount
	}
	return json.Marshal(toSerialize)
}

type NullableGetViralLoopWaitlists struct {
	value *GetViralLoopWaitlists
	isSet bool
}

func (v NullableGetViralLoopWaitlists) Get() *GetViralLoopWaitlists {
	return v.value
}

func (v *NullableGetViralLoopWaitlists) Set(val *GetViralLoopWaitlists) {
	v.value = val
	v.isSet = true
}

func (v NullableGetViralLoopWaitlists) IsSet() bool {
	return v.isSet
}

func (v *NullableGetViralLoopWaitlists) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetViralLoopWaitlists(val *GetViralLoopWaitlists) *NullableGetViralLoopWaitlists {
	return &NullableGetViralLoopWaitlists{value: val, isSet: true}
}

func (v NullableGetViralLoopWaitlists) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetViralLoopWaitlists) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


