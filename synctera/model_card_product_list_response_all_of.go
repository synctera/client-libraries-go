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

// CardProductListResponseAllOf struct for CardProductListResponseAllOf
type CardProductListResponseAllOf struct {
	// Array of Card Products
	CardProducts []CardProductResponse `json:"card_products"`
}

// NewCardProductListResponseAllOf instantiates a new CardProductListResponseAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCardProductListResponseAllOf(cardProducts []CardProductResponse) *CardProductListResponseAllOf {
	this := CardProductListResponseAllOf{}
	this.CardProducts = cardProducts
	return &this
}

// NewCardProductListResponseAllOfWithDefaults instantiates a new CardProductListResponseAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCardProductListResponseAllOfWithDefaults() *CardProductListResponseAllOf {
	this := CardProductListResponseAllOf{}
	return &this
}

// GetCardProducts returns the CardProducts field value
func (o *CardProductListResponseAllOf) GetCardProducts() []CardProductResponse {
	if o == nil {
		var ret []CardProductResponse
		return ret
	}

	return o.CardProducts
}

// GetCardProductsOk returns a tuple with the CardProducts field value
// and a boolean to check if the value has been set.
func (o *CardProductListResponseAllOf) GetCardProductsOk() (*[]CardProductResponse, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.CardProducts, true
}

// SetCardProducts sets field value
func (o *CardProductListResponseAllOf) SetCardProducts(v []CardProductResponse) {
	o.CardProducts = v
}

func (o CardProductListResponseAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["card_products"] = o.CardProducts
	}
	return json.Marshal(toSerialize)
}

type NullableCardProductListResponseAllOf struct {
	value *CardProductListResponseAllOf
	isSet bool
}

func (v NullableCardProductListResponseAllOf) Get() *CardProductListResponseAllOf {
	return v.value
}

func (v *NullableCardProductListResponseAllOf) Set(val *CardProductListResponseAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableCardProductListResponseAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableCardProductListResponseAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCardProductListResponseAllOf(val *CardProductListResponseAllOf) *NullableCardProductListResponseAllOf {
	return &NullableCardProductListResponseAllOf{value: val, isSet: true}
}

func (v NullableCardProductListResponseAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCardProductListResponseAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


