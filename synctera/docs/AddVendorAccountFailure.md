# AddVendorAccountFailure

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Reason** | [**AddVendorAccountsErrorReason**](AddVendorAccountsErrorReason.md) |  | 
**ReasonDescription** | **string** | A human-readable message describing the reason for the failure. | 
**VendorAccountId** | **string** | The vendor account ID for the account that failed. For Plaid, this is an &#x60;account_id&#x60;.  | 
**VendorErrorMessage** | Pointer to **string** | The display_message returned by the vendor. Only returned if reason is set to &#x60;PROVIDER_ERROR&#x60;. For Plaid, this is the &#x60;display_message&#x60;.  | [optional] 
**VendorRequestId** | Pointer to **string** | A unique identifier for the request from the vendor, which can be used for troubleshooting. Only returned if reason is set to &#x60;PROVIDER_ERROR&#x60;.  | [optional] 

## Methods

### NewAddVendorAccountFailure

`func NewAddVendorAccountFailure(reason AddVendorAccountsErrorReason, reasonDescription string, vendorAccountId string, ) *AddVendorAccountFailure`

NewAddVendorAccountFailure instantiates a new AddVendorAccountFailure object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddVendorAccountFailureWithDefaults

`func NewAddVendorAccountFailureWithDefaults() *AddVendorAccountFailure`

NewAddVendorAccountFailureWithDefaults instantiates a new AddVendorAccountFailure object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReason

`func (o *AddVendorAccountFailure) GetReason() AddVendorAccountsErrorReason`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *AddVendorAccountFailure) GetReasonOk() (*AddVendorAccountsErrorReason, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *AddVendorAccountFailure) SetReason(v AddVendorAccountsErrorReason)`

SetReason sets Reason field to given value.


### GetReasonDescription

`func (o *AddVendorAccountFailure) GetReasonDescription() string`

GetReasonDescription returns the ReasonDescription field if non-nil, zero value otherwise.

### GetReasonDescriptionOk

`func (o *AddVendorAccountFailure) GetReasonDescriptionOk() (*string, bool)`

GetReasonDescriptionOk returns a tuple with the ReasonDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReasonDescription

`func (o *AddVendorAccountFailure) SetReasonDescription(v string)`

SetReasonDescription sets ReasonDescription field to given value.


### GetVendorAccountId

`func (o *AddVendorAccountFailure) GetVendorAccountId() string`

GetVendorAccountId returns the VendorAccountId field if non-nil, zero value otherwise.

### GetVendorAccountIdOk

`func (o *AddVendorAccountFailure) GetVendorAccountIdOk() (*string, bool)`

GetVendorAccountIdOk returns a tuple with the VendorAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendorAccountId

`func (o *AddVendorAccountFailure) SetVendorAccountId(v string)`

SetVendorAccountId sets VendorAccountId field to given value.


### GetVendorErrorMessage

`func (o *AddVendorAccountFailure) GetVendorErrorMessage() string`

GetVendorErrorMessage returns the VendorErrorMessage field if non-nil, zero value otherwise.

### GetVendorErrorMessageOk

`func (o *AddVendorAccountFailure) GetVendorErrorMessageOk() (*string, bool)`

GetVendorErrorMessageOk returns a tuple with the VendorErrorMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendorErrorMessage

`func (o *AddVendorAccountFailure) SetVendorErrorMessage(v string)`

SetVendorErrorMessage sets VendorErrorMessage field to given value.

### HasVendorErrorMessage

`func (o *AddVendorAccountFailure) HasVendorErrorMessage() bool`

HasVendorErrorMessage returns a boolean if a field has been set.

### GetVendorRequestId

`func (o *AddVendorAccountFailure) GetVendorRequestId() string`

GetVendorRequestId returns the VendorRequestId field if non-nil, zero value otherwise.

### GetVendorRequestIdOk

`func (o *AddVendorAccountFailure) GetVendorRequestIdOk() (*string, bool)`

GetVendorRequestIdOk returns a tuple with the VendorRequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendorRequestId

`func (o *AddVendorAccountFailure) SetVendorRequestId(v string)`

SetVendorRequestId sets VendorRequestId field to given value.

### HasVendorRequestId

`func (o *AddVendorAccountFailure) HasVendorRequestId() bool`

HasVendorRequestId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


