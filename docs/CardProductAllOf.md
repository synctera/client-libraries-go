# CardProductAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountRangeId** | **string** | Account Range ID | 
**Active** | **bool** | Indicates whether the Card Product is active | 
**AutoAllocateRange** | **bool** | Identifies whether a new account range will be automatically allocated | 
**BankId** | **int32** | Bank ID | 
**BinId** | **string** | BIN ID | 
**CardProgramId** | **string** | Card Program ID | 
**CreationTime** | Pointer to **time.Time** | The timestamp representing when the Card Product was created | [optional] [readonly] 
**EndDate** | Pointer to **time.Time** | The time when the Card Product is decommissioned | [optional] 
**FundingSourceId** | **string** | Funding Source ID | 
**GatewayId** | Pointer to **string** | Gateway ID, used if the Card Product is utilizing FinTech authorization flow | [optional] 
**Id** | Pointer to **string** | Card Product ID | [optional] [readonly] 
**ImageMode** | Pointer to [**CardImageMode**](CardImageMode.md) |  | [optional] 
**LastModifiedTime** | Pointer to **time.Time** | The timestamp representing when the Card Product was last modified | [optional] [readonly] 
**Name** | **string** | The name of the Card Product | 
**PackageId** | Pointer to **string** | Card fulfillment provider’s package ID | [optional] 
**PartnerId** | **int32** | Partner ID | 
**PhysicalCardFormat** | Pointer to [**PhysicalCardFormat**](PhysicalCardFormat.md) |  | [optional] 
**ReturnAddress** | Pointer to [**Shipping**](Shipping.md) |  | [optional] 
**StartDate** | **time.Time** | The time when the Card Product goes live | 

## Methods

### NewCardProductAllOf

`func NewCardProductAllOf(accountRangeId string, active bool, autoAllocateRange bool, bankId int32, binId string, cardProgramId string, fundingSourceId string, name string, partnerId int32, startDate time.Time, ) *CardProductAllOf`

NewCardProductAllOf instantiates a new CardProductAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCardProductAllOfWithDefaults

`func NewCardProductAllOfWithDefaults() *CardProductAllOf`

NewCardProductAllOfWithDefaults instantiates a new CardProductAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountRangeId

`func (o *CardProductAllOf) GetAccountRangeId() string`

GetAccountRangeId returns the AccountRangeId field if non-nil, zero value otherwise.

### GetAccountRangeIdOk

`func (o *CardProductAllOf) GetAccountRangeIdOk() (*string, bool)`

GetAccountRangeIdOk returns a tuple with the AccountRangeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountRangeId

`func (o *CardProductAllOf) SetAccountRangeId(v string)`

SetAccountRangeId sets AccountRangeId field to given value.


### GetActive

`func (o *CardProductAllOf) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *CardProductAllOf) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *CardProductAllOf) SetActive(v bool)`

SetActive sets Active field to given value.


### GetAutoAllocateRange

`func (o *CardProductAllOf) GetAutoAllocateRange() bool`

GetAutoAllocateRange returns the AutoAllocateRange field if non-nil, zero value otherwise.

### GetAutoAllocateRangeOk

`func (o *CardProductAllOf) GetAutoAllocateRangeOk() (*bool, bool)`

GetAutoAllocateRangeOk returns a tuple with the AutoAllocateRange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoAllocateRange

`func (o *CardProductAllOf) SetAutoAllocateRange(v bool)`

SetAutoAllocateRange sets AutoAllocateRange field to given value.


### GetBankId

`func (o *CardProductAllOf) GetBankId() int32`

GetBankId returns the BankId field if non-nil, zero value otherwise.

### GetBankIdOk

`func (o *CardProductAllOf) GetBankIdOk() (*int32, bool)`

GetBankIdOk returns a tuple with the BankId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBankId

`func (o *CardProductAllOf) SetBankId(v int32)`

SetBankId sets BankId field to given value.


### GetBinId

`func (o *CardProductAllOf) GetBinId() string`

GetBinId returns the BinId field if non-nil, zero value otherwise.

### GetBinIdOk

`func (o *CardProductAllOf) GetBinIdOk() (*string, bool)`

GetBinIdOk returns a tuple with the BinId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBinId

`func (o *CardProductAllOf) SetBinId(v string)`

SetBinId sets BinId field to given value.


### GetCardProgramId

`func (o *CardProductAllOf) GetCardProgramId() string`

GetCardProgramId returns the CardProgramId field if non-nil, zero value otherwise.

### GetCardProgramIdOk

`func (o *CardProductAllOf) GetCardProgramIdOk() (*string, bool)`

GetCardProgramIdOk returns a tuple with the CardProgramId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardProgramId

`func (o *CardProductAllOf) SetCardProgramId(v string)`

SetCardProgramId sets CardProgramId field to given value.


### GetCreationTime

`func (o *CardProductAllOf) GetCreationTime() time.Time`

GetCreationTime returns the CreationTime field if non-nil, zero value otherwise.

### GetCreationTimeOk

`func (o *CardProductAllOf) GetCreationTimeOk() (*time.Time, bool)`

GetCreationTimeOk returns a tuple with the CreationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationTime

`func (o *CardProductAllOf) SetCreationTime(v time.Time)`

SetCreationTime sets CreationTime field to given value.

### HasCreationTime

`func (o *CardProductAllOf) HasCreationTime() bool`

HasCreationTime returns a boolean if a field has been set.

### GetEndDate

