# WebhookConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CustomHeader** | Pointer to [**CustomHeaders**](CustomHeaders.md) |  | [optional] 
**Password** | Pointer to **string** | password for access webhook endpoint | [optional] 
**Url** | Pointer to **string** | url of webhook endpoint | [optional] 
**Username** | Pointer to **string** | username for access webhook endpoint | [optional] 

## Methods

### NewWebhookConfig

`func NewWebhookConfig() *WebhookConfig`

NewWebhookConfig instantiates a new WebhookConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebhookConfigWithDefaults

`func NewWebhookConfigWithDefaults() *WebhookConfig`

NewWebhookConfigWithDefaults instantiates a new WebhookConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCustomHeader

`func (o *WebhookConfig) GetCustomHeader() CustomHeaders`

GetCustomHeader returns the CustomHeader field if non-nil, zero value otherwise.

### GetCustomHeaderOk

`func (o *WebhookConfig) GetCustomHeaderOk() (*CustomHeaders, bool)`

GetCustomHeaderOk returns a tuple with the CustomHeader field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomHeader

`func (o *WebhookConfig) SetCustomHeader(v CustomHeaders)`

SetCustomHeader sets CustomHeader field to given value.

### HasCustomHeader

`func (o *WebhookConfig) HasCustomHeader() bool`

HasCustomHeader returns a boolean if a field has been set.

### GetPassword

`func (o *WebhookConfig) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *WebhookConfig) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *WebhookConfig) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *WebhookConfig) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetUrl

`func (o *WebhookConfig) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *WebhookConfig) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *WebhookConfig) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *WebhookConfig) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetUsername

`func (o *WebhookConfig) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *WebhookConfig) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *WebhookConfig) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *WebhookConfig) HasUsername() bool`

HasUsername returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


