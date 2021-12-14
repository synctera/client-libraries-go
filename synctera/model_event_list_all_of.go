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

// EventListAllOf struct for EventListAllOf
type EventListAllOf struct {
	// Array of events
	EventList []Event `json:"event_list"`
}

// NewEventListAllOf instantiates a new EventListAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEventListAllOf(eventList []Event) *EventListAllOf {
	this := EventListAllOf{}
	this.EventList = eventList
	return &this
}

// NewEventListAllOfWithDefaults instantiates a new EventListAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEventListAllOfWithDefaults() *EventListAllOf {
	this := EventListAllOf{}
	return &this
}

// GetEventList returns the EventList field value
func (o *EventListAllOf) GetEventList() []Event {
	if o == nil {
		var ret []Event
		return ret
	}

	return o.EventList
}

// GetEventListOk returns a tuple with the EventList field value
// and a boolean to check if the value has been set.
func (o *EventListAllOf) GetEventListOk() (*[]Event, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.EventList, true
}

// SetEventList sets field value
func (o *EventListAllOf) SetEventList(v []Event) {
	o.EventList = v
}

func (o EventListAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["event_list"] = o.EventList
	}
	return json.Marshal(toSerialize)
}

type NullableEventListAllOf struct {
	value *EventListAllOf
	isSet bool
}

func (v NullableEventListAllOf) Get() *EventListAllOf {
	return v.value
}

func (v *NullableEventListAllOf) Set(val *EventListAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableEventListAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableEventListAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEventListAllOf(val *EventListAllOf) *NullableEventListAllOf {
	return &NullableEventListAllOf{value: val, isSet: true}
}

func (v NullableEventListAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEventListAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


