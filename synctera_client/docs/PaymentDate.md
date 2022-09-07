# PaymentDate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExecutionDate** | **string** | Execution date for the next payment | 
**ScheduledDate** | **string** | Scheduled date for the next payment | 

## Methods

### NewPaymentDate

`func NewPaymentDate(executionDate string, scheduledDate string, ) *PaymentDate`

NewPaymentDate instantiates a new PaymentDate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaymentDateWithDefaults

`func NewPaymentDateWithDefaults() *PaymentDate`

NewPaymentDateWithDefaults instantiates a new PaymentDate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExecutionDate

`func (o *PaymentDate) GetExecutionDate() string`

GetExecutionDate returns the ExecutionDate field if non-nil, zero value otherwise.

### GetExecutionDateOk

`func (o *PaymentDate) GetExecutionDateOk() (*string, bool)`

GetExecutionDateOk returns a tuple with the ExecutionDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutionDate

`func (o *PaymentDate) SetExecutionDate(v string)`

SetExecutionDate sets ExecutionDate field to given value.


### GetScheduledDate

`func (o *PaymentDate) GetScheduledDate() string`

GetScheduledDate returns the ScheduledDate field if non-nil, zero value otherwise.

### GetScheduledDateOk

`func (o *PaymentDate) GetScheduledDateOk() (*string, bool)`

GetScheduledDateOk returns a tuple with the ScheduledDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduledDate

`func (o *PaymentDate) SetScheduledDate(v string)`

SetScheduledDate sets ScheduledDate field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


