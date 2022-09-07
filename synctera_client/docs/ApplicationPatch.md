# ApplicationPatch

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApplicationDetails** | Pointer to **map[string]interface{}** | Details about the applicant. The exact schema is to be determined with your bank. | [optional] 
**Status** | Pointer to [**ApplicationStatus**](ApplicationStatus.md) |  | [optional] 

## Methods

### NewApplicationPatch

`func NewApplicationPatch() *ApplicationPatch`

NewApplicationPatch instantiates a new ApplicationPatch object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplicationPatchWithDefaults

`func NewApplicationPatchWithDefaults() *ApplicationPatch`

NewApplicationPatchWithDefaults instantiates a new ApplicationPatch object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApplicationDetails

`func (o *ApplicationPatch) GetApplicationDetails() map[string]interface{}`

GetApplicationDetails returns the ApplicationDetails field if non-nil, zero value otherwise.

### GetApplicationDetailsOk

`func (o *ApplicationPatch) GetApplicationDetailsOk() (*map[string]interface{}, bool)`

GetApplicationDetailsOk returns a tuple with the ApplicationDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationDetails

`func (o *ApplicationPatch) SetApplicationDetails(v map[string]interface{})`

SetApplicationDetails sets ApplicationDetails field to given value.

### HasApplicationDetails

`func (o *ApplicationPatch) HasApplicationDetails() bool`

HasApplicationDetails returns a boolean if a field has been set.

### GetStatus

`func (o *ApplicationPatch) GetStatus() ApplicationStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ApplicationPatch) GetStatusOk() (*ApplicationStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ApplicationPatch) SetStatus(v ApplicationStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ApplicationPatch) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


