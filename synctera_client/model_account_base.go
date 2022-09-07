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

// AccountBase struct for AccountBase
type AccountBase struct {
	AccessStatus *AccountAccessStatus `json:"access_status,omitempty"`
	// Account number
	AccountNumber *string `json:"account_number,omitempty"`
	// The response will contain the bank fintech ID (3 or 6 digits) plus the last 4 digits, with the digits in between replaced with * characters. Shadow mode account numbers will not be masked.
	AccountNumberMasked *string `json:"account_number_masked,omitempty"`
	// Purpose of the account
	AccountPurpose *string `json:"account_purpose,omitempty"`
	AccountType *AccountType `json:"account_type,omitempty"`
	// A list of balances for account based on different type
	Balances []Balance `json:"balances,omitempty"`
	// Bank routing number
	BankRouting *string `json:"bank_routing,omitempty"`
	// Account creation timestamp in RFC3339 format
	CreationTime *time.Time `json:"creation_time,omitempty"`
	// Account currency or account settlement currency. ISO 4217 alphabetic currency code. Default USD
	Currency *string `json:"currency,omitempty"`
	// A list of the customer IDs of the account holders.
	CustomerIds []string `json:"customer_ids,omitempty"`
	CustomerType *CustomerType `json:"customer_type,omitempty"`
	// Exchange rate type
	ExchangeRateType *string `json:"exchange_rate_type,omitempty"`
	// A list of fee resources from account product that the current account associate with
	FeeProductIds []string `json:"fee_product_ids,omitempty"`
	// International bank account number
	Iban *string `json:"iban,omitempty"`
	// Account ID
	Id *string `json:"id,omitempty"`
	// An interest from account product that the current account associate with
	InterestProductId *string `json:"interest_product_id,omitempty"`
	// Account is investment (variable balance) account or a multi-balance account pool. Default false
	IsAccountPool *bool `json:"is_account_pool,omitempty"`
	// A flag to indicate whether ACH transactions are enabled.
	IsAchEnabled *bool `json:"is_ach_enabled,omitempty"`
	// A flag to indicate whether card transactions are enabled.
	IsCardEnabled *bool `json:"is_card_enabled,omitempty"`
	// A flag to indicate whether P2P transactions are enabled.
	IsP2pEnabled *bool `json:"is_p2p_enabled,omitempty"`
	// A flag to indicate whether wire transactions are enabled.
	IsWireEnabled *bool `json:"is_wire_enabled,omitempty"`
	// Timestamp of the last account modification in RFC3339 format
	LastUpdatedTime *time.Time `json:"last_updated_time,omitempty"`
	// User provided account metadata
	Metadata map[string]interface{} `json:"metadata,omitempty"`
	// User provided account nickname
	Nickname *string `json:"nickname,omitempty"`
	Status *Status `json:"status,omitempty"`
	// SWIFT code
	SwiftCode *string `json:"swift_code,omitempty"`
}

// NewAccountBase instantiates a new AccountBase object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccountBase() *AccountBase {
	this := AccountBase{}
	return &this
}

// NewAccountBaseWithDefaults instantiates a new AccountBase object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccountBaseWithDefaults() *AccountBase {
	this := AccountBase{}
	return &this
}

// GetAccessStatus returns the AccessStatus field value if set, zero value otherwise.
func (o *AccountBase) GetAccessStatus() AccountAccessStatus {
	if o == nil || o.AccessStatus == nil {
		var ret AccountAccessStatus
		return ret
	}
	return *o.AccessStatus
}

// GetAccessStatusOk returns a tuple with the AccessStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountBase) GetAccessStatusOk() (*AccountAccessStatus, bool) {
	if o == nil || o.AccessStatus == nil {
		return nil, false
	}
	return o.AccessStatus, true
}

// HasAccessStatus returns a boolean if a field has been set.
func (o *AccountBase) HasAccessStatus() bool {
	if o != nil && o.AccessStatus != nil {
		return true
	}

	return false
}

// SetAccessStatus gets a reference to the given AccountAccessStatus and assigns it to the AccessStatus field.
func (o *AccountBase) SetAccessStatus(v AccountAccessStatus) {
	o.AccessStatus = &v
}

// GetAccountNumber returns the AccountNumber field value if set, zero value otherwise.
func (o *AccountBase) GetAccountNumber() string {
	if o == nil || o.AccountNumber == nil {
		var ret string
		return ret
	}
	return *o.AccountNumber
}

