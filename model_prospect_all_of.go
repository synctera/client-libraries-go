/*
 * Synctera API
 *
 * <h2>Let's build something great.</h2><p>Welcome to the official reference documentation for Synctera APIs. Our APIs are the best way to automate your company's banking needs and are designed to be easy to understand and implement.</p><p>We're continuously growing this library and what you see here is just the start, but if you need something specific or have a question, <a class='text-blue-600' href='https://synctera.com/contact' target='_blank' rel='noreferrer'>contact us</a>.</p>
 *
 * API version: 0.5.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package synctera

import (
	"encoding/json"

	oapi "github.com/deepmap/oapi-codegen/pkg/types"
)

// ProspectAllOf struct for ProspectAllOf
type ProspectAllOf struct {
	// Customer's date of birth in RFC 3339 full-date format (YYYY-MM-DD)
	Dob *oapi.Date `json:"dob,omitempty"`
	// Customer's first name
	FirstName *string `json:"first_name,omitempty"`
	// Customer's last name
	LastName *string `json:"last_name,omitempty"`
	// Customer's status
	Status string `json:"status"`
}

// NewProspectAllOf instantiates a new ProspectAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProspectAllOf(status string) *ProspectAllOf {
	this := ProspectAllOf{}
	this.Status = status
	return &this
}

// NewProspectAllOfWithDefaults instantiates a new ProspectAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProspectAllOfWithDefaults() *ProspectAllOf {
	this := ProspectAllOf{}
	return &this
}

// GetDob returns the Dob field value if set, zero value otherwise.
func (o *ProspectAllOf) GetDob() oapi.Date {
	if o == nil || o.Dob == nil {
		var ret oapi.Date
		return ret
	}
	return *o.Dob
}

// GetDobOk returns a tuple with the Dob field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProspectAllOf) GetDobOk() (*oapi.Date, bool) {
	if o == nil || o.Dob == nil {
		return nil, false
	}
	return o.Dob, true
}

// HasDob returns a boolean if a field has been set.
func (o *ProspectAllOf) HasDob() bool {
	if o != nil && o.Dob != nil {
		return true
	}

	return false
}

// SetDob gets a reference to the given oapi.Date and assigns it to the Dob field.
func (o *ProspectAllOf) SetDob(v oapi.Date) {
	o.Dob = &v
}

// GetFirstName returns the FirstName field value if set, zero value otherwise.
func (o *ProspectAllOf) GetFirstName() string {
	if o == nil || o.FirstName == nil {
		var ret string
		return ret
	}
	return *o.FirstName
}

// GetFirstNameOk returns a tuple with the FirstName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProspectAllOf) GetFirstNameOk() (*string, bool) {
	if o == nil || o.FirstName == nil {
		return nil, false
	}
	return o.FirstName, true
}

// HasFirstName returns a boolean if a field has been set.
func (o *ProspectAllOf) HasFirstName() bool {
	if o != nil && o.FirstName != nil {
		return true
	}

	return false
}

// SetFirstName gets a reference to the given string and assigns it to the FirstName field.
func (o *ProspectAllOf) SetFirstName(v string) {
	o.FirstName = &v
}

// GetLastName returns the LastName field value if set, zero value otherwise.
func (o *ProspectAllOf) GetLastName() string {
	if o == nil || o.LastName == nil {
		var ret string
		return ret
	}
	return *o.LastName
}

// GetLastNameOk returns a tuple with the LastName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProspectAllOf) GetLastNameOk() (*string, bool) {
	if o == nil || o.LastName == nil {
		return nil, false
	}
	return o.LastName, true
}

// HasLastName returns a boolean if a field has been set.
func (o *ProspectAllOf) HasLastName() bool {
	if o != nil && o.LastName != nil {
		return true
	}

	return false
}

// SetLastName gets a reference to the given string and assigns it to the LastName field.
func (o *ProspectAllOf) SetLastName(v string) {
	o.LastName = &v
}

// GetStatus returns the Status field value
func (o *ProspectAllOf) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *ProspectAllOf) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *ProspectAllOf) SetStatus(v string) {
	o.Status = v
}

func (o ProspectAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Dob != nil {
		toSerialize["dob"] = o.Dob
	}
	if o.FirstName != nil {
		toSerialize["first_name"] = o.FirstName
	}
	if o.LastName != nil {
		toSerialize["last_name"] = o.LastName
	}
	if true {
		toSerialize["status"] = o.Status
	}
	return json.Marshal(toSerialize)
}

type NullableProspectAllOf struct {
	value *ProspectAllOf
	isSet bool
}

func (v NullableProspectAllOf) Get() *ProspectAllOf {
	return v.value
}

func (v *NullableProspectAllOf) Set(val *ProspectAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableProspectAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableProspectAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProspectAllOf(val *ProspectAllOf) *NullableProspectAllOf {
	return &NullableProspectAllOf{value: val, isSet: true}
}

func (v NullableProspectAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProspectAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
