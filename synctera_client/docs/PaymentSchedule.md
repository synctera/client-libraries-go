# PaymentSchedule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | **string** | User provided description for the payment schedule | 
**Id** | Pointer to **string** | Payment schedule ID | [optional] [readonly] 
**Metadata** | Pointer to **map[string]interface{}** | User provided JSON format data | [optional] 
**NextPaymentDate** | Pointer to [**PaymentDate**](PaymentDate.md) |  | [optional] 
**PaymentInstruction** | [**PaymentInstruction**](PaymentInstruction.md) |  | 
**Schedule** | [**ScheduleConfig**](ScheduleConfig.md) |  | 
**Status** | Pointer to [**PaymentScheduleStatus**](PaymentScheduleStatus.md) |  | [optional] 

## Methods

### NewPaymentSchedule

`func NewPaymentSchedule(description string, paymentInstruction PaymentInstruction, schedule ScheduleConfig, ) *PaymentSchedule`

NewPaymentSchedule instantiates a new PaymentSchedule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaymentScheduleWithDefaults

`func NewPaymentScheduleWithDefaults() *PaymentSchedule`

NewPaymentScheduleWithDefaults instantiates a new PaymentSchedule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *PaymentSchedule) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PaymentSchedule) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PaymentSchedule) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetId

`func (o *PaymentSchedule) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PaymentSchedule) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PaymentSchedule) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *PaymentSchedule) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMetadata

`func (o *PaymentSchedule) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *PaymentSchedule) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *PaymentSchedule) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *PaymentSchedule) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetNextPaymentDate

`func (o *PaymentSchedule) GetNextPaymentDate() PaymentDate`

GetNextPaymentDate returns the NextPaymentDate field if non-nil, zero value otherwise.

### GetNextPaymentDateOk

`func (o *PaymentSchedule) GetNextPaymentDateOk() (*PaymentDate, bool)`

GetNextPaymentDateOk returns a tuple with the NextPaymentDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextPaymentDate

`func (o *PaymentSchedule) SetNextPaymentDate(v PaymentDate)`

SetNextPaymentDate sets NextPaymentDate field to given value.

### HasNextPaymentDate

`func (o *PaymentSchedule) HasNextPaymentDate() bool`

HasNextPaymentDate returns a boolean if a field has been set.

### GetPaymentInstruction

`func (o *PaymentSchedule) GetPaymentInstruction() PaymentInstruction`

GetPaymentInstruction returns the PaymentInstruction field if non-nil, zero value otherwise.

### GetPaymentInstructionOk

`func (o *PaymentSchedule) GetPaymentInstructionOk() (*PaymentInstruction, bool)`

GetPaymentInstructionOk returns a tuple with the PaymentInstruction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentInstruction

`func (o *PaymentSchedule) SetPaymentInstruction(v PaymentInstruction)`

SetPaymentInstruction sets PaymentInstruction field to given value.


### GetSchedule

`func (o *PaymentSchedule) GetSchedule() ScheduleConfig`

GetSchedule returns the Schedule field if non-nil, zero value otherwise.

### GetScheduleOk

`func (o *PaymentSchedule) GetScheduleOk() (*ScheduleConfig, bool)`

GetScheduleOk returns a tuple with the Schedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchedule

`func (o *PaymentSchedule) SetSchedule(v ScheduleConfig)`

SetSchedule sets Schedule field to given value.


### GetStatus

`func (o *PaymentSchedule) GetStatus() PaymentScheduleStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *PaymentSchedule) GetStatusOk() (*PaymentScheduleStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *PaymentSchedule) SetStatus(v PaymentScheduleStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *PaymentSchedule) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


