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

// StatementList struct for StatementList
type StatementList struct {
	// Array of statements
	Statements []StatementSummary `json:"statements"`
	// If returned, use the next_page_token to query for the next page of results. Not returned if there are no more rows.
	NextPageToken *string `json:"next_page_token,omitempty"`
}

// NewStatementList instantiates a new StatementList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStatementList(statements []StatementSummary) *StatementList {
	this := StatementList{}
	this.Statements = statements
	return &this
}

// NewStatementListWithDefaults instantiates a new StatementList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStatementListWithDefaults() *StatementList {
	this := StatementList{}
	return &this
}

// GetStatements returns the Statements field value
func (o *StatementList) GetStatements() []StatementSummary {
	if o == nil {
		var ret []StatementSummary
		return ret
	}

	return o.Statements
}

// GetStatementsOk returns a tuple with the Statements field value
// and a boolean to check if the value has been set.
func (o *StatementList) GetStatementsOk() ([]StatementSummary, bool) {
	if o == nil {
		return nil, false
	}
	return o.Statements, true
}

// SetStatements sets field value
func (o *StatementList) SetStatements(v []StatementSummary) {
	o.Statements = v
}

// GetNextPageToken returns the NextPageToken field value if set, zero value otherwise.
func (o *StatementList) GetNextPageToken() string {
	if o == nil || o.NextPageToken == nil {
		var ret string
		return ret
	}
	return *o.NextPageToken
}

// GetNextPageTokenOk returns a tuple with the NextPageToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StatementList) GetNextPageTokenOk() (*string, bool) {
	if o == nil || o.NextPageToken == nil {
		return nil, false
	}
	return o.NextPageToken, true
}

// HasNextPageToken returns a boolean if a field has been set.
func (o *StatementList) HasNextPageToken() bool {
	if o != nil && o.NextPageToken != nil {
		return true
	}

	return false
}

// SetNextPageToken gets a reference to the given string and assigns it to the NextPageToken field.
func (o *StatementList) SetNextPageToken(v string) {
	o.NextPageToken = &v
}

func (o StatementList) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["statements"] = o.Statements
	}
	if o.NextPageToken != nil {
		toSerialize["next_page_token"] = o.NextPageToken
	}
	return json.Marshal(toSerialize)
}

type NullableStatementList struct {
	value *StatementList
	isSet bool
}

func (v NullableStatementList) Get() *StatementList {
	return v.value
}

func (v *NullableStatementList) Set(val *StatementList) {
	v.value = val
	v.isSet = true
}

func (v NullableStatementList) IsSet() bool {
	return v.isSet
}

func (v *NullableStatementList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStatementList(val *StatementList) *NullableStatementList {
	return &NullableStatementList{value: val, isSet: true}
}

func (v NullableStatementList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStatementList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

