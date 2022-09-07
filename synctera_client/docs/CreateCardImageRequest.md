# CreateCardImageRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CardProductId** | **string** | The unique identifier of a cards product | 
**CustomerId** | **string** | The unique identifier of a customer | 

## Methods

### NewCreateCardImageRequest

`func NewCreateCardImageRequest(cardProductId string, customerId string, ) *CreateCardImageRequest`

NewCreateCardImageRequest instantiates a new CreateCardImageRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateCardImageRequestWithDefaults

`func NewCreateCardImageRequestWithDefaults() *CreateCardImageRequest`

NewCreateCardImageRequestWithDefaults instantiates a new CreateCardImageRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCardProductId

`func (o *CreateCardImageRequest) GetCardProductId() string`

GetCardProductId returns the CardProductId field if non-nil, zero value otherwise.

### GetCardProductIdOk

`func (o *CreateCardImageRequest) GetCardProductIdOk() (*string, bool)`

GetCardProductIdOk returns a tuple with the CardProductId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardProductId

`func (o *CreateCardImageRequest) SetCardProductId(v string)`

SetCardProductId sets CardProductId field to given value.


### GetCustomerId

`func (o *CreateCardImageRequest) GetCustomerId() string`

GetCustomerId returns the CustomerId field if non-nil, zero value otherwise.

### GetCustomerIdOk

`func (o *CreateCardImageRequest) GetCustomerIdOk() (*string, bool)`

GetCustomerIdOk returns a tuple with the CustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerId

`func (o *CreateCardImageRequest) SetCustomerId(v string)`

SetCustomerId sets CustomerId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


