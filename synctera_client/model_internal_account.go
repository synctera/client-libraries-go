/*
Synctera API

<h2>Let's build something great.</h2><p>Welcome to the official reference documentation for Synctera APIs. Our APIs are the best way to automate your company's banking needs and are designed to be easy to understand and implement.</p><p>We're continuously growing this library and what you see here is just the start, but if you need something specific or have a question, <a class='text-blue-600' href='https://synctera.com/contact' target='_blank' rel='noreferrer'>contact us</a>.</p>

API version: 0.20.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package synctera_client

import (
	"encoding/json"
	"time"
)

// InternalAccount struct for InternalAccount
type InternalAccount struct {
	// Generated internal account number
	AccountNumber *string `json:"account_number,omitempty"`
	// A list of balances for internal account based on different type
	Balances []Balance `json:"balances,omitempty"`
	// Bank routing number
	BankRouting *string `json:"bank_routing,omitempty"`
	// The date and time the resource was created.
	CreationTime *time.Time `json:"creation_time,omitempty"`
	// Account currency or account settlement currency. ISO 4217 alphabetic currency code.
	Currency string `json:"currency"`
	// A user provided description for the current account
	Description *string `json:"description,omitempty"`
	// Whether the account will represent assets or liabilities
	GlType *string `json:"gl_type,omitempty"`
	// Generated ID for internal account
	Id *string `json:"id,omitempty"`
	// Is a system-controlled internal account
	IsSystemAcc *bool `json:"is_system_acc,omitempty"`
	// The date and time the resource was last updated.
	LastUpdatedTime *time.Time `json:"last_updated_time,omitempty"`
	Status          string     `json:"status"`
}

// NewInternalAccount instantiates a new InternalAccount object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInternalAccount(currency string, status string) *InternalAccount {
	this := InternalAccount{}
	this.Currency = currency
	var isSystemAcc bool = false
	this.IsSystemAcc = &isSystemAcc
	this.Status = status
	return &this
}

// NewInternalAccountWithDefaults instantiates a new InternalAccount object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInternalAccountWithDefaults() *InternalAccount {
	this := InternalAccount{}
	var isSystemAcc bool = false
	this.IsSystemAcc = &isSystemAcc
	return &this
}

// GetAccountNumber returns the AccountNumber field value if set, zero value otherwise.
func (o *InternalAccount) GetAccountNumber() string {
	if o == nil || o.AccountNumber == nil {
		var ret string
		return ret
	}
	return *o.AccountNumber
}

// GetAccountNumberOk returns a tuple with the AccountNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InternalAccount) GetAccountNumberOk() (*string, bool) {
	if o == nil || o.AccountNumber == nil {
		return nil, false
	}
	return o.AccountNumber, true
}

// HasAccountNumber returns a boolean if a field has been set.
func (o *InternalAccount) HasAccountNumber() bool {
	if o != nil && o.AccountNumber != nil {
		return true
	}

	return false
}

// SetAccountNumber gets a reference to the given string and assigns it to the AccountNumber field.
func (o *InternalAccount) SetAccountNumber(v string) {
	o.AccountNumber = &v
}

// GetBalances returns the Balances field value if set, zero value otherwise.
func (o *InternalAccount) GetBalances() []Balance {
	if o == nil || o.Balances == nil {
		var ret []Balance
		return ret
	}
	return o.Balances
}

// GetBalancesOk returns a tuple with the Balances field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InternalAccount) GetBalancesOk() ([]Balance, bool) {
	if o == nil || o.Balances == nil {
		return nil, false
	}
	return o.Balances, true
}

// HasBalances returns a boolean if a field has been set.
func (o *InternalAccount) HasBalances() bool {
	if o != nil && o.Balances != nil {
		return true
	}

	return false
}

// SetBalances gets a reference to the given []Balance and assigns it to the Balances field.
func (o *InternalAccount) SetBalances(v []Balance) {
	o.Balances = v
}

// GetBankRouting returns the BankRouting field value if set, zero value otherwise.
func (o *InternalAccount) GetBankRouting() string {
	if o == nil || o.BankRouting == nil {
		var ret string
		return ret
	}
	return *o.BankRouting
}

// GetBankRoutingOk returns a tuple with the BankRouting field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InternalAccount) GetBankRoutingOk() (*string, bool) {
	if o == nil || o.BankRouting == nil {
		return nil, false
	}
	return o.BankRouting, true
}

// HasBankRouting returns a boolean if a field has been set.
func (o *InternalAccount) HasBankRouting() bool {
	if o != nil && o.BankRouting != nil {
		return true
	}

	return false
}

// SetBankRouting gets a reference to the given string and assigns it to the BankRouting field.
func (o *InternalAccount) SetBankRouting(v string) {
	o.BankRouting = &v
}

// GetCreationTime returns the CreationTime field value if set, zero value otherwise.
func (o *InternalAccount) GetCreationTime() time.Time {
	if o == nil || o.CreationTime == nil {
		var ret time.Time
		return ret
	}
	return *o.CreationTime
}

// GetCreationTimeOk returns a tuple with the CreationTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InternalAccount) GetCreationTimeOk() (*time.Time, bool) {
	if o == nil || o.CreationTime == nil {
		return nil, false
	}
	return o.CreationTime, true
}

// HasCreationTime returns a boolean if a field has been set.
func (o *InternalAccount) HasCreationTime() bool {
	if o != nil && o.CreationTime != nil {
		return true
	}

	return false
}

// SetCreationTime gets a reference to the given time.Time and assigns it to the CreationTime field.
func (o *InternalAccount) SetCreationTime(v time.Time) {
	o.CreationTime = &v
}

// GetCurrency returns the Currency field value
func (o *InternalAccount) GetCurrency() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Currency
}

// GetCurrencyOk returns a tuple with the Currency field value
// and a boolean to check if the value has been set.
func (o *InternalAccount) GetCurrencyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Currency, true
}

// SetCurrency sets field value
func (o *InternalAccount) SetCurrency(v string) {
	o.Currency = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *InternalAccount) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InternalAccount) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *InternalAccount) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *InternalAccount) SetDescription(v string) {
	o.Description = &v
}

// GetGlType returns the GlType field value if set, zero value otherwise.
func (o *InternalAccount) GetGlType() string {
	if o == nil || o.GlType == nil {
		var ret string
		return ret
	}
	return *o.GlType
}

// GetGlTypeOk returns a tuple with the GlType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InternalAccount) GetGlTypeOk() (*string, bool) {
	if o == nil || o.GlType == nil {
		return nil, false
	}
	return o.GlType, true
}

// HasGlType returns a boolean if a field has been set.
func (o *InternalAccount) HasGlType() bool {
	if o != nil && o.GlType != nil {
		return true
	}

	return false
}

// SetGlType gets a reference to the given string and assigns it to the GlType field.
func (o *InternalAccount) SetGlType(v string) {
	o.GlType = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *InternalAccount) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InternalAccount) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *InternalAccount) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *InternalAccount) SetId(v string) {
	o.Id = &v
}

// GetIsSystemAcc returns the IsSystemAcc field value if set, zero value otherwise.
func (o *InternalAccount) GetIsSystemAcc() bool {
	if o == nil || o.IsSystemAcc == nil {
		var ret bool
		return ret
	}
	return *o.IsSystemAcc
}

// GetIsSystemAccOk returns a tuple with the IsSystemAcc field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InternalAccount) GetIsSystemAccOk() (*bool, bool) {
	if o == nil || o.IsSystemAcc == nil {
		return nil, false
	}
	return o.IsSystemAcc, true
}

// HasIsSystemAcc returns a boolean if a field has been set.
func (o *InternalAccount) HasIsSystemAcc() bool {
	if o != nil && o.IsSystemAcc != nil {
		return true
	}

	return false
}

// SetIsSystemAcc gets a reference to the given bool and assigns it to the IsSystemAcc field.
func (o *InternalAccount) SetIsSystemAcc(v bool) {
	o.IsSystemAcc = &v
}

// GetLastUpdatedTime returns the LastUpdatedTime field value if set, zero value otherwise.
func (o *InternalAccount) GetLastUpdatedTime() time.Time {
	if o == nil || o.LastUpdatedTime == nil {
		var ret time.Time
		return ret
	}
	return *o.LastUpdatedTime
}

// GetLastUpdatedTimeOk returns a tuple with the LastUpdatedTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InternalAccount) GetLastUpdatedTimeOk() (*time.Time, bool) {
	if o == nil || o.LastUpdatedTime == nil {
		return nil, false
	}
	return o.LastUpdatedTime, true
}

// HasLastUpdatedTime returns a boolean if a field has been set.
func (o *InternalAccount) HasLastUpdatedTime() bool {
	if o != nil && o.LastUpdatedTime != nil {
		return true
	}

	return false
}

// SetLastUpdatedTime gets a reference to the given time.Time and assigns it to the LastUpdatedTime field.
func (o *InternalAccount) SetLastUpdatedTime(v time.Time) {
	o.LastUpdatedTime = &v
}

// GetStatus returns the Status field value
func (o *InternalAccount) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *InternalAccount) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *InternalAccount) SetStatus(v string) {
	o.Status = v
}

func (o InternalAccount) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AccountNumber != nil {
		toSerialize["account_number"] = o.AccountNumber
	}
	if o.Balances != nil {
		toSerialize["balances"] = o.Balances
	}
	if o.BankRouting != nil {
		toSerialize["bank_routing"] = o.BankRouting
	}
	if o.CreationTime != nil {
		toSerialize["creation_time"] = o.CreationTime
	}
	if true {
		toSerialize["currency"] = o.Currency
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.GlType != nil {
		toSerialize["gl_type"] = o.GlType
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.IsSystemAcc != nil {
		toSerialize["is_system_acc"] = o.IsSystemAcc
	}
	if o.LastUpdatedTime != nil {
		toSerialize["last_updated_time"] = o.LastUpdatedTime
	}
	if true {
		toSerialize["status"] = o.Status
	}
	return json.Marshal(toSerialize)
}

type NullableInternalAccount struct {
	value *InternalAccount
	isSet bool
}

func (v NullableInternalAccount) Get() *InternalAccount {
	return v.value
}

func (v *NullableInternalAccount) Set(val *InternalAccount) {
	v.value = val
	v.isSet = true
}

func (v NullableInternalAccount) IsSet() bool {
	return v.isSet
}

func (v *NullableInternalAccount) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInternalAccount(val *InternalAccount) *NullableInternalAccount {
	return &NullableInternalAccount{value: val, isSet: true}
}

func (v NullableInternalAccount) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInternalAccount) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
