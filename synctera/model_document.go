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

// Document struct for Document
type Document struct {
	// Document ID
	Id *string `json:"id,omitempty"`
	MediaType KycMediaType `json:"media_type"`
	DocumentType DocumentType `json:"document_type"`
	// Base64url encoded image
	ByteData *string `json:"byte_data,omitempty"`
}

// NewDocument instantiates a new Document object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDocument(mediaType KycMediaType, documentType DocumentType) *Document {
	this := Document{}
	this.MediaType = mediaType
	this.DocumentType = documentType
	return &this
}

// NewDocumentWithDefaults instantiates a new Document object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDocumentWithDefaults() *Document {
	this := Document{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Document) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Document) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Document) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *Document) SetId(v string) {
	o.Id = &v
}

// GetMediaType returns the MediaType field value
func (o *Document) GetMediaType() KycMediaType {
	if o == nil {
		var ret KycMediaType
		return ret
	}

	return o.MediaType
}

// GetMediaTypeOk returns a tuple with the MediaType field value
// and a boolean to check if the value has been set.
func (o *Document) GetMediaTypeOk() (*KycMediaType, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.MediaType, true
}

// SetMediaType sets field value
func (o *Document) SetMediaType(v KycMediaType) {
	o.MediaType = v
}

// GetDocumentType returns the DocumentType field value
func (o *Document) GetDocumentType() DocumentType {
	if o == nil {
		var ret DocumentType
		return ret
	}

	return o.DocumentType
}

// GetDocumentTypeOk returns a tuple with the DocumentType field value
// and a boolean to check if the value has been set.
func (o *Document) GetDocumentTypeOk() (*DocumentType, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.DocumentType, true
}

// SetDocumentType sets field value
func (o *Document) SetDocumentType(v DocumentType) {
	o.DocumentType = v
}

// GetByteData returns the ByteData field value if set, zero value otherwise.
func (o *Document) GetByteData() string {
	if o == nil || o.ByteData == nil {
		var ret string
		return ret
	}
	return *o.ByteData
}

// GetByteDataOk returns a tuple with the ByteData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Document) GetByteDataOk() (*string, bool) {
	if o == nil || o.ByteData == nil {
		return nil, false
	}
	return o.ByteData, true
}

// HasByteData returns a boolean if a field has been set.
func (o *Document) HasByteData() bool {
	if o != nil && o.ByteData != nil {
		return true
	}

	return false
}

// SetByteData gets a reference to the given string and assigns it to the ByteData field.
func (o *Document) SetByteData(v string) {
	o.ByteData = &v
}

func (o Document) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["media_type"] = o.MediaType
	}
	if true {
		toSerialize["document_type"] = o.DocumentType
	}
	if o.ByteData != nil {
		toSerialize["byte_data"] = o.ByteData
	}
	return json.Marshal(toSerialize)
}

type NullableDocument struct {
	value *Document
	isSet bool
}

func (v NullableDocument) Get() *Document {
	return v.value
}

func (v *NullableDocument) Set(val *Document) {
	v.value = val
	v.isSet = true
}

func (v NullableDocument) IsSet() bool {
	return v.isSet
}

func (v *NullableDocument) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDocument(val *Document) *NullableDocument {
	return &NullableDocument{value: val, isSet: true}
}

func (v NullableDocument) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDocument) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


