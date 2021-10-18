# AddVendorAccountsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CustomerId** | **string** | The identifier for the customer associated with these accounts. | 
**CustomerType** | **string** | The type of customer. | 
**Vendor** | [**ExternalAccountVendorValues**](ExternalAccountVendorValues.md) |  | 
**VendorAccessToken** | **string** | The token provided to link external accounts. For Plaid, this is their &#x60;access_token&#x60;.  | 
**VendorAccountIds** | **[]string** | The list of vendor account IDs that the customer chose to link. For Plaid, these are &#x60;account_id&#x60;s.  | 

## Methods

### NewAddVendorAccountsRequest

`func NewAddVendorAccountsRequest(customerId string, customerType string, vendor ExternalAccountVendorValues, vendorAccessToken string, vendorAccountIds []string, ) *AddVendorAccountsRequest`

NewAddVendorAccountsRequest instantiates a new AddVendorAccountsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddVendorAccountsRequestWithDefaults

`func NewAddVendorAccountsRequestWithDefaults() *AddVendorAccountsRequest`

NewAddVendorAccountsRequestWithDefaults instantiates a new AddVendorAccountsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

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


### GetCustomerType

`func (o *AddVendorAccountsRequest) GetCustomerType() string`

GetCustomerType returns the CustomerType field if non-nil, zero value otherwise.

### GetCustomerTypeOk

`func (o *AddVendorAccountsRequest) GetCustomerTypeOk() (*string, bool)`

GetCustomerTypeOk returns a tuple with the CustomerType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerType

`func (o *AddVendorAccountsRequest) SetCustomerType(v string)`

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



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


