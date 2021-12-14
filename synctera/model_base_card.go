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

// BaseCard struct for BaseCard
type BaseCard struct {
	// PHYSICAL or VIRTUAL.
	Form string `json:"form"`
	// The ID of the account to which the card will be linked
	AccountId *string `json:"account_id,omitempty"`
	CardBrand *CardBrand `json:"card_brand,omitempty"`
	// The card product to which the card is attached
	CardProductId *string `json:"card_product_id,omitempty"`
	// The timestamp representing when the card issuance request was made
	CreationTime *time.Time `json:"creation_time,omitempty"`
	// The ID of the customer to whom the card will be issued
	CustomerId *string `json:"customer_id,omitempty"`
	EmbossName *EmbossName `json:"emboss_name,omitempty"`
	ExpirationMonth *string `json:"expiration_month,omitempty"`
	// The timestamp representing when the card would expire at
	ExpirationTime *time.Time `json:"expiration_time,omitempty"`
	ExpirationYear *string `json:"expiration_year,omitempty"`
	// Card ID
	Id *string `json:"id,omitempty"`
	// The last 4 digits of the card PAN
	LastFour *string `json:"last_four,omitempty"`
	// The timestamp representing when the card was last modified at
	LastModifiedTime *time.Time `json:"last_modified_time,omitempty"`
	// Additional data to include in the request structured as key-value pairs
	Metadata *map[string]string `json:"metadata,omitempty"`
	// The reason the card needs to be reissued
	ReissueReason *string `json:"reissue_reason,omitempty"`
	// When reissuing a card, specify the card to be replaced here. When getting a card's details, if this card was issued as a reissuance of another card, this ID refers to the card was replaced. 
	ReissuedFromId *string `json:"reissued_from_id,omitempty"`
	// If this card was reissued, this ID refers to the card that replaced it.
	ReissuedToId *string `json:"reissued_to_id,omitempty"`
	Shipping *Shipping `json:"shipping,omitempty"`
	// Indicates the type of card to be issued
	Type *string `json:"type,omitempty"`
}

// NewBaseCard instantiates a new BaseCard object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBaseCard(form string) *BaseCard {
	this := BaseCard{}
	this.Form = form
	return &this
}

// NewBaseCardWithDefaults instantiates a new BaseCard object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBaseCardWithDefaults() *BaseCard {
	this := BaseCard{}
	return &this
}

// GetForm returns the Form field value
func (o *BaseCard) GetForm() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Form
}

// GetFormOk returns a tuple with the Form field value
// and a boolean to check if the value has been set.
func (o *BaseCard) GetFormOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Form, true
}

// SetForm sets field value
func (o *BaseCard) SetForm(v string) {
	o.Form = v
}

// GetAccountId returns the AccountId field value if set, zero value otherwise.
func (o *BaseCard) GetAccountId() string {
	if o == nil || o.AccountId == nil {
		var ret string
		return ret
	}
	return *o.AccountId
}

// GetAccountIdOk returns a tuple with the AccountId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseCard) GetAccountIdOk() (*string, bool) {
	if o == nil || o.AccountId == nil {
		return nil, false
	}
	return o.AccountId, true
}

// HasAccountId returns a boolean if a field has been set.
func (o *BaseCard) HasAccountId() bool {
	if o != nil && o.AccountId != nil {
		return true
	}

	return false
}

// SetAccountId gets a reference to the given string and assigns it to the AccountId field.
func (o *BaseCard) SetAccountId(v string) {
	o.AccountId = &v
}

// GetCardBrand returns the CardBrand field value if set, zero value otherwise.
func (o *BaseCard) GetCardBrand() CardBrand {
	if o == nil || o.CardBrand == nil {
		var ret CardBrand
		return ret
	}
	return *o.CardBrand
}

// GetCardBrandOk returns a tuple with the CardBrand field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseCard) GetCardBrandOk() (*CardBrand, bool) {
	if o == nil || o.CardBrand == nil {
		return nil, false
	}
	return o.CardBrand, true
}

// HasCardBrand returns a boolean if a field has been set.
func (o *BaseCard) HasCardBrand() bool {
	if o != nil && o.CardBrand != nil {
		return true
	}

	return false
}

// SetCardBrand gets a reference to the given CardBrand and assigns it to the CardBrand field.
func (o *BaseCard) SetCardBrand(v CardBrand) {
	o.CardBrand = &v
}

// GetCardProductId returns the CardProductId field value if set, zero value otherwise.
func (o *BaseCard) GetCardProductId() string {
	if o == nil || o.CardProductId == nil {
		var ret string
		return ret
	}
	return *o.CardProductId
}

