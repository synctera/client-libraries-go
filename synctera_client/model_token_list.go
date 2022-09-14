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

// TokenList struct for TokenList
type TokenList struct {
	// Array of Digital Wallet Token information of a Card
	DigitalWalletTokens []DigitalWalletTokenResponse `json:"digital_wallet_tokens,omitempty"`
}

// NewTokenList instantiates a new TokenList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTokenList() *TokenList {
	this := TokenList{}
	return &this
}

// NewTokenListWithDefaults instantiates a new TokenList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTokenListWithDefaults() *TokenList {
	this := TokenList{}
	return &this
}

// GetDigitalWalletTokens returns the DigitalWalletTokens field value if set, zero value otherwise.
func (o *TokenList) GetDigitalWalletTokens() []DigitalWalletTokenResponse {
	if o == nil || o.DigitalWalletTokens == nil {
		var ret []DigitalWalletTokenResponse
		return ret
	}
	return o.DigitalWalletTokens
}

// GetDigitalWalletTokensOk returns a tuple with the DigitalWalletTokens field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenList) GetDigitalWalletTokensOk() ([]DigitalWalletTokenResponse, bool) {
	if o == nil || o.DigitalWalletTokens == nil {
		return nil, false
	}
	return o.DigitalWalletTokens, true
}

// HasDigitalWalletTokens returns a boolean if a field has been set.
func (o *TokenList) HasDigitalWalletTokens() bool {
	if o != nil && o.DigitalWalletTokens != nil {
		return true
	}

	return false
}

// SetDigitalWalletTokens gets a reference to the given []DigitalWalletTokenResponse and assigns it to the DigitalWalletTokens field.
func (o *TokenList) SetDigitalWalletTokens(v []DigitalWalletTokenResponse) {
	o.DigitalWalletTokens = v
}

func (o TokenList) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.DigitalWalletTokens != nil {
		toSerialize["digital_wallet_tokens"] = o.DigitalWalletTokens
	}
	return json.Marshal(toSerialize)
}

type NullableTokenList struct {
	value *TokenList
	isSet bool
}

func (v NullableTokenList) Get() *TokenList {
	return v.value
}

func (v *NullableTokenList) Set(val *TokenList) {
	v.value = val
	v.isSet = true
}

func (v NullableTokenList) IsSet() bool {
	return v.isSet
}

func (v *NullableTokenList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTokenList(val *TokenList) *NullableTokenList {
	return &NullableTokenList{value: val, isSet: true}
}

func (v NullableTokenList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTokenList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
