# BinResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountRangeLength** | **int32** | Account range length | 
**BankId** | **int32** | The bank ID | 
**BillingIca** | **string** | The ICA to which fees will be billed | 
**Bin** | **string** | The bin number | 
**BinStatus** | [**BinStatus**](BinStatus.md) |  | 
**BrandProductCode** | **string** | The Mastercard or Visa Product Code - 3 alpha-numeric characters | 
**CardBrand** | [**CardBrand**](CardBrand.md) |  | 
**CardCategory** | [**CardCategory**](CardCategory.md) |  | 
**CardProductType** | [**CardProductType**](CardProductType.md) |  | 
**Country** | **string** | ISO-3166-1 Alpha-2 country code | 
**CreationTime** | **time.Time** | The timestamp representing when the bin was created | [readonly] 
**Currency** | **string** | ISO 4217  Alpha-3 currency code | 
**DigitalWalletActive** | Pointer to **bool** | Determines if bin supports digital wallet tokenization | [optional] [default to false]
**EndDate** | Pointer to **time.Time** | The time when bin is decommissioned | [optional] 
**IcaBid** | **string** | ICA/BID | 
**Id** | **string** | Bin ID | [readonly] 
**LastModifiedTime** | **time.Time** | The timestamp representing when the bin was last modified | [readonly] 
**PanUtilization** | Pointer to **int32** | Pan utilization | [optional] 
**PartnerId** | **int32** | The partner ID | 
**PhysicalCardFormat** | Pointer to [**PhysicalCardFormat**](PhysicalCardFormat.md) |  | [optional] 
**Processor** | **string** | The name of the card processor | 
**StartDate** | Pointer to **time.Time** | The time when bin goes live | [optional] 

## Methods

### NewBinResponse

`func NewBinResponse(accountRangeLength int32, bankId int32, billingIca string, bin string, binStatus BinStatus, brandProductCode string, cardBrand CardBrand, cardCategory CardCategory, cardProductType CardProductType, country string, creationTime time.Time, currency string, icaBid string, id string, lastModifiedTime time.Time, partnerId int32, processor string, ) *BinResponse`

NewBinResponse instantiates a new BinResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBinResponseWithDefaults

`func NewBinResponseWithDefaults() *BinResponse`

NewBinResponseWithDefaults instantiates a new BinResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountRangeLength

`func (o *BinResponse) GetAccountRangeLength() int32`

GetAccountRangeLength returns the AccountRangeLength field if non-nil, zero value otherwise.

### GetAccountRangeLengthOk

`func (o *BinResponse) GetAccountRangeLengthOk() (*int32, bool)`

GetAccountRangeLengthOk returns a tuple with the AccountRangeLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountRangeLength

`func (o *BinResponse) SetAccountRangeLength(v int32)`

SetAccountRangeLength sets AccountRangeLength field to given value.


### GetBankId

`func (o *BinResponse) GetBankId() int32`

GetBankId returns the BankId field if non-nil, zero value otherwise.

### GetBankIdOk

`func (o *BinResponse) GetBankIdOk() (*int32, bool)`

GetBankIdOk returns a tuple with the BankId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBankId

`func (o *BinResponse) SetBankId(v int32)`

SetBankId sets BankId field to given value.


### GetBillingIca

`func (o *BinResponse) GetBillingIca() string`

GetBillingIca returns the BillingIca field if non-nil, zero value otherwise.

### GetBillingIcaOk

`func (o *BinResponse) GetBillingIcaOk() (*string, bool)`

GetBillingIcaOk returns a tuple with the BillingIca field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingIca

`func (o *BinResponse) SetBillingIca(v string)`

SetBillingIca sets BillingIca field to given value.


### GetBin

`func (o *BinResponse) GetBin() string`

GetBin returns the Bin field if non-nil, zero value otherwise.

### GetBinOk

`func (o *BinResponse) GetBinOk() (*string, bool)`

GetBinOk returns a tuple with the Bin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBin

`func (o *BinResponse) SetBin(v string)`

