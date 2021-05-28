# Balance

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Account ID | [optional] [readonly] 
**PostingDate** | Pointer to **time.Time** | Balance at the posting date | [optional] [readonly] 
**BalanceType** | Pointer to [**BalanceType**](BalanceType.md) |  | [optional] 
**Currency** | Pointer to **string** | Currency of the balance. ISO 4217 alphabetic currency code | [optional] [readonly] 
**Amount** | Pointer to **int32** | amount in ISO 4217 minor currency units | [optional] [readonly] 
**DcSign** | Pointer to [**DcSignType**](DcSignType.md) |  | [optional] 

## Methods

### NewBalance

`func NewBalance() *Balance`

NewBalance instantiates a new Balance object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBalanceWithDefaults

`func NewBalanceWithDefaults() *Balance`

NewBalanceWithDefaults instantiates a new Balance object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Balance) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Balance) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Balance) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Balance) HasId() bool`

HasId returns a boolean if a field has been set.

### GetPostingDate

`func (o *Balance) GetPostingDate() time.Time`

GetPostingDate returns the PostingDate field if non-nil, zero value otherwise.

### GetPostingDateOk

`func (o *Balance) GetPostingDateOk() (*time.Time, bool)`

GetPostingDateOk returns a tuple with the PostingDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostingDate

`func (o *Balance) SetPostingDate(v time.Time)`

SetPostingDate sets PostingDate field to given value.

### HasPostingDate

`func (o *Balance) HasPostingDate() bool`

HasPostingDate returns a boolean if a field has been set.

### GetBalanceType

`func (o *Balance) GetBalanceType() BalanceType`

GetBalanceType returns the BalanceType field if non-nil, zero value otherwise.

### GetBalanceTypeOk

`func (o *Balance) GetBalanceTypeOk() (*BalanceType, bool)`

GetBalanceTypeOk returns a tuple with the BalanceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalanceType

`func (o *Balance) SetBalanceType(v BalanceType)`

SetBalanceType sets BalanceType field to given value.

### HasBalanceType

`func (o *Balance) HasBalanceType() bool`

HasBalanceType returns a boolean if a field has been set.

### GetCurrency

`func (o *Balance) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *Balance) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *Balance) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *Balance) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetAmount

`func (o *Balance) GetAmount() int32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *Balance) GetAmountOk() (*int32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *Balance) SetAmount(v int32)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *Balance) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### GetDcSign

`func (o *Balance) GetDcSign() DcSignType`

GetDcSign returns the DcSign field if non-nil, zero value otherwise.

### GetDcSignOk

`func (o *Balance) GetDcSignOk() (*DcSignType, bool)`

GetDcSignOk returns a tuple with the DcSign field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDcSign

`func (o *Balance) SetDcSign(v DcSignType)`

SetDcSign sets DcSign field to given value.

### HasDcSign

`func (o *Balance) HasDcSign() bool`

HasDcSign returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


