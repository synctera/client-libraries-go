# VirtualCardResponseStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CardStatus** | [**CardStatus**](CardStatus.md) |  | 
**Memo** | Pointer to **string** | Additional details about the reason for the status change | [optional] 
**StatusReason** | [**CardStatusReasonCode**](CardStatusReasonCode.md) |  | 

## Methods

### NewVirtualCardResponseStatus

`func NewVirtualCardResponseStatus(cardStatus CardStatus, statusReason CardStatusReasonCode, ) *VirtualCardResponseStatus`

NewVirtualCardResponseStatus instantiates a new VirtualCardResponseStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVirtualCardResponseStatusWithDefaults

`func NewVirtualCardResponseStatusWithDefaults() *VirtualCardResponseStatus`

NewVirtualCardResponseStatusWithDefaults instantiates a new VirtualCardResponseStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCardStatus

`func (o *VirtualCardResponseStatus) GetCardStatus() CardStatus`

GetCardStatus returns the CardStatus field if non-nil, zero value otherwise.

### GetCardStatusOk

`func (o *VirtualCardResponseStatus) GetCardStatusOk() (*CardStatus, bool)`

GetCardStatusOk returns a tuple with the CardStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardStatus

`func (o *VirtualCardResponseStatus) SetCardStatus(v CardStatus)`

SetCardStatus sets CardStatus field to given value.


### GetMemo

`func (o *VirtualCardResponseStatus) GetMemo() string`

GetMemo returns the Memo field if non-nil, zero value otherwise.

### GetMemoOk

`func (o *VirtualCardResponseStatus) GetMemoOk() (*string, bool)`

GetMemoOk returns a tuple with the Memo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemo

`func (o *VirtualCardResponseStatus) SetMemo(v string)`

SetMemo sets Memo field to given value.

### HasMemo

`func (o *VirtualCardResponseStatus) HasMemo() bool`

HasMemo returns a boolean if a field has been set.

### GetStatusReason

`func (o *VirtualCardResponseStatus) GetStatusReason() CardStatusReasonCode`

GetStatusReason returns the StatusReason field if non-nil, zero value otherwise.

### GetStatusReasonOk

`func (o *VirtualCardResponseStatus) GetStatusReasonOk() (*CardStatusReasonCode, bool)`

GetStatusReasonOk returns a tuple with the StatusReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusReason

`func (o *VirtualCardResponseStatus) SetStatusReason(v CardStatusReasonCode)`

SetStatusReason sets StatusReason field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


