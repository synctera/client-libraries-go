# BinNetworkMapping

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Active** | **bool** | indicates whether mapping is active | 
**BankNetworkId** | **string** | ID debit network uses to identify a bank | 
**BinId** | **string** | The ID of the bank&#39;s BIN that uses this debit network | 
**CreationTime** | Pointer to **time.Time** | The timestamp representing when BIN network mapping was created | [optional] [readonly] 
**EndDate** | Pointer to **time.Time** | The time when mapping becomes inactive | [optional] 
**LastModifiedTime** | Pointer to **time.Time** | The timestamp representing when the BIN network mapping was last modified | [optional] [readonly] 
**NetworkId** | **string** | The ID of the debit_network associated with the BIN of the bank | 
**StartDate** | Pointer to **time.Time** | The time when mapping becomes active | [optional] 

## Methods

### NewBinNetworkMapping

`func NewBinNetworkMapping(active bool, bankNetworkId string, binId string, networkId string, ) *BinNetworkMapping`

NewBinNetworkMapping instantiates a new BinNetworkMapping object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBinNetworkMappingWithDefaults

`func NewBinNetworkMappingWithDefaults() *BinNetworkMapping`

NewBinNetworkMappingWithDefaults instantiates a new BinNetworkMapping object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActive

`func (o *BinNetworkMapping) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *BinNetworkMapping) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *BinNetworkMapping) SetActive(v bool)`

SetActive sets Active field to given value.


### GetBankNetworkId

`func (o *BinNetworkMapping) GetBankNetworkId() string`

GetBankNetworkId returns the BankNetworkId field if non-nil, zero value otherwise.

### GetBankNetworkIdOk

`func (o *BinNetworkMapping) GetBankNetworkIdOk() (*string, bool)`

GetBankNetworkIdOk returns a tuple with the BankNetworkId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBankNetworkId

`func (o *BinNetworkMapping) SetBankNetworkId(v string)`

SetBankNetworkId sets BankNetworkId field to given value.


### GetBinId

`func (o *BinNetworkMapping) GetBinId() string`

GetBinId returns the BinId field if non-nil, zero value otherwise.

### GetBinIdOk

`func (o *BinNetworkMapping) GetBinIdOk() (*string, bool)`

GetBinIdOk returns a tuple with the BinId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBinId

`func (o *BinNetworkMapping) SetBinId(v string)`

SetBinId sets BinId field to given value.


### GetCreationTime

`func (o *BinNetworkMapping) GetCreationTime() time.Time`

GetCreationTime returns the CreationTime field if non-nil, zero value otherwise.

### GetCreationTimeOk

`func (o *BinNetworkMapping) GetCreationTimeOk() (*time.Time, bool)`

GetCreationTimeOk returns a tuple with the CreationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationTime

`func (o *BinNetworkMapping) SetCreationTime(v time.Time)`

SetCreationTime sets CreationTime field to given value.

### HasCreationTime

`func (o *BinNetworkMapping) HasCreationTime() bool`

HasCreationTime returns a boolean if a field has been set.

### GetEndDate

`func (o *BinNetworkMapping) GetEndDate() time.Time`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *BinNetworkMapping) GetEndDateOk() (*time.Time, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *BinNetworkMapping) SetEndDate(v time.Time)`

SetEndDate sets EndDate field to given value.

### HasEndDate

`func (o *BinNetworkMapping) HasEndDate() bool`

HasEndDate returns a boolean if a field has been set.

### GetLastModifiedTime

`func (o *BinNetworkMapping) GetLastModifiedTime() time.Time`

GetLastModifiedTime returns the LastModifiedTime field if non-nil, zero value otherwise.

### GetLastModifiedTimeOk

`func (o *BinNetworkMapping) GetLastModifiedTimeOk() (*time.Time, bool)`

GetLastModifiedTimeOk returns a tuple with the LastModifiedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModifiedTime

`func (o *BinNetworkMapping) SetLastModifiedTime(v time.Time)`

SetLastModifiedTime sets LastModifiedTime field to given value.

### HasLastModifiedTime

`func (o *BinNetworkMapping) HasLastModifiedTime() bool`

HasLastModifiedTime returns a boolean if a field has been set.

### GetNetworkId

`func (o *BinNetworkMapping) GetNetworkId() string`

GetNetworkId returns the NetworkId field if non-nil, zero value otherwise.

### GetNetworkIdOk

`func (o *BinNetworkMapping) GetNetworkIdOk() (*string, bool)`

GetNetworkIdOk returns a tuple with the NetworkId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkId

`func (o *BinNetworkMapping) SetNetworkId(v string)`

SetNetworkId sets NetworkId field to given value.


### GetStartDate

`func (o *BinNetworkMapping) GetStartDate() time.Time`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *BinNetworkMapping) GetStartDateOk() (*time.Time, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *BinNetworkMapping) SetStartDate(v time.Time)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *BinNetworkMapping) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


