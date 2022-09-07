# CustomerList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Customers** | [**[]CustomerInBody**](CustomerInBody.md) | Array of Customers | 
**NextPageToken** | Pointer to **string** | If returned, use the next_page_token to query for the next page of results. Not returned if there are no more rows. | [optional] 

## Methods

### NewCustomerList

`func NewCustomerList(customers []CustomerInBody, ) *CustomerList`

NewCustomerList instantiates a new CustomerList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomerListWithDefaults

`func NewCustomerListWithDefaults() *CustomerList`

NewCustomerListWithDefaults instantiates a new CustomerList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCustomers

`func (o *CustomerList) GetCustomers() []CustomerInBody`

GetCustomers returns the Customers field if non-nil, zero value otherwise.

### GetCustomersOk

`func (o *CustomerList) GetCustomersOk() (*[]CustomerInBody, bool)`

GetCustomersOk returns a tuple with the Customers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomers

`func (o *CustomerList) SetCustomers(v []CustomerInBody)`

SetCustomers sets Customers field to given value.


### GetNextPageToken

`func (o *CustomerList) GetNextPageToken() string`

GetNextPageToken returns the NextPageToken field if non-nil, zero value otherwise.

### GetNextPageTokenOk

`func (o *CustomerList) GetNextPageTokenOk() (*string, bool)`

GetNextPageTokenOk returns a tuple with the NextPageToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextPageToken

`func (o *CustomerList) SetNextPageToken(v string)`

SetNextPageToken sets NextPageToken field to given value.

### HasNextPageToken

`func (o *CustomerList) HasNextPageToken() bool`

HasNextPageToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


