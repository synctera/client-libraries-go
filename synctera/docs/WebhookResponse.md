# WebhookResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Active** | Pointer to **bool** | indicates whether webhook is active | [optional] 
**Config** | Pointer to [**WebhookConfig**](WebhookConfig.md) |  | [optional] 
**Events** | Pointer to **[]string** | list of webhook events, use * to receive all notifications | [optional] 
**Id** | Pointer to **string** | id of the webhook | [optional] 
**Name** | Pointer to **string** | name of the webhook | [optional] 
**CreationTime** | Pointer to **time.Time** | The timestamp representing when the webhook request was made | [optional] [readonly] 
**LastModifiedTime** | Pointer to **time.Time** | The timestamp representing when the webhook was last modified | [optional] [readonly] 

## Methods

### NewWebhookResponse

`func NewWebhookResponse() *WebhookResponse`

NewWebhookResponse instantiates a new WebhookResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebhookResponseWithDefaults

`func NewWebhookResponseWithDefaults() *WebhookResponse`

NewWebhookResponseWithDefaults instantiates a new WebhookResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActive

`func (o *WebhookResponse) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *WebhookResponse) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *WebhookResponse) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *WebhookResponse) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetConfig

`func (o *WebhookResponse) GetConfig() WebhookConfig`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *WebhookResponse) GetConfigOk() (*WebhookConfig, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *WebhookResponse) SetConfig(v WebhookConfig)`

SetConfig sets Config field to given value.

### HasConfig

`func (o *WebhookResponse) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetEvents

`func (o *WebhookResponse) GetEvents() []string`

GetEvents returns the Events field if non-nil, zero value otherwise.

### GetEventsOk

`func (o *WebhookResponse) GetEventsOk() (*[]string, bool)`

GetEventsOk returns a tuple with the Events field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvents

`func (o *WebhookResponse) SetEvents(v []string)`

SetEvents sets Events field to given value.

### HasEvents

`func (o *WebhookResponse) HasEvents() bool`

HasEvents returns a boolean if a field has been set.

### GetId

`func (o *WebhookResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *WebhookResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *WebhookResponse) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *WebhookResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *WebhookResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *WebhookResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *WebhookResponse) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *WebhookResponse) HasName() bool`

HasName returns a boolean if a field has been set.

### GetCreationTime

`func (o *WebhookResponse) GetCreationTime() time.Time`

GetCreationTime returns the CreationTime field if non-nil, zero value otherwise.

### GetCreationTimeOk

`func (o *WebhookResponse) GetCreationTimeOk() (*time.Time, bool)`

GetCreationTimeOk returns a tuple with the CreationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationTime

`func (o *WebhookResponse) SetCreationTime(v time.Time)`

SetCreationTime sets CreationTime field to given value.

### HasCreationTime

`func (o *WebhookResponse) HasCreationTime() bool`

HasCreationTime returns a boolean if a field has been set.

### GetLastModifiedTime

`func (o *WebhookResponse) GetLastModifiedTime() time.Time`

GetLastModifiedTime returns the LastModifiedTime field if non-nil, zero value otherwise.

### GetLastModifiedTimeOk

`func (o *WebhookResponse) GetLastModifiedTimeOk() (*time.Time, bool)`

GetLastModifiedTimeOk returns a tuple with the LastModifiedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModifiedTime

`func (o *WebhookResponse) SetLastModifiedTime(v time.Time)`

SetLastModifiedTime sets LastModifiedTime field to given value.

### HasLastModifiedTime

`func (o *WebhookResponse) HasLastModifiedTime() bool`

HasLastModifiedTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


