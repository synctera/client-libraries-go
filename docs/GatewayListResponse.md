# GatewayListResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Gateways** | Pointer to [**[]GatewayResponse**](GatewayResponse.md) | Array of Authorization gateway configuration | [optional] 
**NextPageToken** | Pointer to **string** | If returned, use the next_page_token to query for the next page of results. Not returned if there are no more rows. | [optional] 

## Methods

### NewGatewayListResponse

`func NewGatewayListResponse() *GatewayListResponse`

NewGatewayListResponse instantiates a new GatewayListResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGatewayListResponseWithDefaults

`func NewGatewayListResponseWithDefaults() *GatewayListResponse`

NewGatewayListResponseWithDefaults instantiates a new GatewayListResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGateways

`func (o *GatewayListResponse) GetGateways() []GatewayResponse`

GetGateways returns the Gateways field if non-nil, zero value otherwise.

### GetGatewaysOk

`func (o *GatewayListResponse) GetGatewaysOk() (*[]GatewayResponse, bool)`

GetGatewaysOk returns a tuple with the Gateways field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGateways

`func (o *GatewayListResponse) SetGateways(v []GatewayResponse)`

SetGateways sets Gateways field to given value.

### HasGateways

`func (o *GatewayListResponse) HasGateways() bool`

HasGateways returns a boolean if a field has been set.

### GetNextPageToken

`func (o *GatewayListResponse) GetNextPageToken() string`

GetNextPageToken returns the NextPageToken field if non-nil, zero value otherwise.

### GetNextPageTokenOk

`func (o *GatewayListResponse) GetNextPageTokenOk() (*string, bool)`

GetNextPageTokenOk returns a tuple with the NextPageToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextPageToken

`func (o *GatewayListResponse) SetNextPageToken(v string)`

SetNextPageToken sets NextPageToken field to given value.

### HasNextPageToken

`func (o *GatewayListResponse) HasNextPageToken() bool`

HasNextPageToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


