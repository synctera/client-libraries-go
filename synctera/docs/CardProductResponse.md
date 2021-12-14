# CardProductResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountRangeId** | **string** | ID of the Account Range for which card product was created | 
**Active** | **bool** | indicates whether program is active | 
**AutoAllocateRange** | **bool** | Identifies whether a new account range will be automatically allocated | 
**BankId** | **int32** | The ID of the bank partner works with within this product | 
**BinId** | **string** | Bin ID | 
**CardFormat** | [**CardFormat**](CardFormat.md) |  | 
**CardProgramId** | **string** | Program ID | 
**CreationTime** | **time.Time** | The timestamp representing when the card product was created | [readonly] 
**EndDate** | **time.Time** | The time when bin is decommissioned | 
**FundingSourceId** | Pointer to **string** | Funding source ID | [optional] 
**Id** | **string** | Card Product ID | [readonly] 
**LastModifiedTime** | **time.Time** | The timestamp representing when the card product was last modified | [readonly] 
**Name** | **string** | The name of the card product | 
**PartnerId** | **int32** | The ID of the partner card product belongs to | 
**PhysicalCardFormat** | Pointer to [**PhysicalCardFormat**](PhysicalCardFormat.md) |  | [optional] 
**StartDate** | **time.Time** | The time when bin goes live | 

## Methods

### NewCardProductResponse

`func NewCardProductResponse(accountRangeId string, active bool, autoAllocateRange bool, bankId int32, binId string, cardFormat CardFormat, cardProgramId string, creationTime time.Time, endDate time.Time, id string, lastModifiedTime time.Time, name string, partnerId int32, startDate time.Time, ) *CardProductResponse`

NewCardProductResponse instantiates a new CardProductResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCardProductResponseWithDefaults

`func NewCardProductResponseWithDefaults() *CardProductResponse`

NewCardProductResponseWithDefaults instantiates a new CardProductResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountRangeId

`func (o *CardProductResponse) GetAccountRangeId() string`

GetAccountRangeId returns the AccountRangeId field if non-nil, zero value otherwise.

### GetAccountRangeIdOk

`func (o *CardProductResponse) GetAccountRangeIdOk() (*string, bool)`

GetAccountRangeIdOk returns a tuple with the AccountRangeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountRangeId

`func (o *CardProductResponse) SetAccountRangeId(v string)`

SetAccountRangeId sets AccountRangeId field to given value.


### GetActive

`func (o *CardProductResponse) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *CardProductResponse) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *CardProductResponse) SetActive(v bool)`

SetActive sets Active field to given value.


### GetAutoAllocateRange

`func (o *CardProductResponse) GetAutoAllocateRange() bool`

GetAutoAllocateRange returns the AutoAllocateRange field if non-nil, zero value otherwise.

### GetAutoAllocateRangeOk

`func (o *CardProductResponse) GetAutoAllocateRangeOk() (*bool, bool)`

GetAutoAllocateRangeOk returns a tuple with the AutoAllocateRange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoAllocateRange

`func (o *CardProductResponse) SetAutoAllocateRange(v bool)`

SetAutoAllocateRange sets AutoAllocateRange field to given value.


### GetBankId

`func (o *CardProductResponse) GetBankId() int32`

GetBankId returns the BankId field if non-nil, zero value otherwise.

### GetBankIdOk

`func (o *CardProductResponse) GetBankIdOk() (*int32, bool)`

GetBankIdOk returns a tuple with the BankId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBankId

`func (o *CardProductResponse) SetBankId(v int32)`

SetBankId sets BankId field to given value.


### GetBinId

`func (o *CardProductResponse) GetBinId() string`

GetBinId returns the BinId field if non-nil, zero value otherwise.

### GetBinIdOk

`func (o *CardProductResponse) GetBinIdOk() (*string, bool)`

GetBinIdOk returns a tuple with the BinId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBinId

`func (o *CardProductResponse) SetBinId(v string)`

SetBinId sets BinId field to given value.


### GetCardFormat

`func (o *CardProductResponse) GetCardFormat() CardFormat`

GetCardFormat returns the CardFormat field if non-nil, zero value otherwise.

### GetCardFormatOk

`func (o *CardProductResponse) GetCardFormatOk() (*CardFormat, bool)`

GetCardFormatOk returns a tuple with the CardFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardFormat

`func (o *CardProductResponse) SetCardFormat(v CardFormat)`

