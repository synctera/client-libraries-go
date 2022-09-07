# ApplicationRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BankId** | **int32** | Bank ID | 
**BusinessAddress** | [**Address1**](Address1.md) |  | 
**BusinessName** | **string** | todo | 
**BusinessPhone** | **string** | todo | 
**BusinessTaxId** | **string** | todo | 
**BusinessType** | **string** | todo | 
**Dob** | Pointer to [**ExternalPaymentDate**](ExternalPaymentDate.md) |  | [optional] 
**DoingBusinessAs** | **string** | todo | 
**Email** | Pointer to **string** | todo | [optional] 
**FirstName** | Pointer to **string** | todo | [optional] 
**IncorporationDate** | Pointer to [**ExternalPaymentDate**](ExternalPaymentDate.md) |  | [optional] 
**LastName** | Pointer to **string** | todo | [optional] 
**MaxTransactionAmount** | **int32** | Maximum amount that can be transacted for a single transaction in cents | 
**PartnerId** | **int32** | Partner ID | 
**PersonalAddress** | Pointer to [**Address1**](Address1.md) |  | [optional] 
**Phone** | Pointer to **string** | todo | [optional] 
**PrincipalPercentageOwnership** | Pointer to **string** | todo | [optional] 
**Processor** | [**Processor**](Processor.md) |  | 
**TaxId** | Pointer to **string** | todo | [optional] 
**Title** | Pointer to **string** | todo | [optional] 
**Url** | Pointer to **string** | todo | [optional] 

## Methods

### NewApplicationRequest

`func NewApplicationRequest(bankId int32, businessAddress Address1, businessName string, businessPhone string, businessTaxId string, businessType string, doingBusinessAs string, maxTransactionAmount int32, partnerId int32, processor Processor, ) *ApplicationRequest`

NewApplicationRequest instantiates a new ApplicationRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplicationRequestWithDefaults

`func NewApplicationRequestWithDefaults() *ApplicationRequest`

NewApplicationRequestWithDefaults instantiates a new ApplicationRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBankId

`func (o *ApplicationRequest) GetBankId() int32`

GetBankId returns the BankId field if non-nil, zero value otherwise.

### GetBankIdOk

`func (o *ApplicationRequest) GetBankIdOk() (*int32, bool)`

GetBankIdOk returns a tuple with the BankId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBankId

`func (o *ApplicationRequest) SetBankId(v int32)`

SetBankId sets BankId field to given value.


### GetBusinessAddress

`func (o *ApplicationRequest) GetBusinessAddress() Address1`

GetBusinessAddress returns the BusinessAddress field if non-nil, zero value otherwise.

### GetBusinessAddressOk

`func (o *ApplicationRequest) GetBusinessAddressOk() (*Address1, bool)`

GetBusinessAddressOk returns a tuple with the BusinessAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessAddress

`func (o *ApplicationRequest) SetBusinessAddress(v Address1)`

SetBusinessAddress sets BusinessAddress field to given value.


### GetBusinessName

`func (o *ApplicationRequest) GetBusinessName() string`

GetBusinessName returns the BusinessName field if non-nil, zero value otherwise.

### GetBusinessNameOk

`func (o *ApplicationRequest) GetBusinessNameOk() (*string, bool)`

GetBusinessNameOk returns a tuple with the BusinessName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessName

`func (o *ApplicationRequest) SetBusinessName(v string)`

SetBusinessName sets BusinessName field to given value.


### GetBusinessPhone

`func (o *ApplicationRequest) GetBusinessPhone() string`

GetBusinessPhone returns the BusinessPhone field if non-nil, zero value otherwise.

### GetBusinessPhoneOk

`func (o *ApplicationRequest) GetBusinessPhoneOk() (*string, bool)`

GetBusinessPhoneOk returns a tuple with the BusinessPhone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessPhone

`func (o *ApplicationRequest) SetBusinessPhone(v string)`

SetBusinessPhone sets BusinessPhone field to given value.


### GetBusinessTaxId

`func (o *ApplicationRequest) GetBusinessTaxId() string`

