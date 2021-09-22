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

// AchOutgoing Represents an outgoing ACH
type AchOutgoing struct {
	// ID of the source account
	AccountId string `json:"account_id"`
	// Amount to transfer in ISO 4217 minor currency units
	Amount int32 `json:"amount"`
	// ISO 4217 alphabetic currency code of the transfer amount
	Currency string `json:"currency"`
	// The date the transfer executes (default today)
	ExecutionDate *string `json:"execution_date,omitempty"`
	// Outgoing ACH
	Id *string `json:"id,omitempty"`
	// Do not perform a fraud check
	NoFraudCheck *bool `json:"no_fraud_check,omitempty"`
	// Overwrite non mandatory posting checks. Mandatory checks will still be processed
	OverwriteChecks *bool `json:"overwrite_checks,omitempty"`
	// The name of the recipient
	RecipientName string `json:"recipient_name"`
	RecurringData *RecurrenceData `json:"recurring_data,omitempty"`
	// Reference information for the payment
	ReferenceInfo *string `json:"reference_info,omitempty"`
	// The account number of the destination account
	TargetAccountNo string `json:"target_account_no"`
	// The routing number of the destination account
	TargetAccountRouting string `json:"target_account_routing"`
	// The ISO-3166-1 Alpha-2 country code in which the target account is registered (default US)
	TargetBankCountry string `json:"target_bank_country"`
	// The type of transaction (DEBIT/CREDIT) for originating account
	TransactionDirection *string `json:"transaction_direction,omitempty"`
}

// NewAchOutgoing instantiates a new AchOutgoing object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAchOutgoing(accountId string, amount int32, currency string, recipientName string, targetAccountNo string, targetAccountRouting string, targetBankCountry string) *AchOutgoing {
	this := AchOutgoing{}
	this.AccountId = accountId
	this.Amount = amount
	this.Currency = currency
	this.RecipientName = recipientName
	this.TargetAccountNo = targetAccountNo
	this.TargetAccountRouting = targetAccountRouting
	this.TargetBankCountry = targetBankCountry
	return &this
}

// NewAchOutgoingWithDefaults instantiates a new AchOutgoing object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAchOutgoingWithDefaults() *AchOutgoing {
	this := AchOutgoing{}
	var targetBankCountry string = "US"
	this.TargetBankCountry = targetBankCountry
	return &this
}

// GetAccountId returns the AccountId field value
func (o *AchOutgoing) GetAccountId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AccountId
}

// GetAccountIdOk returns a tuple with the AccountId field value
// and a boolean to check if the value has been set.
func (o *AchOutgoing) GetAccountIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.AccountId, true
}

// SetAccountId sets field value
func (o *AchOutgoing) SetAccountId(v string) {
	o.AccountId = v
}

// GetAmount returns the Amount field value
func (o *AchOutgoing) GetAmount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Amount
}

// GetAmountOk returns a tuple with the Amount field value
// and a boolean to check if the value has been set.
func (o *AchOutgoing) GetAmountOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Amount, true
}

// SetAmount sets field value
func (o *AchOutgoing) SetAmount(v int32) {
	o.Amount = v
}

// GetCurrency returns the Currency field value
func (o *AchOutgoing) GetCurrency() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Currency
}

// GetCurrencyOk returns a tuple with the Currency field value
// and a boolean to check if the value has been set.
func (o *AchOutgoing) GetCurrencyOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Currency, true
}

// SetCurrency sets field value
func (o *AchOutgoing) SetCurrency(v string) {
	o.Currency = v
}

// GetExecutionDate returns the ExecutionDate field value if set, zero value otherwise.
func (o *AchOutgoing) GetExecutionDate() string {
	if o == nil || o.ExecutionDate == nil {
		var ret string
		return ret
	}
	return *o.ExecutionDate
}

// GetExecutionDateOk returns a tuple with the ExecutionDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AchOutgoing) GetExecutionDateOk() (*string, bool) {
	if o == nil || o.ExecutionDate == nil {
		return nil, false
	}
	return o.ExecutionDate, true
}

// HasExecutionDate returns a boolean if a field has been set.
func (o *AchOutgoing) HasExecutionDate() bool {
	if o != nil && o.ExecutionDate != nil {
		return true
	}

	return false
}

// SetExecutionDate gets a reference to the given string and assigns it to the ExecutionDate field.
func (o *AchOutgoing) SetExecutionDate(v string) {
	o.ExecutionDate = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *AchOutgoing) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AchOutgoing) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *AchOutgoing) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *AchOutgoing) SetId(v string) {
	o.Id = &v
}

