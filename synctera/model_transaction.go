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
	"time"
)

// Transaction struct for Transaction
type Transaction struct {
	// Account ID
	AccountId *string `json:"account_id,omitempty"`
	// Transaction amount in ISO 4217 minor currency units
	Amount *int32 `json:"amount,omitempty"`
	// Reference for the authorization
	AuthorizationRef *string `json:"authorization_ref,omitempty"`
	// Currency of the transaction. ISO 4217 alphabetic currency code
	Currency *string `json:"currency,omitempty"`
	DcSign *DcSignType `json:"dc_sign,omitempty"`
	// The effective date of the transaction (value_date)
	EffectiveDate *time.Time `json:"effective_date,omitempty"`
	// External reference from Synctera for the transaction
	ExtReference *string `json:"ext_reference,omitempty"`
	// Transaction ID
	Id *string `json:"id,omitempty"`
	// Profit center of the transaction
	ProfitCenter *string `json:"profit_center,omitempty"`
	// The status of the transaction
	Status *string `json:"status,omitempty"`
	// Transaction type
	TransactionType *string `json:"transaction_type,omitempty"`
}

// NewTransaction instantiates a new Transaction object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTransaction() *Transaction {
	this := Transaction{}
	return &this
}

// NewTransactionWithDefaults instantiates a new Transaction object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTransactionWithDefaults() *Transaction {
	this := Transaction{}
	return &this
}

// GetAccountId returns the AccountId field value if set, zero value otherwise.
func (o *Transaction) GetAccountId() string {
	if o == nil || o.AccountId == nil {
		var ret string
		return ret
	}
	return *o.AccountId
}

// GetAccountIdOk returns a tuple with the AccountId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Transaction) GetAccountIdOk() (*string, bool) {
	if o == nil || o.AccountId == nil {
		return nil, false
	}
	return o.AccountId, true
}

// HasAccountId returns a boolean if a field has been set.
func (o *Transaction) HasAccountId() bool {
	if o != nil && o.AccountId != nil {
		return true
	}

	return false
}

// SetAccountId gets a reference to the given string and assigns it to the AccountId field.
func (o *Transaction) SetAccountId(v string) {
	o.AccountId = &v
}

// GetAmount returns the Amount field value if set, zero value otherwise.
func (o *Transaction) GetAmount() int32 {
	if o == nil || o.Amount == nil {
		var ret int32
		return ret
	}
	return *o.Amount
}

// GetAmountOk returns a tuple with the Amount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Transaction) GetAmountOk() (*int32, bool) {
	if o == nil || o.Amount == nil {
		return nil, false
	}
	return o.Amount, true
}

// HasAmount returns a boolean if a field has been set.
func (o *Transaction) HasAmount() bool {
	if o != nil && o.Amount != nil {
		return true
	}

	return false
}

// SetAmount gets a reference to the given int32 and assigns it to the Amount field.
func (o *Transaction) SetAmount(v int32) {
	o.Amount = &v
}

// GetAuthorizationRef returns the AuthorizationRef field value if set, zero value otherwise.
func (o *Transaction) GetAuthorizationRef() string {
	if o == nil || o.AuthorizationRef == nil {
		var ret string
		return ret
	}
	return *o.AuthorizationRef
}

// GetAuthorizationRefOk returns a tuple with the AuthorizationRef field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Transaction) GetAuthorizationRefOk() (*string, bool) {
	if o == nil || o.AuthorizationRef == nil {
		return nil, false
	}
	return o.AuthorizationRef, true
}

// HasAuthorizationRef returns a boolean if a field has been set.
func (o *Transaction) HasAuthorizationRef() bool {
	if o != nil && o.AuthorizationRef != nil {
		return true
	}

	return false
}

// SetAuthorizationRef gets a reference to the given string and assigns it to the AuthorizationRef field.
func (o *Transaction) SetAuthorizationRef(v string) {
	o.AuthorizationRef = &v
}

// GetCurrency returns the Currency field value if set, zero value otherwise.
func (o *Transaction) GetCurrency() string {
	if o == nil || o.Currency == nil {
		var ret string
		return ret
	}
	return *o.Currency
}

// GetCurrencyOk returns a tuple with the Currency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Transaction) GetCurrencyOk() (*string, bool) {
	if o == nil || o.Currency == nil {
		return nil, false
	}
	return o.Currency, true
}

// HasCurrency returns a boolean if a field has been set.
func (o *Transaction) HasCurrency() bool {
	if o != nil && o.Currency != nil {
		return true
	}

	return false
}

// SetCurrency gets a reference to the given string and assigns it to the Currency field.
func (o *Transaction) SetCurrency(v string) {
	o.Currency = &v
}

// GetDcSign returns the DcSign field value if set, zero value otherwise.
func (o *Transaction) GetDcSign() DcSignType {
	if o == nil || o.DcSign == nil {
		var ret DcSignType
		return ret
	}
	return *o.DcSign
}

// GetDcSignOk returns a tuple with the DcSign field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Transaction) GetDcSignOk() (*DcSignType, bool) {
	if o == nil || o.DcSign == nil {
		return nil, false
	}
	return o.DcSign, true
}

// HasDcSign returns a boolean if a field has been set.
func (o *Transaction) HasDcSign() bool {
	if o != nil && o.DcSign != nil {
		return true
	}

	return false
}

