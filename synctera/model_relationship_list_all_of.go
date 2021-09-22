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

// RelationshipListAllOf struct for RelationshipListAllOf
type RelationshipListAllOf struct {
	// Array of relationships
	Relationships []Relationship `json:"relationships"`
}

// NewRelationshipListAllOf instantiates a new RelationshipListAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRelationshipListAllOf(relationships []Relationship) *RelationshipListAllOf {
	this := RelationshipListAllOf{}
	this.Relationships = relationships
	return &this
}

// NewRelationshipListAllOfWithDefaults instantiates a new RelationshipListAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRelationshipListAllOfWithDefaults() *RelationshipListAllOf {
	this := RelationshipListAllOf{}
	return &this
}

// GetRelationships returns the Relationships field value
func (o *RelationshipListAllOf) GetRelationships() []Relationship {
	if o == nil {
		var ret []Relationship
		return ret
	}

	return o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value
// and a boolean to check if the value has been set.
func (o *RelationshipListAllOf) GetRelationshipsOk() (*[]Relationship, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Relationships, true
}

// SetRelationships sets field value
func (o *RelationshipListAllOf) SetRelationships(v []Relationship) {
	o.Relationships = v
}

func (o RelationshipListAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["relationships"] = o.Relationships
	}
	return json.Marshal(toSerialize)
}

type NullableRelationshipListAllOf struct {
	value *RelationshipListAllOf
	isSet bool
}

func (v NullableRelationshipListAllOf) Get() *RelationshipListAllOf {
	return v.value
}

func (v *NullableRelationshipListAllOf) Set(val *RelationshipListAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableRelationshipListAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableRelationshipListAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRelationshipListAllOf(val *RelationshipListAllOf) *NullableRelationshipListAllOf {
	return &NullableRelationshipListAllOf{value: val, isSet: true}
}

func (v NullableRelationshipListAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRelationshipListAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


