# TransactionLine

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountNo** | **string** | The account number associated with this transaction line | 
**Amount** | **int32** | The amount (in cents) of the transaction | 
**AvailableBalance** | **int32** | The account \&quot;available balance\&quot; at the point in time this transaction was posted | 
**Balance** | **int32** | The account balance at the point in time this transaction was posted | 
**Created** | **time.Time** | The creation date of the transaction | 
**Currency** | **string** | ISO 4217 alphabetic currency code of the transfer amount | 
**DcSign** | [**DcSign**](DcSign.md) |  | 
**IsFee** | **bool** | Whether or not this line is considered a fee | 
**IsGlAcc** | **bool** | Whether or not this line represents a GL account | 
**IsOffset** | **bool** | Whether or not this line is considered the \&quot;offset\&quot; line | 
**IsPrimary** | **bool** | Whether or not this line is considered the \&quot;primary\&quot; line | 
**Meta** | **map[string]interface{}** |  | 
**Network** | **string** | The network this transaction is associated with | 
**RelatedLine** | **int32** |  | 
**Seq** | **int32** |  | 
**Tenant** | **string** | The tenant associated with this transaction, in the form \&quot;&lt;bankid&gt;_&lt;partnerid&gt;\&quot; | 
**Updated** | **time.Time** | The creation date of the transaction | 
**Uuid** | **string** |  | 

## Methods

### NewTransactionLine

`func NewTransactionLine(accountNo string, amount int32, availableBalance int32, balance int32, created time.Time, currency string, dcSign DcSign, isFee bool, isGlAcc bool, isOffset bool, isPrimary bool, meta map[string]interface{}, network string, relatedLine int32, seq int32, tenant string, updated time.Time, uuid string, ) *TransactionLine`

NewTransactionLine instantiates a new TransactionLine object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransactionLineWithDefaults

`func NewTransactionLineWithDefaults() *TransactionLine`

NewTransactionLineWithDefaults instantiates a new TransactionLine object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountNo

`func (o *TransactionLine) GetAccountNo() string`

GetAccountNo returns the AccountNo field if non-nil, zero value otherwise.

### GetAccountNoOk

`func (o *TransactionLine) GetAccountNoOk() (*string, bool)`

GetAccountNoOk returns a tuple with the AccountNo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountNo

`func (o *TransactionLine) SetAccountNo(v string)`

SetAccountNo sets AccountNo field to given value.


### GetAmount

`func (o *TransactionLine) GetAmount() int32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *TransactionLine) GetAmountOk() (*int32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *TransactionLine) SetAmount(v int32)`

SetAmount sets Amount field to given value.


### GetAvailableBalance

`func (o *TransactionLine) GetAvailableBalance() int32`

GetAvailableBalance returns the AvailableBalance field if non-nil, zero value otherwise.

### GetAvailableBalanceOk

`func (o *TransactionLine) GetAvailableBalanceOk() (*int32, bool)`

GetAvailableBalanceOk returns a tuple with the AvailableBalance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailableBalance

`func (o *TransactionLine) SetAvailableBalance(v int32)`

SetAvailableBalance sets AvailableBalance field to given value.


### GetBalance

`func (o *TransactionLine) GetBalance() int32`

GetBalance returns the Balance field if non-nil, zero value otherwise.

### GetBalanceOk

`func (o *TransactionLine) GetBalanceOk() (*int32, bool)`

GetBalanceOk returns a tuple with the Balance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalance

`func (o *TransactionLine) SetBalance(v int32)`

SetBalance sets Balance field to given value.


### GetCreated

`func (o *TransactionLine) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *TransactionLine) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *TransactionLine) SetCreated(v time.Time)`

SetCreated sets Created field to given value.


### GetCurrency

`func (o *TransactionLine) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *TransactionLine) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *TransactionLine) SetCurrency(v string)`

SetCurrency sets Currency field to given value.


### GetDcSign

`func (o *TransactionLine) GetDcSign() DcSign`

GetDcSign returns the DcSign field if non-nil, zero value otherwise.

### GetDcSignOk

`func (o *TransactionLine) GetDcSignOk() (*DcSign, bool)`

GetDcSignOk returns a tuple with the DcSign field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDcSign

`func (o *TransactionLine) SetDcSign(v DcSign)`

