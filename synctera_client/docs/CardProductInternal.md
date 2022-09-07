# CardProductInternal

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Form** | **string** | PHYSICAL or VIRTUAL. | 
**Active** | **bool** | Indicates whether the Card Product is active | 
**CardProgramId** | **string** | Card Program ID | 
**Color** | Pointer to **string** | Color code for dynamic card elements such as PAN and card holder name | [optional] 
**CreationTime** | Pointer to **time.Time** | The timestamp representing when the Card Product was created | [optional] [readonly] 
**DigitalWalletTokenization** | Pointer to [**DigitalWalletTokenization**](DigitalWalletTokenization.md) |  | [optional] 
**EndDate** | Pointer to **time.Time** | The time when the Card Product is decommissioned | [optional] 
**Id** | Pointer to **string** | Card Product ID | [optional] [readonly] 
**Image** | Pointer to **bool** | Indicates whether or not there is an overlay image of the card product available | [optional] 
**ImageMode** | Pointer to [**CardImageMode**](CardImageMode.md) |  | [optional] 
**LastModifiedTime** | Pointer to **time.Time** | The timestamp representing when the Card Product was last modified | [optional] [readonly] 
**Name** | **string** | The name of the Card Product | 
**Orientation** | Pointer to **string** | Card orientation | [optional] 
**PackageId** | Pointer to **string** | Card fulfillment provider’s package ID | [optional] 
**PhysicalCardFormat** | Pointer to [**PhysicalCardFormat**](PhysicalCardFormat.md) |  | [optional] 
**ReturnAddress** | Pointer to [**Shipping**](Shipping.md) |  | [optional] 
**StartDate** | **time.Time** | The time when the Card Product goes live | 
**TxnEnhancer** | Pointer to [**TxnEnhancer**](TxnEnhancer.md) |  | [optional] [default to TXNENHANCER_NONE]
**AccountRangeId** | **string** | Account Range ID | 
**AutoAllocateRange** | **bool** | Identifies whether a new account range will be automatically allocated | 
**BankId** | **int32** | Bank ID | 
**BinId** | **string** | BIN ID | 
**FundingSourceId** | **string** | Funding Source ID | 
**PartnerId** | **int32** | Partner ID | 

## Methods

### NewCardProductInternal

`func NewCardProductInternal(form string, active bool, cardProgramId string, name string, startDate time.Time, accountRangeId string, autoAllocateRange bool, bankId int32, binId string, fundingSourceId string, partnerId int32, ) *CardProductInternal`

NewCardProductInternal instantiates a new CardProductInternal object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCardProductInternalWithDefaults

`func NewCardProductInternalWithDefaults() *CardProductInternal`

NewCardProductInternalWithDefaults instantiates a new CardProductInternal object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetForm

`func (o *CardProductInternal) GetForm() string`

GetForm returns the Form field if non-nil, zero value otherwise.

### GetFormOk

`func (o *CardProductInternal) GetFormOk() (*string, bool)`

GetFormOk returns a tuple with the Form field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForm

`func (o *CardProductInternal) SetForm(v string)`

SetForm sets Form field to given value.


### GetActive

`func (o *CardProductInternal) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *CardProductInternal) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *CardProductInternal) SetActive(v bool)`

SetActive sets Active field to given value.


### GetCardProgramId

`func (o *CardProductInternal) GetCardProgramId() string`

GetCardProgramId returns the CardProgramId field if non-nil, zero value otherwise.

### GetCardProgramIdOk

`func (o *CardProductInternal) GetCardProgramIdOk() (*string, bool)`

GetCardProgramIdOk returns a tuple with the CardProgramId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardProgramId

`func (o *CardProductInternal) SetCardProgramId(v string)`

SetCardProgramId sets CardProgramId field to given value.


### GetColor

`func (o *CardProductInternal) GetColor() string`

GetColor returns the Color field if non-nil, zero value otherwise.

### GetColorOk

`func (o *CardProductInternal) GetColorOk() (*string, bool)`

GetColorOk returns a tuple with the Color field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColor

`func (o *CardProductInternal) SetColor(v string)`

SetColor sets Color field to given value.

### HasColor

`func (o *CardProductInternal) HasColor() bool`

HasColor returns a boolean if a field has been set.

### GetCreationTime

`func (o *CardProductInternal) GetCreationTime() time.Time`

GetCreationTime returns the CreationTime field if non-nil, zero value otherwise.

### GetCreationTimeOk

`func (o *CardProductInternal) GetCreationTimeOk() (*time.Time, bool)`

GetCreationTimeOk returns a tuple with the CreationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationTime

`func (o *CardProductInternal) SetCreationTime(v time.Time)`

SetCreationTime sets CreationTime field to given value.

### HasCreationTime

`func (o *CardProductInternal) HasCreationTime() bool`

HasCreationTime returns a boolean if a field has been set.

### GetDigitalWalletTokenization

