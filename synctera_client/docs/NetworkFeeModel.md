# NetworkFeeModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | **int32** | The amount of the fee in the smallest whole denomination of the applicable currency (eg. For USD use cents) | 
**CreditDebit** | Pointer to **string** | C &#x3D; credit; D &#x3D; debit | [optional] 
**Type** | **string** |  | 

## Methods

### NewNetworkFeeModel

`func NewNetworkFeeModel(amount int32, type_ string, ) *NetworkFeeModel`

NewNetworkFeeModel instantiates a new NetworkFeeModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkFeeModelWithDefaults

`func NewNetworkFeeModelWithDefaults() *NetworkFeeModel`

NewNetworkFeeModelWithDefaults instantiates a new NetworkFeeModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmount

`func (o *NetworkFeeModel) GetAmount() int32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *NetworkFeeModel) GetAmountOk() (*int32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *NetworkFeeModel) SetAmount(v int32)`

SetAmount sets Amount field to given value.


### GetCreditDebit

`func (o *NetworkFeeModel) GetCreditDebit() string`

GetCreditDebit returns the CreditDebit field if non-nil, zero value otherwise.

### GetCreditDebitOk

`func (o *NetworkFeeModel) GetCreditDebitOk() (*string, bool)`

GetCreditDebitOk returns a tuple with the CreditDebit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreditDebit

`func (o *NetworkFeeModel) SetCreditDebit(v string)`

SetCreditDebit sets CreditDebit field to given value.

### HasCreditDebit

`func (o *NetworkFeeModel) HasCreditDebit() bool`

HasCreditDebit returns a boolean if a field has been set.

### GetType

`func (o *NetworkFeeModel) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *NetworkFeeModel) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *NetworkFeeModel) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


