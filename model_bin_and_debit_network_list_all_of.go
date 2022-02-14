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

// BinAndDebitNetworkListAllOf struct for BinAndDebitNetworkListAllOf
type BinAndDebitNetworkListAllOf struct {
	// Array of BINs and debit networks
	BinAndDebitNetworks []BinAndDebitNetwork `json:"bin_and_debit_networks"`
}

// NewBinAndDebitNetworkListAllOf instantiates a new BinAndDebitNetworkListAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBinAndDebitNetworkListAllOf(binAndDebitNetworks []BinAndDebitNetwork) *BinAndDebitNetworkListAllOf {
	this := BinAndDebitNetworkListAllOf{}
	this.BinAndDebitNetworks = binAndDebitNetworks
	return &this
}

// NewBinAndDebitNetworkListAllOfWithDefaults instantiates a new BinAndDebitNetworkListAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBinAndDebitNetworkListAllOfWithDefaults() *BinAndDebitNetworkListAllOf {
	this := BinAndDebitNetworkListAllOf{}
	return &this
}

// GetBinAndDebitNetworks returns the BinAndDebitNetworks field value
func (o *BinAndDebitNetworkListAllOf) GetBinAndDebitNetworks() []BinAndDebitNetwork {
	if o == nil {
		var ret []BinAndDebitNetwork
		return ret
	}

	return o.BinAndDebitNetworks
}

// GetBinAndDebitNetworksOk returns a tuple with the BinAndDebitNetworks field value
// and a boolean to check if the value has been set.
func (o *BinAndDebitNetworkListAllOf) GetBinAndDebitNetworksOk() (*[]BinAndDebitNetwork, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BinAndDebitNetworks, true
}

// SetBinAndDebitNetworks sets field value
func (o *BinAndDebitNetworkListAllOf) SetBinAndDebitNetworks(v []BinAndDebitNetwork) {
	o.BinAndDebitNetworks = v
}

func (o BinAndDebitNetworkListAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["bin_and_debit_networks"] = o.BinAndDebitNetworks
	}
	return json.Marshal(toSerialize)
}

type NullableBinAndDebitNetworkListAllOf struct {
	value *BinAndDebitNetworkListAllOf
	isSet bool
}

func (v NullableBinAndDebitNetworkListAllOf) Get() *BinAndDebitNetworkListAllOf {
	return v.value
}

func (v *NullableBinAndDebitNetworkListAllOf) Set(val *BinAndDebitNetworkListAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableBinAndDebitNetworkListAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableBinAndDebitNetworkListAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBinAndDebitNetworkListAllOf(val *BinAndDebitNetworkListAllOf) *NullableBinAndDebitNetworkListAllOf {
	return &NullableBinAndDebitNetworkListAllOf{value: val, isSet: true}
}

func (v NullableBinAndDebitNetworkListAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBinAndDebitNetworkListAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}