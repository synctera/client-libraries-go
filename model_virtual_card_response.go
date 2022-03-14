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

// VirtualCardResponse struct for VirtualCardResponse
type VirtualCardResponse struct {
	// PHYSICAL or VIRTUAL.
	Form string `json:"form"`
	// The ID of the account to which the card will be linked
	AccountId string `json:"account_id"`
	// The bin number
	Bin       *string   `json:"bin,omitempty"`
	CardBrand CardBrand `json:"card_brand"`
	// The card product to which the card is attached
	CardProductId string `json:"card_product_id"`
	// The timestamp representing when the card issuance request was made
	CreationTime time.Time `json:"creation_time"`
	// The ID of the customer to whom the card will be issued
	CustomerId      string     `json:"customer_id"`
	EmbossName      EmbossName `json:"emboss_name"`
	ExpirationMonth string     `json:"expiration_month"`
	// The timestamp representing when the card would expire at
	ExpirationTime *time.Time `json:"expiration_time,omitempty"`
	ExpirationYear string     `json:"expiration_year"`
	// Card ID
	Id string `json:"id"`
	// The last 4 digits of the card PAN
	LastFour string `json:"last_four"`
	// The timestamp representing when the card was last modified at
	LastModifiedTime *time.Time `json:"last_modified_time,omitempty"`
	// Additional data to include in the request structured as key-value pairs
	Metadata *map[string]string `json:"metadata,omitempty"`
	// The reason the card needs to be reissued
	ReissueReason *string `json:"reissue_reason,omitempty"`
	// When reissuing a card, specify the card to be replaced here. When getting a card's details, if this card was issued as a reissuance of another card, this ID refers to the card was replaced.
	ReissuedFromId *string `json:"reissued_from_id,omitempty"`
	// If this card was reissued, this ID refers to the card that replaced it.
	ReissuedToId *string   `json:"reissued_to_id,omitempty"`
	Shipping     *Shipping `json:"shipping,omitempty"`
	// Indicates the type of card to be issued
	Type       string     `json:"type"`
	CardStatus CardStatus `json:"card_status"`
	// Additional details about the reason for the status change
	Memo         *string              `json:"memo,omitempty"`
	StatusReason CardStatusReasonCode `json:"status_reason"`
}

// NewVirtualCardResponse instantiates a new VirtualCardResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVirtualCardResponse(form string, accountId string, cardBrand CardBrand, cardProductId string, creationTime time.Time, customerId string, embossName EmbossName, expirationMonth string, expirationYear string, id string, lastFour string, type_ string, cardStatus CardStatus, statusReason CardStatusReasonCode) *VirtualCardResponse {
	this := VirtualCardResponse{}
	this.Form = form
	this.AccountId = accountId
	this.CardBrand = cardBrand
	this.CardProductId = cardProductId
	this.CreationTime = creationTime
	this.CustomerId = customerId
	this.EmbossName = embossName
	this.ExpirationMonth = expirationMonth
	this.ExpirationYear = expirationYear
	this.Id = id
	this.LastFour = lastFour
	this.Type = type_
	this.CardStatus = cardStatus
	this.StatusReason = statusReason
	return &this
}

// NewVirtualCardResponseWithDefaults instantiates a new VirtualCardResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVirtualCardResponseWithDefaults() *VirtualCardResponse {
	this := VirtualCardResponse{}
	return &this
}

// GetForm returns the Form field value
func (o *VirtualCardResponse) GetForm() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Form
}

// GetFormOk returns a tuple with the Form field value
// and a boolean to check if the value has been set.
func (o *VirtualCardResponse) GetFormOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Form, true
}

// SetForm sets field value
func (o *VirtualCardResponse) SetForm(v string) {
	o.Form = v
}

// GetAccountId returns the AccountId field value
func (o *VirtualCardResponse) GetAccountId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AccountId
}

// GetAccountIdOk returns a tuple with the AccountId field value
// and a boolean to check if the value has been set.
func (o *VirtualCardResponse) GetAccountIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AccountId, true
}

// SetAccountId sets field value
func (o *VirtualCardResponse) SetAccountId(v string) {
	o.AccountId = v
}

// GetBin returns the Bin field value if set, zero value otherwise.
func (o *VirtualCardResponse) GetBin() string {
	if o == nil || o.Bin == nil {
		var ret string
		return ret
	}
	return *o.Bin
}

// GetBinOk returns a tuple with the Bin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualCardResponse) GetBinOk() (*string, bool) {
	if o == nil || o.Bin == nil {
		return nil, false
	}
	return o.Bin, true
}

// HasBin returns a boolean if a field has been set.
func (o *VirtualCardResponse) HasBin() bool {
	if o != nil && o.Bin != nil {
		return true
	}

	return false
}

