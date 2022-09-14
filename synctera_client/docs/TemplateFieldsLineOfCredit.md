# TemplateFieldsLineOfCredit

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountType** | [**AccountType**](AccountType.md) |  | 
**BankCountry** | **string** | Bank country of the account. ISO 3166-1 Alpha-2 or Alpha-3 country code. | 
**Currency** | **string** | Account currency. ISO 4217 alphabetic currency code | 
**ChargeoffPeriod** | Pointer to **int32** | The number of days an account can stay delinquent before marking an account as charged-off.  | [optional] [default to 90]
**DelinquencyPeriod** | Pointer to **int32** | The number of days past the due date to wait for a minimum payment before marking an account as delinquent.  | [optional] [default to 30]
**GracePeriod** | Pointer to **int32** | The number of days past the billing period to allow for payment before it is considered due. This directly infers the due date for a payment.  | [optional] [default to 30]
**InterestProductId** | Pointer to **string** | An interest account product that the current account associates with. The account product must have its calculation_method set to COMPOUNDED_DAILY.  | [optional] 
**MinimumPayment** | [**MinimumPayment**](MinimumPayment.md) |  | 

## Methods

### NewTemplateFieldsLineOfCredit

`func NewTemplateFieldsLineOfCredit(accountType AccountType, bankCountry string, currency string, minimumPayment MinimumPayment, ) *TemplateFieldsLineOfCredit`

NewTemplateFieldsLineOfCredit instantiates a new TemplateFieldsLineOfCredit object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTemplateFieldsLineOfCreditWithDefaults

`func NewTemplateFieldsLineOfCreditWithDefaults() *TemplateFieldsLineOfCredit`

NewTemplateFieldsLineOfCreditWithDefaults instantiates a new TemplateFieldsLineOfCredit object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountType

`func (o *TemplateFieldsLineOfCredit) GetAccountType() AccountType`

GetAccountType returns the AccountType field if non-nil, zero value otherwise.

### GetAccountTypeOk

`func (o *TemplateFieldsLineOfCredit) GetAccountTypeOk() (*AccountType, bool)`

GetAccountTypeOk returns a tuple with the AccountType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountType

`func (o *TemplateFieldsLineOfCredit) SetAccountType(v AccountType)`

SetAccountType sets AccountType field to given value.


### GetBankCountry

`func (o *TemplateFieldsLineOfCredit) GetBankCountry() string`

GetBankCountry returns the BankCountry field if non-nil, zero value otherwise.

### GetBankCountryOk

`func (o *TemplateFieldsLineOfCredit) GetBankCountryOk() (*string, bool)`

GetBankCountryOk returns a tuple with the BankCountry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBankCountry

`func (o *TemplateFieldsLineOfCredit) SetBankCountry(v string)`

SetBankCountry sets BankCountry field to given value.


### GetCurrency

`func (o *TemplateFieldsLineOfCredit) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *TemplateFieldsLineOfCredit) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *TemplateFieldsLineOfCredit) SetCurrency(v string)`

SetCurrency sets Currency field to given value.


### GetChargeoffPeriod

`func (o *TemplateFieldsLineOfCredit) GetChargeoffPeriod() int32`

GetChargeoffPeriod returns the ChargeoffPeriod field if non-nil, zero value otherwise.

### GetChargeoffPeriodOk

`func (o *TemplateFieldsLineOfCredit) GetChargeoffPeriodOk() (*int32, bool)`

GetChargeoffPeriodOk returns a tuple with the ChargeoffPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChargeoffPeriod

`func (o *TemplateFieldsLineOfCredit) SetChargeoffPeriod(v int32)`

SetChargeoffPeriod sets ChargeoffPeriod field to given value.

### HasChargeoffPeriod

`func (o *TemplateFieldsLineOfCredit) HasChargeoffPeriod() bool`

HasChargeoffPeriod returns a boolean if a field has been set.

### GetDelinquencyPeriod

`func (o *TemplateFieldsLineOfCredit) GetDelinquencyPeriod() int32`

GetDelinquencyPeriod returns the DelinquencyPeriod field if non-nil, zero value otherwise.

### GetDelinquencyPeriodOk

`func (o *TemplateFieldsLineOfCredit) GetDelinquencyPeriodOk() (*int32, bool)`

GetDelinquencyPeriodOk returns a tuple with the DelinquencyPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelinquencyPeriod

`func (o *TemplateFieldsLineOfCredit) SetDelinquencyPeriod(v int32)`

SetDelinquencyPeriod sets DelinquencyPeriod field to given value.

### HasDelinquencyPeriod

`func (o *TemplateFieldsLineOfCredit) HasDelinquencyPeriod() bool`

HasDelinquencyPeriod returns a boolean if a field has been set.

### GetGracePeriod

`func (o *TemplateFieldsLineOfCredit) GetGracePeriod() int32`

GetGracePeriod returns the GracePeriod field if non-nil, zero value otherwise.

### GetGracePeriodOk

`func (o *TemplateFieldsLineOfCredit) GetGracePeriodOk() (*int32, bool)`

GetGracePeriodOk returns a tuple with the GracePeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGracePeriod

`func (o *TemplateFieldsLineOfCredit) SetGracePeriod(v int32)`

SetGracePeriod sets GracePeriod field to given value.

### HasGracePeriod

`func (o *TemplateFieldsLineOfCredit) HasGracePeriod() bool`

HasGracePeriod returns a boolean if a field has been set.

### GetInterestProductId

`func (o *TemplateFieldsLineOfCredit) GetInterestProductId() string`

GetInterestProductId returns the InterestProductId field if non-nil, zero value otherwise.

### GetInterestProductIdOk

`func (o *TemplateFieldsLineOfCredit) GetInterestProductIdOk() (*string, bool)`

GetInterestProductIdOk returns a tuple with the InterestProductId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterestProductId

`func (o *TemplateFieldsLineOfCredit) SetInterestProductId(v string)`

SetInterestProductId sets InterestProductId field to given value.

### HasInterestProductId

`func (o *TemplateFieldsLineOfCredit) HasInterestProductId() bool`

HasInterestProductId returns a boolean if a field has been set.

### GetMinimumPayment

`func (o *TemplateFieldsLineOfCredit) GetMinimumPayment() MinimumPayment`

GetMinimumPayment returns the MinimumPayment field if non-nil, zero value otherwise.

### GetMinimumPaymentOk

`func (o *TemplateFieldsLineOfCredit) GetMinimumPaymentOk() (*MinimumPayment, bool)`

GetMinimumPaymentOk returns a tuple with the MinimumPayment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinimumPayment

`func (o *TemplateFieldsLineOfCredit) SetMinimumPayment(v MinimumPayment)`

SetMinimumPayment sets MinimumPayment field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


