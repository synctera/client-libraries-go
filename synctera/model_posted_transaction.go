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

// PostedTransaction struct for PostedTransaction
type PostedTransaction struct {
	// The creation date of the transaction
	Created time.Time `json:"created"`
	Data PostedTransactionData `json:"data"`
	// The \"effective date\" of a transaction. This may be earlier than posted_date in some cases (for example, a transaction that occurs on a Saturday may not be posted until the following Monday, but would have an effective date of Saturday)
	EffectiveDate time.Time `json:"effective_date"`
	Id int32 `json:"id"`
	// The idempotency key used when initially creating this transaction.
	Idemkey string `json:"idemkey"`
	// Whether or not this transaction represents a purely informational operation or an actual money movement
	InfoOnly bool `json:"info_only"`
	// Whether or not this transaction was created operating in \"lead ledger\" mode
	LeadMode bool `json:"lead_mode"`
	// The date the transaction was posted. This is the date any money is considered to be added or removed from an account.
	PostedDate time.Time `json:"posted_date"`
	Status string `json:"status"`
	// The specific transaction type. For example, for `ach`, this may be \"outgoing_debit\".
	Subtype string `json:"subtype"`
	// The tenant associated with this transaction, in the form \"<bankid>_<partnerid>\"
	Tenant string `json:"tenant"`
	// The general type of transaction. For example, \"card\" or \"ach\".
	Type string `json:"type"`
	// The date the transaction was last updated
	Updated time.Time `json:"updated"`
	// The unique identifier of the transaction.
	Uuid string `json:"uuid"`
}

// NewPostedTransaction instantiates a new PostedTransaction object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPostedTransaction(created time.Time, data PostedTransactionData, effectiveDate time.Time, id int32, idemkey string, infoOnly bool, leadMode bool, postedDate time.Time, status string, subtype string, tenant string, type_ string, updated time.Time, uuid string) *PostedTransaction {
	this := PostedTransaction{}
	this.Created = created
	this.Data = data
	this.EffectiveDate = effectiveDate
	this.Id = id
	this.Idemkey = idemkey
	this.InfoOnly = infoOnly
	this.LeadMode = leadMode
	this.PostedDate = postedDate
	this.Status = status
	this.Subtype = subtype
	this.Tenant = tenant
	this.Type = type_
	this.Updated = updated
	this.Uuid = uuid
	return &this
}

// NewPostedTransactionWithDefaults instantiates a new PostedTransaction object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPostedTransactionWithDefaults() *PostedTransaction {
	this := PostedTransaction{}
	return &this
}

// GetCreated returns the Created field value
func (o *PostedTransaction) GetCreated() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.Created
}

// GetCreatedOk returns a tuple with the Created field value
// and a boolean to check if the value has been set.
func (o *PostedTransaction) GetCreatedOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Created, true
}

// SetCreated sets field value
func (o *PostedTransaction) SetCreated(v time.Time) {
	o.Created = v
}

// GetData returns the Data field value
func (o *PostedTransaction) GetData() PostedTransactionData {
	if o == nil {
		var ret PostedTransactionData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *PostedTransaction) GetDataOk() (*PostedTransactionData, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *PostedTransaction) SetData(v PostedTransactionData) {
	o.Data = v
}

// GetEffectiveDate returns the EffectiveDate field value
func (o *PostedTransaction) GetEffectiveDate() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.EffectiveDate
}

// GetEffectiveDateOk returns a tuple with the EffectiveDate field value
// and a boolean to check if the value has been set.
func (o *PostedTransaction) GetEffectiveDateOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.EffectiveDate, true
}

// SetEffectiveDate sets field value
func (o *PostedTransaction) SetEffectiveDate(v time.Time) {
	o.EffectiveDate = v
}

// GetId returns the Id field value
func (o *PostedTransaction) GetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *PostedTransaction) GetIdOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *PostedTransaction) SetId(v int32) {
	o.Id = v
}

// GetIdemkey returns the Idemkey field value
func (o *PostedTransaction) GetIdemkey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Idemkey
}

// GetIdemkeyOk returns a tuple with the Idemkey field value
// and a boolean to check if the value has been set.
func (o *PostedTransaction) GetIdemkeyOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Idemkey, true
}

// SetIdemkey sets field value
func (o *PostedTransaction) SetIdemkey(v string) {
	o.Idemkey = v
}

// GetInfoOnly returns the InfoOnly field value
func (o *PostedTransaction) GetInfoOnly() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.InfoOnly
}

// GetInfoOnlyOk returns a tuple with the InfoOnly field value
// and a boolean to check if the value has been set.
func (o *PostedTransaction) GetInfoOnlyOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.InfoOnly, true
}

