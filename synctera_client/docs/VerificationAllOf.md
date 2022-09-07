# VerificationAllOf

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
**Result** | Pointer to [**VerificationResult**](VerificationResult.md) |  | [optional] 
**VendorInfo** | Pointer to [**VendorInfo**](VendorInfo.md) |  | [optional] 
**VerificationTime** | Pointer to **time.Time** | The date and time the verification was completed. | [optional] 
**VerificationType** | Pointer to [**VerificationType1**](VerificationType1.md) |  | [optional] 

## Methods

### NewVerificationAllOf

`func NewVerificationAllOf() *VerificationAllOf`

NewVerificationAllOf instantiates a new VerificationAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVerificationAllOfWithDefaults

`func NewVerificationAllOfWithDefaults() *VerificationAllOf`

NewVerificationAllOfWithDefaults instantiates a new VerificationAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBusinessId

`func (o *VerificationAllOf) GetBusinessId() string`

GetBusinessId returns the BusinessId field if non-nil, zero value otherwise.

### GetBusinessIdOk

`func (o *VerificationAllOf) GetBusinessIdOk() (*string, bool)`

GetBusinessIdOk returns a tuple with the BusinessId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessId

`func (o *VerificationAllOf) SetBusinessId(v string)`

SetBusinessId sets BusinessId field to given value.

### HasBusinessId

`func (o *VerificationAllOf) HasBusinessId() bool`

HasBusinessId returns a boolean if a field has been set.

### GetCreationTime

`func (o *VerificationAllOf) GetCreationTime() time.Time`

GetCreationTime returns the CreationTime field if non-nil, zero value otherwise.

### GetCreationTimeOk

`func (o *VerificationAllOf) GetCreationTimeOk() (*time.Time, bool)`

GetCreationTimeOk returns a tuple with the CreationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationTime

`func (o *VerificationAllOf) SetCreationTime(v time.Time)`

SetCreationTime sets CreationTime field to given value.

### HasCreationTime

`func (o *VerificationAllOf) HasCreationTime() bool`

HasCreationTime returns a boolean if a field has been set.

### GetDetails

`func (o *VerificationAllOf) GetDetails() []Detail`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *VerificationAllOf) GetDetailsOk() (*[]Detail, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetails

`func (o *VerificationAllOf) SetDetails(v []Detail)`

SetDetails sets Details field to given value.

### HasDetails

`func (o *VerificationAllOf) HasDetails() bool`

HasDetails returns a boolean if a field has been set.

### GetId

`func (o *VerificationAllOf) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *VerificationAllOf) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *VerificationAllOf) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *VerificationAllOf) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLastUpdatedTime

`func (o *VerificationAllOf) GetLastUpdatedTime() time.Time`

GetLastUpdatedTime returns the LastUpdatedTime field if non-nil, zero value otherwise.

### GetLastUpdatedTimeOk

`func (o *VerificationAllOf) GetLastUpdatedTimeOk() (*time.Time, bool)`

GetLastUpdatedTimeOk returns a tuple with the LastUpdatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdatedTime

`func (o *VerificationAllOf) SetLastUpdatedTime(v time.Time)`

SetLastUpdatedTime sets LastUpdatedTime field to given value.

### HasLastUpdatedTime

`func (o *VerificationAllOf) HasLastUpdatedTime() bool`

HasLastUpdatedTime returns a boolean if a field has been set.

### GetMetadata

`func (o *VerificationAllOf) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *VerificationAllOf) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *VerificationAllOf) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *VerificationAllOf) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetPersonId

`func (o *VerificationAllOf) GetPersonId() string`

GetPersonId returns the PersonId field if non-nil, zero value otherwise.

### GetPersonIdOk

`func (o *VerificationAllOf) GetPersonIdOk() (*string, bool)`

GetPersonIdOk returns a tuple with the PersonId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPersonId

`func (o *VerificationAllOf) SetPersonId(v string)`

SetPersonId sets PersonId field to given value.

### HasPersonId

`func (o *VerificationAllOf) HasPersonId() bool`

HasPersonId returns a boolean if a field has been set.

### GetResult

`func (o *VerificationAllOf) GetResult() VerificationResult`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *VerificationAllOf) GetResultOk() (*VerificationResult, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *VerificationAllOf) SetResult(v VerificationResult)`

SetResult sets Result field to given value.

### HasResult

`func (o *VerificationAllOf) HasResult() bool`

HasResult returns a boolean if a field has been set.

### GetVendorInfo

`func (o *VerificationAllOf) GetVendorInfo() VendorInfo`

GetVendorInfo returns the VendorInfo field if non-nil, zero value otherwise.

### GetVendorInfoOk

`func (o *VerificationAllOf) GetVendorInfoOk() (*VendorInfo, bool)`

GetVendorInfoOk returns a tuple with the VendorInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendorInfo

`func (o *VerificationAllOf) SetVendorInfo(v VendorInfo)`

SetVendorInfo sets VendorInfo field to given value.

### HasVendorInfo

`func (o *VerificationAllOf) HasVendorInfo() bool`

HasVendorInfo returns a boolean if a field has been set.

### GetVerificationTime

`func (o *VerificationAllOf) GetVerificationTime() time.Time`

GetVerificationTime returns the VerificationTime field if non-nil, zero value otherwise.

### GetVerificationTimeOk

`func (o *VerificationAllOf) GetVerificationTimeOk() (*time.Time, bool)`

GetVerificationTimeOk returns a tuple with the VerificationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerificationTime

`func (o *VerificationAllOf) SetVerificationTime(v time.Time)`

SetVerificationTime sets VerificationTime field to given value.

### HasVerificationTime

`func (o *VerificationAllOf) HasVerificationTime() bool`

HasVerificationTime returns a boolean if a field has been set.

### GetVerificationType

`func (o *VerificationAllOf) GetVerificationType() VerificationType1`

GetVerificationType returns the VerificationType field if non-nil, zero value otherwise.

### GetVerificationTypeOk

`func (o *VerificationAllOf) GetVerificationTypeOk() (*VerificationType1, bool)`

GetVerificationTypeOk returns a tuple with the VerificationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerificationType

`func (o *VerificationAllOf) SetVerificationType(v VerificationType1)`

SetVerificationType sets VerificationType field to given value.

### HasVerificationType

`func (o *VerificationAllOf) HasVerificationType() bool`

HasVerificationType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


