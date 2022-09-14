# PostedTransaction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Created** | **time.Time** | The creation date of the transaction | 
**Data** | [**PostedTransactionData**](PostedTransactionData.md) |  | 
**EffectiveDate** | **time.Time** | The \&quot;effective date\&quot; of a transaction. This may be earlier than posted_date in some cases (for example, a transaction that occurs on a Saturday may not be posted until the following Monday, but would have an effective date of Saturday) | 
**Id** | **int64** |  | 
**Idemkey** | **string** | The idempotency key used when initially creating this transaction. | 
**InfoOnly** | **bool** | Whether or not this transaction represents a purely informational operation or an actual money movement | 
**LeadMode** | **bool** | Whether or not this transaction was created operating in \&quot;lead ledger\&quot; mode | 
**PostedDate** | **time.Time** | The date the transaction was posted. This is the date any money is considered to be added or removed from an account. | 
**ReferenceId** | **NullableString** | An external ID provided by the payment network to represent this transaction. This will always be null for internal transfers. | 
**Status** | **string** |  | 
**Subtype** | **string** | The specific transaction type. For example, for &#x60;ach&#x60;, this may be \&quot;outgoing_debit\&quot;. | 
**Tenant** | **string** | The tenant associated with this transaction, in the form \&quot;&lt;bankid&gt;_&lt;partnerid&gt;\&quot; | 
**TransactionTime** | **time.Time** | The time the transaction occurred. | 
**Type** | **string** | The general type of transaction. For example, \&quot;card\&quot; or \&quot;ach\&quot;. | 
**Updated** | **time.Time** | The date the transaction was last updated | 
**Uuid** | **string** | The unique identifier of the transaction. | 

## Methods

### NewPostedTransaction

`func NewPostedTransaction(created time.Time, data PostedTransactionData, effectiveDate time.Time, id int64, idemkey string, infoOnly bool, leadMode bool, postedDate time.Time, referenceId NullableString, status string, subtype string, tenant string, transactionTime time.Time, type_ string, updated time.Time, uuid string, ) *PostedTransaction`

NewPostedTransaction instantiates a new PostedTransaction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostedTransactionWithDefaults

`func NewPostedTransactionWithDefaults() *PostedTransaction`

NewPostedTransactionWithDefaults instantiates a new PostedTransaction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreated

`func (o *PostedTransaction) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *PostedTransaction) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *PostedTransaction) SetCreated(v time.Time)`

SetCreated sets Created field to given value.


### GetData

`func (o *PostedTransaction) GetData() PostedTransactionData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *PostedTransaction) GetDataOk() (*PostedTransactionData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *PostedTransaction) SetData(v PostedTransactionData)`

SetData sets Data field to given value.


### GetEffectiveDate

`func (o *PostedTransaction) GetEffectiveDate() time.Time`

GetEffectiveDate returns the EffectiveDate field if non-nil, zero value otherwise.

### GetEffectiveDateOk

`func (o *PostedTransaction) GetEffectiveDateOk() (*time.Time, bool)`

GetEffectiveDateOk returns a tuple with the EffectiveDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEffectiveDate

`func (o *PostedTransaction) SetEffectiveDate(v time.Time)`

SetEffectiveDate sets EffectiveDate field to given value.


### GetId

`func (o *PostedTransaction) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PostedTransaction) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PostedTransaction) SetId(v int64)`

SetId sets Id field to given value.


### GetIdemkey

`func (o *PostedTransaction) GetIdemkey() string`

GetIdemkey returns the Idemkey field if non-nil, zero value otherwise.

### GetIdemkeyOk

`func (o *PostedTransaction) GetIdemkeyOk() (*string, bool)`

GetIdemkeyOk returns a tuple with the Idemkey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdemkey

`func (o *PostedTransaction) SetIdemkey(v string)`

SetIdemkey sets Idemkey field to given value.


### GetInfoOnly

`func (o *PostedTransaction) GetInfoOnly() bool`

GetInfoOnly returns the InfoOnly field if non-nil, zero value otherwise.

### GetInfoOnlyOk

`func (o *PostedTransaction) GetInfoOnlyOk() (*bool, bool)`