// SetBin gets a reference to the given string and assigns it to the Bin field.
func (o *VirtualCardResponse) SetBin(v string) {
	o.Bin = &v
}

// GetCardBrand returns the CardBrand field value
func (o *VirtualCardResponse) GetCardBrand() CardBrand {
	if o == nil {
		var ret CardBrand
		return ret
	}

	return o.CardBrand
}

// GetCardBrandOk returns a tuple with the CardBrand field value
// and a boolean to check if the value has been set.
func (o *VirtualCardResponse) GetCardBrandOk() (*CardBrand, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CardBrand, true
}

// SetCardBrand sets field value
func (o *VirtualCardResponse) SetCardBrand(v CardBrand) {
	o.CardBrand = v
}

// GetCardProductId returns the CardProductId field value
func (o *VirtualCardResponse) GetCardProductId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CardProductId
}

// GetCardProductIdOk returns a tuple with the CardProductId field value
// and a boolean to check if the value has been set.
func (o *VirtualCardResponse) GetCardProductIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CardProductId, true
}

// SetCardProductId sets field value
func (o *VirtualCardResponse) SetCardProductId(v string) {
	o.CardProductId = v
}

// GetCreationTime returns the CreationTime field value
func (o *VirtualCardResponse) GetCreationTime() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreationTime
}

// GetCreationTimeOk returns a tuple with the CreationTime field value
// and a boolean to check if the value has been set.
func (o *VirtualCardResponse) GetCreationTimeOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreationTime, true
}

// SetCreationTime sets field value
func (o *VirtualCardResponse) SetCreationTime(v time.Time) {
	o.CreationTime = v
}

// GetCustomerId returns the CustomerId field value
func (o *VirtualCardResponse) GetCustomerId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CustomerId
}

// GetCustomerIdOk returns a tuple with the CustomerId field value
// and a boolean to check if the value has been set.
func (o *VirtualCardResponse) GetCustomerIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CustomerId, true
}

// SetCustomerId sets field value
func (o *VirtualCardResponse) SetCustomerId(v string) {
	o.CustomerId = v
}

// GetEmbossName returns the EmbossName field value
func (o *VirtualCardResponse) GetEmbossName() EmbossName {
	if o == nil {
		var ret EmbossName
		return ret
	}

	return o.EmbossName
}

// GetEmbossNameOk returns a tuple with the EmbossName field value
// and a boolean to check if the value has been set.
func (o *VirtualCardResponse) GetEmbossNameOk() (*EmbossName, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EmbossName, true
}

// SetEmbossName sets field value
func (o *VirtualCardResponse) SetEmbossName(v EmbossName) {
	o.EmbossName = v
}

// GetExpirationMonth returns the ExpirationMonth field value
func (o *VirtualCardResponse) GetExpirationMonth() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ExpirationMonth
}

// GetExpirationMonthOk returns a tuple with the ExpirationMonth field value
// and a boolean to check if the value has been set.
func (o *VirtualCardResponse) GetExpirationMonthOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ExpirationMonth, true
}

// SetExpirationMonth sets field value
func (o *VirtualCardResponse) SetExpirationMonth(v string) {
	o.ExpirationMonth = v
}

// GetExpirationTime returns the ExpirationTime field value if set, zero value otherwise.
func (o *VirtualCardResponse) GetExpirationTime() time.Time {
	if o == nil || o.ExpirationTime == nil {
		var ret time.Time
		return ret
	}
	return *o.ExpirationTime
}

// GetExpirationTimeOk returns a tuple with the ExpirationTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualCardResponse) GetExpirationTimeOk() (*time.Time, bool) {
	if o == nil || o.ExpirationTime == nil {
		return nil, false
	}
	return o.ExpirationTime, true
}

// HasExpirationTime returns a boolean if a field has been set.
func (o *VirtualCardResponse) HasExpirationTime() bool {
	if o != nil && o.ExpirationTime != nil {
		return true
	}

	return false
}

// SetExpirationTime gets a reference to the given time.Time and assigns it to the ExpirationTime field.
func (o *VirtualCardResponse) SetExpirationTime(v time.Time) {
	o.ExpirationTime = &v
}

// GetExpirationYear returns the ExpirationYear field value
func (o *VirtualCardResponse) GetExpirationYear() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ExpirationYear
}

// GetExpirationYearOk returns a tuple with the ExpirationYear field value
// and a boolean to check if the value has been set.
func (o *VirtualCardResponse) GetExpirationYearOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ExpirationYear, true
}

// SetExpirationYear sets field value
func (o *VirtualCardResponse) SetExpirationYear(v string) {
	o.ExpirationYear = v
}

// GetId returns the Id field value
func (o *VirtualCardResponse) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *VirtualCardResponse) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *VirtualCardResponse) SetId(v string) {
	o.Id = v
}

