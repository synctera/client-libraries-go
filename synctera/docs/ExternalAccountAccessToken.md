# ExternalAccountAccessToken

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RequestId** | Pointer to **string** | A unique identifier for the request, which can be used for troubleshooting | [optional] [readonly] 
**VendorAccessToken** | Pointer to **string** | The access token associated with the Item data is being requested for. | [optional] [readonly] 
**VendorCustomerId** | **string** | The account_id value of the account associated with the returned vendor_access_token | [readonly] 
**VendorPublicToken** | **string** | The user&#39;s public token obtained from successful link login.  | 

## Methods

### NewExternalAccountAccessToken

`func NewExternalAccountAccessToken(vendorCustomerId string, vendorPublicToken string, ) *ExternalAccountAccessToken`

NewExternalAccountAccessToken instantiates a new ExternalAccountAccessToken object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExternalAccountAccessTokenWithDefaults

`func NewExternalAccountAccessTokenWithDefaults() *ExternalAccountAccessToken`

NewExternalAccountAccessTokenWithDefaults instantiates a new ExternalAccountAccessToken object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRequestId

`func (o *ExternalAccountAccessToken) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *ExternalAccountAccessToken) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *ExternalAccountAccessToken) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.

### HasRequestId

`func (o *ExternalAccountAccessToken) HasRequestId() bool`

HasRequestId returns a boolean if a field has been set.

### GetVendorAccessToken

`func (o *ExternalAccountAccessToken) GetVendorAccessToken() string`

GetVendorAccessToken returns the VendorAccessToken field if non-nil, zero value otherwise.

### GetVendorAccessTokenOk

`func (o *ExternalAccountAccessToken) GetVendorAccessTokenOk() (*string, bool)`

GetVendorAccessTokenOk returns a tuple with the VendorAccessToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendorAccessToken

`func (o *ExternalAccountAccessToken) SetVendorAccessToken(v string)`

SetVendorAccessToken sets VendorAccessToken field to given value.

### HasVendorAccessToken

`func (o *ExternalAccountAccessToken) HasVendorAccessToken() bool`

HasVendorAccessToken returns a boolean if a field has been set.

### GetVendorCustomerId

`func (o *ExternalAccountAccessToken) GetVendorCustomerId() string`

GetVendorCustomerId returns the VendorCustomerId field if non-nil, zero value otherwise.

### GetVendorCustomerIdOk

`func (o *ExternalAccountAccessToken) GetVendorCustomerIdOk() (*string, bool)`

GetVendorCustomerIdOk returns a tuple with the VendorCustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendorCustomerId

`func (o *ExternalAccountAccessToken) SetVendorCustomerId(v string)`

SetVendorCustomerId sets VendorCustomerId field to given value.


### GetVendorPublicToken

`func (o *ExternalAccountAccessToken) GetVendorPublicToken() string`

GetVendorPublicToken returns the VendorPublicToken field if non-nil, zero value otherwise.

### GetVendorPublicTokenOk

`func (o *ExternalAccountAccessToken) GetVendorPublicTokenOk() (*string, bool)`

GetVendorPublicTokenOk returns a tuple with the VendorPublicToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendorPublicToken

`func (o *ExternalAccountAccessToken) SetVendorPublicToken(v string)`

SetVendorPublicToken sets VendorPublicToken field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


