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

// Application struct for Application
type Application struct {
	// Details about the applicant. The exact schema is to be determined with your bank.
	ApplicationDetails map[string]interface{} `json:"application_details"`
	// Customer ID for the application
	CustomerId string `json:"customer_id"`
	Status *ApplicationStatus `json:"status,omitempty"`
	Type ApplicationType `json:"type"`
}

// NewApplication instantiates a new Application object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApplication(applicationDetails map[string]interface{}, customerId string, type_ ApplicationType) *Application {
	this := Application{}
	this.ApplicationDetails = applicationDetails
	this.CustomerId = customerId
	this.Type = type_
	return &this
}

// NewApplicationWithDefaults instantiates a new Application object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApplicationWithDefaults() *Application {
	this := Application{}
	return &this
}

// GetApplicationDetails returns the ApplicationDetails field value
func (o *Application) GetApplicationDetails() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.ApplicationDetails
}

// GetApplicationDetailsOk returns a tuple with the ApplicationDetails field value
// and a boolean to check if the value has been set.
func (o *Application) GetApplicationDetailsOk() (map[string]interface{}, bool) {
	if o == nil {
		return nil, false
	}
	return o.ApplicationDetails, true
}

// SetApplicationDetails sets field value
func (o *Application) SetApplicationDetails(v map[string]interface{}) {
	o.ApplicationDetails = v
}

// GetCustomerId returns the CustomerId field value
func (o *Application) GetCustomerId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CustomerId
}

// GetCustomerIdOk returns a tuple with the CustomerId field value
// and a boolean to check if the value has been set.
func (o *Application) GetCustomerIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CustomerId, true
}

// SetCustomerId sets field value
func (o *Application) SetCustomerId(v string) {
	o.CustomerId = v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *Application) GetStatus() ApplicationStatus {
	if o == nil || o.Status == nil {
		var ret ApplicationStatus
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Application) GetStatusOk() (*ApplicationStatus, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *Application) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given ApplicationStatus and assigns it to the Status field.
func (o *Application) SetStatus(v ApplicationStatus) {
	o.Status = &v
}

// GetType returns the Type field value
func (o *Application) GetType() ApplicationType {
	if o == nil {
		var ret ApplicationType
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *Application) GetTypeOk() (*ApplicationType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *Application) SetType(v ApplicationType) {
	o.Type = v
}

func (o Application) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["application_details"] = o.ApplicationDetails
	}
	if true {
		toSerialize["customer_id"] = o.CustomerId
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	if true {
		toSerialize["type"] = o.Type
	}
	return json.Marshal(toSerialize)
}

type NullableApplication struct {
	value *Application
	isSet bool
}

func (v NullableApplication) Get() *Application {
	return v.value
}

func (v *NullableApplication) Set(val *Application) {
	v.value = val
	v.isSet = true
}

func (v NullableApplication) IsSet() bool {
	return v.isSet
}

func (v *NullableApplication) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApplication(val *Application) *NullableApplication {
	return &NullableApplication{value: val, isSet: true}
}

func (v NullableApplication) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApplication) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

