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

// PatchAccountsRequestRoutingIdentifiers struct for PatchAccountsRequestRoutingIdentifiers
type PatchAccountsRequestRoutingIdentifiers struct {
	// The routing number used for US ACH payments. 
	AchRoutingNumber *string `json:"ach_routing_number,omitempty"`
	// The name of the bank managing the account
	BankName *string `json:"bank_name,omitempty"`
}

// NewPatchAccountsRequestRoutingIdentifiers instantiates a new PatchAccountsRequestRoutingIdentifiers object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPatchAccountsRequestRoutingIdentifiers() *PatchAccountsRequestRoutingIdentifiers {
	this := PatchAccountsRequestRoutingIdentifiers{}
	return &this
}

// NewPatchAccountsRequestRoutingIdentifiersWithDefaults instantiates a new PatchAccountsRequestRoutingIdentifiers object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPatchAccountsRequestRoutingIdentifiersWithDefaults() *PatchAccountsRequestRoutingIdentifiers {
	this := PatchAccountsRequestRoutingIdentifiers{}
	return &this
}

// GetAchRoutingNumber returns the AchRoutingNumber field value if set, zero value otherwise.
func (o *PatchAccountsRequestRoutingIdentifiers) GetAchRoutingNumber() string {
	if o == nil || o.AchRoutingNumber == nil {
		var ret string
		return ret
	}
	return *o.AchRoutingNumber
}

// GetAchRoutingNumberOk returns a tuple with the AchRoutingNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchAccountsRequestRoutingIdentifiers) GetAchRoutingNumberOk() (*string, bool) {
	if o == nil || o.AchRoutingNumber == nil {
		return nil, false
	}
	return o.AchRoutingNumber, true
}

// HasAchRoutingNumber returns a boolean if a field has been set.
func (o *PatchAccountsRequestRoutingIdentifiers) HasAchRoutingNumber() bool {
	if o != nil && o.AchRoutingNumber != nil {
		return true
	}

	return false
}

// SetAchRoutingNumber gets a reference to the given string and assigns it to the AchRoutingNumber field.
func (o *PatchAccountsRequestRoutingIdentifiers) SetAchRoutingNumber(v string) {
	o.AchRoutingNumber = &v
}

// GetBankName returns the BankName field value if set, zero value otherwise.
func (o *PatchAccountsRequestRoutingIdentifiers) GetBankName() string {
	if o == nil || o.BankName == nil {
		var ret string
		return ret
	}
	return *o.BankName
}

// GetBankNameOk returns a tuple with the BankName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchAccountsRequestRoutingIdentifiers) GetBankNameOk() (*string, bool) {
	if o == nil || o.BankName == nil {
		return nil, false
	}
	return o.BankName, true
}

// HasBankName returns a boolean if a field has been set.
func (o *PatchAccountsRequestRoutingIdentifiers) HasBankName() bool {
	if o != nil && o.BankName != nil {
		return true
	}

	return false
}

// SetBankName gets a reference to the given string and assigns it to the BankName field.
func (o *PatchAccountsRequestRoutingIdentifiers) SetBankName(v string) {
	o.BankName = &v
}

func (o PatchAccountsRequestRoutingIdentifiers) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AchRoutingNumber != nil {
		toSerialize["ach_routing_number"] = o.AchRoutingNumber
	}
	if o.BankName != nil {
		toSerialize["bank_name"] = o.BankName
	}
	return json.Marshal(toSerialize)
}

type NullablePatchAccountsRequestRoutingIdentifiers struct {
	value *PatchAccountsRequestRoutingIdentifiers
	isSet bool
}

func (v NullablePatchAccountsRequestRoutingIdentifiers) Get() *PatchAccountsRequestRoutingIdentifiers {
	return v.value
}

func (v *NullablePatchAccountsRequestRoutingIdentifiers) Set(val *PatchAccountsRequestRoutingIdentifiers) {
	v.value = val
	v.isSet = true
}

func (v NullablePatchAccountsRequestRoutingIdentifiers) IsSet() bool {
	return v.isSet
}

func (v *NullablePatchAccountsRequestRoutingIdentifiers) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePatchAccountsRequestRoutingIdentifiers(val *PatchAccountsRequestRoutingIdentifiers) *NullablePatchAccountsRequestRoutingIdentifiers {
	return &NullablePatchAccountsRequestRoutingIdentifiers{value: val, isSet: true}
}

func (v NullablePatchAccountsRequestRoutingIdentifiers) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePatchAccountsRequestRoutingIdentifiers) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


