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

// Prospect A prospect has a unique identifier. It can be upgrade to a customer with required information
type Prospect struct {
	// Customer's date of birth in RFC 3339 full-date format (YYYY-MM-DD)
	Dob *string `json:"dob,omitempty"`
	// Customer's first name
	FirstName *string `json:"first_name,omitempty"`
	// Customer's last name
	LastName *string `json:"last_name,omitempty"`
	// Customer's status
	Status    string     `json:"status"`
	BanStatus *BanStatus `json:"ban_status,omitempty"`
	// The date and time the resource was created.
	CreationTime *time.Time `json:"creation_time,omitempty"`
	// Customer's email
	Email *string `json:"email,omitempty"`
	// Customer unique identifier
	Id *string `json:"id,omitempty"`
	// Customer's KYC exemption
	KycExempt *bool `json:"kyc_exempt,omitempty"`
	// Date and time KYC was last run on the customer
	KycLastRun *time.Time         `json:"kyc_last_run,omitempty"`
	KycStatus  *CustomerKycStatus `json:"kyc_status,omitempty"`
	// The date and time the resource was last updated.
	LastUpdatedTime *time.Time `json:"last_updated_time,omitempty"`
	LegalAddress    *Address   `json:"legal_address,omitempty"`
	// User-supplied metadata. Do not use to store PII.
	Metadata map[string]interface{} `json:"metadata,omitempty"`
	// Customer's middle name
	MiddleName *string `json:"middle_name,omitempty"`
	// Customer's mobile phone number with country code in E.164 format. Must have a valid country code. Area code and local phone number are not validated.
	PhoneNumber *string `json:"phone_number,omitempty"`
	// Customer's relationships with other accounts eg. guardian
	RelatedCustomers []Relationship1 `json:"related_customers,omitempty"`
	ShippingAddress  *Address        `json:"shipping_address,omitempty"`
	// Customer's full tax ID eg SSN formatted with hyphens. This optional parameter is required when running KYC on a customer. Input must match the pattern ^\\d{3}-\\d{2}-\\d{4}$. The response contains the last 4 digits only (e.g. 6789).
	Ssn       *string    `json:"ssn,omitempty"`
	SsnSource *SsnSource `json:"ssn_source,omitempty"`
}

// NewProspect instantiates a new Prospect object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProspect(status string) *Prospect {
	this := Prospect{}
	this.Status = status
	return &this
}

// NewProspectWithDefaults instantiates a new Prospect object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProspectWithDefaults() *Prospect {
	this := Prospect{}
	return &this
}

// GetDob returns the Dob field value if set, zero value otherwise.
func (o *Prospect) GetDob() string {
	if o == nil || o.Dob == nil {
		var ret string
		return ret
	}
	return *o.Dob
}

// GetDobOk returns a tuple with the Dob field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Prospect) GetDobOk() (*string, bool) {
	if o == nil || o.Dob == nil {
		return nil, false
	}
	return o.Dob, true
}

// HasDob returns a boolean if a field has been set.
func (o *Prospect) HasDob() bool {
	if o != nil && o.Dob != nil {
		return true
	}

	return false
}

// SetDob gets a reference to the given string and assigns it to the Dob field.
func (o *Prospect) SetDob(v string) {
	o.Dob = &v
}

// GetFirstName returns the FirstName field value if set, zero value otherwise.
func (o *Prospect) GetFirstName() string {
	if o == nil || o.FirstName == nil {
		var ret string
		return ret
	}
	return *o.FirstName
}

// GetFirstNameOk returns a tuple with the FirstName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Prospect) GetFirstNameOk() (*string, bool) {
	if o == nil || o.FirstName == nil {
		return nil, false
	}
	return o.FirstName, true
}

// HasFirstName returns a boolean if a field has been set.
func (o *Prospect) HasFirstName() bool {
	if o != nil && o.FirstName != nil {
		return true
	}

	return false
}

// SetFirstName gets a reference to the given string and assigns it to the FirstName field.
func (o *Prospect) SetFirstName(v string) {
	o.FirstName = &v
}

