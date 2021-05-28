/*
 * Synctera API
 *
 * <h2>Let's build something great.</h2><p>Welcome to the official reference documentation for Synctera APIs. Our APIs are the best way to automate your company's banking needs and are designed to be easy to understand and implement.</p><p>We're continuously growing this library and what you see here is just the start, but if you need something specific or have a question, <a class='text-blue-600' href='https://synctera.com/contact' target='_blank' rel='noreferrer'>contact us</a>.</p> 
 *
 * API version: 0.5.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package synctera

import (
	"encoding/json"
	"time"
)

// Webhook Webhook object
type Webhook struct {
	// The unique ID of the webhook
	Id *string `json:"id,omitempty"`
	// URL that the webhook will send request to
	Url string `json:"url"`
	// A description of what the webhook is used for
	Description *string `json:"description,omitempty"`
	// A list of the events that will trigger the webhook
	EnabledEvents []EventType `json:"enabled_events"`
	// Additional information stored to the webhook
	Metadata *string `json:"metadata,omitempty"`
	// Timestamp that this webhook was created or the last time any field was changed
	LastUpdated *time.Time `json:"last_updated,omitempty"`
	// Set the webhook to be enabled or disabled
	IsEnabled bool `json:"is_enabled"`
}

// NewWebhook instantiates a new Webhook object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWebhook(url string, enabledEvents []EventType, isEnabled bool) *Webhook {
	this := Webhook{}
	this.Url = url
	this.EnabledEvents = enabledEvents
	this.IsEnabled = isEnabled
	return &this
}

// NewWebhookWithDefaults instantiates a new Webhook object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWebhookWithDefaults() *Webhook {
	this := Webhook{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Webhook) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Webhook) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Webhook) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *Webhook) SetId(v string) {
	o.Id = &v
}

// GetUrl returns the Url field value
func (o *Webhook) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *Webhook) GetUrlOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *Webhook) SetUrl(v string) {
	o.Url = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *Webhook) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Webhook) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *Webhook) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *Webhook) SetDescription(v string) {
	o.Description = &v
}

// GetEnabledEvents returns the EnabledEvents field value
func (o *Webhook) GetEnabledEvents() []EventType {
	if o == nil {
		var ret []EventType
		return ret
	}

	return o.EnabledEvents
}

// GetEnabledEventsOk returns a tuple with the EnabledEvents field value
// and a boolean to check if the value has been set.
func (o *Webhook) GetEnabledEventsOk() (*[]EventType, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.EnabledEvents, true
}

// SetEnabledEvents sets field value
func (o *Webhook) SetEnabledEvents(v []EventType) {
	o.EnabledEvents = v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *Webhook) GetMetadata() string {
	if o == nil || o.Metadata == nil {
		var ret string
		return ret
	}
	return *o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Webhook) GetMetadataOk() (*string, bool) {
	if o == nil || o.Metadata == nil {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *Webhook) HasMetadata() bool {
	if o != nil && o.Metadata != nil {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given string and assigns it to the Metadata field.
func (o *Webhook) SetMetadata(v string) {
	o.Metadata = &v
}

// GetLastUpdated returns the LastUpdated field value if set, zero value otherwise.
func (o *Webhook) GetLastUpdated() time.Time {
	if o == nil || o.LastUpdated == nil {
		var ret time.Time
		return ret
	}
	return *o.LastUpdated
}

// GetLastUpdatedOk returns a tuple with the LastUpdated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Webhook) GetLastUpdatedOk() (*time.Time, bool) {
	if o == nil || o.LastUpdated == nil {
		return nil, false
	}
	return o.LastUpdated, true
}

// HasLastUpdated returns a boolean if a field has been set.
func (o *Webhook) HasLastUpdated() bool {
	if o != nil && o.LastUpdated != nil {
		return true
	}

	return false
}

// SetLastUpdated gets a reference to the given time.Time and assigns it to the LastUpdated field.
func (o *Webhook) SetLastUpdated(v time.Time) {
	o.LastUpdated = &v
}

// GetIsEnabled returns the IsEnabled field value
func (o *Webhook) GetIsEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsEnabled
}

// GetIsEnabledOk returns a tuple with the IsEnabled field value
// and a boolean to check if the value has been set.
func (o *Webhook) GetIsEnabledOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.IsEnabled, true
}

// SetIsEnabled sets field value
func (o *Webhook) SetIsEnabled(v bool) {
	o.IsEnabled = v
}

func (o Webhook) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["url"] = o.Url
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if true {
		toSerialize["enabled_events"] = o.EnabledEvents
	}
	if o.Metadata != nil {
		toSerialize["metadata"] = o.Metadata
	}
	if o.LastUpdated != nil {
		toSerialize["last_updated"] = o.LastUpdated
	}
	if true {
		toSerialize["is_enabled"] = o.IsEnabled
	}
	return json.Marshal(toSerialize)
}

type NullableWebhook struct {
	value *Webhook
	isSet bool
}

func (v NullableWebhook) Get() *Webhook {
	return v.value
}

func (v *NullableWebhook) Set(val *Webhook) {
	v.value = val
	v.isSet = true
}

func (v NullableWebhook) IsSet() bool {
	return v.isSet
}

func (v *NullableWebhook) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWebhook(val *Webhook) *NullableWebhook {
	return &NullableWebhook{value: val, isSet: true}
}

func (v NullableWebhook) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWebhook) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

