# VendorInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ContentType** | **string** | Describes the content-type encoding received from the vendor. | 
**Details** | [**[]Detail**](Detail.md) | Array of vendor specific information. | 
**Json** | **map[string]interface{}** | Data representation in JSON. | 
**Vendor** | **string** | Name of the vendor used. | 
**Xml** | **string** | Data representaion in XML. | 

## Methods

### NewVendorInfo

`func NewVendorInfo(contentType string, details []Detail, json map[string]interface{}, vendor string, xml string, ) *VendorInfo`

NewVendorInfo instantiates a new VendorInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVendorInfoWithDefaults

`func NewVendorInfoWithDefaults() *VendorInfo`

NewVendorInfoWithDefaults instantiates a new VendorInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContentType

`func (o *VendorInfo) GetContentType() string`

GetContentType returns the ContentType field if non-nil, zero value otherwise.

### GetContentTypeOk

`func (o *VendorInfo) GetContentTypeOk() (*string, bool)`

GetContentTypeOk returns a tuple with the ContentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentType

`func (o *VendorInfo) SetContentType(v string)`

SetContentType sets ContentType field to given value.


### GetDetails

`func (o *VendorInfo) GetDetails() []Detail`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *VendorInfo) GetDetailsOk() (*[]Detail, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetails

`func (o *VendorInfo) SetDetails(v []Detail)`

SetDetails sets Details field to given value.


### GetJson

`func (o *VendorInfo) GetJson() map[string]interface{}`

GetJson returns the Json field if non-nil, zero value otherwise.

### GetJsonOk

`func (o *VendorInfo) GetJsonOk() (*map[string]interface{}, bool)`

GetJsonOk returns a tuple with the Json field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJson

`func (o *VendorInfo) SetJson(v map[string]interface{})`

SetJson sets Json field to given value.


### GetVendor

`func (o *VendorInfo) GetVendor() string`

GetVendor returns the Vendor field if non-nil, zero value otherwise.

### GetVendorOk

`func (o *VendorInfo) GetVendorOk() (*string, bool)`

GetVendorOk returns a tuple with the Vendor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendor

`func (o *VendorInfo) SetVendor(v string)`

SetVendor sets Vendor field to given value.


### GetXml

`func (o *VendorInfo) GetXml() string`

GetXml returns the Xml field if non-nil, zero value otherwise.

### GetXmlOk

`func (o *VendorInfo) GetXmlOk() (*string, bool)`

GetXmlOk returns a tuple with the Xml field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetXml

`func (o *VendorInfo) SetXml(v string)`

SetXml sets Xml field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


