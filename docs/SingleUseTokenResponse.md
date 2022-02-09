# SingleUseTokenResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Token** | Pointer to **string** |  | [optional] 
**Expires** | **time.Time** | yyyy-MM-ddTHH:mm:ssZ | 
**CustomerAccountMappingId** | Pointer to **string** |  | [optional] 

## Methods

### NewSingleUseTokenResponse

`func NewSingleUseTokenResponse(expires time.Time, ) *SingleUseTokenResponse`

NewSingleUseTokenResponse instantiates a new SingleUseTokenResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSingleUseTokenResponseWithDefaults

`func NewSingleUseTokenResponseWithDefaults() *SingleUseTokenResponse`

NewSingleUseTokenResponseWithDefaults instantiates a new SingleUseTokenResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetToken

`func (o *SingleUseTokenResponse) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *SingleUseTokenResponse) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *SingleUseTokenResponse) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *SingleUseTokenResponse) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetExpires

`func (o *SingleUseTokenResponse) GetExpires() time.Time`

GetExpires returns the Expires field if non-nil, zero value otherwise.

### GetExpiresOk

`func (o *SingleUseTokenResponse) GetExpiresOk() (*time.Time, bool)`

GetExpiresOk returns a tuple with the Expires field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpires

`func (o *SingleUseTokenResponse) SetExpires(v time.Time)`

SetExpires sets Expires field to given value.


### GetCustomerAccountMappingId

`func (o *SingleUseTokenResponse) GetCustomerAccountMappingId() string`

GetCustomerAccountMappingId returns the CustomerAccountMappingId field if non-nil, zero value otherwise.

### GetCustomerAccountMappingIdOk

`func (o *SingleUseTokenResponse) GetCustomerAccountMappingIdOk() (*string, bool)`

GetCustomerAccountMappingIdOk returns a tuple with the CustomerAccountMappingId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerAccountMappingId

`func (o *SingleUseTokenResponse) SetCustomerAccountMappingId(v string)`

SetCustomerAccountMappingId sets CustomerAccountMappingId field to given value.

### HasCustomerAccountMappingId

`func (o *SingleUseTokenResponse) HasCustomerAccountMappingId() bool`

HasCustomerAccountMappingId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


