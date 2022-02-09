# Event

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Unique event ID of the webhook request. Use event endpoints to get more event summary data | [optional] [readonly] 
**WebhookId** | Pointer to **string** | Webhook the current event belongs to | [optional] 
**Type** | Pointer to [**EventTypeExplicit**](EventTypeExplicit.md) |  | [optional] 
**EventTime** | Pointer to **time.Time** | Timestamp of the current event raised | [optional] 
**Metadata** | Pointer to **string** | Metadata that stored in the webhook subscription | [optional] 
**EventResource** | Pointer to **string** | Json string of object associated with the event. For example, if your event is ACCOUNT.CREATED, You can refer to Acccount to parse the account event to obtain the ID, status etc.  | [optional] 
**EventResourceChangedFields** | Pointer to **string** | Json string of object associated with the event related to a resource change. This only contains those fields that have value changed on the event, and the field values are prior to the resource change event.  | [optional] 
**Status** | Pointer to **string** | Current event status. Failing event will keep retry until it is purged. | [optional] 
**ResponseHistory** | Pointer to [**[]EventResponseHistory**](EventResponseHistory.md) | Response history of the webhook request | [optional] 

## Methods

### NewEvent

`func NewEvent() *Event`

NewEvent instantiates a new Event object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEventWithDefaults

`func NewEventWithDefaults() *Event`

NewEventWithDefaults instantiates a new Event object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Event) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Event) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Event) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Event) HasId() bool`

HasId returns a boolean if a field has been set.

### GetWebhookId

`func (o *Event) GetWebhookId() string`

GetWebhookId returns the WebhookId field if non-nil, zero value otherwise.

### GetWebhookIdOk

`func (o *Event) GetWebhookIdOk() (*string, bool)`

GetWebhookIdOk returns a tuple with the WebhookId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhookId

`func (o *Event) SetWebhookId(v string)`

SetWebhookId sets WebhookId field to given value.

### HasWebhookId

`func (o *Event) HasWebhookId() bool`

HasWebhookId returns a boolean if a field has been set.

### GetType

`func (o *Event) GetType() EventTypeExplicit`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Event) GetTypeOk() (*EventTypeExplicit, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Event) SetType(v EventTypeExplicit)`

SetType sets Type field to given value.

### HasType

`func (o *Event) HasType() bool`

HasType returns a boolean if a field has been set.

### GetEventTime

`func (o *Event) GetEventTime() time.Time`

GetEventTime returns the EventTime field if non-nil, zero value otherwise.

### GetEventTimeOk

`func (o *Event) GetEventTimeOk() (*time.Time, bool)`

GetEventTimeOk returns a tuple with the EventTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventTime

`func (o *Event) SetEventTime(v time.Time)`

SetEventTime sets EventTime field to given value.

### HasEventTime

`func (o *Event) HasEventTime() bool`

HasEventTime returns a boolean if a field has been set.

### GetMetadata

`func (o *Event) GetMetadata() string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *Event) GetMetadataOk() (*string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *Event) SetMetadata(v string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *Event) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetEventResource

`func (o *Event) GetEventResource() string`

GetEventResource returns the EventResource field if non-nil, zero value otherwise.

### GetEventResourceOk

`func (o *Event) GetEventResourceOk() (*string, bool)`

GetEventResourceOk returns a tuple with the EventResource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventResource

`func (o *Event) SetEventResource(v string)`

SetEventResource sets EventResource field to given value.

### HasEventResource

`func (o *Event) HasEventResource() bool`

HasEventResource returns a boolean if a field has been set.

### GetEventResourceChangedFields

`func (o *Event) GetEventResourceChangedFields() string`

GetEventResourceChangedFields returns the EventResourceChangedFields field if non-nil, zero value otherwise.

### GetEventResourceChangedFieldsOk

`func (o *Event) GetEventResourceChangedFieldsOk() (*string, bool)`

GetEventResourceChangedFieldsOk returns a tuple with the EventResourceChangedFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventResourceChangedFields

`func (o *Event) SetEventResourceChangedFields(v string)`

SetEventResourceChangedFields sets EventResourceChangedFields field to given value.

### HasEventResourceChangedFields

`func (o *Event) HasEventResourceChangedFields() bool`

HasEventResourceChangedFields returns a boolean if a field has been set.

### GetStatus

`func (o *Event) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Event) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Event) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *Event) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetResponseHistory

`func (o *Event) GetResponseHistory() []EventResponseHistory`

GetResponseHistory returns the ResponseHistory field if non-nil, zero value otherwise.

### GetResponseHistoryOk

`func (o *Event) GetResponseHistoryOk() (*[]EventResponseHistory, bool)`

GetResponseHistoryOk returns a tuple with the ResponseHistory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseHistory

`func (o *Event) SetResponseHistory(v []EventResponseHistory)`

SetResponseHistory sets ResponseHistory field to given value.

### HasResponseHistory

`func (o *Event) HasResponseHistory() bool`

HasResponseHistory returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


