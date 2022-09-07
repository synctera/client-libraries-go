# Workspace

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BankId** | **int32** | Bank ID | 
**BankName** | **string** | First Name | 
**Environment** | [**Environment**](Environment.md) |  | 
**PartnerId** | **int32** | Partner ID | 
**PartnerName** | **string** | Last Name | 
**Rank** | **int32** | Each workspace has a rank. The highest-ranked (lowest numerical value) workspace is intended to be presented first within its environment.  | 
**VerificationStatus** | **string** |  | 

## Methods

### NewWorkspace

`func NewWorkspace(bankId int32, bankName string, environment Environment, partnerId int32, partnerName string, rank int32, verificationStatus string, ) *Workspace`

NewWorkspace instantiates a new Workspace object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkspaceWithDefaults

`func NewWorkspaceWithDefaults() *Workspace`

NewWorkspaceWithDefaults instantiates a new Workspace object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBankId

`func (o *Workspace) GetBankId() int32`

GetBankId returns the BankId field if non-nil, zero value otherwise.

### GetBankIdOk

`func (o *Workspace) GetBankIdOk() (*int32, bool)`

GetBankIdOk returns a tuple with the BankId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBankId

`func (o *Workspace) SetBankId(v int32)`

SetBankId sets BankId field to given value.


### GetBankName

`func (o *Workspace) GetBankName() string`

GetBankName returns the BankName field if non-nil, zero value otherwise.

### GetBankNameOk

`func (o *Workspace) GetBankNameOk() (*string, bool)`

GetBankNameOk returns a tuple with the BankName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBankName

`func (o *Workspace) SetBankName(v string)`

SetBankName sets BankName field to given value.


### GetEnvironment

`func (o *Workspace) GetEnvironment() Environment`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *Workspace) GetEnvironmentOk() (*Environment, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *Workspace) SetEnvironment(v Environment)`

SetEnvironment sets Environment field to given value.


### GetPartnerId

`func (o *Workspace) GetPartnerId() int32`

GetPartnerId returns the PartnerId field if non-nil, zero value otherwise.

### GetPartnerIdOk

`func (o *Workspace) GetPartnerIdOk() (*int32, bool)`

GetPartnerIdOk returns a tuple with the PartnerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartnerId

`func (o *Workspace) SetPartnerId(v int32)`

SetPartnerId sets PartnerId field to given value.


### GetPartnerName

`func (o *Workspace) GetPartnerName() string`

GetPartnerName returns the PartnerName field if non-nil, zero value otherwise.

### GetPartnerNameOk

`func (o *Workspace) GetPartnerNameOk() (*string, bool)`

GetPartnerNameOk returns a tuple with the PartnerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartnerName

`func (o *Workspace) SetPartnerName(v string)`

SetPartnerName sets PartnerName field to given value.


### GetRank

`func (o *Workspace) GetRank() int32`

GetRank returns the Rank field if non-nil, zero value otherwise.

### GetRankOk

`func (o *Workspace) GetRankOk() (*int32, bool)`

GetRankOk returns a tuple with the Rank field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRank

`func (o *Workspace) SetRank(v int32)`

SetRank sets Rank field to given value.


### GetVerificationStatus

`func (o *Workspace) GetVerificationStatus() string`

GetVerificationStatus returns the VerificationStatus field if non-nil, zero value otherwise.

### GetVerificationStatusOk

`func (o *Workspace) GetVerificationStatusOk() (*string, bool)`

GetVerificationStatusOk returns a tuple with the VerificationStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerificationStatus

`func (o *Workspace) SetVerificationStatus(v string)`

SetVerificationStatus sets VerificationStatus field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


