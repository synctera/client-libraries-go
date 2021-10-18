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

// CardChangesList struct for CardChangesList
type CardChangesList struct {
	// List of changes in descending chronological order
	Changes []CardChange `json:"changes"`
	// If returned, use the next_page_token to query for the next page of results. Not returned if there are no more rows.
	NextPageToken *string `json:"next_page_token,omitempty"`
}

// NewCardChangesList instantiates a new CardChangesList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCardChangesList(changes []CardChange) *CardChangesList {
	this := CardChangesList{}
	this.Changes = changes
	return &this
}

// NewCardChangesListWithDefaults instantiates a new CardChangesList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCardChangesListWithDefaults() *CardChangesList {
	this := CardChangesList{}
	return &this
}

// GetChanges returns the Changes field value
func (o *CardChangesList) GetChanges() []CardChange {
	if o == nil {
		var ret []CardChange
		return ret
	}

	return o.Changes
}

// GetChangesOk returns a tuple with the Changes field value
// and a boolean to check if the value has been set.
func (o *CardChangesList) GetChangesOk() (*[]CardChange, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Changes, true
}

// SetChanges sets field value
func (o *CardChangesList) SetChanges(v []CardChange) {
	o.Changes = v
}

// GetNextPageToken returns the NextPageToken field value if set, zero value otherwise.
func (o *CardChangesList) GetNextPageToken() string {
	if o == nil || o.NextPageToken == nil {
		var ret string
		return ret
	}
	return *o.NextPageToken
}

// GetNextPageTokenOk returns a tuple with the NextPageToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CardChangesList) GetNextPageTokenOk() (*string, bool) {
	if o == nil || o.NextPageToken == nil {
		return nil, false
	}
	return o.NextPageToken, true
}

// HasNextPageToken returns a boolean if a field has been set.
func (o *CardChangesList) HasNextPageToken() bool {
	if o != nil && o.NextPageToken != nil {
		return true
	}

	return false
}

// SetNextPageToken gets a reference to the given string and assigns it to the NextPageToken field.
func (o *CardChangesList) SetNextPageToken(v string) {
	o.NextPageToken = &v
}

func (o CardChangesList) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["changes"] = o.Changes
	}
	if o.NextPageToken != nil {
		toSerialize["next_page_token"] = o.NextPageToken
	}
	return json.Marshal(toSerialize)
}

type NullableCardChangesList struct {
	value *CardChangesList
	isSet bool
}

func (v NullableCardChangesList) Get() *CardChangesList {
	return v.value
}

func (v *NullableCardChangesList) Set(val *CardChangesList) {
	v.value = val
	v.isSet = true
}

func (v NullableCardChangesList) IsSet() bool {
	return v.isSet
}

func (v *NullableCardChangesList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCardChangesList(val *CardChangesList) *NullableCardChangesList {
	return &NullableCardChangesList{value: val, isSet: true}
}

func (v NullableCardChangesList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCardChangesList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


