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

// RecipientName The name of the recipient to whom the card will be shipped
type RecipientName struct {
	FirstName  string  `json:"first_name"`
	LastName   string  `json:"last_name"`
	MiddleName *string `json:"middle_name,omitempty"`
}

// NewRecipientName instantiates a new RecipientName object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRecipientName(firstName string, lastName string) *RecipientName {
	this := RecipientName{}
	this.FirstName = firstName
	this.LastName = lastName
	return &this
}

// NewRecipientNameWithDefaults instantiates a new RecipientName object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRecipientNameWithDefaults() *RecipientName {
	this := RecipientName{}
	return &this
}

// GetFirstName returns the FirstName field value
func (o *RecipientName) GetFirstName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FirstName
}

// GetFirstNameOk returns a tuple with the FirstName field value
// and a boolean to check if the value has been set.
func (o *RecipientName) GetFirstNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FirstName, true
}

// SetFirstName sets field value
func (o *RecipientName) SetFirstName(v string) {
	o.FirstName = v
}

// GetLastName returns the LastName field value
func (o *RecipientName) GetLastName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.LastName
}

// GetLastNameOk returns a tuple with the LastName field value
// and a boolean to check if the value has been set.
func (o *RecipientName) GetLastNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LastName, true
}

// SetLastName sets field value
func (o *RecipientName) SetLastName(v string) {
	o.LastName = v
}

// GetMiddleName returns the MiddleName field value if set, zero value otherwise.
func (o *RecipientName) GetMiddleName() string {
	if o == nil || o.MiddleName == nil {
		var ret string
		return ret
	}
	return *o.MiddleName
}

// GetMiddleNameOk returns a tuple with the MiddleName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecipientName) GetMiddleNameOk() (*string, bool) {
	if o == nil || o.MiddleName == nil {
		return nil, false
	}
	return o.MiddleName, true
}

// HasMiddleName returns a boolean if a field has been set.
func (o *RecipientName) HasMiddleName() bool {
	if o != nil && o.MiddleName != nil {
		return true
	}

	return false
}

// SetMiddleName gets a reference to the given string and assigns it to the MiddleName field.
func (o *RecipientName) SetMiddleName(v string) {
	o.MiddleName = &v
}

func (o RecipientName) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["first_name"] = o.FirstName
	}
	if true {
		toSerialize["last_name"] = o.LastName
	}
	if o.MiddleName != nil {
		toSerialize["middle_name"] = o.MiddleName
	}
	return json.Marshal(toSerialize)
}

type NullableRecipientName struct {
	value *RecipientName
	isSet bool
}

func (v NullableRecipientName) Get() *RecipientName {
	return v.value
}

func (v *NullableRecipientName) Set(val *RecipientName) {
	v.value = val
	v.isSet = true
}

func (v NullableRecipientName) IsSet() bool {
	return v.isSet
}

func (v *NullableRecipientName) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRecipientName(val *RecipientName) *NullableRecipientName {
	return &NullableRecipientName{value: val, isSet: true}
}

func (v NullableRecipientName) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRecipientName) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
