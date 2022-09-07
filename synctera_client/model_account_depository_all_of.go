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

// AccountDepositoryAllOf struct for AccountDepositoryAllOf
type AccountDepositoryAllOf struct {
	BalanceCeiling *BalanceCeiling `json:"balance_ceiling,omitempty"`
	BalanceFloor *BalanceFloor `json:"balance_floor,omitempty"`
	// Account's overdraft limit
	OverdraftLimit *int64 `json:"overdraft_limit,omitempty"`
	SpendingLimits *SpendingLimits `json:"spending_limits,omitempty"`
}

// NewAccountDepositoryAllOf instantiates a new AccountDepositoryAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccountDepositoryAllOf() *AccountDepositoryAllOf {
	this := AccountDepositoryAllOf{}
	return &this
}

// NewAccountDepositoryAllOfWithDefaults instantiates a new AccountDepositoryAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccountDepositoryAllOfWithDefaults() *AccountDepositoryAllOf {
	this := AccountDepositoryAllOf{}
	return &this
}

// GetBalanceCeiling returns the BalanceCeiling field value if set, zero value otherwise.
func (o *AccountDepositoryAllOf) GetBalanceCeiling() BalanceCeiling {
	if o == nil || o.BalanceCeiling == nil {
		var ret BalanceCeiling
		return ret
	}
	return *o.BalanceCeiling
}

// GetBalanceCeilingOk returns a tuple with the BalanceCeiling field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountDepositoryAllOf) GetBalanceCeilingOk() (*BalanceCeiling, bool) {
	if o == nil || o.BalanceCeiling == nil {
		return nil, false
	}
	return o.BalanceCeiling, true
}

// HasBalanceCeiling returns a boolean if a field has been set.
func (o *AccountDepositoryAllOf) HasBalanceCeiling() bool {
	if o != nil && o.BalanceCeiling != nil {
		return true
	}

	return false
}

// SetBalanceCeiling gets a reference to the given BalanceCeiling and assigns it to the BalanceCeiling field.
func (o *AccountDepositoryAllOf) SetBalanceCeiling(v BalanceCeiling) {
	o.BalanceCeiling = &v
}

// GetBalanceFloor returns the BalanceFloor field value if set, zero value otherwise.
func (o *AccountDepositoryAllOf) GetBalanceFloor() BalanceFloor {
	if o == nil || o.BalanceFloor == nil {
		var ret BalanceFloor
		return ret
	}
	return *o.BalanceFloor
}

// GetBalanceFloorOk returns a tuple with the BalanceFloor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountDepositoryAllOf) GetBalanceFloorOk() (*BalanceFloor, bool) {
	if o == nil || o.BalanceFloor == nil {
		return nil, false
	}
	return o.BalanceFloor, true
}

// HasBalanceFloor returns a boolean if a field has been set.
func (o *AccountDepositoryAllOf) HasBalanceFloor() bool {
	if o != nil && o.BalanceFloor != nil {
		return true
	}

	return false
}

// SetBalanceFloor gets a reference to the given BalanceFloor and assigns it to the BalanceFloor field.
func (o *AccountDepositoryAllOf) SetBalanceFloor(v BalanceFloor) {
	o.BalanceFloor = &v
}

// GetOverdraftLimit returns the OverdraftLimit field value if set, zero value otherwise.
func (o *AccountDepositoryAllOf) GetOverdraftLimit() int64 {
	if o == nil || o.OverdraftLimit == nil {
		var ret int64
		return ret
	}
	return *o.OverdraftLimit
}

// GetOverdraftLimitOk returns a tuple with the OverdraftLimit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountDepositoryAllOf) GetOverdraftLimitOk() (*int64, bool) {
	if o == nil || o.OverdraftLimit == nil {
		return nil, false
	}
	return o.OverdraftLimit, true
}

// HasOverdraftLimit returns a boolean if a field has been set.
func (o *AccountDepositoryAllOf) HasOverdraftLimit() bool {
	if o != nil && o.OverdraftLimit != nil {
		return true
	}

	return false
}

// SetOverdraftLimit gets a reference to the given int64 and assigns it to the OverdraftLimit field.
func (o *AccountDepositoryAllOf) SetOverdraftLimit(v int64) {
	o.OverdraftLimit = &v
}

// GetSpendingLimits returns the SpendingLimits field value if set, zero value otherwise.
func (o *AccountDepositoryAllOf) GetSpendingLimits() SpendingLimits {
	if o == nil || o.SpendingLimits == nil {
		var ret SpendingLimits
		return ret
	}
	return *o.SpendingLimits
}

// GetSpendingLimitsOk returns a tuple with the SpendingLimits field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountDepositoryAllOf) GetSpendingLimitsOk() (*SpendingLimits, bool) {
	if o == nil || o.SpendingLimits == nil {
		return nil, false
	}
	return o.SpendingLimits, true
}

// HasSpendingLimits returns a boolean if a field has been set.
func (o *AccountDepositoryAllOf) HasSpendingLimits() bool {
	if o != nil && o.SpendingLimits != nil {
		return true
	}

	return false
}

// SetSpendingLimits gets a reference to the given SpendingLimits and assigns it to the SpendingLimits field.
func (o *AccountDepositoryAllOf) SetSpendingLimits(v SpendingLimits) {
	o.SpendingLimits = &v
}

func (o AccountDepositoryAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.BalanceCeiling != nil {
		toSerialize["balance_ceiling"] = o.BalanceCeiling
	}
	if o.BalanceFloor != nil {
		toSerialize["balance_floor"] = o.BalanceFloor
	}
	if o.OverdraftLimit != nil {
		toSerialize["overdraft_limit"] = o.OverdraftLimit
	}
	if o.SpendingLimits != nil {
		toSerialize["spending_limits"] = o.SpendingLimits
	}
	return json.Marshal(toSerialize)
}

type NullableAccountDepositoryAllOf struct {
	value *AccountDepositoryAllOf
	isSet bool
}

func (v NullableAccountDepositoryAllOf) Get() *AccountDepositoryAllOf {
	return v.value
}

func (v *NullableAccountDepositoryAllOf) Set(val *AccountDepositoryAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableAccountDepositoryAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableAccountDepositoryAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccountDepositoryAllOf(val *AccountDepositoryAllOf) *NullableAccountDepositoryAllOf {
	return &NullableAccountDepositoryAllOf{value: val, isSet: true}
}

func (v NullableAccountDepositoryAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccountDepositoryAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

