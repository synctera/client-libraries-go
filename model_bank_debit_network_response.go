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

// BankDebitNetworkResponse struct for BankDebitNetworkResponse
type BankDebitNetworkResponse struct {
	// indicates whether debit network is active
	Active *bool `json:"active,omitempty"`
	// The timestamp representing when the debit network was created
	CreationTime *time.Time `json:"creation_time,omitempty"`
	// The time when debit network became inactive
	EndDate *time.Time `json:"end_date,omitempty"`
	// Debit Network ID
	Id *string `json:"id,omitempty"`
	// The timestamp representing when the debit network was last modified
	LastModifiedTime *time.Time `json:"last_modified_time,omitempty"`
	// The name describing the debit network
	Name *string `json:"name,omitempty"`
	// The time when debit network goes live
	StartDate *time.Time `json:"start_date,omitempty"`
	// The ID of the bank network
	BankNetworkId *string `json:"bank_network_id,omitempty"`
	// The ID of the bank's bin that uses this debit network
	BinId *string `json:"bin_id,omitempty"`
}

// NewBankDebitNetworkResponse instantiates a new BankDebitNetworkResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBankDebitNetworkResponse() *BankDebitNetworkResponse {
	this := BankDebitNetworkResponse{}
	return &this
}

// NewBankDebitNetworkResponseWithDefaults instantiates a new BankDebitNetworkResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBankDebitNetworkResponseWithDefaults() *BankDebitNetworkResponse {
	this := BankDebitNetworkResponse{}
	return &this
}

// GetActive returns the Active field value if set, zero value otherwise.
func (o *BankDebitNetworkResponse) GetActive() bool {
	if o == nil || o.Active == nil {
		var ret bool
		return ret
	}
	return *o.Active
}

// GetActiveOk returns a tuple with the Active field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BankDebitNetworkResponse) GetActiveOk() (*bool, bool) {
	if o == nil || o.Active == nil {
		return nil, false
	}
	return o.Active, true
}

// HasActive returns a boolean if a field has been set.
func (o *BankDebitNetworkResponse) HasActive() bool {
	if o != nil && o.Active != nil {
		return true
	}

	return false
}

// SetActive gets a reference to the given bool and assigns it to the Active field.
func (o *BankDebitNetworkResponse) SetActive(v bool) {
	o.Active = &v
}

// GetCreationTime returns the CreationTime field value if set, zero value otherwise.
func (o *BankDebitNetworkResponse) GetCreationTime() time.Time {
	if o == nil || o.CreationTime == nil {
		var ret time.Time
		return ret
	}
	return *o.CreationTime
}

// GetCreationTimeOk returns a tuple with the CreationTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BankDebitNetworkResponse) GetCreationTimeOk() (*time.Time, bool) {
	if o == nil || o.CreationTime == nil {
		return nil, false
	}
	return o.CreationTime, true
}

// HasCreationTime returns a boolean if a field has been set.
func (o *BankDebitNetworkResponse) HasCreationTime() bool {
	if o != nil && o.CreationTime != nil {
		return true
	}

	return false
}

// SetCreationTime gets a reference to the given time.Time and assigns it to the CreationTime field.
func (o *BankDebitNetworkResponse) SetCreationTime(v time.Time) {
	o.CreationTime = &v
}

// GetEndDate returns the EndDate field value if set, zero value otherwise.
func (o *BankDebitNetworkResponse) GetEndDate() time.Time {
	if o == nil || o.EndDate == nil {
		var ret time.Time
		return ret
	}
	return *o.EndDate
}

// GetEndDateOk returns a tuple with the EndDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BankDebitNetworkResponse) GetEndDateOk() (*time.Time, bool) {
	if o == nil || o.EndDate == nil {
		return nil, false
	}
	return o.EndDate, true
}

// HasEndDate returns a boolean if a field has been set.
func (o *BankDebitNetworkResponse) HasEndDate() bool {
	if o != nil && o.EndDate != nil {
		return true
	}

	return false
}

// SetEndDate gets a reference to the given time.Time and assigns it to the EndDate field.
func (o *BankDebitNetworkResponse) SetEndDate(v time.Time) {
	o.EndDate = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *BankDebitNetworkResponse) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BankDebitNetworkResponse) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *BankDebitNetworkResponse) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *BankDebitNetworkResponse) SetId(v string) {
	o.Id = &v
}

// GetLastModifiedTime returns the LastModifiedTime field value if set, zero value otherwise.
func (o *BankDebitNetworkResponse) GetLastModifiedTime() time.Time {
	if o == nil || o.LastModifiedTime == nil {
		var ret time.Time
		return ret
	}
	return *o.LastModifiedTime
}

// GetLastModifiedTimeOk returns a tuple with the LastModifiedTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BankDebitNetworkResponse) GetLastModifiedTimeOk() (*time.Time, bool) {
	if o == nil || o.LastModifiedTime == nil {
		return nil, false
	}
	return o.LastModifiedTime, true
}

