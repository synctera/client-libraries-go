# WebhookRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | id of the webhook | [optional] 
**Name** | Pointer to **string** | name of the webhook | [optional] 
**Active** | Pointer to **bool** | indicates whether webhook is active | [optional] 
**Events** | Pointer to **[]string** | list of webhook events, use * to receive all notifications | [optional] 
**Config** | Pointer to [**WebhookConfig**](WebhookConfig.md) |  | [optional] 

## Methods

### NewWebhookRequest

`func NewWebhookRequest() *WebhookRequest`

NewWebhookRequest instantiates a new WebhookRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebhookRequestWithDefaults

`func NewWebhookRequestWithDefaults() *WebhookRequest`

NewWebhookRequestWithDefaults instantiates a new WebhookRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *WebhookRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *WebhookRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *WebhookRequest) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *WebhookRequest) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *WebhookRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *WebhookRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *WebhookRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *WebhookRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetActive

`func (o *WebhookRequest) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *WebhookRequest) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *WebhookRequest) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *WebhookRequest) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetEvents

`func (o *WebhookRequest) GetEvents() []string`

GetEvents returns the Events field if non-nil, zero value otherwise.

### GetEventsOk

`func (o *WebhookRequest) GetEventsOk() (*[]string, bool)`

GetEventsOk returns a tuple with the Events field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvents

`func (o *WebhookRequest) SetEvents(v []string)`

SetEvents sets Events field to given value.

### HasEvents

`func (o *WebhookRequest) HasEvents() bool`

HasEvents returns a boolean if a field has been set.

### GetConfig

`func (o *WebhookRequest) GetConfig() WebhookConfig`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *WebhookRequest) GetConfigOk() (*WebhookConfig, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *WebhookRequest) SetConfig(v WebhookConfig)`

SetConfig sets Config field to given value.

### HasConfig

`func (o *WebhookRequest) HasConfig() bool`

HasConfig returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