GetBusinessTaxId returns the BusinessTaxId field if non-nil, zero value otherwise.

### GetBusinessTaxIdOk

`func (o *ApplicationRequest) GetBusinessTaxIdOk() (*string, bool)`

GetBusinessTaxIdOk returns a tuple with the BusinessTaxId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessTaxId

`func (o *ApplicationRequest) SetBusinessTaxId(v string)`

SetBusinessTaxId sets BusinessTaxId field to given value.


### GetBusinessType

`func (o *ApplicationRequest) GetBusinessType() string`

GetBusinessType returns the BusinessType field if non-nil, zero value otherwise.

### GetBusinessTypeOk

`func (o *ApplicationRequest) GetBusinessTypeOk() (*string, bool)`

GetBusinessTypeOk returns a tuple with the BusinessType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessType

`func (o *ApplicationRequest) SetBusinessType(v string)`

SetBusinessType sets BusinessType field to given value.


### GetDob

`func (o *ApplicationRequest) GetDob() ExternalPaymentDate`

GetDob returns the Dob field if non-nil, zero value otherwise.

### GetDobOk

`func (o *ApplicationRequest) GetDobOk() (*ExternalPaymentDate, bool)`

GetDobOk returns a tuple with the Dob field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDob

`func (o *ApplicationRequest) SetDob(v ExternalPaymentDate)`

SetDob sets Dob field to given value.

### HasDob

`func (o *ApplicationRequest) HasDob() bool`

HasDob returns a boolean if a field has been set.

### GetDoingBusinessAs

`func (o *ApplicationRequest) GetDoingBusinessAs() string`

GetDoingBusinessAs returns the DoingBusinessAs field if non-nil, zero value otherwise.

### GetDoingBusinessAsOk

`func (o *ApplicationRequest) GetDoingBusinessAsOk() (*string, bool)`

GetDoingBusinessAsOk returns a tuple with the DoingBusinessAs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDoingBusinessAs

`func (o *ApplicationRequest) SetDoingBusinessAs(v string)`

SetDoingBusinessAs sets DoingBusinessAs field to given value.


### GetEmail

`func (o *ApplicationRequest) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *ApplicationRequest) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *ApplicationRequest) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *ApplicationRequest) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetFirstName

`func (o *ApplicationRequest) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *ApplicationRequest) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *ApplicationRequest) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.

### HasFirstName

`func (o *ApplicationRequest) HasFirstName() bool`

HasFirstName returns a boolean if a field has been set.

### GetIncorporationDate

`func (o *ApplicationRequest) GetIncorporationDate() ExternalPaymentDate`

GetIncorporationDate returns the IncorporationDate field if non-nil, zero value otherwise.

### GetIncorporationDateOk

`func (o *ApplicationRequest) GetIncorporationDateOk() (*ExternalPaymentDate, bool)`

GetIncorporationDateOk returns a tuple with the IncorporationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncorporationDate

`func (o *ApplicationRequest) SetIncorporationDate(v ExternalPaymentDate)`

SetIncorporationDate sets IncorporationDate field to given value.

### HasIncorporationDate

`func (o *ApplicationRequest) HasIncorporationDate() bool`

HasIncorporationDate returns a boolean if a field has been set.

### GetLastName

`func (o *ApplicationRequest) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *ApplicationRequest) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *ApplicationRequest) SetLastName(v string)`

SetLastName sets LastName field to given value.

### HasLastName

`func (o *ApplicationRequest) HasLastName() bool`

HasLastName returns a boolean if a field has been set.

### GetMaxTransactionAmount

`func (o *ApplicationRequest) GetMaxTransactionAmount() int32`

GetMaxTransactionAmount returns the MaxTransactionAmount field if non-nil, zero value otherwise.

### GetMaxTransactionAmountOk

`func (o *ApplicationRequest) GetMaxTransactionAmountOk() (*int32, bool)`

GetMaxTransactionAmountOk returns a tuple with the MaxTransactionAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxTransactionAmount

`func (o *ApplicationRequest) SetMaxTransactionAmount(v int32)`

