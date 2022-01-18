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

// AccountCreation struct for AccountCreation
type AccountCreation struct {
	AccessStatus *AccountAccessStatus `json:"access_status,omitempty"`
	// Account number
	AccountNumber *string `json:"account_number,omitempty"`
	// Purpose of the account
	AccountPurpose *string `json:"account_purpose,omitempty"`
	AccountType *AccountType `json:"account_type,omitempty"`
	// A list of balances for account based on different type
	Balances *[]Balance `json:"balances,omitempty"`
	// Bank routing number
	BankRouting *string `json:"bank_routing,omitempty"`
	// Account creation timestamp in RFC3337 format
	CreationTime *time.Time `json:"creation_time,omitempty"`
	// Account currency or account settlement currency. ISO 4217 alphabetic currency code. Default USD
	Currency *string `json:"currency,omitempty"`
	// A list of the customer IDs of the account holders.
	CustomerIds *[]string `json:"customer_ids,omitempty"`
	// Exchange rate type
	ExchangeRateType *string `json:"exchange_rate_type,omitempty"`
	// A list of fee resources from account product that the current account associate with
	FeeProductIds *[]string `json:"fee_product_ids,omitempty"`
	// International bank account number
	Iban *string `json:"iban,omitempty"`
	// Account ID
	Id *string `json:"id,omitempty"`
	// An interest from account product that the current account associate with
	InterestProductId *string `json:"interest_product_id,omitempty"`
	// Account is investment (variable balance) account or a multi-balance account pool. Default false
	IsAccountPool *bool `json:"is_account_pool,omitempty"`
	// Timestamp of the last account modification in RFC3337 format
	LastUpdatedTime *time.Time `json:"last_updated_time,omitempty"`
	// Account's overdraft limit
	OverdraftLimit *int64 `json:"overdraft_limit,omitempty"`
	SpendingLimits *SpendingLimits `json:"spending_limits,omitempty"`
	Status *Status `json:"status,omitempty"`
	// SWIFT code
	SwiftCode *string `json:"swift_code,omitempty"`
	// Account template ID
	AccountTemplateId *string `json:"account_template_id,omitempty"`
	// List of the relationship for this account to the parties
	Relationships *[]Relationship `json:"relationships,omitempty"`
}

// NewAccountCreation instantiates a new AccountCreation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccountCreation() *AccountCreation {
	this := AccountCreation{}
	return &this
}

// NewAccountCreationWithDefaults instantiates a new AccountCreation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccountCreationWithDefaults() *AccountCreation {
	this := AccountCreation{}
	return &this
}

// GetAccessStatus returns the AccessStatus field value if set, zero value otherwise.
func (o *AccountCreation) GetAccessStatus() AccountAccessStatus {
	if o == nil || o.AccessStatus == nil {
		var ret AccountAccessStatus
		return ret
	}
	return *o.AccessStatus
}

// GetAccessStatusOk returns a tuple with the AccessStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountCreation) GetAccessStatusOk() (*AccountAccessStatus, bool) {
	if o == nil || o.AccessStatus == nil {
		return nil, false
	}
	return o.AccessStatus, true
}

// HasAccessStatus returns a boolean if a field has been set.
func (o *AccountCreation) HasAccessStatus() bool {
	if o != nil && o.AccessStatus != nil {
		return true
	}

	return false
}

// SetAccessStatus gets a reference to the given AccountAccessStatus and assigns it to the AccessStatus field.
func (o *AccountCreation) SetAccessStatus(v AccountAccessStatus) {
	o.AccessStatus = &v
}

// GetAccountNumber returns the AccountNumber field value if set, zero value otherwise.
func (o *AccountCreation) GetAccountNumber() string {
	if o == nil || o.AccountNumber == nil {
		var ret string
		return ret
	}
	return *o.AccountNumber
}

