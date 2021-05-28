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

// Reconciliation Reconciliation
type Reconciliation struct {
	// Reconciliation ID
	Id string `json:"id"`
	// Filename of the data to be reconciled
	FileName string `json:"file_name"`
	IngestionStatus IngestionStatus `json:"ingestion_status"`
}

// NewReconciliation instantiates a new Reconciliation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReconciliation(id string, fileName string, ingestionStatus IngestionStatus) *Reconciliation {
	this := Reconciliation{}
	this.Id = id
	this.FileName = fileName
	this.IngestionStatus = ingestionStatus
	return &this
}

// NewReconciliationWithDefaults instantiates a new Reconciliation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReconciliationWithDefaults() *Reconciliation {
	this := Reconciliation{}
	return &this
}

// GetId returns the Id field value
func (o *Reconciliation) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *Reconciliation) GetIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *Reconciliation) SetId(v string) {
	o.Id = v
}

// GetFileName returns the FileName field value
func (o *Reconciliation) GetFileName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FileName
}

// GetFileNameOk returns a tuple with the FileName field value
// and a boolean to check if the value has been set.
func (o *Reconciliation) GetFileNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.FileName, true
}

// SetFileName sets field value
func (o *Reconciliation) SetFileName(v string) {
	o.FileName = v
}

// GetIngestionStatus returns the IngestionStatus field value
func (o *Reconciliation) GetIngestionStatus() IngestionStatus {
	if o == nil {
		var ret IngestionStatus
		return ret
	}

	return o.IngestionStatus
}

// GetIngestionStatusOk returns a tuple with the IngestionStatus field value
// and a boolean to check if the value has been set.
func (o *Reconciliation) GetIngestionStatusOk() (*IngestionStatus, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.IngestionStatus, true
}

// SetIngestionStatus sets field value
func (o *Reconciliation) SetIngestionStatus(v IngestionStatus) {
	o.IngestionStatus = v
}

func (o Reconciliation) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["file_name"] = o.FileName
	}
	if true {
		toSerialize["ingestion_status"] = o.IngestionStatus
	}
	return json.Marshal(toSerialize)
}

type NullableReconciliation struct {
	value *Reconciliation
	isSet bool
}

func (v NullableReconciliation) Get() *Reconciliation {
	return v.value
}

func (v *NullableReconciliation) Set(val *Reconciliation) {
	v.value = val
	v.isSet = true
}

func (v NullableReconciliation) IsSet() bool {
	return v.isSet
}

func (v *NullableReconciliation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReconciliation(val *Reconciliation) *NullableReconciliation {
	return &NullableReconciliation{value: val, isSet: true}
}

func (v NullableReconciliation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReconciliation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


