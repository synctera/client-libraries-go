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

// Relationship1 struct for Relationship1
type Relationship1 struct {
	// ID of related entity
	Id               string           `json:"id"`
	RelationshipRole RelationshipRole `json:"relationship_role"`
}

// NewRelationship1 instantiates a new Relationship1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRelationship1(id string, relationshipRole RelationshipRole) *Relationship1 {
	this := Relationship1{}
	this.Id = id
	this.RelationshipRole = relationshipRole
	return &this
}

// NewRelationship1WithDefaults instantiates a new Relationship1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRelationship1WithDefaults() *Relationship1 {
	this := Relationship1{}
	return &this
}

// GetId returns the Id field value
func (o *Relationship1) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *Relationship1) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *Relationship1) SetId(v string) {
	o.Id = v
}

// GetRelationshipRole returns the RelationshipRole field value
func (o *Relationship1) GetRelationshipRole() RelationshipRole {
	if o == nil {
		var ret RelationshipRole
		return ret
	}

	return o.RelationshipRole
}

// GetRelationshipRoleOk returns a tuple with the RelationshipRole field value
// and a boolean to check if the value has been set.
func (o *Relationship1) GetRelationshipRoleOk() (*RelationshipRole, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RelationshipRole, true
}

// SetRelationshipRole sets field value
func (o *Relationship1) SetRelationshipRole(v RelationshipRole) {
	o.RelationshipRole = v
}

func (o Relationship1) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["relationship_role"] = o.RelationshipRole
	}
	return json.Marshal(toSerialize)
}

type NullableRelationship1 struct {
	value *Relationship1
	isSet bool
}

func (v NullableRelationship1) Get() *Relationship1 {
	return v.value
}

func (v *NullableRelationship1) Set(val *Relationship1) {
	v.value = val
	v.isSet = true
}

func (v NullableRelationship1) IsSet() bool {
	return v.isSet
}

func (v *NullableRelationship1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRelationship1(val *Relationship1) *NullableRelationship1 {
	return &NullableRelationship1{value: val, isSet: true}
}

func (v NullableRelationship1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRelationship1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
