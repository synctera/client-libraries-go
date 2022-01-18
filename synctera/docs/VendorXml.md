# VendorXml

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ContentType** | **string** | Describes the content-type encoding received from the vendor. | 
**Details** | Pointer to [**[]Detail**](Detail.md) | Array of vendor specific information. | [optional] [readonly] 
**Vendor** | **string** | Name of the vendor used. | 
**Xml** | **string** | Data representaion in XML. | 

## Methods

### NewVendorXml

`func NewVendorXml(contentType string, vendor string, xml string, ) *VendorXml`

NewVendorXml instantiates a new VendorXml object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVendorXmlWithDefaults

`func NewVendorXmlWithDefaults() *VendorXml`

NewVendorXmlWithDefaults instantiates a new VendorXml object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContentType

`func (o *VendorXml) GetContentType() string`

GetContentType returns the ContentType field if non-nil, zero value otherwise.

### GetContentTypeOk

`func (o *VendorXml) GetContentTypeOk() (*string, bool)`

GetContentTypeOk returns a tuple with the ContentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentType

`func (o *VendorXml) SetContentType(v string)`

SetContentType sets ContentType field to given value.


### GetDetails

`func (o *VendorXml) GetDetails() []Detail`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *VendorXml) GetDetailsOk() (*[]Detail, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetails

`func (o *VendorXml) SetDetails(v []Detail)`

SetDetails sets Details field to given value.

### HasDetails

`func (o *VendorXml) HasDetails() bool`

HasDetails returns a boolean if a field has been set.

### GetVendor

`func (o *VendorXml) GetVendor() string`

GetVendor returns the Vendor field if non-nil, zero value otherwise.

### GetVendorOk

`func (o *VendorXml) GetVendorOk() (*string, bool)`

GetVendorOk returns a tuple with the Vendor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendor

`func (o *VendorXml) SetVendor(v string)`

SetVendor sets Vendor field to given value.


### GetXml

`func (o *VendorXml) GetXml() string`

GetXml returns the Xml field if non-nil, zero value otherwise.

### GetXmlOk

`func (o *VendorXml) GetXmlOk() (*string, bool)`

GetXmlOk returns a tuple with the Xml field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetXml

`func (o *VendorXml) SetXml(v string)`

SetXml sets Xml field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


