# SocureEventBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Unique identifier for the monitoring event | 
**EnvironmentName** | **string** | Environment the event belongs to | 
**Event** | [**SocureWatchlistResult**](SocureWatchlistResult.md) |  | 

## Methods

### NewSocureEventBody

`func NewSocureEventBody(id string, environmentName string, event SocureWatchlistResult, ) *SocureEventBody`

NewSocureEventBody instantiates a new SocureEventBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSocureEventBodyWithDefaults

`func NewSocureEventBodyWithDefaults() *SocureEventBody`

NewSocureEventBodyWithDefaults instantiates a new SocureEventBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SocureEventBody) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SocureEventBody) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SocureEventBody) SetId(v string)`

SetId sets Id field to given value.


### GetEnvironmentName

`func (o *SocureEventBody) GetEnvironmentName() string`

GetEnvironmentName returns the EnvironmentName field if non-nil, zero value otherwise.

### GetEnvironmentNameOk

`func (o *SocureEventBody) GetEnvironmentNameOk() (*string, bool)`

GetEnvironmentNameOk returns a tuple with the EnvironmentName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentName

`func (o *SocureEventBody) SetEnvironmentName(v string)`

SetEnvironmentName sets EnvironmentName field to given value.


### GetEvent

`func (o *SocureEventBody) GetEvent() SocureWatchlistResult`

GetEvent returns the Event field if non-nil, zero value otherwise.

### GetEventOk

`func (o *SocureEventBody) GetEventOk() (*SocureWatchlistResult, bool)`

GetEventOk returns a tuple with the Event field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvent

`func (o *SocureEventBody) SetEvent(v SocureWatchlistResult)`

SetEvent sets Event field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