// GetAccountNumberOk returns a tuple with the AccountNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountCreation) GetAccountNumberOk() (*string, bool) {
	if o == nil || o.AccountNumber == nil {
		return nil, false
	}
	return o.AccountNumber, true
}

// HasAccountNumber returns a boolean if a field has been set.
func (o *AccountCreation) HasAccountNumber() bool {
	if o != nil && o.AccountNumber != nil {
		return true
	}

	return false
}

// SetAccountNumber gets a reference to the given string and assigns it to the AccountNumber field.
func (o *AccountCreation) SetAccountNumber(v string) {
	o.AccountNumber = &v
}

// GetAccountPurpose returns the AccountPurpose field value if set, zero value otherwise.
func (o *AccountCreation) GetAccountPurpose() string {
	if o == nil || o.AccountPurpose == nil {
		var ret string
		return ret
	}
	return *o.AccountPurpose
}

// GetAccountPurposeOk returns a tuple with the AccountPurpose field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountCreation) GetAccountPurposeOk() (*string, bool) {
	if o == nil || o.AccountPurpose == nil {
		return nil, false
	}
	return o.AccountPurpose, true
}

// HasAccountPurpose returns a boolean if a field has been set.
func (o *AccountCreation) HasAccountPurpose() bool {
	if o != nil && o.AccountPurpose != nil {
		return true
	}

	return false
}

// SetAccountPurpose gets a reference to the given string and assigns it to the AccountPurpose field.
func (o *AccountCreation) SetAccountPurpose(v string) {
	o.AccountPurpose = &v
}

// GetAccountType returns the AccountType field value if set, zero value otherwise.
func (o *AccountCreation) GetAccountType() AccountType {
	if o == nil || o.AccountType == nil {
		var ret AccountType
		return ret
	}
	return *o.AccountType
}

// GetAccountTypeOk returns a tuple with the AccountType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountCreation) GetAccountTypeOk() (*AccountType, bool) {
	if o == nil || o.AccountType == nil {
		return nil, false
	}
	return o.AccountType, true
}

// HasAccountType returns a boolean if a field has been set.
func (o *AccountCreation) HasAccountType() bool {
	if o != nil && o.AccountType != nil {
		return true
	}

	return false
}

// SetAccountType gets a reference to the given AccountType and assigns it to the AccountType field.
func (o *AccountCreation) SetAccountType(v AccountType) {
	o.AccountType = &v
}

// GetBalances returns the Balances field value if set, zero value otherwise.
func (o *AccountCreation) GetBalances() []Balance {
	if o == nil || o.Balances == nil {
		var ret []Balance
		return ret
	}
	return *o.Balances
}

// GetBalancesOk returns a tuple with the Balances field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountCreation) GetBalancesOk() (*[]Balance, bool) {
	if o == nil || o.Balances == nil {
		return nil, false
	}
	return o.Balances, true
}

// HasBalances returns a boolean if a field has been set.
func (o *AccountCreation) HasBalances() bool {
	if o != nil && o.Balances != nil {
		return true
	}

	return false
}

// SetBalances gets a reference to the given []Balance and assigns it to the Balances field.
func (o *AccountCreation) SetBalances(v []Balance) {
	o.Balances = &v
}

// GetBankRouting returns the BankRouting field value if set, zero value otherwise.
func (o *AccountCreation) GetBankRouting() string {
	if o == nil || o.BankRouting == nil {
		var ret string
		return ret
	}
	return *o.BankRouting
}

// GetBankRoutingOk returns a tuple with the BankRouting field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountCreation) GetBankRoutingOk() (*string, bool) {
	if o == nil || o.BankRouting == nil {
		return nil, false
	}
	return o.BankRouting, true
}

// HasBankRouting returns a boolean if a field has been set.
func (o *AccountCreation) HasBankRouting() bool {
	if o != nil && o.BankRouting != nil {
		return true
	}

	return false
}

// SetBankRouting gets a reference to the given string and assigns it to the BankRouting field.
func (o *AccountCreation) SetBankRouting(v string) {
	o.BankRouting = &v
}

