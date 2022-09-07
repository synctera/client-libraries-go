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

// TemplateFieldsDepositoryAllOf struct for TemplateFieldsDepositoryAllOf
type TemplateFieldsDepositoryAllOf struct {
	BalanceCeiling *BalanceCeiling `json:"balance_ceiling,omitempty"`
	BalanceFloor *BalanceFloor `json:"balance_floor,omitempty"`
	// A list of fee resources from account product that new accounts will associate with
	FeeProductIds []string `json:"fee_product_ids,omitempty"`
	// Interest from account product that new accounts will associate with
	InterestProductId *string `json:"interest_product_id,omitempty"`
	// Enable ACH transaction.
	IsAchEnabled *bool `json:"is_ach_enabled,omitempty"`
	// Enable card transaction.
	IsCardEnabled *bool `json:"is_card_enabled,omitempty"`
	// Enable P2P transaction.
	IsP2pEnabled *bool `json:"is_p2p_enabled,omitempty"`
	// Enable wire transaction.
	IsWireEnabled *bool `json:"is_wire_enabled,omitempty"`
	// Account's overdraft limit. Default is 0. Unit in cents.
	OverdraftLimit *int64 `json:"overdraft_limit,omitempty"`
	SpendingLimits *SpendingLimits `json:"spending_limits,omitempty"`
}

// NewTemplateFieldsDepositoryAllOf instantiates a new TemplateFieldsDepositoryAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTemplateFieldsDepositoryAllOf() *TemplateFieldsDepositoryAllOf {
	this := TemplateFieldsDepositoryAllOf{}
	var isAchEnabled bool = false
	this.IsAchEnabled = &isAchEnabled
	var isCardEnabled bool = false
	this.IsCardEnabled = &isCardEnabled
	var isP2pEnabled bool = false
	this.IsP2pEnabled = &isP2pEnabled
	var isWireEnabled bool = false
	this.IsWireEnabled = &isWireEnabled
	return &this
}

// NewTemplateFieldsDepositoryAllOfWithDefaults instantiates a new TemplateFieldsDepositoryAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTemplateFieldsDepositoryAllOfWithDefaults() *TemplateFieldsDepositoryAllOf {
	this := TemplateFieldsDepositoryAllOf{}
	var isAchEnabled bool = false
	this.IsAchEnabled = &isAchEnabled
	var isCardEnabled bool = false
	this.IsCardEnabled = &isCardEnabled
	var isP2pEnabled bool = false
	this.IsP2pEnabled = &isP2pEnabled
	var isWireEnabled bool = false
	this.IsWireEnabled = &isWireEnabled
	return &this
}

// GetBalanceCeiling returns the BalanceCeiling field value if set, zero value otherwise.
func (o *TemplateFieldsDepositoryAllOf) GetBalanceCeiling() BalanceCeiling {
	if o == nil || o.BalanceCeiling == nil {
		var ret BalanceCeiling
		return ret
	}
	return *o.BalanceCeiling
}

// GetBalanceCeilingOk returns a tuple with the BalanceCeiling field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TemplateFieldsDepositoryAllOf) GetBalanceCeilingOk() (*BalanceCeiling, bool) {
	if o == nil || o.BalanceCeiling == nil {
		return nil, false
	}
	return o.BalanceCeiling, true
}

// HasBalanceCeiling returns a boolean if a field has been set.
func (o *TemplateFieldsDepositoryAllOf) HasBalanceCeiling() bool {
	if o != nil && o.BalanceCeiling != nil {
		return true
	}

	return false
}

// SetBalanceCeiling gets a reference to the given BalanceCeiling and assigns it to the BalanceCeiling field.
func (o *TemplateFieldsDepositoryAllOf) SetBalanceCeiling(v BalanceCeiling) {
	o.BalanceCeiling = &v
}

// GetBalanceFloor returns the BalanceFloor field value if set, zero value otherwise.
func (o *TemplateFieldsDepositoryAllOf) GetBalanceFloor() BalanceFloor {
	if o == nil || o.BalanceFloor == nil {
		var ret BalanceFloor
		return ret
	}
	return *o.BalanceFloor
}

// GetBalanceFloorOk returns a tuple with the BalanceFloor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TemplateFieldsDepositoryAllOf) GetBalanceFloorOk() (*BalanceFloor, bool) {
	if o == nil || o.BalanceFloor == nil {
		return nil, false
	}
	return o.BalanceFloor, true
}

// HasBalanceFloor returns a boolean if a field has been set.
func (o *TemplateFieldsDepositoryAllOf) HasBalanceFloor() bool {
	if o != nil && o.BalanceFloor != nil {
		return true
	}

	return false
}

// SetBalanceFloor gets a reference to the given BalanceFloor and assigns it to the BalanceFloor field.
func (o *TemplateFieldsDepositoryAllOf) SetBalanceFloor(v BalanceFloor) {
	o.BalanceFloor = &v
}

