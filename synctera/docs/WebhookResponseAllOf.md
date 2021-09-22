# WebhookResponseAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreationTime** | Pointer to **time.Time** | The timestamp representing when the webhook request was made | [optional] [readonly] 
**LastModifiedTime** | Pointer to **time.Time** | The timestamp representing when the webhook was last modified | [optional] [readonly] 

## Methods

### NewWebhookResponseAllOf

`func NewWebhookResponseAllOf() *WebhookResponseAllOf`

NewWebhookResponseAllOf instantiates a new WebhookResponseAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebhookResponseAllOfWithDefaults

`func NewWebhookResponseAllOfWithDefaults() *WebhookResponseAllOf`

NewWebhookResponseAllOfWithDefaults instantiates a new WebhookResponseAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreationTime

`func (o *WebhookResponseAllOf) GetCreationTime() time.Time`

GetCreationTime returns the CreationTime field if non-nil, zero value otherwise.

### GetCreationTimeOk

`func (o *WebhookResponseAllOf) GetCreationTimeOk() (*time.Time, bool)`

GetCreationTimeOk returns a tuple with the CreationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationTime

`func (o *WebhookResponseAllOf) SetCreationTime(v time.Time)`

SetCreationTime sets CreationTime field to given value.

### HasCreationTime

`func (o *WebhookResponseAllOf) HasCreationTime() bool`

HasCreationTime returns a boolean if a field has been set.

### GetLastModifiedTime

`func (o *WebhookResponseAllOf) GetLastModifiedTime() time.Time`

GetLastModifiedTime returns the LastModifiedTime field if non-nil, zero value otherwise.

### GetLastModifiedTimeOk

`func (o *WebhookResponseAllOf) GetLastModifiedTimeOk() (*time.Time, bool)`

GetLastModifiedTimeOk returns a tuple with the LastModifiedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModifiedTime

`func (o *WebhookResponseAllOf) SetLastModifiedTime(v time.Time)`

SetLastModifiedTime sets LastModifiedTime field to given value.

### HasLastModifiedTime

`func (o *WebhookResponseAllOf) HasLastModifiedTime() bool`

HasLastModifiedTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


