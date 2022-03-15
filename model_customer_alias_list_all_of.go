/*
Synctera API

<h2>Let's build something great.</h2><p>Welcome to the official reference documentation for Synctera APIs. Our APIs are the best way to automate your company's banking needs and are designed to be easy to understand and implement.</p><p>We're continuously growing this library and what you see here is just the start, but if you need something specific or have a question, <a class='text-blue-600' href='https://synctera.com/contact' target='_blank' rel='noreferrer'>contact us</a>.</p>

API version: 0.5.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package synctera

import (
	"encoding/json"
)

// CustomerAliasListAllOf struct for CustomerAliasListAllOf
type CustomerAliasListAllOf struct {
	// Array of customer alias
	CustomerAlias []CustomerAlias `json:"customer_alias"`
}

// NewCustomerAliasListAllOf instantiates a new CustomerAliasListAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCustomerAliasListAllOf(customerAlias []CustomerAlias) *CustomerAliasListAllOf {
	this := CustomerAliasListAllOf{}
	this.CustomerAlias = customerAlias
	return &this
}

// NewCustomerAliasListAllOfWithDefaults instantiates a new CustomerAliasListAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCustomerAliasListAllOfWithDefaults() *CustomerAliasListAllOf {
	this := CustomerAliasListAllOf{}
	return &this
}

// GetCustomerAlias returns the CustomerAlias field value
func (o *CustomerAliasListAllOf) GetCustomerAlias() []CustomerAlias {
	if o == nil {
		var ret []CustomerAlias
		return ret
	}

	return o.CustomerAlias
}

// GetCustomerAliasOk returns a tuple with the CustomerAlias field value
// and a boolean to check if the value has been set.
func (o *CustomerAliasListAllOf) GetCustomerAliasOk() ([]CustomerAlias, bool) {
	if o == nil {
		return nil, false
	}
	return o.CustomerAlias, true
}

// SetCustomerAlias sets field value
func (o *CustomerAliasListAllOf) SetCustomerAlias(v []CustomerAlias) {
	o.CustomerAlias = v
}

func (o CustomerAliasListAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["customer_alias"] = o.CustomerAlias
	}
	return json.Marshal(toSerialize)
}

type NullableCustomerAliasListAllOf struct {
	value *CustomerAliasListAllOf
	isSet bool
}

func (v NullableCustomerAliasListAllOf) Get() *CustomerAliasListAllOf {
	return v.value
}

func (v *NullableCustomerAliasListAllOf) Set(val *CustomerAliasListAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableCustomerAliasListAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableCustomerAliasListAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCustomerAliasListAllOf(val *CustomerAliasListAllOf) *NullableCustomerAliasListAllOf {
	return &NullableCustomerAliasListAllOf{value: val, isSet: true}
}

func (v NullableCustomerAliasListAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCustomerAliasListAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