// GetLastFour returns the LastFour field value
func (o *VirtualCardResponse) GetLastFour() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.LastFour
}

// GetLastFourOk returns a tuple with the LastFour field value
// and a boolean to check if the value has been set.
func (o *VirtualCardResponse) GetLastFourOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LastFour, true
}

// SetLastFour sets field value
func (o *VirtualCardResponse) SetLastFour(v string) {
	o.LastFour = v
}

// GetLastModifiedTime returns the LastModifiedTime field value if set, zero value otherwise.
func (o *VirtualCardResponse) GetLastModifiedTime() time.Time {
	if o == nil || o.LastModifiedTime == nil {
		var ret time.Time
		return ret
	}
	return *o.LastModifiedTime
}

// GetLastModifiedTimeOk returns a tuple with the LastModifiedTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualCardResponse) GetLastModifiedTimeOk() (*time.Time, bool) {
	if o == nil || o.LastModifiedTime == nil {
		return nil, false
	}
	return o.LastModifiedTime, true
}

// HasLastModifiedTime returns a boolean if a field has been set.
func (o *VirtualCardResponse) HasLastModifiedTime() bool {
	if o != nil && o.LastModifiedTime != nil {
		return true
	}

	return false
}

// SetLastModifiedTime gets a reference to the given time.Time and assigns it to the LastModifiedTime field.
func (o *VirtualCardResponse) SetLastModifiedTime(v time.Time) {
	o.LastModifiedTime = &v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *VirtualCardResponse) GetMetadata() map[string]string {
	if o == nil || o.Metadata == nil {
		var ret map[string]string
		return ret
	}
	return *o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualCardResponse) GetMetadataOk() (*map[string]string, bool) {
	if o == nil || o.Metadata == nil {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *VirtualCardResponse) HasMetadata() bool {
	if o != nil && o.Metadata != nil {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given map[string]string and assigns it to the Metadata field.
func (o *VirtualCardResponse) SetMetadata(v map[string]string) {
	o.Metadata = &v
}

// GetReissueReason returns the ReissueReason field value if set, zero value otherwise.
func (o *VirtualCardResponse) GetReissueReason() string {
	if o == nil || o.ReissueReason == nil {
		var ret string
		return ret
	}
	return *o.ReissueReason
}

// GetReissueReasonOk returns a tuple with the ReissueReason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualCardResponse) GetReissueReasonOk() (*string, bool) {
	if o == nil || o.ReissueReason == nil {
		return nil, false
	}
	return o.ReissueReason, true
}

// HasReissueReason returns a boolean if a field has been set.
func (o *VirtualCardResponse) HasReissueReason() bool {
	if o != nil && o.ReissueReason != nil {
		return true
	}

	return false
}

// SetReissueReason gets a reference to the given string and assigns it to the ReissueReason field.
func (o *VirtualCardResponse) SetReissueReason(v string) {
	o.ReissueReason = &v
}

// GetReissuedFromId returns the ReissuedFromId field value if set, zero value otherwise.
func (o *VirtualCardResponse) GetReissuedFromId() string {
	if o == nil || o.ReissuedFromId == nil {
		var ret string
		return ret
	}
	return *o.ReissuedFromId
}

// GetReissuedFromIdOk returns a tuple with the ReissuedFromId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualCardResponse) GetReissuedFromIdOk() (*string, bool) {
	if o == nil || o.ReissuedFromId == nil {
		return nil, false
	}
	return o.ReissuedFromId, true
}

// HasReissuedFromId returns a boolean if a field has been set.
func (o *VirtualCardResponse) HasReissuedFromId() bool {
	if o != nil && o.ReissuedFromId != nil {
		return true
	}

	return false
}

// SetReissuedFromId gets a reference to the given string and assigns it to the ReissuedFromId field.
func (o *VirtualCardResponse) SetReissuedFromId(v string) {
	o.ReissuedFromId = &v
}

// GetReissuedToId returns the ReissuedToId field value if set, zero value otherwise.
func (o *VirtualCardResponse) GetReissuedToId() string {
	if o == nil || o.ReissuedToId == nil {
		var ret string
		return ret
	}
	return *o.ReissuedToId
}

// GetReissuedToIdOk returns a tuple with the ReissuedToId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualCardResponse) GetReissuedToIdOk() (*string, bool) {
	if o == nil || o.ReissuedToId == nil {
		return nil, false
	}
	return o.ReissuedToId, true
}

// HasReissuedToId returns a boolean if a field has been set.
func (o *VirtualCardResponse) HasReissuedToId() bool {
	if o != nil && o.ReissuedToId != nil {
		return true
	}

	return false
}

// SetReissuedToId gets a reference to the given string and assigns it to the ReissuedToId field.
func (o *VirtualCardResponse) SetReissuedToId(v string) {
	o.ReissuedToId = &v
}

