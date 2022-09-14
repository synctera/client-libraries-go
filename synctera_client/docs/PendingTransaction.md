# PendingTransaction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountId** | **string** | The account id associated with the hold | 
**AccountNo** | **string** | The account number associated with the hold | 
**Created** | **time.Time** | The creation date of the hold | 
**Data** | [**PendingTransactionData**](PendingTransactionData.md) |  | 
**Id** | **int64** |  | 
**Idemkey** | **string** | The idempotency key used when initially creating this hold. | 
**OffsetAccountId** | **string** | The offset account id associated with the hold | 
**OffsetAccountNo** | **string** | The offset account number associated with the hold | 
**ReferenceId** | **NullableString** | An external ID provided by the payment network to represent this transaction. | 
**Tenant** | **string** | The tenant associated with this hold, in the form \&quot;&lt;bankid&gt;_&lt;partnerid&gt;\&quot; | 
**Updated** | **time.Time** | The date the hold was last update | 
**Uuid** | **string** | The unique identifier of the hold transaction. | 

## Methods

### NewPendingTransaction

`func NewPendingTransaction(accountId string, accountNo string, created time.Time, data PendingTransactionData, id int64, idemkey string, offsetAccountId string, offsetAccountNo string, referenceId NullableString, tenant string, updated time.Time, uuid string, ) *PendingTransaction`

NewPendingTransaction instantiates a new PendingTransaction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPendingTransactionWithDefaults

`func NewPendingTransactionWithDefaults() *PendingTransaction`

NewPendingTransactionWithDefaults instantiates a new PendingTransaction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountId

`func (o *PendingTransaction) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *PendingTransaction) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *PendingTransaction) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.


### GetAccountNo

`func (o *PendingTransaction) GetAccountNo() string`

GetAccountNo returns the AccountNo field if non-nil, zero value otherwise.

### GetAccountNoOk

`func (o *PendingTransaction) GetAccountNoOk() (*string, bool)`

GetAccountNoOk returns a tuple with the AccountNo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountNo

`func (o *PendingTransaction) SetAccountNo(v string)`

SetAccountNo sets AccountNo field to given value.


### GetCreated

`func (o *PendingTransaction) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *PendingTransaction) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *PendingTransaction) SetCreated(v time.Time)`

SetCreated sets Created field to given value.


### GetData

`func (o *PendingTransaction) GetData() PendingTransactionData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *PendingTransaction) GetDataOk() (*PendingTransactionData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *PendingTransaction) SetData(v PendingTransactionData)`

SetData sets Data field to given value.


### GetId

`func (o *PendingTransaction) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PendingTransaction) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PendingTransaction) SetId(v int64)`

SetId sets Id field to given value.


### GetIdemkey

`func (o *PendingTransaction) GetIdemkey() string`

GetIdemkey returns the Idemkey field if non-nil, zero value otherwise.

### GetIdemkeyOk

`func (o *PendingTransaction) GetIdemkeyOk() (*string, bool)`

GetIdemkeyOk returns a tuple with the Idemkey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdemkey

`func (o *PendingTransaction) SetIdemkey(v string)`

SetIdemkey sets Idemkey field to given value.


### GetOffsetAccountId

`func (o *PendingTransaction) GetOffsetAccountId() string`

GetOffsetAccountId returns the OffsetAccountId field if non-nil, zero value otherwise.

### GetOffsetAccountIdOk

`func (o *PendingTransaction) GetOffsetAccountIdOk() (*string, bool)`

GetOffsetAccountIdOk returns a tuple with the OffsetAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffsetAccountId

`func (o *PendingTransaction) SetOffsetAccountId(v string)`

SetOffsetAccountId sets OffsetAccountId field to given value.


### GetOffsetAccountNo

`func (o *PendingTransaction) GetOffsetAccountNo() string`

GetOffsetAccountNo returns the OffsetAccountNo field if non-nil, zero value otherwise.

### GetOffsetAccountNoOk

`func (o *PendingTransaction) GetOffsetAccountNoOk() (*string, bool)`

GetOffsetAccountNoOk returns a tuple with the OffsetAccountNo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffsetAccountNo

`func (o *PendingTransaction) SetOffsetAccountNo(v string)`

SetOffsetAccountNo sets OffsetAccountNo field to given value.


### GetReferenceId

`func (o *PendingTransaction) GetReferenceId() string`

GetReferenceId returns the ReferenceId field if non-nil, zero value otherwise.

### GetReferenceIdOk

`func (o *PendingTransaction) GetReferenceIdOk() (*string, bool)`

GetReferenceIdOk returns a tuple with the ReferenceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceId

`func (o *PendingTransaction) SetReferenceId(v string)`

SetReferenceId sets ReferenceId field to given value.


### SetReferenceIdNil

`func (o *PendingTransaction) SetReferenceIdNil(b bool)`

 SetReferenceIdNil sets the value for ReferenceId to be an explicit nil

### UnsetReferenceId
`func (o *PendingTransaction) UnsetReferenceId()`

UnsetReferenceId ensures that no value is present for ReferenceId, not even an explicit nil
### GetTenant

`func (o *PendingTransaction) GetTenant() string`

GetTenant returns the Tenant field if non-nil, zero value otherwise.

### GetTenantOk

`func (o *PendingTransaction) GetTenantOk() (*string, bool)`

GetTenantOk returns a tuple with the Tenant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenant

`func (o *PendingTransaction) SetTenant(v string)`

SetTenant sets Tenant field to given value.


### GetUpdated

`func (o *PendingTransaction) GetUpdated() time.Time`

GetUpdated returns the Updated field if non-nil, zero value otherwise.

### GetUpdatedOk

`func (o *PendingTransaction) GetUpdatedOk() (*time.Time, bool)`

GetUpdatedOk returns a tuple with the Updated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdated

`func (o *PendingTransaction) SetUpdated(v time.Time)`

SetUpdated sets Updated field to given value.


### GetUuid

`func (o *PendingTransaction) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *PendingTransaction) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *PendingTransaction) SetUuid(v string)`

SetUuid sets Uuid field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


