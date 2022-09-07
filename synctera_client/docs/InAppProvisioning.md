# InAppProvisioning

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AddressVerification** | Pointer to [**DigitalWalletTokenAddressVerification**](DigitalWalletTokenAddressVerification.md) |  | [optional] 
**Enabled** | Pointer to **bool** |  | [optional] [default to false]

## Methods

### NewInAppProvisioning

`func NewInAppProvisioning() *InAppProvisioning`

NewInAppProvisioning instantiates a new InAppProvisioning object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInAppProvisioningWithDefaults

`func NewInAppProvisioningWithDefaults() *InAppProvisioning`

NewInAppProvisioningWithDefaults instantiates a new InAppProvisioning object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddressVerification

`func (o *InAppProvisioning) GetAddressVerification() DigitalWalletTokenAddressVerification`

GetAddressVerification returns the AddressVerification field if non-nil, zero value otherwise.

### GetAddressVerificationOk

`func (o *InAppProvisioning) GetAddressVerificationOk() (*DigitalWalletTokenAddressVerification, bool)`

GetAddressVerificationOk returns a tuple with the AddressVerification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressVerification

`func (o *InAppProvisioning) SetAddressVerification(v DigitalWalletTokenAddressVerification)`

SetAddressVerification sets AddressVerification field to given value.

### HasAddressVerification

`func (o *InAppProvisioning) HasAddressVerification() bool`

HasAddressVerification returns a boolean if a field has been set.

### GetEnabled

`func (o *InAppProvisioning) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *InAppProvisioning) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *InAppProvisioning) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *InAppProvisioning) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


