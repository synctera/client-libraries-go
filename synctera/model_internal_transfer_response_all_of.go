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

// InternalTransferResponseAllOf struct for InternalTransferResponseAllOf
type InternalTransferResponseAllOf struct {
	// The transaction id associated with the transfer
	Id string `json:"id"`
}

// NewInternalTransferResponseAllOf instantiates a new InternalTransferResponseAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInternalTransferResponseAllOf(id string) *InternalTransferResponseAllOf {
	this := InternalTransferResponseAllOf{}
	this.Id = id
	return &this
}

// NewInternalTransferResponseAllOfWithDefaults instantiates a new InternalTransferResponseAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInternalTransferResponseAllOfWithDefaults() *InternalTransferResponseAllOf {
	this := InternalTransferResponseAllOf{}
	return &this
}

// GetId returns the Id field value
func (o *InternalTransferResponseAllOf) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *InternalTransferResponseAllOf) GetIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *InternalTransferResponseAllOf) SetId(v string) {
	o.Id = v
}

func (o InternalTransferResponseAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	return json.Marshal(toSerialize)
}

type NullableInternalTransferResponseAllOf struct {
	value *InternalTransferResponseAllOf
	isSet bool
}

func (v NullableInternalTransferResponseAllOf) Get() *InternalTransferResponseAllOf {
	return v.value
}

func (v *NullableInternalTransferResponseAllOf) Set(val *InternalTransferResponseAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableInternalTransferResponseAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableInternalTransferResponseAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInternalTransferResponseAllOf(val *InternalTransferResponseAllOf) *NullableInternalTransferResponseAllOf {
	return &NullableInternalTransferResponseAllOf{value: val, isSet: true}
}

func (v NullableInternalTransferResponseAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInternalTransferResponseAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


