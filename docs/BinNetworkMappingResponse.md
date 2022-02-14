# BinNetworkMappingResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Active** | **bool** | indicates whether mapping is active | 
**BankNetworkId** | **string** | ID debit network uses to identify a bank | 
**BinId** | **string** | The ID of the bank&#39;s BIN that uses this debit network | 
**CreationTime** | **time.Time** | The timestamp representing when BIN network mapping was created | [readonly] 
**EndDate** | **time.Time** | The time when mapping becomes inactive | 
**LastModifiedTime** | **time.Time** | The timestamp representing when the BIN network mapping was last modified | [readonly] 
**NetworkId** | **string** | The ID of the debit_network associated with the BIN of the bank | 
**StartDate** | **time.Time** | The time when mapping becomes active | 

## Methods

### NewBinNetworkMappingResponse

`func NewBinNetworkMappingResponse(active bool, bankNetworkId string, binId string, creationTime time.Time, endDate time.Time, lastModifiedTime time.Time, networkId string, startDate time.Time, ) *BinNetworkMappingResponse`

NewBinNetworkMappingResponse instantiates a new BinNetworkMappingResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBinNetworkMappingResponseWithDefaults

`func NewBinNetworkMappingResponseWithDefaults() *BinNetworkMappingResponse`

NewBinNetworkMappingResponseWithDefaults instantiates a new BinNetworkMappingResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActive

`func (o *BinNetworkMappingResponse) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *BinNetworkMappingResponse) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *BinNetworkMappingResponse) SetActive(v bool)`

SetActive sets Active field to given value.


### GetBankNetworkId

`func (o *BinNetworkMappingResponse) GetBankNetworkId() string`

GetBankNetworkId returns the BankNetworkId field if non-nil, zero value otherwise.

### GetBankNetworkIdOk

`func (o *BinNetworkMappingResponse) GetBankNetworkIdOk() (*string, bool)`

GetBankNetworkIdOk returns a tuple with the BankNetworkId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBankNetworkId

`func (o *BinNetworkMappingResponse) SetBankNetworkId(v string)`

SetBankNetworkId sets BankNetworkId field to given value.


### GetBinId

`func (o *BinNetworkMappingResponse) GetBinId() string`

GetBinId returns the BinId field if non-nil, zero value otherwise.

### GetBinIdOk

`func (o *BinNetworkMappingResponse) GetBinIdOk() (*string, bool)`

GetBinIdOk returns a tuple with the BinId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBinId

`func (o *BinNetworkMappingResponse) SetBinId(v string)`

SetBinId sets BinId field to given value.


### GetCreationTime

`func (o *BinNetworkMappingResponse) GetCreationTime() time.Time`

GetCreationTime returns the CreationTime field if non-nil, zero value otherwise.

### GetCreationTimeOk

`func (o *BinNetworkMappingResponse) GetCreationTimeOk() (*time.Time, bool)`

GetCreationTimeOk returns a tuple with the CreationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationTime

`func (o *BinNetworkMappingResponse) SetCreationTime(v time.Time)`

SetCreationTime sets CreationTime field to given value.


### GetEndDate

`func (o *BinNetworkMappingResponse) GetEndDate() time.Time`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *BinNetworkMappingResponse) GetEndDateOk() (*time.Time, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *BinNetworkMappingResponse) SetEndDate(v time.Time)`

SetEndDate sets EndDate field to given value.


### GetLastModifiedTime

`func (o *BinNetworkMappingResponse) GetLastModifiedTime() time.Time`

GetLastModifiedTime returns the LastModifiedTime field if non-nil, zero value otherwise.

### GetLastModifiedTimeOk

`func (o *BinNetworkMappingResponse) GetLastModifiedTimeOk() (*time.Time, bool)`

GetLastModifiedTimeOk returns a tuple with the LastModifiedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModifiedTime

`func (o *BinNetworkMappingResponse) SetLastModifiedTime(v time.Time)`

SetLastModifiedTime sets LastModifiedTime field to given value.


### GetNetworkId

`func (o *BinNetworkMappingResponse) GetNetworkId() string`

GetNetworkId returns the NetworkId field if non-nil, zero value otherwise.

### GetNetworkIdOk

`func (o *BinNetworkMappingResponse) GetNetworkIdOk() (*string, bool)`

GetNetworkIdOk returns a tuple with the NetworkId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkId

`func (o *BinNetworkMappingResponse) SetNetworkId(v string)`

SetNetworkId sets NetworkId field to given value.


### GetStartDate

`func (o *BinNetworkMappingResponse) GetStartDate() time.Time`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *BinNetworkMappingResponse) GetStartDateOk() (*time.Time, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *BinNetworkMappingResponse) SetStartDate(v time.Time)`

SetStartDate sets StartDate field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


