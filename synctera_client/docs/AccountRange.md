# AccountRange

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountRange** | **[]int64** |  | 
**BankId** | **int32** | The bank ID | 
**BinId** | **string** | The ID of the BIN this account range belogns to | 
**CreationTime** | Pointer to **time.Time** | The timestamp representing when the account range was created | [optional] [readonly] 
**EndDate** | Pointer to **time.Time** | The time when account range becomes inactive | [optional] 
**Id** | Pointer to **string** | Account Range Id | [optional] [readonly] 
**IsTokenizationEnabled** | Pointer to **bool** | Controls whether account range allows tokenization | [optional] [default to false]
**LastModifiedTime** | Pointer to **time.Time** | The timestamp representing when the account range was last modified | [optional] [readonly] 
**PartnerId** | **int32** | The partner ID | 
**PhysicalCardFormat** | Pointer to [**PhysicalCardFormat**](PhysicalCardFormat.md) |  | [optional] 
**StartDate** | Pointer to **time.Time** | The time when account range becomes active | [optional] 

## Methods

### NewAccountRange

`func NewAccountRange(accountRange []int64, bankId int32, binId string, partnerId int32, ) *AccountRange`

NewAccountRange instantiates a new AccountRange object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountRangeWithDefaults

`func NewAccountRangeWithDefaults() *AccountRange`

NewAccountRangeWithDefaults instantiates a new AccountRange object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountRange

`func (o *AccountRange) GetAccountRange() []int64`

GetAccountRange returns the AccountRange field if non-nil, zero value otherwise.

### GetAccountRangeOk

`func (o *AccountRange) GetAccountRangeOk() (*[]int64, bool)`

GetAccountRangeOk returns a tuple with the AccountRange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountRange

`func (o *AccountRange) SetAccountRange(v []int64)`

SetAccountRange sets AccountRange field to given value.


### GetBankId

`func (o *AccountRange) GetBankId() int32`

GetBankId returns the BankId field if non-nil, zero value otherwise.

### GetBankIdOk

`func (o *AccountRange) GetBankIdOk() (*int32, bool)`

GetBankIdOk returns a tuple with the BankId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBankId

`func (o *AccountRange) SetBankId(v int32)`

SetBankId sets BankId field to given value.


### GetBinId

`func (o *AccountRange) GetBinId() string`

GetBinId returns the BinId field if non-nil, zero value otherwise.

### GetBinIdOk

`func (o *AccountRange) GetBinIdOk() (*string, bool)`

GetBinIdOk returns a tuple with the BinId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBinId

`func (o *AccountRange) SetBinId(v string)`

SetBinId sets BinId field to given value.


### GetCreationTime

`func (o *AccountRange) GetCreationTime() time.Time`

GetCreationTime returns the CreationTime field if non-nil, zero value otherwise.

### GetCreationTimeOk

`func (o *AccountRange) GetCreationTimeOk() (*time.Time, bool)`

GetCreationTimeOk returns a tuple with the CreationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationTime

`func (o *AccountRange) SetCreationTime(v time.Time)`

SetCreationTime sets CreationTime field to given value.

### HasCreationTime

`func (o *AccountRange) HasCreationTime() bool`

HasCreationTime returns a boolean if a field has been set.

### GetEndDate

`func (o *AccountRange) GetEndDate() time.Time`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *AccountRange) GetEndDateOk() (*time.Time, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *AccountRange) SetEndDate(v time.Time)`

SetEndDate sets EndDate field to given value.

### HasEndDate

`func (o *AccountRange) HasEndDate() bool`

HasEndDate returns a boolean if a field has been set.

### GetId

`func (o *AccountRange) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AccountRange) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AccountRange) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *AccountRange) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIsTokenizationEnabled

`func (o *AccountRange) GetIsTokenizationEnabled() bool`

GetIsTokenizationEnabled returns the IsTokenizationEnabled field if non-nil, zero value otherwise.

### GetIsTokenizationEnabledOk

`func (o *AccountRange) GetIsTokenizationEnabledOk() (*bool, bool)`

GetIsTokenizationEnabledOk returns a tuple with the IsTokenizationEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsTokenizationEnabled

`func (o *AccountRange) SetIsTokenizationEnabled(v bool)`

SetIsTokenizationEnabled sets IsTokenizationEnabled field to given value.

### HasIsTokenizationEnabled

`func (o *AccountRange) HasIsTokenizationEnabled() bool`

HasIsTokenizationEnabled returns a boolean if a field has been set.

### GetLastModifiedTime

`func (o *AccountRange) GetLastModifiedTime() time.Time`

GetLastModifiedTime returns the LastModifiedTime field if non-nil, zero value otherwise.

### GetLastModifiedTimeOk

`func (o *AccountRange) GetLastModifiedTimeOk() (*time.Time, bool)`

GetLastModifiedTimeOk returns a tuple with the LastModifiedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModifiedTime

`func (o *AccountRange) SetLastModifiedTime(v time.Time)`

SetLastModifiedTime sets LastModifiedTime field to given value.

### HasLastModifiedTime

`func (o *AccountRange) HasLastModifiedTime() bool`

HasLastModifiedTime returns a boolean if a field has been set.

### GetPartnerId

`func (o *AccountRange) GetPartnerId() int32`

GetPartnerId returns the PartnerId field if non-nil, zero value otherwise.

### GetPartnerIdOk

`func (o *AccountRange) GetPartnerIdOk() (*int32, bool)`

GetPartnerIdOk returns a tuple with the PartnerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartnerId

`func (o *AccountRange) SetPartnerId(v int32)`

SetPartnerId sets PartnerId field to given value.


### GetPhysicalCardFormat

`func (o *AccountRange) GetPhysicalCardFormat() PhysicalCardFormat`

GetPhysicalCardFormat returns the PhysicalCardFormat field if non-nil, zero value otherwise.

### GetPhysicalCardFormatOk

`func (o *AccountRange) GetPhysicalCardFormatOk() (*PhysicalCardFormat, bool)`

GetPhysicalCardFormatOk returns a tuple with the PhysicalCardFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhysicalCardFormat

`func (o *AccountRange) SetPhysicalCardFormat(v PhysicalCardFormat)`

SetPhysicalCardFormat sets PhysicalCardFormat field to given value.

### HasPhysicalCardFormat

`func (o *AccountRange) HasPhysicalCardFormat() bool`

HasPhysicalCardFormat returns a boolean if a field has been set.

### GetStartDate

`func (o *AccountRange) GetStartDate() time.Time`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *AccountRange) GetStartDateOk() (*time.Time, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *AccountRange) SetStartDate(v time.Time)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *AccountRange) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