// GetAccountNumberOk returns a tuple with the AccountNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountBase) GetAccountNumberOk() (*string, bool) {
	if o == nil || o.AccountNumber == nil {
		return nil, false
	}
	return o.AccountNumber, true
}

// HasAccountNumber returns a boolean if a field has been set.
func (o *AccountBase) HasAccountNumber() bool {
	if o != nil && o.AccountNumber != nil {
		return true
	}

	return false
}

// SetAccountNumber gets a reference to the given string and assigns it to the AccountNumber field.
func (o *AccountBase) SetAccountNumber(v string) {
	o.AccountNumber = &v
}

// GetAccountNumberMasked returns the AccountNumberMasked field value if set, zero value otherwise.
func (o *AccountBase) GetAccountNumberMasked() string {
	if o == nil || o.AccountNumberMasked == nil {
		var ret string
		return ret
	}
	return *o.AccountNumberMasked
}

// GetAccountNumberMaskedOk returns a tuple with the AccountNumberMasked field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountBase) GetAccountNumberMaskedOk() (*string, bool) {
	if o == nil || o.AccountNumberMasked == nil {
		return nil, false
	}
	return o.AccountNumberMasked, true
}

// HasAccountNumberMasked returns a boolean if a field has been set.
func (o *AccountBase) HasAccountNumberMasked() bool {
	if o != nil && o.AccountNumberMasked != nil {
		return true
	}

	return false
}

// SetAccountNumberMasked gets a reference to the given string and assigns it to the AccountNumberMasked field.
func (o *AccountBase) SetAccountNumberMasked(v string) {
	o.AccountNumberMasked = &v
}

// GetAccountPurpose returns the AccountPurpose field value if set, zero value otherwise.
func (o *AccountBase) GetAccountPurpose() string {
	if o == nil || o.AccountPurpose == nil {
		var ret string
		return ret
	}
	return *o.AccountPurpose
}

// GetAccountPurposeOk returns a tuple with the AccountPurpose field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountBase) GetAccountPurposeOk() (*string, bool) {
	if o == nil || o.AccountPurpose == nil {
		return nil, false
	}
	return o.AccountPurpose, true
}

// HasAccountPurpose returns a boolean if a field has been set.
func (o *AccountBase) HasAccountPurpose() bool {
	if o != nil && o.AccountPurpose != nil {
		return true
	}

	return false
}

// SetAccountPurpose gets a reference to the given string and assigns it to the AccountPurpose field.
func (o *AccountBase) SetAccountPurpose(v string) {
	o.AccountPurpose = &v
}

// GetAccountType returns the AccountType field value if set, zero value otherwise.
func (o *AccountBase) GetAccountType() AccountType {
	if o == nil || o.AccountType == nil {
		var ret AccountType
		return ret
	}
	return *o.AccountType
}

// GetAccountTypeOk returns a tuple with the AccountType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountBase) GetAccountTypeOk() (*AccountType, bool) {
	if o == nil || o.AccountType == nil {
		return nil, false
	}
	return o.AccountType, true
}

// HasAccountType returns a boolean if a field has been set.
func (o *AccountBase) HasAccountType() bool {
	if o != nil && o.AccountType != nil {
		return true
	}

	return false
}

// SetAccountType gets a reference to the given AccountType and assigns it to the AccountType field.
func (o *AccountBase) SetAccountType(v AccountType) {
	o.AccountType = &v
}

// GetBalances returns the Balances field value if set, zero value otherwise.
func (o *AccountBase) GetBalances() []Balance {
	if o == nil || o.Balances == nil {
		var ret []Balance
		return ret
	}
	return o.Balances
}

// GetBalancesOk returns a tuple with the Balances field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountBase) GetBalancesOk() ([]Balance, bool) {
	if o == nil || o.Balances == nil {
		return nil, false
	}
	return o.Balances, true
}

// HasBalances returns a boolean if a field has been set.
func (o *AccountBase) HasBalances() bool {
	if o != nil && o.Balances != nil {
		return true
	}

	return false
}

// SetBalances gets a reference to the given []Balance and assigns it to the Balances field.
func (o *AccountBase) SetBalances(v []Balance) {
	o.Balances = v
}

// GetBankRouting returns the BankRouting field value if set, zero value otherwise.
func (o *AccountBase) GetBankRouting() string {
	if o == nil || o.BankRouting == nil {
		var ret string
		return ret
	}
	return *o.BankRouting
}

