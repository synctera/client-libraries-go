/*
Synctera API

<h2>Let's build something great.</h2><p>Welcome to the official reference documentation for Synctera APIs. Our APIs are the best way to automate your company's banking needs and are designed to be easy to understand and implement.</p><p>We're continuously growing this library and what you see here is just the start, but if you need something specific or have a question, <a class='text-blue-600' href='https://synctera.com/contact' target='_blank' rel='noreferrer'>contact us</a>.</p>

API version: 0.20.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package synctera_client

import (
	"encoding/json"
)

// BankDebitNetworkResponseAllOf struct for BankDebitNetworkResponseAllOf
type BankDebitNetworkResponseAllOf struct {
	// The ID of the bank network
	BankNetworkId *string `json:"bank_network_id,omitempty"`
	// The ID of the bank's bin that uses this debit network
	BinId *string `json:"bin_id,omitempty"`
}

// NewBankDebitNetworkResponseAllOf instantiates a new BankDebitNetworkResponseAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBankDebitNetworkResponseAllOf() *BankDebitNetworkResponseAllOf {
	this := BankDebitNetworkResponseAllOf{}
	return &this
}

// NewBankDebitNetworkResponseAllOfWithDefaults instantiates a new BankDebitNetworkResponseAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBankDebitNetworkResponseAllOfWithDefaults() *BankDebitNetworkResponseAllOf {
	this := BankDebitNetworkResponseAllOf{}
	return &this
}

// GetBankNetworkId returns the BankNetworkId field value if set, zero value otherwise.
func (o *BankDebitNetworkResponseAllOf) GetBankNetworkId() string {
	if o == nil || o.BankNetworkId == nil {
		var ret string
		return ret
	}
	return *o.BankNetworkId
}

// GetBankNetworkIdOk returns a tuple with the BankNetworkId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BankDebitNetworkResponseAllOf) GetBankNetworkIdOk() (*string, bool) {
	if o == nil || o.BankNetworkId == nil {
		return nil, false
	}
	return o.BankNetworkId, true
}

// HasBankNetworkId returns a boolean if a field has been set.
func (o *BankDebitNetworkResponseAllOf) HasBankNetworkId() bool {
	if o != nil && o.BankNetworkId != nil {
		return true
	}

	return false
}

// SetBankNetworkId gets a reference to the given string and assigns it to the BankNetworkId field.
func (o *BankDebitNetworkResponseAllOf) SetBankNetworkId(v string) {
	o.BankNetworkId = &v
}

// GetBinId returns the BinId field value if set, zero value otherwise.
func (o *BankDebitNetworkResponseAllOf) GetBinId() string {
	if o == nil || o.BinId == nil {
		var ret string
		return ret
	}
	return *o.BinId
}

// GetBinIdOk returns a tuple with the BinId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BankDebitNetworkResponseAllOf) GetBinIdOk() (*string, bool) {
	if o == nil || o.BinId == nil {
		return nil, false
	}
	return o.BinId, true
}

// HasBinId returns a boolean if a field has been set.
func (o *BankDebitNetworkResponseAllOf) HasBinId() bool {
	if o != nil && o.BinId != nil {
		return true
	}

	return false
}

// SetBinId gets a reference to the given string and assigns it to the BinId field.
func (o *BankDebitNetworkResponseAllOf) SetBinId(v string) {
	o.BinId = &v
}

func (o BankDebitNetworkResponseAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.BankNetworkId != nil {
		toSerialize["bank_network_id"] = o.BankNetworkId
	}
	if o.BinId != nil {
		toSerialize["bin_id"] = o.BinId
	}
	return json.Marshal(toSerialize)
}

type NullableBankDebitNetworkResponseAllOf struct {
	value *BankDebitNetworkResponseAllOf
	isSet bool
}

func (v NullableBankDebitNetworkResponseAllOf) Get() *BankDebitNetworkResponseAllOf {
	return v.value
}

func (v *NullableBankDebitNetworkResponseAllOf) Set(val *BankDebitNetworkResponseAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableBankDebitNetworkResponseAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableBankDebitNetworkResponseAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBankDebitNetworkResponseAllOf(val *BankDebitNetworkResponseAllOf) *NullableBankDebitNetworkResponseAllOf {
	return &NullableBankDebitNetworkResponseAllOf{value: val, isSet: true}
}

func (v NullableBankDebitNetworkResponseAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBankDebitNetworkResponseAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
