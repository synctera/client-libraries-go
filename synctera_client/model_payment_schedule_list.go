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

// PaymentScheduleList struct for PaymentScheduleList
type PaymentScheduleList struct {
	// Array of payment schedules.
	PaymentSchedules []PaymentSchedule `json:"payment_schedules"`
	// If returned, use the next_page_token to query for the next page of results. Not returned if there are no more rows.
	NextPageToken *string `json:"next_page_token,omitempty"`
}

// NewPaymentScheduleList instantiates a new PaymentScheduleList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaymentScheduleList(paymentSchedules []PaymentSchedule) *PaymentScheduleList {
	this := PaymentScheduleList{}
	this.PaymentSchedules = paymentSchedules
	return &this
}

// NewPaymentScheduleListWithDefaults instantiates a new PaymentScheduleList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaymentScheduleListWithDefaults() *PaymentScheduleList {
	this := PaymentScheduleList{}
	return &this
}

// GetPaymentSchedules returns the PaymentSchedules field value
func (o *PaymentScheduleList) GetPaymentSchedules() []PaymentSchedule {
	if o == nil {
		var ret []PaymentSchedule
		return ret
	}

	return o.PaymentSchedules
}

// GetPaymentSchedulesOk returns a tuple with the PaymentSchedules field value
// and a boolean to check if the value has been set.
func (o *PaymentScheduleList) GetPaymentSchedulesOk() ([]PaymentSchedule, bool) {
	if o == nil {
		return nil, false
	}
	return o.PaymentSchedules, true
}

// SetPaymentSchedules sets field value
func (o *PaymentScheduleList) SetPaymentSchedules(v []PaymentSchedule) {
	o.PaymentSchedules = v
}

// GetNextPageToken returns the NextPageToken field value if set, zero value otherwise.
func (o *PaymentScheduleList) GetNextPageToken() string {
	if o == nil || o.NextPageToken == nil {
		var ret string
		return ret
	}
	return *o.NextPageToken
}

// GetNextPageTokenOk returns a tuple with the NextPageToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentScheduleList) GetNextPageTokenOk() (*string, bool) {
	if o == nil || o.NextPageToken == nil {
		return nil, false
	}
	return o.NextPageToken, true
}

// HasNextPageToken returns a boolean if a field has been set.
func (o *PaymentScheduleList) HasNextPageToken() bool {
	if o != nil && o.NextPageToken != nil {
		return true
	}

	return false
}

// SetNextPageToken gets a reference to the given string and assigns it to the NextPageToken field.
func (o *PaymentScheduleList) SetNextPageToken(v string) {
	o.NextPageToken = &v
}

func (o PaymentScheduleList) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["payment_schedules"] = o.PaymentSchedules
	}
	if o.NextPageToken != nil {
		toSerialize["next_page_token"] = o.NextPageToken
	}
	return json.Marshal(toSerialize)
}

type NullablePaymentScheduleList struct {
	value *PaymentScheduleList
	isSet bool
}

func (v NullablePaymentScheduleList) Get() *PaymentScheduleList {
	return v.value
}

func (v *NullablePaymentScheduleList) Set(val *PaymentScheduleList) {
	v.value = val
	v.isSet = true
}

func (v NullablePaymentScheduleList) IsSet() bool {
	return v.isSet
}

func (v *NullablePaymentScheduleList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaymentScheduleList(val *PaymentScheduleList) *NullablePaymentScheduleList {
	return &NullablePaymentScheduleList{value: val, isSet: true}
}

func (v NullablePaymentScheduleList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaymentScheduleList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
