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

// ProvisioningControls struct for ProvisioningControls
type ProvisioningControls struct {
	InAppProvisioning        *InAppProvisioning        `json:"in_app_provisioning,omitempty"`
	ManualEntry              *ManualEntry              `json:"manual_entry,omitempty"`
	WalletProviderCardOnFile *WalletProviderCardOnFile `json:"wallet_provider_card_on_file,omitempty"`
}

// NewProvisioningControls instantiates a new ProvisioningControls object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProvisioningControls() *ProvisioningControls {
	this := ProvisioningControls{}
	return &this
}

// NewProvisioningControlsWithDefaults instantiates a new ProvisioningControls object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProvisioningControlsWithDefaults() *ProvisioningControls {
	this := ProvisioningControls{}
	return &this
}

// GetInAppProvisioning returns the InAppProvisioning field value if set, zero value otherwise.
func (o *ProvisioningControls) GetInAppProvisioning() InAppProvisioning {
	if o == nil || o.InAppProvisioning == nil {
		var ret InAppProvisioning
		return ret
	}
	return *o.InAppProvisioning
}

// GetInAppProvisioningOk returns a tuple with the InAppProvisioning field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProvisioningControls) GetInAppProvisioningOk() (*InAppProvisioning, bool) {
	if o == nil || o.InAppProvisioning == nil {
		return nil, false
	}
	return o.InAppProvisioning, true
}

// HasInAppProvisioning returns a boolean if a field has been set.
func (o *ProvisioningControls) HasInAppProvisioning() bool {
	if o != nil && o.InAppProvisioning != nil {
		return true
	}

	return false
}

// SetInAppProvisioning gets a reference to the given InAppProvisioning and assigns it to the InAppProvisioning field.
func (o *ProvisioningControls) SetInAppProvisioning(v InAppProvisioning) {
	o.InAppProvisioning = &v
}

// GetManualEntry returns the ManualEntry field value if set, zero value otherwise.
func (o *ProvisioningControls) GetManualEntry() ManualEntry {
	if o == nil || o.ManualEntry == nil {
		var ret ManualEntry
		return ret
	}
	return *o.ManualEntry
}

// GetManualEntryOk returns a tuple with the ManualEntry field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProvisioningControls) GetManualEntryOk() (*ManualEntry, bool) {
	if o == nil || o.ManualEntry == nil {
		return nil, false
	}
	return o.ManualEntry, true
}

// HasManualEntry returns a boolean if a field has been set.
func (o *ProvisioningControls) HasManualEntry() bool {
	if o != nil && o.ManualEntry != nil {
		return true
	}

	return false
}

// SetManualEntry gets a reference to the given ManualEntry and assigns it to the ManualEntry field.
func (o *ProvisioningControls) SetManualEntry(v ManualEntry) {
	o.ManualEntry = &v
}

// GetWalletProviderCardOnFile returns the WalletProviderCardOnFile field value if set, zero value otherwise.
func (o *ProvisioningControls) GetWalletProviderCardOnFile() WalletProviderCardOnFile {
	if o == nil || o.WalletProviderCardOnFile == nil {
		var ret WalletProviderCardOnFile
		return ret
	}
	return *o.WalletProviderCardOnFile
}

// GetWalletProviderCardOnFileOk returns a tuple with the WalletProviderCardOnFile field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProvisioningControls) GetWalletProviderCardOnFileOk() (*WalletProviderCardOnFile, bool) {
	if o == nil || o.WalletProviderCardOnFile == nil {
		return nil, false
	}
	return o.WalletProviderCardOnFile, true
}

// HasWalletProviderCardOnFile returns a boolean if a field has been set.
func (o *ProvisioningControls) HasWalletProviderCardOnFile() bool {
	if o != nil && o.WalletProviderCardOnFile != nil {
		return true
	}

	return false
}

// SetWalletProviderCardOnFile gets a reference to the given WalletProviderCardOnFile and assigns it to the WalletProviderCardOnFile field.
func (o *ProvisioningControls) SetWalletProviderCardOnFile(v WalletProviderCardOnFile) {
	o.WalletProviderCardOnFile = &v
}

func (o ProvisioningControls) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.InAppProvisioning != nil {
		toSerialize["in_app_provisioning"] = o.InAppProvisioning
	}
	if o.ManualEntry != nil {
		toSerialize["manual_entry"] = o.ManualEntry
	}
	if o.WalletProviderCardOnFile != nil {
		toSerialize["wallet_provider_card_on_file"] = o.WalletProviderCardOnFile
	}
	return json.Marshal(toSerialize)
}

type NullableProvisioningControls struct {
	value *ProvisioningControls
	isSet bool
}

func (v NullableProvisioningControls) Get() *ProvisioningControls {
	return v.value
}

func (v *NullableProvisioningControls) Set(val *ProvisioningControls) {
	v.value = val
	v.isSet = true
}

func (v NullableProvisioningControls) IsSet() bool {
	return v.isSet
}

func (v *NullableProvisioningControls) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProvisioningControls(val *ProvisioningControls) *NullableProvisioningControls {
	return &NullableProvisioningControls{value: val, isSet: true}
}

func (v NullableProvisioningControls) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProvisioningControls) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
