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

// Fee struct for Fee
type Fee struct {
	// Fee amount
	Amount int64 `json:"amount"`
	// Fee currency code in ISO 4217
	Currency string `json:"currency"`
	// Fee type
	FeeType string `json:"fee_type"`
	// Fee ID
	Id          *string `json:"id,omitempty"`
	ProductType string  `json:"product_type"`
}

// NewFee instantiates a new Fee object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFee(amount int64, currency string, feeType string, productType string) *Fee {
	this := Fee{}
	this.Amount = amount
	this.Currency = currency
	this.FeeType = feeType
	this.ProductType = productType
	return &this
}

// NewFeeWithDefaults instantiates a new Fee object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFeeWithDefaults() *Fee {
	this := Fee{}
	return &this
}

// GetAmount returns the Amount field value
func (o *Fee) GetAmount() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Amount
}

// GetAmountOk returns a tuple with the Amount field value
// and a boolean to check if the value has been set.
func (o *Fee) GetAmountOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Amount, true
}

// SetAmount sets field value
func (o *Fee) SetAmount(v int64) {
	o.Amount = v
}

// GetCurrency returns the Currency field value
func (o *Fee) GetCurrency() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Currency
}

// GetCurrencyOk returns a tuple with the Currency field value
// and a boolean to check if the value has been set.
func (o *Fee) GetCurrencyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Currency, true
}

// SetCurrency sets field value
func (o *Fee) SetCurrency(v string) {
	o.Currency = v
}

// GetFeeType returns the FeeType field value
func (o *Fee) GetFeeType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FeeType
}

// GetFeeTypeOk returns a tuple with the FeeType field value
// and a boolean to check if the value has been set.
func (o *Fee) GetFeeTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FeeType, true
}

// SetFeeType sets field value
func (o *Fee) SetFeeType(v string) {
	o.FeeType = v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Fee) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Fee) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Fee) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *Fee) SetId(v string) {
	o.Id = &v
}

// GetProductType returns the ProductType field value
func (o *Fee) GetProductType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ProductType
}

// GetProductTypeOk returns a tuple with the ProductType field value
// and a boolean to check if the value has been set.
func (o *Fee) GetProductTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProductType, true
}

// SetProductType sets field value
func (o *Fee) SetProductType(v string) {
	o.ProductType = v
}

func (o Fee) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["amount"] = o.Amount
	}
	if true {
		toSerialize["currency"] = o.Currency
	}
	if true {
		toSerialize["fee_type"] = o.FeeType
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["product_type"] = o.ProductType
	}
	return json.Marshal(toSerialize)
}

type NullableFee struct {
	value *Fee
	isSet bool
}

func (v NullableFee) Get() *Fee {
	return v.value
}

func (v *NullableFee) Set(val *Fee) {
	v.value = val
	v.isSet = true
}

func (v NullableFee) IsSet() bool {
	return v.isSet
}

func (v *NullableFee) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFee(val *Fee) *NullableFee {
	return &NullableFee{value: val, isSet: true}
}

func (v NullableFee) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFee) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
