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

// WaitlistEditable Waitlist fields that can be specified in a PATCH operation
type WaitlistEditable struct {
	// Waitlist Name
	Description *string `json:"description,omitempty"`
	// Number of points earned per verified referral
	Increment *int32 `json:"increment,omitempty"`
	// Max number of prospects allowed on this waitlist. Default is 10,000,000.
	MaxProspects *int32 `json:"max_prospects,omitempty"`
	// Waitlist Name
	WaitlistName *string `json:"waitlist_name,omitempty"`
}

// NewWaitlistEditable instantiates a new WaitlistEditable object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWaitlistEditable() *WaitlistEditable {
	this := WaitlistEditable{}
	return &this
}

// NewWaitlistEditableWithDefaults instantiates a new WaitlistEditable object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWaitlistEditableWithDefaults() *WaitlistEditable {
	this := WaitlistEditable{}
	return &this
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *WaitlistEditable) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WaitlistEditable) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *WaitlistEditable) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *WaitlistEditable) SetDescription(v string) {
	o.Description = &v
}

// GetIncrement returns the Increment field value if set, zero value otherwise.
func (o *WaitlistEditable) GetIncrement() int32 {
	if o == nil || o.Increment == nil {
		var ret int32
		return ret
	}
	return *o.Increment
}

// GetIncrementOk returns a tuple with the Increment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WaitlistEditable) GetIncrementOk() (*int32, bool) {
	if o == nil || o.Increment == nil {
		return nil, false
	}
	return o.Increment, true
}

// HasIncrement returns a boolean if a field has been set.
func (o *WaitlistEditable) HasIncrement() bool {
	if o != nil && o.Increment != nil {
		return true
	}

	return false
}

// SetIncrement gets a reference to the given int32 and assigns it to the Increment field.
func (o *WaitlistEditable) SetIncrement(v int32) {
	o.Increment = &v
}

// GetMaxProspects returns the MaxProspects field value if set, zero value otherwise.
func (o *WaitlistEditable) GetMaxProspects() int32 {
	if o == nil || o.MaxProspects == nil {
		var ret int32
		return ret
	}
	return *o.MaxProspects
}

// GetMaxProspectsOk returns a tuple with the MaxProspects field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WaitlistEditable) GetMaxProspectsOk() (*int32, bool) {
	if o == nil || o.MaxProspects == nil {
		return nil, false
	}
	return o.MaxProspects, true
}

// HasMaxProspects returns a boolean if a field has been set.
func (o *WaitlistEditable) HasMaxProspects() bool {
	if o != nil && o.MaxProspects != nil {
		return true
	}

	return false
}

// SetMaxProspects gets a reference to the given int32 and assigns it to the MaxProspects field.
func (o *WaitlistEditable) SetMaxProspects(v int32) {
	o.MaxProspects = &v
}

// GetWaitlistName returns the WaitlistName field value if set, zero value otherwise.
func (o *WaitlistEditable) GetWaitlistName() string {
	if o == nil || o.WaitlistName == nil {
		var ret string
		return ret
	}
	return *o.WaitlistName
}

// GetWaitlistNameOk returns a tuple with the WaitlistName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WaitlistEditable) GetWaitlistNameOk() (*string, bool) {
	if o == nil || o.WaitlistName == nil {
		return nil, false
	}
	return o.WaitlistName, true
}

// HasWaitlistName returns a boolean if a field has been set.
func (o *WaitlistEditable) HasWaitlistName() bool {
	if o != nil && o.WaitlistName != nil {
		return true
	}

	return false
}

// SetWaitlistName gets a reference to the given string and assigns it to the WaitlistName field.
func (o *WaitlistEditable) SetWaitlistName(v string) {
	o.WaitlistName = &v
}

func (o WaitlistEditable) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.Increment != nil {
		toSerialize["increment"] = o.Increment
	}
	if o.MaxProspects != nil {
		toSerialize["max_prospects"] = o.MaxProspects
	}
	if o.WaitlistName != nil {
		toSerialize["waitlist_name"] = o.WaitlistName
	}
	return json.Marshal(toSerialize)
}

type NullableWaitlistEditable struct {
	value *WaitlistEditable
	isSet bool
}

func (v NullableWaitlistEditable) Get() *WaitlistEditable {
	return v.value
}

func (v *NullableWaitlistEditable) Set(val *WaitlistEditable) {
	v.value = val
	v.isSet = true
}

func (v NullableWaitlistEditable) IsSet() bool {
	return v.isSet
}

func (v *NullableWaitlistEditable) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWaitlistEditable(val *WaitlistEditable) *NullableWaitlistEditable {
	return &NullableWaitlistEditable{value: val, isSet: true}
}

func (v NullableWaitlistEditable) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWaitlistEditable) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

