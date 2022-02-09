# BinUpdateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CardBrand** | Pointer to [**CardBrand**](CardBrand.md) |  | [optional] 
**CardCategory** | Pointer to [**CardCategory**](CardCategory.md) |  | [optional] 
**CardProductType** | Pointer to [**CardProductType**](CardProductType.md) |  | [optional] 
**BinStatus** | Pointer to [**BinStatus**](BinStatus.md) |  | [optional] 
**StartDate** | Pointer to **time.Time** | The time when bin goes live | [optional] 
**EndDate** | Pointer to **time.Time** | The time when bin is decommissioned | [optional] 
**PhysicalCardFormat** | Pointer to [**PhysicalCardFormat**](PhysicalCardFormat.md) |  | [optional] 

## Methods

### NewBinUpdateRequest

`func NewBinUpdateRequest() *BinUpdateRequest`

NewBinUpdateRequest instantiates a new BinUpdateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBinUpdateRequestWithDefaults

`func NewBinUpdateRequestWithDefaults() *BinUpdateRequest`

NewBinUpdateRequestWithDefaults instantiates a new BinUpdateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCardBrand

`func (o *BinUpdateRequest) GetCardBrand() CardBrand`

GetCardBrand returns the CardBrand field if non-nil, zero value otherwise.

### GetCardBrandOk

`func (o *BinUpdateRequest) GetCardBrandOk() (*CardBrand, bool)`

GetCardBrandOk returns a tuple with the CardBrand field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardBrand

`func (o *BinUpdateRequest) SetCardBrand(v CardBrand)`

SetCardBrand sets CardBrand field to given value.

### HasCardBrand

`func (o *BinUpdateRequest) HasCardBrand() bool`

HasCardBrand returns a boolean if a field has been set.

### GetCardCategory

`func (o *BinUpdateRequest) GetCardCategory() CardCategory`

GetCardCategory returns the CardCategory field if non-nil, zero value otherwise.

### GetCardCategoryOk

`func (o *BinUpdateRequest) GetCardCategoryOk() (*CardCategory, bool)`

GetCardCategoryOk returns a tuple with the CardCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardCategory

`func (o *BinUpdateRequest) SetCardCategory(v CardCategory)`

SetCardCategory sets CardCategory field to given value.

### HasCardCategory

`func (o *BinUpdateRequest) HasCardCategory() bool`

HasCardCategory returns a boolean if a field has been set.

### GetCardProductType

`func (o *BinUpdateRequest) GetCardProductType() CardProductType`

GetCardProductType returns the CardProductType field if non-nil, zero value otherwise.

### GetCardProductTypeOk

`func (o *BinUpdateRequest) GetCardProductTypeOk() (*CardProductType, bool)`

GetCardProductTypeOk returns a tuple with the CardProductType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardProductType

`func (o *BinUpdateRequest) SetCardProductType(v CardProductType)`

SetCardProductType sets CardProductType field to given value.

### HasCardProductType

`func (o *BinUpdateRequest) HasCardProductType() bool`

HasCardProductType returns a boolean if a field has been set.

### GetBinStatus

`func (o *BinUpdateRequest) GetBinStatus() BinStatus`

GetBinStatus returns the BinStatus field if non-nil, zero value otherwise.

### GetBinStatusOk

`func (o *BinUpdateRequest) GetBinStatusOk() (*BinStatus, bool)`

GetBinStatusOk returns a tuple with the BinStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBinStatus

`func (o *BinUpdateRequest) SetBinStatus(v BinStatus)`

SetBinStatus sets BinStatus field to given value.

### HasBinStatus

`func (o *BinUpdateRequest) HasBinStatus() bool`

HasBinStatus returns a boolean if a field has been set.

### GetStartDate

`func (o *BinUpdateRequest) GetStartDate() time.Time`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *BinUpdateRequest) GetStartDateOk() (*time.Time, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *BinUpdateRequest) SetStartDate(v time.Time)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *BinUpdateRequest) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### GetEndDate

`func (o *BinUpdateRequest) GetEndDate() time.Time`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *BinUpdateRequest) GetEndDateOk() (*time.Time, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *BinUpdateRequest) SetEndDate(v time.Time)`

SetEndDate sets EndDate field to given value.

### HasEndDate

`func (o *BinUpdateRequest) HasEndDate() bool`

HasEndDate returns a boolean if a field has been set.

### GetPhysicalCardFormat

`func (o *BinUpdateRequest) GetPhysicalCardFormat() PhysicalCardFormat`

GetPhysicalCardFormat returns the PhysicalCardFormat field if non-nil, zero value otherwise.

### GetPhysicalCardFormatOk

`func (o *BinUpdateRequest) GetPhysicalCardFormatOk() (*PhysicalCardFormat, bool)`

GetPhysicalCardFormatOk returns a tuple with the PhysicalCardFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhysicalCardFormat

`func (o *BinUpdateRequest) SetPhysicalCardFormat(v PhysicalCardFormat)`

SetPhysicalCardFormat sets PhysicalCardFormat field to given value.

### HasPhysicalCardFormat

`func (o *BinUpdateRequest) HasPhysicalCardFormat() bool`

HasPhysicalCardFormat returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


