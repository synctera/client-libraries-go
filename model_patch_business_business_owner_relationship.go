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

// PatchBusinessBusinessOwnerRelationship Denotes the relationship between specified businesses.
type PatchBusinessBusinessOwnerRelationship struct {
	AdditionalData *AdditionalOwnerData `json:"additional_data,omitempty"`
	// The date and time the resource was created.
	CreationTime *time.Time `json:"creation_time,omitempty"`
	// Unique ID for the subject business.
	FromBusinessId *string `json:"from_business_id,omitempty"`
	// Relationship unique identifier.
	Id *string `json:"id,omitempty"`
	// The date and time the resource was last updated.
	LastUpdatedTime *time.Time `json:"last_updated_time,omitempty"`
	// Optional field to store additional information about the resource. Intended to be used by the integrator to store non-sensitive data.
	Metadata *map[string]interface{} `json:"metadata,omitempty"`
	// The relationship type. One of the following: * `BENEFICIAL_OWNER_OF` – a person who directly or indirectly owns a portion of the business. * `MANAGING_PERSON_OF` – a person who is an officer, director, or other notable person of an organization. * `OWNER_OF` – a business with ownership of another business.
	RelationshipType string `json:"relationship_type"`
	// Unique ID for the related business.
	ToBusinessId *string `json:"to_business_id,omitempty"`
}

// NewPatchBusinessBusinessOwnerRelationship instantiates a new PatchBusinessBusinessOwnerRelationship object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPatchBusinessBusinessOwnerRelationship(relationshipType string) *PatchBusinessBusinessOwnerRelationship {
	this := PatchBusinessBusinessOwnerRelationship{}
	this.RelationshipType = relationshipType
	return &this
}

// NewPatchBusinessBusinessOwnerRelationshipWithDefaults instantiates a new PatchBusinessBusinessOwnerRelationship object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPatchBusinessBusinessOwnerRelationshipWithDefaults() *PatchBusinessBusinessOwnerRelationship {
	this := PatchBusinessBusinessOwnerRelationship{}
	return &this
}

// GetAdditionalData returns the AdditionalData field value if set, zero value otherwise.
func (o *PatchBusinessBusinessOwnerRelationship) GetAdditionalData() AdditionalOwnerData {
	if o == nil || o.AdditionalData == nil {
		var ret AdditionalOwnerData
		return ret
	}
	return *o.AdditionalData
}

// GetAdditionalDataOk returns a tuple with the AdditionalData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchBusinessBusinessOwnerRelationship) GetAdditionalDataOk() (*AdditionalOwnerData, bool) {
	if o == nil || o.AdditionalData == nil {
		return nil, false
	}
	return o.AdditionalData, true
}

// HasAdditionalData returns a boolean if a field has been set.
func (o *PatchBusinessBusinessOwnerRelationship) HasAdditionalData() bool {
	if o != nil && o.AdditionalData != nil {
		return true
	}

	return false
}

// SetAdditionalData gets a reference to the given AdditionalOwnerData and assigns it to the AdditionalData field.
func (o *PatchBusinessBusinessOwnerRelationship) SetAdditionalData(v AdditionalOwnerData) {
	o.AdditionalData = &v
}

// GetCreationTime returns the CreationTime field value if set, zero value otherwise.
func (o *PatchBusinessBusinessOwnerRelationship) GetCreationTime() time.Time {
	if o == nil || o.CreationTime == nil {
		var ret time.Time
		return ret
	}
	return *o.CreationTime
}

// GetCreationTimeOk returns a tuple with the CreationTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchBusinessBusinessOwnerRelationship) GetCreationTimeOk() (*time.Time, bool) {
	if o == nil || o.CreationTime == nil {
		return nil, false
	}
	return o.CreationTime, true
}

// HasCreationTime returns a boolean if a field has been set.
func (o *PatchBusinessBusinessOwnerRelationship) HasCreationTime() bool {
	if o != nil && o.CreationTime != nil {
		return true
	}

	return false
}

// SetCreationTime gets a reference to the given time.Time and assigns it to the CreationTime field.
func (o *PatchBusinessBusinessOwnerRelationship) SetCreationTime(v time.Time) {
	o.CreationTime = &v
}

// GetFromBusinessId returns the FromBusinessId field value if set, zero value otherwise.
func (o *PatchBusinessBusinessOwnerRelationship) GetFromBusinessId() string {
	if o == nil || o.FromBusinessId == nil {
		var ret string
		return ret
	}
	return *o.FromBusinessId
}

// GetFromBusinessIdOk returns a tuple with the FromBusinessId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchBusinessBusinessOwnerRelationship) GetFromBusinessIdOk() (*string, bool) {
	if o == nil || o.FromBusinessId == nil {
		return nil, false
	}
	return o.FromBusinessId, true
}

// HasFromBusinessId returns a boolean if a field has been set.
func (o *PatchBusinessBusinessOwnerRelationship) HasFromBusinessId() bool {
	if o != nil && o.FromBusinessId != nil {
		return true
	}

	return false
}

// SetFromBusinessId gets a reference to the given string and assigns it to the FromBusinessId field.
func (o *PatchBusinessBusinessOwnerRelationship) SetFromBusinessId(v string) {
	o.FromBusinessId = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *PatchBusinessBusinessOwnerRelationship) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchBusinessBusinessOwnerRelationship) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *PatchBusinessBusinessOwnerRelationship) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *PatchBusinessBusinessOwnerRelationship) SetId(v string) {
	o.Id = &v
}

