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

// AppleDigitalWalletProvisionResponse struct for AppleDigitalWalletProvisionResponse
type AppleDigitalWalletProvisionResponse struct {
	ActivationData *string `json:"activation_data,omitempty"`
	// The unique identifier of a card
	CardId             *string    `json:"card_id,omitempty"`
	CreatedTime        *time.Time `json:"created_time,omitempty"`
	EncryptedPassData  *string    `json:"encrypted_pass_data,omitempty"`
	EphemeralPublicKey *string    `json:"ephemeral_public_key,omitempty"`
	LastModifiedTime   *time.Time `json:"last_modified_time,omitempty"`
}

// NewAppleDigitalWalletProvisionResponse instantiates a new AppleDigitalWalletProvisionResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppleDigitalWalletProvisionResponse() *AppleDigitalWalletProvisionResponse {
	this := AppleDigitalWalletProvisionResponse{}
	return &this
}

// NewAppleDigitalWalletProvisionResponseWithDefaults instantiates a new AppleDigitalWalletProvisionResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppleDigitalWalletProvisionResponseWithDefaults() *AppleDigitalWalletProvisionResponse {
	this := AppleDigitalWalletProvisionResponse{}
	return &this
}

// GetActivationData returns the ActivationData field value if set, zero value otherwise.
func (o *AppleDigitalWalletProvisionResponse) GetActivationData() string {
	if o == nil || o.ActivationData == nil {
		var ret string
		return ret
	}
	return *o.ActivationData
}

// GetActivationDataOk returns a tuple with the ActivationData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppleDigitalWalletProvisionResponse) GetActivationDataOk() (*string, bool) {
	if o == nil || o.ActivationData == nil {
		return nil, false
	}
	return o.ActivationData, true
}

// HasActivationData returns a boolean if a field has been set.
func (o *AppleDigitalWalletProvisionResponse) HasActivationData() bool {
	if o != nil && o.ActivationData != nil {
		return true
	}

	return false
}

// SetActivationData gets a reference to the given string and assigns it to the ActivationData field.
func (o *AppleDigitalWalletProvisionResponse) SetActivationData(v string) {
	o.ActivationData = &v
}

// GetCardId returns the CardId field value if set, zero value otherwise.
func (o *AppleDigitalWalletProvisionResponse) GetCardId() string {
	if o == nil || o.CardId == nil {
		var ret string
		return ret
	}
	return *o.CardId
}

// GetCardIdOk returns a tuple with the CardId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppleDigitalWalletProvisionResponse) GetCardIdOk() (*string, bool) {
	if o == nil || o.CardId == nil {
		return nil, false
	}
	return o.CardId, true
}

// HasCardId returns a boolean if a field has been set.
func (o *AppleDigitalWalletProvisionResponse) HasCardId() bool {
	if o != nil && o.CardId != nil {
		return true
	}

	return false
}

// SetCardId gets a reference to the given string and assigns it to the CardId field.
func (o *AppleDigitalWalletProvisionResponse) SetCardId(v string) {
	o.CardId = &v
}

// GetCreatedTime returns the CreatedTime field value if set, zero value otherwise.
func (o *AppleDigitalWalletProvisionResponse) GetCreatedTime() time.Time {
	if o == nil || o.CreatedTime == nil {
		var ret time.Time
		return ret
	}
	return *o.CreatedTime
}

// GetCreatedTimeOk returns a tuple with the CreatedTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppleDigitalWalletProvisionResponse) GetCreatedTimeOk() (*time.Time, bool) {
	if o == nil || o.CreatedTime == nil {
		return nil, false
	}
	return o.CreatedTime, true
}

// HasCreatedTime returns a boolean if a field has been set.
func (o *AppleDigitalWalletProvisionResponse) HasCreatedTime() bool {
	if o != nil && o.CreatedTime != nil {
		return true
	}

	return false
}

// SetCreatedTime gets a reference to the given time.Time and assigns it to the CreatedTime field.
func (o *AppleDigitalWalletProvisionResponse) SetCreatedTime(v time.Time) {
	o.CreatedTime = &v
}

// GetEncryptedPassData returns the EncryptedPassData field value if set, zero value otherwise.
func (o *AppleDigitalWalletProvisionResponse) GetEncryptedPassData() string {
	if o == nil || o.EncryptedPassData == nil {
		var ret string
		return ret
	}
	return *o.EncryptedPassData
}

