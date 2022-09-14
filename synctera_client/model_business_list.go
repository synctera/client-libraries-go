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

// BusinessList struct for BusinessList
type BusinessList struct {
	// Array of businesses.
	Businesses []Business `json:"businesses"`
	// If returned, use the next_page_token to query for the next page of results. Not returned if there are no more rows.
	NextPageToken *string `json:"next_page_token,omitempty"`
}

// NewBusinessList instantiates a new BusinessList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBusinessList(businesses []Business) *BusinessList {
	this := BusinessList{}
	this.Businesses = businesses
	return &this
}

// NewBusinessListWithDefaults instantiates a new BusinessList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBusinessListWithDefaults() *BusinessList {
	this := BusinessList{}
	return &this
}

// GetBusinesses returns the Businesses field value
func (o *BusinessList) GetBusinesses() []Business {
	if o == nil {
		var ret []Business
		return ret
	}

	return o.Businesses
}

// GetBusinessesOk returns a tuple with the Businesses field value
// and a boolean to check if the value has been set.
func (o *BusinessList) GetBusinessesOk() ([]Business, bool) {
	if o == nil {
		return nil, false
	}
	return o.Businesses, true
}

// SetBusinesses sets field value
func (o *BusinessList) SetBusinesses(v []Business) {
	o.Businesses = v
}

// GetNextPageToken returns the NextPageToken field value if set, zero value otherwise.
func (o *BusinessList) GetNextPageToken() string {
	if o == nil || o.NextPageToken == nil {
		var ret string
		return ret
	}
	return *o.NextPageToken
}

// GetNextPageTokenOk returns a tuple with the NextPageToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BusinessList) GetNextPageTokenOk() (*string, bool) {
	if o == nil || o.NextPageToken == nil {
		return nil, false
	}
	return o.NextPageToken, true
}

// HasNextPageToken returns a boolean if a field has been set.
func (o *BusinessList) HasNextPageToken() bool {
	if o != nil && o.NextPageToken != nil {
		return true
	}

	return false
}

// SetNextPageToken gets a reference to the given string and assigns it to the NextPageToken field.
func (o *BusinessList) SetNextPageToken(v string) {
	o.NextPageToken = &v
}

func (o BusinessList) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["businesses"] = o.Businesses
	}
	if o.NextPageToken != nil {
		toSerialize["next_page_token"] = o.NextPageToken
	}
	return json.Marshal(toSerialize)
}

type NullableBusinessList struct {
	value *BusinessList
	isSet bool
}

func (v NullableBusinessList) Get() *BusinessList {
	return v.value
}

func (v *NullableBusinessList) Set(val *BusinessList) {
	v.value = val
	v.isSet = true
}

func (v NullableBusinessList) IsSet() bool {
	return v.isSet
}

func (v *NullableBusinessList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBusinessList(val *BusinessList) *NullableBusinessList {
	return &NullableBusinessList{value: val, isSet: true}
}

func (v NullableBusinessList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBusinessList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
