# Transaction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountId** | Pointer to **string** | Account ID | [optional] 
**Amount** | Pointer to **int32** | Transaction amount in ISO 4217 minor currency units | [optional] 
**AuthorizationRef** | Pointer to **string** | Reference for the authorization | [optional] 
**Currency** | Pointer to **string** | Currency of the transaction. ISO 4217 alphabetic currency code | [optional] 
**DcSign** | Pointer to [**DcSignType**](DcSignType.md) |  | [optional] 
**EffectiveDate** | Pointer to **time.Time** | The effective date of the transaction (value_date) | [optional] 
**ExtReference** | Pointer to **string** | External reference from Synctera for the transaction | [optional] 
**Id** | Pointer to **string** | Transaction ID | [optional] 
**ProfitCenter** | Pointer to **string** | Profit center of the transaction | [optional] 
**Status** | Pointer to **string** | The status of the transaction | [optional] 
**TransactionType** | Pointer to **string** | Transaction type | [optional] 

## Methods

### NewTransaction

`func NewTransaction() *Transaction`

NewTransaction instantiates a new Transaction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransactionWithDefaults

`func NewTransactionWithDefaults() *Transaction`

NewTransactionWithDefaults instantiates a new Transaction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountId

`func (o *Transaction) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *Transaction) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *Transaction) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.

### HasAccountId

`func (o *Transaction) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.

### GetAmount

`func (o *Transaction) GetAmount() int32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *Transaction) GetAmountOk() (*int32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *Transaction) SetAmount(v int32)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *Transaction) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### GetAuthorizationRef

`func (o *Transaction) GetAuthorizationRef() string`

GetAuthorizationRef returns the AuthorizationRef field if non-nil, zero value otherwise.

### GetAuthorizationRefOk

`func (o *Transaction) GetAuthorizationRefOk() (*string, bool)`

GetAuthorizationRefOk returns a tuple with the AuthorizationRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorizationRef

`func (o *Transaction) SetAuthorizationRef(v string)`

SetAuthorizationRef sets AuthorizationRef field to given value.

### HasAuthorizationRef

`func (o *Transaction) HasAuthorizationRef() bool`

HasAuthorizationRef returns a boolean if a field has been set.

### GetCurrency

`func (o *Transaction) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *Transaction) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *Transaction) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *Transaction) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetDcSign

`func (o *Transaction) GetDcSign() DcSignType`

GetDcSign returns the DcSign field if non-nil, zero value otherwise.

### GetDcSignOk

`func (o *Transaction) GetDcSignOk() (*DcSignType, bool)`

GetDcSignOk returns a tuple with the DcSign field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDcSign

`func (o *Transaction) SetDcSign(v DcSignType)`

SetDcSign sets DcSign field to given value.

### HasDcSign

`func (o *Transaction) HasDcSign() bool`

HasDcSign returns a boolean if a field has been set.

### GetEffectiveDate

`func (o *Transaction) GetEffectiveDate() time.Time`

GetEffectiveDate returns the EffectiveDate field if non-nil, zero value otherwise.

### GetEffectiveDateOk

`func (o *Transaction) GetEffectiveDateOk() (*time.Time, bool)`

GetEffectiveDateOk returns a tuple with the EffectiveDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEffectiveDate

`func (o *Transaction) SetEffectiveDate(v time.Time)`

SetEffectiveDate sets EffectiveDate field to given value.

### HasEffectiveDate

`func (o *Transaction) HasEffectiveDate() bool`

HasEffectiveDate returns a boolean if a field has been set.

### GetExtReference

`func (o *Transaction) GetExtReference() string`

GetExtReference returns the ExtReference field if non-nil, zero value otherwise.

### GetExtReferenceOk

`func (o *Transaction) GetExtReferenceOk() (*string, bool)`

GetExtReferenceOk returns a tuple with the ExtReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtReference

`func (o *Transaction) SetExtReference(v string)`

SetExtReference sets ExtReference field to given value.

### HasExtReference

`func (o *Transaction) HasExtReference() bool`

HasExtReference returns a boolean if a field has been set.

### GetId

`func (o *Transaction) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Transaction) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Transaction) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Transaction) HasId() bool`

HasId returns a boolean if a field has been set.

### GetProfitCenter

`func (o *Transaction) GetProfitCenter() string`

GetProfitCenter returns the ProfitCenter field if non-nil, zero value otherwise.

### GetProfitCenterOk

`func (o *Transaction) GetProfitCenterOk() (*string, bool)`

GetProfitCenterOk returns a tuple with the ProfitCenter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfitCenter

`func (o *Transaction) SetProfitCenter(v string)`

SetProfitCenter sets ProfitCenter field to given value.

### HasProfitCenter

`func (o *Transaction) HasProfitCenter() bool`

HasProfitCenter returns a boolean if a field has been set.

### GetStatus

`func (o *Transaction) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Transaction) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Transaction) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *Transaction) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTransactionType

`func (o *Transaction) GetTransactionType() string`

GetTransactionType returns the TransactionType field if non-nil, zero value otherwise.

### GetTransactionTypeOk

`func (o *Transaction) GetTransactionTypeOk() (*string, bool)`

GetTransactionTypeOk returns a tuple with the TransactionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionType

`func (o *Transaction) SetTransactionType(v string)`

SetTransactionType sets TransactionType field to given value.

### HasTransactionType

`func (o *Transaction) HasTransactionType() bool`

HasTransactionType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


