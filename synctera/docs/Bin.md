# Bin

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountRangeLength** | **int32** | Account range length | 
**BankId** | **int32** | The bank ID | 
**BillingIca** | **string** | The ICA to which fees will be billed | 
**Bin** | **string** | The bin number | 
**BinStatus** | Pointer to [**BinStatus**](BinStatus.md) |  | [optional] 
**BrandProductCode** | **string** | The Mastercard or Visa Product Code - 3 alpha-numeric characters | 
**CardBrand** | [**CardBrand**](CardBrand.md) |  | 
**CardCategory** | [**CardCategory**](CardCategory.md) |  | 
**CardProductType** | [**CardProductType**](CardProductType.md) |  | 
**Country** | **string** | ISO-3166-1 Alpha-2 country code | 
**CreationTime** | Pointer to **time.Time** | The timestamp representing when the bin was created | [optional] [readonly] 
**Currency** | **string** | ISO 4217  Alpha-3 currency code | 
**EndDate** | Pointer to **time.Time** | The time when bin is decommissioned | [optional] 
**IcaBid** | **string** | ICA/BID | 
**Id** | Pointer to **string** | Bin ID | [optional] [readonly] 
**LastModifiedTime** | Pointer to **time.Time** | The timestamp representing when the bin was last modified | [optional] [readonly] 
**PanUtilization** | Pointer to **int32** | Pan utilization | [optional] 
**PartnerId** | **int32** | The partner ID | 
**PhysicalCardFormat** | [**PhysicalCardFormat**](PhysicalCardFormat.md) |  | 
**Processor** | **string** | The name of the card processor | 
**StartDate** | Pointer to **time.Time** | The time when bin goes live | [optional] 

## Methods

### NewBin

`func NewBin(accountRangeLength int32, bankId int32, billingIca string, bin string, brandProductCode string, cardBrand CardBrand, cardCategory CardCategory, cardProductType CardProductType, country string, currency string, icaBid string, partnerId int32, physicalCardFormat PhysicalCardFormat, processor string, ) *Bin`

NewBin instantiates a new Bin object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBinWithDefaults

`func NewBinWithDefaults() *Bin`

NewBinWithDefaults instantiates a new Bin object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountRangeLength

`func (o *Bin) GetAccountRangeLength() int32`

GetAccountRangeLength returns the AccountRangeLength field if non-nil, zero value otherwise.

### GetAccountRangeLengthOk

`func (o *Bin) GetAccountRangeLengthOk() (*int32, bool)`

GetAccountRangeLengthOk returns a tuple with the AccountRangeLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountRangeLength

`func (o *Bin) SetAccountRangeLength(v int32)`

SetAccountRangeLength sets AccountRangeLength field to given value.


### GetBankId

`func (o *Bin) GetBankId() int32`

GetBankId returns the BankId field if non-nil, zero value otherwise.

### GetBankIdOk

`func (o *Bin) GetBankIdOk() (*int32, bool)`

GetBankIdOk returns a tuple with the BankId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBankId

`func (o *Bin) SetBankId(v int32)`

SetBankId sets BankId field to given value.


### GetBillingIca

`func (o *Bin) GetBillingIca() string`

GetBillingIca returns the BillingIca field if non-nil, zero value otherwise.

### GetBillingIcaOk

`func (o *Bin) GetBillingIcaOk() (*string, bool)`

GetBillingIcaOk returns a tuple with the BillingIca field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingIca

`func (o *Bin) SetBillingIca(v string)`

SetBillingIca sets BillingIca field to given value.


### GetBin

`func (o *Bin) GetBin() string`

GetBin returns the Bin field if non-nil, zero value otherwise.

### GetBinOk

`func (o *Bin) GetBinOk() (*string, bool)`

GetBinOk returns a tuple with the Bin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBin

`func (o *Bin) SetBin(v string)`

SetBin sets Bin field to given value.


### GetBinStatus

`func (o *Bin) GetBinStatus() BinStatus`

GetBinStatus returns the BinStatus field if non-nil, zero value otherwise.

### GetBinStatusOk

`func (o *Bin) GetBinStatusOk() (*BinStatus, bool)`

GetBinStatusOk returns a tuple with the BinStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBinStatus

`func (o *Bin) SetBinStatus(v BinStatus)`

SetBinStatus sets BinStatus field to given value.

### HasBinStatus

`func (o *Bin) HasBinStatus() bool`

HasBinStatus returns a boolean if a field has been set.

### GetBrandProductCode

`func (o *Bin) GetBrandProductCode() string`

GetBrandProductCode returns the BrandProductCode field if non-nil, zero value otherwise.

### GetBrandProductCodeOk

`func (o *Bin) GetBrandProductCodeOk() (*string, bool)`

GetBrandProductCodeOk returns a tuple with the BrandProductCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrandProductCode

`func (o *Bin) SetBrandProductCode(v string)`

SetBrandProductCode sets BrandProductCode field to given value.


### GetCardBrand

`func (o *Bin) GetCardBrand() CardBrand`

GetCardBrand returns the CardBrand field if non-nil, zero value otherwise.

### GetCardBrandOk

`func (o *Bin) GetCardBrandOk() (*CardBrand, bool)`

GetCardBrandOk returns a tuple with the CardBrand field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardBrand

`func (o *Bin) SetCardBrand(v CardBrand)`

SetCardBrand sets CardBrand field to given value.


### GetCardCategory

`func (o *Bin) GetCardCategory() CardCategory`

