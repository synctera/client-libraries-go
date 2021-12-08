# CardProgramCreateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Active** | Pointer to **bool** | indicates whether program is active | [optional] 
**BankId** | **int32** | The ID of the bank partner works with within this program | 
**CardBrand** | [**CardBrand**](CardBrand.md) |  | 
**CardCategory** | [**CardCategory**](CardCategory.md) |  | 
**CardProductType** | [**CardProductType**](CardProductType.md) |  | 
**CreationTime** | Pointer to **time.Time** | The timestamp representing when the program was created | [optional] [readonly] 
**EndDate** | Pointer to **time.Time** | The time when program became inactive | [optional] 
**Id** | Pointer to **string** | Program ID | [optional] [readonly] 
**LastModifiedTime** | Pointer to **time.Time** | The timestamp representing when the program was last modified | [optional] [readonly] 
**Name** | **string** | Program name | 
**PartnerId** | **int32** | The ID of the partner program belongs to | 
**StartDate** | Pointer to **time.Time** | The time when program becomes active | [optional] 

## Methods

### NewCardProgramCreateRequest

`func NewCardProgramCreateRequest(bankId int32, cardBrand CardBrand, cardCategory CardCategory, cardProductType CardProductType, name string, partnerId int32, ) *CardProgramCreateRequest`

NewCardProgramCreateRequest instantiates a new CardProgramCreateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCardProgramCreateRequestWithDefaults

`func NewCardProgramCreateRequestWithDefaults() *CardProgramCreateRequest`

NewCardProgramCreateRequestWithDefaults instantiates a new CardProgramCreateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActive

`func (o *CardProgramCreateRequest) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *CardProgramCreateRequest) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *CardProgramCreateRequest) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *CardProgramCreateRequest) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetBankId

`func (o *CardProgramCreateRequest) GetBankId() int32`

GetBankId returns the BankId field if non-nil, zero value otherwise.

### GetBankIdOk

`func (o *CardProgramCreateRequest) GetBankIdOk() (*int32, bool)`

GetBankIdOk returns a tuple with the BankId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBankId

`func (o *CardProgramCreateRequest) SetBankId(v int32)`

SetBankId sets BankId field to given value.


### GetCardBrand

`func (o *CardProgramCreateRequest) GetCardBrand() CardBrand`

GetCardBrand returns the CardBrand field if non-nil, zero value otherwise.

### GetCardBrandOk

`func (o *CardProgramCreateRequest) GetCardBrandOk() (*CardBrand, bool)`

GetCardBrandOk returns a tuple with the CardBrand field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardBrand

`func (o *CardProgramCreateRequest) SetCardBrand(v CardBrand)`

SetCardBrand sets CardBrand field to given value.


### GetCardCategory

`func (o *CardProgramCreateRequest) GetCardCategory() CardCategory`

GetCardCategory returns the CardCategory field if non-nil, zero value otherwise.

### GetCardCategoryOk

`func (o *CardProgramCreateRequest) GetCardCategoryOk() (*CardCategory, bool)`

GetCardCategoryOk returns a tuple with the CardCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardCategory

`func (o *CardProgramCreateRequest) SetCardCategory(v CardCategory)`

SetCardCategory sets CardCategory field to given value.


### GetCardProductType

`func (o *CardProgramCreateRequest) GetCardProductType() CardProductType`

GetCardProductType returns the CardProductType field if non-nil, zero value otherwise.

### GetCardProductTypeOk

`func (o *CardProgramCreateRequest) GetCardProductTypeOk() (*CardProductType, bool)`

GetCardProductTypeOk returns a tuple with the CardProductType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardProductType

`func (o *CardProgramCreateRequest) SetCardProductType(v CardProductType)`

SetCardProductType sets CardProductType field to given value.


### GetCreationTime

`func (o *CardProgramCreateRequest) GetCreationTime() time.Time`

GetCreationTime returns the CreationTime field if non-nil, zero value otherwise.

### GetCreationTimeOk

`func (o *CardProgramCreateRequest) GetCreationTimeOk() (*time.Time, bool)`

GetCreationTimeOk returns a tuple with the CreationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationTime

`func (o *CardProgramCreateRequest) SetCreationTime(v time.Time)`

SetCreationTime sets CreationTime field to given value.

### HasCreationTime

`func (o *CardProgramCreateRequest) HasCreationTime() bool`

HasCreationTime returns a boolean if a field has been set.

### GetEndDate

`func (o *CardProgramCreateRequest) GetEndDate() time.Time`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *CardProgramCreateRequest) GetEndDateOk() (*time.Time, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *CardProgramCreateRequest) SetEndDate(v time.Time)`

SetEndDate sets EndDate field to given value.

### HasEndDate

`func (o *CardProgramCreateRequest) HasEndDate() bool`

HasEndDate returns a boolean if a field has been set.

### GetId

`func (o *CardProgramCreateRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CardProgramCreateRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CardProgramCreateRequest) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CardProgramCreateRequest) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLastModifiedTime

`func (o *CardProgramCreateRequest) GetLastModifiedTime() time.Time`

GetLastModifiedTime returns the LastModifiedTime field if non-nil, zero value otherwise.

### GetLastModifiedTimeOk

`func (o *CardProgramCreateRequest) GetLastModifiedTimeOk() (*time.Time, bool)`

GetLastModifiedTimeOk returns a tuple with the LastModifiedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModifiedTime

`func (o *CardProgramCreateRequest) SetLastModifiedTime(v time.Time)`

SetLastModifiedTime sets LastModifiedTime field to given value.

### HasLastModifiedTime

`func (o *CardProgramCreateRequest) HasLastModifiedTime() bool`

HasLastModifiedTime returns a boolean if a field has been set.

### GetName

`func (o *CardProgramCreateRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CardProgramCreateRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CardProgramCreateRequest) SetName(v string)`

SetName sets Name field to given value.


### GetPartnerId

`func (o *CardProgramCreateRequest) GetPartnerId() int32`

GetPartnerId returns the PartnerId field if non-nil, zero value otherwise.

### GetPartnerIdOk

`func (o *CardProgramCreateRequest) GetPartnerIdOk() (*int32, bool)`

GetPartnerIdOk returns a tuple with the PartnerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartnerId

`func (o *CardProgramCreateRequest) SetPartnerId(v int32)`

SetPartnerId sets PartnerId field to given value.


### GetStartDate

`func (o *CardProgramCreateRequest) GetStartDate() time.Time`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *CardProgramCreateRequest) GetStartDateOk() (*time.Time, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *CardProgramCreateRequest) SetStartDate(v time.Time)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *CardProgramCreateRequest) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