// GetCardProductIdOk returns a tuple with the CardProductId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseCard) GetCardProductIdOk() (*string, bool) {
	if o == nil || o.CardProductId == nil {
		return nil, false
	}
	return o.CardProductId, true
}

// HasCardProductId returns a boolean if a field has been set.
func (o *BaseCard) HasCardProductId() bool {
	if o != nil && o.CardProductId != nil {
		return true
	}

	return false
}

// SetCardProductId gets a reference to the given string and assigns it to the CardProductId field.
func (o *BaseCard) SetCardProductId(v string) {
	o.CardProductId = &v
}

// GetCreationTime returns the CreationTime field value if set, zero value otherwise.
func (o *BaseCard) GetCreationTime() time.Time {
	if o == nil || o.CreationTime == nil {
		var ret time.Time
		return ret
	}
	return *o.CreationTime
}

// GetCreationTimeOk returns a tuple with the CreationTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseCard) GetCreationTimeOk() (*time.Time, bool) {
	if o == nil || o.CreationTime == nil {
		return nil, false
	}
	return o.CreationTime, true
}

// HasCreationTime returns a boolean if a field has been set.
func (o *BaseCard) HasCreationTime() bool {
	if o != nil && o.CreationTime != nil {
		return true
	}

	return false
}

// SetCreationTime gets a reference to the given time.Time and assigns it to the CreationTime field.
func (o *BaseCard) SetCreationTime(v time.Time) {
	o.CreationTime = &v
}

// GetCustomerId returns the CustomerId field value if set, zero value otherwise.
func (o *BaseCard) GetCustomerId() string {
	if o == nil || o.CustomerId == nil {
		var ret string
		return ret
	}
	return *o.CustomerId
}

// GetCustomerIdOk returns a tuple with the CustomerId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseCard) GetCustomerIdOk() (*string, bool) {
	if o == nil || o.CustomerId == nil {
		return nil, false
	}
	return o.CustomerId, true
}

// HasCustomerId returns a boolean if a field has been set.
func (o *BaseCard) HasCustomerId() bool {
	if o != nil && o.CustomerId != nil {
		return true
	}

	return false
}

// SetCustomerId gets a reference to the given string and assigns it to the CustomerId field.
func (o *BaseCard) SetCustomerId(v string) {
	o.CustomerId = &v
}

// GetEmbossName returns the EmbossName field value if set, zero value otherwise.
func (o *BaseCard) GetEmbossName() EmbossName {
	if o == nil || o.EmbossName == nil {
		var ret EmbossName
		return ret
	}
	return *o.EmbossName
}

// GetEmbossNameOk returns a tuple with the EmbossName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseCard) GetEmbossNameOk() (*EmbossName, bool) {
	if o == nil || o.EmbossName == nil {
		return nil, false
	}
	return o.EmbossName, true
}

// HasEmbossName returns a boolean if a field has been set.
func (o *BaseCard) HasEmbossName() bool {
	if o != nil && o.EmbossName != nil {
		return true
	}

	return false
}

// SetEmbossName gets a reference to the given EmbossName and assigns it to the EmbossName field.
func (o *BaseCard) SetEmbossName(v EmbossName) {
	o.EmbossName = &v
}

// GetExpirationMonth returns the ExpirationMonth field value if set, zero value otherwise.
func (o *BaseCard) GetExpirationMonth() string {
	if o == nil || o.ExpirationMonth == nil {
		var ret string
		return ret
	}
	return *o.ExpirationMonth
}

// GetExpirationMonthOk returns a tuple with the ExpirationMonth field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseCard) GetExpirationMonthOk() (*string, bool) {
	if o == nil || o.ExpirationMonth == nil {
		return nil, false
	}
	return o.ExpirationMonth, true
}

// HasExpirationMonth returns a boolean if a field has been set.
func (o *BaseCard) HasExpirationMonth() bool {
	if o != nil && o.ExpirationMonth != nil {
		return true
	}

	return false
}

// SetExpirationMonth gets a reference to the given string and assigns it to the ExpirationMonth field.
func (o *BaseCard) SetExpirationMonth(v string) {
	o.ExpirationMonth = &v
}

// GetExpirationTime returns the ExpirationTime field value if set, zero value otherwise.
func (o *BaseCard) GetExpirationTime() time.Time {
	if o == nil || o.ExpirationTime == nil {
		var ret time.Time
		return ret
	}
	return *o.ExpirationTime
}

// GetExpirationTimeOk returns a tuple with the ExpirationTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseCard) GetExpirationTimeOk() (*time.Time, bool) {
	if o == nil || o.ExpirationTime == nil {
		return nil, false
	}
	return o.ExpirationTime, true
}

// HasExpirationTime returns a boolean if a field has been set.
func (o *BaseCard) HasExpirationTime() bool {
	if o != nil && o.ExpirationTime != nil {
		return true
	}

	return false
}

