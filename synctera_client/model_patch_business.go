/*
Synctera API

<h2>Let's build something great.</h2><p>Welcome to the official reference documentation for Synctera APIs. Our APIs are the best way to automate your company's banking needs and are designed to be easy to understand and implement.</p><p>We're continuously growing this library and what you see here is just the start, but if you need something specific or have a question, <a class='text-blue-600' href='https://synctera.com/contact' target='_blank' rel='noreferrer'>contact us</a>.</p> 

API version: 0.19.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package synctera_client

import (
	"encoding/json"
	"time"
)

// PatchBusiness Represents a business customer.
type PatchBusiness struct {
	// The date and time the resource was created.
	CreationTime *time.Time `json:"creation_time,omitempty"`
	// U.S. Employer Identification Number (EIN) for this business, in the format xx-xxxxxxx.
	Ein *string `json:"ein,omitempty"`
	// Business's email.
	Email *string `json:"email,omitempty"`
	// Business's legal name.
	EntityName *string `json:"entity_name,omitempty"`
	// Date the business was legally registered in RFC 3339 full-date format (YYYY-MM-DD).
	FormationDate *string `json:"formation_date,omitempty"`
	// U.S. state where the business is legally registered (2-letter abbreviation).
	FormationState *string `json:"formation_state,omitempty"`
	// Business's unique identifier.
	Id *string `json:"id,omitempty"`
	// True for personal and business customers with a direct relationship with the fintech or bank.
	IsCustomer *bool `json:"is_customer,omitempty"`
	// The date and time the resource was last updated.
	LastUpdatedTime *time.Time `json:"last_updated_time,omitempty"`
	LegalAddress *Address `json:"legal_address,omitempty"`
	// Optional field to store additional information about the resource. Intended to be used by the integrator to store non-sensitive data. 
	Metadata map[string]interface{} `json:"metadata,omitempty"`
	// Business's phone number with country code in E.164 format. Must have a valid country code. Area code and local phone number are not validated.
	PhoneNumber *string `json:"phone_number,omitempty"`
	// Status of the business. One of the following: * `PROSPECT` – a potential customer, used for information-gathering and disclosures. * `ACTIVE` –  is an integrator defined status.  Integrators should set a business to active if they believe the person to be qualified for conducting business.  Synctera will combine this status with other statuses such a verification to determine if the business is eligible for specific actions such as initiating transactions or issuing a card. * `FROZEN` – business's actions are blocked for security, legal, or other reasons. * `SANCTION` – business is on a sanctions list and should be carefully monitored. * `DISSOLVED` – an inactive status indicating a business entity has filed articles of dissolution or a certificate of termination to terminate its existence. * `CANCELLED` – an inactive status indicating that a business entity has filed a cancellation or has failed to file its periodic report after notice of forfeiture of its rights to do business. * `SUSPENDED` – an inactive status indicating that the business entity has lost the right to operate in it's registered jurisdiction. * `MERGED` – an inactive status indicating that the business entity has terminated existence by merging into another entity. * `INACTIVE` – an inactive status indicating that the business entity is no longer active. * `CONVERTED` – An inactive status indicating that the business entity has been converted to another type of business entity in the same jurisdiction. 
	Status *string `json:"status,omitempty"`
	// Business's legal structure.
	Structure *string `json:"structure,omitempty"`
	// Other names by which this business is known.
	TradeNames []string `json:"trade_names,omitempty"`
	// Date and time KYB verification was last run on the business.
	VerificationLastRun *time.Time `json:"verification_last_run,omitempty"`
	VerificationStatus *VerificationStatus `json:"verification_status,omitempty"`
	// Business's website.
	Website *string `json:"website,omitempty"`
}

// NewPatchBusiness instantiates a new PatchBusiness object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPatchBusiness() *PatchBusiness {
	this := PatchBusiness{}
	return &this
}

// NewPatchBusinessWithDefaults instantiates a new PatchBusiness object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPatchBusinessWithDefaults() *PatchBusiness {
	this := PatchBusiness{}
	return &this
}

// GetCreationTime returns the CreationTime field value if set, zero value otherwise.
func (o *PatchBusiness) GetCreationTime() time.Time {
	if o == nil || o.CreationTime == nil {
		var ret time.Time
		return ret
	}
	return *o.CreationTime
}

// GetCreationTimeOk returns a tuple with the CreationTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchBusiness) GetCreationTimeOk() (*time.Time, bool) {
	if o == nil || o.CreationTime == nil {
		return nil, false
	}
	return o.CreationTime, true
}

// HasCreationTime returns a boolean if a field has been set.
func (o *PatchBusiness) HasCreationTime() bool {
	if o != nil && o.CreationTime != nil {
		return true
	}

	return false
}

// SetCreationTime gets a reference to the given time.Time and assigns it to the CreationTime field.
func (o *PatchBusiness) SetCreationTime(v time.Time) {
	o.CreationTime = &v
}

// GetEin returns the Ein field value if set, zero value otherwise.
func (o *PatchBusiness) GetEin() string {
	if o == nil || o.Ein == nil {
		var ret string
		return ret
	}
	return *o.Ein
}

// GetEinOk returns a tuple with the Ein field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchBusiness) GetEinOk() (*string, bool) {
	if o == nil || o.Ein == nil {
		return nil, false
	}
	return o.Ein, true
}

// HasEin returns a boolean if a field has been set.
func (o *PatchBusiness) HasEin() bool {
	if o != nil && o.Ein != nil {
		return true
	}

	return false
}

// SetEin gets a reference to the given string and assigns it to the Ein field.
func (o *PatchBusiness) SetEin(v string) {
	o.Ein = &v
}

// GetEmail returns the Email field value if set, zero value otherwise.
func (o *PatchBusiness) GetEmail() string {
	if o == nil || o.Email == nil {
		var ret string
		return ret
	}
	return *o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchBusiness) GetEmailOk() (*string, bool) {
	if o == nil || o.Email == nil {
		return nil, false
	}
	return o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *PatchBusiness) HasEmail() bool {
	if o != nil && o.Email != nil {
		return true
	}

	return false
}

// SetEmail gets a reference to the given string and assigns it to the Email field.
func (o *PatchBusiness) SetEmail(v string) {
	o.Email = &v
}

// GetEntityName returns the EntityName field value if set, zero value otherwise.
func (o *PatchBusiness) GetEntityName() string {
	if o == nil || o.EntityName == nil {
		var ret string
		return ret
	}
	return *o.EntityName
}

// GetEntityNameOk returns a tuple with the EntityName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchBusiness) GetEntityNameOk() (*string, bool) {
	if o == nil || o.EntityName == nil {
		return nil, false
	}
	return o.EntityName, true
}

// HasEntityName returns a boolean if a field has been set.
func (o *PatchBusiness) HasEntityName() bool {
	if o != nil && o.EntityName != nil {
		return true
	}

	return false
}

// SetEntityName gets a reference to the given string and assigns it to the EntityName field.
func (o *PatchBusiness) SetEntityName(v string) {
	o.EntityName = &v
}

// GetFormationDate returns the FormationDate field value if set, zero value otherwise.
func (o *PatchBusiness) GetFormationDate() string {
	if o == nil || o.FormationDate == nil {
		var ret string
		return ret
	}
	return *o.FormationDate
}

// GetFormationDateOk returns a tuple with the FormationDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchBusiness) GetFormationDateOk() (*string, bool) {
	if o == nil || o.FormationDate == nil {
		return nil, false
	}
	return o.FormationDate, true
}

// HasFormationDate returns a boolean if a field has been set.
func (o *PatchBusiness) HasFormationDate() bool {
	if o != nil && o.FormationDate != nil {
		return true
	}

	return false
}

// SetFormationDate gets a reference to the given string and assigns it to the FormationDate field.
func (o *PatchBusiness) SetFormationDate(v string) {
	o.FormationDate = &v
}

// GetFormationState returns the FormationState field value if set, zero value otherwise.
func (o *PatchBusiness) GetFormationState() string {
	if o == nil || o.FormationState == nil {
		var ret string
		return ret
	}
	return *o.FormationState
}

// GetFormationStateOk returns a tuple with the FormationState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchBusiness) GetFormationStateOk() (*string, bool) {
	if o == nil || o.FormationState == nil {
		return nil, false
	}
	return o.FormationState, true
}

// HasFormationState returns a boolean if a field has been set.
func (o *PatchBusiness) HasFormationState() bool {
	if o != nil && o.FormationState != nil {
		return true
	}

	return false
}

// SetFormationState gets a reference to the given string and assigns it to the FormationState field.
func (o *PatchBusiness) SetFormationState(v string) {
	o.FormationState = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *PatchBusiness) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchBusiness) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *PatchBusiness) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *PatchBusiness) SetId(v string) {
	o.Id = &v
}

// GetIsCustomer returns the IsCustomer field value if set, zero value otherwise.
func (o *PatchBusiness) GetIsCustomer() bool {
	if o == nil || o.IsCustomer == nil {
		var ret bool
		return ret
	}
	return *o.IsCustomer
}

// GetIsCustomerOk returns a tuple with the IsCustomer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchBusiness) GetIsCustomerOk() (*bool, bool) {
	if o == nil || o.IsCustomer == nil {
		return nil, false
	}
	return o.IsCustomer, true
}

// HasIsCustomer returns a boolean if a field has been set.
func (o *PatchBusiness) HasIsCustomer() bool {
	if o != nil && o.IsCustomer != nil {
		return true
	}

	return false
}

// SetIsCustomer gets a reference to the given bool and assigns it to the IsCustomer field.
func (o *PatchBusiness) SetIsCustomer(v bool) {
	o.IsCustomer = &v
}

// GetLastUpdatedTime returns the LastUpdatedTime field value if set, zero value otherwise.
func (o *PatchBusiness) GetLastUpdatedTime() time.Time {
	if o == nil || o.LastUpdatedTime == nil {
		var ret time.Time
		return ret
	}
	return *o.LastUpdatedTime
}

// GetLastUpdatedTimeOk returns a tuple with the LastUpdatedTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchBusiness) GetLastUpdatedTimeOk() (*time.Time, bool) {
	if o == nil || o.LastUpdatedTime == nil {
		return nil, false
	}
	return o.LastUpdatedTime, true
}

// HasLastUpdatedTime returns a boolean if a field has been set.
func (o *PatchBusiness) HasLastUpdatedTime() bool {
	if o != nil && o.LastUpdatedTime != nil {
		return true
	}

	return false
}

// SetLastUpdatedTime gets a reference to the given time.Time and assigns it to the LastUpdatedTime field.
func (o *PatchBusiness) SetLastUpdatedTime(v time.Time) {
	o.LastUpdatedTime = &v
}

// GetLegalAddress returns the LegalAddress field value if set, zero value otherwise.
func (o *PatchBusiness) GetLegalAddress() Address {
	if o == nil || o.LegalAddress == nil {
		var ret Address
		return ret
	}
	return *o.LegalAddress
}

// GetLegalAddressOk returns a tuple with the LegalAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchBusiness) GetLegalAddressOk() (*Address, bool) {
	if o == nil || o.LegalAddress == nil {
		return nil, false
	}
	return o.LegalAddress, true
}

// HasLegalAddress returns a boolean if a field has been set.
func (o *PatchBusiness) HasLegalAddress() bool {
	if o != nil && o.LegalAddress != nil {
		return true
	}

	return false
}

// SetLegalAddress gets a reference to the given Address and assigns it to the LegalAddress field.
func (o *PatchBusiness) SetLegalAddress(v Address) {
	o.LegalAddress = &v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *PatchBusiness) GetMetadata() map[string]interface{} {
	if o == nil || o.Metadata == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchBusiness) GetMetadataOk() (map[string]interface{}, bool) {
	if o == nil || o.Metadata == nil {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *PatchBusiness) HasMetadata() bool {
	if o != nil && o.Metadata != nil {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given map[string]interface{} and assigns it to the Metadata field.
func (o *PatchBusiness) SetMetadata(v map[string]interface{}) {
	o.Metadata = v
}

// GetPhoneNumber returns the PhoneNumber field value if set, zero value otherwise.
func (o *PatchBusiness) GetPhoneNumber() string {
	if o == nil || o.PhoneNumber == nil {
		var ret string
		return ret
	}
	return *o.PhoneNumber
}

// GetPhoneNumberOk returns a tuple with the PhoneNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchBusiness) GetPhoneNumberOk() (*string, bool) {
	if o == nil || o.PhoneNumber == nil {
		return nil, false
	}
	return o.PhoneNumber, true
}

// HasPhoneNumber returns a boolean if a field has been set.
func (o *PatchBusiness) HasPhoneNumber() bool {
	if o != nil && o.PhoneNumber != nil {
		return true
	}

	return false
}

// SetPhoneNumber gets a reference to the given string and assigns it to the PhoneNumber field.
func (o *PatchBusiness) SetPhoneNumber(v string) {
	o.PhoneNumber = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *PatchBusiness) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchBusiness) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *PatchBusiness) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *PatchBusiness) SetStatus(v string) {
	o.Status = &v
}

// GetStructure returns the Structure field value if set, zero value otherwise.
func (o *PatchBusiness) GetStructure() string {
	if o == nil || o.Structure == nil {
		var ret string
		return ret
	}
	return *o.Structure
}

// GetStructureOk returns a tuple with the Structure field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchBusiness) GetStructureOk() (*string, bool) {
	if o == nil || o.Structure == nil {
		return nil, false
	}
	return o.Structure, true
}

// HasStructure returns a boolean if a field has been set.
func (o *PatchBusiness) HasStructure() bool {
	if o != nil && o.Structure != nil {
		return true
	}

	return false
}

// SetStructure gets a reference to the given string and assigns it to the Structure field.
func (o *PatchBusiness) SetStructure(v string) {
	o.Structure = &v
}

// GetTradeNames returns the TradeNames field value if set, zero value otherwise.
func (o *PatchBusiness) GetTradeNames() []string {
	if o == nil || o.TradeNames == nil {
		var ret []string
		return ret
	}
	return o.TradeNames
}

// GetTradeNamesOk returns a tuple with the TradeNames field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchBusiness) GetTradeNamesOk() ([]string, bool) {
	if o == nil || o.TradeNames == nil {
		return nil, false
	}
	return o.TradeNames, true
}

// HasTradeNames returns a boolean if a field has been set.
func (o *PatchBusiness) HasTradeNames() bool {
	if o != nil && o.TradeNames != nil {
		return true
	}

	return false
}

// SetTradeNames gets a reference to the given []string and assigns it to the TradeNames field.
func (o *PatchBusiness) SetTradeNames(v []string) {
	o.TradeNames = v
}

// GetVerificationLastRun returns the VerificationLastRun field value if set, zero value otherwise.
func (o *PatchBusiness) GetVerificationLastRun() time.Time {
	if o == nil || o.VerificationLastRun == nil {
		var ret time.Time
		return ret
	}
	return *o.VerificationLastRun
}

// GetVerificationLastRunOk returns a tuple with the VerificationLastRun field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchBusiness) GetVerificationLastRunOk() (*time.Time, bool) {
	if o == nil || o.VerificationLastRun == nil {
		return nil, false
	}
	return o.VerificationLastRun, true
}

// HasVerificationLastRun returns a boolean if a field has been set.
func (o *PatchBusiness) HasVerificationLastRun() bool {
	if o != nil && o.VerificationLastRun != nil {
		return true
	}

	return false
}

// SetVerificationLastRun gets a reference to the given time.Time and assigns it to the VerificationLastRun field.
func (o *PatchBusiness) SetVerificationLastRun(v time.Time) {
	o.VerificationLastRun = &v
}

// GetVerificationStatus returns the VerificationStatus field value if set, zero value otherwise.
func (o *PatchBusiness) GetVerificationStatus() VerificationStatus {
	if o == nil || o.VerificationStatus == nil {
		var ret VerificationStatus
		return ret
	}
	return *o.VerificationStatus
}

// GetVerificationStatusOk returns a tuple with the VerificationStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchBusiness) GetVerificationStatusOk() (*VerificationStatus, bool) {
	if o == nil || o.VerificationStatus == nil {
		return nil, false
	}
	return o.VerificationStatus, true
}

// HasVerificationStatus returns a boolean if a field has been set.
func (o *PatchBusiness) HasVerificationStatus() bool {
	if o != nil && o.VerificationStatus != nil {
		return true
	}

	return false
}

// SetVerificationStatus gets a reference to the given VerificationStatus and assigns it to the VerificationStatus field.
func (o *PatchBusiness) SetVerificationStatus(v VerificationStatus) {
	o.VerificationStatus = &v
}

// GetWebsite returns the Website field value if set, zero value otherwise.
func (o *PatchBusiness) GetWebsite() string {
	if o == nil || o.Website == nil {
		var ret string
		return ret
	}
	return *o.Website
}

// GetWebsiteOk returns a tuple with the Website field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchBusiness) GetWebsiteOk() (*string, bool) {
	if o == nil || o.Website == nil {
		return nil, false
	}
	return o.Website, true
}

// HasWebsite returns a boolean if a field has been set.
func (o *PatchBusiness) HasWebsite() bool {
	if o != nil && o.Website != nil {
		return true
	}

	return false
}

// SetWebsite gets a reference to the given string and assigns it to the Website field.
func (o *PatchBusiness) SetWebsite(v string) {
	o.Website = &v
}

func (o PatchBusiness) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.CreationTime != nil {
		toSerialize["creation_time"] = o.CreationTime
	}
	if o.Ein != nil {
		toSerialize["ein"] = o.Ein
	}
	if o.Email != nil {
		toSerialize["email"] = o.Email
	}
	if o.EntityName != nil {
		toSerialize["entity_name"] = o.EntityName
	}
	if o.FormationDate != nil {
		toSerialize["formation_date"] = o.FormationDate
	}
	if o.FormationState != nil {
		toSerialize["formation_state"] = o.FormationState
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.IsCustomer != nil {
		toSerialize["is_customer"] = o.IsCustomer
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
	if o.PhoneNumber != nil {
		toSerialize["phone_number"] = o.PhoneNumber
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	if o.Structure != nil {
		toSerialize["structure"] = o.Structure
	}
	if o.TradeNames != nil {
		toSerialize["trade_names"] = o.TradeNames
	}
	if o.VerificationLastRun != nil {
		toSerialize["verification_last_run"] = o.VerificationLastRun
	}
	if o.VerificationStatus != nil {
		toSerialize["verification_status"] = o.VerificationStatus
	}
	if o.Website != nil {
		toSerialize["website"] = o.Website
	}
	return json.Marshal(toSerialize)
}

type NullablePatchBusiness struct {
	value *PatchBusiness
	isSet bool
}

func (v NullablePatchBusiness) Get() *PatchBusiness {
	return v.value
}

func (v *NullablePatchBusiness) Set(val *PatchBusiness) {
	v.value = val
	v.isSet = true
}

func (v NullablePatchBusiness) IsSet() bool {
	return v.isSet
}

func (v *NullablePatchBusiness) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePatchBusiness(val *PatchBusiness) *NullablePatchBusiness {
	return &NullablePatchBusiness{value: val, isSet: true}
}

func (v NullablePatchBusiness) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePatchBusiness) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

