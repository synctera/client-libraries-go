# AddAccountsRequestRoutingIdentifiers

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BankName** | **string** | The name of the bank managing the account | 
**BankCountries** | **[]string** | The countries that this bank operates the account in | 
**AchRoutingNumber** | **string** | The routing number used for US ACH payments.  | 

## Methods

### NewAddAccountsRequestRoutingIdentifiers

`func NewAddAccountsRequestRoutingIdentifiers(bankName string, bankCountries []string, achRoutingNumber string, ) *AddAccountsRequestRoutingIdentifiers`

NewAddAccountsRequestRoutingIdentifiers instantiates a new AddAccountsRequestRoutingIdentifiers object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddAccountsRequestRoutingIdentifiersWithDefaults

`func NewAddAccountsRequestRoutingIdentifiersWithDefaults() *AddAccountsRequestRoutingIdentifiers`

NewAddAccountsRequestRoutingIdentifiersWithDefaults instantiates a new AddAccountsRequestRoutingIdentifiers object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBankName

`func (o *AddAccountsRequestRoutingIdentifiers) GetBankName() string`

GetBankName returns the BankName field if non-nil, zero value otherwise.

### GetBankNameOk

`func (o *AddAccountsRequestRoutingIdentifiers) GetBankNameOk() (*string, bool)`

GetBankNameOk returns a tuple with the BankName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBankName

`func (o *AddAccountsRequestRoutingIdentifiers) SetBankName(v string)`

SetBankName sets BankName field to given value.


### GetBankCountries

`func (o *AddAccountsRequestRoutingIdentifiers) GetBankCountries() []string`

GetBankCountries returns the BankCountries field if non-nil, zero value otherwise.

### GetBankCountriesOk

`func (o *AddAccountsRequestRoutingIdentifiers) GetBankCountriesOk() (*[]string, bool)`

GetBankCountriesOk returns a tuple with the BankCountries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBankCountries

`func (o *AddAccountsRequestRoutingIdentifiers) SetBankCountries(v []string)`

SetBankCountries sets BankCountries field to given value.


### GetAchRoutingNumber

`func (o *AddAccountsRequestRoutingIdentifiers) GetAchRoutingNumber() string`

GetAchRoutingNumber returns the AchRoutingNumber field if non-nil, zero value otherwise.

### GetAchRoutingNumberOk

`func (o *AddAccountsRequestRoutingIdentifiers) GetAchRoutingNumberOk() (*string, bool)`

GetAchRoutingNumberOk returns a tuple with the AchRoutingNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAchRoutingNumber

`func (o *AddAccountsRequestRoutingIdentifiers) SetAchRoutingNumber(v string)`

SetAchRoutingNumber sets AchRoutingNumber field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


