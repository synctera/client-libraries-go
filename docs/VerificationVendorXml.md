# VerificationVendorXml

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Vendor** | **string** | Name of the vendor used. | 
**ContentType** | **string** | Describes the content-type encoding received from the vendor. | 
**Xml** | **string** | Data representaion in XML. | 
**Details** | Pointer to [**[]VerificationVendorInfoDetail**](VerificationVendorInfoDetail.md) | Array of vendor specific information. | [optional] [readonly] 

## Methods

### NewVerificationVendorXml

`func NewVerificationVendorXml(vendor string, contentType string, xml string, ) *VerificationVendorXml`

NewVerificationVendorXml instantiates a new VerificationVendorXml object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVerificationVendorXmlWithDefaults

`func NewVerificationVendorXmlWithDefaults() *VerificationVendorXml`

NewVerificationVendorXmlWithDefaults instantiates a new VerificationVendorXml object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVendor

`func (o *VerificationVendorXml) GetVendor() string`

GetVendor returns the Vendor field if non-nil, zero value otherwise.

### GetVendorOk

`func (o *VerificationVendorXml) GetVendorOk() (*string, bool)`

GetVendorOk returns a tuple with the Vendor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendor

`func (o *VerificationVendorXml) SetVendor(v string)`

SetVendor sets Vendor field to given value.


### GetContentType

`func (o *VerificationVendorXml) GetContentType() string`

GetContentType returns the ContentType field if non-nil, zero value otherwise.

### GetContentTypeOk

`func (o *VerificationVendorXml) GetContentTypeOk() (*string, bool)`

GetContentTypeOk returns a tuple with the ContentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentType

`func (o *VerificationVendorXml) SetContentType(v string)`

SetContentType sets ContentType field to given value.


### GetXml

`func (o *VerificationVendorXml) GetXml() string`

GetXml returns the Xml field if non-nil, zero value otherwise.

### GetXmlOk

`func (o *VerificationVendorXml) GetXmlOk() (*string, bool)`

GetXmlOk returns a tuple with the Xml field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetXml

`func (o *VerificationVendorXml) SetXml(v string)`

SetXml sets Xml field to given value.


### GetDetails

`func (o *VerificationVendorXml) GetDetails() []VerificationVendorInfoDetail`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *VerificationVendorXml) GetDetailsOk() (*[]VerificationVendorInfoDetail, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetails

`func (o *VerificationVendorXml) SetDetails(v []VerificationVendorInfoDetail)`

SetDetails sets Details field to given value.

### HasDetails

`func (o *VerificationVendorXml) HasDetails() bool`

HasDetails returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


