/*
 * Synctera API
 *
 * <h2>Let's build something great.</h2><p>Welcome to the official reference documentation for Synctera APIs. Our APIs are the best way to automate your company's banking needs and are designed to be easy to understand and implement.</p><p>We're continuously growing this library and what you see here is just the start, but if you need something specific or have a question, <a class='text-blue-600' href='https://synctera.com/contact' target='_blank' rel='noreferrer'>contact us</a>.</p> 
 *
 * API version: 0.5.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package synctera

import (
	"encoding/json"
)

// Device struct for Device
type Device struct {
	// Resource ID for device
	Id *string `json:"id,omitempty"`
	// IP address
	Ip *string `json:"ip,omitempty"`
	// Operation system
	Os *string `json:"os,omitempty"`
	// Operation system version
	OsVersion *string `json:"os_version,omitempty"`
	// Device screen width
	ScreenWidth *int32 `json:"screen_width,omitempty"`
	// Device screen height
	ScreenHeight *int32 `json:"screen_height,omitempty"`
}

// NewDevice instantiates a new Device object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDevice() *Device {
	this := Device{}
	return &this
}

// NewDeviceWithDefaults instantiates a new Device object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeviceWithDefaults() *Device {
	this := Device{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Device) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Device) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Device) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *Device) SetId(v string) {
	o.Id = &v
}

// GetIp returns the Ip field value if set, zero value otherwise.
func (o *Device) GetIp() string {
	if o == nil || o.Ip == nil {
		var ret string
		return ret
	}
	return *o.Ip
}

// GetIpOk returns a tuple with the Ip field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Device) GetIpOk() (*string, bool) {
	if o == nil || o.Ip == nil {
		return nil, false
	}
	return o.Ip, true
}

// HasIp returns a boolean if a field has been set.
func (o *Device) HasIp() bool {
	if o != nil && o.Ip != nil {
		return true
	}

	return false
}

// SetIp gets a reference to the given string and assigns it to the Ip field.
func (o *Device) SetIp(v string) {
	o.Ip = &v
}

// GetOs returns the Os field value if set, zero value otherwise.
func (o *Device) GetOs() string {
	if o == nil || o.Os == nil {
		var ret string
		return ret
	}
	return *o.Os
}

// GetOsOk returns a tuple with the Os field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Device) GetOsOk() (*string, bool) {
	if o == nil || o.Os == nil {
		return nil, false
	}
	return o.Os, true
}

// HasOs returns a boolean if a field has been set.
func (o *Device) HasOs() bool {
	if o != nil && o.Os != nil {
		return true
	}

	return false
}

// SetOs gets a reference to the given string and assigns it to the Os field.
func (o *Device) SetOs(v string) {
	o.Os = &v
}

// GetOsVersion returns the OsVersion field value if set, zero value otherwise.
func (o *Device) GetOsVersion() string {
	if o == nil || o.OsVersion == nil {
		var ret string
		return ret
	}
	return *o.OsVersion
}

// GetOsVersionOk returns a tuple with the OsVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Device) GetOsVersionOk() (*string, bool) {
	if o == nil || o.OsVersion == nil {
		return nil, false
	}
	return o.OsVersion, true
}

// HasOsVersion returns a boolean if a field has been set.
func (o *Device) HasOsVersion() bool {
	if o != nil && o.OsVersion != nil {
		return true
	}

	return false
}

// SetOsVersion gets a reference to the given string and assigns it to the OsVersion field.
func (o *Device) SetOsVersion(v string) {
	o.OsVersion = &v
}

// GetScreenWidth returns the ScreenWidth field value if set, zero value otherwise.
func (o *Device) GetScreenWidth() int32 {
	if o == nil || o.ScreenWidth == nil {
		var ret int32
		return ret
	}
	return *o.ScreenWidth
}

// GetScreenWidthOk returns a tuple with the ScreenWidth field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Device) GetScreenWidthOk() (*int32, bool) {
	if o == nil || o.ScreenWidth == nil {
		return nil, false
	}
	return o.ScreenWidth, true
}

// HasScreenWidth returns a boolean if a field has been set.
func (o *Device) HasScreenWidth() bool {
	if o != nil && o.ScreenWidth != nil {
		return true
	}

	return false
}

// SetScreenWidth gets a reference to the given int32 and assigns it to the ScreenWidth field.
func (o *Device) SetScreenWidth(v int32) {
	o.ScreenWidth = &v
}

// GetScreenHeight returns the ScreenHeight field value if set, zero value otherwise.
func (o *Device) GetScreenHeight() int32 {
	if o == nil || o.ScreenHeight == nil {
		var ret int32
		return ret
	}
	return *o.ScreenHeight
}

// GetScreenHeightOk returns a tuple with the ScreenHeight field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Device) GetScreenHeightOk() (*int32, bool) {
	if o == nil || o.ScreenHeight == nil {
		return nil, false
	}
	return o.ScreenHeight, true
}

// HasScreenHeight returns a boolean if a field has been set.
func (o *Device) HasScreenHeight() bool {
	if o != nil && o.ScreenHeight != nil {
		return true
	}

	return false
}

// SetScreenHeight gets a reference to the given int32 and assigns it to the ScreenHeight field.
func (o *Device) SetScreenHeight(v int32) {
	o.ScreenHeight = &v
}

func (o Device) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Ip != nil {
		toSerialize["ip"] = o.Ip
	}
	if o.Os != nil {
		toSerialize["os"] = o.Os
	}
	if o.OsVersion != nil {
		toSerialize["os_version"] = o.OsVersion
	}
	if o.ScreenWidth != nil {
		toSerialize["screen_width"] = o.ScreenWidth
	}
	if o.ScreenHeight != nil {
		toSerialize["screen_height"] = o.ScreenHeight
	}
	return json.Marshal(toSerialize)
}

type NullableDevice struct {
	value *Device
	isSet bool
}

func (v NullableDevice) Get() *Device {
	return v.value
}

func (v *NullableDevice) Set(val *Device) {
	v.value = val
	v.isSet = true
}

func (v NullableDevice) IsSet() bool {
	return v.isSet
}

func (v *NullableDevice) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDevice(val *Device) *NullableDevice {
	return &NullableDevice{value: val, isSet: true}
}

func (v NullableDevice) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDevice) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

