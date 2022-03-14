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

// BinNetworkMapping struct for BinNetworkMapping
type BinNetworkMapping struct {
	// indicates whether mapping is active
	Active bool `json:"active"`
	// ID debit network uses to identify a bank
	BankNetworkId string `json:"bank_network_id"`
	// The ID of the bank's BIN that uses this debit network
	BinId string `json:"bin_id"`
	// The timestamp representing when BIN network mapping was created
	CreationTime *time.Time `json:"creation_time,omitempty"`
	// The time when mapping becomes inactive
	EndDate *time.Time `json:"end_date,omitempty"`
	// The timestamp representing when the BIN network mapping was last modified
	LastModifiedTime *time.Time `json:"last_modified_time,omitempty"`
	// The ID of the debit_network associated with the BIN of the bank
	NetworkId string `json:"network_id"`
	// The time when mapping becomes active
	StartDate *time.Time `json:"start_date,omitempty"`
}

// NewBinNetworkMapping instantiates a new BinNetworkMapping object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBinNetworkMapping(active bool, bankNetworkId string, binId string, networkId string) *BinNetworkMapping {
	this := BinNetworkMapping{}
	this.Active = active
	this.BankNetworkId = bankNetworkId
	this.BinId = binId
	this.NetworkId = networkId
	return &this
}

// NewBinNetworkMappingWithDefaults instantiates a new BinNetworkMapping object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBinNetworkMappingWithDefaults() *BinNetworkMapping {
	this := BinNetworkMapping{}
	return &this
}

// GetActive returns the Active field value
func (o *BinNetworkMapping) GetActive() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Active
}

// GetActiveOk returns a tuple with the Active field value
// and a boolean to check if the value has been set.
func (o *BinNetworkMapping) GetActiveOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Active, true
}

// SetActive sets field value
func (o *BinNetworkMapping) SetActive(v bool) {
	o.Active = v
}

// GetBankNetworkId returns the BankNetworkId field value
func (o *BinNetworkMapping) GetBankNetworkId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.BankNetworkId
}

// GetBankNetworkIdOk returns a tuple with the BankNetworkId field value
// and a boolean to check if the value has been set.
func (o *BinNetworkMapping) GetBankNetworkIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BankNetworkId, true
}

// SetBankNetworkId sets field value
func (o *BinNetworkMapping) SetBankNetworkId(v string) {
	o.BankNetworkId = v
}

// GetBinId returns the BinId field value
func (o *BinNetworkMapping) GetBinId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.BinId
}

// GetBinIdOk returns a tuple with the BinId field value
// and a boolean to check if the value has been set.
func (o *BinNetworkMapping) GetBinIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BinId, true
}

// SetBinId sets field value
func (o *BinNetworkMapping) SetBinId(v string) {
	o.BinId = v
}

// GetCreationTime returns the CreationTime field value if set, zero value otherwise.
func (o *BinNetworkMapping) GetCreationTime() time.Time {
	if o == nil || o.CreationTime == nil {
		var ret time.Time
		return ret
	}
	return *o.CreationTime
}

// GetCreationTimeOk returns a tuple with the CreationTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BinNetworkMapping) GetCreationTimeOk() (*time.Time, bool) {
	if o == nil || o.CreationTime == nil {
		return nil, false
	}
	return o.CreationTime, true
}

// HasCreationTime returns a boolean if a field has been set.
func (o *BinNetworkMapping) HasCreationTime() bool {
	if o != nil && o.CreationTime != nil {
		return true
	}

	return false
}

// SetCreationTime gets a reference to the given time.Time and assigns it to the CreationTime field.
func (o *BinNetworkMapping) SetCreationTime(v time.Time) {
	o.CreationTime = &v
}

// GetEndDate returns the EndDate field value if set, zero value otherwise.
func (o *BinNetworkMapping) GetEndDate() time.Time {
	if o == nil || o.EndDate == nil {
		var ret time.Time
		return ret
	}
	return *o.EndDate
}

// GetEndDateOk returns a tuple with the EndDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BinNetworkMapping) GetEndDateOk() (*time.Time, bool) {
	if o == nil || o.EndDate == nil {
		return nil, false
	}
	return o.EndDate, true
}

