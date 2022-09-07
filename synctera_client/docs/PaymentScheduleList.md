# PaymentScheduleList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PaymentSchedules** | [**[]PaymentSchedule**](PaymentSchedule.md) | Array of payment schedules. | 
**NextPageToken** | Pointer to **string** | If returned, use the next_page_token to query for the next page of results. Not returned if there are no more rows. | [optional] 

## Methods

### NewPaymentScheduleList

`func NewPaymentScheduleList(paymentSchedules []PaymentSchedule, ) *PaymentScheduleList`

NewPaymentScheduleList instantiates a new PaymentScheduleList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaymentScheduleListWithDefaults

`func NewPaymentScheduleListWithDefaults() *PaymentScheduleList`

NewPaymentScheduleListWithDefaults instantiates a new PaymentScheduleList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPaymentSchedules

`func (o *PaymentScheduleList) GetPaymentSchedules() []PaymentSchedule`

GetPaymentSchedules returns the PaymentSchedules field if non-nil, zero value otherwise.

### GetPaymentSchedulesOk

`func (o *PaymentScheduleList) GetPaymentSchedulesOk() (*[]PaymentSchedule, bool)`

GetPaymentSchedulesOk returns a tuple with the PaymentSchedules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentSchedules

`func (o *PaymentScheduleList) SetPaymentSchedules(v []PaymentSchedule)`

SetPaymentSchedules sets PaymentSchedules field to given value.


### GetNextPageToken

`func (o *PaymentScheduleList) GetNextPageToken() string`

GetNextPageToken returns the NextPageToken field if non-nil, zero value otherwise.

### GetNextPageTokenOk

`func (o *PaymentScheduleList) GetNextPageTokenOk() (*string, bool)`

GetNextPageTokenOk returns a tuple with the NextPageToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextPageToken

`func (o *PaymentScheduleList) SetNextPageToken(v string)`

SetNextPageToken sets NextPageToken field to given value.

### HasNextPageToken

`func (o *PaymentScheduleList) HasNextPageToken() bool`

HasNextPageToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


