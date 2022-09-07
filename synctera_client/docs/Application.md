# Application

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApplicationDetails** | **map[string]interface{}** | Details about the applicant. The exact schema is to be determined with your bank. | 
**CustomerId** | **string** | Customer ID for the application | 
**Status** | Pointer to [**ApplicationStatus**](ApplicationStatus.md) |  | [optional] 
**Type** | [**ApplicationType**](ApplicationType.md) |  | 

## Methods

### NewApplication

`func NewApplication(applicationDetails map[string]interface{}, customerId string, type_ ApplicationType, ) *Application`

NewApplication instantiates a new Application object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplicationWithDefaults

`func NewApplicationWithDefaults() *Application`

NewApplicationWithDefaults instantiates a new Application object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApplicationDetails

`func (o *Application) GetApplicationDetails() map[string]interface{}`

GetApplicationDetails returns the ApplicationDetails field if non-nil, zero value otherwise.

### GetApplicationDetailsOk

`func (o *Application) GetApplicationDetailsOk() (*map[string]interface{}, bool)`

GetApplicationDetailsOk returns a tuple with the ApplicationDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationDetails

`func (o *Application) SetApplicationDetails(v map[string]interface{})`

SetApplicationDetails sets ApplicationDetails field to given value.


### GetCustomerId

`func (o *Application) GetCustomerId() string`

GetCustomerId returns the CustomerId field if non-nil, zero value otherwise.

### GetCustomerIdOk

`func (o *Application) GetCustomerIdOk() (*string, bool)`

GetCustomerIdOk returns a tuple with the CustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerId

`func (o *Application) SetCustomerId(v string)`

SetCustomerId sets CustomerId field to given value.


### GetStatus

`func (o *Application) GetStatus() ApplicationStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Application) GetStatusOk() (*ApplicationStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Application) SetStatus(v ApplicationStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *Application) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetType

`func (o *Application) GetType() ApplicationType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Application) GetTypeOk() (*ApplicationType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Application) SetType(v ApplicationType)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


