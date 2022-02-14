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

// InlineResponse2001 struct for InlineResponse2001
type InlineResponse2001 struct {
	// Timestamp that the old secret is delete
	DeleteAt *time.Time `json:"delete_at,omitempty"`
	// Generated secret. Do not share. This secret will be used to verify that webhook requests were sent from Synctera.
	Secret *string `json:"secret,omitempty"`
}

// NewInlineResponse2001 instantiates a new InlineResponse2001 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse2001() *InlineResponse2001 {
	this := InlineResponse2001{}
	return &this
}

// NewInlineResponse2001WithDefaults instantiates a new InlineResponse2001 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse2001WithDefaults() *InlineResponse2001 {
	this := InlineResponse2001{}
	return &this
}

// GetDeleteAt returns the DeleteAt field value if set, zero value otherwise.
func (o *InlineResponse2001) GetDeleteAt() time.Time {
	if o == nil || o.DeleteAt == nil {
		var ret time.Time
		return ret
	}
	return *o.DeleteAt
}

// GetDeleteAtOk returns a tuple with the DeleteAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse2001) GetDeleteAtOk() (*time.Time, bool) {
	if o == nil || o.DeleteAt == nil {
		return nil, false
	}
	return o.DeleteAt, true
}

// HasDeleteAt returns a boolean if a field has been set.
func (o *InlineResponse2001) HasDeleteAt() bool {
	if o != nil && o.DeleteAt != nil {
		return true
	}

	return false
}

// SetDeleteAt gets a reference to the given time.Time and assigns it to the DeleteAt field.
func (o *InlineResponse2001) SetDeleteAt(v time.Time) {
	o.DeleteAt = &v
}

// GetSecret returns the Secret field value if set, zero value otherwise.
func (o *InlineResponse2001) GetSecret() string {
	if o == nil || o.Secret == nil {
		var ret string
		return ret
	}
	return *o.Secret
}

// GetSecretOk returns a tuple with the Secret field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse2001) GetSecretOk() (*string, bool) {
	if o == nil || o.Secret == nil {
		return nil, false
	}
	return o.Secret, true
}

// HasSecret returns a boolean if a field has been set.
func (o *InlineResponse2001) HasSecret() bool {
	if o != nil && o.Secret != nil {
		return true
	}

	return false
}

// SetSecret gets a reference to the given string and assigns it to the Secret field.
func (o *InlineResponse2001) SetSecret(v string) {
	o.Secret = &v
}

func (o InlineResponse2001) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.DeleteAt != nil {
		toSerialize["delete_at"] = o.DeleteAt
	}
	if o.Secret != nil {
		toSerialize["secret"] = o.Secret
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse2001 struct {
	value *InlineResponse2001
	isSet bool
}

func (v NullableInlineResponse2001) Get() *InlineResponse2001 {
	return v.value
}

func (v *NullableInlineResponse2001) Set(val *InlineResponse2001) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse2001) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse2001) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse2001(val *InlineResponse2001) *NullableInlineResponse2001 {
	return &NullableInlineResponse2001{value: val, isSet: true}
}

func (v NullableInlineResponse2001) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse2001) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
