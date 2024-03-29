/*
Synctera API

<h2>Let's build something great.</h2><p>Welcome to the official reference documentation for Synctera APIs. Our APIs are the best way to automate your company's banking needs and are designed to be easy to understand and implement.</p><p>We're continuously growing this library and what you see here is just the start, but if you need something specific or have a question, <a class='text-blue-600' href='https://synctera.com/contact' target='_blank' rel='noreferrer'>contact us</a>.</p>

API version: 0.20.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package synctera_client

import (
	"encoding/json"
)

// VerificationVendorXml struct for VerificationVendorXml
type VerificationVendorXml struct {
	// Describes the content-type encoding received from the vendor.
	ContentType string `json:"content_type"`
	// Array of vendor specific information.
	Details []VerificationVendorInfoDetail `json:"details,omitempty"`
	// Name of the vendor used.
	Vendor string `json:"vendor"`
	// Data representaion in XML.
	Xml string `json:"xml"`
}

// NewVerificationVendorXml instantiates a new VerificationVendorXml object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVerificationVendorXml(contentType string, vendor string, xml string) *VerificationVendorXml {
	this := VerificationVendorXml{}
	this.ContentType = contentType
	this.Vendor = vendor
	this.Xml = xml
	return &this
}

// NewVerificationVendorXmlWithDefaults instantiates a new VerificationVendorXml object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVerificationVendorXmlWithDefaults() *VerificationVendorXml {
	this := VerificationVendorXml{}
	return &this
}

// GetContentType returns the ContentType field value
func (o *VerificationVendorXml) GetContentType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ContentType
}

// GetContentTypeOk returns a tuple with the ContentType field value
// and a boolean to check if the value has been set.
func (o *VerificationVendorXml) GetContentTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ContentType, true
}

// SetContentType sets field value
func (o *VerificationVendorXml) SetContentType(v string) {
	o.ContentType = v
}

// GetDetails returns the Details field value if set, zero value otherwise.
func (o *VerificationVendorXml) GetDetails() []VerificationVendorInfoDetail {
	if o == nil || o.Details == nil {
		var ret []VerificationVendorInfoDetail
		return ret
	}
	return o.Details
}

// GetDetailsOk returns a tuple with the Details field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VerificationVendorXml) GetDetailsOk() ([]VerificationVendorInfoDetail, bool) {
	if o == nil || o.Details == nil {
		return nil, false
	}
	return o.Details, true
}

// HasDetails returns a boolean if a field has been set.
func (o *VerificationVendorXml) HasDetails() bool {
	if o != nil && o.Details != nil {
		return true
	}

	return false
}

// SetDetails gets a reference to the given []VerificationVendorInfoDetail and assigns it to the Details field.
func (o *VerificationVendorXml) SetDetails(v []VerificationVendorInfoDetail) {
	o.Details = v
}

// GetVendor returns the Vendor field value
func (o *VerificationVendorXml) GetVendor() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Vendor
}

// GetVendorOk returns a tuple with the Vendor field value
// and a boolean to check if the value has been set.
func (o *VerificationVendorXml) GetVendorOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Vendor, true
}

// SetVendor sets field value
func (o *VerificationVendorXml) SetVendor(v string) {
	o.Vendor = v
}

// GetXml returns the Xml field value
func (o *VerificationVendorXml) GetXml() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Xml
}

// GetXmlOk returns a tuple with the Xml field value
// and a boolean to check if the value has been set.
func (o *VerificationVendorXml) GetXmlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Xml, true
}

// SetXml sets field value
func (o *VerificationVendorXml) SetXml(v string) {
	o.Xml = v
}

func (o VerificationVendorXml) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["content_type"] = o.ContentType
	}
	if o.Details != nil {
		toSerialize["details"] = o.Details
	}
	if true {
		toSerialize["vendor"] = o.Vendor
	}
	if true {
		toSerialize["xml"] = o.Xml
	}
	return json.Marshal(toSerialize)
}

type NullableVerificationVendorXml struct {
	value *VerificationVendorXml
	isSet bool
}

func (v NullableVerificationVendorXml) Get() *VerificationVendorXml {
	return v.value
}

func (v *NullableVerificationVendorXml) Set(val *VerificationVendorXml) {
	v.value = val
	v.isSet = true
}

func (v NullableVerificationVendorXml) IsSet() bool {
	return v.isSet
}

func (v *NullableVerificationVendorXml) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVerificationVendorXml(val *VerificationVendorXml) *NullableVerificationVendorXml {
	return &NullableVerificationVendorXml{value: val, isSet: true}
}

func (v NullableVerificationVendorXml) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVerificationVendorXml) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
