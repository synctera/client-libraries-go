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

// AccountAddress struct for AccountAddress
type AccountAddress struct {
	Address *Address `json:"address,omitempty"`
	// Address unique ID
	AddressId *int32 `json:"address_id,omitempty"`
	// Connection ID of the account
	ConnectId *string `json:"connect_id,omitempty"`
	// Customer ID
	CustomerId *string `json:"customer_id,omitempty"`
	// Document
	DocumentTypeId *int32 `json:"document_type_id,omitempty"`
	// Indicator of duplicate of the address
	Duplicate *bool `json:"duplicate,omitempty"`
}

// NewAccountAddress instantiates a new AccountAddress object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccountAddress() *AccountAddress {
	this := AccountAddress{}
	return &this
}

// NewAccountAddressWithDefaults instantiates a new AccountAddress object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccountAddressWithDefaults() *AccountAddress {
	this := AccountAddress{}
	return &this
}

// GetAddress returns the Address field value if set, zero value otherwise.
func (o *AccountAddress) GetAddress() Address {
	if o == nil || o.Address == nil {
		var ret Address
		return ret
	}
	return *o.Address
}

// GetAddressOk returns a tuple with the Address field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountAddress) GetAddressOk() (*Address, bool) {
	if o == nil || o.Address == nil {
		return nil, false
	}
	return o.Address, true
}

// HasAddress returns a boolean if a field has been set.
func (o *AccountAddress) HasAddress() bool {
	if o != nil && o.Address != nil {
		return true
	}

	return false
}

// SetAddress gets a reference to the given Address and assigns it to the Address field.
func (o *AccountAddress) SetAddress(v Address) {
	o.Address = &v
}

// GetAddressId returns the AddressId field value if set, zero value otherwise.
func (o *AccountAddress) GetAddressId() int32 {
	if o == nil || o.AddressId == nil {
		var ret int32
		return ret
	}
	return *o.AddressId
}

// GetAddressIdOk returns a tuple with the AddressId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountAddress) GetAddressIdOk() (*int32, bool) {
	if o == nil || o.AddressId == nil {
		return nil, false
	}
	return o.AddressId, true
}

// HasAddressId returns a boolean if a field has been set.
func (o *AccountAddress) HasAddressId() bool {
	if o != nil && o.AddressId != nil {
		return true
	}

	return false
}

// SetAddressId gets a reference to the given int32 and assigns it to the AddressId field.
func (o *AccountAddress) SetAddressId(v int32) {
	o.AddressId = &v
}

// GetConnectId returns the ConnectId field value if set, zero value otherwise.
func (o *AccountAddress) GetConnectId() string {
	if o == nil || o.ConnectId == nil {
		var ret string
		return ret
	}
	return *o.ConnectId
}

// GetConnectIdOk returns a tuple with the ConnectId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountAddress) GetConnectIdOk() (*string, bool) {
	if o == nil || o.ConnectId == nil {
		return nil, false
	}
	return o.ConnectId, true
}

// HasConnectId returns a boolean if a field has been set.
func (o *AccountAddress) HasConnectId() bool {
	if o != nil && o.ConnectId != nil {
		return true
	}

	return false
}

// SetConnectId gets a reference to the given string and assigns it to the ConnectId field.
func (o *AccountAddress) SetConnectId(v string) {
	o.ConnectId = &v
}

// GetCustomerId returns the CustomerId field value if set, zero value otherwise.
func (o *AccountAddress) GetCustomerId() string {
	if o == nil || o.CustomerId == nil {
		var ret string
		return ret
	}
	return *o.CustomerId
}

// GetCustomerIdOk returns a tuple with the CustomerId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountAddress) GetCustomerIdOk() (*string, bool) {
	if o == nil || o.CustomerId == nil {
		return nil, false
	}
	return o.CustomerId, true
}

// HasCustomerId returns a boolean if a field has been set.
func (o *AccountAddress) HasCustomerId() bool {
	if o != nil && o.CustomerId != nil {
		return true
	}

	return false
}

// SetCustomerId gets a reference to the given string and assigns it to the CustomerId field.
func (o *AccountAddress) SetCustomerId(v string) {
	o.CustomerId = &v
}

// GetDocumentTypeId returns the DocumentTypeId field value if set, zero value otherwise.
func (o *AccountAddress) GetDocumentTypeId() int32 {
	if o == nil || o.DocumentTypeId == nil {
		var ret int32
		return ret
	}
	return *o.DocumentTypeId
}

// GetDocumentTypeIdOk returns a tuple with the DocumentTypeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountAddress) GetDocumentTypeIdOk() (*int32, bool) {
	if o == nil || o.DocumentTypeId == nil {
		return nil, false
	}
	return o.DocumentTypeId, true
}

// HasDocumentTypeId returns a boolean if a field has been set.
func (o *AccountAddress) HasDocumentTypeId() bool {
	if o != nil && o.DocumentTypeId != nil {
		return true
	}

	return false
}

// SetDocumentTypeId gets a reference to the given int32 and assigns it to the DocumentTypeId field.
func (o *AccountAddress) SetDocumentTypeId(v int32) {
	o.DocumentTypeId = &v
}

// GetDuplicate returns the Duplicate field value if set, zero value otherwise.
func (o *AccountAddress) GetDuplicate() bool {
	if o == nil || o.Duplicate == nil {
		var ret bool
		return ret
	}
	return *o.Duplicate
}

// GetDuplicateOk returns a tuple with the Duplicate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountAddress) GetDuplicateOk() (*bool, bool) {
	if o == nil || o.Duplicate == nil {
		return nil, false
	}
	return o.Duplicate, true
}

// HasDuplicate returns a boolean if a field has been set.
func (o *AccountAddress) HasDuplicate() bool {
	if o != nil && o.Duplicate != nil {
		return true
	}

	return false
}

// SetDuplicate gets a reference to the given bool and assigns it to the Duplicate field.
func (o *AccountAddress) SetDuplicate(v bool) {
	o.Duplicate = &v
}

func (o AccountAddress) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Address != nil {
		toSerialize["address"] = o.Address
	}
	if o.AddressId != nil {
		toSerialize["address_id"] = o.AddressId
	}
	if o.ConnectId != nil {
		toSerialize["connect_id"] = o.ConnectId
	}
	if o.CustomerId != nil {
		toSerialize["customer_id"] = o.CustomerId
	}
	if o.DocumentTypeId != nil {
		toSerialize["document_type_id"] = o.DocumentTypeId
	}
	if o.Duplicate != nil {
		toSerialize["duplicate"] = o.Duplicate
	}
	return json.Marshal(toSerialize)
}

type NullableAccountAddress struct {
	value *AccountAddress
	isSet bool
}

func (v NullableAccountAddress) Get() *AccountAddress {
	return v.value
}

func (v *NullableAccountAddress) Set(val *AccountAddress) {
	v.value = val
	v.isSet = true
}

func (v NullableAccountAddress) IsSet() bool {
	return v.isSet
}

func (v *NullableAccountAddress) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccountAddress(val *AccountAddress) *NullableAccountAddress {
	return &NullableAccountAddress{value: val, isSet: true}
}

func (v NullableAccountAddress) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccountAddress) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


