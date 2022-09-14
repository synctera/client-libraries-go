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

// AuthRequestModel struct for AuthRequestModel
type AuthRequestModel struct {
	// The amount of the transaction in the smallest whole denomination of the applicable currency (eg. For USD use cents)
	Amount             int32               `json:"amount"`
	CardAcceptor       *CardAcceptorModel  `json:"card_acceptor,omitempty"`
	CardId             string              `json:"card_id"`
	CardOptions        *CardOptions        `json:"card_options,omitempty"`
	CashBackAmount     *int32              `json:"cash_back_amount,omitempty"`
	IsPreAuth          *bool               `json:"is_pre_auth,omitempty"`
	Mid                string              `json:"mid"`
	NetworkFees        []NetworkFeeModel   `json:"network_fees,omitempty"`
	Pin                *string             `json:"pin,omitempty"`
	TransactionOptions *TransactionOptions `json:"transaction_options,omitempty"`
}

// NewAuthRequestModel instantiates a new AuthRequestModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAuthRequestModel(amount int32, cardId string, mid string) *AuthRequestModel {
	this := AuthRequestModel{}
	this.Amount = amount
	this.CardId = cardId
	var isPreAuth bool = false
	this.IsPreAuth = &isPreAuth
	this.Mid = mid
	return &this
}

// NewAuthRequestModelWithDefaults instantiates a new AuthRequestModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAuthRequestModelWithDefaults() *AuthRequestModel {
	this := AuthRequestModel{}
	var isPreAuth bool = false
	this.IsPreAuth = &isPreAuth
	return &this
}

// GetAmount returns the Amount field value
func (o *AuthRequestModel) GetAmount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Amount
}

// GetAmountOk returns a tuple with the Amount field value
// and a boolean to check if the value has been set.
func (o *AuthRequestModel) GetAmountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Amount, true
}

// SetAmount sets field value
func (o *AuthRequestModel) SetAmount(v int32) {
	o.Amount = v
}

// GetCardAcceptor returns the CardAcceptor field value if set, zero value otherwise.
func (o *AuthRequestModel) GetCardAcceptor() CardAcceptorModel {
	if o == nil || o.CardAcceptor == nil {
		var ret CardAcceptorModel
		return ret
	}
	return *o.CardAcceptor
}

// GetCardAcceptorOk returns a tuple with the CardAcceptor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthRequestModel) GetCardAcceptorOk() (*CardAcceptorModel, bool) {
	if o == nil || o.CardAcceptor == nil {
		return nil, false
	}
	return o.CardAcceptor, true
}

// HasCardAcceptor returns a boolean if a field has been set.
func (o *AuthRequestModel) HasCardAcceptor() bool {
	if o != nil && o.CardAcceptor != nil {
		return true
	}

	return false
}

// SetCardAcceptor gets a reference to the given CardAcceptorModel and assigns it to the CardAcceptor field.
func (o *AuthRequestModel) SetCardAcceptor(v CardAcceptorModel) {
	o.CardAcceptor = &v
}

// GetCardId returns the CardId field value
func (o *AuthRequestModel) GetCardId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CardId
}

// GetCardIdOk returns a tuple with the CardId field value
// and a boolean to check if the value has been set.
func (o *AuthRequestModel) GetCardIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CardId, true
}

// SetCardId sets field value
func (o *AuthRequestModel) SetCardId(v string) {
	o.CardId = v
}

// GetCardOptions returns the CardOptions field value if set, zero value otherwise.
func (o *AuthRequestModel) GetCardOptions() CardOptions {
	if o == nil || o.CardOptions == nil {
		var ret CardOptions
		return ret
	}
	return *o.CardOptions
}

// GetCardOptionsOk returns a tuple with the CardOptions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthRequestModel) GetCardOptionsOk() (*CardOptions, bool) {
	if o == nil || o.CardOptions == nil {
		return nil, false
	}
	return o.CardOptions, true
}

// HasCardOptions returns a boolean if a field has been set.
func (o *AuthRequestModel) HasCardOptions() bool {
	if o != nil && o.CardOptions != nil {
		return true
	}

	return false
}

// SetCardOptions gets a reference to the given CardOptions and assigns it to the CardOptions field.
func (o *AuthRequestModel) SetCardOptions(v CardOptions) {
	o.CardOptions = &v
}

// GetCashBackAmount returns the CashBackAmount field value if set, zero value otherwise.
func (o *AuthRequestModel) GetCashBackAmount() int32 {
	if o == nil || o.CashBackAmount == nil {
		var ret int32
		return ret
	}
	return *o.CashBackAmount
}

// GetCashBackAmountOk returns a tuple with the CashBackAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthRequestModel) GetCashBackAmountOk() (*int32, bool) {
	if o == nil || o.CashBackAmount == nil {
		return nil, false
	}
	return o.CashBackAmount, true
}

// HasCashBackAmount returns a boolean if a field has been set.
func (o *AuthRequestModel) HasCashBackAmount() bool {
	if o != nil && o.CashBackAmount != nil {
		return true
	}

	return false
}

// SetCashBackAmount gets a reference to the given int32 and assigns it to the CashBackAmount field.
func (o *AuthRequestModel) SetCashBackAmount(v int32) {
	o.CashBackAmount = &v
}

// GetIsPreAuth returns the IsPreAuth field value if set, zero value otherwise.
func (o *AuthRequestModel) GetIsPreAuth() bool {
	if o == nil || o.IsPreAuth == nil {
		var ret bool
		return ret
	}
	return *o.IsPreAuth
}

