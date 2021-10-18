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

// AddVendorAccountsResponse struct for AddVendorAccountsResponse
type AddVendorAccountsResponse struct {
	AddedAccounts []ExternalAccount `json:"added_accounts"`
	FailedAccounts []AddVendorAccountFailure `json:"failed_accounts"`
}

// NewAddVendorAccountsResponse instantiates a new AddVendorAccountsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddVendorAccountsResponse(addedAccounts []ExternalAccount, failedAccounts []AddVendorAccountFailure) *AddVendorAccountsResponse {
	this := AddVendorAccountsResponse{}
	this.AddedAccounts = addedAccounts
	this.FailedAccounts = failedAccounts
	return &this
}

// NewAddVendorAccountsResponseWithDefaults instantiates a new AddVendorAccountsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddVendorAccountsResponseWithDefaults() *AddVendorAccountsResponse {
	this := AddVendorAccountsResponse{}
	return &this
}

// GetAddedAccounts returns the AddedAccounts field value
func (o *AddVendorAccountsResponse) GetAddedAccounts() []ExternalAccount {
	if o == nil {
		var ret []ExternalAccount
		return ret
	}

	return o.AddedAccounts
}

// GetAddedAccountsOk returns a tuple with the AddedAccounts field value
// and a boolean to check if the value has been set.
func (o *AddVendorAccountsResponse) GetAddedAccountsOk() (*[]ExternalAccount, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.AddedAccounts, true
}

// SetAddedAccounts sets field value
func (o *AddVendorAccountsResponse) SetAddedAccounts(v []ExternalAccount) {
	o.AddedAccounts = v
}

// GetFailedAccounts returns the FailedAccounts field value
func (o *AddVendorAccountsResponse) GetFailedAccounts() []AddVendorAccountFailure {
	if o == nil {
		var ret []AddVendorAccountFailure
		return ret
	}

	return o.FailedAccounts
}

// GetFailedAccountsOk returns a tuple with the FailedAccounts field value
// and a boolean to check if the value has been set.
func (o *AddVendorAccountsResponse) GetFailedAccountsOk() (*[]AddVendorAccountFailure, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.FailedAccounts, true
}

// SetFailedAccounts sets field value
func (o *AddVendorAccountsResponse) SetFailedAccounts(v []AddVendorAccountFailure) {
	o.FailedAccounts = v
}

func (o AddVendorAccountsResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["added_accounts"] = o.AddedAccounts
	}
	if true {
		toSerialize["failed_accounts"] = o.FailedAccounts
	}
	return json.Marshal(toSerialize)
}

type NullableAddVendorAccountsResponse struct {
	value *AddVendorAccountsResponse
	isSet bool
}

func (v NullableAddVendorAccountsResponse) Get() *AddVendorAccountsResponse {
	return v.value
}

func (v *NullableAddVendorAccountsResponse) Set(val *AddVendorAccountsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAddVendorAccountsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAddVendorAccountsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddVendorAccountsResponse(val *AddVendorAccountsResponse) *NullableAddVendorAccountsResponse {
	return &NullableAddVendorAccountsResponse{value: val, isSet: true}
}

func (v NullableAddVendorAccountsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddVendorAccountsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


