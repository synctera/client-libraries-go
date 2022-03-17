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

// AccountCreationAllOf struct for AccountCreationAllOf
type AccountCreationAllOf struct {
	// Account template ID
	AccountTemplateId *string `json:"account_template_id,omitempty"`
	// List of the relationship for this account to the parties
	Relationships *[]Relationship `json:"relationships,omitempty"`
}

// NewAccountCreationAllOf instantiates a new AccountCreationAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccountCreationAllOf() *AccountCreationAllOf {
	this := AccountCreationAllOf{}
	return &this
}

// NewAccountCreationAllOfWithDefaults instantiates a new AccountCreationAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccountCreationAllOfWithDefaults() *AccountCreationAllOf {
	this := AccountCreationAllOf{}
	return &this
}

// GetAccountTemplateId returns the AccountTemplateId field value if set, zero value otherwise.
func (o *AccountCreationAllOf) GetAccountTemplateId() string {
	if o == nil || o.AccountTemplateId == nil {
		var ret string
		return ret
	}
	return *o.AccountTemplateId
}

// GetAccountTemplateIdOk returns a tuple with the AccountTemplateId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountCreationAllOf) GetAccountTemplateIdOk() (*string, bool) {
	if o == nil || o.AccountTemplateId == nil {
		return nil, false
	}
	return o.AccountTemplateId, true
}

// HasAccountTemplateId returns a boolean if a field has been set.
func (o *AccountCreationAllOf) HasAccountTemplateId() bool {
	if o != nil && o.AccountTemplateId != nil {
		return true
	}

	return false
}

// SetAccountTemplateId gets a reference to the given string and assigns it to the AccountTemplateId field.
func (o *AccountCreationAllOf) SetAccountTemplateId(v string) {
	o.AccountTemplateId = &v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *AccountCreationAllOf) GetRelationships() []Relationship {
	if o == nil || o.Relationships == nil {
		var ret []Relationship
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountCreationAllOf) GetRelationshipsOk() (*[]Relationship, bool) {
	if o == nil || o.Relationships == nil {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *AccountCreationAllOf) HasRelationships() bool {
	if o != nil && o.Relationships != nil {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given []Relationship and assigns it to the Relationships field.
func (o *AccountCreationAllOf) SetRelationships(v []Relationship) {
	o.Relationships = &v
}

func (o AccountCreationAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AccountTemplateId != nil {
		toSerialize["account_template_id"] = o.AccountTemplateId
	}
	if o.Relationships != nil {
		toSerialize["relationships"] = o.Relationships
	}
	return json.Marshal(toSerialize)
}

type NullableAccountCreationAllOf struct {
	value *AccountCreationAllOf
	isSet bool
}

func (v NullableAccountCreationAllOf) Get() *AccountCreationAllOf {
	return v.value
}

func (v *NullableAccountCreationAllOf) Set(val *AccountCreationAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableAccountCreationAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableAccountCreationAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccountCreationAllOf(val *AccountCreationAllOf) *NullableAccountCreationAllOf {
	return &NullableAccountCreationAllOf{value: val, isSet: true}
}

func (v NullableAccountCreationAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccountCreationAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
