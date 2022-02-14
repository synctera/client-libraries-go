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
)

// CardImageDetails struct for CardImageDetails
type CardImageDetails struct {
	// The unique identifier of a customer
	CustomerId string `json:"customer_id"`
	// The unique identifier of a card image
	Id     string          `json:"id"`
	Status CardImageStatus `json:"status"`
}

// NewCardImageDetails instantiates a new CardImageDetails object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCardImageDetails(customerId string, id string, status CardImageStatus) *CardImageDetails {
	this := CardImageDetails{}
	this.CustomerId = customerId
	this.Id = id
	this.Status = status
	return &this
}

// NewCardImageDetailsWithDefaults instantiates a new CardImageDetails object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCardImageDetailsWithDefaults() *CardImageDetails {
	this := CardImageDetails{}
	return &this
}

// GetCustomerId returns the CustomerId field value
func (o *CardImageDetails) GetCustomerId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CustomerId
}

// GetCustomerIdOk returns a tuple with the CustomerId field value
// and a boolean to check if the value has been set.
func (o *CardImageDetails) GetCustomerIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CustomerId, true
}

// SetCustomerId sets field value
func (o *CardImageDetails) SetCustomerId(v string) {
	o.CustomerId = v
}

// GetId returns the Id field value
func (o *CardImageDetails) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *CardImageDetails) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *CardImageDetails) SetId(v string) {
	o.Id = v
}

// GetStatus returns the Status field value
func (o *CardImageDetails) GetStatus() CardImageStatus {
	if o == nil {
		var ret CardImageStatus
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *CardImageDetails) GetStatusOk() (*CardImageStatus, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *CardImageDetails) SetStatus(v CardImageStatus) {
	o.Status = v
}

func (o CardImageDetails) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["customer_id"] = o.CustomerId
	}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["status"] = o.Status
	}
	return json.Marshal(toSerialize)
}

type NullableCardImageDetails struct {
	value *CardImageDetails
	isSet bool
}

func (v NullableCardImageDetails) Get() *CardImageDetails {
	return v.value
}

func (v *NullableCardImageDetails) Set(val *CardImageDetails) {
	v.value = val
	v.isSet = true
}

func (v NullableCardImageDetails) IsSet() bool {
	return v.isSet
}

func (v *NullableCardImageDetails) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCardImageDetails(val *CardImageDetails) *NullableCardImageDetails {
	return &NullableCardImageDetails{value: val, isSet: true}
}

func (v NullableCardImageDetails) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCardImageDetails) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