// GetBankRoutingOk returns a tuple with the BankRouting field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountBase) GetBankRoutingOk() (*string, bool) {
	if o == nil || o.BankRouting == nil {
		return nil, false
	}
	return o.BankRouting, true
}

// HasBankRouting returns a boolean if a field has been set.
func (o *AccountBase) HasBankRouting() bool {
	if o != nil && o.BankRouting != nil {
		return true
	}

	return false
}

// SetBankRouting gets a reference to the given string and assigns it to the BankRouting field.
func (o *AccountBase) SetBankRouting(v string) {
	o.BankRouting = &v
}

// GetCreationTime returns the CreationTime field value if set, zero value otherwise.
func (o *AccountBase) GetCreationTime() time.Time {
	if o == nil || o.CreationTime == nil {
		var ret time.Time
		return ret
	}
	return *o.CreationTime
}

// GetCreationTimeOk returns a tuple with the CreationTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountBase) GetCreationTimeOk() (*time.Time, bool) {
	if o == nil || o.CreationTime == nil {
		return nil, false
	}
	return o.CreationTime, true
}

// HasCreationTime returns a boolean if a field has been set.
func (o *AccountBase) HasCreationTime() bool {
	if o != nil && o.CreationTime != nil {
		return true
	}

	return false
}

// SetCreationTime gets a reference to the given time.Time and assigns it to the CreationTime field.
func (o *AccountBase) SetCreationTime(v time.Time) {
	o.CreationTime = &v
}

// GetCurrency returns the Currency field value if set, zero value otherwise.
func (o *AccountBase) GetCurrency() string {
	if o == nil || o.Currency == nil {
		var ret string
		return ret
	}
	return *o.Currency
}

// GetCurrencyOk returns a tuple with the Currency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountBase) GetCurrencyOk() (*string, bool) {
	if o == nil || o.Currency == nil {
		return nil, false
	}
	return o.Currency, true
}

// HasCurrency returns a boolean if a field has been set.
func (o *AccountBase) HasCurrency() bool {
	if o != nil && o.Currency != nil {
		return true
	}

	return false
}

// SetCurrency gets a reference to the given string and assigns it to the Currency field.
func (o *AccountBase) SetCurrency(v string) {
	o.Currency = &v
}

// GetCustomerIds returns the CustomerIds field value if set, zero value otherwise.
func (o *AccountBase) GetCustomerIds() []string {
	if o == nil || o.CustomerIds == nil {
		var ret []string
		return ret
	}
	return o.CustomerIds
}

// GetCustomerIdsOk returns a tuple with the CustomerIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountBase) GetCustomerIdsOk() ([]string, bool) {
	if o == nil || o.CustomerIds == nil {
		return nil, false
	}
	return o.CustomerIds, true
}

// HasCustomerIds returns a boolean if a field has been set.
func (o *AccountBase) HasCustomerIds() bool {
	if o != nil && o.CustomerIds != nil {
		return true
	}

	return false
}

// SetCustomerIds gets a reference to the given []string and assigns it to the CustomerIds field.
func (o *AccountBase) SetCustomerIds(v []string) {
	o.CustomerIds = v
}

// GetCustomerType returns the CustomerType field value if set, zero value otherwise.
func (o *AccountBase) GetCustomerType() CustomerType {
	if o == nil || o.CustomerType == nil {
		var ret CustomerType
		return ret
	}
	return *o.CustomerType
}

// GetCustomerTypeOk returns a tuple with the CustomerType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountBase) GetCustomerTypeOk() (*CustomerType, bool) {
	if o == nil || o.CustomerType == nil {
		return nil, false
	}
	return o.CustomerType, true
}

// HasCustomerType returns a boolean if a field has been set.
func (o *AccountBase) HasCustomerType() bool {
	if o != nil && o.CustomerType != nil {
		return true
	}

	return false
}

// SetCustomerType gets a reference to the given CustomerType and assigns it to the CustomerType field.
func (o *AccountBase) SetCustomerType(v CustomerType) {
	o.CustomerType = &v
}

// GetExchangeRateType returns the ExchangeRateType field value if set, zero value otherwise.
func (o *AccountBase) GetExchangeRateType() string {
	if o == nil || o.ExchangeRateType == nil {
		var ret string
		return ret
	}
	return *o.ExchangeRateType
}

// GetExchangeRateTypeOk returns a tuple with the ExchangeRateType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountBase) GetExchangeRateTypeOk() (*string, bool) {
	if o == nil || o.ExchangeRateType == nil {
		return nil, false
	}
	return o.ExchangeRateType, true
}

