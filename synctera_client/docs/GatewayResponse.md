# GatewayResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Active** | **bool** | Current status of the Authorization gateway | 
**CardProducts** | **[]string** | List of Card Product unique identifiers that will utilize the Gateway | 
**CreationTime** | **time.Time** | The timestamp representing when the gateway config request was made | [readonly] 
**CustomHeaders** | Pointer to **map[string]string** | Custom Headers of the Authorization gateway | [optional] 
**Id** | **string** | Gateway ID | 
**LastModifiedTime** | **time.Time** | The timestamp representing when the gateway config was last modified at | [readonly] 
**Url** | **string** | URL of the Authorization gateway | 

## Methods

### NewGatewayResponse

`func NewGatewayResponse(active bool, cardProducts []string, creationTime time.Time, id string, lastModifiedTime time.Time, url string, ) *GatewayResponse`

NewGatewayResponse instantiates a new GatewayResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGatewayResponseWithDefaults

`func NewGatewayResponseWithDefaults() *GatewayResponse`

NewGatewayResponseWithDefaults instantiates a new GatewayResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActive

`func (o *GatewayResponse) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *GatewayResponse) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *GatewayResponse) SetActive(v bool)`

SetActive sets Active field to given value.


### GetCardProducts

`func (o *GatewayResponse) GetCardProducts() []string`

GetCardProducts returns the CardProducts field if non-nil, zero value otherwise.

### GetCardProductsOk

`func (o *GatewayResponse) GetCardProductsOk() (*[]string, bool)`

GetCardProductsOk returns a tuple with the CardProducts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardProducts

`func (o *GatewayResponse) SetCardProducts(v []string)`

SetCardProducts sets CardProducts field to given value.


### GetCreationTime

`func (o *GatewayResponse) GetCreationTime() time.Time`

GetCreationTime returns the CreationTime field if non-nil, zero value otherwise.

### GetCreationTimeOk

`func (o *GatewayResponse) GetCreationTimeOk() (*time.Time, bool)`

GetCreationTimeOk returns a tuple with the CreationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationTime

`func (o *GatewayResponse) SetCreationTime(v time.Time)`

SetCreationTime sets CreationTime field to given value.


### GetCustomHeaders

`func (o *GatewayResponse) GetCustomHeaders() map[string]string`

GetCustomHeaders returns the CustomHeaders field if non-nil, zero value otherwise.

### GetCustomHeadersOk

`func (o *GatewayResponse) GetCustomHeadersOk() (*map[string]string, bool)`

GetCustomHeadersOk returns a tuple with the CustomHeaders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomHeaders

`func (o *GatewayResponse) SetCustomHeaders(v map[string]string)`

SetCustomHeaders sets CustomHeaders field to given value.

### HasCustomHeaders

`func (o *GatewayResponse) HasCustomHeaders() bool`

HasCustomHeaders returns a boolean if a field has been set.

### GetId

`func (o *GatewayResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GatewayResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GatewayResponse) SetId(v string)`

SetId sets Id field to given value.


### GetLastModifiedTime

`func (o *GatewayResponse) GetLastModifiedTime() time.Time`

GetLastModifiedTime returns the LastModifiedTime field if non-nil, zero value otherwise.

### GetLastModifiedTimeOk

`func (o *GatewayResponse) GetLastModifiedTimeOk() (*time.Time, bool)`

GetLastModifiedTimeOk returns a tuple with the LastModifiedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModifiedTime

`func (o *GatewayResponse) SetLastModifiedTime(v time.Time)`

SetLastModifiedTime sets LastModifiedTime field to given value.


### GetUrl

`func (o *GatewayResponse) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *GatewayResponse) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *GatewayResponse) SetUrl(v string)`

SetUrl sets Url field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


