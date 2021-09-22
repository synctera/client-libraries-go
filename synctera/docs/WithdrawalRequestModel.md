# WithdrawalRequestModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountType** | Pointer to **string** |  | [optional] 
**Amount** | **float32** |  | 
**CardAcceptor** | Pointer to [**CardAcceptorModel**](CardAcceptorModel.md) |  | [optional] 
**CardToken** | **string** |  | 
**Mid** | **string** |  | 
**Pin** | Pointer to **string** |  | [optional] 

## Methods

### NewWithdrawalRequestModel

`func NewWithdrawalRequestModel(amount float32, cardToken string, mid string, ) *WithdrawalRequestModel`

NewWithdrawalRequestModel instantiates a new WithdrawalRequestModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWithdrawalRequestModelWithDefaults

`func NewWithdrawalRequestModelWithDefaults() *WithdrawalRequestModel`

NewWithdrawalRequestModelWithDefaults instantiates a new WithdrawalRequestModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountType

`func (o *WithdrawalRequestModel) GetAccountType() string`

GetAccountType returns the AccountType field if non-nil, zero value otherwise.

### GetAccountTypeOk

`func (o *WithdrawalRequestModel) GetAccountTypeOk() (*string, bool)`

GetAccountTypeOk returns a tuple with the AccountType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountType

`func (o *WithdrawalRequestModel) SetAccountType(v string)`

SetAccountType sets AccountType field to given value.

### HasAccountType

`func (o *WithdrawalRequestModel) HasAccountType() bool`

HasAccountType returns a boolean if a field has been set.

### GetAmount

`func (o *WithdrawalRequestModel) GetAmount() float32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *WithdrawalRequestModel) GetAmountOk() (*float32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *WithdrawalRequestModel) SetAmount(v float32)`

SetAmount sets Amount field to given value.


### GetCardAcceptor

`func (o *WithdrawalRequestModel) GetCardAcceptor() CardAcceptorModel`

GetCardAcceptor returns the CardAcceptor field if non-nil, zero value otherwise.

### GetCardAcceptorOk

`func (o *WithdrawalRequestModel) GetCardAcceptorOk() (*CardAcceptorModel, bool)`

GetCardAcceptorOk returns a tuple with the CardAcceptor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardAcceptor

`func (o *WithdrawalRequestModel) SetCardAcceptor(v CardAcceptorModel)`

SetCardAcceptor sets CardAcceptor field to given value.

### HasCardAcceptor

`func (o *WithdrawalRequestModel) HasCardAcceptor() bool`

HasCardAcceptor returns a boolean if a field has been set.

### GetCardToken

`func (o *WithdrawalRequestModel) GetCardToken() string`

GetCardToken returns the CardToken field if non-nil, zero value otherwise.

### GetCardTokenOk

`func (o *WithdrawalRequestModel) GetCardTokenOk() (*string, bool)`

GetCardTokenOk returns a tuple with the CardToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardToken

`func (o *WithdrawalRequestModel) SetCardToken(v string)`

SetCardToken sets CardToken field to given value.


### GetMid

`func (o *WithdrawalRequestModel) GetMid() string`

GetMid returns the Mid field if non-nil, zero value otherwise.

### GetMidOk

`func (o *WithdrawalRequestModel) GetMidOk() (*string, bool)`

GetMidOk returns a tuple with the Mid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMid

`func (o *WithdrawalRequestModel) SetMid(v string)`

SetMid sets Mid field to given value.


### GetPin

`func (o *WithdrawalRequestModel) GetPin() string`

GetPin returns the Pin field if non-nil, zero value otherwise.

### GetPinOk

`func (o *WithdrawalRequestModel) GetPinOk() (*string, bool)`

GetPinOk returns a tuple with the Pin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPin

`func (o *WithdrawalRequestModel) SetPin(v string)`

SetPin sets Pin field to given value.

### HasPin

`func (o *WithdrawalRequestModel) HasPin() bool`

HasPin returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


