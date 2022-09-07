# Verification

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BusinessId** | Pointer to **string** | Unique ID for the business. Exactly one of &#x60;business_id&#x60; or &#x60;person_id&#x60; must be set.  | [optional] 
**CreationTime** | Pointer to **time.Time** | The date and time the resource was created. | [optional] [readonly] 
**Details** | Pointer to [**[]Detail**](Detail.md) | A list of individual checks done as part of the due diligence process for the verification type.  | [optional] [readonly] 
**Id** | Pointer to **string** | Unique ID for this verification result. | [optional] [readonly] 
**LastUpdatedTime** | Pointer to **time.Time** | The date and time the resource was last updated. | [optional] [readonly] 
**Metadata** | Pointer to **map[string]interface{}** | Optional field to store additional information about the resource. Intended to be used by the integrator to store non-sensitive data.  | [optional] 
**PersonId** | Pointer to **string** | Unique ID for the person. Exactly one of &#x60;person_id&#x60; or &#x60;business_id&#x60; must be set.  | [optional] 
**Result** | [**VerificationResult**](VerificationResult.md) |  | 
**VendorInfo** | Pointer to [**VendorInfo**](VendorInfo.md) |  | [optional] 
**VerificationTime** | **time.Time** | The date and time the verification was completed. | 
**VerificationType** | [**VerificationType1**](VerificationType1.md) |  | 

## Methods

### NewVerification

`func NewVerification(result VerificationResult, verificationTime time.Time, verificationType VerificationType1, ) *Verification`

NewVerification instantiates a new Verification object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVerificationWithDefaults

`func NewVerificationWithDefaults() *Verification`

NewVerificationWithDefaults instantiates a new Verification object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBusinessId

`func (o *Verification) GetBusinessId() string`

GetBusinessId returns the BusinessId field if non-nil, zero value otherwise.

### GetBusinessIdOk

`func (o *Verification) GetBusinessIdOk() (*string, bool)`

GetBusinessIdOk returns a tuple with the BusinessId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessId

`func (o *Verification) SetBusinessId(v string)`

SetBusinessId sets BusinessId field to given value.

### HasBusinessId

`func (o *Verification) HasBusinessId() bool`

HasBusinessId returns a boolean if a field has been set.

### GetCreationTime

`func (o *Verification) GetCreationTime() time.Time`

GetCreationTime returns the CreationTime field if non-nil, zero value otherwise.

### GetCreationTimeOk

`func (o *Verification) GetCreationTimeOk() (*time.Time, bool)`

GetCreationTimeOk returns a tuple with the CreationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationTime

`func (o *Verification) SetCreationTime(v time.Time)`

SetCreationTime sets CreationTime field to given value.

### HasCreationTime

`func (o *Verification) HasCreationTime() bool`

HasCreationTime returns a boolean if a field has been set.

### GetDetails

`func (o *Verification) GetDetails() []Detail`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *Verification) GetDetailsOk() (*[]Detail, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetails

`func (o *Verification) SetDetails(v []Detail)`

SetDetails sets Details field to given value.

### HasDetails

`func (o *Verification) HasDetails() bool`

HasDetails returns a boolean if a field has been set.

### GetId

`func (o *Verification) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Verification) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Verification) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Verification) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLastUpdatedTime

`func (o *Verification) GetLastUpdatedTime() time.Time`

GetLastUpdatedTime returns the LastUpdatedTime field if non-nil, zero value otherwise.

### GetLastUpdatedTimeOk

`func (o *Verification) GetLastUpdatedTimeOk() (*time.Time, bool)`

GetLastUpdatedTimeOk returns a tuple with the LastUpdatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdatedTime

`func (o *Verification) SetLastUpdatedTime(v time.Time)`

SetLastUpdatedTime sets LastUpdatedTime field to given value.

### HasLastUpdatedTime

`func (o *Verification) HasLastUpdatedTime() bool`

HasLastUpdatedTime returns a boolean if a field has been set.

### GetMetadata

`func (o *Verification) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *Verification) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *Verification) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *Verification) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetPersonId

`func (o *Verification) GetPersonId() string`

GetPersonId returns the PersonId field if non-nil, zero value otherwise.

### GetPersonIdOk

`func (o *Verification) GetPersonIdOk() (*string, bool)`

GetPersonIdOk returns a tuple with the PersonId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPersonId

`func (o *Verification) SetPersonId(v string)`

SetPersonId sets PersonId field to given value.

### HasPersonId

`func (o *Verification) HasPersonId() bool`

HasPersonId returns a boolean if a field has been set.

### GetResult

`func (o *Verification) GetResult() VerificationResult`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *Verification) GetResultOk() (*VerificationResult, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *Verification) SetResult(v VerificationResult)`

SetResult sets Result field to given value.


### GetVendorInfo

`func (o *Verification) GetVendorInfo() VendorInfo`

GetVendorInfo returns the VendorInfo field if non-nil, zero value otherwise.

### GetVendorInfoOk

`func (o *Verification) GetVendorInfoOk() (*VendorInfo, bool)`

GetVendorInfoOk returns a tuple with the VendorInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendorInfo

`func (o *Verification) SetVendorInfo(v VendorInfo)`

SetVendorInfo sets VendorInfo field to given value.

### HasVendorInfo

`func (o *Verification) HasVendorInfo() bool`

HasVendorInfo returns a boolean if a field has been set.

### GetVerificationTime

`func (o *Verification) GetVerificationTime() time.Time`

GetVerificationTime returns the VerificationTime field if non-nil, zero value otherwise.

### GetVerificationTimeOk

`func (o *Verification) GetVerificationTimeOk() (*time.Time, bool)`

GetVerificationTimeOk returns a tuple with the VerificationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerificationTime

`func (o *Verification) SetVerificationTime(v time.Time)`

SetVerificationTime sets VerificationTime field to given value.


### GetVerificationType

`func (o *Verification) GetVerificationType() VerificationType1`

GetVerificationType returns the VerificationType field if non-nil, zero value otherwise.

### GetVerificationTypeOk

`func (o *Verification) GetVerificationTypeOk() (*VerificationType1, bool)`

GetVerificationTypeOk returns a tuple with the VerificationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerificationType

`func (o *Verification) SetVerificationType(v VerificationType1)`

SetVerificationType sets VerificationType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


