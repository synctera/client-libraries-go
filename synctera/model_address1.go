/*
Synctera API

<h2>Let's build something great.</h2><p>Welcome to the official reference documentation for Synctera APIs. Our APIs are the best way to automate your company's banking needs and are designed to be easy to understand and implement.</p><p>We're continuously growing this library and what you see here is just the start, but if you need something specific or have a question, <a class='text-blue-600' href='https://synctera.com/contact' target='_blank' rel='noreferrer'>contact us</a>.</p> 

API version: 0.5.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// Address1 struct for Address1
type Address1 struct {
	// Street address line 1
	AddressLine1 string `json:"address_line_1"`
	// String address line 2
	AddressLine2 *string `json:"address_line_2,omitempty"`
	// City
	City string `json:"city"`
	// ISO-3166-1 Alpha-2 country code
	CountryCode string `json:"country_code"`
	// Postal code
	PostalCode string `json:"postal_code"`
	// State, region, province, or prefecture
	State string `json:"state"`
}

// NewAddress1 instantiates a new Address1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddress1(addressLine1 string, city string, countryCode string, postalCode string, state string) *Address1 {
	this := Address1{}
	this.AddressLine1 = addressLine1
	this.City = city
	this.CountryCode = countryCode
	this.PostalCode = postalCode
	this.State = state
	return &this
}

// NewAddress1WithDefaults instantiates a new Address1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddress1WithDefaults() *Address1 {
	this := Address1{}
	return &this
}

// GetAddressLine1 returns the AddressLine1 field value
func (o *Address1) GetAddressLine1() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AddressLine1
}

// GetAddressLine1Ok returns a tuple with the AddressLine1 field value
// and a boolean to check if the value has been set.
func (o *Address1) GetAddressLine1Ok() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.AddressLine1, true
}

// SetAddressLine1 sets field value
func (o *Address1) SetAddressLine1(v string) {
	o.AddressLine1 = v
}

// GetAddressLine2 returns the AddressLine2 field value if set, zero value otherwise.
func (o *Address1) GetAddressLine2() string {
	if o == nil || o.AddressLine2 == nil {
		var ret string
		return ret
	}
	return *o.AddressLine2
}

// GetAddressLine2Ok returns a tuple with the AddressLine2 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Address1) GetAddressLine2Ok() (*string, bool) {
	if o == nil || o.AddressLine2 == nil {
		return nil, false
	}
	return o.AddressLine2, true
}

// HasAddressLine2 returns a boolean if a field has been set.
func (o *Address1) HasAddressLine2() bool {
	if o != nil && o.AddressLine2 != nil {
		return true
	}

	return false
}

// SetAddressLine2 gets a reference to the given string and assigns it to the AddressLine2 field.
func (o *Address1) SetAddressLine2(v string) {
	o.AddressLine2 = &v
}

// GetCity returns the City field value
func (o *Address1) GetCity() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.City
}

// GetCityOk returns a tuple with the City field value
// and a boolean to check if the value has been set.
func (o *Address1) GetCityOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.City, true
}

// SetCity sets field value
func (o *Address1) SetCity(v string) {
	o.City = v
}

// GetCountryCode returns the CountryCode field value
func (o *Address1) GetCountryCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CountryCode
}

// GetCountryCodeOk returns a tuple with the CountryCode field value
// and a boolean to check if the value has been set.
func (o *Address1) GetCountryCodeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.CountryCode, true
}

// SetCountryCode sets field value
func (o *Address1) SetCountryCode(v string) {
	o.CountryCode = v
}

// GetPostalCode returns the PostalCode field value
func (o *Address1) GetPostalCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PostalCode
}

// GetPostalCodeOk returns a tuple with the PostalCode field value
// and a boolean to check if the value has been set.
func (o *Address1) GetPostalCodeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.PostalCode, true
}

// SetPostalCode sets field value
func (o *Address1) SetPostalCode(v string) {
	o.PostalCode = v
}

// GetState returns the State field value
func (o *Address1) GetState() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.State
}

// GetStateOk returns a tuple with the State field value
// and a boolean to check if the value has been set.
func (o *Address1) GetStateOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.State, true
}

// SetState sets field value
func (o *Address1) SetState(v string) {
	o.State = v
}

func (o Address1) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["address_line_1"] = o.AddressLine1
	}
	if o.AddressLine2 != nil {
		toSerialize["address_line_2"] = o.AddressLine2
	}
	if true {
		toSerialize["city"] = o.City
	}
	if true {
		toSerialize["country_code"] = o.CountryCode
	}
	if true {
		toSerialize["postal_code"] = o.PostalCode
	}
	if true {
		toSerialize["state"] = o.State
	}
	return json.Marshal(toSerialize)
}

type NullableAddress1 struct {
	value *Address1
	isSet bool
}

func (v NullableAddress1) Get() *Address1 {
	return v.value
}

func (v *NullableAddress1) Set(val *Address1) {
	v.value = val
	v.isSet = true
}

func (v NullableAddress1) IsSet() bool {
	return v.isSet
}

func (v *NullableAddress1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddress1(val *Address1) *NullableAddress1 {
	return &NullableAddress1{value: val, isSet: true}
}

func (v NullableAddress1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddress1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


