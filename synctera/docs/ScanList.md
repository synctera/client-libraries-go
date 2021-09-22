# ScanList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Scans** | [**[]Scan**](Scan.md) | Array of RDC scans | 
**NextPageToken** | Pointer to **string** | If returned, use the next_page_token to query for the next page of results. Not returned if there are no more rows. | [optional] 

## Methods

### NewScanList

`func NewScanList(scans []Scan, ) *ScanList`

NewScanList instantiates a new ScanList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewScanListWithDefaults

`func NewScanListWithDefaults() *ScanList`

NewScanListWithDefaults instantiates a new ScanList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetScans

`func (o *ScanList) GetScans() []Scan`

GetScans returns the Scans field if non-nil, zero value otherwise.

### GetScansOk

`func (o *ScanList) GetScansOk() (*[]Scan, bool)`

GetScansOk returns a tuple with the Scans field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScans

`func (o *ScanList) SetScans(v []Scan)`

SetScans sets Scans field to given value.


### GetNextPageToken

`func (o *ScanList) GetNextPageToken() string`

GetNextPageToken returns the NextPageToken field if non-nil, zero value otherwise.

### GetNextPageTokenOk

`func (o *ScanList) GetNextPageTokenOk() (*string, bool)`

GetNextPageTokenOk returns a tuple with the NextPageToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextPageToken

`func (o *ScanList) SetNextPageToken(v string)`

SetNextPageToken sets NextPageToken field to given value.

### HasNextPageToken

`func (o *ScanList) HasNextPageToken() bool`

HasNextPageToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


