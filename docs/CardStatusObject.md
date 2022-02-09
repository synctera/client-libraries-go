# CardStatusObject

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CardStatus** | [**CardStatus**](CardStatus.md) |  | 
**StatusReason** | Pointer to [**CardStatusReasonCode**](CardStatusReasonCode.md) |  | [optional] 
**Memo** | Pointer to **string** | Additional details about the reason for the status change | [optional] 

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


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


