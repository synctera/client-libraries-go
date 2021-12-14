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

// Employment A period of time in which a customer is (was) employed by a particular employer. 
type Employment struct {
	// Name of customer's employer.
	EmployerName string `json:"employer_name"`
	// First day of employment.
	EmploymentFrom *time.Time `json:"employment_from,omitempty"`
	// Number of hours spent per week working for specified employment.
	EmploymentHours *float32 `json:"employment_hours,omitempty"`
	// Annual income in cents.
	EmploymentIncome *int32 `json:"employment_income,omitempty"`
	// The 3-letter alphabetic ISO 4217 code for the currency in which the employee was paid. 
	EmploymentIncomeCurrency *string `json:"employment_income_currency,omitempty"`
	// A collection of arbitrary key-value pairs providing additional information about this employment relationship. 
	EmploymentInfo *map[string]interface{} `json:"employment_info,omitempty"`
	// Customer's work title, profession, or field.
	EmploymentOccupation *string `json:"employment_occupation,omitempty"`
	// Last day of employment.
	EmploymentTo *time.Time `json:"employment_to,omitempty"`
	// Unique ID for this employment relationship.
	Id *string `json:"id,omitempty"`
}

// NewEmployment instantiates a new Employment object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEmployment(employerName string) *Employment {
	this := Employment{}
	this.EmployerName = employerName
	return &this
}

// NewEmploymentWithDefaults instantiates a new Employment object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEmploymentWithDefaults() *Employment {
	this := Employment{}
	return &this
}

// GetEmployerName returns the EmployerName field value
func (o *Employment) GetEmployerName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EmployerName
}

// GetEmployerNameOk returns a tuple with the EmployerName field value
// and a boolean to check if the value has been set.
func (o *Employment) GetEmployerNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.EmployerName, true
}

// SetEmployerName sets field value
func (o *Employment) SetEmployerName(v string) {
	o.EmployerName = v
}

// GetEmploymentFrom returns the EmploymentFrom field value if set, zero value otherwise.
func (o *Employment) GetEmploymentFrom() time.Time {
	if o == nil || o.EmploymentFrom == nil {
		var ret time.Time
		return ret
	}
	return *o.EmploymentFrom
}

// GetEmploymentFromOk returns a tuple with the EmploymentFrom field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Employment) GetEmploymentFromOk() (*time.Time, bool) {
	if o == nil || o.EmploymentFrom == nil {
		return nil, false
	}
	return o.EmploymentFrom, true
}

// HasEmploymentFrom returns a boolean if a field has been set.
func (o *Employment) HasEmploymentFrom() bool {
	if o != nil && o.EmploymentFrom != nil {
		return true
	}

	return false
}

// SetEmploymentFrom gets a reference to the given time.Time and assigns it to the EmploymentFrom field.
func (o *Employment) SetEmploymentFrom(v time.Time) {
	o.EmploymentFrom = &v
}

// GetEmploymentHours returns the EmploymentHours field value if set, zero value otherwise.
func (o *Employment) GetEmploymentHours() float32 {
	if o == nil || o.EmploymentHours == nil {
		var ret float32
		return ret
	}
	return *o.EmploymentHours
}

// GetEmploymentHoursOk returns a tuple with the EmploymentHours field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Employment) GetEmploymentHoursOk() (*float32, bool) {
	if o == nil || o.EmploymentHours == nil {
		return nil, false
	}
	return o.EmploymentHours, true
}

// HasEmploymentHours returns a boolean if a field has been set.
func (o *Employment) HasEmploymentHours() bool {
	if o != nil && o.EmploymentHours != nil {
		return true
	}

	return false
}

// SetEmploymentHours gets a reference to the given float32 and assigns it to the EmploymentHours field.
func (o *Employment) SetEmploymentHours(v float32) {
	o.EmploymentHours = &v
}

// GetEmploymentIncome returns the EmploymentIncome field value if set, zero value otherwise.
func (o *Employment) GetEmploymentIncome() int32 {
	if o == nil || o.EmploymentIncome == nil {
		var ret int32
		return ret
	}
	return *o.EmploymentIncome
}

// GetEmploymentIncomeOk returns a tuple with the EmploymentIncome field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Employment) GetEmploymentIncomeOk() (*int32, bool) {
	if o == nil || o.EmploymentIncome == nil {
		return nil, false
	}
	return o.EmploymentIncome, true
}

// HasEmploymentIncome returns a boolean if a field has been set.
func (o *Employment) HasEmploymentIncome() bool {
	if o != nil && o.EmploymentIncome != nil {
		return true
	}

	return false
}

// SetEmploymentIncome gets a reference to the given int32 and assigns it to the EmploymentIncome field.
func (o *Employment) SetEmploymentIncome(v int32) {
	o.EmploymentIncome = &v
}

// GetEmploymentIncomeCurrency returns the EmploymentIncomeCurrency field value if set, zero value otherwise.
func (o *Employment) GetEmploymentIncomeCurrency() string {
	if o == nil || o.EmploymentIncomeCurrency == nil {
		var ret string
		return ret
	}
	return *o.EmploymentIncomeCurrency
}