GetCardCategory returns the CardCategory field if non-nil, zero value otherwise.

### GetCardCategoryOk

`func (o *Bin) GetCardCategoryOk() (*CardCategory, bool)`

GetCardCategoryOk returns a tuple with the CardCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardCategory

`func (o *Bin) SetCardCategory(v CardCategory)`

SetCardCategory sets CardCategory field to given value.


### GetCardProductType

`func (o *Bin) GetCardProductType() CardProductType`

GetCardProductType returns the CardProductType field if non-nil, zero value otherwise.

### GetCardProductTypeOk

`func (o *Bin) GetCardProductTypeOk() (*CardProductType, bool)`

GetCardProductTypeOk returns a tuple with the CardProductType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardProductType

`func (o *Bin) SetCardProductType(v CardProductType)`

SetCardProductType sets CardProductType field to given value.


### GetCountry

`func (o *Bin) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *Bin) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *Bin) SetCountry(v string)`

SetCountry sets Country field to given value.


### GetCreationTime

`func (o *Bin) GetCreationTime() time.Time`

GetCreationTime returns the CreationTime field if non-nil, zero value otherwise.

### GetCreationTimeOk

`func (o *Bin) GetCreationTimeOk() (*time.Time, bool)`

GetCreationTimeOk returns a tuple with the CreationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationTime

`func (o *Bin) SetCreationTime(v time.Time)`

SetCreationTime sets CreationTime field to given value.

### HasCreationTime

`func (o *Bin) HasCreationTime() bool`

HasCreationTime returns a boolean if a field has been set.

### GetCurrency

`func (o *Bin) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *Bin) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *Bin) SetCurrency(v string)`

SetCurrency sets Currency field to given value.


### GetEndDate

`func (o *Bin) GetEndDate() time.Time`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *Bin) GetEndDateOk() (*time.Time, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *Bin) SetEndDate(v time.Time)`

SetEndDate sets EndDate field to given value.

### HasEndDate

`func (o *Bin) HasEndDate() bool`

HasEndDate returns a boolean if a field has been set.

### GetIcaBid

`func (o *Bin) GetIcaBid() string`

GetIcaBid returns the IcaBid field if non-nil, zero value otherwise.

### GetIcaBidOk

`func (o *Bin) GetIcaBidOk() (*string, bool)`

GetIcaBidOk returns a tuple with the IcaBid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIcaBid

`func (o *Bin) SetIcaBid(v string)`

SetIcaBid sets IcaBid field to given value.


### GetId

`func (o *Bin) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Bin) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Bin) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Bin) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLastModifiedTime

`func (o *Bin) GetLastModifiedTime() time.Time`

GetLastModifiedTime returns the LastModifiedTime field if non-nil, zero value otherwise.

### GetLastModifiedTimeOk

`func (o *Bin) GetLastModifiedTimeOk() (*time.Time, bool)`

GetLastModifiedTimeOk returns a tuple with the LastModifiedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModifiedTime

`func (o *Bin) SetLastModifiedTime(v time.Time)`

SetLastModifiedTime sets LastModifiedTime field to given value.

### HasLastModifiedTime

`func (o *Bin) HasLastModifiedTime() bool`

HasLastModifiedTime returns a boolean if a field has been set.

### GetPanUtilization

`func (o *Bin) GetPanUtilization() int32`

GetPanUtilization returns the PanUtilization field if non-nil, zero value otherwise.

### GetPanUtilizationOk

`func (o *Bin) GetPanUtilizationOk() (*int32, bool)`

GetPanUtilizationOk returns a tuple with the PanUtilization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPanUtilization

`func (o *Bin) SetPanUtilization(v int32)`

SetPanUtilization sets PanUtilization field to given value.

### HasPanUtilization

`func (o *Bin) HasPanUtilization() bool`

HasPanUtilization returns a boolean if a field has been set.

### GetPartnerId

`func (o *Bin) GetPartnerId() int32`

GetPartnerId returns the PartnerId field if non-nil, zero value otherwise.

### GetPartnerIdOk

`func (o *Bin) GetPartnerIdOk() (*int32, bool)`

GetPartnerIdOk returns a tuple with the PartnerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartnerId

`func (o *Bin) SetPartnerId(v int32)`

SetPartnerId sets PartnerId field to given value.


### GetPhysicalCardFormat

`func (o *Bin) GetPhysicalCardFormat() PhysicalCardFormat`

GetPhysicalCardFormat returns the PhysicalCardFormat field if non-nil, zero value otherwise.

### GetPhysicalCardFormatOk

`func (o *Bin) GetPhysicalCardFormatOk() (*PhysicalCardFormat, bool)`

GetPhysicalCardFormatOk returns a tuple with the PhysicalCardFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhysicalCardFormat

`func (o *Bin) SetPhysicalCardFormat(v PhysicalCardFormat)`

SetPhysicalCardFormat sets PhysicalCardFormat field to given value.


### GetProcessor

`func (o *Bin) GetProcessor() string`

GetProcessor returns the Processor field if non-nil, zero value otherwise.

### GetProcessorOk

`func (o *Bin) GetProcessorOk() (*string, bool)`

GetProcessorOk returns a tuple with the Processor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessor

`func (o *Bin) SetProcessor(v string)`

SetProcessor sets Processor field to given value.


### GetStartDate

`func (o *Bin) GetStartDate() time.Time`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *Bin) GetStartDateOk() (*time.Time, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *Bin) SetStartDate(v time.Time)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *Bin) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


