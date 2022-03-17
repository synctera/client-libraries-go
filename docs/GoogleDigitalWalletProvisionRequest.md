# GoogleDigitalWalletProvisionRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CardId** | **string** | The unique identifier of a card | 
**DeviceId** | **string** | The user’s Android device ID; the device’s unique identifier. | 
**DeviceType** | [**DeviceType**](DeviceType.md) |  | 
**ProvisioningAppVersion** | **string** | Version of the application making the provisioning request. | 
**WalletAccountId** | **string** | The user’s Google wallet account ID. | 

## Methods

### NewGoogleDigitalWalletProvisionRequest

`func NewGoogleDigitalWalletProvisionRequest(cardId string, deviceId string, deviceType DeviceType, provisioningAppVersion string, walletAccountId string, ) *GoogleDigitalWalletProvisionRequest`

NewGoogleDigitalWalletProvisionRequest instantiates a new GoogleDigitalWalletProvisionRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGoogleDigitalWalletProvisionRequestWithDefaults

`func NewGoogleDigitalWalletProvisionRequestWithDefaults() *GoogleDigitalWalletProvisionRequest`

NewGoogleDigitalWalletProvisionRequestWithDefaults instantiates a new GoogleDigitalWalletProvisionRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCardId

`func (o *GoogleDigitalWalletProvisionRequest) GetCardId() string`

GetCardId returns the CardId field if non-nil, zero value otherwise.

### GetCardIdOk

`func (o *GoogleDigitalWalletProvisionRequest) GetCardIdOk() (*string, bool)`

GetCardIdOk returns a tuple with the CardId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardId

`func (o *GoogleDigitalWalletProvisionRequest) SetCardId(v string)`

SetCardId sets CardId field to given value.


### GetDeviceId

`func (o *GoogleDigitalWalletProvisionRequest) GetDeviceId() string`

GetDeviceId returns the DeviceId field if non-nil, zero value otherwise.

### GetDeviceIdOk

`func (o *GoogleDigitalWalletProvisionRequest) GetDeviceIdOk() (*string, bool)`

GetDeviceIdOk returns a tuple with the DeviceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceId

`func (o *GoogleDigitalWalletProvisionRequest) SetDeviceId(v string)`

SetDeviceId sets DeviceId field to given value.


### GetDeviceType

`func (o *GoogleDigitalWalletProvisionRequest) GetDeviceType() DeviceType`

GetDeviceType returns the DeviceType field if non-nil, zero value otherwise.

### GetDeviceTypeOk

`func (o *GoogleDigitalWalletProvisionRequest) GetDeviceTypeOk() (*DeviceType, bool)`

GetDeviceTypeOk returns a tuple with the DeviceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceType

`func (o *GoogleDigitalWalletProvisionRequest) SetDeviceType(v DeviceType)`

SetDeviceType sets DeviceType field to given value.


### GetProvisioningAppVersion

`func (o *GoogleDigitalWalletProvisionRequest) GetProvisioningAppVersion() string`

GetProvisioningAppVersion returns the ProvisioningAppVersion field if non-nil, zero value otherwise.

### GetProvisioningAppVersionOk

`func (o *GoogleDigitalWalletProvisionRequest) GetProvisioningAppVersionOk() (*string, bool)`

GetProvisioningAppVersionOk returns a tuple with the ProvisioningAppVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvisioningAppVersion

`func (o *GoogleDigitalWalletProvisionRequest) SetProvisioningAppVersion(v string)`

SetProvisioningAppVersion sets ProvisioningAppVersion field to given value.


### GetWalletAccountId

`func (o *GoogleDigitalWalletProvisionRequest) GetWalletAccountId() string`

GetWalletAccountId returns the WalletAccountId field if non-nil, zero value otherwise.

### GetWalletAccountIdOk

`func (o *GoogleDigitalWalletProvisionRequest) GetWalletAccountIdOk() (*string, bool)`

GetWalletAccountIdOk returns a tuple with the WalletAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWalletAccountId

`func (o *GoogleDigitalWalletProvisionRequest) SetWalletAccountId(v string)`

SetWalletAccountId sets WalletAccountId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


