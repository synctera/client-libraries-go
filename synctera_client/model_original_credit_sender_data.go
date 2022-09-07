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

// OriginalCreditSenderData struct for OriginalCreditSenderData
type OriginalCreditSenderData struct {
	FundingSource string `json:"funding_source"`
	SenderAccountNumber *string `json:"sender_account_number,omitempty"`
	SenderAccountType *string `json:"sender_account_type,omitempty"`
	SenderAddress *string `json:"sender_address,omitempty"`
	SenderCity *string `json:"sender_city,omitempty"`
	SenderCountry *string `json:"sender_country,omitempty"`
	SenderName *string `json:"sender_name,omitempty"`
	SenderReferenceNumber *string `json:"sender_reference_number,omitempty"`
	SenderState *string `json:"sender_state,omitempty"`
	TransactionPurpose *string `json:"transaction_purpose,omitempty"`
	UniqueTransactionReferenceNumber *string `json:"unique_transaction_reference_number,omitempty"`
}

// NewOriginalCreditSenderData instantiates a new OriginalCreditSenderData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOriginalCreditSenderData(fundingSource string) *OriginalCreditSenderData {
	this := OriginalCreditSenderData{}
	this.FundingSource = fundingSource
	return &this
}

// NewOriginalCreditSenderDataWithDefaults instantiates a new OriginalCreditSenderData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOriginalCreditSenderDataWithDefaults() *OriginalCreditSenderData {
	this := OriginalCreditSenderData{}
	return &this
}

// GetFundingSource returns the FundingSource field value
func (o *OriginalCreditSenderData) GetFundingSource() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FundingSource
}

// GetFundingSourceOk returns a tuple with the FundingSource field value
// and a boolean to check if the value has been set.
func (o *OriginalCreditSenderData) GetFundingSourceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FundingSource, true
}

// SetFundingSource sets field value
func (o *OriginalCreditSenderData) SetFundingSource(v string) {
	o.FundingSource = v
}

// GetSenderAccountNumber returns the SenderAccountNumber field value if set, zero value otherwise.
func (o *OriginalCreditSenderData) GetSenderAccountNumber() string {
	if o == nil || o.SenderAccountNumber == nil {
		var ret string
		return ret
	}
	return *o.SenderAccountNumber
}

// GetSenderAccountNumberOk returns a tuple with the SenderAccountNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OriginalCreditSenderData) GetSenderAccountNumberOk() (*string, bool) {
	if o == nil || o.SenderAccountNumber == nil {
		return nil, false
	}
	return o.SenderAccountNumber, true
}

// HasSenderAccountNumber returns a boolean if a field has been set.
func (o *OriginalCreditSenderData) HasSenderAccountNumber() bool {
	if o != nil && o.SenderAccountNumber != nil {
		return true
	}

	return false
}

// SetSenderAccountNumber gets a reference to the given string and assigns it to the SenderAccountNumber field.
func (o *OriginalCreditSenderData) SetSenderAccountNumber(v string) {
	o.SenderAccountNumber = &v
}

// GetSenderAccountType returns the SenderAccountType field value if set, zero value otherwise.
func (o *OriginalCreditSenderData) GetSenderAccountType() string {
	if o == nil || o.SenderAccountType == nil {
		var ret string
		return ret
	}
	return *o.SenderAccountType
}

// GetSenderAccountTypeOk returns a tuple with the SenderAccountType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OriginalCreditSenderData) GetSenderAccountTypeOk() (*string, bool) {
	if o == nil || o.SenderAccountType == nil {
		return nil, false
	}
	return o.SenderAccountType, true
}

// HasSenderAccountType returns a boolean if a field has been set.
func (o *OriginalCreditSenderData) HasSenderAccountType() bool {
	if o != nil && o.SenderAccountType != nil {
		return true
	}

	return false
}

// SetSenderAccountType gets a reference to the given string and assigns it to the SenderAccountType field.
func (o *OriginalCreditSenderData) SetSenderAccountType(v string) {
	o.SenderAccountType = &v
}

// GetSenderAddress returns the SenderAddress field value if set, zero value otherwise.
func (o *OriginalCreditSenderData) GetSenderAddress() string {
	if o == nil || o.SenderAddress == nil {
		var ret string
		return ret
	}
	return *o.SenderAddress
}

// GetSenderAddressOk returns a tuple with the SenderAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OriginalCreditSenderData) GetSenderAddressOk() (*string, bool) {
	if o == nil || o.SenderAddress == nil {
		return nil, false
	}
	return o.SenderAddress, true
}

// HasSenderAddress returns a boolean if a field has been set.
func (o *OriginalCreditSenderData) HasSenderAddress() bool {
	if o != nil && o.SenderAddress != nil {
		return true
	}

	return false
}

