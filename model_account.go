/*
 * Synctera API
 *
 * <h2>Let's build something great.</h2><p>Welcome to the official reference documentation for Synctera APIs. Our APIs are the best way to automate your company's banking needs and are designed to be easy to understand and implement.</p><p>We're continuously growing this library and what you see here is just the start, but if you need something specific or have a question, <a class='text-blue-600' href='https://synctera.com/contact' target='_blank' rel='noreferrer'>contact us</a>.</p>
 *
 * API version: 0.5.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package synctera

import (
	"encoding/json"
	"time"
)

// Account Account
type Account struct {
	AccessStatus *AccountAccessStatus `json:"access_status,omitempty"`
	// Account number
	AccountNumber *string `json:"account_number,omitempty"`
	// Purpose of the account
	AccountPurpose *string         `json:"account_purpose,omitempty"`
	AccountType    *AccountType    `json:"account_type,omitempty"`
	BalanceCeiling *BalanceCeiling `json:"balance_ceiling,omitempty"`
	BalanceFloor   *BalanceFloor   `json:"balance_floor,omitempty"`
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
	// User provided account metadata
	Metadata *map[string]interface{} `json:"metadata,omitempty"`
	// User provided account nickname
	Nickname *string `json:"nickname,omitempty"`
	// Account's overdraft limit
	OverdraftLimit *int64          `json:"overdraft_limit,omitempty"`
	SpendingLimits *SpendingLimits `json:"spending_limits,omitempty"`
	Status         *Status         `json:"status,omitempty"`
	// SWIFT code
	SwiftCode *string `json:"swift_code,omitempty"`
}

// NewAccount instantiates a new Account object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccount() *Account {
	this := Account{}
	return &this
}

// NewAccountWithDefaults instantiates a new Account object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccountWithDefaults() *Account {
	this := Account{}
	return &this
}

// GetAccessStatus returns the AccessStatus field value if set, zero value otherwise.
func (o *Account) GetAccessStatus() AccountAccessStatus {
	if o == nil || o.AccessStatus == nil {
		var ret AccountAccessStatus
		return ret
	}
	return *o.AccessStatus
}

// GetAccessStatusOk returns a tuple with the AccessStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Account) GetAccessStatusOk() (*AccountAccessStatus, bool) {
	if o == nil || o.AccessStatus == nil {
		return nil, false
	}
	return o.AccessStatus, true
}

// HasAccessStatus returns a boolean if a field has been set.
func (o *Account) HasAccessStatus() bool {
	if o != nil && o.AccessStatus != nil {
		return true
	}

	return false
}

// SetAccessStatus gets a reference to the given AccountAccessStatus and assigns it to the AccessStatus field.
func (o *Account) SetAccessStatus(v AccountAccessStatus) {
	o.AccessStatus = &v
}

// GetAccountNumber returns the AccountNumber field value if set, zero value otherwise.
func (o *Account) GetAccountNumber() string {
	if o == nil || o.AccountNumber == nil {
		var ret string
		return ret
	}
	return *o.AccountNumber
}

// GetAccountNumberOk returns a tuple with the AccountNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Account) GetAccountNumberOk() (*string, bool) {
	if o == nil || o.AccountNumber == nil {
		return nil, false
	}
	return o.AccountNumber, true
}

// HasAccountNumber returns a boolean if a field has been set.
func (o *Account) HasAccountNumber() bool {
	if o != nil && o.AccountNumber != nil {
		return true
	}

	return false
}

// SetAccountNumber gets a reference to the given string and assigns it to the AccountNumber field.
func (o *Account) SetAccountNumber(v string) {
	o.AccountNumber = &v
}

// GetAccountPurpose returns the AccountPurpose field value if set, zero value otherwise.
func (o *Account) GetAccountPurpose() string {
	if o == nil || o.AccountPurpose == nil {
		var ret string
		return ret
	}
	return *o.AccountPurpose
}

// GetAccountPurposeOk returns a tuple with the AccountPurpose field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Account) GetAccountPurposeOk() (*string, bool) {
	if o == nil || o.AccountPurpose == nil {
		return nil, false
	}
	return o.AccountPurpose, true
}

// HasAccountPurpose returns a boolean if a field has been set.
func (o *Account) HasAccountPurpose() bool {
	if o != nil && o.AccountPurpose != nil {
		return true
	}

	return false
}

// SetAccountPurpose gets a reference to the given string and assigns it to the AccountPurpose field.
func (o *Account) SetAccountPurpose(v string) {
	o.AccountPurpose = &v
}

// GetAccountType returns the AccountType field value if set, zero value otherwise.
func (o *Account) GetAccountType() AccountType {
	if o == nil || o.AccountType == nil {
		var ret AccountType
		return ret
	}
	return *o.AccountType
}

// GetAccountTypeOk returns a tuple with the AccountType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Account) GetAccountTypeOk() (*AccountType, bool) {
	if o == nil || o.AccountType == nil {
		return nil, false
	}
	return o.AccountType, true
}

// HasAccountType returns a boolean if a field has been set.
func (o *Account) HasAccountType() bool {
	if o != nil && o.AccountType != nil {
		return true
	}

	return false
}

// SetAccountType gets a reference to the given AccountType and assigns it to the AccountType field.
func (o *Account) SetAccountType(v AccountType) {
	o.AccountType = &v
}

// GetBalanceCeiling returns the BalanceCeiling field value if set, zero value otherwise.
func (o *Account) GetBalanceCeiling() BalanceCeiling {
	if o == nil || o.BalanceCeiling == nil {
		var ret BalanceCeiling
		return ret
	}
	return *o.BalanceCeiling
}

// GetBalanceCeilingOk returns a tuple with the BalanceCeiling field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Account) GetBalanceCeilingOk() (*BalanceCeiling, bool) {
	if o == nil || o.BalanceCeiling == nil {
		return nil, false
	}
	return o.BalanceCeiling, true
}

// HasBalanceCeiling returns a boolean if a field has been set.
func (o *Account) HasBalanceCeiling() bool {
	if o != nil && o.BalanceCeiling != nil {
		return true
	}

	return false
}

// SetBalanceCeiling gets a reference to the given BalanceCeiling and assigns it to the BalanceCeiling field.
func (o *Account) SetBalanceCeiling(v BalanceCeiling) {
	o.BalanceCeiling = &v
}

// GetBalanceFloor returns the BalanceFloor field value if set, zero value otherwise.
func (o *Account) GetBalanceFloor() BalanceFloor {
	if o == nil || o.BalanceFloor == nil {
		var ret BalanceFloor
		return ret
	}
	return *o.BalanceFloor
}

// GetBalanceFloorOk returns a tuple with the BalanceFloor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Account) GetBalanceFloorOk() (*BalanceFloor, bool) {
	if o == nil || o.BalanceFloor == nil {
		return nil, false
	}
	return o.BalanceFloor, true
}

// HasBalanceFloor returns a boolean if a field has been set.
func (o *Account) HasBalanceFloor() bool {
	if o != nil && o.BalanceFloor != nil {
		return true
	}

	return false
}

// SetBalanceFloor gets a reference to the given BalanceFloor and assigns it to the BalanceFloor field.
func (o *Account) SetBalanceFloor(v BalanceFloor) {
	o.BalanceFloor = &v
}

// GetBalances returns the Balances field value if set, zero value otherwise.
func (o *Account) GetBalances() []Balance {
	if o == nil || o.Balances == nil {
		var ret []Balance
		return ret
	}
	return *o.Balances
}

// GetBalancesOk returns a tuple with the Balances field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Account) GetBalancesOk() (*[]Balance, bool) {
	if o == nil || o.Balances == nil {
		return nil, false
	}
	return o.Balances, true
}

// HasBalances returns a boolean if a field has been set.
func (o *Account) HasBalances() bool {
	if o != nil && o.Balances != nil {
		return true
	}

	return false
}

// SetBalances gets a reference to the given []Balance and assigns it to the Balances field.
func (o *Account) SetBalances(v []Balance) {
	o.Balances = &v
}

// GetBankRouting returns the BankRouting field value if set, zero value otherwise.
func (o *Account) GetBankRouting() string {
	if o == nil || o.BankRouting == nil {
		var ret string
		return ret
	}
	return *o.BankRouting
}

// GetBankRoutingOk returns a tuple with the BankRouting field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Account) GetBankRoutingOk() (*string, bool) {
	if o == nil || o.BankRouting == nil {
		return nil, false
	}
	return o.BankRouting, true
}

// HasBankRouting returns a boolean if a field has been set.
func (o *Account) HasBankRouting() bool {
	if o != nil && o.BankRouting != nil {
		return true
	}

	return false
}

// SetBankRouting gets a reference to the given string and assigns it to the BankRouting field.
func (o *Account) SetBankRouting(v string) {
	o.BankRouting = &v
}

// GetCreationTime returns the CreationTime field value if set, zero value otherwise.
func (o *Account) GetCreationTime() time.Time {
	if o == nil || o.CreationTime == nil {
		var ret time.Time
		return ret
	}
	return *o.CreationTime
}

// GetCreationTimeOk returns a tuple with the CreationTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Account) GetCreationTimeOk() (*time.Time, bool) {
	if o == nil || o.CreationTime == nil {
		return nil, false
	}
	return o.CreationTime, true
}

// HasCreationTime returns a boolean if a field has been set.
func (o *Account) HasCreationTime() bool {
	if o != nil && o.CreationTime != nil {
		return true
	}

	return false
}

// SetCreationTime gets a reference to the given time.Time and assigns it to the CreationTime field.
func (o *Account) SetCreationTime(v time.Time) {
	o.CreationTime = &v
}

// GetCurrency returns the Currency field value if set, zero value otherwise.
func (o *Account) GetCurrency() string {
	if o == nil || o.Currency == nil {
		var ret string
		return ret
	}
	return *o.Currency
}

// GetCurrencyOk returns a tuple with the Currency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Account) GetCurrencyOk() (*string, bool) {
	if o == nil || o.Currency == nil {
		return nil, false
	}
	return o.Currency, true
}

// HasCurrency returns a boolean if a field has been set.
func (o *Account) HasCurrency() bool {
	if o != nil && o.Currency != nil {
		return true
	}

	return false
}

// SetCurrency gets a reference to the given string and assigns it to the Currency field.
func (o *Account) SetCurrency(v string) {
	o.Currency = &v
}

// GetCustomerIds returns the CustomerIds field value if set, zero value otherwise.
func (o *Account) GetCustomerIds() []string {
	if o == nil || o.CustomerIds == nil {
		var ret []string
		return ret
	}
	return *o.CustomerIds
}

// GetCustomerIdsOk returns a tuple with the CustomerIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Account) GetCustomerIdsOk() (*[]string, bool) {
	if o == nil || o.CustomerIds == nil {
		return nil, false
	}
	return o.CustomerIds, true
}

// HasCustomerIds returns a boolean if a field has been set.
func (o *Account) HasCustomerIds() bool {
	if o != nil && o.CustomerIds != nil {
		return true
	}

	return false
}

// SetCustomerIds gets a reference to the given []string and assigns it to the CustomerIds field.
func (o *Account) SetCustomerIds(v []string) {
	o.CustomerIds = &v
}

// GetExchangeRateType returns the ExchangeRateType field value if set, zero value otherwise.
func (o *Account) GetExchangeRateType() string {
	if o == nil || o.ExchangeRateType == nil {
		var ret string
		return ret
	}
	return *o.ExchangeRateType
}

// GetExchangeRateTypeOk returns a tuple with the ExchangeRateType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Account) GetExchangeRateTypeOk() (*string, bool) {
	if o == nil || o.ExchangeRateType == nil {
		return nil, false
	}
	return o.ExchangeRateType, true
}

// HasExchangeRateType returns a boolean if a field has been set.
func (o *Account) HasExchangeRateType() bool {
	if o != nil && o.ExchangeRateType != nil {
		return true
	}

	return false
}

// SetExchangeRateType gets a reference to the given string and assigns it to the ExchangeRateType field.
func (o *Account) SetExchangeRateType(v string) {
	o.ExchangeRateType = &v
}

// GetFeeProductIds returns the FeeProductIds field value if set, zero value otherwise.
func (o *Account) GetFeeProductIds() []string {
	if o == nil || o.FeeProductIds == nil {
		var ret []string
		return ret
	}
	return *o.FeeProductIds
}

// GetFeeProductIdsOk returns a tuple with the FeeProductIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Account) GetFeeProductIdsOk() (*[]string, bool) {
	if o == nil || o.FeeProductIds == nil {
		return nil, false
	}
	return o.FeeProductIds, true
}

// HasFeeProductIds returns a boolean if a field has been set.
func (o *Account) HasFeeProductIds() bool {
	if o != nil && o.FeeProductIds != nil {
		return true
	}

	return false
}

// SetFeeProductIds gets a reference to the given []string and assigns it to the FeeProductIds field.
func (o *Account) SetFeeProductIds(v []string) {
	o.FeeProductIds = &v
}

// GetIban returns the Iban field value if set, zero value otherwise.
func (o *Account) GetIban() string {
	if o == nil || o.Iban == nil {
		var ret string
		return ret
	}
	return *o.Iban
}

// GetIbanOk returns a tuple with the Iban field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Account) GetIbanOk() (*string, bool) {
	if o == nil || o.Iban == nil {
		return nil, false
	}
	return o.Iban, true
}

// HasIban returns a boolean if a field has been set.
func (o *Account) HasIban() bool {
	if o != nil && o.Iban != nil {
		return true
	}

	return false
}

// SetIban gets a reference to the given string and assigns it to the Iban field.
func (o *Account) SetIban(v string) {
	o.Iban = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Account) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Account) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Account) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *Account) SetId(v string) {
	o.Id = &v
}

// GetInterestProductId returns the InterestProductId field value if set, zero value otherwise.
func (o *Account) GetInterestProductId() string {
	if o == nil || o.InterestProductId == nil {
		var ret string
		return ret
	}
	return *o.InterestProductId
}

// GetInterestProductIdOk returns a tuple with the InterestProductId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Account) GetInterestProductIdOk() (*string, bool) {
	if o == nil || o.InterestProductId == nil {
		return nil, false
	}
	return o.InterestProductId, true
}

// HasInterestProductId returns a boolean if a field has been set.
func (o *Account) HasInterestProductId() bool {
	if o != nil && o.InterestProductId != nil {
		return true
	}

	return false
}

// SetInterestProductId gets a reference to the given string and assigns it to the InterestProductId field.
func (o *Account) SetInterestProductId(v string) {
	o.InterestProductId = &v
}

// GetIsAccountPool returns the IsAccountPool field value if set, zero value otherwise.
func (o *Account) GetIsAccountPool() bool {
	if o == nil || o.IsAccountPool == nil {
		var ret bool
		return ret
	}
	return *o.IsAccountPool
}

// GetIsAccountPoolOk returns a tuple with the IsAccountPool field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Account) GetIsAccountPoolOk() (*bool, bool) {
	if o == nil || o.IsAccountPool == nil {
		return nil, false
	}
	return o.IsAccountPool, true
}

// HasIsAccountPool returns a boolean if a field has been set.
func (o *Account) HasIsAccountPool() bool {
	if o != nil && o.IsAccountPool != nil {
		return true
	}

	return false
}

// SetIsAccountPool gets a reference to the given bool and assigns it to the IsAccountPool field.
func (o *Account) SetIsAccountPool(v bool) {
	o.IsAccountPool = &v
}

// GetLastUpdatedTime returns the LastUpdatedTime field value if set, zero value otherwise.
func (o *Account) GetLastUpdatedTime() time.Time {
	if o == nil || o.LastUpdatedTime == nil {
		var ret time.Time
		return ret
	}
	return *o.LastUpdatedTime
}

// GetLastUpdatedTimeOk returns a tuple with the LastUpdatedTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Account) GetLastUpdatedTimeOk() (*time.Time, bool) {
	if o == nil || o.LastUpdatedTime == nil {
		return nil, false
	}
	return o.LastUpdatedTime, true
}

// HasLastUpdatedTime returns a boolean if a field has been set.
func (o *Account) HasLastUpdatedTime() bool {
	if o != nil && o.LastUpdatedTime != nil {
		return true
	}

	return false
}

// SetLastUpdatedTime gets a reference to the given time.Time and assigns it to the LastUpdatedTime field.
func (o *Account) SetLastUpdatedTime(v time.Time) {
	o.LastUpdatedTime = &v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *Account) GetMetadata() map[string]interface{} {
	if o == nil || o.Metadata == nil {
		var ret map[string]interface{}
		return ret
	}
	return *o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Account) GetMetadataOk() (*map[string]interface{}, bool) {
	if o == nil || o.Metadata == nil {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *Account) HasMetadata() bool {
	if o != nil && o.Metadata != nil {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given map[string]interface{} and assigns it to the Metadata field.
func (o *Account) SetMetadata(v map[string]interface{}) {
	o.Metadata = &v
}

// GetNickname returns the Nickname field value if set, zero value otherwise.
func (o *Account) GetNickname() string {
	if o == nil || o.Nickname == nil {
		var ret string
		return ret
	}
	return *o.Nickname
}

// GetNicknameOk returns a tuple with the Nickname field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Account) GetNicknameOk() (*string, bool) {
	if o == nil || o.Nickname == nil {
		return nil, false
	}
	return o.Nickname, true
}

// HasNickname returns a boolean if a field has been set.
func (o *Account) HasNickname() bool {
	if o != nil && o.Nickname != nil {
		return true
	}

	return false
}

// SetNickname gets a reference to the given string and assigns it to the Nickname field.
func (o *Account) SetNickname(v string) {
	o.Nickname = &v
}

// GetOverdraftLimit returns the OverdraftLimit field value if set, zero value otherwise.
func (o *Account) GetOverdraftLimit() int64 {
	if o == nil || o.OverdraftLimit == nil {
		var ret int64
		return ret
	}
	return *o.OverdraftLimit
}

// GetOverdraftLimitOk returns a tuple with the OverdraftLimit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Account) GetOverdraftLimitOk() (*int64, bool) {
	if o == nil || o.OverdraftLimit == nil {
		return nil, false
	}
	return o.OverdraftLimit, true
}

// HasOverdraftLimit returns a boolean if a field has been set.
func (o *Account) HasOverdraftLimit() bool {
	if o != nil && o.OverdraftLimit != nil {
		return true
	}

	return false
}

// SetOverdraftLimit gets a reference to the given int64 and assigns it to the OverdraftLimit field.
func (o *Account) SetOverdraftLimit(v int64) {
	o.OverdraftLimit = &v
}

// GetSpendingLimits returns the SpendingLimits field value if set, zero value otherwise.
func (o *Account) GetSpendingLimits() SpendingLimits {
	if o == nil || o.SpendingLimits == nil {
		var ret SpendingLimits
		return ret
	}
	return *o.SpendingLimits
}

// GetSpendingLimitsOk returns a tuple with the SpendingLimits field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Account) GetSpendingLimitsOk() (*SpendingLimits, bool) {
	if o == nil || o.SpendingLimits == nil {
		return nil, false
	}
	return o.SpendingLimits, true
}

// HasSpendingLimits returns a boolean if a field has been set.
func (o *Account) HasSpendingLimits() bool {
	if o != nil && o.SpendingLimits != nil {
		return true
	}

	return false
}

// SetSpendingLimits gets a reference to the given SpendingLimits and assigns it to the SpendingLimits field.
func (o *Account) SetSpendingLimits(v SpendingLimits) {
	o.SpendingLimits = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *Account) GetStatus() Status {
	if o == nil || o.Status == nil {
		var ret Status
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Account) GetStatusOk() (*Status, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *Account) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given Status and assigns it to the Status field.
func (o *Account) SetStatus(v Status) {
	o.Status = &v
}

// GetSwiftCode returns the SwiftCode field value if set, zero value otherwise.
func (o *Account) GetSwiftCode() string {
	if o == nil || o.SwiftCode == nil {
		var ret string
		return ret
	}
	return *o.SwiftCode
}

// GetSwiftCodeOk returns a tuple with the SwiftCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Account) GetSwiftCodeOk() (*string, bool) {
	if o == nil || o.SwiftCode == nil {
		return nil, false
	}
	return o.SwiftCode, true
}

// HasSwiftCode returns a boolean if a field has been set.
func (o *Account) HasSwiftCode() bool {
	if o != nil && o.SwiftCode != nil {
		return true
	}

	return false
}

// SetSwiftCode gets a reference to the given string and assigns it to the SwiftCode field.
func (o *Account) SetSwiftCode(v string) {
	o.SwiftCode = &v
}

func (o Account) MarshalJSON() ([]byte, error) {
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
	if o.BalanceCeiling != nil {
		toSerialize["balance_ceiling"] = o.BalanceCeiling
	}
	if o.BalanceFloor != nil {
		toSerialize["balance_floor"] = o.BalanceFloor
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
	if o.Metadata != nil {
		toSerialize["metadata"] = o.Metadata
	}
	if o.Nickname != nil {
		toSerialize["nickname"] = o.Nickname
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
	return json.Marshal(toSerialize)
}

type NullableAccount struct {
	value *Account
	isSet bool
}

func (v NullableAccount) Get() *Account {
	return v.value
}

func (v *NullableAccount) Set(val *Account) {
	v.value = val
	v.isSet = true
}

func (v NullableAccount) IsSet() bool {
	return v.isSet
}

func (v *NullableAccount) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccount(val *Account) *NullableAccount {
	return &NullableAccount{value: val, isSet: true}
}

func (v NullableAccount) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccount) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}