# CardChange

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ChangeType** | [**ChangeType**](ChangeType.md) |  | 
**Channel** | [**ChangeChannel**](ChangeChannel.md) |  | 
**Id** | **string** | Unique token | [readonly] 
**Memo** | Pointer to **string** | Additional details about the reason for the status change | [optional] 
**Reason** | Pointer to [**CardStatusReasonCode**](CardStatusReasonCode.md) |  | [optional] 
**StateNew** | [**CardChangeState**](CardChangeState.md) |  | 
**StateOld** | [**CardChangeState**](CardChangeState.md) |  | 
**UpdatedAt** | **time.Time** | Date of change | [readonly] 
**UpdatedBy** | **string** | ID of user who initiated the change, if done via Synctera Admin System | 

## Methods

### NewCardChange

`func NewCardChange(changeType ChangeType, channel ChangeChannel, id string, stateNew CardChangeState, stateOld CardChangeState, updatedAt time.Time, updatedBy string, ) *CardChange`

NewCardChange instantiates a new CardChange object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCardChangeWithDefaults

`func NewCardChangeWithDefaults() *CardChange`

NewCardChangeWithDefaults instantiates a new CardChange object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChangeType

`func (o *CardChange) GetChangeType() ChangeType`

GetChangeType returns the ChangeType field if non-nil, zero value otherwise.

### GetChangeTypeOk

`func (o *CardChange) GetChangeTypeOk() (*ChangeType, bool)`

GetChangeTypeOk returns a tuple with the ChangeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChangeType

`func (o *CardChange) SetChangeType(v ChangeType)`

SetChangeType sets ChangeType field to given value.


### GetChannel

`func (o *CardChange) GetChannel() ChangeChannel`

GetChannel returns the Channel field if non-nil, zero value otherwise.

### GetChannelOk

`func (o *CardChange) GetChannelOk() (*ChangeChannel, bool)`

GetChannelOk returns a tuple with the Channel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannel

`func (o *CardChange) SetChannel(v ChangeChannel)`

SetChannel sets Channel field to given value.


### GetId

`func (o *CardChange) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CardChange) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CardChange) SetId(v string)`

SetId sets Id field to given value.


### GetMemo

`func (o *CardChange) GetMemo() string`

GetMemo returns the Memo field if non-nil, zero value otherwise.

### GetMemoOk

`func (o *CardChange) GetMemoOk() (*string, bool)`

GetMemoOk returns a tuple with the Memo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemo

`func (o *CardChange) SetMemo(v string)`

SetMemo sets Memo field to given value.

### HasMemo

`func (o *CardChange) HasMemo() bool`

HasMemo returns a boolean if a field has been set.

### GetReason

`func (o *CardChange) GetReason() CardStatusReasonCode`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *CardChange) GetReasonOk() (*CardStatusReasonCode, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *CardChange) SetReason(v CardStatusReasonCode)`

SetReason sets Reason field to given value.

### HasReason

`func (o *CardChange) HasReason() bool`

HasReason returns a boolean if a field has been set.

### GetStateNew

`func (o *CardChange) GetStateNew() CardChangeState`

GetStateNew returns the StateNew field if non-nil, zero value otherwise.

### GetStateNewOk

`func (o *CardChange) GetStateNewOk() (*CardChangeState, bool)`

GetStateNewOk returns a tuple with the StateNew field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStateNew

`func (o *CardChange) SetStateNew(v CardChangeState)`

SetStateNew sets StateNew field to given value.


### GetStateOld

`func (o *CardChange) GetStateOld() CardChangeState`

GetStateOld returns the StateOld field if non-nil, zero value otherwise.

### GetStateOldOk

`func (o *CardChange) GetStateOldOk() (*CardChangeState, bool)`

GetStateOldOk returns a tuple with the StateOld field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStateOld

`func (o *CardChange) SetStateOld(v CardChangeState)`

SetStateOld sets StateOld field to given value.


### GetUpdatedAt

`func (o *CardChange) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *CardChange) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *CardChange) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetUpdatedBy

`func (o *CardChange) GetUpdatedBy() string`

GetUpdatedBy returns the UpdatedBy field if non-nil, zero value otherwise.

### GetUpdatedByOk

`func (o *CardChange) GetUpdatedByOk() (*string, bool)`

GetUpdatedByOk returns a tuple with the UpdatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedBy

`func (o *CardChange) SetUpdatedBy(v string)`

SetUpdatedBy sets UpdatedBy field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