// SetSenderAddress gets a reference to the given string and assigns it to the SenderAddress field.
func (o *OriginalCreditSenderData) SetSenderAddress(v string) {
	o.SenderAddress = &v
}

// GetSenderCity returns the SenderCity field value if set, zero value otherwise.
func (o *OriginalCreditSenderData) GetSenderCity() string {
	if o == nil || o.SenderCity == nil {
		var ret string
		return ret
	}
	return *o.SenderCity
}

// GetSenderCityOk returns a tuple with the SenderCity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OriginalCreditSenderData) GetSenderCityOk() (*string, bool) {
	if o == nil || o.SenderCity == nil {
		return nil, false
	}
	return o.SenderCity, true
}

// HasSenderCity returns a boolean if a field has been set.
func (o *OriginalCreditSenderData) HasSenderCity() bool {
	if o != nil && o.SenderCity != nil {
		return true
	}

	return false
}

// SetSenderCity gets a reference to the given string and assigns it to the SenderCity field.
func (o *OriginalCreditSenderData) SetSenderCity(v string) {
	o.SenderCity = &v
}

// GetSenderCountry returns the SenderCountry field value if set, zero value otherwise.
func (o *OriginalCreditSenderData) GetSenderCountry() string {
	if o == nil || o.SenderCountry == nil {
		var ret string
		return ret
	}
	return *o.SenderCountry
}

// GetSenderCountryOk returns a tuple with the SenderCountry field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OriginalCreditSenderData) GetSenderCountryOk() (*string, bool) {
	if o == nil || o.SenderCountry == nil {
		return nil, false
	}
	return o.SenderCountry, true
}

// HasSenderCountry returns a boolean if a field has been set.
func (o *OriginalCreditSenderData) HasSenderCountry() bool {
	if o != nil && o.SenderCountry != nil {
		return true
	}

	return false
}

// SetSenderCountry gets a reference to the given string and assigns it to the SenderCountry field.
func (o *OriginalCreditSenderData) SetSenderCountry(v string) {
	o.SenderCountry = &v
}

// GetSenderName returns the SenderName field value if set, zero value otherwise.
func (o *OriginalCreditSenderData) GetSenderName() string {
	if o == nil || o.SenderName == nil {
		var ret string
		return ret
	}
	return *o.SenderName
}

// GetSenderNameOk returns a tuple with the SenderName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OriginalCreditSenderData) GetSenderNameOk() (*string, bool) {
	if o == nil || o.SenderName == nil {
		return nil, false
	}
	return o.SenderName, true
}

// HasSenderName returns a boolean if a field has been set.
func (o *OriginalCreditSenderData) HasSenderName() bool {
	if o != nil && o.SenderName != nil {
		return true
	}

	return false
}

// SetSenderName gets a reference to the given string and assigns it to the SenderName field.
func (o *OriginalCreditSenderData) SetSenderName(v string) {
	o.SenderName = &v
}

// GetSenderReferenceNumber returns the SenderReferenceNumber field value if set, zero value otherwise.
func (o *OriginalCreditSenderData) GetSenderReferenceNumber() string {
	if o == nil || o.SenderReferenceNumber == nil {
		var ret string
		return ret
	}
	return *o.SenderReferenceNumber
}

// GetSenderReferenceNumberOk returns a tuple with the SenderReferenceNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OriginalCreditSenderData) GetSenderReferenceNumberOk() (*string, bool) {
	if o == nil || o.SenderReferenceNumber == nil {
		return nil, false
	}
	return o.SenderReferenceNumber, true
}

// HasSenderReferenceNumber returns a boolean if a field has been set.
func (o *OriginalCreditSenderData) HasSenderReferenceNumber() bool {
	if o != nil && o.SenderReferenceNumber != nil {
		return true
	}

	return false
}

// SetSenderReferenceNumber gets a reference to the given string and assigns it to the SenderReferenceNumber field.
func (o *OriginalCreditSenderData) SetSenderReferenceNumber(v string) {
	o.SenderReferenceNumber = &v
}

// GetSenderState returns the SenderState field value if set, zero value otherwise.
func (o *OriginalCreditSenderData) GetSenderState() string {
	if o == nil || o.SenderState == nil {
		var ret string
		return ret
	}
	return *o.SenderState
}

// GetSenderStateOk returns a tuple with the SenderState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OriginalCreditSenderData) GetSenderStateOk() (*string, bool) {
	if o == nil || o.SenderState == nil {
		return nil, false
	}
	return o.SenderState, true
}