// GetLastName returns the LastName field value if set, zero value otherwise.
func (o *Prospect) GetLastName() string {
	if o == nil || o.LastName == nil {
		var ret string
		return ret
	}
	return *o.LastName
}

// GetLastNameOk returns a tuple with the LastName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Prospect) GetLastNameOk() (*string, bool) {
	if o == nil || o.LastName == nil {
		return nil, false
	}
	return o.LastName, true
}

// HasLastName returns a boolean if a field has been set.
func (o *Prospect) HasLastName() bool {
	if o != nil && o.LastName != nil {
		return true
	}

	return false
}

// SetLastName gets a reference to the given string and assigns it to the LastName field.
func (o *Prospect) SetLastName(v string) {
	o.LastName = &v
}

// GetStatus returns the Status field value
func (o *Prospect) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *Prospect) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *Prospect) SetStatus(v string) {
	o.Status = v
}

// GetBanStatus returns the BanStatus field value if set, zero value otherwise.
func (o *Prospect) GetBanStatus() BanStatus {
	if o == nil || o.BanStatus == nil {
		var ret BanStatus
		return ret
	}
	return *o.BanStatus
}

// GetBanStatusOk returns a tuple with the BanStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Prospect) GetBanStatusOk() (*BanStatus, bool) {
	if o == nil || o.BanStatus == nil {
		return nil, false
	}
	return o.BanStatus, true
}

// HasBanStatus returns a boolean if a field has been set.
func (o *Prospect) HasBanStatus() bool {
	if o != nil && o.BanStatus != nil {
		return true
	}

	return false
}

// SetBanStatus gets a reference to the given BanStatus and assigns it to the BanStatus field.
func (o *Prospect) SetBanStatus(v BanStatus) {
	o.BanStatus = &v
}

// GetCreationTime returns the CreationTime field value if set, zero value otherwise.
func (o *Prospect) GetCreationTime() time.Time {
	if o == nil || o.CreationTime == nil {
		var ret time.Time
		return ret
	}
	return *o.CreationTime
}

// GetCreationTimeOk returns a tuple with the CreationTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Prospect) GetCreationTimeOk() (*time.Time, bool) {
	if o == nil || o.CreationTime == nil {
		return nil, false
	}
	return o.CreationTime, true
}

// HasCreationTime returns a boolean if a field has been set.
func (o *Prospect) HasCreationTime() bool {
	if o != nil && o.CreationTime != nil {
		return true
	}

	return false
}

// SetCreationTime gets a reference to the given time.Time and assigns it to the CreationTime field.
func (o *Prospect) SetCreationTime(v time.Time) {
	o.CreationTime = &v
}

// GetEmail returns the Email field value if set, zero value otherwise.
func (o *Prospect) GetEmail() string {
	if o == nil || o.Email == nil {
		var ret string
		return ret
	}
	return *o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Prospect) GetEmailOk() (*string, bool) {
	if o == nil || o.Email == nil {
		return nil, false
	}
	return o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *Prospect) HasEmail() bool {
	if o != nil && o.Email != nil {
		return true
	}

	return false
}

// SetEmail gets a reference to the given string and assigns it to the Email field.
func (o *Prospect) SetEmail(v string) {
	o.Email = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Prospect) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Prospect) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Prospect) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *Prospect) SetId(v string) {
	o.Id = &v
}

// GetKycExempt returns the KycExempt field value if set, zero value otherwise.
func (o *Prospect) GetKycExempt() bool {
	if o == nil || o.KycExempt == nil {
		var ret bool
		return ret
	}
	return *o.KycExempt
}

// GetKycExemptOk returns a tuple with the KycExempt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Prospect) GetKycExemptOk() (*bool, bool) {
	if o == nil || o.KycExempt == nil {
		return nil, false
	}
	return o.KycExempt, true
}

// HasKycExempt returns a boolean if a field has been set.
func (o *Prospect) HasKycExempt() bool {
	if o != nil && o.KycExempt != nil {
		return true
	}

	return false
}

