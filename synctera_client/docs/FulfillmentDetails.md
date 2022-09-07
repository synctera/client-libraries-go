# FulfillmentDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ShipDate** | Pointer to **string** | The date that the card was shipped as reported by the card fulfillment provider | [optional] [readonly] 
**ShippingMethod** | Pointer to **string** | The specific shipping method as reported by the card fulfillment provider | [optional] [readonly] 
**TrackingNumber** | Pointer to **string** | The shipment tracking number | [optional] [readonly] 

## Methods

### NewFulfillmentDetails

`func NewFulfillmentDetails() *FulfillmentDetails`

NewFulfillmentDetails instantiates a new FulfillmentDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFulfillmentDetailsWithDefaults

`func NewFulfillmentDetailsWithDefaults() *FulfillmentDetails`

NewFulfillmentDetailsWithDefaults instantiates a new FulfillmentDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetShipDate

`func (o *FulfillmentDetails) GetShipDate() string`

GetShipDate returns the ShipDate field if non-nil, zero value otherwise.

### GetShipDateOk

`func (o *FulfillmentDetails) GetShipDateOk() (*string, bool)`

GetShipDateOk returns a tuple with the ShipDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShipDate

`func (o *FulfillmentDetails) SetShipDate(v string)`

SetShipDate sets ShipDate field to given value.

### HasShipDate

`func (o *FulfillmentDetails) HasShipDate() bool`

HasShipDate returns a boolean if a field has been set.

### GetShippingMethod

`func (o *FulfillmentDetails) GetShippingMethod() string`

GetShippingMethod returns the ShippingMethod field if non-nil, zero value otherwise.

### GetShippingMethodOk

`func (o *FulfillmentDetails) GetShippingMethodOk() (*string, bool)`

GetShippingMethodOk returns a tuple with the ShippingMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShippingMethod

`func (o *FulfillmentDetails) SetShippingMethod(v string)`

SetShippingMethod sets ShippingMethod field to given value.

### HasShippingMethod

`func (o *FulfillmentDetails) HasShippingMethod() bool`

HasShippingMethod returns a boolean if a field has been set.

### GetTrackingNumber

`func (o *FulfillmentDetails) GetTrackingNumber() string`

GetTrackingNumber returns the TrackingNumber field if non-nil, zero value otherwise.

### GetTrackingNumberOk

`func (o *FulfillmentDetails) GetTrackingNumberOk() (*string, bool)`

GetTrackingNumberOk returns a tuple with the TrackingNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrackingNumber

`func (o *FulfillmentDetails) SetTrackingNumber(v string)`

SetTrackingNumber sets TrackingNumber field to given value.

### HasTrackingNumber

`func (o *FulfillmentDetails) HasTrackingNumber() bool`

HasTrackingNumber returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