// GetEmploymentIncomeCurrencyOk returns a tuple with the EmploymentIncomeCurrency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Employment) GetEmploymentIncomeCurrencyOk() (*string, bool) {
	if o == nil || o.EmploymentIncomeCurrency == nil {
		return nil, false
	}
	return o.EmploymentIncomeCurrency, true
}

// HasEmploymentIncomeCurrency returns a boolean if a field has been set.
func (o *Employment) HasEmploymentIncomeCurrency() bool {
	if o != nil && o.EmploymentIncomeCurrency != nil {
		return true
	}

	return false
}

// SetEmploymentIncomeCurrency gets a reference to the given string and assigns it to the EmploymentIncomeCurrency field.
func (o *Employment) SetEmploymentIncomeCurrency(v string) {
	o.EmploymentIncomeCurrency = &v
}

// GetEmploymentInfo returns the EmploymentInfo field value if set, zero value otherwise.
func (o *Employment) GetEmploymentInfo() map[string]interface{} {
	if o == nil || o.EmploymentInfo == nil {
		var ret map[string]interface{}
		return ret
	}
	return *o.EmploymentInfo
}

// GetEmploymentInfoOk returns a tuple with the EmploymentInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Employment) GetEmploymentInfoOk() (*map[string]interface{}, bool) {
	if o == nil || o.EmploymentInfo == nil {
		return nil, false
	}
	return o.EmploymentInfo, true
}

// HasEmploymentInfo returns a boolean if a field has been set.
func (o *Employment) HasEmploymentInfo() bool {
	if o != nil && o.EmploymentInfo != nil {
		return true
	}

	return false
}

// SetEmploymentInfo gets a reference to the given map[string]interface{} and assigns it to the EmploymentInfo field.
func (o *Employment) SetEmploymentInfo(v map[string]interface{}) {
	o.EmploymentInfo = &v
}

// GetEmploymentOccupation returns the EmploymentOccupation field value if set, zero value otherwise.
func (o *Employment) GetEmploymentOccupation() string {
	if o == nil || o.EmploymentOccupation == nil {
		var ret string
		return ret
	}
	return *o.EmploymentOccupation
}

// GetEmploymentOccupationOk returns a tuple with the EmploymentOccupation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Employment) GetEmploymentOccupationOk() (*string, bool) {
	if o == nil || o.EmploymentOccupation == nil {
		return nil, false
	}
	return o.EmploymentOccupation, true
}

// HasEmploymentOccupation returns a boolean if a field has been set.
func (o *Employment) HasEmploymentOccupation() bool {
	if o != nil && o.EmploymentOccupation != nil {
		return true
	}

	return false
}

// SetEmploymentOccupation gets a reference to the given string and assigns it to the EmploymentOccupation field.
func (o *Employment) SetEmploymentOccupation(v string) {
	o.EmploymentOccupation = &v
}

// GetEmploymentTo returns the EmploymentTo field value if set, zero value otherwise.
func (o *Employment) GetEmploymentTo() time.Time {
	if o == nil || o.EmploymentTo == nil {
		var ret time.Time
		return ret
	}
	return *o.EmploymentTo
}

// GetEmploymentToOk returns a tuple with the EmploymentTo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Employment) GetEmploymentToOk() (*time.Time, bool) {
	if o == nil || o.EmploymentTo == nil {
		return nil, false
	}
	return o.EmploymentTo, true
}

// HasEmploymentTo returns a boolean if a field has been set.
func (o *Employment) HasEmploymentTo() bool {
	if o != nil && o.EmploymentTo != nil {
		return true
	}

	return false
}

// SetEmploymentTo gets a reference to the given time.Time and assigns it to the EmploymentTo field.
func (o *Employment) SetEmploymentTo(v time.Time) {
	o.EmploymentTo = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Employment) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Employment) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Employment) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *Employment) SetId(v string) {
	o.Id = &v
}

func (o Employment) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["employer_name"] = o.EmployerName
	}
	if o.EmploymentFrom != nil {
		toSerialize["employment_from"] = o.EmploymentFrom
	}
	if o.EmploymentHours != nil {
		toSerialize["employment_hours"] = o.EmploymentHours
	}
	if o.EmploymentIncome != nil {
		toSerialize["employment_income"] = o.EmploymentIncome
	}
	if o.EmploymentIncomeCurrency != nil {
		toSerialize["employment_income_currency"] = o.EmploymentIncomeCurrency
	}
	if o.EmploymentInfo != nil {
		toSerialize["employment_info"] = o.EmploymentInfo
	}
	if o.EmploymentOccupation != nil {
		toSerialize["employment_occupation"] = o.EmploymentOccupation
	}
	if o.EmploymentTo != nil {
		toSerialize["employment_to"] = o.EmploymentTo
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	return json.Marshal(toSerialize)
}

type NullableEmployment struct {
	value *Employment
	isSet bool
}

func (v NullableEmployment) Get() *Employment {
	return v.value
}

func (v *NullableEmployment) Set(val *Employment) {
	v.value = val
	v.isSet = true
}

func (v NullableEmployment) IsSet() bool {
	return v.isSet
}

func (v *NullableEmployment) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEmployment(val *Employment) *NullableEmployment {
	return &NullableEmployment{value: val, isSet: true}
}

func (v NullableEmployment) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEmployment) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