// SetExpirationTime gets a reference to the given time.Time and assigns it to the ExpirationTime field.
func (o *BaseCard) SetExpirationTime(v time.Time) {
	o.ExpirationTime = &v
}

// GetExpirationYear returns the ExpirationYear field value if set, zero value otherwise.
func (o *BaseCard) GetExpirationYear() string {
	if o == nil || o.ExpirationYear == nil {
		var ret string
		return ret
	}
	return *o.ExpirationYear
}

// GetExpirationYearOk returns a tuple with the ExpirationYear field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseCard) GetExpirationYearOk() (*string, bool) {
	if o == nil || o.ExpirationYear == nil {
		return nil, false
	}
	return o.ExpirationYear, true
}

// HasExpirationYear returns a boolean if a field has been set.
func (o *BaseCard) HasExpirationYear() bool {
	if o != nil && o.ExpirationYear != nil {
		return true
	}

	return false
}

// SetExpirationYear gets a reference to the given string and assigns it to the ExpirationYear field.
func (o *BaseCard) SetExpirationYear(v string) {
	o.ExpirationYear = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *BaseCard) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseCard) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *BaseCard) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *BaseCard) SetId(v string) {
	o.Id = &v
}

// GetLastFour returns the LastFour field value if set, zero value otherwise.
func (o *BaseCard) GetLastFour() string {
	if o == nil || o.LastFour == nil {
		var ret string
		return ret
	}
	return *o.LastFour
}

// GetLastFourOk returns a tuple with the LastFour field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseCard) GetLastFourOk() (*string, bool) {
	if o == nil || o.LastFour == nil {
		return nil, false
	}
	return o.LastFour, true
}

// HasLastFour returns a boolean if a field has been set.
func (o *BaseCard) HasLastFour() bool {
	if o != nil && o.LastFour != nil {
		return true
	}

	return false
}

// SetLastFour gets a reference to the given string and assigns it to the LastFour field.
func (o *BaseCard) SetLastFour(v string) {
	o.LastFour = &v
}

// GetLastModifiedTime returns the LastModifiedTime field value if set, zero value otherwise.
func (o *BaseCard) GetLastModifiedTime() time.Time {
	if o == nil || o.LastModifiedTime == nil {
		var ret time.Time
		return ret
	}
	return *o.LastModifiedTime
}

// GetLastModifiedTimeOk returns a tuple with the LastModifiedTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseCard) GetLastModifiedTimeOk() (*time.Time, bool) {
	if o == nil || o.LastModifiedTime == nil {
		return nil, false
	}
	return o.LastModifiedTime, true
}

// HasLastModifiedTime returns a boolean if a field has been set.
func (o *BaseCard) HasLastModifiedTime() bool {
	if o != nil && o.LastModifiedTime != nil {
		return true
	}

	return false
}

// SetLastModifiedTime gets a reference to the given time.Time and assigns it to the LastModifiedTime field.
func (o *BaseCard) SetLastModifiedTime(v time.Time) {
	o.LastModifiedTime = &v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *BaseCard) GetMetadata() map[string]string {
	if o == nil || o.Metadata == nil {
		var ret map[string]string
		return ret
	}
	return *o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseCard) GetMetadataOk() (*map[string]string, bool) {
	if o == nil || o.Metadata == nil {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *BaseCard) HasMetadata() bool {
	if o != nil && o.Metadata != nil {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given map[string]string and assigns it to the Metadata field.
func (o *BaseCard) SetMetadata(v map[string]string) {
	o.Metadata = &v
}

// GetReissueReason returns the ReissueReason field value if set, zero value otherwise.
func (o *BaseCard) GetReissueReason() string {
	if o == nil || o.ReissueReason == nil {
		var ret string
		return ret
	}
	return *o.ReissueReason
}

// GetReissueReasonOk returns a tuple with the ReissueReason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseCard) GetReissueReasonOk() (*string, bool) {
	if o == nil || o.ReissueReason == nil {
		return nil, false
	}
	return o.ReissueReason, true
}

// HasReissueReason returns a boolean if a field has been set.
func (o *BaseCard) HasReissueReason() bool {
	if o != nil && o.ReissueReason != nil {
		return true
	}

	return false
}

// SetReissueReason gets a reference to the given string and assigns it to the ReissueReason field.
func (o *BaseCard) SetReissueReason(v string) {
	o.ReissueReason = &v
}

// GetReissuedFromId returns the ReissuedFromId field value if set, zero value otherwise.
func (o *BaseCard) GetReissuedFromId() string {
	if o == nil || o.ReissuedFromId == nil {
		var ret string
		return ret
	}
	return *o.ReissuedFromId
}

// GetReissuedFromIdOk returns a tuple with the ReissuedFromId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseCard) GetReissuedFromIdOk() (*string, bool) {
	if o == nil || o.ReissuedFromId == nil {
		return nil, false
	}
	return o.ReissuedFromId, true
}

// HasReissuedFromId returns a boolean if a field has been set.
func (o *BaseCard) HasReissuedFromId() bool {
	if o != nil && o.ReissuedFromId != nil {
		return true
	}

	return false
}

// SetReissuedFromId gets a reference to the given string and assigns it to the ReissuedFromId field.
func (o *BaseCard) SetReissuedFromId(v string) {
	o.ReissuedFromId = &v
}

// GetReissuedToId returns the ReissuedToId field value if set, zero value otherwise.
func (o *BaseCard) GetReissuedToId() string {
	if o == nil || o.ReissuedToId == nil {
		var ret string
		return ret
	}
	return *o.ReissuedToId
}

// GetReissuedToIdOk returns a tuple with the ReissuedToId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseCard) GetReissuedToIdOk() (*string, bool) {
	if o == nil || o.ReissuedToId == nil {
		return nil, false
	}
	return o.ReissuedToId, true
}

