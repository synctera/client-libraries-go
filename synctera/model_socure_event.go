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

// SocureEvent struct for SocureEvent
type SocureEvent struct {
	GlobalWatchlist *SocureGlobalWatchlist `json:"globalWatchlist,omitempty"`
	// A 36 character reference ID included with every ID+ response.
	ReferenceId string `json:"referenceId"`
}

// NewSocureEvent instantiates a new SocureEvent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSocureEvent(referenceId string) *SocureEvent {
	this := SocureEvent{}
	this.ReferenceId = referenceId
	return &this
}

// NewSocureEventWithDefaults instantiates a new SocureEvent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSocureEventWithDefaults() *SocureEvent {
	this := SocureEvent{}
	return &this
}

// GetGlobalWatchlist returns the GlobalWatchlist field value if set, zero value otherwise.
func (o *SocureEvent) GetGlobalWatchlist() SocureGlobalWatchlist {
	if o == nil || o.GlobalWatchlist == nil {
		var ret SocureGlobalWatchlist
		return ret
	}
	return *o.GlobalWatchlist
}

// GetGlobalWatchlistOk returns a tuple with the GlobalWatchlist field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SocureEvent) GetGlobalWatchlistOk() (*SocureGlobalWatchlist, bool) {
	if o == nil || o.GlobalWatchlist == nil {
		return nil, false
	}
	return o.GlobalWatchlist, true
}

// HasGlobalWatchlist returns a boolean if a field has been set.
func (o *SocureEvent) HasGlobalWatchlist() bool {
	if o != nil && o.GlobalWatchlist != nil {
		return true
	}

	return false
}

// SetGlobalWatchlist gets a reference to the given SocureGlobalWatchlist and assigns it to the GlobalWatchlist field.
func (o *SocureEvent) SetGlobalWatchlist(v SocureGlobalWatchlist) {
	o.GlobalWatchlist = &v
}

// GetReferenceId returns the ReferenceId field value
func (o *SocureEvent) GetReferenceId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ReferenceId
}

// GetReferenceIdOk returns a tuple with the ReferenceId field value
// and a boolean to check if the value has been set.
func (o *SocureEvent) GetReferenceIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ReferenceId, true
}

// SetReferenceId sets field value
func (o *SocureEvent) SetReferenceId(v string) {
	o.ReferenceId = v
}

func (o SocureEvent) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.GlobalWatchlist != nil {
		toSerialize["globalWatchlist"] = o.GlobalWatchlist
	}
	if true {
		toSerialize["referenceId"] = o.ReferenceId
	}
	return json.Marshal(toSerialize)
}

type NullableSocureEvent struct {
	value *SocureEvent
	isSet bool
}

func (v NullableSocureEvent) Get() *SocureEvent {
	return v.value
}

func (v *NullableSocureEvent) Set(val *SocureEvent) {
	v.value = val
	v.isSet = true
}

func (v NullableSocureEvent) IsSet() bool {
	return v.isSet
}

func (v *NullableSocureEvent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSocureEvent(val *SocureEvent) *NullableSocureEvent {
	return &NullableSocureEvent{value: val, isSet: true}
}

func (v NullableSocureEvent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSocureEvent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


