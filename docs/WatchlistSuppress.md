# WatchlistSuppress

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProviderSubjectId** | **string** | The id of the subject (person) for whom future alerts should be suppressed.  | 
**ProviderSubscriptionId** | **string** | The provider&#39;s id for the subscription that caused the alert(s) that are being suppressed.  | 
**Status** | **string** | The status of this suppression | 

## Methods

### NewWatchlistSuppress

`func NewWatchlistSuppress(providerSubjectId string, providerSubscriptionId string, status string, ) *WatchlistSuppress`

NewWatchlistSuppress instantiates a new WatchlistSuppress object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWatchlistSuppressWithDefaults

`func NewWatchlistSuppressWithDefaults() *WatchlistSuppress`

NewWatchlistSuppressWithDefaults instantiates a new WatchlistSuppress object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProviderSubjectId

`func (o *WatchlistSuppress) GetProviderSubjectId() string`

GetProviderSubjectId returns the ProviderSubjectId field if non-nil, zero value otherwise.

### GetProviderSubjectIdOk

`func (o *WatchlistSuppress) GetProviderSubjectIdOk() (*string, bool)`

GetProviderSubjectIdOk returns a tuple with the ProviderSubjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderSubjectId

`func (o *WatchlistSuppress) SetProviderSubjectId(v string)`

SetProviderSubjectId sets ProviderSubjectId field to given value.


### GetProviderSubscriptionId

`func (o *WatchlistSuppress) GetProviderSubscriptionId() string`

GetProviderSubscriptionId returns the ProviderSubscriptionId field if non-nil, zero value otherwise.

### GetProviderSubscriptionIdOk

`func (o *WatchlistSuppress) GetProviderSubscriptionIdOk() (*string, bool)`

GetProviderSubscriptionIdOk returns a tuple with the ProviderSubscriptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderSubscriptionId

`func (o *WatchlistSuppress) SetProviderSubscriptionId(v string)`

SetProviderSubscriptionId sets ProviderSubscriptionId field to given value.


### GetStatus

`func (o *WatchlistSuppress) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *WatchlistSuppress) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *WatchlistSuppress) SetStatus(v string)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


