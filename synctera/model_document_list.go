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

// DocumentList struct for DocumentList
type DocumentList struct {
	// Array of documents
	Documents []Document `json:"documents"`
	// If returned, use the next_page_token to query for the next page of results. Not returned if there are no more rows.
	NextPageToken *string `json:"next_page_token,omitempty"`
}

// NewDocumentList instantiates a new DocumentList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDocumentList(documents []Document) *DocumentList {
	this := DocumentList{}
	this.Documents = documents
	return &this
}

// NewDocumentListWithDefaults instantiates a new DocumentList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDocumentListWithDefaults() *DocumentList {
	this := DocumentList{}
	return &this
}

// GetDocuments returns the Documents field value
func (o *DocumentList) GetDocuments() []Document {
	if o == nil {
		var ret []Document
		return ret
	}

	return o.Documents
}

// GetDocumentsOk returns a tuple with the Documents field value
// and a boolean to check if the value has been set.
func (o *DocumentList) GetDocumentsOk() (*[]Document, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Documents, true
}

// SetDocuments sets field value
func (o *DocumentList) SetDocuments(v []Document) {
	o.Documents = v
}

// GetNextPageToken returns the NextPageToken field value if set, zero value otherwise.
func (o *DocumentList) GetNextPageToken() string {
	if o == nil || o.NextPageToken == nil {
		var ret string
		return ret
	}
	return *o.NextPageToken
}

// GetNextPageTokenOk returns a tuple with the NextPageToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DocumentList) GetNextPageTokenOk() (*string, bool) {
	if o == nil || o.NextPageToken == nil {
		return nil, false
	}
	return o.NextPageToken, true
}

// HasNextPageToken returns a boolean if a field has been set.
func (o *DocumentList) HasNextPageToken() bool {
	if o != nil && o.NextPageToken != nil {
		return true
	}

	return false
}

// SetNextPageToken gets a reference to the given string and assigns it to the NextPageToken field.
func (o *DocumentList) SetNextPageToken(v string) {
	o.NextPageToken = &v
}

func (o DocumentList) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["documents"] = o.Documents
	}
	if o.NextPageToken != nil {
		toSerialize["next_page_token"] = o.NextPageToken
	}
	return json.Marshal(toSerialize)
}

type NullableDocumentList struct {
	value *DocumentList
	isSet bool
}

func (v NullableDocumentList) Get() *DocumentList {
	return v.value
}

func (v *NullableDocumentList) Set(val *DocumentList) {
	v.value = val
	v.isSet = true
}

func (v NullableDocumentList) IsSet() bool {
	return v.isSet
}

func (v *NullableDocumentList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDocumentList(val *DocumentList) *NullableDocumentList {
	return &NullableDocumentList{value: val, isSet: true}
}

func (v NullableDocumentList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDocumentList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


