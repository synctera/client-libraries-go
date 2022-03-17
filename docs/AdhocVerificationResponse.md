# AdhocVerificationResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Unique ID for this verification result. | 
**MatchingWatchlists** | **[]string** | list of watchlists that the subject of the request matched  | 
**Result** | [**VerificationResult**](VerificationResult.md) |  | 
**VendorInfo** | Pointer to [**VendorInfo**](VendorInfo.md) |  | [optional] 

## Methods

### NewAdhocVerificationResponse

`func NewAdhocVerificationResponse(id string, matchingWatchlists []string, result VerificationResult, ) *AdhocVerificationResponse`

NewAdhocVerificationResponse instantiates a new AdhocVerificationResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdhocVerificationResponseWithDefaults

`func NewAdhocVerificationResponseWithDefaults() *AdhocVerificationResponse`

NewAdhocVerificationResponseWithDefaults instantiates a new AdhocVerificationResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AdhocVerificationResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AdhocVerificationResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AdhocVerificationResponse) SetId(v string)`

SetId sets Id field to given value.


### GetMatchingWatchlists

`func (o *AdhocVerificationResponse) GetMatchingWatchlists() []string`

GetMatchingWatchlists returns the MatchingWatchlists field if non-nil, zero value otherwise.

### GetMatchingWatchlistsOk

`func (o *AdhocVerificationResponse) GetMatchingWatchlistsOk() (*[]string, bool)`

GetMatchingWatchlistsOk returns a tuple with the MatchingWatchlists field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatchingWatchlists

`func (o *AdhocVerificationResponse) SetMatchingWatchlists(v []string)`

SetMatchingWatchlists sets MatchingWatchlists field to given value.


### GetResult

`func (o *AdhocVerificationResponse) GetResult() VerificationResult`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *AdhocVerificationResponse) GetResultOk() (*VerificationResult, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *AdhocVerificationResponse) SetResult(v VerificationResult)`

SetResult sets Result field to given value.


### GetVendorInfo

`func (o *AdhocVerificationResponse) GetVendorInfo() VendorInfo`

GetVendorInfo returns the VendorInfo field if non-nil, zero value otherwise.

### GetVendorInfoOk

`func (o *AdhocVerificationResponse) GetVendorInfoOk() (*VendorInfo, bool)`

GetVendorInfoOk returns a tuple with the VendorInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendorInfo

`func (o *AdhocVerificationResponse) SetVendorInfo(v VendorInfo)`

SetVendorInfo sets VendorInfo field to given value.

### HasVendorInfo

`func (o *AdhocVerificationResponse) HasVendorInfo() bool`

HasVendorInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


