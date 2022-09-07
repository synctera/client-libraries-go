# AddVendorAccountsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BusinessId** | Pointer to **string** | The identifier for the business customer associated with this external account. Exactly one of &#x60;business_id&#x60; or &#x60;customer_id&#x60; must be specified.  | [optional] 
**CustomerId** | Pointer to **string** | The identifier for the personal customer associated with this external account. Exactly one of &#x60;customer_id&#x60; or &#x60;business_id&#x60; must be specified.  | [optional] 
**CustomerType** | [**ExtAccountCustomerType**](ExtAccountCustomerType.md) |  | 
**Vendor** | [**ExternalAccountVendorValues**](ExternalAccountVendorValues.md) |  | 
**VendorAccessToken** | Pointer to **string** | The token provided to link external accounts. For Plaid, this is their &#x60;access_token&#x60;.  | [optional] 
**VendorAccountIds** | Pointer to **[]string** | The list of vendor account IDs that the customer chose to link. For Plaid, these are &#x60;account_id&#x60;s.  | [optional] 
**VendorCustomerId** | Pointer to **string** | The identifier provided by the vendor for the customer associated with this external account.  | [optional] 
**VerifyOwner** | Pointer to **bool** | If true, Synctera will attempt to verify that the external account owner is the same as the customer by comparing external account data to customer data. At least 2 of the following fields must match: name, phone number, email, address. Verification is disabled by default.  | [optional] [default to false]

## Methods

### NewAddVendorAccountsRequest

`func NewAddVendorAccountsRequest(customerType ExtAccountCustomerType, vendor ExternalAccountVendorValues, ) *AddVendorAccountsRequest`

NewAddVendorAccountsRequest instantiates a new AddVendorAccountsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddVendorAccountsRequestWithDefaults

`func NewAddVendorAccountsRequestWithDefaults() *AddVendorAccountsRequest`

NewAddVendorAccountsRequestWithDefaults instantiates a new AddVendorAccountsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBusinessId

`func (o *AddVendorAccountsRequest) GetBusinessId() string`

GetBusinessId returns the BusinessId field if non-nil, zero value otherwise.

### GetBusinessIdOk

`func (o *AddVendorAccountsRequest) GetBusinessIdOk() (*string, bool)`

GetBusinessIdOk returns a tuple with the BusinessId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessId

`func (o *AddVendorAccountsRequest) SetBusinessId(v string)`

SetBusinessId sets BusinessId field to given value.

### HasBusinessId

`func (o *AddVendorAccountsRequest) HasBusinessId() bool`

HasBusinessId returns a boolean if a field has been set.

### GetCustomerId

`func (o *AddVendorAccountsRequest) GetCustomerId() string`

GetCustomerId returns the CustomerId field if non-nil, zero value otherwise.

### GetCustomerIdOk

`func (o *AddVendorAccountsRequest) GetCustomerIdOk() (*string, bool)`

GetCustomerIdOk returns a tuple with the CustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerId

`func (o *AddVendorAccountsRequest) SetCustomerId(v string)`

SetCustomerId sets CustomerId field to given value.

### HasCustomerId

`func (o *AddVendorAccountsRequest) HasCustomerId() bool`

HasCustomerId returns a boolean if a field has been set.

### GetCustomerType

`func (o *AddVendorAccountsRequest) GetCustomerType() ExtAccountCustomerType`

GetCustomerType returns the CustomerType field if non-nil, zero value otherwise.

### GetCustomerTypeOk

`func (o *AddVendorAccountsRequest) GetCustomerTypeOk() (*ExtAccountCustomerType, bool)`

GetCustomerTypeOk returns a tuple with the CustomerType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerType

`func (o *AddVendorAccountsRequest) SetCustomerType(v ExtAccountCustomerType)`

SetCustomerType sets CustomerType field to given value.


### GetVendor

`func (o *AddVendorAccountsRequest) GetVendor() ExternalAccountVendorValues`

GetVendor returns the Vendor field if non-nil, zero value otherwise.

### GetVendorOk

`func (o *AddVendorAccountsRequest) GetVendorOk() (*ExternalAccountVendorValues, bool)`

GetVendorOk returns a tuple with the Vendor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendor

`func (o *AddVendorAccountsRequest) SetVendor(v ExternalAccountVendorValues)`

SetVendor sets Vendor field to given value.


### GetVendorAccessToken

`func (o *AddVendorAccountsRequest) GetVendorAccessToken() string`

GetVendorAccessToken returns the VendorAccessToken field if non-nil, zero value otherwise.

### GetVendorAccessTokenOk

`func (o *AddVendorAccountsRequest) GetVendorAccessTokenOk() (*string, bool)`

GetVendorAccessTokenOk returns a tuple with the VendorAccessToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendorAccessToken

`func (o *AddVendorAccountsRequest) SetVendorAccessToken(v string)`

SetVendorAccessToken sets VendorAccessToken field to given value.

### HasVendorAccessToken

`func (o *AddVendorAccountsRequest) HasVendorAccessToken() bool`

HasVendorAccessToken returns a boolean if a field has been set.

### GetVendorAccountIds

`func (o *AddVendorAccountsRequest) GetVendorAccountIds() []string`

GetVendorAccountIds returns the VendorAccountIds field if non-nil, zero value otherwise.

### GetVendorAccountIdsOk

`func (o *AddVendorAccountsRequest) GetVendorAccountIdsOk() (*[]string, bool)`

GetVendorAccountIdsOk returns a tuple with the VendorAccountIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendorAccountIds

`func (o *AddVendorAccountsRequest) SetVendorAccountIds(v []string)`

SetVendorAccountIds sets VendorAccountIds field to given value.

### HasVendorAccountIds

`func (o *AddVendorAccountsRequest) HasVendorAccountIds() bool`

HasVendorAccountIds returns a boolean if a field has been set.

### GetVendorCustomerId

`func (o *AddVendorAccountsRequest) GetVendorCustomerId() string`

GetVendorCustomerId returns the VendorCustomerId field if non-nil, zero value otherwise.

### GetVendorCustomerIdOk

`func (o *AddVendorAccountsRequest) GetVendorCustomerIdOk() (*string, bool)`

GetVendorCustomerIdOk returns a tuple with the VendorCustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendorCustomerId

`func (o *AddVendorAccountsRequest) SetVendorCustomerId(v string)`

SetVendorCustomerId sets VendorCustomerId field to given value.

### HasVendorCustomerId

`func (o *AddVendorAccountsRequest) HasVendorCustomerId() bool`

HasVendorCustomerId returns a boolean if a field has been set.

### GetVerifyOwner

`func (o *AddVendorAccountsRequest) GetVerifyOwner() bool`

GetVerifyOwner returns the VerifyOwner field if non-nil, zero value otherwise.

### GetVerifyOwnerOk

`func (o *AddVendorAccountsRequest) GetVerifyOwnerOk() (*bool, bool)`

GetVerifyOwnerOk returns a tuple with the VerifyOwner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerifyOwner

`func (o *AddVendorAccountsRequest) SetVerifyOwner(v bool)`

SetVerifyOwner sets VerifyOwner field to given value.

### HasVerifyOwner

`func (o *AddVendorAccountsRequest) HasVerifyOwner() bool`

HasVerifyOwner returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