// GetCreationTime returns the CreationTime field value if set, zero value otherwise.
func (o *AccountCreation) GetCreationTime() time.Time {
	if o == nil || o.CreationTime == nil {
		var ret time.Time
		return ret
	}
	return *o.CreationTime
}

// GetCreationTimeOk returns a tuple with the CreationTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountCreation) GetCreationTimeOk() (*time.Time, bool) {
	if o == nil || o.CreationTime == nil {
		return nil, false
	}
	return o.CreationTime, true
}

// HasCreationTime returns a boolean if a field has been set.
func (o *AccountCreation) HasCreationTime() bool {
	if o != nil && o.CreationTime != nil {
		return true
	}

	return false
}

// SetCreationTime gets a reference to the given time.Time and assigns it to the CreationTime field.
func (o *AccountCreation) SetCreationTime(v time.Time) {
	o.CreationTime = &v
}

// GetCurrency returns the Currency field value if set, zero value otherwise.
func (o *AccountCreation) GetCurrency() string {
	if o == nil || o.Currency == nil {
		var ret string
		return ret
	}
	return *o.Currency
}

// GetCurrencyOk returns a tuple with the Currency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountCreation) GetCurrencyOk() (*string, bool) {
	if o == nil || o.Currency == nil {
		return nil, false
	}
	return o.Currency, true
}

// HasCurrency returns a boolean if a field has been set.
func (o *AccountCreation) HasCurrency() bool {
	if o != nil && o.Currency != nil {
		return true
	}

	return false
}

// SetCurrency gets a reference to the given string and assigns it to the Currency field.
func (o *AccountCreation) SetCurrency(v string) {
	o.Currency = &v
}

// GetCustomerIds returns the CustomerIds field value if set, zero value otherwise.
func (o *AccountCreation) GetCustomerIds() []string {
	if o == nil || o.CustomerIds == nil {
		var ret []string
		return ret
	}
	return *o.CustomerIds
}

// GetCustomerIdsOk returns a tuple with the CustomerIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountCreation) GetCustomerIdsOk() (*[]string, bool) {
	if o == nil || o.CustomerIds == nil {
		return nil, false
	}
	return o.CustomerIds, true
}

// HasCustomerIds returns a boolean if a field has been set.
func (o *AccountCreation) HasCustomerIds() bool {
	if o != nil && o.CustomerIds != nil {
		return true
	}

	return false
}

// SetCustomerIds gets a reference to the given []string and assigns it to the CustomerIds field.
func (o *AccountCreation) SetCustomerIds(v []string) {
	o.CustomerIds = &v
}

// GetExchangeRateType returns the ExchangeRateType field value if set, zero value otherwise.
func (o *AccountCreation) GetExchangeRateType() string {
	if o == nil || o.ExchangeRateType == nil {
		var ret string
		return ret
	}
	return *o.ExchangeRateType
}

// GetExchangeRateTypeOk returns a tuple with the ExchangeRateType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountCreation) GetExchangeRateTypeOk() (*string, bool) {
	if o == nil || o.ExchangeRateType == nil {
		return nil, false
	}
	return o.ExchangeRateType, true
}

// HasExchangeRateType returns a boolean if a field has been set.
func (o *AccountCreation) HasExchangeRateType() bool {
	if o != nil && o.ExchangeRateType != nil {
		return true
	}

	return false
}

// SetExchangeRateType gets a reference to the given string and assigns it to the ExchangeRateType field.
func (o *AccountCreation) SetExchangeRateType(v string) {
	o.ExchangeRateType = &v
}

// GetFeeProductIds returns the FeeProductIds field value if set, zero value otherwise.
func (o *AccountCreation) GetFeeProductIds() []string {
	if o == nil || o.FeeProductIds == nil {
		var ret []string
		return ret
	}
	return *o.FeeProductIds
}

