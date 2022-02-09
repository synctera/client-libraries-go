# CardProgram

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Program ID | [optional] [readonly] 
**Name** | **string** | Program name | 
**BankId** | **int32** | The ID of the bank partner works with within this program | 
**PartnerId** | **int32** | The ID of the partner program belongs to | 
**CardBrand** | [**CardBrand**](CardBrand.md) |  | 
**CardCategory** | [**CardCategory**](CardCategory.md) |  | 
**CardProductType** | [**CardProductType**](CardProductType.md) |  | 
**Active** | Pointer to **bool** | indicates whether program is active | [optional] 
**StartDate** | Pointer to **time.Time** | The time when program becomes active | [optional] 
**EndDate** | Pointer to **time.Time** | The time when program became inactive | [optional] 
**CreationTime** | Pointer to **time.Time** | The timestamp representing when the program was created | [optional] [readonly] 
**LastModifiedTime** | Pointer to **time.Time** | The timestamp representing when the program was last modified | [optional] [readonly] 

## Methods

### NewCardProgram

`func NewCardProgram(name string, bankId int32, partnerId int32, cardBrand CardBrand, cardCategory CardCategory, cardProductType CardProductType, ) *CardProgram`

NewCardProgram instantiates a new CardProgram object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCardProgramWithDefaults

`func NewCardProgramWithDefaults() *CardProgram`

NewCardProgramWithDefaults instantiates a new CardProgram object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CardProgram) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CardProgram) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CardProgram) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CardProgram) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *CardProgram) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CardProgram) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CardProgram) SetName(v string)`

SetName sets Name field to given value.


### GetBankId

`func (o *CardProgram) GetBankId() int32`

GetBankId returns the BankId field if non-nil, zero value otherwise.

### GetBankIdOk

`func (o *CardProgram) GetBankIdOk() (*int32, bool)`

GetBankIdOk returns a tuple with the BankId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBankId

`func (o *CardProgram) SetBankId(v int32)`

SetBankId sets BankId field to given value.


### GetPartnerId

`func (o *CardProgram) GetPartnerId() int32`

GetPartnerId returns the PartnerId field if non-nil, zero value otherwise.

### GetPartnerIdOk

`func (o *CardProgram) GetPartnerIdOk() (*int32, bool)`

GetPartnerIdOk returns a tuple with the PartnerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartnerId

`func (o *CardProgram) SetPartnerId(v int32)`

SetPartnerId sets PartnerId field to given value.


### GetCardBrand

`func (o *CardProgram) GetCardBrand() CardBrand`

GetCardBrand returns the CardBrand field if non-nil, zero value otherwise.

### GetCardBrandOk

`func (o *CardProgram) GetCardBrandOk() (*CardBrand, bool)`

GetCardBrandOk returns a tuple with the CardBrand field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardBrand

`func (o *CardProgram) SetCardBrand(v CardBrand)`

SetCardBrand sets CardBrand field to given value.


### GetCardCategory

`func (o *CardProgram) GetCardCategory() CardCategory`

GetCardCategory returns the CardCategory field if non-nil, zero value otherwise.

### GetCardCategoryOk

`func (o *CardProgram) GetCardCategoryOk() (*CardCategory, bool)`

GetCardCategoryOk returns a tuple with the CardCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardCategory

`func (o *CardProgram) SetCardCategory(v CardCategory)`

SetCardCategory sets CardCategory field to given value.


### GetCardProductType

`func (o *CardProgram) GetCardProductType() CardProductType`

GetCardProductType returns the CardProductType field if non-nil, zero value otherwise.

### GetCardProductTypeOk

`func (o *CardProgram) GetCardProductTypeOk() (*CardProductType, bool)`

GetCardProductTypeOk returns a tuple with the CardProductType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardProductType

`func (o *CardProgram) SetCardProductType(v CardProductType)`

SetCardProductType sets CardProductType field to given value.


### GetActive

`func (o *CardProgram) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *CardProgram) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *CardProgram) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *CardProgram) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetStartDate

`func (o *CardProgram) GetStartDate() time.Time`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *CardProgram) GetStartDateOk() (*time.Time, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *CardProgram) SetStartDate(v time.Time)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *CardProgram) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### GetEndDate

`func (o *CardProgram) GetEndDate() time.Time`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *CardProgram) GetEndDateOk() (*time.Time, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *CardProgram) SetEndDate(v time.Time)`

SetEndDate sets EndDate field to given value.

### HasEndDate

`func (o *CardProgram) HasEndDate() bool`

HasEndDate returns a boolean if a field has been set.

### GetCreationTime

`func (o *CardProgram) GetCreationTime() time.Time`

GetCreationTime returns the CreationTime field if non-nil, zero value otherwise.

### GetCreationTimeOk

`func (o *CardProgram) GetCreationTimeOk() (*time.Time, bool)`

GetCreationTimeOk returns a tuple with the CreationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationTime

`func (o *CardProgram) SetCreationTime(v time.Time)`

SetCreationTime sets CreationTime field to given value.

### HasCreationTime

`func (o *CardProgram) HasCreationTime() bool`

HasCreationTime returns a boolean if a field has been set.

### GetLastModifiedTime

`func (o *CardProgram) GetLastModifiedTime() time.Time`

GetLastModifiedTime returns the LastModifiedTime field if non-nil, zero value otherwise.

### GetLastModifiedTimeOk

`func (o *CardProgram) GetLastModifiedTimeOk() (*time.Time, bool)`

GetLastModifiedTimeOk returns a tuple with the LastModifiedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModifiedTime

`func (o *CardProgram) SetLastModifiedTime(v time.Time)`

SetLastModifiedTime sets LastModifiedTime field to given value.

### HasLastModifiedTime

`func (o *CardProgram) HasLastModifiedTime() bool`

HasLastModifiedTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