// GetEncryptedPassDataOk returns a tuple with the EncryptedPassData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppleDigitalWalletProvisionResponse) GetEncryptedPassDataOk() (*string, bool) {
	if o == nil || o.EncryptedPassData == nil {
		return nil, false
	}
	return o.EncryptedPassData, true
}

// HasEncryptedPassData returns a boolean if a field has been set.
func (o *AppleDigitalWalletProvisionResponse) HasEncryptedPassData() bool {
	if o != nil && o.EncryptedPassData != nil {
		return true
	}

	return false
}

// SetEncryptedPassData gets a reference to the given string and assigns it to the EncryptedPassData field.
func (o *AppleDigitalWalletProvisionResponse) SetEncryptedPassData(v string) {
	o.EncryptedPassData = &v
}

// GetEphemeralPublicKey returns the EphemeralPublicKey field value if set, zero value otherwise.
func (o *AppleDigitalWalletProvisionResponse) GetEphemeralPublicKey() string {
	if o == nil || o.EphemeralPublicKey == nil {
		var ret string
		return ret
	}
	return *o.EphemeralPublicKey
}

// GetEphemeralPublicKeyOk returns a tuple with the EphemeralPublicKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppleDigitalWalletProvisionResponse) GetEphemeralPublicKeyOk() (*string, bool) {
	if o == nil || o.EphemeralPublicKey == nil {
		return nil, false
	}
	return o.EphemeralPublicKey, true
}

// HasEphemeralPublicKey returns a boolean if a field has been set.
func (o *AppleDigitalWalletProvisionResponse) HasEphemeralPublicKey() bool {
	if o != nil && o.EphemeralPublicKey != nil {
		return true
	}

	return false
}

// SetEphemeralPublicKey gets a reference to the given string and assigns it to the EphemeralPublicKey field.
func (o *AppleDigitalWalletProvisionResponse) SetEphemeralPublicKey(v string) {
	o.EphemeralPublicKey = &v
}

// GetLastModifiedTime returns the LastModifiedTime field value if set, zero value otherwise.
func (o *AppleDigitalWalletProvisionResponse) GetLastModifiedTime() time.Time {
	if o == nil || o.LastModifiedTime == nil {
		var ret time.Time
		return ret
	}
	return *o.LastModifiedTime
}

// GetLastModifiedTimeOk returns a tuple with the LastModifiedTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppleDigitalWalletProvisionResponse) GetLastModifiedTimeOk() (*time.Time, bool) {
	if o == nil || o.LastModifiedTime == nil {
		return nil, false
	}
	return o.LastModifiedTime, true
}

// HasLastModifiedTime returns a boolean if a field has been set.
func (o *AppleDigitalWalletProvisionResponse) HasLastModifiedTime() bool {
	if o != nil && o.LastModifiedTime != nil {
		return true
	}

	return false
}

// SetLastModifiedTime gets a reference to the given time.Time and assigns it to the LastModifiedTime field.
func (o *AppleDigitalWalletProvisionResponse) SetLastModifiedTime(v time.Time) {
	o.LastModifiedTime = &v
}

func (o AppleDigitalWalletProvisionResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ActivationData != nil {
		toSerialize["activation_data"] = o.ActivationData
	}
	if o.CardId != nil {
		toSerialize["card_id"] = o.CardId
	}
	if o.CreatedTime != nil {
		toSerialize["created_time"] = o.CreatedTime
	}
	if o.EncryptedPassData != nil {
		toSerialize["encrypted_pass_data"] = o.EncryptedPassData
	}
	if o.EphemeralPublicKey != nil {
		toSerialize["ephemeral_public_key"] = o.EphemeralPublicKey
	}
	if o.LastModifiedTime != nil {
		toSerialize["last_modified_time"] = o.LastModifiedTime
	}
	return json.Marshal(toSerialize)
}

type NullableAppleDigitalWalletProvisionResponse struct {
	value *AppleDigitalWalletProvisionResponse
	isSet bool
}

func (v NullableAppleDigitalWalletProvisionResponse) Get() *AppleDigitalWalletProvisionResponse {
	return v.value
}

func (v *NullableAppleDigitalWalletProvisionResponse) Set(val *AppleDigitalWalletProvisionResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAppleDigitalWalletProvisionResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAppleDigitalWalletProvisionResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppleDigitalWalletProvisionResponse(val *AppleDigitalWalletProvisionResponse) *NullableAppleDigitalWalletProvisionResponse {
	return &NullableAppleDigitalWalletProvisionResponse{value: val, isSet: true}
}

func (v NullableAppleDigitalWalletProvisionResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppleDigitalWalletProvisionResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