// GetFeeProductIdsOk returns a tuple with the FeeProductIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountCreation) GetFeeProductIdsOk() (*[]string, bool) {
	if o == nil || o.FeeProductIds == nil {
		return nil, false
	}
	return o.FeeProductIds, true
}

// HasFeeProductIds returns a boolean if a field has been set.
func (o *AccountCreation) HasFeeProductIds() bool {
	if o != nil && o.FeeProductIds != nil {
		return true
	}

	return false
}

// SetFeeProductIds gets a reference to the given []string and assigns it to the FeeProductIds field.
func (o *AccountCreation) SetFeeProductIds(v []string) {
	o.FeeProductIds = &v
}

// GetIban returns the Iban field value if set, zero value otherwise.
func (o *AccountCreation) GetIban() string {
	if o == nil || o.Iban == nil {
		var ret string
		return ret
	}
	return *o.Iban
}

// GetIbanOk returns a tuple with the Iban field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountCreation) GetIbanOk() (*string, bool) {
	if o == nil || o.Iban == nil {
		return nil, false
	}
	return o.Iban, true
}

// HasIban returns a boolean if a field has been set.
func (o *AccountCreation) HasIban() bool {
	if o != nil && o.Iban != nil {
		return true
	}

	return false
}

// SetIban gets a reference to the given string and assigns it to the Iban field.
func (o *AccountCreation) SetIban(v string) {
	o.Iban = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *AccountCreation) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountCreation) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *AccountCreation) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *AccountCreation) SetId(v string) {
	o.Id = &v
}

// GetInterestProductId returns the InterestProductId field value if set, zero value otherwise.
func (o *AccountCreation) GetInterestProductId() string {
	if o == nil || o.InterestProductId == nil {
		var ret string
		return ret
	}
	return *o.InterestProductId
}

// GetInterestProductIdOk returns a tuple with the InterestProductId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountCreation) GetInterestProductIdOk() (*string, bool) {
	if o == nil || o.InterestProductId == nil {
		return nil, false
	}
	return o.InterestProductId, true
}

// HasInterestProductId returns a boolean if a field has been set.
func (o *AccountCreation) HasInterestProductId() bool {
	if o != nil && o.InterestProductId != nil {
		return true
	}

	return false
}

// SetInterestProductId gets a reference to the given string and assigns it to the InterestProductId field.
func (o *AccountCreation) SetInterestProductId(v string) {
	o.InterestProductId = &v
}

// GetIsAccountPool returns the IsAccountPool field value if set, zero value otherwise.
func (o *AccountCreation) GetIsAccountPool() bool {
	if o == nil || o.IsAccountPool == nil {
		var ret bool
		return ret
	}
	return *o.IsAccountPool
}

// GetIsAccountPoolOk returns a tuple with the IsAccountPool field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountCreation) GetIsAccountPoolOk() (*bool, bool) {
	if o == nil || o.IsAccountPool == nil {
		return nil, false
	}
	return o.IsAccountPool, true
}

// HasIsAccountPool returns a boolean if a field has been set.
func (o *AccountCreation) HasIsAccountPool() bool {
	if o != nil && o.IsAccountPool != nil {
		return true
	}

	return false
}

// SetIsAccountPool gets a reference to the given bool and assigns it to the IsAccountPool field.
func (o *AccountCreation) SetIsAccountPool(v bool) {
	o.IsAccountPool = &v
}

// GetLastUpdatedTime returns the LastUpdatedTime field value if set, zero value otherwise.
func (o *AccountCreation) GetLastUpdatedTime() time.Time {
	if o == nil || o.LastUpdatedTime == nil {
		var ret time.Time
		return ret
	}
	return *o.LastUpdatedTime
}

// GetLastUpdatedTimeOk returns a tuple with the LastUpdatedTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountCreation) GetLastUpdatedTimeOk() (*time.Time, bool) {
	if o == nil || o.LastUpdatedTime == nil {
		return nil, false
	}
	return o.LastUpdatedTime, true
}

