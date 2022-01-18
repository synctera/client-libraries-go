# VendorJson

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ContentType** | **string** | Describes the content-type encoding received from the vendor. | 
**Details** | Pointer to [**[]Detail**](Detail.md) | Array of vendor specific information. | [optional] [readonly] 
**Json** | **map[string]interface{}** | Data representation in JSON. | 
**Vendor** | **string** | Name of the vendor used. | 

## Methods

### NewVendorJson

`func NewVendorJson(contentType string, json map[string]interface{}, vendor string, ) *VendorJson`

NewVendorJson instantiates a new VendorJson object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVendorJsonWithDefaults

`func NewVendorJsonWithDefaults() *VendorJson`

NewVendorJsonWithDefaults instantiates a new VendorJson object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContentType

`func (o *VendorJson) GetContentType() string`

GetContentType returns the ContentType field if non-nil, zero value otherwise.

### GetContentTypeOk

`func (o *VendorJson) GetContentTypeOk() (*string, bool)`

GetContentTypeOk returns a tuple with the ContentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentType

`func (o *VendorJson) SetContentType(v string)`

SetContentType sets ContentType field to given value.


### GetDetails

`func (o *VendorJson) GetDetails() []Detail`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *VendorJson) GetDetailsOk() (*[]Detail, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetails

`func (o *VendorJson) SetDetails(v []Detail)`

SetDetails sets Details field to given value.

### HasDetails

`func (o *VendorJson) HasDetails() bool`

HasDetails returns a boolean if a field has been set.

### GetJson

`func (o *VendorJson) GetJson() map[string]interface{}`

GetJson returns the Json field if non-nil, zero value otherwise.

### GetJsonOk

`func (o *VendorJson) GetJsonOk() (*map[string]interface{}, bool)`

GetJsonOk returns a tuple with the Json field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJson

`func (o *VendorJson) SetJson(v map[string]interface{})`

SetJson sets Json field to given value.


### GetVendor

`func (o *VendorJson) GetVendor() string`

GetVendor returns the Vendor field if non-nil, zero value otherwise.

### GetVendorOk

`func (o *VendorJson) GetVendorOk() (*string, bool)`

GetVendorOk returns a tuple with the Vendor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendor

`func (o *VendorJson) SetVendor(v string)`

SetVendor sets Vendor field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