// HasLastModifiedTime returns a boolean if a field has been set.
func (o *BankDebitNetworkResponse) HasLastModifiedTime() bool {
	if o != nil && o.LastModifiedTime != nil {
		return true
	}

	return false
}

// SetLastModifiedTime gets a reference to the given time.Time and assigns it to the LastModifiedTime field.
func (o *BankDebitNetworkResponse) SetLastModifiedTime(v time.Time) {
	o.LastModifiedTime = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *BankDebitNetworkResponse) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BankDebitNetworkResponse) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *BankDebitNetworkResponse) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *BankDebitNetworkResponse) SetName(v string) {
	o.Name = &v
}

// GetStartDate returns the StartDate field value if set, zero value otherwise.
func (o *BankDebitNetworkResponse) GetStartDate() time.Time {
	if o == nil || o.StartDate == nil {
		var ret time.Time
		return ret
	}
	return *o.StartDate
}

// GetStartDateOk returns a tuple with the StartDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BankDebitNetworkResponse) GetStartDateOk() (*time.Time, bool) {
	if o == nil || o.StartDate == nil {
		return nil, false
	}
	return o.StartDate, true
}

// HasStartDate returns a boolean if a field has been set.
func (o *BankDebitNetworkResponse) HasStartDate() bool {
	if o != nil && o.StartDate != nil {
		return true
	}

	return false
}

// SetStartDate gets a reference to the given time.Time and assigns it to the StartDate field.
func (o *BankDebitNetworkResponse) SetStartDate(v time.Time) {
	o.StartDate = &v
}

// GetBankNetworkId returns the BankNetworkId field value if set, zero value otherwise.
func (o *BankDebitNetworkResponse) GetBankNetworkId() string {
	if o == nil || o.BankNetworkId == nil {
		var ret string
		return ret
	}
	return *o.BankNetworkId
}

// GetBankNetworkIdOk returns a tuple with the BankNetworkId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BankDebitNetworkResponse) GetBankNetworkIdOk() (*string, bool) {
	if o == nil || o.BankNetworkId == nil {
		return nil, false
	}
	return o.BankNetworkId, true
}

// HasBankNetworkId returns a boolean if a field has been set.
func (o *BankDebitNetworkResponse) HasBankNetworkId() bool {
	if o != nil && o.BankNetworkId != nil {
		return true
	}

	return false
}

// SetBankNetworkId gets a reference to the given string and assigns it to the BankNetworkId field.
func (o *BankDebitNetworkResponse) SetBankNetworkId(v string) {
	o.BankNetworkId = &v
}

// GetBinId returns the BinId field value if set, zero value otherwise.
func (o *BankDebitNetworkResponse) GetBinId() string {
	if o == nil || o.BinId == nil {
		var ret string
		return ret
	}
	return *o.BinId
}

// GetBinIdOk returns a tuple with the BinId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BankDebitNetworkResponse) GetBinIdOk() (*string, bool) {
	if o == nil || o.BinId == nil {
		return nil, false
	}
	return o.BinId, true
}

// HasBinId returns a boolean if a field has been set.
func (o *BankDebitNetworkResponse) HasBinId() bool {
	if o != nil && o.BinId != nil {
		return true
	}

	return false
}

// SetBinId gets a reference to the given string and assigns it to the BinId field.
func (o *BankDebitNetworkResponse) SetBinId(v string) {
	o.BinId = &v
}

func (o BankDebitNetworkResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Active != nil {
		toSerialize["active"] = o.Active
	}
	if o.CreationTime != nil {
		toSerialize["creation_time"] = o.CreationTime
	}
	if o.EndDate != nil {
		toSerialize["end_date"] = o.EndDate
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.LastModifiedTime != nil {
		toSerialize["last_modified_time"] = o.LastModifiedTime
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.StartDate != nil {
		toSerialize["start_date"] = o.StartDate
	}
	if o.BankNetworkId != nil {
		toSerialize["bank_network_id"] = o.BankNetworkId
	}
	if o.BinId != nil {
		toSerialize["bin_id"] = o.BinId
	}
	return json.Marshal(toSerialize)
}

type NullableBankDebitNetworkResponse struct {
	value *BankDebitNetworkResponse
	isSet bool
}

func (v NullableBankDebitNetworkResponse) Get() *BankDebitNetworkResponse {
	return v.value
}

func (v *NullableBankDebitNetworkResponse) Set(val *BankDebitNetworkResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableBankDebitNetworkResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableBankDebitNetworkResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBankDebitNetworkResponse(val *BankDebitNetworkResponse) *NullableBankDebitNetworkResponse {
	return &NullableBankDebitNetworkResponse{value: val, isSet: true}
}

func (v NullableBankDebitNetworkResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBankDebitNetworkResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
