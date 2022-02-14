# PatchExternalAccount

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountIdentifiers** | Pointer to [**PatchAccountsRequestAccountIdentifiers**](PatchAccountsRequestAccountIdentifiers.md) |  | [optional] 
**AccountOwnerNames** | Pointer to **[]string** | The names of the account owners. | [optional] 
**Nickname** | Pointer to **string** | A user-meaningful name for the account | [optional] 
**RoutingIdentifiers** | Pointer to [**PatchAccountsRequestRoutingIdentifiers**](PatchAccountsRequestRoutingIdentifiers.md) |  | [optional] 
**Type** | Pointer to **string** | The type of the account | [optional] 
**Verification** | Pointer to [**NullableAccountVerification**](AccountVerification.md) |  | [optional] 

## Methods

### NewPatchExternalAccount

`func NewPatchExternalAccount() *PatchExternalAccount`

NewPatchExternalAccount instantiates a new PatchExternalAccount object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchExternalAccountWithDefaults

`func NewPatchExternalAccountWithDefaults() *PatchExternalAccount`

NewPatchExternalAccountWithDefaults instantiates a new PatchExternalAccount object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountIdentifiers

`func (o *PatchExternalAccount) GetAccountIdentifiers() PatchAccountsRequestAccountIdentifiers`

GetAccountIdentifiers returns the AccountIdentifiers field if non-nil, zero value otherwise.

### GetAccountIdentifiersOk

`func (o *PatchExternalAccount) GetAccountIdentifiersOk() (*PatchAccountsRequestAccountIdentifiers, bool)`

GetAccountIdentifiersOk returns a tuple with the AccountIdentifiers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountIdentifiers

`func (o *PatchExternalAccount) SetAccountIdentifiers(v PatchAccountsRequestAccountIdentifiers)`

SetAccountIdentifiers sets AccountIdentifiers field to given value.

### HasAccountIdentifiers

`func (o *PatchExternalAccount) HasAccountIdentifiers() bool`

HasAccountIdentifiers returns a boolean if a field has been set.

### GetAccountOwnerNames

`func (o *PatchExternalAccount) GetAccountOwnerNames() []string`

GetAccountOwnerNames returns the AccountOwnerNames field if non-nil, zero value otherwise.

### GetAccountOwnerNamesOk

`func (o *PatchExternalAccount) GetAccountOwnerNamesOk() (*[]string, bool)`

GetAccountOwnerNamesOk returns a tuple with the AccountOwnerNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountOwnerNames

`func (o *PatchExternalAccount) SetAccountOwnerNames(v []string)`

SetAccountOwnerNames sets AccountOwnerNames field to given value.

### HasAccountOwnerNames

`func (o *PatchExternalAccount) HasAccountOwnerNames() bool`

HasAccountOwnerNames returns a boolean if a field has been set.

### GetNickname

`func (o *PatchExternalAccount) GetNickname() string`

GetNickname returns the Nickname field if non-nil, zero value otherwise.

### GetNicknameOk

`func (o *PatchExternalAccount) GetNicknameOk() (*string, bool)`

GetNicknameOk returns a tuple with the Nickname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNickname

`func (o *PatchExternalAccount) SetNickname(v string)`

SetNickname sets Nickname field to given value.

### HasNickname

`func (o *PatchExternalAccount) HasNickname() bool`

HasNickname returns a boolean if a field has been set.

### GetRoutingIdentifiers

`func (o *PatchExternalAccount) GetRoutingIdentifiers() PatchAccountsRequestRoutingIdentifiers`

GetRoutingIdentifiers returns the RoutingIdentifiers field if non-nil, zero value otherwise.

### GetRoutingIdentifiersOk

`func (o *PatchExternalAccount) GetRoutingIdentifiersOk() (*PatchAccountsRequestRoutingIdentifiers, bool)`

GetRoutingIdentifiersOk returns a tuple with the RoutingIdentifiers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoutingIdentifiers

`func (o *PatchExternalAccount) SetRoutingIdentifiers(v PatchAccountsRequestRoutingIdentifiers)`

SetRoutingIdentifiers sets RoutingIdentifiers field to given value.

### HasRoutingIdentifiers

`func (o *PatchExternalAccount) HasRoutingIdentifiers() bool`

HasRoutingIdentifiers returns a boolean if a field has been set.

### GetType

`func (o *PatchExternalAccount) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PatchExternalAccount) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PatchExternalAccount) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *PatchExternalAccount) HasType() bool`

HasType returns a boolean if a field has been set.

### GetVerification

`func (o *PatchExternalAccount) GetVerification() AccountVerification`

GetVerification returns the Verification field if non-nil, zero value otherwise.

### GetVerificationOk

`func (o *PatchExternalAccount) GetVerificationOk() (*AccountVerification, bool)`

GetVerificationOk returns a tuple with the Verification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerification

`func (o *PatchExternalAccount) SetVerification(v AccountVerification)`

SetVerification sets Verification field to given value.

### HasVerification

`func (o *PatchExternalAccount) HasVerification() bool`

HasVerification returns a boolean if a field has been set.

### SetVerificationNil

`func (o *PatchExternalAccount) SetVerificationNil(b bool)`

 SetVerificationNil sets the value for Verification to be an explicit nil

### UnsetVerification
`func (o *PatchExternalAccount) UnsetVerification()`

UnsetVerification ensures that no value is present for Verification, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


