# ExternalAccount

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountIdentifiers** | [**AccountIdentifiers**](AccountIdentifiers.md) |  | 
**AccountOwnerNames** | **[]string** | The names of the account owners. Values may be masked, in which case the array will be empty.  | 
**BusinessId** | Pointer to **string** | The identifier for the business customer associated with this external account. Exactly one of &#x60;business_id&#x60; or &#x60;customer_id&#x60; will be returned.  | [optional] 
**CreationTime** | **time.Time** |  | 
**CustomerId** | Pointer to **string** | The identifier for the personal customer associated with this external account. Exactly one of &#x60;customer_id&#x60; or &#x60;business_id&#x60; will be returned.  | [optional] 
**Id** | **string** | External account unique identifier | [readonly] 
**LastUpdatedTime** | **time.Time** |  | 
**Metadata** | Pointer to **map[string]interface{}** | User-supplied JSON format metadata. | [optional] 
**Name** | Pointer to **string** | The official name of the account | [optional] 
**Nickname** | Pointer to **NullableString** | A user-meaningful name for the account | [optional] 
**RoutingIdentifiers** | [**AccountRouting**](AccountRouting.md) |  | 
**Status** | **string** | The current state of the account | 
**Type** | **string** | The type of the account | 
**VendorData** | Pointer to [**ExternalAccountVendorData**](ExternalAccountVendorData.md) |  | [optional] 
**VendorInfo** | Pointer to [**VendorInfo**](VendorInfo.md) |  | [optional] 
**Verification** | [**NullableAccountVerification**](AccountVerification.md) |  | 

## Methods

### NewExternalAccount

`func NewExternalAccount(accountIdentifiers AccountIdentifiers, accountOwnerNames []string, creationTime time.Time, id string, lastUpdatedTime time.Time, routingIdentifiers AccountRouting, status string, type_ string, verification NullableAccountVerification, ) *ExternalAccount`

NewExternalAccount instantiates a new ExternalAccount object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExternalAccountWithDefaults

`func NewExternalAccountWithDefaults() *ExternalAccount`

NewExternalAccountWithDefaults instantiates a new ExternalAccount object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountIdentifiers

`func (o *ExternalAccount) GetAccountIdentifiers() AccountIdentifiers`

GetAccountIdentifiers returns the AccountIdentifiers field if non-nil, zero value otherwise.

### GetAccountIdentifiersOk

`func (o *ExternalAccount) GetAccountIdentifiersOk() (*AccountIdentifiers, bool)`

GetAccountIdentifiersOk returns a tuple with the AccountIdentifiers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountIdentifiers

`func (o *ExternalAccount) SetAccountIdentifiers(v AccountIdentifiers)`

SetAccountIdentifiers sets AccountIdentifiers field to given value.


### GetAccountOwnerNames

`func (o *ExternalAccount) GetAccountOwnerNames() []string`

GetAccountOwnerNames returns the AccountOwnerNames field if non-nil, zero value otherwise.

### GetAccountOwnerNamesOk

`func (o *ExternalAccount) GetAccountOwnerNamesOk() (*[]string, bool)`

GetAccountOwnerNamesOk returns a tuple with the AccountOwnerNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountOwnerNames

`func (o *ExternalAccount) SetAccountOwnerNames(v []string)`

SetAccountOwnerNames sets AccountOwnerNames field to given value.


### GetBusinessId

`func (o *ExternalAccount) GetBusinessId() string`

GetBusinessId returns the BusinessId field if non-nil, zero value otherwise.

### GetBusinessIdOk

`func (o *ExternalAccount) GetBusinessIdOk() (*string, bool)`

GetBusinessIdOk returns a tuple with the BusinessId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessId

`func (o *ExternalAccount) SetBusinessId(v string)`

SetBusinessId sets BusinessId field to given value.

### HasBusinessId

`func (o *ExternalAccount) HasBusinessId() bool`

HasBusinessId returns a boolean if a field has been set.

### GetCreationTime

`func (o *ExternalAccount) GetCreationTime() time.Time`

GetCreationTime returns the CreationTime field if non-nil, zero value otherwise.

### GetCreationTimeOk

`func (o *ExternalAccount) GetCreationTimeOk() (*time.Time, bool)`

GetCreationTimeOk returns a tuple with the CreationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationTime

`func (o *ExternalAccount) SetCreationTime(v time.Time)`

SetCreationTime sets CreationTime field to given value.


### GetCustomerId

`func (o *ExternalAccount) GetCustomerId() string`

GetCustomerId returns the CustomerId field if non-nil, zero value otherwise.

### GetCustomerIdOk

`func (o *ExternalAccount) GetCustomerIdOk() (*string, bool)`

GetCustomerIdOk returns a tuple with the CustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerId

`func (o *ExternalAccount) SetCustomerId(v string)`

SetCustomerId sets CustomerId field to given value.

