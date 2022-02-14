# EventResponseHistory

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | Pointer to **int32** | Response code from the request | [optional] 
**ResponseBody** | Pointer to **string** | Response body from the request | [optional] 
**ResponseTime** | Pointer to **time.Time** | Timestamp that the response is received | [optional] 
**SentTime** | Pointer to **time.Time** | Timestamp that the request is sent | [optional] 

## Methods

### NewEventResponseHistory

`func NewEventResponseHistory() *EventResponseHistory`

NewEventResponseHistory instantiates a new EventResponseHistory object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEventResponseHistoryWithDefaults

`func NewEventResponseHistoryWithDefaults() *EventResponseHistory`

NewEventResponseHistoryWithDefaults instantiates a new EventResponseHistory object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *EventResponseHistory) GetCode() int32`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *EventResponseHistory) GetCodeOk() (*int32, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *EventResponseHistory) SetCode(v int32)`

SetCode sets Code field to given value.

### HasCode

`func (o *EventResponseHistory) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetResponseBody

`func (o *EventResponseHistory) GetResponseBody() string`

GetResponseBody returns the ResponseBody field if non-nil, zero value otherwise.

### GetResponseBodyOk

`func (o *EventResponseHistory) GetResponseBodyOk() (*string, bool)`

GetResponseBodyOk returns a tuple with the ResponseBody field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseBody

`func (o *EventResponseHistory) SetResponseBody(v string)`

SetResponseBody sets ResponseBody field to given value.

### HasResponseBody

`func (o *EventResponseHistory) HasResponseBody() bool`

HasResponseBody returns a boolean if a field has been set.

### GetResponseTime

`func (o *EventResponseHistory) GetResponseTime() time.Time`

GetResponseTime returns the ResponseTime field if non-nil, zero value otherwise.

### GetResponseTimeOk

`func (o *EventResponseHistory) GetResponseTimeOk() (*time.Time, bool)`

GetResponseTimeOk returns a tuple with the ResponseTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseTime

`func (o *EventResponseHistory) SetResponseTime(v time.Time)`

SetResponseTime sets ResponseTime field to given value.

### HasResponseTime

`func (o *EventResponseHistory) HasResponseTime() bool`

HasResponseTime returns a boolean if a field has been set.

### GetSentTime

`func (o *EventResponseHistory) GetSentTime() time.Time`

GetSentTime returns the SentTime field if non-nil, zero value otherwise.

### GetSentTimeOk

`func (o *EventResponseHistory) GetSentTimeOk() (*time.Time, bool)`

GetSentTimeOk returns a tuple with the SentTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSentTime

`func (o *EventResponseHistory) SetSentTime(v time.Time)`

SetSentTime sets SentTime field to given value.

### HasSentTime

`func (o *EventResponseHistory) HasSentTime() bool`

HasSentTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


