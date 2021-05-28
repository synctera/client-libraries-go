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

// ScanListAllOf struct for ScanListAllOf
type ScanListAllOf struct {
	// Array of RDC scans
	Scans []Scan `json:"scans"`
}

// NewScanListAllOf instantiates a new ScanListAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewScanListAllOf(scans []Scan) *ScanListAllOf {
	this := ScanListAllOf{}
	this.Scans = scans
	return &this
}

// NewScanListAllOfWithDefaults instantiates a new ScanListAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewScanListAllOfWithDefaults() *ScanListAllOf {
	this := ScanListAllOf{}
	return &this
}

// GetScans returns the Scans field value
func (o *ScanListAllOf) GetScans() []Scan {
	if o == nil {
		var ret []Scan
		return ret
	}

	return o.Scans
}

// GetScansOk returns a tuple with the Scans field value
// and a boolean to check if the value has been set.
func (o *ScanListAllOf) GetScansOk() (*[]Scan, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Scans, true
}

// SetScans sets field value
func (o *ScanListAllOf) SetScans(v []Scan) {
	o.Scans = v
}

func (o ScanListAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["scans"] = o.Scans
	}
	return json.Marshal(toSerialize)
}

type NullableScanListAllOf struct {
	value *ScanListAllOf
	isSet bool
}

func (v NullableScanListAllOf) Get() *ScanListAllOf {
	return v.value
}

func (v *NullableScanListAllOf) Set(val *ScanListAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableScanListAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableScanListAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableScanListAllOf(val *ScanListAllOf) *NullableScanListAllOf {
	return &NullableScanListAllOf{value: val, isSet: true}
}

func (v NullableScanListAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableScanListAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


