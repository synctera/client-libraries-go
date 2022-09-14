/*
Synctera API

<h2>Let's build something great.</h2><p>Welcome to the official reference documentation for Synctera APIs. Our APIs are the best way to automate your company's banking needs and are designed to be easy to understand and implement.</p><p>We're continuously growing this library and what you see here is just the start, but if you need something specific or have a question, <a class='text-blue-600' href='https://synctera.com/contact' target='_blank' rel='noreferrer'>contact us</a>.</p>

API version: 0.20.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package synctera_client

import (
	"encoding/json"
	"time"
)

// Prospect1AllOf struct for Prospect1AllOf
type Prospect1AllOf struct {
	// Creation time
	CreationTime *time.Time `json:"creation_time,omitempty"`
	// Email
	Email           string  `json:"email"`
	EmailValidation *string `json:"email_validation,omitempty"`
	// Prospect ID
	Id *string `json:"id,omitempty"`
	// Most recent updated time
	LastUpdatedTime *time.Time `json:"last_updated_time,omitempty"`
	// Number of referral points
	Points *int32 `json:"points,omitempty"`
	// A token used to verify the user is not a robot
	RecaptchaToken *string `json:"recaptcha_token,omitempty"`
	// Referral code
	ReferralCode *string `json:"referral_code,omitempty"`
	// UUID of the referrer
	ReferredBy *string `json:"referred_by,omitempty"`
	// Referral code of the referrer
	ReferredByCode *string `json:"referred_by_code,omitempty"`
	// Number of people this prospect referred
	ReferredProspects *int32      `json:"referred_prospects,omitempty"`
	VendorInfo        *VendorInfo `json:"vendor_info,omitempty"`
	// Verification token sent to prospect
	VerificationToken *string `json:"verification_token,omitempty"`
	// Waitlist ID
	WaitlistId *string `json:"waitlist_id,omitempty"`
	// Prospect's number in the waitlist
	WaitlistPosition NullableInt32 `json:"waitlist_position,omitempty"`
}

// NewProspect1AllOf instantiates a new Prospect1AllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProspect1AllOf(email string) *Prospect1AllOf {
	this := Prospect1AllOf{}
	this.Email = email
	return &this
}

// NewProspect1AllOfWithDefaults instantiates a new Prospect1AllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProspect1AllOfWithDefaults() *Prospect1AllOf {
	this := Prospect1AllOf{}
	return &this
}

// GetCreationTime returns the CreationTime field value if set, zero value otherwise.
func (o *Prospect1AllOf) GetCreationTime() time.Time {
	if o == nil || o.CreationTime == nil {
		var ret time.Time
		return ret
	}
	return *o.CreationTime
}

// GetCreationTimeOk returns a tuple with the CreationTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Prospect1AllOf) GetCreationTimeOk() (*time.Time, bool) {
	if o == nil || o.CreationTime == nil {
		return nil, false
	}
	return o.CreationTime, true
}

// HasCreationTime returns a boolean if a field has been set.
func (o *Prospect1AllOf) HasCreationTime() bool {
	if o != nil && o.CreationTime != nil {
		return true
	}

	return false
}

// SetCreationTime gets a reference to the given time.Time and assigns it to the CreationTime field.
func (o *Prospect1AllOf) SetCreationTime(v time.Time) {
	o.CreationTime = &v
}

// GetEmail returns the Email field value
func (o *Prospect1AllOf) GetEmail() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Email
}

// GetEmailOk returns a tuple with the Email field value
// and a boolean to check if the value has been set.
func (o *Prospect1AllOf) GetEmailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Email, true
}

// SetEmail sets field value
func (o *Prospect1AllOf) SetEmail(v string) {
	o.Email = v
}

// GetEmailValidation returns the EmailValidation field value if set, zero value otherwise.
func (o *Prospect1AllOf) GetEmailValidation() string {
	if o == nil || o.EmailValidation == nil {
		var ret string
		return ret
	}
	return *o.EmailValidation
}

// GetEmailValidationOk returns a tuple with the EmailValidation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Prospect1AllOf) GetEmailValidationOk() (*string, bool) {
	if o == nil || o.EmailValidation == nil {
		return nil, false
	}
	return o.EmailValidation, true
}

// HasEmailValidation returns a boolean if a field has been set.
func (o *Prospect1AllOf) HasEmailValidation() bool {
	if o != nil && o.EmailValidation != nil {
		return true
	}

	return false
}

// SetEmailValidation gets a reference to the given string and assigns it to the EmailValidation field.
func (o *Prospect1AllOf) SetEmailValidation(v string) {
	o.EmailValidation = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Prospect1AllOf) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Prospect1AllOf) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Prospect1AllOf) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *Prospect1AllOf) SetId(v string) {
	o.Id = &v
}

// GetLastUpdatedTime returns the LastUpdatedTime field value if set, zero value otherwise.
func (o *Prospect1AllOf) GetLastUpdatedTime() time.Time {
	if o == nil || o.LastUpdatedTime == nil {
		var ret time.Time
		return ret
	}
	return *o.LastUpdatedTime
}

// GetLastUpdatedTimeOk returns a tuple with the LastUpdatedTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Prospect1AllOf) GetLastUpdatedTimeOk() (*time.Time, bool) {
	if o == nil || o.LastUpdatedTime == nil {
		return nil, false
	}
	return o.LastUpdatedTime, true
}

// HasLastUpdatedTime returns a boolean if a field has been set.
func (o *Prospect1AllOf) HasLastUpdatedTime() bool {
	if o != nil && o.LastUpdatedTime != nil {
		return true
	}

	return false
}

// SetLastUpdatedTime gets a reference to the given time.Time and assigns it to the LastUpdatedTime field.
func (o *Prospect1AllOf) SetLastUpdatedTime(v time.Time) {
	o.LastUpdatedTime = &v
}

// GetPoints returns the Points field value if set, zero value otherwise.
func (o *Prospect1AllOf) GetPoints() int32 {
	if o == nil || o.Points == nil {
		var ret int32
		return ret
	}
	return *o.Points
}

// GetPointsOk returns a tuple with the Points field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Prospect1AllOf) GetPointsOk() (*int32, bool) {
	if o == nil || o.Points == nil {
		return nil, false
	}
	return o.Points, true
}

// HasPoints returns a boolean if a field has been set.
func (o *Prospect1AllOf) HasPoints() bool {
	if o != nil && o.Points != nil {
		return true
	}

	return false
}

// SetPoints gets a reference to the given int32 and assigns it to the Points field.
func (o *Prospect1AllOf) SetPoints(v int32) {
	o.Points = &v
}

// GetRecaptchaToken returns the RecaptchaToken field value if set, zero value otherwise.
func (o *Prospect1AllOf) GetRecaptchaToken() string {
	if o == nil || o.RecaptchaToken == nil {
		var ret string
		return ret
	}
	return *o.RecaptchaToken
}

// GetRecaptchaTokenOk returns a tuple with the RecaptchaToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Prospect1AllOf) GetRecaptchaTokenOk() (*string, bool) {
	if o == nil || o.RecaptchaToken == nil {
		return nil, false
	}
	return o.RecaptchaToken, true
}

// HasRecaptchaToken returns a boolean if a field has been set.
func (o *Prospect1AllOf) HasRecaptchaToken() bool {
	if o != nil && o.RecaptchaToken != nil {
		return true
	}

	return false
}

// SetRecaptchaToken gets a reference to the given string and assigns it to the RecaptchaToken field.
func (o *Prospect1AllOf) SetRecaptchaToken(v string) {
	o.RecaptchaToken = &v
}

// GetReferralCode returns the ReferralCode field value if set, zero value otherwise.
func (o *Prospect1AllOf) GetReferralCode() string {
	if o == nil || o.ReferralCode == nil {
		var ret string
		return ret
	}
	return *o.ReferralCode
}

// GetReferralCodeOk returns a tuple with the ReferralCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Prospect1AllOf) GetReferralCodeOk() (*string, bool) {
	if o == nil || o.ReferralCode == nil {
		return nil, false
	}
	return o.ReferralCode, true
}

// HasReferralCode returns a boolean if a field has been set.
func (o *Prospect1AllOf) HasReferralCode() bool {
	if o != nil && o.ReferralCode != nil {
		return true
	}

	return false
}

// SetReferralCode gets a reference to the given string and assigns it to the ReferralCode field.
func (o *Prospect1AllOf) SetReferralCode(v string) {
	o.ReferralCode = &v
}

// GetReferredBy returns the ReferredBy field value if set, zero value otherwise.
func (o *Prospect1AllOf) GetReferredBy() string {
	if o == nil || o.ReferredBy == nil {
		var ret string
		return ret
	}
	return *o.ReferredBy
}

// GetReferredByOk returns a tuple with the ReferredBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Prospect1AllOf) GetReferredByOk() (*string, bool) {
	if o == nil || o.ReferredBy == nil {
		return nil, false
	}
	return o.ReferredBy, true
}

// HasReferredBy returns a boolean if a field has been set.
func (o *Prospect1AllOf) HasReferredBy() bool {
	if o != nil && o.ReferredBy != nil {
		return true
	}

	return false
}

// SetReferredBy gets a reference to the given string and assigns it to the ReferredBy field.
func (o *Prospect1AllOf) SetReferredBy(v string) {
	o.ReferredBy = &v
}

// GetReferredByCode returns the ReferredByCode field value if set, zero value otherwise.
func (o *Prospect1AllOf) GetReferredByCode() string {
	if o == nil || o.ReferredByCode == nil {
		var ret string
		return ret
	}
	return *o.ReferredByCode
}

// GetReferredByCodeOk returns a tuple with the ReferredByCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Prospect1AllOf) GetReferredByCodeOk() (*string, bool) {
	if o == nil || o.ReferredByCode == nil {
		return nil, false
	}
	return o.ReferredByCode, true
}

// HasReferredByCode returns a boolean if a field has been set.
func (o *Prospect1AllOf) HasReferredByCode() bool {
	if o != nil && o.ReferredByCode != nil {
		return true
	}

	return false
}

// SetReferredByCode gets a reference to the given string and assigns it to the ReferredByCode field.
func (o *Prospect1AllOf) SetReferredByCode(v string) {
	o.ReferredByCode = &v
}

// GetReferredProspects returns the ReferredProspects field value if set, zero value otherwise.
func (o *Prospect1AllOf) GetReferredProspects() int32 {
	if o == nil || o.ReferredProspects == nil {
		var ret int32
		return ret
	}
	return *o.ReferredProspects
}

// GetReferredProspectsOk returns a tuple with the ReferredProspects field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Prospect1AllOf) GetReferredProspectsOk() (*int32, bool) {
	if o == nil || o.ReferredProspects == nil {
		return nil, false
	}
	return o.ReferredProspects, true
}

// HasReferredProspects returns a boolean if a field has been set.
func (o *Prospect1AllOf) HasReferredProspects() bool {
	if o != nil && o.ReferredProspects != nil {
		return true
	}

	return false
}

// SetReferredProspects gets a reference to the given int32 and assigns it to the ReferredProspects field.
func (o *Prospect1AllOf) SetReferredProspects(v int32) {
	o.ReferredProspects = &v
}

// GetVendorInfo returns the VendorInfo field value if set, zero value otherwise.
func (o *Prospect1AllOf) GetVendorInfo() VendorInfo {
	if o == nil || o.VendorInfo == nil {
		var ret VendorInfo
		return ret
	}
	return *o.VendorInfo
}

// GetVendorInfoOk returns a tuple with the VendorInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Prospect1AllOf) GetVendorInfoOk() (*VendorInfo, bool) {
	if o == nil || o.VendorInfo == nil {
		return nil, false
	}
	return o.VendorInfo, true
}

// HasVendorInfo returns a boolean if a field has been set.
func (o *Prospect1AllOf) HasVendorInfo() bool {
	if o != nil && o.VendorInfo != nil {
		return true
	}

	return false
}

// SetVendorInfo gets a reference to the given VendorInfo and assigns it to the VendorInfo field.
func (o *Prospect1AllOf) SetVendorInfo(v VendorInfo) {
	o.VendorInfo = &v
}

// GetVerificationToken returns the VerificationToken field value if set, zero value otherwise.
func (o *Prospect1AllOf) GetVerificationToken() string {
	if o == nil || o.VerificationToken == nil {
		var ret string
		return ret
	}
	return *o.VerificationToken
}

// GetVerificationTokenOk returns a tuple with the VerificationToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Prospect1AllOf) GetVerificationTokenOk() (*string, bool) {
	if o == nil || o.VerificationToken == nil {
		return nil, false
	}
	return o.VerificationToken, true
}

// HasVerificationToken returns a boolean if a field has been set.
func (o *Prospect1AllOf) HasVerificationToken() bool {
	if o != nil && o.VerificationToken != nil {
		return true
	}

	return false
}

// SetVerificationToken gets a reference to the given string and assigns it to the VerificationToken field.
func (o *Prospect1AllOf) SetVerificationToken(v string) {
	o.VerificationToken = &v
}

// GetWaitlistId returns the WaitlistId field value if set, zero value otherwise.
func (o *Prospect1AllOf) GetWaitlistId() string {
	if o == nil || o.WaitlistId == nil {
		var ret string
		return ret
	}
	return *o.WaitlistId
}

// GetWaitlistIdOk returns a tuple with the WaitlistId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Prospect1AllOf) GetWaitlistIdOk() (*string, bool) {
	if o == nil || o.WaitlistId == nil {
		return nil, false
	}
	return o.WaitlistId, true
}

// HasWaitlistId returns a boolean if a field has been set.
func (o *Prospect1AllOf) HasWaitlistId() bool {
	if o != nil && o.WaitlistId != nil {
		return true
	}

	return false
}

// SetWaitlistId gets a reference to the given string and assigns it to the WaitlistId field.
func (o *Prospect1AllOf) SetWaitlistId(v string) {
	o.WaitlistId = &v
}

// GetWaitlistPosition returns the WaitlistPosition field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Prospect1AllOf) GetWaitlistPosition() int32 {
	if o == nil || o.WaitlistPosition.Get() == nil {
		var ret int32
		return ret
	}
	return *o.WaitlistPosition.Get()
}

// GetWaitlistPositionOk returns a tuple with the WaitlistPosition field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Prospect1AllOf) GetWaitlistPositionOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.WaitlistPosition.Get(), o.WaitlistPosition.IsSet()
}

// HasWaitlistPosition returns a boolean if a field has been set.
func (o *Prospect1AllOf) HasWaitlistPosition() bool {
	if o != nil && o.WaitlistPosition.IsSet() {
		return true
	}

	return false
}

// SetWaitlistPosition gets a reference to the given NullableInt32 and assigns it to the WaitlistPosition field.
func (o *Prospect1AllOf) SetWaitlistPosition(v int32) {
	o.WaitlistPosition.Set(&v)
}

// SetWaitlistPositionNil sets the value for WaitlistPosition to be an explicit nil
func (o *Prospect1AllOf) SetWaitlistPositionNil() {
	o.WaitlistPosition.Set(nil)
}

// UnsetWaitlistPosition ensures that no value is present for WaitlistPosition, not even an explicit nil
func (o *Prospect1AllOf) UnsetWaitlistPosition() {
	o.WaitlistPosition.Unset()
}

func (o Prospect1AllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.CreationTime != nil {
		toSerialize["creation_time"] = o.CreationTime
	}
	if true {
		toSerialize["email"] = o.Email
	}
	if o.EmailValidation != nil {
		toSerialize["email_validation"] = o.EmailValidation
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.LastUpdatedTime != nil {
		toSerialize["last_updated_time"] = o.LastUpdatedTime
	}
	if o.Points != nil {
		toSerialize["points"] = o.Points
	}
	if o.RecaptchaToken != nil {
		toSerialize["recaptcha_token"] = o.RecaptchaToken
	}
	if o.ReferralCode != nil {
		toSerialize["referral_code"] = o.ReferralCode
	}
	if o.ReferredBy != nil {
		toSerialize["referred_by"] = o.ReferredBy
	}
	if o.ReferredByCode != nil {
		toSerialize["referred_by_code"] = o.ReferredByCode
	}
	if o.ReferredProspects != nil {
		toSerialize["referred_prospects"] = o.ReferredProspects
	}
	if o.VendorInfo != nil {
		toSerialize["vendor_info"] = o.VendorInfo
	}
	if o.VerificationToken != nil {
		toSerialize["verification_token"] = o.VerificationToken
	}
	if o.WaitlistId != nil {
		toSerialize["waitlist_id"] = o.WaitlistId
	}
	if o.WaitlistPosition.IsSet() {
		toSerialize["waitlist_position"] = o.WaitlistPosition.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableProspect1AllOf struct {
	value *Prospect1AllOf
	isSet bool
}

func (v NullableProspect1AllOf) Get() *Prospect1AllOf {
	return v.value
}

func (v *NullableProspect1AllOf) Set(val *Prospect1AllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableProspect1AllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableProspect1AllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProspect1AllOf(val *Prospect1AllOf) *NullableProspect1AllOf {
	return &NullableProspect1AllOf{value: val, isSet: true}
}

func (v NullableProspect1AllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProspect1AllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