// GetNoFraudCheck returns the NoFraudCheck field value if set, zero value otherwise.
func (o *AchOutgoing) GetNoFraudCheck() bool {
	if o == nil || o.NoFraudCheck == nil {
		var ret bool
		return ret
	}
	return *o.NoFraudCheck
}

// GetNoFraudCheckOk returns a tuple with the NoFraudCheck field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AchOutgoing) GetNoFraudCheckOk() (*bool, bool) {
	if o == nil || o.NoFraudCheck == nil {
		return nil, false
	}
	return o.NoFraudCheck, true
}

// HasNoFraudCheck returns a boolean if a field has been set.
func (o *AchOutgoing) HasNoFraudCheck() bool {
	if o != nil && o.NoFraudCheck != nil {
		return true
	}

	return false
}

// SetNoFraudCheck gets a reference to the given bool and assigns it to the NoFraudCheck field.
func (o *AchOutgoing) SetNoFraudCheck(v bool) {
	o.NoFraudCheck = &v
}

// GetOverwriteChecks returns the OverwriteChecks field value if set, zero value otherwise.
func (o *AchOutgoing) GetOverwriteChecks() bool {
	if o == nil || o.OverwriteChecks == nil {
		var ret bool
		return ret
	}
	return *o.OverwriteChecks
}

// GetOverwriteChecksOk returns a tuple with the OverwriteChecks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AchOutgoing) GetOverwriteChecksOk() (*bool, bool) {
	if o == nil || o.OverwriteChecks == nil {
		return nil, false
	}
	return o.OverwriteChecks, true
}

// HasOverwriteChecks returns a boolean if a field has been set.
func (o *AchOutgoing) HasOverwriteChecks() bool {
	if o != nil && o.OverwriteChecks != nil {
		return true
	}

	return false
}

// SetOverwriteChecks gets a reference to the given bool and assigns it to the OverwriteChecks field.
func (o *AchOutgoing) SetOverwriteChecks(v bool) {
	o.OverwriteChecks = &v
}

// GetRecipientName returns the RecipientName field value
func (o *AchOutgoing) GetRecipientName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RecipientName
}

// GetRecipientNameOk returns a tuple with the RecipientName field value
// and a boolean to check if the value has been set.
func (o *AchOutgoing) GetRecipientNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.RecipientName, true
}

// SetRecipientName sets field value
func (o *AchOutgoing) SetRecipientName(v string) {
	o.RecipientName = v
}

// GetRecurringData returns the RecurringData field value if set, zero value otherwise.
func (o *AchOutgoing) GetRecurringData() RecurrenceData {
	if o == nil || o.RecurringData == nil {
		var ret RecurrenceData
		return ret
	}
	return *o.RecurringData
}

// GetRecurringDataOk returns a tuple with the RecurringData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AchOutgoing) GetRecurringDataOk() (*RecurrenceData, bool) {
	if o == nil || o.RecurringData == nil {
		return nil, false
	}
	return o.RecurringData, true
}

// HasRecurringData returns a boolean if a field has been set.
func (o *AchOutgoing) HasRecurringData() bool {
	if o != nil && o.RecurringData != nil {
		return true
	}

	return false
}

// SetRecurringData gets a reference to the given RecurrenceData and assigns it to the RecurringData field.
func (o *AchOutgoing) SetRecurringData(v RecurrenceData) {
	o.RecurringData = &v
}

// GetReferenceInfo returns the ReferenceInfo field value if set, zero value otherwise.
func (o *AchOutgoing) GetReferenceInfo() string {
	if o == nil || o.ReferenceInfo == nil {
		var ret string
		return ret
	}
	return *o.ReferenceInfo
}

// GetReferenceInfoOk returns a tuple with the ReferenceInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AchOutgoing) GetReferenceInfoOk() (*string, bool) {
	if o == nil || o.ReferenceInfo == nil {
		return nil, false
	}
	return o.ReferenceInfo, true
}

// HasReferenceInfo returns a boolean if a field has been set.
func (o *AchOutgoing) HasReferenceInfo() bool {
	if o != nil && o.ReferenceInfo != nil {
		return true
	}

	return false
}

// SetReferenceInfo gets a reference to the given string and assigns it to the ReferenceInfo field.
func (o *AchOutgoing) SetReferenceInfo(v string) {
	o.ReferenceInfo = &v
}

