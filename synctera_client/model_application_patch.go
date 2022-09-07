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

// ApplicationPatch struct for ApplicationPatch
type ApplicationPatch struct {
	// Details about the applicant. The exact schema is to be determined with your bank.
	ApplicationDetails map[string]interface{} `json:"application_details,omitempty"`
	Status *ApplicationStatus `json:"status,omitempty"`
}

// NewApplicationPatch instantiates a new ApplicationPatch object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApplicationPatch() *ApplicationPatch {
	this := ApplicationPatch{}
	return &this
}

// NewApplicationPatchWithDefaults instantiates a new ApplicationPatch object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApplicationPatchWithDefaults() *ApplicationPatch {
	this := ApplicationPatch{}
	return &this
}

// GetApplicationDetails returns the ApplicationDetails field value if set, zero value otherwise.
func (o *ApplicationPatch) GetApplicationDetails() map[string]interface{} {
	if o == nil || o.ApplicationDetails == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.ApplicationDetails
}

// GetApplicationDetailsOk returns a tuple with the ApplicationDetails field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationPatch) GetApplicationDetailsOk() (map[string]interface{}, bool) {
	if o == nil || o.ApplicationDetails == nil {
		return nil, false
	}
	return o.ApplicationDetails, true
}

// HasApplicationDetails returns a boolean if a field has been set.
func (o *ApplicationPatch) HasApplicationDetails() bool {
	if o != nil && o.ApplicationDetails != nil {
		return true
	}

	return false
}

// SetApplicationDetails gets a reference to the given map[string]interface{} and assigns it to the ApplicationDetails field.
func (o *ApplicationPatch) SetApplicationDetails(v map[string]interface{}) {
	o.ApplicationDetails = v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *ApplicationPatch) GetStatus() ApplicationStatus {
	if o == nil || o.Status == nil {
		var ret ApplicationStatus
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationPatch) GetStatusOk() (*ApplicationStatus, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *ApplicationPatch) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given ApplicationStatus and assigns it to the Status field.
func (o *ApplicationPatch) SetStatus(v ApplicationStatus) {
	o.Status = &v
}

func (o ApplicationPatch) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ApplicationDetails != nil {
		toSerialize["application_details"] = o.ApplicationDetails
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	return json.Marshal(toSerialize)
}

type NullableApplicationPatch struct {
	value *ApplicationPatch
	isSet bool
}

func (v NullableApplicationPatch) Get() *ApplicationPatch {
	return v.value
}

func (v *NullableApplicationPatch) Set(val *ApplicationPatch) {
	v.value = val
	v.isSet = true
}

func (v NullableApplicationPatch) IsSet() bool {
	return v.isSet
}

func (v *NullableApplicationPatch) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApplicationPatch(val *ApplicationPatch) *NullableApplicationPatch {
	return &NullableApplicationPatch{value: val, isSet: true}
}

func (v NullableApplicationPatch) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApplicationPatch) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

