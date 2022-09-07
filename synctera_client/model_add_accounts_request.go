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

// AddAccountsRequest struct for AddAccountsRequest
type AddAccountsRequest struct {
	AccountIdentifiers AddAccountsRequestAccountIdentifiers `json:"account_identifiers"`
	// The names of the account owners.
	AccountOwnerNames []string `json:"account_owner_names"`
	// The identifier for the business customer associated with this external account. Exactly one of `business_id` or `customer_id` must be specified. 
	BusinessId *string `json:"business_id,omitempty"`
	// The identifier for the personal customer associated with this external account. Exactly one of `customer_id` or `business_id` must be specified. 
	CustomerId *string `json:"customer_id,omitempty"`
	CustomerType ExtAccountCustomerType `json:"customer_type"`
	// User-supplied metadata
	Metadata map[string]interface{} `json:"metadata,omitempty"`
	// A user-meaningful name for the account
	Nickname *string `json:"nickname,omitempty"`
	RoutingIdentifiers AddAccountsRequestRoutingIdentifiers `json:"routing_identifiers"`
	// The type of the account
	Type string `json:"type"`
	// The ID of the vendor account, will be empty for MANUAL vendor
	VendorAccountId *string `json:"vendor_account_id,omitempty"`
	Verification NullableAccountVerification `json:"verification,omitempty"`
}

// NewAddAccountsRequest instantiates a new AddAccountsRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddAccountsRequest(accountIdentifiers AddAccountsRequestAccountIdentifiers, accountOwnerNames []string, customerType ExtAccountCustomerType, routingIdentifiers AddAccountsRequestRoutingIdentifiers, type_ string) *AddAccountsRequest {
	this := AddAccountsRequest{}
	this.AccountIdentifiers = accountIdentifiers
	this.AccountOwnerNames = accountOwnerNames
	this.CustomerType = customerType
	this.RoutingIdentifiers = routingIdentifiers
	this.Type = type_
	return &this
}

// NewAddAccountsRequestWithDefaults instantiates a new AddAccountsRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddAccountsRequestWithDefaults() *AddAccountsRequest {
	this := AddAccountsRequest{}
	return &this
}

// GetAccountIdentifiers returns the AccountIdentifiers field value
func (o *AddAccountsRequest) GetAccountIdentifiers() AddAccountsRequestAccountIdentifiers {
	if o == nil {
		var ret AddAccountsRequestAccountIdentifiers
		return ret
	}

	return o.AccountIdentifiers
}

// GetAccountIdentifiersOk returns a tuple with the AccountIdentifiers field value
// and a boolean to check if the value has been set.
func (o *AddAccountsRequest) GetAccountIdentifiersOk() (*AddAccountsRequestAccountIdentifiers, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AccountIdentifiers, true
}

// SetAccountIdentifiers sets field value
func (o *AddAccountsRequest) SetAccountIdentifiers(v AddAccountsRequestAccountIdentifiers) {
	o.AccountIdentifiers = v
}

// GetAccountOwnerNames returns the AccountOwnerNames field value
func (o *AddAccountsRequest) GetAccountOwnerNames() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.AccountOwnerNames
}

// GetAccountOwnerNamesOk returns a tuple with the AccountOwnerNames field value
// and a boolean to check if the value has been set.
func (o *AddAccountsRequest) GetAccountOwnerNamesOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AccountOwnerNames, true
}

// SetAccountOwnerNames sets field value
func (o *AddAccountsRequest) SetAccountOwnerNames(v []string) {
	o.AccountOwnerNames = v
}

// GetBusinessId returns the BusinessId field value if set, zero value otherwise.
func (o *AddAccountsRequest) GetBusinessId() string {
	if o == nil || o.BusinessId == nil {
		var ret string
		return ret
	}
	return *o.BusinessId
}

// GetBusinessIdOk returns a tuple with the BusinessId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddAccountsRequest) GetBusinessIdOk() (*string, bool) {
	if o == nil || o.BusinessId == nil {
		return nil, false
	}
	return o.BusinessId, true
}

