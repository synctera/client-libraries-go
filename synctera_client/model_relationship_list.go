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

// RelationshipList struct for RelationshipList
type RelationshipList struct {
	// Array of relationships
	Relationships []Relationship `json:"relationships"`
	// If returned, use the next_page_token to query for the next page of results. Not returned if there are no more rows.
	NextPageToken *string `json:"next_page_token,omitempty"`
}

// NewRelationshipList instantiates a new RelationshipList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRelationshipList(relationships []Relationship) *RelationshipList {
	this := RelationshipList{}
	this.Relationships = relationships
	return &this
}

// NewRelationshipListWithDefaults instantiates a new RelationshipList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRelationshipListWithDefaults() *RelationshipList {
	this := RelationshipList{}
	return &this
}

// GetRelationships returns the Relationships field value
func (o *RelationshipList) GetRelationships() []Relationship {
	if o == nil {
		var ret []Relationship
		return ret
	}

	return o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value
// and a boolean to check if the value has been set.
func (o *RelationshipList) GetRelationshipsOk() ([]Relationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.Relationships, true
}

// SetRelationships sets field value
func (o *RelationshipList) SetRelationships(v []Relationship) {
	o.Relationships = v
}

// GetNextPageToken returns the NextPageToken field value if set, zero value otherwise.
func (o *RelationshipList) GetNextPageToken() string {
	if o == nil || o.NextPageToken == nil {
		var ret string
		return ret
	}
	return *o.NextPageToken
}

// GetNextPageTokenOk returns a tuple with the NextPageToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RelationshipList) GetNextPageTokenOk() (*string, bool) {
	if o == nil || o.NextPageToken == nil {
		return nil, false
	}
	return o.NextPageToken, true
}

// HasNextPageToken returns a boolean if a field has been set.
func (o *RelationshipList) HasNextPageToken() bool {
	if o != nil && o.NextPageToken != nil {
		return true
	}

	return false
}

// SetNextPageToken gets a reference to the given string and assigns it to the NextPageToken field.
func (o *RelationshipList) SetNextPageToken(v string) {
	o.NextPageToken = &v
}

func (o RelationshipList) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["relationships"] = o.Relationships
	}
	if o.NextPageToken != nil {
		toSerialize["next_page_token"] = o.NextPageToken
	}
	return json.Marshal(toSerialize)
}

type NullableRelationshipList struct {
	value *RelationshipList
	isSet bool
}

func (v NullableRelationshipList) Get() *RelationshipList {
	return v.value
}

func (v *NullableRelationshipList) Set(val *RelationshipList) {
	v.value = val
	v.isSet = true
}

func (v NullableRelationshipList) IsSet() bool {
	return v.isSet
}

func (v *NullableRelationshipList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRelationshipList(val *RelationshipList) *NullableRelationshipList {
	return &NullableRelationshipList{value: val, isSet: true}
}

func (v NullableRelationshipList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRelationshipList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

