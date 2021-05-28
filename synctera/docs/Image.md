# Image

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | RDC image id | [readonly] 
**MediaType** | [**RdcMediaType**](RdcMediaType.md) |  | 
**ByteData** | **string** | Base64url encoded image | 
**DateUploaded** | **string** | Date the image was uploaded, in RFC 3339 format | [readonly] 

## Methods

### NewImage

`func NewImage(id string, mediaType RdcMediaType, byteData string, dateUploaded string, ) *Image`

NewImage instantiates a new Image object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewImageWithDefaults

`func NewImageWithDefaults() *Image`

NewImageWithDefaults instantiates a new Image object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Image) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Image) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Image) SetId(v string)`

SetId sets Id field to given value.


### GetMediaType

`func (o *Image) GetMediaType() RdcMediaType`

GetMediaType returns the MediaType field if non-nil, zero value otherwise.

### GetMediaTypeOk

`func (o *Image) GetMediaTypeOk() (*RdcMediaType, bool)`

GetMediaTypeOk returns a tuple with the MediaType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMediaType

`func (o *Image) SetMediaType(v RdcMediaType)`

SetMediaType sets MediaType field to given value.


### GetByteData

`func (o *Image) GetByteData() string`

GetByteData returns the ByteData field if non-nil, zero value otherwise.

### GetByteDataOk

`func (o *Image) GetByteDataOk() (*string, bool)`

GetByteDataOk returns a tuple with the ByteData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetByteData

`func (o *Image) SetByteData(v string)`

SetByteData sets ByteData field to given value.


### GetDateUploaded

`func (o *Image) GetDateUploaded() string`

GetDateUploaded returns the DateUploaded field if non-nil, zero value otherwise.

### GetDateUploadedOk

`func (o *Image) GetDateUploadedOk() (*string, bool)`

GetDateUploadedOk returns a tuple with the DateUploaded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateUploaded

`func (o *Image) SetDateUploaded(v string)`

SetDateUploaded sets DateUploaded field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