// HasSenderState returns a boolean if a field has been set.
func (o *OriginalCreditSenderData) HasSenderState() bool {
	if o != nil && o.SenderState != nil {
		return true
	}

	return false
}

// SetSenderState gets a reference to the given string and assigns it to the SenderState field.
func (o *OriginalCreditSenderData) SetSenderState(v string) {
	o.SenderState = &v
}

// GetTransactionPurpose returns the TransactionPurpose field value if set, zero value otherwise.
func (o *OriginalCreditSenderData) GetTransactionPurpose() string {
	if o == nil || o.TransactionPurpose == nil {
		var ret string
		return ret
	}
	return *o.TransactionPurpose
}

// GetTransactionPurposeOk returns a tuple with the TransactionPurpose field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OriginalCreditSenderData) GetTransactionPurposeOk() (*string, bool) {
	if o == nil || o.TransactionPurpose == nil {
		return nil, false
	}
	return o.TransactionPurpose, true
}

// HasTransactionPurpose returns a boolean if a field has been set.
func (o *OriginalCreditSenderData) HasTransactionPurpose() bool {
	if o != nil && o.TransactionPurpose != nil {
		return true
	}

	return false
}

// SetTransactionPurpose gets a reference to the given string and assigns it to the TransactionPurpose field.
func (o *OriginalCreditSenderData) SetTransactionPurpose(v string) {
	o.TransactionPurpose = &v
}

// GetUniqueTransactionReferenceNumber returns the UniqueTransactionReferenceNumber field value if set, zero value otherwise.
func (o *OriginalCreditSenderData) GetUniqueTransactionReferenceNumber() string {
	if o == nil || o.UniqueTransactionReferenceNumber == nil {
		var ret string
		return ret
	}
	return *o.UniqueTransactionReferenceNumber
}

// GetUniqueTransactionReferenceNumberOk returns a tuple with the UniqueTransactionReferenceNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OriginalCreditSenderData) GetUniqueTransactionReferenceNumberOk() (*string, bool) {
	if o == nil || o.UniqueTransactionReferenceNumber == nil {
		return nil, false
	}
	return o.UniqueTransactionReferenceNumber, true
}

// HasUniqueTransactionReferenceNumber returns a boolean if a field has been set.
func (o *OriginalCreditSenderData) HasUniqueTransactionReferenceNumber() bool {
	if o != nil && o.UniqueTransactionReferenceNumber != nil {
		return true
	}

	return false
}

// SetUniqueTransactionReferenceNumber gets a reference to the given string and assigns it to the UniqueTransactionReferenceNumber field.
func (o *OriginalCreditSenderData) SetUniqueTransactionReferenceNumber(v string) {
	o.UniqueTransactionReferenceNumber = &v
}

func (o OriginalCreditSenderData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["funding_source"] = o.FundingSource
	}
	if o.SenderAccountNumber != nil {
		toSerialize["sender_account_number"] = o.SenderAccountNumber
	}
	if o.SenderAccountType != nil {
		toSerialize["sender_account_type"] = o.SenderAccountType
	}
	if o.SenderAddress != nil {
		toSerialize["sender_address"] = o.SenderAddress
	}
	if o.SenderCity != nil {
		toSerialize["sender_city"] = o.SenderCity
	}
	if o.SenderCountry != nil {
		toSerialize["sender_country"] = o.SenderCountry
	}
	if o.SenderName != nil {
		toSerialize["sender_name"] = o.SenderName
	}
	if o.SenderReferenceNumber != nil {
		toSerialize["sender_reference_number"] = o.SenderReferenceNumber
	}
	if o.SenderState != nil {
		toSerialize["sender_state"] = o.SenderState
	}
	if o.TransactionPurpose != nil {
		toSerialize["transaction_purpose"] = o.TransactionPurpose
	}
	if o.UniqueTransactionReferenceNumber != nil {
		toSerialize["unique_transaction_reference_number"] = o.UniqueTransactionReferenceNumber
	}
	return json.Marshal(toSerialize)
}

type NullableOriginalCreditSenderData struct {
	value *OriginalCreditSenderData
	isSet bool
}

func (v NullableOriginalCreditSenderData) Get() *OriginalCreditSenderData {
	return v.value
}

func (v *NullableOriginalCreditSenderData) Set(val *OriginalCreditSenderData) {
	v.value = val
	v.isSet = true
}

func (v NullableOriginalCreditSenderData) IsSet() bool {
	return v.isSet
}

func (v *NullableOriginalCreditSenderData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOriginalCreditSenderData(val *OriginalCreditSenderData) *NullableOriginalCreditSenderData {
	return &NullableOriginalCreditSenderData{value: val, isSet: true}
}

func (v NullableOriginalCreditSenderData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOriginalCreditSenderData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


