/*
Synctera API

<h2>Let's build something great.</h2><p>Welcome to the official reference documentation for Synctera APIs. Our APIs are the best way to automate your company's banking needs and are designed to be easy to understand and implement.</p><p>We're continuously growing this library and what you see here is just the start, but if you need something specific or have a question, <a class='text-blue-600' href='https://synctera.com/contact' target='_blank' rel='noreferrer'>contact us</a>.</p> 

API version: 0.19.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package synctera_client

import (
	"encoding/json"
	"time"
)

// ViralLoopWaitlists Viral Loop Waitlists
type ViralLoopWaitlists struct {
	// Bank ID
	BankId *int32 `json:"bank_id,omitempty"`
	// Creation time
	CreationTime *time.Time `json:"creation_time,omitempty"`
	// Viral Loop Waitlist ID
	Id string `json:"id"`
	// Most recent updated time
	LastUpdatedTime *time.Time `json:"last_updated_time,omitempty"`
	// Total number of leads/people on the waitlist
	LeadCount *int32 `json:"lead_count,omitempty"`
	// Partner ID
	PartnerId *int32 `json:"partner_id,omitempty"`
	// viral loop api token
	ViralLoopApiToken string `json:"viral_loop_api_token"`
}

// NewViralLoopWaitlists instantiates a new ViralLoopWaitlists object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewViralLoopWaitlists(id string, viralLoopApiToken string) *ViralLoopWaitlists {
	this := ViralLoopWaitlists{}
	this.Id = id
	this.ViralLoopApiToken = viralLoopApiToken
	return &this
}

// NewViralLoopWaitlistsWithDefaults instantiates a new ViralLoopWaitlists object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewViralLoopWaitlistsWithDefaults() *ViralLoopWaitlists {
	this := ViralLoopWaitlists{}
	return &this
}

// GetBankId returns the BankId field value if set, zero value otherwise.
func (o *ViralLoopWaitlists) GetBankId() int32 {
	if o == nil || o.BankId == nil {
		var ret int32
		return ret
	}
	return *o.BankId
}

// GetBankIdOk returns a tuple with the BankId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ViralLoopWaitlists) GetBankIdOk() (*int32, bool) {
	if o == nil || o.BankId == nil {
		return nil, false
	}
	return o.BankId, true
}

// HasBankId returns a boolean if a field has been set.
func (o *ViralLoopWaitlists) HasBankId() bool {
	if o != nil && o.BankId != nil {
		return true
	}

	return false
}

// SetBankId gets a reference to the given int32 and assigns it to the BankId field.
func (o *ViralLoopWaitlists) SetBankId(v int32) {
	o.BankId = &v
}

// GetCreationTime returns the CreationTime field value if set, zero value otherwise.
func (o *ViralLoopWaitlists) GetCreationTime() time.Time {
	if o == nil || o.CreationTime == nil {
		var ret time.Time
		return ret
	}
	return *o.CreationTime
}

// GetCreationTimeOk returns a tuple with the CreationTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ViralLoopWaitlists) GetCreationTimeOk() (*time.Time, bool) {
	if o == nil || o.CreationTime == nil {
		return nil, false
	}
	return o.CreationTime, true
}

// HasCreationTime returns a boolean if a field has been set.
func (o *ViralLoopWaitlists) HasCreationTime() bool {
	if o != nil && o.CreationTime != nil {
		return true
	}

	return false
}

// SetCreationTime gets a reference to the given time.Time and assigns it to the CreationTime field.
func (o *ViralLoopWaitlists) SetCreationTime(v time.Time) {
	o.CreationTime = &v
}

// GetId returns the Id field value
func (o *ViralLoopWaitlists) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ViralLoopWaitlists) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ViralLoopWaitlists) SetId(v string) {
	o.Id = v
}

// GetLastUpdatedTime returns the LastUpdatedTime field value if set, zero value otherwise.
func (o *ViralLoopWaitlists) GetLastUpdatedTime() time.Time {
	if o == nil || o.LastUpdatedTime == nil {
		var ret time.Time
		return ret
	}
	return *o.LastUpdatedTime
}

// GetLastUpdatedTimeOk returns a tuple with the LastUpdatedTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ViralLoopWaitlists) GetLastUpdatedTimeOk() (*time.Time, bool) {
	if o == nil || o.LastUpdatedTime == nil {
		return nil, false
	}
	return o.LastUpdatedTime, true
}

// HasLastUpdatedTime returns a boolean if a field has been set.
func (o *ViralLoopWaitlists) HasLastUpdatedTime() bool {
	if o != nil && o.LastUpdatedTime != nil {
		return true
	}

	return false
}

// SetLastUpdatedTime gets a reference to the given time.Time and assigns it to the LastUpdatedTime field.
func (o *ViralLoopWaitlists) SetLastUpdatedTime(v time.Time) {
	o.LastUpdatedTime = &v
}

// GetLeadCount returns the LeadCount field value if set, zero value otherwise.
func (o *ViralLoopWaitlists) GetLeadCount() int32 {
	if o == nil || o.LeadCount == nil {
		var ret int32
		return ret
	}
	return *o.LeadCount
}

// GetLeadCountOk returns a tuple with the LeadCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ViralLoopWaitlists) GetLeadCountOk() (*int32, bool) {
	if o == nil || o.LeadCount == nil {
		return nil, false
	}
	return o.LeadCount, true
}

// HasLeadCount returns a boolean if a field has been set.
func (o *ViralLoopWaitlists) HasLeadCount() bool {
	if o != nil && o.LeadCount != nil {
		return true
	}

	return false
}

// SetLeadCount gets a reference to the given int32 and assigns it to the LeadCount field.
func (o *ViralLoopWaitlists) SetLeadCount(v int32) {
	o.LeadCount = &v
}

// GetPartnerId returns the PartnerId field value if set, zero value otherwise.
func (o *ViralLoopWaitlists) GetPartnerId() int32 {
	if o == nil || o.PartnerId == nil {
		var ret int32
		return ret
	}
	return *o.PartnerId
}

// GetPartnerIdOk returns a tuple with the PartnerId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ViralLoopWaitlists) GetPartnerIdOk() (*int32, bool) {
	if o == nil || o.PartnerId == nil {
		return nil, false
	}
	return o.PartnerId, true
}

// HasPartnerId returns a boolean if a field has been set.
func (o *ViralLoopWaitlists) HasPartnerId() bool {
	if o != nil && o.PartnerId != nil {
		return true
	}

	return false
}

// SetPartnerId gets a reference to the given int32 and assigns it to the PartnerId field.
func (o *ViralLoopWaitlists) SetPartnerId(v int32) {
	o.PartnerId = &v
}

// GetViralLoopApiToken returns the ViralLoopApiToken field value
func (o *ViralLoopWaitlists) GetViralLoopApiToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ViralLoopApiToken
}

// GetViralLoopApiTokenOk returns a tuple with the ViralLoopApiToken field value
// and a boolean to check if the value has been set.
func (o *ViralLoopWaitlists) GetViralLoopApiTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ViralLoopApiToken, true
}

// SetViralLoopApiToken sets field value
func (o *ViralLoopWaitlists) SetViralLoopApiToken(v string) {
	o.ViralLoopApiToken = v
}

func (o ViralLoopWaitlists) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.BankId != nil {
		toSerialize["bank_id"] = o.BankId
	}
	if o.CreationTime != nil {
		toSerialize["creation_time"] = o.CreationTime
	}
	if true {
		toSerialize["id"] = o.Id
	}
	if o.LastUpdatedTime != nil {
		toSerialize["last_updated_time"] = o.LastUpdatedTime
	}
	if o.LeadCount != nil {
		toSerialize["lead_count"] = o.LeadCount
	}
	if o.PartnerId != nil {
		toSerialize["partner_id"] = o.PartnerId
	}
	if true {
		toSerialize["viral_loop_api_token"] = o.ViralLoopApiToken
	}
	return json.Marshal(toSerialize)
}

type NullableViralLoopWaitlists struct {
	value *ViralLoopWaitlists
	isSet bool
}

func (v NullableViralLoopWaitlists) Get() *ViralLoopWaitlists {
	return v.value
}

func (v *NullableViralLoopWaitlists) Set(val *ViralLoopWaitlists) {
	v.value = val
	v.isSet = true
}

func (v NullableViralLoopWaitlists) IsSet() bool {
	return v.isSet
}

func (v *NullableViralLoopWaitlists) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableViralLoopWaitlists(val *ViralLoopWaitlists) *NullableViralLoopWaitlists {
	return &NullableViralLoopWaitlists{value: val, isSet: true}
}

func (v NullableViralLoopWaitlists) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableViralLoopWaitlists) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