// GetLastUpdatedTime returns the LastUpdatedTime field value if set, zero value otherwise.
func (o *PatchBusinessBusinessOwnerRelationship) GetLastUpdatedTime() time.Time {
	if o == nil || o.LastUpdatedTime == nil {
		var ret time.Time
		return ret
	}
	return *o.LastUpdatedTime
}

// GetLastUpdatedTimeOk returns a tuple with the LastUpdatedTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchBusinessBusinessOwnerRelationship) GetLastUpdatedTimeOk() (*time.Time, bool) {
	if o == nil || o.LastUpdatedTime == nil {
		return nil, false
	}
	return o.LastUpdatedTime, true
}

// HasLastUpdatedTime returns a boolean if a field has been set.
func (o *PatchBusinessBusinessOwnerRelationship) HasLastUpdatedTime() bool {
	if o != nil && o.LastUpdatedTime != nil {
		return true
	}

	return false
}

// SetLastUpdatedTime gets a reference to the given time.Time and assigns it to the LastUpdatedTime field.
func (o *PatchBusinessBusinessOwnerRelationship) SetLastUpdatedTime(v time.Time) {
	o.LastUpdatedTime = &v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *PatchBusinessBusinessOwnerRelationship) GetMetadata() map[string]interface{} {
	if o == nil || o.Metadata == nil {
		var ret map[string]interface{}
		return ret
	}
	return *o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchBusinessBusinessOwnerRelationship) GetMetadataOk() (*map[string]interface{}, bool) {
	if o == nil || o.Metadata == nil {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *PatchBusinessBusinessOwnerRelationship) HasMetadata() bool {
	if o != nil && o.Metadata != nil {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given map[string]interface{} and assigns it to the Metadata field.
func (o *PatchBusinessBusinessOwnerRelationship) SetMetadata(v map[string]interface{}) {
	o.Metadata = &v
}

// GetRelationshipType returns the RelationshipType field value
func (o *PatchBusinessBusinessOwnerRelationship) GetRelationshipType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RelationshipType
}

// GetRelationshipTypeOk returns a tuple with the RelationshipType field value
// and a boolean to check if the value has been set.
func (o *PatchBusinessBusinessOwnerRelationship) GetRelationshipTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RelationshipType, true
}

// SetRelationshipType sets field value
func (o *PatchBusinessBusinessOwnerRelationship) SetRelationshipType(v string) {
	o.RelationshipType = v
}

// GetToBusinessId returns the ToBusinessId field value if set, zero value otherwise.
func (o *PatchBusinessBusinessOwnerRelationship) GetToBusinessId() string {
	if o == nil || o.ToBusinessId == nil {
		var ret string
		return ret
	}
	return *o.ToBusinessId
}

// GetToBusinessIdOk returns a tuple with the ToBusinessId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchBusinessBusinessOwnerRelationship) GetToBusinessIdOk() (*string, bool) {
	if o == nil || o.ToBusinessId == nil {
		return nil, false
	}
	return o.ToBusinessId, true
}

// HasToBusinessId returns a boolean if a field has been set.
func (o *PatchBusinessBusinessOwnerRelationship) HasToBusinessId() bool {
	if o != nil && o.ToBusinessId != nil {
		return true
	}

	return false
}

// SetToBusinessId gets a reference to the given string and assigns it to the ToBusinessId field.
func (o *PatchBusinessBusinessOwnerRelationship) SetToBusinessId(v string) {
	o.ToBusinessId = &v
}

func (o PatchBusinessBusinessOwnerRelationship) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AdditionalData != nil {
		toSerialize["additional_data"] = o.AdditionalData
	}
	if o.CreationTime != nil {
		toSerialize["creation_time"] = o.CreationTime
	}
	if o.FromBusinessId != nil {
		toSerialize["from_business_id"] = o.FromBusinessId
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
	if true {
		toSerialize["relationship_type"] = o.RelationshipType
	}
	if o.ToBusinessId != nil {
		toSerialize["to_business_id"] = o.ToBusinessId
	}
	return json.Marshal(toSerialize)
}

type NullablePatchBusinessBusinessOwnerRelationship struct {
	value *PatchBusinessBusinessOwnerRelationship
	isSet bool
}

func (v NullablePatchBusinessBusinessOwnerRelationship) Get() *PatchBusinessBusinessOwnerRelationship {
	return v.value
}

func (v *NullablePatchBusinessBusinessOwnerRelationship) Set(val *PatchBusinessBusinessOwnerRelationship) {
	v.value = val
	v.isSet = true
}

func (v NullablePatchBusinessBusinessOwnerRelationship) IsSet() bool {
	return v.isSet
}

func (v *NullablePatchBusinessBusinessOwnerRelationship) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePatchBusinessBusinessOwnerRelationship(val *PatchBusinessBusinessOwnerRelationship) *NullablePatchBusinessBusinessOwnerRelationship {
	return &NullablePatchBusinessBusinessOwnerRelationship{value: val, isSet: true}
}

func (v NullablePatchBusinessBusinessOwnerRelationship) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePatchBusinessBusinessOwnerRelationship) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
