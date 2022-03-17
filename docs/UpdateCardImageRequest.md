# UpdateCardImageRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RejectionMemo** | Pointer to **string** |  | [optional] 
**RejectionReason** | Pointer to [**CardImageRejectionReason**](CardImageRejectionReason.md) |  | [optional] 
**Status** | [**CardImageStatus**](CardImageStatus.md) |  | 

## Methods

### NewUpdateCardImageRequest

`func NewUpdateCardImageRequest(status CardImageStatus, ) *UpdateCardImageRequest`

NewUpdateCardImageRequest instantiates a new UpdateCardImageRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateCardImageRequestWithDefaults

`func NewUpdateCardImageRequestWithDefaults() *UpdateCardImageRequest`

NewUpdateCardImageRequestWithDefaults instantiates a new UpdateCardImageRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRejectionMemo

`func (o *UpdateCardImageRequest) GetRejectionMemo() string`

GetRejectionMemo returns the RejectionMemo field if non-nil, zero value otherwise.

### GetRejectionMemoOk

`func (o *UpdateCardImageRequest) GetRejectionMemoOk() (*string, bool)`

GetRejectionMemoOk returns a tuple with the RejectionMemo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRejectionMemo

`func (o *UpdateCardImageRequest) SetRejectionMemo(v string)`

SetRejectionMemo sets RejectionMemo field to given value.

### HasRejectionMemo

`func (o *UpdateCardImageRequest) HasRejectionMemo() bool`

HasRejectionMemo returns a boolean if a field has been set.

### GetRejectionReason

`func (o *UpdateCardImageRequest) GetRejectionReason() CardImageRejectionReason`

GetRejectionReason returns the RejectionReason field if non-nil, zero value otherwise.

### GetRejectionReasonOk

`func (o *UpdateCardImageRequest) GetRejectionReasonOk() (*CardImageRejectionReason, bool)`

GetRejectionReasonOk returns a tuple with the RejectionReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRejectionReason

`func (o *UpdateCardImageRequest) SetRejectionReason(v CardImageRejectionReason)`

SetRejectionReason sets RejectionReason field to given value.

### HasRejectionReason

`func (o *UpdateCardImageRequest) HasRejectionReason() bool`

HasRejectionReason returns a boolean if a field has been set.

### GetStatus

`func (o *UpdateCardImageRequest) GetStatus() CardImageStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *UpdateCardImageRequest) GetStatusOk() (*CardImageStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *UpdateCardImageRequest) SetStatus(v CardImageStatus)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


