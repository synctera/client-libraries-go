# Device

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Resource ID for device | [optional] 
**Ip** | Pointer to **string** | IP address | [optional] 
**Os** | Pointer to **string** | Operation system | [optional] 
**OsVersion** | Pointer to **string** | Operation system version | [optional] 
**ScreenWidth** | Pointer to **int32** | Device screen width | [optional] 
**ScreenHeight** | Pointer to **int32** | Device screen height | [optional] 

## Methods

### NewDevice

`func NewDevice() *Device`

NewDevice instantiates a new Device object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeviceWithDefaults

`func NewDeviceWithDefaults() *Device`

NewDeviceWithDefaults instantiates a new Device object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Device) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Device) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Device) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Device) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIp

`func (o *Device) GetIp() string`

GetIp returns the Ip field if non-nil, zero value otherwise.

### GetIpOk

`func (o *Device) GetIpOk() (*string, bool)`

GetIpOk returns a tuple with the Ip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp

`func (o *Device) SetIp(v string)`

SetIp sets Ip field to given value.

### HasIp

`func (o *Device) HasIp() bool`

HasIp returns a boolean if a field has been set.

### GetOs

`func (o *Device) GetOs() string`

GetOs returns the Os field if non-nil, zero value otherwise.

### GetOsOk

`func (o *Device) GetOsOk() (*string, bool)`

GetOsOk returns a tuple with the Os field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOs

`func (o *Device) SetOs(v string)`

SetOs sets Os field to given value.

### HasOs

`func (o *Device) HasOs() bool`

HasOs returns a boolean if a field has been set.

### GetOsVersion

`func (o *Device) GetOsVersion() string`

GetOsVersion returns the OsVersion field if non-nil, zero value otherwise.

### GetOsVersionOk

`func (o *Device) GetOsVersionOk() (*string, bool)`

GetOsVersionOk returns a tuple with the OsVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsVersion

`func (o *Device) SetOsVersion(v string)`

SetOsVersion sets OsVersion field to given value.

### HasOsVersion

`func (o *Device) HasOsVersion() bool`

HasOsVersion returns a boolean if a field has been set.

### GetScreenWidth

`func (o *Device) GetScreenWidth() int32`

GetScreenWidth returns the ScreenWidth field if non-nil, zero value otherwise.

### GetScreenWidthOk

`func (o *Device) GetScreenWidthOk() (*int32, bool)`

GetScreenWidthOk returns a tuple with the ScreenWidth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScreenWidth

`func (o *Device) SetScreenWidth(v int32)`

SetScreenWidth sets ScreenWidth field to given value.

### HasScreenWidth

`func (o *Device) HasScreenWidth() bool`

HasScreenWidth returns a boolean if a field has been set.

### GetScreenHeight

`func (o *Device) GetScreenHeight() int32`

GetScreenHeight returns the ScreenHeight field if non-nil, zero value otherwise.

### GetScreenHeightOk

`func (o *Device) GetScreenHeightOk() (*int32, bool)`

GetScreenHeightOk returns a tuple with the ScreenHeight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScreenHeight

`func (o *Device) SetScreenHeight(v int32)`

SetScreenHeight sets ScreenHeight field to given value.

### HasScreenHeight

`func (o *Device) HasScreenHeight() bool`

HasScreenHeight returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


