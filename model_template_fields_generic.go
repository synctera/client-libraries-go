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

// TemplateFieldsGeneric struct for TemplateFieldsGeneric
type TemplateFieldsGeneric struct {
	AccountType AccountType `json:"account_type"`
	// Bank country of the account
	BankCountry string `json:"bank_country"`
	// Account currency. ISO 4217 alphabetic currency code
	Currency       string          `json:"currency"`
	BalanceCeiling *BalanceCeiling `json:"balance_ceiling,omitempty"`
	BalanceFloor   *BalanceFloor   `json:"balance_floor,omitempty"`
	// A list of fee resources from account product that new accounts will associate with
	FeeProductIds *[]string `json:"fee_product_ids,omitempty"`
	// Interest from account product that new accounts will associate with
	InterestProductId *string `json:"interest_product_id,omitempty"`
	// Enable ACH transaction on ledger. Default is false
	IsAchEnabled *bool `json:"is_ach_enabled,omitempty"`
	// Enable card transaction on ledger. Default is false
	IsCardEnabled *bool `json:"is_card_enabled,omitempty"`
	// Enable P2P transaction on ledger. Default is false
	IsP2pEnabled *bool `json:"is_p2p_enabled,omitempty"`
	// Account's overdraft limit. Default is 0
	OverdraftLimit *int64          `json:"overdraft_limit,omitempty"`
	SpendingLimits *SpendingLimits `json:"spending_limits,omitempty"`
}

// NewTemplateFieldsGeneric instantiates a new TemplateFieldsGeneric object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTemplateFieldsGeneric(accountType AccountType, bankCountry string, currency string) *TemplateFieldsGeneric {
	this := TemplateFieldsGeneric{}
	this.AccountType = accountType
	this.BankCountry = bankCountry
	this.Currency = currency
	return &this
}

// NewTemplateFieldsGenericWithDefaults instantiates a new TemplateFieldsGeneric object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTemplateFieldsGenericWithDefaults() *TemplateFieldsGeneric {
	this := TemplateFieldsGeneric{}
	return &this
}

// GetAccountType returns the AccountType field value
func (o *TemplateFieldsGeneric) GetAccountType() AccountType {
	if o == nil {
		var ret AccountType
		return ret
	}

	return o.AccountType
}

// GetAccountTypeOk returns a tuple with the AccountType field value
// and a boolean to check if the value has been set.
func (o *TemplateFieldsGeneric) GetAccountTypeOk() (*AccountType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AccountType, true
}

// SetAccountType sets field value
func (o *TemplateFieldsGeneric) SetAccountType(v AccountType) {
	o.AccountType = v
}

// GetBankCountry returns the BankCountry field value
func (o *TemplateFieldsGeneric) GetBankCountry() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.BankCountry
}

// GetBankCountryOk returns a tuple with the BankCountry field value
// and a boolean to check if the value has been set.
func (o *TemplateFieldsGeneric) GetBankCountryOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BankCountry, true
}

// SetBankCountry sets field value
func (o *TemplateFieldsGeneric) SetBankCountry(v string) {
	o.BankCountry = v
}

// GetCurrency returns the Currency field value
func (o *TemplateFieldsGeneric) GetCurrency() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Currency
}

// GetCurrencyOk returns a tuple with the Currency field value
// and a boolean to check if the value has been set.
func (o *TemplateFieldsGeneric) GetCurrencyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Currency, true
}

// SetCurrency sets field value
func (o *TemplateFieldsGeneric) SetCurrency(v string) {
	o.Currency = v
}

// GetBalanceCeiling returns the BalanceCeiling field value if set, zero value otherwise.
func (o *TemplateFieldsGeneric) GetBalanceCeiling() BalanceCeiling {
	if o == nil || o.BalanceCeiling == nil {
		var ret BalanceCeiling
		return ret
	}
	return *o.BalanceCeiling
}

// GetBalanceCeilingOk returns a tuple with the BalanceCeiling field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TemplateFieldsGeneric) GetBalanceCeilingOk() (*BalanceCeiling, bool) {
	if o == nil || o.BalanceCeiling == nil {
		return nil, false
	}
	return o.BalanceCeiling, true
}

// HasBalanceCeiling returns a boolean if a field has been set.
func (o *TemplateFieldsGeneric) HasBalanceCeiling() bool {
	if o != nil && o.BalanceCeiling != nil {
		return true
	}

	return false
}

// SetBalanceCeiling gets a reference to the given BalanceCeiling and assigns it to the BalanceCeiling field.
func (o *TemplateFieldsGeneric) SetBalanceCeiling(v BalanceCeiling) {
	o.BalanceCeiling = &v
}

// GetBalanceFloor returns the BalanceFloor field value if set, zero value otherwise.
func (o *TemplateFieldsGeneric) GetBalanceFloor() BalanceFloor {
	if o == nil || o.BalanceFloor == nil {
		var ret BalanceFloor
		return ret
	}
	return *o.BalanceFloor
}

