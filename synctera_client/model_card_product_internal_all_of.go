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

// CardProductInternalAllOf struct for CardProductInternalAllOf
type CardProductInternalAllOf struct {
	// Account Range ID
	AccountRangeId string `json:"account_range_id"`
	// Identifies whether a new account range will be automatically allocated
	AutoAllocateRange bool `json:"auto_allocate_range"`
	// Bank ID
	BankId int32 `json:"bank_id"`
	// BIN ID
	BinId string `json:"bin_id"`
	// Funding Source ID
	FundingSourceId string `json:"funding_source_id"`
	// Partner ID
	PartnerId int32 `json:"partner_id"`
}

// NewCardProductInternalAllOf instantiates a new CardProductInternalAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCardProductInternalAllOf(accountRangeId string, autoAllocateRange bool, bankId int32, binId string, fundingSourceId string, partnerId int32) *CardProductInternalAllOf {
	this := CardProductInternalAllOf{}
	this.AccountRangeId = accountRangeId
	this.AutoAllocateRange = autoAllocateRange
	this.BankId = bankId
	this.BinId = binId
	this.FundingSourceId = fundingSourceId
	this.PartnerId = partnerId
	return &this
}

// NewCardProductInternalAllOfWithDefaults instantiates a new CardProductInternalAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCardProductInternalAllOfWithDefaults() *CardProductInternalAllOf {
	this := CardProductInternalAllOf{}
	return &this
}

// GetAccountRangeId returns the AccountRangeId field value
func (o *CardProductInternalAllOf) GetAccountRangeId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AccountRangeId
}

// GetAccountRangeIdOk returns a tuple with the AccountRangeId field value
// and a boolean to check if the value has been set.
func (o *CardProductInternalAllOf) GetAccountRangeIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AccountRangeId, true
}

// SetAccountRangeId sets field value
func (o *CardProductInternalAllOf) SetAccountRangeId(v string) {
	o.AccountRangeId = v
}

// GetAutoAllocateRange returns the AutoAllocateRange field value
func (o *CardProductInternalAllOf) GetAutoAllocateRange() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.AutoAllocateRange
}

// GetAutoAllocateRangeOk returns a tuple with the AutoAllocateRange field value
// and a boolean to check if the value has been set.
func (o *CardProductInternalAllOf) GetAutoAllocateRangeOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AutoAllocateRange, true
}

// SetAutoAllocateRange sets field value
func (o *CardProductInternalAllOf) SetAutoAllocateRange(v bool) {
	o.AutoAllocateRange = v
}

// GetBankId returns the BankId field value
func (o *CardProductInternalAllOf) GetBankId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.BankId
}

// GetBankIdOk returns a tuple with the BankId field value
// and a boolean to check if the value has been set.
func (o *CardProductInternalAllOf) GetBankIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BankId, true
}

// SetBankId sets field value
func (o *CardProductInternalAllOf) SetBankId(v int32) {
	o.BankId = v
}

// GetBinId returns the BinId field value
func (o *CardProductInternalAllOf) GetBinId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.BinId
}

// GetBinIdOk returns a tuple with the BinId field value
// and a boolean to check if the value has been set.
func (o *CardProductInternalAllOf) GetBinIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BinId, true
}

// SetBinId sets field value
func (o *CardProductInternalAllOf) SetBinId(v string) {
	o.BinId = v
}

// GetFundingSourceId returns the FundingSourceId field value
func (o *CardProductInternalAllOf) GetFundingSourceId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FundingSourceId
}

// GetFundingSourceIdOk returns a tuple with the FundingSourceId field value
// and a boolean to check if the value has been set.
func (o *CardProductInternalAllOf) GetFundingSourceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FundingSourceId, true
}

// SetFundingSourceId sets field value
func (o *CardProductInternalAllOf) SetFundingSourceId(v string) {
	o.FundingSourceId = v
}

// GetPartnerId returns the PartnerId field value
func (o *CardProductInternalAllOf) GetPartnerId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.PartnerId
}

// GetPartnerIdOk returns a tuple with the PartnerId field value
// and a boolean to check if the value has been set.
func (o *CardProductInternalAllOf) GetPartnerIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PartnerId, true
}

// SetPartnerId sets field value
func (o *CardProductInternalAllOf) SetPartnerId(v int32) {
	o.PartnerId = v
}

func (o CardProductInternalAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["account_range_id"] = o.AccountRangeId
	}
	if true {
		toSerialize["auto_allocate_range"] = o.AutoAllocateRange
	}
	if true {
		toSerialize["bank_id"] = o.BankId
	}
	if true {
		toSerialize["bin_id"] = o.BinId
	}
	if true {
		toSerialize["funding_source_id"] = o.FundingSourceId
	}
	if true {
		toSerialize["partner_id"] = o.PartnerId
	}
	return json.Marshal(toSerialize)
}

type NullableCardProductInternalAllOf struct {
	value *CardProductInternalAllOf
	isSet bool
}

func (v NullableCardProductInternalAllOf) Get() *CardProductInternalAllOf {
	return v.value
}

func (v *NullableCardProductInternalAllOf) Set(val *CardProductInternalAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableCardProductInternalAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableCardProductInternalAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCardProductInternalAllOf(val *CardProductInternalAllOf) *NullableCardProductInternalAllOf {
	return &NullableCardProductInternalAllOf{value: val, isSet: true}
}

func (v NullableCardProductInternalAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCardProductInternalAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
