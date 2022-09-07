# BankDebitNetworkResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Active** | Pointer to **bool** | indicates whether debit network is active | [optional] 
**CreationTime** | Pointer to **time.Time** | The timestamp representing when the debit network was created | [optional] [readonly] 
**EndDate** | Pointer to **time.Time** | The time when debit network became inactive | [optional] 
**Id** | Pointer to **string** | Debit Network ID | [optional] [readonly] 
**LastModifiedTime** | Pointer to **time.Time** | The timestamp representing when the debit network was last modified | [optional] [readonly] 
**Name** | Pointer to **string** | The name describing the debit network | [optional] 
**StartDate** | Pointer to **time.Time** | The time when debit network goes live | [optional] 
**BankNetworkId** | Pointer to **string** | The ID of the bank network | [optional] 
**BinId** | Pointer to **string** | The ID of the bank&#39;s bin that uses this debit network | [optional] 

## Methods

### NewBankDebitNetworkResponse

`func NewBankDebitNetworkResponse() *BankDebitNetworkResponse`

NewBankDebitNetworkResponse instantiates a new BankDebitNetworkResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBankDebitNetworkResponseWithDefaults

`func NewBankDebitNetworkResponseWithDefaults() *BankDebitNetworkResponse`

NewBankDebitNetworkResponseWithDefaults instantiates a new BankDebitNetworkResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActive

`func (o *BankDebitNetworkResponse) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *BankDebitNetworkResponse) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *BankDebitNetworkResponse) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *BankDebitNetworkResponse) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetCreationTime

`func (o *BankDebitNetworkResponse) GetCreationTime() time.Time`

GetCreationTime returns the CreationTime field if non-nil, zero value otherwise.

### GetCreationTimeOk

`func (o *BankDebitNetworkResponse) GetCreationTimeOk() (*time.Time, bool)`

GetCreationTimeOk returns a tuple with the CreationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationTime

`func (o *BankDebitNetworkResponse) SetCreationTime(v time.Time)`

SetCreationTime sets CreationTime field to given value.

### HasCreationTime

`func (o *BankDebitNetworkResponse) HasCreationTime() bool`

HasCreationTime returns a boolean if a field has been set.

### GetEndDate

`func (o *BankDebitNetworkResponse) GetEndDate() time.Time`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *BankDebitNetworkResponse) GetEndDateOk() (*time.Time, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *BankDebitNetworkResponse) SetEndDate(v time.Time)`

SetEndDate sets EndDate field to given value.

### HasEndDate

`func (o *BankDebitNetworkResponse) HasEndDate() bool`

HasEndDate returns a boolean if a field has been set.

### GetId

`func (o *BankDebitNetworkResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BankDebitNetworkResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BankDebitNetworkResponse) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *BankDebitNetworkResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLastModifiedTime

`func (o *BankDebitNetworkResponse) GetLastModifiedTime() time.Time`

GetLastModifiedTime returns the LastModifiedTime field if non-nil, zero value otherwise.

### GetLastModifiedTimeOk

`func (o *BankDebitNetworkResponse) GetLastModifiedTimeOk() (*time.Time, bool)`

GetLastModifiedTimeOk returns a tuple with the LastModifiedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModifiedTime

`func (o *BankDebitNetworkResponse) SetLastModifiedTime(v time.Time)`

SetLastModifiedTime sets LastModifiedTime field to given value.

### HasLastModifiedTime

`func (o *BankDebitNetworkResponse) HasLastModifiedTime() bool`

HasLastModifiedTime returns a boolean if a field has been set.

### GetName

`func (o *BankDebitNetworkResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BankDebitNetworkResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BankDebitNetworkResponse) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *BankDebitNetworkResponse) HasName() bool`

HasName returns a boolean if a field has been set.

### GetStartDate

`func (o *BankDebitNetworkResponse) GetStartDate() time.Time`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *BankDebitNetworkResponse) GetStartDateOk() (*time.Time, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *BankDebitNetworkResponse) SetStartDate(v time.Time)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *BankDebitNetworkResponse) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### GetBankNetworkId

`func (o *BankDebitNetworkResponse) GetBankNetworkId() string`

GetBankNetworkId returns the BankNetworkId field if non-nil, zero value otherwise.

### GetBankNetworkIdOk

`func (o *BankDebitNetworkResponse) GetBankNetworkIdOk() (*string, bool)`

GetBankNetworkIdOk returns a tuple with the BankNetworkId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBankNetworkId

`func (o *BankDebitNetworkResponse) SetBankNetworkId(v string)`

SetBankNetworkId sets BankNetworkId field to given value.

### HasBankNetworkId

`func (o *BankDebitNetworkResponse) HasBankNetworkId() bool`

HasBankNetworkId returns a boolean if a field has been set.

### GetBinId

`func (o *BankDebitNetworkResponse) GetBinId() string`

GetBinId returns the BinId field if non-nil, zero value otherwise.

### GetBinIdOk

`func (o *BankDebitNetworkResponse) GetBinIdOk() (*string, bool)`

GetBinIdOk returns a tuple with the BinId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBinId

`func (o *BankDebitNetworkResponse) SetBinId(v string)`

SetBinId sets BinId field to given value.

### HasBinId

`func (o *BankDebitNetworkResponse) HasBinId() bool`

HasBinId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


