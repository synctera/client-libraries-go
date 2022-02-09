# CardProgramResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Program ID | [readonly] 
**Name** | **string** | Program name | 
**BankId** | **int32** | The ID of the bank partner works with within this program | 
**PartnerId** | **int32** | The ID of the partner program belongs to | 
**CardBrand** | [**CardBrand**](CardBrand.md) |  | 
**CardCategory** | [**CardCategory**](CardCategory.md) |  | 
**CardProductType** | [**CardProductType**](CardProductType.md) |  | 
**Active** | **bool** | indicates whether program is active | 
**StartDate** | **time.Time** | The time when program becomes active | 
**EndDate** | **time.Time** | The time when program became inactive | 
**CreationTime** | **time.Time** | The timestamp representing when the program was created | [readonly] 
**LastModifiedTime** | **time.Time** | The timestamp representing when the program was last modified | [readonly] 

## Methods

### NewCardProgramResponse

`func NewCardProgramResponse(id string, name string, bankId int32, partnerId int32, cardBrand CardBrand, cardCategory CardCategory, cardProductType CardProductType, active bool, startDate time.Time, endDate time.Time, creationTime time.Time, lastModifiedTime time.Time, ) *CardProgramResponse`

NewCardProgramResponse instantiates a new CardProgramResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCardProgramResponseWithDefaults

`func NewCardProgramResponseWithDefaults() *CardProgramResponse`

NewCardProgramResponseWithDefaults instantiates a new CardProgramResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CardProgramResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CardProgramResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CardProgramResponse) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *CardProgramResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CardProgramResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CardProgramResponse) SetName(v string)`

SetName sets Name field to given value.


### GetBankId

`func (o *CardProgramResponse) GetBankId() int32`

GetBankId returns the BankId field if non-nil, zero value otherwise.

### GetBankIdOk

`func (o *CardProgramResponse) GetBankIdOk() (*int32, bool)`

GetBankIdOk returns a tuple with the BankId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBankId

`func (o *CardProgramResponse) SetBankId(v int32)`

SetBankId sets BankId field to given value.


### GetPartnerId

`func (o *CardProgramResponse) GetPartnerId() int32`

GetPartnerId returns the PartnerId field if non-nil, zero value otherwise.

### GetPartnerIdOk

`func (o *CardProgramResponse) GetPartnerIdOk() (*int32, bool)`

GetPartnerIdOk returns a tuple with the PartnerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartnerId

`func (o *CardProgramResponse) SetPartnerId(v int32)`

SetPartnerId sets PartnerId field to given value.


### GetCardBrand

`func (o *CardProgramResponse) GetCardBrand() CardBrand`

GetCardBrand returns the CardBrand field if non-nil, zero value otherwise.

### GetCardBrandOk

`func (o *CardProgramResponse) GetCardBrandOk() (*CardBrand, bool)`

GetCardBrandOk returns a tuple with the CardBrand field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardBrand

`func (o *CardProgramResponse) SetCardBrand(v CardBrand)`

SetCardBrand sets CardBrand field to given value.


### GetCardCategory

`func (o *CardProgramResponse) GetCardCategory() CardCategory`

GetCardCategory returns the CardCategory field if non-nil, zero value otherwise.

### GetCardCategoryOk

`func (o *CardProgramResponse) GetCardCategoryOk() (*CardCategory, bool)`

GetCardCategoryOk returns a tuple with the CardCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardCategory

`func (o *CardProgramResponse) SetCardCategory(v CardCategory)`

SetCardCategory sets CardCategory field to given value.


### GetCardProductType

`func (o *CardProgramResponse) GetCardProductType() CardProductType`

GetCardProductType returns the CardProductType field if non-nil, zero value otherwise.

### GetCardProductTypeOk

`func (o *CardProgramResponse) GetCardProductTypeOk() (*CardProductType, bool)`

GetCardProductTypeOk returns a tuple with the CardProductType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardProductType

`func (o *CardProgramResponse) SetCardProductType(v CardProductType)`

SetCardProductType sets CardProductType field to given value.


### GetActive

`func (o *CardProgramResponse) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *CardProgramResponse) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *CardProgramResponse) SetActive(v bool)`

SetActive sets Active field to given value.


### GetStartDate

`func (o *CardProgramResponse) GetStartDate() time.Time`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *CardProgramResponse) GetStartDateOk() (*time.Time, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *CardProgramResponse) SetStartDate(v time.Time)`

SetStartDate sets StartDate field to given value.


### GetEndDate

`func (o *CardProgramResponse) GetEndDate() time.Time`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *CardProgramResponse) GetEndDateOk() (*time.Time, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *CardProgramResponse) SetEndDate(v time.Time)`

SetEndDate sets EndDate field to given value.


### GetCreationTime

`func (o *CardProgramResponse) GetCreationTime() time.Time`

GetCreationTime returns the CreationTime field if non-nil, zero value otherwise.

### GetCreationTimeOk

`func (o *CardProgramResponse) GetCreationTimeOk() (*time.Time, bool)`

GetCreationTimeOk returns a tuple with the CreationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationTime

`func (o *CardProgramResponse) SetCreationTime(v time.Time)`

SetCreationTime sets CreationTime field to given value.


### GetLastModifiedTime

`func (o *CardProgramResponse) GetLastModifiedTime() time.Time`

GetLastModifiedTime returns the LastModifiedTime field if non-nil, zero value otherwise.

### GetLastModifiedTimeOk

`func (o *CardProgramResponse) GetLastModifiedTimeOk() (*time.Time, bool)`

GetLastModifiedTimeOk returns a tuple with the LastModifiedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModifiedTime

`func (o *CardProgramResponse) SetLastModifiedTime(v time.Time)`

SetLastModifiedTime sets LastModifiedTime field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


