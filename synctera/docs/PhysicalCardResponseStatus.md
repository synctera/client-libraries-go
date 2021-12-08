# PhysicalCardResponseStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CardStatus** | [**CardStatus**](CardStatus.md) |  | 
**Memo** | Pointer to **string** | Additional details about the reason for the status change | [optional] 
**StatusReason** | [**CardStatusReasonCode**](CardStatusReasonCode.md) |  | 
**CardFulfillmentStatus** | [**CardFulfillmentStatus**](CardFulfillmentStatus.md) |  | 
**TrackingNumber** | Pointer to **string** | The tracking number | [optional] [readonly] 

## Methods

### NewPhysicalCardResponseStatus

`func NewPhysicalCardResponseStatus(cardStatus CardStatus, statusReason CardStatusReasonCode, cardFulfillmentStatus CardFulfillmentStatus, ) *PhysicalCardResponseStatus`

NewPhysicalCardResponseStatus instantiates a new PhysicalCardResponseStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPhysicalCardResponseStatusWithDefaults

`func NewPhysicalCardResponseStatusWithDefaults() *PhysicalCardResponseStatus`

NewPhysicalCardResponseStatusWithDefaults instantiates a new PhysicalCardResponseStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCardStatus

`func (o *PhysicalCardResponseStatus) GetCardStatus() CardStatus`

GetCardStatus returns the CardStatus field if non-nil, zero value otherwise.

### GetCardStatusOk

`func (o *PhysicalCardResponseStatus) GetCardStatusOk() (*CardStatus, bool)`

GetCardStatusOk returns a tuple with the CardStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardStatus

`func (o *PhysicalCardResponseStatus) SetCardStatus(v CardStatus)`

SetCardStatus sets CardStatus field to given value.


### GetMemo

`func (o *PhysicalCardResponseStatus) GetMemo() string`

GetMemo returns the Memo field if non-nil, zero value otherwise.

### GetMemoOk

`func (o *PhysicalCardResponseStatus) GetMemoOk() (*string, bool)`

GetMemoOk returns a tuple with the Memo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemo

`func (o *PhysicalCardResponseStatus) SetMemo(v string)`

SetMemo sets Memo field to given value.

### HasMemo

`func (o *PhysicalCardResponseStatus) HasMemo() bool`

HasMemo returns a boolean if a field has been set.

### GetStatusReason

`func (o *PhysicalCardResponseStatus) GetStatusReason() CardStatusReasonCode`

GetStatusReason returns the StatusReason field if non-nil, zero value otherwise.

### GetStatusReasonOk

`func (o *PhysicalCardResponseStatus) GetStatusReasonOk() (*CardStatusReasonCode, bool)`

GetStatusReasonOk returns a tuple with the StatusReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusReason

`func (o *PhysicalCardResponseStatus) SetStatusReason(v CardStatusReasonCode)`

SetStatusReason sets StatusReason field to given value.


### GetCardFulfillmentStatus

`func (o *PhysicalCardResponseStatus) GetCardFulfillmentStatus() CardFulfillmentStatus`

GetCardFulfillmentStatus returns the CardFulfillmentStatus field if non-nil, zero value otherwise.

### GetCardFulfillmentStatusOk

`func (o *PhysicalCardResponseStatus) GetCardFulfillmentStatusOk() (*CardFulfillmentStatus, bool)`

GetCardFulfillmentStatusOk returns a tuple with the CardFulfillmentStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardFulfillmentStatus

`func (o *PhysicalCardResponseStatus) SetCardFulfillmentStatus(v CardFulfillmentStatus)`

SetCardFulfillmentStatus sets CardFulfillmentStatus field to given value.


### GetTrackingNumber

`func (o *PhysicalCardResponseStatus) GetTrackingNumber() string`

GetTrackingNumber returns the TrackingNumber field if non-nil, zero value otherwise.

### GetTrackingNumberOk

`func (o *PhysicalCardResponseStatus) GetTrackingNumberOk() (*string, bool)`

GetTrackingNumberOk returns a tuple with the TrackingNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrackingNumber

`func (o *PhysicalCardResponseStatus) SetTrackingNumber(v string)`

SetTrackingNumber sets TrackingNumber field to given value.

### HasTrackingNumber

`func (o *PhysicalCardResponseStatus) HasTrackingNumber() bool`

HasTrackingNumber returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