// GetBalanceFloorOk returns a tuple with the BalanceFloor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TemplateFieldsGeneric) GetBalanceFloorOk() (*BalanceFloor, bool) {
	if o == nil || o.BalanceFloor == nil {
		return nil, false
	}
	return o.BalanceFloor, true
}

// HasBalanceFloor returns a boolean if a field has been set.
func (o *TemplateFieldsGeneric) HasBalanceFloor() bool {
	if o != nil && o.BalanceFloor != nil {
		return true
	}

	return false
}

// SetBalanceFloor gets a reference to the given BalanceFloor and assigns it to the BalanceFloor field.
func (o *TemplateFieldsGeneric) SetBalanceFloor(v BalanceFloor) {
	o.BalanceFloor = &v
}

// GetFeeProductIds returns the FeeProductIds field value if set, zero value otherwise.
func (o *TemplateFieldsGeneric) GetFeeProductIds() []string {
	if o == nil || o.FeeProductIds == nil {
		var ret []string
		return ret
	}
	return *o.FeeProductIds
}

// GetFeeProductIdsOk returns a tuple with the FeeProductIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TemplateFieldsGeneric) GetFeeProductIdsOk() (*[]string, bool) {
	if o == nil || o.FeeProductIds == nil {
		return nil, false
	}
	return o.FeeProductIds, true
}

// HasFeeProductIds returns a boolean if a field has been set.
func (o *TemplateFieldsGeneric) HasFeeProductIds() bool {
	if o != nil && o.FeeProductIds != nil {
		return true
	}

	return false
}

// SetFeeProductIds gets a reference to the given []string and assigns it to the FeeProductIds field.
func (o *TemplateFieldsGeneric) SetFeeProductIds(v []string) {
	o.FeeProductIds = &v
}

// GetInterestProductId returns the InterestProductId field value if set, zero value otherwise.
func (o *TemplateFieldsGeneric) GetInterestProductId() string {
	if o == nil || o.InterestProductId == nil {
		var ret string
		return ret
	}
	return *o.InterestProductId
}

// GetInterestProductIdOk returns a tuple with the InterestProductId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TemplateFieldsGeneric) GetInterestProductIdOk() (*string, bool) {
	if o == nil || o.InterestProductId == nil {
		return nil, false
	}
	return o.InterestProductId, true
}

// HasInterestProductId returns a boolean if a field has been set.
func (o *TemplateFieldsGeneric) HasInterestProductId() bool {
	if o != nil && o.InterestProductId != nil {
		return true
	}

	return false
}

// SetInterestProductId gets a reference to the given string and assigns it to the InterestProductId field.
func (o *TemplateFieldsGeneric) SetInterestProductId(v string) {
	o.InterestProductId = &v
}

// GetIsAchEnabled returns the IsAchEnabled field value if set, zero value otherwise.
func (o *TemplateFieldsGeneric) GetIsAchEnabled() bool {
	if o == nil || o.IsAchEnabled == nil {
		var ret bool
		return ret
	}
	return *o.IsAchEnabled
}

// GetIsAchEnabledOk returns a tuple with the IsAchEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TemplateFieldsGeneric) GetIsAchEnabledOk() (*bool, bool) {
	if o == nil || o.IsAchEnabled == nil {
		return nil, false
	}
	return o.IsAchEnabled, true
}

// HasIsAchEnabled returns a boolean if a field has been set.
func (o *TemplateFieldsGeneric) HasIsAchEnabled() bool {
	if o != nil && o.IsAchEnabled != nil {
		return true
	}

	return false
}

// SetIsAchEnabled gets a reference to the given bool and assigns it to the IsAchEnabled field.
func (o *TemplateFieldsGeneric) SetIsAchEnabled(v bool) {
	o.IsAchEnabled = &v
}

// GetIsCardEnabled returns the IsCardEnabled field value if set, zero value otherwise.
func (o *TemplateFieldsGeneric) GetIsCardEnabled() bool {
	if o == nil || o.IsCardEnabled == nil {
		var ret bool
		return ret
	}
	return *o.IsCardEnabled
}

// GetIsCardEnabledOk returns a tuple with the IsCardEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TemplateFieldsGeneric) GetIsCardEnabledOk() (*bool, bool) {
	if o == nil || o.IsCardEnabled == nil {
		return nil, false
	}
	return o.IsCardEnabled, true
}

// HasIsCardEnabled returns a boolean if a field has been set.
func (o *TemplateFieldsGeneric) HasIsCardEnabled() bool {
	if o != nil && o.IsCardEnabled != nil {
		return true
	}

	return false
}

// SetIsCardEnabled gets a reference to the given bool and assigns it to the IsCardEnabled field.
func (o *TemplateFieldsGeneric) SetIsCardEnabled(v bool) {
	o.IsCardEnabled = &v
}

// GetIsP2pEnabled returns the IsP2pEnabled field value if set, zero value otherwise.
func (o *TemplateFieldsGeneric) GetIsP2pEnabled() bool {
	if o == nil || o.IsP2pEnabled == nil {
		var ret bool
		return ret
	}
	return *o.IsP2pEnabled
}

