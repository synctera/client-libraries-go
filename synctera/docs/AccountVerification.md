# AccountVerification

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreationTime** | Pointer to **time.Time** | The time at which verification was first completed. | [optional] 
**LastUpdatedTime** | Pointer to **time.Time** | The time at which verification was last updated. | [optional] 
**Status** | **string** | The status of verification | 
**Vendor** | **string** | The vendor used for verifying the account | 

## Methods

### NewAccountVerification

`func NewAccountVerification(status string, vendor string, ) *AccountVerification`

NewAccountVerification instantiates a new AccountVerification object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountVerificationWithDefaults

`func NewAccountVerificationWithDefaults() *AccountVerification`

NewAccountVerificationWithDefaults instantiates a new AccountVerification object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreationTime

`func (o *AccountVerification) GetCreationTime() time.Time`

GetCreationTime returns the CreationTime field if non-nil, zero value otherwise.

### GetCreationTimeOk

`func (o *AccountVerification) GetCreationTimeOk() (*time.Time, bool)`

GetCreationTimeOk returns a tuple with the CreationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationTime

`func (o *AccountVerification) SetCreationTime(v time.Time)`

SetCreationTime sets CreationTime field to given value.

### HasCreationTime

`func (o *AccountVerification) HasCreationTime() bool`

HasCreationTime returns a boolean if a field has been set.

### GetLastUpdatedTime

`func (o *AccountVerification) GetLastUpdatedTime() time.Time`

GetLastUpdatedTime returns the LastUpdatedTime field if non-nil, zero value otherwise.

### GetLastUpdatedTimeOk

`func (o *AccountVerification) GetLastUpdatedTimeOk() (*time.Time, bool)`

GetLastUpdatedTimeOk returns a tuple with the LastUpdatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdatedTime

`func (o *AccountVerification) SetLastUpdatedTime(v time.Time)`

SetLastUpdatedTime sets LastUpdatedTime field to given value.

### HasLastUpdatedTime

`func (o *AccountVerification) HasLastUpdatedTime() bool`

HasLastUpdatedTime returns a boolean if a field has been set.

### GetStatus

`func (o *AccountVerification) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *AccountVerification) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *AccountVerification) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetVendor

`func (o *AccountVerification) GetVendor() string`

GetVendor returns the Vendor field if non-nil, zero value otherwise.

### GetVendorOk

`func (o *AccountVerification) GetVendorOk() (*string, bool)`

GetVendorOk returns a tuple with the Vendor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendor

`func (o *AccountVerification) SetVendor(v string)`

SetVendor sets Vendor field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


