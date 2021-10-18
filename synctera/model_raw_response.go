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

// RawResponse struct for RawResponse
type RawResponse struct {
	Provider *ProviderType `json:"provider,omitempty"`
	// the raw data from an external provider, as a JSON string
	RawData *string `json:"raw_data,omitempty"`
}

// NewRawResponse instantiates a new RawResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRawResponse() *RawResponse {
	this := RawResponse{}
	return &this
}

// NewRawResponseWithDefaults instantiates a new RawResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRawResponseWithDefaults() *RawResponse {
	this := RawResponse{}
	return &this
}

// GetProvider returns the Provider field value if set, zero value otherwise.
func (o *RawResponse) GetProvider() ProviderType {
	if o == nil || o.Provider == nil {
		var ret ProviderType
		return ret
	}
	return *o.Provider
}

// GetProviderOk returns a tuple with the Provider field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RawResponse) GetProviderOk() (*ProviderType, bool) {
	if o == nil || o.Provider == nil {
		return nil, false
	}
	return o.Provider, true
}

// HasProvider returns a boolean if a field has been set.
func (o *RawResponse) HasProvider() bool {
	if o != nil && o.Provider != nil {
		return true
	}

	return false
}

// SetProvider gets a reference to the given ProviderType and assigns it to the Provider field.
func (o *RawResponse) SetProvider(v ProviderType) {
	o.Provider = &v
}

// GetRawData returns the RawData field value if set, zero value otherwise.
func (o *RawResponse) GetRawData() string {
	if o == nil || o.RawData == nil {
		var ret string
		return ret
	}
	return *o.RawData
}

// GetRawDataOk returns a tuple with the RawData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RawResponse) GetRawDataOk() (*string, bool) {
	if o == nil || o.RawData == nil {
		return nil, false
	}
	return o.RawData, true
}

// HasRawData returns a boolean if a field has been set.
func (o *RawResponse) HasRawData() bool {
	if o != nil && o.RawData != nil {
		return true
	}

	return false
}

// SetRawData gets a reference to the given string and assigns it to the RawData field.
func (o *RawResponse) SetRawData(v string) {
	o.RawData = &v
}

func (o RawResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Provider != nil {
		toSerialize["provider"] = o.Provider
	}
	if o.RawData != nil {
		toSerialize["raw_data"] = o.RawData
	}
	return json.Marshal(toSerialize)
}

type NullableRawResponse struct {
	value *RawResponse
	isSet bool
}

func (v NullableRawResponse) Get() *RawResponse {
	return v.value
}

func (v *NullableRawResponse) Set(val *RawResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableRawResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableRawResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRawResponse(val *RawResponse) *NullableRawResponse {
	return &NullableRawResponse{value: val, isSet: true}
}

func (v NullableRawResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRawResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


