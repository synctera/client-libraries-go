# OriginalCreditRequestModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | **int32** | The amount of the transaction in the smallest whole denomination of the applicable currency (eg. For USD use cents) | 
**CardAcceptor** | Pointer to [**CardAcceptorModel**](CardAcceptorModel.md) |  | [optional] 
**CardId** | **string** |  | 
**Mid** | **string** |  | 
**ScreeningScore** | Pointer to **string** |  | [optional] 
**SenderData** | Pointer to [**OriginalCreditSenderData**](OriginalCreditSenderData.md) |  | [optional] 
**TransactionPurpose** | Pointer to **string** |  | [optional] 
**Type** | **string** |  | 

## Methods

### NewOriginalCreditRequestModel

`func NewOriginalCreditRequestModel(amount int32, cardId string, mid string, type_ string, ) *OriginalCreditRequestModel`

NewOriginalCreditRequestModel instantiates a new OriginalCreditRequestModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOriginalCreditRequestModelWithDefaults

`func NewOriginalCreditRequestModelWithDefaults() *OriginalCreditRequestModel`

NewOriginalCreditRequestModelWithDefaults instantiates a new OriginalCreditRequestModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmount

`func (o *OriginalCreditRequestModel) GetAmount() int32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *OriginalCreditRequestModel) GetAmountOk() (*int32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *OriginalCreditRequestModel) SetAmount(v int32)`

SetAmount sets Amount field to given value.


### GetCardAcceptor

`func (o *OriginalCreditRequestModel) GetCardAcceptor() CardAcceptorModel`

GetCardAcceptor returns the CardAcceptor field if non-nil, zero value otherwise.

### GetCardAcceptorOk

`func (o *OriginalCreditRequestModel) GetCardAcceptorOk() (*CardAcceptorModel, bool)`

GetCardAcceptorOk returns a tuple with the CardAcceptor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardAcceptor

`func (o *OriginalCreditRequestModel) SetCardAcceptor(v CardAcceptorModel)`

SetCardAcceptor sets CardAcceptor field to given value.

### HasCardAcceptor

`func (o *OriginalCreditRequestModel) HasCardAcceptor() bool`

HasCardAcceptor returns a boolean if a field has been set.

### GetCardId

`func (o *OriginalCreditRequestModel) GetCardId() string`

GetCardId returns the CardId field if non-nil, zero value otherwise.

### GetCardIdOk

`func (o *OriginalCreditRequestModel) GetCardIdOk() (*string, bool)`

GetCardIdOk returns a tuple with the CardId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardId

`func (o *OriginalCreditRequestModel) SetCardId(v string)`

SetCardId sets CardId field to given value.


### GetMid

`func (o *OriginalCreditRequestModel) GetMid() string`

GetMid returns the Mid field if non-nil, zero value otherwise.

### GetMidOk

`func (o *OriginalCreditRequestModel) GetMidOk() (*string, bool)`

GetMidOk returns a tuple with the Mid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMid

`func (o *OriginalCreditRequestModel) SetMid(v string)`

SetMid sets Mid field to given value.


### GetScreeningScore

`func (o *OriginalCreditRequestModel) GetScreeningScore() string`

GetScreeningScore returns the ScreeningScore field if non-nil, zero value otherwise.

### GetScreeningScoreOk

`func (o *OriginalCreditRequestModel) GetScreeningScoreOk() (*string, bool)`

GetScreeningScoreOk returns a tuple with the ScreeningScore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScreeningScore

`func (o *OriginalCreditRequestModel) SetScreeningScore(v string)`

SetScreeningScore sets ScreeningScore field to given value.

### HasScreeningScore

`func (o *OriginalCreditRequestModel) HasScreeningScore() bool`

HasScreeningScore returns a boolean if a field has been set.

### GetSenderData

`func (o *OriginalCreditRequestModel) GetSenderData() OriginalCreditSenderData`

GetSenderData returns the SenderData field if non-nil, zero value otherwise.

### GetSenderDataOk

`func (o *OriginalCreditRequestModel) GetSenderDataOk() (*OriginalCreditSenderData, bool)`

GetSenderDataOk returns a tuple with the SenderData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSenderData

`func (o *OriginalCreditRequestModel) SetSenderData(v OriginalCreditSenderData)`

SetSenderData sets SenderData field to given value.

### HasSenderData

`func (o *OriginalCreditRequestModel) HasSenderData() bool`

HasSenderData returns a boolean if a field has been set.

### GetTransactionPurpose

`func (o *OriginalCreditRequestModel) GetTransactionPurpose() string`

GetTransactionPurpose returns the TransactionPurpose field if non-nil, zero value otherwise.

### GetTransactionPurposeOk

`func (o *OriginalCreditRequestModel) GetTransactionPurposeOk() (*string, bool)`

GetTransactionPurposeOk returns a tuple with the TransactionPurpose field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionPurpose

`func (o *OriginalCreditRequestModel) SetTransactionPurpose(v string)`

SetTransactionPurpose sets TransactionPurpose field to given value.

### HasTransactionPurpose

`func (o *OriginalCreditRequestModel) HasTransactionPurpose() bool`

HasTransactionPurpose returns a boolean if a field has been set.

### GetType

`func (o *OriginalCreditRequestModel) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *OriginalCreditRequestModel) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *OriginalCreditRequestModel) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


