# WatchlistSubscription

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AutoRenew** | Pointer to **bool** | Whether this subscription should automatically renew when the subscription period is over (default: vendor-dependent).  | [optional] 
**Created** | Pointer to **time.Time** | When this subscription was created | [optional] 
**Id** | Pointer to **string** | Unique identifier for this subscription | [optional] 
**PeriodEnd** | Pointer to **string** | The date when monitoring of this individual should end. | [optional] 
**PeriodStart** | Pointer to **string** | The date when monitoring of this individual should begin (default: today). | [optional] 
**ProviderSubscriptionId** | Pointer to **string** | External provider subscription id | [optional] 
**Status** | Pointer to **string** |  | [optional] 

## Methods

### NewWatchlistSubscription

`func NewWatchlistSubscription() *WatchlistSubscription`

NewWatchlistSubscription instantiates a new WatchlistSubscription object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWatchlistSubscriptionWithDefaults

`func NewWatchlistSubscriptionWithDefaults() *WatchlistSubscription`

NewWatchlistSubscriptionWithDefaults instantiates a new WatchlistSubscription object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAutoRenew

`func (o *WatchlistSubscription) GetAutoRenew() bool`

GetAutoRenew returns the AutoRenew field if non-nil, zero value otherwise.

### GetAutoRenewOk

`func (o *WatchlistSubscription) GetAutoRenewOk() (*bool, bool)`

GetAutoRenewOk returns a tuple with the AutoRenew field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoRenew

`func (o *WatchlistSubscription) SetAutoRenew(v bool)`

SetAutoRenew sets AutoRenew field to given value.

### HasAutoRenew

`func (o *WatchlistSubscription) HasAutoRenew() bool`

HasAutoRenew returns a boolean if a field has been set.

### GetCreated

`func (o *WatchlistSubscription) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *WatchlistSubscription) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *WatchlistSubscription) SetCreated(v time.Time)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *WatchlistSubscription) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetId

`func (o *WatchlistSubscription) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *WatchlistSubscription) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *WatchlistSubscription) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *WatchlistSubscription) HasId() bool`

HasId returns a boolean if a field has been set.

### GetPeriodEnd

`func (o *WatchlistSubscription) GetPeriodEnd() string`

GetPeriodEnd returns the PeriodEnd field if non-nil, zero value otherwise.

### GetPeriodEndOk

`func (o *WatchlistSubscription) GetPeriodEndOk() (*string, bool)`

GetPeriodEndOk returns a tuple with the PeriodEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriodEnd

`func (o *WatchlistSubscription) SetPeriodEnd(v string)`

SetPeriodEnd sets PeriodEnd field to given value.

### HasPeriodEnd

`func (o *WatchlistSubscription) HasPeriodEnd() bool`

HasPeriodEnd returns a boolean if a field has been set.

### GetPeriodStart

`func (o *WatchlistSubscription) GetPeriodStart() string`

GetPeriodStart returns the PeriodStart field if non-nil, zero value otherwise.

### GetPeriodStartOk

`func (o *WatchlistSubscription) GetPeriodStartOk() (*string, bool)`

GetPeriodStartOk returns a tuple with the PeriodStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriodStart

`func (o *WatchlistSubscription) SetPeriodStart(v string)`

SetPeriodStart sets PeriodStart field to given value.

### HasPeriodStart

`func (o *WatchlistSubscription) HasPeriodStart() bool`

HasPeriodStart returns a boolean if a field has been set.

### GetProviderSubscriptionId

`func (o *WatchlistSubscription) GetProviderSubscriptionId() string`

GetProviderSubscriptionId returns the ProviderSubscriptionId field if non-nil, zero value otherwise.

### GetProviderSubscriptionIdOk

`func (o *WatchlistSubscription) GetProviderSubscriptionIdOk() (*string, bool)`

GetProviderSubscriptionIdOk returns a tuple with the ProviderSubscriptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderSubscriptionId

`func (o *WatchlistSubscription) SetProviderSubscriptionId(v string)`

SetProviderSubscriptionId sets ProviderSubscriptionId field to given value.

### HasProviderSubscriptionId

`func (o *WatchlistSubscription) HasProviderSubscriptionId() bool`

HasProviderSubscriptionId returns a boolean if a field has been set.

### GetStatus

`func (o *WatchlistSubscription) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *WatchlistSubscription) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *WatchlistSubscription) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *WatchlistSubscription) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


