# ExternalAccountTransaction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | Pointer to **int64** | Transaction amount. Number in cents. E.g. 1000 represents $10.00 | [optional] 
**AuthorizedDate** | Pointer to **string** | Date that the transaction is authorized. ISO 8601 format ( YYYY-MM-DD ). | [optional] 
**Category** | Pointer to **[]string** | Category of the transaction | [optional] 
**CheckNumber** | Pointer to **string** | Check number of the transaction. This field will be null if not a check transaction. | [optional] 
**Currency** | Pointer to **string** | ISO 4217 alphabetic currency code | [optional] 
**Date** | Pointer to **string** | For pending transactions, this represents the date of the transaction occurred; for posted transactions, this represents the date of the transaction posted. ISO 8601 format ( YYYY-MM-DD ).  | [optional] 
**IsPending** | Pointer to **bool** | Indicates the transaction is pending or unsettled if true. | [optional] 
**MerchantName** | Pointer to **string** | Merchant name of the transaction | [optional] 
**PaymentChannel** | Pointer to **string** | channel used to make a payment | [optional] 
**PaymentMethod** | Pointer to **string** | Transfer type of the transaction, e.g. ACH | [optional] 
**TransactionId** | Pointer to **string** | case-sensitive transaction ID | [optional] 

## Methods

### NewExternalAccountTransaction

`func NewExternalAccountTransaction() *ExternalAccountTransaction`

NewExternalAccountTransaction instantiates a new ExternalAccountTransaction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExternalAccountTransactionWithDefaults

`func NewExternalAccountTransactionWithDefaults() *ExternalAccountTransaction`

NewExternalAccountTransactionWithDefaults instantiates a new ExternalAccountTransaction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmount

`func (o *ExternalAccountTransaction) GetAmount() int64`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *ExternalAccountTransaction) GetAmountOk() (*int64, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *ExternalAccountTransaction) SetAmount(v int64)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *ExternalAccountTransaction) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### GetAuthorizedDate

`func (o *ExternalAccountTransaction) GetAuthorizedDate() string`

GetAuthorizedDate returns the AuthorizedDate field if non-nil, zero value otherwise.

### GetAuthorizedDateOk

`func (o *ExternalAccountTransaction) GetAuthorizedDateOk() (*string, bool)`

GetAuthorizedDateOk returns a tuple with the AuthorizedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorizedDate

`func (o *ExternalAccountTransaction) SetAuthorizedDate(v string)`

SetAuthorizedDate sets AuthorizedDate field to given value.

### HasAuthorizedDate

`func (o *ExternalAccountTransaction) HasAuthorizedDate() bool`

HasAuthorizedDate returns a boolean if a field has been set.

### GetCategory

`func (o *ExternalAccountTransaction) GetCategory() []string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *ExternalAccountTransaction) GetCategoryOk() (*[]string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *ExternalAccountTransaction) SetCategory(v []string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *ExternalAccountTransaction) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### GetCheckNumber

`func (o *ExternalAccountTransaction) GetCheckNumber() string`

GetCheckNumber returns the CheckNumber field if non-nil, zero value otherwise.

### GetCheckNumberOk

`func (o *ExternalAccountTransaction) GetCheckNumberOk() (*string, bool)`

GetCheckNumberOk returns a tuple with the CheckNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckNumber

`func (o *ExternalAccountTransaction) SetCheckNumber(v string)`

SetCheckNumber sets CheckNumber field to given value.

### HasCheckNumber

`func (o *ExternalAccountTransaction) HasCheckNumber() bool`

HasCheckNumber returns a boolean if a field has been set.

### GetCurrency

`func (o *ExternalAccountTransaction) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *ExternalAccountTransaction) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *ExternalAccountTransaction) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *ExternalAccountTransaction) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetDate

`func (o *ExternalAccountTransaction) GetDate() string`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *ExternalAccountTransaction) GetDateOk() (*string, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *ExternalAccountTransaction) SetDate(v string)`

SetDate sets Date field to given value.

### HasDate

`func (o *ExternalAccountTransaction) HasDate() bool`

HasDate returns a boolean if a field has been set.

### GetIsPending

`func (o *ExternalAccountTransaction) GetIsPending() bool`

GetIsPending returns the IsPending field if non-nil, zero value otherwise.

### GetIsPendingOk

`func (o *ExternalAccountTransaction) GetIsPendingOk() (*bool, bool)`

GetIsPendingOk returns a tuple with the IsPending field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPending

`func (o *ExternalAccountTransaction) SetIsPending(v bool)`

SetIsPending sets IsPending field to given value.

### HasIsPending

`func (o *ExternalAccountTransaction) HasIsPending() bool`

HasIsPending returns a boolean if a field has been set.

### GetMerchantName

`func (o *ExternalAccountTransaction) GetMerchantName() string`

GetMerchantName returns the MerchantName field if non-nil, zero value otherwise.

### GetMerchantNameOk

`func (o *ExternalAccountTransaction) GetMerchantNameOk() (*string, bool)`

GetMerchantNameOk returns a tuple with the MerchantName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantName

`func (o *ExternalAccountTransaction) SetMerchantName(v string)`

SetMerchantName sets MerchantName field to given value.

### HasMerchantName

`func (o *ExternalAccountTransaction) HasMerchantName() bool`

HasMerchantName returns a boolean if a field has been set.

### GetPaymentChannel

`func (o *ExternalAccountTransaction) GetPaymentChannel() string`

GetPaymentChannel returns the PaymentChannel field if non-nil, zero value otherwise.

### GetPaymentChannelOk

`func (o *ExternalAccountTransaction) GetPaymentChannelOk() (*string, bool)`

GetPaymentChannelOk returns a tuple with the PaymentChannel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentChannel

`func (o *ExternalAccountTransaction) SetPaymentChannel(v string)`

SetPaymentChannel sets PaymentChannel field to given value.

### HasPaymentChannel

`func (o *ExternalAccountTransaction) HasPaymentChannel() bool`

HasPaymentChannel returns a boolean if a field has been set.

### GetPaymentMethod

`func (o *ExternalAccountTransaction) GetPaymentMethod() string`

GetPaymentMethod returns the PaymentMethod field if non-nil, zero value otherwise.

### GetPaymentMethodOk

`func (o *ExternalAccountTransaction) GetPaymentMethodOk() (*string, bool)`

GetPaymentMethodOk returns a tuple with the PaymentMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentMethod

`func (o *ExternalAccountTransaction) SetPaymentMethod(v string)`

SetPaymentMethod sets PaymentMethod field to given value.

### HasPaymentMethod

`func (o *ExternalAccountTransaction) HasPaymentMethod() bool`

HasPaymentMethod returns a boolean if a field has been set.

### GetTransactionId

`func (o *ExternalAccountTransaction) GetTransactionId() string`

GetTransactionId returns the TransactionId field if non-nil, zero value otherwise.

### GetTransactionIdOk

`func (o *ExternalAccountTransaction) GetTransactionIdOk() (*string, bool)`

GetTransactionIdOk returns a tuple with the TransactionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionId

`func (o *ExternalAccountTransaction) SetTransactionId(v string)`

SetTransactionId sets TransactionId field to given value.

### HasTransactionId

`func (o *ExternalAccountTransaction) HasTransactionId() bool`

HasTransactionId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


