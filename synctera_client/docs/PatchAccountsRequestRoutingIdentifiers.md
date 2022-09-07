# PatchAccountsRequestRoutingIdentifiers

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AchRoutingNumber** | Pointer to **string** | The routing number used for US ACH payments. On write, Synctera will store the entire routing number; on read, we only return the last 4 characters.  | [optional] 
**BankName** | Pointer to **string** | The name of the bank managing the account | [optional] 

## Methods

### NewPatchAccountsRequestRoutingIdentifiers

`func NewPatchAccountsRequestRoutingIdentifiers() *PatchAccountsRequestRoutingIdentifiers`

NewPatchAccountsRequestRoutingIdentifiers instantiates a new PatchAccountsRequestRoutingIdentifiers object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchAccountsRequestRoutingIdentifiersWithDefaults

`func NewPatchAccountsRequestRoutingIdentifiersWithDefaults() *PatchAccountsRequestRoutingIdentifiers`

NewPatchAccountsRequestRoutingIdentifiersWithDefaults instantiates a new PatchAccountsRequestRoutingIdentifiers object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAchRoutingNumber

`func (o *PatchAccountsRequestRoutingIdentifiers) GetAchRoutingNumber() string`

GetAchRoutingNumber returns the AchRoutingNumber field if non-nil, zero value otherwise.

### GetAchRoutingNumberOk

`func (o *PatchAccountsRequestRoutingIdentifiers) GetAchRoutingNumberOk() (*string, bool)`

GetAchRoutingNumberOk returns a tuple with the AchRoutingNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAchRoutingNumber

`func (o *PatchAccountsRequestRoutingIdentifiers) SetAchRoutingNumber(v string)`

SetAchRoutingNumber sets AchRoutingNumber field to given value.

### HasAchRoutingNumber

`func (o *PatchAccountsRequestRoutingIdentifiers) HasAchRoutingNumber() bool`

HasAchRoutingNumber returns a boolean if a field has been set.

### GetBankName

`func (o *PatchAccountsRequestRoutingIdentifiers) GetBankName() string`

GetBankName returns the BankName field if non-nil, zero value otherwise.

### GetBankNameOk

`func (o *PatchAccountsRequestRoutingIdentifiers) GetBankNameOk() (*string, bool)`

GetBankNameOk returns a tuple with the BankName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBankName

`func (o *PatchAccountsRequestRoutingIdentifiers) SetBankName(v string)`

SetBankName sets BankName field to given value.

### HasBankName

`func (o *PatchAccountsRequestRoutingIdentifiers) HasBankName() bool`

HasBankName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


