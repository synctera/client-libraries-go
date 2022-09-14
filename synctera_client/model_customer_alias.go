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

// CustomerAlias Represents a customer's alias
type CustomerAlias struct {
	AliasInfo       map[string]interface{} `json:"alias_info,omitempty"`
	AliasName       *string                `json:"alias_name,omitempty"`
	AliasSource     *string                `json:"alias_source,omitempty"`
	AliasType       *string                `json:"alias_type,omitempty"`
	CustomerAliasId *string                `json:"customer_alias_id,omitempty"`
	CustomerId      *string                `json:"customer_id,omitempty"`
}

// NewCustomerAlias instantiates a new CustomerAlias object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCustomerAlias() *CustomerAlias {
	this := CustomerAlias{}
	return &this
}

// NewCustomerAliasWithDefaults instantiates a new CustomerAlias object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCustomerAliasWithDefaults() *CustomerAlias {
	this := CustomerAlias{}
	return &this
}

// GetAliasInfo returns the AliasInfo field value if set, zero value otherwise.
func (o *CustomerAlias) GetAliasInfo() map[string]interface{} {
	if o == nil || o.AliasInfo == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.AliasInfo
}

// GetAliasInfoOk returns a tuple with the AliasInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerAlias) GetAliasInfoOk() (map[string]interface{}, bool) {
	if o == nil || o.AliasInfo == nil {
		return nil, false
	}
	return o.AliasInfo, true
}

// HasAliasInfo returns a boolean if a field has been set.
func (o *CustomerAlias) HasAliasInfo() bool {
	if o != nil && o.AliasInfo != nil {
		return true
	}

	return false
}

// SetAliasInfo gets a reference to the given map[string]interface{} and assigns it to the AliasInfo field.
func (o *CustomerAlias) SetAliasInfo(v map[string]interface{}) {
	o.AliasInfo = v
}

// GetAliasName returns the AliasName field value if set, zero value otherwise.
func (o *CustomerAlias) GetAliasName() string {
	if o == nil || o.AliasName == nil {
		var ret string
		return ret
	}
	return *o.AliasName
}

// GetAliasNameOk returns a tuple with the AliasName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerAlias) GetAliasNameOk() (*string, bool) {
	if o == nil || o.AliasName == nil {
		return nil, false
	}
	return o.AliasName, true
}

// HasAliasName returns a boolean if a field has been set.
func (o *CustomerAlias) HasAliasName() bool {
	if o != nil && o.AliasName != nil {
		return true
	}

	return false
}

// SetAliasName gets a reference to the given string and assigns it to the AliasName field.
func (o *CustomerAlias) SetAliasName(v string) {
	o.AliasName = &v
}

// GetAliasSource returns the AliasSource field value if set, zero value otherwise.
func (o *CustomerAlias) GetAliasSource() string {
	if o == nil || o.AliasSource == nil {
		var ret string
		return ret
	}
	return *o.AliasSource
}

// GetAliasSourceOk returns a tuple with the AliasSource field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerAlias) GetAliasSourceOk() (*string, bool) {
	if o == nil || o.AliasSource == nil {
		return nil, false
	}
	return o.AliasSource, true
}

// HasAliasSource returns a boolean if a field has been set.
func (o *CustomerAlias) HasAliasSource() bool {
	if o != nil && o.AliasSource != nil {
		return true
	}

	return false
}

// SetAliasSource gets a reference to the given string and assigns it to the AliasSource field.
func (o *CustomerAlias) SetAliasSource(v string) {
	o.AliasSource = &v
}

// GetAliasType returns the AliasType field value if set, zero value otherwise.
func (o *CustomerAlias) GetAliasType() string {
	if o == nil || o.AliasType == nil {
		var ret string
		return ret
	}
	return *o.AliasType
}

// GetAliasTypeOk returns a tuple with the AliasType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerAlias) GetAliasTypeOk() (*string, bool) {
	if o == nil || o.AliasType == nil {
		return nil, false
	}
	return o.AliasType, true
}

// HasAliasType returns a boolean if a field has been set.
func (o *CustomerAlias) HasAliasType() bool {
	if o != nil && o.AliasType != nil {
		return true
	}

	return false
}

// SetAliasType gets a reference to the given string and assigns it to the AliasType field.
func (o *CustomerAlias) SetAliasType(v string) {
	o.AliasType = &v
}

// GetCustomerAliasId returns the CustomerAliasId field value if set, zero value otherwise.
func (o *CustomerAlias) GetCustomerAliasId() string {
	if o == nil || o.CustomerAliasId == nil {
		var ret string
		return ret
	}
	return *o.CustomerAliasId
}

// GetCustomerAliasIdOk returns a tuple with the CustomerAliasId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerAlias) GetCustomerAliasIdOk() (*string, bool) {
	if o == nil || o.CustomerAliasId == nil {
		return nil, false
	}
	return o.CustomerAliasId, true
}

// HasCustomerAliasId returns a boolean if a field has been set.
func (o *CustomerAlias) HasCustomerAliasId() bool {
	if o != nil && o.CustomerAliasId != nil {
		return true
	}

	return false
}

// SetCustomerAliasId gets a reference to the given string and assigns it to the CustomerAliasId field.
func (o *CustomerAlias) SetCustomerAliasId(v string) {
	o.CustomerAliasId = &v
}

// GetCustomerId returns the CustomerId field value if set, zero value otherwise.
func (o *CustomerAlias) GetCustomerId() string {
	if o == nil || o.CustomerId == nil {
		var ret string
		return ret
	}
	return *o.CustomerId
}

// GetCustomerIdOk returns a tuple with the CustomerId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerAlias) GetCustomerIdOk() (*string, bool) {
	if o == nil || o.CustomerId == nil {
		return nil, false
	}
	return o.CustomerId, true
}

// HasCustomerId returns a boolean if a field has been set.
func (o *CustomerAlias) HasCustomerId() bool {
	if o != nil && o.CustomerId != nil {
		return true
	}

	return false
}

// SetCustomerId gets a reference to the given string and assigns it to the CustomerId field.
func (o *CustomerAlias) SetCustomerId(v string) {
	o.CustomerId = &v
}

func (o CustomerAlias) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AliasInfo != nil {
		toSerialize["alias_info"] = o.AliasInfo
	}
	if o.AliasName != nil {
		toSerialize["alias_name"] = o.AliasName
	}
	if o.AliasSource != nil {
		toSerialize["alias_source"] = o.AliasSource
	}
	if o.AliasType != nil {
		toSerialize["alias_type"] = o.AliasType
	}
	if o.CustomerAliasId != nil {
		toSerialize["customer_alias_id"] = o.CustomerAliasId
	}
	if o.CustomerId != nil {
		toSerialize["customer_id"] = o.CustomerId
	}
	return json.Marshal(toSerialize)
}

type NullableCustomerAlias struct {
	value *CustomerAlias
	isSet bool
}

func (v NullableCustomerAlias) Get() *CustomerAlias {
	return v.value
}

func (v *NullableCustomerAlias) Set(val *CustomerAlias) {
	v.value = val
	v.isSet = true
}

func (v NullableCustomerAlias) IsSet() bool {
	return v.isSet
}

func (v *NullableCustomerAlias) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCustomerAlias(val *CustomerAlias) *NullableCustomerAlias {
	return &NullableCustomerAlias{value: val, isSet: true}
}

func (v NullableCustomerAlias) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCustomerAlias) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
