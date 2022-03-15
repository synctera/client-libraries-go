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

// ManualAccountVerification struct for ManualAccountVerification
type ManualAccountVerification struct {
	// The time at which verification was first completed.
	CreationTime *time.Time `json:"creation_time,omitempty"`
	// The time at which verification was last updated.
	LastUpdatedTime *time.Time `json:"last_updated_time,omitempty"`
	// The status of verification
	Status string `json:"status"`
	// The vendor used for verifying the account
	Vendor string `json:"vendor"`
}

// NewManualAccountVerification instantiates a new ManualAccountVerification object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewManualAccountVerification(status string, vendor string) *ManualAccountVerification {
	this := ManualAccountVerification{}
	this.Status = status
	this.Vendor = vendor
	return &this
}

// NewManualAccountVerificationWithDefaults instantiates a new ManualAccountVerification object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewManualAccountVerificationWithDefaults() *ManualAccountVerification {
	this := ManualAccountVerification{}
	return &this
}

// GetCreationTime returns the CreationTime field value if set, zero value otherwise.
func (o *ManualAccountVerification) GetCreationTime() time.Time {
	if o == nil || o.CreationTime == nil {
		var ret time.Time
		return ret
	}
	return *o.CreationTime
}

// GetCreationTimeOk returns a tuple with the CreationTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManualAccountVerification) GetCreationTimeOk() (*time.Time, bool) {
	if o == nil || o.CreationTime == nil {
		return nil, false
	}
	return o.CreationTime, true
}

// HasCreationTime returns a boolean if a field has been set.
func (o *ManualAccountVerification) HasCreationTime() bool {
	if o != nil && o.CreationTime != nil {
		return true
	}

	return false
}

// SetCreationTime gets a reference to the given time.Time and assigns it to the CreationTime field.
func (o *ManualAccountVerification) SetCreationTime(v time.Time) {
	o.CreationTime = &v
}

// GetLastUpdatedTime returns the LastUpdatedTime field value if set, zero value otherwise.
func (o *ManualAccountVerification) GetLastUpdatedTime() time.Time {
	if o == nil || o.LastUpdatedTime == nil {
		var ret time.Time
		return ret
	}
	return *o.LastUpdatedTime
}

// GetLastUpdatedTimeOk returns a tuple with the LastUpdatedTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManualAccountVerification) GetLastUpdatedTimeOk() (*time.Time, bool) {
	if o == nil || o.LastUpdatedTime == nil {
		return nil, false
	}
	return o.LastUpdatedTime, true
}

// HasLastUpdatedTime returns a boolean if a field has been set.
func (o *ManualAccountVerification) HasLastUpdatedTime() bool {
	if o != nil && o.LastUpdatedTime != nil {
		return true
	}

	return false
}

// SetLastUpdatedTime gets a reference to the given time.Time and assigns it to the LastUpdatedTime field.
func (o *ManualAccountVerification) SetLastUpdatedTime(v time.Time) {
	o.LastUpdatedTime = &v
}

// GetStatus returns the Status field value
func (o *ManualAccountVerification) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *ManualAccountVerification) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *ManualAccountVerification) SetStatus(v string) {
	o.Status = v
}

// GetVendor returns the Vendor field value
func (o *ManualAccountVerification) GetVendor() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Vendor
}

// GetVendorOk returns a tuple with the Vendor field value
// and a boolean to check if the value has been set.
func (o *ManualAccountVerification) GetVendorOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Vendor, true
}

// SetVendor sets field value
func (o *ManualAccountVerification) SetVendor(v string) {
	o.Vendor = v
}

func (o ManualAccountVerification) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.CreationTime != nil {
		toSerialize["creation_time"] = o.CreationTime
	}
	if o.LastUpdatedTime != nil {
		toSerialize["last_updated_time"] = o.LastUpdatedTime
	}
	if true {
		toSerialize["status"] = o.Status
	}
	if true {
		toSerialize["vendor"] = o.Vendor
	}
	return json.Marshal(toSerialize)
}

type NullableManualAccountVerification struct {
	value *ManualAccountVerification
	isSet bool
}

func (v NullableManualAccountVerification) Get() *ManualAccountVerification {
	return v.value
}

func (v *NullableManualAccountVerification) Set(val *ManualAccountVerification) {
	v.value = val
	v.isSet = true
}

func (v NullableManualAccountVerification) IsSet() bool {
	return v.isSet
}

func (v *NullableManualAccountVerification) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableManualAccountVerification(val *ManualAccountVerification) *NullableManualAccountVerification {
	return &NullableManualAccountVerification{value: val, isSet: true}
}

func (v NullableManualAccountVerification) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableManualAccountVerification) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
