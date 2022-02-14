# CreateCardImageRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CustomerId** | **string** | The unique identifier of a customer | 
**ImageData** | Pointer to **string** |  | [optional] 

## Methods

### NewCreateCardImageRequest

`func NewCreateCardImageRequest(customerId string, ) *CreateCardImageRequest`

NewCreateCardImageRequest instantiates a new CreateCardImageRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateCardImageRequestWithDefaults

`func NewCreateCardImageRequestWithDefaults() *CreateCardImageRequest`

NewCreateCardImageRequestWithDefaults instantiates a new CreateCardImageRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

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


### GetImageData

`func (o *CreateCardImageRequest) GetImageData() string`

GetImageData returns the ImageData field if non-nil, zero value otherwise.

### GetImageDataOk

`func (o *CreateCardImageRequest) GetImageDataOk() (*string, bool)`

GetImageDataOk returns a tuple with the ImageData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageData

`func (o *CreateCardImageRequest) SetImageData(v string)`

SetImageData sets ImageData field to given value.

### HasImageData

`func (o *CreateCardImageRequest) HasImageData() bool`

HasImageData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


