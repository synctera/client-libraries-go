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

// CardWidgetUrlResponse A URL for activate card and set pin widgets
type CardWidgetUrlResponse struct {
	Url *string `json:"url,omitempty"`
}

// NewCardWidgetUrlResponse instantiates a new CardWidgetUrlResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCardWidgetUrlResponse() *CardWidgetUrlResponse {
	this := CardWidgetUrlResponse{}
	return &this
}

// NewCardWidgetUrlResponseWithDefaults instantiates a new CardWidgetUrlResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCardWidgetUrlResponseWithDefaults() *CardWidgetUrlResponse {
	this := CardWidgetUrlResponse{}
	return &this
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *CardWidgetUrlResponse) GetUrl() string {
	if o == nil || o.Url == nil {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CardWidgetUrlResponse) GetUrlOk() (*string, bool) {
	if o == nil || o.Url == nil {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *CardWidgetUrlResponse) HasUrl() bool {
	if o != nil && o.Url != nil {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *CardWidgetUrlResponse) SetUrl(v string) {
	o.Url = &v
}

func (o CardWidgetUrlResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Url != nil {
		toSerialize["url"] = o.Url
	}
	return json.Marshal(toSerialize)
}

type NullableCardWidgetUrlResponse struct {
	value *CardWidgetUrlResponse
	isSet bool
}

func (v NullableCardWidgetUrlResponse) Get() *CardWidgetUrlResponse {
	return v.value
}

func (v *NullableCardWidgetUrlResponse) Set(val *CardWidgetUrlResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCardWidgetUrlResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCardWidgetUrlResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCardWidgetUrlResponse(val *CardWidgetUrlResponse) *NullableCardWidgetUrlResponse {
	return &NullableCardWidgetUrlResponse{value: val, isSet: true}
}

func (v NullableCardWidgetUrlResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCardWidgetUrlResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