### HasCustomerId

`func (o *ExternalAccount) HasCustomerId() bool`

HasCustomerId returns a boolean if a field has been set.

### GetId

`func (o *ExternalAccount) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ExternalAccount) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ExternalAccount) SetId(v string)`

SetId sets Id field to given value.


### GetLastUpdatedTime

`func (o *ExternalAccount) GetLastUpdatedTime() time.Time`

GetLastUpdatedTime returns the LastUpdatedTime field if non-nil, zero value otherwise.

### GetLastUpdatedTimeOk

`func (o *ExternalAccount) GetLastUpdatedTimeOk() (*time.Time, bool)`

GetLastUpdatedTimeOk returns a tuple with the LastUpdatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdatedTime

`func (o *ExternalAccount) SetLastUpdatedTime(v time.Time)`

SetLastUpdatedTime sets LastUpdatedTime field to given value.


### GetMetadata

`func (o *ExternalAccount) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *ExternalAccount) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *ExternalAccount) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *ExternalAccount) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetName

`func (o *ExternalAccount) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ExternalAccount) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ExternalAccount) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ExternalAccount) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNickname

`func (o *ExternalAccount) GetNickname() string`

GetNickname returns the Nickname field if non-nil, zero value otherwise.

### GetNicknameOk

`func (o *ExternalAccount) GetNicknameOk() (*string, bool)`

GetNicknameOk returns a tuple with the Nickname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNickname

`func (o *ExternalAccount) SetNickname(v string)`

SetNickname sets Nickname field to given value.

### HasNickname

`func (o *ExternalAccount) HasNickname() bool`

HasNickname returns a boolean if a field has been set.

### SetNicknameNil

`func (o *ExternalAccount) SetNicknameNil(b bool)`

 SetNicknameNil sets the value for Nickname to be an explicit nil

### UnsetNickname
`func (o *ExternalAccount) UnsetNickname()`

UnsetNickname ensures that no value is present for Nickname, not even an explicit nil
### GetRoutingIdentifiers

`func (o *ExternalAccount) GetRoutingIdentifiers() AccountRouting`

GetRoutingIdentifiers returns the RoutingIdentifiers field if non-nil, zero value otherwise.

### GetRoutingIdentifiersOk

`func (o *ExternalAccount) GetRoutingIdentifiersOk() (*AccountRouting, bool)`

GetRoutingIdentifiersOk returns a tuple with the RoutingIdentifiers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoutingIdentifiers

`func (o *ExternalAccount) SetRoutingIdentifiers(v AccountRouting)`

SetRoutingIdentifiers sets RoutingIdentifiers field to given value.


### GetStatus

`func (o *ExternalAccount) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ExternalAccount) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ExternalAccount) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetType

`func (o *ExternalAccount) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ExternalAccount) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ExternalAccount) SetType(v string)`

SetType sets Type field to given value.


### GetVendorData

`func (o *ExternalAccount) GetVendorData() ExternalAccountVendorData`

GetVendorData returns the VendorData field if non-nil, zero value otherwise.

### GetVendorDataOk

`func (o *ExternalAccount) GetVendorDataOk() (*ExternalAccountVendorData, bool)`

GetVendorDataOk returns a tuple with the VendorData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendorData

`func (o *ExternalAccount) SetVendorData(v ExternalAccountVendorData)`

SetVendorData sets VendorData field to given value.

### HasVendorData

`func (o *ExternalAccount) HasVendorData() bool`

HasVendorData returns a boolean if a field has been set.

### GetVendorInfo

`func (o *ExternalAccount) GetVendorInfo() VendorInfo`

GetVendorInfo returns the VendorInfo field if non-nil, zero value otherwise.

### GetVendorInfoOk

`func (o *ExternalAccount) GetVendorInfoOk() (*VendorInfo, bool)`

GetVendorInfoOk returns a tuple with the VendorInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendorInfo

`func (o *ExternalAccount) SetVendorInfo(v VendorInfo)`

SetVendorInfo sets VendorInfo field to given value.

### HasVendorInfo

`func (o *ExternalAccount) HasVendorInfo() bool`

HasVendorInfo returns a boolean if a field has been set.

### GetVerification

`func (o *ExternalAccount) GetVerification() AccountVerification`

GetVerification returns the Verification field if non-nil, zero value otherwise.

### GetVerificationOk

`func (o *ExternalAccount) GetVerificationOk() (*AccountVerification, bool)`

GetVerificationOk returns a tuple with the Verification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerification

`func (o *ExternalAccount) SetVerification(v AccountVerification)`

SetVerification sets Verification field to given value.


### SetVerificationNil

`func (o *ExternalAccount) SetVerificationNil(b bool)`

 SetVerificationNil sets the value for Verification to be an explicit nil

### UnsetVerification
`func (o *ExternalAccount) UnsetVerification()`

UnsetVerification ensures that no value is present for Verification, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


