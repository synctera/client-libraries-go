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

// PhysicalCardResponseStatus struct for PhysicalCardResponseStatus
type PhysicalCardResponseStatus struct {
	CardFulfillmentStatus CardFulfillmentStatus `json:"card_fulfillment_status"`
	CardStatus CardStatus `json:"card_status"`
	// The carrier with whom the card is shipped
	Carrier *string `json:"carrier,omitempty"`
	// Additional details about the reason for the status change
	Memo *string `json:"memo,omitempty"`
	// The status of indicating the shipping status of the card
	ShippingStatus *string `json:"shipping_status,omitempty"`
	StatusReason CardStatusReasonCode `json:"status_reason"`
	// The tracking number
	TrackingNumber *string `json:"tracking_number,omitempty"`
}

// NewPhysicalCardResponseStatus instantiates a new PhysicalCardResponseStatus object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPhysicalCardResponseStatus(cardFulfillmentStatus CardFulfillmentStatus, cardStatus CardStatus, statusReason CardStatusReasonCode) *PhysicalCardResponseStatus {
	this := PhysicalCardResponseStatus{}
	this.CardFulfillmentStatus = cardFulfillmentStatus
	this.CardStatus = cardStatus
	this.StatusReason = statusReason
	return &this
}

// NewPhysicalCardResponseStatusWithDefaults instantiates a new PhysicalCardResponseStatus object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPhysicalCardResponseStatusWithDefaults() *PhysicalCardResponseStatus {
	this := PhysicalCardResponseStatus{}
	return &this
}

// GetCardFulfillmentStatus returns the CardFulfillmentStatus field value
func (o *PhysicalCardResponseStatus) GetCardFulfillmentStatus() CardFulfillmentStatus {
	if o == nil {
		var ret CardFulfillmentStatus
		return ret
	}

	return o.CardFulfillmentStatus
}

// GetCardFulfillmentStatusOk returns a tuple with the CardFulfillmentStatus field value
// and a boolean to check if the value has been set.
func (o *PhysicalCardResponseStatus) GetCardFulfillmentStatusOk() (*CardFulfillmentStatus, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.CardFulfillmentStatus, true
}

// SetCardFulfillmentStatus sets field value
func (o *PhysicalCardResponseStatus) SetCardFulfillmentStatus(v CardFulfillmentStatus) {
	o.CardFulfillmentStatus = v
}

// GetCardStatus returns the CardStatus field value
func (o *PhysicalCardResponseStatus) GetCardStatus() CardStatus {
	if o == nil {
		var ret CardStatus
		return ret
	}

	return o.CardStatus
}

// GetCardStatusOk returns a tuple with the CardStatus field value
// and a boolean to check if the value has been set.
func (o *PhysicalCardResponseStatus) GetCardStatusOk() (*CardStatus, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.CardStatus, true
}

// SetCardStatus sets field value
func (o *PhysicalCardResponseStatus) SetCardStatus(v CardStatus) {
	o.CardStatus = v
}

// GetCarrier returns the Carrier field value if set, zero value otherwise.
func (o *PhysicalCardResponseStatus) GetCarrier() string {
	if o == nil || o.Carrier == nil {
		var ret string
		return ret
	}
	return *o.Carrier
}

// GetCarrierOk returns a tuple with the Carrier field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PhysicalCardResponseStatus) GetCarrierOk() (*string, bool) {
	if o == nil || o.Carrier == nil {
		return nil, false
	}
	return o.Carrier, true
}

// HasCarrier returns a boolean if a field has been set.
func (o *PhysicalCardResponseStatus) HasCarrier() bool {
	if o != nil && o.Carrier != nil {
		return true
	}

	return false
}

// SetCarrier gets a reference to the given string and assigns it to the Carrier field.
func (o *PhysicalCardResponseStatus) SetCarrier(v string) {
	o.Carrier = &v
}

// GetMemo returns the Memo field value if set, zero value otherwise.
func (o *PhysicalCardResponseStatus) GetMemo() string {
	if o == nil || o.Memo == nil {
		var ret string
		return ret
	}
	return *o.Memo
}

// GetMemoOk returns a tuple with the Memo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PhysicalCardResponseStatus) GetMemoOk() (*string, bool) {
	if o == nil || o.Memo == nil {
		return nil, false
	}
	return o.Memo, true
}

// HasMemo returns a boolean if a field has been set.
func (o *PhysicalCardResponseStatus) HasMemo() bool {
	if o != nil && o.Memo != nil {
		return true
	}

	return false
}

