/*
Synctera API

<h2>Let's build something great.</h2><p>Welcome to the official reference documentation for Synctera APIs. Our APIs are the best way to automate your company's banking needs and are designed to be easy to understand and implement.</p><p>We're continuously growing this library and what you see here is just the start, but if you need something specific or have a question, <a class='text-blue-600' href='https://synctera.com/contact' target='_blank' rel='noreferrer'>contact us</a>.</p>

API version: 0.5.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package synctera

import (
	"encoding/json"
	"fmt"
)

// VendorInfo - The information provided to Synctera from the vendor. Interpretation of this object is up to the client.
type VendorInfo struct {
	VendorJson *VendorJson
	VendorXml  *VendorXml
}

// VendorJsonAsVendorInfo is a convenience function that returns VendorJson wrapped in VendorInfo
func VendorJsonAsVendorInfo(v *VendorJson) VendorInfo {
	return VendorInfo{VendorJson: v}
}

// VendorXmlAsVendorInfo is a convenience function that returns VendorXml wrapped in VendorInfo
func VendorXmlAsVendorInfo(v *VendorXml) VendorInfo {
	return VendorInfo{VendorXml: v}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *VendorInfo) UnmarshalJSON(data []byte) error {
	var err error
	// use discriminator value to speed up the lookup
	var jsonDict map[string]interface{}
	err = json.Unmarshal(data, &jsonDict)
	if err != nil {
		return fmt.Errorf("Failed to unmarshal JSON into map for the discriminator lookup.")
	}

	// check if the discriminator value is 'application/json'
	if jsonDict["content_type"] == "application/json" {
		// try to unmarshal JSON data into VendorJson
		err = json.Unmarshal(data, &dst.VendorJson)
		if err == nil {
			return nil // data stored in dst.VendorJson, return on the first match
		} else {
			dst.VendorJson = nil
			return fmt.Errorf("Failed to unmarshal VendorInfo as VendorJson: %s", err.Error())
		}
	}

	// check if the discriminator value is 'text/xml'
	if jsonDict["content_type"] == "text/xml" {
		// try to unmarshal JSON data into VendorXml
		err = json.Unmarshal(data, &dst.VendorXml)
		if err == nil {
			return nil // data stored in dst.VendorXml, return on the first match
		} else {
			dst.VendorXml = nil
			return fmt.Errorf("Failed to unmarshal VendorInfo as VendorXml: %s", err.Error())
		}
	}

	// check if the discriminator value is 'vendor_json'
	if jsonDict["content_type"] == "vendor_json" {
		// try to unmarshal JSON data into VendorJson
		err = json.Unmarshal(data, &dst.VendorJson)
		if err == nil {
			return nil // data stored in dst.VendorJson, return on the first match
		} else {
			dst.VendorJson = nil
			return fmt.Errorf("Failed to unmarshal VendorInfo as VendorJson: %s", err.Error())
		}
	}

	// check if the discriminator value is 'vendor_xml'
	if jsonDict["content_type"] == "vendor_xml" {
		// try to unmarshal JSON data into VendorXml
		err = json.Unmarshal(data, &dst.VendorXml)
		if err == nil {
			return nil // data stored in dst.VendorXml, return on the first match
		} else {
			dst.VendorXml = nil
			return fmt.Errorf("Failed to unmarshal VendorInfo as VendorXml: %s", err.Error())
		}
	}

	return nil
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src VendorInfo) MarshalJSON() ([]byte, error) {
	if src.VendorJson != nil {
		return json.Marshal(&src.VendorJson)
	}

	if src.VendorXml != nil {
		return json.Marshal(&src.VendorXml)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *VendorInfo) GetActualInstance() interface{} {
	if obj.VendorJson != nil {
		return obj.VendorJson
	}

	if obj.VendorXml != nil {
		return obj.VendorXml
	}

	// all schemas are nil
	return nil
}

type NullableVendorInfo struct {
	value *VendorInfo
	isSet bool
}

func (v NullableVendorInfo) Get() *VendorInfo {
	return v.value
}

func (v *NullableVendorInfo) Set(val *VendorInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableVendorInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableVendorInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVendorInfo(val *VendorInfo) *NullableVendorInfo {
	return &NullableVendorInfo{value: val, isSet: true}
}

func (v NullableVendorInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVendorInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
