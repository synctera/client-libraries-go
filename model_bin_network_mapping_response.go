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

// BinNetworkMappingResponse struct for BinNetworkMappingResponse
type BinNetworkMappingResponse struct {
	// indicates whether mapping is active
	Active bool `json:"active"`
	// ID debit network uses to identify a bank
	BankNetworkId string `json:"bank_network_id"`
	// The ID of the bank's BIN that uses this debit network
	BinId string `json:"bin_id"`
	// The timestamp representing when BIN network mapping was created
	CreationTime time.Time `json:"creation_time"`
	// The time when mapping becomes inactive
	EndDate time.Time `json:"end_date"`
	// The timestamp representing when the BIN network mapping was last modified
	LastModifiedTime time.Time `json:"last_modified_time"`
	// The ID of the debit_network associated with the BIN of the bank
	NetworkId string `json:"network_id"`
	// The time when mapping becomes active
	StartDate time.Time `json:"start_date"`
}

// NewBinNetworkMappingResponse instantiates a new BinNetworkMappingResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBinNetworkMappingResponse(active bool, bankNetworkId string, binId string, creationTime time.Time, endDate time.Time, lastModifiedTime time.Time, networkId string, startDate time.Time) *BinNetworkMappingResponse {
	this := BinNetworkMappingResponse{}
	this.Active = active
	this.BankNetworkId = bankNetworkId
	this.BinId = binId
	this.CreationTime = creationTime
	this.EndDate = endDate
	this.LastModifiedTime = lastModifiedTime
	this.NetworkId = networkId
	this.StartDate = startDate
	return &this
}

// NewBinNetworkMappingResponseWithDefaults instantiates a new BinNetworkMappingResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBinNetworkMappingResponseWithDefaults() *BinNetworkMappingResponse {
	this := BinNetworkMappingResponse{}
	return &this
}

// GetActive returns the Active field value
func (o *BinNetworkMappingResponse) GetActive() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Active
}

// GetActiveOk returns a tuple with the Active field value
// and a boolean to check if the value has been set.
func (o *BinNetworkMappingResponse) GetActiveOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Active, true
}

// SetActive sets field value
func (o *BinNetworkMappingResponse) SetActive(v bool) {
	o.Active = v
}

// GetBankNetworkId returns the BankNetworkId field value
func (o *BinNetworkMappingResponse) GetBankNetworkId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.BankNetworkId
}

// GetBankNetworkIdOk returns a tuple with the BankNetworkId field value
// and a boolean to check if the value has been set.
func (o *BinNetworkMappingResponse) GetBankNetworkIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BankNetworkId, true
}

// SetBankNetworkId sets field value
func (o *BinNetworkMappingResponse) SetBankNetworkId(v string) {
	o.BankNetworkId = v
}

// GetBinId returns the BinId field value
func (o *BinNetworkMappingResponse) GetBinId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.BinId
}

// GetBinIdOk returns a tuple with the BinId field value
// and a boolean to check if the value has been set.
func (o *BinNetworkMappingResponse) GetBinIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BinId, true
}

// SetBinId sets field value
func (o *BinNetworkMappingResponse) SetBinId(v string) {
	o.BinId = v
}

// GetCreationTime returns the CreationTime field value
func (o *BinNetworkMappingResponse) GetCreationTime() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreationTime
}

// GetCreationTimeOk returns a tuple with the CreationTime field value
// and a boolean to check if the value has been set.
func (o *BinNetworkMappingResponse) GetCreationTimeOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreationTime, true
}

// SetCreationTime sets field value
func (o *BinNetworkMappingResponse) SetCreationTime(v time.Time) {
	o.CreationTime = v
}

// GetEndDate returns the EndDate field value
func (o *BinNetworkMappingResponse) GetEndDate() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.EndDate
}

// GetEndDateOk returns a tuple with the EndDate field value
// and a boolean to check if the value has been set.
func (o *BinNetworkMappingResponse) GetEndDateOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EndDate, true
}

// SetEndDate sets field value
func (o *BinNetworkMappingResponse) SetEndDate(v time.Time) {
	o.EndDate = v
}

// GetLastModifiedTime returns the LastModifiedTime field value
func (o *BinNetworkMappingResponse) GetLastModifiedTime() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.LastModifiedTime
}

// GetLastModifiedTimeOk returns a tuple with the LastModifiedTime field value
// and a boolean to check if the value has been set.
func (o *BinNetworkMappingResponse) GetLastModifiedTimeOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LastModifiedTime, true
}

// SetLastModifiedTime sets field value
func (o *BinNetworkMappingResponse) SetLastModifiedTime(v time.Time) {
	o.LastModifiedTime = v
}

// GetNetworkId returns the NetworkId field value
func (o *BinNetworkMappingResponse) GetNetworkId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.NetworkId
}

// GetNetworkIdOk returns a tuple with the NetworkId field value
// and a boolean to check if the value has been set.
func (o *BinNetworkMappingResponse) GetNetworkIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NetworkId, true
}

// SetNetworkId sets field value
func (o *BinNetworkMappingResponse) SetNetworkId(v string) {
	o.NetworkId = v
}

// GetStartDate returns the StartDate field value
func (o *BinNetworkMappingResponse) GetStartDate() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.StartDate
}

// GetStartDateOk returns a tuple with the StartDate field value
// and a boolean to check if the value has been set.
func (o *BinNetworkMappingResponse) GetStartDateOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StartDate, true
}

// SetStartDate sets field value
func (o *BinNetworkMappingResponse) SetStartDate(v time.Time) {
	o.StartDate = v
}

func (o BinNetworkMappingResponse) MarshalJSON() ([]byte, error) {
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
	if true {
		toSerialize["creation_time"] = o.CreationTime
	}
	if true {
		toSerialize["end_date"] = o.EndDate
	}
	if true {
		toSerialize["last_modified_time"] = o.LastModifiedTime
	}
	if true {
		toSerialize["network_id"] = o.NetworkId
	}
	if true {
		toSerialize["start_date"] = o.StartDate
	}
	return json.Marshal(toSerialize)
}

type NullableBinNetworkMappingResponse struct {
	value *BinNetworkMappingResponse
	isSet bool
}

func (v NullableBinNetworkMappingResponse) Get() *BinNetworkMappingResponse {
	return v.value
}

func (v *NullableBinNetworkMappingResponse) Set(val *BinNetworkMappingResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableBinNetworkMappingResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableBinNetworkMappingResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBinNetworkMappingResponse(val *BinNetworkMappingResponse) *NullableBinNetworkMappingResponse {
	return &NullableBinNetworkMappingResponse{value: val, isSet: true}
}

func (v NullableBinNetworkMappingResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBinNetworkMappingResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
