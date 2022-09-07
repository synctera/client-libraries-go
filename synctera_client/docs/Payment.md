# Payment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** | User provided description for the payment schedule | [optional] 
**ErrorDetails** | Pointer to [**PaymentErrorDetails**](PaymentErrorDetails.md) |  | [optional] 
**Id** | Pointer to **string** | Payment ID | [optional] 
**Metadata** | Pointer to **map[string]interface{}** | User provided JSON format data for the payment schedule | [optional] 
**PaymentDate** | Pointer to [**PaymentDate**](PaymentDate.md) |  | [optional] 
**PaymentInstruction** | Pointer to [**PaymentInstruction**](PaymentInstruction.md) |  | [optional] 
**PaymentScheduleId** | Pointer to **string** | ID of the payment schedule that executed this payment | [optional] 
**Status** | Pointer to [**PaymentStatus**](PaymentStatus.md) |  | [optional] 
**TransactionId** | Pointer to **string** | Transaction ID. It will be included only when status is COMPLETED | [optional] 

## Methods

### NewPayment

`func NewPayment() *Payment`

NewPayment instantiates a new Payment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaymentWithDefaults

`func NewPaymentWithDefaults() *Payment`

NewPaymentWithDefaults instantiates a new Payment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *Payment) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Payment) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Payment) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Payment) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetErrorDetails

`func (o *Payment) GetErrorDetails() PaymentErrorDetails`

GetErrorDetails returns the ErrorDetails field if non-nil, zero value otherwise.

### GetErrorDetailsOk

`func (o *Payment) GetErrorDetailsOk() (*PaymentErrorDetails, bool)`

GetErrorDetailsOk returns a tuple with the ErrorDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorDetails

`func (o *Payment) SetErrorDetails(v PaymentErrorDetails)`

SetErrorDetails sets ErrorDetails field to given value.

### HasErrorDetails

`func (o *Payment) HasErrorDetails() bool`

HasErrorDetails returns a boolean if a field has been set.

### GetId

`func (o *Payment) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Payment) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Payment) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Payment) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMetadata

`func (o *Payment) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *Payment) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *Payment) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *Payment) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetPaymentDate

`func (o *Payment) GetPaymentDate() PaymentDate`

GetPaymentDate returns the PaymentDate field if non-nil, zero value otherwise.

### GetPaymentDateOk

`func (o *Payment) GetPaymentDateOk() (*PaymentDate, bool)`

GetPaymentDateOk returns a tuple with the PaymentDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentDate

`func (o *Payment) SetPaymentDate(v PaymentDate)`

SetPaymentDate sets PaymentDate field to given value.

### HasPaymentDate

`func (o *Payment) HasPaymentDate() bool`

HasPaymentDate returns a boolean if a field has been set.

### GetPaymentInstruction

`func (o *Payment) GetPaymentInstruction() PaymentInstruction`

GetPaymentInstruction returns the PaymentInstruction field if non-nil, zero value otherwise.

### GetPaymentInstructionOk

`func (o *Payment) GetPaymentInstructionOk() (*PaymentInstruction, bool)`

GetPaymentInstructionOk returns a tuple with the PaymentInstruction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentInstruction

`func (o *Payment) SetPaymentInstruction(v PaymentInstruction)`

SetPaymentInstruction sets PaymentInstruction field to given value.

### HasPaymentInstruction

`func (o *Payment) HasPaymentInstruction() bool`

HasPaymentInstruction returns a boolean if a field has been set.

### GetPaymentScheduleId

`func (o *Payment) GetPaymentScheduleId() string`

GetPaymentScheduleId returns the PaymentScheduleId field if non-nil, zero value otherwise.

### GetPaymentScheduleIdOk

`func (o *Payment) GetPaymentScheduleIdOk() (*string, bool)`

GetPaymentScheduleIdOk returns a tuple with the PaymentScheduleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentScheduleId

`func (o *Payment) SetPaymentScheduleId(v string)`

SetPaymentScheduleId sets PaymentScheduleId field to given value.

### HasPaymentScheduleId

`func (o *Payment) HasPaymentScheduleId() bool`

HasPaymentScheduleId returns a boolean if a field has been set.

### GetStatus

`func (o *Payment) GetStatus() PaymentStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Payment) GetStatusOk() (*PaymentStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Payment) SetStatus(v PaymentStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *Payment) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTransactionId

`func (o *Payment) GetTransactionId() string`

GetTransactionId returns the TransactionId field if non-nil, zero value otherwise.

### GetTransactionIdOk

`func (o *Payment) GetTransactionIdOk() (*string, bool)`

GetTransactionIdOk returns a tuple with the TransactionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionId

`func (o *Payment) SetTransactionId(v string)`

SetTransactionId sets TransactionId field to given value.

### HasTransactionId

`func (o *Payment) HasTransactionId() bool`

HasTransactionId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