// SetDcSign gets a reference to the given DcSignType and assigns it to the DcSign field.
func (o *Transaction) SetDcSign(v DcSignType) {
	o.DcSign = &v
}

// GetEffectiveDate returns the EffectiveDate field value if set, zero value otherwise.
func (o *Transaction) GetEffectiveDate() time.Time {
	if o == nil || o.EffectiveDate == nil {
		var ret time.Time
		return ret
	}
	return *o.EffectiveDate
}

// GetEffectiveDateOk returns a tuple with the EffectiveDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Transaction) GetEffectiveDateOk() (*time.Time, bool) {
	if o == nil || o.EffectiveDate == nil {
		return nil, false
	}
	return o.EffectiveDate, true
}

// HasEffectiveDate returns a boolean if a field has been set.
func (o *Transaction) HasEffectiveDate() bool {
	if o != nil && o.EffectiveDate != nil {
		return true
	}

	return false
}

// SetEffectiveDate gets a reference to the given time.Time and assigns it to the EffectiveDate field.
func (o *Transaction) SetEffectiveDate(v time.Time) {
	o.EffectiveDate = &v
}

// GetExtReference returns the ExtReference field value if set, zero value otherwise.
func (o *Transaction) GetExtReference() string {
	if o == nil || o.ExtReference == nil {
		var ret string
		return ret
	}
	return *o.ExtReference
}

// GetExtReferenceOk returns a tuple with the ExtReference field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Transaction) GetExtReferenceOk() (*string, bool) {
	if o == nil || o.ExtReference == nil {
		return nil, false
	}
	return o.ExtReference, true
}

// HasExtReference returns a boolean if a field has been set.
func (o *Transaction) HasExtReference() bool {
	if o != nil && o.ExtReference != nil {
		return true
	}

	return false
}

// SetExtReference gets a reference to the given string and assigns it to the ExtReference field.
func (o *Transaction) SetExtReference(v string) {
	o.ExtReference = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Transaction) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Transaction) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Transaction) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *Transaction) SetId(v string) {
	o.Id = &v
}

// GetProfitCenter returns the ProfitCenter field value if set, zero value otherwise.
func (o *Transaction) GetProfitCenter() string {
	if o == nil || o.ProfitCenter == nil {
		var ret string
		return ret
	}
	return *o.ProfitCenter
}

// GetProfitCenterOk returns a tuple with the ProfitCenter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Transaction) GetProfitCenterOk() (*string, bool) {
	if o == nil || o.ProfitCenter == nil {
		return nil, false
	}
	return o.ProfitCenter, true
}

// HasProfitCenter returns a boolean if a field has been set.
func (o *Transaction) HasProfitCenter() bool {
	if o != nil && o.ProfitCenter != nil {
		return true
	}

	return false
}

// SetProfitCenter gets a reference to the given string and assigns it to the ProfitCenter field.
func (o *Transaction) SetProfitCenter(v string) {
	o.ProfitCenter = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *Transaction) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Transaction) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *Transaction) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *Transaction) SetStatus(v string) {
	o.Status = &v
}

// GetTransactionType returns the TransactionType field value if set, zero value otherwise.
func (o *Transaction) GetTransactionType() string {
	if o == nil || o.TransactionType == nil {
		var ret string
		return ret
	}
	return *o.TransactionType
}

// GetTransactionTypeOk returns a tuple with the TransactionType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Transaction) GetTransactionTypeOk() (*string, bool) {
	if o == nil || o.TransactionType == nil {
		return nil, false
	}
	return o.TransactionType, true
}

// HasTransactionType returns a boolean if a field has been set.
func (o *Transaction) HasTransactionType() bool {
	if o != nil && o.TransactionType != nil {
		return true
	}

	return false
}

// SetTransactionType gets a reference to the given string and assigns it to the TransactionType field.
func (o *Transaction) SetTransactionType(v string) {
	o.TransactionType = &v
}

func (o Transaction) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AccountId != nil {
		toSerialize["account_id"] = o.AccountId
	}
	if o.Amount != nil {
		toSerialize["amount"] = o.Amount
	}
	if o.AuthorizationRef != nil {
		toSerialize["authorization_ref"] = o.AuthorizationRef
	}
	if o.Currency != nil {
		toSerialize["currency"] = o.Currency
	}
	if o.DcSign != nil {
		toSerialize["dc_sign"] = o.DcSign
	}
	if o.EffectiveDate != nil {
		toSerialize["effective_date"] = o.EffectiveDate
	}
	if o.ExtReference != nil {
		toSerialize["ext_reference"] = o.ExtReference
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.ProfitCenter != nil {
		toSerialize["profit_center"] = o.ProfitCenter
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	if o.TransactionType != nil {
		toSerialize["transaction_type"] = o.TransactionType
	}
	return json.Marshal(toSerialize)
}

type NullableTransaction struct {
	value *Transaction
	isSet bool
}

func (v NullableTransaction) Get() *Transaction {
	return v.value
}

func (v *NullableTransaction) Set(val *Transaction) {
	v.value = val
	v.isSet = true
}

func (v NullableTransaction) IsSet() bool {
	return v.isSet
}

func (v *NullableTransaction) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransaction(val *Transaction) *NullableTransaction {
	return &NullableTransaction{value: val, isSet: true}
}

func (v NullableTransaction) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTransaction) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