// SetKycExempt gets a reference to the given bool and assigns it to the KycExempt field.
func (o *Prospect) SetKycExempt(v bool) {
	o.KycExempt = &v
}

// GetKycLastRun returns the KycLastRun field value if set, zero value otherwise.
func (o *Prospect) GetKycLastRun() time.Time {
	if o == nil || o.KycLastRun == nil {
		var ret time.Time
		return ret
	}
	return *o.KycLastRun
}

// GetKycLastRunOk returns a tuple with the KycLastRun field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Prospect) GetKycLastRunOk() (*time.Time, bool) {
	if o == nil || o.KycLastRun == nil {
		return nil, false
	}
	return o.KycLastRun, true
}

// HasKycLastRun returns a boolean if a field has been set.
func (o *Prospect) HasKycLastRun() bool {
	if o != nil && o.KycLastRun != nil {
		return true
	}

	return false
}

// SetKycLastRun gets a reference to the given time.Time and assigns it to the KycLastRun field.
func (o *Prospect) SetKycLastRun(v time.Time) {
	o.KycLastRun = &v
}

// GetKycStatus returns the KycStatus field value if set, zero value otherwise.
func (o *Prospect) GetKycStatus() CustomerKycStatus {
	if o == nil || o.KycStatus == nil {
		var ret CustomerKycStatus
		return ret
	}
	return *o.KycStatus
}

// GetKycStatusOk returns a tuple with the KycStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Prospect) GetKycStatusOk() (*CustomerKycStatus, bool) {
	if o == nil || o.KycStatus == nil {
		return nil, false
	}
	return o.KycStatus, true
}

// HasKycStatus returns a boolean if a field has been set.
func (o *Prospect) HasKycStatus() bool {
	if o != nil && o.KycStatus != nil {
		return true
	}

	return false
}

// SetKycStatus gets a reference to the given CustomerKycStatus and assigns it to the KycStatus field.
func (o *Prospect) SetKycStatus(v CustomerKycStatus) {
	o.KycStatus = &v
}

// GetLastUpdatedTime returns the LastUpdatedTime field value if set, zero value otherwise.
func (o *Prospect) GetLastUpdatedTime() time.Time {
	if o == nil || o.LastUpdatedTime == nil {
		var ret time.Time
		return ret
	}
	return *o.LastUpdatedTime
}

// GetLastUpdatedTimeOk returns a tuple with the LastUpdatedTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Prospect) GetLastUpdatedTimeOk() (*time.Time, bool) {
	if o == nil || o.LastUpdatedTime == nil {
		return nil, false
	}
	return o.LastUpdatedTime, true
}

// HasLastUpdatedTime returns a boolean if a field has been set.
func (o *Prospect) HasLastUpdatedTime() bool {
	if o != nil && o.LastUpdatedTime != nil {
		return true
	}

	return false
}

// SetLastUpdatedTime gets a reference to the given time.Time and assigns it to the LastUpdatedTime field.
func (o *Prospect) SetLastUpdatedTime(v time.Time) {
	o.LastUpdatedTime = &v
}

// GetLegalAddress returns the LegalAddress field value if set, zero value otherwise.
func (o *Prospect) GetLegalAddress() Address {
	if o == nil || o.LegalAddress == nil {
		var ret Address
		return ret
	}
	return *o.LegalAddress
}

// GetLegalAddressOk returns a tuple with the LegalAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Prospect) GetLegalAddressOk() (*Address, bool) {
	if o == nil || o.LegalAddress == nil {
		return nil, false
	}
	return o.LegalAddress, true
}

// HasLegalAddress returns a boolean if a field has been set.
func (o *Prospect) HasLegalAddress() bool {
	if o != nil && o.LegalAddress != nil {
		return true
	}

	return false
}