// GetFeeProductIds returns the FeeProductIds field value if set, zero value otherwise.
func (o *TemplateFieldsDepositoryAllOf) GetFeeProductIds() []string {
	if o == nil || o.FeeProductIds == nil {
		var ret []string
		return ret
	}
	return o.FeeProductIds
}

// GetFeeProductIdsOk returns a tuple with the FeeProductIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TemplateFieldsDepositoryAllOf) GetFeeProductIdsOk() ([]string, bool) {
	if o == nil || o.FeeProductIds == nil {
		return nil, false
	}
	return o.FeeProductIds, true
}

// HasFeeProductIds returns a boolean if a field has been set.
func (o *TemplateFieldsDepositoryAllOf) HasFeeProductIds() bool {
	if o != nil && o.FeeProductIds != nil {
		return true
	}

	return false
}

// SetFeeProductIds gets a reference to the given []string and assigns it to the FeeProductIds field.
func (o *TemplateFieldsDepositoryAllOf) SetFeeProductIds(v []string) {
	o.FeeProductIds = v
}

// GetInterestProductId returns the InterestProductId field value if set, zero value otherwise.
func (o *TemplateFieldsDepositoryAllOf) GetInterestProductId() string {
	if o == nil || o.InterestProductId == nil {
		var ret string
		return ret
	}
	return *o.InterestProductId
}

// GetInterestProductIdOk returns a tuple with the InterestProductId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TemplateFieldsDepositoryAllOf) GetInterestProductIdOk() (*string, bool) {
	if o == nil || o.InterestProductId == nil {
		return nil, false
	}
	return o.InterestProductId, true
}

// HasInterestProductId returns a boolean if a field has been set.
func (o *TemplateFieldsDepositoryAllOf) HasInterestProductId() bool {
	if o != nil && o.InterestProductId != nil {
		return true
	}

	return false
}

// SetInterestProductId gets a reference to the given string and assigns it to the InterestProductId field.
func (o *TemplateFieldsDepositoryAllOf) SetInterestProductId(v string) {
	o.InterestProductId = &v
}

// GetIsAchEnabled returns the IsAchEnabled field value if set, zero value otherwise.
func (o *TemplateFieldsDepositoryAllOf) GetIsAchEnabled() bool {
	if o == nil || o.IsAchEnabled == nil {
		var ret bool
		return ret
	}
	return *o.IsAchEnabled
}

// GetIsAchEnabledOk returns a tuple with the IsAchEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TemplateFieldsDepositoryAllOf) GetIsAchEnabledOk() (*bool, bool) {
	if o == nil || o.IsAchEnabled == nil {
		return nil, false
	}
	return o.IsAchEnabled, true
}

// HasIsAchEnabled returns a boolean if a field has been set.
func (o *TemplateFieldsDepositoryAllOf) HasIsAchEnabled() bool {
	if o != nil && o.IsAchEnabled != nil {
		return true
	}

	return false
}

// SetIsAchEnabled gets a reference to the given bool and assigns it to the IsAchEnabled field.
func (o *TemplateFieldsDepositoryAllOf) SetIsAchEnabled(v bool) {
	o.IsAchEnabled = &v
}

// GetIsCardEnabled returns the IsCardEnabled field value if set, zero value otherwise.
func (o *TemplateFieldsDepositoryAllOf) GetIsCardEnabled() bool {
	if o == nil || o.IsCardEnabled == nil {
		var ret bool
		return ret
	}
	return *o.IsCardEnabled
}

// GetIsCardEnabledOk returns a tuple with the IsCardEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TemplateFieldsDepositoryAllOf) GetIsCardEnabledOk() (*bool, bool) {
	if o == nil || o.IsCardEnabled == nil {
		return nil, false
	}
	return o.IsCardEnabled, true
}

// HasIsCardEnabled returns a boolean if a field has been set.
func (o *TemplateFieldsDepositoryAllOf) HasIsCardEnabled() bool {
	if o != nil && o.IsCardEnabled != nil {
		return true
	}

	return false
}

// SetIsCardEnabled gets a reference to the given bool and assigns it to the IsCardEnabled field.
func (o *TemplateFieldsDepositoryAllOf) SetIsCardEnabled(v bool) {
	o.IsCardEnabled = &v
}

// GetIsP2pEnabled returns the IsP2pEnabled field value if set, zero value otherwise.
func (o *TemplateFieldsDepositoryAllOf) GetIsP2pEnabled() bool {
	if o == nil || o.IsP2pEnabled == nil {
		var ret bool
		return ret
	}
	return *o.IsP2pEnabled
}

// GetIsP2pEnabledOk returns a tuple with the IsP2pEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TemplateFieldsDepositoryAllOf) GetIsP2pEnabledOk() (*bool, bool) {
	if o == nil || o.IsP2pEnabled == nil {
		return nil, false
	}
	return o.IsP2pEnabled, true
}

// HasIsP2pEnabled returns a boolean if a field has been set.
func (o *TemplateFieldsDepositoryAllOf) HasIsP2pEnabled() bool {
	if o != nil && o.IsP2pEnabled != nil {
		return true
	}

	return false
}

