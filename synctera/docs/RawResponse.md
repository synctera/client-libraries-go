# RawResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Provider** | Pointer to [**ProviderType**](ProviderType.md) |  | [optional] 
**RawData** | Pointer to **string** |  | [optional] 

## Methods

### NewRawResponse

`func NewRawResponse() *RawResponse`

NewRawResponse instantiates a new RawResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRawResponseWithDefaults

`func NewRawResponseWithDefaults() *RawResponse`

NewRawResponseWithDefaults instantiates a new RawResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProvider

`func (o *RawResponse) GetProvider() ProviderType`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *RawResponse) GetProviderOk() (*ProviderType, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *RawResponse) SetProvider(v ProviderType)`

SetProvider sets Provider field to given value.

### HasProvider

`func (o *RawResponse) HasProvider() bool`

HasProvider returns a boolean if a field has been set.

### GetRawData

`func (o *RawResponse) GetRawData() string`

GetRawData returns the RawData field if non-nil, zero value otherwise.

### GetRawDataOk

`func (o *RawResponse) GetRawDataOk() (*string, bool)`

GetRawDataOk returns a tuple with the RawData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRawData

`func (o *RawResponse) SetRawData(v string)`

SetRawData sets RawData field to given value.

### HasRawData

`func (o *RawResponse) HasRawData() bool`

HasRawData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