// GetShipping returns the Shipping field value if set, zero value otherwise.
func (o *VirtualCardResponse) GetShipping() Shipping {
	if o == nil || o.Shipping == nil {
		var ret Shipping
		return ret
	}
	return *o.Shipping
}

// GetShippingOk returns a tuple with the Shipping field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualCardResponse) GetShippingOk() (*Shipping, bool) {
	if o == nil || o.Shipping == nil {
		return nil, false
	}
	return o.Shipping, true
}

// HasShipping returns a boolean if a field has been set.
func (o *VirtualCardResponse) HasShipping() bool {
	if o != nil && o.Shipping != nil {
		return true
	}

	return false
}

// SetShipping gets a reference to the given Shipping and assigns it to the Shipping field.
func (o *VirtualCardResponse) SetShipping(v Shipping) {
	o.Shipping = &v
}

// GetType returns the Type field value
func (o *VirtualCardResponse) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *VirtualCardResponse) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *VirtualCardResponse) SetType(v string) {
	o.Type = v
}

// GetCardStatus returns the CardStatus field value
func (o *VirtualCardResponse) GetCardStatus() CardStatus {
	if o == nil {
		var ret CardStatus
		return ret
	}

	return o.CardStatus
}

// GetCardStatusOk returns a tuple with the CardStatus field value
// and a boolean to check if the value has been set.
func (o *VirtualCardResponse) GetCardStatusOk() (*CardStatus, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CardStatus, true
}

// SetCardStatus sets field value
func (o *VirtualCardResponse) SetCardStatus(v CardStatus) {
	o.CardStatus = v
}

// GetMemo returns the Memo field value if set, zero value otherwise.
func (o *VirtualCardResponse) GetMemo() string {
	if o == nil || o.Memo == nil {
		var ret string
		return ret
	}
	return *o.Memo
}

// GetMemoOk returns a tuple with the Memo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualCardResponse) GetMemoOk() (*string, bool) {
	if o == nil || o.Memo == nil {
		return nil, false
	}
	return o.Memo, true
}

// HasMemo returns a boolean if a field has been set.
func (o *VirtualCardResponse) HasMemo() bool {
	if o != nil && o.Memo != nil {
		return true
	}

	return false
}

// SetMemo gets a reference to the given string and assigns it to the Memo field.
func (o *VirtualCardResponse) SetMemo(v string) {
	o.Memo = &v
}

// GetStatusReason returns the StatusReason field value
func (o *VirtualCardResponse) GetStatusReason() CardStatusReasonCode {
	if o == nil {
		var ret CardStatusReasonCode
		return ret
	}

	return o.StatusReason
}

// GetStatusReasonOk returns a tuple with the StatusReason field value
// and a boolean to check if the value has been set.
func (o *VirtualCardResponse) GetStatusReasonOk() (*CardStatusReasonCode, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StatusReason, true
}

// SetStatusReason sets field value
func (o *VirtualCardResponse) SetStatusReason(v CardStatusReasonCode) {
	o.StatusReason = v
}

func (o VirtualCardResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["form"] = o.Form
	}
	if true {
		toSerialize["account_id"] = o.AccountId
	}
	if o.Bin != nil {
		toSerialize["bin"] = o.Bin
	}
	if true {
		toSerialize["card_brand"] = o.CardBrand
	}
	if true {
		toSerialize["card_product_id"] = o.CardProductId
	}
	if true {
		toSerialize["creation_time"] = o.CreationTime
	}
	if true {
		toSerialize["customer_id"] = o.CustomerId
	}
	if true {
		toSerialize["emboss_name"] = o.EmbossName
	}
	if true {
		toSerialize["expiration_month"] = o.ExpirationMonth
	}
	if o.ExpirationTime != nil {
		toSerialize["expiration_time"] = o.ExpirationTime
	}
	if true {
		toSerialize["expiration_year"] = o.ExpirationYear
	}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
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
	if true {
		toSerialize["type"] = o.Type
	}
	if true {
		toSerialize["card_status"] = o.CardStatus
	}
	if o.Memo != nil {
		toSerialize["memo"] = o.Memo
	}
	if true {
		toSerialize["status_reason"] = o.StatusReason
	}
	return json.Marshal(toSerialize)
}

type NullableVirtualCardResponse struct {
	value *VirtualCardResponse
	isSet bool
}

func (v NullableVirtualCardResponse) Get() *VirtualCardResponse {
	return v.value
}

func (v *NullableVirtualCardResponse) Set(val *VirtualCardResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableVirtualCardResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableVirtualCardResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVirtualCardResponse(val *VirtualCardResponse) *NullableVirtualCardResponse {
	return &NullableVirtualCardResponse{value: val, isSet: true}
}

func (v NullableVirtualCardResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVirtualCardResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
