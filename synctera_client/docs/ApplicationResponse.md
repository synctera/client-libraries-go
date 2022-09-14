# ApplicationResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApplicationDetails** | **map[string]interface{}** | Details about the applicant. The exact schema is to be determined with your bank. | 
**CreationTime** | **time.Time** | Application creation timestamp in RFC3339 format | [readonly] 
**CustomerId** | **string** | Customer ID for the application | 
**Id** | **string** | Generated ID for the application | [readonly] 
**LastUpdatedTime** | **time.Time** | Timestamp of the last application modification in RFC3339 format | [readonly] 
**Status** | [**ApplicationStatus**](ApplicationStatus.md) |  | 
**Type** | [**ApplicationType1**](ApplicationType1.md) |  | 

## Methods

### NewApplicationResponse

`func NewApplicationResponse(applicationDetails map[string]interface{}, creationTime time.Time, customerId string, id string, lastUpdatedTime time.Time, status ApplicationStatus, type_ ApplicationType1, ) *ApplicationResponse`

NewApplicationResponse instantiates a new ApplicationResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplicationResponseWithDefaults

`func NewApplicationResponseWithDefaults() *ApplicationResponse`

NewApplicationResponseWithDefaults instantiates a new ApplicationResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApplicationDetails

`func (o *ApplicationResponse) GetApplicationDetails() map[string]interface{}`

GetApplicationDetails returns the ApplicationDetails field if non-nil, zero value otherwise.

### GetApplicationDetailsOk

`func (o *ApplicationResponse) GetApplicationDetailsOk() (*map[string]interface{}, bool)`

GetApplicationDetailsOk returns a tuple with the ApplicationDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationDetails

`func (o *ApplicationResponse) SetApplicationDetails(v map[string]interface{})`

SetApplicationDetails sets ApplicationDetails field to given value.


### GetCreationTime

`func (o *ApplicationResponse) GetCreationTime() time.Time`

GetCreationTime returns the CreationTime field if non-nil, zero value otherwise.

### GetCreationTimeOk

`func (o *ApplicationResponse) GetCreationTimeOk() (*time.Time, bool)`

GetCreationTimeOk returns a tuple with the CreationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationTime

`func (o *ApplicationResponse) SetCreationTime(v time.Time)`

SetCreationTime sets CreationTime field to given value.


### GetCustomerId

`func (o *ApplicationResponse) GetCustomerId() string`

GetCustomerId returns the CustomerId field if non-nil, zero value otherwise.

### GetCustomerIdOk

`func (o *ApplicationResponse) GetCustomerIdOk() (*string, bool)`

GetCustomerIdOk returns a tuple with the CustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerId

`func (o *ApplicationResponse) SetCustomerId(v string)`

SetCustomerId sets CustomerId field to given value.


### GetId

`func (o *ApplicationResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ApplicationResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ApplicationResponse) SetId(v string)`

SetId sets Id field to given value.


### GetLastUpdatedTime

`func (o *ApplicationResponse) GetLastUpdatedTime() time.Time`

GetLastUpdatedTime returns the LastUpdatedTime field if non-nil, zero value otherwise.

### GetLastUpdatedTimeOk

`func (o *ApplicationResponse) GetLastUpdatedTimeOk() (*time.Time, bool)`

GetLastUpdatedTimeOk returns a tuple with the LastUpdatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdatedTime

`func (o *ApplicationResponse) SetLastUpdatedTime(v time.Time)`

SetLastUpdatedTime sets LastUpdatedTime field to given value.


### GetStatus

`func (o *ApplicationResponse) GetStatus() ApplicationStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ApplicationResponse) GetStatusOk() (*ApplicationStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ApplicationResponse) SetStatus(v ApplicationStatus)`

SetStatus sets Status field to given value.


### GetType

`func (o *ApplicationResponse) GetType() ApplicationType1`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ApplicationResponse) GetTypeOk() (*ApplicationType1, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ApplicationResponse) SetType(v ApplicationType1)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