// GetTargetAccountNo returns the TargetAccountNo field value
func (o *AchOutgoing) GetTargetAccountNo() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TargetAccountNo
}

// GetTargetAccountNoOk returns a tuple with the TargetAccountNo field value
// and a boolean to check if the value has been set.
func (o *AchOutgoing) GetTargetAccountNoOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.TargetAccountNo, true
}

// SetTargetAccountNo sets field value
func (o *AchOutgoing) SetTargetAccountNo(v string) {
	o.TargetAccountNo = v
}

// GetTargetAccountRouting returns the TargetAccountRouting field value
func (o *AchOutgoing) GetTargetAccountRouting() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TargetAccountRouting
}

// GetTargetAccountRoutingOk returns a tuple with the TargetAccountRouting field value
// and a boolean to check if the value has been set.
func (o *AchOutgoing) GetTargetAccountRoutingOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.TargetAccountRouting, true
}

// SetTargetAccountRouting sets field value
func (o *AchOutgoing) SetTargetAccountRouting(v string) {
	o.TargetAccountRouting = v
}

// GetTargetBankCountry returns the TargetBankCountry field value
func (o *AchOutgoing) GetTargetBankCountry() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TargetBankCountry
}

// GetTargetBankCountryOk returns a tuple with the TargetBankCountry field value
// and a boolean to check if the value has been set.
func (o *AchOutgoing) GetTargetBankCountryOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.TargetBankCountry, true
}

// SetTargetBankCountry sets field value
func (o *AchOutgoing) SetTargetBankCountry(v string) {
	o.TargetBankCountry = v
}

// GetTransactionDirection returns the TransactionDirection field value if set, zero value otherwise.
func (o *AchOutgoing) GetTransactionDirection() string {
	if o == nil || o.TransactionDirection == nil {
		var ret string
		return ret
	}
	return *o.TransactionDirection
}

// GetTransactionDirectionOk returns a tuple with the TransactionDirection field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AchOutgoing) GetTransactionDirectionOk() (*string, bool) {
	if o == nil || o.TransactionDirection == nil {
		return nil, false
	}
	return o.TransactionDirection, true
}

// HasTransactionDirection returns a boolean if a field has been set.
func (o *AchOutgoing) HasTransactionDirection() bool {
	if o != nil && o.TransactionDirection != nil {
		return true
	}

	return false
}

// SetTransactionDirection gets a reference to the given string and assigns it to the TransactionDirection field.
func (o *AchOutgoing) SetTransactionDirection(v string) {
	o.TransactionDirection = &v
}

func (o AchOutgoing) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["account_id"] = o.AccountId
	}
	if true {
		toSerialize["amount"] = o.Amount
	}
	if true {
		toSerialize["currency"] = o.Currency
	}
	if o.ExecutionDate != nil {
		toSerialize["execution_date"] = o.ExecutionDate
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.NoFraudCheck != nil {
		toSerialize["no_fraud_check"] = o.NoFraudCheck
	}
	if o.OverwriteChecks != nil {
		toSerialize["overwrite_checks"] = o.OverwriteChecks
	}
	if true {
		toSerialize["recipient_name"] = o.RecipientName
	}
	if o.RecurringData != nil {
		toSerialize["recurring_data"] = o.RecurringData
	}
	if o.ReferenceInfo != nil {
		toSerialize["reference_info"] = o.ReferenceInfo
	}
	if true {
		toSerialize["target_account_no"] = o.TargetAccountNo
	}
	if true {
		toSerialize["target_account_routing"] = o.TargetAccountRouting
	}
	if true {
		toSerialize["target_bank_country"] = o.TargetBankCountry
	}
	if o.TransactionDirection != nil {
		toSerialize["transaction_direction"] = o.TransactionDirection
	}
	return json.Marshal(toSerialize)
}

type NullableAchOutgoing struct {
	value *AchOutgoing
	isSet bool
}

func (v NullableAchOutgoing) Get() *AchOutgoing {
	return v.value
}

func (v *NullableAchOutgoing) Set(val *AchOutgoing) {
	v.value = val
	v.isSet = true
}

func (v NullableAchOutgoing) IsSet() bool {
	return v.isSet
}

func (v *NullableAchOutgoing) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAchOutgoing(val *AchOutgoing) *NullableAchOutgoing {
	return &NullableAchOutgoing{value: val, isSet: true}
}

func (v NullableAchOutgoing) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAchOutgoing) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


