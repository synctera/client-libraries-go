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

// EmploymentList struct for EmploymentList
type EmploymentList struct {
	// Array of customer employment records.
	Employment []Employment `json:"employment"`
	// If returned, use the next_page_token to query for the next page of results. Not returned if there are no more rows.
	NextPageToken *string `json:"next_page_token,omitempty"`
}

// NewEmploymentList instantiates a new EmploymentList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEmploymentList(employment []Employment) *EmploymentList {
	this := EmploymentList{}
	this.Employment = employment
	return &this
}

// NewEmploymentListWithDefaults instantiates a new EmploymentList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEmploymentListWithDefaults() *EmploymentList {
	this := EmploymentList{}
	return &this
}

// GetEmployment returns the Employment field value
func (o *EmploymentList) GetEmployment() []Employment {
	if o == nil {
		var ret []Employment
		return ret
	}

	return o.Employment
}

// GetEmploymentOk returns a tuple with the Employment field value
// and a boolean to check if the value has been set.
func (o *EmploymentList) GetEmploymentOk() (*[]Employment, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Employment, true
}

// SetEmployment sets field value
func (o *EmploymentList) SetEmployment(v []Employment) {
	o.Employment = v
}

// GetNextPageToken returns the NextPageToken field value if set, zero value otherwise.
func (o *EmploymentList) GetNextPageToken() string {
	if o == nil || o.NextPageToken == nil {
		var ret string
		return ret
	}
	return *o.NextPageToken
}

// GetNextPageTokenOk returns a tuple with the NextPageToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmploymentList) GetNextPageTokenOk() (*string, bool) {
	if o == nil || o.NextPageToken == nil {
		return nil, false
	}
	return o.NextPageToken, true
}

// HasNextPageToken returns a boolean if a field has been set.
func (o *EmploymentList) HasNextPageToken() bool {
	if o != nil && o.NextPageToken != nil {
		return true
	}

	return false
}

// SetNextPageToken gets a reference to the given string and assigns it to the NextPageToken field.
func (o *EmploymentList) SetNextPageToken(v string) {
	o.NextPageToken = &v
}

func (o EmploymentList) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["employment"] = o.Employment
	}
	if o.NextPageToken != nil {
		toSerialize["next_page_token"] = o.NextPageToken
	}
	return json.Marshal(toSerialize)
}

type NullableEmploymentList struct {
	value *EmploymentList
	isSet bool
}

func (v NullableEmploymentList) Get() *EmploymentList {
	return v.value
}

func (v *NullableEmploymentList) Set(val *EmploymentList) {
	v.value = val
	v.isSet = true
}

func (v NullableEmploymentList) IsSet() bool {
	return v.isSet
}

func (v *NullableEmploymentList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEmploymentList(val *EmploymentList) *NullableEmploymentList {
	return &NullableEmploymentList{value: val, isSet: true}
}

func (v NullableEmploymentList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEmploymentList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


