/*
Synctera API

<h2>Let's build something great.</h2><p>Welcome to the official reference documentation for Synctera APIs. Our APIs are the best way to automate your company's banking needs and are designed to be easy to understand and implement.</p><p>We're continuously growing this library and what you see here is just the start, but if you need something specific or have a question, <a class='text-blue-600' href='https://synctera.com/contact' target='_blank' rel='noreferrer'>contact us</a>.</p> 

API version: 0.5.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package synctera

import (
	"encoding/json"
	"time"
)

// InlineObject1 struct for InlineObject1
type InlineObject1 struct {
	BalanceType *BalanceType `json:"balance_type,omitempty"`
	// Posting date of the balance. Default is today's date
	PostingDate *time.Time `json:"posting_date,omitempty"`
}

// NewInlineObject1 instantiates a new InlineObject1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject1() *InlineObject1 {
	this := InlineObject1{}
	return &this
}

// NewInlineObject1WithDefaults instantiates a new InlineObject1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject1WithDefaults() *InlineObject1 {
	this := InlineObject1{}
	return &this
}

// GetBalanceType returns the BalanceType field value if set, zero value otherwise.
func (o *InlineObject1) GetBalanceType() BalanceType {
	if o == nil || o.BalanceType == nil {
		var ret BalanceType
		return ret
	}
	return *o.BalanceType
}

// GetBalanceTypeOk returns a tuple with the BalanceType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject1) GetBalanceTypeOk() (*BalanceType, bool) {
	if o == nil || o.BalanceType == nil {
		return nil, false
	}
	return o.BalanceType, true
}

// HasBalanceType returns a boolean if a field has been set.
func (o *InlineObject1) HasBalanceType() bool {
	if o != nil && o.BalanceType != nil {
		return true
	}

	return false
}

// SetBalanceType gets a reference to the given BalanceType and assigns it to the BalanceType field.
func (o *InlineObject1) SetBalanceType(v BalanceType) {
	o.BalanceType = &v
}

// GetPostingDate returns the PostingDate field value if set, zero value otherwise.
func (o *InlineObject1) GetPostingDate() time.Time {
	if o == nil || o.PostingDate == nil {
		var ret time.Time
		return ret
	}
	return *o.PostingDate
}

// GetPostingDateOk returns a tuple with the PostingDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject1) GetPostingDateOk() (*time.Time, bool) {
	if o == nil || o.PostingDate == nil {
		return nil, false
	}
	return o.PostingDate, true
}

// HasPostingDate returns a boolean if a field has been set.
func (o *InlineObject1) HasPostingDate() bool {
	if o != nil && o.PostingDate != nil {
		return true
	}

	return false
}

// SetPostingDate gets a reference to the given time.Time and assigns it to the PostingDate field.
func (o *InlineObject1) SetPostingDate(v time.Time) {
	o.PostingDate = &v
}

func (o InlineObject1) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.BalanceType != nil {
		toSerialize["balance_type"] = o.BalanceType
	}
	if o.PostingDate != nil {
		toSerialize["posting_date"] = o.PostingDate
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject1 struct {
	value *InlineObject1
	isSet bool
}

func (v NullableInlineObject1) Get() *InlineObject1 {
	return v.value
}

func (v *NullableInlineObject1) Set(val *InlineObject1) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject1) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject1(val *InlineObject1) *NullableInlineObject1 {
	return &NullableInlineObject1{value: val, isSet: true}
}

func (v NullableInlineObject1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