SetBin sets Bin field to given value.


### GetBinStatus

`func (o *BinResponse) GetBinStatus() BinStatus`

GetBinStatus returns the BinStatus field if non-nil, zero value otherwise.

### GetBinStatusOk

`func (o *BinResponse) GetBinStatusOk() (*BinStatus, bool)`

GetBinStatusOk returns a tuple with the BinStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBinStatus

`func (o *BinResponse) SetBinStatus(v BinStatus)`

SetBinStatus sets BinStatus field to given value.


### GetBrandProductCode

`func (o *BinResponse) GetBrandProductCode() string`

GetBrandProductCode returns the BrandProductCode field if non-nil, zero value otherwise.

### GetBrandProductCodeOk

`func (o *BinResponse) GetBrandProductCodeOk() (*string, bool)`

GetBrandProductCodeOk returns a tuple with the BrandProductCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrandProductCode

`func (o *BinResponse) SetBrandProductCode(v string)`

SetBrandProductCode sets BrandProductCode field to given value.


### GetCardBrand

`func (o *BinResponse) GetCardBrand() CardBrand`

GetCardBrand returns the CardBrand field if non-nil, zero value otherwise.

### GetCardBrandOk

`func (o *BinResponse) GetCardBrandOk() (*CardBrand, bool)`

GetCardBrandOk returns a tuple with the CardBrand field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardBrand

`func (o *BinResponse) SetCardBrand(v CardBrand)`

SetCardBrand sets CardBrand field to given value.


### GetCardCategory

`func (o *BinResponse) GetCardCategory() CardCategory`

GetCardCategory returns the CardCategory field if non-nil, zero value otherwise.

### GetCardCategoryOk

`func (o *BinResponse) GetCardCategoryOk() (*CardCategory, bool)`

GetCardCategoryOk returns a tuple with the CardCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardCategory

`func (o *BinResponse) SetCardCategory(v CardCategory)`

SetCardCategory sets CardCategory field to given value.


### GetCardProductType

`func (o *BinResponse) GetCardProductType() CardProductType`

GetCardProductType returns the CardProductType field if non-nil, zero value otherwise.

### GetCardProductTypeOk

`func (o *BinResponse) GetCardProductTypeOk() (*CardProductType, bool)`

GetCardProductTypeOk returns a tuple with the CardProductType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardProductType

`func (o *BinResponse) SetCardProductType(v CardProductType)`

SetCardProductType sets CardProductType field to given value.


### GetCountry

`func (o *BinResponse) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *BinResponse) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *BinResponse) SetCountry(v string)`

SetCountry sets Country field to given value.


### GetCreationTime

`func (o *BinResponse) GetCreationTime() time.Time`

GetCreationTime returns the CreationTime field if non-nil, zero value otherwise.

### GetCreationTimeOk

`func (o *BinResponse) GetCreationTimeOk() (*time.Time, bool)`

GetCreationTimeOk returns a tuple with the CreationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationTime

`func (o *BinResponse) SetCreationTime(v time.Time)`

SetCreationTime sets CreationTime field to given value.


### GetCurrency

`func (o *BinResponse) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *BinResponse) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *BinResponse) SetCurrency(v string)`

SetCurrency sets Currency field to given value.


### GetDigitalWalletActive

`func (o *BinResponse) GetDigitalWalletActive() bool`

GetDigitalWalletActive returns the DigitalWalletActive field if non-nil, zero value otherwise.

### GetDigitalWalletActiveOk

`func (o *BinResponse) GetDigitalWalletActiveOk() (*bool, bool)`

GetDigitalWalletActiveOk returns a tuple with the DigitalWalletActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDigitalWalletActive

`func (o *BinResponse) SetDigitalWalletActive(v bool)`

SetDigitalWalletActive sets DigitalWalletActive field to given value.

### HasDigitalWalletActive

`func (o *BinResponse) HasDigitalWalletActive() bool`

HasDigitalWalletActive returns a boolean if a field has been set.