// SetIsP2pEnabled gets a reference to the given bool and assigns it to the IsP2pEnabled field.
func (o *TemplateFieldsDepositoryAllOf) SetIsP2pEnabled(v bool) {
	o.IsP2pEnabled = &v
}

// GetIsWireEnabled returns the IsWireEnabled field value if set, zero value otherwise.
func (o *TemplateFieldsDepositoryAllOf) GetIsWireEnabled() bool {
	if o == nil || o.IsWireEnabled == nil {
		var ret bool
		return ret
	}
	return *o.IsWireEnabled
}

// GetIsWireEnabledOk returns a tuple with the IsWireEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TemplateFieldsDepositoryAllOf) GetIsWireEnabledOk() (*bool, bool) {
	if o == nil || o.IsWireEnabled == nil {
		return nil, false
	}
	return o.IsWireEnabled, true
}

// HasIsWireEnabled returns a boolean if a field has been set.
func (o *TemplateFieldsDepositoryAllOf) HasIsWireEnabled() bool {
	if o != nil && o.IsWireEnabled != nil {
		return true
	}

	return false
}

// SetIsWireEnabled gets a reference to the given bool and assigns it to the IsWireEnabled field.
func (o *TemplateFieldsDepositoryAllOf) SetIsWireEnabled(v bool) {
	o.IsWireEnabled = &v
}

// GetOverdraftLimit returns the OverdraftLimit field value if set, zero value otherwise.
func (o *TemplateFieldsDepositoryAllOf) GetOverdraftLimit() int64 {
	if o == nil || o.OverdraftLimit == nil {
		var ret int64
		return ret
	}
	return *o.OverdraftLimit
}

// GetOverdraftLimitOk returns a tuple with the OverdraftLimit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TemplateFieldsDepositoryAllOf) GetOverdraftLimitOk() (*int64, bool) {
	if o == nil || o.OverdraftLimit == nil {
		return nil, false
	}
	return o.OverdraftLimit, true
}

// HasOverdraftLimit returns a boolean if a field has been set.
func (o *TemplateFieldsDepositoryAllOf) HasOverdraftLimit() bool {
	if o != nil && o.OverdraftLimit != nil {
		return true
	}

	return false
}

// SetOverdraftLimit gets a reference to the given int64 and assigns it to the OverdraftLimit field.
func (o *TemplateFieldsDepositoryAllOf) SetOverdraftLimit(v int64) {
	o.OverdraftLimit = &v
}

// GetSpendingLimits returns the SpendingLimits field value if set, zero value otherwise.
func (o *TemplateFieldsDepositoryAllOf) GetSpendingLimits() SpendingLimits {
	if o == nil || o.SpendingLimits == nil {
		var ret SpendingLimits
		return ret
	}
	return *o.SpendingLimits
}

// GetSpendingLimitsOk returns a tuple with the SpendingLimits field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TemplateFieldsDepositoryAllOf) GetSpendingLimitsOk() (*SpendingLimits, bool) {
	if o == nil || o.SpendingLimits == nil {
		return nil, false
	}
	return o.SpendingLimits, true
}

// HasSpendingLimits returns a boolean if a field has been set.
func (o *TemplateFieldsDepositoryAllOf) HasSpendingLimits() bool {
	if o != nil && o.SpendingLimits != nil {
		return true
	}

	return false
}

// SetSpendingLimits gets a reference to the given SpendingLimits and assigns it to the SpendingLimits field.
func (o *TemplateFieldsDepositoryAllOf) SetSpendingLimits(v SpendingLimits) {
	o.SpendingLimits = &v
}

func (o TemplateFieldsDepositoryAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
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
	if o.IsWireEnabled != nil {
		toSerialize["is_wire_enabled"] = o.IsWireEnabled
	}
	if o.OverdraftLimit != nil {
		toSerialize["overdraft_limit"] = o.OverdraftLimit
	}
	if o.SpendingLimits != nil {
		toSerialize["spending_limits"] = o.SpendingLimits
	}
	return json.Marshal(toSerialize)
}

type NullableTemplateFieldsDepositoryAllOf struct {
	value *TemplateFieldsDepositoryAllOf
	isSet bool
}

func (v NullableTemplateFieldsDepositoryAllOf) Get() *TemplateFieldsDepositoryAllOf {
	return v.value
}

func (v *NullableTemplateFieldsDepositoryAllOf) Set(val *TemplateFieldsDepositoryAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableTemplateFieldsDepositoryAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableTemplateFieldsDepositoryAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTemplateFieldsDepositoryAllOf(val *TemplateFieldsDepositoryAllOf) *NullableTemplateFieldsDepositoryAllOf {
	return &NullableTemplateFieldsDepositoryAllOf{value: val, isSet: true}
}

func (v NullableTemplateFieldsDepositoryAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTemplateFieldsDepositoryAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

