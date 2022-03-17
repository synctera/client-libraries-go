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

// VerificationAllOf struct for VerificationAllOf
type VerificationAllOf struct {
	// Unique ID for the business. Exactly one of `business_id` or `person_id` must be set.
	BusinessId *string `json:"business_id,omitempty"`
	// The date and time the resource was created.
	CreationTime *time.Time `json:"creation_time,omitempty"`
	// A list of individual checks done as part of the due diligence process  for the verification type.
	Details *[]Detail `json:"details,omitempty"`
	// Unique ID for this verification result.
	Id *string `json:"id,omitempty"`
	// The date and time the resource was last updated.
	LastUpdatedTime *time.Time `json:"last_updated_time,omitempty"`
	// Optional field to store additional information about the resource. Intended to be used by the integrator to store non-sensitive data.
	Metadata *map[string]interface{} `json:"metadata,omitempty"`
	// Unique ID for the person. Exactly one of `person_id` or `business_id` must be set.
	PersonId   *string                 `json:"person_id,omitempty"`
	Result     *VerificationResult     `json:"result,omitempty"`
	VendorInfo *VerificationVendorInfo `json:"vendor_info,omitempty"`
	// The date and time the verification was completed.
	VerificationTime *time.Time         `json:"verification_time,omitempty"`
	VerificationType *VerificationType1 `json:"verification_type,omitempty"`
}

// NewVerificationAllOf instantiates a new VerificationAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVerificationAllOf() *VerificationAllOf {
	this := VerificationAllOf{}
	return &this
}

// NewVerificationAllOfWithDefaults instantiates a new VerificationAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVerificationAllOfWithDefaults() *VerificationAllOf {
	this := VerificationAllOf{}
	return &this
}

// GetBusinessId returns the BusinessId field value if set, zero value otherwise.
func (o *VerificationAllOf) GetBusinessId() string {
	if o == nil || o.BusinessId == nil {
		var ret string
		return ret
	}
	return *o.BusinessId
}

// GetBusinessIdOk returns a tuple with the BusinessId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VerificationAllOf) GetBusinessIdOk() (*string, bool) {
	if o == nil || o.BusinessId == nil {
		return nil, false
	}
	return o.BusinessId, true
}

// HasBusinessId returns a boolean if a field has been set.
func (o *VerificationAllOf) HasBusinessId() bool {
	if o != nil && o.BusinessId != nil {
		return true
	}

	return false
}

// SetBusinessId gets a reference to the given string and assigns it to the BusinessId field.
func (o *VerificationAllOf) SetBusinessId(v string) {
	o.BusinessId = &v
}

// GetCreationTime returns the CreationTime field value if set, zero value otherwise.
func (o *VerificationAllOf) GetCreationTime() time.Time {
	if o == nil || o.CreationTime == nil {
		var ret time.Time
		return ret
	}
	return *o.CreationTime
}

// GetCreationTimeOk returns a tuple with the CreationTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VerificationAllOf) GetCreationTimeOk() (*time.Time, bool) {
	if o == nil || o.CreationTime == nil {
		return nil, false
	}
	return o.CreationTime, true
}

// HasCreationTime returns a boolean if a field has been set.
func (o *VerificationAllOf) HasCreationTime() bool {
	if o != nil && o.CreationTime != nil {
		return true
	}

	return false
}

// SetCreationTime gets a reference to the given time.Time and assigns it to the CreationTime field.
func (o *VerificationAllOf) SetCreationTime(v time.Time) {
	o.CreationTime = &v
}

// GetDetails returns the Details field value if set, zero value otherwise.
func (o *VerificationAllOf) GetDetails() []Detail {
	if o == nil || o.Details == nil {
		var ret []Detail
		return ret
	}
	return *o.Details
}

// GetDetailsOk returns a tuple with the Details field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VerificationAllOf) GetDetailsOk() (*[]Detail, bool) {
	if o == nil || o.Details == nil {
		return nil, false
	}
	return o.Details, true
}

// HasDetails returns a boolean if a field has been set.
func (o *VerificationAllOf) HasDetails() bool {
	if o != nil && o.Details != nil {
		return true
	}

	return false
}

