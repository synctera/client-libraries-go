/*
Synctera API

<h2>Let's build something great.</h2><p>Welcome to the official reference documentation for Synctera APIs. Our APIs are the best way to automate your company's banking needs and are designed to be easy to understand and implement.</p><p>We're continuously growing this library and what you see here is just the start, but if you need something specific or have a question, <a class='text-blue-600' href='https://synctera.com/contact' target='_blank' rel='noreferrer'>contact us</a>.</p>

API version: 0.5.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package synctera

import (
	"encoding/json"
)

// WebhookList struct for WebhookList
type WebhookList struct {
	// Array of webhooks
	Webhooks []Webhook `json:"webhooks"`
	// If returned, use the next_page_token to query for the next page of results. Not returned if there are no more rows.
	NextPageToken *string `json:"next_page_token,omitempty"`
}

// NewWebhookList instantiates a new WebhookList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWebhookList(webhooks []Webhook) *WebhookList {
	this := WebhookList{}
	this.Webhooks = webhooks
	return &this
}

// NewWebhookListWithDefaults instantiates a new WebhookList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWebhookListWithDefaults() *WebhookList {
	this := WebhookList{}
	return &this
}

// GetWebhooks returns the Webhooks field value
func (o *WebhookList) GetWebhooks() []Webhook {
	if o == nil {
		var ret []Webhook
		return ret
	}

	return o.Webhooks
}

// GetWebhooksOk returns a tuple with the Webhooks field value
// and a boolean to check if the value has been set.
func (o *WebhookList) GetWebhooksOk() (*[]Webhook, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Webhooks, true
}

// SetWebhooks sets field value
func (o *WebhookList) SetWebhooks(v []Webhook) {
	o.Webhooks = v
}

// GetNextPageToken returns the NextPageToken field value if set, zero value otherwise.
func (o *WebhookList) GetNextPageToken() string {
	if o == nil || o.NextPageToken == nil {
		var ret string
		return ret
	}
	return *o.NextPageToken
}

// GetNextPageTokenOk returns a tuple with the NextPageToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebhookList) GetNextPageTokenOk() (*string, bool) {
	if o == nil || o.NextPageToken == nil {
		return nil, false
	}
	return o.NextPageToken, true
}

// HasNextPageToken returns a boolean if a field has been set.
func (o *WebhookList) HasNextPageToken() bool {
	if o != nil && o.NextPageToken != nil {
		return true
	}

	return false
}

// SetNextPageToken gets a reference to the given string and assigns it to the NextPageToken field.
func (o *WebhookList) SetNextPageToken(v string) {
	o.NextPageToken = &v
}

func (o WebhookList) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["webhooks"] = o.Webhooks
	}
	if o.NextPageToken != nil {
		toSerialize["next_page_token"] = o.NextPageToken
	}
	return json.Marshal(toSerialize)
}

type NullableWebhookList struct {
	value *WebhookList
	isSet bool
}

func (v NullableWebhookList) Get() *WebhookList {
	return v.value
}

func (v *NullableWebhookList) Set(val *WebhookList) {
	v.value = val
	v.isSet = true
}

func (v NullableWebhookList) IsSet() bool {
	return v.isSet
}

func (v *NullableWebhookList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWebhookList(val *WebhookList) *NullableWebhookList {
	return &NullableWebhookList{value: val, isSet: true}
}

func (v NullableWebhookList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWebhookList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