// HasBusinessId returns a boolean if a field has been set.
func (o *AddAccountsRequest) HasBusinessId() bool {
	if o != nil && o.BusinessId != nil {
		return true
	}

	return false
}

// SetBusinessId gets a reference to the given string and assigns it to the BusinessId field.
func (o *AddAccountsRequest) SetBusinessId(v string) {
	o.BusinessId = &v
}

// GetCustomerId returns the CustomerId field value if set, zero value otherwise.
func (o *AddAccountsRequest) GetCustomerId() string {
	if o == nil || o.CustomerId == nil {
		var ret string
		return ret
	}
	return *o.CustomerId
}

// GetCustomerIdOk returns a tuple with the CustomerId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddAccountsRequest) GetCustomerIdOk() (*string, bool) {
	if o == nil || o.CustomerId == nil {
		return nil, false
	}
	return o.CustomerId, true
}

// HasCustomerId returns a boolean if a field has been set.
func (o *AddAccountsRequest) HasCustomerId() bool {
	if o != nil && o.CustomerId != nil {
		return true
	}

	return false
}

// SetCustomerId gets a reference to the given string and assigns it to the CustomerId field.
func (o *AddAccountsRequest) SetCustomerId(v string) {
	o.CustomerId = &v
}

// GetCustomerType returns the CustomerType field value
func (o *AddAccountsRequest) GetCustomerType() ExtAccountCustomerType {
	if o == nil {
		var ret ExtAccountCustomerType
		return ret
	}

	return o.CustomerType
}

// GetCustomerTypeOk returns a tuple with the CustomerType field value
// and a boolean to check if the value has been set.
func (o *AddAccountsRequest) GetCustomerTypeOk() (*ExtAccountCustomerType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CustomerType, true
}

// SetCustomerType sets field value
func (o *AddAccountsRequest) SetCustomerType(v ExtAccountCustomerType) {
	o.CustomerType = v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *AddAccountsRequest) GetMetadata() map[string]interface{} {
	if o == nil || o.Metadata == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddAccountsRequest) GetMetadataOk() (map[string]interface{}, bool) {
	if o == nil || o.Metadata == nil {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *AddAccountsRequest) HasMetadata() bool {
	if o != nil && o.Metadata != nil {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given map[string]interface{} and assigns it to the Metadata field.
func (o *AddAccountsRequest) SetMetadata(v map[string]interface{}) {
	o.Metadata = v
}

// GetNickname returns the Nickname field value if set, zero value otherwise.
func (o *AddAccountsRequest) GetNickname() string {
	if o == nil || o.Nickname == nil {
		var ret string
		return ret
	}
	return *o.Nickname
}

// GetNicknameOk returns a tuple with the Nickname field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddAccountsRequest) GetNicknameOk() (*string, bool) {
	if o == nil || o.Nickname == nil {
		return nil, false
	}
	return o.Nickname, true
}

// HasNickname returns a boolean if a field has been set.
func (o *AddAccountsRequest) HasNickname() bool {
	if o != nil && o.Nickname != nil {
		return true
	}

	return false
}

// SetNickname gets a reference to the given string and assigns it to the Nickname field.
func (o *AddAccountsRequest) SetNickname(v string) {
	o.Nickname = &v
}

// GetRoutingIdentifiers returns the RoutingIdentifiers field value
func (o *AddAccountsRequest) GetRoutingIdentifiers() AddAccountsRequestRoutingIdentifiers {
	if o == nil {
		var ret AddAccountsRequestRoutingIdentifiers
		return ret
	}

	return o.RoutingIdentifiers
}

// GetRoutingIdentifiersOk returns a tuple with the RoutingIdentifiers field value
// and a boolean to check if the value has been set.
func (o *AddAccountsRequest) GetRoutingIdentifiersOk() (*AddAccountsRequestRoutingIdentifiers, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RoutingIdentifiers, true
}

// SetRoutingIdentifiers sets field value
func (o *AddAccountsRequest) SetRoutingIdentifiers(v AddAccountsRequestRoutingIdentifiers) {
	o.RoutingIdentifiers = v
}

// GetType returns the Type field value
func (o *AddAccountsRequest) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *AddAccountsRequest) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *AddAccountsRequest) SetType(v string) {
	o.Type = v
}

