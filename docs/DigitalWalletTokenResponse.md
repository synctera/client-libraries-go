# DigitalWalletTokenResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Digital Wallet Token ID | [optional] 
**CardId** | Pointer to **string** | Card ID of the Digital wallet Token | [optional] 
**State** | Pointer to **string** | Current status of the Digital Wallet Token | [optional] 
**Type** | Pointer to **string** | Type of the Digital Wallet | [optional] 
**DeviceType** | Pointer to **string** | Type of the device where the Digital Wallet Token is used in | [optional] 
**DeviceId** | Pointer to **string** | The user’s Android device ID; the device’s unique identifier. | [optional] 
**RequestedTime** | Pointer to **time.Time** |  | [optional] 
**ApprovedTime** | Pointer to **time.Time** |  | [optional] 
**LastModifiedTime** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewDigitalWalletTokenResponse

`func NewDigitalWalletTokenResponse() *DigitalWalletTokenResponse`

NewDigitalWalletTokenResponse instantiates a new DigitalWalletTokenResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDigitalWalletTokenResponseWithDefaults

`func NewDigitalWalletTokenResponseWithDefaults() *DigitalWalletTokenResponse`

NewDigitalWalletTokenResponseWithDefaults instantiates a new DigitalWalletTokenResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *DigitalWalletTokenResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DigitalWalletTokenResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DigitalWalletTokenResponse) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *DigitalWalletTokenResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCardId

`func (o *DigitalWalletTokenResponse) GetCardId() string`

GetCardId returns the CardId field if non-nil, zero value otherwise.

### GetCardIdOk

`func (o *DigitalWalletTokenResponse) GetCardIdOk() (*string, bool)`

GetCardIdOk returns a tuple with the CardId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardId

`func (o *DigitalWalletTokenResponse) SetCardId(v string)`

SetCardId sets CardId field to given value.

### HasCardId

`func (o *DigitalWalletTokenResponse) HasCardId() bool`

HasCardId returns a boolean if a field has been set.

### GetState

`func (o *DigitalWalletTokenResponse) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *DigitalWalletTokenResponse) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *DigitalWalletTokenResponse) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *DigitalWalletTokenResponse) HasState() bool`

HasState returns a boolean if a field has been set.

### GetType

`func (o *DigitalWalletTokenResponse) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *DigitalWalletTokenResponse) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *DigitalWalletTokenResponse) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *DigitalWalletTokenResponse) HasType() bool`

HasType returns a boolean if a field has been set.

### GetDeviceType

`func (o *DigitalWalletTokenResponse) GetDeviceType() string`

GetDeviceType returns the DeviceType field if non-nil, zero value otherwise.

### GetDeviceTypeOk

`func (o *DigitalWalletTokenResponse) GetDeviceTypeOk() (*string, bool)`

GetDeviceTypeOk returns a tuple with the DeviceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceType

`func (o *DigitalWalletTokenResponse) SetDeviceType(v string)`

SetDeviceType sets DeviceType field to given value.

### HasDeviceType

`func (o *DigitalWalletTokenResponse) HasDeviceType() bool`

HasDeviceType returns a boolean if a field has been set.

### GetDeviceId

`func (o *DigitalWalletTokenResponse) GetDeviceId() string`

GetDeviceId returns the DeviceId field if non-nil, zero value otherwise.

### GetDeviceIdOk

`func (o *DigitalWalletTokenResponse) GetDeviceIdOk() (*string, bool)`

GetDeviceIdOk returns a tuple with the DeviceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceId

`func (o *DigitalWalletTokenResponse) SetDeviceId(v string)`

SetDeviceId sets DeviceId field to given value.

### HasDeviceId

`func (o *DigitalWalletTokenResponse) HasDeviceId() bool`

HasDeviceId returns a boolean if a field has been set.

### GetRequestedTime

`func (o *DigitalWalletTokenResponse) GetRequestedTime() time.Time`

GetRequestedTime returns the RequestedTime field if non-nil, zero value otherwise.

### GetRequestedTimeOk

`func (o *DigitalWalletTokenResponse) GetRequestedTimeOk() (*time.Time, bool)`

GetRequestedTimeOk returns a tuple with the RequestedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestedTime

`func (o *DigitalWalletTokenResponse) SetRequestedTime(v time.Time)`

SetRequestedTime sets RequestedTime field to given value.

### HasRequestedTime

`func (o *DigitalWalletTokenResponse) HasRequestedTime() bool`

HasRequestedTime returns a boolean if a field has been set.

### GetApprovedTime

`func (o *DigitalWalletTokenResponse) GetApprovedTime() time.Time`

GetApprovedTime returns the ApprovedTime field if non-nil, zero value otherwise.

### GetApprovedTimeOk

`func (o *DigitalWalletTokenResponse) GetApprovedTimeOk() (*time.Time, bool)`

GetApprovedTimeOk returns a tuple with the ApprovedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApprovedTime

`func (o *DigitalWalletTokenResponse) SetApprovedTime(v time.Time)`

SetApprovedTime sets ApprovedTime field to given value.

### HasApprovedTime

`func (o *DigitalWalletTokenResponse) HasApprovedTime() bool`

HasApprovedTime returns a boolean if a field has been set.

### GetLastModifiedTime

`func (o *DigitalWalletTokenResponse) GetLastModifiedTime() time.Time`

GetLastModifiedTime returns the LastModifiedTime field if non-nil, zero value otherwise.

### GetLastModifiedTimeOk

`func (o *DigitalWalletTokenResponse) GetLastModifiedTimeOk() (*time.Time, bool)`

GetLastModifiedTimeOk returns a tuple with the LastModifiedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModifiedTime

`func (o *DigitalWalletTokenResponse) SetLastModifiedTime(v time.Time)`

SetLastModifiedTime sets LastModifiedTime field to given value.

### HasLastModifiedTime

`func (o *DigitalWalletTokenResponse) HasLastModifiedTime() bool`

HasLastModifiedTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


