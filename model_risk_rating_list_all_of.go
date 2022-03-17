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

// RiskRatingListAllOf struct for RiskRatingListAllOf
type RiskRatingListAllOf struct {
	// Array of customer risk ratings
	RiskRatings []RiskRating `json:"risk_ratings"`
}

// NewRiskRatingListAllOf instantiates a new RiskRatingListAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRiskRatingListAllOf(riskRatings []RiskRating) *RiskRatingListAllOf {
	this := RiskRatingListAllOf{}
	this.RiskRatings = riskRatings
	return &this
}

// NewRiskRatingListAllOfWithDefaults instantiates a new RiskRatingListAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRiskRatingListAllOfWithDefaults() *RiskRatingListAllOf {
	this := RiskRatingListAllOf{}
	return &this
}

// GetRiskRatings returns the RiskRatings field value
func (o *RiskRatingListAllOf) GetRiskRatings() []RiskRating {
	if o == nil {
		var ret []RiskRating
		return ret
	}

	return o.RiskRatings
}

// GetRiskRatingsOk returns a tuple with the RiskRatings field value
// and a boolean to check if the value has been set.
func (o *RiskRatingListAllOf) GetRiskRatingsOk() (*[]RiskRating, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RiskRatings, true
}

// SetRiskRatings sets field value
func (o *RiskRatingListAllOf) SetRiskRatings(v []RiskRating) {
	o.RiskRatings = v
}

func (o RiskRatingListAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["risk_ratings"] = o.RiskRatings
	}
	return json.Marshal(toSerialize)
}

type NullableRiskRatingListAllOf struct {
	value *RiskRatingListAllOf
	isSet bool
}

func (v NullableRiskRatingListAllOf) Get() *RiskRatingListAllOf {
	return v.value
}

func (v *NullableRiskRatingListAllOf) Set(val *RiskRatingListAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableRiskRatingListAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableRiskRatingListAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRiskRatingListAllOf(val *RiskRatingListAllOf) *NullableRiskRatingListAllOf {
	return &NullableRiskRatingListAllOf{value: val, isSet: true}
}

func (v NullableRiskRatingListAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRiskRatingListAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
