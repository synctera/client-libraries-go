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

// Address2 struct for Address2
type Address2 struct {
	AddressLine1 *string `json:"address_line_1,omitempty"`
	AddressLine2 *string `json:"address_line_2,omitempty"`
	City *string `json:"city,omitempty"`
	CountryCode *string `json:"country_code,omitempty"`
	PostalCode *string `json:"postal_code,omitempty"`
	State *string `json:"state,omitempty"`
}

// NewAddress2 instantiates a new Address2 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddress2() *Address2 {
	this := Address2{}
	return &this
}

// NewAddress2WithDefaults instantiates a new Address2 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddress2WithDefaults() *Address2 {
	this := Address2{}
	return &this
}

// GetAddressLine1 returns the AddressLine1 field value if set, zero value otherwise.
func (o *Address2) GetAddressLine1() string {
	if o == nil || o.AddressLine1 == nil {
		var ret string
		return ret
	}
	return *o.AddressLine1
}

// GetAddressLine1Ok returns a tuple with the AddressLine1 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Address2) GetAddressLine1Ok() (*string, bool) {
	if o == nil || o.AddressLine1 == nil {
		return nil, false
	}
	return o.AddressLine1, true
}

// HasAddressLine1 returns a boolean if a field has been set.
func (o *Address2) HasAddressLine1() bool {
	if o != nil && o.AddressLine1 != nil {
		return true
	}

	return false
}

// SetAddressLine1 gets a reference to the given string and assigns it to the AddressLine1 field.
func (o *Address2) SetAddressLine1(v string) {
	o.AddressLine1 = &v
}

// GetAddressLine2 returns the AddressLine2 field value if set, zero value otherwise.
func (o *Address2) GetAddressLine2() string {
	if o == nil || o.AddressLine2 == nil {
		var ret string
		return ret
	}
	return *o.AddressLine2
}

// GetAddressLine2Ok returns a tuple with the AddressLine2 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Address2) GetAddressLine2Ok() (*string, bool) {
	if o == nil || o.AddressLine2 == nil {
		return nil, false
	}
	return o.AddressLine2, true
}

// HasAddressLine2 returns a boolean if a field has been set.
func (o *Address2) HasAddressLine2() bool {
	if o != nil && o.AddressLine2 != nil {
		return true
	}

	return false
}

// SetAddressLine2 gets a reference to the given string and assigns it to the AddressLine2 field.
func (o *Address2) SetAddressLine2(v string) {
	o.AddressLine2 = &v
}

// GetCity returns the City field value if set, zero value otherwise.
func (o *Address2) GetCity() string {
	if o == nil || o.City == nil {
		var ret string
		return ret
	}
	return *o.City
}

// GetCityOk returns a tuple with the City field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Address2) GetCityOk() (*string, bool) {
	if o == nil || o.City == nil {
		return nil, false
	}
	return o.City, true
}

// HasCity returns a boolean if a field has been set.
func (o *Address2) HasCity() bool {
	if o != nil && o.City != nil {
		return true
	}

	return false
}

// SetCity gets a reference to the given string and assigns it to the City field.
func (o *Address2) SetCity(v string) {
	o.City = &v
}

// GetCountryCode returns the CountryCode field value if set, zero value otherwise.
func (o *Address2) GetCountryCode() string {
	if o == nil || o.CountryCode == nil {
		var ret string
		return ret
	}
	return *o.CountryCode
}

// GetCountryCodeOk returns a tuple with the CountryCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Address2) GetCountryCodeOk() (*string, bool) {
	if o == nil || o.CountryCode == nil {
		return nil, false
	}
	return o.CountryCode, true
}

// HasCountryCode returns a boolean if a field has been set.
func (o *Address2) HasCountryCode() bool {
	if o != nil && o.CountryCode != nil {
		return true
	}

	return false
}

// SetCountryCode gets a reference to the given string and assigns it to the CountryCode field.
func (o *Address2) SetCountryCode(v string) {
	o.CountryCode = &v
}

// GetPostalCode returns the PostalCode field value if set, zero value otherwise.
func (o *Address2) GetPostalCode() string {
	if o == nil || o.PostalCode == nil {
		var ret string
		return ret
	}
	return *o.PostalCode
}

// GetPostalCodeOk returns a tuple with the PostalCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Address2) GetPostalCodeOk() (*string, bool) {
	if o == nil || o.PostalCode == nil {
		return nil, false
	}
	return o.PostalCode, true
}

// HasPostalCode returns a boolean if a field has been set.
func (o *Address2) HasPostalCode() bool {
	if o != nil && o.PostalCode != nil {
		return true
	}

	return false
}

// SetPostalCode gets a reference to the given string and assigns it to the PostalCode field.
func (o *Address2) SetPostalCode(v string) {
	o.PostalCode = &v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *Address2) GetState() string {
	if o == nil || o.State == nil {
		var ret string
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Address2) GetStateOk() (*string, bool) {
	if o == nil || o.State == nil {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *Address2) HasState() bool {
	if o != nil && o.State != nil {
		return true
	}

	return false
}

// SetState gets a reference to the given string and assigns it to the State field.
func (o *Address2) SetState(v string) {
	o.State = &v
}

func (o Address2) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AddressLine1 != nil {
		toSerialize["address_line_1"] = o.AddressLine1
	}
	if o.AddressLine2 != nil {
		toSerialize["address_line_2"] = o.AddressLine2
	}
	if o.City != nil {
		toSerialize["city"] = o.City
	}
	if o.CountryCode != nil {
		toSerialize["country_code"] = o.CountryCode
	}
	if o.PostalCode != nil {
		toSerialize["postal_code"] = o.PostalCode
	}
	if o.State != nil {
		toSerialize["state"] = o.State
	}
	return json.Marshal(toSerialize)
}

type NullableAddress2 struct {
	value *Address2
	isSet bool
}

func (v NullableAddress2) Get() *Address2 {
	return v.value
}

func (v *NullableAddress2) Set(val *Address2) {
	v.value = val
	v.isSet = true
}

func (v NullableAddress2) IsSet() bool {
	return v.isSet
}

func (v *NullableAddress2) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddress2(val *Address2) *NullableAddress2 {
	return &NullableAddress2{value: val, isSet: true}
}

func (v NullableAddress2) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddress2) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


