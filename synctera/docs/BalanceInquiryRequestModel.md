# BalanceInquiryRequestModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountType** | **string** |  | 
**CardAcceptor** | [**CardAcceptorModel**](CardAcceptorModel.md) |  | 
**CardToken** | **string** |  | 
**Mid** | **string** |  | 
**NetworkFees** | Pointer to [**[]NetworkFeeModel**](NetworkFeeModel.md) |  | [optional] 
**Pin** | Pointer to **string** |  | [optional] 

## Methods

### NewBalanceInquiryRequestModel

`func NewBalanceInquiryRequestModel(accountType string, cardAcceptor CardAcceptorModel, cardToken string, mid string, ) *BalanceInquiryRequestModel`

NewBalanceInquiryRequestModel instantiates a new BalanceInquiryRequestModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBalanceInquiryRequestModelWithDefaults

`func NewBalanceInquiryRequestModelWithDefaults() *BalanceInquiryRequestModel`

NewBalanceInquiryRequestModelWithDefaults instantiates a new BalanceInquiryRequestModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountType

`func (o *BalanceInquiryRequestModel) GetAccountType() string`

GetAccountType returns the AccountType field if non-nil, zero value otherwise.

### GetAccountTypeOk

`func (o *BalanceInquiryRequestModel) GetAccountTypeOk() (*string, bool)`

GetAccountTypeOk returns a tuple with the AccountType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountType

`func (o *BalanceInquiryRequestModel) SetAccountType(v string)`

SetAccountType sets AccountType field to given value.


### GetCardAcceptor

`func (o *BalanceInquiryRequestModel) GetCardAcceptor() CardAcceptorModel`

GetCardAcceptor returns the CardAcceptor field if non-nil, zero value otherwise.

### GetCardAcceptorOk

`func (o *BalanceInquiryRequestModel) GetCardAcceptorOk() (*CardAcceptorModel, bool)`

GetCardAcceptorOk returns a tuple with the CardAcceptor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardAcceptor

`func (o *BalanceInquiryRequestModel) SetCardAcceptor(v CardAcceptorModel)`

SetCardAcceptor sets CardAcceptor field to given value.


### GetCardToken

`func (o *BalanceInquiryRequestModel) GetCardToken() string`

GetCardToken returns the CardToken field if non-nil, zero value otherwise.

### GetCardTokenOk

`func (o *BalanceInquiryRequestModel) GetCardTokenOk() (*string, bool)`

GetCardTokenOk returns a tuple with the CardToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardToken

`func (o *BalanceInquiryRequestModel) SetCardToken(v string)`

SetCardToken sets CardToken field to given value.


### GetMid

`func (o *BalanceInquiryRequestModel) GetMid() string`

GetMid returns the Mid field if non-nil, zero value otherwise.

### GetMidOk

`func (o *BalanceInquiryRequestModel) GetMidOk() (*string, bool)`

GetMidOk returns a tuple with the Mid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMid

`func (o *BalanceInquiryRequestModel) SetMid(v string)`

SetMid sets Mid field to given value.


### GetNetworkFees

`func (o *BalanceInquiryRequestModel) GetNetworkFees() []NetworkFeeModel`

GetNetworkFees returns the NetworkFees field if non-nil, zero value otherwise.

### GetNetworkFeesOk

`func (o *BalanceInquiryRequestModel) GetNetworkFeesOk() (*[]NetworkFeeModel, bool)`

GetNetworkFeesOk returns a tuple with the NetworkFees field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkFees

`func (o *BalanceInquiryRequestModel) SetNetworkFees(v []NetworkFeeModel)`

SetNetworkFees sets NetworkFees field to given value.

### HasNetworkFees

`func (o *BalanceInquiryRequestModel) HasNetworkFees() bool`

HasNetworkFees returns a boolean if a field has been set.

### GetPin

`func (o *BalanceInquiryRequestModel) GetPin() string`

GetPin returns the Pin field if non-nil, zero value otherwise.

### GetPinOk

`func (o *BalanceInquiryRequestModel) GetPinOk() (*string, bool)`

GetPinOk returns a tuple with the Pin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPin

`func (o *BalanceInquiryRequestModel) SetPin(v string)`

SetPin sets Pin field to given value.

### HasPin

`func (o *BalanceInquiryRequestModel) HasPin() bool`

HasPin returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


