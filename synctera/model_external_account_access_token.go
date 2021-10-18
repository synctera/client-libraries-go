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

// ExternalAccountAccessToken struct for ExternalAccountAccessToken
type ExternalAccountAccessToken struct {
	// A unique identifier for the request, which can be used for troubleshooting
	RequestId *string `json:"request_id,omitempty"`
	// The access token associated with the Item data is being requested for.
	VendorAccessToken *string `json:"vendor_access_token,omitempty"`
	// The account_id value of the account associated with the returned vendor_access_token
	VendorCustomerId string `json:"vendor_customer_id"`
	// The user's public token obtained from successful link login. 
	VendorPublicToken string `json:"vendor_public_token"`
}

// NewExternalAccountAccessToken instantiates a new ExternalAccountAccessToken object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewExternalAccountAccessToken(vendorCustomerId string, vendorPublicToken string) *ExternalAccountAccessToken {
	this := ExternalAccountAccessToken{}
	this.VendorCustomerId = vendorCustomerId
	this.VendorPublicToken = vendorPublicToken
	return &this
}

// NewExternalAccountAccessTokenWithDefaults instantiates a new ExternalAccountAccessToken object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewExternalAccountAccessTokenWithDefaults() *ExternalAccountAccessToken {
	this := ExternalAccountAccessToken{}
	return &this
}

// GetRequestId returns the RequestId field value if set, zero value otherwise.
func (o *ExternalAccountAccessToken) GetRequestId() string {
	if o == nil || o.RequestId == nil {
		var ret string
		return ret
	}
	return *o.RequestId
}

// GetRequestIdOk returns a tuple with the RequestId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExternalAccountAccessToken) GetRequestIdOk() (*string, bool) {
	if o == nil || o.RequestId == nil {
		return nil, false
	}
	return o.RequestId, true
}

// HasRequestId returns a boolean if a field has been set.
func (o *ExternalAccountAccessToken) HasRequestId() bool {
	if o != nil && o.RequestId != nil {
		return true
	}

	return false
}

// SetRequestId gets a reference to the given string and assigns it to the RequestId field.
func (o *ExternalAccountAccessToken) SetRequestId(v string) {
	o.RequestId = &v
}

// GetVendorAccessToken returns the VendorAccessToken field value if set, zero value otherwise.
func (o *ExternalAccountAccessToken) GetVendorAccessToken() string {
	if o == nil || o.VendorAccessToken == nil {
		var ret string
		return ret
	}
	return *o.VendorAccessToken
}

// GetVendorAccessTokenOk returns a tuple with the VendorAccessToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExternalAccountAccessToken) GetVendorAccessTokenOk() (*string, bool) {
	if o == nil || o.VendorAccessToken == nil {
		return nil, false
	}
	return o.VendorAccessToken, true
}

// HasVendorAccessToken returns a boolean if a field has been set.
func (o *ExternalAccountAccessToken) HasVendorAccessToken() bool {
	if o != nil && o.VendorAccessToken != nil {
		return true
	}

	return false
}

// SetVendorAccessToken gets a reference to the given string and assigns it to the VendorAccessToken field.
func (o *ExternalAccountAccessToken) SetVendorAccessToken(v string) {
	o.VendorAccessToken = &v
}

// GetVendorCustomerId returns the VendorCustomerId field value
func (o *ExternalAccountAccessToken) GetVendorCustomerId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.VendorCustomerId
}

// GetVendorCustomerIdOk returns a tuple with the VendorCustomerId field value
// and a boolean to check if the value has been set.
func (o *ExternalAccountAccessToken) GetVendorCustomerIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.VendorCustomerId, true
}

// SetVendorCustomerId sets field value
func (o *ExternalAccountAccessToken) SetVendorCustomerId(v string) {
	o.VendorCustomerId = v
}

// GetVendorPublicToken returns the VendorPublicToken field value
func (o *ExternalAccountAccessToken) GetVendorPublicToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.VendorPublicToken
}

// GetVendorPublicTokenOk returns a tuple with the VendorPublicToken field value
// and a boolean to check if the value has been set.
func (o *ExternalAccountAccessToken) GetVendorPublicTokenOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.VendorPublicToken, true
}

// SetVendorPublicToken sets field value
func (o *ExternalAccountAccessToken) SetVendorPublicToken(v string) {
	o.VendorPublicToken = v
}

func (o ExternalAccountAccessToken) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.RequestId != nil {
		toSerialize["request_id"] = o.RequestId
	}
	if o.VendorAccessToken != nil {
		toSerialize["vendor_access_token"] = o.VendorAccessToken
	}
	if true {
		toSerialize["vendor_customer_id"] = o.VendorCustomerId
	}
	if true {
		toSerialize["vendor_public_token"] = o.VendorPublicToken
	}
	return json.Marshal(toSerialize)
}

type NullableExternalAccountAccessToken struct {
	value *ExternalAccountAccessToken
	isSet bool
}

func (v NullableExternalAccountAccessToken) Get() *ExternalAccountAccessToken {
	return v.value
}

func (v *NullableExternalAccountAccessToken) Set(val *ExternalAccountAccessToken) {
	v.value = val
	v.isSet = true
}

func (v NullableExternalAccountAccessToken) IsSet() bool {
	return v.isSet
}

func (v *NullableExternalAccountAccessToken) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExternalAccountAccessToken(val *ExternalAccountAccessToken) *NullableExternalAccountAccessToken {
	return &NullableExternalAccountAccessToken{value: val, isSet: true}
}

func (v NullableExternalAccountAccessToken) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExternalAccountAccessToken) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

