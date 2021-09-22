# PingResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**WebhookStatus** | **string** | status of webhook endpoint | 

## Methods

### NewPingResponse

`func NewPingResponse(webhookStatus string, ) *PingResponse`

NewPingResponse instantiates a new PingResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPingResponseWithDefaults

`func NewPingResponseWithDefaults() *PingResponse`

NewPingResponseWithDefaults instantiates a new PingResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWebhookStatus

`func (o *PingResponse) GetWebhookStatus() string`

GetWebhookStatus returns the WebhookStatus field if non-nil, zero value otherwise.

### GetWebhookStatusOk

`func (o *PingResponse) GetWebhookStatusOk() (*string, bool)`

GetWebhookStatusOk returns a tuple with the WebhookStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhookStatus

`func (o *PingResponse) SetWebhookStatus(v string)`

SetWebhookStatus sets WebhookStatus field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


