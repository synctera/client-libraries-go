/*
Synctera API

<h2>Let's build something great.</h2><p>Welcome to the official reference documentation for Synctera APIs. Our APIs are the best way to automate your company's banking needs and are designed to be easy to understand and implement.</p><p>We're continuously growing this library and what you see here is just the start, but if you need something specific or have a question, <a class='text-blue-600' href='https://synctera.com/contact' target='_blank' rel='noreferrer'>contact us</a>.</p> 

API version: 0.5.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// CardOptions struct for CardOptions
type CardOptions struct {
	BillingAddress *BillingAddress `json:"billing_address,omitempty"`
	CardPresent *bool `json:"card_present,omitempty"`
	Cvv *string `json:"cvv,omitempty"`
	Expiration *string `json:"expiration,omitempty"`
}

// NewCardOptions instantiates a new CardOptions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCardOptions() *CardOptions {
	this := CardOptions{}
	var cardPresent bool = false
	this.CardPresent = &cardPresent
	return &this
}

// NewCardOptionsWithDefaults instantiates a new CardOptions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCardOptionsWithDefaults() *CardOptions {
	this := CardOptions{}
	var cardPresent bool = false
	this.CardPresent = &cardPresent
	return &this
}

// GetBillingAddress returns the BillingAddress field value if set, zero value otherwise.
func (o *CardOptions) GetBillingAddress() BillingAddress {
	if o == nil || o.BillingAddress == nil {
		var ret BillingAddress
		return ret
	}
	return *o.BillingAddress
}

// GetBillingAddressOk returns a tuple with the BillingAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CardOptions) GetBillingAddressOk() (*BillingAddress, bool) {
	if o == nil || o.BillingAddress == nil {
		return nil, false
	}
	return o.BillingAddress, true
}

// HasBillingAddress returns a boolean if a field has been set.
func (o *CardOptions) HasBillingAddress() bool {
	if o != nil && o.BillingAddress != nil {
		return true
	}

	return false
}

// SetBillingAddress gets a reference to the given BillingAddress and assigns it to the BillingAddress field.
func (o *CardOptions) SetBillingAddress(v BillingAddress) {
	o.BillingAddress = &v
}

// GetCardPresent returns the CardPresent field value if set, zero value otherwise.
func (o *CardOptions) GetCardPresent() bool {
	if o == nil || o.CardPresent == nil {
		var ret bool
		return ret
	}
	return *o.CardPresent
}

// GetCardPresentOk returns a tuple with the CardPresent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CardOptions) GetCardPresentOk() (*bool, bool) {
	if o == nil || o.CardPresent == nil {
		return nil, false
	}
	return o.CardPresent, true
}

// HasCardPresent returns a boolean if a field has been set.
func (o *CardOptions) HasCardPresent() bool {
	if o != nil && o.CardPresent != nil {
		return true
	}

	return false
}

// SetCardPresent gets a reference to the given bool and assigns it to the CardPresent field.
func (o *CardOptions) SetCardPresent(v bool) {
	o.CardPresent = &v
}

// GetCvv returns the Cvv field value if set, zero value otherwise.
func (o *CardOptions) GetCvv() string {
	if o == nil || o.Cvv == nil {
		var ret string
		return ret
	}
	return *o.Cvv
}

// GetCvvOk returns a tuple with the Cvv field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CardOptions) GetCvvOk() (*string, bool) {
	if o == nil || o.Cvv == nil {
		return nil, false
	}
	return o.Cvv, true
}

// HasCvv returns a boolean if a field has been set.
func (o *CardOptions) HasCvv() bool {
	if o != nil && o.Cvv != nil {
		return true
	}

	return false
}

// SetCvv gets a reference to the given string and assigns it to the Cvv field.
func (o *CardOptions) SetCvv(v string) {
	o.Cvv = &v
}

// GetExpiration returns the Expiration field value if set, zero value otherwise.
func (o *CardOptions) GetExpiration() string {
	if o == nil || o.Expiration == nil {
		var ret string
		return ret
	}
	return *o.Expiration
}

// GetExpirationOk returns a tuple with the Expiration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CardOptions) GetExpirationOk() (*string, bool) {
	if o == nil || o.Expiration == nil {
		return nil, false
	}
	return o.Expiration, true
}

// HasExpiration returns a boolean if a field has been set.
func (o *CardOptions) HasExpiration() bool {
	if o != nil && o.Expiration != nil {
		return true
	}

	return false
}

// SetExpiration gets a reference to the given string and assigns it to the Expiration field.
func (o *CardOptions) SetExpiration(v string) {
	o.Expiration = &v
}

func (o CardOptions) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.BillingAddress != nil {
		toSerialize["billing_address"] = o.BillingAddress
	}
	if o.CardPresent != nil {
		toSerialize["card_present"] = o.CardPresent
	}
	if o.Cvv != nil {
		toSerialize["cvv"] = o.Cvv
	}
	if o.Expiration != nil {
		toSerialize["expiration"] = o.Expiration
	}
	return json.Marshal(toSerialize)
}

type NullableCardOptions struct {
	value *CardOptions
	isSet bool
}

func (v NullableCardOptions) Get() *CardOptions {
	return v.value
}

func (v *NullableCardOptions) Set(val *CardOptions) {
	v.value = val
	v.isSet = true
}

func (v NullableCardOptions) IsSet() bool {
	return v.isSet
}

func (v *NullableCardOptions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCardOptions(val *CardOptions) *NullableCardOptions {
	return &NullableCardOptions{value: val, isSet: true}
}

func (v NullableCardOptions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCardOptions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