SetCardFormat sets CardFormat field to given value.


### GetCardProgramId

`func (o *CardProductResponse) GetCardProgramId() string`

GetCardProgramId returns the CardProgramId field if non-nil, zero value otherwise.

### GetCardProgramIdOk

`func (o *CardProductResponse) GetCardProgramIdOk() (*string, bool)`

GetCardProgramIdOk returns a tuple with the CardProgramId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardProgramId

`func (o *CardProductResponse) SetCardProgramId(v string)`

SetCardProgramId sets CardProgramId field to given value.


### GetCreationTime

`func (o *CardProductResponse) GetCreationTime() time.Time`

GetCreationTime returns the CreationTime field if non-nil, zero value otherwise.

### GetCreationTimeOk

`func (o *CardProductResponse) GetCreationTimeOk() (*time.Time, bool)`

GetCreationTimeOk returns a tuple with the CreationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationTime

`func (o *CardProductResponse) SetCreationTime(v time.Time)`

SetCreationTime sets CreationTime field to given value.


### GetEndDate

`func (o *CardProductResponse) GetEndDate() time.Time`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *CardProductResponse) GetEndDateOk() (*time.Time, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *CardProductResponse) SetEndDate(v time.Time)`

SetEndDate sets EndDate field to given value.


### GetFundingSourceId

`func (o *CardProductResponse) GetFundingSourceId() string`

GetFundingSourceId returns the FundingSourceId field if non-nil, zero value otherwise.

### GetFundingSourceIdOk

`func (o *CardProductResponse) GetFundingSourceIdOk() (*string, bool)`

GetFundingSourceIdOk returns a tuple with the FundingSourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFundingSourceId

`func (o *CardProductResponse) SetFundingSourceId(v string)`

SetFundingSourceId sets FundingSourceId field to given value.

### HasFundingSourceId

`func (o *CardProductResponse) HasFundingSourceId() bool`

HasFundingSourceId returns a boolean if a field has been set.

### GetId

`func (o *CardProductResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CardProductResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CardProductResponse) SetId(v string)`

SetId sets Id field to given value.


### GetLastModifiedTime

`func (o *CardProductResponse) GetLastModifiedTime() time.Time`

GetLastModifiedTime returns the LastModifiedTime field if non-nil, zero value otherwise.

### GetLastModifiedTimeOk

`func (o *CardProductResponse) GetLastModifiedTimeOk() (*time.Time, bool)`

GetLastModifiedTimeOk returns a tuple with the LastModifiedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModifiedTime

`func (o *CardProductResponse) SetLastModifiedTime(v time.Time)`

SetLastModifiedTime sets LastModifiedTime field to given value.


### GetName

`func (o *CardProductResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CardProductResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CardProductResponse) SetName(v string)`

SetName sets Name field to given value.


### GetPartnerId

`func (o *CardProductResponse) GetPartnerId() int32`

GetPartnerId returns the PartnerId field if non-nil, zero value otherwise.

### GetPartnerIdOk

`func (o *CardProductResponse) GetPartnerIdOk() (*int32, bool)`

GetPartnerIdOk returns a tuple with the PartnerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartnerId

`func (o *CardProductResponse) SetPartnerId(v int32)`

SetPartnerId sets PartnerId field to given value.


### GetPhysicalCardFormat

`func (o *CardProductResponse) GetPhysicalCardFormat() PhysicalCardFormat`

GetPhysicalCardFormat returns the PhysicalCardFormat field if non-nil, zero value otherwise.

### GetPhysicalCardFormatOk

`func (o *CardProductResponse) GetPhysicalCardFormatOk() (*PhysicalCardFormat, bool)`

GetPhysicalCardFormatOk returns a tuple with the PhysicalCardFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhysicalCardFormat

`func (o *CardProductResponse) SetPhysicalCardFormat(v PhysicalCardFormat)`

SetPhysicalCardFormat sets PhysicalCardFormat field to given value.

### HasPhysicalCardFormat

`func (o *CardProductResponse) HasPhysicalCardFormat() bool`

HasPhysicalCardFormat returns a boolean if a field has been set.

### GetStartDate

`func (o *CardProductResponse) GetStartDate() time.Time`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *CardProductResponse) GetStartDateOk() (*time.Time, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *CardProductResponse) SetStartDate(v time.Time)`

SetStartDate sets StartDate field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