`func (o *CardProductInternal) GetDigitalWalletTokenization() DigitalWalletTokenization`

GetDigitalWalletTokenization returns the DigitalWalletTokenization field if non-nil, zero value otherwise.

### GetDigitalWalletTokenizationOk

`func (o *CardProductInternal) GetDigitalWalletTokenizationOk() (*DigitalWalletTokenization, bool)`

GetDigitalWalletTokenizationOk returns a tuple with the DigitalWalletTokenization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDigitalWalletTokenization

`func (o *CardProductInternal) SetDigitalWalletTokenization(v DigitalWalletTokenization)`

SetDigitalWalletTokenization sets DigitalWalletTokenization field to given value.

### HasDigitalWalletTokenization

`func (o *CardProductInternal) HasDigitalWalletTokenization() bool`

HasDigitalWalletTokenization returns a boolean if a field has been set.

### GetEndDate

`func (o *CardProductInternal) GetEndDate() time.Time`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *CardProductInternal) GetEndDateOk() (*time.Time, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *CardProductInternal) SetEndDate(v time.Time)`

SetEndDate sets EndDate field to given value.

### HasEndDate

`func (o *CardProductInternal) HasEndDate() bool`

HasEndDate returns a boolean if a field has been set.

### GetId

`func (o *CardProductInternal) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CardProductInternal) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CardProductInternal) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CardProductInternal) HasId() bool`

HasId returns a boolean if a field has been set.

### GetImage

`func (o *CardProductInternal) GetImage() bool`

GetImage returns the Image field if non-nil, zero value otherwise.

### GetImageOk

`func (o *CardProductInternal) GetImageOk() (*bool, bool)`

GetImageOk returns a tuple with the Image field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImage

`func (o *CardProductInternal) SetImage(v bool)`

SetImage sets Image field to given value.

### HasImage

`func (o *CardProductInternal) HasImage() bool`

HasImage returns a boolean if a field has been set.

### GetImageMode

`func (o *CardProductInternal) GetImageMode() CardImageMode`

GetImageMode returns the ImageMode field if non-nil, zero value otherwise.

### GetImageModeOk

`func (o *CardProductInternal) GetImageModeOk() (*CardImageMode, bool)`

GetImageModeOk returns a tuple with the ImageMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageMode

`func (o *CardProductInternal) SetImageMode(v CardImageMode)`

SetImageMode sets ImageMode field to given value.

### HasImageMode

`func (o *CardProductInternal) HasImageMode() bool`

HasImageMode returns a boolean if a field has been set.

### GetLastModifiedTime

`func (o *CardProductInternal) GetLastModifiedTime() time.Time`

GetLastModifiedTime returns the LastModifiedTime field if non-nil, zero value otherwise.

### GetLastModifiedTimeOk

`func (o *CardProductInternal) GetLastModifiedTimeOk() (*time.Time, bool)`

GetLastModifiedTimeOk returns a tuple with the LastModifiedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModifiedTime

`func (o *CardProductInternal) SetLastModifiedTime(v time.Time)`

SetLastModifiedTime sets LastModifiedTime field to given value.

### HasLastModifiedTime

`func (o *CardProductInternal) HasLastModifiedTime() bool`

HasLastModifiedTime returns a boolean if a field has been set.

### GetName

`func (o *CardProductInternal) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CardProductInternal) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CardProductInternal) SetName(v string)`

SetName sets Name field to given value.


### GetOrientation

`func (o *CardProductInternal) GetOrientation() string`

GetOrientation returns the Orientation field if non-nil, zero value otherwise.

### GetOrientationOk

`func (o *CardProductInternal) GetOrientationOk() (*string, bool)`

GetOrientationOk returns a tuple with the Orientation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrientation

`func (o *CardProductInternal) SetOrientation(v string)`

SetOrientation sets Orientation field to given value.

### HasOrientation

`func (o *CardProductInternal) HasOrientation() bool`

HasOrientation returns a boolean if a field has been set.

### GetPackageId

`func (o *CardProductInternal) GetPackageId() string`

GetPackageId returns the PackageId field if non-nil, zero value otherwise.

### GetPackageIdOk

`func (o *CardProductInternal) GetPackageIdOk() (*string, bool)`

GetPackageIdOk returns a tuple with the PackageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackageId

`func (o *CardProductInternal) SetPackageId(v string)`

SetPackageId sets PackageId field to given value.

### HasPackageId

`func (o *CardProductInternal) HasPackageId() bool`

HasPackageId returns a boolean if a field has been set.

### GetPhysicalCardFormat

`func (o *CardProductInternal) GetPhysicalCardFormat() PhysicalCardFormat`

GetPhysicalCardFormat returns the PhysicalCardFormat field if non-nil, zero value otherwise.

### GetPhysicalCardFormatOk

`func (o *CardProductInternal) GetPhysicalCardFormatOk() (*PhysicalCardFormat, bool)`

GetPhysicalCardFormatOk returns a tuple with the PhysicalCardFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhysicalCardFormat

`func (o *CardProductInternal) SetPhysicalCardFormat(v PhysicalCardFormat)`

SetPhysicalCardFormat sets PhysicalCardFormat field to given value.

### HasPhysicalCardFormat

`func (o *CardProductInternal) HasPhysicalCardFormat() bool`

HasPhysicalCardFormat returns a boolean if a field has been set.

### GetReturnAddress

`func (o *CardProductInternal) GetReturnAddress() Shipping`

GetReturnAddress returns the ReturnAddress field if non-nil, zero value otherwise.

### GetReturnAddressOk

`func (o *CardProductInternal) GetReturnAddressOk() (*Shipping, bool)`

GetReturnAddressOk returns a tuple with the ReturnAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturnAddress

`func (o *CardProductInternal) SetReturnAddress(v Shipping)`

SetReturnAddress sets ReturnAddress field to given value.

### HasReturnAddress

`func (o *CardProductInternal) HasReturnAddress() bool`

HasReturnAddress returns a boolean if a field has been set.

### GetStartDate

`func (o *CardProductInternal) GetStartDate() time.Time`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *CardProductInternal) GetStartDateOk() (*time.Time, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *CardProductInternal) SetStartDate(v time.Time)`