// SetInfoOnly sets field value
func (o *PostedTransaction) SetInfoOnly(v bool) {
	o.InfoOnly = v
}

// GetLeadMode returns the LeadMode field value
func (o *PostedTransaction) GetLeadMode() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.LeadMode
}

// GetLeadModeOk returns a tuple with the LeadMode field value
// and a boolean to check if the value has been set.
func (o *PostedTransaction) GetLeadModeOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.LeadMode, true
}

// SetLeadMode sets field value
func (o *PostedTransaction) SetLeadMode(v bool) {
	o.LeadMode = v
}

// GetPostedDate returns the PostedDate field value
func (o *PostedTransaction) GetPostedDate() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.PostedDate
}

// GetPostedDateOk returns a tuple with the PostedDate field value
// and a boolean to check if the value has been set.
func (o *PostedTransaction) GetPostedDateOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.PostedDate, true
}

// SetPostedDate sets field value
func (o *PostedTransaction) SetPostedDate(v time.Time) {
	o.PostedDate = v
}

// GetStatus returns the Status field value
func (o *PostedTransaction) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *PostedTransaction) GetStatusOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *PostedTransaction) SetStatus(v string) {
	o.Status = v
}

// GetSubtype returns the Subtype field value
func (o *PostedTransaction) GetSubtype() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Subtype
}

// GetSubtypeOk returns a tuple with the Subtype field value
// and a boolean to check if the value has been set.
func (o *PostedTransaction) GetSubtypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Subtype, true
}

// SetSubtype sets field value
func (o *PostedTransaction) SetSubtype(v string) {
	o.Subtype = v
}

// GetTenant returns the Tenant field value
func (o *PostedTransaction) GetTenant() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Tenant
}

// GetTenantOk returns a tuple with the Tenant field value
// and a boolean to check if the value has been set.
func (o *PostedTransaction) GetTenantOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Tenant, true
}

// SetTenant sets field value
func (o *PostedTransaction) SetTenant(v string) {
	o.Tenant = v
}

// GetType returns the Type field value
func (o *PostedTransaction) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *PostedTransaction) GetTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *PostedTransaction) SetType(v string) {
	o.Type = v
}

// GetUpdated returns the Updated field value
func (o *PostedTransaction) GetUpdated() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.Updated
}

// GetUpdatedOk returns a tuple with the Updated field value
// and a boolean to check if the value has been set.
func (o *PostedTransaction) GetUpdatedOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Updated, true
}

// SetUpdated sets field value
func (o *PostedTransaction) SetUpdated(v time.Time) {
	o.Updated = v
}

// GetUuid returns the Uuid field value
func (o *PostedTransaction) GetUuid() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Uuid
}

// GetUuidOk returns a tuple with the Uuid field value
// and a boolean to check if the value has been set.
func (o *PostedTransaction) GetUuidOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Uuid, true
}

// SetUuid sets field value
func (o *PostedTransaction) SetUuid(v string) {
	o.Uuid = v
}

func (o PostedTransaction) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["created"] = o.Created
	}
	if true {
		toSerialize["data"] = o.Data
	}
	if true {
		toSerialize["effective_date"] = o.EffectiveDate
	}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["idemkey"] = o.Idemkey
	}
	if true {
		toSerialize["info_only"] = o.InfoOnly
	}
	if true {
		toSerialize["lead_mode"] = o.LeadMode
	}
	if true {
		toSerialize["posted_date"] = o.PostedDate
	}
	if true {
		toSerialize["status"] = o.Status
	}
	if true {
		toSerialize["subtype"] = o.Subtype
	}
	if true {
		toSerialize["tenant"] = o.Tenant
	}
	if true {
		toSerialize["type"] = o.Type
	}
	if true {
		toSerialize["updated"] = o.Updated
	}
	if true {
		toSerialize["uuid"] = o.Uuid
	}
	return json.Marshal(toSerialize)
}

type NullablePostedTransaction struct {
	value *PostedTransaction
	isSet bool
}

func (v NullablePostedTransaction) Get() *PostedTransaction {
	return v.value
}

func (v *NullablePostedTransaction) Set(val *PostedTransaction) {
	v.value = val
	v.isSet = true
}

func (v NullablePostedTransaction) IsSet() bool {
	return v.isSet
}

func (v *NullablePostedTransaction) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePostedTransaction(val *PostedTransaction) *NullablePostedTransaction {
	return &NullablePostedTransaction{value: val, isSet: true}
}

func (v NullablePostedTransaction) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePostedTransaction) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


