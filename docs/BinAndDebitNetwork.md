# BinAndDebitNetwork

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Bin** | [**Bin**](Bin.md) |  | 
**DebitNetwork** | [**DebitNetwork**](DebitNetwork.md) |  | 
**BankNetworkId** | **string** | The ID of the bank network | 

## Methods

### NewBinAndDebitNetwork

`func NewBinAndDebitNetwork(bin Bin, debitNetwork DebitNetwork, bankNetworkId string, ) *BinAndDebitNetwork`

NewBinAndDebitNetwork instantiates a new BinAndDebitNetwork object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBinAndDebitNetworkWithDefaults

`func NewBinAndDebitNetworkWithDefaults() *BinAndDebitNetwork`

NewBinAndDebitNetworkWithDefaults instantiates a new BinAndDebitNetwork object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBin

`func (o *BinAndDebitNetwork) GetBin() Bin`

GetBin returns the Bin field if non-nil, zero value otherwise.

### GetBinOk

`func (o *BinAndDebitNetwork) GetBinOk() (*Bin, bool)`

GetBinOk returns a tuple with the Bin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBin

`func (o *BinAndDebitNetwork) SetBin(v Bin)`

SetBin sets Bin field to given value.


### GetDebitNetwork

`func (o *BinAndDebitNetwork) GetDebitNetwork() DebitNetwork`

GetDebitNetwork returns the DebitNetwork field if non-nil, zero value otherwise.

### GetDebitNetworkOk

`func (o *BinAndDebitNetwork) GetDebitNetworkOk() (*DebitNetwork, bool)`

GetDebitNetworkOk returns a tuple with the DebitNetwork field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDebitNetwork

`func (o *BinAndDebitNetwork) SetDebitNetwork(v DebitNetwork)`

SetDebitNetwork sets DebitNetwork field to given value.


### GetBankNetworkId

`func (o *BinAndDebitNetwork) GetBankNetworkId() string`

GetBankNetworkId returns the BankNetworkId field if non-nil, zero value otherwise.

### GetBankNetworkIdOk

`func (o *BinAndDebitNetwork) GetBankNetworkIdOk() (*string, bool)`

GetBankNetworkIdOk returns a tuple with the BankNetworkId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBankNetworkId

`func (o *BinAndDebitNetwork) SetBankNetworkId(v string)`

SetBankNetworkId sets BankNetworkId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


