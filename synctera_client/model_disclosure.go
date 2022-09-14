/*
Synctera API

<h2>Let's build something great.</h2><p>Welcome to the official reference documentation for Synctera APIs. Our APIs are the best way to automate your company's banking needs and are designed to be easy to understand and implement.</p><p>We're continuously growing this library and what you see here is just the start, but if you need something specific or have a question, <a class='text-blue-600' href='https://synctera.com/contact' target='_blank' rel='noreferrer'>contact us</a>.</p>

API version: 0.20.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package synctera_client

import (
	"encoding/json"
	"time"
)

// Disclosure Represents a disclosure.
type Disclosure struct {
	// Unique ID for the business. Exactly one of `business_id` or `person_id` must be set.
	BusinessId *string `json:"business_id,omitempty"`
	// The date and time the resource was created.
	CreationTime *time.Time `json:"creation_time,omitempty"`
	// Date and time the disclosure was made.
	DisclosureDate time.Time `json:"disclosure_date"`
	// Describes how the disclosure was shown and what the user did as a result. One of the following: * `DISPLAYED` —     The document was made visible to the user,     but they did not interact with it. * `VIEWED` —     The document was made visible to the user,     and they interacted enough to see the whole document (e.g. scrolled to the bottom). * `ACKNOWLEDGED` —     The document was made visible to the user,     and they took positive action to confirm that they have read and accepted the document.
	EventType string `json:"event_type"`
	// The unique identifier for this resource.
	Id *string `json:"id,omitempty"`
	// The date and time the resource was last updated.
	LastUpdatedTime *time.Time `json:"last_updated_time,omitempty"`
	// Optional field to store additional information about the resource. Intended to be used by the integrator to store non-sensitive data.
	Metadata map[string]interface{} `json:"metadata,omitempty"`
	// Unique ID for the person. Exactly one of `person_id` or `business_id` must be set.
	PersonId *string        `json:"person_id,omitempty"`
	Type     DisclosureType `json:"type"`
	// Version of the disclosure document.
	Version string `json:"version"`
}

// NewDisclosure instantiates a new Disclosure object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDisclosure(disclosureDate time.Time, eventType string, type_ DisclosureType, version string) *Disclosure {
	this := Disclosure{}
	this.DisclosureDate = disclosureDate
	this.EventType = eventType
	this.Type = type_
	this.Version = version
	return &this
}

// NewDisclosureWithDefaults instantiates a new Disclosure object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDisclosureWithDefaults() *Disclosure {
	this := Disclosure{}
	return &this
}

// GetBusinessId returns the BusinessId field value if set, zero value otherwise.
func (o *Disclosure) GetBusinessId() string {
	if o == nil || o.BusinessId == nil {
		var ret string
		return ret
	}
	return *o.BusinessId
}

// GetBusinessIdOk returns a tuple with the BusinessId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Disclosure) GetBusinessIdOk() (*string, bool) {
	if o == nil || o.BusinessId == nil {
		return nil, false
	}
	return o.BusinessId, true
}

// HasBusinessId returns a boolean if a field has been set.
func (o *Disclosure) HasBusinessId() bool {
	if o != nil && o.BusinessId != nil {
		return true
	}

	return false
}

// SetBusinessId gets a reference to the given string and assigns it to the BusinessId field.
func (o *Disclosure) SetBusinessId(v string) {
	o.BusinessId = &v
}

// GetCreationTime returns the CreationTime field value if set, zero value otherwise.
func (o *Disclosure) GetCreationTime() time.Time {
	if o == nil || o.CreationTime == nil {
		var ret time.Time
		return ret
	}
	return *o.CreationTime
}

// GetCreationTimeOk returns a tuple with the CreationTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Disclosure) GetCreationTimeOk() (*time.Time, bool) {
	if o == nil || o.CreationTime == nil {
		return nil, false
	}
	return o.CreationTime, true
}

// HasCreationTime returns a boolean if a field has been set.
func (o *Disclosure) HasCreationTime() bool {
	if o != nil && o.CreationTime != nil {
		return true
	}

	return false
}

// SetCreationTime gets a reference to the given time.Time and assigns it to the CreationTime field.
func (o *Disclosure) SetCreationTime(v time.Time) {
	o.CreationTime = &v
}

// GetDisclosureDate returns the DisclosureDate field value
func (o *Disclosure) GetDisclosureDate() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.DisclosureDate
}

// GetDisclosureDateOk returns a tuple with the DisclosureDate field value
// and a boolean to check if the value has been set.
func (o *Disclosure) GetDisclosureDateOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DisclosureDate, true
}

// SetDisclosureDate sets field value
func (o *Disclosure) SetDisclosureDate(v time.Time) {
	o.DisclosureDate = v
}

// GetEventType returns the EventType field value
func (o *Disclosure) GetEventType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EventType
}

// GetEventTypeOk returns a tuple with the EventType field value
// and a boolean to check if the value has been set.
func (o *Disclosure) GetEventTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EventType, true
}

// SetEventType sets field value
func (o *Disclosure) SetEventType(v string) {
	o.EventType = v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Disclosure) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Disclosure) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Disclosure) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *Disclosure) SetId(v string) {
	o.Id = &v
}

// GetLastUpdatedTime returns the LastUpdatedTime field value if set, zero value otherwise.
func (o *Disclosure) GetLastUpdatedTime() time.Time {
	if o == nil || o.LastUpdatedTime == nil {
		var ret time.Time
		return ret
	}
	return *o.LastUpdatedTime
}

// GetLastUpdatedTimeOk returns a tuple with the LastUpdatedTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Disclosure) GetLastUpdatedTimeOk() (*time.Time, bool) {
	if o == nil || o.LastUpdatedTime == nil {
		return nil, false
	}
	return o.LastUpdatedTime, true
}

// HasLastUpdatedTime returns a boolean if a field has been set.
func (o *Disclosure) HasLastUpdatedTime() bool {
	if o != nil && o.LastUpdatedTime != nil {
		return true
	}

	return false
}

// SetLastUpdatedTime gets a reference to the given time.Time and assigns it to the LastUpdatedTime field.
func (o *Disclosure) SetLastUpdatedTime(v time.Time) {
	o.LastUpdatedTime = &v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *Disclosure) GetMetadata() map[string]interface{} {
	if o == nil || o.Metadata == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Disclosure) GetMetadataOk() (map[string]interface{}, bool) {
	if o == nil || o.Metadata == nil {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *Disclosure) HasMetadata() bool {
	if o != nil && o.Metadata != nil {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given map[string]interface{} and assigns it to the Metadata field.
func (o *Disclosure) SetMetadata(v map[string]interface{}) {
	o.Metadata = v
}

// GetPersonId returns the PersonId field value if set, zero value otherwise.
func (o *Disclosure) GetPersonId() string {
	if o == nil || o.PersonId == nil {
		var ret string
		return ret
	}
	return *o.PersonId
}

// GetPersonIdOk returns a tuple with the PersonId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Disclosure) GetPersonIdOk() (*string, bool) {
	if o == nil || o.PersonId == nil {
		return nil, false
	}
	return o.PersonId, true
}

// HasPersonId returns a boolean if a field has been set.
func (o *Disclosure) HasPersonId() bool {
	if o != nil && o.PersonId != nil {
		return true
	}

	return false
}

// SetPersonId gets a reference to the given string and assigns it to the PersonId field.
func (o *Disclosure) SetPersonId(v string) {
	o.PersonId = &v
}

// GetType returns the Type field value
func (o *Disclosure) GetType() DisclosureType {
	if o == nil {
		var ret DisclosureType
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *Disclosure) GetTypeOk() (*DisclosureType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *Disclosure) SetType(v DisclosureType) {
	o.Type = v
}

// GetVersion returns the Version field value
func (o *Disclosure) GetVersion() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Version
}

// GetVersionOk returns a tuple with the Version field value
// and a boolean to check if the value has been set.
func (o *Disclosure) GetVersionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Version, true
}

// SetVersion sets field value
func (o *Disclosure) SetVersion(v string) {
	o.Version = v
}

func (o Disclosure) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.BusinessId != nil {
		toSerialize["business_id"] = o.BusinessId
	}
	if o.CreationTime != nil {
		toSerialize["creation_time"] = o.CreationTime
	}
	if true {
		toSerialize["disclosure_date"] = o.DisclosureDate
	}
	if true {
		toSerialize["event_type"] = o.EventType
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
	if true {
		toSerialize["type"] = o.Type
	}
	if true {
		toSerialize["version"] = o.Version
	}
	return json.Marshal(toSerialize)
}

type NullableDisclosure struct {
	value *Disclosure
	isSet bool
}

func (v NullableDisclosure) Get() *Disclosure {
	return v.value
}

func (v *NullableDisclosure) Set(val *Disclosure) {
	v.value = val
	v.isSet = true
}

func (v NullableDisclosure) IsSet() bool {
	return v.isSet
}

func (v *NullableDisclosure) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDisclosure(val *Disclosure) *NullableDisclosure {
	return &NullableDisclosure{value: val, isSet: true}
}

func (v NullableDisclosure) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDisclosure) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