// SetLegalAddress gets a reference to the given Address and assigns it to the LegalAddress field.
func (o *Prospect) SetLegalAddress(v Address) {
	o.LegalAddress = &v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *Prospect) GetMetadata() map[string]interface{} {
	if o == nil || o.Metadata == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Prospect) GetMetadataOk() (map[string]interface{}, bool) {
	if o == nil || o.Metadata == nil {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *Prospect) HasMetadata() bool {
	if o != nil && o.Metadata != nil {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given map[string]interface{} and assigns it to the Metadata field.
func (o *Prospect) SetMetadata(v map[string]interface{}) {
	o.Metadata = v
}

// GetMiddleName returns the MiddleName field value if set, zero value otherwise.
func (o *Prospect) GetMiddleName() string {
	if o == nil || o.MiddleName == nil {
		var ret string
		return ret
	}
	return *o.MiddleName
}

// GetMiddleNameOk returns a tuple with the MiddleName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Prospect) GetMiddleNameOk() (*string, bool) {
	if o == nil || o.MiddleName == nil {
		return nil, false
	}
	return o.MiddleName, true
}

// HasMiddleName returns a boolean if a field has been set.
func (o *Prospect) HasMiddleName() bool {
	if o != nil && o.MiddleName != nil {
		return true
	}

	return false
}

// SetMiddleName gets a reference to the given string and assigns it to the MiddleName field.
func (o *Prospect) SetMiddleName(v string) {
	o.MiddleName = &v
}

// GetPhoneNumber returns the PhoneNumber field value if set, zero value otherwise.
func (o *Prospect) GetPhoneNumber() string {
	if o == nil || o.PhoneNumber == nil {
		var ret string
		return ret
	}
	return *o.PhoneNumber
}

// GetPhoneNumberOk returns a tuple with the PhoneNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Prospect) GetPhoneNumberOk() (*string, bool) {
	if o == nil || o.PhoneNumber == nil {
		return nil, false
	}
	return o.PhoneNumber, true
}

// HasPhoneNumber returns a boolean if a field has been set.
func (o *Prospect) HasPhoneNumber() bool {
	if o != nil && o.PhoneNumber != nil {
		return true
	}

	return false
}

// SetPhoneNumber gets a reference to the given string and assigns it to the PhoneNumber field.
func (o *Prospect) SetPhoneNumber(v string) {
	o.PhoneNumber = &v
}

// GetRelatedCustomers returns the RelatedCustomers field value if set, zero value otherwise.
func (o *Prospect) GetRelatedCustomers() []Relationship1 {
	if o == nil || o.RelatedCustomers == nil {
		var ret []Relationship1
		return ret
	}
	return o.RelatedCustomers
}

// GetRelatedCustomersOk returns a tuple with the RelatedCustomers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Prospect) GetRelatedCustomersOk() ([]Relationship1, bool) {
	if o == nil || o.RelatedCustomers == nil {
		return nil, false
	}
	return o.RelatedCustomers, true
}

// HasRelatedCustomers returns a boolean if a field has been set.
func (o *Prospect) HasRelatedCustomers() bool {
	if o != nil && o.RelatedCustomers != nil {
		return true
	}

	return false
}

// SetRelatedCustomers gets a reference to the given []Relationship1 and assigns it to the RelatedCustomers field.
func (o *Prospect) SetRelatedCustomers(v []Relationship1) {
	o.RelatedCustomers = v
}

// GetShippingAddress returns the ShippingAddress field value if set, zero value otherwise.
func (o *Prospect) GetShippingAddress() Address {
	if o == nil || o.ShippingAddress == nil {
		var ret Address
		return ret
	}
	return *o.ShippingAddress
}

// GetShippingAddressOk returns a tuple with the ShippingAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Prospect) GetShippingAddressOk() (*Address, bool) {
	if o == nil || o.ShippingAddress == nil {
		return nil, false
	}
	return o.ShippingAddress, true
}

// HasShippingAddress returns a boolean if a field has been set.
func (o *Prospect) HasShippingAddress() bool {
	if o != nil && o.ShippingAddress != nil {
		return true
	}

	return false
}

// SetShippingAddress gets a reference to the given Address and assigns it to the ShippingAddress field.
func (o *Prospect) SetShippingAddress(v Address) {
	o.ShippingAddress = &v
}

