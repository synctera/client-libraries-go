# CardProduct

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Form** | **string** | PHYSICAL or VIRTUAL. | 
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

### NewCardProduct

`func NewCardProduct(form string, accountRangeId string, active bool, autoAllocateRange bool, bankId int32, binId string, cardProgramId string, fundingSourceId string, name string, partnerId int32, startDate time.Time, ) *CardProduct`

NewCardProduct instantiates a new CardProduct object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCardProductWithDefaults

`func NewCardProductWithDefaults() *CardProduct`

NewCardProductWithDefaults instantiates a new CardProduct object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetForm

`func (o *CardProduct) GetForm() string`

GetForm returns the Form field if non-nil, zero value otherwise.

### GetFormOk

`func (o *CardProduct) GetFormOk() (*string, bool)`

GetFormOk returns a tuple with the Form field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForm

`func (o *CardProduct) SetForm(v string)`

SetForm sets Form field to given value.


### GetAccountRangeId

`func (o *CardProduct) GetAccountRangeId() string`

GetAccountRangeId returns the AccountRangeId field if non-nil, zero value otherwise.

### GetAccountRangeIdOk

`func (o *CardProduct) GetAccountRangeIdOk() (*string, bool)`

GetAccountRangeIdOk returns a tuple with the AccountRangeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountRangeId

`func (o *CardProduct) SetAccountRangeId(v string)`

SetAccountRangeId sets AccountRangeId field to given value.


### GetActive

`func (o *CardProduct) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *CardProduct) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *CardProduct) SetActive(v bool)`

SetActive sets Active field to given value.


### GetAutoAllocateRange

`func (o *CardProduct) GetAutoAllocateRange() bool`

GetAutoAllocateRange returns the AutoAllocateRange field if non-nil, zero value otherwise.

### GetAutoAllocateRangeOk

`func (o *CardProduct) GetAutoAllocateRangeOk() (*bool, bool)`

GetAutoAllocateRangeOk returns a tuple with the AutoAllocateRange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoAllocateRange

`func (o *CardProduct) SetAutoAllocateRange(v bool)`

SetAutoAllocateRange sets AutoAllocateRange field to given value.


### GetBankId

`func (o *CardProduct) GetBankId() int32`

GetBankId returns the BankId field if non-nil, zero value otherwise.

### GetBankIdOk

`func (o *CardProduct) GetBankIdOk() (*int32, bool)`

GetBankIdOk returns a tuple with the BankId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBankId

`func (o *CardProduct) SetBankId(v int32)`

SetBankId sets BankId field to given value.


### GetBinId

`func (o *CardProduct) GetBinId() string`

GetBinId returns the BinId field if non-nil, zero value otherwise.

### GetBinIdOk

`func (o *CardProduct) GetBinIdOk() (*string, bool)`

GetBinIdOk returns a tuple with the BinId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBinId

`func (o *CardProduct) SetBinId(v string)`

SetBinId sets BinId field to given value.


### GetCardProgramId

`func (o *CardProduct) GetCardProgramId() string`

GetCardProgramId returns the CardProgramId field if non-nil, zero value otherwise.

### GetCardProgramIdOk

`func (o *CardProduct) GetCardProgramIdOk() (*string, bool)`

GetCardProgramIdOk returns a tuple with the CardProgramId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardProgramId

`func (o *CardProduct) SetCardProgramId(v string)`

SetCardProgramId sets CardProgramId field to given value.


### GetCreationTime

`func (o *CardProduct) GetCreationTime() time.Time`

GetCreationTime returns the CreationTime field if non-nil, zero value otherwise.

### GetCreationTimeOk

`func (o *CardProduct) GetCreationTimeOk() (*time.Time, bool)`

GetCreationTimeOk returns a tuple with the CreationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationTime

`func (o *CardProduct) SetCreationTime(v time.Time)`

SetCreationTime sets CreationTime field to given value.

### HasCreationTime

`func (o *CardProduct) HasCreationTime() bool`

HasCreationTime returns a boolean if a field has been set.

### GetEndDate

