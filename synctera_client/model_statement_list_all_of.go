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

// StatementListAllOf struct for StatementListAllOf
type StatementListAllOf struct {
	// Array of statements
	Statements []StatementSummary `json:"statements"`
}

// NewStatementListAllOf instantiates a new StatementListAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStatementListAllOf(statements []StatementSummary) *StatementListAllOf {
	this := StatementListAllOf{}
	this.Statements = statements
	return &this
}

// NewStatementListAllOfWithDefaults instantiates a new StatementListAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStatementListAllOfWithDefaults() *StatementListAllOf {
	this := StatementListAllOf{}
	return &this
}

// GetStatements returns the Statements field value
func (o *StatementListAllOf) GetStatements() []StatementSummary {
	if o == nil {
		var ret []StatementSummary
		return ret
	}

	return o.Statements
}

// GetStatementsOk returns a tuple with the Statements field value
// and a boolean to check if the value has been set.
func (o *StatementListAllOf) GetStatementsOk() ([]StatementSummary, bool) {
	if o == nil {
		return nil, false
	}
	return o.Statements, true
}

// SetStatements sets field value
func (o *StatementListAllOf) SetStatements(v []StatementSummary) {
	o.Statements = v
}

func (o StatementListAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["statements"] = o.Statements
	}
	return json.Marshal(toSerialize)
}

type NullableStatementListAllOf struct {
	value *StatementListAllOf
	isSet bool
}

func (v NullableStatementListAllOf) Get() *StatementListAllOf {
	return v.value
}

func (v *NullableStatementListAllOf) Set(val *StatementListAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableStatementListAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableStatementListAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStatementListAllOf(val *StatementListAllOf) *NullableStatementListAllOf {
	return &NullableStatementListAllOf{value: val, isSet: true}
}

func (v NullableStatementListAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStatementListAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


