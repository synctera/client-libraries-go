# WatchlistAlert

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Created** | Pointer to **time.Time** | When this alert was created | [optional] 
**Id** | Pointer to **string** | Unique identifier for this alert | [optional] 
**ProviderInfo** | Pointer to **map[string]interface{}** | The information provided to Synctera that triggered this alert, as an arbitrary JSON object. Interpretation of this object is up to the client.  | [optional] 
**ProviderSubjectId** | Pointer to **string** | The id of the provider subject for this alert | [optional] 
**ProviderSubscriptionId** | Pointer to **string** | The id of the provider subscription for this alert | [optional] 
**ProviderWatchlistName** | Pointer to **string** | The name of the provider for this alert | [optional] 
**Status** | **string** | The status of this alert | 
**Urls** | Pointer to **[]string** | Where to get more information about this alert (according to our third-party data provider).  | [optional] 
**VendorInfo** | Pointer to [**VendorInfo1**](VendorInfo1.md) |  | [optional] 

## Methods

### NewWatchlistAlert

`func NewWatchlistAlert(status string, ) *WatchlistAlert`

NewWatchlistAlert instantiates a new WatchlistAlert object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWatchlistAlertWithDefaults

`func NewWatchlistAlertWithDefaults() *WatchlistAlert`

NewWatchlistAlertWithDefaults instantiates a new WatchlistAlert object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreated

`func (o *WatchlistAlert) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *WatchlistAlert) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *WatchlistAlert) SetCreated(v time.Time)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *WatchlistAlert) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetId

`func (o *WatchlistAlert) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *WatchlistAlert) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *WatchlistAlert) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *WatchlistAlert) HasId() bool`

HasId returns a boolean if a field has been set.

### GetProviderInfo

`func (o *WatchlistAlert) GetProviderInfo() map[string]interface{}`

GetProviderInfo returns the ProviderInfo field if non-nil, zero value otherwise.

### GetProviderInfoOk

`func (o *WatchlistAlert) GetProviderInfoOk() (*map[string]interface{}, bool)`

GetProviderInfoOk returns a tuple with the ProviderInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderInfo

`func (o *WatchlistAlert) SetProviderInfo(v map[string]interface{})`

SetProviderInfo sets ProviderInfo field to given value.

### HasProviderInfo

`func (o *WatchlistAlert) HasProviderInfo() bool`

HasProviderInfo returns a boolean if a field has been set.

### GetProviderSubjectId

`func (o *WatchlistAlert) GetProviderSubjectId() string`

GetProviderSubjectId returns the ProviderSubjectId field if non-nil, zero value otherwise.

### GetProviderSubjectIdOk

`func (o *WatchlistAlert) GetProviderSubjectIdOk() (*string, bool)`

GetProviderSubjectIdOk returns a tuple with the ProviderSubjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderSubjectId

`func (o *WatchlistAlert) SetProviderSubjectId(v string)`

SetProviderSubjectId sets ProviderSubjectId field to given value.

### HasProviderSubjectId

`func (o *WatchlistAlert) HasProviderSubjectId() bool`

HasProviderSubjectId returns a boolean if a field has been set.

### GetProviderSubscriptionId

`func (o *WatchlistAlert) GetProviderSubscriptionId() string`

GetProviderSubscriptionId returns the ProviderSubscriptionId field if non-nil, zero value otherwise.

### GetProviderSubscriptionIdOk

`func (o *WatchlistAlert) GetProviderSubscriptionIdOk() (*string, bool)`

GetProviderSubscriptionIdOk returns a tuple with the ProviderSubscriptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderSubscriptionId

`func (o *WatchlistAlert) SetProviderSubscriptionId(v string)`

SetProviderSubscriptionId sets ProviderSubscriptionId field to given value.

### HasProviderSubscriptionId

`func (o *WatchlistAlert) HasProviderSubscriptionId() bool`

HasProviderSubscriptionId returns a boolean if a field has been set.

### GetProviderWatchlistName

`func (o *WatchlistAlert) GetProviderWatchlistName() string`

GetProviderWatchlistName returns the ProviderWatchlistName field if non-nil, zero value otherwise.

### GetProviderWatchlistNameOk

`func (o *WatchlistAlert) GetProviderWatchlistNameOk() (*string, bool)`

GetProviderWatchlistNameOk returns a tuple with the ProviderWatchlistName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderWatchlistName

`func (o *WatchlistAlert) SetProviderWatchlistName(v string)`

SetProviderWatchlistName sets ProviderWatchlistName field to given value.

### HasProviderWatchlistName

`func (o *WatchlistAlert) HasProviderWatchlistName() bool`

HasProviderWatchlistName returns a boolean if a field has been set.

### GetStatus

`func (o *WatchlistAlert) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *WatchlistAlert) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *WatchlistAlert) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetUrls

`func (o *WatchlistAlert) GetUrls() []string`

GetUrls returns the Urls field if non-nil, zero value otherwise.

### GetUrlsOk

`func (o *WatchlistAlert) GetUrlsOk() (*[]string, bool)`

GetUrlsOk returns a tuple with the Urls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrls

`func (o *WatchlistAlert) SetUrls(v []string)`

SetUrls sets Urls field to given value.

### HasUrls

`func (o *WatchlistAlert) HasUrls() bool`

HasUrls returns a boolean if a field has been set.

### GetVendorInfo

`func (o *WatchlistAlert) GetVendorInfo() VendorInfo1`

GetVendorInfo returns the VendorInfo field if non-nil, zero value otherwise.

### GetVendorInfoOk

`func (o *WatchlistAlert) GetVendorInfoOk() (*VendorInfo1, bool)`

GetVendorInfoOk returns a tuple with the VendorInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendorInfo

`func (o *WatchlistAlert) SetVendorInfo(v VendorInfo1)`

SetVendorInfo sets VendorInfo field to given value.

### HasVendorInfo

`func (o *WatchlistAlert) HasVendorInfo() bool`

HasVendorInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