SetMaxTransactionAmount sets MaxTransactionAmount field to given value.


### GetPartnerId

`func (o *ApplicationRequest) GetPartnerId() int32`

GetPartnerId returns the PartnerId field if non-nil, zero value otherwise.

### GetPartnerIdOk

`func (o *ApplicationRequest) GetPartnerIdOk() (*int32, bool)`

GetPartnerIdOk returns a tuple with the PartnerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartnerId

`func (o *ApplicationRequest) SetPartnerId(v int32)`

SetPartnerId sets PartnerId field to given value.


### GetPersonalAddress

`func (o *ApplicationRequest) GetPersonalAddress() Address1`

GetPersonalAddress returns the PersonalAddress field if non-nil, zero value otherwise.

### GetPersonalAddressOk

`func (o *ApplicationRequest) GetPersonalAddressOk() (*Address1, bool)`

GetPersonalAddressOk returns a tuple with the PersonalAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPersonalAddress

`func (o *ApplicationRequest) SetPersonalAddress(v Address1)`

SetPersonalAddress sets PersonalAddress field to given value.

### HasPersonalAddress

`func (o *ApplicationRequest) HasPersonalAddress() bool`

HasPersonalAddress returns a boolean if a field has been set.

### GetPhone

`func (o *ApplicationRequest) GetPhone() string`

GetPhone returns the Phone field if non-nil, zero value otherwise.

### GetPhoneOk

`func (o *ApplicationRequest) GetPhoneOk() (*string, bool)`

GetPhoneOk returns a tuple with the Phone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhone

`func (o *ApplicationRequest) SetPhone(v string)`

SetPhone sets Phone field to given value.

### HasPhone

`func (o *ApplicationRequest) HasPhone() bool`

HasPhone returns a boolean if a field has been set.

### GetPrincipalPercentageOwnership

`func (o *ApplicationRequest) GetPrincipalPercentageOwnership() string`

GetPrincipalPercentageOwnership returns the PrincipalPercentageOwnership field if non-nil, zero value otherwise.

### GetPrincipalPercentageOwnershipOk

`func (o *ApplicationRequest) GetPrincipalPercentageOwnershipOk() (*string, bool)`

GetPrincipalPercentageOwnershipOk returns a tuple with the PrincipalPercentageOwnership field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrincipalPercentageOwnership

`func (o *ApplicationRequest) SetPrincipalPercentageOwnership(v string)`

SetPrincipalPercentageOwnership sets PrincipalPercentageOwnership field to given value.

### HasPrincipalPercentageOwnership

`func (o *ApplicationRequest) HasPrincipalPercentageOwnership() bool`

HasPrincipalPercentageOwnership returns a boolean if a field has been set.

### GetProcessor

`func (o *ApplicationRequest) GetProcessor() Processor`

GetProcessor returns the Processor field if non-nil, zero value otherwise.

### GetProcessorOk

`func (o *ApplicationRequest) GetProcessorOk() (*Processor, bool)`

GetProcessorOk returns a tuple with the Processor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessor

`func (o *ApplicationRequest) SetProcessor(v Processor)`

SetProcessor sets Processor field to given value.


### GetTaxId

`func (o *ApplicationRequest) GetTaxId() string`

GetTaxId returns the TaxId field if non-nil, zero value otherwise.

### GetTaxIdOk

`func (o *ApplicationRequest) GetTaxIdOk() (*string, bool)`

GetTaxIdOk returns a tuple with the TaxId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxId

`func (o *ApplicationRequest) SetTaxId(v string)`

SetTaxId sets TaxId field to given value.

### HasTaxId

`func (o *ApplicationRequest) HasTaxId() bool`

HasTaxId returns a boolean if a field has been set.

### GetTitle

`func (o *ApplicationRequest) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *ApplicationRequest) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *ApplicationRequest) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *ApplicationRequest) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetUrl

`func (o *ApplicationRequest) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *ApplicationRequest) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *ApplicationRequest) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *ApplicationRequest) HasUrl() bool`

HasUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


