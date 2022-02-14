# CustomerAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Dob** | [**oapi.Date**](oapi.Date.md) | Customer&#39;s date of birth in RFC 3339 full-date format (YYYY-MM-DD) | 
**FirstName** | **string** | Customer&#39;s first name | 
**LastName** | **string** | Customer&#39;s last name | 
**Status** | **string** | Customer&#39;s status | 

## Methods

### NewCustomerAllOf

`func NewCustomerAllOf(dob oapi.Date, firstName string, lastName string, status string, ) *CustomerAllOf`

NewCustomerAllOf instantiates a new CustomerAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomerAllOfWithDefaults

`func NewCustomerAllOfWithDefaults() *CustomerAllOf`

NewCustomerAllOfWithDefaults instantiates a new CustomerAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDob

`func (o *CustomerAllOf) GetDob() oapi.Date`

GetDob returns the Dob field if non-nil, zero value otherwise.

### GetDobOk

`func (o *CustomerAllOf) GetDobOk() (*oapi.Date, bool)`

GetDobOk returns a tuple with the Dob field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDob

`func (o *CustomerAllOf) SetDob(v oapi.Date)`

SetDob sets Dob field to given value.


### GetFirstName

`func (o *CustomerAllOf) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *CustomerAllOf) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *CustomerAllOf) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.


### GetLastName

`func (o *CustomerAllOf) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *CustomerAllOf) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *CustomerAllOf) SetLastName(v string)`

SetLastName sets LastName field to given value.


### GetStatus

`func (o *CustomerAllOf) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CustomerAllOf) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CustomerAllOf) SetStatus(v string)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


