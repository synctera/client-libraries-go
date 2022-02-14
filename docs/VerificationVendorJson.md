# VerificationVendorJson

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ContentType** | **string** | Describes the content-type encoding received from the vendor. | 
**Details** | Pointer to [**[]VerificationVendorInfoDetail**](VerificationVendorInfoDetail.md) | Array of vendor specific information. | [optional] [readonly] 
**Json** | **map[string]interface{}** | Data representation in JSON. | 
**Vendor** | **string** | Name of the vendor used. | 

## Methods

### NewVerificationVendorJson

`func NewVerificationVendorJson(contentType string, json map[string]interface{}, vendor string, ) *VerificationVendorJson`

NewVerificationVendorJson instantiates a new VerificationVendorJson object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVerificationVendorJsonWithDefaults

`func NewVerificationVendorJsonWithDefaults() *VerificationVendorJson`

NewVerificationVendorJsonWithDefaults instantiates a new VerificationVendorJson object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContentType

`func (o *VerificationVendorJson) GetContentType() string`

GetContentType returns the ContentType field if non-nil, zero value otherwise.

### GetContentTypeOk

`func (o *VerificationVendorJson) GetContentTypeOk() (*string, bool)`

GetContentTypeOk returns a tuple with the ContentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentType

`func (o *VerificationVendorJson) SetContentType(v string)`

SetContentType sets ContentType field to given value.


### GetDetails

`func (o *VerificationVendorJson) GetDetails() []VerificationVendorInfoDetail`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *VerificationVendorJson) GetDetailsOk() (*[]VerificationVendorInfoDetail, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetails

`func (o *VerificationVendorJson) SetDetails(v []VerificationVendorInfoDetail)`

SetDetails sets Details field to given value.

### HasDetails

`func (o *VerificationVendorJson) HasDetails() bool`

HasDetails returns a boolean if a field has been set.

### GetJson

`func (o *VerificationVendorJson) GetJson() map[string]interface{}`

GetJson returns the Json field if non-nil, zero value otherwise.

### GetJsonOk

`func (o *VerificationVendorJson) GetJsonOk() (*map[string]interface{}, bool)`

GetJsonOk returns a tuple with the Json field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJson

`func (o *VerificationVendorJson) SetJson(v map[string]interface{})`

SetJson sets Json field to given value.


### GetVendor

`func (o *VerificationVendorJson) GetVendor() string`

GetVendor returns the Vendor field if non-nil, zero value otherwise.

### GetVendorOk

`func (o *VerificationVendorJson) GetVendorOk() (*string, bool)`

GetVendorOk returns a tuple with the Vendor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendor

`func (o *VerificationVendorJson) SetVendor(v string)`

SetVendor sets Vendor field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


