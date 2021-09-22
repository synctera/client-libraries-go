# AchOutgoing

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountId** | **string** | ID of the source account | [readonly] 
**Amount** | **int32** | Amount to transfer in ISO 4217 minor currency units | 
**Currency** | **string** | ISO 4217 alphabetic currency code of the transfer amount | 
**ExecutionDate** | Pointer to **string** | The date the transfer executes (default today) | [optional] 
**Id** | Pointer to **string** | Outgoing ACH | [optional] [readonly] 
**NoFraudCheck** | Pointer to **bool** | Do not perform a fraud check | [optional] 
**OverwriteChecks** | Pointer to **bool** | Overwrite non mandatory posting checks. Mandatory checks will still be processed | [optional] 
**RecipientName** | **string** | The name of the recipient | 
**RecurringData** | Pointer to [**RecurrenceData**](RecurrenceData.md) |  | [optional] 
**ReferenceInfo** | Pointer to **string** | Reference information for the payment | [optional] 
**TargetAccountNo** | **string** | The account number of the destination account | 
**TargetAccountRouting** | **string** | The routing number of the destination account | 
**TargetBankCountry** | **string** | The ISO-3166-1 Alpha-2 country code in which the target account is registered (default US) | [default to "US"]
**TransactionDirection** | Pointer to **string** | The type of transaction (DEBIT/CREDIT) for originating account | [optional] 

## Methods

### NewAchOutgoing

`func NewAchOutgoing(accountId string, amount int32, currency string, recipientName string, targetAccountNo string, targetAccountRouting string, targetBankCountry string, ) *AchOutgoing`

NewAchOutgoing instantiates a new AchOutgoing object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAchOutgoingWithDefaults

`func NewAchOutgoingWithDefaults() *AchOutgoing`

NewAchOutgoingWithDefaults instantiates a new AchOutgoing object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountId

