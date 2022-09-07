# MonitoringSubscriptionList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Subscriptions** | [**[]MonitoringSubscription**](MonitoringSubscription.md) |  | 
**NextPageToken** | Pointer to **string** | If returned, use the next_page_token to query for the next page of results. Not returned if there are no more rows. | [optional] 

## Methods

### NewMonitoringSubscriptionList

`func NewMonitoringSubscriptionList(subscriptions []MonitoringSubscription, ) *MonitoringSubscriptionList`

NewMonitoringSubscriptionList instantiates a new MonitoringSubscriptionList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMonitoringSubscriptionListWithDefaults

`func NewMonitoringSubscriptionListWithDefaults() *MonitoringSubscriptionList`

NewMonitoringSubscriptionListWithDefaults instantiates a new MonitoringSubscriptionList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSubscriptions

`func (o *MonitoringSubscriptionList) GetSubscriptions() []MonitoringSubscription`

GetSubscriptions returns the Subscriptions field if non-nil, zero value otherwise.

### GetSubscriptionsOk

`func (o *MonitoringSubscriptionList) GetSubscriptionsOk() (*[]MonitoringSubscription, bool)`

GetSubscriptionsOk returns a tuple with the Subscriptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptions

`func (o *MonitoringSubscriptionList) SetSubscriptions(v []MonitoringSubscription)`

SetSubscriptions sets Subscriptions field to given value.


### GetNextPageToken

`func (o *MonitoringSubscriptionList) GetNextPageToken() string`

GetNextPageToken returns the NextPageToken field if non-nil, zero value otherwise.

### GetNextPageTokenOk

`func (o *MonitoringSubscriptionList) GetNextPageTokenOk() (*string, bool)`

GetNextPageTokenOk returns a tuple with the NextPageToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextPageToken

`func (o *MonitoringSubscriptionList) SetNextPageToken(v string)`

SetNextPageToken sets NextPageToken field to given value.

### HasNextPageToken

`func (o *MonitoringSubscriptionList) HasNextPageToken() bool`

HasNextPageToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


