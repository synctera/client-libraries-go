/*
Synctera API

<h2>Let's build something great.</h2><p>Welcome to the official reference documentation for Synctera APIs. Our APIs are the best way to automate your company's banking needs and are designed to be easy to understand and implement.</p><p>We're continuously growing this library and what you see here is just the start, but if you need something specific or have a question, <a class='text-blue-600' href='https://synctera.com/contact' target='_blank' rel='noreferrer'>contact us</a>.</p>

API version: 0.5.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package synctera

import (
	"encoding/json"

	oapi "github.com/deepmap/oapi-codegen/pkg/types"
)

// CustomerAllOf struct for CustomerAllOf
type CustomerAllOf struct {
	// Customer's date of birth in RFC 3339 full-date format (YYYY-MM-DD)
	Dob oapi.Date `json:"dob"`
	// Customer's first name
	FirstName string `json:"first_name"`
	// Customer's last name
	LastName string `json:"last_name"`
	// Customer's status
	Status string `json:"status"`
}

// NewCustomerAllOf instantiates a new CustomerAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCustomerAllOf(dob oapi.Date, firstName string, lastName string, status string) *CustomerAllOf {
	this := CustomerAllOf{}
	this.Dob = dob
	this.FirstName = firstName
	this.LastName = lastName
	this.Status = status
	return &this
}

// NewCustomerAllOfWithDefaults instantiates a new CustomerAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCustomerAllOfWithDefaults() *CustomerAllOf {
	this := CustomerAllOf{}
	return &this
}

// GetDob returns the Dob field value
func (o *CustomerAllOf) GetDob() oapi.Date {
	if o == nil {
		var ret oapi.Date
		return ret
	}

	return o.Dob
}

// GetDobOk returns a tuple with the Dob field value
// and a boolean to check if the value has been set.
func (o *CustomerAllOf) GetDobOk() (*oapi.Date, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Dob, true
}

// SetDob sets field value
func (o *CustomerAllOf) SetDob(v oapi.Date) {
	o.Dob = v
}

// GetFirstName returns the FirstName field value
func (o *CustomerAllOf) GetFirstName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FirstName
}

// GetFirstNameOk returns a tuple with the FirstName field value
// and a boolean to check if the value has been set.
func (o *CustomerAllOf) GetFirstNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FirstName, true
}

// SetFirstName sets field value
func (o *CustomerAllOf) SetFirstName(v string) {
	o.FirstName = v
}

// GetLastName returns the LastName field value
func (o *CustomerAllOf) GetLastName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.LastName
}

// GetLastNameOk returns a tuple with the LastName field value
// and a boolean to check if the value has been set.
func (o *CustomerAllOf) GetLastNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LastName, true
}

// SetLastName sets field value
func (o *CustomerAllOf) SetLastName(v string) {
	o.LastName = v
}

// GetStatus returns the Status field value
func (o *CustomerAllOf) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *CustomerAllOf) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *CustomerAllOf) SetStatus(v string) {
	o.Status = v
}

func (o CustomerAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["dob"] = o.Dob
	}
	if true {
		toSerialize["first_name"] = o.FirstName
	}
	if true {
		toSerialize["last_name"] = o.LastName
	}
	if true {
		toSerialize["status"] = o.Status
	}
	return json.Marshal(toSerialize)
}

type NullableCustomerAllOf struct {
	value *CustomerAllOf
	isSet bool
}

func (v NullableCustomerAllOf) Get() *CustomerAllOf {
	return v.value
}

func (v *NullableCustomerAllOf) Set(val *CustomerAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableCustomerAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableCustomerAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCustomerAllOf(val *CustomerAllOf) *NullableCustomerAllOf {
	return &NullableCustomerAllOf{value: val, isSet: true}
}

func (v NullableCustomerAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCustomerAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
