# CardEditRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CardStatus** | [**CardStatus**](CardStatus.md) |  | 
**Memo** | Pointer to **string** | Additional details about the reason for the status change | [optional] 
**Metadata** | Pointer to **map[string]string** | Additional data to include in the request structured as key-value pairs | [optional] 
**Reason** | [**CardChangeReasonCode**](CardChangeReasonCode.md) |  | 
**Shipping** | Pointer to [**Shipping**](Shipping.md) |  | [optional] 

## Methods

### NewCardEditRequest

`func NewCardEditRequest(cardStatus CardStatus, reason CardChangeReasonCode, ) *CardEditRequest`

NewCardEditRequest instantiates a new CardEditRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCardEditRequestWithDefaults

`func NewCardEditRequestWithDefaults() *CardEditRequest`

NewCardEditRequestWithDefaults instantiates a new CardEditRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCardStatus

`func (o *CardEditRequest) GetCardStatus() CardStatus`

GetCardStatus returns the CardStatus field if non-nil, zero value otherwise.

### GetCardStatusOk

`func (o *CardEditRequest) GetCardStatusOk() (*CardStatus, bool)`

GetCardStatusOk returns a tuple with the CardStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardStatus

`func (o *CardEditRequest) SetCardStatus(v CardStatus)`

SetCardStatus sets CardStatus field to given value.


### GetMemo

`func (o *CardEditRequest) GetMemo() string`

GetMemo returns the Memo field if non-nil, zero value otherwise.

### GetMemoOk

`func (o *CardEditRequest) GetMemoOk() (*string, bool)`

GetMemoOk returns a tuple with the Memo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemo

`func (o *CardEditRequest) SetMemo(v string)`

SetMemo sets Memo field to given value.

### HasMemo

`func (o *CardEditRequest) HasMemo() bool`

HasMemo returns a boolean if a field has been set.

### GetMetadata

`func (o *CardEditRequest) GetMetadata() map[string]string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *CardEditRequest) GetMetadataOk() (*map[string]string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *CardEditRequest) SetMetadata(v map[string]string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *CardEditRequest) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetReason

`func (o *CardEditRequest) GetReason() CardChangeReasonCode`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *CardEditRequest) GetReasonOk() (*CardChangeReasonCode, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *CardEditRequest) SetReason(v CardChangeReasonCode)`

SetReason sets Reason field to given value.


### GetShipping

`func (o *CardEditRequest) GetShipping() Shipping`

GetShipping returns the Shipping field if non-nil, zero value otherwise.

### GetShippingOk

`func (o *CardEditRequest) GetShippingOk() (*Shipping, bool)`

GetShippingOk returns a tuple with the Shipping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShipping

`func (o *CardEditRequest) SetShipping(v Shipping)`

SetShipping sets Shipping field to given value.

### HasShipping

`func (o *CardEditRequest) HasShipping() bool`

HasShipping returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


