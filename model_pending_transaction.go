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

// PendingTransaction struct for PendingTransaction
type PendingTransaction struct {
	// The account id associated with the hold
	AccountId string `json:"account_id"`
	// The account number associated with the hold
	AccountNo string `json:"account_no"`
	// The creation date of the hold
	Created time.Time              `json:"created"`
	Data    PendingTransactionData `json:"data"`
	Id      int32                  `json:"id"`
	// The idempotency key used when initially creating this hold.
	Idemkey string `json:"idemkey"`
	// The tenant associated with this hold, in the form \"<bankid>_<partnerid>\"
	Tenant string `json:"tenant"`
	// The date the hold was last update
	Updated time.Time `json:"updated"`
	// The unique identifier of the hold transaction.
	Uuid string `json:"uuid"`
}

// NewPendingTransaction instantiates a new PendingTransaction object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPendingTransaction(accountId string, accountNo string, created time.Time, data PendingTransactionData, id int32, idemkey string, tenant string, updated time.Time, uuid string) *PendingTransaction {
	this := PendingTransaction{}
	this.AccountId = accountId
	this.AccountNo = accountNo
	this.Created = created
	this.Data = data
	this.Id = id
	this.Idemkey = idemkey
	this.Tenant = tenant
	this.Updated = updated
	this.Uuid = uuid
	return &this
}

// NewPendingTransactionWithDefaults instantiates a new PendingTransaction object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPendingTransactionWithDefaults() *PendingTransaction {
	this := PendingTransaction{}
	return &this
}

// GetAccountId returns the AccountId field value
func (o *PendingTransaction) GetAccountId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AccountId
}

// GetAccountIdOk returns a tuple with the AccountId field value
// and a boolean to check if the value has been set.
func (o *PendingTransaction) GetAccountIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AccountId, true
}

// SetAccountId sets field value
func (o *PendingTransaction) SetAccountId(v string) {
	o.AccountId = v
}

// GetAccountNo returns the AccountNo field value
func (o *PendingTransaction) GetAccountNo() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AccountNo
}

// GetAccountNoOk returns a tuple with the AccountNo field value
// and a boolean to check if the value has been set.
func (o *PendingTransaction) GetAccountNoOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AccountNo, true
}

// SetAccountNo sets field value
func (o *PendingTransaction) SetAccountNo(v string) {
	o.AccountNo = v
}

// GetCreated returns the Created field value
func (o *PendingTransaction) GetCreated() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.Created
}

// GetCreatedOk returns a tuple with the Created field value
// and a boolean to check if the value has been set.
func (o *PendingTransaction) GetCreatedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Created, true
}

// SetCreated sets field value
func (o *PendingTransaction) SetCreated(v time.Time) {
	o.Created = v
}

// GetData returns the Data field value
func (o *PendingTransaction) GetData() PendingTransactionData {
	if o == nil {
		var ret PendingTransactionData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *PendingTransaction) GetDataOk() (*PendingTransactionData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *PendingTransaction) SetData(v PendingTransactionData) {
	o.Data = v
}

// GetId returns the Id field value
func (o *PendingTransaction) GetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *PendingTransaction) GetIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *PendingTransaction) SetId(v int32) {
	o.Id = v
}

// GetIdemkey returns the Idemkey field value
func (o *PendingTransaction) GetIdemkey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Idemkey
}

// GetIdemkeyOk returns a tuple with the Idemkey field value
// and a boolean to check if the value has been set.
func (o *PendingTransaction) GetIdemkeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Idemkey, true
}

// SetIdemkey sets field value
func (o *PendingTransaction) SetIdemkey(v string) {
	o.Idemkey = v
}

// GetTenant returns the Tenant field value
func (o *PendingTransaction) GetTenant() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Tenant
}

// GetTenantOk returns a tuple with the Tenant field value
// and a boolean to check if the value has been set.
func (o *PendingTransaction) GetTenantOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Tenant, true
}

// SetTenant sets field value
func (o *PendingTransaction) SetTenant(v string) {
	o.Tenant = v
}

// GetUpdated returns the Updated field value
func (o *PendingTransaction) GetUpdated() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.Updated
}

// GetUpdatedOk returns a tuple with the Updated field value
// and a boolean to check if the value has been set.
func (o *PendingTransaction) GetUpdatedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Updated, true
}

// SetUpdated sets field value
func (o *PendingTransaction) SetUpdated(v time.Time) {
	o.Updated = v
}

// GetUuid returns the Uuid field value
func (o *PendingTransaction) GetUuid() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Uuid
}

// GetUuidOk returns a tuple with the Uuid field value
// and a boolean to check if the value has been set.
func (o *PendingTransaction) GetUuidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Uuid, true
}

// SetUuid sets field value
func (o *PendingTransaction) SetUuid(v string) {
	o.Uuid = v
}

func (o PendingTransaction) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["account_id"] = o.AccountId
	}
	if true {
		toSerialize["account_no"] = o.AccountNo
	}
	if true {
		toSerialize["created"] = o.Created
	}
	if true {
		toSerialize["data"] = o.Data
	}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["idemkey"] = o.Idemkey
	}
	if true {
		toSerialize["tenant"] = o.Tenant
	}
	if true {
		toSerialize["updated"] = o.Updated
	}
	if true {
		toSerialize["uuid"] = o.Uuid
	}
	return json.Marshal(toSerialize)
}

type NullablePendingTransaction struct {
	value *PendingTransaction
	isSet bool
}

func (v NullablePendingTransaction) Get() *PendingTransaction {
	return v.value
}

func (v *NullablePendingTransaction) Set(val *PendingTransaction) {
	v.value = val
	v.isSet = true
}

func (v NullablePendingTransaction) IsSet() bool {
	return v.isSet
}

func (v *NullablePendingTransaction) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePendingTransaction(val *PendingTransaction) *NullablePendingTransaction {
	return &NullablePendingTransaction{value: val, isSet: true}
}

func (v NullablePendingTransaction) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePendingTransaction) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