### GetEndDate

`func (o *BinResponse) GetEndDate() time.Time`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *BinResponse) GetEndDateOk() (*time.Time, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *BinResponse) SetEndDate(v time.Time)`

SetEndDate sets EndDate field to given value.

### HasEndDate

`func (o *BinResponse) HasEndDate() bool`

HasEndDate returns a boolean if a field has been set.

### GetIcaBid

`func (o *BinResponse) GetIcaBid() string`

GetIcaBid returns the IcaBid field if non-nil, zero value otherwise.

### GetIcaBidOk

`func (o *BinResponse) GetIcaBidOk() (*string, bool)`

GetIcaBidOk returns a tuple with the IcaBid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIcaBid

`func (o *BinResponse) SetIcaBid(v string)`

SetIcaBid sets IcaBid field to given value.


### GetId

`func (o *BinResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BinResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BinResponse) SetId(v string)`

SetId sets Id field to given value.


### GetLastModifiedTime

`func (o *BinResponse) GetLastModifiedTime() time.Time`

GetLastModifiedTime returns the LastModifiedTime field if non-nil, zero value otherwise.

### GetLastModifiedTimeOk

`func (o *BinResponse) GetLastModifiedTimeOk() (*time.Time, bool)`

GetLastModifiedTimeOk returns a tuple with the LastModifiedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModifiedTime

`func (o *BinResponse) SetLastModifiedTime(v time.Time)`

SetLastModifiedTime sets LastModifiedTime field to given value.


### GetPanUtilization

`func (o *BinResponse) GetPanUtilization() int32`

GetPanUtilization returns the PanUtilization field if non-nil, zero value otherwise.

### GetPanUtilizationOk

`func (o *BinResponse) GetPanUtilizationOk() (*int32, bool)`

GetPanUtilizationOk returns a tuple with the PanUtilization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPanUtilization

`func (o *BinResponse) SetPanUtilization(v int32)`

SetPanUtilization sets PanUtilization field to given value.

### HasPanUtilization

`func (o *BinResponse) HasPanUtilization() bool`

HasPanUtilization returns a boolean if a field has been set.

### GetPartnerId

`func (o *BinResponse) GetPartnerId() int32`

GetPartnerId returns the PartnerId field if non-nil, zero value otherwise.

### GetPartnerIdOk

`func (o *BinResponse) GetPartnerIdOk() (*int32, bool)`

GetPartnerIdOk returns a tuple with the PartnerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartnerId

`func (o *BinResponse) SetPartnerId(v int32)`

SetPartnerId sets PartnerId field to given value.


### GetPhysicalCardFormat

`func (o *BinResponse) GetPhysicalCardFormat() PhysicalCardFormat`

GetPhysicalCardFormat returns the PhysicalCardFormat field if non-nil, zero value otherwise.

### GetPhysicalCardFormatOk

`func (o *BinResponse) GetPhysicalCardFormatOk() (*PhysicalCardFormat, bool)`

GetPhysicalCardFormatOk returns a tuple with the PhysicalCardFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhysicalCardFormat

`func (o *BinResponse) SetPhysicalCardFormat(v PhysicalCardFormat)`

SetPhysicalCardFormat sets PhysicalCardFormat field to given value.

### HasPhysicalCardFormat

`func (o *BinResponse) HasPhysicalCardFormat() bool`

HasPhysicalCardFormat returns a boolean if a field has been set.

### GetProcessor

`func (o *BinResponse) GetProcessor() string`

GetProcessor returns the Processor field if non-nil, zero value otherwise.

### GetProcessorOk

`func (o *BinResponse) GetProcessorOk() (*string, bool)`

GetProcessorOk returns a tuple with the Processor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessor

`func (o *BinResponse) SetProcessor(v string)`

SetProcessor sets Processor field to given value.


### GetStartDate

`func (o *BinResponse) GetStartDate() time.Time`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *BinResponse) GetStartDateOk() (*time.Time, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *BinResponse) SetStartDate(v time.Time)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *BinResponse) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


