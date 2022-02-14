# CardImageDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CustomerId** | **string** | The unique identifier of a customer | 
**Id** | **string** | The unique identifier of a card image | 
**Status** | [**CardImageStatus**](CardImageStatus.md) |  | 

## Methods

### NewCardImageDetails

`func NewCardImageDetails(customerId string, id string, status CardImageStatus, ) *CardImageDetails`

NewCardImageDetails instantiates a new CardImageDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCardImageDetailsWithDefaults

`func NewCardImageDetailsWithDefaults() *CardImageDetails`

NewCardImageDetailsWithDefaults instantiates a new CardImageDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCustomerId

`func (o *CardImageDetails) GetCustomerId() string`

GetCustomerId returns the CustomerId field if non-nil, zero value otherwise.

### GetCustomerIdOk

`func (o *CardImageDetails) GetCustomerIdOk() (*string, bool)`

GetCustomerIdOk returns a tuple with the CustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerId

`func (o *CardImageDetails) SetCustomerId(v string)`

SetCustomerId sets CustomerId field to given value.


### GetId

`func (o *CardImageDetails) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CardImageDetails) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CardImageDetails) SetId(v string)`

SetId sets Id field to given value.


### GetStatus

`func (o *CardImageDetails) GetStatus() CardImageStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CardImageDetails) GetStatusOk() (*CardImageStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CardImageDetails) SetStatus(v CardImageStatus)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


