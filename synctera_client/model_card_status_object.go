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

// CardStatusObject The status of the card
type CardStatusObject struct {
	CardStatus CardStatus `json:"card_status"`
	// Additional details about the reason for the status change
	Memo         *string               `json:"memo,omitempty"`
	StatusReason *CardStatusReasonCode `json:"status_reason,omitempty"`
}

// NewCardStatusObject instantiates a new CardStatusObject object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCardStatusObject(cardStatus CardStatus) *CardStatusObject {
	this := CardStatusObject{}
	this.CardStatus = cardStatus
	return &this
}

// NewCardStatusObjectWithDefaults instantiates a new CardStatusObject object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCardStatusObjectWithDefaults() *CardStatusObject {
	this := CardStatusObject{}
	return &this
}

// GetCardStatus returns the CardStatus field value
func (o *CardStatusObject) GetCardStatus() CardStatus {
	if o == nil {
		var ret CardStatus
		return ret
	}

	return o.CardStatus
}

// GetCardStatusOk returns a tuple with the CardStatus field value
// and a boolean to check if the value has been set.
func (o *CardStatusObject) GetCardStatusOk() (*CardStatus, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CardStatus, true
}

// SetCardStatus sets field value
func (o *CardStatusObject) SetCardStatus(v CardStatus) {
	o.CardStatus = v
}

// GetMemo returns the Memo field value if set, zero value otherwise.
func (o *CardStatusObject) GetMemo() string {
	if o == nil || o.Memo == nil {
		var ret string
		return ret
	}
	return *o.Memo
}

// GetMemoOk returns a tuple with the Memo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CardStatusObject) GetMemoOk() (*string, bool) {
	if o == nil || o.Memo == nil {
		return nil, false
	}
	return o.Memo, true
}

// HasMemo returns a boolean if a field has been set.
func (o *CardStatusObject) HasMemo() bool {
	if o != nil && o.Memo != nil {
		return true
	}

	return false
}

// SetMemo gets a reference to the given string and assigns it to the Memo field.
func (o *CardStatusObject) SetMemo(v string) {
	o.Memo = &v
}

// GetStatusReason returns the StatusReason field value if set, zero value otherwise.
func (o *CardStatusObject) GetStatusReason() CardStatusReasonCode {
	if o == nil || o.StatusReason == nil {
		var ret CardStatusReasonCode
		return ret
	}
	return *o.StatusReason
}

// GetStatusReasonOk returns a tuple with the StatusReason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CardStatusObject) GetStatusReasonOk() (*CardStatusReasonCode, bool) {
	if o == nil || o.StatusReason == nil {
		return nil, false
	}
	return o.StatusReason, true
}

// HasStatusReason returns a boolean if a field has been set.
func (o *CardStatusObject) HasStatusReason() bool {
	if o != nil && o.StatusReason != nil {
		return true
	}

	return false
}

// SetStatusReason gets a reference to the given CardStatusReasonCode and assigns it to the StatusReason field.
func (o *CardStatusObject) SetStatusReason(v CardStatusReasonCode) {
	o.StatusReason = &v
}

func (o CardStatusObject) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["card_status"] = o.CardStatus
	}
	if o.Memo != nil {
		toSerialize["memo"] = o.Memo
	}
	if o.StatusReason != nil {
		toSerialize["status_reason"] = o.StatusReason
	}
	return json.Marshal(toSerialize)
}

type NullableCardStatusObject struct {
	value *CardStatusObject
	isSet bool
}

func (v NullableCardStatusObject) Get() *CardStatusObject {
	return v.value
}

func (v *NullableCardStatusObject) Set(val *CardStatusObject) {
	v.value = val
	v.isSet = true
}

func (v NullableCardStatusObject) IsSet() bool {
	return v.isSet
}

func (v *NullableCardStatusObject) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCardStatusObject(val *CardStatusObject) *NullableCardStatusObject {
	return &NullableCardStatusObject{value: val, isSet: true}
}

func (v NullableCardStatusObject) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCardStatusObject) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