// GetIsP2pEnabledOk returns a tuple with the IsP2pEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TemplateFieldsGeneric) GetIsP2pEnabledOk() (*bool, bool) {
	if o == nil || o.IsP2pEnabled == nil {
		return nil, false
	}
	return o.IsP2pEnabled, true
}

// HasIsP2pEnabled returns a boolean if a field has been set.
func (o *TemplateFieldsGeneric) HasIsP2pEnabled() bool {
	if o != nil && o.IsP2pEnabled != nil {
		return true
	}

	return false
}

// SetIsP2pEnabled gets a reference to the given bool and assigns it to the IsP2pEnabled field.
func (o *TemplateFieldsGeneric) SetIsP2pEnabled(v bool) {
	o.IsP2pEnabled = &v
}

// GetOverdraftLimit returns the OverdraftLimit field value if set, zero value otherwise.
func (o *TemplateFieldsGeneric) GetOverdraftLimit() int64 {
	if o == nil || o.OverdraftLimit == nil {
		var ret int64
		return ret
	}
	return *o.OverdraftLimit
}

// GetOverdraftLimitOk returns a tuple with the OverdraftLimit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TemplateFieldsGeneric) GetOverdraftLimitOk() (*int64, bool) {
	if o == nil || o.OverdraftLimit == nil {
		return nil, false
	}
	return o.OverdraftLimit, true
}

// HasOverdraftLimit returns a boolean if a field has been set.
func (o *TemplateFieldsGeneric) HasOverdraftLimit() bool {
	if o != nil && o.OverdraftLimit != nil {
		return true
	}

	return false
}

// SetOverdraftLimit gets a reference to the given int64 and assigns it to the OverdraftLimit field.
func (o *TemplateFieldsGeneric) SetOverdraftLimit(v int64) {
	o.OverdraftLimit = &v
}

// GetSpendingLimits returns the SpendingLimits field value if set, zero value otherwise.
func (o *TemplateFieldsGeneric) GetSpendingLimits() SpendingLimits {
	if o == nil || o.SpendingLimits == nil {
		var ret SpendingLimits
		return ret
	}
	return *o.SpendingLimits
}

// GetSpendingLimitsOk returns a tuple with the SpendingLimits field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TemplateFieldsGeneric) GetSpendingLimitsOk() (*SpendingLimits, bool) {
	if o == nil || o.SpendingLimits == nil {
		return nil, false
	}
	return o.SpendingLimits, true
}

// HasSpendingLimits returns a boolean if a field has been set.
func (o *TemplateFieldsGeneric) HasSpendingLimits() bool {
	if o != nil && o.SpendingLimits != nil {
		return true
	}

	return false
}

// SetSpendingLimits gets a reference to the given SpendingLimits and assigns it to the SpendingLimits field.
func (o *TemplateFieldsGeneric) SetSpendingLimits(v SpendingLimits) {
	o.SpendingLimits = &v
}

func (o TemplateFieldsGeneric) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["account_type"] = o.AccountType
	}
	if true {
		toSerialize["bank_country"] = o.BankCountry
	}
	if true {
		toSerialize["currency"] = o.Currency
	}
	if o.BalanceCeiling != nil {
		toSerialize["balance_ceiling"] = o.BalanceCeiling
	}
	if o.BalanceFloor != nil {
		toSerialize["balance_floor"] = o.BalanceFloor
	}
	if o.FeeProductIds != nil {
		toSerialize["fee_product_ids"] = o.FeeProductIds
	}
	if o.InterestProductId != nil {
		toSerialize["interest_product_id"] = o.InterestProductId
	}
	if o.IsAchEnabled != nil {
		toSerialize["is_ach_enabled"] = o.IsAchEnabled
	}
	if o.IsCardEnabled != nil {
		toSerialize["is_card_enabled"] = o.IsCardEnabled
	}
	if o.IsP2pEnabled != nil {
		toSerialize["is_p2p_enabled"] = o.IsP2pEnabled
	}
	if o.OverdraftLimit != nil {
		toSerialize["overdraft_limit"] = o.OverdraftLimit
	}
	if o.SpendingLimits != nil {
		toSerialize["spending_limits"] = o.SpendingLimits
	}
	return json.Marshal(toSerialize)
}

type NullableTemplateFieldsGeneric struct {
	value *TemplateFieldsGeneric
	isSet bool
}

func (v NullableTemplateFieldsGeneric) Get() *TemplateFieldsGeneric {
	return v.value
}

func (v *NullableTemplateFieldsGeneric) Set(val *TemplateFieldsGeneric) {
	v.value = val
	v.isSet = true
}

func (v NullableTemplateFieldsGeneric) IsSet() bool {
	return v.isSet
}

func (v *NullableTemplateFieldsGeneric) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTemplateFieldsGeneric(val *TemplateFieldsGeneric) *NullableTemplateFieldsGeneric {
	return &NullableTemplateFieldsGeneric{value: val, isSet: true}
}

func (v NullableTemplateFieldsGeneric) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTemplateFieldsGeneric) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
