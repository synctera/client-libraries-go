# AddVendorAccountsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AddedAccounts** | [**[]ExternalAccount**](ExternalAccount.md) |  | 
**FailedAccounts** | [**[]AddVendorAccountFailure**](AddVendorAccountFailure.md) |  | 

## Methods

### NewAddVendorAccountsResponse

`func NewAddVendorAccountsResponse(addedAccounts []ExternalAccount, failedAccounts []AddVendorAccountFailure, ) *AddVendorAccountsResponse`

NewAddVendorAccountsResponse instantiates a new AddVendorAccountsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddVendorAccountsResponseWithDefaults

`func NewAddVendorAccountsResponseWithDefaults() *AddVendorAccountsResponse`

NewAddVendorAccountsResponseWithDefaults instantiates a new AddVendorAccountsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddedAccounts

`func (o *AddVendorAccountsResponse) GetAddedAccounts() []ExternalAccount`

GetAddedAccounts returns the AddedAccounts field if non-nil, zero value otherwise.

### GetAddedAccountsOk

`func (o *AddVendorAccountsResponse) GetAddedAccountsOk() (*[]ExternalAccount, bool)`

GetAddedAccountsOk returns a tuple with the AddedAccounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddedAccounts

`func (o *AddVendorAccountsResponse) SetAddedAccounts(v []ExternalAccount)`

SetAddedAccounts sets AddedAccounts field to given value.


### GetFailedAccounts

`func (o *AddVendorAccountsResponse) GetFailedAccounts() []AddVendorAccountFailure`

GetFailedAccounts returns the FailedAccounts field if non-nil, zero value otherwise.

### GetFailedAccountsOk

`func (o *AddVendorAccountsResponse) GetFailedAccountsOk() (*[]AddVendorAccountFailure, bool)`

GetFailedAccountsOk returns a tuple with the FailedAccounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailedAccounts

`func (o *AddVendorAccountsResponse) SetFailedAccounts(v []AddVendorAccountFailure)`

SetFailedAccounts sets FailedAccounts field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


