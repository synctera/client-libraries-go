# CardStatusObject

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CardFulfillmentStatus** | Pointer to [**CardFulfillmentStatus**](CardFulfillmentStatus.md) |  | [optional] 
**CardStatus** | [**CardStatus**](CardStatus.md) |  | 
**Carrier** | Pointer to **string** | The carrier with whom the card is shipped | [optional] [readonly] 
**Memo** | Pointer to **string** | Additional details about the reason for the status change | [optional] 
**ShippingStatus** | Pointer to **string** | The status of indicating the shipping status of the card | [optional] [readonly] 
**StatusReason** | Pointer to [**CardStatusReasonCode**](CardStatusReasonCode.md) |  | [optional] 
**TrackingNumber** | Pointer to **string** | The tracking number | [optional] [readonly] 

## Methods

### NewCardStatusObject

`func NewCardStatusObject(cardStatus CardStatus, ) *CardStatusObject`

NewCardStatusObject instantiates a new CardStatusObject object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCardStatusObjectWithDefaults

`func NewCardStatusObjectWithDefaults() *CardStatusObject`

NewCardStatusObjectWithDefaults instantiates a new CardStatusObject object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCardFulfillmentStatus

`func (o *CardStatusObject) GetCardFulfillmentStatus() CardFulfillmentStatus`

GetCardFulfillmentStatus returns the CardFulfillmentStatus field if non-nil, zero value otherwise.

### GetCardFulfillmentStatusOk

`func (o *CardStatusObject) GetCardFulfillmentStatusOk() (*CardFulfillmentStatus, bool)`

GetCardFulfillmentStatusOk returns a tuple with the CardFulfillmentStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardFulfillmentStatus

`func (o *CardStatusObject) SetCardFulfillmentStatus(v CardFulfillmentStatus)`

SetCardFulfillmentStatus sets CardFulfillmentStatus field to given value.

### HasCardFulfillmentStatus

`func (o *CardStatusObject) HasCardFulfillmentStatus() bool`

HasCardFulfillmentStatus returns a boolean if a field has been set.

### GetCardStatus

`func (o *CardStatusObject) GetCardStatus() CardStatus`

GetCardStatus returns the CardStatus field if non-nil, zero value otherwise.

### GetCardStatusOk

`func (o *CardStatusObject) GetCardStatusOk() (*CardStatus, bool)`

GetCardStatusOk returns a tuple with the CardStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardStatus

`func (o *CardStatusObject) SetCardStatus(v CardStatus)`

SetCardStatus sets CardStatus field to given value.


### GetCarrier

`func (o *CardStatusObject) GetCarrier() string`

GetCarrier returns the Carrier field if non-nil, zero value otherwise.

### GetCarrierOk

`func (o *CardStatusObject) GetCarrierOk() (*string, bool)`

GetCarrierOk returns a tuple with the Carrier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCarrier

`func (o *CardStatusObject) SetCarrier(v string)`

SetCarrier sets Carrier field to given value.

### HasCarrier

`func (o *CardStatusObject) HasCarrier() bool`

HasCarrier returns a boolean if a field has been set.

### GetMemo

`func (o *CardStatusObject) GetMemo() string`

GetMemo returns the Memo field if non-nil, zero value otherwise.

### GetMemoOk

`func (o *CardStatusObject) GetMemoOk() (*string, bool)`

GetMemoOk returns a tuple with the Memo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemo

`func (o *CardStatusObject) SetMemo(v string)`

SetMemo sets Memo field to given value.

### HasMemo

`func (o *CardStatusObject) HasMemo() bool`

HasMemo returns a boolean if a field has been set.

### GetShippingStatus

`func (o *CardStatusObject) GetShippingStatus() string`

GetShippingStatus returns the ShippingStatus field if non-nil, zero value otherwise.

### GetShippingStatusOk

`func (o *CardStatusObject) GetShippingStatusOk() (*string, bool)`

GetShippingStatusOk returns a tuple with the ShippingStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShippingStatus

`func (o *CardStatusObject) SetShippingStatus(v string)`

SetShippingStatus sets ShippingStatus field to given value.

### HasShippingStatus

`func (o *CardStatusObject) HasShippingStatus() bool`

HasShippingStatus returns a boolean if a field has been set.

### GetStatusReason

`func (o *CardStatusObject) GetStatusReason() CardStatusReasonCode`

GetStatusReason returns the StatusReason field if non-nil, zero value otherwise.

### GetStatusReasonOk

`func (o *CardStatusObject) GetStatusReasonOk() (*CardStatusReasonCode, bool)`

GetStatusReasonOk returns a tuple with the StatusReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusReason

`func (o *CardStatusObject) SetStatusReason(v CardStatusReasonCode)`

SetStatusReason sets StatusReason field to given value.

### HasStatusReason

`func (o *CardStatusObject) HasStatusReason() bool`

HasStatusReason returns a boolean if a field has been set.

### GetTrackingNumber

`func (o *CardStatusObject) GetTrackingNumber() string`

GetTrackingNumber returns the TrackingNumber field if non-nil, zero value otherwise.

### GetTrackingNumberOk

`func (o *CardStatusObject) GetTrackingNumberOk() (*string, bool)`

GetTrackingNumberOk returns a tuple with the TrackingNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrackingNumber

`func (o *CardStatusObject) SetTrackingNumber(v string)`

SetTrackingNumber sets TrackingNumber field to given value.

### HasTrackingNumber

`func (o *CardStatusObject) HasTrackingNumber() bool`

HasTrackingNumber returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


