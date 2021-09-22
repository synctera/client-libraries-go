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
	"time"
)

// EventResponseHistory struct for EventResponseHistory
type EventResponseHistory struct {
	// Response code from the request
	Code *int32 `json:"code,omitempty"`
	// Response body from the request
	ResponseBody *string `json:"response_body,omitempty"`
	// Timestamp that the response is received
	ResponseTime *time.Time `json:"response_time,omitempty"`
	// Timestamp that the request is sent
	SentTime *time.Time `json:"sent_time,omitempty"`
}

// NewEventResponseHistory instantiates a new EventResponseHistory object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEventResponseHistory() *EventResponseHistory {
	this := EventResponseHistory{}
	return &this
}

// NewEventResponseHistoryWithDefaults instantiates a new EventResponseHistory object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEventResponseHistoryWithDefaults() *EventResponseHistory {
	this := EventResponseHistory{}
	return &this
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *EventResponseHistory) GetCode() int32 {
	if o == nil || o.Code == nil {
		var ret int32
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventResponseHistory) GetCodeOk() (*int32, bool) {
	if o == nil || o.Code == nil {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *EventResponseHistory) HasCode() bool {
	if o != nil && o.Code != nil {
		return true
	}

	return false
}

// SetCode gets a reference to the given int32 and assigns it to the Code field.
func (o *EventResponseHistory) SetCode(v int32) {
	o.Code = &v
}

// GetResponseBody returns the ResponseBody field value if set, zero value otherwise.
func (o *EventResponseHistory) GetResponseBody() string {
	if o == nil || o.ResponseBody == nil {
		var ret string
		return ret
	}
	return *o.ResponseBody
}

// GetResponseBodyOk returns a tuple with the ResponseBody field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventResponseHistory) GetResponseBodyOk() (*string, bool) {
	if o == nil || o.ResponseBody == nil {
		return nil, false
	}
	return o.ResponseBody, true
}

// HasResponseBody returns a boolean if a field has been set.
func (o *EventResponseHistory) HasResponseBody() bool {
	if o != nil && o.ResponseBody != nil {
		return true
	}

	return false
}

// SetResponseBody gets a reference to the given string and assigns it to the ResponseBody field.
func (o *EventResponseHistory) SetResponseBody(v string) {
	o.ResponseBody = &v
}

// GetResponseTime returns the ResponseTime field value if set, zero value otherwise.
func (o *EventResponseHistory) GetResponseTime() time.Time {
	if o == nil || o.ResponseTime == nil {
		var ret time.Time
		return ret
	}
	return *o.ResponseTime
}

// GetResponseTimeOk returns a tuple with the ResponseTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventResponseHistory) GetResponseTimeOk() (*time.Time, bool) {
	if o == nil || o.ResponseTime == nil {
		return nil, false
	}
	return o.ResponseTime, true
}

// HasResponseTime returns a boolean if a field has been set.
func (o *EventResponseHistory) HasResponseTime() bool {
	if o != nil && o.ResponseTime != nil {
		return true
	}

	return false
}

// SetResponseTime gets a reference to the given time.Time and assigns it to the ResponseTime field.
func (o *EventResponseHistory) SetResponseTime(v time.Time) {
	o.ResponseTime = &v
}

// GetSentTime returns the SentTime field value if set, zero value otherwise.
func (o *EventResponseHistory) GetSentTime() time.Time {
	if o == nil || o.SentTime == nil {
		var ret time.Time
		return ret
	}
	return *o.SentTime
}

// GetSentTimeOk returns a tuple with the SentTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventResponseHistory) GetSentTimeOk() (*time.Time, bool) {
	if o == nil || o.SentTime == nil {
		return nil, false
	}
	return o.SentTime, true
}

// HasSentTime returns a boolean if a field has been set.
func (o *EventResponseHistory) HasSentTime() bool {
	if o != nil && o.SentTime != nil {
		return true
	}

	return false
}

// SetSentTime gets a reference to the given time.Time and assigns it to the SentTime field.
func (o *EventResponseHistory) SetSentTime(v time.Time) {
	o.SentTime = &v
}

func (o EventResponseHistory) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Code != nil {
		toSerialize["code"] = o.Code
	}
	if o.ResponseBody != nil {
		toSerialize["response_body"] = o.ResponseBody
	}
	if o.ResponseTime != nil {
		toSerialize["response_time"] = o.ResponseTime
	}
	if o.SentTime != nil {
		toSerialize["sent_time"] = o.SentTime
	}
	return json.Marshal(toSerialize)
}

type NullableEventResponseHistory struct {
	value *EventResponseHistory
	isSet bool
}

func (v NullableEventResponseHistory) Get() *EventResponseHistory {
	return v.value
}

func (v *NullableEventResponseHistory) Set(val *EventResponseHistory) {
	v.value = val
	v.isSet = true
}

func (v NullableEventResponseHistory) IsSet() bool {
	return v.isSet
}

func (v *NullableEventResponseHistory) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEventResponseHistory(val *EventResponseHistory) *NullableEventResponseHistory {
	return &NullableEventResponseHistory{value: val, isSet: true}
}

func (v NullableEventResponseHistory) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEventResponseHistory) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


