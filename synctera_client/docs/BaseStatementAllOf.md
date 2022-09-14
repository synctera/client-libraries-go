# BaseStatementAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountSummary** | Pointer to [**AccountSummary**](AccountSummary.md) |  | [optional] 
**AuthorizedSigner** | Pointer to [**[]Person1**](Person1.md) |  | [optional] [readonly] 
**Disclosure** | Pointer to **string** |  | [optional] 
**JointAccountHolders** | Pointer to [**[]Person1**](Person1.md) |  | [optional] [readonly] 
**PrimaryAccountHolderBusiness** | Pointer to [**Business1**](Business1.md) |  | [optional] 
**PrimaryAccountHolderPersonal** | Pointer to [**Person1**](Person1.md) |  | [optional] 
**Transactions** | Pointer to [**[]Transaction**](Transaction.md) |  | [optional] 

## Methods

### NewBaseStatementAllOf

`func NewBaseStatementAllOf() *BaseStatementAllOf`

NewBaseStatementAllOf instantiates a new BaseStatementAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBaseStatementAllOfWithDefaults

`func NewBaseStatementAllOfWithDefaults() *BaseStatementAllOf`

NewBaseStatementAllOfWithDefaults instantiates a new BaseStatementAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountSummary

`func (o *BaseStatementAllOf) GetAccountSummary() AccountSummary`

GetAccountSummary returns the AccountSummary field if non-nil, zero value otherwise.

### GetAccountSummaryOk

`func (o *BaseStatementAllOf) GetAccountSummaryOk() (*AccountSummary, bool)`

GetAccountSummaryOk returns a tuple with the AccountSummary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountSummary

`func (o *BaseStatementAllOf) SetAccountSummary(v AccountSummary)`

SetAccountSummary sets AccountSummary field to given value.

### HasAccountSummary

`func (o *BaseStatementAllOf) HasAccountSummary() bool`

HasAccountSummary returns a boolean if a field has been set.

### GetAuthorizedSigner

`func (o *BaseStatementAllOf) GetAuthorizedSigner() []Person1`

GetAuthorizedSigner returns the AuthorizedSigner field if non-nil, zero value otherwise.

### GetAuthorizedSignerOk

`func (o *BaseStatementAllOf) GetAuthorizedSignerOk() (*[]Person1, bool)`

GetAuthorizedSignerOk returns a tuple with the AuthorizedSigner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorizedSigner

`func (o *BaseStatementAllOf) SetAuthorizedSigner(v []Person1)`

SetAuthorizedSigner sets AuthorizedSigner field to given value.

### HasAuthorizedSigner

`func (o *BaseStatementAllOf) HasAuthorizedSigner() bool`

HasAuthorizedSigner returns a boolean if a field has been set.

### GetDisclosure

`func (o *BaseStatementAllOf) GetDisclosure() string`

GetDisclosure returns the Disclosure field if non-nil, zero value otherwise.

### GetDisclosureOk

`func (o *BaseStatementAllOf) GetDisclosureOk() (*string, bool)`

GetDisclosureOk returns a tuple with the Disclosure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisclosure

`func (o *BaseStatementAllOf) SetDisclosure(v string)`

SetDisclosure sets Disclosure field to given value.

### HasDisclosure

`func (o *BaseStatementAllOf) HasDisclosure() bool`

HasDisclosure returns a boolean if a field has been set.

### GetJointAccountHolders

`func (o *BaseStatementAllOf) GetJointAccountHolders() []Person1`

GetJointAccountHolders returns the JointAccountHolders field if non-nil, zero value otherwise.

### GetJointAccountHoldersOk

`func (o *BaseStatementAllOf) GetJointAccountHoldersOk() (*[]Person1, bool)`

GetJointAccountHoldersOk returns a tuple with the JointAccountHolders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJointAccountHolders

`func (o *BaseStatementAllOf) SetJointAccountHolders(v []Person1)`

SetJointAccountHolders sets JointAccountHolders field to given value.

### HasJointAccountHolders

`func (o *BaseStatementAllOf) HasJointAccountHolders() bool`

HasJointAccountHolders returns a boolean if a field has been set.

### GetPrimaryAccountHolderBusiness

`func (o *BaseStatementAllOf) GetPrimaryAccountHolderBusiness() Business1`

GetPrimaryAccountHolderBusiness returns the PrimaryAccountHolderBusiness field if non-nil, zero value otherwise.

### GetPrimaryAccountHolderBusinessOk

`func (o *BaseStatementAllOf) GetPrimaryAccountHolderBusinessOk() (*Business1, bool)`

GetPrimaryAccountHolderBusinessOk returns a tuple with the PrimaryAccountHolderBusiness field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryAccountHolderBusiness

`func (o *BaseStatementAllOf) SetPrimaryAccountHolderBusiness(v Business1)`

SetPrimaryAccountHolderBusiness sets PrimaryAccountHolderBusiness field to given value.

### HasPrimaryAccountHolderBusiness

`func (o *BaseStatementAllOf) HasPrimaryAccountHolderBusiness() bool`

HasPrimaryAccountHolderBusiness returns a boolean if a field has been set.

### GetPrimaryAccountHolderPersonal

`func (o *BaseStatementAllOf) GetPrimaryAccountHolderPersonal() Person1`

GetPrimaryAccountHolderPersonal returns the PrimaryAccountHolderPersonal field if non-nil, zero value otherwise.

### GetPrimaryAccountHolderPersonalOk

`func (o *BaseStatementAllOf) GetPrimaryAccountHolderPersonalOk() (*Person1, bool)`

GetPrimaryAccountHolderPersonalOk returns a tuple with the PrimaryAccountHolderPersonal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryAccountHolderPersonal

`func (o *BaseStatementAllOf) SetPrimaryAccountHolderPersonal(v Person1)`

SetPrimaryAccountHolderPersonal sets PrimaryAccountHolderPersonal field to given value.

### HasPrimaryAccountHolderPersonal

`func (o *BaseStatementAllOf) HasPrimaryAccountHolderPersonal() bool`

HasPrimaryAccountHolderPersonal returns a boolean if a field has been set.

### GetTransactions

`func (o *BaseStatementAllOf) GetTransactions() []Transaction`

GetTransactions returns the Transactions field if non-nil, zero value otherwise.

### GetTransactionsOk

`func (o *BaseStatementAllOf) GetTransactionsOk() (*[]Transaction, bool)`

GetTransactionsOk returns a tuple with the Transactions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactions

`func (o *BaseStatementAllOf) SetTransactions(v []Transaction)`

SetTransactions sets Transactions field to given value.

### HasTransactions

`func (o *BaseStatementAllOf) HasTransactions() bool`

HasTransactions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


