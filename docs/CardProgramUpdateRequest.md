# CardProgramUpdateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Program name | [optional] 
**CardBrand** | Pointer to [**CardBrand**](CardBrand.md) |  | [optional] 
**CardCategory** | Pointer to [**CardCategory**](CardCategory.md) |  | [optional] 
**CardProductType** | Pointer to [**CardProductType**](CardProductType.md) |  | [optional] 
**StartDate** | Pointer to **time.Time** | The time when program becomes active | [optional] 
**EndDate** | Pointer to **time.Time** | The time when program became inactive | [optional] 

## Methods

### NewCardProgramUpdateRequest

`func NewCardProgramUpdateRequest() *CardProgramUpdateRequest`

NewCardProgramUpdateRequest instantiates a new CardProgramUpdateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCardProgramUpdateRequestWithDefaults

`func NewCardProgramUpdateRequestWithDefaults() *CardProgramUpdateRequest`

NewCardProgramUpdateRequestWithDefaults instantiates a new CardProgramUpdateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CardProgramUpdateRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CardProgramUpdateRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CardProgramUpdateRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CardProgramUpdateRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetCardBrand

`func (o *CardProgramUpdateRequest) GetCardBrand() CardBrand`

GetCardBrand returns the CardBrand field if non-nil, zero value otherwise.

### GetCardBrandOk

`func (o *CardProgramUpdateRequest) GetCardBrandOk() (*CardBrand, bool)`

GetCardBrandOk returns a tuple with the CardBrand field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardBrand

`func (o *CardProgramUpdateRequest) SetCardBrand(v CardBrand)`

SetCardBrand sets CardBrand field to given value.

### HasCardBrand

`func (o *CardProgramUpdateRequest) HasCardBrand() bool`

HasCardBrand returns a boolean if a field has been set.

### GetCardCategory

`func (o *CardProgramUpdateRequest) GetCardCategory() CardCategory`

GetCardCategory returns the CardCategory field if non-nil, zero value otherwise.

### GetCardCategoryOk

`func (o *CardProgramUpdateRequest) GetCardCategoryOk() (*CardCategory, bool)`

GetCardCategoryOk returns a tuple with the CardCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardCategory

`func (o *CardProgramUpdateRequest) SetCardCategory(v CardCategory)`

SetCardCategory sets CardCategory field to given value.

### HasCardCategory

`func (o *CardProgramUpdateRequest) HasCardCategory() bool`

HasCardCategory returns a boolean if a field has been set.

### GetCardProductType

`func (o *CardProgramUpdateRequest) GetCardProductType() CardProductType`

GetCardProductType returns the CardProductType field if non-nil, zero value otherwise.

### GetCardProductTypeOk

`func (o *CardProgramUpdateRequest) GetCardProductTypeOk() (*CardProductType, bool)`

GetCardProductTypeOk returns a tuple with the CardProductType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardProductType

`func (o *CardProgramUpdateRequest) SetCardProductType(v CardProductType)`

SetCardProductType sets CardProductType field to given value.

### HasCardProductType

`func (o *CardProgramUpdateRequest) HasCardProductType() bool`

HasCardProductType returns a boolean if a field has been set.

### GetStartDate

`func (o *CardProgramUpdateRequest) GetStartDate() time.Time`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *CardProgramUpdateRequest) GetStartDateOk() (*time.Time, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *CardProgramUpdateRequest) SetStartDate(v time.Time)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *CardProgramUpdateRequest) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### GetEndDate

`func (o *CardProgramUpdateRequest) GetEndDate() time.Time`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *CardProgramUpdateRequest) GetEndDateOk() (*time.Time, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *CardProgramUpdateRequest) SetEndDate(v time.Time)`

SetEndDate sets EndDate field to given value.

### HasEndDate

`func (o *CardProgramUpdateRequest) HasEndDate() bool`

HasEndDate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


