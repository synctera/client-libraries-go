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

// DebitNetworkResponseListAllOf struct for DebitNetworkResponseListAllOf
type DebitNetworkResponseListAllOf struct {
	// Array of debit networks
	DebitNetworks []DebitNetworkResponse `json:"debit_networks"`
}

// NewDebitNetworkResponseListAllOf instantiates a new DebitNetworkResponseListAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDebitNetworkResponseListAllOf(debitNetworks []DebitNetworkResponse) *DebitNetworkResponseListAllOf {
	this := DebitNetworkResponseListAllOf{}
	this.DebitNetworks = debitNetworks
	return &this
}

// NewDebitNetworkResponseListAllOfWithDefaults instantiates a new DebitNetworkResponseListAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDebitNetworkResponseListAllOfWithDefaults() *DebitNetworkResponseListAllOf {
	this := DebitNetworkResponseListAllOf{}
	return &this
}

// GetDebitNetworks returns the DebitNetworks field value
func (o *DebitNetworkResponseListAllOf) GetDebitNetworks() []DebitNetworkResponse {
	if o == nil {
		var ret []DebitNetworkResponse
		return ret
	}

	return o.DebitNetworks
}

// GetDebitNetworksOk returns a tuple with the DebitNetworks field value
// and a boolean to check if the value has been set.
func (o *DebitNetworkResponseListAllOf) GetDebitNetworksOk() (*[]DebitNetworkResponse, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.DebitNetworks, true
}

// SetDebitNetworks sets field value
func (o *DebitNetworkResponseListAllOf) SetDebitNetworks(v []DebitNetworkResponse) {
	o.DebitNetworks = v
}

func (o DebitNetworkResponseListAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["debit_networks"] = o.DebitNetworks
	}
	return json.Marshal(toSerialize)
}

type NullableDebitNetworkResponseListAllOf struct {
	value *DebitNetworkResponseListAllOf
	isSet bool
}

func (v NullableDebitNetworkResponseListAllOf) Get() *DebitNetworkResponseListAllOf {
	return v.value
}

func (v *NullableDebitNetworkResponseListAllOf) Set(val *DebitNetworkResponseListAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableDebitNetworkResponseListAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableDebitNetworkResponseListAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDebitNetworkResponseListAllOf(val *DebitNetworkResponseListAllOf) *NullableDebitNetworkResponseListAllOf {
	return &NullableDebitNetworkResponseListAllOf{value: val, isSet: true}
}

func (v NullableDebitNetworkResponseListAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDebitNetworkResponseListAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


