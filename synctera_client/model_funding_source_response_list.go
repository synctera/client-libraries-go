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

// FundingSourceResponseList struct for FundingSourceResponseList
type FundingSourceResponseList struct {
	// Array of funding sources
	FundingSources []FundingSourceResponse `json:"funding_sources"`
}

// NewFundingSourceResponseList instantiates a new FundingSourceResponseList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFundingSourceResponseList(fundingSources []FundingSourceResponse) *FundingSourceResponseList {
	this := FundingSourceResponseList{}
	this.FundingSources = fundingSources
	return &this
}

// NewFundingSourceResponseListWithDefaults instantiates a new FundingSourceResponseList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFundingSourceResponseListWithDefaults() *FundingSourceResponseList {
	this := FundingSourceResponseList{}
	return &this
}

// GetFundingSources returns the FundingSources field value
func (o *FundingSourceResponseList) GetFundingSources() []FundingSourceResponse {
	if o == nil {
		var ret []FundingSourceResponse
		return ret
	}

	return o.FundingSources
}

// GetFundingSourcesOk returns a tuple with the FundingSources field value
// and a boolean to check if the value has been set.
func (o *FundingSourceResponseList) GetFundingSourcesOk() ([]FundingSourceResponse, bool) {
	if o == nil {
		return nil, false
	}
	return o.FundingSources, true
}

// SetFundingSources sets field value
func (o *FundingSourceResponseList) SetFundingSources(v []FundingSourceResponse) {
	o.FundingSources = v
}

func (o FundingSourceResponseList) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["funding_sources"] = o.FundingSources
	}
	return json.Marshal(toSerialize)
}

type NullableFundingSourceResponseList struct {
	value *FundingSourceResponseList
	isSet bool
}

func (v NullableFundingSourceResponseList) Get() *FundingSourceResponseList {
	return v.value
}

func (v *NullableFundingSourceResponseList) Set(val *FundingSourceResponseList) {
	v.value = val
	v.isSet = true
}

func (v NullableFundingSourceResponseList) IsSet() bool {
	return v.isSet
}

func (v *NullableFundingSourceResponseList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFundingSourceResponseList(val *FundingSourceResponseList) *NullableFundingSourceResponseList {
	return &NullableFundingSourceResponseList{value: val, isSet: true}
}

func (v NullableFundingSourceResponseList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFundingSourceResponseList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

