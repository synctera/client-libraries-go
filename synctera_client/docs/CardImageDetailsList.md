# CardImageDetailsList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Images** | [**[]CardImageDetails**](CardImageDetails.md) | Array of image details | 
**NextPageToken** | Pointer to **string** | If returned, use the next_page_token to query for the next page of results. Not returned if there are no more rows. | [optional] 

## Methods

### NewCardImageDetailsList

`func NewCardImageDetailsList(images []CardImageDetails, ) *CardImageDetailsList`

NewCardImageDetailsList instantiates a new CardImageDetailsList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCardImageDetailsListWithDefaults

`func NewCardImageDetailsListWithDefaults() *CardImageDetailsList`

NewCardImageDetailsListWithDefaults instantiates a new CardImageDetailsList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetImages

`func (o *CardImageDetailsList) GetImages() []CardImageDetails`

GetImages returns the Images field if non-nil, zero value otherwise.

### GetImagesOk

`func (o *CardImageDetailsList) GetImagesOk() (*[]CardImageDetails, bool)`

GetImagesOk returns a tuple with the Images field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImages

`func (o *CardImageDetailsList) SetImages(v []CardImageDetails)`

SetImages sets Images field to given value.


### GetNextPageToken

`func (o *CardImageDetailsList) GetNextPageToken() string`

GetNextPageToken returns the NextPageToken field if non-nil, zero value otherwise.

### GetNextPageTokenOk

`func (o *CardImageDetailsList) GetNextPageTokenOk() (*string, bool)`

GetNextPageTokenOk returns a tuple with the NextPageToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextPageToken

`func (o *CardImageDetailsList) SetNextPageToken(v string)`

SetNextPageToken sets NextPageToken field to given value.

### HasNextPageToken

`func (o *CardImageDetailsList) HasNextPageToken() bool`

HasNextPageToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


