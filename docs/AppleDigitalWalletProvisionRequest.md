# AppleDigitalWalletProvisionRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CardId** | **string** | The unique identifier of a card | 
**Certificates** | **[]string** | Leaf and sub-CA certificates provided by Apple | 
**DeviceType** | [**DeviceType**](DeviceType.md) |  | 
**Nonce** | **string** | One-time-use nonce provided by Apple for security purposes. | 
**NonceSignature** | **string** | Apple-provided signature to the nonce. | 
**ProvisioningAppVersion** | **string** | Version of the application making the provisioning request. | 

## Methods

### NewAppleDigitalWalletProvisionRequest

`func NewAppleDigitalWalletProvisionRequest(cardId string, certificates []string, deviceType DeviceType, nonce string, nonceSignature string, provisioningAppVersion string, ) *AppleDigitalWalletProvisionRequest`

NewAppleDigitalWalletProvisionRequest instantiates a new AppleDigitalWalletProvisionRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppleDigitalWalletProvisionRequestWithDefaults

`func NewAppleDigitalWalletProvisionRequestWithDefaults() *AppleDigitalWalletProvisionRequest`

NewAppleDigitalWalletProvisionRequestWithDefaults instantiates a new AppleDigitalWalletProvisionRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCardId

`func (o *AppleDigitalWalletProvisionRequest) GetCardId() string`

GetCardId returns the CardId field if non-nil, zero value otherwise.

### GetCardIdOk

`func (o *AppleDigitalWalletProvisionRequest) GetCardIdOk() (*string, bool)`

GetCardIdOk returns a tuple with the CardId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardId

`func (o *AppleDigitalWalletProvisionRequest) SetCardId(v string)`

SetCardId sets CardId field to given value.


### GetCertificates

`func (o *AppleDigitalWalletProvisionRequest) GetCertificates() []string`

GetCertificates returns the Certificates field if non-nil, zero value otherwise.

### GetCertificatesOk

`func (o *AppleDigitalWalletProvisionRequest) GetCertificatesOk() (*[]string, bool)`

GetCertificatesOk returns a tuple with the Certificates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificates

`func (o *AppleDigitalWalletProvisionRequest) SetCertificates(v []string)`

SetCertificates sets Certificates field to given value.


### GetDeviceType

`func (o *AppleDigitalWalletProvisionRequest) GetDeviceType() DeviceType`

GetDeviceType returns the DeviceType field if non-nil, zero value otherwise.

### GetDeviceTypeOk

`func (o *AppleDigitalWalletProvisionRequest) GetDeviceTypeOk() (*DeviceType, bool)`

GetDeviceTypeOk returns a tuple with the DeviceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceType

`func (o *AppleDigitalWalletProvisionRequest) SetDeviceType(v DeviceType)`

SetDeviceType sets DeviceType field to given value.


### GetNonce

`func (o *AppleDigitalWalletProvisionRequest) GetNonce() string`

GetNonce returns the Nonce field if non-nil, zero value otherwise.

### GetNonceOk

`func (o *AppleDigitalWalletProvisionRequest) GetNonceOk() (*string, bool)`

GetNonceOk returns a tuple with the Nonce field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNonce

`func (o *AppleDigitalWalletProvisionRequest) SetNonce(v string)`

SetNonce sets Nonce field to given value.


### GetNonceSignature

`func (o *AppleDigitalWalletProvisionRequest) GetNonceSignature() string`

GetNonceSignature returns the NonceSignature field if non-nil, zero value otherwise.

### GetNonceSignatureOk

`func (o *AppleDigitalWalletProvisionRequest) GetNonceSignatureOk() (*string, bool)`

GetNonceSignatureOk returns a tuple with the NonceSignature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNonceSignature

`func (o *AppleDigitalWalletProvisionRequest) SetNonceSignature(v string)`

SetNonceSignature sets NonceSignature field to given value.


### GetProvisioningAppVersion

`func (o *AppleDigitalWalletProvisionRequest) GetProvisioningAppVersion() string`

GetProvisioningAppVersion returns the ProvisioningAppVersion field if non-nil, zero value otherwise.

### GetProvisioningAppVersionOk

`func (o *AppleDigitalWalletProvisionRequest) GetProvisioningAppVersionOk() (*string, bool)`

GetProvisioningAppVersionOk returns a tuple with the ProvisioningAppVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvisioningAppVersion

`func (o *AppleDigitalWalletProvisionRequest) SetProvisioningAppVersion(v string)`

SetProvisioningAppVersion sets ProvisioningAppVersion field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


