# AccountRouting

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AchRoutingNumber** | Pointer to **string** | The routing number used for US ACH payments. Only appears if &#x60;bank_countries&#x60; contains &#x60;US&#x60;.  | [optional] 
**BankCountries** | **[]string** | The countries that this bank operates the account in | 
**BankName** | **string** | The name of the bank managing the account | 
**EftBranchNumber** | Pointer to **string** | The branch number used for EFT payments, identifying a branch at a Canadian bank. Only appears if &#x60;bank_countries&#x60; contains &#x60;CA&#x60;.  | [optional] 
**EftInstitutionNumber** | Pointer to **string** | The institution number used for EFT payments, identifying a Canadian bank. Only appears if &#x60;bank_countries&#x60; contains &#x60;CA&#x60;.  | [optional] 
**SwiftCode** | **string** | The SWIFT code for the bank | 
**WireRoutingNumber** | Pointer to **string** | The routing number used for domestic wire payments. Only appears if &#x60;bank_countries&#x60; contains &#x60;US&#x60;.  | [optional] 

## Methods

### NewAccountRouting

`func NewAccountRouting(bankCountries []string, bankName string, swiftCode string, ) *AccountRouting`

NewAccountRouting instantiates a new AccountRouting object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountRoutingWithDefaults

`func NewAccountRoutingWithDefaults() *AccountRouting`

NewAccountRoutingWithDefaults instantiates a new AccountRouting object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAchRoutingNumber

`func (o *AccountRouting) GetAchRoutingNumber() string`

GetAchRoutingNumber returns the AchRoutingNumber field if non-nil, zero value otherwise.

### GetAchRoutingNumberOk

`func (o *AccountRouting) GetAchRoutingNumberOk() (*string, bool)`

GetAchRoutingNumberOk returns a tuple with the AchRoutingNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAchRoutingNumber

`func (o *AccountRouting) SetAchRoutingNumber(v string)`

SetAchRoutingNumber sets AchRoutingNumber field to given value.

### HasAchRoutingNumber

`func (o *AccountRouting) HasAchRoutingNumber() bool`

HasAchRoutingNumber returns a boolean if a field has been set.

### GetBankCountries

`func (o *AccountRouting) GetBankCountries() []string`

GetBankCountries returns the BankCountries field if non-nil, zero value otherwise.

### GetBankCountriesOk

`func (o *AccountRouting) GetBankCountriesOk() (*[]string, bool)`

GetBankCountriesOk returns a tuple with the BankCountries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBankCountries

`func (o *AccountRouting) SetBankCountries(v []string)`

SetBankCountries sets BankCountries field to given value.


### GetBankName

`func (o *AccountRouting) GetBankName() string`

GetBankName returns the BankName field if non-nil, zero value otherwise.

### GetBankNameOk

`func (o *AccountRouting) GetBankNameOk() (*string, bool)`

GetBankNameOk returns a tuple with the BankName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBankName

`func (o *AccountRouting) SetBankName(v string)`

SetBankName sets BankName field to given value.


### GetEftBranchNumber

`func (o *AccountRouting) GetEftBranchNumber() string`

GetEftBranchNumber returns the EftBranchNumber field if non-nil, zero value otherwise.

### GetEftBranchNumberOk

`func (o *AccountRouting) GetEftBranchNumberOk() (*string, bool)`

GetEftBranchNumberOk returns a tuple with the EftBranchNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEftBranchNumber

`func (o *AccountRouting) SetEftBranchNumber(v string)`

SetEftBranchNumber sets EftBranchNumber field to given value.

### HasEftBranchNumber

`func (o *AccountRouting) HasEftBranchNumber() bool`

HasEftBranchNumber returns a boolean if a field has been set.

### GetEftInstitutionNumber

`func (o *AccountRouting) GetEftInstitutionNumber() string`

GetEftInstitutionNumber returns the EftInstitutionNumber field if non-nil, zero value otherwise.

### GetEftInstitutionNumberOk

`func (o *AccountRouting) GetEftInstitutionNumberOk() (*string, bool)`

GetEftInstitutionNumberOk returns a tuple with the EftInstitutionNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEftInstitutionNumber

`func (o *AccountRouting) SetEftInstitutionNumber(v string)`

SetEftInstitutionNumber sets EftInstitutionNumber field to given value.

### HasEftInstitutionNumber

`func (o *AccountRouting) HasEftInstitutionNumber() bool`

HasEftInstitutionNumber returns a boolean if a field has been set.

### GetSwiftCode

`func (o *AccountRouting) GetSwiftCode() string`

GetSwiftCode returns the SwiftCode field if non-nil, zero value otherwise.

### GetSwiftCodeOk

`func (o *AccountRouting) GetSwiftCodeOk() (*string, bool)`

GetSwiftCodeOk returns a tuple with the SwiftCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSwiftCode

`func (o *AccountRouting) SetSwiftCode(v string)`

SetSwiftCode sets SwiftCode field to given value.


### GetWireRoutingNumber

`func (o *AccountRouting) GetWireRoutingNumber() string`

GetWireRoutingNumber returns the WireRoutingNumber field if non-nil, zero value otherwise.

### GetWireRoutingNumberOk

`func (o *AccountRouting) GetWireRoutingNumberOk() (*string, bool)`

GetWireRoutingNumberOk returns a tuple with the WireRoutingNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWireRoutingNumber

`func (o *AccountRouting) SetWireRoutingNumber(v string)`

SetWireRoutingNumber sets WireRoutingNumber field to given value.

### HasWireRoutingNumber

`func (o *AccountRouting) HasWireRoutingNumber() bool`

HasWireRoutingNumber returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


