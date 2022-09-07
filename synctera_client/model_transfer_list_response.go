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

// TransferListResponse struct for TransferListResponse
type TransferListResponse struct {
	// Array of External transfer
	ExternalTransfers []TransferResponse `json:"external_transfers"`
	// If returned, use the next_page_token to query for the next page of results. Not returned if there are no more rows.
	NextPageToken *string `json:"next_page_token,omitempty"`
}

// NewTransferListResponse instantiates a new TransferListResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTransferListResponse(externalTransfers []TransferResponse) *TransferListResponse {
	this := TransferListResponse{}
	this.ExternalTransfers = externalTransfers
	return &this
}

// NewTransferListResponseWithDefaults instantiates a new TransferListResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTransferListResponseWithDefaults() *TransferListResponse {
	this := TransferListResponse{}
	return &this
}

// GetExternalTransfers returns the ExternalTransfers field value
func (o *TransferListResponse) GetExternalTransfers() []TransferResponse {
	if o == nil {
		var ret []TransferResponse
		return ret
	}

	return o.ExternalTransfers
}

// GetExternalTransfersOk returns a tuple with the ExternalTransfers field value
// and a boolean to check if the value has been set.
func (o *TransferListResponse) GetExternalTransfersOk() ([]TransferResponse, bool) {
	if o == nil {
		return nil, false
	}
	return o.ExternalTransfers, true
}

// SetExternalTransfers sets field value
func (o *TransferListResponse) SetExternalTransfers(v []TransferResponse) {
	o.ExternalTransfers = v
}

// GetNextPageToken returns the NextPageToken field value if set, zero value otherwise.
func (o *TransferListResponse) GetNextPageToken() string {
	if o == nil || o.NextPageToken == nil {
		var ret string
		return ret
	}
	return *o.NextPageToken
}

// GetNextPageTokenOk returns a tuple with the NextPageToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransferListResponse) GetNextPageTokenOk() (*string, bool) {
	if o == nil || o.NextPageToken == nil {
		return nil, false
	}
	return o.NextPageToken, true
}

// HasNextPageToken returns a boolean if a field has been set.
func (o *TransferListResponse) HasNextPageToken() bool {
	if o != nil && o.NextPageToken != nil {
		return true
	}

	return false
}

// SetNextPageToken gets a reference to the given string and assigns it to the NextPageToken field.
func (o *TransferListResponse) SetNextPageToken(v string) {
	o.NextPageToken = &v
}

func (o TransferListResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["external_transfers"] = o.ExternalTransfers
	}
	if o.NextPageToken != nil {
		toSerialize["next_page_token"] = o.NextPageToken
	}
	return json.Marshal(toSerialize)
}

type NullableTransferListResponse struct {
	value *TransferListResponse
	isSet bool
}

func (v NullableTransferListResponse) Get() *TransferListResponse {
	return v.value
}

func (v *NullableTransferListResponse) Set(val *TransferListResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableTransferListResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableTransferListResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransferListResponse(val *TransferListResponse) *NullableTransferListResponse {
	return &NullableTransferListResponse{value: val, isSet: true}
}

func (v NullableTransferListResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTransferListResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


