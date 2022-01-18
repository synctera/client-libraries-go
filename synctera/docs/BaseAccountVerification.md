# BaseAccountVerification

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreationTime** | Pointer to **time.Time** | The time at which verification was first completed. | [optional] 
**LastUpdatedTime** | Pointer to **time.Time** | The time at which verification was last updated. | [optional] 
**Status** | **string** | The status of verification | 
**Vendor** | **string** | The vendor used for verifying the account | 

## Methods

### NewBaseAccountVerification

`func NewBaseAccountVerification(status string, vendor string, ) *BaseAccountVerification`

NewBaseAccountVerification instantiates a new BaseAccountVerification object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBaseAccountVerificationWithDefaults

`func NewBaseAccountVerificationWithDefaults() *BaseAccountVerification`

NewBaseAccountVerificationWithDefaults instantiates a new BaseAccountVerification object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreationTime

`func (o *BaseAccountVerification) GetCreationTime() time.Time`

GetCreationTime returns the CreationTime field if non-nil, zero value otherwise.

### GetCreationTimeOk

`func (o *BaseAccountVerification) GetCreationTimeOk() (*time.Time, bool)`

GetCreationTimeOk returns a tuple with the CreationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationTime

`func (o *BaseAccountVerification) SetCreationTime(v time.Time)`

SetCreationTime sets CreationTime field to given value.

### HasCreationTime

`func (o *BaseAccountVerification) HasCreationTime() bool`

HasCreationTime returns a boolean if a field has been set.

### GetLastUpdatedTime

`func (o *BaseAccountVerification) GetLastUpdatedTime() time.Time`

GetLastUpdatedTime returns the LastUpdatedTime field if non-nil, zero value otherwise.

### GetLastUpdatedTimeOk

`func (o *BaseAccountVerification) GetLastUpdatedTimeOk() (*time.Time, bool)`

GetLastUpdatedTimeOk returns a tuple with the LastUpdatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdatedTime

`func (o *BaseAccountVerification) SetLastUpdatedTime(v time.Time)`

SetLastUpdatedTime sets LastUpdatedTime field to given value.

### HasLastUpdatedTime

`func (o *BaseAccountVerification) HasLastUpdatedTime() bool`

HasLastUpdatedTime returns a boolean if a field has been set.

### GetStatus

`func (o *BaseAccountVerification) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *BaseAccountVerification) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *BaseAccountVerification) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetVendor

`func (o *BaseAccountVerification) GetVendor() string`

GetVendor returns the Vendor field if non-nil, zero value otherwise.

### GetVendorOk

`func (o *BaseAccountVerification) GetVendorOk() (*string, bool)`

GetVendorOk returns a tuple with the Vendor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendor

`func (o *BaseAccountVerification) SetVendor(v string)`

SetVendor sets Vendor field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


