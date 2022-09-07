# CardProductResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Form** | **string** | PHYSICAL or VIRTUAL. | 
**Active** | **bool** | Indicates whether the Card Product is active | 
**CardProgramId** | **string** | Card Program ID | 
**Color** | Pointer to **string** | Color code for dynamic card elements such as PAN and card holder name | [optional] 
**CreationTime** | **time.Time** | The timestamp representing when the Card Product was created | [readonly] 
**DigitalWalletTokenization** | [**DigitalWalletTokenization**](DigitalWalletTokenization.md) |  | 
**EndDate** | **time.Time** | The time when the Card Product is decommissioned | 
**Id** | **string** | Card Product ID | [readonly] 
**Image** | Pointer to **bool** | Indicates whether or not there is an overlay image of the card product available | [optional] 
**ImageMode** | Pointer to [**CardImageMode**](CardImageMode.md) |  | [optional] 
**LastModifiedTime** | **time.Time** | The timestamp representing when the Card Product was last modified | [readonly] 
**Name** | **string** | The name of the Card Product | 
**Orientation** | Pointer to **string** | Card orientation | [optional] 
**PackageId** | Pointer to **string** | Card fulfillment provider’s package ID | [optional] 
**PhysicalCardFormat** | Pointer to [**PhysicalCardFormat**](PhysicalCardFormat.md) |  | [optional] 
**ReturnAddress** | Pointer to [**Shipping**](Shipping.md) |  | [optional] 
**StartDate** | **time.Time** | The time when the Card Product goes live | 
**TxnEnhancer** | Pointer to [**TxnEnhancer**](TxnEnhancer.md) |  | [optional] [default to TXNENHANCER_NONE]

## Methods

### NewCardProductResponse

`func NewCardProductResponse(form string, active bool, cardProgramId string, creationTime time.Time, digitalWalletTokenization DigitalWalletTokenization, endDate time.Time, id string, lastModifiedTime time.Time, name string, startDate time.Time, ) *CardProductResponse`

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


### GetColor

`func (o *CardProductResponse) GetColor() string`

GetColor returns the Color field if non-nil, zero value otherwise.

### GetColorOk

`func (o *CardProductResponse) GetColorOk() (*string, bool)`

GetColorOk returns a tuple with the Color field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColor

`func (o *CardProductResponse) SetColor(v string)`

SetColor sets Color field to given value.

### HasColor

`func (o *CardProductResponse) HasColor() bool`

HasColor returns a boolean if a field has been set.

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


### GetDigitalWalletTokenization

`func (o *CardProductResponse) GetDigitalWalletTokenization() DigitalWalletTokenization`

GetDigitalWalletTokenization returns the DigitalWalletTokenization field if non-nil, zero value otherwise.

### GetDigitalWalletTokenizationOk

`func (o *CardProductResponse) GetDigitalWalletTokenizationOk() (*DigitalWalletTokenization, bool)`

GetDigitalWalletTokenizationOk returns a tuple with the DigitalWalletTokenization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDigitalWalletTokenization

`func (o *CardProductResponse) SetDigitalWalletTokenization(v DigitalWalletTokenization)`

SetDigitalWalletTokenization sets DigitalWalletTokenization field to given value.


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


### GetImage

`func (o *CardProductResponse) GetImage() bool`

GetImage returns the Image field if non-nil, zero value otherwise.

### GetImageOk

`func (o *CardProductResponse) GetImageOk() (*bool, bool)`

GetImageOk returns a tuple with the Image field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImage

`func (o *CardProductResponse) SetImage(v bool)`

SetImage sets Image field to given value.

### HasImage

`func (o *CardProductResponse) HasImage() bool`

HasImage returns a boolean if a field has been set.

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


### GetOrientation

`func (o *CardProductResponse) GetOrientation() string`

GetOrientation returns the Orientation field if non-nil, zero value otherwise.

### GetOrientationOk

`func (o *CardProductResponse) GetOrientationOk() (*string, bool)`

GetOrientationOk returns a tuple with the Orientation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrientation

`func (o *CardProductResponse) SetOrientation(v string)`

SetOrientation sets Orientation field to given value.

### HasOrientation

`func (o *CardProductResponse) HasOrientation() bool`

HasOrientation returns a boolean if a field has been set.

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


### GetTxnEnhancer

`func (o *CardProductResponse) GetTxnEnhancer() TxnEnhancer`

GetTxnEnhancer returns the TxnEnhancer field if non-nil, zero value otherwise.

### GetTxnEnhancerOk

`func (o *CardProductResponse) GetTxnEnhancerOk() (*TxnEnhancer, bool)`

GetTxnEnhancerOk returns a tuple with the TxnEnhancer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxnEnhancer

`func (o *CardProductResponse) SetTxnEnhancer(v TxnEnhancer)`

SetTxnEnhancer sets TxnEnhancer field to given value.

### HasTxnEnhancer

`func (o *CardProductResponse) HasTxnEnhancer() bool`

HasTxnEnhancer returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


