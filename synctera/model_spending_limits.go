/*
Synctera API

<h2>Let's build something great.</h2><p>Welcome to the official reference documentation for Synctera APIs. Our APIs are the best way to automate your company's banking needs and are designed to be easy to understand and implement.</p><p>We're continuously growing this library and what you see here is just the start, but if you need something specific or have a question, <a class='text-blue-600' href='https://synctera.com/contact' target='_blank' rel='noreferrer'>contact us</a>.</p> 

API version: 0.5.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// SpendingLimits Account spending limits
type SpendingLimits struct {
	Day *SpendingLimitWithTime `json:"day,omitempty"`
	Lifetime *SpendingLimitWithTime `json:"lifetime,omitempty"`
	Month *SpendingLimitWithTime `json:"month,omitempty"`
	Transaction *SpendingLimitsTransaction `json:"transaction,omitempty"`
	Week *SpendingLimitWithTime `json:"week,omitempty"`
}

// NewSpendingLimits instantiates a new SpendingLimits object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSpendingLimits() *SpendingLimits {
	this := SpendingLimits{}
	return &this
}

// NewSpendingLimitsWithDefaults instantiates a new SpendingLimits object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSpendingLimitsWithDefaults() *SpendingLimits {
	this := SpendingLimits{}
	return &this
}

// GetDay returns the Day field value if set, zero value otherwise.
func (o *SpendingLimits) GetDay() SpendingLimitWithTime {
	if o == nil || o.Day == nil {
		var ret SpendingLimitWithTime
		return ret
	}
	return *o.Day
}

// GetDayOk returns a tuple with the Day field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpendingLimits) GetDayOk() (*SpendingLimitWithTime, bool) {
	if o == nil || o.Day == nil {
		return nil, false
	}
	return o.Day, true
}

// HasDay returns a boolean if a field has been set.
func (o *SpendingLimits) HasDay() bool {
	if o != nil && o.Day != nil {
		return true
	}

	return false
}

// SetDay gets a reference to the given SpendingLimitWithTime and assigns it to the Day field.
func (o *SpendingLimits) SetDay(v SpendingLimitWithTime) {
	o.Day = &v
}

// GetLifetime returns the Lifetime field value if set, zero value otherwise.
func (o *SpendingLimits) GetLifetime() SpendingLimitWithTime {
	if o == nil || o.Lifetime == nil {
		var ret SpendingLimitWithTime
		return ret
	}
	return *o.Lifetime
}

// GetLifetimeOk returns a tuple with the Lifetime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpendingLimits) GetLifetimeOk() (*SpendingLimitWithTime, bool) {
	if o == nil || o.Lifetime == nil {
		return nil, false
	}
	return o.Lifetime, true
}

// HasLifetime returns a boolean if a field has been set.
func (o *SpendingLimits) HasLifetime() bool {
	if o != nil && o.Lifetime != nil {
		return true
	}

	return false
}

// SetLifetime gets a reference to the given SpendingLimitWithTime and assigns it to the Lifetime field.
func (o *SpendingLimits) SetLifetime(v SpendingLimitWithTime) {
	o.Lifetime = &v
}

// GetMonth returns the Month field value if set, zero value otherwise.
func (o *SpendingLimits) GetMonth() SpendingLimitWithTime {
	if o == nil || o.Month == nil {
		var ret SpendingLimitWithTime
		return ret
	}
	return *o.Month
}

// GetMonthOk returns a tuple with the Month field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpendingLimits) GetMonthOk() (*SpendingLimitWithTime, bool) {
	if o == nil || o.Month == nil {
		return nil, false
	}
	return o.Month, true
}

// HasMonth returns a boolean if a field has been set.
func (o *SpendingLimits) HasMonth() bool {
	if o != nil && o.Month != nil {
		return true
	}

	return false
}

// SetMonth gets a reference to the given SpendingLimitWithTime and assigns it to the Month field.
func (o *SpendingLimits) SetMonth(v SpendingLimitWithTime) {
	o.Month = &v
}

// GetTransaction returns the Transaction field value if set, zero value otherwise.
func (o *SpendingLimits) GetTransaction() SpendingLimitsTransaction {
	if o == nil || o.Transaction == nil {
		var ret SpendingLimitsTransaction
		return ret
	}
	return *o.Transaction
}

// GetTransactionOk returns a tuple with the Transaction field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpendingLimits) GetTransactionOk() (*SpendingLimitsTransaction, bool) {
	if o == nil || o.Transaction == nil {
		return nil, false
	}
	return o.Transaction, true
}

// HasTransaction returns a boolean if a field has been set.
func (o *SpendingLimits) HasTransaction() bool {
	if o != nil && o.Transaction != nil {
		return true
	}

	return false
}

// SetTransaction gets a reference to the given SpendingLimitsTransaction and assigns it to the Transaction field.
func (o *SpendingLimits) SetTransaction(v SpendingLimitsTransaction) {
	o.Transaction = &v
}

// GetWeek returns the Week field value if set, zero value otherwise.
func (o *SpendingLimits) GetWeek() SpendingLimitWithTime {
	if o == nil || o.Week == nil {
		var ret SpendingLimitWithTime
		return ret
	}
	return *o.Week
}

// GetWeekOk returns a tuple with the Week field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpendingLimits) GetWeekOk() (*SpendingLimitWithTime, bool) {
	if o == nil || o.Week == nil {
		return nil, false
	}
	return o.Week, true
}

// HasWeek returns a boolean if a field has been set.
func (o *SpendingLimits) HasWeek() bool {
	if o != nil && o.Week != nil {
		return true
	}

	return false
}

// SetWeek gets a reference to the given SpendingLimitWithTime and assigns it to the Week field.
func (o *SpendingLimits) SetWeek(v SpendingLimitWithTime) {
	o.Week = &v
}

func (o SpendingLimits) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Day != nil {
		toSerialize["day"] = o.Day
	}
	if o.Lifetime != nil {
		toSerialize["lifetime"] = o.Lifetime
	}
	if o.Month != nil {
		toSerialize["month"] = o.Month
	}
	if o.Transaction != nil {
		toSerialize["transaction"] = o.Transaction
	}
	if o.Week != nil {
		toSerialize["week"] = o.Week
	}
	return json.Marshal(toSerialize)
}

type NullableSpendingLimits struct {
	value *SpendingLimits
	isSet bool
}

func (v NullableSpendingLimits) Get() *SpendingLimits {
	return v.value
}

func (v *NullableSpendingLimits) Set(val *SpendingLimits) {
	v.value = val
	v.isSet = true
}

func (v NullableSpendingLimits) IsSet() bool {
	return v.isSet
}

func (v *NullableSpendingLimits) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSpendingLimits(val *SpendingLimits) *NullableSpendingLimits {
	return &NullableSpendingLimits{value: val, isSet: true}
}

func (v NullableSpendingLimits) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSpendingLimits) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


