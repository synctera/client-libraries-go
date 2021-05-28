# Document

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Document ID | [optional] [readonly] 
**MediaType** | [**KycMediaType**](KycMediaType.md) |  | 
**DocumentType** | [**DocumentType**](DocumentType.md) |  | 
**ByteData** | Pointer to **string** | Base64url encoded image | [optional] 

## Methods

### NewDocument

`func NewDocument(mediaType KycMediaType, documentType DocumentType, ) *Document`

NewDocument instantiates a new Document object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDocumentWithDefaults

`func NewDocumentWithDefaults() *Document`

NewDocumentWithDefaults instantiates a new Document object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Document) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Document) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Document) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Document) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMediaType

`func (o *Document) GetMediaType() KycMediaType`

GetMediaType returns the MediaType field if non-nil, zero value otherwise.

### GetMediaTypeOk

`func (o *Document) GetMediaTypeOk() (*KycMediaType, bool)`

GetMediaTypeOk returns a tuple with the MediaType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMediaType

`func (o *Document) SetMediaType(v KycMediaType)`

SetMediaType sets MediaType field to given value.


### GetDocumentType

`func (o *Document) GetDocumentType() DocumentType`

GetDocumentType returns the DocumentType field if non-nil, zero value otherwise.

### GetDocumentTypeOk

`func (o *Document) GetDocumentTypeOk() (*DocumentType, bool)`

GetDocumentTypeOk returns a tuple with the DocumentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentType

`func (o *Document) SetDocumentType(v DocumentType)`

SetDocumentType sets DocumentType field to given value.


### GetByteData

`func (o *Document) GetByteData() string`

GetByteData returns the ByteData field if non-nil, zero value otherwise.

### GetByteDataOk

`func (o *Document) GetByteDataOk() (*string, bool)`

GetByteDataOk returns a tuple with the ByteData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetByteData

`func (o *Document) SetByteData(v string)`

SetByteData sets ByteData field to given value.

### HasByteData

`func (o *Document) HasByteData() bool`

HasByteData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


