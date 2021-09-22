# ReversalModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | **float32** |  | 
**FindOriginalWindowDays** | Pointer to **int32** |  | [optional] 
**IsAdvice** | Pointer to **bool** |  | [optional] [default to false]
**NetworkFees** | Pointer to [**[]NetworkFeeModel**](NetworkFeeModel.md) |  | [optional] 
**OriginalTransactionToken** | **string** |  | 

## Methods

### NewReversalModel

`func NewReversalModel(amount float32, originalTransactionToken string, ) *ReversalModel`

NewReversalModel instantiates a new ReversalModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReversalModelWithDefaults

`func NewReversalModelWithDefaults() *ReversalModel`

NewReversalModelWithDefaults instantiates a new ReversalModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmount

`func (o *ReversalModel) GetAmount() float32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *ReversalModel) GetAmountOk() (*float32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *ReversalModel) SetAmount(v float32)`

SetAmount sets Amount field to given value.


### GetFindOriginalWindowDays

`func (o *ReversalModel) GetFindOriginalWindowDays() int32`

GetFindOriginalWindowDays returns the FindOriginalWindowDays field if non-nil, zero value otherwise.

### GetFindOriginalWindowDaysOk

`func (o *ReversalModel) GetFindOriginalWindowDaysOk() (*int32, bool)`

GetFindOriginalWindowDaysOk returns a tuple with the FindOriginalWindowDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFindOriginalWindowDays

`func (o *ReversalModel) SetFindOriginalWindowDays(v int32)`

SetFindOriginalWindowDays sets FindOriginalWindowDays field to given value.

### HasFindOriginalWindowDays

`func (o *ReversalModel) HasFindOriginalWindowDays() bool`

HasFindOriginalWindowDays returns a boolean if a field has been set.

### GetIsAdvice

`func (o *ReversalModel) GetIsAdvice() bool`

GetIsAdvice returns the IsAdvice field if non-nil, zero value otherwise.

### GetIsAdviceOk

`func (o *ReversalModel) GetIsAdviceOk() (*bool, bool)`

GetIsAdviceOk returns a tuple with the IsAdvice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAdvice

`func (o *ReversalModel) SetIsAdvice(v bool)`

SetIsAdvice sets IsAdvice field to given value.

### HasIsAdvice

`func (o *ReversalModel) HasIsAdvice() bool`

HasIsAdvice returns a boolean if a field has been set.

### GetNetworkFees

`func (o *ReversalModel) GetNetworkFees() []NetworkFeeModel`

GetNetworkFees returns the NetworkFees field if non-nil, zero value otherwise.

### GetNetworkFeesOk

`func (o *ReversalModel) GetNetworkFeesOk() (*[]NetworkFeeModel, bool)`

GetNetworkFeesOk returns a tuple with the NetworkFees field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkFees

`func (o *ReversalModel) SetNetworkFees(v []NetworkFeeModel)`

SetNetworkFees sets NetworkFees field to given value.

### HasNetworkFees

`func (o *ReversalModel) HasNetworkFees() bool`

HasNetworkFees returns a boolean if a field has been set.

### GetOriginalTransactionToken

`func (o *ReversalModel) GetOriginalTransactionToken() string`

GetOriginalTransactionToken returns the OriginalTransactionToken field if non-nil, zero value otherwise.

### GetOriginalTransactionTokenOk

`func (o *ReversalModel) GetOriginalTransactionTokenOk() (*string, bool)`

GetOriginalTransactionTokenOk returns a tuple with the OriginalTransactionToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginalTransactionToken

`func (o *ReversalModel) SetOriginalTransactionToken(v string)`

SetOriginalTransactionToken sets OriginalTransactionToken field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