// SetDetails gets a reference to the given []Detail and assigns it to the Details field.
func (o *VerificationAllOf) SetDetails(v []Detail) {
	o.Details = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *VerificationAllOf) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VerificationAllOf) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *VerificationAllOf) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *VerificationAllOf) SetId(v string) {
	o.Id = &v
}

// GetLastUpdatedTime returns the LastUpdatedTime field value if set, zero value otherwise.
func (o *VerificationAllOf) GetLastUpdatedTime() time.Time {
	if o == nil || o.LastUpdatedTime == nil {
		var ret time.Time
		return ret
	}
	return *o.LastUpdatedTime
}

// GetLastUpdatedTimeOk returns a tuple with the LastUpdatedTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VerificationAllOf) GetLastUpdatedTimeOk() (*time.Time, bool) {
	if o == nil || o.LastUpdatedTime == nil {
		return nil, false
	}
	return o.LastUpdatedTime, true
}

// HasLastUpdatedTime returns a boolean if a field has been set.
func (o *VerificationAllOf) HasLastUpdatedTime() bool {
	if o != nil && o.LastUpdatedTime != nil {
		return true
	}

	return false
}

// SetLastUpdatedTime gets a reference to the given time.Time and assigns it to the LastUpdatedTime field.
func (o *VerificationAllOf) SetLastUpdatedTime(v time.Time) {
	o.LastUpdatedTime = &v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *VerificationAllOf) GetMetadata() map[string]interface{} {
	if o == nil || o.Metadata == nil {
		var ret map[string]interface{}
		return ret
	}
	return *o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VerificationAllOf) GetMetadataOk() (*map[string]interface{}, bool) {
	if o == nil || o.Metadata == nil {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *VerificationAllOf) HasMetadata() bool {
	if o != nil && o.Metadata != nil {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given map[string]interface{} and assigns it to the Metadata field.
func (o *VerificationAllOf) SetMetadata(v map[string]interface{}) {
	o.Metadata = &v
}

// GetPersonId returns the PersonId field value if set, zero value otherwise.
func (o *VerificationAllOf) GetPersonId() string {
	if o == nil || o.PersonId == nil {
		var ret string
		return ret
	}
	return *o.PersonId
}

// GetPersonIdOk returns a tuple with the PersonId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VerificationAllOf) GetPersonIdOk() (*string, bool) {
	if o == nil || o.PersonId == nil {
		return nil, false
	}
	return o.PersonId, true
}

// HasPersonId returns a boolean if a field has been set.
func (o *VerificationAllOf) HasPersonId() bool {
	if o != nil && o.PersonId != nil {
		return true
	}

	return false
}

// SetPersonId gets a reference to the given string and assigns it to the PersonId field.
func (o *VerificationAllOf) SetPersonId(v string) {
	o.PersonId = &v
}

// GetResult returns the Result field value if set, zero value otherwise.
func (o *VerificationAllOf) GetResult() VerificationResult {
	if o == nil || o.Result == nil {
		var ret VerificationResult
		return ret
	}
	return *o.Result
}

// GetResultOk returns a tuple with the Result field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VerificationAllOf) GetResultOk() (*VerificationResult, bool) {
	if o == nil || o.Result == nil {
		return nil, false
	}
	return o.Result, true
}

// HasResult returns a boolean if a field has been set.
func (o *VerificationAllOf) HasResult() bool {
	if o != nil && o.Result != nil {
		return true
	}

	return false
}

// SetResult gets a reference to the given VerificationResult and assigns it to the Result field.
func (o *VerificationAllOf) SetResult(v VerificationResult) {
	o.Result = &v
}

// GetVendorInfo returns the VendorInfo field value if set, zero value otherwise.
func (o *VerificationAllOf) GetVendorInfo() VerificationVendorInfo {
	if o == nil || o.VendorInfo == nil {
		var ret VerificationVendorInfo
		return ret
	}
	return *o.VendorInfo
}

// GetVendorInfoOk returns a tuple with the VendorInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VerificationAllOf) GetVendorInfoOk() (*VerificationVendorInfo, bool) {
	if o == nil || o.VendorInfo == nil {
		return nil, false
	}
	return o.VendorInfo, true
}

