# EventResend

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EventType** | [**EventType**](EventType.md) |  | 
**EventId** | **string** | id of the event notification you want to resend | 

## Methods

### NewEventResend

`func NewEventResend(eventType EventType, eventId string, ) *EventResend`

NewEventResend instantiates a new EventResend object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEventResendWithDefaults

`func NewEventResendWithDefaults() *EventResend`

NewEventResendWithDefaults instantiates a new EventResend object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEventType

`func (o *EventResend) GetEventType() EventType`

GetEventType returns the EventType field if non-nil, zero value otherwise.

### GetEventTypeOk

`func (o *EventResend) GetEventTypeOk() (*EventType, bool)`

GetEventTypeOk returns a tuple with the EventType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventType

`func (o *EventResend) SetEventType(v EventType)`

SetEventType sets EventType field to given value.


### GetEventId

`func (o *EventResend) GetEventId() string`

GetEventId returns the EventId field if non-nil, zero value otherwise.

### GetEventIdOk

`func (o *EventResend) GetEventIdOk() (*string, bool)`

GetEventIdOk returns a tuple with the EventId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventId

`func (o *EventResend) SetEventId(v string)`

SetEventId sets EventId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


