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

// SchemasRelationship struct for SchemasRelationship
type SchemasRelationship struct {
	// Connection ID of the account
	ConnectId *string `json:"connect_id,omitempty"`
	// Relationship type
	RelationshipType string `json:"relationship_type"`
	// Customer that the current account is associated with
	CustomerId string `json:"customer_id"`
}

// NewSchemasRelationship instantiates a new SchemasRelationship object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSchemasRelationship(relationshipType string, customerId string) *SchemasRelationship {
	this := SchemasRelationship{}
	this.RelationshipType = relationshipType
	this.CustomerId = customerId
	return &this
}

// NewSchemasRelationshipWithDefaults instantiates a new SchemasRelationship object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSchemasRelationshipWithDefaults() *SchemasRelationship {
	this := SchemasRelationship{}
	return &this
}

// GetConnectId returns the ConnectId field value if set, zero value otherwise.
func (o *SchemasRelationship) GetConnectId() string {
	if o == nil || o.ConnectId == nil {
		var ret string
		return ret
	}
	return *o.ConnectId
}

// GetConnectIdOk returns a tuple with the ConnectId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SchemasRelationship) GetConnectIdOk() (*string, bool) {
	if o == nil || o.ConnectId == nil {
		return nil, false
	}
	return o.ConnectId, true
}

// HasConnectId returns a boolean if a field has been set.
func (o *SchemasRelationship) HasConnectId() bool {
	if o != nil && o.ConnectId != nil {
		return true
	}

	return false
}

// SetConnectId gets a reference to the given string and assigns it to the ConnectId field.
func (o *SchemasRelationship) SetConnectId(v string) {
	o.ConnectId = &v
}

// GetRelationshipType returns the RelationshipType field value
func (o *SchemasRelationship) GetRelationshipType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RelationshipType
}

// GetRelationshipTypeOk returns a tuple with the RelationshipType field value
// and a boolean to check if the value has been set.
func (o *SchemasRelationship) GetRelationshipTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.RelationshipType, true
}

// SetRelationshipType sets field value
func (o *SchemasRelationship) SetRelationshipType(v string) {
	o.RelationshipType = v
}

// GetCustomerId returns the CustomerId field value
func (o *SchemasRelationship) GetCustomerId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CustomerId
}

// GetCustomerIdOk returns a tuple with the CustomerId field value
// and a boolean to check if the value has been set.
func (o *SchemasRelationship) GetCustomerIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.CustomerId, true
}

// SetCustomerId sets field value
func (o *SchemasRelationship) SetCustomerId(v string) {
	o.CustomerId = v
}

func (o SchemasRelationship) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ConnectId != nil {
		toSerialize["connect_id"] = o.ConnectId
	}
	if true {
		toSerialize["relationship_type"] = o.RelationshipType
	}
	if true {
		toSerialize["customer_id"] = o.CustomerId
	}
	return json.Marshal(toSerialize)
}

type NullableSchemasRelationship struct {
	value *SchemasRelationship
	isSet bool
}

func (v NullableSchemasRelationship) Get() *SchemasRelationship {
	return v.value
}

func (v *NullableSchemasRelationship) Set(val *SchemasRelationship) {
	v.value = val
	v.isSet = true
}

func (v NullableSchemasRelationship) IsSet() bool {
	return v.isSet
}

func (v *NullableSchemasRelationship) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSchemasRelationship(val *SchemasRelationship) *NullableSchemasRelationship {
	return &NullableSchemasRelationship{value: val, isSet: true}
}

func (v NullableSchemasRelationship) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSchemasRelationship) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

