/*
Synctera API

<h2>Let's build something great.</h2><p>Welcome to the official reference documentation for Synctera APIs. Our APIs are the best way to automate your company's banking needs and are designed to be easy to understand and implement.</p><p>We're continuously growing this library and what you see here is just the start, but if you need something specific or have a question, <a class='text-blue-600' href='https://synctera.com/contact' target='_blank' rel='noreferrer'>contact us</a>.</p>

API version: 0.20.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package synctera_client

import (
	"encoding/json"
	"time"
)

// AccountDepository Account representing either a checking or saving account.
type AccountDepository struct {
	AccessStatus *AccountAccessStatus `json:"access_status,omitempty"`
	// Account number
	AccountNumber *string `json:"account_number,omitempty"`
	// The response will contain the bank fintech ID (3 or 6 digits) plus the last 4 digits, with the digits in between replaced with * characters. Shadow mode account numbers will not be masked.
	AccountNumberMasked *string `json:"account_number_masked,omitempty"`
	// Purpose of the account
	AccountPurpose *string      `json:"account_purpose,omitempty"`
	AccountType    *AccountType `json:"account_type,omitempty"`
	// The application ID for this account.
	ApplicationId *string `json:"application_id,omitempty"`
	// A list of balances for account based on different type
	Balances []Balance `json:"balances,omitempty"`
	// Bank routing number
	BankRouting *string `json:"bank_routing,omitempty"`
	// Account creation timestamp in RFC3339 format
	CreationTime *time.Time `json:"creation_time,omitempty"`
	// Account currency or account settlement currency. ISO 4217 alphabetic currency code. Default USD
	Currency *string `json:"currency,omitempty"`
	// A list of the customer IDs of the account holders.
	CustomerIds  []string      `json:"customer_ids,omitempty"`
	CustomerType *CustomerType `json:"customer_type,omitempty"`
	// Exchange rate type
	ExchangeRateType *string `json:"exchange_rate_type,omitempty"`
	// International bank account number
	Iban *string `json:"iban,omitempty"`
	// Account ID
	Id *string `json:"id,omitempty"`
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
	Status   *Status `json:"status,omitempty"`
	// SWIFT code
	SwiftCode      *string         `json:"swift_code,omitempty"`
	BalanceCeiling *BalanceCeiling `json:"balance_ceiling,omitempty"`
	BalanceFloor   *BalanceFloor   `json:"balance_floor,omitempty"`
	// A list of fee account products that the current account associates with.
	FeeProductIds []string `json:"fee_product_ids,omitempty"`
	// An interest account product that the current account associates with. The account product must have its calculation_method set to COMPOUNDED_MONTHLY.
	InterestProductId *string `json:"interest_product_id,omitempty"`
	// Account's overdraft limit
	OverdraftLimit *int64          `json:"overdraft_limit,omitempty"`
	SpendingLimits *SpendingLimits `json:"spending_limits,omitempty"`
}

// NewAccountDepository instantiates a new AccountDepository object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccountDepository() *AccountDepository {
	this := AccountDepository{}
	return &this
}

// NewAccountDepositoryWithDefaults instantiates a new AccountDepository object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccountDepositoryWithDefaults() *AccountDepository {
	this := AccountDepository{}
	return &this
}

// GetAccessStatus returns the AccessStatus field value if set, zero value otherwise.
func (o *AccountDepository) GetAccessStatus() AccountAccessStatus {
	if o == nil || o.AccessStatus == nil {
		var ret AccountAccessStatus
		return ret
	}
	return *o.AccessStatus
}

// GetAccessStatusOk returns a tuple with the AccessStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountDepository) GetAccessStatusOk() (*AccountAccessStatus, bool) {
	if o == nil || o.AccessStatus == nil {
		return nil, false
	}
	return o.AccessStatus, true
}

// HasAccessStatus returns a boolean if a field has been set.
func (o *AccountDepository) HasAccessStatus() bool {
	if o != nil && o.AccessStatus != nil {
		return true
	}

	return false
}

// SetAccessStatus gets a reference to the given AccountAccessStatus and assigns it to the AccessStatus field.
func (o *AccountDepository) SetAccessStatus(v AccountAccessStatus) {
	o.AccessStatus = &v
}

// GetAccountNumber returns the AccountNumber field value if set, zero value otherwise.
func (o *AccountDepository) GetAccountNumber() string {
	if o == nil || o.AccountNumber == nil {
		var ret string
		return ret
	}
	return *o.AccountNumber
}

// GetAccountNumberOk returns a tuple with the AccountNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountDepository) GetAccountNumberOk() (*string, bool) {
	if o == nil || o.AccountNumber == nil {
		return nil, false
	}
	return o.AccountNumber, true
}

// HasAccountNumber returns a boolean if a field has been set.
func (o *AccountDepository) HasAccountNumber() bool {
	if o != nil && o.AccountNumber != nil {
		return true
	}

	return false
}

// SetAccountNumber gets a reference to the given string and assigns it to the AccountNumber field.
func (o *AccountDepository) SetAccountNumber(v string) {
	o.AccountNumber = &v
}

// GetAccountNumberMasked returns the AccountNumberMasked field value if set, zero value otherwise.
func (o *AccountDepository) GetAccountNumberMasked() string {
	if o == nil || o.AccountNumberMasked == nil {
		var ret string
		return ret
	}
	return *o.AccountNumberMasked
}

// GetAccountNumberMaskedOk returns a tuple with the AccountNumberMasked field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountDepository) GetAccountNumberMaskedOk() (*string, bool) {
	if o == nil || o.AccountNumberMasked == nil {
		return nil, false
	}
	return o.AccountNumberMasked, true
}

// HasAccountNumberMasked returns a boolean if a field has been set.
func (o *AccountDepository) HasAccountNumberMasked() bool {
	if o != nil && o.AccountNumberMasked != nil {
		return true
	}

	return false
}

// SetAccountNumberMasked gets a reference to the given string and assigns it to the AccountNumberMasked field.
func (o *AccountDepository) SetAccountNumberMasked(v string) {
	o.AccountNumberMasked = &v
}

// GetAccountPurpose returns the AccountPurpose field value if set, zero value otherwise.
func (o *AccountDepository) GetAccountPurpose() string {
	if o == nil || o.AccountPurpose == nil {
		var ret string
		return ret
	}
	return *o.AccountPurpose
}

// GetAccountPurposeOk returns a tuple with the AccountPurpose field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountDepository) GetAccountPurposeOk() (*string, bool) {
	if o == nil || o.AccountPurpose == nil {
		return nil, false
	}
	return o.AccountPurpose, true
}

// HasAccountPurpose returns a boolean if a field has been set.
func (o *AccountDepository) HasAccountPurpose() bool {
	if o != nil && o.AccountPurpose != nil {
		return true
	}

	return false
}

// SetAccountPurpose gets a reference to the given string and assigns it to the AccountPurpose field.
func (o *AccountDepository) SetAccountPurpose(v string) {
	o.AccountPurpose = &v
}

// GetAccountType returns the AccountType field value if set, zero value otherwise.
func (o *AccountDepository) GetAccountType() AccountType {
	if o == nil || o.AccountType == nil {
		var ret AccountType
		return ret
	}
	return *o.AccountType
}

// GetAccountTypeOk returns a tuple with the AccountType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountDepository) GetAccountTypeOk() (*AccountType, bool) {
	if o == nil || o.AccountType == nil {
		return nil, false
	}
	return o.AccountType, true
}

// HasAccountType returns a boolean if a field has been set.
func (o *AccountDepository) HasAccountType() bool {
	if o != nil && o.AccountType != nil {
		return true
	}

	return false
}

// SetAccountType gets a reference to the given AccountType and assigns it to the AccountType field.
func (o *AccountDepository) SetAccountType(v AccountType) {
	o.AccountType = &v
}

// GetApplicationId returns the ApplicationId field value if set, zero value otherwise.
func (o *AccountDepository) GetApplicationId() string {
	if o == nil || o.ApplicationId == nil {
		var ret string
		return ret
	}
	return *o.ApplicationId
}

// GetApplicationIdOk returns a tuple with the ApplicationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountDepository) GetApplicationIdOk() (*string, bool) {
	if o == nil || o.ApplicationId == nil {
		return nil, false
	}
	return o.ApplicationId, true
}

// HasApplicationId returns a boolean if a field has been set.
func (o *AccountDepository) HasApplicationId() bool {
	if o != nil && o.ApplicationId != nil {
		return true
	}

	return false
}

// SetApplicationId gets a reference to the given string and assigns it to the ApplicationId field.
func (o *AccountDepository) SetApplicationId(v string) {
	o.ApplicationId = &v
}

// GetBalances returns the Balances field value if set, zero value otherwise.
func (o *AccountDepository) GetBalances() []Balance {
	if o == nil || o.Balances == nil {
		var ret []Balance
		return ret
	}
	return o.Balances
}

// GetBalancesOk returns a tuple with the Balances field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountDepository) GetBalancesOk() ([]Balance, bool) {
	if o == nil || o.Balances == nil {
		return nil, false
	}
	return o.Balances, true
}

// HasBalances returns a boolean if a field has been set.
func (o *AccountDepository) HasBalances() bool {
	if o != nil && o.Balances != nil {
		return true
	}

	return false
}

// SetBalances gets a reference to the given []Balance and assigns it to the Balances field.
func (o *AccountDepository) SetBalances(v []Balance) {
	o.Balances = v
}

// GetBankRouting returns the BankRouting field value if set, zero value otherwise.
func (o *AccountDepository) GetBankRouting() string {
	if o == nil || o.BankRouting == nil {
		var ret string
		return ret
	}
	return *o.BankRouting
}

// GetBankRoutingOk returns a tuple with the BankRouting field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountDepository) GetBankRoutingOk() (*string, bool) {
	if o == nil || o.BankRouting == nil {
		return nil, false
	}
	return o.BankRouting, true
}

// HasBankRouting returns a boolean if a field has been set.
func (o *AccountDepository) HasBankRouting() bool {
	if o != nil && o.BankRouting != nil {
		return true
	}

	return false
}

// SetBankRouting gets a reference to the given string and assigns it to the BankRouting field.
func (o *AccountDepository) SetBankRouting(v string) {
	o.BankRouting = &v
}

// GetCreationTime returns the CreationTime field value if set, zero value otherwise.
func (o *AccountDepository) GetCreationTime() time.Time {
	if o == nil || o.CreationTime == nil {
		var ret time.Time
		return ret
	}
	return *o.CreationTime
}

// GetCreationTimeOk returns a tuple with the CreationTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountDepository) GetCreationTimeOk() (*time.Time, bool) {
	if o == nil || o.CreationTime == nil {
		return nil, false
	}
	return o.CreationTime, true
}

// HasCreationTime returns a boolean if a field has been set.
func (o *AccountDepository) HasCreationTime() bool {
	if o != nil && o.CreationTime != nil {
		return true
	}

	return false
}

// SetCreationTime gets a reference to the given time.Time and assigns it to the CreationTime field.
func (o *AccountDepository) SetCreationTime(v time.Time) {
	o.CreationTime = &v
}

// GetCurrency returns the Currency field value if set, zero value otherwise.
func (o *AccountDepository) GetCurrency() string {
	if o == nil || o.Currency == nil {
		var ret string
		return ret
	}
	return *o.Currency
}

// GetCurrencyOk returns a tuple with the Currency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountDepository) GetCurrencyOk() (*string, bool) {
	if o == nil || o.Currency == nil {
		return nil, false
	}
	return o.Currency, true
}

// HasCurrency returns a boolean if a field has been set.
func (o *AccountDepository) HasCurrency() bool {
	if o != nil && o.Currency != nil {
		return true
	}

	return false
}

// SetCurrency gets a reference to the given string and assigns it to the Currency field.
func (o *AccountDepository) SetCurrency(v string) {
	o.Currency = &v
}

// GetCustomerIds returns the CustomerIds field value if set, zero value otherwise.
func (o *AccountDepository) GetCustomerIds() []string {
	if o == nil || o.CustomerIds == nil {
		var ret []string
		return ret
	}
	return o.CustomerIds
}

// GetCustomerIdsOk returns a tuple with the CustomerIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountDepository) GetCustomerIdsOk() ([]string, bool) {
	if o == nil || o.CustomerIds == nil {
		return nil, false
	}
	return o.CustomerIds, true
}

// HasCustomerIds returns a boolean if a field has been set.
func (o *AccountDepository) HasCustomerIds() bool {
	if o != nil && o.CustomerIds != nil {
		return true
	}

	return false
}

// SetCustomerIds gets a reference to the given []string and assigns it to the CustomerIds field.
func (o *AccountDepository) SetCustomerIds(v []string) {
	o.CustomerIds = v
}

// GetCustomerType returns the CustomerType field value if set, zero value otherwise.
func (o *AccountDepository) GetCustomerType() CustomerType {
	if o == nil || o.CustomerType == nil {
		var ret CustomerType
		return ret
	}
	return *o.CustomerType
}

// GetCustomerTypeOk returns a tuple with the CustomerType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountDepository) GetCustomerTypeOk() (*CustomerType, bool) {
	if o == nil || o.CustomerType == nil {
		return nil, false
	}
	return o.CustomerType, true
}

// HasCustomerType returns a boolean if a field has been set.
func (o *AccountDepository) HasCustomerType() bool {
	if o != nil && o.CustomerType != nil {
		return true
	}

	return false
}

// SetCustomerType gets a reference to the given CustomerType and assigns it to the CustomerType field.
func (o *AccountDepository) SetCustomerType(v CustomerType) {
	o.CustomerType = &v
}

// GetExchangeRateType returns the ExchangeRateType field value if set, zero value otherwise.
func (o *AccountDepository) GetExchangeRateType() string {
	if o == nil || o.ExchangeRateType == nil {
		var ret string
		return ret
	}
	return *o.ExchangeRateType
}

// GetExchangeRateTypeOk returns a tuple with the ExchangeRateType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountDepository) GetExchangeRateTypeOk() (*string, bool) {
	if o == nil || o.ExchangeRateType == nil {
		return nil, false
	}
	return o.ExchangeRateType, true
}

// HasExchangeRateType returns a boolean if a field has been set.
func (o *AccountDepository) HasExchangeRateType() bool {
	if o != nil && o.ExchangeRateType != nil {
		return true
	}

	return false
}

// SetExchangeRateType gets a reference to the given string and assigns it to the ExchangeRateType field.
func (o *AccountDepository) SetExchangeRateType(v string) {
	o.ExchangeRateType = &v
}

// GetIban returns the Iban field value if set, zero value otherwise.
func (o *AccountDepository) GetIban() string {
	if o == nil || o.Iban == nil {
		var ret string
		return ret
	}
	return *o.Iban
}

// GetIbanOk returns a tuple with the Iban field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountDepository) GetIbanOk() (*string, bool) {
	if o == nil || o.Iban == nil {
		return nil, false
	}
	return o.Iban, true
}

// HasIban returns a boolean if a field has been set.
func (o *AccountDepository) HasIban() bool {
	if o != nil && o.Iban != nil {
		return true
	}

	return false
}

// SetIban gets a reference to the given string and assigns it to the Iban field.
func (o *AccountDepository) SetIban(v string) {
	o.Iban = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *AccountDepository) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountDepository) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *AccountDepository) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *AccountDepository) SetId(v string) {
	o.Id = &v
}

// GetIsAccountPool returns the IsAccountPool field value if set, zero value otherwise.
func (o *AccountDepository) GetIsAccountPool() bool {
	if o == nil || o.IsAccountPool == nil {
		var ret bool
		return ret
	}
	return *o.IsAccountPool
}

// GetIsAccountPoolOk returns a tuple with the IsAccountPool field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountDepository) GetIsAccountPoolOk() (*bool, bool) {
	if o == nil || o.IsAccountPool == nil {
		return nil, false
	}
	return o.IsAccountPool, true
}

// HasIsAccountPool returns a boolean if a field has been set.
func (o *AccountDepository) HasIsAccountPool() bool {
	if o != nil && o.IsAccountPool != nil {
		return true
	}

	return false
}

// SetIsAccountPool gets a reference to the given bool and assigns it to the IsAccountPool field.
func (o *AccountDepository) SetIsAccountPool(v bool) {
	o.IsAccountPool = &v
}

// GetIsAchEnabled returns the IsAchEnabled field value if set, zero value otherwise.
func (o *AccountDepository) GetIsAchEnabled() bool {
	if o == nil || o.IsAchEnabled == nil {
		var ret bool
		return ret
	}
	return *o.IsAchEnabled
}

// GetIsAchEnabledOk returns a tuple with the IsAchEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountDepository) GetIsAchEnabledOk() (*bool, bool) {
	if o == nil || o.IsAchEnabled == nil {
		return nil, false
	}
	return o.IsAchEnabled, true
}

// HasIsAchEnabled returns a boolean if a field has been set.
func (o *AccountDepository) HasIsAchEnabled() bool {
	if o != nil && o.IsAchEnabled != nil {
		return true
	}

	return false
}

// SetIsAchEnabled gets a reference to the given bool and assigns it to the IsAchEnabled field.
func (o *AccountDepository) SetIsAchEnabled(v bool) {
	o.IsAchEnabled = &v
}

// GetIsCardEnabled returns the IsCardEnabled field value if set, zero value otherwise.
func (o *AccountDepository) GetIsCardEnabled() bool {
	if o == nil || o.IsCardEnabled == nil {
		var ret bool
		return ret
	}
	return *o.IsCardEnabled
}

// GetIsCardEnabledOk returns a tuple with the IsCardEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountDepository) GetIsCardEnabledOk() (*bool, bool) {
	if o == nil || o.IsCardEnabled == nil {
		return nil, false
	}
	return o.IsCardEnabled, true
}

// HasIsCardEnabled returns a boolean if a field has been set.
func (o *AccountDepository) HasIsCardEnabled() bool {
	if o != nil && o.IsCardEnabled != nil {
		return true
	}

	return false
}

// SetIsCardEnabled gets a reference to the given bool and assigns it to the IsCardEnabled field.
func (o *AccountDepository) SetIsCardEnabled(v bool) {
	o.IsCardEnabled = &v
}

// GetIsP2pEnabled returns the IsP2pEnabled field value if set, zero value otherwise.
func (o *AccountDepository) GetIsP2pEnabled() bool {
	if o == nil || o.IsP2pEnabled == nil {
		var ret bool
		return ret
	}
	return *o.IsP2pEnabled
}

// GetIsP2pEnabledOk returns a tuple with the IsP2pEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountDepository) GetIsP2pEnabledOk() (*bool, bool) {
	if o == nil || o.IsP2pEnabled == nil {
		return nil, false
	}
	return o.IsP2pEnabled, true
}

// HasIsP2pEnabled returns a boolean if a field has been set.
func (o *AccountDepository) HasIsP2pEnabled() bool {
	if o != nil && o.IsP2pEnabled != nil {
		return true
	}

	return false
}

// SetIsP2pEnabled gets a reference to the given bool and assigns it to the IsP2pEnabled field.
func (o *AccountDepository) SetIsP2pEnabled(v bool) {
	o.IsP2pEnabled = &v
}

// GetIsWireEnabled returns the IsWireEnabled field value if set, zero value otherwise.
func (o *AccountDepository) GetIsWireEnabled() bool {
	if o == nil || o.IsWireEnabled == nil {
		var ret bool
		return ret
	}
	return *o.IsWireEnabled
}

// GetIsWireEnabledOk returns a tuple with the IsWireEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountDepository) GetIsWireEnabledOk() (*bool, bool) {
	if o == nil || o.IsWireEnabled == nil {
		return nil, false
	}
	return o.IsWireEnabled, true
}

// HasIsWireEnabled returns a boolean if a field has been set.
func (o *AccountDepository) HasIsWireEnabled() bool {
	if o != nil && o.IsWireEnabled != nil {
		return true
	}

	return false
}

// SetIsWireEnabled gets a reference to the given bool and assigns it to the IsWireEnabled field.
func (o *AccountDepository) SetIsWireEnabled(v bool) {
	o.IsWireEnabled = &v
}

// GetLastUpdatedTime returns the LastUpdatedTime field value if set, zero value otherwise.
func (o *AccountDepository) GetLastUpdatedTime() time.Time {
	if o == nil || o.LastUpdatedTime == nil {
		var ret time.Time
		return ret
	}
	return *o.LastUpdatedTime
}

// GetLastUpdatedTimeOk returns a tuple with the LastUpdatedTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountDepository) GetLastUpdatedTimeOk() (*time.Time, bool) {
	if o == nil || o.LastUpdatedTime == nil {
		return nil, false
	}
	return o.LastUpdatedTime, true
}

// HasLastUpdatedTime returns a boolean if a field has been set.
func (o *AccountDepository) HasLastUpdatedTime() bool {
	if o != nil && o.LastUpdatedTime != nil {
		return true
	}

	return false
}

// SetLastUpdatedTime gets a reference to the given time.Time and assigns it to the LastUpdatedTime field.
func (o *AccountDepository) SetLastUpdatedTime(v time.Time) {
	o.LastUpdatedTime = &v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *AccountDepository) GetMetadata() map[string]interface{} {
	if o == nil || o.Metadata == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountDepository) GetMetadataOk() (map[string]interface{}, bool) {
	if o == nil || o.Metadata == nil {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *AccountDepository) HasMetadata() bool {
	if o != nil && o.Metadata != nil {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given map[string]interface{} and assigns it to the Metadata field.
func (o *AccountDepository) SetMetadata(v map[string]interface{}) {
	o.Metadata = v
}

// GetNickname returns the Nickname field value if set, zero value otherwise.
func (o *AccountDepository) GetNickname() string {
	if o == nil || o.Nickname == nil {
		var ret string
		return ret
	}
	return *o.Nickname
}

// GetNicknameOk returns a tuple with the Nickname field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountDepository) GetNicknameOk() (*string, bool) {
	if o == nil || o.Nickname == nil {
		return nil, false
	}
	return o.Nickname, true
}

// HasNickname returns a boolean if a field has been set.
func (o *AccountDepository) HasNickname() bool {
	if o != nil && o.Nickname != nil {
		return true
	}

	return false
}

// SetNickname gets a reference to the given string and assigns it to the Nickname field.
func (o *AccountDepository) SetNickname(v string) {
	o.Nickname = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *AccountDepository) GetStatus() Status {
	if o == nil || o.Status == nil {
		var ret Status
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountDepository) GetStatusOk() (*Status, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *AccountDepository) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given Status and assigns it to the Status field.
func (o *AccountDepository) SetStatus(v Status) {
	o.Status = &v
}

// GetSwiftCode returns the SwiftCode field value if set, zero value otherwise.
func (o *AccountDepository) GetSwiftCode() string {
	if o == nil || o.SwiftCode == nil {
		var ret string
		return ret
	}
	return *o.SwiftCode
}

// GetSwiftCodeOk returns a tuple with the SwiftCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountDepository) GetSwiftCodeOk() (*string, bool) {
	if o == nil || o.SwiftCode == nil {
		return nil, false
	}
	return o.SwiftCode, true
}

// HasSwiftCode returns a boolean if a field has been set.
func (o *AccountDepository) HasSwiftCode() bool {
	if o != nil && o.SwiftCode != nil {
		return true
	}

	return false
}

// SetSwiftCode gets a reference to the given string and assigns it to the SwiftCode field.
func (o *AccountDepository) SetSwiftCode(v string) {
	o.SwiftCode = &v
}

// GetBalanceCeiling returns the BalanceCeiling field value if set, zero value otherwise.
func (o *AccountDepository) GetBalanceCeiling() BalanceCeiling {
	if o == nil || o.BalanceCeiling == nil {
		var ret BalanceCeiling
		return ret
	}
	return *o.BalanceCeiling
}

// GetBalanceCeilingOk returns a tuple with the BalanceCeiling field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountDepository) GetBalanceCeilingOk() (*BalanceCeiling, bool) {
	if o == nil || o.BalanceCeiling == nil {
		return nil, false
	}
	return o.BalanceCeiling, true
}

// HasBalanceCeiling returns a boolean if a field has been set.
func (o *AccountDepository) HasBalanceCeiling() bool {
	if o != nil && o.BalanceCeiling != nil {
		return true
	}

	return false
}

// SetBalanceCeiling gets a reference to the given BalanceCeiling and assigns it to the BalanceCeiling field.
func (o *AccountDepository) SetBalanceCeiling(v BalanceCeiling) {
	o.BalanceCeiling = &v
}

// GetBalanceFloor returns the BalanceFloor field value if set, zero value otherwise.
func (o *AccountDepository) GetBalanceFloor() BalanceFloor {
	if o == nil || o.BalanceFloor == nil {
		var ret BalanceFloor
		return ret
	}
	return *o.BalanceFloor
}

// GetBalanceFloorOk returns a tuple with the BalanceFloor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountDepository) GetBalanceFloorOk() (*BalanceFloor, bool) {
	if o == nil || o.BalanceFloor == nil {
		return nil, false
	}
	return o.BalanceFloor, true
}

// HasBalanceFloor returns a boolean if a field has been set.
func (o *AccountDepository) HasBalanceFloor() bool {
	if o != nil && o.BalanceFloor != nil {
		return true
	}

	return false
}

// SetBalanceFloor gets a reference to the given BalanceFloor and assigns it to the BalanceFloor field.
func (o *AccountDepository) SetBalanceFloor(v BalanceFloor) {
	o.BalanceFloor = &v
}

// GetFeeProductIds returns the FeeProductIds field value if set, zero value otherwise.
func (o *AccountDepository) GetFeeProductIds() []string {
	if o == nil || o.FeeProductIds == nil {
		var ret []string
		return ret
	}
	return o.FeeProductIds
}

// GetFeeProductIdsOk returns a tuple with the FeeProductIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountDepository) GetFeeProductIdsOk() ([]string, bool) {
	if o == nil || o.FeeProductIds == nil {
		return nil, false
	}
	return o.FeeProductIds, true
}

// HasFeeProductIds returns a boolean if a field has been set.
func (o *AccountDepository) HasFeeProductIds() bool {
	if o != nil && o.FeeProductIds != nil {
		return true
	}

	return false
}

// SetFeeProductIds gets a reference to the given []string and assigns it to the FeeProductIds field.
func (o *AccountDepository) SetFeeProductIds(v []string) {
	o.FeeProductIds = v
}

// GetInterestProductId returns the InterestProductId field value if set, zero value otherwise.
func (o *AccountDepository) GetInterestProductId() string {
	if o == nil || o.InterestProductId == nil {
		var ret string
		return ret
	}
	return *o.InterestProductId
}

// GetInterestProductIdOk returns a tuple with the InterestProductId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountDepository) GetInterestProductIdOk() (*string, bool) {
	if o == nil || o.InterestProductId == nil {
		return nil, false
	}
	return o.InterestProductId, true
}

// HasInterestProductId returns a boolean if a field has been set.
func (o *AccountDepository) HasInterestProductId() bool {
	if o != nil && o.InterestProductId != nil {
		return true
	}

	return false
}

// SetInterestProductId gets a reference to the given string and assigns it to the InterestProductId field.
func (o *AccountDepository) SetInterestProductId(v string) {
	o.InterestProductId = &v
}

// GetOverdraftLimit returns the OverdraftLimit field value if set, zero value otherwise.
func (o *AccountDepository) GetOverdraftLimit() int64 {
	if o == nil || o.OverdraftLimit == nil {
		var ret int64
		return ret
	}
	return *o.OverdraftLimit
}

// GetOverdraftLimitOk returns a tuple with the OverdraftLimit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountDepository) GetOverdraftLimitOk() (*int64, bool) {
	if o == nil || o.OverdraftLimit == nil {
		return nil, false
	}
	return o.OverdraftLimit, true
}

// HasOverdraftLimit returns a boolean if a field has been set.
func (o *AccountDepository) HasOverdraftLimit() bool {
	if o != nil && o.OverdraftLimit != nil {
		return true
	}

	return false
}

// SetOverdraftLimit gets a reference to the given int64 and assigns it to the OverdraftLimit field.
func (o *AccountDepository) SetOverdraftLimit(v int64) {
	o.OverdraftLimit = &v
}

// GetSpendingLimits returns the SpendingLimits field value if set, zero value otherwise.
func (o *AccountDepository) GetSpendingLimits() SpendingLimits {
	if o == nil || o.SpendingLimits == nil {
		var ret SpendingLimits
		return ret
	}
	return *o.SpendingLimits
}

// GetSpendingLimitsOk returns a tuple with the SpendingLimits field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountDepository) GetSpendingLimitsOk() (*SpendingLimits, bool) {
	if o == nil || o.SpendingLimits == nil {
		return nil, false
	}
	return o.SpendingLimits, true
}

// HasSpendingLimits returns a boolean if a field has been set.
func (o *AccountDepository) HasSpendingLimits() bool {
	if o != nil && o.SpendingLimits != nil {
		return true
	}

	return false
}

// SetSpendingLimits gets a reference to the given SpendingLimits and assigns it to the SpendingLimits field.
func (o *AccountDepository) SetSpendingLimits(v SpendingLimits) {
	o.SpendingLimits = &v
}

func (o AccountDepository) MarshalJSON() ([]byte, error) {
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
	if o.ApplicationId != nil {
		toSerialize["application_id"] = o.ApplicationId
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
	if o.Iban != nil {
		toSerialize["iban"] = o.Iban
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
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
	if o.BalanceCeiling != nil {
		toSerialize["balance_ceiling"] = o.BalanceCeiling
	}
	if o.BalanceFloor != nil {
		toSerialize["balance_floor"] = o.BalanceFloor
	}
	if o.FeeProductIds != nil {
		toSerialize["fee_product_ids"] = o.FeeProductIds
	}
	if o.InterestProductId != nil {
		toSerialize["interest_product_id"] = o.InterestProductId
	}
	if o.OverdraftLimit != nil {
		toSerialize["overdraft_limit"] = o.OverdraftLimit
	}
	if o.SpendingLimits != nil {
		toSerialize["spending_limits"] = o.SpendingLimits
	}
	return json.Marshal(toSerialize)
}

type NullableAccountDepository struct {
	value *AccountDepository
	isSet bool
}

func (v NullableAccountDepository) Get() *AccountDepository {
	return v.value
}

func (v *NullableAccountDepository) Set(val *AccountDepository) {
	v.value = val
	v.isSet = true
}

func (v NullableAccountDepository) IsSet() bool {
	return v.isSet
}

func (v *NullableAccountDepository) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccountDepository(val *AccountDepository) *NullableAccountDepository {
	return &NullableAccountDepository{value: val, isSet: true}
}

func (v NullableAccountDepository) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccountDepository) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
