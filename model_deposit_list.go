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

// DepositList struct for DepositList
type DepositList struct {
	// Array of RDC deposits
	Deposits []Deposit `json:"deposits"`
	// If returned, use the next_page_token to query for the next page of results. Not returned if there are no more rows.
	NextPageToken *string `json:"next_page_token,omitempty"`
}

// NewDepositList instantiates a new DepositList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDepositList(deposits []Deposit) *DepositList {
	this := DepositList{}
	this.Deposits = deposits
	return &this
}

// NewDepositListWithDefaults instantiates a new DepositList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDepositListWithDefaults() *DepositList {
	this := DepositList{}
	return &this
}

// GetDeposits returns the Deposits field value
func (o *DepositList) GetDeposits() []Deposit {
	if o == nil {
		var ret []Deposit
		return ret
	}

	return o.Deposits
}

// GetDepositsOk returns a tuple with the Deposits field value
// and a boolean to check if the value has been set.
func (o *DepositList) GetDepositsOk() (*[]Deposit, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Deposits, true
}

// SetDeposits sets field value
func (o *DepositList) SetDeposits(v []Deposit) {
	o.Deposits = v
}

// GetNextPageToken returns the NextPageToken field value if set, zero value otherwise.
func (o *DepositList) GetNextPageToken() string {
	if o == nil || o.NextPageToken == nil {
		var ret string
		return ret
	}
	return *o.NextPageToken
}

// GetNextPageTokenOk returns a tuple with the NextPageToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DepositList) GetNextPageTokenOk() (*string, bool) {
	if o == nil || o.NextPageToken == nil {
		return nil, false
	}
	return o.NextPageToken, true
}

// HasNextPageToken returns a boolean if a field has been set.
func (o *DepositList) HasNextPageToken() bool {
	if o != nil && o.NextPageToken != nil {
		return true
	}

	return false
}

// SetNextPageToken gets a reference to the given string and assigns it to the NextPageToken field.
func (o *DepositList) SetNextPageToken(v string) {
	o.NextPageToken = &v
}

func (o DepositList) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["deposits"] = o.Deposits
	}
	if o.NextPageToken != nil {
		toSerialize["next_page_token"] = o.NextPageToken
	}
	return json.Marshal(toSerialize)
}

type NullableDepositList struct {
	value *DepositList
	isSet bool
}

func (v NullableDepositList) Get() *DepositList {
	return v.value
}

func (v *NullableDepositList) Set(val *DepositList) {
	v.value = val
	v.isSet = true
}

func (v NullableDepositList) IsSet() bool {
	return v.isSet
}

func (v *NullableDepositList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDepositList(val *DepositList) *NullableDepositList {
	return &NullableDepositList{value: val, isSet: true}
}

func (v NullableDepositList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDepositList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
