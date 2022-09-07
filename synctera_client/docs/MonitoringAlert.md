# MonitoringAlert

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BusinessId** | Pointer to **string** | Unique ID for the business. Exactly one of &#x60;business_id&#x60; or &#x60;person_id&#x60; must be set.  | [optional] 
**CreationTime** | Pointer to **time.Time** | The date and time the resource was created. | [optional] [readonly] 
**Id** | Pointer to **string** | Unique identifier for this alert. | [optional] [readonly] 
**LastUpdatedTime** | Pointer to **time.Time** | The date and time the resource was last update. | [optional] [readonly] 
**Metadata** | Pointer to **map[string]interface{}** | Optional field to store additional information about the resource. Intended to be used by the integrator to store non-sensitive data.  | [optional] 
**PersonId** | Pointer to **string** | Unique ID for the person. Exactly one of &#x60;person_id&#x60; or &#x60;business_id&#x60; must be set.  | [optional] 
**Status** | Pointer to [**MonitoringStatus**](MonitoringStatus.md) |  | [optional] 
**Type** | Pointer to **string** | The type of customer alert. Any of the following: * &#x60;WATCHLIST&#x60; – the customer was added to a known watchlist. * &#x60;BANKRUPTCY&#x60; – the customer filed for bankruptcy.  | [optional] [readonly] 
**Urls** | Pointer to **[]string** | Where to get more information about this alert. | [optional] [readonly] 
**VendorInfo** | Pointer to [**VendorInfo**](VendorInfo.md) |  | [optional] 

## Methods

### NewMonitoringAlert

`func NewMonitoringAlert() *MonitoringAlert`

NewMonitoringAlert instantiates a new MonitoringAlert object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMonitoringAlertWithDefaults

`func NewMonitoringAlertWithDefaults() *MonitoringAlert`

NewMonitoringAlertWithDefaults instantiates a new MonitoringAlert object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBusinessId

`func (o *MonitoringAlert) GetBusinessId() string`

GetBusinessId returns the BusinessId field if non-nil, zero value otherwise.

### GetBusinessIdOk

`func (o *MonitoringAlert) GetBusinessIdOk() (*string, bool)`

GetBusinessIdOk returns a tuple with the BusinessId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessId

`func (o *MonitoringAlert) SetBusinessId(v string)`

SetBusinessId sets BusinessId field to given value.

### HasBusinessId

`func (o *MonitoringAlert) HasBusinessId() bool`

HasBusinessId returns a boolean if a field has been set.

### GetCreationTime

`func (o *MonitoringAlert) GetCreationTime() time.Time`

GetCreationTime returns the CreationTime field if non-nil, zero value otherwise.

### GetCreationTimeOk

`func (o *MonitoringAlert) GetCreationTimeOk() (*time.Time, bool)`

GetCreationTimeOk returns a tuple with the CreationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationTime

`func (o *MonitoringAlert) SetCreationTime(v time.Time)`

SetCreationTime sets CreationTime field to given value.

### HasCreationTime

`func (o *MonitoringAlert) HasCreationTime() bool`

HasCreationTime returns a boolean if a field has been set.

### GetId

`func (o *MonitoringAlert) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MonitoringAlert) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MonitoringAlert) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *MonitoringAlert) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLastUpdatedTime

`func (o *MonitoringAlert) GetLastUpdatedTime() time.Time`

GetLastUpdatedTime returns the LastUpdatedTime field if non-nil, zero value otherwise.

### GetLastUpdatedTimeOk

`func (o *MonitoringAlert) GetLastUpdatedTimeOk() (*time.Time, bool)`

GetLastUpdatedTimeOk returns a tuple with the LastUpdatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdatedTime

`func (o *MonitoringAlert) SetLastUpdatedTime(v time.Time)`

SetLastUpdatedTime sets LastUpdatedTime field to given value.

### HasLastUpdatedTime

`func (o *MonitoringAlert) HasLastUpdatedTime() bool`

HasLastUpdatedTime returns a boolean if a field has been set.

### GetMetadata

`func (o *MonitoringAlert) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *MonitoringAlert) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *MonitoringAlert) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *MonitoringAlert) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetPersonId

`func (o *MonitoringAlert) GetPersonId() string`

GetPersonId returns the PersonId field if non-nil, zero value otherwise.

### GetPersonIdOk

`func (o *MonitoringAlert) GetPersonIdOk() (*string, bool)`

GetPersonIdOk returns a tuple with the PersonId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPersonId

`func (o *MonitoringAlert) SetPersonId(v string)`

SetPersonId sets PersonId field to given value.

### HasPersonId

`func (o *MonitoringAlert) HasPersonId() bool`

HasPersonId returns a boolean if a field has been set.

### GetStatus

`func (o *MonitoringAlert) GetStatus() MonitoringStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *MonitoringAlert) GetStatusOk() (*MonitoringStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *MonitoringAlert) SetStatus(v MonitoringStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *MonitoringAlert) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetType

`func (o *MonitoringAlert) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *MonitoringAlert) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *MonitoringAlert) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *MonitoringAlert) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUrls

`func (o *MonitoringAlert) GetUrls() []string`

GetUrls returns the Urls field if non-nil, zero value otherwise.

### GetUrlsOk

`func (o *MonitoringAlert) GetUrlsOk() (*[]string, bool)`

GetUrlsOk returns a tuple with the Urls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrls

`func (o *MonitoringAlert) SetUrls(v []string)`

SetUrls sets Urls field to given value.

### HasUrls

`func (o *MonitoringAlert) HasUrls() bool`

HasUrls returns a boolean if a field has been set.

### GetVendorInfo

`func (o *MonitoringAlert) GetVendorInfo() VendorInfo`

GetVendorInfo returns the VendorInfo field if non-nil, zero value otherwise.

### GetVendorInfoOk

`func (o *MonitoringAlert) GetVendorInfoOk() (*VendorInfo, bool)`

GetVendorInfoOk returns a tuple with the VendorInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendorInfo

`func (o *MonitoringAlert) SetVendorInfo(v VendorInfo)`

SetVendorInfo sets VendorInfo field to given value.

### HasVendorInfo

`func (o *MonitoringAlert) HasVendorInfo() bool`

HasVendorInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