// GetIsPreAuthOk returns a tuple with the IsPreAuth field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthRequestModel) GetIsPreAuthOk() (*bool, bool) {
	if o == nil || o.IsPreAuth == nil {
		return nil, false
	}
	return o.IsPreAuth, true
}

// HasIsPreAuth returns a boolean if a field has been set.
func (o *AuthRequestModel) HasIsPreAuth() bool {
	if o != nil && o.IsPreAuth != nil {
		return true
	}

	return false
}

// SetIsPreAuth gets a reference to the given bool and assigns it to the IsPreAuth field.
func (o *AuthRequestModel) SetIsPreAuth(v bool) {
	o.IsPreAuth = &v
}

// GetMid returns the Mid field value
func (o *AuthRequestModel) GetMid() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Mid
}

// GetMidOk returns a tuple with the Mid field value
// and a boolean to check if the value has been set.
func (o *AuthRequestModel) GetMidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Mid, true
}

// SetMid sets field value
func (o *AuthRequestModel) SetMid(v string) {
	o.Mid = v
}

// GetNetworkFees returns the NetworkFees field value if set, zero value otherwise.
func (o *AuthRequestModel) GetNetworkFees() []NetworkFeeModel {
	if o == nil || o.NetworkFees == nil {
		var ret []NetworkFeeModel
		return ret
	}
	return o.NetworkFees
}

// GetNetworkFeesOk returns a tuple with the NetworkFees field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthRequestModel) GetNetworkFeesOk() ([]NetworkFeeModel, bool) {
	if o == nil || o.NetworkFees == nil {
		return nil, false
	}
	return o.NetworkFees, true
}

// HasNetworkFees returns a boolean if a field has been set.
func (o *AuthRequestModel) HasNetworkFees() bool {
	if o != nil && o.NetworkFees != nil {
		return true
	}

	return false
}

// SetNetworkFees gets a reference to the given []NetworkFeeModel and assigns it to the NetworkFees field.
func (o *AuthRequestModel) SetNetworkFees(v []NetworkFeeModel) {
	o.NetworkFees = v
}

// GetPin returns the Pin field value if set, zero value otherwise.
func (o *AuthRequestModel) GetPin() string {
	if o == nil || o.Pin == nil {
		var ret string
		return ret
	}
	return *o.Pin
}

// GetPinOk returns a tuple with the Pin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthRequestModel) GetPinOk() (*string, bool) {
	if o == nil || o.Pin == nil {
		return nil, false
	}
	return o.Pin, true
}

// HasPin returns a boolean if a field has been set.
func (o *AuthRequestModel) HasPin() bool {
	if o != nil && o.Pin != nil {
		return true
	}

	return false
}

// SetPin gets a reference to the given string and assigns it to the Pin field.
func (o *AuthRequestModel) SetPin(v string) {
	o.Pin = &v
}

// GetTransactionOptions returns the TransactionOptions field value if set, zero value otherwise.
func (o *AuthRequestModel) GetTransactionOptions() TransactionOptions {
	if o == nil || o.TransactionOptions == nil {
		var ret TransactionOptions
		return ret
	}
	return *o.TransactionOptions
}

// GetTransactionOptionsOk returns a tuple with the TransactionOptions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthRequestModel) GetTransactionOptionsOk() (*TransactionOptions, bool) {
	if o == nil || o.TransactionOptions == nil {
		return nil, false
	}
	return o.TransactionOptions, true
}

// HasTransactionOptions returns a boolean if a field has been set.
func (o *AuthRequestModel) HasTransactionOptions() bool {
	if o != nil && o.TransactionOptions != nil {
		return true
	}

	return false
}

// SetTransactionOptions gets a reference to the given TransactionOptions and assigns it to the TransactionOptions field.
func (o *AuthRequestModel) SetTransactionOptions(v TransactionOptions) {
	o.TransactionOptions = &v
}

func (o AuthRequestModel) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["amount"] = o.Amount
	}
	if o.CardAcceptor != nil {
		toSerialize["card_acceptor"] = o.CardAcceptor
	}
	if true {
		toSerialize["card_id"] = o.CardId
	}
	if o.CardOptions != nil {
		toSerialize["card_options"] = o.CardOptions
	}
	if o.CashBackAmount != nil {
		toSerialize["cash_back_amount"] = o.CashBackAmount
	}
	if o.IsPreAuth != nil {
		toSerialize["is_pre_auth"] = o.IsPreAuth
	}
	if true {
		toSerialize["mid"] = o.Mid
	}
	if o.NetworkFees != nil {
		toSerialize["network_fees"] = o.NetworkFees
	}
	if o.Pin != nil {
		toSerialize["pin"] = o.Pin
	}
	if o.TransactionOptions != nil {
		toSerialize["transaction_options"] = o.TransactionOptions
	}
	return json.Marshal(toSerialize)
}

type NullableAuthRequestModel struct {
	value *AuthRequestModel
	isSet bool
}

func (v NullableAuthRequestModel) Get() *AuthRequestModel {
	return v.value
}

func (v *NullableAuthRequestModel) Set(val *AuthRequestModel) {
	v.value = val
	v.isSet = true
}

func (v NullableAuthRequestModel) IsSet() bool {
	return v.isSet
}

func (v *NullableAuthRequestModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuthRequestModel(val *AuthRequestModel) *NullableAuthRequestModel {
	return &NullableAuthRequestModel{value: val, isSet: true}
}

func (v NullableAuthRequestModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAuthRequestModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