// GetVendorAccountId returns the VendorAccountId field value if set, zero value otherwise.
func (o *AddAccountsRequest) GetVendorAccountId() string {
	if o == nil || o.VendorAccountId == nil {
		var ret string
		return ret
	}
	return *o.VendorAccountId
}

// GetVendorAccountIdOk returns a tuple with the VendorAccountId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddAccountsRequest) GetVendorAccountIdOk() (*string, bool) {
	if o == nil || o.VendorAccountId == nil {
		return nil, false
	}
	return o.VendorAccountId, true
}

// HasVendorAccountId returns a boolean if a field has been set.
func (o *AddAccountsRequest) HasVendorAccountId() bool {
	if o != nil && o.VendorAccountId != nil {
		return true
	}

	return false
}

// SetVendorAccountId gets a reference to the given string and assigns it to the VendorAccountId field.
func (o *AddAccountsRequest) SetVendorAccountId(v string) {
	o.VendorAccountId = &v
}

// GetVerification returns the Verification field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AddAccountsRequest) GetVerification() AccountVerification {
	if o == nil || o.Verification.Get() == nil {
		var ret AccountVerification
		return ret
	}
	return *o.Verification.Get()
}

// GetVerificationOk returns a tuple with the Verification field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AddAccountsRequest) GetVerificationOk() (*AccountVerification, bool) {
	if o == nil {
		return nil, false
	}
	return o.Verification.Get(), o.Verification.IsSet()
}

// HasVerification returns a boolean if a field has been set.
func (o *AddAccountsRequest) HasVerification() bool {
	if o != nil && o.Verification.IsSet() {
		return true
	}

	return false
}

// SetVerification gets a reference to the given NullableAccountVerification and assigns it to the Verification field.
func (o *AddAccountsRequest) SetVerification(v AccountVerification) {
	o.Verification.Set(&v)
}
// SetVerificationNil sets the value for Verification to be an explicit nil
func (o *AddAccountsRequest) SetVerificationNil() {
	o.Verification.Set(nil)
}

// UnsetVerification ensures that no value is present for Verification, not even an explicit nil
func (o *AddAccountsRequest) UnsetVerification() {
	o.Verification.Unset()
}

func (o AddAccountsRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["account_identifiers"] = o.AccountIdentifiers
	}
	if true {
		toSerialize["account_owner_names"] = o.AccountOwnerNames
	}
	if o.BusinessId != nil {
		toSerialize["business_id"] = o.BusinessId
	}
	if o.CustomerId != nil {
		toSerialize["customer_id"] = o.CustomerId
	}
	if true {
		toSerialize["customer_type"] = o.CustomerType
	}
	if o.Metadata != nil {
		toSerialize["metadata"] = o.Metadata
	}
	if o.Nickname != nil {
		toSerialize["nickname"] = o.Nickname
	}
	if true {
		toSerialize["routing_identifiers"] = o.RoutingIdentifiers
	}
	if true {
		toSerialize["type"] = o.Type
	}
	if o.VendorAccountId != nil {
		toSerialize["vendor_account_id"] = o.VendorAccountId
	}
	if o.Verification.IsSet() {
		toSerialize["verification"] = o.Verification.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableAddAccountsRequest struct {
	value *AddAccountsRequest
	isSet bool
}

func (v NullableAddAccountsRequest) Get() *AddAccountsRequest {
	return v.value
}

func (v *NullableAddAccountsRequest) Set(val *AddAccountsRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAddAccountsRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAddAccountsRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddAccountsRequest(val *AddAccountsRequest) *NullableAddAccountsRequest {
	return &NullableAddAccountsRequest{value: val, isSet: true}
}

func (v NullableAddAccountsRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddAccountsRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