// HasExchangeRateType returns a boolean if a field has been set.
func (o *AccountBase) HasExchangeRateType() bool {
	if o != nil && o.ExchangeRateType != nil {
		return true
	}

	return false
}

// SetExchangeRateType gets a reference to the given string and assigns it to the ExchangeRateType field.
func (o *AccountBase) SetExchangeRateType(v string) {
	o.ExchangeRateType = &v
}

// GetFeeProductIds returns the FeeProductIds field value if set, zero value otherwise.
func (o *AccountBase) GetFeeProductIds() []string {
	if o == nil || o.FeeProductIds == nil {
		var ret []string
		return ret
	}
	return o.FeeProductIds
}

// GetFeeProductIdsOk returns a tuple with the FeeProductIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountBase) GetFeeProductIdsOk() ([]string, bool) {
	if o == nil || o.FeeProductIds == nil {
		return nil, false
	}
	return o.FeeProductIds, true
}

// HasFeeProductIds returns a boolean if a field has been set.
func (o *AccountBase) HasFeeProductIds() bool {
	if o != nil && o.FeeProductIds != nil {
		return true
	}

	return false
}

// SetFeeProductIds gets a reference to the given []string and assigns it to the FeeProductIds field.
func (o *AccountBase) SetFeeProductIds(v []string) {
	o.FeeProductIds = v
}

// GetIban returns the Iban field value if set, zero value otherwise.
func (o *AccountBase) GetIban() string {
	if o == nil || o.Iban == nil {
		var ret string
		return ret
	}
	return *o.Iban
}

// GetIbanOk returns a tuple with the Iban field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountBase) GetIbanOk() (*string, bool) {
	if o == nil || o.Iban == nil {
		return nil, false
	}
	return o.Iban, true
}

// HasIban returns a boolean if a field has been set.
func (o *AccountBase) HasIban() bool {
	if o != nil && o.Iban != nil {
		return true
	}

	return false
}

// SetIban gets a reference to the given string and assigns it to the Iban field.
func (o *AccountBase) SetIban(v string) {
	o.Iban = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *AccountBase) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountBase) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *AccountBase) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *AccountBase) SetId(v string) {
	o.Id = &v
}

// GetInterestProductId returns the InterestProductId field value if set, zero value otherwise.
func (o *AccountBase) GetInterestProductId() string {
	if o == nil || o.InterestProductId == nil {
		var ret string
		return ret
	}
	return *o.InterestProductId
}

// GetInterestProductIdOk returns a tuple with the InterestProductId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountBase) GetInterestProductIdOk() (*string, bool) {
	if o == nil || o.InterestProductId == nil {
		return nil, false
	}
	return o.InterestProductId, true
}

// HasInterestProductId returns a boolean if a field has been set.
func (o *AccountBase) HasInterestProductId() bool {
	if o != nil && o.InterestProductId != nil {
		return true
	}

	return false
}

// SetInterestProductId gets a reference to the given string and assigns it to the InterestProductId field.
func (o *AccountBase) SetInterestProductId(v string) {
	o.InterestProductId = &v
}

// GetIsAccountPool returns the IsAccountPool field value if set, zero value otherwise.
func (o *AccountBase) GetIsAccountPool() bool {
	if o == nil || o.IsAccountPool == nil {
		var ret bool
		return ret
	}
	return *o.IsAccountPool
}

// GetIsAccountPoolOk returns a tuple with the IsAccountPool field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountBase) GetIsAccountPoolOk() (*bool, bool) {
	if o == nil || o.IsAccountPool == nil {
		return nil, false
	}
	return o.IsAccountPool, true
}

// HasIsAccountPool returns a boolean if a field has been set.
func (o *AccountBase) HasIsAccountPool() bool {
	if o != nil && o.IsAccountPool != nil {
		return true
	}

	return false
}

// SetIsAccountPool gets a reference to the given bool and assigns it to the IsAccountPool field.
func (o *AccountBase) SetIsAccountPool(v bool) {
	o.IsAccountPool = &v
}

// GetIsAchEnabled returns the IsAchEnabled field value if set, zero value otherwise.
func (o *AccountBase) GetIsAchEnabled() bool {
	if o == nil || o.IsAchEnabled == nil {
		var ret bool
		return ret
	}
	return *o.IsAchEnabled
}

