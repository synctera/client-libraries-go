# AuthRequestModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | **int32** | The amount of the transaction in the smallest whole denomination of the applicable currency (eg. For USD use cents) | 
**CardAcceptor** | Pointer to [**CardAcceptorModel**](CardAcceptorModel.md) |  | [optional] 
**CardId** | **string** |  | 
**CardOptions** | Pointer to [**CardOptions**](CardOptions.md) |  | [optional] 
**CashBackAmount** | Pointer to **int32** |  | [optional] 
**IsPreAuth** | Pointer to **bool** |  | [optional] [default to false]
**Mid** | **string** |  | 
**NetworkFees** | Pointer to [**[]NetworkFeeModel**](NetworkFeeModel.md) |  | [optional] 
**Pin** | Pointer to **string** |  | [optional] 
**TransactionOptions** | Pointer to [**TransactionOptions**](TransactionOptions.md) |  | [optional] 

## Methods

### NewAuthRequestModel

`func NewAuthRequestModel(amount int32, cardId string, mid string, ) *AuthRequestModel`

NewAuthRequestModel instantiates a new AuthRequestModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthRequestModelWithDefaults

`func NewAuthRequestModelWithDefaults() *AuthRequestModel`

NewAuthRequestModelWithDefaults instantiates a new AuthRequestModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmount

`func (o *AuthRequestModel) GetAmount() int32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *AuthRequestModel) GetAmountOk() (*int32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *AuthRequestModel) SetAmount(v int32)`

SetAmount sets Amount field to given value.


### GetCardAcceptor

`func (o *AuthRequestModel) GetCardAcceptor() CardAcceptorModel`

GetCardAcceptor returns the CardAcceptor field if non-nil, zero value otherwise.

### GetCardAcceptorOk

`func (o *AuthRequestModel) GetCardAcceptorOk() (*CardAcceptorModel, bool)`

GetCardAcceptorOk returns a tuple with the CardAcceptor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardAcceptor

`func (o *AuthRequestModel) SetCardAcceptor(v CardAcceptorModel)`

SetCardAcceptor sets CardAcceptor field to given value.

### HasCardAcceptor

`func (o *AuthRequestModel) HasCardAcceptor() bool`

HasCardAcceptor returns a boolean if a field has been set.

### GetCardId

`func (o *AuthRequestModel) GetCardId() string`

GetCardId returns the CardId field if non-nil, zero value otherwise.

### GetCardIdOk

`func (o *AuthRequestModel) GetCardIdOk() (*string, bool)`

GetCardIdOk returns a tuple with the CardId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardId

`func (o *AuthRequestModel) SetCardId(v string)`

SetCardId sets CardId field to given value.


### GetCardOptions

`func (o *AuthRequestModel) GetCardOptions() CardOptions`

GetCardOptions returns the CardOptions field if non-nil, zero value otherwise.

### GetCardOptionsOk

`func (o *AuthRequestModel) GetCardOptionsOk() (*CardOptions, bool)`

GetCardOptionsOk returns a tuple with the CardOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardOptions

`func (o *AuthRequestModel) SetCardOptions(v CardOptions)`

SetCardOptions sets CardOptions field to given value.

### HasCardOptions

`func (o *AuthRequestModel) HasCardOptions() bool`

HasCardOptions returns a boolean if a field has been set.

### GetCashBackAmount

`func (o *AuthRequestModel) GetCashBackAmount() int32`

GetCashBackAmount returns the CashBackAmount field if non-nil, zero value otherwise.

### GetCashBackAmountOk

`func (o *AuthRequestModel) GetCashBackAmountOk() (*int32, bool)`

GetCashBackAmountOk returns a tuple with the CashBackAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCashBackAmount

`func (o *AuthRequestModel) SetCashBackAmount(v int32)`

SetCashBackAmount sets CashBackAmount field to given value.

### HasCashBackAmount

`func (o *AuthRequestModel) HasCashBackAmount() bool`

HasCashBackAmount returns a boolean if a field has been set.

### GetIsPreAuth

`func (o *AuthRequestModel) GetIsPreAuth() bool`

GetIsPreAuth returns the IsPreAuth field if non-nil, zero value otherwise.

### GetIsPreAuthOk

`func (o *AuthRequestModel) GetIsPreAuthOk() (*bool, bool)`

GetIsPreAuthOk returns a tuple with the IsPreAuth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPreAuth

`func (o *AuthRequestModel) SetIsPreAuth(v bool)`

SetIsPreAuth sets IsPreAuth field to given value.

### HasIsPreAuth

`func (o *AuthRequestModel) HasIsPreAuth() bool`

HasIsPreAuth returns a boolean if a field has been set.

### GetMid

`func (o *AuthRequestModel) GetMid() string`

GetMid returns the Mid field if non-nil, zero value otherwise.

### GetMidOk

`func (o *AuthRequestModel) GetMidOk() (*string, bool)`

GetMidOk returns a tuple with the Mid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMid

`func (o *AuthRequestModel) SetMid(v string)`

SetMid sets Mid field to given value.


### GetNetworkFees

`func (o *AuthRequestModel) GetNetworkFees() []NetworkFeeModel`

GetNetworkFees returns the NetworkFees field if non-nil, zero value otherwise.

### GetNetworkFeesOk

`func (o *AuthRequestModel) GetNetworkFeesOk() (*[]NetworkFeeModel, bool)`

GetNetworkFeesOk returns a tuple with the NetworkFees field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkFees

`func (o *AuthRequestModel) SetNetworkFees(v []NetworkFeeModel)`

SetNetworkFees sets NetworkFees field to given value.

### HasNetworkFees

`func (o *AuthRequestModel) HasNetworkFees() bool`

HasNetworkFees returns a boolean if a field has been set.

### GetPin

`func (o *AuthRequestModel) GetPin() string`

GetPin returns the Pin field if non-nil, zero value otherwise.

### GetPinOk

`func (o *AuthRequestModel) GetPinOk() (*string, bool)`

GetPinOk returns a tuple with the Pin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPin

`func (o *AuthRequestModel) SetPin(v string)`

SetPin sets Pin field to given value.

### HasPin

`func (o *AuthRequestModel) HasPin() bool`

HasPin returns a boolean if a field has been set.

### GetTransactionOptions

`func (o *AuthRequestModel) GetTransactionOptions() TransactionOptions`

GetTransactionOptions returns the TransactionOptions field if non-nil, zero value otherwise.

### GetTransactionOptionsOk

`func (o *AuthRequestModel) GetTransactionOptionsOk() (*TransactionOptions, bool)`

GetTransactionOptionsOk returns a tuple with the TransactionOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionOptions

`func (o *AuthRequestModel) SetTransactionOptions(v TransactionOptions)`

SetTransactionOptions sets TransactionOptions field to given value.

### HasTransactionOptions

`func (o *AuthRequestModel) HasTransactionOptions() bool`

HasTransactionOptions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