// HasLastUpdatedTime returns a boolean if a field has been set.
func (o *AccountCreation) HasLastUpdatedTime() bool {
	if o != nil && o.LastUpdatedTime != nil {
		return true
	}

	return false
}

// SetLastUpdatedTime gets a reference to the given time.Time and assigns it to the LastUpdatedTime field.
func (o *AccountCreation) SetLastUpdatedTime(v time.Time) {
	o.LastUpdatedTime = &v
}

// GetOverdraftLimit returns the OverdraftLimit field value if set, zero value otherwise.
func (o *AccountCreation) GetOverdraftLimit() int64 {
	if o == nil || o.OverdraftLimit == nil {
		var ret int64
		return ret
	}
	return *o.OverdraftLimit
}

// GetOverdraftLimitOk returns a tuple with the OverdraftLimit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountCreation) GetOverdraftLimitOk() (*int64, bool) {
	if o == nil || o.OverdraftLimit == nil {
		return nil, false
	}
	return o.OverdraftLimit, true
}

// HasOverdraftLimit returns a boolean if a field has been set.
func (o *AccountCreation) HasOverdraftLimit() bool {
	if o != nil && o.OverdraftLimit != nil {
		return true
	}

	return false
}

// SetOverdraftLimit gets a reference to the given int64 and assigns it to the OverdraftLimit field.
func (o *AccountCreation) SetOverdraftLimit(v int64) {
	o.OverdraftLimit = &v
}

// GetSpendingLimits returns the SpendingLimits field value if set, zero value otherwise.
func (o *AccountCreation) GetSpendingLimits() SpendingLimits {
	if o == nil || o.SpendingLimits == nil {
		var ret SpendingLimits
		return ret
	}
	return *o.SpendingLimits
}

// GetSpendingLimitsOk returns a tuple with the SpendingLimits field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountCreation) GetSpendingLimitsOk() (*SpendingLimits, bool) {
	if o == nil || o.SpendingLimits == nil {
		return nil, false
	}
	return o.SpendingLimits, true
}

// HasSpendingLimits returns a boolean if a field has been set.
func (o *AccountCreation) HasSpendingLimits() bool {
	if o != nil && o.SpendingLimits != nil {
		return true
	}

	return false
}

// SetSpendingLimits gets a reference to the given SpendingLimits and assigns it to the SpendingLimits field.
func (o *AccountCreation) SetSpendingLimits(v SpendingLimits) {
	o.SpendingLimits = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *AccountCreation) GetStatus() Status {
	if o == nil || o.Status == nil {
		var ret Status
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountCreation) GetStatusOk() (*Status, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *AccountCreation) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given Status and assigns it to the Status field.
func (o *AccountCreation) SetStatus(v Status) {
	o.Status = &v
}

// GetSwiftCode returns the SwiftCode field value if set, zero value otherwise.
func (o *AccountCreation) GetSwiftCode() string {
	if o == nil || o.SwiftCode == nil {
		var ret string
		return ret
	}
	return *o.SwiftCode
}

// GetSwiftCodeOk returns a tuple with the SwiftCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountCreation) GetSwiftCodeOk() (*string, bool) {
	if o == nil || o.SwiftCode == nil {
		return nil, false
	}
	return o.SwiftCode, true
}

// HasSwiftCode returns a boolean if a field has been set.
func (o *AccountCreation) HasSwiftCode() bool {
	if o != nil && o.SwiftCode != nil {
		return true
	}

	return false
}

// SetSwiftCode gets a reference to the given string and assigns it to the SwiftCode field.
func (o *AccountCreation) SetSwiftCode(v string) {
	o.SwiftCode = &v
}

// GetAccountTemplateId returns the AccountTemplateId field value if set, zero value otherwise.
func (o *AccountCreation) GetAccountTemplateId() string {
	if o == nil || o.AccountTemplateId == nil {
		var ret string
		return ret
	}
	return *o.AccountTemplateId
}

