/*
 * Synctera API
 *
 * <h2>Let's build something great.</h2><p>Welcome to the official reference documentation for Synctera APIs. Our APIs are the best way to automate your company's banking needs and are designed to be easy to understand and implement.</p><p>We're continuously growing this library and what you see here is just the start, but if you need something specific or have a question, <a class='text-blue-600' href='https://synctera.com/contact' target='_blank' rel='noreferrer'>contact us</a>.</p> 
 *
 * API version: 0.5.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package synctera

import (
	"encoding/json"
)

// CardProductListAllOf struct for CardProductListAllOf
type CardProductListAllOf struct {
	// Array of Card Products
	CardProducts []CardProduct `json:"card_products"`
}

// NewCardProductListAllOf instantiates a new CardProductListAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCardProductListAllOf(cardProducts []CardProduct) *CardProductListAllOf {
	this := CardProductListAllOf{}
	this.CardProducts = cardProducts
	return &this
}

// NewCardProductListAllOfWithDefaults instantiates a new CardProductListAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCardProductListAllOfWithDefaults() *CardProductListAllOf {
	this := CardProductListAllOf{}
	return &this
}

// GetCardProducts returns the CardProducts field value
func (o *CardProductListAllOf) GetCardProducts() []CardProduct {
	if o == nil {
		var ret []CardProduct
		return ret
	}

	return o.CardProducts
}

// GetCardProductsOk returns a tuple with the CardProducts field value
// and a boolean to check if the value has been set.
func (o *CardProductListAllOf) GetCardProductsOk() (*[]CardProduct, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.CardProducts, true
}

// SetCardProducts sets field value
func (o *CardProductListAllOf) SetCardProducts(v []CardProduct) {
	o.CardProducts = v
}

func (o CardProductListAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["card_products"] = o.CardProducts
	}
	return json.Marshal(toSerialize)
}

type NullableCardProductListAllOf struct {
	value *CardProductListAllOf
	isSet bool
}

func (v NullableCardProductListAllOf) Get() *CardProductListAllOf {
	return v.value
}

func (v *NullableCardProductListAllOf) Set(val *CardProductListAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableCardProductListAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableCardProductListAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCardProductListAllOf(val *CardProductListAllOf) *NullableCardProductListAllOf {
	return &NullableCardProductListAllOf{value: val, isSet: true}
}

func (v NullableCardProductListAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCardProductListAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


