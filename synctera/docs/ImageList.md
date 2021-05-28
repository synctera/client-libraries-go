# ImageList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Images** | [**[]Image**](Image.md) | Array of RDC images | 
**NextPageToken** | Pointer to **string** | If returned, use the next_page_token to query for the next page of results. | [optional] 

## Methods

### NewImageList

`func NewImageList(images []Image, ) *ImageList`

NewImageList instantiates a new ImageList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewImageListWithDefaults

`func NewImageListWithDefaults() *ImageList`

NewImageListWithDefaults instantiates a new ImageList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetImages

`func (o *ImageList) GetImages() []Image`

GetImages returns the Images field if non-nil, zero value otherwise.

### GetImagesOk

`func (o *ImageList) GetImagesOk() (*[]Image, bool)`

GetImagesOk returns a tuple with the Images field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImages

`func (o *ImageList) SetImages(v []Image)`

SetImages sets Images field to given value.


### GetNextPageToken

`func (o *ImageList) GetNextPageToken() string`

GetNextPageToken returns the NextPageToken field if non-nil, zero value otherwise.

### GetNextPageTokenOk

`func (o *ImageList) GetNextPageTokenOk() (*string, bool)`

GetNextPageTokenOk returns a tuple with the NextPageToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextPageToken

`func (o *ImageList) SetNextPageToken(v string)`

SetNextPageToken sets NextPageToken field to given value.

### HasNextPageToken

`func (o *ImageList) HasNextPageToken() bool`

HasNextPageToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


