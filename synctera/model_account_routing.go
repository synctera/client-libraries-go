/*
Synctera API

<h2>Let's build something great.</h2><p>Welcome to the official reference documentation for Synctera APIs. Our APIs are the best way to automate your company's banking needs and are designed to be easy to understand and implement.</p><p>We're continuously growing this library and what you see here is just the start, but if you need something specific or have a question, <a class='text-blue-600' href='https://synctera.com/contact' target='_blank' rel='noreferrer'>contact us</a>.</p> 

API version: 0.5.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package synctera

import (
	"encoding/json"
)

// AccountRouting struct for AccountRouting
type AccountRouting struct {
	// The routing number used for US ACH payments. Only appears if `bank_countries` contains `US`. 
	AchRoutingNumber *string `json:"ach_routing_number,omitempty"`
	// The countries that this bank operates the account in
	BankCountries []string `json:"bank_countries"`
	// The name of the bank managing the account
	BankName string `json:"bank_name"`
	// The branch number used for EFT payments, identifying a branch at a Canadian bank. Only appears if `bank_countries` contains `CA`. 
	EftBranchNumber *string `json:"eft_branch_number,omitempty"`
	// The institution number used for EFT payments, identifying a Canadian bank. Only appears if `bank_countries` contains `CA`. 
	EftInstitutionNumber *string `json:"eft_institution_number,omitempty"`
	// The SWIFT code for the bank
	SwiftCode string `json:"swift_code"`
	// The routing number used for domestic wire payments. Only appears if `bank_countries` contains `US`. 
	WireRoutingNumber *string `json:"wire_routing_number,omitempty"`
}

// NewAccountRouting instantiates a new AccountRouting object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccountRouting(bankCountries []string, bankName string, swiftCode string) *AccountRouting {
	this := AccountRouting{}
	this.BankCountries = bankCountries
	this.BankName = bankName
	this.SwiftCode = swiftCode
	return &this
}

// NewAccountRoutingWithDefaults instantiates a new AccountRouting object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccountRoutingWithDefaults() *AccountRouting {
	this := AccountRouting{}
	return &this
}

// GetAchRoutingNumber returns the AchRoutingNumber field value if set, zero value otherwise.
func (o *AccountRouting) GetAchRoutingNumber() string {
	if o == nil || o.AchRoutingNumber == nil {
		var ret string
		return ret
	}
	return *o.AchRoutingNumber
}

// GetAchRoutingNumberOk returns a tuple with the AchRoutingNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountRouting) GetAchRoutingNumberOk() (*string, bool) {
	if o == nil || o.AchRoutingNumber == nil {
		return nil, false
	}
	return o.AchRoutingNumber, true
}

// HasAchRoutingNumber returns a boolean if a field has been set.
func (o *AccountRouting) HasAchRoutingNumber() bool {
	if o != nil && o.AchRoutingNumber != nil {
		return true
	}

	return false
}

// SetAchRoutingNumber gets a reference to the given string and assigns it to the AchRoutingNumber field.
func (o *AccountRouting) SetAchRoutingNumber(v string) {
	o.AchRoutingNumber = &v
}

// GetBankCountries returns the BankCountries field value
func (o *AccountRouting) GetBankCountries() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.BankCountries
}

// GetBankCountriesOk returns a tuple with the BankCountries field value
// and a boolean to check if the value has been set.
func (o *AccountRouting) GetBankCountriesOk() (*[]string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.BankCountries, true
}

// SetBankCountries sets field value
func (o *AccountRouting) SetBankCountries(v []string) {
	o.BankCountries = v
}

// GetBankName returns the BankName field value
func (o *AccountRouting) GetBankName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.BankName
}

// GetBankNameOk returns a tuple with the BankName field value
// and a boolean to check if the value has been set.
func (o *AccountRouting) GetBankNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.BankName, true
}

// SetBankName sets field value
func (o *AccountRouting) SetBankName(v string) {
	o.BankName = v
}

// GetEftBranchNumber returns the EftBranchNumber field value if set, zero value otherwise.
func (o *AccountRouting) GetEftBranchNumber() string {
	if o == nil || o.EftBranchNumber == nil {
		var ret string
		return ret
	}
	return *o.EftBranchNumber
}

// GetEftBranchNumberOk returns a tuple with the EftBranchNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountRouting) GetEftBranchNumberOk() (*string, bool) {
	if o == nil || o.EftBranchNumber == nil {
		return nil, false
	}
	return o.EftBranchNumber, true
}

// HasEftBranchNumber returns a boolean if a field has been set.
func (o *AccountRouting) HasEftBranchNumber() bool {
	if o != nil && o.EftBranchNumber != nil {
		return true
	}

	return false
}

// SetEftBranchNumber gets a reference to the given string and assigns it to the EftBranchNumber field.
func (o *AccountRouting) SetEftBranchNumber(v string) {
	o.EftBranchNumber = &v
}

// GetEftInstitutionNumber returns the EftInstitutionNumber field value if set, zero value otherwise.
func (o *AccountRouting) GetEftInstitutionNumber() string {
	if o == nil || o.EftInstitutionNumber == nil {
		var ret string
		return ret
	}
	return *o.EftInstitutionNumber
}

// GetEftInstitutionNumberOk returns a tuple with the EftInstitutionNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountRouting) GetEftInstitutionNumberOk() (*string, bool) {
	if o == nil || o.EftInstitutionNumber == nil {
		return nil, false
	}
	return o.EftInstitutionNumber, true
}

// HasEftInstitutionNumber returns a boolean if a field has been set.
func (o *AccountRouting) HasEftInstitutionNumber() bool {
	if o != nil && o.EftInstitutionNumber != nil {
		return true
	}

	return false
}

// SetEftInstitutionNumber gets a reference to the given string and assigns it to the EftInstitutionNumber field.
func (o *AccountRouting) SetEftInstitutionNumber(v string) {
	o.EftInstitutionNumber = &v
}

// GetSwiftCode returns the SwiftCode field value
func (o *AccountRouting) GetSwiftCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SwiftCode
}

// GetSwiftCodeOk returns a tuple with the SwiftCode field value
// and a boolean to check if the value has been set.
func (o *AccountRouting) GetSwiftCodeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.SwiftCode, true
}

// SetSwiftCode sets field value
func (o *AccountRouting) SetSwiftCode(v string) {
	o.SwiftCode = v
}

// GetWireRoutingNumber returns the WireRoutingNumber field value if set, zero value otherwise.
func (o *AccountRouting) GetWireRoutingNumber() string {
	if o == nil || o.WireRoutingNumber == nil {
		var ret string
		return ret
	}
	return *o.WireRoutingNumber
}

// GetWireRoutingNumberOk returns a tuple with the WireRoutingNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountRouting) GetWireRoutingNumberOk() (*string, bool) {
	if o == nil || o.WireRoutingNumber == nil {
		return nil, false
	}
	return o.WireRoutingNumber, true
}

// HasWireRoutingNumber returns a boolean if a field has been set.
func (o *AccountRouting) HasWireRoutingNumber() bool {
	if o != nil && o.WireRoutingNumber != nil {
		return true
	}

	return false
}

// SetWireRoutingNumber gets a reference to the given string and assigns it to the WireRoutingNumber field.
func (o *AccountRouting) SetWireRoutingNumber(v string) {
	o.WireRoutingNumber = &v
}

func (o AccountRouting) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AchRoutingNumber != nil {
		toSerialize["ach_routing_number"] = o.AchRoutingNumber
	}
	if true {
		toSerialize["bank_countries"] = o.BankCountries
	}
	if true {
		toSerialize["bank_name"] = o.BankName
	}
	if o.EftBranchNumber != nil {
		toSerialize["eft_branch_number"] = o.EftBranchNumber
	}
	if o.EftInstitutionNumber != nil {
		toSerialize["eft_institution_number"] = o.EftInstitutionNumber
	}
	if true {
		toSerialize["swift_code"] = o.SwiftCode
	}
	if o.WireRoutingNumber != nil {
		toSerialize["wire_routing_number"] = o.WireRoutingNumber
	}
	return json.Marshal(toSerialize)
}

type NullableAccountRouting struct {
	value *AccountRouting
	isSet bool
}

func (v NullableAccountRouting) Get() *AccountRouting {
	return v.value
}

func (v *NullableAccountRouting) Set(val *AccountRouting) {
	v.value = val
	v.isSet = true
}

func (v NullableAccountRouting) IsSet() bool {
	return v.isSet
}

func (v *NullableAccountRouting) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccountRouting(val *AccountRouting) *NullableAccountRouting {
	return &NullableAccountRouting{value: val, isSet: true}
}

func (v NullableAccountRouting) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccountRouting) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

