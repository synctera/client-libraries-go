/*
Synctera API

<h2>Let's build something great.</h2><p>Welcome to the official reference documentation for Synctera APIs. Our APIs are the best way to automate your company's banking needs and are designed to be easy to understand and implement.</p><p>We're continuously growing this library and what you see here is just the start, but if you need something specific or have a question, <a class='text-blue-600' href='https://synctera.com/contact' target='_blank' rel='noreferrer'>contact us</a>.</p> 

API version: 0.19.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package synctera_client

import (
	"encoding/json"
)

// WebhookListAllOf struct for WebhookListAllOf
type WebhookListAllOf struct {
	// Array of webhooks
	Webhooks []Webhook `json:"webhooks"`
}

// NewWebhookListAllOf instantiates a new WebhookListAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWebhookListAllOf(webhooks []Webhook) *WebhookListAllOf {
	this := WebhookListAllOf{}
	this.Webhooks = webhooks
	return &this
}

// NewWebhookListAllOfWithDefaults instantiates a new WebhookListAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWebhookListAllOfWithDefaults() *WebhookListAllOf {
	this := WebhookListAllOf{}
	return &this
}

// GetWebhooks returns the Webhooks field value
func (o *WebhookListAllOf) GetWebhooks() []Webhook {
	if o == nil {
		var ret []Webhook
		return ret
	}

	return o.Webhooks
}

// GetWebhooksOk returns a tuple with the Webhooks field value
// and a boolean to check if the value has been set.
func (o *WebhookListAllOf) GetWebhooksOk() ([]Webhook, bool) {
	if o == nil {
		return nil, false
	}
	return o.Webhooks, true
}

// SetWebhooks sets field value
func (o *WebhookListAllOf) SetWebhooks(v []Webhook) {
	o.Webhooks = v
}

func (o WebhookListAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["webhooks"] = o.Webhooks
	}
	return json.Marshal(toSerialize)
}

type NullableWebhookListAllOf struct {
	value *WebhookListAllOf
	isSet bool
}

func (v NullableWebhookListAllOf) Get() *WebhookListAllOf {
	return v.value
}

func (v *NullableWebhookListAllOf) Set(val *WebhookListAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableWebhookListAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableWebhookListAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWebhookListAllOf(val *WebhookListAllOf) *NullableWebhookListAllOf {
	return &NullableWebhookListAllOf{value: val, isSet: true}
}

func (v NullableWebhookListAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWebhookListAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