SetStartDate sets StartDate field to given value.


### GetTxnEnhancer

`func (o *CardProductInternal) GetTxnEnhancer() TxnEnhancer`

GetTxnEnhancer returns the TxnEnhancer field if non-nil, zero value otherwise.

### GetTxnEnhancerOk

`func (o *CardProductInternal) GetTxnEnhancerOk() (*TxnEnhancer, bool)`

GetTxnEnhancerOk returns a tuple with the TxnEnhancer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxnEnhancer

`func (o *CardProductInternal) SetTxnEnhancer(v TxnEnhancer)`

SetTxnEnhancer sets TxnEnhancer field to given value.

### HasTxnEnhancer

`func (o *CardProductInternal) HasTxnEnhancer() bool`

HasTxnEnhancer returns a boolean if a field has been set.

### GetAccountRangeId

`func (o *CardProductInternal) GetAccountRangeId() string`

GetAccountRangeId returns the AccountRangeId field if non-nil, zero value otherwise.

### GetAccountRangeIdOk

`func (o *CardProductInternal) GetAccountRangeIdOk() (*string, bool)`

GetAccountRangeIdOk returns a tuple with the AccountRangeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountRangeId

`func (o *CardProductInternal) SetAccountRangeId(v string)`

SetAccountRangeId sets AccountRangeId field to given value.


### GetAutoAllocateRange

`func (o *CardProductInternal) GetAutoAllocateRange() bool`

GetAutoAllocateRange returns the AutoAllocateRange field if non-nil, zero value otherwise.

### GetAutoAllocateRangeOk

`func (o *CardProductInternal) GetAutoAllocateRangeOk() (*bool, bool)`

GetAutoAllocateRangeOk returns a tuple with the AutoAllocateRange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoAllocateRange

`func (o *CardProductInternal) SetAutoAllocateRange(v bool)`

SetAutoAllocateRange sets AutoAllocateRange field to given value.


### GetBankId

`func (o *CardProductInternal) GetBankId() int32`

GetBankId returns the BankId field if non-nil, zero value otherwise.

### GetBankIdOk

`func (o *CardProductInternal) GetBankIdOk() (*int32, bool)`

GetBankIdOk returns a tuple with the BankId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBankId

`func (o *CardProductInternal) SetBankId(v int32)`

SetBankId sets BankId field to given value.


### GetBinId

`func (o *CardProductInternal) GetBinId() string`

GetBinId returns the BinId field if non-nil, zero value otherwise.

### GetBinIdOk

`func (o *CardProductInternal) GetBinIdOk() (*string, bool)`

GetBinIdOk returns a tuple with the BinId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBinId

`func (o *CardProductInternal) SetBinId(v string)`

SetBinId sets BinId field to given value.


### GetFundingSourceId

`func (o *CardProductInternal) GetFundingSourceId() string`

GetFundingSourceId returns the FundingSourceId field if non-nil, zero value otherwise.

### GetFundingSourceIdOk

`func (o *CardProductInternal) GetFundingSourceIdOk() (*string, bool)`

GetFundingSourceIdOk returns a tuple with the FundingSourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFundingSourceId

`func (o *CardProductInternal) SetFundingSourceId(v string)`

SetFundingSourceId sets FundingSourceId field to given value.


### GetPartnerId

`func (o *CardProductInternal) GetPartnerId() int32`

GetPartnerId returns the PartnerId field if non-nil, zero value otherwise.

### GetPartnerIdOk

`func (o *CardProductInternal) GetPartnerIdOk() (*int32, bool)`

GetPartnerIdOk returns a tuple with the PartnerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartnerId

`func (o *CardProductInternal) SetPartnerId(v int32)`

SetPartnerId sets PartnerId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


