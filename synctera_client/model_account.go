/*
Synctera API

<h2>Let's build something great.</h2><p>Welcome to the official reference documentation for Synctera APIs. Our APIs are the best way to automate your company's banking needs and are designed to be easy to understand and implement.</p><p>We're continuously growing this library and what you see here is just the start, but if you need something specific or have a question, <a class='text-blue-600' href='https://synctera.com/contact' target='_blank' rel='noreferrer'>contact us</a>.</p>

API version: 0.20.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package synctera_client

import (
	"encoding/json"
	"fmt"
)

// Account Account
type Account struct {
	AccountDepository   *AccountDepository
	AccountLineOfCredit *AccountLineOfCredit
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *Account) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into AccountDepository
	err = json.Unmarshal(data, &dst.AccountDepository)
	if err == nil {
		jsonAccountDepository, _ := json.Marshal(dst.AccountDepository)
		if string(jsonAccountDepository) == "{}" { // empty struct
			dst.AccountDepository = nil
		} else {
			return nil // data stored in dst.AccountDepository, return on the first match
		}
	} else {
		dst.AccountDepository = nil
	}

	// try to unmarshal JSON data into AccountLineOfCredit
	err = json.Unmarshal(data, &dst.AccountLineOfCredit)
	if err == nil {
		jsonAccountLineOfCredit, _ := json.Marshal(dst.AccountLineOfCredit)
		if string(jsonAccountLineOfCredit) == "{}" { // empty struct
			dst.AccountLineOfCredit = nil
		} else {
			return nil // data stored in dst.AccountLineOfCredit, return on the first match
		}
	} else {
		dst.AccountLineOfCredit = nil
	}

	return fmt.Errorf("Data failed to match schemas in anyOf(Account)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *Account) MarshalJSON() ([]byte, error) {
	if src.AccountDepository != nil {
		return json.Marshal(&src.AccountDepository)
	}

	if src.AccountLineOfCredit != nil {
		return json.Marshal(&src.AccountLineOfCredit)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableAccount struct {
	value *Account
	isSet bool
}

func (v NullableAccount) Get() *Account {
	return v.value
}

func (v *NullableAccount) Set(val *Account) {
	v.value = val
	v.isSet = true
}

func (v NullableAccount) IsSet() bool {
	return v.isSet
}

func (v *NullableAccount) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccount(val *Account) *NullableAccount {
	return &NullableAccount{value: val, isSet: true}
}

func (v NullableAccount) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccount) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