// GetIsAchEnabledOk returns a tuple with the IsAchEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountBase) GetIsAchEnabledOk() (*bool, bool) {
	if o == nil || o.IsAchEnabled == nil {
		return nil, false
	}
	return o.IsAchEnabled, true
}

// HasIsAchEnabled returns a boolean if a field has been set.
func (o *AccountBase) HasIsAchEnabled() bool {
	if o != nil && o.IsAchEnabled != nil {
		return true
	}

	return false
}

// SetIsAchEnabled gets a reference to the given bool and assigns it to the IsAchEnabled field.
func (o *AccountBase) SetIsAchEnabled(v bool) {
	o.IsAchEnabled = &v
}

// GetIsCardEnabled returns the IsCardEnabled field value if set, zero value otherwise.
func (o *AccountBase) GetIsCardEnabled() bool {
	if o == nil || o.IsCardEnabled == nil {
		var ret bool
		return ret
	}
	return *o.IsCardEnabled
}

// GetIsCardEnabledOk returns a tuple with the IsCardEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountBase) GetIsCardEnabledOk() (*bool, bool) {
	if o == nil || o.IsCardEnabled == nil {
		return nil, false
	}
	return o.IsCardEnabled, true
}

// HasIsCardEnabled returns a boolean if a field has been set.
func (o *AccountBase) HasIsCardEnabled() bool {
	if o != nil && o.IsCardEnabled != nil {
		return true
	}

	return false
}

// SetIsCardEnabled gets a reference to the given bool and assigns it to the IsCardEnabled field.
func (o *AccountBase) SetIsCardEnabled(v bool) {
	o.IsCardEnabled = &v
}

// GetIsP2pEnabled returns the IsP2pEnabled field value if set, zero value otherwise.
func (o *AccountBase) GetIsP2pEnabled() bool {
	if o == nil || o.IsP2pEnabled == nil {
		var ret bool
		return ret
	}
	return *o.IsP2pEnabled
}

// GetIsP2pEnabledOk returns a tuple with the IsP2pEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountBase) GetIsP2pEnabledOk() (*bool, bool) {
	if o == nil || o.IsP2pEnabled == nil {
		return nil, false
	}
	return o.IsP2pEnabled, true
}

// HasIsP2pEnabled returns a boolean if a field has been set.
func (o *AccountBase) HasIsP2pEnabled() bool {
	if o != nil && o.IsP2pEnabled != nil {
		return true
	}

	return false
}

// SetIsP2pEnabled gets a reference to the given bool and assigns it to the IsP2pEnabled field.
func (o *AccountBase) SetIsP2pEnabled(v bool) {
	o.IsP2pEnabled = &v
}

// GetIsWireEnabled returns the IsWireEnabled field value if set, zero value otherwise.
func (o *AccountBase) GetIsWireEnabled() bool {
	if o == nil || o.IsWireEnabled == nil {
		var ret bool
		return ret
	}
	return *o.IsWireEnabled
}

// GetIsWireEnabledOk returns a tuple with the IsWireEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountBase) GetIsWireEnabledOk() (*bool, bool) {
	if o == nil || o.IsWireEnabled == nil {
		return nil, false
	}
	return o.IsWireEnabled, true
}

// HasIsWireEnabled returns a boolean if a field has been set.
func (o *AccountBase) HasIsWireEnabled() bool {
	if o != nil && o.IsWireEnabled != nil {
		return true
	}

	return false
}

// SetIsWireEnabled gets a reference to the given bool and assigns it to the IsWireEnabled field.
func (o *AccountBase) SetIsWireEnabled(v bool) {
	o.IsWireEnabled = &v
}

// GetLastUpdatedTime returns the LastUpdatedTime field value if set, zero value otherwise.
func (o *AccountBase) GetLastUpdatedTime() time.Time {
	if o == nil || o.LastUpdatedTime == nil {
		var ret time.Time
		return ret
	}
	return *o.LastUpdatedTime
}

// GetLastUpdatedTimeOk returns a tuple with the LastUpdatedTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountBase) GetLastUpdatedTimeOk() (*time.Time, bool) {
	if o == nil || o.LastUpdatedTime == nil {
		return nil, false
	}
	return o.LastUpdatedTime, true
}

// HasLastUpdatedTime returns a boolean if a field has been set.
func (o *AccountBase) HasLastUpdatedTime() bool {
	if o != nil && o.LastUpdatedTime != nil {
		return true
	}

	return false
}

