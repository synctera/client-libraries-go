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

// CardChange Card change details
type CardChange struct {
	ChangeType ChangeType `json:"change_type"`
	Channel ChangeChannel `json:"channel"`
	// Unique token
	Id string `json:"id"`
	// Additional details about the reason for the status change
	Memo *string `json:"memo,omitempty"`
	Reason *CardStatusReasonCode `json:"reason,omitempty"`
	StateNew CardChangeState `json:"state_new"`
	StateOld CardChangeState `json:"state_old"`
	// Date of change
	UpdatedAt time.Time `json:"updated_at"`
	// ID of user who initiated the change, if done via Synctera Admin System
	UpdatedBy string `json:"updated_by"`
}

// NewCardChange instantiates a new CardChange object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCardChange(changeType ChangeType, channel ChangeChannel, id string, stateNew CardChangeState, stateOld CardChangeState, updatedAt time.Time, updatedBy string) *CardChange {
	this := CardChange{}
	this.ChangeType = changeType
	this.Channel = channel
	this.Id = id
	this.StateNew = stateNew
	this.StateOld = stateOld
	this.UpdatedAt = updatedAt
	this.UpdatedBy = updatedBy
	return &this
}

// NewCardChangeWithDefaults instantiates a new CardChange object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCardChangeWithDefaults() *CardChange {
	this := CardChange{}
	return &this
}

// GetChangeType returns the ChangeType field value
func (o *CardChange) GetChangeType() ChangeType {
	if o == nil {
		var ret ChangeType
		return ret
	}

	return o.ChangeType
}

// GetChangeTypeOk returns a tuple with the ChangeType field value
// and a boolean to check if the value has been set.
func (o *CardChange) GetChangeTypeOk() (*ChangeType, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ChangeType, true
}

// SetChangeType sets field value
func (o *CardChange) SetChangeType(v ChangeType) {
	o.ChangeType = v
}

// GetChannel returns the Channel field value
func (o *CardChange) GetChannel() ChangeChannel {
	if o == nil {
		var ret ChangeChannel
		return ret
	}

	return o.Channel
}

// GetChannelOk returns a tuple with the Channel field value
// and a boolean to check if the value has been set.
func (o *CardChange) GetChannelOk() (*ChangeChannel, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Channel, true
}

// SetChannel sets field value
func (o *CardChange) SetChannel(v ChangeChannel) {
	o.Channel = v
}

// GetId returns the Id field value
func (o *CardChange) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *CardChange) GetIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *CardChange) SetId(v string) {
	o.Id = v
}

// GetMemo returns the Memo field value if set, zero value otherwise.
func (o *CardChange) GetMemo() string {
	if o == nil || o.Memo == nil {
		var ret string
		return ret
	}
	return *o.Memo
}

// GetMemoOk returns a tuple with the Memo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CardChange) GetMemoOk() (*string, bool) {
	if o == nil || o.Memo == nil {
		return nil, false
	}
	return o.Memo, true
}

// HasMemo returns a boolean if a field has been set.
func (o *CardChange) HasMemo() bool {
	if o != nil && o.Memo != nil {
		return true
	}

	return false
}

// SetMemo gets a reference to the given string and assigns it to the Memo field.
func (o *CardChange) SetMemo(v string) {
	o.Memo = &v
}

// GetReason returns the Reason field value if set, zero value otherwise.
func (o *CardChange) GetReason() CardStatusReasonCode {
	if o == nil || o.Reason == nil {
		var ret CardStatusReasonCode
		return ret
	}
	return *o.Reason
}

// GetReasonOk returns a tuple with the Reason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CardChange) GetReasonOk() (*CardStatusReasonCode, bool) {
	if o == nil || o.Reason == nil {
		return nil, false
	}
	return o.Reason, true
}

// HasReason returns a boolean if a field has been set.
func (o *CardChange) HasReason() bool {
	if o != nil && o.Reason != nil {
		return true
	}

	return false
}

// SetReason gets a reference to the given CardStatusReasonCode and assigns it to the Reason field.
func (o *CardChange) SetReason(v CardStatusReasonCode) {
	o.Reason = &v
}

// GetStateNew returns the StateNew field value
func (o *CardChange) GetStateNew() CardChangeState {
	if o == nil {
		var ret CardChangeState
		return ret
	}

	return o.StateNew
}

// GetStateNewOk returns a tuple with the StateNew field value
// and a boolean to check if the value has been set.
func (o *CardChange) GetStateNewOk() (*CardChangeState, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.StateNew, true
}

// SetStateNew sets field value
func (o *CardChange) SetStateNew(v CardChangeState) {
	o.StateNew = v
}

// GetStateOld returns the StateOld field value
func (o *CardChange) GetStateOld() CardChangeState {
	if o == nil {
		var ret CardChangeState
		return ret
	}

	return o.StateOld
}

// GetStateOldOk returns a tuple with the StateOld field value
// and a boolean to check if the value has been set.
func (o *CardChange) GetStateOldOk() (*CardChangeState, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.StateOld, true
}

// SetStateOld sets field value
func (o *CardChange) SetStateOld(v CardChangeState) {
	o.StateOld = v
}

// GetUpdatedAt returns the UpdatedAt field value
func (o *CardChange) GetUpdatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value
// and a boolean to check if the value has been set.
func (o *CardChange) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// SetUpdatedAt sets field value
func (o *CardChange) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = v
}

// GetUpdatedBy returns the UpdatedBy field value
func (o *CardChange) GetUpdatedBy() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UpdatedBy
}

// GetUpdatedByOk returns a tuple with the UpdatedBy field value
// and a boolean to check if the value has been set.
func (o *CardChange) GetUpdatedByOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.UpdatedBy, true
}

// SetUpdatedBy sets field value
func (o *CardChange) SetUpdatedBy(v string) {
	o.UpdatedBy = v
}

func (o CardChange) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["change_type"] = o.ChangeType
	}
	if true {
		toSerialize["channel"] = o.Channel
	}
	if true {
		toSerialize["id"] = o.Id
	}
	if o.Memo != nil {
		toSerialize["memo"] = o.Memo
	}
	if o.Reason != nil {
		toSerialize["reason"] = o.Reason
	}
	if true {
		toSerialize["state_new"] = o.StateNew
	}
	if true {
		toSerialize["state_old"] = o.StateOld
	}
	if true {
		toSerialize["updated_at"] = o.UpdatedAt
	}
	if true {
		toSerialize["updated_by"] = o.UpdatedBy
	}
	return json.Marshal(toSerialize)
}

type NullableCardChange struct {
	value *CardChange
	isSet bool
}

func (v NullableCardChange) Get() *CardChange {
	return v.value
}

func (v *NullableCardChange) Set(val *CardChange) {
	v.value = val
	v.isSet = true
}

func (v NullableCardChange) IsSet() bool {
	return v.isSet
}

func (v *NullableCardChange) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCardChange(val *CardChange) *NullableCardChange {
	return &NullableCardChange{value: val, isSet: true}
}

func (v NullableCardChange) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCardChange) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


