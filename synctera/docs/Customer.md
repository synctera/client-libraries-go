# Customer

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Customer unique identifier | [optional] [readonly] 
**FirstName** | Pointer to **string** | Customer&#39;s first name | [optional] 
**LastName** | Pointer to **string** | Customer&#39;s last name | [optional] 
**MiddleName** | Pointer to **string** | Customer&#39;s middle name | [optional] 
**LegalAddress** | Pointer to [**Address**](Address.md) |  | [optional] 
**ShippingAddress** | Pointer to [**Address**](Address.md) |  | [optional] 
**Dob** | Pointer to **string** | Customer&#39;s date of birth in ISO-8601 date format YYYY-MM-DD | [optional] 
**Ssn** | Pointer to **string** | Customer&#39;s full tax ID eg SSN formatted with hyphens 123-45-6789 | [optional] 
**SsnLastFour** | Pointer to **string** | Customer&#39;s masked tax ID eg SSN formatted with hyphens ***-**-6789 | [optional] [readonly] 
**Email** | Pointer to **string** | Customer&#39;s email | [optional] 
**MobilePhoneNumber** | Pointer to **string** | Customer&#39;s mobile phone number in E.164 format e.g. +19178675309 | [optional] 
**AltPhoneNumber** | Pointer to **string** | Customer&#39;s alternate phone number in E.164 format e.g. +19178675309 | [optional] 
**CreationTime** | Pointer to **time.Time** |  | [optional] [readonly] 
**LastUpdatedTime** | Pointer to **time.Time** |  | [optional] [readonly] 
**Devices** | Pointer to [**[]Device**](Device.md) | List of the Customer&#39;s associated devices | [optional] 
**RelatedCustomers** | Pointer to [**[]Relationship1**](Relationship1.md) | Customer&#39;s relationships with other accounts eg. guardian | [optional] 
**Accounts** | Pointer to [**[]Account**](Account.md) | List of accounts that belong to the customer | [optional] [readonly] 

## Methods

### NewCustomer

`func NewCustomer() *Customer`

NewCustomer instantiates a new Customer object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomerWithDefaults

`func NewCustomerWithDefaults() *Customer`

NewCustomerWithDefaults instantiates a new Customer object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Customer) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Customer) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Customer) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Customer) HasId() bool`

HasId returns a boolean if a field has been set.

### GetFirstName

`func (o *Customer) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *Customer) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *Customer) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.

### HasFirstName

`func (o *Customer) HasFirstName() bool`

HasFirstName returns a boolean if a field has been set.

### GetLastName

`func (o *Customer) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *Customer) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *Customer) SetLastName(v string)`

SetLastName sets LastName field to given value.

### HasLastName

`func (o *Customer) HasLastName() bool`

HasLastName returns a boolean if a field has been set.

### GetMiddleName

`func (o *Customer) GetMiddleName() string`

GetMiddleName returns the MiddleName field if non-nil, zero value otherwise.

### GetMiddleNameOk

`func (o *Customer) GetMiddleNameOk() (*string, bool)`

GetMiddleNameOk returns a tuple with the MiddleName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMiddleName

`func (o *Customer) SetMiddleName(v string)`

SetMiddleName sets MiddleName field to given value.

### HasMiddleName

`func (o *Customer) HasMiddleName() bool`

HasMiddleName returns a boolean if a field has been set.

### GetLegalAddress

`func (o *Customer) GetLegalAddress() Address`

GetLegalAddress returns the LegalAddress field if non-nil, zero value otherwise.

### GetLegalAddressOk

`func (o *Customer) GetLegalAddressOk() (*Address, bool)`

GetLegalAddressOk returns a tuple with the LegalAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLegalAddress

`func (o *Customer) SetLegalAddress(v Address)`

SetLegalAddress sets LegalAddress field to given value.

### HasLegalAddress

`func (o *Customer) HasLegalAddress() bool`

HasLegalAddress returns a boolean if a field has been set.

### GetShippingAddress

`func (o *Customer) GetShippingAddress() Address`

GetShippingAddress returns the ShippingAddress field if non-nil, zero value otherwise.

### GetShippingAddressOk

`func (o *Customer) GetShippingAddressOk() (*Address, bool)`

GetShippingAddressOk returns a tuple with the ShippingAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShippingAddress

`func (o *Customer) SetShippingAddress(v Address)`

SetShippingAddress sets ShippingAddress field to given value.

### HasShippingAddress

`func (o *Customer) HasShippingAddress() bool`

HasShippingAddress returns a boolean if a field has been set.

### GetDob

`func (o *Customer) GetDob() string`

GetDob returns the Dob field if non-nil, zero value otherwise.

### GetDobOk

`func (o *Customer) GetDobOk() (*string, bool)`

GetDobOk returns a tuple with the Dob field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDob

`func (o *Customer) SetDob(v string)`

SetDob sets Dob field to given value.

### HasDob

`func (o *Customer) HasDob() bool`

HasDob returns a boolean if a field has been set.

### GetSsn

`func (o *Customer) GetSsn() string`

GetSsn returns the Ssn field if non-nil, zero value otherwise.

### GetSsnOk

`func (o *Customer) GetSsnOk() (*string, bool)`

GetSsnOk returns a tuple with the Ssn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSsn

`func (o *Customer) SetSsn(v string)`