// GetAccountTemplateIdOk returns a tuple with the AccountTemplateId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountCreation) GetAccountTemplateIdOk() (*string, bool) {
	if o == nil || o.AccountTemplateId == nil {
		return nil, false
	}
	return o.AccountTemplateId, true
}

// HasAccountTemplateId returns a boolean if a field has been set.
func (o *AccountCreation) HasAccountTemplateId() bool {
	if o != nil && o.AccountTemplateId != nil {
		return true
	}

	return false
}

// SetAccountTemplateId gets a reference to the given string and assigns it to the AccountTemplateId field.
func (o *AccountCreation) SetAccountTemplateId(v string) {
	o.AccountTemplateId = &v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *AccountCreation) GetRelationships() []Relationship {
	if o == nil || o.Relationships == nil {
		var ret []Relationship
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountCreation) GetRelationshipsOk() (*[]Relationship, bool) {
	if o == nil || o.Relationships == nil {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *AccountCreation) HasRelationships() bool {
	if o != nil && o.Relationships != nil {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given []Relationship and assigns it to the Relationships field.
func (o *AccountCreation) SetRelationships(v []Relationship) {
	o.Relationships = &v
}

func (o AccountCreation) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AccessStatus != nil {
		toSerialize["access_status"] = o.AccessStatus
	}
	if o.AccountNumber != nil {
		toSerialize["account_number"] = o.AccountNumber
	}
	if o.AccountPurpose != nil {
		toSerialize["account_purpose"] = o.AccountPurpose
	}
	if o.AccountType != nil {
		toSerialize["account_type"] = o.AccountType
	}
	if o.Balances != nil {
		toSerialize["balances"] = o.Balances
	}
	if o.BankRouting != nil {
		toSerialize["bank_routing"] = o.BankRouting
	}
	if o.CreationTime != nil {
		toSerialize["creation_time"] = o.CreationTime
	}
	if o.Currency != nil {
		toSerialize["currency"] = o.Currency
	}
	if o.CustomerIds != nil {
		toSerialize["customer_ids"] = o.CustomerIds
	}
	if o.ExchangeRateType != nil {
		toSerialize["exchange_rate_type"] = o.ExchangeRateType
	}
	if o.FeeProductIds != nil {
		toSerialize["fee_product_ids"] = o.FeeProductIds
	}
	if o.Iban != nil {
		toSerialize["iban"] = o.Iban
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.InterestProductId != nil {
		toSerialize["interest_product_id"] = o.InterestProductId
	}
	if o.IsAccountPool != nil {
		toSerialize["is_account_pool"] = o.IsAccountPool
	}
	if o.LastUpdatedTime != nil {
		toSerialize["last_updated_time"] = o.LastUpdatedTime
	}
	if o.OverdraftLimit != nil {
		toSerialize["overdraft_limit"] = o.OverdraftLimit
	}
	if o.SpendingLimits != nil {
		toSerialize["spending_limits"] = o.SpendingLimits
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	if o.SwiftCode != nil {
		toSerialize["swift_code"] = o.SwiftCode
	}
	if o.AccountTemplateId != nil {
		toSerialize["account_template_id"] = o.AccountTemplateId
	}
	if o.Relationships != nil {
		toSerialize["relationships"] = o.Relationships
	}
	return json.Marshal(toSerialize)
}

type NullableAccountCreation struct {
	value *AccountCreation
	isSet bool
}

func (v NullableAccountCreation) Get() *AccountCreation {
	return v.value
}

func (v *NullableAccountCreation) Set(val *AccountCreation) {
	v.value = val
	v.isSet = true
}

func (v NullableAccountCreation) IsSet() bool {
	return v.isSet
}

func (v *NullableAccountCreation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccountCreation(val *AccountCreation) *NullableAccountCreation {
	return &NullableAccountCreation{value: val, isSet: true}
}

func (v NullableAccountCreation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccountCreation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