// GetSsn returns the Ssn field value if set, zero value otherwise.
func (o *Prospect) GetSsn() string {
	if o == nil || o.Ssn == nil {
		var ret string
		return ret
	}
	return *o.Ssn
}

// GetSsnOk returns a tuple with the Ssn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Prospect) GetSsnOk() (*string, bool) {
	if o == nil || o.Ssn == nil {
		return nil, false
	}
	return o.Ssn, true
}

// HasSsn returns a boolean if a field has been set.
func (o *Prospect) HasSsn() bool {
	if o != nil && o.Ssn != nil {
		return true
	}

	return false
}

// SetSsn gets a reference to the given string and assigns it to the Ssn field.
func (o *Prospect) SetSsn(v string) {
	o.Ssn = &v
}

// GetSsnSource returns the SsnSource field value if set, zero value otherwise.
func (o *Prospect) GetSsnSource() SsnSource {
	if o == nil || o.SsnSource == nil {
		var ret SsnSource
		return ret
	}
	return *o.SsnSource
}

// GetSsnSourceOk returns a tuple with the SsnSource field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Prospect) GetSsnSourceOk() (*SsnSource, bool) {
	if o == nil || o.SsnSource == nil {
		return nil, false
	}
	return o.SsnSource, true
}

// HasSsnSource returns a boolean if a field has been set.
func (o *Prospect) HasSsnSource() bool {
	if o != nil && o.SsnSource != nil {
		return true
	}

	return false
}

// SetSsnSource gets a reference to the given SsnSource and assigns it to the SsnSource field.
func (o *Prospect) SetSsnSource(v SsnSource) {
	o.SsnSource = &v
}

func (o Prospect) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Dob != nil {
		toSerialize["dob"] = o.Dob
	}
	if o.FirstName != nil {
		toSerialize["first_name"] = o.FirstName
	}
	if o.LastName != nil {
		toSerialize["last_name"] = o.LastName
	}
	if true {
		toSerialize["status"] = o.Status
	}
	if o.BanStatus != nil {
		toSerialize["ban_status"] = o.BanStatus
	}
	if o.CreationTime != nil {
		toSerialize["creation_time"] = o.CreationTime
	}
	if o.Email != nil {
		toSerialize["email"] = o.Email
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.KycExempt != nil {
		toSerialize["kyc_exempt"] = o.KycExempt
	}
	if o.KycLastRun != nil {
		toSerialize["kyc_last_run"] = o.KycLastRun
	}
	if o.KycStatus != nil {
		toSerialize["kyc_status"] = o.KycStatus
	}
	if o.LastUpdatedTime != nil {
		toSerialize["last_updated_time"] = o.LastUpdatedTime
	}
	if o.LegalAddress != nil {
		toSerialize["legal_address"] = o.LegalAddress
	}
	if o.Metadata != nil {
		toSerialize["metadata"] = o.Metadata
	}
	if o.MiddleName != nil {
		toSerialize["middle_name"] = o.MiddleName
	}
	if o.PhoneNumber != nil {
		toSerialize["phone_number"] = o.PhoneNumber
	}
	if o.RelatedCustomers != nil {
		toSerialize["related_customers"] = o.RelatedCustomers
	}
	if o.ShippingAddress != nil {
		toSerialize["shipping_address"] = o.ShippingAddress
	}
	if o.Ssn != nil {
		toSerialize["ssn"] = o.Ssn
	}
	if o.SsnSource != nil {
		toSerialize["ssn_source"] = o.SsnSource
	}
	return json.Marshal(toSerialize)
}

type NullableProspect struct {
	value *Prospect
	isSet bool
}

func (v NullableProspect) Get() *Prospect {
	return v.value
}

func (v *NullableProspect) Set(val *Prospect) {
	v.value = val
	v.isSet = true
}

func (v NullableProspect) IsSet() bool {
	return v.isSet
}

func (v *NullableProspect) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProspect(val *Prospect) *NullableProspect {
	return &NullableProspect{value: val, isSet: true}
}

func (v NullableProspect) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProspect) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
