# CardProductAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
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

## Methods

### NewCardProductAllOf

`func NewCardProductAllOf(active bool, cardProgramId string, name string, startDate time.Time, ) *CardProductAllOf`

NewCardProductAllOf instantiates a new CardProductAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCardProductAllOfWithDefaults

`func NewCardProductAllOfWithDefaults() *CardProductAllOf`

NewCardProductAllOfWithDefaults instantiates a new CardProductAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

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


### GetColor

`func (o *CardProductAllOf) GetColor() string`

GetColor returns the Color field if non-nil, zero value otherwise.

### GetColorOk

`func (o *CardProductAllOf) GetColorOk() (*string, bool)`

GetColorOk returns a tuple with the Color field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColor

`func (o *CardProductAllOf) SetColor(v string)`

SetColor sets Color field to given value.

### HasColor

`func (o *CardProductAllOf) HasColor() bool`

HasColor returns a boolean if a field has been set.

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

### GetDigitalWalletTokenization

`func (o *CardProductAllOf) GetDigitalWalletTokenization() DigitalWalletTokenization`

GetDigitalWalletTokenization returns the DigitalWalletTokenization field if non-nil, zero value otherwise.

### GetDigitalWalletTokenizationOk

`func (o *CardProductAllOf) GetDigitalWalletTokenizationOk() (*DigitalWalletTokenization, bool)`

GetDigitalWalletTokenizationOk returns a tuple with the DigitalWalletTokenization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDigitalWalletTokenization

`func (o *CardProductAllOf) SetDigitalWalletTokenization(v DigitalWalletTokenization)`

SetDigitalWalletTokenization sets DigitalWalletTokenization field to given value.

### HasDigitalWalletTokenization

`func (o *CardProductAllOf) HasDigitalWalletTokenization() bool`

HasDigitalWalletTokenization returns a boolean if a field has been set.

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

### GetImage

`func (o *CardProductAllOf) GetImage() bool`

GetImage returns the Image field if non-nil, zero value otherwise.

### GetImageOk

`func (o *CardProductAllOf) GetImageOk() (*bool, bool)`

GetImageOk returns a tuple with the Image field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImage

`func (o *CardProductAllOf) SetImage(v bool)`

SetImage sets Image field to given value.

### HasImage

`func (o *CardProductAllOf) HasImage() bool`

HasImage returns a boolean if a field has been set.

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


### GetOrientation

`func (o *CardProductAllOf) GetOrientation() string`

GetOrientation returns the Orientation field if non-nil, zero value otherwise.

### GetOrientationOk

`func (o *CardProductAllOf) GetOrientationOk() (*string, bool)`

GetOrientationOk returns a tuple with the Orientation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrientation

`func (o *CardProductAllOf) SetOrientation(v string)`

SetOrientation sets Orientation field to given value.

### HasOrientation

`func (o *CardProductAllOf) HasOrientation() bool`

HasOrientation returns a boolean if a field has been set.

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


### GetTxnEnhancer

`func (o *CardProductAllOf) GetTxnEnhancer() TxnEnhancer`

GetTxnEnhancer returns the TxnEnhancer field if non-nil, zero value otherwise.

### GetTxnEnhancerOk

`func (o *CardProductAllOf) GetTxnEnhancerOk() (*TxnEnhancer, bool)`

GetTxnEnhancerOk returns a tuple with the TxnEnhancer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxnEnhancer

`func (o *CardProductAllOf) SetTxnEnhancer(v TxnEnhancer)`

SetTxnEnhancer sets TxnEnhancer field to given value.

### HasTxnEnhancer

`func (o *CardProductAllOf) HasTxnEnhancer() bool`

HasTxnEnhancer returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


