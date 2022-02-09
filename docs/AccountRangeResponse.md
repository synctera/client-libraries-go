# AccountRangeResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Account Range Id | [readonly] 
**BankId** | **int32** | The bank ID | 
**PartnerId** | **int32** | The partner ID | 
**BinId** | **string** | The ID of the BIN this account range belogns to | 
**AccountRange** | **[]int32** |  | 
**PhysicalCardFormat** | Pointer to [**PhysicalCardFormat**](PhysicalCardFormat.md) |  | [optional] 
**StartDate** | Pointer to **time.Time** | The time when account range becomes active | [optional] 
**EndDate** | Pointer to **time.Time** | The time when account range becomes inactive | [optional] 
**CreationTime** | **time.Time** | The timestamp representing when the account range was created | [readonly] 
**LastModifiedTime** | **time.Time** | The timestamp representing when the account range was last modified | [readonly] 

## Methods

### NewAccountRangeResponse

`func NewAccountRangeResponse(id string, bankId int32, partnerId int32, binId string, accountRange []int32, creationTime time.Time, lastModifiedTime time.Time, ) *AccountRangeResponse`

NewAccountRangeResponse instantiates a new AccountRangeResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountRangeResponseWithDefaults

`func NewAccountRangeResponseWithDefaults() *AccountRangeResponse`

NewAccountRangeResponseWithDefaults instantiates a new AccountRangeResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AccountRangeResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AccountRangeResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AccountRangeResponse) SetId(v string)`

SetId sets Id field to given value.


### GetBankId

`func (o *AccountRangeResponse) GetBankId() int32`

GetBankId returns the BankId field if non-nil, zero value otherwise.

### GetBankIdOk

`func (o *AccountRangeResponse) GetBankIdOk() (*int32, bool)`

GetBankIdOk returns a tuple with the BankId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBankId

`func (o *AccountRangeResponse) SetBankId(v int32)`

SetBankId sets BankId field to given value.


### GetPartnerId

`func (o *AccountRangeResponse) GetPartnerId() int32`

GetPartnerId returns the PartnerId field if non-nil, zero value otherwise.

### GetPartnerIdOk

`func (o *AccountRangeResponse) GetPartnerIdOk() (*int32, bool)`

GetPartnerIdOk returns a tuple with the PartnerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartnerId

`func (o *AccountRangeResponse) SetPartnerId(v int32)`

SetPartnerId sets PartnerId field to given value.


### GetBinId

`func (o *AccountRangeResponse) GetBinId() string`

GetBinId returns the BinId field if non-nil, zero value otherwise.

### GetBinIdOk

`func (o *AccountRangeResponse) GetBinIdOk() (*string, bool)`

GetBinIdOk returns a tuple with the BinId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBinId

`func (o *AccountRangeResponse) SetBinId(v string)`

SetBinId sets BinId field to given value.


### GetAccountRange

`func (o *AccountRangeResponse) GetAccountRange() []int32`

GetAccountRange returns the AccountRange field if non-nil, zero value otherwise.

### GetAccountRangeOk

`func (o *AccountRangeResponse) GetAccountRangeOk() (*[]int32, bool)`

GetAccountRangeOk returns a tuple with the AccountRange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountRange

`func (o *AccountRangeResponse) SetAccountRange(v []int32)`

SetAccountRange sets AccountRange field to given value.


### GetPhysicalCardFormat

`func (o *AccountRangeResponse) GetPhysicalCardFormat() PhysicalCardFormat`

GetPhysicalCardFormat returns the PhysicalCardFormat field if non-nil, zero value otherwise.

### GetPhysicalCardFormatOk

`func (o *AccountRangeResponse) GetPhysicalCardFormatOk() (*PhysicalCardFormat, bool)`

GetPhysicalCardFormatOk returns a tuple with the PhysicalCardFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhysicalCardFormat

`func (o *AccountRangeResponse) SetPhysicalCardFormat(v PhysicalCardFormat)`

SetPhysicalCardFormat sets PhysicalCardFormat field to given value.

### HasPhysicalCardFormat

`func (o *AccountRangeResponse) HasPhysicalCardFormat() bool`

HasPhysicalCardFormat returns a boolean if a field has been set.

### GetStartDate

`func (o *AccountRangeResponse) GetStartDate() time.Time`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *AccountRangeResponse) GetStartDateOk() (*time.Time, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *AccountRangeResponse) SetStartDate(v time.Time)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *AccountRangeResponse) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### GetEndDate

`func (o *AccountRangeResponse) GetEndDate() time.Time`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *AccountRangeResponse) GetEndDateOk() (*time.Time, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *AccountRangeResponse) SetEndDate(v time.Time)`

SetEndDate sets EndDate field to given value.

### HasEndDate

`func (o *AccountRangeResponse) HasEndDate() bool`

HasEndDate returns a boolean if a field has been set.

### GetCreationTime

`func (o *AccountRangeResponse) GetCreationTime() time.Time`

GetCreationTime returns the CreationTime field if non-nil, zero value otherwise.

### GetCreationTimeOk

`func (o *AccountRangeResponse) GetCreationTimeOk() (*time.Time, bool)`

GetCreationTimeOk returns a tuple with the CreationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationTime

`func (o *AccountRangeResponse) SetCreationTime(v time.Time)`

SetCreationTime sets CreationTime field to given value.


### GetLastModifiedTime

`func (o *AccountRangeResponse) GetLastModifiedTime() time.Time`

GetLastModifiedTime returns the LastModifiedTime field if non-nil, zero value otherwise.

### GetLastModifiedTimeOk

`func (o *AccountRangeResponse) GetLastModifiedTimeOk() (*time.Time, bool)`

GetLastModifiedTimeOk returns a tuple with the LastModifiedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModifiedTime

`func (o *AccountRangeResponse) SetLastModifiedTime(v time.Time)`

SetLastModifiedTime sets LastModifiedTime field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


