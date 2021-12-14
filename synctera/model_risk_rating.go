/*
Synctera API

<h2>Let's build something great.</h2><p>Welcome to the official reference documentation for Synctera APIs. Our APIs are the best way to automate your company's banking needs and are designed to be easy to understand and implement.</p><p>We're continuously growing this library and what you see here is just the start, but if you need something specific or have a question, <a class='text-blue-600' href='https://synctera.com/contact' target='_blank' rel='noreferrer'>contact us</a>.</p> 

API version: 0.5.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"time"
)

// RiskRating Represents a customer's risk profile
type RiskRating struct {
	// The risk configuration id used in risk score calculation
	ConfigurationId *string `json:"configuration_id,omitempty"`
	// Risk rating ID
	Id *string `json:"id,omitempty"`
	// The next review date where customer risk will be calculated
	NextReview *time.Time `json:"next_review,omitempty"`
	// A textual representation of the customer risk score
	RiskLevel *string `json:"risk_level,omitempty"`
	// The date the customer risk rating was calculated
	RiskReview *time.Time `json:"risk_review,omitempty"`
	// The cumulative score of all risk rating fields
	RiskScore *float32 `json:"risk_score,omitempty"`
}

// NewRiskRating instantiates a new RiskRating object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRiskRating() *RiskRating {
	this := RiskRating{}
	return &this
}

// NewRiskRatingWithDefaults instantiates a new RiskRating object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRiskRatingWithDefaults() *RiskRating {
	this := RiskRating{}
	return &this
}

// GetConfigurationId returns the ConfigurationId field value if set, zero value otherwise.
func (o *RiskRating) GetConfigurationId() string {
	if o == nil || o.ConfigurationId == nil {
		var ret string
		return ret
	}
	return *o.ConfigurationId
}

// GetConfigurationIdOk returns a tuple with the ConfigurationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RiskRating) GetConfigurationIdOk() (*string, bool) {
	if o == nil || o.ConfigurationId == nil {
		return nil, false
	}
	return o.ConfigurationId, true
}

// HasConfigurationId returns a boolean if a field has been set.
func (o *RiskRating) HasConfigurationId() bool {
	if o != nil && o.ConfigurationId != nil {
		return true
	}

	return false
}

// SetConfigurationId gets a reference to the given string and assigns it to the ConfigurationId field.
func (o *RiskRating) SetConfigurationId(v string) {
	o.ConfigurationId = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *RiskRating) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RiskRating) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *RiskRating) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *RiskRating) SetId(v string) {
	o.Id = &v
}

// GetNextReview returns the NextReview field value if set, zero value otherwise.
func (o *RiskRating) GetNextReview() time.Time {
	if o == nil || o.NextReview == nil {
		var ret time.Time
		return ret
	}
	return *o.NextReview
}

// GetNextReviewOk returns a tuple with the NextReview field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RiskRating) GetNextReviewOk() (*time.Time, bool) {
	if o == nil || o.NextReview == nil {
		return nil, false
	}
	return o.NextReview, true
}

// HasNextReview returns a boolean if a field has been set.
func (o *RiskRating) HasNextReview() bool {
	if o != nil && o.NextReview != nil {
		return true
	}

	return false
}

// SetNextReview gets a reference to the given time.Time and assigns it to the NextReview field.
func (o *RiskRating) SetNextReview(v time.Time) {
	o.NextReview = &v
}

// GetRiskLevel returns the RiskLevel field value if set, zero value otherwise.
func (o *RiskRating) GetRiskLevel() string {
	if o == nil || o.RiskLevel == nil {
		var ret string
		return ret
	}
	return *o.RiskLevel
}

// GetRiskLevelOk returns a tuple with the RiskLevel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RiskRating) GetRiskLevelOk() (*string, bool) {
	if o == nil || o.RiskLevel == nil {
		return nil, false
	}
	return o.RiskLevel, true
}

// HasRiskLevel returns a boolean if a field has been set.
func (o *RiskRating) HasRiskLevel() bool {
	if o != nil && o.RiskLevel != nil {
		return true
	}

	return false
}

// SetRiskLevel gets a reference to the given string and assigns it to the RiskLevel field.
func (o *RiskRating) SetRiskLevel(v string) {
	o.RiskLevel = &v
}

// GetRiskReview returns the RiskReview field value if set, zero value otherwise.
func (o *RiskRating) GetRiskReview() time.Time {
	if o == nil || o.RiskReview == nil {
		var ret time.Time
		return ret
	}
	return *o.RiskReview
}

// GetRiskReviewOk returns a tuple with the RiskReview field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RiskRating) GetRiskReviewOk() (*time.Time, bool) {
	if o == nil || o.RiskReview == nil {
		return nil, false
	}
	return o.RiskReview, true
}

// HasRiskReview returns a boolean if a field has been set.
func (o *RiskRating) HasRiskReview() bool {
	if o != nil && o.RiskReview != nil {
		return true
	}

	return false
}

// SetRiskReview gets a reference to the given time.Time and assigns it to the RiskReview field.
func (o *RiskRating) SetRiskReview(v time.Time) {
	o.RiskReview = &v
}

// GetRiskScore returns the RiskScore field value if set, zero value otherwise.
func (o *RiskRating) GetRiskScore() float32 {
	if o == nil || o.RiskScore == nil {
		var ret float32
		return ret
	}
	return *o.RiskScore
}

// GetRiskScoreOk returns a tuple with the RiskScore field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RiskRating) GetRiskScoreOk() (*float32, bool) {
	if o == nil || o.RiskScore == nil {
		return nil, false
	}
	return o.RiskScore, true
}

// HasRiskScore returns a boolean if a field has been set.
func (o *RiskRating) HasRiskScore() bool {
	if o != nil && o.RiskScore != nil {
		return true
	}

	return false
}

// SetRiskScore gets a reference to the given float32 and assigns it to the RiskScore field.
func (o *RiskRating) SetRiskScore(v float32) {
	o.RiskScore = &v
}

func (o RiskRating) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ConfigurationId != nil {
		toSerialize["configuration_id"] = o.ConfigurationId
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.NextReview != nil {
		toSerialize["next_review"] = o.NextReview
	}
	if o.RiskLevel != nil {
		toSerialize["risk_level"] = o.RiskLevel
	}
	if o.RiskReview != nil {
		toSerialize["risk_review"] = o.RiskReview
	}
	if o.RiskScore != nil {
		toSerialize["risk_score"] = o.RiskScore
	}
	return json.Marshal(toSerialize)
}

type NullableRiskRating struct {
	value *RiskRating
	isSet bool
}

func (v NullableRiskRating) Get() *RiskRating {
	return v.value
}

func (v *NullableRiskRating) Set(val *RiskRating) {
	v.value = val
	v.isSet = true
}

func (v NullableRiskRating) IsSet() bool {
	return v.isSet
}

func (v *NullableRiskRating) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRiskRating(val *RiskRating) *NullableRiskRating {
	return &NullableRiskRating{value: val, isSet: true}
}

func (v NullableRiskRating) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRiskRating) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