GetInfoOnlyOk returns a tuple with the InfoOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfoOnly

`func (o *PostedTransaction) SetInfoOnly(v bool)`

SetInfoOnly sets InfoOnly field to given value.


### GetLeadMode

`func (o *PostedTransaction) GetLeadMode() bool`

GetLeadMode returns the LeadMode field if non-nil, zero value otherwise.

### GetLeadModeOk

`func (o *PostedTransaction) GetLeadModeOk() (*bool, bool)`

GetLeadModeOk returns a tuple with the LeadMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeadMode

`func (o *PostedTransaction) SetLeadMode(v bool)`

SetLeadMode sets LeadMode field to given value.


### GetPostedDate

`func (o *PostedTransaction) GetPostedDate() time.Time`

GetPostedDate returns the PostedDate field if non-nil, zero value otherwise.

### GetPostedDateOk

`func (o *PostedTransaction) GetPostedDateOk() (*time.Time, bool)`

GetPostedDateOk returns a tuple with the PostedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostedDate

`func (o *PostedTransaction) SetPostedDate(v time.Time)`

SetPostedDate sets PostedDate field to given value.


### GetReferenceId

`func (o *PostedTransaction) GetReferenceId() string`

GetReferenceId returns the ReferenceId field if non-nil, zero value otherwise.

### GetReferenceIdOk

`func (o *PostedTransaction) GetReferenceIdOk() (*string, bool)`

GetReferenceIdOk returns a tuple with the ReferenceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceId

`func (o *PostedTransaction) SetReferenceId(v string)`

SetReferenceId sets ReferenceId field to given value.


### SetReferenceIdNil

`func (o *PostedTransaction) SetReferenceIdNil(b bool)`

 SetReferenceIdNil sets the value for ReferenceId to be an explicit nil

### UnsetReferenceId
`func (o *PostedTransaction) UnsetReferenceId()`

UnsetReferenceId ensures that no value is present for ReferenceId, not even an explicit nil
### GetStatus

`func (o *PostedTransaction) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *PostedTransaction) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *PostedTransaction) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetSubtype

`func (o *PostedTransaction) GetSubtype() string`

GetSubtype returns the Subtype field if non-nil, zero value otherwise.

### GetSubtypeOk

`func (o *PostedTransaction) GetSubtypeOk() (*string, bool)`

GetSubtypeOk returns a tuple with the Subtype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubtype

`func (o *PostedTransaction) SetSubtype(v string)`

SetSubtype sets Subtype field to given value.


### GetTenant

`func (o *PostedTransaction) GetTenant() string`

GetTenant returns the Tenant field if non-nil, zero value otherwise.

### GetTenantOk

`func (o *PostedTransaction) GetTenantOk() (*string, bool)`

GetTenantOk returns a tuple with the Tenant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenant

`func (o *PostedTransaction) SetTenant(v string)`

SetTenant sets Tenant field to given value.


### GetTransactionTime

`func (o *PostedTransaction) GetTransactionTime() time.Time`

GetTransactionTime returns the TransactionTime field if non-nil, zero value otherwise.

### GetTransactionTimeOk

`func (o *PostedTransaction) GetTransactionTimeOk() (*time.Time, bool)`

GetTransactionTimeOk returns a tuple with the TransactionTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionTime

`func (o *PostedTransaction) SetTransactionTime(v time.Time)`

SetTransactionTime sets TransactionTime field to given value.


### GetType

`func (o *PostedTransaction) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PostedTransaction) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PostedTransaction) SetType(v string)`

SetType sets Type field to given value.


### GetUpdated

`func (o *PostedTransaction) GetUpdated() time.Time`

GetUpdated returns the Updated field if non-nil, zero value otherwise.

### GetUpdatedOk

`func (o *PostedTransaction) GetUpdatedOk() (*time.Time, bool)`

GetUpdatedOk returns a tuple with the Updated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdated

`func (o *PostedTransaction) SetUpdated(v time.Time)`

SetUpdated sets Updated field to given value.


### GetUuid

`func (o *PostedTransaction) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *PostedTransaction) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *PostedTransaction) SetUuid(v string)`

SetUuid sets Uuid field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


