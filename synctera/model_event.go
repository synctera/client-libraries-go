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

// Event Webhook event object
type Event struct {
	// Json string of object associated with the event. For example, if your event is ACCOUNT.STATUS_CHANGE, You can refer to Acccount to parse the account event to obtain the ID, status etc.  
	EventResource *map[string]interface{} `json:"event_resource,omitempty"`
	// Timestamp of the current event raised
	EventTime *time.Time `json:"event_time,omitempty"`
	// Unique event ID of the webhook request. Use event endpoints to get more event summary data
	Id *string `json:"id,omitempty"`
	// Metadata that stored in the webhook subscription
	Metadata *string `json:"metadata,omitempty"`
	// Response history of the webhook request
	ResponseHistory *[]EventResponseHistory `json:"response_history,omitempty"`
	// Current event status. Failing event will keep retry until it is purged.
	Status *string `json:"status,omitempty"`
	Type *EventTypeExplicit `json:"type,omitempty"`
	// Webhook the current event belongs to
	WebhookId *string `json:"webhook_id,omitempty"`
}

// NewEvent instantiates a new Event object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEvent() *Event {
	this := Event{}
	return &this
}

// NewEventWithDefaults instantiates a new Event object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEventWithDefaults() *Event {
	this := Event{}
	return &this
}

// GetEventResource returns the EventResource field value if set, zero value otherwise.
func (o *Event) GetEventResource() map[string]interface{} {
	if o == nil || o.EventResource == nil {
		var ret map[string]interface{}
		return ret
	}
	return *o.EventResource
}

// GetEventResourceOk returns a tuple with the EventResource field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Event) GetEventResourceOk() (*map[string]interface{}, bool) {
	if o == nil || o.EventResource == nil {
		return nil, false
	}
	return o.EventResource, true
}

// HasEventResource returns a boolean if a field has been set.
func (o *Event) HasEventResource() bool {
	if o != nil && o.EventResource != nil {
		return true
	}

	return false
}

// SetEventResource gets a reference to the given map[string]interface{} and assigns it to the EventResource field.
func (o *Event) SetEventResource(v map[string]interface{}) {
	o.EventResource = &v
}

// GetEventTime returns the EventTime field value if set, zero value otherwise.
func (o *Event) GetEventTime() time.Time {
	if o == nil || o.EventTime == nil {
		var ret time.Time
		return ret
	}
	return *o.EventTime
}

// GetEventTimeOk returns a tuple with the EventTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Event) GetEventTimeOk() (*time.Time, bool) {
	if o == nil || o.EventTime == nil {
		return nil, false
	}
	return o.EventTime, true
}

// HasEventTime returns a boolean if a field has been set.
func (o *Event) HasEventTime() bool {
	if o != nil && o.EventTime != nil {
		return true
	}

	return false
}

// SetEventTime gets a reference to the given time.Time and assigns it to the EventTime field.
func (o *Event) SetEventTime(v time.Time) {
	o.EventTime = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Event) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Event) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Event) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *Event) SetId(v string) {
	o.Id = &v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *Event) GetMetadata() string {
	if o == nil || o.Metadata == nil {
		var ret string
		return ret
	}
	return *o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Event) GetMetadataOk() (*string, bool) {
	if o == nil || o.Metadata == nil {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *Event) HasMetadata() bool {
	if o != nil && o.Metadata != nil {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given string and assigns it to the Metadata field.
func (o *Event) SetMetadata(v string) {
	o.Metadata = &v
}

// GetResponseHistory returns the ResponseHistory field value if set, zero value otherwise.
func (o *Event) GetResponseHistory() []EventResponseHistory {
	if o == nil || o.ResponseHistory == nil {
		var ret []EventResponseHistory
		return ret
	}
	return *o.ResponseHistory
}

// GetResponseHistoryOk returns a tuple with the ResponseHistory field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Event) GetResponseHistoryOk() (*[]EventResponseHistory, bool) {
	if o == nil || o.ResponseHistory == nil {
		return nil, false
	}
	return o.ResponseHistory, true
}

// HasResponseHistory returns a boolean if a field has been set.
func (o *Event) HasResponseHistory() bool {
	if o != nil && o.ResponseHistory != nil {
		return true
	}

	return false
}

// SetResponseHistory gets a reference to the given []EventResponseHistory and assigns it to the ResponseHistory field.
func (o *Event) SetResponseHistory(v []EventResponseHistory) {
	o.ResponseHistory = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *Event) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Event) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *Event) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *Event) SetStatus(v string) {
	o.Status = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *Event) GetType() EventTypeExplicit {
	if o == nil || o.Type == nil {
		var ret EventTypeExplicit
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Event) GetTypeOk() (*EventTypeExplicit, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *Event) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given EventTypeExplicit and assigns it to the Type field.
func (o *Event) SetType(v EventTypeExplicit) {
	o.Type = &v
}

// GetWebhookId returns the WebhookId field value if set, zero value otherwise.
func (o *Event) GetWebhookId() string {
	if o == nil || o.WebhookId == nil {
		var ret string
		return ret
	}
	return *o.WebhookId
}

// GetWebhookIdOk returns a tuple with the WebhookId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Event) GetWebhookIdOk() (*string, bool) {
	if o == nil || o.WebhookId == nil {
		return nil, false
	}
	return o.WebhookId, true
}

// HasWebhookId returns a boolean if a field has been set.
func (o *Event) HasWebhookId() bool {
	if o != nil && o.WebhookId != nil {
		return true
	}

	return false
}

// SetWebhookId gets a reference to the given string and assigns it to the WebhookId field.
func (o *Event) SetWebhookId(v string) {
	o.WebhookId = &v
}

func (o Event) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.EventResource != nil {
		toSerialize["event_resource"] = o.EventResource
	}
	if o.EventTime != nil {
		toSerialize["event_time"] = o.EventTime
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Metadata != nil {
		toSerialize["metadata"] = o.Metadata
	}
	if o.ResponseHistory != nil {
		toSerialize["response_history"] = o.ResponseHistory
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.WebhookId != nil {
		toSerialize["webhook_id"] = o.WebhookId
	}
	return json.Marshal(toSerialize)
}

type NullableEvent struct {
	value *Event
	isSet bool
}

func (v NullableEvent) Get() *Event {
	return v.value
}

func (v *NullableEvent) Set(val *Event) {
	v.value = val
	v.isSet = true
}

func (v NullableEvent) IsSet() bool {
	return v.isSet
}

func (v *NullableEvent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEvent(val *Event) *NullableEvent {
	return &NullableEvent{value: val, isSet: true}
}

func (v NullableEvent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEvent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


