# TransactionReverseRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** | The description of the transaction | [optional] 
**ForcePost** | **bool** | Whether or not the hold was forced (spending controls ignored) | 
**OffsetDescription** | Pointer to **string** | The description of the offset transaction | [optional] 
**Reason** | **string** | The reason for the reversal | 
**ReferenceId** | **NullableString** | An external ID provided by the payment network to represent this transaction. This will always be null for internal transfers. | 
**UserData** | **map[string]interface{}** | An unstructured JSON blob representing additional transaction information specific to each payment rail. | 

## Methods

### NewTransactionReverseRequest

`func NewTransactionReverseRequest(forcePost bool, reason string, referenceId NullableString, userData map[string]interface{}, ) *TransactionReverseRequest`

NewTransactionReverseRequest instantiates a new TransactionReverseRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransactionReverseRequestWithDefaults

`func NewTransactionReverseRequestWithDefaults() *TransactionReverseRequest`

NewTransactionReverseRequestWithDefaults instantiates a new TransactionReverseRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *TransactionReverseRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *TransactionReverseRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *TransactionReverseRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *TransactionReverseRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetForcePost

`func (o *TransactionReverseRequest) GetForcePost() bool`

GetForcePost returns the ForcePost field if non-nil, zero value otherwise.

### GetForcePostOk

`func (o *TransactionReverseRequest) GetForcePostOk() (*bool, bool)`

GetForcePostOk returns a tuple with the ForcePost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForcePost

`func (o *TransactionReverseRequest) SetForcePost(v bool)`

SetForcePost sets ForcePost field to given value.


### GetOffsetDescription

`func (o *TransactionReverseRequest) GetOffsetDescription() string`

GetOffsetDescription returns the OffsetDescription field if non-nil, zero value otherwise.

### GetOffsetDescriptionOk

`func (o *TransactionReverseRequest) GetOffsetDescriptionOk() (*string, bool)`

GetOffsetDescriptionOk returns a tuple with the OffsetDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffsetDescription

`func (o *TransactionReverseRequest) SetOffsetDescription(v string)`

SetOffsetDescription sets OffsetDescription field to given value.

### HasOffsetDescription

`func (o *TransactionReverseRequest) HasOffsetDescription() bool`

HasOffsetDescription returns a boolean if a field has been set.

### GetReason

`func (o *TransactionReverseRequest) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *TransactionReverseRequest) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *TransactionReverseRequest) SetReason(v string)`

SetReason sets Reason field to given value.


### GetReferenceId

`func (o *TransactionReverseRequest) GetReferenceId() string`

GetReferenceId returns the ReferenceId field if non-nil, zero value otherwise.

### GetReferenceIdOk

`func (o *TransactionReverseRequest) GetReferenceIdOk() (*string, bool)`

GetReferenceIdOk returns a tuple with the ReferenceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceId

`func (o *TransactionReverseRequest) SetReferenceId(v string)`

SetReferenceId sets ReferenceId field to given value.


### SetReferenceIdNil

`func (o *TransactionReverseRequest) SetReferenceIdNil(b bool)`

 SetReferenceIdNil sets the value for ReferenceId to be an explicit nil

### UnsetReferenceId
`func (o *TransactionReverseRequest) UnsetReferenceId()`

UnsetReferenceId ensures that no value is present for ReferenceId, not even an explicit nil
### GetUserData

`func (o *TransactionReverseRequest) GetUserData() map[string]interface{}`

GetUserData returns the UserData field if non-nil, zero value otherwise.

### GetUserDataOk

`func (o *TransactionReverseRequest) GetUserDataOk() (*map[string]interface{}, bool)`

GetUserDataOk returns a tuple with the UserData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserData

`func (o *TransactionReverseRequest) SetUserData(v map[string]interface{})`

SetUserData sets UserData field to given value.


### SetUserDataNil

`func (o *TransactionReverseRequest) SetUserDataNil(b bool)`

 SetUserDataNil sets the value for UserData to be an explicit nil

### UnsetUserData
`func (o *TransactionReverseRequest) UnsetUserData()`

UnsetUserData ensures that no value is present for UserData, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


