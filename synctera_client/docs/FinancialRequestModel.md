# FinancialRequestModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | **int32** | The amount of the transaction in the smallest whole denomination of the applicable currency (eg. For USD use cents) | 
**CardAcceptor** | [**CardAcceptorModel**](CardAcceptorModel.md) |  | 
**CardId** | **string** |  | 
**CashBackAmount** | Pointer to **int32** | Cash back amount in the smallest whole denomination of the applicable currency (eg. For USD use cents) | [optional] 
**IsPreAuth** | Pointer to **bool** |  | [optional] [default to false]
**Mid** | **string** |  | 
**Pin** | Pointer to **string** |  | [optional] 
**TransactionOptions** | Pointer to [**TransactionOptions**](TransactionOptions.md) |  | [optional] 

## Methods

### NewFinancialRequestModel

`func NewFinancialRequestModel(amount int32, cardAcceptor CardAcceptorModel, cardId string, mid string, ) *FinancialRequestModel`

NewFinancialRequestModel instantiates a new FinancialRequestModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFinancialRequestModelWithDefaults

`func NewFinancialRequestModelWithDefaults() *FinancialRequestModel`

NewFinancialRequestModelWithDefaults instantiates a new FinancialRequestModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmount

`func (o *FinancialRequestModel) GetAmount() int32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *FinancialRequestModel) GetAmountOk() (*int32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *FinancialRequestModel) SetAmount(v int32)`

SetAmount sets Amount field to given value.


### GetCardAcceptor

`func (o *FinancialRequestModel) GetCardAcceptor() CardAcceptorModel`

GetCardAcceptor returns the CardAcceptor field if non-nil, zero value otherwise.

### GetCardAcceptorOk

`func (o *FinancialRequestModel) GetCardAcceptorOk() (*CardAcceptorModel, bool)`

GetCardAcceptorOk returns a tuple with the CardAcceptor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardAcceptor

`func (o *FinancialRequestModel) SetCardAcceptor(v CardAcceptorModel)`

SetCardAcceptor sets CardAcceptor field to given value.


### GetCardId

`func (o *FinancialRequestModel) GetCardId() string`

GetCardId returns the CardId field if non-nil, zero value otherwise.

### GetCardIdOk

`func (o *FinancialRequestModel) GetCardIdOk() (*string, bool)`

GetCardIdOk returns a tuple with the CardId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardId

`func (o *FinancialRequestModel) SetCardId(v string)`

SetCardId sets CardId field to given value.


### GetCashBackAmount

`func (o *FinancialRequestModel) GetCashBackAmount() int32`

GetCashBackAmount returns the CashBackAmount field if non-nil, zero value otherwise.

### GetCashBackAmountOk

`func (o *FinancialRequestModel) GetCashBackAmountOk() (*int32, bool)`

GetCashBackAmountOk returns a tuple with the CashBackAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCashBackAmount

`func (o *FinancialRequestModel) SetCashBackAmount(v int32)`

SetCashBackAmount sets CashBackAmount field to given value.

### HasCashBackAmount

`func (o *FinancialRequestModel) HasCashBackAmount() bool`

HasCashBackAmount returns a boolean if a field has been set.

### GetIsPreAuth

`func (o *FinancialRequestModel) GetIsPreAuth() bool`

GetIsPreAuth returns the IsPreAuth field if non-nil, zero value otherwise.

### GetIsPreAuthOk

`func (o *FinancialRequestModel) GetIsPreAuthOk() (*bool, bool)`

GetIsPreAuthOk returns a tuple with the IsPreAuth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPreAuth

`func (o *FinancialRequestModel) SetIsPreAuth(v bool)`

SetIsPreAuth sets IsPreAuth field to given value.

### HasIsPreAuth

`func (o *FinancialRequestModel) HasIsPreAuth() bool`

HasIsPreAuth returns a boolean if a field has been set.

### GetMid

`func (o *FinancialRequestModel) GetMid() string`

GetMid returns the Mid field if non-nil, zero value otherwise.

### GetMidOk

`func (o *FinancialRequestModel) GetMidOk() (*string, bool)`

GetMidOk returns a tuple with the Mid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMid

`func (o *FinancialRequestModel) SetMid(v string)`

SetMid sets Mid field to given value.


### GetPin

`func (o *FinancialRequestModel) GetPin() string`

GetPin returns the Pin field if non-nil, zero value otherwise.

### GetPinOk

`func (o *FinancialRequestModel) GetPinOk() (*string, bool)`

GetPinOk returns a tuple with the Pin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPin

`func (o *FinancialRequestModel) SetPin(v string)`

SetPin sets Pin field to given value.

### HasPin

`func (o *FinancialRequestModel) HasPin() bool`

HasPin returns a boolean if a field has been set.

### GetTransactionOptions

`func (o *FinancialRequestModel) GetTransactionOptions() TransactionOptions`

GetTransactionOptions returns the TransactionOptions field if non-nil, zero value otherwise.

### GetTransactionOptionsOk

`func (o *FinancialRequestModel) GetTransactionOptionsOk() (*TransactionOptions, bool)`

GetTransactionOptionsOk returns a tuple with the TransactionOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionOptions

`func (o *FinancialRequestModel) SetTransactionOptions(v TransactionOptions)`

SetTransactionOptions sets TransactionOptions field to given value.

### HasTransactionOptions

`func (o *FinancialRequestModel) HasTransactionOptions() bool`

HasTransactionOptions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