SetDcSign sets DcSign field to given value.


### GetIsFee

`func (o *TransactionLine) GetIsFee() bool`

GetIsFee returns the IsFee field if non-nil, zero value otherwise.

### GetIsFeeOk

`func (o *TransactionLine) GetIsFeeOk() (*bool, bool)`

GetIsFeeOk returns a tuple with the IsFee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsFee

`func (o *TransactionLine) SetIsFee(v bool)`

SetIsFee sets IsFee field to given value.


### GetIsGlAcc

`func (o *TransactionLine) GetIsGlAcc() bool`

GetIsGlAcc returns the IsGlAcc field if non-nil, zero value otherwise.

### GetIsGlAccOk

`func (o *TransactionLine) GetIsGlAccOk() (*bool, bool)`

GetIsGlAccOk returns a tuple with the IsGlAcc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsGlAcc

`func (o *TransactionLine) SetIsGlAcc(v bool)`

SetIsGlAcc sets IsGlAcc field to given value.


### GetIsOffset

`func (o *TransactionLine) GetIsOffset() bool`

GetIsOffset returns the IsOffset field if non-nil, zero value otherwise.

### GetIsOffsetOk

`func (o *TransactionLine) GetIsOffsetOk() (*bool, bool)`

GetIsOffsetOk returns a tuple with the IsOffset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsOffset

`func (o *TransactionLine) SetIsOffset(v bool)`

SetIsOffset sets IsOffset field to given value.


### GetIsPrimary

`func (o *TransactionLine) GetIsPrimary() bool`

GetIsPrimary returns the IsPrimary field if non-nil, zero value otherwise.

### GetIsPrimaryOk

`func (o *TransactionLine) GetIsPrimaryOk() (*bool, bool)`

GetIsPrimaryOk returns a tuple with the IsPrimary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPrimary

`func (o *TransactionLine) SetIsPrimary(v bool)`

SetIsPrimary sets IsPrimary field to given value.


### GetMeta

`func (o *TransactionLine) GetMeta() map[string]interface{}`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *TransactionLine) GetMetaOk() (*map[string]interface{}, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *TransactionLine) SetMeta(v map[string]interface{})`

SetMeta sets Meta field to given value.


### GetNetwork

`func (o *TransactionLine) GetNetwork() string`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *TransactionLine) GetNetworkOk() (*string, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *TransactionLine) SetNetwork(v string)`

SetNetwork sets Network field to given value.


### GetRelatedLine

`func (o *TransactionLine) GetRelatedLine() int32`

GetRelatedLine returns the RelatedLine field if non-nil, zero value otherwise.

### GetRelatedLineOk

`func (o *TransactionLine) GetRelatedLineOk() (*int32, bool)`

GetRelatedLineOk returns a tuple with the RelatedLine field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelatedLine

`func (o *TransactionLine) SetRelatedLine(v int32)`

SetRelatedLine sets RelatedLine field to given value.


### GetSeq

`func (o *TransactionLine) GetSeq() int32`

GetSeq returns the Seq field if non-nil, zero value otherwise.

### GetSeqOk

`func (o *TransactionLine) GetSeqOk() (*int32, bool)`

GetSeqOk returns a tuple with the Seq field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeq

`func (o *TransactionLine) SetSeq(v int32)`

SetSeq sets Seq field to given value.


### GetTenant

`func (o *TransactionLine) GetTenant() string`

GetTenant returns the Tenant field if non-nil, zero value otherwise.

### GetTenantOk

`func (o *TransactionLine) GetTenantOk() (*string, bool)`

GetTenantOk returns a tuple with the Tenant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenant

`func (o *TransactionLine) SetTenant(v string)`

SetTenant sets Tenant field to given value.


### GetUpdated

`func (o *TransactionLine) GetUpdated() time.Time`

GetUpdated returns the Updated field if non-nil, zero value otherwise.

### GetUpdatedOk

`func (o *TransactionLine) GetUpdatedOk() (*time.Time, bool)`

GetUpdatedOk returns a tuple with the Updated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdated

`func (o *TransactionLine) SetUpdated(v time.Time)`

SetUpdated sets Updated field to given value.


### GetUuid

`func (o *TransactionLine) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *TransactionLine) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *TransactionLine) SetUuid(v string)`

SetUuid sets Uuid field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


