/*
Synctera API

<h2>Let's build something great.</h2><p>Welcome to the official reference documentation for Synctera APIs. Our APIs are the best way to automate your company's banking needs and are designed to be easy to understand and implement.</p><p>We're continuously growing this library and what you see here is just the start, but if you need something specific or have a question, <a class='text-blue-600' href='https://synctera.com/contact' target='_blank' rel='noreferrer'>contact us</a>.</p> 

API version: 0.5.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package synctera

import (
	"encoding/json"
	"time"
)

// ExternalAccount struct for ExternalAccount
type ExternalAccount struct {
	AccountIdentifiers AccountIdentifiers `json:"account_identifiers"`
	CreationTime time.Time `json:"creation_time"`
	// The identifier for the customer associated with this account
	CustomerId string `json:"customer_id"`
	// External account unique identifier
	Id string `json:"id"`
	LastUpdatedTime time.Time `json:"last_updated_time"`
	// User-supplied metadata
	Metadata map[string]interface{} `json:"metadata"`
	// A user-meaningful name for the account
	Nickname *string `json:"nickname,omitempty"`
	RoutingIdentifiers AccountRouting `json:"routing_identifiers"`
	// The current state of the account
	Status string `json:"status"`
	// The type of the account
	Type string `json:"type"`
	Verification AccountVerification `json:"verification"`
}

// NewExternalAccount instantiates a new ExternalAccount object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewExternalAccount(accountIdentifiers AccountIdentifiers, creationTime time.Time, customerId string, id string, lastUpdatedTime time.Time, metadata map[string]interface{}, routingIdentifiers AccountRouting, status string, type_ string, verification AccountVerification) *ExternalAccount {
	this := ExternalAccount{}
	this.AccountIdentifiers = accountIdentifiers
	this.CreationTime = creationTime
	this.CustomerId = customerId
	this.Id = id
	this.LastUpdatedTime = lastUpdatedTime
	this.Metadata = metadata
	this.RoutingIdentifiers = routingIdentifiers
	this.Status = status
	this.Type = type_
	this.Verification = verification
	return &this
}

// NewExternalAccountWithDefaults instantiates a new ExternalAccount object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewExternalAccountWithDefaults() *ExternalAccount {
	this := ExternalAccount{}
	return &this
}

// GetAccountIdentifiers returns the AccountIdentifiers field value
func (o *ExternalAccount) GetAccountIdentifiers() AccountIdentifiers {
	if o == nil {
		var ret AccountIdentifiers
		return ret
	}

	return o.AccountIdentifiers
}

// GetAccountIdentifiersOk returns a tuple with the AccountIdentifiers field value
// and a boolean to check if the value has been set.
func (o *ExternalAccount) GetAccountIdentifiersOk() (*AccountIdentifiers, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.AccountIdentifiers, true
}

// SetAccountIdentifiers sets field value
func (o *ExternalAccount) SetAccountIdentifiers(v AccountIdentifiers) {
	o.AccountIdentifiers = v
}

// GetCreationTime returns the CreationTime field value
func (o *ExternalAccount) GetCreationTime() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreationTime
}

// GetCreationTimeOk returns a tuple with the CreationTime field value
// and a boolean to check if the value has been set.
func (o *ExternalAccount) GetCreationTimeOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.CreationTime, true
}

// SetCreationTime sets field value
func (o *ExternalAccount) SetCreationTime(v time.Time) {
	o.CreationTime = v
}

// GetCustomerId returns the CustomerId field value
func (o *ExternalAccount) GetCustomerId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CustomerId
}

// GetCustomerIdOk returns a tuple with the CustomerId field value
// and a boolean to check if the value has been set.
func (o *ExternalAccount) GetCustomerIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.CustomerId, true
}

// SetCustomerId sets field value
func (o *ExternalAccount) SetCustomerId(v string) {
	o.CustomerId = v
}

// GetId returns the Id field value
func (o *ExternalAccount) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ExternalAccount) GetIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ExternalAccount) SetId(v string) {
	o.Id = v
}

// GetLastUpdatedTime returns the LastUpdatedTime field value
func (o *ExternalAccount) GetLastUpdatedTime() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.LastUpdatedTime
}

// GetLastUpdatedTimeOk returns a tuple with the LastUpdatedTime field value
// and a boolean to check if the value has been set.
func (o *ExternalAccount) GetLastUpdatedTimeOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.LastUpdatedTime, true
}

// SetLastUpdatedTime sets field value
func (o *ExternalAccount) SetLastUpdatedTime(v time.Time) {
	o.LastUpdatedTime = v
}

// GetMetadata returns the Metadata field value
func (o *ExternalAccount) GetMetadata() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value
// and a boolean to check if the value has been set.
func (o *ExternalAccount) GetMetadataOk() (*map[string]interface{}, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Metadata, true
}

// SetMetadata sets field value
func (o *ExternalAccount) SetMetadata(v map[string]interface{}) {
	o.Metadata = v
}

// GetNickname returns the Nickname field value if set, zero value otherwise.
func (o *ExternalAccount) GetNickname() string {
	if o == nil || o.Nickname == nil {
		var ret string
		return ret
	}
	return *o.Nickname
}

// GetNicknameOk returns a tuple with the Nickname field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExternalAccount) GetNicknameOk() (*string, bool) {
	if o == nil || o.Nickname == nil {
		return nil, false
	}
	return o.Nickname, true
}

// HasNickname returns a boolean if a field has been set.
func (o *ExternalAccount) HasNickname() bool {
	if o != nil && o.Nickname != nil {
		return true
	}

	return false
}

// SetNickname gets a reference to the given string and assigns it to the Nickname field.
func (o *ExternalAccount) SetNickname(v string) {
	o.Nickname = &v
}

// GetRoutingIdentifiers returns the RoutingIdentifiers field value
func (o *ExternalAccount) GetRoutingIdentifiers() AccountRouting {
	if o == nil {
		var ret AccountRouting
		return ret
	}

	return o.RoutingIdentifiers
}

// GetRoutingIdentifiersOk returns a tuple with the RoutingIdentifiers field value
// and a boolean to check if the value has been set.
func (o *ExternalAccount) GetRoutingIdentifiersOk() (*AccountRouting, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.RoutingIdentifiers, true
}

// SetRoutingIdentifiers sets field value
func (o *ExternalAccount) SetRoutingIdentifiers(v AccountRouting) {
	o.RoutingIdentifiers = v
}

// GetStatus returns the Status field value
func (o *ExternalAccount) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *ExternalAccount) GetStatusOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *ExternalAccount) SetStatus(v string) {
	o.Status = v
}

// GetType returns the Type field value
func (o *ExternalAccount) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ExternalAccount) GetTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *ExternalAccount) SetType(v string) {
	o.Type = v
}

// GetVerification returns the Verification field value
func (o *ExternalAccount) GetVerification() AccountVerification {
	if o == nil {
		var ret AccountVerification
		return ret
	}

	return o.Verification
}

// GetVerificationOk returns a tuple with the Verification field value
// and a boolean to check if the value has been set.
func (o *ExternalAccount) GetVerificationOk() (*AccountVerification, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Verification, true
}

// SetVerification sets field value
func (o *ExternalAccount) SetVerification(v AccountVerification) {
	o.Verification = v
}

func (o ExternalAccount) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["account_identifiers"] = o.AccountIdentifiers
	}
	if true {
		toSerialize["creation_time"] = o.CreationTime
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
		toSerialize["metadata"] = o.Metadata
	}
	if o.Nickname != nil {
		toSerialize["nickname"] = o.Nickname
	}
	if true {
		toSerialize["routing_identifiers"] = o.RoutingIdentifiers
	}
	if true {
		toSerialize["status"] = o.Status
	}
	if true {
		toSerialize["type"] = o.Type
	}
	if true {
		toSerialize["verification"] = o.Verification
	}
	return json.Marshal(toSerialize)
}

type NullableExternalAccount struct {
	value *ExternalAccount
	isSet bool
}

func (v NullableExternalAccount) Get() *ExternalAccount {
	return v.value
}

func (v *NullableExternalAccount) Set(val *ExternalAccount) {
	v.value = val
	v.isSet = true
}

func (v NullableExternalAccount) IsSet() bool {
	return v.isSet
}

func (v *NullableExternalAccount) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExternalAccount(val *ExternalAccount) *NullableExternalAccount {
	return &NullableExternalAccount{value: val, isSet: true}
}

func (v NullableExternalAccount) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExternalAccount) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


