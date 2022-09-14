/*
Synctera API

<h2>Let's build something great.</h2><p>Welcome to the official reference documentation for Synctera APIs. Our APIs are the best way to automate your company's banking needs and are designed to be easy to understand and implement.</p><p>We're continuously growing this library and what you see here is just the start, but if you need something specific or have a question, <a class='text-blue-600' href='https://synctera.com/contact' target='_blank' rel='noreferrer'>contact us</a>.</p>

API version: 0.20.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package synctera_client

import (
	"encoding/json"
)

// ApplicationRequest struct for ApplicationRequest
type ApplicationRequest struct {
	// Bank ID
	BankId          int32    `json:"bank_id"`
	BusinessAddress Address1 `json:"business_address"`
	// todo
	BusinessName string `json:"business_name"`
	// todo
	BusinessPhone string `json:"business_phone"`
	// todo
	BusinessTaxId string `json:"business_tax_id"`
	// todo
	BusinessType string               `json:"business_type"`
	Dob          *ExternalPaymentDate `json:"dob,omitempty"`
	// todo
	DoingBusinessAs string `json:"doing_business_as"`
	// todo
	Email *string `json:"email,omitempty"`
	// To enable or disable aft/oct feature
	Enabled bool `json:"enabled"`
	// todo
	FirstName         *string              `json:"first_name,omitempty"`
	IncorporationDate *ExternalPaymentDate `json:"incorporation_date,omitempty"`
	// todo
	LastName *string `json:"last_name,omitempty"`
	// Maximum amount that can be transacted for a single transaction in cents
	MaxTransactionAmount int32 `json:"max_transaction_amount"`
	// Partner ID
	PartnerId       int32     `json:"partner_id"`
	PersonalAddress *Address1 `json:"personal_address,omitempty"`
	// todo
	Phone *string `json:"phone,omitempty"`
	// todo
	PrincipalPercentageOwnership *string   `json:"principal_percentage_ownership,omitempty"`
	Processor                    Processor `json:"processor"`
	// todo
	TaxId *string `json:"tax_id,omitempty"`
	// todo
	Title *string `json:"title,omitempty"`
	// todo
	Url *string `json:"url,omitempty"`
}

// NewApplicationRequest instantiates a new ApplicationRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApplicationRequest(bankId int32, businessAddress Address1, businessName string, businessPhone string, businessTaxId string, businessType string, doingBusinessAs string, enabled bool, maxTransactionAmount int32, partnerId int32, processor Processor) *ApplicationRequest {
	this := ApplicationRequest{}
	this.BankId = bankId
	this.BusinessAddress = businessAddress
	this.BusinessName = businessName
	this.BusinessPhone = businessPhone
	this.BusinessTaxId = businessTaxId
	this.BusinessType = businessType
	this.DoingBusinessAs = doingBusinessAs
	this.Enabled = enabled
	this.MaxTransactionAmount = maxTransactionAmount
	this.PartnerId = partnerId
	this.Processor = processor
	return &this
}

// NewApplicationRequestWithDefaults instantiates a new ApplicationRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApplicationRequestWithDefaults() *ApplicationRequest {
	this := ApplicationRequest{}
	return &this
}

// GetBankId returns the BankId field value
func (o *ApplicationRequest) GetBankId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.BankId
}

// GetBankIdOk returns a tuple with the BankId field value
// and a boolean to check if the value has been set.
func (o *ApplicationRequest) GetBankIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BankId, true
}

// SetBankId sets field value
func (o *ApplicationRequest) SetBankId(v int32) {
	o.BankId = v
}

// GetBusinessAddress returns the BusinessAddress field value
func (o *ApplicationRequest) GetBusinessAddress() Address1 {
	if o == nil {
		var ret Address1
		return ret
	}

	return o.BusinessAddress
}

// GetBusinessAddressOk returns a tuple with the BusinessAddress field value
// and a boolean to check if the value has been set.
func (o *ApplicationRequest) GetBusinessAddressOk() (*Address1, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BusinessAddress, true
}

// SetBusinessAddress sets field value
func (o *ApplicationRequest) SetBusinessAddress(v Address1) {
	o.BusinessAddress = v
}

// GetBusinessName returns the BusinessName field value
func (o *ApplicationRequest) GetBusinessName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.BusinessName
}

// GetBusinessNameOk returns a tuple with the BusinessName field value
// and a boolean to check if the value has been set.
func (o *ApplicationRequest) GetBusinessNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BusinessName, true
}

// SetBusinessName sets field value
func (o *ApplicationRequest) SetBusinessName(v string) {
	o.BusinessName = v
}

// GetBusinessPhone returns the BusinessPhone field value
func (o *ApplicationRequest) GetBusinessPhone() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.BusinessPhone
}

// GetBusinessPhoneOk returns a tuple with the BusinessPhone field value
// and a boolean to check if the value has been set.
func (o *ApplicationRequest) GetBusinessPhoneOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BusinessPhone, true
}

// SetBusinessPhone sets field value
func (o *ApplicationRequest) SetBusinessPhone(v string) {
	o.BusinessPhone = v
}

// GetBusinessTaxId returns the BusinessTaxId field value
func (o *ApplicationRequest) GetBusinessTaxId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.BusinessTaxId
}

// GetBusinessTaxIdOk returns a tuple with the BusinessTaxId field value
// and a boolean to check if the value has been set.
func (o *ApplicationRequest) GetBusinessTaxIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BusinessTaxId, true
}

// SetBusinessTaxId sets field value
func (o *ApplicationRequest) SetBusinessTaxId(v string) {
	o.BusinessTaxId = v
}

// GetBusinessType returns the BusinessType field value
func (o *ApplicationRequest) GetBusinessType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.BusinessType
}

// GetBusinessTypeOk returns a tuple with the BusinessType field value
// and a boolean to check if the value has been set.
func (o *ApplicationRequest) GetBusinessTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BusinessType, true
}

// SetBusinessType sets field value
func (o *ApplicationRequest) SetBusinessType(v string) {
	o.BusinessType = v
}

// GetDob returns the Dob field value if set, zero value otherwise.
func (o *ApplicationRequest) GetDob() ExternalPaymentDate {
	if o == nil || o.Dob == nil {
		var ret ExternalPaymentDate
		return ret
	}
	return *o.Dob
}

// GetDobOk returns a tuple with the Dob field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationRequest) GetDobOk() (*ExternalPaymentDate, bool) {
	if o == nil || o.Dob == nil {
		return nil, false
	}
	return o.Dob, true
}

// HasDob returns a boolean if a field has been set.
func (o *ApplicationRequest) HasDob() bool {
	if o != nil && o.Dob != nil {
		return true
	}

	return false
}

// SetDob gets a reference to the given ExternalPaymentDate and assigns it to the Dob field.
func (o *ApplicationRequest) SetDob(v ExternalPaymentDate) {
	o.Dob = &v
}

// GetDoingBusinessAs returns the DoingBusinessAs field value
func (o *ApplicationRequest) GetDoingBusinessAs() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DoingBusinessAs
}

// GetDoingBusinessAsOk returns a tuple with the DoingBusinessAs field value
// and a boolean to check if the value has been set.
func (o *ApplicationRequest) GetDoingBusinessAsOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DoingBusinessAs, true
}

// SetDoingBusinessAs sets field value
func (o *ApplicationRequest) SetDoingBusinessAs(v string) {
	o.DoingBusinessAs = v
}

// GetEmail returns the Email field value if set, zero value otherwise.
func (o *ApplicationRequest) GetEmail() string {
	if o == nil || o.Email == nil {
		var ret string
		return ret
	}
	return *o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationRequest) GetEmailOk() (*string, bool) {
	if o == nil || o.Email == nil {
		return nil, false
	}
	return o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *ApplicationRequest) HasEmail() bool {
	if o != nil && o.Email != nil {
		return true
	}

	return false
}

// SetEmail gets a reference to the given string and assigns it to the Email field.
func (o *ApplicationRequest) SetEmail(v string) {
	o.Email = &v
}

// GetEnabled returns the Enabled field value
func (o *ApplicationRequest) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *ApplicationRequest) GetEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value
func (o *ApplicationRequest) SetEnabled(v bool) {
	o.Enabled = v
}

// GetFirstName returns the FirstName field value if set, zero value otherwise.
func (o *ApplicationRequest) GetFirstName() string {
	if o == nil || o.FirstName == nil {
		var ret string
		return ret
	}
	return *o.FirstName
}

// GetFirstNameOk returns a tuple with the FirstName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationRequest) GetFirstNameOk() (*string, bool) {
	if o == nil || o.FirstName == nil {
		return nil, false
	}
	return o.FirstName, true
}

// HasFirstName returns a boolean if a field has been set.
func (o *ApplicationRequest) HasFirstName() bool {
	if o != nil && o.FirstName != nil {
		return true
	}

	return false
}

// SetFirstName gets a reference to the given string and assigns it to the FirstName field.
func (o *ApplicationRequest) SetFirstName(v string) {
	o.FirstName = &v
}

// GetIncorporationDate returns the IncorporationDate field value if set, zero value otherwise.
func (o *ApplicationRequest) GetIncorporationDate() ExternalPaymentDate {
	if o == nil || o.IncorporationDate == nil {
		var ret ExternalPaymentDate
		return ret
	}
	return *o.IncorporationDate
}

// GetIncorporationDateOk returns a tuple with the IncorporationDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationRequest) GetIncorporationDateOk() (*ExternalPaymentDate, bool) {
	if o == nil || o.IncorporationDate == nil {
		return nil, false
	}
	return o.IncorporationDate, true
}

// HasIncorporationDate returns a boolean if a field has been set.
func (o *ApplicationRequest) HasIncorporationDate() bool {
	if o != nil && o.IncorporationDate != nil {
		return true
	}

	return false
}

// SetIncorporationDate gets a reference to the given ExternalPaymentDate and assigns it to the IncorporationDate field.
func (o *ApplicationRequest) SetIncorporationDate(v ExternalPaymentDate) {
	o.IncorporationDate = &v
}

// GetLastName returns the LastName field value if set, zero value otherwise.
func (o *ApplicationRequest) GetLastName() string {
	if o == nil || o.LastName == nil {
		var ret string
		return ret
	}
	return *o.LastName
}

// GetLastNameOk returns a tuple with the LastName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationRequest) GetLastNameOk() (*string, bool) {
	if o == nil || o.LastName == nil {
		return nil, false
	}
	return o.LastName, true
}

// HasLastName returns a boolean if a field has been set.
func (o *ApplicationRequest) HasLastName() bool {
	if o != nil && o.LastName != nil {
		return true
	}

	return false
}

// SetLastName gets a reference to the given string and assigns it to the LastName field.
func (o *ApplicationRequest) SetLastName(v string) {
	o.LastName = &v
}

// GetMaxTransactionAmount returns the MaxTransactionAmount field value
func (o *ApplicationRequest) GetMaxTransactionAmount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.MaxTransactionAmount
}

// GetMaxTransactionAmountOk returns a tuple with the MaxTransactionAmount field value
// and a boolean to check if the value has been set.
func (o *ApplicationRequest) GetMaxTransactionAmountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MaxTransactionAmount, true
}

// SetMaxTransactionAmount sets field value
func (o *ApplicationRequest) SetMaxTransactionAmount(v int32) {
	o.MaxTransactionAmount = v
}

// GetPartnerId returns the PartnerId field value
func (o *ApplicationRequest) GetPartnerId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.PartnerId
}

// GetPartnerIdOk returns a tuple with the PartnerId field value
// and a boolean to check if the value has been set.
func (o *ApplicationRequest) GetPartnerIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PartnerId, true
}

// SetPartnerId sets field value
func (o *ApplicationRequest) SetPartnerId(v int32) {
	o.PartnerId = v
}

// GetPersonalAddress returns the PersonalAddress field value if set, zero value otherwise.
func (o *ApplicationRequest) GetPersonalAddress() Address1 {
	if o == nil || o.PersonalAddress == nil {
		var ret Address1
		return ret
	}
	return *o.PersonalAddress
}

// GetPersonalAddressOk returns a tuple with the PersonalAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationRequest) GetPersonalAddressOk() (*Address1, bool) {
	if o == nil || o.PersonalAddress == nil {
		return nil, false
	}
	return o.PersonalAddress, true
}

// HasPersonalAddress returns a boolean if a field has been set.
func (o *ApplicationRequest) HasPersonalAddress() bool {
	if o != nil && o.PersonalAddress != nil {
		return true
	}

	return false
}

// SetPersonalAddress gets a reference to the given Address1 and assigns it to the PersonalAddress field.
func (o *ApplicationRequest) SetPersonalAddress(v Address1) {
	o.PersonalAddress = &v
}

// GetPhone returns the Phone field value if set, zero value otherwise.
func (o *ApplicationRequest) GetPhone() string {
	if o == nil || o.Phone == nil {
		var ret string
		return ret
	}
	return *o.Phone
}

// GetPhoneOk returns a tuple with the Phone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationRequest) GetPhoneOk() (*string, bool) {
	if o == nil || o.Phone == nil {
		return nil, false
	}
	return o.Phone, true
}

// HasPhone returns a boolean if a field has been set.
func (o *ApplicationRequest) HasPhone() bool {
	if o != nil && o.Phone != nil {
		return true
	}

	return false
}

// SetPhone gets a reference to the given string and assigns it to the Phone field.
func (o *ApplicationRequest) SetPhone(v string) {
	o.Phone = &v
}

// GetPrincipalPercentageOwnership returns the PrincipalPercentageOwnership field value if set, zero value otherwise.
func (o *ApplicationRequest) GetPrincipalPercentageOwnership() string {
	if o == nil || o.PrincipalPercentageOwnership == nil {
		var ret string
		return ret
	}
	return *o.PrincipalPercentageOwnership
}

// GetPrincipalPercentageOwnershipOk returns a tuple with the PrincipalPercentageOwnership field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationRequest) GetPrincipalPercentageOwnershipOk() (*string, bool) {
	if o == nil || o.PrincipalPercentageOwnership == nil {
		return nil, false
	}
	return o.PrincipalPercentageOwnership, true
}

// HasPrincipalPercentageOwnership returns a boolean if a field has been set.
func (o *ApplicationRequest) HasPrincipalPercentageOwnership() bool {
	if o != nil && o.PrincipalPercentageOwnership != nil {
		return true
	}

	return false
}

// SetPrincipalPercentageOwnership gets a reference to the given string and assigns it to the PrincipalPercentageOwnership field.
func (o *ApplicationRequest) SetPrincipalPercentageOwnership(v string) {
	o.PrincipalPercentageOwnership = &v
}

// GetProcessor returns the Processor field value
func (o *ApplicationRequest) GetProcessor() Processor {
	if o == nil {
		var ret Processor
		return ret
	}

	return o.Processor
}

// GetProcessorOk returns a tuple with the Processor field value
// and a boolean to check if the value has been set.
func (o *ApplicationRequest) GetProcessorOk() (*Processor, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Processor, true
}

// SetProcessor sets field value
func (o *ApplicationRequest) SetProcessor(v Processor) {
	o.Processor = v
}

// GetTaxId returns the TaxId field value if set, zero value otherwise.
func (o *ApplicationRequest) GetTaxId() string {
	if o == nil || o.TaxId == nil {
		var ret string
		return ret
	}
	return *o.TaxId
}

// GetTaxIdOk returns a tuple with the TaxId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationRequest) GetTaxIdOk() (*string, bool) {
	if o == nil || o.TaxId == nil {
		return nil, false
	}
	return o.TaxId, true
}

// HasTaxId returns a boolean if a field has been set.
func (o *ApplicationRequest) HasTaxId() bool {
	if o != nil && o.TaxId != nil {
		return true
	}

	return false
}

// SetTaxId gets a reference to the given string and assigns it to the TaxId field.
func (o *ApplicationRequest) SetTaxId(v string) {
	o.TaxId = &v
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *ApplicationRequest) GetTitle() string {
	if o == nil || o.Title == nil {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationRequest) GetTitleOk() (*string, bool) {
	if o == nil || o.Title == nil {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *ApplicationRequest) HasTitle() bool {
	if o != nil && o.Title != nil {
		return true
	}

	return false
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *ApplicationRequest) SetTitle(v string) {
	o.Title = &v
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *ApplicationRequest) GetUrl() string {
	if o == nil || o.Url == nil {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationRequest) GetUrlOk() (*string, bool) {
	if o == nil || o.Url == nil {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *ApplicationRequest) HasUrl() bool {
	if o != nil && o.Url != nil {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *ApplicationRequest) SetUrl(v string) {
	o.Url = &v
}

func (o ApplicationRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["bank_id"] = o.BankId
	}
	if true {
		toSerialize["business_address"] = o.BusinessAddress
	}
	if true {
		toSerialize["business_name"] = o.BusinessName
	}
	if true {
		toSerialize["business_phone"] = o.BusinessPhone
	}
	if true {
		toSerialize["business_tax_id"] = o.BusinessTaxId
	}
	if true {
		toSerialize["business_type"] = o.BusinessType
	}
	if o.Dob != nil {
		toSerialize["dob"] = o.Dob
	}
	if true {
		toSerialize["doing_business_as"] = o.DoingBusinessAs
	}
	if o.Email != nil {
		toSerialize["email"] = o.Email
	}
	if true {
		toSerialize["enabled"] = o.Enabled
	}
	if o.FirstName != nil {
		toSerialize["first_name"] = o.FirstName
	}
	if o.IncorporationDate != nil {
		toSerialize["incorporation_date"] = o.IncorporationDate
	}
	if o.LastName != nil {
		toSerialize["last_name"] = o.LastName
	}
	if true {
		toSerialize["max_transaction_amount"] = o.MaxTransactionAmount
	}
	if true {
		toSerialize["partner_id"] = o.PartnerId
	}
	if o.PersonalAddress != nil {
		toSerialize["personal_address"] = o.PersonalAddress
	}
	if o.Phone != nil {
		toSerialize["phone"] = o.Phone
	}
	if o.PrincipalPercentageOwnership != nil {
		toSerialize["principal_percentage_ownership"] = o.PrincipalPercentageOwnership
	}
	if true {
		toSerialize["processor"] = o.Processor
	}
	if o.TaxId != nil {
		toSerialize["tax_id"] = o.TaxId
	}
	if o.Title != nil {
		toSerialize["title"] = o.Title
	}
	if o.Url != nil {
		toSerialize["url"] = o.Url
	}
	return json.Marshal(toSerialize)
}

type NullableApplicationRequest struct {
	value *ApplicationRequest
	isSet bool
}

func (v NullableApplicationRequest) Get() *ApplicationRequest {
	return v.value
}

func (v *NullableApplicationRequest) Set(val *ApplicationRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableApplicationRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableApplicationRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApplicationRequest(val *ApplicationRequest) *NullableApplicationRequest {
	return &NullableApplicationRequest{value: val, isSet: true}
}

func (v NullableApplicationRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApplicationRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