`func (o *CardProductAllOf) GetEndDate() time.Time`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *CardProductAllOf) GetEndDateOk() (*time.Time, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *CardProductAllOf) SetEndDate(v time.Time)`

SetEndDate sets EndDate field to given value.

### HasEndDate

`func (o *CardProductAllOf) HasEndDate() bool`

HasEndDate returns a boolean if a field has been set.

### GetFundingSourceId

`func (o *CardProductAllOf) GetFundingSourceId() string`

GetFundingSourceId returns the FundingSourceId field if non-nil, zero value otherwise.

### GetFundingSourceIdOk

`func (o *CardProductAllOf) GetFundingSourceIdOk() (*string, bool)`

GetFundingSourceIdOk returns a tuple with the FundingSourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFundingSourceId

`func (o *CardProductAllOf) SetFundingSourceId(v string)`

SetFundingSourceId sets FundingSourceId field to given value.


### GetGatewayId

`func (o *CardProductAllOf) GetGatewayId() string`

GetGatewayId returns the GatewayId field if non-nil, zero value otherwise.

### GetGatewayIdOk

`func (o *CardProductAllOf) GetGatewayIdOk() (*string, bool)`

GetGatewayIdOk returns a tuple with the GatewayId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewayId

`func (o *CardProductAllOf) SetGatewayId(v string)`

SetGatewayId sets GatewayId field to given value.

### HasGatewayId

`func (o *CardProductAllOf) HasGatewayId() bool`

HasGatewayId returns a boolean if a field has been set.

### GetId

`func (o *CardProductAllOf) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CardProductAllOf) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CardProductAllOf) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CardProductAllOf) HasId() bool`

HasId returns a boolean if a field has been set.

### GetImageMode

`func (o *CardProductAllOf) GetImageMode() CardImageMode`

GetImageMode returns the ImageMode field if non-nil, zero value otherwise.

### GetImageModeOk

`func (o *CardProductAllOf) GetImageModeOk() (*CardImageMode, bool)`

GetImageModeOk returns a tuple with the ImageMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageMode

`func (o *CardProductAllOf) SetImageMode(v CardImageMode)`

SetImageMode sets ImageMode field to given value.

### HasImageMode

`func (o *CardProductAllOf) HasImageMode() bool`

HasImageMode returns a boolean if a field has been set.

### GetLastModifiedTime

`func (o *CardProductAllOf) GetLastModifiedTime() time.Time`

GetLastModifiedTime returns the LastModifiedTime field if non-nil, zero value otherwise.

### GetLastModifiedTimeOk

`func (o *CardProductAllOf) GetLastModifiedTimeOk() (*time.Time, bool)`

GetLastModifiedTimeOk returns a tuple with the LastModifiedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModifiedTime

`func (o *CardProductAllOf) SetLastModifiedTime(v time.Time)`

SetLastModifiedTime sets LastModifiedTime field to given value.

### HasLastModifiedTime

`func (o *CardProductAllOf) HasLastModifiedTime() bool`

HasLastModifiedTime returns a boolean if a field has been set.

### GetName

`func (o *CardProductAllOf) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CardProductAllOf) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CardProductAllOf) SetName(v string)`

SetName sets Name field to given value.


### GetPackageId

`func (o *CardProductAllOf) GetPackageId() string`

GetPackageId returns the PackageId field if non-nil, zero value otherwise.

### GetPackageIdOk

`func (o *CardProductAllOf) GetPackageIdOk() (*string, bool)`

GetPackageIdOk returns a tuple with the PackageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackageId

`func (o *CardProductAllOf) SetPackageId(v string)`

SetPackageId sets PackageId field to given value.

### HasPackageId

`func (o *CardProductAllOf) HasPackageId() bool`

HasPackageId returns a boolean if a field has been set.

### GetPartnerId

`func (o *CardProductAllOf) GetPartnerId() int32`

GetPartnerId returns the PartnerId field if non-nil, zero value otherwise.

### GetPartnerIdOk

`func (o *CardProductAllOf) GetPartnerIdOk() (*int32, bool)`

GetPartnerIdOk returns a tuple with the PartnerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartnerId

`func (o *CardProductAllOf) SetPartnerId(v int32)`

SetPartnerId sets PartnerId field to given value.


### GetPhysicalCardFormat

`func (o *CardProductAllOf) GetPhysicalCardFormat() PhysicalCardFormat`

GetPhysicalCardFormat returns the PhysicalCardFormat field if non-nil, zero value otherwise.

### GetPhysicalCardFormatOk

`func (o *CardProductAllOf) GetPhysicalCardFormatOk() (*PhysicalCardFormat, bool)`

GetPhysicalCardFormatOk returns a tuple with the PhysicalCardFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhysicalCardFormat

`func (o *CardProductAllOf) SetPhysicalCardFormat(v PhysicalCardFormat)`

SetPhysicalCardFormat sets PhysicalCardFormat field to given value.

### HasPhysicalCardFormat

`func (o *CardProductAllOf) HasPhysicalCardFormat() bool`

HasPhysicalCardFormat returns a boolean if a field has been set.

### GetReturnAddress

`func (o *CardProductAllOf) GetReturnAddress() Shipping`

GetReturnAddress returns the ReturnAddress field if non-nil, zero value otherwise.

### GetReturnAddressOk

`func (o *CardProductAllOf) GetReturnAddressOk() (*Shipping, bool)`

GetReturnAddressOk returns a tuple with the ReturnAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturnAddress

`func (o *CardProductAllOf) SetReturnAddress(v Shipping)`

SetReturnAddress sets ReturnAddress field to given value.

### HasReturnAddress

`func (o *CardProductAllOf) HasReturnAddress() bool`

HasReturnAddress returns a boolean if a field has been set.

### GetStartDate

`func (o *CardProductAllOf) GetStartDate() time.Time`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *CardProductAllOf) GetStartDateOk() (*time.Time, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *CardProductAllOf) SetStartDate(v time.Time)`

SetStartDate sets StartDate field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


