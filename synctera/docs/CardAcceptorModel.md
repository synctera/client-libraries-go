# CardAcceptorModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address** | Pointer to **string** |  | [optional] 
**City** | Pointer to **string** |  | [optional] 
**Country** | Pointer to **string** |  | [optional] 
**EcommerceSecurityLevelIndicator** | Pointer to **string** |  | [optional] 
**Mcc** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**PartialApprovalCapable** | Pointer to **bool** |  | [optional] [default to false]
**State** | Pointer to **string** |  | [optional] 
**Zip** | Pointer to **string** |  | [optional] 

## Methods

### NewCardAcceptorModel

`func NewCardAcceptorModel() *CardAcceptorModel`

NewCardAcceptorModel instantiates a new CardAcceptorModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCardAcceptorModelWithDefaults

`func NewCardAcceptorModelWithDefaults() *CardAcceptorModel`

NewCardAcceptorModelWithDefaults instantiates a new CardAcceptorModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddress

`func (o *CardAcceptorModel) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *CardAcceptorModel) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *CardAcceptorModel) SetAddress(v string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *CardAcceptorModel) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetCity

`func (o *CardAcceptorModel) GetCity() string`

GetCity returns the City field if non-nil, zero value otherwise.

### GetCityOk

`func (o *CardAcceptorModel) GetCityOk() (*string, bool)`

GetCityOk returns a tuple with the City field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCity

`func (o *CardAcceptorModel) SetCity(v string)`

SetCity sets City field to given value.

### HasCity

`func (o *CardAcceptorModel) HasCity() bool`

HasCity returns a boolean if a field has been set.

### GetCountry

`func (o *CardAcceptorModel) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *CardAcceptorModel) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *CardAcceptorModel) SetCountry(v string)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *CardAcceptorModel) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### GetEcommerceSecurityLevelIndicator

`func (o *CardAcceptorModel) GetEcommerceSecurityLevelIndicator() string`

GetEcommerceSecurityLevelIndicator returns the EcommerceSecurityLevelIndicator field if non-nil, zero value otherwise.

### GetEcommerceSecurityLevelIndicatorOk

`func (o *CardAcceptorModel) GetEcommerceSecurityLevelIndicatorOk() (*string, bool)`

GetEcommerceSecurityLevelIndicatorOk returns a tuple with the EcommerceSecurityLevelIndicator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEcommerceSecurityLevelIndicator

`func (o *CardAcceptorModel) SetEcommerceSecurityLevelIndicator(v string)`

SetEcommerceSecurityLevelIndicator sets EcommerceSecurityLevelIndicator field to given value.

### HasEcommerceSecurityLevelIndicator

`func (o *CardAcceptorModel) HasEcommerceSecurityLevelIndicator() bool`

HasEcommerceSecurityLevelIndicator returns a boolean if a field has been set.

### GetMcc

`func (o *CardAcceptorModel) GetMcc() string`

GetMcc returns the Mcc field if non-nil, zero value otherwise.

### GetMccOk

`func (o *CardAcceptorModel) GetMccOk() (*string, bool)`

GetMccOk returns a tuple with the Mcc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMcc

`func (o *CardAcceptorModel) SetMcc(v string)`

SetMcc sets Mcc field to given value.

### HasMcc

`func (o *CardAcceptorModel) HasMcc() bool`

HasMcc returns a boolean if a field has been set.

### GetName

`func (o *CardAcceptorModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CardAcceptorModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CardAcceptorModel) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CardAcceptorModel) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPartialApprovalCapable

`func (o *CardAcceptorModel) GetPartialApprovalCapable() bool`

GetPartialApprovalCapable returns the PartialApprovalCapable field if non-nil, zero value otherwise.

### GetPartialApprovalCapableOk

`func (o *CardAcceptorModel) GetPartialApprovalCapableOk() (*bool, bool)`

GetPartialApprovalCapableOk returns a tuple with the PartialApprovalCapable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartialApprovalCapable

`func (o *CardAcceptorModel) SetPartialApprovalCapable(v bool)`

SetPartialApprovalCapable sets PartialApprovalCapable field to given value.

### HasPartialApprovalCapable

`func (o *CardAcceptorModel) HasPartialApprovalCapable() bool`

HasPartialApprovalCapable returns a boolean if a field has been set.

### GetState

`func (o *CardAcceptorModel) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *CardAcceptorModel) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *CardAcceptorModel) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *CardAcceptorModel) HasState() bool`

HasState returns a boolean if a field has been set.

### GetZip

`func (o *CardAcceptorModel) GetZip() string`

GetZip returns the Zip field if non-nil, zero value otherwise.

### GetZipOk

`func (o *CardAcceptorModel) GetZipOk() (*string, bool)`

GetZipOk returns a tuple with the Zip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZip

`func (o *CardAcceptorModel) SetZip(v string)`

SetZip sets Zip field to given value.

### HasZip

`func (o *CardAcceptorModel) HasZip() bool`

HasZip returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