// SetLastUpdatedTime gets a reference to the given time.Time and assigns it to the LastUpdatedTime field.
func (o *AccountBase) SetLastUpdatedTime(v time.Time) {
	o.LastUpdatedTime = &v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *AccountBase) GetMetadata() map[string]interface{} {
	if o == nil || o.Metadata == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountBase) GetMetadataOk() (map[string]interface{}, bool) {
	if o == nil || o.Metadata == nil {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *AccountBase) HasMetadata() bool {
	if o != nil && o.Metadata != nil {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given map[string]interface{} and assigns it to the Metadata field.
func (o *AccountBase) SetMetadata(v map[string]interface{}) {
	o.Metadata = v
}

// GetNickname returns the Nickname field value if set, zero value otherwise.
func (o *AccountBase) GetNickname() string {
	if o == nil || o.Nickname == nil {
		var ret string
		return ret
	}
	return *o.Nickname
}

// GetNicknameOk returns a tuple with the Nickname field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountBase) GetNicknameOk() (*string, bool) {
	if o == nil || o.Nickname == nil {
		return nil, false
	}
	return o.Nickname, true
}

// HasNickname returns a boolean if a field has been set.
func (o *AccountBase) HasNickname() bool {
	if o != nil && o.Nickname != nil {
		return true
	}

	return false
}

// SetNickname gets a reference to the given string and assigns it to the Nickname field.
func (o *AccountBase) SetNickname(v string) {
	o.Nickname = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *AccountBase) GetStatus() Status {
	if o == nil || o.Status == nil {
		var ret Status
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountBase) GetStatusOk() (*Status, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *AccountBase) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given Status and assigns it to the Status field.
func (o *AccountBase) SetStatus(v Status) {
	o.Status = &v
}

// GetSwiftCode returns the SwiftCode field value if set, zero value otherwise.
func (o *AccountBase) GetSwiftCode() string {
	if o == nil || o.SwiftCode == nil {
		var ret string
		return ret
	}
	return *o.SwiftCode
}

// GetSwiftCodeOk returns a tuple with the SwiftCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountBase) GetSwiftCodeOk() (*string, bool) {
	if o == nil || o.SwiftCode == nil {
		return nil, false
	}
	return o.SwiftCode, true
}

// HasSwiftCode returns a boolean if a field has been set.
func (o *AccountBase) HasSwiftCode() bool {
	if o != nil && o.SwiftCode != nil {
		return true
	}

	return false
}

// SetSwiftCode gets a reference to the given string and assigns it to the SwiftCode field.
func (o *AccountBase) SetSwiftCode(v string) {
	o.SwiftCode = &v
}

func (o AccountBase) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AccessStatus != nil {
		toSerialize["access_status"] = o.AccessStatus
	}
	if o.AccountNumber != nil {
		toSerialize["account_number"] = o.AccountNumber
	}
	if o.AccountNumberMasked != nil {
		toSerialize["account_number_masked"] = o.AccountNumberMasked
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
	if o.CustomerType != nil {
		toSerialize["customer_type"] = o.CustomerType
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
	if o.IsAchEnabled != nil {
		toSerialize["is_ach_enabled"] = o.IsAchEnabled
	}
	if o.IsCardEnabled != nil {
		toSerialize["is_card_enabled"] = o.IsCardEnabled
	}
	if o.IsP2pEnabled != nil {
		toSerialize["is_p2p_enabled"] = o.IsP2pEnabled
	}
	if o.IsWireEnabled != nil {
		toSerialize["is_wire_enabled"] = o.IsWireEnabled
	}
	if o.LastUpdatedTime != nil {
		toSerialize["last_updated_time"] = o.LastUpdatedTime
	}
	if o.Metadata != nil {
		toSerialize["metadata"] = o.Metadata
	}
	if o.Nickname != nil {
		toSerialize["nickname"] = o.Nickname
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	if o.SwiftCode != nil {
		toSerialize["swift_code"] = o.SwiftCode
	}
	return json.Marshal(toSerialize)
}

type NullableAccountBase struct {
	value *AccountBase
	isSet bool
}

func (v NullableAccountBase) Get() *AccountBase {
	return v.value
}

func (v *NullableAccountBase) Set(val *AccountBase) {
	v.value = val
	v.isSet = true
}

func (v NullableAccountBase) IsSet() bool {
	return v.isSet
}

func (v *NullableAccountBase) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccountBase(val *AccountBase) *NullableAccountBase {
	return &NullableAccountBase{value: val, isSet: true}
}

func (v NullableAccountBase) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccountBase) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

