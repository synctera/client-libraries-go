# UpdateGatewayRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Active** | Pointer to **bool** | Current status of the Authorization gateway | [optional] 
**CardProducts** | Pointer to **[]string** | List of Card Product unique identifiers that will utilize the Gateway | [optional] 
**CustomHeaders** | Pointer to **map[string]string** | Custom Headers of the Authorization gateway | [optional] 
**Url** | Pointer to **string** | URL of the Authorization gateway | [optional] 

## Methods

### NewUpdateGatewayRequest

`func NewUpdateGatewayRequest() *UpdateGatewayRequest`

NewUpdateGatewayRequest instantiates a new UpdateGatewayRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateGatewayRequestWithDefaults

`func NewUpdateGatewayRequestWithDefaults() *UpdateGatewayRequest`

NewUpdateGatewayRequestWithDefaults instantiates a new UpdateGatewayRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActive

`func (o *UpdateGatewayRequest) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *UpdateGatewayRequest) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *UpdateGatewayRequest) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *UpdateGatewayRequest) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetCardProducts

`func (o *UpdateGatewayRequest) GetCardProducts() []string`

GetCardProducts returns the CardProducts field if non-nil, zero value otherwise.

### GetCardProductsOk

`func (o *UpdateGatewayRequest) GetCardProductsOk() (*[]string, bool)`

GetCardProductsOk returns a tuple with the CardProducts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardProducts

`func (o *UpdateGatewayRequest) SetCardProducts(v []string)`

SetCardProducts sets CardProducts field to given value.

### HasCardProducts

`func (o *UpdateGatewayRequest) HasCardProducts() bool`

HasCardProducts returns a boolean if a field has been set.

### GetCustomHeaders

`func (o *UpdateGatewayRequest) GetCustomHeaders() map[string]string`

GetCustomHeaders returns the CustomHeaders field if non-nil, zero value otherwise.

### GetCustomHeadersOk

`func (o *UpdateGatewayRequest) GetCustomHeadersOk() (*map[string]string, bool)`

GetCustomHeadersOk returns a tuple with the CustomHeaders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomHeaders

`func (o *UpdateGatewayRequest) SetCustomHeaders(v map[string]string)`

SetCustomHeaders sets CustomHeaders field to given value.

### HasCustomHeaders

`func (o *UpdateGatewayRequest) HasCustomHeaders() bool`

HasCustomHeaders returns a boolean if a field has been set.

### GetUrl

`func (o *UpdateGatewayRequest) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *UpdateGatewayRequest) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *UpdateGatewayRequest) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *UpdateGatewayRequest) HasUrl() bool`

HasUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


