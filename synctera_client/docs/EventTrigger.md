# EventTrigger

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EventResource** | Pointer to **string** | Json string of object associated with the event. For example, if your event is ACCOUNT.CREATED, You can refer to Acccount to parse the account event to obtain the ID, status etc.  | [optional] 
**EventResourceChangedFields** | Pointer to **string** | Json string of object associated with the event related to a resource change. This only contains those fields that have value changed on the event, and the field values are prior to the resource change event.  | [optional] 
**EventTime** | Pointer to **time.Time** | Timestamp of the current event raised | [optional] 
**Id** | Pointer to **string** | Unique event ID of the webhook request. Use event endpoints to get more event summary data | [optional] [readonly] 
**Metadata** | Pointer to **string** | Metadata that stored in the webhook subscription | [optional] 
**ResponseHistory** | Pointer to [**[]ResponseHistoryItem**](ResponseHistoryItem.md) | Response history of the webhook request | [optional] 
**Status** | Pointer to **string** | Current event status. Failing event will keep retry until it is purged. | [optional] 
**Type** | Pointer to [**EventTypeExplicit**](EventTypeExplicit.md) |  | [optional] 
**Url** | Pointer to **string** | URL that the current event will be sent to | [optional] 
**WebhookId** | Pointer to **string** | Webhook the current event belongs to | [optional] 

## Methods

### NewEventTrigger

`func NewEventTrigger() *EventTrigger`

NewEventTrigger instantiates a new EventTrigger object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEventTriggerWithDefaults

`func NewEventTriggerWithDefaults() *EventTrigger`

NewEventTriggerWithDefaults instantiates a new EventTrigger object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEventResource

`func (o *EventTrigger) GetEventResource() string`

GetEventResource returns the EventResource field if non-nil, zero value otherwise.

### GetEventResourceOk

`func (o *EventTrigger) GetEventResourceOk() (*string, bool)`

GetEventResourceOk returns a tuple with the EventResource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventResource

`func (o *EventTrigger) SetEventResource(v string)`

SetEventResource sets EventResource field to given value.

### HasEventResource

`func (o *EventTrigger) HasEventResource() bool`

HasEventResource returns a boolean if a field has been set.

### GetEventResourceChangedFields

`func (o *EventTrigger) GetEventResourceChangedFields() string`

GetEventResourceChangedFields returns the EventResourceChangedFields field if non-nil, zero value otherwise.

### GetEventResourceChangedFieldsOk

`func (o *EventTrigger) GetEventResourceChangedFieldsOk() (*string, bool)`

GetEventResourceChangedFieldsOk returns a tuple with the EventResourceChangedFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventResourceChangedFields

`func (o *EventTrigger) SetEventResourceChangedFields(v string)`

SetEventResourceChangedFields sets EventResourceChangedFields field to given value.

### HasEventResourceChangedFields

`func (o *EventTrigger) HasEventResourceChangedFields() bool`

HasEventResourceChangedFields returns a boolean if a field has been set.

### GetEventTime

`func (o *EventTrigger) GetEventTime() time.Time`

GetEventTime returns the EventTime field if non-nil, zero value otherwise.

### GetEventTimeOk

`func (o *EventTrigger) GetEventTimeOk() (*time.Time, bool)`

GetEventTimeOk returns a tuple with the EventTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventTime

`func (o *EventTrigger) SetEventTime(v time.Time)`

SetEventTime sets EventTime field to given value.

### HasEventTime

`func (o *EventTrigger) HasEventTime() bool`

HasEventTime returns a boolean if a field has been set.

### GetId

`func (o *EventTrigger) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *EventTrigger) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *EventTrigger) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *EventTrigger) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMetadata

`func (o *EventTrigger) GetMetadata() string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *EventTrigger) GetMetadataOk() (*string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *EventTrigger) SetMetadata(v string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *EventTrigger) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetResponseHistory

`func (o *EventTrigger) GetResponseHistory() []ResponseHistoryItem`

GetResponseHistory returns the ResponseHistory field if non-nil, zero value otherwise.

### GetResponseHistoryOk

`func (o *EventTrigger) GetResponseHistoryOk() (*[]ResponseHistoryItem, bool)`

GetResponseHistoryOk returns a tuple with the ResponseHistory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseHistory

`func (o *EventTrigger) SetResponseHistory(v []ResponseHistoryItem)`

SetResponseHistory sets ResponseHistory field to given value.

### HasResponseHistory

`func (o *EventTrigger) HasResponseHistory() bool`

HasResponseHistory returns a boolean if a field has been set.

### GetStatus

`func (o *EventTrigger) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *EventTrigger) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *EventTrigger) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *EventTrigger) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetType

`func (o *EventTrigger) GetType() EventTypeExplicit`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *EventTrigger) GetTypeOk() (*EventTypeExplicit, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *EventTrigger) SetType(v EventTypeExplicit)`

SetType sets Type field to given value.

### HasType

`func (o *EventTrigger) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUrl

`func (o *EventTrigger) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *EventTrigger) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *EventTrigger) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *EventTrigger) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetWebhookId

`func (o *EventTrigger) GetWebhookId() string`

GetWebhookId returns the WebhookId field if non-nil, zero value otherwise.

### GetWebhookIdOk

`func (o *EventTrigger) GetWebhookIdOk() (*string, bool)`

GetWebhookIdOk returns a tuple with the WebhookId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhookId

`func (o *EventTrigger) SetWebhookId(v string)`

SetWebhookId sets WebhookId field to given value.

### HasWebhookId

`func (o *EventTrigger) HasWebhookId() bool`

HasWebhookId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