// HasVendorInfo returns a boolean if a field has been set.
func (o *VerificationAllOf) HasVendorInfo() bool {
	if o != nil && o.VendorInfo != nil {
		return true
	}

	return false
}

// SetVendorInfo gets a reference to the given VerificationVendorInfo and assigns it to the VendorInfo field.
func (o *VerificationAllOf) SetVendorInfo(v VerificationVendorInfo) {
	o.VendorInfo = &v
}

// GetVerificationTime returns the VerificationTime field value if set, zero value otherwise.
func (o *VerificationAllOf) GetVerificationTime() time.Time {
	if o == nil || o.VerificationTime == nil {
		var ret time.Time
		return ret
	}
	return *o.VerificationTime
}

// GetVerificationTimeOk returns a tuple with the VerificationTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VerificationAllOf) GetVerificationTimeOk() (*time.Time, bool) {
	if o == nil || o.VerificationTime == nil {
		return nil, false
	}
	return o.VerificationTime, true
}

// HasVerificationTime returns a boolean if a field has been set.
func (o *VerificationAllOf) HasVerificationTime() bool {
	if o != nil && o.VerificationTime != nil {
		return true
	}

	return false
}

// SetVerificationTime gets a reference to the given time.Time and assigns it to the VerificationTime field.
func (o *VerificationAllOf) SetVerificationTime(v time.Time) {
	o.VerificationTime = &v
}

// GetVerificationType returns the VerificationType field value if set, zero value otherwise.
func (o *VerificationAllOf) GetVerificationType() VerificationType1 {
	if o == nil || o.VerificationType == nil {
		var ret VerificationType1
		return ret
	}
	return *o.VerificationType
}

// GetVerificationTypeOk returns a tuple with the VerificationType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VerificationAllOf) GetVerificationTypeOk() (*VerificationType1, bool) {
	if o == nil || o.VerificationType == nil {
		return nil, false
	}
	return o.VerificationType, true
}

// HasVerificationType returns a boolean if a field has been set.
func (o *VerificationAllOf) HasVerificationType() bool {
	if o != nil && o.VerificationType != nil {
		return true
	}

	return false
}

// SetVerificationType gets a reference to the given VerificationType1 and assigns it to the VerificationType field.
func (o *VerificationAllOf) SetVerificationType(v VerificationType1) {
	o.VerificationType = &v
}

func (o VerificationAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.BusinessId != nil {
		toSerialize["business_id"] = o.BusinessId
	}
	if o.CreationTime != nil {
		toSerialize["creation_time"] = o.CreationTime
	}
	if o.Details != nil {
		toSerialize["details"] = o.Details
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.LastUpdatedTime != nil {
		toSerialize["last_updated_time"] = o.LastUpdatedTime
	}
	if o.Metadata != nil {
		toSerialize["metadata"] = o.Metadata
	}
	if o.PersonId != nil {
		toSerialize["person_id"] = o.PersonId
	}
	if o.Result != nil {
		toSerialize["result"] = o.Result
	}
	if o.VendorInfo != nil {
		toSerialize["vendor_info"] = o.VendorInfo
	}
	if o.VerificationTime != nil {
		toSerialize["verification_time"] = o.VerificationTime
	}
	if o.VerificationType != nil {
		toSerialize["verification_type"] = o.VerificationType
	}
	return json.Marshal(toSerialize)
}

type NullableVerificationAllOf struct {
	value *VerificationAllOf
	isSet bool
}

func (v NullableVerificationAllOf) Get() *VerificationAllOf {
	return v.value
}

func (v *NullableVerificationAllOf) Set(val *VerificationAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableVerificationAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableVerificationAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVerificationAllOf(val *VerificationAllOf) *NullableVerificationAllOf {
	return &NullableVerificationAllOf{value: val, isSet: true}
}

func (v NullableVerificationAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVerificationAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