// HasEndDate returns a boolean if a field has been set.
func (o *BinNetworkMapping) HasEndDate() bool {
	if o != nil && o.EndDate != nil {
		return true
	}

	return false
}

// SetEndDate gets a reference to the given time.Time and assigns it to the EndDate field.
func (o *BinNetworkMapping) SetEndDate(v time.Time) {
	o.EndDate = &v
}

// GetLastModifiedTime returns the LastModifiedTime field value if set, zero value otherwise.
func (o *BinNetworkMapping) GetLastModifiedTime() time.Time {
	if o == nil || o.LastModifiedTime == nil {
		var ret time.Time
		return ret
	}
	return *o.LastModifiedTime
}

// GetLastModifiedTimeOk returns a tuple with the LastModifiedTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BinNetworkMapping) GetLastModifiedTimeOk() (*time.Time, bool) {
	if o == nil || o.LastModifiedTime == nil {
		return nil, false
	}
	return o.LastModifiedTime, true
}

// HasLastModifiedTime returns a boolean if a field has been set.
func (o *BinNetworkMapping) HasLastModifiedTime() bool {
	if o != nil && o.LastModifiedTime != nil {
		return true
	}

	return false
}

// SetLastModifiedTime gets a reference to the given time.Time and assigns it to the LastModifiedTime field.
func (o *BinNetworkMapping) SetLastModifiedTime(v time.Time) {
	o.LastModifiedTime = &v
}

// GetNetworkId returns the NetworkId field value
func (o *BinNetworkMapping) GetNetworkId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.NetworkId
}

// GetNetworkIdOk returns a tuple with the NetworkId field value
// and a boolean to check if the value has been set.
func (o *BinNetworkMapping) GetNetworkIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NetworkId, true
}

// SetNetworkId sets field value
func (o *BinNetworkMapping) SetNetworkId(v string) {
	o.NetworkId = v
}

// GetStartDate returns the StartDate field value if set, zero value otherwise.
func (o *BinNetworkMapping) GetStartDate() time.Time {
	if o == nil || o.StartDate == nil {
		var ret time.Time
		return ret
	}
	return *o.StartDate
}

// GetStartDateOk returns a tuple with the StartDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BinNetworkMapping) GetStartDateOk() (*time.Time, bool) {
	if o == nil || o.StartDate == nil {
		return nil, false
	}
	return o.StartDate, true
}

// HasStartDate returns a boolean if a field has been set.
func (o *BinNetworkMapping) HasStartDate() bool {
	if o != nil && o.StartDate != nil {
		return true
	}

	return false
}

// SetStartDate gets a reference to the given time.Time and assigns it to the StartDate field.
func (o *BinNetworkMapping) SetStartDate(v time.Time) {
	o.StartDate = &v
}

func (o BinNetworkMapping) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["active"] = o.Active
	}
	if true {
		toSerialize["bank_network_id"] = o.BankNetworkId
	}
	if true {
		toSerialize["bin_id"] = o.BinId
	}
	if o.CreationTime != nil {
		toSerialize["creation_time"] = o.CreationTime
	}
	if o.EndDate != nil {
		toSerialize["end_date"] = o.EndDate
	}
	if o.LastModifiedTime != nil {
		toSerialize["last_modified_time"] = o.LastModifiedTime
	}
	if true {
		toSerialize["network_id"] = o.NetworkId
	}
	if o.StartDate != nil {
		toSerialize["start_date"] = o.StartDate
	}
	return json.Marshal(toSerialize)
}

type NullableBinNetworkMapping struct {
	value *BinNetworkMapping
	isSet bool
}

func (v NullableBinNetworkMapping) Get() *BinNetworkMapping {
	return v.value
}

func (v *NullableBinNetworkMapping) Set(val *BinNetworkMapping) {
	v.value = val
	v.isSet = true
}

func (v NullableBinNetworkMapping) IsSet() bool {
	return v.isSet
}

func (v *NullableBinNetworkMapping) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBinNetworkMapping(val *BinNetworkMapping) *NullableBinNetworkMapping {
	return &NullableBinNetworkMapping{value: val, isSet: true}
}

func (v NullableBinNetworkMapping) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBinNetworkMapping) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