// SetMemo gets a reference to the given string and assigns it to the Memo field.
func (o *PhysicalCardResponseStatus) SetMemo(v string) {
	o.Memo = &v
}

// GetShippingStatus returns the ShippingStatus field value if set, zero value otherwise.
func (o *PhysicalCardResponseStatus) GetShippingStatus() string {
	if o == nil || o.ShippingStatus == nil {
		var ret string
		return ret
	}
	return *o.ShippingStatus
}

// GetShippingStatusOk returns a tuple with the ShippingStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PhysicalCardResponseStatus) GetShippingStatusOk() (*string, bool) {
	if o == nil || o.ShippingStatus == nil {
		return nil, false
	}
	return o.ShippingStatus, true
}

// HasShippingStatus returns a boolean if a field has been set.
func (o *PhysicalCardResponseStatus) HasShippingStatus() bool {
	if o != nil && o.ShippingStatus != nil {
		return true
	}

	return false
}

// SetShippingStatus gets a reference to the given string and assigns it to the ShippingStatus field.
func (o *PhysicalCardResponseStatus) SetShippingStatus(v string) {
	o.ShippingStatus = &v
}

// GetStatusReason returns the StatusReason field value
func (o *PhysicalCardResponseStatus) GetStatusReason() CardStatusReasonCode {
	if o == nil {
		var ret CardStatusReasonCode
		return ret
	}

	return o.StatusReason
}

// GetStatusReasonOk returns a tuple with the StatusReason field value
// and a boolean to check if the value has been set.
func (o *PhysicalCardResponseStatus) GetStatusReasonOk() (*CardStatusReasonCode, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.StatusReason, true
}

// SetStatusReason sets field value
func (o *PhysicalCardResponseStatus) SetStatusReason(v CardStatusReasonCode) {
	o.StatusReason = v
}

// GetTrackingNumber returns the TrackingNumber field value if set, zero value otherwise.
func (o *PhysicalCardResponseStatus) GetTrackingNumber() string {
	if o == nil || o.TrackingNumber == nil {
		var ret string
		return ret
	}
	return *o.TrackingNumber
}

// GetTrackingNumberOk returns a tuple with the TrackingNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PhysicalCardResponseStatus) GetTrackingNumberOk() (*string, bool) {
	if o == nil || o.TrackingNumber == nil {
		return nil, false
	}
	return o.TrackingNumber, true
}

// HasTrackingNumber returns a boolean if a field has been set.
func (o *PhysicalCardResponseStatus) HasTrackingNumber() bool {
	if o != nil && o.TrackingNumber != nil {
		return true
	}

	return false
}

// SetTrackingNumber gets a reference to the given string and assigns it to the TrackingNumber field.
func (o *PhysicalCardResponseStatus) SetTrackingNumber(v string) {
	o.TrackingNumber = &v
}

func (o PhysicalCardResponseStatus) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["card_fulfillment_status"] = o.CardFulfillmentStatus
	}
	if true {
		toSerialize["card_status"] = o.CardStatus
	}
	if o.Carrier != nil {
		toSerialize["carrier"] = o.Carrier
	}
	if o.Memo != nil {
		toSerialize["memo"] = o.Memo
	}
	if o.ShippingStatus != nil {
		toSerialize["shipping_status"] = o.ShippingStatus
	}
	if true {
		toSerialize["status_reason"] = o.StatusReason
	}
	if o.TrackingNumber != nil {
		toSerialize["tracking_number"] = o.TrackingNumber
	}
	return json.Marshal(toSerialize)
}

type NullablePhysicalCardResponseStatus struct {
	value *PhysicalCardResponseStatus
	isSet bool
}

func (v NullablePhysicalCardResponseStatus) Get() *PhysicalCardResponseStatus {
	return v.value
}

func (v *NullablePhysicalCardResponseStatus) Set(val *PhysicalCardResponseStatus) {
	v.value = val
	v.isSet = true
}

func (v NullablePhysicalCardResponseStatus) IsSet() bool {
	return v.isSet
}

func (v *NullablePhysicalCardResponseStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePhysicalCardResponseStatus(val *PhysicalCardResponseStatus) *NullablePhysicalCardResponseStatus {
	return &NullablePhysicalCardResponseStatus{value: val, isSet: true}
}

func (v NullablePhysicalCardResponseStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePhysicalCardResponseStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


