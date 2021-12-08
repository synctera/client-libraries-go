# VendorJson1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ContentType** | **string** | Describes the content-type encoding received from the vendor | 
**Json** | **map[string]interface{}** | Data representation in JSON | 
**Vendor** | **string** |  | 

## Methods

### NewVendorJson1

`func NewVendorJson1(contentType string, json map[string]interface{}, vendor string, ) *VendorJson1`

NewVendorJson1 instantiates a new VendorJson1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVendorJson1WithDefaults

`func NewVendorJson1WithDefaults() *VendorJson1`

NewVendorJson1WithDefaults instantiates a new VendorJson1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContentType

`func (o *VendorJson1) GetContentType() string`

GetContentType returns the ContentType field if non-nil, zero value otherwise.

### GetContentTypeOk

`func (o *VendorJson1) GetContentTypeOk() (*string, bool)`

GetContentTypeOk returns a tuple with the ContentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentType

`func (o *VendorJson1) SetContentType(v string)`

SetContentType sets ContentType field to given value.


### GetJson

`func (o *VendorJson1) GetJson() map[string]interface{}`

GetJson returns the Json field if non-nil, zero value otherwise.

### GetJsonOk

`func (o *VendorJson1) GetJsonOk() (*map[string]interface{}, bool)`

GetJsonOk returns a tuple with the Json field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJson

`func (o *VendorJson1) SetJson(v map[string]interface{})`

SetJson sets Json field to given value.


### GetVendor

`func (o *VendorJson1) GetVendor() string`

GetVendor returns the Vendor field if non-nil, zero value otherwise.

### GetVendorOk

`func (o *VendorJson1) GetVendorOk() (*string, bool)`

GetVendorOk returns a tuple with the Vendor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendor

`func (o *VendorJson1) SetVendor(v string)`

SetVendor sets Vendor field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


