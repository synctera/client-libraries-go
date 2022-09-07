# ApplicationListAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Applications** | [**[]ApplicationResponse**](ApplicationResponse.md) | Array of credit applications. | 

## Methods

### NewApplicationListAllOf

`func NewApplicationListAllOf(applications []ApplicationResponse, ) *ApplicationListAllOf`

NewApplicationListAllOf instantiates a new ApplicationListAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplicationListAllOfWithDefaults

`func NewApplicationListAllOfWithDefaults() *ApplicationListAllOf`

NewApplicationListAllOfWithDefaults instantiates a new ApplicationListAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApplications

`func (o *ApplicationListAllOf) GetApplications() []ApplicationResponse`

GetApplications returns the Applications field if non-nil, zero value otherwise.

### GetApplicationsOk

`func (o *ApplicationListAllOf) GetApplicationsOk() (*[]ApplicationResponse, bool)`

GetApplicationsOk returns a tuple with the Applications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplications

`func (o *ApplicationListAllOf) SetApplications(v []ApplicationResponse)`

SetApplications sets Applications field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