`func (o *AchOutgoing) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *AchOutgoing) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *AchOutgoing) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.


### GetAmount

`func (o *AchOutgoing) GetAmount() int32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *AchOutgoing) GetAmountOk() (*int32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *AchOutgoing) SetAmount(v int32)`

SetAmount sets Amount field to given value.


### GetCurrency

`func (o *AchOutgoing) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *AchOutgoing) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *AchOutgoing) SetCurrency(v string)`

SetCurrency sets Currency field to given value.


### GetExecutionDate

`func (o *AchOutgoing) GetExecutionDate() string`

GetExecutionDate returns the ExecutionDate field if non-nil, zero value otherwise.

### GetExecutionDateOk

`func (o *AchOutgoing) GetExecutionDateOk() (*string, bool)`

GetExecutionDateOk returns a tuple with the ExecutionDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutionDate

`func (o *AchOutgoing) SetExecutionDate(v string)`

SetExecutionDate sets ExecutionDate field to given value.

### HasExecutionDate

`func (o *AchOutgoing) HasExecutionDate() bool`

HasExecutionDate returns a boolean if a field has been set.

### GetId

`func (o *AchOutgoing) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AchOutgoing) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AchOutgoing) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *AchOutgoing) HasId() bool`

HasId returns a boolean if a field has been set.

### GetNoFraudCheck

`func (o *AchOutgoing) GetNoFraudCheck() bool`

GetNoFraudCheck returns the NoFraudCheck field if non-nil, zero value otherwise.

### GetNoFraudCheckOk

`func (o *AchOutgoing) GetNoFraudCheckOk() (*bool, bool)`

GetNoFraudCheckOk returns a tuple with the NoFraudCheck field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoFraudCheck

`func (o *AchOutgoing) SetNoFraudCheck(v bool)`

SetNoFraudCheck sets NoFraudCheck field to given value.

### HasNoFraudCheck

`func (o *AchOutgoing) HasNoFraudCheck() bool`

HasNoFraudCheck returns a boolean if a field has been set.

### GetOverwriteChecks

`func (o *AchOutgoing) GetOverwriteChecks() bool`

GetOverwriteChecks returns the OverwriteChecks field if non-nil, zero value otherwise.

### GetOverwriteChecksOk

`func (o *AchOutgoing) GetOverwriteChecksOk() (*bool, bool)`

GetOverwriteChecksOk returns a tuple with the OverwriteChecks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverwriteChecks

`func (o *AchOutgoing) SetOverwriteChecks(v bool)`

SetOverwriteChecks sets OverwriteChecks field to given value.

### HasOverwriteChecks

`func (o *AchOutgoing) HasOverwriteChecks() bool`

HasOverwriteChecks returns a boolean if a field has been set.

### GetRecipientName

`func (o *AchOutgoing) GetRecipientName() string`

GetRecipientName returns the RecipientName field if non-nil, zero value otherwise.

### GetRecipientNameOk

`func (o *AchOutgoing) GetRecipientNameOk() (*string, bool)`

GetRecipientNameOk returns a tuple with the RecipientName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipientName

`func (o *AchOutgoing) SetRecipientName(v string)`

SetRecipientName sets RecipientName field to given value.


### GetRecurringData

`func (o *AchOutgoing) GetRecurringData() RecurrenceData`

GetRecurringData returns the RecurringData field if non-nil, zero value otherwise.

### GetRecurringDataOk

`func (o *AchOutgoing) GetRecurringDataOk() (*RecurrenceData, bool)`

GetRecurringDataOk returns a tuple with the RecurringData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecurringData

`func (o *AchOutgoing) SetRecurringData(v RecurrenceData)`

SetRecurringData sets RecurringData field to given value.

### HasRecurringData

`func (o *AchOutgoing) HasRecurringData() bool`

HasRecurringData returns a boolean if a field has been set.

### GetReferenceInfo

`func (o *AchOutgoing) GetReferenceInfo() string`

GetReferenceInfo returns the ReferenceInfo field if non-nil, zero value otherwise.

### GetReferenceInfoOk

`func (o *AchOutgoing) GetReferenceInfoOk() (*string, bool)`

GetReferenceInfoOk returns a tuple with the ReferenceInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceInfo

`func (o *AchOutgoing) SetReferenceInfo(v string)`

SetReferenceInfo sets ReferenceInfo field to given value.

### HasReferenceInfo

`func (o *AchOutgoing) HasReferenceInfo() bool`

HasReferenceInfo returns a boolean if a field has been set.

### GetTargetAccountNo

`func (o *AchOutgoing) GetTargetAccountNo() string`

GetTargetAccountNo returns the TargetAccountNo field if non-nil, zero value otherwise.

### GetTargetAccountNoOk

`func (o *AchOutgoing) GetTargetAccountNoOk() (*string, bool)`

GetTargetAccountNoOk returns a tuple with the TargetAccountNo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetAccountNo

`func (o *AchOutgoing) SetTargetAccountNo(v string)`

SetTargetAccountNo sets TargetAccountNo field to given value.


### GetTargetAccountRouting

`func (o *AchOutgoing) GetTargetAccountRouting() string`

GetTargetAccountRouting returns the TargetAccountRouting field if non-nil, zero value otherwise.

### GetTargetAccountRoutingOk

`func (o *AchOutgoing) GetTargetAccountRoutingOk() (*string, bool)`

GetTargetAccountRoutingOk returns a tuple with the TargetAccountRouting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetAccountRouting

`func (o *AchOutgoing) SetTargetAccountRouting(v string)`

SetTargetAccountRouting sets TargetAccountRouting field to given value.


### GetTargetBankCountry

`func (o *AchOutgoing) GetTargetBankCountry() string`

GetTargetBankCountry returns the TargetBankCountry field if non-nil, zero value otherwise.

### GetTargetBankCountryOk

`func (o *AchOutgoing) GetTargetBankCountryOk() (*string, bool)`

GetTargetBankCountryOk returns a tuple with the TargetBankCountry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetBankCountry

`func (o *AchOutgoing) SetTargetBankCountry(v string)`

SetTargetBankCountry sets TargetBankCountry field to given value.


### GetTransactionDirection

`func (o *AchOutgoing) GetTransactionDirection() string`

GetTransactionDirection returns the TransactionDirection field if non-nil, zero value otherwise.

### GetTransactionDirectionOk

`func (o *AchOutgoing) GetTransactionDirectionOk() (*string, bool)`

GetTransactionDirectionOk returns a tuple with the TransactionDirection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionDirection

`func (o *AchOutgoing) SetTransactionDirection(v string)`

SetTransactionDirection sets TransactionDirection field to given value.

### HasTransactionDirection

`func (o *AchOutgoing) HasTransactionDirection() bool`

HasTransactionDirection returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