SetSsn sets Ssn field to given value.

### HasSsn

`func (o *Customer) HasSsn() bool`

HasSsn returns a boolean if a field has been set.

### GetSsnLastFour

`func (o *Customer) GetSsnLastFour() string`

GetSsnLastFour returns the SsnLastFour field if non-nil, zero value otherwise.

### GetSsnLastFourOk

`func (o *Customer) GetSsnLastFourOk() (*string, bool)`

GetSsnLastFourOk returns a tuple with the SsnLastFour field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSsnLastFour

`func (o *Customer) SetSsnLastFour(v string)`

SetSsnLastFour sets SsnLastFour field to given value.

### HasSsnLastFour

`func (o *Customer) HasSsnLastFour() bool`

HasSsnLastFour returns a boolean if a field has been set.

### GetEmail

`func (o *Customer) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *Customer) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *Customer) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *Customer) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetMobilePhoneNumber

`func (o *Customer) GetMobilePhoneNumber() string`

GetMobilePhoneNumber returns the MobilePhoneNumber field if non-nil, zero value otherwise.

### GetMobilePhoneNumberOk

`func (o *Customer) GetMobilePhoneNumberOk() (*string, bool)`

GetMobilePhoneNumberOk returns a tuple with the MobilePhoneNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMobilePhoneNumber

`func (o *Customer) SetMobilePhoneNumber(v string)`

SetMobilePhoneNumber sets MobilePhoneNumber field to given value.

### HasMobilePhoneNumber

`func (o *Customer) HasMobilePhoneNumber() bool`

HasMobilePhoneNumber returns a boolean if a field has been set.

### GetAltPhoneNumber

`func (o *Customer) GetAltPhoneNumber() string`

GetAltPhoneNumber returns the AltPhoneNumber field if non-nil, zero value otherwise.

### GetAltPhoneNumberOk

`func (o *Customer) GetAltPhoneNumberOk() (*string, bool)`

GetAltPhoneNumberOk returns a tuple with the AltPhoneNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAltPhoneNumber

`func (o *Customer) SetAltPhoneNumber(v string)`

SetAltPhoneNumber sets AltPhoneNumber field to given value.

### HasAltPhoneNumber

`func (o *Customer) HasAltPhoneNumber() bool`

HasAltPhoneNumber returns a boolean if a field has been set.

### GetCreationTime

`func (o *Customer) GetCreationTime() time.Time`

GetCreationTime returns the CreationTime field if non-nil, zero value otherwise.

### GetCreationTimeOk

`func (o *Customer) GetCreationTimeOk() (*time.Time, bool)`

GetCreationTimeOk returns a tuple with the CreationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationTime

`func (o *Customer) SetCreationTime(v time.Time)`

SetCreationTime sets CreationTime field to given value.

### HasCreationTime

`func (o *Customer) HasCreationTime() bool`

HasCreationTime returns a boolean if a field has been set.

### GetLastUpdatedTime

`func (o *Customer) GetLastUpdatedTime() time.Time`

GetLastUpdatedTime returns the LastUpdatedTime field if non-nil, zero value otherwise.

### GetLastUpdatedTimeOk

`func (o *Customer) GetLastUpdatedTimeOk() (*time.Time, bool)`

GetLastUpdatedTimeOk returns a tuple with the LastUpdatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdatedTime

`func (o *Customer) SetLastUpdatedTime(v time.Time)`

SetLastUpdatedTime sets LastUpdatedTime field to given value.

### HasLastUpdatedTime

`func (o *Customer) HasLastUpdatedTime() bool`

HasLastUpdatedTime returns a boolean if a field has been set.

### GetDevices

`func (o *Customer) GetDevices() []Device`

GetDevices returns the Devices field if non-nil, zero value otherwise.

### GetDevicesOk

`func (o *Customer) GetDevicesOk() (*[]Device, bool)`

GetDevicesOk returns a tuple with the Devices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevices

`func (o *Customer) SetDevices(v []Device)`

SetDevices sets Devices field to given value.

### HasDevices

`func (o *Customer) HasDevices() bool`

HasDevices returns a boolean if a field has been set.

### GetRelatedCustomers

`func (o *Customer) GetRelatedCustomers() []Relationship1`

GetRelatedCustomers returns the RelatedCustomers field if non-nil, zero value otherwise.

### GetRelatedCustomersOk

`func (o *Customer) GetRelatedCustomersOk() (*[]Relationship1, bool)`

GetRelatedCustomersOk returns a tuple with the RelatedCustomers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelatedCustomers

`func (o *Customer) SetRelatedCustomers(v []Relationship1)`

SetRelatedCustomers sets RelatedCustomers field to given value.

### HasRelatedCustomers

`func (o *Customer) HasRelatedCustomers() bool`

HasRelatedCustomers returns a boolean if a field has been set.

### GetAccounts

`func (o *Customer) GetAccounts() []Account`

GetAccounts returns the Accounts field if non-nil, zero value otherwise.

### GetAccountsOk

`func (o *Customer) GetAccountsOk() (*[]Account, bool)`

GetAccountsOk returns a tuple with the Accounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccounts

`func (o *Customer) SetAccounts(v []Account)`

SetAccounts sets Accounts field to given value.

### HasAccounts

`func (o *Customer) HasAccounts() bool`

HasAccounts returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


