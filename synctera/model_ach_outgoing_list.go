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

// AchOutgoingList struct for AchOutgoingList
type AchOutgoingList struct {
	// Array of ACH
	AchOutgoings []AchOutgoing `json:"ach_outgoings"`
	// If returned, use the next_page_token to query for the next page of results. Not returned if there are no more rows.
	NextPageToken *string `json:"next_page_token,omitempty"`
}

// NewAchOutgoingList instantiates a new AchOutgoingList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAchOutgoingList(achOutgoings []AchOutgoing) *AchOutgoingList {
	this := AchOutgoingList{}
	this.AchOutgoings = achOutgoings
	return &this
}

// NewAchOutgoingListWithDefaults instantiates a new AchOutgoingList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAchOutgoingListWithDefaults() *AchOutgoingList {
	this := AchOutgoingList{}
	return &this
}

// GetAchOutgoings returns the AchOutgoings field value
func (o *AchOutgoingList) GetAchOutgoings() []AchOutgoing {
	if o == nil {
		var ret []AchOutgoing
		return ret
	}

	return o.AchOutgoings
}

// GetAchOutgoingsOk returns a tuple with the AchOutgoings field value
// and a boolean to check if the value has been set.
func (o *AchOutgoingList) GetAchOutgoingsOk() (*[]AchOutgoing, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.AchOutgoings, true
}

// SetAchOutgoings sets field value
func (o *AchOutgoingList) SetAchOutgoings(v []AchOutgoing) {
	o.AchOutgoings = v
}

// GetNextPageToken returns the NextPageToken field value if set, zero value otherwise.
func (o *AchOutgoingList) GetNextPageToken() string {
	if o == nil || o.NextPageToken == nil {
		var ret string
		return ret
	}
	return *o.NextPageToken
}

// GetNextPageTokenOk returns a tuple with the NextPageToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AchOutgoingList) GetNextPageTokenOk() (*string, bool) {
	if o == nil || o.NextPageToken == nil {
		return nil, false
	}
	return o.NextPageToken, true
}

// HasNextPageToken returns a boolean if a field has been set.
func (o *AchOutgoingList) HasNextPageToken() bool {
	if o != nil && o.NextPageToken != nil {
		return true
	}

	return false
}

// SetNextPageToken gets a reference to the given string and assigns it to the NextPageToken field.
func (o *AchOutgoingList) SetNextPageToken(v string) {
	o.NextPageToken = &v
}

func (o AchOutgoingList) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ach_outgoings"] = o.AchOutgoings
	}
	if o.NextPageToken != nil {
		toSerialize["next_page_token"] = o.NextPageToken
	}
	return json.Marshal(toSerialize)
}

type NullableAchOutgoingList struct {
	value *AchOutgoingList
	isSet bool
}

func (v NullableAchOutgoingList) Get() *AchOutgoingList {
	return v.value
}

func (v *NullableAchOutgoingList) Set(val *AchOutgoingList) {
	v.value = val
	v.isSet = true
}

func (v NullableAchOutgoingList) IsSet() bool {
	return v.isSet
}

func (v *NullableAchOutgoingList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAchOutgoingList(val *AchOutgoingList) *NullableAchOutgoingList {
	return &NullableAchOutgoingList{value: val, isSet: true}
}

func (v NullableAchOutgoingList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAchOutgoingList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


