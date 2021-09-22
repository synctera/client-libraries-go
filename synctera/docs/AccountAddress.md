# AccountAddress

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address** | Pointer to [**Address**](address.md) |  | [optional] [readonly] 
**AddressId** | Pointer to **int32** | Address unique ID | [optional] 
**ConnectId** | Pointer to **string** | Connection ID of the account | [optional] 
**CustomerId** | Pointer to **string** | Customer ID | [optional] 
**DocumentTypeId** | Pointer to **int32** | Document | [optional] 
**Duplicate** | Pointer to **bool** | Indicator of duplicate of the address | [optional] 

## Methods

### NewAccountAddress

`func NewAccountAddress() *AccountAddress`

NewAccountAddress instantiates a new AccountAddress object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountAddressWithDefaults

`func NewAccountAddressWithDefaults() *AccountAddress`

NewAccountAddressWithDefaults instantiates a new AccountAddress object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddress

`func (o *AccountAddress) GetAddress() Address`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *AccountAddress) GetAddressOk() (*Address, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *AccountAddress) SetAddress(v Address)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *AccountAddress) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetAddressId

`func (o *AccountAddress) GetAddressId() int32`

GetAddressId returns the AddressId field if non-nil, zero value otherwise.

### GetAddressIdOk

`func (o *AccountAddress) GetAddressIdOk() (*int32, bool)`

GetAddressIdOk returns a tuple with the AddressId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressId

`func (o *AccountAddress) SetAddressId(v int32)`

SetAddressId sets AddressId field to given value.

### HasAddressId

`func (o *AccountAddress) HasAddressId() bool`

HasAddressId returns a boolean if a field has been set.

### GetConnectId

`func (o *AccountAddress) GetConnectId() string`

GetConnectId returns the ConnectId field if non-nil, zero value otherwise.

### GetConnectIdOk

`func (o *AccountAddress) GetConnectIdOk() (*string, bool)`

GetConnectIdOk returns a tuple with the ConnectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectId

`func (o *AccountAddress) SetConnectId(v string)`

SetConnectId sets ConnectId field to given value.

### HasConnectId

`func (o *AccountAddress) HasConnectId() bool`

HasConnectId returns a boolean if a field has been set.

### GetCustomerId

`func (o *AccountAddress) GetCustomerId() string`

GetCustomerId returns the CustomerId field if non-nil, zero value otherwise.

### GetCustomerIdOk

`func (o *AccountAddress) GetCustomerIdOk() (*string, bool)`

GetCustomerIdOk returns a tuple with the CustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerId

`func (o *AccountAddress) SetCustomerId(v string)`

SetCustomerId sets CustomerId field to given value.

### HasCustomerId

`func (o *AccountAddress) HasCustomerId() bool`

HasCustomerId returns a boolean if a field has been set.

### GetDocumentTypeId

`func (o *AccountAddress) GetDocumentTypeId() int32`

GetDocumentTypeId returns the DocumentTypeId field if non-nil, zero value otherwise.

### GetDocumentTypeIdOk

`func (o *AccountAddress) GetDocumentTypeIdOk() (*int32, bool)`

GetDocumentTypeIdOk returns a tuple with the DocumentTypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentTypeId

`func (o *AccountAddress) SetDocumentTypeId(v int32)`

SetDocumentTypeId sets DocumentTypeId field to given value.

### HasDocumentTypeId

`func (o *AccountAddress) HasDocumentTypeId() bool`

HasDocumentTypeId returns a boolean if a field has been set.

### GetDuplicate

`func (o *AccountAddress) GetDuplicate() bool`

GetDuplicate returns the Duplicate field if non-nil, zero value otherwise.

### GetDuplicateOk

`func (o *AccountAddress) GetDuplicateOk() (*bool, bool)`

GetDuplicateOk returns a tuple with the Duplicate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuplicate

`func (o *AccountAddress) SetDuplicate(v bool)`

SetDuplicate sets Duplicate field to given value.

### HasDuplicate

`func (o *AccountAddress) HasDuplicate() bool`

HasDuplicate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