// HasReissuedToId returns a boolean if a field has been set.
func (o *BaseCard) HasReissuedToId() bool {
	if o != nil && o.ReissuedToId != nil {
		return true
	}

	return false
}

// SetReissuedToId gets a reference to the given string and assigns it to the ReissuedToId field.
func (o *BaseCard) SetReissuedToId(v string) {
	o.ReissuedToId = &v
}

// GetShipping returns the Shipping field value if set, zero value otherwise.
func (o *BaseCard) GetShipping() Shipping {
	if o == nil || o.Shipping == nil {
		var ret Shipping
		return ret
	}
	return *o.Shipping
}

// GetShippingOk returns a tuple with the Shipping field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseCard) GetShippingOk() (*Shipping, bool) {
	if o == nil || o.Shipping == nil {
		return nil, false
	}
	return o.Shipping, true
}

// HasShipping returns a boolean if a field has been set.
func (o *BaseCard) HasShipping() bool {
	if o != nil && o.Shipping != nil {
		return true
	}

	return false
}

// SetShipping gets a reference to the given Shipping and assigns it to the Shipping field.
func (o *BaseCard) SetShipping(v Shipping) {
	o.Shipping = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *BaseCard) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseCard) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *BaseCard) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *BaseCard) SetType(v string) {
	o.Type = &v
}

func (o BaseCard) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["form"] = o.Form
	}
	if o.AccountId != nil {
		toSerialize["account_id"] = o.AccountId
	}
	if o.CardBrand != nil {
		toSerialize["card_brand"] = o.CardBrand
	}
	if o.CardProductId != nil {
		toSerialize["card_product_id"] = o.CardProductId
	}
	if o.CreationTime != nil {
		toSerialize["creation_time"] = o.CreationTime
	}
	if o.CustomerId != nil {
		toSerialize["customer_id"] = o.CustomerId
	}
	if o.EmbossName != nil {
		toSerialize["emboss_name"] = o.EmbossName
	}
	if o.ExpirationMonth != nil {
		toSerialize["expiration_month"] = o.ExpirationMonth
	}
	if o.ExpirationTime != nil {
		toSerialize["expiration_time"] = o.ExpirationTime
	}
	if o.ExpirationYear != nil {
		toSerialize["expiration_year"] = o.ExpirationYear
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.LastFour != nil {
		toSerialize["last_four"] = o.LastFour
	}
	if o.LastModifiedTime != nil {
		toSerialize["last_modified_time"] = o.LastModifiedTime
	}
	if o.Metadata != nil {
		toSerialize["metadata"] = o.Metadata
	}
	if o.ReissueReason != nil {
		toSerialize["reissue_reason"] = o.ReissueReason
	}
	if o.ReissuedFromId != nil {
		toSerialize["reissued_from_id"] = o.ReissuedFromId
	}
	if o.ReissuedToId != nil {
		toSerialize["reissued_to_id"] = o.ReissuedToId
	}
	if o.Shipping != nil {
		toSerialize["shipping"] = o.Shipping
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	return json.Marshal(toSerialize)
}

type NullableBaseCard struct {
	value *BaseCard
	isSet bool
}

func (v NullableBaseCard) Get() *BaseCard {
	return v.value
}

func (v *NullableBaseCard) Set(val *BaseCard) {
	v.value = val
	v.isSet = true
}

func (v NullableBaseCard) IsSet() bool {
	return v.isSet
}

func (v *NullableBaseCard) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBaseCard(val *BaseCard) *NullableBaseCard {
	return &NullableBaseCard{value: val, isSet: true}
}

func (v NullableBaseCard) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBaseCard) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


