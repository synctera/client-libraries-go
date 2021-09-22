# AuthorizationAdviceModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | **float32** |  | 
**NetworkFees** | Pointer to [**[]NetworkFeeModel**](NetworkFeeModel.md) |  | [optional] 
**OriginalTransactionToken** | **string** |  | 
**TransactionOptions** | Pointer to [**TransactionOptions**](TransactionOptions.md) |  | [optional] 

## Methods

### NewAuthorizationAdviceModel

`func NewAuthorizationAdviceModel(amount float32, originalTransactionToken string, ) *AuthorizationAdviceModel`

NewAuthorizationAdviceModel instantiates a new AuthorizationAdviceModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthorizationAdviceModelWithDefaults

`func NewAuthorizationAdviceModelWithDefaults() *AuthorizationAdviceModel`

NewAuthorizationAdviceModelWithDefaults instantiates a new AuthorizationAdviceModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmount

`func (o *AuthorizationAdviceModel) GetAmount() float32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *AuthorizationAdviceModel) GetAmountOk() (*float32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *AuthorizationAdviceModel) SetAmount(v float32)`

SetAmount sets Amount field to given value.


### GetNetworkFees

`func (o *AuthorizationAdviceModel) GetNetworkFees() []NetworkFeeModel`

GetNetworkFees returns the NetworkFees field if non-nil, zero value otherwise.

### GetNetworkFeesOk

`func (o *AuthorizationAdviceModel) GetNetworkFeesOk() (*[]NetworkFeeModel, bool)`

GetNetworkFeesOk returns a tuple with the NetworkFees field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkFees

`func (o *AuthorizationAdviceModel) SetNetworkFees(v []NetworkFeeModel)`

SetNetworkFees sets NetworkFees field to given value.

### HasNetworkFees

`func (o *AuthorizationAdviceModel) HasNetworkFees() bool`

HasNetworkFees returns a boolean if a field has been set.

### GetOriginalTransactionToken

`func (o *AuthorizationAdviceModel) GetOriginalTransactionToken() string`

GetOriginalTransactionToken returns the OriginalTransactionToken field if non-nil, zero value otherwise.

### GetOriginalTransactionTokenOk

`func (o *AuthorizationAdviceModel) GetOriginalTransactionTokenOk() (*string, bool)`

GetOriginalTransactionTokenOk returns a tuple with the OriginalTransactionToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginalTransactionToken

`func (o *AuthorizationAdviceModel) SetOriginalTransactionToken(v string)`

SetOriginalTransactionToken sets OriginalTransactionToken field to given value.


### GetTransactionOptions

`func (o *AuthorizationAdviceModel) GetTransactionOptions() TransactionOptions`

GetTransactionOptions returns the TransactionOptions field if non-nil, zero value otherwise.

### GetTransactionOptionsOk

`func (o *AuthorizationAdviceModel) GetTransactionOptionsOk() (*TransactionOptions, bool)`

GetTransactionOptionsOk returns a tuple with the TransactionOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionOptions

`func (o *AuthorizationAdviceModel) SetTransactionOptions(v TransactionOptions)`

SetTransactionOptions sets TransactionOptions field to given value.

### HasTransactionOptions

`func (o *AuthorizationAdviceModel) HasTransactionOptions() bool`

HasTransactionOptions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


