# A2aTransfer

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | **int32** | amount in ISO 4217 minor currency units | 
**Currency** | **string** | Account currency or account settlement currency. ISO 4217 alphabetic currency code. | 
**DcSign** | [**DcSignType**](DcSignType.md) |  | 
**ExecutionDate** | Pointer to **string** | Execution date of the transfer. Default is the current date | [optional] 
**IsOverwriteChecks** | Pointer to **bool** | Flag to indicate the overwrite checks | [optional] [readonly] 
**PaymentId** | Pointer to **string** | Payment ID | [optional] [readonly] 
**RecurringData** | Pointer to [**RecurrenceData**](RecurrenceData.md) |  | [optional] 
**ReferenceInfo** | Pointer to **string** | User specified information of this transfer | [optional] 
**SourceAccount** | **string** | Account ID that is instructing the transfer | 
**TargetAccount** | **string** | Account ID that is receiving the transfer | 
**TransferReversal** | Pointer to [**A2aTransferTransferReversal**](A2aTransferTransferReversal.md) |  | [optional] 

## Methods

### NewA2aTransfer

`func NewA2aTransfer(amount int32, currency string, dcSign DcSignType, sourceAccount string, targetAccount string, ) *A2aTransfer`

NewA2aTransfer instantiates a new A2aTransfer object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewA2aTransferWithDefaults

`func NewA2aTransferWithDefaults() *A2aTransfer`

NewA2aTransferWithDefaults instantiates a new A2aTransfer object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmount

`func (o *A2aTransfer) GetAmount() int32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *A2aTransfer) GetAmountOk() (*int32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *A2aTransfer) SetAmount(v int32)`

SetAmount sets Amount field to given value.


### GetCurrency

`func (o *A2aTransfer) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *A2aTransfer) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *A2aTransfer) SetCurrency(v string)`

SetCurrency sets Currency field to given value.


### GetDcSign

`func (o *A2aTransfer) GetDcSign() DcSignType`

GetDcSign returns the DcSign field if non-nil, zero value otherwise.

### GetDcSignOk

`func (o *A2aTransfer) GetDcSignOk() (*DcSignType, bool)`

GetDcSignOk returns a tuple with the DcSign field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDcSign

`func (o *A2aTransfer) SetDcSign(v DcSignType)`

SetDcSign sets DcSign field to given value.


### GetExecutionDate

`func (o *A2aTransfer) GetExecutionDate() string`

GetExecutionDate returns the ExecutionDate field if non-nil, zero value otherwise.

### GetExecutionDateOk

`func (o *A2aTransfer) GetExecutionDateOk() (*string, bool)`

GetExecutionDateOk returns a tuple with the ExecutionDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutionDate

`func (o *A2aTransfer) SetExecutionDate(v string)`

SetExecutionDate sets ExecutionDate field to given value.

### HasExecutionDate

`func (o *A2aTransfer) HasExecutionDate() bool`

HasExecutionDate returns a boolean if a field has been set.

### GetIsOverwriteChecks

`func (o *A2aTransfer) GetIsOverwriteChecks() bool`

GetIsOverwriteChecks returns the IsOverwriteChecks field if non-nil, zero value otherwise.

### GetIsOverwriteChecksOk

`func (o *A2aTransfer) GetIsOverwriteChecksOk() (*bool, bool)`

GetIsOverwriteChecksOk returns a tuple with the IsOverwriteChecks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsOverwriteChecks

`func (o *A2aTransfer) SetIsOverwriteChecks(v bool)`

SetIsOverwriteChecks sets IsOverwriteChecks field to given value.

### HasIsOverwriteChecks

`func (o *A2aTransfer) HasIsOverwriteChecks() bool`

HasIsOverwriteChecks returns a boolean if a field has been set.

### GetPaymentId

`func (o *A2aTransfer) GetPaymentId() string`

GetPaymentId returns the PaymentId field if non-nil, zero value otherwise.

### GetPaymentIdOk

`func (o *A2aTransfer) GetPaymentIdOk() (*string, bool)`

GetPaymentIdOk returns a tuple with the PaymentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentId

`func (o *A2aTransfer) SetPaymentId(v string)`

SetPaymentId sets PaymentId field to given value.

### HasPaymentId

`func (o *A2aTransfer) HasPaymentId() bool`

HasPaymentId returns a boolean if a field has been set.

### GetRecurringData

`func (o *A2aTransfer) GetRecurringData() RecurrenceData`

GetRecurringData returns the RecurringData field if non-nil, zero value otherwise.

### GetRecurringDataOk

`func (o *A2aTransfer) GetRecurringDataOk() (*RecurrenceData, bool)`

GetRecurringDataOk returns a tuple with the RecurringData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecurringData

`func (o *A2aTransfer) SetRecurringData(v RecurrenceData)`

SetRecurringData sets RecurringData field to given value.

### HasRecurringData

`func (o *A2aTransfer) HasRecurringData() bool`

HasRecurringData returns a boolean if a field has been set.

### GetReferenceInfo

`func (o *A2aTransfer) GetReferenceInfo() string`

GetReferenceInfo returns the ReferenceInfo field if non-nil, zero value otherwise.

### GetReferenceInfoOk

`func (o *A2aTransfer) GetReferenceInfoOk() (*string, bool)`

GetReferenceInfoOk returns a tuple with the ReferenceInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceInfo

`func (o *A2aTransfer) SetReferenceInfo(v string)`

SetReferenceInfo sets ReferenceInfo field to given value.

### HasReferenceInfo

`func (o *A2aTransfer) HasReferenceInfo() bool`

HasReferenceInfo returns a boolean if a field has been set.

### GetSourceAccount

`func (o *A2aTransfer) GetSourceAccount() string`

GetSourceAccount returns the SourceAccount field if non-nil, zero value otherwise.

### GetSourceAccountOk

`func (o *A2aTransfer) GetSourceAccountOk() (*string, bool)`

GetSourceAccountOk returns a tuple with the SourceAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceAccount

`func (o *A2aTransfer) SetSourceAccount(v string)`

SetSourceAccount sets SourceAccount field to given value.


### GetTargetAccount

`func (o *A2aTransfer) GetTargetAccount() string`

GetTargetAccount returns the TargetAccount field if non-nil, zero value otherwise.

### GetTargetAccountOk

`func (o *A2aTransfer) GetTargetAccountOk() (*string, bool)`

GetTargetAccountOk returns a tuple with the TargetAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetAccount

`func (o *A2aTransfer) SetTargetAccount(v string)`

SetTargetAccount sets TargetAccount field to given value.


### GetTransferReversal

`func (o *A2aTransfer) GetTransferReversal() A2aTransferTransferReversal`

GetTransferReversal returns the TransferReversal field if non-nil, zero value otherwise.

### GetTransferReversalOk

`func (o *A2aTransfer) GetTransferReversalOk() (*A2aTransferTransferReversal, bool)`

GetTransferReversalOk returns a tuple with the TransferReversal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransferReversal

`func (o *A2aTransfer) SetTransferReversal(v A2aTransferTransferReversal)`

SetTransferReversal sets TransferReversal field to given value.

### HasTransferReversal

`func (o *A2aTransfer) HasTransferReversal() bool`

HasTransferReversal returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