`func (o *CardProduct) GetEndDate() time.Time`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *CardProduct) GetEndDateOk() (*time.Time, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *CardProduct) SetEndDate(v time.Time)`

SetEndDate sets EndDate field to given value.

### HasEndDate

`func (o *CardProduct) HasEndDate() bool`

HasEndDate returns a boolean if a field has been set.

### GetFundingSourceId

`func (o *CardProduct) GetFundingSourceId() string`

GetFundingSourceId returns the FundingSourceId field if non-nil, zero value otherwise.

### GetFundingSourceIdOk

`func (o *CardProduct) GetFundingSourceIdOk() (*string, bool)`

GetFundingSourceIdOk returns a tuple with the FundingSourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFundingSourceId

`func (o *CardProduct) SetFundingSourceId(v string)`

SetFundingSourceId sets FundingSourceId field to given value.


### GetGatewayId

`func (o *CardProduct) GetGatewayId() string`

GetGatewayId returns the GatewayId field if non-nil, zero value otherwise.

### GetGatewayIdOk

`func (o *CardProduct) GetGatewayIdOk() (*string, bool)`

GetGatewayIdOk returns a tuple with the GatewayId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewayId

`func (o *CardProduct) SetGatewayId(v string)`

SetGatewayId sets GatewayId field to given value.

### HasGatewayId

`func (o *CardProduct) HasGatewayId() bool`

HasGatewayId returns a boolean if a field has been set.

### GetId

`func (o *CardProduct) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CardProduct) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CardProduct) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CardProduct) HasId() bool`

HasId returns a boolean if a field has been set.

### GetImageMode

`func (o *CardProduct) GetImageMode() CardImageMode`

GetImageMode returns the ImageMode field if non-nil, zero value otherwise.

### GetImageModeOk

`func (o *CardProduct) GetImageModeOk() (*CardImageMode, bool)`

GetImageModeOk returns a tuple with the ImageMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageMode

`func (o *CardProduct) SetImageMode(v CardImageMode)`

SetImageMode sets ImageMode field to given value.

### HasImageMode

`func (o *CardProduct) HasImageMode() bool`

HasImageMode returns a boolean if a field has been set.

### GetLastModifiedTime

`func (o *CardProduct) GetLastModifiedTime() time.Time`

GetLastModifiedTime returns the LastModifiedTime field if non-nil, zero value otherwise.

### GetLastModifiedTimeOk

`func (o *CardProduct) GetLastModifiedTimeOk() (*time.Time, bool)`

GetLastModifiedTimeOk returns a tuple with the LastModifiedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModifiedTime

`func (o *CardProduct) SetLastModifiedTime(v time.Time)`

SetLastModifiedTime sets LastModifiedTime field to given value.

### HasLastModifiedTime

`func (o *CardProduct) HasLastModifiedTime() bool`

HasLastModifiedTime returns a boolean if a field has been set.

### GetName

`func (o *CardProduct) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CardProduct) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CardProduct) SetName(v string)`

SetName sets Name field to given value.


### GetPackageId

`func (o *CardProduct) GetPackageId() string`

GetPackageId returns the PackageId field if non-nil, zero value otherwise.

### GetPackageIdOk

`func (o *CardProduct) GetPackageIdOk() (*string, bool)`

GetPackageIdOk returns a tuple with the PackageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackageId

`func (o *CardProduct) SetPackageId(v string)`

SetPackageId sets PackageId field to given value.

### HasPackageId

`func (o *CardProduct) HasPackageId() bool`

HasPackageId returns a boolean if a field has been set.

### GetPartnerId

`func (o *CardProduct) GetPartnerId() int32`

GetPartnerId returns the PartnerId field if non-nil, zero value otherwise.

### GetPartnerIdOk

`func (o *CardProduct) GetPartnerIdOk() (*int32, bool)`

GetPartnerIdOk returns a tuple with the PartnerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartnerId

`func (o *CardProduct) SetPartnerId(v int32)`

SetPartnerId sets PartnerId field to given value.


### GetPhysicalCardFormat

`func (o *CardProduct) GetPhysicalCardFormat() PhysicalCardFormat`

GetPhysicalCardFormat returns the PhysicalCardFormat field if non-nil, zero value otherwise.

### GetPhysicalCardFormatOk

`func (o *CardProduct) GetPhysicalCardFormatOk() (*PhysicalCardFormat, bool)`

GetPhysicalCardFormatOk returns a tuple with the PhysicalCardFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhysicalCardFormat

`func (o *CardProduct) SetPhysicalCardFormat(v PhysicalCardFormat)`

SetPhysicalCardFormat sets PhysicalCardFormat field to given value.

### HasPhysicalCardFormat

`func (o *CardProduct) HasPhysicalCardFormat() bool`

HasPhysicalCardFormat returns a boolean if a field has been set.

### GetReturnAddress

`func (o *CardProduct) GetReturnAddress() Shipping`

GetReturnAddress returns the ReturnAddress field if non-nil, zero value otherwise.

### GetReturnAddressOk

`func (o *CardProduct) GetReturnAddressOk() (*Shipping, bool)`

GetReturnAddressOk returns a tuple with the ReturnAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturnAddress

`func (o *CardProduct) SetReturnAddress(v Shipping)`

SetReturnAddress sets ReturnAddress field to given value.

### HasReturnAddress

`func (o *CardProduct) HasReturnAddress() bool`

HasReturnAddress returns a boolean if a field has been set.

### GetStartDate

`func (o *CardProduct) GetStartDate() time.Time`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *CardProduct) GetStartDateOk() (*time.Time, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *CardProduct) SetStartDate(v time.Time)`

SetStartDate sets StartDate field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


