# MonitoringAlertList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Alerts** | [**[]MonitoringAlert**](MonitoringAlert.md) |  | 
**NextPageToken** | Pointer to **string** | If returned, use the next_page_token to query for the next page of results. Not returned if there are no more rows. | [optional] 

## Methods

### NewMonitoringAlertList

`func NewMonitoringAlertList(alerts []MonitoringAlert, ) *MonitoringAlertList`

NewMonitoringAlertList instantiates a new MonitoringAlertList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMonitoringAlertListWithDefaults

`func NewMonitoringAlertListWithDefaults() *MonitoringAlertList`

NewMonitoringAlertListWithDefaults instantiates a new MonitoringAlertList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAlerts

`func (o *MonitoringAlertList) GetAlerts() []MonitoringAlert`

GetAlerts returns the Alerts field if non-nil, zero value otherwise.

### GetAlertsOk

`func (o *MonitoringAlertList) GetAlertsOk() (*[]MonitoringAlert, bool)`

GetAlertsOk returns a tuple with the Alerts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlerts

`func (o *MonitoringAlertList) SetAlerts(v []MonitoringAlert)`

SetAlerts sets Alerts field to given value.


### GetNextPageToken

`func (o *MonitoringAlertList) GetNextPageToken() string`

GetNextPageToken returns the NextPageToken field if non-nil, zero value otherwise.

### GetNextPageTokenOk

`func (o *MonitoringAlertList) GetNextPageTokenOk() (*string, bool)`

GetNextPageTokenOk returns a tuple with the NextPageToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextPageToken

`func (o *MonitoringAlertList) SetNextPageToken(v string)`

SetNextPageToken sets NextPageToken field to given value.

### HasNextPageToken

`func (o *MonitoringAlertList) HasNextPageToken() bool`

HasNextPageToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


