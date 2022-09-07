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

// GatewayResponse struct for GatewayResponse
type GatewayResponse struct {
	// Current status of the Authorization gateway
	Active bool `json:"active"`
	// List of Card Product unique identifiers that will utilize the Gateway
	CardProducts []string `json:"card_products"`
	// The timestamp representing when the gateway config request was made
	CreationTime time.Time `json:"creation_time"`
	// Custom Headers of the Authorization gateway
	CustomHeaders *map[string]string `json:"custom_headers,omitempty"`
	// Gateway ID
	Id string `json:"id"`
	// The timestamp representing when the gateway config was last modified at
	LastModifiedTime time.Time `json:"last_modified_time"`
	// URL of the Authorization gateway
	Url string `json:"url"`
}

// NewGatewayResponse instantiates a new GatewayResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGatewayResponse(active bool, cardProducts []string, creationTime time.Time, id string, lastModifiedTime time.Time, url string) *GatewayResponse {
	this := GatewayResponse{}
	this.Active = active
	this.CardProducts = cardProducts
	this.CreationTime = creationTime
	this.Id = id
	this.LastModifiedTime = lastModifiedTime
	this.Url = url
	return &this
}

// NewGatewayResponseWithDefaults instantiates a new GatewayResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGatewayResponseWithDefaults() *GatewayResponse {
	this := GatewayResponse{}
	return &this
}

// GetActive returns the Active field value
func (o *GatewayResponse) GetActive() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Active
}

// GetActiveOk returns a tuple with the Active field value
// and a boolean to check if the value has been set.
func (o *GatewayResponse) GetActiveOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Active, true
}

// SetActive sets field value
func (o *GatewayResponse) SetActive(v bool) {
	o.Active = v
}

// GetCardProducts returns the CardProducts field value
func (o *GatewayResponse) GetCardProducts() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.CardProducts
}

// GetCardProductsOk returns a tuple with the CardProducts field value
// and a boolean to check if the value has been set.
func (o *GatewayResponse) GetCardProductsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CardProducts, true
}

// SetCardProducts sets field value
func (o *GatewayResponse) SetCardProducts(v []string) {
	o.CardProducts = v
}

// GetCreationTime returns the CreationTime field value
func (o *GatewayResponse) GetCreationTime() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreationTime
}

// GetCreationTimeOk returns a tuple with the CreationTime field value
// and a boolean to check if the value has been set.
func (o *GatewayResponse) GetCreationTimeOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreationTime, true
}

// SetCreationTime sets field value
func (o *GatewayResponse) SetCreationTime(v time.Time) {
	o.CreationTime = v
}

// GetCustomHeaders returns the CustomHeaders field value if set, zero value otherwise.
func (o *GatewayResponse) GetCustomHeaders() map[string]string {
	if o == nil || o.CustomHeaders == nil {
		var ret map[string]string
		return ret
	}
	return *o.CustomHeaders
}

// GetCustomHeadersOk returns a tuple with the CustomHeaders field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayResponse) GetCustomHeadersOk() (*map[string]string, bool) {
	if o == nil || o.CustomHeaders == nil {
		return nil, false
	}
	return o.CustomHeaders, true
}

// HasCustomHeaders returns a boolean if a field has been set.
func (o *GatewayResponse) HasCustomHeaders() bool {
	if o != nil && o.CustomHeaders != nil {
		return true
	}

	return false
}

// SetCustomHeaders gets a reference to the given map[string]string and assigns it to the CustomHeaders field.
func (o *GatewayResponse) SetCustomHeaders(v map[string]string) {
	o.CustomHeaders = &v
}

// GetId returns the Id field value
func (o *GatewayResponse) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *GatewayResponse) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *GatewayResponse) SetId(v string) {
	o.Id = v
}

// GetLastModifiedTime returns the LastModifiedTime field value
func (o *GatewayResponse) GetLastModifiedTime() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.LastModifiedTime
}

// GetLastModifiedTimeOk returns a tuple with the LastModifiedTime field value
// and a boolean to check if the value has been set.
func (o *GatewayResponse) GetLastModifiedTimeOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LastModifiedTime, true
}

// SetLastModifiedTime sets field value
func (o *GatewayResponse) SetLastModifiedTime(v time.Time) {
	o.LastModifiedTime = v
}

// GetUrl returns the Url field value
func (o *GatewayResponse) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *GatewayResponse) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *GatewayResponse) SetUrl(v string) {
	o.Url = v
}

func (o GatewayResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["active"] = o.Active
	}
	if true {
		toSerialize["card_products"] = o.CardProducts
	}
	if true {
		toSerialize["creation_time"] = o.CreationTime
	}
	if o.CustomHeaders != nil {
		toSerialize["custom_headers"] = o.CustomHeaders
	}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["last_modified_time"] = o.LastModifiedTime
	}
	if true {
		toSerialize["url"] = o.Url
	}
	return json.Marshal(toSerialize)
}

type NullableGatewayResponse struct {
	value *GatewayResponse
	isSet bool
}

func (v NullableGatewayResponse) Get() *GatewayResponse {
	return v.value
}

func (v *NullableGatewayResponse) Set(val *GatewayResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGatewayResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGatewayResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGatewayResponse(val *GatewayResponse) *NullableGatewayResponse {
	return &NullableGatewayResponse{value: val, isSet: true}
}

func (v NullableGatewayResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGatewayResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

