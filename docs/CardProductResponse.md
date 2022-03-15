# CardProductResponse

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
**CreationTime** | **time.Time** | The timestamp representing when the Card Product was created | [readonly] 
**EndDate** | **time.Time** | The time when the Card Product is decommissioned | 
**FundingSourceId** | **string** | Funding Source ID | 
**GatewayId** | Pointer to **string** | Gateway ID, used if the Card Product is utilizing FinTech authorization flow | [optional] 
**Id** | **string** | Card Product ID | [readonly] 
**ImageMode** | Pointer to [**CardImageMode**](CardImageMode.md) |  | [optional] 
**LastModifiedTime** | **time.Time** | The timestamp representing when the Card Product was last modified | [readonly] 
**Name** | **string** | The name of the Card Product | 
**PackageId** | Pointer to **string** | Card fulfillment provider’s package ID | [optional] 
**PartnerId** | **int32** | Partner ID | 
**PhysicalCardFormat** | Pointer to [**PhysicalCardFormat**](PhysicalCardFormat.md) |  | [optional] 
**ReturnAddress** | Pointer to [**Shipping**](Shipping.md) |  | [optional] 
**StartDate** | **time.Time** | The time when the Card Product goes live | 

## Methods

### NewCardProductResponse

`func NewCardProductResponse(form string, accountRangeId string, active bool, autoAllocateRange bool, bankId int32, binId string, cardProgramId string, creationTime time.Time, endDate time.Time, fundingSourceId string, id string, lastModifiedTime time.Time, name string, partnerId int32, startDate time.Time, ) *CardProductResponse`

NewCardProductResponse instantiates a new CardProductResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCardProductResponseWithDefaults

`func NewCardProductResponseWithDefaults() *CardProductResponse`

NewCardProductResponseWithDefaults instantiates a new CardProductResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetForm

`func (o *CardProductResponse) GetForm() string`

GetForm returns the Form field if non-nil, zero value otherwise.

### GetFormOk

`func (o *CardProductResponse) GetFormOk() (*string, bool)`

GetFormOk returns a tuple with the Form field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForm

`func (o *CardProductResponse) SetForm(v string)`

SetForm sets Form field to given value.


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


### GetGatewayId

`func (o *CardProductResponse) GetGatewayId() string`

GetGatewayId returns the GatewayId field if non-nil, zero value otherwise.

### GetGatewayIdOk

`func (o *CardProductResponse) GetGatewayIdOk() (*string, bool)`

GetGatewayIdOk returns a tuple with the GatewayId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewayId

`func (o *CardProductResponse) SetGatewayId(v string)`

SetGatewayId sets GatewayId field to given value.

### HasGatewayId

`func (o *CardProductResponse) HasGatewayId() bool`

HasGatewayId returns a boolean if a field has been set.

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


### GetImageMode

`func (o *CardProductResponse) GetImageMode() CardImageMode`

GetImageMode returns the ImageMode field if non-nil, zero value otherwise.

### GetImageModeOk

`func (o *CardProductResponse) GetImageModeOk() (*CardImageMode, bool)`

GetImageModeOk returns a tuple with the ImageMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageMode

`func (o *CardProductResponse) SetImageMode(v CardImageMode)`

SetImageMode sets ImageMode field to given value.

### HasImageMode

`func (o *CardProductResponse) HasImageMode() bool`

HasImageMode returns a boolean if a field has been set.

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


### GetPackageId

`func (o *CardProductResponse) GetPackageId() string`

GetPackageId returns the PackageId field if non-nil, zero value otherwise.

### GetPackageIdOk

`func (o *CardProductResponse) GetPackageIdOk() (*string, bool)`

GetPackageIdOk returns a tuple with the PackageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackageId

`func (o *CardProductResponse) SetPackageId(v string)`

SetPackageId sets PackageId field to given value.

### HasPackageId

`func (o *CardProductResponse) HasPackageId() bool`

HasPackageId returns a boolean if a field has been set.

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

### GetReturnAddress

`func (o *CardProductResponse) GetReturnAddress() Shipping`

GetReturnAddress returns the ReturnAddress field if non-nil, zero value otherwise.

### GetReturnAddressOk

`func (o *CardProductResponse) GetReturnAddressOk() (*Shipping, bool)`

GetReturnAddressOk returns a tuple with the ReturnAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturnAddress

`func (o *CardProductResponse) SetReturnAddress(v Shipping)`

SetReturnAddress sets ReturnAddress field to given value.

### HasReturnAddress

`func (o *CardProductResponse) HasReturnAddress() bool`

HasReturnAddress returns a boolean if a field has been set.

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


