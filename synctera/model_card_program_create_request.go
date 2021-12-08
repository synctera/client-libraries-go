/*
Synctera API

<h2>Let's build something great.</h2><p>Welcome to the official reference documentation for Synctera APIs. Our APIs are the best way to automate your company's banking needs and are designed to be easy to understand and implement.</p><p>We're continuously growing this library and what you see here is just the start, but if you need something specific or have a question, <a class='text-blue-600' href='https://synctera.com/contact' target='_blank' rel='noreferrer'>contact us</a>.</p> 

API version: 0.5.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package synctera

import (
	"encoding/json"
	"time"
)

// CardProgramCreateRequest struct for CardProgramCreateRequest
type CardProgramCreateRequest struct {
	// indicates whether program is active
	Active *bool `json:"active,omitempty"`
	// The ID of the bank partner works with within this program
	BankId int32 `json:"bank_id"`
	CardBrand CardBrand `json:"card_brand"`
	CardCategory CardCategory `json:"card_category"`
	CardProductType CardProductType `json:"card_product_type"`
	// The timestamp representing when the program was created
	CreationTime *time.Time `json:"creation_time,omitempty"`
	// The time when program became inactive
	EndDate *time.Time `json:"end_date,omitempty"`
	// Program ID
	Id *string `json:"id,omitempty"`
	// The timestamp representing when the program was last modified
	LastModifiedTime *time.Time `json:"last_modified_time,omitempty"`
	// Program name
	Name string `json:"name"`
	// The ID of the partner program belongs to
	PartnerId int32 `json:"partner_id"`
	// The time when program becomes active
	StartDate *time.Time `json:"start_date,omitempty"`
}

// NewCardProgramCreateRequest instantiates a new CardProgramCreateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCardProgramCreateRequest(bankId int32, cardBrand CardBrand, cardCategory CardCategory, cardProductType CardProductType, name string, partnerId int32) *CardProgramCreateRequest {
	this := CardProgramCreateRequest{}
	this.BankId = bankId
	this.CardBrand = cardBrand
	this.CardCategory = cardCategory
	this.CardProductType = cardProductType
	this.Name = name
	this.PartnerId = partnerId
	return &this
}

// NewCardProgramCreateRequestWithDefaults instantiates a new CardProgramCreateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCardProgramCreateRequestWithDefaults() *CardProgramCreateRequest {
	this := CardProgramCreateRequest{}
	return &this
}

// GetActive returns the Active field value if set, zero value otherwise.
func (o *CardProgramCreateRequest) GetActive() bool {
	if o == nil || o.Active == nil {
		var ret bool
		return ret
	}
	return *o.Active
}

// GetActiveOk returns a tuple with the Active field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CardProgramCreateRequest) GetActiveOk() (*bool, bool) {
	if o == nil || o.Active == nil {
		return nil, false
	}
	return o.Active, true
}

// HasActive returns a boolean if a field has been set.
func (o *CardProgramCreateRequest) HasActive() bool {
	if o != nil && o.Active != nil {
		return true
	}

	return false
}

// SetActive gets a reference to the given bool and assigns it to the Active field.
func (o *CardProgramCreateRequest) SetActive(v bool) {
	o.Active = &v
}

// GetBankId returns the BankId field value
func (o *CardProgramCreateRequest) GetBankId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.BankId
}

// GetBankIdOk returns a tuple with the BankId field value
// and a boolean to check if the value has been set.
func (o *CardProgramCreateRequest) GetBankIdOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.BankId, true
}

// SetBankId sets field value
func (o *CardProgramCreateRequest) SetBankId(v int32) {
	o.BankId = v
}

// GetCardBrand returns the CardBrand field value
func (o *CardProgramCreateRequest) GetCardBrand() CardBrand {
	if o == nil {
		var ret CardBrand
		return ret
	}

	return o.CardBrand
}

// GetCardBrandOk returns a tuple with the CardBrand field value
// and a boolean to check if the value has been set.
func (o *CardProgramCreateRequest) GetCardBrandOk() (*CardBrand, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.CardBrand, true
}

// SetCardBrand sets field value
func (o *CardProgramCreateRequest) SetCardBrand(v CardBrand) {
	o.CardBrand = v
}

// GetCardCategory returns the CardCategory field value
func (o *CardProgramCreateRequest) GetCardCategory() CardCategory {
	if o == nil {
		var ret CardCategory
		return ret
	}

	return o.CardCategory
}

// GetCardCategoryOk returns a tuple with the CardCategory field value
// and a boolean to check if the value has been set.
func (o *CardProgramCreateRequest) GetCardCategoryOk() (*CardCategory, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.CardCategory, true
}

// SetCardCategory sets field value
func (o *CardProgramCreateRequest) SetCardCategory(v CardCategory) {
	o.CardCategory = v
}

// GetCardProductType returns the CardProductType field value
func (o *CardProgramCreateRequest) GetCardProductType() CardProductType {
	if o == nil {
		var ret CardProductType
		return ret
	}

	return o.CardProductType
}

// GetCardProductTypeOk returns a tuple with the CardProductType field value
// and a boolean to check if the value has been set.
func (o *CardProgramCreateRequest) GetCardProductTypeOk() (*CardProductType, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.CardProductType, true
}

// SetCardProductType sets field value
func (o *CardProgramCreateRequest) SetCardProductType(v CardProductType) {
	o.CardProductType = v
}

// GetCreationTime returns the CreationTime field value if set, zero value otherwise.
func (o *CardProgramCreateRequest) GetCreationTime() time.Time {
	if o == nil || o.CreationTime == nil {
		var ret time.Time
		return ret
	}
	return *o.CreationTime
}

// GetCreationTimeOk returns a tuple with the CreationTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CardProgramCreateRequest) GetCreationTimeOk() (*time.Time, bool) {
	if o == nil || o.CreationTime == nil {
		return nil, false
	}
	return o.CreationTime, true
}

// HasCreationTime returns a boolean if a field has been set.
func (o *CardProgramCreateRequest) HasCreationTime() bool {
	if o != nil && o.CreationTime != nil {
		return true
	}

	return false
}

// SetCreationTime gets a reference to the given time.Time and assigns it to the CreationTime field.
func (o *CardProgramCreateRequest) SetCreationTime(v time.Time) {
	o.CreationTime = &v
}

// GetEndDate returns the EndDate field value if set, zero value otherwise.
func (o *CardProgramCreateRequest) GetEndDate() time.Time {
	if o == nil || o.EndDate == nil {
		var ret time.Time
		return ret
	}
	return *o.EndDate
}

// GetEndDateOk returns a tuple with the EndDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CardProgramCreateRequest) GetEndDateOk() (*time.Time, bool) {
	if o == nil || o.EndDate == nil {
		return nil, false
	}
	return o.EndDate, true
}

// HasEndDate returns a boolean if a field has been set.
func (o *CardProgramCreateRequest) HasEndDate() bool {
	if o != nil && o.EndDate != nil {
		return true
	}

	return false
}

// SetEndDate gets a reference to the given time.Time and assigns it to the EndDate field.
func (o *CardProgramCreateRequest) SetEndDate(v time.Time) {
	o.EndDate = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *CardProgramCreateRequest) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CardProgramCreateRequest) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *CardProgramCreateRequest) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *CardProgramCreateRequest) SetId(v string) {
	o.Id = &v
}

// GetLastModifiedTime returns the LastModifiedTime field value if set, zero value otherwise.
func (o *CardProgramCreateRequest) GetLastModifiedTime() time.Time {
	if o == nil || o.LastModifiedTime == nil {
		var ret time.Time
		return ret
	}
	return *o.LastModifiedTime
}

// GetLastModifiedTimeOk returns a tuple with the LastModifiedTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CardProgramCreateRequest) GetLastModifiedTimeOk() (*time.Time, bool) {
	if o == nil || o.LastModifiedTime == nil {
		return nil, false
	}
	return o.LastModifiedTime, true
}

// HasLastModifiedTime returns a boolean if a field has been set.
func (o *CardProgramCreateRequest) HasLastModifiedTime() bool {
	if o != nil && o.LastModifiedTime != nil {
		return true
	}

	return false
}

// SetLastModifiedTime gets a reference to the given time.Time and assigns it to the LastModifiedTime field.
func (o *CardProgramCreateRequest) SetLastModifiedTime(v time.Time) {
	o.LastModifiedTime = &v
}

// GetName returns the Name field value
func (o *CardProgramCreateRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CardProgramCreateRequest) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *CardProgramCreateRequest) SetName(v string) {
	o.Name = v
}

// GetPartnerId returns the PartnerId field value
func (o *CardProgramCreateRequest) GetPartnerId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.PartnerId
}

// GetPartnerIdOk returns a tuple with the PartnerId field value
// and a boolean to check if the value has been set.
func (o *CardProgramCreateRequest) GetPartnerIdOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.PartnerId, true
}

// SetPartnerId sets field value
func (o *CardProgramCreateRequest) SetPartnerId(v int32) {
	o.PartnerId = v
}

// GetStartDate returns the StartDate field value if set, zero value otherwise.
func (o *CardProgramCreateRequest) GetStartDate() time.Time {
	if o == nil || o.StartDate == nil {
		var ret time.Time
		return ret
	}
	return *o.StartDate
}

// GetStartDateOk returns a tuple with the StartDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CardProgramCreateRequest) GetStartDateOk() (*time.Time, bool) {
	if o == nil || o.StartDate == nil {
		return nil, false
	}
	return o.StartDate, true
}

// HasStartDate returns a boolean if a field has been set.
func (o *CardProgramCreateRequest) HasStartDate() bool {
	if o != nil && o.StartDate != nil {
		return true
	}

	return false
}

// SetStartDate gets a reference to the given time.Time and assigns it to the StartDate field.
func (o *CardProgramCreateRequest) SetStartDate(v time.Time) {
	o.StartDate = &v
}

func (o CardProgramCreateRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Active != nil {
		toSerialize["active"] = o.Active
	}
	if true {
		toSerialize["bank_id"] = o.BankId
	}
	if true {
		toSerialize["card_brand"] = o.CardBrand
	}
	if true {
		toSerialize["card_category"] = o.CardCategory
	}
	if true {
		toSerialize["card_product_type"] = o.CardProductType
	}
	if o.CreationTime != nil {
		toSerialize["creation_time"] = o.CreationTime
	}
	if o.EndDate != nil {
		toSerialize["end_date"] = o.EndDate
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.LastModifiedTime != nil {
		toSerialize["last_modified_time"] = o.LastModifiedTime
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["partner_id"] = o.PartnerId
	}
	if o.StartDate != nil {
		toSerialize["start_date"] = o.StartDate
	}
	return json.Marshal(toSerialize)
}

type NullableCardProgramCreateRequest struct {
	value *CardProgramCreateRequest
	isSet bool
}

func (v NullableCardProgramCreateRequest) Get() *CardProgramCreateRequest {
	return v.value
}

func (v *NullableCardProgramCreateRequest) Set(val *CardProgramCreateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCardProgramCreateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCardProgramCreateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCardProgramCreateRequest(val *CardProgramCreateRequest) *NullableCardProgramCreateRequest {
	return &NullableCardProgramCreateRequest{value: val, isSet: true}
}

func (v NullableCardProgramCreateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCardProgramCreateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


