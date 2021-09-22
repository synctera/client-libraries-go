# CreateWebhookRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Active** | Pointer to **bool** | indicates whether webhook is active | [optional] 
**Config** | [**WebhookConfig**](WebhookConfig.md) |  | 
**Events** | **[]string** | list of webhook events, use * to receive all notifications | 
**Id** | Pointer to **string** | id of the webhook | [optional] 
**Name** | **string** | name of the webhook | 

## Methods

### NewCreateWebhookRequest

`func NewCreateWebhookRequest(config WebhookConfig, events []string, name string, ) *CreateWebhookRequest`

NewCreateWebhookRequest instantiates a new CreateWebhookRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateWebhookRequestWithDefaults

`func NewCreateWebhookRequestWithDefaults() *CreateWebhookRequest`

NewCreateWebhookRequestWithDefaults instantiates a new CreateWebhookRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActive

`func (o *CreateWebhookRequest) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *CreateWebhookRequest) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *CreateWebhookRequest) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *CreateWebhookRequest) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetConfig

`func (o *CreateWebhookRequest) GetConfig() WebhookConfig`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *CreateWebhookRequest) GetConfigOk() (*WebhookConfig, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *CreateWebhookRequest) SetConfig(v WebhookConfig)`

SetConfig sets Config field to given value.


### GetEvents

`func (o *CreateWebhookRequest) GetEvents() []string`

GetEvents returns the Events field if non-nil, zero value otherwise.

### GetEventsOk

`func (o *CreateWebhookRequest) GetEventsOk() (*[]string, bool)`

GetEventsOk returns a tuple with the Events field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvents

`func (o *CreateWebhookRequest) SetEvents(v []string)`

SetEvents sets Events field to given value.


### GetId

`func (o *CreateWebhookRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CreateWebhookRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CreateWebhookRequest) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CreateWebhookRequest) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *CreateWebhookRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateWebhookRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateWebhookRequest) SetName(v string)`

SetName sets Name field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


