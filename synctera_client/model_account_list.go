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

// AccountList struct for AccountList
type AccountList struct {
	// Array of Accounts
	Accounts []AccountGenericResponse `json:"accounts"`
	// If returned, use the next_page_token to query for the next page of results. Not returned if there are no more rows.
	NextPageToken *string `json:"next_page_token,omitempty"`
}

// NewAccountList instantiates a new AccountList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccountList(accounts []AccountGenericResponse) *AccountList {
	this := AccountList{}
	this.Accounts = accounts
	return &this
}

// NewAccountListWithDefaults instantiates a new AccountList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccountListWithDefaults() *AccountList {
	this := AccountList{}
	return &this
}

// GetAccounts returns the Accounts field value
func (o *AccountList) GetAccounts() []AccountGenericResponse {
	if o == nil {
		var ret []AccountGenericResponse
		return ret
	}

	return o.Accounts
}

// GetAccountsOk returns a tuple with the Accounts field value
// and a boolean to check if the value has been set.
func (o *AccountList) GetAccountsOk() ([]AccountGenericResponse, bool) {
	if o == nil {
		return nil, false
	}
	return o.Accounts, true
}

// SetAccounts sets field value
func (o *AccountList) SetAccounts(v []AccountGenericResponse) {
	o.Accounts = v
}

// GetNextPageToken returns the NextPageToken field value if set, zero value otherwise.
func (o *AccountList) GetNextPageToken() string {
	if o == nil || o.NextPageToken == nil {
		var ret string
		return ret
	}
	return *o.NextPageToken
}

// GetNextPageTokenOk returns a tuple with the NextPageToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountList) GetNextPageTokenOk() (*string, bool) {
	if o == nil || o.NextPageToken == nil {
		return nil, false
	}
	return o.NextPageToken, true
}

// HasNextPageToken returns a boolean if a field has been set.
func (o *AccountList) HasNextPageToken() bool {
	if o != nil && o.NextPageToken != nil {
		return true
	}

	return false
}

// SetNextPageToken gets a reference to the given string and assigns it to the NextPageToken field.
func (o *AccountList) SetNextPageToken(v string) {
	o.NextPageToken = &v
}

func (o AccountList) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["accounts"] = o.Accounts
	}
	if o.NextPageToken != nil {
		toSerialize["next_page_token"] = o.NextPageToken
	}
	return json.Marshal(toSerialize)
}

type NullableAccountList struct {
	value *AccountList
	isSet bool
}

func (v NullableAccountList) Get() *AccountList {
	return v.value
}

func (v *NullableAccountList) Set(val *AccountList) {
	v.value = val
	v.isSet = true
}

func (v NullableAccountList) IsSet() bool {
	return v.isSet
}

func (v *NullableAccountList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccountList(val *AccountList) *NullableAccountList {
	return &NullableAccountList{value: val, isSet: true}
}

func (v NullableAccountList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccountList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
