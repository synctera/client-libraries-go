/*
Synctera API

<h2>Let's build something great.</h2><p>Welcome to the official reference documentation for Synctera APIs. Our APIs are the best way to automate your company's banking needs and are designed to be easy to understand and implement.</p><p>We're continuously growing this library and what you see here is just the start, but if you need something specific or have a question, <a class='text-blue-600' href='https://synctera.com/contact' target='_blank' rel='noreferrer'>contact us</a>.</p> 

API version: 0.5.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"time"
)

// SingleUseTokenResponse struct for SingleUseTokenResponse
type SingleUseTokenResponse struct {
	CustomerAccountMappingId *string `json:"customer_account_mapping_id,omitempty"`
	// yyyy-MM-ddTHH:mm:ssZ
	Expires time.Time `json:"expires"`
	Token *string `json:"token,omitempty"`
}

// NewSingleUseTokenResponse instantiates a new SingleUseTokenResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSingleUseTokenResponse(expires time.Time) *SingleUseTokenResponse {
	this := SingleUseTokenResponse{}
	this.Expires = expires
	return &this
}

// NewSingleUseTokenResponseWithDefaults instantiates a new SingleUseTokenResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSingleUseTokenResponseWithDefaults() *SingleUseTokenResponse {
	this := SingleUseTokenResponse{}
	return &this
}

// GetCustomerAccountMappingId returns the CustomerAccountMappingId field value if set, zero value otherwise.
func (o *SingleUseTokenResponse) GetCustomerAccountMappingId() string {
	if o == nil || o.CustomerAccountMappingId == nil {
		var ret string
		return ret
	}
	return *o.CustomerAccountMappingId
}

// GetCustomerAccountMappingIdOk returns a tuple with the CustomerAccountMappingId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SingleUseTokenResponse) GetCustomerAccountMappingIdOk() (*string, bool) {
	if o == nil || o.CustomerAccountMappingId == nil {
		return nil, false
	}
	return o.CustomerAccountMappingId, true
}

// HasCustomerAccountMappingId returns a boolean if a field has been set.
func (o *SingleUseTokenResponse) HasCustomerAccountMappingId() bool {
	if o != nil && o.CustomerAccountMappingId != nil {
		return true
	}

	return false
}

// SetCustomerAccountMappingId gets a reference to the given string and assigns it to the CustomerAccountMappingId field.
func (o *SingleUseTokenResponse) SetCustomerAccountMappingId(v string) {
	o.CustomerAccountMappingId = &v
}

// GetExpires returns the Expires field value
func (o *SingleUseTokenResponse) GetExpires() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.Expires
}

// GetExpiresOk returns a tuple with the Expires field value
// and a boolean to check if the value has been set.
func (o *SingleUseTokenResponse) GetExpiresOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Expires, true
}

// SetExpires sets field value
func (o *SingleUseTokenResponse) SetExpires(v time.Time) {
	o.Expires = v
}

// GetToken returns the Token field value if set, zero value otherwise.
func (o *SingleUseTokenResponse) GetToken() string {
	if o == nil || o.Token == nil {
		var ret string
		return ret
	}
	return *o.Token
}

// GetTokenOk returns a tuple with the Token field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SingleUseTokenResponse) GetTokenOk() (*string, bool) {
	if o == nil || o.Token == nil {
		return nil, false
	}
	return o.Token, true
}

// HasToken returns a boolean if a field has been set.
func (o *SingleUseTokenResponse) HasToken() bool {
	if o != nil && o.Token != nil {
		return true
	}

	return false
}

// SetToken gets a reference to the given string and assigns it to the Token field.
func (o *SingleUseTokenResponse) SetToken(v string) {
	o.Token = &v
}

func (o SingleUseTokenResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.CustomerAccountMappingId != nil {
		toSerialize["customer_account_mapping_id"] = o.CustomerAccountMappingId
	}
	if true {
		toSerialize["expires"] = o.Expires
	}
	if o.Token != nil {
		toSerialize["token"] = o.Token
	}
	return json.Marshal(toSerialize)
}

type NullableSingleUseTokenResponse struct {
	value *SingleUseTokenResponse
	isSet bool
}

func (v NullableSingleUseTokenResponse) Get() *SingleUseTokenResponse {
	return v.value
}

func (v *NullableSingleUseTokenResponse) Set(val *SingleUseTokenResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableSingleUseTokenResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableSingleUseTokenResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSingleUseTokenResponse(val *SingleUseTokenResponse) *NullableSingleUseTokenResponse {
	return &NullableSingleUseTokenResponse{value: val, isSet: true}
}

func (v NullableSingleUseTokenResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSingleUseTokenResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


