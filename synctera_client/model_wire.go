/*
Synctera API

<h2>Let's build something great.</h2><p>Welcome to the official reference documentation for Synctera APIs. Our APIs are the best way to automate your company's banking needs and are designed to be easy to understand and implement.</p><p>We're continuously growing this library and what you see here is just the start, but if you need something specific or have a question, <a class='text-blue-600' href='https://synctera.com/contact' target='_blank' rel='noreferrer'>contact us</a>.</p> 

API version: 0.19.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package synctera_client

import (
	"encoding/json"
	"time"
)

// Wire struct for Wire
type Wire struct {
	// Transfer amount in cents ($100 would be 10000)
	Amount int32 `json:"amount"`
	// Instructions intended for the financial institutions that are processing the wire.
	BankMessage *string `json:"bank_message,omitempty"`
	CreationTime time.Time `json:"creation_time"`
	// 3-character currency code
	Currency string `json:"currency"`
	// The customer UUID representing the person initiating the Wire transfer
	CustomerId string `json:"customer_id"`
	// wire ID
	Id string `json:"id"`
	LastUpdatedTime time.Time `json:"last_updated_time"`
	// Sender account ID
	OriginatingAccountId string `json:"originating_account_id"`
	// The external account uuid representing the recipient of the wire.
	ReceivingAccountId string `json:"receiving_account_id"`
	// Information from the originator to the beneficiary (recipient).
	RecipientMessage *string `json:"recipient_message,omitempty"`
	// Sender's id associated with fedwire transfer
	SenderReferenceId string `json:"sender_reference_id"`
	// The current status of the transfer
	Status string `json:"status"`
	// ID of the resulting transaction resource
	TransactionId string `json:"transaction_id"`
}

// NewWire instantiates a new Wire object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWire(amount int32, creationTime time.Time, currency string, customerId string, id string, lastUpdatedTime time.Time, originatingAccountId string, receivingAccountId string, senderReferenceId string, status string, transactionId string) *Wire {
	this := Wire{}
	this.Amount = amount
	this.CreationTime = creationTime
	this.Currency = currency
	this.CustomerId = customerId
	this.Id = id
	this.LastUpdatedTime = lastUpdatedTime
	this.OriginatingAccountId = originatingAccountId
	this.ReceivingAccountId = receivingAccountId
	this.SenderReferenceId = senderReferenceId
	this.Status = status
	this.TransactionId = transactionId
	return &this
}

// NewWireWithDefaults instantiates a new Wire object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWireWithDefaults() *Wire {
	this := Wire{}
	return &this
}

// GetAmount returns the Amount field value
func (o *Wire) GetAmount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Amount
}

// GetAmountOk returns a tuple with the Amount field value
// and a boolean to check if the value has been set.
func (o *Wire) GetAmountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Amount, true
}

// SetAmount sets field value
func (o *Wire) SetAmount(v int32) {
	o.Amount = v
}

// GetBankMessage returns the BankMessage field value if set, zero value otherwise.
func (o *Wire) GetBankMessage() string {
	if o == nil || o.BankMessage == nil {
		var ret string
		return ret
	}
	return *o.BankMessage
}

// GetBankMessageOk returns a tuple with the BankMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Wire) GetBankMessageOk() (*string, bool) {
	if o == nil || o.BankMessage == nil {
		return nil, false
	}
	return o.BankMessage, true
}

// HasBankMessage returns a boolean if a field has been set.
func (o *Wire) HasBankMessage() bool {
	if o != nil && o.BankMessage != nil {
		return true
	}

	return false
}

// SetBankMessage gets a reference to the given string and assigns it to the BankMessage field.
func (o *Wire) SetBankMessage(v string) {
	o.BankMessage = &v
}

// GetCreationTime returns the CreationTime field value
func (o *Wire) GetCreationTime() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreationTime
}

// GetCreationTimeOk returns a tuple with the CreationTime field value
// and a boolean to check if the value has been set.
func (o *Wire) GetCreationTimeOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreationTime, true
}

// SetCreationTime sets field value
func (o *Wire) SetCreationTime(v time.Time) {
	o.CreationTime = v
}

// GetCurrency returns the Currency field value
func (o *Wire) GetCurrency() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Currency
}

// GetCurrencyOk returns a tuple with the Currency field value
// and a boolean to check if the value has been set.
func (o *Wire) GetCurrencyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Currency, true
}

// SetCurrency sets field value
func (o *Wire) SetCurrency(v string) {
	o.Currency = v
}

// GetCustomerId returns the CustomerId field value
func (o *Wire) GetCustomerId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CustomerId
}

// GetCustomerIdOk returns a tuple with the CustomerId field value
// and a boolean to check if the value has been set.
func (o *Wire) GetCustomerIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CustomerId, true
}

// SetCustomerId sets field value
func (o *Wire) SetCustomerId(v string) {
	o.CustomerId = v
}

// GetId returns the Id field value
func (o *Wire) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *Wire) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *Wire) SetId(v string) {
	o.Id = v
}

// GetLastUpdatedTime returns the LastUpdatedTime field value
func (o *Wire) GetLastUpdatedTime() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.LastUpdatedTime
}

// GetLastUpdatedTimeOk returns a tuple with the LastUpdatedTime field value
// and a boolean to check if the value has been set.
func (o *Wire) GetLastUpdatedTimeOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LastUpdatedTime, true
}

// SetLastUpdatedTime sets field value
func (o *Wire) SetLastUpdatedTime(v time.Time) {
	o.LastUpdatedTime = v
}

// GetOriginatingAccountId returns the OriginatingAccountId field value
func (o *Wire) GetOriginatingAccountId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OriginatingAccountId
}

// GetOriginatingAccountIdOk returns a tuple with the OriginatingAccountId field value
// and a boolean to check if the value has been set.
func (o *Wire) GetOriginatingAccountIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OriginatingAccountId, true
}

// SetOriginatingAccountId sets field value
func (o *Wire) SetOriginatingAccountId(v string) {
	o.OriginatingAccountId = v
}

// GetReceivingAccountId returns the ReceivingAccountId field value
func (o *Wire) GetReceivingAccountId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ReceivingAccountId
}

// GetReceivingAccountIdOk returns a tuple with the ReceivingAccountId field value
// and a boolean to check if the value has been set.
func (o *Wire) GetReceivingAccountIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ReceivingAccountId, true
}

// SetReceivingAccountId sets field value
func (o *Wire) SetReceivingAccountId(v string) {
	o.ReceivingAccountId = v
}

// GetRecipientMessage returns the RecipientMessage field value if set, zero value otherwise.
func (o *Wire) GetRecipientMessage() string {
	if o == nil || o.RecipientMessage == nil {
		var ret string
		return ret
	}
	return *o.RecipientMessage
}

// GetRecipientMessageOk returns a tuple with the RecipientMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Wire) GetRecipientMessageOk() (*string, bool) {
	if o == nil || o.RecipientMessage == nil {
		return nil, false
	}
	return o.RecipientMessage, true
}

// HasRecipientMessage returns a boolean if a field has been set.
func (o *Wire) HasRecipientMessage() bool {
	if o != nil && o.RecipientMessage != nil {
		return true
	}

	return false
}

// SetRecipientMessage gets a reference to the given string and assigns it to the RecipientMessage field.
func (o *Wire) SetRecipientMessage(v string) {
	o.RecipientMessage = &v
}

// GetSenderReferenceId returns the SenderReferenceId field value
func (o *Wire) GetSenderReferenceId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SenderReferenceId
}

// GetSenderReferenceIdOk returns a tuple with the SenderReferenceId field value
// and a boolean to check if the value has been set.
func (o *Wire) GetSenderReferenceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SenderReferenceId, true
}

// SetSenderReferenceId sets field value
func (o *Wire) SetSenderReferenceId(v string) {
	o.SenderReferenceId = v
}

// GetStatus returns the Status field value
func (o *Wire) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *Wire) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *Wire) SetStatus(v string) {
	o.Status = v
}

// GetTransactionId returns the TransactionId field value
func (o *Wire) GetTransactionId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TransactionId
}

// GetTransactionIdOk returns a tuple with the TransactionId field value
// and a boolean to check if the value has been set.
func (o *Wire) GetTransactionIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TransactionId, true
}

// SetTransactionId sets field value
func (o *Wire) SetTransactionId(v string) {
	o.TransactionId = v
}

func (o Wire) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["amount"] = o.Amount
	}
	if o.BankMessage != nil {
		toSerialize["bank_message"] = o.BankMessage
	}
	if true {
		toSerialize["creation_time"] = o.CreationTime
	}
	if true {
		toSerialize["currency"] = o.Currency
	}
	if true {
		toSerialize["customer_id"] = o.CustomerId
	}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["last_updated_time"] = o.LastUpdatedTime
	}
	if true {
		toSerialize["originating_account_id"] = o.OriginatingAccountId
	}
	if true {
		toSerialize["receiving_account_id"] = o.ReceivingAccountId
	}
	if o.RecipientMessage != nil {
		toSerialize["recipient_message"] = o.RecipientMessage
	}
	if true {
		toSerialize["sender_reference_id"] = o.SenderReferenceId
	}
	if true {
		toSerialize["status"] = o.Status
	}
	if true {
		toSerialize["transaction_id"] = o.TransactionId
	}
	return json.Marshal(toSerialize)
}

type NullableWire struct {
	value *Wire
	isSet bool
}

func (v NullableWire) Get() *Wire {
	return v.value
}

func (v *NullableWire) Set(val *Wire) {
	v.value = val
	v.isSet = true
}

func (v NullableWire) IsSet() bool {
	return v.isSet
}

func (v *NullableWire) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWire(val *Wire) *NullableWire {
	return &NullableWire{value: val, isSet: true}
}

func (v NullableWire) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWire) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

