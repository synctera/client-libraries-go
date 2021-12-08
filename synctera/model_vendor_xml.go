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

// VendorXml struct for VendorXml
type VendorXml struct {
	// Describes the content-type encoding received from the vendor.
	ContentType string `json:"content_type"`
	// Array of vendor specific information.
	Details []Detail `json:"details"`
	// Name of the vendor used.
	Vendor string `json:"vendor"`
	// Data representaion in XML.
	Xml string `json:"xml"`
}

// NewVendorXml instantiates a new VendorXml object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVendorXml(contentType string, details []Detail, vendor string, xml string) *VendorXml {
	this := VendorXml{}
	this.ContentType = contentType
	this.Details = details
	this.Vendor = vendor
	this.Xml = xml
	return &this
}

// NewVendorXmlWithDefaults instantiates a new VendorXml object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVendorXmlWithDefaults() *VendorXml {
	this := VendorXml{}
	return &this
}

// GetContentType returns the ContentType field value
func (o *VendorXml) GetContentType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ContentType
}

// GetContentTypeOk returns a tuple with the ContentType field value
// and a boolean to check if the value has been set.
func (o *VendorXml) GetContentTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ContentType, true
}

// SetContentType sets field value
func (o *VendorXml) SetContentType(v string) {
	o.ContentType = v
}

// GetDetails returns the Details field value
func (o *VendorXml) GetDetails() []Detail {
	if o == nil {
		var ret []Detail
		return ret
	}

	return o.Details
}

// GetDetailsOk returns a tuple with the Details field value
// and a boolean to check if the value has been set.
func (o *VendorXml) GetDetailsOk() (*[]Detail, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Details, true
}

// SetDetails sets field value
func (o *VendorXml) SetDetails(v []Detail) {
	o.Details = v
}

// GetVendor returns the Vendor field value
func (o *VendorXml) GetVendor() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Vendor
}

// GetVendorOk returns a tuple with the Vendor field value
// and a boolean to check if the value has been set.
func (o *VendorXml) GetVendorOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Vendor, true
}

// SetVendor sets field value
func (o *VendorXml) SetVendor(v string) {
	o.Vendor = v
}

// GetXml returns the Xml field value
func (o *VendorXml) GetXml() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Xml
}

// GetXmlOk returns a tuple with the Xml field value
// and a boolean to check if the value has been set.
func (o *VendorXml) GetXmlOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Xml, true
}

// SetXml sets field value
func (o *VendorXml) SetXml(v string) {
	o.Xml = v
}

func (o VendorXml) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["content_type"] = o.ContentType
	}
	if true {
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

type NullableVendorXml struct {
	value *VendorXml
	isSet bool
}

func (v NullableVendorXml) Get() *VendorXml {
	return v.value
}

func (v *NullableVendorXml) Set(val *VendorXml) {
	v.value = val
	v.isSet = true
}

func (v NullableVendorXml) IsSet() bool {
	return v.isSet
}

func (v *NullableVendorXml) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVendorXml(val *VendorXml) *NullableVendorXml {
	return &NullableVendorXml{value: val, isSet: true}
}

func (v NullableVendorXml) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVendorXml) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


